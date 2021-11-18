package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"github.com/SaWLeaDeR/key-value-store/internal/envvar"
	"github.com/SaWLeaDeR/key-value-store/internal/envvar/vault"
	"github.com/SaWLeaDeR/key-value-store/internal/rest"
	"github.com/SaWLeaDeR/key-value-store/internal/service"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

//go:embed static
var content embed.FS

func main() {
	var env, address string

	flag.StringVar(&env, "env", "", "Environment Variables filename")
	flag.StringVar(&address, "address", ":9234", "HTTP Server Address")
	flag.Parse()

	errC, err := run(env, address)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(env, address string) (<-chan error, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("zap.NewProduction %w", err)
	}

	if err := envvar.Load(env); err != nil {
		return nil, fmt.Errorf("envvar.Load %w", err)
	}

	vault, err := newVaultProvider()
	if err != nil {
		return nil, fmt.Errorf("newVaultProvider %w", err)
	}

	conf := envvar.New(vault)

	promExporter, err := newOTExporter(conf)
	if err != nil {
		return nil, fmt.Errorf("newOTExporter %w", err)
	}

	logging := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info(r.Method,
				zap.Time("time", time.Now()),
				zap.String("url", r.URL.String()),
			)

			h.ServeHTTP(w, r)
		})
	}

	errC := make(chan error, 1)

	srv := newServer(logger, address, promExporter, otelmux.Middleware("store-data-api-server"), logging)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		logger.Info("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			logger.Sync()
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)

		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		logger.Info("Shutdown completed")
	}()

	go func() {
		logger.Info("Listening and serving", zap.String("address", address))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, nil
}

func newServer(logger *zap.Logger, address string, metrics http.Handler, mws ...mux.MiddlewareFunc) *http.Server {
	r := mux.NewRouter()

	for _, mw := range mws {
		r.Use(mw)
	}

	fsys, _ := fs.Sub(content, "static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(fsys))))

	// There is a problem here i am trying to solve
	// https://github.com/open-telemetry/opentelemetry-go/blob/main/CHANGELOG.md
	r.Handle("/metrics", metrics)
	handler := createServices()

	rest.RegisterOpenAPI(r)
	handler.Register(r)

	return &http.Server{
		Handler:           r,
		Addr:              address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
}

func createServices() *rest.StoreHandler {
	storage := make(map[string]string)

	storeService := service.NewStoreService(storage)

	return rest.NewStoreHandler(storeService)
}

func newVaultProvider() (*vault.Provider, error) {
	vaultPath := os.Getenv("VAULT_PATH")
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddress := os.Getenv("VAULT_ADDRESS")

	provider, err := vault.New(vaultToken, vaultAddress, vaultPath)
	if err != nil {
		return nil, fmt.Errorf("vault.New %w", err)
	}

	return provider, nil
}

func newOTExporter(conf *envvar.Configuration) (*prometheus.Exporter, error) {
	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		return nil, fmt.Errorf("runtime.Start %w", err)
	}
	// https://github.com/open-telemetry/opentelemetry-go/blob/main/CHANGELOG.md
	// The prometheus exporter now uses the new pull controller. (#751)
	metricExporter, err := stdoutmetric.New(
		stdoutmetric.WithPrettyPrint(),
	)

	if err != nil {
		return nil, fmt.Errorf("metricExporter.New %w", err)
	}

	pusher := controller.New(
		processor.New(
			simple.NewWithExactDistribution(),
			metricExporter,
		),
		controller.WithExporter(metricExporter),
		controller.WithCollectPeriod(5*time.Second),
	)
	promExporter, err := prometheus.New(prometheus.Config{}, pusher)
	if err != nil {
		return nil, fmt.Errorf("prometheus.NewExportPipeline %w", err)
	}

	global.SetMeterProvider(promExporter.MeterProvider())

	jaegerEndpoint, _ := conf.Get("JAEGER_ENDPOINT")
	collectorEndpointOption := jaeger.WithEndpoint(jaegerEndpoint)
	jaegerExporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(collectorEndpointOption),
	)
	if err != nil {
		return nil, fmt.Errorf("jaeger.NewRawExporter %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(jaegerExporter),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return promExporter, nil
}

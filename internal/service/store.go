package service

import (
	"context"
	"github.com/SaWLeaDeR/key-value-store/internal"
	"go.opentelemetry.io/otel/trace"
)

// StoreHandler
type StoreService struct {
	storage map[string]string
}

// NewStoreHandler
func NewStoreService(storage map[string]string) *StoreService {
	return &StoreService{
		storage: storage,
	}
}

func (s *StoreService) StoreData(ctx context.Context, key string, value string) internal.Data {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("LinksTracer").Start(ctx, "Data.StoreData")
	defer span.End()
	s.storage[key] = value
	return internal.Data{
		Key:   key,
		Value: value,
	}
}

func (s *StoreService) GetStoreData(ctx context.Context, key string) (internal.Data, error) {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("LinksTracer").Start(ctx, "Data.GetStoreData")
	defer span.End()

	value := s.storage[key]
	if value == "" {
		return internal.Data{}, internal.NewErrorf(internal.ErrorCodeUnknown, "Couldn't get value for %s", key)
	}
	return internal.Data{
		Key:   key,
		Value: value,
	}, nil
}

func (s *StoreService) Flush(ctx context.Context) error {
	ctx, span := trace.SpanFromContext(ctx).TracerProvider().Tracer("LinksTracer").Start(ctx, "Data.Flush")
	defer span.End()

	if len(s.storage) == 0 {
		return internal.NewErrorf(internal.ErrorCodeUnknown, "Map is Already Empty")
	}

	for k := range s.storage {
		delete(s.storage, k)
	}

	if len(s.storage) != 0 {
		return internal.NewErrorf(internal.ErrorCodeUnknown, "Flush operation occurs an error")
	}
	return nil
}

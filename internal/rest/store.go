package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SaWLeaDeR/key-value-store/internal"
)

//go:generate counterfeiter -o resttesting/store_service.gen.go . StoreService

const uuidRegEx string = `[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`

// StoreService
type StoreService interface {
	StoreData(ctx context.Context, key string, value string) internal.Data
	GetStoreData(ctx context.Context, key string) (internal.Data, error)
	Flush(ctx context.Context) error
}

// StoreHandler
type StoreHandler struct {
	svc StoreService
}

// NewStoreHandler
func NewStoreHandler(svc StoreService) *StoreHandler {
	return &StoreHandler{
		svc: svc,
	}
}

// Register connects the handlers to the router.
func (u *StoreHandler) Register(r *mux.Router) {
	r.HandleFunc("/store", u.saveStoreData).Methods(http.MethodPost)
	r.HandleFunc("/store/{key}", u.getStoredData).Methods(http.MethodGet)
	r.HandleFunc("/flush", u.flush).Methods(http.MethodGet)
}

// Data is an actor
type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SaveStoreDataRequest defines the request used for adding data to map.
type SaveStoreDataRequest struct {
	Value string `json:"value,omitempty"`
	Key   string `json:"key,omitempty"`
}

// GetKeyRequest defines the request used for get data from map.
type GetKeyRequest struct {
	KeyValue string `json:"key"`
}

// CreateStoreDataResponse defines the response will return back after data added to map.
type CreateStoreDataResponse struct {
	StoreData Data `json:"data"`
}

func (s *StoreHandler) saveStoreData(w http.ResponseWriter, r *http.Request) {
	var req SaveStoreDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		renderErrorResponse(r.Context(), w, "invalid request", internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder"))
		return
	}
	defer r.Body.Close()
	//
	data := s.svc.StoreData(r.Context(), req.Key, req.Value)

	renderResponse(w,
		&CreateStoreDataResponse{
			StoreData: Data{
				Key:   data.Key,
				Value: data.Value,
			},
		},
		http.StatusCreated)
}

func (u *StoreHandler) getStoredData(w http.ResponseWriter, r *http.Request) {
	key, _ := mux.Vars(r)["key"]

	data, err := u.svc.GetStoreData(r.Context(), key)
	fmt.Println(err)
	if err != nil {
		renderErrorResponse(r.Context(), w, "get store data occurs an error with given key", err)
		return
	}
	//
	renderResponse(w,
		&CreateStoreDataResponse{
			StoreData: Data{
				Key:   data.Key,
				Value: data.Value,
			},
		},
		http.StatusOK)
}

func (u *StoreHandler) flush(w http.ResponseWriter, r *http.Request) {
	err := u.svc.Flush(r.Context())
	fmt.Println(err)
	if err != nil {
		renderErrorResponse(r.Context(), w, "flush operation failed", err)
		return
	}
	//
	renderResponse(w, nil, http.StatusOK)
}

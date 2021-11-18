package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SaWLeaDeR/key-value-store/internal"
	"github.com/SaWLeaDeR/key-value-store/internal/rest"
	"github.com/SaWLeaDeR/key-value-store/internal/rest/resttesting"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/gorilla/mux"
)

func TestStoreData_Post(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(s *resttesting.FakeStoreService)
		input  []byte
		output output
	}{
		{
			"OK: 201",
			func(s *resttesting.FakeStoreService) {
				s.StoreDataReturns(
					internal.Data{
						Key:   "test_key",
						Value: "test_value",
					})
			},
			func() []byte {
				b, _ := json.Marshal(&rest.SaveStoreDataRequest{
					Key:   "test_key",
					Value: "test_value",
				})

				return b
			}(),
			output{
				http.StatusCreated,
				&rest.CreateStoreDataResponse{
					StoreData: rest.Data{
						Key:   "test_key",
						Value: "test_value",
					},
				},
				&rest.CreateStoreDataResponse{},
			},
		},
		{
			"ERR: 400",
			func(s *resttesting.FakeStoreService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Error: "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *resttesting.FakeStoreService) {
				s.StoreDataReturns(internal.Data{})
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Error: "",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	//-

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttesting.FakeStoreService{}
			tt.setup(svc)

			rest.NewStoreHandler(svc).Register(router)

			//-

			res := doRequest(router,
				httptest.NewRequest(http.MethodPost, "/store", bytes.NewReader(tt.input)))

			//-

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

type test struct {
	expected interface{}
	target   interface{}
}

func doRequest(router *mux.Router, req *http.Request) *http.Response {
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr.Result()
}

func assertResponse(t *testing.T, res *http.Response, test test) {
	t.Helper()

	if err := json.NewDecoder(res.Body).Decode(test.target); err != nil {
		t.Fatalf("couldn't decode %s", err)
	}
	defer res.Body.Close()

	if !cmp.Equal(test.expected, test.target, cmpopts.IgnoreUnexported()) {
		t.Fatalf("expected results don't match: %s", cmp.Diff(test.expected, test.target, cmpopts.IgnoreUnexported()))
	}
}

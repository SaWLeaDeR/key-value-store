package rest

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/gorilla/mux"
)

//go:generate go run ../../cmd/openapi-gen/main.go -path .
//go:generate oapi-codegen -package openapi3 -generate types  -o ../../pkg/openapi3/store.gen.go openapi3.yaml
//go:generate oapi-codegen -package openapi3 -generate client -o ../../pkg/openapi3/store.client.gen.go     openapi3.yaml

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "StoreData API",
			Description: "REST APIs used for interacting with the StoreData Service",
			Version:     "0.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/SaWLeaDeR/key-value-store/",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://localhost:9234",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Data": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("value", openapi3.NewStringSchema()).
				WithProperty("key", openapi3.NewStringSchema())),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"GetKeyRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for get value using key").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("key", openapi3.NewStringSchema().
						WithMinLength(1)),
				),
		},
		"SaveStoreDataRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for storing a data").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("value", openapi3.NewStringSchema().WithMinLength(1)).
					WithProperty("key", openapi3.NewStringSchema().WithMinLength(1)),
				),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"CreateStoreDataResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after store data or get data from data store").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("data", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Data",
					}))),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/store": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "SaveData",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/SaveStoreDataRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateStoreDataResponse",
					},
				},
			},
		},
		"/flush": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "Flush",
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateStoreDataResponse",
					},
				},
			},
		},
		"/store/{key}": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "GetValueUsingKey",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("key").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateStoreDataResponse",
					},
				},
			},
		},
	}

	return swagger
}

func RegisterOpenAPI(r *mux.Router) {
	swagger := NewOpenAPI3()

	r.HandleFunc("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		renderResponse(w, &swagger, http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger)

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}

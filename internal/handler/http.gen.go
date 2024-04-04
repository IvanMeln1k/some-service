// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package handler

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	domain "github.com/IvanMeln1k/some-service/internal/domain"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Message defines model for Message.
type Message struct {
	Message string `json:"message"`
}

// User defines model for User.
type User = domain.User

// PostApiV1UserJSONBody defines parameters for PostApiV1User.
type PostApiV1UserJSONBody struct {
	Email    openapi_types.Email `json:"email"`
	Name     string              `json:"name"`
	Username string              `json:"username"`
}

// PostApiV1UserJSONRequestBody defines body for PostApiV1User for application/json ContentType.
type PostApiV1UserJSONRequestBody PostApiV1UserJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/v1/photo)
	GetApiV1Photo(ctx echo.Context) error

	// (GET /api/v1/user)
	GetApiV1User(ctx echo.Context) error

	// (POST /api/v1/user)
	PostApiV1User(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetApiV1Photo converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiV1Photo(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiV1Photo(ctx)
	return err
}

// GetApiV1User converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiV1User(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiV1User(ctx)
	return err
}

// PostApiV1User converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiV1User(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiV1User(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/v1/photo", wrapper.GetApiV1Photo)
	router.GET(baseURL+"/api/v1/user", wrapper.GetApiV1User)
	router.POST(baseURL+"/api/v1/user", wrapper.PostApiV1User)

}

type GetApiV1PhotoRequestObject struct {
}

type GetApiV1PhotoResponseObject interface {
	VisitGetApiV1PhotoResponse(w http.ResponseWriter) error
}

type GetApiV1Photo200ImagepngResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response GetApiV1Photo200ImagepngResponse) VisitGetApiV1PhotoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "image/png")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type GetApiV1Photo500JSONResponse Message

func (response GetApiV1Photo500JSONResponse) VisitGetApiV1PhotoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetApiV1UserRequestObject struct {
}

type GetApiV1UserResponseObject interface {
	VisitGetApiV1UserResponse(w http.ResponseWriter) error
}

type GetApiV1User200JSONResponse User

func (response GetApiV1User200JSONResponse) VisitGetApiV1UserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetApiV1User500JSONResponse Message

func (response GetApiV1User500JSONResponse) VisitGetApiV1UserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostApiV1UserRequestObject struct {
	Body *PostApiV1UserJSONRequestBody
}

type PostApiV1UserResponseObject interface {
	VisitPostApiV1UserResponse(w http.ResponseWriter) error
}

type PostApiV1User200JSONResponse struct {
	Id *openapi_types.UUID `json:"id,omitempty"`
}

func (response PostApiV1User200JSONResponse) VisitPostApiV1UserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostApiV1User500JSONResponse Message

func (response PostApiV1User500JSONResponse) VisitPostApiV1UserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /api/v1/photo)
	GetApiV1Photo(ctx context.Context, request GetApiV1PhotoRequestObject) (GetApiV1PhotoResponseObject, error)

	// (GET /api/v1/user)
	GetApiV1User(ctx context.Context, request GetApiV1UserRequestObject) (GetApiV1UserResponseObject, error)

	// (POST /api/v1/user)
	PostApiV1User(ctx context.Context, request PostApiV1UserRequestObject) (PostApiV1UserResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetApiV1Photo operation middleware
func (sh *strictHandler) GetApiV1Photo(ctx echo.Context) error {
	var request GetApiV1PhotoRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiV1Photo(ctx.Request().Context(), request.(GetApiV1PhotoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiV1Photo")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetApiV1PhotoResponseObject); ok {
		return validResponse.VisitGetApiV1PhotoResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetApiV1User operation middleware
func (sh *strictHandler) GetApiV1User(ctx echo.Context) error {
	var request GetApiV1UserRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetApiV1User(ctx.Request().Context(), request.(GetApiV1UserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetApiV1User")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetApiV1UserResponseObject); ok {
		return validResponse.VisitGetApiV1UserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostApiV1User operation middleware
func (sh *strictHandler) PostApiV1User(ctx echo.Context) error {
	var request PostApiV1UserRequestObject

	var body PostApiV1UserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostApiV1User(ctx.Request().Context(), request.(PostApiV1UserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostApiV1User")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostApiV1UserResponseObject); ok {
		return validResponse.VisitPostApiV1UserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xVzWobSRB+FVG7x5FG2mVhmZv3svhgMCy7l0WH1qg8amemu93dIyLMgH8OMcRgyAuE",
	"kBdwTIRNjORXqH6j0D2j2JYUojg5hFykprqqvqr6vq45hFQWSgoU1kByCCYdYcHCcQeNYRn6o9JSobYc",
	"w0Vxf2EnCiEBYzUXGVRVBBoPSq5xCMn/nxz70cJRDvYxtVBF8K9BvZoaC8Zzf9iTumAWksYSLSNFwIeP",
	"/MqSD9e5CVasqzSC0qD+zOVSGyHxoo4Q8iB6pbcInrcz2W6MQ1kwLjqh2wc3bV4oqa3HrmtoHCECxewI",
	"Esi4HZWDTiqLeHvMxA7movcsNrLAtkE95inGXFhfRB43sVXlK+diT/q8QzSp5spyKSCBrYEsbesf1BlO",
	"/Ji4zT1qbdja3YYIxqhN7dzrdDtdPyOpUDDFIYHfg6muLjAVM8XjcS9WI2kDXoZ2FZZe05xu3al7QTfu",
	"xJ23fAOtOiak18x7bg8hgb/Rbin+X2+3udVolBSmFsZv3a7/S6WwKAIQL1iGsRLZvWwfCWLABdOTVUn4",
	"GS0V+dYd0x1N3RnNaO77/mMFjSmV8zQUG+8bH/YQ9FeNe5DAL/H9Y4qblxQvntE64Fc0c6fuxB3RlGY0",
	"cxfuokVzd0Y39I4+0GXLHdPUHdFV+L0MyrQsM16V9Zj63rTgomze1GZU0Hu6DKgvadqieU1NWQt1PTON",
	"ir9AzNNHFfL/2ARtTE/opV9FoKRZx8YbmtO1p6Dm4i6Qc07XNKcrb6Qp3bqLFSp2pVni4qBEY/+Sw8lX",
	"DeTJa/c77NONV2m90L5Jbo/73OibUa0p4+dYGY0mvcl/Q1B76yGUOocERtaqJI5zmbJ8JI1N/ux2u1D1",
	"q48BAAD//6ePk8wjCAAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
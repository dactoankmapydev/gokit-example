package appsvc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	// "fmt"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

var (
	ErrParameter = errors.New("The input parameter must be an integer")
)

// NewHandler returns new http.Handler that routes http request to service
func NewHandler(s Service, router *httprouter.Router) http.Handler {

	router.Handler(http.MethodGet, "/api/v1/main-apps", kithttp.NewServer(
		makeGetMainAppEndpoint(s),
		decodeGetMainAppRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/main-apps/", kithttp.NewServer(
		makeRegisterAppEndpoint(s),
		decodeRegisterAppRequests,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/main-apps/:id", kithttp.NewServer(
		makeGetAppDetailEndpoint(s),
		decodeGetAppDetailRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/main-apps/:id/mini-apps", kithttp.NewServer(
		makeGetMiniofMainAppEndpoint(s),
		decodeGetMiniofMainAppRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/mini-apps", kithttp.NewServer(
		makeGetMiniAppEndpoint(s),
		decodeGetMiniAppRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/mini-apps/:id", kithttp.NewServer(
		makeGetMiniAppDetailEndpoint(s),
		decodeGetMiniAppDetailRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/mini-apps/", kithttp.NewServer(
		makeCreateMiniAppEndpoint(s),
		decodeCreateMiniAppRequests,
		kithttp.EncodeJSONResponse,
	))

	return router
}

func decodeGetMainAppRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	req := GetMainAppRequest{Cursor: q.Get("cursor")}
	if qLimit := q.Get("limit"); qLimit != "" {
		intLimit, err := strconv.Atoi(qLimit)
		if err != nil {
			return nil, ErrParameter
		}
		req.Limit = intLimit
	}
	return req, nil
}

func decodeRegisterAppRequests(ctx context.Context, r *http.Request) (interface{}, error) {
	var request RegisterAppRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetAppDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return GetAppDetailRequest{}, nil
}

func decodeGetMiniofMainAppRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	req := GetMiniAppOfAppRequest{Cursor: q.Get("cursor")}
	if qLimit := q.Get("limit"); qLimit != "" {
		intLimit, err := strconv.Atoi(qLimit)
		if err != nil {
			return nil, ErrParameter
		}
		req.Limit = intLimit
	}
	return req, nil
}

func decodeGetMiniAppRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	req := GetMiniAppRequest{Cursor: q.Get("cursor")}
	if qLimit := q.Get("limit"); qLimit != "" {
		intLimit, err := strconv.Atoi(qLimit)
		if err != nil {
			return nil, ErrParameter
		}
		req.Limit = intLimit
	}
	return req, nil
}

func decodeGetMiniAppDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return GetMiniAppDetailRequest{}, nil
}

func decodeCreateMiniAppRequests(ctx context.Context, r *http.Request) (interface{}, error) {
	var request CreateMiniAppRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

package appsvc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

var (
	ErrParameter = errors.New("The input parameter must be an integer")
	ErrEmpty     = errors.New("No input parameter")
)

// NewHandler returns new http.Handler that routes http request to service
func NewHandler(s Service, router *httprouter.Router) http.Handler {

	router.Handler(http.MethodGet, "/api/:ver/apps", kithttp.NewServer(
		makeGetAppsEndpoint(s),
		decodeGetAppsRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/main-apps/", kithttp.NewServer(
		makeRegisterAppEndpoint(s),
		decodeRegisterAppRequests,
		kithttp.EncodeJSONResponse,
	))

	// router.Handler(http.MethodPost, "/api/v1/main-apps/{id}", kithttp.NewServer(
	// 	makeGetAppDetailEndpoint(s),
	// 	decodeGetAppDetailRequest,
	// 	kithttp.EncodeJSONResponse,
	// ))

	return router
}

func decodeGetAppsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	req := GetAppsRequest{Cursor: q.Get("cursor")}
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

// func decodeGetAppDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
// 	reId := r.URL.Query.Get("id")
// 	return GetAppDetailRequest{Id: reId}, nil

// }

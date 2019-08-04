package appsvc

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

// NewHandler returns new http.Handler that routes http request to service
func NewHandler(s Service, router *httprouter.Router) http.Handler {

	router.Handler(http.MethodGet, "/api/:ver/apps", kithttp.NewServer(
		makeGetAppsEndpoint(s),
		decodeGetAppsRequest,
		kithttp.EncodeJSONResponse,
	))

	return router
}

func decodeGetAppsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
		// limit
	str_limit := r.URL.Query().Get("limit")
	int_limit, err := strconv.Atoi(str_limit)
	if err != nil {
		// log.Fatal(err.Error())
		log.Println(err)
	}

	// cursor
	str_cursor := r.URL.Query().Get("cursor")
	return GetAppsRequest{Limit: int_limit, Cursor: str_cursor}, nil
}

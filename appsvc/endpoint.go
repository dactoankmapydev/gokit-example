package appsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetAppsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetApps(ctx, req.(GetAppsRequest))
	}
}

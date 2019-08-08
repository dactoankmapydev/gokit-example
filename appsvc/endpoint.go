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

func makeRegisterAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.RegisterApp(ctx, req.(RegisterAppRequest))
	}
}

// func makeGetAppDetailEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, req interface{}) (interface{}, error) {
// 		return s.GetAppDetail(ctx, req.(GetAppDetailRequest))
// 	}
// }

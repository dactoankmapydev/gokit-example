package appsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetMainAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetMainApp(ctx, req.(GetMainAppRequest))
	}
}

func makeRegisterAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.RegisterApp(ctx, req.(RegisterAppRequest))
	}
}

func makeGetAppDetailEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAppDetail(ctx, req.(GetAppDetailRequest))
	}
}

func makeGetMiniofMainAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetMiniofMainApp(ctx, req.(GetMiniAppOfAppRequest))
	}
}

func makeGetMiniAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetMiniApp(ctx, req.(GetMiniAppRequest))
	}
}

func makeGetMiniAppDetailEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetMiniAppDetail(ctx, req.(GetMiniAppDetailRequest))
	}
}

func makeCreateMiniAppEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CreateMiniApp(ctx, req.(CreateMiniAppRequest))
	}
}

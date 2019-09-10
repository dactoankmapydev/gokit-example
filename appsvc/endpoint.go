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

func makeGetMininAppEndpoint(s Service) endpoint.Endpoint {
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

func makeUpdateMiniAppOfMainApp(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateMiniAppOfMainApp(ctx, req.(UpdateMiniAppOfMainAppRequest))
	}
}

func makeUpdateMiniApp(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateMiniApp(ctx, req.(UpdateMiniAppRequest))
	}
}

func makeDeployMiniApp(s Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeployMiniApp(ctx, req.(DeployMiniAppRequest))
	}
}

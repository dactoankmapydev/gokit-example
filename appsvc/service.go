package appsvc

import (
	"context"
	"miniapp_backend/app"
)

type Service interface {
	RegisterApp(context.Context, RegisterAppRequest) (*RegisterAppResponse, error)
	GetApps(context.Context, GetAppsRequest) (*GetAppsResponse, error)
	// GetAppDetail(context.Context, GetAppDetailRequest) (*GetAppDetailResponse, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) GetApps(ctx context.Context, r GetAppsRequest) (*GetAppsResponse, error) {

	var sliceApp []app.App
	sliceApp = append(sliceApp, app.App{})

	return &GetAppsResponse{
		Total:  r.Limit,
		Apps:   sliceApp,
		Cursor: Cursor{},
	}, nil
}

func (s *service) RegisterApp(Ctx context.Context, r RegisterAppRequest) (*RegisterAppResponse, error) {

	return &RegisterAppResponse{
		App: app.App{
			r.Id,
			r.Platform,
			r.BundleId,
			r.PackageName,
			r.Name,
			r.GooglePlayUrl,
			r.AppStoreUrl,
			r.Icon,
			r.Version,
		},
		Ica: "4",
	}, nil
}

// func (s *service) GetAppDetail(Ctx context.Context, r GetAppDetailRequest) (*GetAppDetailResponse, error) {
// 	return &GetAppDetailResponse{}
// }

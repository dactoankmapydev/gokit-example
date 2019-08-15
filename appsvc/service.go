package appsvc

import (
	"context"
	"miniapp_backend/app"
	"reflect"
)

type Service interface {
	RegisterApp(context.Context, RegisterAppRequest) (*RegisterAppResponse, error)
	GetMainApp(context.Context, GetMainAppRequest) (*GetMainAppResponse, error)
	GetAppDetail(context.Context, GetAppDetailRequest) (*GetAppDetailResponse, error)
	GetMiniofMainApp(context.Context, GetMiniAppOfAppRequest) (*GetMiniAppOfAppResponse, error)
	GetMiniApp(context.Context, GetMiniAppRequest) (*GetMiniAppResponse, error)
	GetMiniAppDetail(context.Context, GetMiniAppDetailRequest) (*GetMiniAppDetailResponse, error)
	CreateMiniApp(context.Context, CreateMiniAppRequest) (*CreateMiniAppResponse, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) GetMainApp(ctx context.Context, r GetMainAppRequest) (*GetMainAppResponse, error) {
	var sliceApp []app.MainApp
	sliceApp = append(sliceApp, app.MainApp{
		Id:          "1",
		Platform:    "ios",
		PackageName: "appsvc",
		Version:     "v1",
		Events:      []string{"event 1"},
	})
	return &GetMainAppResponse{
		Total:  r.Limit,
		Apps:   sliceApp,
		Cursor: Cursor{},
	}, nil
}

func (s *service) RegisterApp(Ctx context.Context, r RegisterAppRequest) (*RegisterAppResponse, error) {
	v := reflect.ValueOf(r)
	typeOfs := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfs.Field(i).Name
		if valueField := v.Field(i).Interface(); valueField != "" {
			// Todo save to db
			println(field, "exist")
		}
	}
	return &RegisterAppResponse{
		App: app.MainApp{
			BundleId:      r.BundleId,
			Platform:      r.Platform,
			PackageName:   r.PackageName,
			Name:          r.Name,
			GooglePlayUrl: r.GooglePlayUrl,
			AppStoreUrl:   r.AppStoreUrl,
			Icon:          r.Icon,
			Version:       r.Version,
		},
		Ica: "4",
	}, nil
}

func (s *service) GetAppDetail(Ctx context.Context, r GetAppDetailRequest) (*GetAppDetailResponse, error) {
	return &GetAppDetailResponse{
		App: app.MainApp{},
		Ica: "5",
	}, nil
}

func (s *service) GetMiniofMainApp(Ctx context.Context, r GetMiniAppOfAppRequest) (*GetMiniAppOfAppResponse, error) {
	var sliceMiniApp []app.MiniApp
	sliceMiniApp = append(sliceMiniApp, app.MiniApp{
		Id:          "2",
		Platform:    "android",
		PackageName: "appsvc",
		Version:     "v1",
		Permissions: []string{},
	})
	return &GetMiniAppOfAppResponse{
		Total:  r.Limit,
		Apps:   sliceMiniApp,
		Cursor: Cursor{},
	}, nil
}

func (s *service) GetMiniApp(ctx context.Context, r GetMiniAppRequest) (*GetMiniAppResponse, error) {
	var sliceApp []app.MiniApp
	sliceApp = append(sliceApp, app.MiniApp{
		Id:          "5",
		Platform:    "android",
		PackageName: "appsvc",
		Version:     "v3",
		Permissions: []string{},
	})
	return &GetMiniAppResponse{
		Total:  r.Limit,
		Apps:   sliceApp,
		Cursor: Cursor{},
	}, nil
}

func (s *service) GetMiniAppDetail(Ctx context.Context, r GetMiniAppDetailRequest) (*GetMiniAppDetailResponse, error) {
	return &GetMiniAppDetailResponse{
		App: app.MiniApp{
			Permissions: []string{},
		},
		History: []string{},
	}, nil
}

func (s *service) CreateMiniApp(Ctx context.Context, r CreateMiniAppRequest) (*CreateMiniAppResponse, error) {
	v := reflect.ValueOf(r)
	typeOfs := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfs.Field(i).Name
		if valueField := v.Field(i).Interface(); valueField != "" {
			// Todo save to db
			println(field, "exist")
		}
	}
	return &CreateMiniAppResponse{
		App: app.MiniApp{
			BundleId:      r.BundleId,
			Platform:      r.Platform,
			PackageName:   r.PackageName,
			DisplayName:   r.DisplayName,
			AppName:       r.AppName,
			Type:          r.Type,
			TargetVersion: r.TargetVersion,
			Icon:          r.Icon,
			Version:       r.Version,
			Permissions:   r.Permissions,
			Bundle:        r.Bundle,
		},
	}, nil
}

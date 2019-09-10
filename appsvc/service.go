package appsvc

import (
	"context"
	"fmt"
	"miniapp_backend/app"

	"github.com/julienschmidt/httprouter"
)

type Service interface {
	RegisterApp(context.Context, RegisterAppRequest) (*RegisterAppResponse, error)
	GetMainApp(context.Context, GetMainAppRequest) (*GetMainAppResponse, error)
	GetAppDetail(context.Context, GetAppDetailRequest) (*GetAppDetailResponse, error)
	GetMiniofMainApp(context.Context, GetMiniAppOfAppRequest) (*GetMiniAppOfAppResponse, error)
	UpdateMiniAppOfMainApp(context.Context, UpdateMiniAppOfMainAppRequest) (*UpdateMiniAppOfMainAppResponse, error)
	GetMiniApp(context.Context, GetMiniAppRequest) (*GetMiniAppResponse, error)
	GetMiniAppDetail(context.Context, GetMiniAppDetailRequest) (*GetMiniAppDetailResponse, error)
	CreateMiniApp(context.Context, CreateMiniAppRequest) (*CreateMiniAppResponse, error)
	UpdateMiniApp(context.Context, UpdateMiniAppRequest) (*UpdateMiniAppResponse, error)
	DeployMiniApp(context.Context, DeployMiniAppRequest) (*DeployMiniAppResponse, error)
}

type service struct {
	AppRepo app.Repository
}

func New(appRepo app.Repository) Service {

	return &service{
		AppRepo: appRepo,
	}
}

func (s *service) GetMainApp(Ctx context.Context, r GetMainAppRequest) (*GetMainAppResponse, error) {
	listResult, err := s.AppRepo.GetAllMainApp()
	if err != nil {
		panic(err)
	}
	return &GetMainAppResponse{
		Total:  r.Limit,
		Apps:   listResult,
		Cursor: Cursor{},
	}, nil
}

func (s *service) GetAppDetail(Ctx context.Context, r GetAppDetailRequest) (*GetAppDetailResponse, error) {
	fmt.Println(r.Id)
	result, err := s.AppRepo.GetMainAppID(r.Id)
	if err != nil {
		panic(err)
	}
	return &GetAppDetailResponse{
		App: app.MainApp{
			Id:            result.Id,
			Platform:      result.Platform,
			AppStoreUrl:   result.AppStoreUrl,
			PackageName:   result.PackageName,
			Name:          result.Name,
			GooglePlayUrl: result.GooglePlayUrl,
			Icon:          result.Icon,
			Version:       result.Version,
			BundleId:      result.BundleId,
			Events:        result.Events,
		},
		Ica: "",
	}, nil
}

func (s *service) RegisterApp(Ctx context.Context, r RegisterAppRequest) (*RegisterAppResponse, error) {
	registerApp := app.MainApp{
		BundleId:      r.BundleId,
		Platform:      r.Platform,
		PackageName:   r.PackageName,
		Name:          r.Name,
		GooglePlayUrl: r.GooglePlayUrl,
		AppStoreUrl:   r.AppStoreUrl,
		Icon:          r.Icon,
		Version:       r.Version,
	}
	id, err := s.AppRepo.CreateMainApp(registerApp)
	if err != nil {
		panic(err)
	}
	return &RegisterAppResponse{
		App: app.MainApp{
			Id:            id,
			BundleId:      r.BundleId,
			Platform:      r.Platform,
			PackageName:   r.PackageName,
			Name:          r.Name,
			GooglePlayUrl: r.GooglePlayUrl,
			AppStoreUrl:   r.AppStoreUrl,
			Icon:          r.Icon,
			Version:       r.Version,
		},
		Ica: "",
	}, nil
}

func (s *service) GetMiniofMainApp(Ctx context.Context, r GetMiniAppOfAppRequest) (*GetMiniAppOfAppResponse, error) {
	id := httprouter.ParamsFromContext(Ctx).ByName("id")
	result, err := s.AppRepo.GetMiniofMainApp(id)
	if err != nil {
		panic(err)
	}
	return &GetMiniAppOfAppResponse{
		Total: r.Limit,
		Apps: app.MiniApp{
			Id:            result.Id,
			Platform:      result.Platform,
			BundleId:      result.BundleId,
			PackageName:   result.PackageName,
			DisplayName:   result.DisplayName,
			AppName:       result.AppName,
			Status:        result.Status,
			Type:          result.Type,
			Icon:          result.Icon,
			TargetVersion: result.TargetVersion,
			Version:       result.Version,
			Permissions:   result.Permissions,
		},
		Cursor: Cursor{},
	}, nil
}

func (s *service) GetMiniApp(Ctx context.Context, r GetMiniAppRequest) (*GetMiniAppResponse, error) {
	listResult, err := s.AppRepo.GetAllMiniApp()
	if err != nil {
		panic(err)
	}
	return &GetMiniAppResponse{
		Total:  r.Limit,
		Apps:   listResult,
		Cursor: Cursor{},
	}, nil
}

func (s *service) GetMiniAppDetail(Ctx context.Context, r GetMiniAppDetailRequest) (*GetMiniAppDetailResponse, error) {
	fmt.Println(r.Id)
	result, err := s.AppRepo.GetMiniAppID(r.Id)
	if err != nil {
		panic(err)
	}
	return &GetMiniAppDetailResponse{
		App: app.MiniApp{
			Id:            result.Id,
			Platform:      result.Platform,
			BundleId:      result.BundleId,
			PackageName:   result.PackageName,
			DisplayName:   result.DisplayName,
			AppName:       result.AppName,
			Status:        result.Status,
			Type:          result.Type,
			Icon:          result.Icon,
			TargetVersion: result.TargetVersion,
			Version:       result.Version,
			Permissions:   result.Permissions,
		},
		History: []string{},
	}, nil

}

func (s *service) CreateMiniApp(Ctx context.Context, r CreateMiniAppRequest) (*CreateMiniAppResponse, error) {
	createMiniApp := app.MiniApp{
		Platform:      r.Platform,
		BundleId:      r.BundleId,
		PackageName:   r.PackageName,
		DisplayName:   r.DisplayName,
		AppName:       r.AppName,
		Type:          r.Type,
		TargetVersion: r.TargetVersion,
		Icon:          r.Icon,
		Version:       r.Version,
		Permissions:   r.Permissions,
	}
	id, err := s.AppRepo.CreateMiniApp(createMiniApp)
	if err != nil {
		panic(err)
	}
	return &CreateMiniAppResponse{
		App: app.MiniApp{
			Id:            id,
			Platform:      r.Platform,
			BundleId:      r.BundleId,
			PackageName:   r.PackageName,
			DisplayName:   r.DisplayName,
			AppName:       r.AppName,
			Type:          r.Type,
			TargetVersion: r.TargetVersion,
			Icon:          r.Icon,
			Version:       r.Version,
			Permissions:   r.Permissions,
		},
	}, nil
}

func (s *service) UpdateMiniAppOfMainApp(Ctx context.Context, r UpdateMiniAppOfMainAppRequest) (*UpdateMiniAppOfMainAppResponse, error) {
	return &UpdateMiniAppOfMainAppResponse{
		App: app.MiniApp{
			Status:      r.Status,
			Permissions: []string{},
		},
	}, nil
}

func (s *service) UpdateMiniApp(Ctx context.Context, r UpdateMiniAppRequest) (*UpdateMiniAppResponse, error) {
	fmt.Println(r.Id)

	updateMiniApp := app.MiniApp{
		Platform:      r.Platform,
		BundleId:      r.BundleId,
		PackageName:   r.PackageName,
		DisplayName:   r.DisplayName,
		AppName:       r.AppName,
		Type:          r.Type,
		TargetVersion: r.TargetVersion,
		Icon:          r.Icon,
		Version:       r.Version,
		Permissions:   r.Permissions,
	}
	result, resultID, err := s.AppRepo.Update(r.Id, updateMiniApp)
	if err != nil {
		panic(err)
	}
	return &UpdateMiniAppResponse{
		App: app.MiniApp{
			Id:            resultID,
			Platform:      result.Platform,
			BundleId:      result.BundleId,
			PackageName:   result.PackageName,
			DisplayName:   result.DisplayName,
			AppName:       result.AppName,
			Type:          result.Type,
			TargetVersion: result.TargetVersion,
			Icon:          result.Icon,
			Version:       result.Version,
			Permissions:   result.Permissions,
		},
	}, nil
}

func (s *service) DeployMiniApp(Ctx context.Context, r DeployMiniAppRequest) (*DeployMiniAppResponse, error) {
	fmt.Println(r.Id)
	result, err := s.AppRepo.GetMiniofMainApp(r.Id)
	if err != nil {
		panic(err)
	}
	return &DeployMiniAppResponse{
		App: app.MiniApp{
			Id:       result.Id,
			Platform: result.Platform,
			Version:  result.Version,
		},
	}, nil
}

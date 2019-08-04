package appsvc

import (
	"context"
	"miniapp_backend/app"
)

type Service interface {
	GetApps(context.Context, GetAppsRequest) (*GetAppsResponse, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) GetApps(ctx context.Context, r GetAppsRequest) (*GetAppsResponse, error) {

	return &GetAppsResponse{
		Total: r.Limit,
		Apps:  []app.App{},
	}, nil
}

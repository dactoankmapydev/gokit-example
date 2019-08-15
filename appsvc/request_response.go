package appsvc

import (
	"miniapp_backend/app"
)

type Cursor struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type GetMainAppRequest struct {
	Limit     int
	Offset    int
	Cursor    string
	Ownership string
}

type GetMainAppResponse struct {
	Total  int           `json:"total"`
	Apps   []app.MainApp `json:"apps"`
	Cursor Cursor        `json:"cursor"`
}

type RegisterAppRequest struct {
	Platform      app.Platform
	BundleId      string
	PackageName   string
	Name          string
	GooglePlayUrl string
	AppStoreUrl   string
	Icon          string
	Version       string
}

type RegisterAppResponse struct {
	App app.MainApp `json:"app"`
	Ica string      `json:"ica"` // intermediate cert
}

type GetAppDetailRequest struct {
	Id string
}

type GetAppDetailResponse struct {
	App app.MainApp `json:"app"`
	Ica string      `json:"ica"`
}

type GetMiniAppOfAppRequest struct {
	Limit  int
	Cursor string
	Status app.Status
}

type GetMiniAppOfAppResponse struct {
	Total  int           `json:"total"`
	Apps   []app.MiniApp `json:"apps"`
	Cursor Cursor        `json:"cursor"`
}

type UpdateMiniAppOfAppRequest struct {
	MainAppID string
	MiniAppID string
	Status    app.MiniAppType
}

type GetMiniAppRequest struct {
	Limit  int
	Offset int
	Cursor string
	Status app.MiniAppType
}

type GetMiniAppResponse struct {
	Total  int           `json:"total"`
	Apps   []app.MiniApp `json:"apps"`
	Cursor Cursor        `json:"cursor"`
}

type GetMiniAppDetailRequest struct {
	Id string
}

type GetMiniAppDetailResponse struct {
	App     app.MiniApp `json:"app"`
	History []string    `json:"history"`
}

type CreateMiniAppRequest struct {
	Platform      app.Platform
	BundleId      string
	PackageName   string
	DisplayName   string
	AppName       string
	Type          app.MiniAppType
	TargetVersion string
	Icon          string
	Version       string
	Permissions   []string
	Bundle        app.BundleType
}

type CreateMiniAppResponse struct {
	App app.MiniApp `json:"app"`
}

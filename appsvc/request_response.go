package appsvc

import (
	"miniapp_backend/app"
)

type Cursor struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type GetAppsRequest struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type GetAppsResponse struct {
	Total  int       `json:"total"`
	Apps   []app.App `json:"apps"`
	Cursor Cursor    `json:"cursor"`
}

type RegisterAppRequest struct {
	Id            string `json:"id"`
	Platform      string `json:"platform"`
	BundleId      string `json:"bundleId"`
	PackageName   string `json:"packageName"`
	Name          string `json:"name"`
	GooglePlayUrl string `json:"googlePlayUrl"`
	AppStoreUrl   string `json:"appStoreUrl"`
	Icon          string `json:"icon"`
	Version       string `json:"version"`
}

type RegisterAppResponse struct {
	App app.App `json:"app"`
	Ica string  `json:"ica"`
}

// type GetAppDetailRequest struct {
// 	Id string `json:"string"`
// }

// type GetAppDetailResponse struct {
// 	App app.App `json:"app"`
// 	ica string  `json:"string"`
// }

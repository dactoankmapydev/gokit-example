package appsvc

import (
	"miniapp_backend/app"
)

type GetAppsRequest struct {
	Limit int
}
type GetAppsResponse struct {
	Total int
	Apps  []app.App
}

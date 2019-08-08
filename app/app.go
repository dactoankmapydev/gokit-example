package app

type App struct {
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

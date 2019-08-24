package app

// App
type MainApp struct {
	Id            string   `json:"id"`
	Platform      string   `json:"platform"`
	BundleId      string   `json:"bundleId"`
	PackageName   string   `json:"packageName"`
	Name          string   `json:"name"`
	GooglePlayUrl string   `json:"googlePlayUrl"`
	AppStoreUrl   string   `json:"appStoreUrl"`
	Icon          string   `json:"icon"`
	Version       string   `json:"version"`
	Events        []string `json:"events"`
}

// MiniApp
type MiniApp struct {
	Id            string   `json:"id"`
	Platform      string   `json:"platform"`
	BundleId      string   `json:"bundleId"`
	PackageName   string   `json:"packageName"`
	DisplayName   string   `json:"displayname"`
	AppName       string   `json:"appName"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	Icon          string   `json:"icon"`
	TargetVersion string   `json:"targetVersion"`
	Version       string   `json:"version"`
	Permissions   []string `json:"permissions"`
}

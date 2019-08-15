package app

// Platform
type Platform string

const (
	Android Platform = "android"
	Ios     Platform = "ios"
)

// Status
type Status string

const (
	Develop   Status = "develop"
	Staging   Status = "staging"
	Published Status = "published"
)

// MiniAppType
type MiniAppType string

const (
	InApp      MiniAppType = "in-app"
	Background MiniAppType = "background"
)

// MiniAppStatus
type MiniAppStatus string

const (
	Approved MiniAppStatus = "approved"
	Rejected MiniAppStatus = "rejected"
	Pending  MiniAppStatus = "pending"
)

// BundleType
type BundleType string

const (
	Multipart BundleType = "multipart"
	Formdata  BundleType = "form-data"
)

// App
type MainApp struct {
	Id            string   `json:"id"`
	Platform      Platform `json:"platform"`
	BundleId      string   `json:"bundleId"`
	PackageName   string   `json:"packageName"`
	Name          string   `json:"name"`
	GooglePlayUrl string   `json:"googlePlayUrl"`
	AppStoreUrl   string   `json:"appStoreUrl"`
	Icon          string   `json:"icon"`
	Version       string   `json:"version"`
	Events        []string `json:"event"`
}

// MiniApp
type MiniApp struct {
	Id            string      `json:"id"`
	Platform      Platform    `json:"platform"`
	BundleId      string      `json:"bundleId"`
	PackageName   string      `json:"packageName"`
	DisplayName   string      `json:"displayname"`
	AppName       string      `json:"appName"`
	Status        Status      `json:"status"`
	Type          MiniAppType `json:"type"`
	Icon          string      `json:"icon"`
	TargetVersion string      `json:"targetVersion"`
	Version       string      `json:'version"`
	Bundle        BundleType  `json:'bundle"`
	Permissions   []string    `json:"permissions"`
}

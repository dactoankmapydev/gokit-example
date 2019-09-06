package app

import "go.mongodb.org/mongo-driver/bson/primitive"

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// primitive.ObjectID
// App
type MainApp struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Platform      string             `json:"platform" bson:"platform"`
	BundleId      string             `json:"bundleId" bson:"bundleId"`
	PackageName   string             `json:"packageName" bson:"packageName"`
	Name          string             `json:"name" bson:"name"`
	GooglePlayUrl string             `json:"googlePlayUrl" bson:"googlePlayUrl"`
	AppStoreUrl   string             `json:"appStoreUrl" bson:"appStoreUrl"`
	Icon          string             `json:"icon" bson:"icon"`
	Version       string             `json:"version" bson:"version"`
	Events        []string           `json:"events" bson:"events"`
}

// MiniApp
type MiniApp struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Platform      string             `json:"platform" bson:"platform"`
	BundleId      string             `json:"bundleId" bson:"bundleId"`
	PackageName   string             `json:"packageName" bson:"packageName"`
	DisplayName   string             `json:"displayName" bson:"displayName"`
	AppName       string             `json:"appName" bson:"appName"`
	Status        string             `json:"status" bson:"status"`
	Type          string             `json:"type" bson:"type"`
	Icon          string             `json:"icon" bson:"icon"`
	TargetVersion string             `json:"targetVersion" bson:"targetVersion"`
	Version       string             `json:"version" bson:"version"`
	Permissions   []string           `json:"permissions" bson:"permissions"`
}

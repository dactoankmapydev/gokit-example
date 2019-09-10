package app

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repoImpl struct {
	Db *mongo.Database
}

type Repository interface {
	CreateMainApp(main MainApp) (string, error)
	GetMainAppID(id string) (MainApp, error)
	GetAllMainApp() ([]MainApp, error)
	CreateMiniApp(mini MiniApp) (string, error)
	GetMiniAppID(id string) (MiniApp, error)
	GetAllMiniApp() ([]MiniApp, error)
	GetMiniofMainApp(id string) (MiniApp, error)
	DeployMiniApp(id string) (MiniApp, error)
	Update(id string, mini MiniApp) (MiniApp, string, error)
	// UpdateMiniAppOfMainApp(mainId,miniId string, mini MiniApp) (MiniApp, string, error)
}

func NewRepo(db *mongo.Database) Repository {
	return &repoImpl{
		Db: db,
	}
}

func (mongo *repoImpl) Update(id string, mini MiniApp) (MiniApp, string, error) {
	filter := bson.M{"_id": id}
	fmt.Println(filter)
	update := bson.M{
		"$set": bson.M{
			"appName":       mini.AppName,
			"platform":      mini.Platform,
			"bundleId":      mini.BundleId,
			"packageName":   mini.PackageName,
			"displayName":   mini.DisplayName,
			"type":          mini.Type,
			"targetVersion": mini.TargetVersion,
			"icon":          mini.Icon,
			"version":       mini.Version,
			"permissions":   mini.Permissions,
		},
	}
	_, err := mongo.Db.Collection("mini_app").UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		panic(err)
	}
	return mini, id, nil

}

func (mongo *repoImpl) CreateMainApp(main MainApp) (string, error) {
	objid := primitive.NewObjectID()
	strid := objid.Hex()
	update := bson.M{
		"_id":           strid,
		"name":          main.Name,
		"platform":      main.Platform,
		"bundleId":      main.BundleId,
		"packageName":   main.PackageName,
		"googlePlayUrl": main.GooglePlayUrl,
		"appStoreUrl":   main.AppStoreUrl,
		"icon":          main.Icon,
		"version":       main.Version,
		"events":        main.Events,
	}
	_, err := mongo.Db.Collection("main_app").InsertOne(
		context.Background(),
		update)
	if err != nil {
		panic(err)
	}
	return strid, nil
}

func (mongo *repoImpl) GetMainAppID(id string) (MainApp, error) {
	result := mongo.Db.Collection("main_app").FindOne(
		context.Background(),
		bson.M{"_id": id})
	mainApp := MainApp{}
	if err := result.Decode(&mainApp); err != nil {
		panic(err)
	}
	response := MainApp{
		Id:            mainApp.Id,
		Platform:      mainApp.Platform,
		AppStoreUrl:   mainApp.AppStoreUrl,
		PackageName:   mainApp.PackageName,
		Name:          mainApp.Name,
		GooglePlayUrl: mainApp.GooglePlayUrl,
		Icon:          mainApp.Icon,
		Version:       mainApp.Version,
		BundleId:      mainApp.BundleId,
		Events:        mainApp.Events,
	}
	return response, nil
}

func (mongo *repoImpl) GetMiniAppID(id string) (MiniApp, error) {
	result := mongo.Db.Collection("mini_app").FindOne(
		context.Background(),
		bson.M{"_id": id})
	miniApp := MiniApp{}
	if err := result.Decode(&miniApp); err != nil {
		panic(err)
	}
	response := MiniApp{
		Id:            miniApp.Id,
		Platform:      miniApp.Platform,
		AppName:       miniApp.AppName,
		PackageName:   miniApp.PackageName,
		TargetVersion: miniApp.TargetVersion,
		DisplayName:   miniApp.DisplayName,
		Icon:          miniApp.Icon,
		Version:       miniApp.Version,
		BundleId:      miniApp.BundleId,
		Type:          miniApp.Type,
		Permissions:   miniApp.Permissions,
	}
	return response, nil
}

func (mongo *repoImpl) GetAllMiniApp() ([]MiniApp, error) {
	listResult := []MiniApp{}
	cursor, err := mongo.Db.Collection("mini_app").Find(
		context.Background(),
		bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		result := MiniApp{}
		cursor.Decode(&result)
		listResult = append(listResult, result)
	}
	if err := cursor.Err(); err != nil {
		panic(err)
	}
	return listResult, nil
}

func (mongo *repoImpl) GetAllMainApp() ([]MainApp, error) {
	listResult := []MainApp{}
	cursor, err := mongo.Db.Collection("main_app").Find(
		context.Background(),
		bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		result := MainApp{}
		cursor.Decode(&result)
		listResult = append(listResult, result)
	}
	if err := cursor.Err(); err != nil {
		panic(err)
	}
	return listResult, nil
}

func (mongo *repoImpl) CreateMiniApp(mini MiniApp) (string, error) {
	objid := primitive.NewObjectID()
	strid := objid.Hex()
	update := bson.M{
		"_id":           strid,
		"appName":       mini.AppName,
		"platform":      mini.Platform,
		"bundleId":      mini.BundleId,
		"packageName":   mini.PackageName,
		"displayName":   mini.DisplayName,
		"type":          mini.Type,
		"targetVersion": mini.TargetVersion,
		"icon":          mini.Icon,
		"version":       mini.Version,
		"permissions":   mini.Permissions,
	}
	_, err := mongo.Db.Collection("mini_app").InsertOne(
		context.Background(),
		update)
	if err != nil {
		panic(err)
	}
	return strid, nil
}

func (mongo *repoImpl) GetMiniofMainApp(id string) (MiniApp, error) {
	miniApp := MiniApp{}
	err := mongo.Db.Collection("mini_app").FindOne(
		context.Background(),
		bson.M{"_id": id}).Decode(&miniApp)
	if err != nil {
		panic(err)
	}
	return miniApp, nil
}

func (mongo *repoImpl) DeployMiniApp(id string) (MiniApp, error) {
	miniApp := MiniApp{}
	err := mongo.Db.Collection("mini_app").FindOne(
		context.Background(),
		bson.M{"_id": id}).Decode(&miniApp)
	if err != nil {
		panic(err)
	}
	return miniApp, nil
}

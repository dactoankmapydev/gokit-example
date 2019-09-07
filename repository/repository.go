package repository

import (
	"context"
	"fmt"
	"miniapp_backend/app"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repoImpl struct {
	Db *mongo.Database
}

type Repository interface {
	RegisterApp(main app.MainApp) (string, error)
	GetAppDetail(id string) (app.MainApp, error)
	GetMainApp() ([]app.MainApp, error)
	CreateMiniApp(mini app.MiniApp) (primitive.ObjectID, error)
	GetMiniAppDetail(id string) (app.MiniApp, error)
	GetMiniApp() ([]app.MiniApp, error)
	GetMiniofMainApp(id string) (app.MiniApp, error)
	DeployMiniApp(id string) (app.MiniApp, error)
	UpdateMiniApp(id string, mini app.MiniApp) (app.MiniApp, primitive.ObjectID, error)
	// UpdateMiniAppOfMainApp(mainId,miniId string, mini app.MiniApp) (app.MiniApp, primitive.ObjectID, error)
}

func NewRepo(db *mongo.Database) Repository {
	return &repoImpl{
		Db: db,
	}
}

func (mongo *repoImpl) UpdateMiniApp(id string, mini app.MiniApp) (app.MiniApp, primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": objID}
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
	_, err = mongo.Db.Collection("mini_app").UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		panic(err)
	}
	return mini, objID, nil

}

func (mongo *repoImpl) RegisterApp(main app.MainApp) (string, error) {
	result, err := mongo.Db.Collection("main_app").InsertOne(context.Background(), &main)
	if err != nil {
		panic(err)
	}
	newID := result.InsertedID.(primitive.ObjectID).Hex()
	return newID, nil

}

func (mongo *repoImpl) GetAppDetail(id string) (app.MainApp, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)

	}
	result := mongo.Db.Collection("main_app").FindOne(context.Background(), bson.M{"_id": objID})
	mainApp := app.MainApp{}
	fmt.Println(objID.Hex())
	if err := result.Decode(&mainApp); err != nil {

		panic(err)
	}
	response := app.MainApp{
		Id:            objID.Hex(),
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

func (mongo *repoImpl) GetMiniAppDetail(id string) (app.MiniApp, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	miniApp := app.MiniApp{}
	err = mongo.Db.Collection("mini_app").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&miniApp)
	if err != nil {
		panic(err)
	}
	return miniApp, nil
}

func (mongo *repoImpl) GetMiniApp() ([]app.MiniApp, error) {
	listResult := []app.MiniApp{}
	cursor, err := mongo.Db.Collection("mini_app").Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		result := app.MiniApp{}
		cursor.Decode(&result)
		listResult = append(listResult, result)
	}
	if err := cursor.Err(); err != nil {
		panic(err)
	}
	return listResult, nil
}

func (mongo *repoImpl) GetMainApp() ([]app.MainApp, error) {
	listResult := []app.MainApp{}
	cursor, err := mongo.Db.Collection("main_app").Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		result := app.MainApp{}
		cursor.Decode(&result)
		listResult = append(listResult, result)
	}
	if err := cursor.Err(); err != nil {
		panic(err)
	}
	return listResult, nil
}

func (mongo *repoImpl) CreateMiniApp(mini app.MiniApp) (primitive.ObjectID, error) {
	result, err := mongo.Db.Collection("mini_app").InsertOne(context.Background(), &mini)
	if err != nil {
		panic(err)
	}
	newID := result.InsertedID.(primitive.ObjectID)
	return newID, nil

}

func (mongo *repoImpl) GetMiniofMainApp(id string) (app.MiniApp, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	miniApp := app.MiniApp{}
	err = mongo.Db.Collection("mini_app").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&miniApp)
	if err != nil {
		panic(err)
	}
	return miniApp, nil
}

func (mongo *repoImpl) DeployMiniApp(id string) (app.MiniApp, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	miniApp := app.MiniApp{}
	err = mongo.Db.Collection("mini_app").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&miniApp)
	if err != nil {
		panic(err)
	}
	return miniApp, nil
}

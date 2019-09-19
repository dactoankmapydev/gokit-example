package app

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repoImpl struct {
	Db *mongo.Database
}

type Repository interface {
	CreateMainApp(main MainApp) (string, error)
	GetMainAppID(id string) (MainApp, error)
	GetAllMainApp(limit int64, cursor string) ([]MainApp, int64, string, string, error)

	CreateMiniApp(mini MiniApp) (string, error)
	GetMiniAppID(id string) (MiniApp, error)
	GetAllMiniApp(limit int64) ([]MiniApp, int64, error)
	GetMiniofMainApp(id string) (MiniApp, int64, error)
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

func (mongo *repoImpl) GetAllMiniApp(limit int64) ([]MiniApp, int64, error) {
	listResult := []MiniApp{}
	options := options.Find()
	lm := options.SetLimit(limit)
	cursor, err := mongo.Db.Collection("mini_app").Find(
		context.Background(),
		bson.M{},
		lm)
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
	total, err := mongo.Db.Collection("mini_app").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	return listResult, total, nil
}

func (mongo *repoImpl) GetAllMainApp(limit int64, cursor string) ([]MainApp, int64, string, string, error) {
	listResult := []MainApp{}
	result := MainApp{}

	if cursor == "" {
		lastRecords, err := mongo.Db.Collection("main_app").Find(
			context.Background(),
			bson.M{},
			&options.FindOptions{Limit: &limit},
		)
		if err != nil {
			panic(err)
		}
		defer lastRecords.Close(context.Background())
		for lastRecords.Next(context.Background()) {
			if err := lastRecords.Decode(&result); err != nil {
				panic(err)
			}
			listResult = append(listResult, result)
		}
		if err := lastRecords.Err(); err != nil {
			panic(err)
		}
	}

	if cursor != "" {
		if (cursor[strings.LastIndex(cursor, ".")+1:]) == "next" {
			id := strings.Split(cursor, "."+cursor[strings.LastIndex(cursor, ".")+1:])
			nextPage, err := mongo.Db.Collection("main_app").Find(
				context.Background(),
				bson.M{
					"_id": bson.M{"$gt": id[0]},
				},
				&options.FindOptions{Limit: &limit},
			)
			if err != nil {
				panic(err)
			}
			defer nextPage.Close(context.Background())
			for nextPage.Next(context.Background()) {
				if err := nextPage.Decode(&result); err != nil {
					panic(err)
				}
				listResult = append(listResult, result)

			}
			if err := nextPage.Err(); err != nil {
				panic(err)
			}
		}
		if (cursor[strings.LastIndex(cursor, ".")+1:]) == "prev" {
			id := strings.Split(cursor, "."+cursor[strings.LastIndex(cursor, ".")+1:])
			prevPage, err := mongo.Db.Collection("main_app").Find(
				context.Background(),
				bson.M{
					"_id": bson.M{"$lt": id[0]},
				},
				&options.FindOptions{
					Sort:  map[string]int{"_id": -1},
					Limit: &limit,
				},
			)
			if err != nil {
				panic(err)
			}
			defer prevPage.Close(context.Background())
			for prevPage.Next(context.Background()) {
				if err := prevPage.Decode(&result); err != nil {
					panic(err)
				}
				listResult = append(listResult, result)

			}
			if err := prevPage.Err(); err != nil {
				panic(err)
			}
		}

	}
	lastIDN := result.Id + ".next"
	lastIDP := result.Id + ".prev"

	// Total documents
	total, err := mongo.Db.Collection("main_app").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	return listResult, total, lastIDN, lastIDP, nil
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

func (mongo *repoImpl) GetMiniofMainApp(id string) (MiniApp, int64, error) {
	miniApp := MiniApp{}
	err := mongo.Db.Collection("mini_app").FindOne(
		context.Background(),
		bson.M{"_id": id}).Decode(&miniApp)
	if err != nil {
		panic(err)
	}
	total, err := mongo.Db.Collection("mini_app").CountDocuments(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	return miniApp, total, nil
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

package appsvc

import (
	"context"
	"fmt"
	"log"
	"miniapp_backend/app"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Service interface {
	RegisterApp(context.Context, RegisterAppRequest) (*RegisterAppResponse, error)
	GetMainApp(context.Context, GetMainAppRequest) (*GetMainAppResponse, error)
	GetAppDetail(context.Context, GetAppDetailRequest) (*GetAppDetailResponse, error)
	GetMiniofMainApp(context.Context, GetMiniAppOfAppRequest) (*GetMiniAppOfAppResponse, error)
	UpdateMiniAppOfMainApp(context.Context, UpdateMiniAppOfMainAppRequest) (*UpdateMiniAppOfMainAppResponse, error)
	GetMiniApp(context.Context, GetMiniAppRequest) (*GetMiniAppResponse, error)
	GetMiniAppDetail(context.Context, GetMiniAppDetailRequest) (*GetMiniAppDetailResponse, error)
	CreateMiniApp(context.Context, CreateMiniAppRequest) (*CreateMiniAppResponse, error)
	UpdateMiniApp(context.Context, UpdateMiniAppRequest) (*UpdateMiniAppResponse, error)
	DeployMiniApp(context.Context, DeployMiniAppRequest) (*DeployMiniAppResponse, error)
}

type service struct {
}

type MongoFields struct {
	FieldStr  string `json:"Field Str"`
	FieldInt  int    `json:"Field Int"`
	FieldBool bool   `json:"Field Bool"`
}

func New() Service {

	return &service{}
}

func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("mainapp").Collection("register")
	fmt.Println("Collection type:", reflect.TypeOf(collection), "\n")
}

func (s *service) GetMainApp(ctx context.Context, r GetMainAppRequest) (*GetMainAppResponse, error) {
	var sliceApp []app.MainApp
	sliceApp = append(sliceApp, app.MainApp{
		Id:          "1",
		Platform:    "ios",
		PackageName: "appsvc",
		Version:     "v1",
		Events:      []string{},
	})
	return &GetMainAppResponse{
		Total:  r.Limit,
		Apps:   sliceApp,
		Cursor: Cursor{},
	}, nil
}

func (s *service) RegisterApp(Ctx context.Context, r RegisterAppRequest) (*RegisterAppResponse, error) {

	oneDoc := MongoFields{
		FieldStr:  "Some Value",
		FieldInt:  12345,
		FieldBool: true,
	}
	fmt.Println("oneDoc TYPE:", reflect.TypeOf(oneDoc), "\n")

	// InsertOne() method Returns mongo.InsertOneResult
	result, insertErr := collection.InsertOne(context.Background(), oneDoc)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() API result:", result)
	}
	return &RegisterAppResponse{
		App: app.MainApp{
			BundleId:      r.BundleId,
			Platform:      r.Platform,
			PackageName:   r.PackageName,
			Name:          r.Name,
			GooglePlayUrl: r.GooglePlayUrl,
			AppStoreUrl:   r.AppStoreUrl,
			Icon:          r.Icon,
			Version:       r.Version,
		},
		Ica: "1",
	}, nil
}

func (s *service) GetAppDetail(Ctx context.Context, r GetAppDetailRequest) (*GetAppDetailResponse, error) {
	return &GetAppDetailResponse{
		App: app.MainApp{Id: r.Id},
		Ica: "5",
	}, nil
}

func (s *service) GetMiniofMainApp(Ctx context.Context, r GetMiniAppOfAppRequest) (*GetMiniAppOfAppResponse, error) {
	var sliceMiniApp []app.MiniApp
	sliceMiniApp = append(sliceMiniApp, app.MiniApp{
		Id:          "2",
		Platform:    "android",
		PackageName: "appsvc",
		Version:     "v1",
		Permissions: []string{},
	})
	return &GetMiniAppOfAppResponse{
		Total:  r.Limit,
		Apps:   sliceMiniApp,
		Cursor: Cursor{},
	}, nil
}

func (s *service) GetMiniApp(ctx context.Context, r GetMiniAppRequest) (*GetMiniAppResponse, error) {
	var sliceApp []app.MiniApp
	sliceApp = append(sliceApp, app.MiniApp{
		Id:          "5",
		Platform:    "android",
		PackageName: "appsvc",
		Version:     "v3",
		Permissions: []string{},
	})
	return &GetMiniAppResponse{
		Total:  r.Limit,
		Apps:   sliceApp,
		Cursor: Cursor{},
	}, nil
}

func (s *service) GetMiniAppDetail(Ctx context.Context, r GetMiniAppDetailRequest) (*GetMiniAppDetailResponse, error) {
	return &GetMiniAppDetailResponse{
		App: app.MiniApp{
			Id:          r.Id,
			Permissions: []string{},
		},
		History: []string{},
	}, nil
}

func (s *service) CreateMiniApp(Ctx context.Context, r CreateMiniAppRequest) (*CreateMiniAppResponse, error) {
	v := reflect.ValueOf(r)
	typeOfs := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfs.Field(i).Name
		if valueField := v.Field(i).Interface(); valueField != "" {
			// Todo save to db
			fmt.Println(field, "exist")
		}
	}
	return &CreateMiniAppResponse{
		App: app.MiniApp{
			Platform:      r.Platform,
			BundleId:      r.BundleId,
			PackageName:   r.PackageName,
			DisplayName:   r.DisplayName,
			AppName:       r.AppName,
			Type:          r.Type,
			TargetVersion: r.TargetVersion,
			Icon:          r.Icon,
			Version:       r.Version,
			Permissions:   r.Permissions,
		},
	}, nil
}

func (s *service) UpdateMiniAppOfMainApp(Ctx context.Context, r UpdateMiniAppOfMainAppRequest) (*UpdateMiniAppOfMainAppResponse, error) {
	v := reflect.ValueOf(r)
	typeOfs := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfs.Field(i).Name
		if valueField := v.Field(i).Interface(); valueField != "" {
			// Todo save to db
			println(field, "exist")
		}
	}
	return &UpdateMiniAppOfMainAppResponse{
		App: app.MiniApp{
			Status:      r.Status,
			Permissions: []string{},
		},
	}, nil
}

func (s *service) UpdateMiniApp(Ctx context.Context, r UpdateMiniAppRequest) (*UpdateMiniAppResponse, error) {
	v := reflect.ValueOf(r)
	typeOfs := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfs.Field(i).Name
		if valueField := v.Field(i).Interface(); valueField != "" {
			// Todo save to db
			println(field, "exist")
		}
	}
	return &UpdateMiniAppResponse{
		App: app.MiniApp{
			Platform:      r.Platform,
			BundleId:      r.BundleId,
			PackageName:   r.PackageName,
			DisplayName:   r.DisplayName,
			AppName:       r.AppName,
			Type:          r.Type,
			TargetVersion: r.TargetVersion,
			Icon:          r.Icon,
			Version:       r.Version,
			Permissions:   r.Permissions,
			Id:            r.Id,
		},
	}, nil
}

func (s *service) DeployMiniApp(Ctx context.Context, r DeployMiniAppRequest) (*DeployMiniAppResponse, error) {
	v := reflect.ValueOf(r)
	typeOfs := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typeOfs.Field(i).Name
		if valueField := v.Field(i).Interface(); valueField != "" {
			// Todo save to db
			println(field, "exist")
		}
	}

	return &DeployMiniAppResponse{
		App: app.MiniApp{
			Id:       r.Id,
			Platform: r.Platform,
			Version:  r.Version,
		},
	}, nil
}

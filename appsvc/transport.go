package appsvc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
)

var (
	ErrParameter = errors.New("The input parameter must be an integer")
	ErrUploadFile      = errors.New("Error Retrieving the File")
	ErrCreateTempFile = errors.New("Cannot create temporary file")
	ErrWriteTempFile  = errors.New("Failed to write to temporary file")
)

// NewHandler returns new http.Handler that routes http request to service
func NewHandler(s Service, router *httprouter.Router) http.Handler {

	router.Handler(http.MethodGet, "/api/v1/main-apps", kithttp.NewServer(
		makeGetMainAppEndpoint(s),
		decodeGetMainAppRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/main-apps/", kithttp.NewServer(
		makeRegisterAppEndpoint(s),
		decodeRegisterAppRequests,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/main-apps/:id", kithttp.NewServer(
		makeGetAppDetailEndpoint(s),
		decodeGetAppDetailRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/main-apps/:id/mini-apps", kithttp.NewServer(
		makeGetMiniofMainAppEndpoint(s),
		decodeGetMiniofMainAppRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/mini-apps", kithttp.NewServer(
		makeGetMininAppEndpoint(s),
		decodeGetMiniAppRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodGet, "/api/v1/mini-apps/:id", kithttp.NewServer(
		makeGetMiniAppDetailEndpoint(s),
		decodeGetMiniAppDetailRequest,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/mini-apps/", kithttp.NewServer(
		makeCreateMiniAppEndpoint(s),
		decodeCreateMiniAppRequests,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/main-apps/:mainAppId/mini-apps/:miniAppId", kithttp.NewServer(
		makeUpdateMiniAppOfMainApp(s),
		decodeUpdateMiniAppOfMainApp,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/mini-apps/:id", kithttp.NewServer(
		makeUpdateMiniApp(s),
		decodeUpdateMiniApp,
		kithttp.EncodeJSONResponse,
	))

	router.Handler(http.MethodPost, "/api/v1/mini-apps/:id/deploy", kithttp.NewServer(
		makeDeployMiniApp(s),
		decodeDeployMiniApp,
		kithttp.EncodeJSONResponse,
	))

	return router
}

func decodeGetMainAppRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	req := GetMainAppRequest{Cursor: q.Get("cursor")}
	if qLimit := q.Get("limit"); qLimit != "" {
		intLimit, err := strconv.Atoi(qLimit)
		if err != nil {
			return nil, ErrParameter
		}
		req.Limit = intLimit
	}
	return req, nil
}

func decodeRegisterAppRequests(ctx context.Context, r *http.Request) (interface{}, error) {
	var request RegisterAppRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeGetAppDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	pid := httprouter.ParamsFromContext(ctx).ByName("id")
	return GetAppDetailRequest{Id: pid}, nil
}

func decodeGetMiniofMainAppRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	req := GetMiniAppOfAppRequest{Cursor: q.Get("cursor")}
	if qLimit := q.Get("limit"); qLimit != "" {
		intLimit, err := strconv.Atoi(qLimit)
		if err != nil {
			return nil, ErrParameter
		}
		req.Limit = intLimit
	}
	return req, nil
}

func decodeGetMiniAppRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	req := GetMiniAppRequest{Cursor: q.Get("cursor")}
	if qLimit := q.Get("limit"); qLimit != "" {
		intLimit, err := strconv.Atoi(qLimit)
		if err != nil {
			return nil, ErrParameter
		}
		req.Limit = intLimit
	}
	return req, nil
}

func decodeGetMiniAppDetailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	pid := httprouter.ParamsFromContext(ctx).ByName("id")
	return GetMiniAppDetailRequest{Id: pid}, nil
}

func decodeCreateMiniAppRequests(ctx context.Context, r *http.Request) (interface{}, error) {

	file, handler, err := r.FormFile("bundle")
	if err != nil {
		return nil, ErrUploadFile
	}
	defer file.Close()

	// Create a temp file
	fileType := (handler.Filename)[strings.LastIndex(handler.Filename, ".")+1:]
	tempFile, err := ioutil.TempFile(os.TempDir(), "*."+fileType)
	if err != nil {
		return nil, ErrCreateTempFile
	}
	
	// Cleaning up by removing the file
	defer os.Remove(tempFile.Name())
	fmt.Println("Create a temp file: ", tempFile.Name())

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err, nil
	}

	// Write to the file
	if _, err = tempFile.Write(fileBytes); err != nil {
		return nil, ErrWriteTempFile
	}

	// Close the file
	if err := tempFile.Close(); err != nil {
		panic(err)
	}

	return CreateMiniAppRequest{
		Platform:      r.FormValue("platform"),
		BundleId:      r.FormValue("bundleId"),
		PackageName:   r.FormValue("packageName"),
		DisplayName:   r.FormValue("displayName"),
		AppName:       r.FormValue("appName"),
		Type:          r.FormValue("type"),
		TargetVersion: r.FormValue("targetVersion"),
		Icon:          r.FormValue("icon"),
		Version:       r.FormValue("version"),
		// Permissions:   r.FormValue("permissions"),
	}, nil
}

func decodeUpdateMiniAppOfMainApp(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateMiniAppOfMainAppRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateMiniApp(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateMiniAppRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeployMiniApp(ctx context.Context, r *http.Request) (interface{}, error) {
	pid := httprouter.ParamsFromContext(ctx).ByName("id")
	file, handler, err := r.FormFile("bundle")
	if err != nil {
		return nil, ErrUploadFile
	}
	defer file.Close()

	fileType := (handler.Filename)[strings.LastIndex(handler.Filename, ".")+1:]
	tempFile, err := ioutil.TempFile(os.TempDir(), "*."+fileType)
	if err != nil {
		return nil, ErrCreateTempFile
	}

	defer os.Remove(tempFile.Name())

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err, nil
	}

	if _, err = tempFile.Write(fileBytes); err != nil {
		return nil, ErrWriteTempFile
	}

	if err := tempFile.Close(); err != nil {
		panic(err)
	}

	return DeployMiniAppRequest{
		Id:       pid,
		Platform: r.FormValue("platform"),
		Version:  r.FormValue("version"),
	}, nil
}


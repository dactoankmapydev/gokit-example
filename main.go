package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"

	"miniapp_backend/app"
	"miniapp_backend/appsvc"
	"miniapp_backend/driver"
)

func main() {
	router := httprouter.New()
	mongoConn := driver.ConnectMongoDB()

	{
		appRepo := app.NewRepo(mongoConn.Client.Database("app"))
		s := appsvc.New(appRepo)
		appsvc.NewHandler(s, router)
	}

	// router.PanicHandler = func(w http.ResponseWriter, r *http.Request, panic interface{}) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("500: Internal Server Error"))
	// 	if p, ok := panic.(string); ok {
	// 		w.Write([]byte("\nPanic: " + p))
	// 	}
	// }

	errs := make(chan error, 2)
	go func() {
		errs <- http.ListenAndServe(":9090", router)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		<-c
		// TODO: received exit signal: do graceful shutdown
		errs <- nil
	}()

	<-errs
	// TODO: handle error
}

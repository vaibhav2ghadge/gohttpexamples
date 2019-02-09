package main

import (
	"fmt"
	logger "log"
	"net/http"
	"os"
	"time"

	dbrepo "github.com/gohttpexamples/restaurant/dao/dbrepository"
	mongoutils "github.com/gohttpexamples/restaurant/dao/utils"
	handlerlib "github.com/gohttpexamples/restaurant/delivery/restapplication/packages/httphandlers"
	"github.com/gohttpexamples/restaurant/delivery/restapplication/usercrudhandler"
	"github.com/gorilla/mux"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)
	//clone mongo session in resturant structure
	restHandler := usercrudhandler.NewRestCrudHandler(repoaccess)
	fmt.Println(restHandler.Mongo1)
	pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	h.Handle("/ping/", pingHandler)
	//h.Handle("/restaurantservice/restaurant/{typeOfFood}", restHandler)
	//h.Handle("/restaurantservice/restaurant/?typeOfFood={food}", restHandler)
	h.Handle("/restaurantservice/restaurant/{id}", restHandler)
	h.Handle("/restaurantservice/restaurant/", restHandler)
	logger.Println("Resource Setup Done.")
	logger.Fatal(http.ListenAndServe(":8080", h))
}

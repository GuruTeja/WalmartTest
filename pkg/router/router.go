package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"walmartTest/pkg/controller"
)

func InitRouter() {
	log.Println("Initiating Routers")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc(
		"/alerts",
		controller.SaveAlert,
	).Methods("POST")

	myRouter.HandleFunc(
		"/alerts",
		controller.GetAlerts,
	).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

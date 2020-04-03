package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//Health check controller
	myRouter.HandleFunc("/ping", Ping).Methods("GET")

	//Save User details
	myRouter.HandleFunc("/saveUser", SaveUser).Methods("POST")

	//Autheniticate User
	myRouter.HandleFunc("/authenticateUser", AuthenticateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

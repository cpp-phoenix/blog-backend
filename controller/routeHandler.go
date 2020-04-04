package controller

import (
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

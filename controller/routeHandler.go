package controller

import (
	"blog_backend/properties"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	port := strconv.Itoa(properties.CUSTOM_PORT)
	fmt.Println("Port No.: " + port)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

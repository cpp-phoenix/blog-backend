package controller

import (
	"fmt"
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

	//Email Service
	myRouter.HandleFunc("/triggerEmail", TriggerEmail).Methods("POST")

	//Save Post
	myRouter.HandleFunc("/savePost", SavePost).Methods("POST")

	fmt.Println("Port No.: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), myRouter))
	//log.Fatal(http.ListenAndServe(":5000", myRouter))
}

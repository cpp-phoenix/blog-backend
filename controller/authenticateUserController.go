package controller

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"blog_backend/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func AuthenticateUser(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(res, msg, http.StatusUnsupportedMediaType)
		return
	}

	//Check if the size of body is not greater than allowed ranged
	req.Body = http.MaxBytesReader(res, req.Body, properties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var user dto.UserDetails

	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}
	status := services.LogIn(user)
	if status {
		res.WriteHeader(http.StatusOK)
		fmt.Fprintf(res, "User Authentication Successfully!!")
	} else {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		fmt.Fprintf(res, "User Not Present. Please Create Account")
	}
}

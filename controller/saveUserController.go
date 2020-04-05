package controller

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"blog_backend/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func SaveUser(w http.ResponseWriter, req *http.Request) {
	//Check for the content-type
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	//Check if the size of body is not greater than allowed ranged
	req.Body = http.MaxBytesReader(w, req.Body, properties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var user dto.UserDetails

	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}
	status := services.SignUp(user)
	if status == "success" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User Updated Successfully!")
	}
}

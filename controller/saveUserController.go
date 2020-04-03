package controller

import (
	"blog_backend/applicationProperties"
	"blog_backend/dto"
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
	req.Body = http.MaxBytesReader(w, req.Body, applicationProperties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var user dto.UserDetails

	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Person: %+v", user)
}

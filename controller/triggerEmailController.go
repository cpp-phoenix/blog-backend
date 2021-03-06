package controller

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"blog_backend/services"
	"encoding/json"
	"net/http"
)

func TriggerEmail(res http.ResponseWriter, req *http.Request) {
	//Check for the content-type
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(res, msg, http.StatusNotFound)
		return
	}

	//Check if the size of body is not greater than allowed ranged
	req.Body = http.MaxBytesReader(res, req.Body, properties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)

	var user dto.UserDetails

	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}

	status := services.TriggerEmail(user)

	//Response Object
	var response dto.UserResponse
	response.Status = status
	jsonResponse, _ := json.Marshal(response)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(jsonResponse)
}

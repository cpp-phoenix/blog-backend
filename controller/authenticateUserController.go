package controller

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"blog_backend/services"
	"encoding/json"
	"net/http"
)

func AuthenticateUser(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(res, msg, http.StatusNotFound)
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
		var response dto.UserResponse
		response.Message = "User Authentication Successfully!!"
		jsonResponse, _ := json.Marshal(response)
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Write(jsonResponse)
	} else {
		var response dto.UserResponse
		response.Message = "User Not Present. Please Create Account"
		jsonResponse, _ := json.Marshal(response)
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Write(jsonResponse)
	}
}

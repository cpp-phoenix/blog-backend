package controller

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"blog_backend/services"
	"encoding/json"
	"net/http"
)

func SavePost(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(res, msg, http.StatusNotFound)
		return
	}

	//Check if the size of body is not greater than allowed ranged
	req.Body = http.MaxBytesReader(res, req.Body, properties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var post dto.PostDetails

	err := dec.Decode(&post)
	if err != nil {
		panic(err)
	}
	status := services.SavePost(post)
	var response dto.UserResponse
	response.Status = status
	jsonResponse, _ := json.Marshal(response)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(jsonResponse)
}

func SaveLikes(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(res, msg, http.StatusNotFound)
		return
	}

	//Check if the size of body is not greater than allowed ranged
	req.Body = http.MaxBytesReader(res, req.Body, properties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var post dto.PostUpdation

	err := dec.Decode(&post)
	if err != nil {
		panic(err)
	}
	status := services.SaveLikes(post)
	var response dto.UserResponse
	response.Status = status
	jsonResponse, _ := json.Marshal(response)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(jsonResponse)
}

func SaveBookmark(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(res, msg, http.StatusNotFound)
		return
	}

	//Check if the size of body is not greater than allowed ranged
	req.Body = http.MaxBytesReader(res, req.Body, properties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var post dto.PostUpdation

	err := dec.Decode(&post)
	if err != nil {
		panic(err)
	}
	status := services.SaveBookmarks(post)
	var response dto.UserResponse
	response.Status = status
	jsonResponse, _ := json.Marshal(response)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(jsonResponse)
}

func UpdateAvatar(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(res, msg, http.StatusNotFound)
		return
	}

	//Check if the size of body is not greater than allowed ranged
	req.Body = http.MaxBytesReader(res, req.Body, properties.MAX_SIZE_OF_INPUT_REQUEST_PERMITTED)

	dec := json.NewDecoder(req.Body)

	var post dto.PostDetails

	err := dec.Decode(&post)
	if err != nil {
		panic(err)
	}
	status := services.UpdateAvatar(post)
	var response dto.UserResponse
	response.Status = status
	jsonResponse, _ := json.Marshal(response)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(jsonResponse)
}

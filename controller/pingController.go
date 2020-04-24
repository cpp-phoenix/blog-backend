package controller

import (
	"blog_backend/dto"
	"encoding/json"
	"net/http"
)

func Ping(res http.ResponseWriter, r *http.Request) {
	var response dto.UserResponse
	response.Message = "The service is running fine"
	jsonResponse, _ := json.Marshal(response)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(jsonResponse)
}

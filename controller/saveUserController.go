package controller

import (
	"blog_backend/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func SaveUser(w http.ResponseWriter, req *http.Request) {
	//Check for the content-type
	if req.Header.Get("Content-Type") != "application/json" {
		msg := "Content type is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	//Check if the size of body is not greater than allowed ranged

	maxSizeInputRequestPermitted, _ := os.LookupEnv("MAX_SIZE_OF_INPUT_REQUEST_PERMITTED")
	inputRequestVar, _ := strconv.ParseInt(maxSizeInputRequestPermitted, 10, 64)
	req.Body = http.MaxBytesReader(w, req.Body, inputRequestVar)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var user dto.UserDetails

	err := dec.Decode(&user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Person: %+v", user)
}

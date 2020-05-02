package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	b64 "encoding/base64"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func validatePassword(password string, savedPassword string) bool {
	sEnc := b64.StdEncoding.EncodeToString([]byte(password))
	if savedPassword == sEnc {
		return true
	}
	return false
}

func ValidateUser(userName string, password string) int {
	searchRequest := searchRequestBuilderForUserName(userName)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.UserName != userName {
		return 3001
	} else if !validatePassword(password, userDetails.Password) {
		return 3007
	}
	return 3000
}

func ValidateEmail(email string, password string) int {
	searchRequest := searchRequestBuilderForEmailAddress(email)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.Email != email {
		return 3003
	} else if !validatePassword(password, userDetails.Password) {
		return 3007
	}
	return 3000
}

func LogIn(user dto.UserDetails) int {
	authentication := 3016
	user.Email = strings.ToLower(user.Email)
	if user.UserName != "" {
		authentication = ValidateUser(user.UserName, user.Password)
	}
	if user.Email != "" {
		authentication = ValidateEmail(user.Email, user.Password)
	}
	return authentication
}

func FetchUser(user dto.UserDetails) dto.UserResponse {
	authentication := 3016
	var response dto.UserResponse
	if user.UserName != "" {
		dbDetails := FetchUserFromDB(user.UserName)
		if dbDetails.UserName == "" {
			response.Status = 3001
			return response
		}
		response.Status = 3000
		dbDetails.Password = ""
		response.Data = dbDetails
		return response
	}
	response.Status = authentication
	return response
}

func FetchUserFromDB(userName string) dto.UserDetails {
	searchRequest := searchRequestBuilderForUserName(userName)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	return userDetails
}

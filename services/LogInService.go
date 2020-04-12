package services

import (
	"blog_backend/dto"
	"blog_backend/properties"

	"go.mongodb.org/mongo-driver/bson"
)

func validateUser(userName string) bool {
	searchRequest := searchRequestBuilderForUserName(userName)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.UserName == userName {
		return true
	}
	return false
}

func validateEmail(email string) bool {
	searchRequest := searchRequestBuilderForEmailAddress(email)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.Email == email {
		return true
	}
	return false
}

func LogIn(user dto.UserDetails) bool {
	authentication := false
	if user.UserName != "" {
		authentication = validateUser(user.UserName)
	}
	if user.Email != "" {
		authentication = validateEmail(user.Email)
	}
	return authentication
}

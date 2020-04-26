package services

import (
	"blog_backend/dto"
	"blog_backend/properties"

	"go.mongodb.org/mongo-driver/bson"
)

func validateUser(userName string) int {
	searchRequest := searchRequestBuilderForUserName(userName)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.UserName == userName {
		return 3000
	}
	return 3001
}

func validateEmail(email string) int {
	searchRequest := searchRequestBuilderForEmailAddress(email)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.Email == email {
		return 3000
	}
	return 3003
}

func LogIn(user dto.UserDetails) int {
	authentication := 3000
	if user.UserName != "" {
		authentication = validateUser(user.UserName)
	}
	if user.Email != "" {
		authentication = validateEmail(user.Email)
	}
	return authentication
}

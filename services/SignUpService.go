package services

import (
	"blog_backend/dto"
	"blog_backend/properties"

	"go.mongodb.org/mongo-driver/bson"
)

func saveToDb(user dto.UserDetails) bool {
	return saveSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, user)
}

func CheckUserName(userName string) bool {
	searchRequest := searchRequestBuilderForUserName(userName)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.UserName != userName {
		return true
	}
	return false
}

func CheckEmailAddress(email string) bool {
	searchRequest := searchRequestBuilderForEmailAddress(email)
	var userDetails dto.UserDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &userDetails)
	if userDetails.Email != email {
		return true
	}
	return false
}

func SignUp(user dto.UserDetails) int {
	isUnique := CheckUserName(user.UserName)
	if !isUnique {
		return 3002
	}
	isUnique = CheckEmailAddress(user.Email)
	if !isUnique {
		return 3004
	}
	isSuccessfull := saveToDb(user)
	if isSuccessfull {
		return 3000
	} else {
		return 3005
	}
}

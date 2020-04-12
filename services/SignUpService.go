package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"fmt"
)

func saveToDb(user dto.UserDetails) bool {
	//return saveSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, user)
	return true
}

func CheckUserName(userName string) bool {
	searchRequest := searchRequestBuilderForUserName(userName)
	fmt.Println(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest))
	if len(searchRequest.SearchCriteria) == 0 {
		return false
	}
	return false
}

func CheckEmailAddress(email string) bool {
	searchRequest := searchRequestBuilderForUserName(email)
	if len(searchRequest.SearchCriteria) == 0 {
		return false
	}
	return false
}

func SignUp(user dto.UserDetails) bool {
	isUnique := CheckUserName(user.UserName)
	if !isUnique {
		return false
	}
	isUnique = CheckEmailAddress(user.Email)
	if !isUnique {
		return false
	}
	isSuccessfull := saveToDb(user)
	return isSuccessfull
}

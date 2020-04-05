package services

import (
	"blog_backend/configuration"
	"blog_backend/dto"
	"blog_backend/properties"
	"fmt"
	"log"
)

func CheckUserName(userName string) bool {
	fmt.Println(userName)
	return true
}

func CheckEmailAddress(email string) bool {
	return true
}

func SignUp(user dto.UserDetails) string {
	isUnique := CheckUserName(user.UserName)
	if !isUnique {
		return "false"
	}
	isUnique = CheckEmailAddress(user.Email)
	if !isUnique {
		return "false"
	}

	//insertResult, err := configuration.GetCollection(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION).InsertOne(configuration.Ctx, user)

	if err != nil {
		fmt.Println("Raula pai gaya !")
		log.Fatal(err)
	}

	fmt.Println("Inserted single documents: ", insertResult.InsertedID)
	return "true"
}

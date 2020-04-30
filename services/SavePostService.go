package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func verifyRandomisation(randomNumber int) bool {
	searchRequest := searchRequestBuilderForPostId(randomNumber)
	var post dto.PostDetails
	bsonBytes, _ := bson.Marshal(executeSearch(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, searchRequest))
	bson.Unmarshal(bsonBytes, &post)
	if post.PostId == randomNumber {
		return false
	}
	return true
}

func randomNumberGenerator() int {
	var randomNo int
	var status bool
	for {
		randomNo = GenerateRandomNumber(1, 100000)
		status = verifyRandomisation(randomNo)
		if status {
			break
		}
	}
	return randomNo
}

func savePostToDb(post dto.PostDetails) bool {
	post.PostId = randomNumberGenerator()
	post.CreatedTimeStamp = time.Now().Unix()
	post.UpdatedTimeStamp = time.Now().Unix()
	return saveSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, post)
}

func SavePost(post dto.PostDetails) int {
	if post.Title == "" || post.Description == "" {
		return 3014
	}
	status := savePostToDb(post)
	if status {
		return 3015
	}
	return 3005
}

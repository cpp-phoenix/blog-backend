package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"encoding/json"
	"log"
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
	_ = updatePostsEntryInUserDb(post.UserName, post.PostId)
	loc, _ := time.LoadLocation("UTC")
	post.CreatedTimeStamp = time.Now().In(loc).Unix()
	post.UpdatedTimeStamp = time.Now().In(loc).Unix()
	return saveSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, post)
}

func createBsonObjectForUserDataUpdation(postId interface{}, key string) bson.M {
	query := make(map[string]interface{})

	internalquery := make(map[string]interface{})
	internalquery[key] = postId

	query["$push"] = internalquery
	var bsonMap bson.M
	foo_marshalled, _ := json.Marshal(query)
	err := json.Unmarshal([]byte(string(foo_marshalled)), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	}
	return bsonMap
}

func createBsonObjectForUserDataRewrite(postId interface{}, key string) bson.M {
	query := make(map[string]interface{})

	internalquery := make(map[string]interface{})
	internalquery[key] = postId

	query["$set"] = internalquery
	var bsonMap bson.M
	foo_marshalled, _ := json.Marshal(query)
	err := json.Unmarshal([]byte(string(foo_marshalled)), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	}
	return bsonMap
}

func updatePostsEntryInUserDb(username string, postid int) int {

	searchRequest := searchRequestBuilderForUserName(username)
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest, createBsonObjectForUserDataUpdation(postid, "createdposts"))
	return status
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

func SaveLikes(updationRequest dto.PostUpdation) int {
	searchRequestForUserName := searchRequestBuilderForUserName(updationRequest.UserName)
	searchRequestForPostId := searchRequestBuilderForPostId(updationRequest.PostId)
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequestForUserName, createBsonObjectForUserDataUpdation(updationRequest.PostId, "likedposts"))
	status = updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, searchRequestForPostId, createBsonObjectForUserDataUpdation(updationRequest.UserName, "likedby"))
	return status
}

func SaveBookmarks(updationRequest dto.PostUpdation) int {
	searchRequestForUserName := searchRequestBuilderForUserName(updationRequest.UserName)
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequestForUserName, createBsonObjectForUserDataUpdation(updationRequest.PostId, "savedposts"))
	return status
}

func UpdateAvatar(request dto.PostDetails) int {
	searchRequestForUserName := searchRequestBuilderForUserName(request.UserName)
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequestForUserName, createBsonObjectForUserDataRewrite(request.Avatar, "avatar"))
	status = updateManyDocuments(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, searchRequestForUserName, createBsonObjectForUserDataRewrite(request.Avatar, "avatar"))
	return status
}

package services

import (
	"blog_backend/dto"
	"blog_backend/properties"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func createBsonObjectForUserDataDeletion(postId interface{}, key string) bson.M {
	query := make(map[string]interface{})

	internalquery := make(map[string]interface{})
	internalquery[key] = postId

	query["$pull"] = internalquery
	var bsonMap bson.M
	foo_marshalled, _ := json.Marshal(query)
	err := json.Unmarshal([]byte(string(foo_marshalled)), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	}
	return bsonMap
}

func DeleteLike(post dto.PostUpdation, postkey string, userkey string) int {
	searchRequestForUserName := searchRequestBuilderForUserName(post.UserName)
	searchRequestForPostId := searchRequestBuilderForPostId(post.PostId)
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, searchRequestForPostId, createBsonObjectForUserDataDeletion(post.UserName, postkey))
	status = updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequestForUserName, createBsonObjectForUserDataDeletion(post.PostId, userkey))
	return status
}

func DeleteSave(post dto.PostUpdation, userkey string) int {
	searchRequestForUserName := searchRequestBuilderForUserName(post.UserName)
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequestForUserName, createBsonObjectForUserDataDeletion(post.PostId, userkey))
	return status
}

func UnFollow(request dto.FnF) int {
	searchRequestForUserName := searchRequestBuilderForUserName(request.UserName)
	searchRequestForFollowingUserName := searchRequestBuilderForUserName(request.FollowingUserName)
	status := updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequestForUserName, createBsonObjectForUserDataDeletion(request.FollowingUserName, "following"))
	status = updateSingleDocument(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequestForFollowingUserName, createBsonObjectForUserDataDeletion(request.UserName, "followers"))
	return status
}

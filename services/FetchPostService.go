package services

import (
	"blog_backend/dto"
	"blog_backend/properties"

	"go.mongodb.org/mongo-driver/bson"
)

func FetchPost(postSearch dto.PostSearch) dto.UserResponse {
	var posts []dto.PostDetails
	var searchRequest bson.M
	bsonResponse := executeSearchFetchMultiple(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, searchRequest, "UpdatedTimeStamp", int64(postSearch.Page), int64(postSearch.Size))
	var response dto.UserResponse
	response.Status = 3000
	response.Data = bsonResponse
	return response
}

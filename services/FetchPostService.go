package services

import (
	"blog_backend/dto"
	"blog_backend/properties"

	"go.mongodb.org/mongo-driver/bson"
)

func FetchPost(postSearch dto.PostSearch) dto.UserResponse {
	var searchRequest bson.M
	bsonResponse := executeSearchFetchMultiple(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, searchRequest, "UpdatedTimeStamp", int64(postSearch.Page), int64(postSearch.Size))
	var response dto.UserResponse
	response.Status = 3000
	response.Data = bsonResponse
	return response
}

func FetchPostByPostIds(postSearchRequest dto.PostSearch) dto.UserResponse {
	var searchRequest = searchRequestBuilderForPostIds(postSearchRequest.PostIds)
	bsonSearchRequest := createDTOToBsonRequest(searchRequest)
	bsonResponse := executeSearchFetchMultiple(properties.BLOG_BACKEND_DATABASE, properties.POST_DETAILS_COLLECTION, bsonSearchRequest, "UpdatedTimeStamp", int64(postSearchRequest.Page), int64(postSearchRequest.Size))
	var response dto.UserResponse
	response.Status = 3000
	response.Data = bsonResponse
	return response
}

func searchByUserIdsText(value interface{}) bson.M {
	query := bson.M{
		"username": bson.M{
			"$regex": value,
		},
	}
	return query
}

func SearchUsers(request dto.UserSearch) dto.UserResponse {
	searchRequest := searchByUserIdsText(request.UserName)
	bsonResponse := executeSearchFetchMultiple(properties.BLOG_BACKEND_DATABASE, properties.USER_DETAILS_COLLECTION, searchRequest, "username", int64(request.Page), int64(request.Size))
	var response dto.UserResponse
	response.Status = 3000
	response.Data = bsonResponse
	return response
}

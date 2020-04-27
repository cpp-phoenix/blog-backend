package services

import (
	"blog_backend/configuration"
	"blog_backend/dto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func saveSingleDocument(db string, collection string, document interface{}) bool {
	_, err := configuration.GetCollection(db, collection).InsertOne(context.TODO(), document)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func updateSingleDocument(db string, collection string, searchRequest dto.SearchRequest, savedocument bson.M) int {
	//converting search request to Bson request
	bsonSearchRequest := createDTOToBsonRequest(searchRequest)

	_, err := configuration.GetCollection(db, collection).UpdateOne(context.TODO(), bsonSearchRequest, savedocument)

	if err != nil {
		log.Fatal(err)
		return 3012
	}
	return 3013
}

func executeSearch(db string, collection string, searchRequest dto.SearchRequest) bson.M {

	//converting search request to Bson request
	bsonSearchRequest := createDTOToBsonRequest(searchRequest)
	var document bson.M
	// Create a string using ` string escape ticks
	// Declare an empty BSON Map object

	err := configuration.GetCollection(db, collection).FindOne(configuration.Ctx(), bsonSearchRequest).Decode(&document)
	if err != nil {
		return document
	}
	return document
}

package services

import (
	"blog_backend/configuration"
	"blog_backend/dto"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func updateManyDocuments(db string, collection string, searchRequest dto.SearchRequest, savedocument bson.M) int {
	//converting search request to Bson request
	bsonSearchRequest := createDTOToBsonRequest(searchRequest)

	_, err := configuration.GetCollection(db, collection).UpdateMany(context.TODO(), bsonSearchRequest, savedocument)

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

func executeSearchFetchMultiple(db string, collection string, searchRequest bson.M, sortBy string, page int64, limit int64) []bson.M {
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(limit * (page - 1))
	findOptions.SetSort(bson.D{{sortBy, 1}})

	var documents []bson.M

	cursor, err := configuration.GetCollection(db, collection).Find(configuration.Ctx(), searchRequest, findOptions)
	err = cursor.All(configuration.Ctx(), &documents)
	if err != nil {
		return documents
	}
	return documents
}

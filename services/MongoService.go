package services

import (
	"blog_backend/configuration"
	"blog_backend/dto"
	"context"
	"encoding/json"
	"fmt"
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

func executeSearch(db string, collection string, searchRequest dto.SearchRequest) bson.M {

	//converting search request to Bson request
	bsonSearchRequest := createDTOToBsonRequest(searchRequest)
	fmt.Println(bsonSearchRequest)
	var document bson.M
	// Create a string using ` string escape ticks
	query11 := `{ "username" : { $eq: "pappu" } }`
	fmt.Println(query11)
	var bsonMap bson.M
	err := json.Unmarshal([]byte(query11), &bsonMap)
	fmt.Println(bsonMap)
	// Declare an empty BSON Map object

	err = configuration.GetCollection(db, collection).FindOne(configuration.Ctx(), bsonSearchRequest).Decode(&document)
	if err != nil {
		fmt.Println("rrrrrrr")
		log.Fatal(err)
	}
	fmt.Println(document)
	return document
}

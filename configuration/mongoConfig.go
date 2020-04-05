package configuration

import (
	"blog_backend/properties"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client, err = mongo.NewClient(options.Client().ApplyURI(properties.MONGODB_URL))
var Ctx, _ = context.WithTimeout(context.Background(), properties.MONGO_CONNECTION_TIMEOUT*time.Second)

func ConnectToMongo() {
	createClient()
}

func createClient() {
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Connect(Ctx)
	if err != nil {
		fmt.Println("Unable to create connection !")
		log.Fatal(err)
	}
	defer Client.Disconnect(Ctx)
}

func GetDatabase(database string) *mongo.Database {
	err = Client.Connect(Ctx)
	return Client.Database(database)
}

func GetCollection(database string, collection string) *mongo.Collection {
	return GetDatabase(database).Collection(collection)
}

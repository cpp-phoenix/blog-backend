package main

import (
	"blog_backend/configuration"
	"blog_backend/controller"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Connecting to mongoDB!!")
	configuration.ConnectToMongo()
	fmt.Println("Successfully Connected to mongoDB!!")
	os.Setenv("PORT", "5000")
	fmt.Println("Starting the server!!")
	controller.Routes()
}

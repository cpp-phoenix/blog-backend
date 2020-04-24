package main

import (
	"blog_backend/configuration"
	"blog_backend/controller"
	"fmt"
)

func main() {

	fmt.Println("Connecting to mongoDB!!")
	configuration.ConnectToMongo()
	fmt.Println("Successfully Connected to mongoDB!!")

	fmt.Println("Starting the server!!")
	controller.Routes()
}

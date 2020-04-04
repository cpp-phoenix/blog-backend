package main

import (
	"blog_backend/controller"
	"fmt"
)
func main() {
	fmt.Println("Starting the server!!")
    controller.Routes()
}
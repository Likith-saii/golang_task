package main

import (
	"fmt"
	"golang_task/golang_task/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Define GET endpoint
	r.GET("/person/:person_id/info", service.GetPersonInfo)

	// Define POST endpoint
	r.POST("/person/create", service.CreatePerson)

	// Start server
	port := ":8080"
	fmt.Println("Server running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

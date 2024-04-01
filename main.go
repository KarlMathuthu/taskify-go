package main

import (
	"log"
	"os"

	"github.com/KarlMathuthu/taskify-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Get the port from .env file, if it's not there then set PORT to 8080
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// Create a new gin instance.
	router := gin.Default()

	// API Routes
	router.GET("/tasks/", routes.GetAllTasks)
	router.POST("/tasks/", routes.AddTask)
	router.GET("/tasks/:id", routes.GetEachTask)

	// Run gin, also catch the error, if it's there then log.Fatal(err)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

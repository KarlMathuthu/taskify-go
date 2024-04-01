package main

import (
	"log"
	"os"

	"github.com/KarlMathuthu/taskify-go/database"
	"github.com/KarlMathuthu/taskify-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// Get the port & dbUrl from .env file, Handle errors.
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DBURL")

	if dbUrl == "" {
		dbUrl = "mongodb://localhost:27017"
	}
	if port == "" {
		port = "8080"
	}

	// Create a new gin instance.
	router := gin.Default()

	// Init MongoDB.
	database.MongoDBInit(dbUrl)

	// API Routes
	router.GET("/tasks/", routes.GetAllTasks)
	router.POST("/tasks/", routes.AddTask)
	router.GET("/tasks/:id", routes.GetEachTask)
	router.PUT("/tasks/:id", routes.UpdateTask)
	router.DELETE("/tasks/:id", routes.DeleteTask)

	// Run gin, also catch the error, if it's there then log.Fatal(err)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

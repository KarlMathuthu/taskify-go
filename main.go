package main

import (
	"log"
	"os"

	"github.com/KarlMathuthu/taskify-go/controllers"
	"github.com/KarlMathuthu/taskify-go/database"
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
	router.GET("/", controllers.Welcome)
	router.GET("/tasks/", controllers.GetAllTasks)
	router.POST("/tasks/", controllers.AddTask)
	router.GET("/tasks/:id", controllers.GetEachTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.PATCH("/tasks/:id", controllers.UpdateIsDoneField)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	// Run gin, also catch the error, if it's there then log.Fatal(err)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

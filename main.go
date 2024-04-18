package main

import (
	"log"
	"os"

	"github.com/KarlMathuthu/taskify-go/database"
	"github.com/KarlMathuthu/taskify-go/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	dbUrl := os.Getenv("DBURL")

	if dbUrl == "" {
		dbUrl = "mongodb://localhost:27017"
	}

	// Init MongoDB.
	database.MongoDBInit(dbUrl)

}

func main() {

	// Get the port  from .env file, Handle errors.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new gin instance.
	router := gin.Default()

	// API Routes from routes package
	routers.TasksRoutes(router)

	// Run gin, also catch the error, if it's there then log.Fatal(err)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

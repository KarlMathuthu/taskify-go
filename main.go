package main

import (
	"log"
	"net/http"
	"os"

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

	router.GET("/api/", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")

		type Response struct {
			Message string `json:"message"`
		}
		message := Response{
			Message: "Welcome to Gin!",
		}
		ctx.JSON(http.StatusOK, message)
	})

	// Run gin, also catch the error, if it's there then log.Fatal(err)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

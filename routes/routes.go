package routes

import (
	"net/http"

	"github.com/KarlMathuthu/taskify-go/models"
	"github.com/gin-gonic/gin"
)

// slice where we store our tasks
// var, name, type and value
var tasks []models.Task

// CRUD Functions.

// Get tasks
func GetAllTasks(ctx *gin.Context) {
	// Set a custom header
	ctx.Header("Content-Type", "application/json")
	// Get all tasks in the slice "tasks"
	ctx.JSON(http.StatusOK, tasks)
}

// Create task
func AddTask(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	task := models.Task{
		ID:          "1",
		Title:       "Hello!",
		Description: "Nothing",
		IsDone:      false,
	}
	tasks = append(tasks, task)
	ctx.JSON(http.StatusOK, task)
}

// Update task
// Delete task

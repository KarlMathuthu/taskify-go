package controllers

import (
	"net/http"

	"github.com/KarlMathuthu/taskify-go/database"
	"github.com/KarlMathuthu/taskify-go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// slice where we store our tasks
// var, name, type and value
var tasks []models.Task

// CRUD Functions.

// Get tasks
func GetAllTasks(ctx *gin.Context) {
	// Set a custom header
	ctx.Header("Content-Type", "application/json")
	myTasks := database.GetAllTasksDB()
	// Get all tasks from the database.
	ctx.JSON(http.StatusOK, myTasks)
}

// Create task
func AddTask(ctx *gin.Context) {
	// Set a custom header
	ctx.Header("Content-Type", "application/json")
	// Get an instance of Task Model
	var task models.Task
	// make ID the doc Object ID.
	task.ID = primitive.NewObjectID()
	// make the taskId the same as the above ID.
	task.TaskId = task.ID.Hex()

	if err := ctx.BindJSON(&task); err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}
	result := database.AddTaskToDB(task)
	response := bson.M{
		"message": result,
	}

	ctx.JSON(http.StatusCreated, response)
}

// Find the task id as int
func FindTaskID(id string) *int {
	for i, t := range tasks {
		if t.TaskId == id {
			return &i
		}
	}
	return nil
}

// Get Each Task
func GetEachTask(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	taskId := ctx.Param("id")
	task, err := database.FindTaskDB(taskId)
	if err != nil {
		response := models.Response{
			Message: "task Not Found",
		}
		ctx.JSON(http.StatusNotFound, response)
	}
	ctx.JSON(http.StatusOK, task)
}

// Update task
func UpdateTask(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	// Get task ID from the Params
	taskId := ctx.Param("id")
	taskIntId := FindTaskID(taskId)
	var newTask models.Task

	newTask.TaskId = taskId

	if err := ctx.BindJSON(&newTask); err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	tasks[*taskIntId] = newTask
	ctx.JSON(http.StatusOK, newTask)
}

// Delete task
func DeleteTask(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	taskId := ctx.Param("id")

	result, err := database.DeleteTaskDB(taskId)

	if err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
	}
	ctx.JSON(http.StatusOK, bson.M{"message": result})
}

package controllers

import (
	"net/http"

	"github.com/KarlMathuthu/taskify-go/database"
	"github.com/KarlMathuthu/taskify-go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Welcome(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{"message": "Hi, Welcome to Taskify!"})
}

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
	var newTask models.Task

	newTask.TaskId = taskId

	if err := ctx.BindJSON(&newTask); err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}
	result, err := database.UpdateTaskDB(taskId, newTask)

	if err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})
}

// Update is done field
func UpdateIsDoneField(ctx *gin.Context) {
	// set custom headers
	ctx.Header("Content-Type", "application/json")
	taskId := ctx.Param("id")
	var newTask models.Task

	if err := ctx.BindJSON(&newTask); err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}
	result, err := database.UpdateTaskDone(taskId, newTask)
	if err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": result})
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

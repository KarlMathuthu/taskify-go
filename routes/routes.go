package routes

import (
	"errors"
	"net/http"

	"github.com/KarlMathuthu/taskify-go/database"
	"github.com/KarlMathuthu/taskify-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	// Set a custom header
	ctx.Header("Content-Type", "application/json")
	// Get an instance of Task Model
	var task models.Task
	// Asign task ID to the UUID.
	task.ID = uuid.NewString()

	if err := ctx.BindJSON(&task); err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}
	result := database.AddTaskToDB(task)

	ctx.JSON(http.StatusCreated, result)
}

// Find the task with the ID
func FindTask(id string) (*models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task Not found")
}

// Find the task id as int
func FindTaskID(id string) *int {
	for i, t := range tasks {
		if t.ID == id {
			return &i
		}
	}
	return nil
}

// Get Each Task
func GetEachTask(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	taskId := ctx.Param("id")
	task, err := FindTask(taskId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
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

	newTask.ID = taskId

	if err := ctx.BindJSON(&newTask); err != nil {
		response := models.Response{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
	}

	tasks[*taskIntId] = newTask
	ctx.JSON(http.StatusOK, newTask)
}

// Remove Task from slice
func RemoveTask(slice []models.Task, index int) []models.Task {
	return append(slice[:index], slice[index+1:]...)
}

// Delete task
func DeleteTask(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	taskId := ctx.Param("id")
	taskIndex := FindTaskID(taskId)
	newTasks := RemoveTask(tasks, *taskIndex)

	ctx.JSON(http.StatusOK, newTasks)
}

package routers

import (
	"github.com/KarlMathuthu/taskify-go/controllers"
	"github.com/gin-gonic/gin"
)

func TasksRoutes(incommingRoutes *gin.Engine) {
	router := incommingRoutes.Group("/api/")
	router.GET("/", controllers.Welcome)
	router.GET("/tasks/", controllers.GetAllTasks)
	router.POST("/tasks/", controllers.AddTask)
	router.GET("/tasks/:id", controllers.GetEachTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.PATCH("/tasks/:id", controllers.UpdateIsDoneField)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}

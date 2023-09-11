package routes

import (
	"ginGonic/learn/controllers"

	"github.com/gin-gonic/gin"
)

func ClassRoute(route *gin.Engine) {
	route.POST("/class", controllers.CreateClass())
	route.GET("/class/:classId", controllers.GetClass())
	route.PUT("/class/:classId", controllers.EditClass())
	route.DELETE("/class/:classId", controllers.DeleteClass())
	route.GET("/class", controllers.GetAllClass())
}

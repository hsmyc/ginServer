package routes

import (
	"ginGonic/learn/controllers"

	"github.com/gin-gonic/gin"
)

func CharRoute(router *gin.Engine) {
	router.POST("/char", controllers.CreateChar())
	router.GET("/char/:charId", controllers.GetChar())
	router.PUT("/char/:charId", controllers.EditChar())
	router.DELETE("/char/:charId", controllers.DeleteChar())
	router.GET("/chars", controllers.GetAllChars())
}

package routes

import (
	"ginGonic/learn/controllers"
	"ginGonic/learn/services"

	"github.com/gin-gonic/gin"
)

func CharRoute(router *gin.Engine) {
	router.POST("/char", controllers.CreateChar())
	router.GET("/char/:charId", controllers.GetChar())
	router.PUT("/char/:charId", controllers.EditChar())
	router.DELETE("/char/:charId", controllers.DeleteChar())
	router.GET("/chars", controllers.GetAllChars())
	router.PUT("/chars/:charId/class/:classId", services.LinkCharToClass())
	router.PUT("/chars/:charId/item/:itemId", services.LinkCharToItem())
}

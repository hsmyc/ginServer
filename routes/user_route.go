package routes

import (
	"ginGonic/learn/controllers"
	"ginGonic/learn/services"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.GET("/users", controllers.GetAllUsers())
	router.PUT("/users/:userId/character/:charId", services.LinkUserToCharacter())

}

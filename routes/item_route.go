package routes

import (
	"ginGonic/learn/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoute(router *gin.Engine) {
	router.POST("/item", controllers.CreateItem())
	router.GET("/item/:itemId", controllers.GetItem())
	router.PUT("/item/:itemId", controllers.EditItem())
	router.DELETE("/item/:itemId", controllers.DeleteItem())
	router.GET("/items", controllers.GetAllItems())
}

package main

import (
	"ginGonic/learn/configs"
	"ginGonic/learn/routes"

	"github.com/gin-gonic/gin"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	osman := Human{"Osman", 25}
	router := gin.Default()
	//db
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)
	routes.ItemRoute(router)
	routes.CharRoute(router)
	routes.ClassRoute(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": osman,
		})
	})
	router.Run("localhost:3131")
}

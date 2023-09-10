package main

import (
	"github.com/gin-gonic/gin"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	osman := Human{"Osman", 25}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": osman,
		})
	})
	router.Run("localhost:3131")
}

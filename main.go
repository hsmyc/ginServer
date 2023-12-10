package main

import (
	"ginGonic/learn/configs"
	"ginGonic/learn/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Token"}
	router.Use(cors.New(config))
	//db
	configs.ConnectDB()
	//routes
	routes.UserRoute(router)
	routes.ItemRoute(router)
	routes.CharRoute(router)
	routes.ClassRoute(router)

	router.Run("localhost:3131")
}

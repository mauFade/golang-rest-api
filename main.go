package main

import (
	"example.com/web-service-gin/config"
	"example.com/web-service-gin/routes"
	"github.com/gin-gonic/gin"
)


func main()  {
	router := gin.Default()

	config.Connect()
	
	routes.AlbumRoutes(router)

	router.Run("localhost:777")
}
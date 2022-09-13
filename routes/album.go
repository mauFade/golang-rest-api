package routes

import (
	"example.com/web-service-gin/controllers"
	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.Engine) {
	router.GET("/albums", controllers.FetchAlbums)
	router.POST("/albums", controllers.CreateAlbum)
	router.PUT("/albums", controllers.UpdateAlbum)
	router.DELETE("/albums", controllers.DeleteAlbum)
}
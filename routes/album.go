package routes

import (
	"example.com/web-service-gin/controllers"
	"github.com/gin-gonic/gin"
)

func AlbumRoutes(router *gin.Engine) {
	router.GET("/albums", controllers.FetchAlbums)
	router.GET("/albums/:id", controllers.FetchAlbumById)
	router.POST("/albums", controllers.CreateAlbum)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
}
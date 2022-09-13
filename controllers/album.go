package controllers

import (
	"example.com/web-service-gin/config"
	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func FetchAlbums(c *gin.Context) {
	albums := []models.Album{}

	config.DB.Find(&albums)
	
	c.JSON(200, &albums)
}

func CreateAlbum(c *gin.Context) {
	var album models.Album

	c.BindJSON(&album)

	config.DB.Create(&album)

	c.JSON(200, &album)
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album

	config.DB.Where("id = ?", c.Param("id")).Delete(&album)

	c.String(200, "Album deleted successfully")
}

func UpdateAlbum(c *gin.Context) {	
	var album models.Album

	config.DB.Where("id = ?", c.Param("id")).First(&album)

	c.BindJSON(&album)

	config.DB.Save(&album)

	c.JSON(200, &album)
}
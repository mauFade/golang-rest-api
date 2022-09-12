package controllers

import (
	"github.com/gin-gonic/gin"
)

func AlbumController(c *gin.Context) {
	c.String(200, "Album controller")
}
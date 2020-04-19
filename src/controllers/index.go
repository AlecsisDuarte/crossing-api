package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Crossing Restful server",
		"title":   "Crossing API",
	})
}

func V1Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Crossing API v1",
		"title":   "Crossing API v1",
	})
}

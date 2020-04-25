package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index returns the root enpoint of the Restful API
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Crossing Restful server",
		"title":   "Crossing API",
	})
}

// V1Index returns the root enpoint of the V1 Restful API
func V1Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Crossing API v1",
		"title":   "Crossing API v1",
	})
}

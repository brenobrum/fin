package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupHealthCheck(router *gin.Engine) {
	// Define a route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

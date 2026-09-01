package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "1.0.0"

func main() {
	router := gin.Default()

	router.GET("/health", health)
	router.GET("/ready", ready)
	router.GET("/version", getVersion)

	router.Run(":8080")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
	})
}

func getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
}

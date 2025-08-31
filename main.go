package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisshihi/golang-lessons/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	
	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"events":  events,
	})
}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error": "bad request body",
		})
		return
	}
}
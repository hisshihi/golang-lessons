package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisshihi/golang-lessons/db"
	"github.com/hisshihi/golang-lessons/models"
)

func main() {
	db.InitDB()
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
		log.Printf("body: %v, err: %v", event, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "bad request body",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"event":   event,
	})
}

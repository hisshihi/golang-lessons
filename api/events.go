// Package api отвечает за маршрутизацию и обработку HTTP-запросов
package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisshihi/golang-lessons/models"
	"gorm.io/gorm"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"events":  events,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "error with get events",
		})
		return
	}

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

	event.ID = 1
	event.UserID = 1

	err := event.Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "don`t save a event " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"event":   event,
	})
}

type eventByID struct {
	ID int64 `uri:"id" binding:"required"`
}

func getEventByID(c *gin.Context) {
	var req eventByID
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "bad request body",
		})
		return
	}

	event, err := models.GetEventByID(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"event":   event,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "error with get event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"event":   event,
	})
}

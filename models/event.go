// Package models - endpoint`s loginc
package models

import (
	"github.com/gin-gonic/gin"
)

type Event struct {
	ID          int    `binding:"required" json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	Location    string `binding:"required" json:"location"`
	// DateTime    time.Time `binding:"required" json:"datetime"`
	UserID int `binding:"required" json:"user_id"`
}

var events = []Event{}

func (e *Event) Save(c *gin.Context) {
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}

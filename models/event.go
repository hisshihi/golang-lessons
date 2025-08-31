// Package models - endpoint`s loginc
package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save(c *gin.Context) {
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}

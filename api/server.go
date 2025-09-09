package api

import "github.com/gin-gonic/gin"

const eventsID = "/events/:id"

func Server(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET(eventsID, getEventByID)
	server.PUT(eventsID, updateEventByID)
	server.DELETE(eventsID, deleteEventByID)
}

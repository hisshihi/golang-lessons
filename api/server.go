package api

import "github.com/gin-gonic/gin"

func Server(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEventByID)
}

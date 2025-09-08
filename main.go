package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hisshihi/golang-lessons/api"
	"github.com/hisshihi/golang-lessons/db"
)

func main() {
	db.InitDB()
	server := gin.Default()

	api.Server(server)

	server.Run(":8080")
}

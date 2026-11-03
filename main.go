package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/db"
	"github.com/pankajy-dev/go-rest-app/models"
)

func main() {
	db.InitDb()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

// getEvent handler
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	// context.JSON(http.StatusOK, gin.H{"message":"Hello"})
	context.JSON(http.StatusOK, events)
}

// createEvent handler
func createEvent(ctx *gin.Context) {

	var event models.Event

	// event will be extracted from the request and fills the pointer just like fmt.Scan
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cannot process request"})
		return
	}

	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{"message": "New event created", "event": event})
}

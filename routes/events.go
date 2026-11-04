package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/models"
)

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get eventId"})
		return
	}
	event, err := models.GetEventId(eventId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No event id found"})
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Event found", "event": event})
}

// getEvent handler
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		panic(err)
	}
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

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	_, err = models.GetEventId(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch."})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	updatedEvent.Id = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

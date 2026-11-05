package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/models"
	"github.com/pankajy-dev/go-rest-app/utils"
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

	token := ctx.Request.Header.Get("Authorization")

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	if userId == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User not found."})
		return
	}

	var event models.Event

	// event will be extracted from the request and fills the pointer just like fmt.Scan
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cannot process request"})
		return
	}

	event.UserId = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event."})
		return
	}

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

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	event, err := models.GetEventId(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch."})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "Error during record deletion."})
		return
	}
}

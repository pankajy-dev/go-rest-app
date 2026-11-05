package routes

import (
	"log"
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
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Event found", "event": event})
}

// getEvent handler
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Not able to get the events."})
		return
	}
	// context.JSON(http.StatusOK, gin.H{"message":"Hello"})
	context.JSON(http.StatusOK, events)
}

// createEvent handler
func createEvent(ctx *gin.Context) {

	userId := ctx.GetInt64("userId")

	if userId == 0 {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "User not found."})
		return
	}

	var event models.Event

	// event will be extracted from the request and fills the pointer just like fmt.Scan
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot process request"})
		return
	}

	event.UserId = int64(userId)

	err = event.Save()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event."})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": "New event created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	event, err := models.GetEventId(eventId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch."})
		return
	}

	userId := ctx.GetInt64("userId")

	if event.UserId != userId {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to delete the event."})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	updatedEvent.Id = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	userId := ctx.GetInt64("userId")

	event, err := models.GetEventId(eventId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch."})
		return
	}

	if event.UserId != userId {
		log.Printf("user %d not authorized for event %d\n", userId, eventId)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized."})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error during record deletion."})
		return
	}
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	event, err := models.GetEventId(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Event could not be retrieved"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful."})

}

func cancelEventRegistration(ctx *gin.Context) {

	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	event, err := models.GetEventId(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Event could not be retrieved"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration could not be cancelled."})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "Registration deleted."})
}

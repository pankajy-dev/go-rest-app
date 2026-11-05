package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/internal/repository"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	event, err := repository.GetEventId(eventId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Event could not be retrieved"})
		return
	}

	err = repository.RegisterForEvent(*event, userId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful."})

}

func cancelEventRegistration(ctx *gin.Context) {

	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse."})
		return
	}

	event, err := repository.GetEventId(eventId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Event could not be retrieved"})
		return
	}

	err = repository.CancelRegistration(*event, userId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration could not be cancelled."})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "Registration deleted."})
}

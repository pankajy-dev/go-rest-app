package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	server := gin.Default()

	server.GET("/events", sendEvents)

	server.Run(":8080")
}

func sendEvents(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{"message":"Hello"})

}
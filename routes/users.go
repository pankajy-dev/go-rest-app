package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/models"
)

func signup(ctx *gin.Context) {
	var usr models.User

	// event will be extracted from the request and fills the pointer just like fmt.Scan
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cannot process request"})
		return
	}

	err = usr.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User cannot be saved."})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "New event created", "user": usr})
}

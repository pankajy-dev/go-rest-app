package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/models"
	"github.com/pankajy-dev/go-rest-app/utils"
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
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "New event created", "user": usr})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request."})
		return
	}

	if !user.ValidateCredentials() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
		return
	}

	jwtToken, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "JWT token creation errored."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Authenticated", "token": jwtToken})

}

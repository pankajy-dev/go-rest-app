package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/internal/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	userId, err := utils.VerifyToken(token)
	if err != nil {
		log.Printf("User not authorized %v\n", userId)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}

	// We can set a value in the context that can be retrieved later.
	ctx.Set("userId", userId)

	// Next handler will get the control
	ctx.Next()
}

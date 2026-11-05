package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pankajy-dev/go-rest-app/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	userId, err := utils.VerifyToken(token)
	if err != nil {
	}

	// We can set a value in the context that can be retrieved later.
	ctx.Set("userId", userId)

	// Next handler will get the control
	ctx.Next()
}

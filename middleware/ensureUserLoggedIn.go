package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/metaphi-org/go-infra-sdk/helpers"
)

func EnsureUserLoggedIn() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		user := GetUserFromContext(ctx)

		if len(user.UserId) == 0 {
			helpers.RenderUnauthorizedError(*ctx)
			ctx.Abort()
			return
		}

	}
}

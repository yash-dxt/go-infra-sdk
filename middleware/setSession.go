package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/metaphi-org/go-infra-sdk/helpers"
)

const sessionParam = "session"

const guest_session = "guest_session"

func SetSession() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		sessionId := helpers.GetSessionFromRequest(ctx)

		if len(sessionId) == 0 {
			ctx.Set(sessionParam, guest_session)
		} else {
			ctx.Set(sessionParam, sessionId)
		}

	}

}

func GetSessionIdFromSetSessionMiddleware(ctx gin.Context) string {
	sessionId, _ := ctx.Get(sessionParam)
	return sessionId.(string)
}

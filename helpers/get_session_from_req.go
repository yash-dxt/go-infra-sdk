package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/yash-dxt/go-infra-sdk/constants"
)

func GetSessionFromRequest(ctx *gin.Context) string {
	cookie, err := ctx.Request.Cookie(constants.AuthHeaderKey)

	var sessionId string
	if err == nil {
		sessionId = cookie.Value
	} else {
		sessionId = ctx.Request.Header.Get(constants.AuthHeaderKey)
	}

	return sessionId
}

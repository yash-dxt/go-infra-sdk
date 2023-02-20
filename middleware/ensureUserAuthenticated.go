package middleware

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/metaphi-org/go-infra-sdk/constants"
	"github.com/metaphi-org/go-infra-sdk/helpers"
	"github.com/metaphi-org/go-infra-sdk/utils"
)

func EnsureUserAuthenticated(config aws.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		sessionId := helpers.GetSessionFromRequest(ctx)

		if len(sessionId) == 0 {
			helpers.RenderUnauthorizedError(*ctx)
			ctx.Abort()
			return
		}

		userId, errApiReq := utils.GetUserIdFromAuthService(ctx, config, sessionId)

		if errApiReq != nil {

			if errApiReq.Error() == constants.ErrorInvalidSessionId {

				helpers.RenderUnauthorizedError(*ctx)
				ctx.Abort()
				return

			} else {

				helpers.RenderInternalServerError(*ctx, errApiReq.Error())
				ctx.Abort()
				return
			}

		}

		ctx.Request.Header.Set("userId", userId)
	}
}

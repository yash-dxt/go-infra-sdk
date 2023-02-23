package middleware

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/metaphi-org/go-infra-sdk/constants"
	"github.com/metaphi-org/go-infra-sdk/helpers"
	"github.com/metaphi-org/go-infra-sdk/utils"
)

type AuthenticatedUserFields struct {
	UserId string
}

const userParam = "user"

func EnsureUserAuthenticated(config aws.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		sessionId := GetSessionIdFromSetSessionMiddleware(*ctx) // comes from previous setSessionId middleware.

		if sessionId == guest_session {
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

		ctx.Set(userParam, AuthenticatedUserFields{
			UserId: userId,
		})
	}
}

func GetUserIdFromAuthenticateUserMiddleware(ctx *gin.Context) string {
	userGet, _ := ctx.Get(userParam)
	user := userGet.(AuthenticatedUserFields)
	return user.UserId
}

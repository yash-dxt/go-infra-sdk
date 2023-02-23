package middleware

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/metaphi-org/go-infra-sdk/constants"
	"github.com/metaphi-org/go-infra-sdk/helpers"
	"github.com/metaphi-org/go-infra-sdk/utils"
)

const userParam = "user"

type UserFields struct {
	UserId string
}

func RetrieveUserIdFromAuthServiceMiddleware(config aws.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		sessionId := helpers.GetSessionFromRequest(ctx)

		if len(sessionId) == 0 {
			return
		}

		userId, err := utils.GetUserIdFromAuthService(ctx, config, sessionId)

		if err != nil {
			if err.Error() != constants.ErrorInvalidSessionId {
				helpers.RenderInternalServerError(*ctx, err.Error())
				ctx.Abort()
				return
			}
		}

		ctx.Set(userParam, UserFields{
			UserId: userId,
		})

	}

}

func GetUserIdFromContext(ctx *gin.Context) string {
	userGet, ok := ctx.Get(userParam)

	if !ok {
		return ""
	}
	user, ok := userGet.(UserFields)

	if !ok {
		return ""
	}

	return user.UserId
}

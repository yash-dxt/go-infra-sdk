package middleware

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gin-gonic/gin"
	"github.com/yash-dxt/go-infra-sdk/constants"
	"github.com/yash-dxt/go-infra-sdk/helpers"
	"github.com/yash-dxt/go-infra-sdk/utils"
)

const userParam = "user"

type UserFields struct {
	UserId string
}

func RetrieveUser(config aws.Config) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		sessionId := helpers.GetSessionFromRequest(ctx)

		if len(sessionId) == 0 {
			return
		}

		userId, err := utils.GetUserIdFromAuthService(ctx, config, sessionId)

		if err != nil && err.Error() != constants.ErrorInvalidSessionId {
			helpers.RenderInternalServerError(*ctx, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set(userParam, UserFields{
			UserId: userId,
		})

	}

}

func GetUserFromContext(ctx *gin.Context) UserFields {
	userGet, ok := ctx.Get(userParam)

	if !ok {
		return UserFields{}
	}
	user, ok := userGet.(UserFields)

	if !ok {
		return UserFields{}
	}

	return user
}

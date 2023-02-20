package utils

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/metaphi-org/go-infra-sdk/constants"
	"github.com/metaphi-org/go-infra-sdk/lambda"
)

const auth_service_userid_endpoint = "/auth/userId"

var auth_service_lambda = os.Getenv("AUTH_SERVICE_LAMBDA_NAME")

type CheckSessionAuthRequest struct {
	UserId string `json:"userId"`
}

// ensure AUTH_SERVICE_LAMBDA is set as environment variable.
func GetUserIdFromAuthService(ctx context.Context, config aws.Config, sessionId string) (string, error) {

	var userId string
	res, err := lambda.InvokeLambda[CheckSessionAuthRequest](ctx, config, auth_service_lambda, lambda.RequestParams{
		Endpoint: auth_service_userid_endpoint,
		Method:   "GET",
		Headers: map[string]string{
			constants.AuthHeaderKey: sessionId,
		},
	})

	if err != nil {
		return userId, err
	}

	if len(res.UserId) == 0 {
		return userId, errors.New(constants.ErrorInvalidSessionId)
	}

	userId = res.UserId
	return userId, nil
}

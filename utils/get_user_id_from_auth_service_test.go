package utils_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yash-dxt/go-infra-sdk/config"
	"github.com/yash-dxt/go-infra-sdk/constants"
	"github.com/yash-dxt/go-infra-sdk/utils"
)

var REGION = os.Getenv("AWS_REGION")
var LAMBDA_NAME = os.Getenv("TEST_LAMBDA")

func TestGetUserIdFromAuthService(t *testing.T) {
	user_id, err := utils.GetUserIdFromAuthService(context.Background(), config.CreateAWSConfig(REGION), "41941e6345fb2a2c33f037f84d33b7f1fbae")
	assert.Nil(t, err)
	assert.NotNil(t, user_id)

	_, err = utils.GetUserIdFromAuthService(context.Background(), config.CreateAWSConfig(REGION), "wrong_session_id")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), constants.ErrorInvalidSessionId)
}

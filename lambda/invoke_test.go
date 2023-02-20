package lambda_test

import (
	"context"
	"os"
	"testing"

	"github.com/metaphi-org/go-infra-sdk/config"
	"github.com/metaphi-org/go-infra-sdk/lambda"
	"github.com/stretchr/testify/assert"
)

type HelloWorld struct {
	Msg string `json:"msg"`
}

var REGION = os.Getenv("AWS_REGION")
var LAMBDA_NAME = os.Getenv("AUTH_SERVICE_LAMBDA")

func TestInvokeLambda(t *testing.T) {
	res, err := lambda.InvokeLambda[HelloWorld](context.Background(), config.CreateAWSConfig(REGION), LAMBDA_NAME, lambda.RequestParams{
		Endpoint: "/",
		Method:   "GET",
	})

	assert.Nil(t, err)
	assert.Equal(t, res.Msg, "helloworld")
}

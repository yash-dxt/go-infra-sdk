package lambda_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yash-dxt/go-infra-sdk/config"
	"github.com/yash-dxt/go-infra-sdk/lambda"
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

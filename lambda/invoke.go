package lambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

type RequestParams struct {
	Endpoint string
	Body     []byte
	Method   string
	Headers  map[string]string
	Cookies  []string
}

type lambdaResponse struct {
	Body string `json:"body"`
}

func InvokeLambda[I any](ctx context.Context, config aws.Config, lambda_name string, request_param RequestParams) (I, error) {

	var res I
	svc := lambda.New(lambda.Options{
		Region:      config.Region,
		Credentials: config.Credentials,
	})

	mapBody := events.APIGatewayV2HTTPRequest{
		RawPath: request_param.Endpoint, // the endpoint of the API goes here.
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: request_param.Method, // Method of API req goes here.
			},
		},
		Body:    string(request_param.Body),
		Headers: request_param.Headers,
		Cookies: request_param.Cookies,
	}

	payload, err := json.Marshal(mapBody)

	if err != nil {
		return res, err
	}

	result, err := svc.Invoke(ctx, &lambda.InvokeInput{
		FunctionName:   &lambda_name,
		InvocationType: types.InvocationTypeRequestResponse,
		Payload:        payload,
	})

	if err != nil {
		return res, err
	}

	var m lambdaResponse
	err = json.Unmarshal(result.Payload, &m)

	if err != nil {
		return res, err
	}

	var data I
	err = json.Unmarshal([]byte(m.Body), &data)
	return data, err
}

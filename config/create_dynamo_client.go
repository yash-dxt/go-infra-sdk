package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Ideally should be used only once throughout the app.
// and then should be cached.
func CreateNewDynamoClient(config aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(config)
}

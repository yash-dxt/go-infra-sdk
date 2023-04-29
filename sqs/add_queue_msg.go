package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func AddMessageToQueue(config aws.Config, ctx context.Context, queueName string, message string, delay int) error {
	client := sqs.NewFromConfig(config)

	result, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		return err
	}

	queueURL := result.QueueUrl

	_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:     queueURL,
		MessageBody:  &message,
		DelaySeconds: int32(delay),
	})

	return err
}

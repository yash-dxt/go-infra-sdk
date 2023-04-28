package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func DeleteMessageFromQueue(config aws.Config, ctx context.Context, queueName string, receiptHandle string) error {
	client := sqs.NewFromConfig(config)

	qUrlOutput, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})

	if err != nil {
		return err
	}

	_, err = client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      qUrlOutput.QueueUrl,
		ReceiptHandle: &receiptHandle,
	})

	return err
}

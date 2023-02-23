package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/metaphi-org/go-infra-sdk/helpers"
)

func TransactWriteItems(ctx context.Context, db dynamodb.Client, transactItems []types.TransactWriteItem) error {
	// Limitation of 100 requests per transaction.
	chunks := helpers.ArrayChunk(transactItems, 100)

	for _, chunk := range chunks {

		_, err := db.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
			TransactItems: chunk,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

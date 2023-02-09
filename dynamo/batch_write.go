package dynamo

import (
	"context"

	"github.com/metaphi-org/go-infra-sdk/helpers"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Batch write items to a single table.
func WriteAllItemsUsingBatchWrite[I any](dynamoClient dynamodb.Client, ctx context.Context, tableName *string, arr []I) error {
	return batchWrite(dynamoClient, ctx, tableName, arr)
}

// currently only made specific to tableName.
func batchWrite[I any](db dynamodb.Client, ctx context.Context, tableName *string, arr []I) error {

	chunks := helpers.ArrayChunk(arr, 25)

	for _, chunk := range chunks {

		writes, err := helpers.MakeArrayOfWriteRequests(chunk)

		if err != nil {
			return err
		}

		_, errBatchWrite := db.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				*tableName: writes,
			}})

		if errBatchWrite != nil {
			return errBatchWrite
		}

	}

	return nil
}

package dynamo

import (
	"context"

	"github.com/metaphi-org/go-infra-sdk/helpers"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Batch Delete (which is actually a type of write) items to a single table.
func DeleteAllItemsUsingBatchWrite[I any](dynamoClient dynamodb.Client, ctx context.Context, tableName *string, arr []I) error {
	return batchDelete(dynamoClient, ctx, tableName, arr)
}

// takes in an array of keys for a table to be deleted.
func batchDelete[I any](db dynamodb.Client, ctx context.Context, tableName *string, arr []I) error {

	// Limitation of 25 items in a batch write query.
	chunks := helpers.ArrayChunk(arr, 25)

	for _, chunk := range chunks {

		writes, err := helpers.MakeArrayOfDeleteWriteRequests(chunk)

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

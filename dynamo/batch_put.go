package dynamo

import (
	"context"

	"github.com/yash-dxt/go-infra-sdk/helpers"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Batch write items to a single table.
func PutAllItemsUsingBatchWrite[I any](dynamoClient dynamodb.Client, ctx context.Context, tableName *string, arr []I) error {
	return batchPut(dynamoClient, ctx, tableName, arr)
}

// currently only made specific to tableName.
func batchPut[I any](db dynamodb.Client, ctx context.Context, tableName *string, arr []I) error {

	// limitation of 25 items in a btach write query.
	chunks := helpers.ArrayChunk(arr, 25)

	for _, chunk := range chunks {

		writes, err := helpers.MakeArrayOfPutWriteRequests(chunk)

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

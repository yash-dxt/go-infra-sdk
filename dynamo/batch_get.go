package dynamo

import (
	"context"

	"github.com/yash-dxt/go-infra-sdk/helpers"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Batch get items from a single table.
func GetAllItemsUsingBatchGet(dynamoClient dynamodb.Client, ctx context.Context, tableName *string, requestItems []map[string]types.AttributeValue) ([]map[string]types.AttributeValue, error) {

	// batch get has a limitation of 100 items per call.
	chunks := helpers.ArrayChunk(requestItems, 100)
	var res []map[string]types.AttributeValue = make([]map[string]types.AttributeValue, 0)

	for _, chunk := range chunks {

		keysAndAttributes := map[string]types.KeysAndAttributes{
			*tableName: {
				Keys: chunk,
			},
		}

		batchGetResult, err := batchGet(dynamoClient, ctx, tableName, keysAndAttributes)

		if err != nil {
			return res, err
		}

		res = append(res, batchGetResult...)
	}

	return res, nil
}

func batchGet(db dynamodb.Client, ctx context.Context, tableName *string, requestItems map[string]types.KeysAndAttributes) ([]map[string]types.AttributeValue, error) {

	items := []map[string]types.AttributeValue{}

	result, errBatchGet := db.BatchGetItem(ctx, &dynamodb.BatchGetItemInput{
		RequestItems: requestItems,
	})

	if errBatchGet != nil {
		return items, errBatchGet
	}

	items = append(items, result.Responses[*tableName]...)

	// From Docs (https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchGetItem.html) :
	// If there are no unprocessed keys remaining, the response contains an empty UnprocessedKeys map.
	// We'll be checking for empty map here.

	if len(result.UnprocessedKeys) != 0 {
		moreItems, err := batchGet(
			db,
			ctx, tableName, result.UnprocessedKeys,
		)
		if err != nil {
			return items, err
		}
		items = append(items, moreItems...)
	}

	return items, nil
}

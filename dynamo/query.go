package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetAllItemsUsingQuery(
	dynamoClient dynamodb.Client,
	ctx context.Context,
	tableName *string,
	partitionKeyParam string,
	partitionKey string,
	indexName *string,
) ([]map[string]types.AttributeValue, error) {

	return getAllItemsUsingQuery(
		dynamoClient,
		ctx,
		tableName,
		partitionKeyParam,
		partitionKey,
		nil,
		indexName,
	)
}

func getAllItemsUsingQuery(
	dbInst dynamodb.Client,
	ctx context.Context,
	tableName *string,
	partitionKeyParam string,
	partitionKey string,
	startKey map[string]types.AttributeValue,
	indexName *string,
) ([]map[string]types.AttributeValue, error) {

	items := []map[string]types.AttributeValue{}

	queryInput := dynamodb.QueryInput{
		TableName: tableName,
		IndexName: indexName,
		KeyConditions: map[string]types.Condition{
			partitionKeyParam: {
				ComparisonOperator: types.ComparisonOperatorEq,
				AttributeValueList: []types.AttributeValue{
					&types.AttributeValueMemberS{
						Value: partitionKey,
					},
				},
			},
		},
		ExclusiveStartKey: startKey,
	}

	result, err := dbInst.Query(ctx, &queryInput)

	if err != nil {
		return items, err
	}

	items = append(items, result.Items...)

	if len(result.LastEvaluatedKey) > 0 {
		moreItems, err := getAllItemsUsingQuery(
			dbInst, ctx, tableName, partitionKeyParam, partitionKey, result.LastEvaluatedKey, indexName,
		)
		if err != nil {
			return items, err
		}
		items = append(items, moreItems...)
	}

	return items, nil
}

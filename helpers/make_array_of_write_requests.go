package helpers

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func MakeArrayOfWriteRequests[T any](entities []T) ([]types.WriteRequest, error) {
	var itemWrites []types.WriteRequest
	for i := 0; i < len(entities); i++ {

		entityItem, err := attributevalue.MarshalMap(entities[i])

		if err != nil {
			return nil, err
		}

		itemWrites = append(
			itemWrites,
			types.WriteRequest{PutRequest: &types.PutRequest{Item: entityItem}},
		)
	}

	return itemWrites, nil

}

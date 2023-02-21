package email

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/metaphi-org/go-infra-sdk/helpers"
)

// the 'from' email should be authorised in SES for the credentials in config.
func SendBulkEmail(config aws.Config, ctx context.Context, fromEmail string, defaultContent types.BulkEmailContent, bulkEmailEntries []types.BulkEmailEntry) error {

	chunks := helpers.ArrayChunk(bulkEmailEntries, 50)

	for _, chunk := range chunks {
		_, err := sesv2.NewFromConfig(config).SendBulkEmail(ctx, &sesv2.SendBulkEmailInput{
			FromEmailAddress: &fromEmail,
			DefaultContent:   &defaultContent,
			BulkEmailEntries: chunk,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

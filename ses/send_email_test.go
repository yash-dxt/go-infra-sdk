package email_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/metaphi-org/go-infra-sdk/config"
	email "github.com/metaphi-org/go-infra-sdk/ses"
)

func TestSendBulkEmail(t *testing.T) {
	email.SendBulkEmail(config.CreateAWSConfig("us-east-1"), context.Background(), "yashdixitsq@gmail.com", types.BulkEmailContent{
		Template: &types.Template{
			TemplateName: aws.String("Test"),
			TemplateData: aws.String("{ \"name\":\"yash\", \"favouriteanimal\": \"dog\" }"),
		},
	}, []types.BulkEmailEntry{
		{
			Destination: &types.Destination{
				ToAddresses: []string{"yash@metaphi.xyz"},
			},
			ReplacementEmailContent: &types.ReplacementEmailContent{
				ReplacementTemplate: &types.ReplacementTemplate{
					ReplacementTemplateData: aws.String("{ \"name\":\"yash\", \"favouriteanimal\": \"dog\" }"),
				},
			},
		},
		{
			Destination: &types.Destination{
				ToAddresses: []string{"yashdixit.work@gmail.com"},
			},
			ReplacementEmailContent: &types.ReplacementEmailContent{
				ReplacementTemplate: &types.ReplacementTemplate{
					ReplacementTemplateData: aws.String("{ \"name\":\"yash\", \"favouriteanimal\": \"dog\" }"),
				},
			},
		},
	})
}

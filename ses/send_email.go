package email

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

// the 'from' email should be authroised in SES for the credentials in config.
func SendEmail(config aws.Config, from string, to string, subject string, body string, ctx context.Context) (*sesv2.SendEmailOutput, error) {

	return sesv2.NewFromConfig(config).SendEmail(
		ctx,
		&sesv2.SendEmailInput{
			FromEmailAddress: &from,
			Destination: &types.Destination{
				CcAddresses: []string{},
				ToAddresses: []string{to},
			},
			Content: &types.EmailContent{
				Simple: &types.Message{
					Subject: &types.Content{
						Data: &subject,
					},
					Body: &types.Body{
						Text: &types.Content{
							Data: &body,
						},
					},
				},
			},
		},
	)
}

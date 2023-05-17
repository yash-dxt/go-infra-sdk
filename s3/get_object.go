package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// returns get object result in the form of byte array
func GetObject(config aws.Config, ctx context.Context, bucket string, key string) ([]byte, error) {
	s3Client := s3.NewFromConfig(config)

	out, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if err != nil {
		return nil, err
	}

	outBytes, err := io.ReadAll(out.Body)
	return outBytes, err
}

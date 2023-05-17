package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteObject(config aws.Config, ctx context.Context, bucket string, key string) (*s3.DeleteObjectOutput, error) {
	s3Client := s3.NewFromConfig(config)

	return s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
}

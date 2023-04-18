package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetFileMetadata(config aws.Config, ctx context.Context, bucket string, key string) (map[string]string, error) {
	s3Client := s3.NewFromConfig(config)

	out, err := s3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	return out.Metadata, err
}

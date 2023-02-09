package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateGetPresignedUri(config aws.Config, ctx context.Context, bucket string, key string, seconds int) (string, error) {

	s3Client := s3.NewFromConfig(config)

	presignClient := s3.NewPresignClient(s3Client)

	res, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(int64(seconds) * int64(time.Second))
	})

	if err != nil {
		return "", err
	}

	return res.URL, nil
}

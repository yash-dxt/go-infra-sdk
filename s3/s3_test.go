package s3_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/metaphi-org/go-infra-sdk/config"
	"github.com/metaphi-org/go-infra-sdk/s3"

	"github.com/stretchr/testify/assert"
)

var region = os.Getenv("AWS_REGION")
var bucket = os.Getenv("TEST_BUCKET")
var file_upload_path = os.Getenv("FILE_UPLOAD_PATH")
var test_file_key = "test_file_key"

func TestCreatePutPresignedUri(t *testing.T) {
	ctx := context.Background()

	presignedUri, err := s3.CreatePutPresignedUri(config.CreateAWSConfig(region), ctx, bucket, test_file_key, 60*1)

	assert.NotNil(t, presignedUri)
	assert.NoError(t, err)

	file, err := os.Open(file_upload_path)
	assert.NoError(t, err)

	defer file.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, file)
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", presignedUri, &buffer)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	defer resp.Body.Close()

	fmt.Print("\n Status Code: ", resp.StatusCode, "\n")
	assert.Equal(t, true, resp.StatusCode >= 200 && resp.StatusCode < 300)

}

func TestGetPresignedUri(t *testing.T) {
	ctx := context.Background()

	presignedUri, err := s3.CreateGetPresignedUri(config.CreateAWSConfig(region), ctx, bucket, test_file_key, 60*1)

	assert.NotNil(t, presignedUri)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", presignedUri, nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)

	defer resp.Body.Close()

	fmt.Print("\n Status Code: ", resp.StatusCode, "\n")
	assert.Equal(t, resp.StatusCode, 200)

}

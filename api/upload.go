package api

import "github.com/go-resty/resty/v2"

// This can be used for places where you want to upload using a PUT req.
// One example of this is in the s3 signed url upload.
func PUTUpload(client resty.Client, url string, data []byte) error {
	_, err := client.R().
		SetBody(data).
		Put(url)
	return err
}

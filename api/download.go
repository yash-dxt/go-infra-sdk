package api

import "github.com/go-resty/resty/v2"

// download data as bytes.
func Download(client resty.Client, url string) ([]byte, error) {
	resp, err := client.R().Get(url)
	return resp.Body(), err
}

package httpclient

import (
	"bytes"
	"net/http"
	"time"
)

// Get sends a GET request to the specified URL and returns the http.Response
func Get(url string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // You can adjust the timeout as needed
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Post sends a POST request with the given payload and returns the http.Response
func Post(url string, payload []byte) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // You can adjust the timeout as needed
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

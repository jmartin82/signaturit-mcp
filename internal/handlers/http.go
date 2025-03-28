package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	sandboxURL    = "https://api.sandbox.signaturit.com/v3"
	productionURL = "https://api.signaturit.com/v3"
)

// SignaturitClient represents the HTTP client for Signaturit API
type SignaturitClient struct {
	client  *http.Client
	apiKey  string
	baseURL string
	debug   bool
}

// NewSignaturitClient creates a new Signaturit API client
func NewSignaturitClient(apiKey string, debug bool) *SignaturitClient {
	baseURL := productionURL
	if debug {
		baseURL = sandboxURL
	}

	return &SignaturitClient{
		client:  &http.Client{},
		apiKey:  apiKey,
		baseURL: baseURL,
		debug:   debug,
	}
}

// doRequest performs the HTTP request with authentication header
func (c *SignaturitClient) doRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add authentication header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	return resp, nil
}

// Get performs a GET request
func (c *SignaturitClient) Get(endpoint string) (*http.Response, error) {
	return c.doRequest(http.MethodGet, endpoint, nil)
}

// Post performs a POST request
func (c *SignaturitClient) Post(endpoint string, body interface{}) (*http.Response, error) {
	return c.doRequest(http.MethodPost, endpoint, body)
}

// Patch performs a PATCH request
func (c *SignaturitClient) Patch(endpoint string, body interface{}) (*http.Response, error) {
	return c.doRequest(http.MethodPatch, endpoint, body)
}

// Delete performs a DELETE request
func (c *SignaturitClient) Delete(endpoint string) (*http.Response, error) {
	return c.doRequest(http.MethodDelete, endpoint, nil)
}

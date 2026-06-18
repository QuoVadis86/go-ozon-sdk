package ozon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const defaultBaseURL = "https://api-seller.ozon.ru"
const defaultTimeout = 30 * time.Second

type Transport struct {
	client   *http.Client
	baseURL  string
	clientID string
	apiKey   string
}

func NewTransport(clientID, apiKey string, opts *ClientOptions) *Transport {
	baseURL := defaultBaseURL
	timeout := defaultTimeout
	if opts != nil {
		if opts.BaseURL != "" {
			baseURL = opts.BaseURL
		}
		if opts.Timeout > 0 {
			timeout = opts.Timeout
		}
	}
	return &Transport{
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:    20,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout: 60 * time.Second,
			},
		},
		baseURL:  baseURL,
		clientID: clientID,
		apiKey:   apiKey,
	}
}

func (t *Transport) Post(ctx context.Context, path string, reqBody, respBody interface{}) (*http.Response, error) {
	return t.doRequest(ctx, "POST", path, reqBody, respBody)
}

func (t *Transport) Get(ctx context.Context, path string, reqBody, respBody interface{}) (*http.Response, error) {
	return t.doRequest(ctx, "GET", path, reqBody, respBody)
}

func (t *Transport) doRequest(ctx context.Context, method, path string, reqBody, respBody interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if reqBody != nil {
		data, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("ozon: marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	u, err := url.JoinPath(t.baseURL, path)
	if err != nil {
		return nil, fmt.Errorf("ozon: build url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("ozon: create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", t.clientID)
	req.Header.Set("Api-Key", t.apiKey)

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ozon: do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ozon: read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{StatusCode: resp.StatusCode}
		json.Unmarshal(body, apiErr)
		return resp, apiErr
	}

	if respBody != nil && len(body) > 0 {
		if err := json.Unmarshal(body, respBody); err != nil {
			return resp, fmt.Errorf("ozon: unmarshal response: %w", err)
		}
	}

	return resp, nil
}

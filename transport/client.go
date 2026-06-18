package transport

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

const (
	DefaultBaseURL = "https://api-seller.ozon.ru"
	DefaultTimeout = 30 * time.Second

	DefaultMaxIdleConns    = 20
	DefaultMaxIdlePerHost  = 10
	DefaultIdleConnTimeout = 60 * time.Second

	HeaderContentType = "application/json"
	HeaderClientID    = "Client-Id"
	HeaderAPIKey      = "Api-Key"

	// Common error codes (applies to all endpoints)
	ErrCircleIsOpen        = "Circle is open"
	ErrInternal            = "Internal error"
	ErrInvalidAPIKey       = "Invalid Api-Key"
	ErrAPIKeyDeactivated   = "Api-key is deactivated"
	ErrMissingRole         = "Api-Key is missing a required role"
	ErrIPRestricted        = "Api-Key is restricted to specific IP addresses"
	ErrRateLimit           = "You have reached request rate limit per second"
	ErrPriceUpdateLimit    = "error limiting: acquire limit per item"
	ErrMethodNotAllowed    = "method is not allowed"
	ErrOfferNotFound       = "offer_id_not_found"
	ErrWarehouseNotFound   = "WAREHOUSE_NOT_FOUND"
	ErrProductNotCreated   = "product_is_not_created"
	ErrStockTooBig         = "STOCK_TOO_BIG"
	ErrNotFound            = "NOT_FOUND_ERROR"
	ErrTooManyRequests     = "TOO_MANY_REQUESTS"
	ErrInvalidCategoryPrice = "invalid_category_price"
	ErrPriceNegative       = "price_negative"

	// Product import error codes
	ErrSPUAlreadyExists    = "SPU_already_exists"
	ErrInvalidDensity      = "Incorrect_density"
	ErrMissingDimension    = "missing_dimension"
	ErrCategoryInvalid     = "description_category_invalid"
	ErrNameTooLong         = "name_too_long"
	ErrAllImageFailed      = "all_image_failed"
	ErrPriceOutOfRange     = "price_out_of_range"

	// FBS posting error codes
	ErrPostingNotFound      = "POSTING_NOT_FOUND"
	ErrPostingAlreadyShip   = "POSTING_ALREADY_SHIPPED"
	ErrPostingAlreadyCancel = "POSTING_ALREADY_CANCELLED"
	ErrIncorrectStatus      = "HAS_INCORRECT_STATUS"
	ErrTransitionNotPossible = "TRANSITION_IS_NOT_POSSIBLE"
	ErrInvalidCancelReason  = "HAS_INCORRECT_CANCEL_REASON"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	clientID   string
	apiKey     string
}

type Options struct {
	BaseURL    string
	Timeout    time.Duration
	HTTPClient *http.Client
}

func New(clientID, apiKey string, opts *Options) *Client {
	baseURL := DefaultBaseURL
	timeout := DefaultTimeout
	if opts != nil {
		if opts.BaseURL != "" { baseURL = opts.BaseURL }
		if opts.Timeout > 0 { timeout = opts.Timeout }
	}
	hc := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConns:        DefaultMaxIdleConns,
			MaxIdleConnsPerHost: DefaultMaxIdlePerHost,
			IdleConnTimeout:     DefaultIdleConnTimeout,
		},
	}
	if opts != nil && opts.HTTPClient != nil { hc = opts.HTTPClient }
	return &Client{httpClient: hc, baseURL: baseURL, clientID: clientID, apiKey: apiKey}
}

type APIError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Details    []struct {
		TypeURL string `json:"typeUrl"`
		Value   string `json:"value"`
	} `json:"details"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("ozon API error (status=%d, code=%d): %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("ozon API error (status=%d)", e.StatusCode)
}

func (c *Client) Post(ctx context.Context, path string, reqBody, respBody interface{}) error {
	return c.doRequest(ctx, "POST", path, reqBody, respBody)
}

func (c *Client) Get(ctx context.Context, path string, reqBody, respBody interface{}) error {
	return c.doRequest(ctx, "GET", path, reqBody, respBody)
}

func (c *Client) doRequest(ctx context.Context, method, path string, reqBody, respBody interface{}) error {
	var bodyReader io.Reader
	if reqBody != nil {
		data, err := json.Marshal(reqBody)
		if err != nil { return fmt.Errorf("ozon: marshal request: %w", err) }
		bodyReader = bytes.NewReader(data)
	}
	u, err := url.JoinPath(c.baseURL, path)
	if err != nil { return fmt.Errorf("ozon: build url: %w", err) }
	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil { return fmt.Errorf("ozon: create request: %w", err) }
	req.Header.Set("Content-Type", HeaderContentType)
	req.Header.Set(HeaderClientID, c.clientID)
	req.Header.Set(HeaderAPIKey, c.apiKey)
	resp, err := c.httpClient.Do(req)
	if err != nil { return fmt.Errorf("ozon: do request: %w", err) }
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil { return fmt.Errorf("ozon: read response: %w", err) }
	if resp.StatusCode >= 400 {
		apiErr := &APIError{StatusCode: resp.StatusCode}
		json.Unmarshal(body, apiErr)
		return apiErr
	}
	if respBody != nil && len(body) > 0 {
		if err := json.Unmarshal(body, respBody); err != nil {
			return fmt.Errorf("ozon: unmarshal response: %w", err)
		}
	}
	return nil
}

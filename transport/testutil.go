package transport

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// MockHandler creates an http.HandlerFunc that returns a fixed JSON response.
func MockHandler(statusCode int, body interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		if body != nil {
			json.NewEncoder(w).Encode(body)
		}
	}
}

// NewTestClient creates a Client pointed at a test server backed by the given handler.
func NewTestClient(handler http.HandlerFunc) (*Client, *httptest.Server) {
	srv := httptest.NewServer(handler)
	return &Client{
		httpClient: srv.Client(),
		baseURL:    srv.URL,
		clientID:   "test-client",
		apiKey:     "test-key",
	}, srv
}

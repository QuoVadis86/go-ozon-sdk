package premium

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestSearchQueriesTop(t *testing.T) {
	handler := transport.MockHandler(200, V1SearchQueriesTopResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.SearchQueriesTop(ctx, &V1SearchQueriesTopRequest{})
	if err != nil {
		t.Fatalf("SearchQueriesTop() error: %v", err)
	}
	if resp == nil {
		t.Fatal("SearchQueriesTop() returned nil")
	}
}

func TestAPIError(t *testing.T) {
	handler := transport.MockHandler(400, map[string]interface{}{
		"code":    400,
		"message": "test error",
	})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	_, err := svc.SearchQueriesTop(ctx, &V1SearchQueriesTopRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

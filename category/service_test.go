package category

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestSearchAttributeValues(t *testing.T) {
	handler := transport.MockHandler(200, V1SearchAttributeValuesResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.SearchAttributeValues(ctx, &V1SearchAttributeValuesRequest{})
	if err != nil {
		t.Fatalf("SearchAttributeValues() error: %v", err)
	}
	if resp == nil {
		t.Fatal("SearchAttributeValues() returned nil")
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
	_, err := svc.SearchAttributeValues(ctx, &V1SearchAttributeValuesRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

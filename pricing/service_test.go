package pricing

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestItemsDelete(t *testing.T) {
	handler := transport.MockHandler(200, V1DeleteStrategyItemsResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ItemsDelete(ctx, &V1ItemIDsRequest{})
	if err != nil {
		t.Fatalf("ItemsDelete() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ItemsDelete() returned nil")
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
	_, err := svc.ItemsDelete(ctx, &V1ItemIDsRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

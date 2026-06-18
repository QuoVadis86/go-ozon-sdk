package pricing

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestCreate(t *testing.T) {
	handler := transport.MockHandler(200, V1CreatePricingStrategyResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.Create(ctx, &V1CreatePricingStrategyRequest{})
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}
	if resp == nil {
		t.Fatal("Create() returned nil")
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
	_, err := svc.Create(ctx, &V1CreatePricingStrategyRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

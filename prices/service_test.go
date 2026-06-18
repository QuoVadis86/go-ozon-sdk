package prices

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetProductInfoPrices(t *testing.T) {
	handler := transport.MockHandler(200, V5GetProductInfoPricesV5Response{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoPrices(ctx, &V5GetProductInfoPricesV5Request{})
	if err != nil {
		t.Fatalf("GetProductInfoPrices() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetProductInfoPrices() returned nil")
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
	_, err := svc.GetProductInfoPrices(ctx, &V5GetProductInfoPricesV5Request{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

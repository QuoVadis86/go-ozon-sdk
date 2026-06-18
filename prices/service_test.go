package prices

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetProductInfoDiscounted(t *testing.T) {
	handler := transport.MockHandler(200, V1GetProductInfoDiscountedResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoDiscounted(ctx, &V1GetProductInfoDiscountedRequest{})
	if err != nil {
		t.Fatalf("GetProductInfoDiscounted() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetProductInfoDiscounted() returned nil")
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
	_, err := svc.GetProductInfoDiscounted(ctx, &V1GetProductInfoDiscountedRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

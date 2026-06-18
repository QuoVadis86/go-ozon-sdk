package product

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetProductList(t *testing.T) {
	handler := transport.MockHandler(200, V3GetProductListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetProductList(ctx, &V3GetProductListRequest{})
	if err != nil {
		t.Fatalf("GetProductList() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetProductList() returned nil")
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
	_, err := svc.GetProductList(ctx, &V3GetProductListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

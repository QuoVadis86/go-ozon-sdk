package prices

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestProductStocksByWarehouseFbs(t *testing.T) {
	handler := transport.MockHandler(200, Sv1GetProductInfoStocksByWarehouseFbsResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ProductStocksByWarehouseFbs(ctx, &Sv1GetProductInfoStocksByWarehouseFbsRequest{})
	if err != nil {
		t.Fatalf("ProductStocksByWarehouseFbs() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ProductStocksByWarehouseFbs() returned nil")
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
	_, err := svc.ProductStocksByWarehouseFbs(ctx, &Sv1GetProductInfoStocksByWarehouseFbsRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

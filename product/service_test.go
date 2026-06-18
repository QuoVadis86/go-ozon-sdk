package product

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestImportProductsV3(t *testing.T) {
	handler := transport.MockHandler(200, V3ImportProductsResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ImportProductsV3(ctx, &V3ImportProductsRequest{})
	if err != nil {
		t.Fatalf("ImportProductsV3() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ImportProductsV3() returned nil")
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
	_, err := svc.ImportProductsV3(ctx, &V3ImportProductsRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

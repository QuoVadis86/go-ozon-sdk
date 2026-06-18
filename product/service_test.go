package product

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetProductInfoList(t *testing.T) {
	handler := transport.MockHandler(200, V3GetProductInfoListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoList(ctx, &V3GetProductInfoListRequest{})
	if err != nil {
		t.Fatalf("GetProductInfoList() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetProductInfoList() returned nil")
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
	_, err := svc.GetProductInfoList(ctx, &V3GetProductInfoListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

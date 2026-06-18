package fbo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestFbpOrderList(t *testing.T) {
	handler := transport.MockHandler(200, V1FbpOrderListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderList(ctx, &V1FbpOrderListRequest{})
	if err != nil {
		t.Fatalf("FbpOrderList() error: %v", err)
	}
	if resp == nil {
		t.Fatal("FbpOrderList() returned nil")
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
	_, err := svc.FbpOrderList(ctx, &V1FbpOrderListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

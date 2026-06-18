package returns

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestReturnsRfbsListV2(t *testing.T) {
	handler := transport.MockHandler(200, V2ReturnsRfbsListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ReturnsRfbsListV2(ctx, &V2ReturnsRfbsListRequest{})
	if err != nil {
		t.Fatalf("ReturnsRfbsListV2() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ReturnsRfbsListV2() returned nil")
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
	_, err := svc.ReturnsRfbsListV2(ctx, &V2ReturnsRfbsListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

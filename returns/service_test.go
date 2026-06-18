package returns

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestReturnsRfbsGetV2(t *testing.T) {
	handler := transport.MockHandler(200, V2ReturnsRfbsGetResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ReturnsRfbsGetV2(ctx, &V2ReturnsRfbsGetRequest{})
	if err != nil {
		t.Fatalf("ReturnsRfbsGetV2() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ReturnsRfbsGetV2() returned nil")
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
	_, err := svc.ReturnsRfbsGetV2(ctx, &V2ReturnsRfbsGetRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

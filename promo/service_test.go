package promo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestPromosCandidates(t *testing.T) {
	handler := transport.MockHandler(200, SellerApiGetSellerProductV1Response{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.PromosCandidates(ctx, &SellerApiGetSellerProductV1Request{})
	if err != nil {
		t.Fatalf("PromosCandidates() error: %v", err)
	}
	if resp == nil {
		t.Fatal("PromosCandidates() returned nil")
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
	_, err := svc.PromosCandidates(ctx, &SellerApiGetSellerProductV1Request{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

package promo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestSellerActionsCreateMultiLevelDiscount(t *testing.T) {
	handler := transport.MockHandler(200, V1SellerActionsCreateMultiLevelDiscountResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsCreateMultiLevelDiscount(ctx, &V1SellerActionsCreateMultiLevelDiscountRequest{})
	if err != nil {
		t.Fatalf("SellerActionsCreateMultiLevelDiscount() error: %v", err)
	}
	if resp == nil {
		t.Fatal("SellerActionsCreateMultiLevelDiscount() returned nil")
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
	_, err := svc.SellerActionsCreateMultiLevelDiscount(ctx, &V1SellerActionsCreateMultiLevelDiscountRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

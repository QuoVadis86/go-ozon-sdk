package promo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestActionsAutoAddProductsList(t *testing.T) {
	handler := transport.MockHandler(200, ActionsV1ActionsAutoAddProductsListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsList(ctx, &ActionsV1ActionsAutoAddProductsListRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsList() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ActionsAutoAddProductsList() returned nil")
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
	_, err := svc.ActionsAutoAddProductsList(ctx, &ActionsV1ActionsAutoAddProductsListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

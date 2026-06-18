package fbs

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestCarriageCreate(t *testing.T) {
	handler := transport.MockHandler(200, V1CarriageCreateResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.CarriageCreate(ctx, &V1CarriageCreateRequest{})
	if err != nil {
		t.Fatalf("CarriageCreate() error: %v", err)
	}
	if resp == nil {
		t.Fatal("CarriageCreate() returned nil")
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
	_, err := svc.CarriageCreate(ctx, &V1CarriageCreateRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

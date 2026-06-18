package category

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetTree(t *testing.T) {
	handler := transport.MockHandler(200, V1GetTreeResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetTree(ctx, &V1GetTreeRequest{})
	if err != nil {
		t.Fatalf("GetTree() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetTree() returned nil")
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
	_, err := svc.GetTree(ctx, &V1GetTreeRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

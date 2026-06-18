package category

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetAttributes(t *testing.T) {
	handler := transport.MockHandler(200, V1GetAttributesResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetAttributes(ctx, &V1GetAttributesRequest{})
	if err != nil {
		t.Fatalf("GetAttributes() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetAttributes() returned nil")
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
	_, err := svc.GetAttributes(ctx, &V1GetAttributesRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

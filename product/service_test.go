package product

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestImportProductsBySKU(t *testing.T) {
	handler := transport.MockHandler(200, ImportProductsBySKUResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ImportProductsBySKU(ctx, &ImportProductsBySKURequest{})
	if err != nil {
		t.Fatalf("ImportProductsBySKU() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ImportProductsBySKU() returned nil")
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
	_, err := svc.ImportProductsBySKU(ctx, &ImportProductsBySKURequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

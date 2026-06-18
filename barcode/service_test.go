package barcode

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGenerateBarcode(t *testing.T) {
	handler := transport.MockHandler(200, V1GenerateBarcodeResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GenerateBarcode(ctx, &V1GenerateBarcodeRequest{})
	if err != nil {
		t.Fatalf("GenerateBarcode() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GenerateBarcode() returned nil")
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
	_, err := svc.GenerateBarcode(ctx, &V1GenerateBarcodeRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

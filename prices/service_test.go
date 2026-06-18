package prices

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestImportProductsPrices(t *testing.T) {
	handler := transport.MockHandler(200, ImportProductsPricesResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ImportProductsPrices(ctx, &ImportProductsPricesRequest{})
	if err != nil {
		t.Fatalf("ImportProductsPrices() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ImportProductsPrices() returned nil")
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
	_, err := svc.ImportProductsPrices(ctx, &ImportProductsPricesRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

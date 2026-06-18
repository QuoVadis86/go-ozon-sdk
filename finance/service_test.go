package finance

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestFinanceTransactionListV3(t *testing.T) {
	handler := transport.MockHandler(200, V3FinanceTransactionListV3Response{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.FinanceTransactionListV3(ctx, &V3FinanceTransactionListV3Request{})
	if err != nil {
		t.Fatalf("FinanceTransactionListV3() error: %v", err)
	}
	if resp == nil {
		t.Fatal("FinanceTransactionListV3() returned nil")
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
	_, err := svc.FinanceTransactionListV3(ctx, &V3FinanceTransactionListV3Request{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

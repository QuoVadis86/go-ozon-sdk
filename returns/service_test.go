package returns

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetConditionalCancellationListV2(t *testing.T) {
	handler := transport.MockHandler(200, V2GetConditionalCancellationListV2Response{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetConditionalCancellationListV2(ctx, &V2GetConditionalCancellationListV2Request{})
	if err != nil {
		t.Fatalf("GetConditionalCancellationListV2() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetConditionalCancellationListV2() returned nil")
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
	_, err := svc.GetConditionalCancellationListV2(ctx, &V2GetConditionalCancellationListV2Request{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

package warehouse

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestDeliveryMethodListV2(t *testing.T) {
	handler := transport.MockHandler(200, V2DeliveryMethodListV2Response{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.DeliveryMethodListV2(ctx, &V2DeliveryMethodListV2Request{})
	if err != nil {
		t.Fatalf("DeliveryMethodListV2() error: %v", err)
	}
	if resp == nil {
		t.Fatal("DeliveryMethodListV2() returned nil")
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
	_, err := svc.DeliveryMethodListV2(ctx, &V2DeliveryMethodListV2Request{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

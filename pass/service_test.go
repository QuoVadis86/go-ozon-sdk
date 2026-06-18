package pass

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestReturnPassCreate(t *testing.T) {
	handler := transport.MockHandler(200, ArrivalpassArrivalPassCreateResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ReturnPassCreate(ctx, &ArrivalpassArrivalPassCreateRequest{})
	if err != nil {
		t.Fatalf("ReturnPassCreate() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ReturnPassCreate() returned nil")
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
	_, err := svc.ReturnPassCreate(ctx, &ArrivalpassArrivalPassCreateRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

package pass

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestPassList(t *testing.T) {
	handler := transport.MockHandler(200, ArrivalpassArrivalPassListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.PassList(ctx, &ArrivalpassArrivalPassListRequest{})
	if err != nil {
		t.Fatalf("PassList() error: %v", err)
	}
	if resp == nil {
		t.Fatal("PassList() returned nil")
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
	_, err := svc.PassList(ctx, &ArrivalpassArrivalPassListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

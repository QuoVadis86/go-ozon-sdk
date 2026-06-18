package promo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestTaskDecline(t *testing.T) {
	handler := transport.MockHandler(200, V1ApproveDeclineDiscountTasksResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.TaskDecline(ctx, &V1DeclineDiscountTasksRequest{})
	if err != nil {
		t.Fatalf("TaskDecline() error: %v", err)
	}
	if resp == nil {
		t.Fatal("TaskDecline() returned nil")
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
	_, err := svc.TaskDecline(ctx, &V1DeclineDiscountTasksRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

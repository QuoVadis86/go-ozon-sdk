package chat

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestChatReadV2(t *testing.T) {
	handler := transport.MockHandler(200, V2ChatReadResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ChatReadV2(ctx, &Read{})
	if err != nil {
		t.Fatalf("ChatReadV2() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ChatReadV2() returned nil")
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
	_, err := svc.ChatReadV2(ctx, &Read{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

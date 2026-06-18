package chat

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestChatSendMessage(t *testing.T) {
	handler := transport.MockHandler(200, ChatSendMessageResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ChatSendMessage(ctx, &ChatSendMessageRequest{})
	if err != nil {
		t.Fatalf("ChatSendMessage() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ChatSendMessage() returned nil")
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
	_, err := svc.ChatSendMessage(ctx, &ChatSendMessageRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

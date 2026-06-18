package chat

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestChatSendFile(t *testing.T) {
	handler := transport.MockHandler(200, ChatSendFileResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ChatSendFile(ctx, &ChatSendFileRequest{})
	if err != nil {
		t.Fatalf("ChatSendFile() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ChatSendFile() returned nil")
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
	_, err := svc.ChatSendFile(ctx, &ChatSendFileRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

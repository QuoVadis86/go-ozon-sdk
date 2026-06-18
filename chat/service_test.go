package chat

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"os"
	"testing"
)

var ctx = context.Background()

func skipNoCreds(t *testing.T) *transport.Client {
	t.Helper()
	if os.Getenv("OZON_CLIENT_ID") == "" || os.Getenv("OZON_API_KEY") == "" {
		t.Skip("set OZON_CLIENT_ID and OZON_API_KEY to run tests")
	}
	return transport.New(os.Getenv("OZON_CLIENT_ID"), os.Getenv("OZON_API_KEY"), nil)
}

func TestChatSendFile(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ChatSendFile(ctx, &ChatSendFileRequest{})
	if err != nil {
		t.Fatalf("ChatSendFile() error: %v", err)
	}
	_ = resp
}

func TestChatStart(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ChatStart(ctx, &ChatStartRequest{})
	if err != nil {
		t.Fatalf("ChatStart() error: %v", err)
	}
	_ = resp
}

func TestChatListV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ChatListV3(ctx, &V3Chat{})
	if err != nil {
		t.Fatalf("ChatListV3() error: %v", err)
	}
	_ = resp
}

func TestChatSendMessage(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ChatSendMessage(ctx, &ChatSendMessageRequest{})
	if err != nil {
		t.Fatalf("ChatSendMessage() error: %v", err)
	}
	_ = resp
}

func TestChatReadV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ChatReadV2(ctx, &Read{})
	if err != nil {
		t.Fatalf("ChatReadV2() error: %v", err)
	}
	_ = resp
}

func TestChatHistoryV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ChatHistoryV3(ctx, &V3ChatHistoryRequest{})
	if err != nil {
		t.Fatalf("ChatHistoryV3() error: %v", err)
	}
	_ = resp
}

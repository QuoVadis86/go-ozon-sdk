package review

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestInfo(t *testing.T) {
	handler := transport.MockHandler(200, V1QuestionInfoResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.Info(ctx, &V1QuestionInfoRequest{})
	if err != nil {
		t.Fatalf("Info() error: %v", err)
	}
	if resp == nil {
		t.Fatal("Info() returned nil")
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
	_, err := svc.Info(ctx, &V1QuestionInfoRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

package review

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestList(t *testing.T) {
	handler := transport.MockHandler(200, V1QuestionAnswerListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.List(ctx, &V1QuestionAnswerListRequest{})
	if err != nil {
		t.Fatalf("List() error: %v", err)
	}
	if resp == nil {
		t.Fatal("List() returned nil")
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
	_, err := svc.List(ctx, &V1QuestionAnswerListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

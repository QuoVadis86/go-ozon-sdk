package review

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestCommentList(t *testing.T) {
	handler := transport.MockHandler(200, V1CommentListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.CommentList(ctx, &V1CommentListRequest{})
	if err != nil {
		t.Fatalf("CommentList() error: %v", err)
	}
	if resp == nil {
		t.Fatal("CommentList() returned nil")
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
	_, err := svc.CommentList(ctx, &V1CommentListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

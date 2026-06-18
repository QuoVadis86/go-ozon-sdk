package fbs

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestPostingFbsUnfulfilledList(t *testing.T) {
	handler := transport.MockHandler(200, PostingV4PostingFbsUnfulfilledListResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.PostingFbsUnfulfilledList(ctx, &PostingV4PostingFbsUnfulfilledListRequest{})
	if err != nil {
		t.Fatalf("PostingFbsUnfulfilledList() error: %v", err)
	}
	if resp == nil {
		t.Fatal("PostingFbsUnfulfilledList() returned nil")
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
	_, err := svc.PostingFbsUnfulfilledList(ctx, &PostingV4PostingFbsUnfulfilledListRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

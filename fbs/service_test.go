package fbs

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestFbsPostingTrackingNumberSet(t *testing.T) {
	handler := transport.MockHandler(200, PostingFbsPostingMoveStatusResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingTrackingNumberSet(ctx, &PostingFbsPostingTrackingNumberSetRequest{})
	if err != nil {
		t.Fatalf("FbsPostingTrackingNumberSet() error: %v", err)
	}
	if resp == nil {
		t.Fatal("FbsPostingTrackingNumberSet() returned nil")
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
	_, err := svc.FbsPostingTrackingNumberSet(ctx, &PostingFbsPostingTrackingNumberSetRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

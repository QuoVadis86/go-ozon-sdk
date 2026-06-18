package fbo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestFbpDraftDropOffDelete(t *testing.T) {
	handler := transport.MockHandler(200, V1FbpDraftDropOffDeleteResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffDelete(ctx, &V1FbpDraftDropOffDeleteRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffDelete() error: %v", err)
	}
	if resp == nil {
		t.Fatal("FbpDraftDropOffDelete() returned nil")
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
	_, err := svc.FbpDraftDropOffDelete(ctx, &V1FbpDraftDropOffDeleteRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

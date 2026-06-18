package fbo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestFbpDraftDirectDelete(t *testing.T) {
	handler := transport.MockHandler(200, V1FbpDraftDirectDeleteResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectDelete(ctx, &V1FbpDraftDirectDeleteRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectDelete() error: %v", err)
	}
	if resp == nil {
		t.Fatal("FbpDraftDirectDelete() returned nil")
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
	_, err := svc.FbpDraftDirectDelete(ctx, &V1FbpDraftDirectDeleteRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

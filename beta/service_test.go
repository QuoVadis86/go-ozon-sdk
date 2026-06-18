package beta

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestFbsPostingProductExemplarStatusV5(t *testing.T) {
	handler := transport.MockHandler(200, V5FbsPostingProductExemplarStatusV5Response{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingProductExemplarStatusV5(ctx, &V5FbsPostingProductExemplarStatusV5Request{})
	if err != nil {
		t.Fatalf("FbsPostingProductExemplarStatusV5() error: %v", err)
	}
	if resp == nil {
		t.Fatal("FbsPostingProductExemplarStatusV5() returned nil")
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
	_, err := svc.FbsPostingProductExemplarStatusV5(ctx, &V5FbsPostingProductExemplarStatusV5Request{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

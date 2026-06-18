package rating

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetFBSRatingIndexInfoV1(t *testing.T) {
	handler := transport.MockHandler(200, V1GetFBSRatingIndexInfoV1Response{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetFBSRatingIndexInfoV1(ctx)
	if err != nil {
		t.Fatalf("GetFBSRatingIndexInfoV1() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetFBSRatingIndexInfoV1() returned nil")
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
	_, err := svc.GetFBSRatingIndexInfoV1(ctx)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

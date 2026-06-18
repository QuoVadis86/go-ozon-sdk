package pass

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestReturnsCompanyFBSInfo(t *testing.T) {
	handler := transport.MockHandler(200, V1ReturnsCompanyFbsInfoResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ReturnsCompanyFBSInfo(ctx, &V1ReturnsCompanyFbsInfoRequest{})
	if err != nil {
		t.Fatalf("ReturnsCompanyFBSInfo() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ReturnsCompanyFBSInfo() returned nil")
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
	_, err := svc.ReturnsCompanyFBSInfo(ctx, &V1ReturnsCompanyFbsInfoRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

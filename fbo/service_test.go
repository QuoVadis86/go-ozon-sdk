package fbo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"os"
	"testing"
)

var ctx = context.Background()

func skipNoCreds(t *testing.T) *transport.Client {
	t.Helper()
	if os.Getenv("OZON_CLIENT_ID") == "" || os.Getenv("OZON_API_KEY") == "" {
		t.Skip("set OZON_CLIENT_ID and OZON_API_KEY to run tests")
	}
	return transport.New(os.Getenv("OZON_CLIENT_ID"), os.Getenv("OZON_API_KEY"), nil)
}

func TestFbpOrderDirectSellerDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderDirectSellerDlvEdit(ctx, &V1FbpOrderDirectSellerDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpOrderDirectSellerDlvEdit() error: %v", err)
	}
	if resp == nil {
		t.Fatal("FbpOrderDirectSellerDlvEdit() returned nil")
	}
}

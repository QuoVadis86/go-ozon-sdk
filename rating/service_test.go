package rating

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

func TestGetFBSRatingIndexInfoV1(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFBSRatingIndexInfoV1(ctx)
	if err != nil {
		t.Fatalf("GetFBSRatingIndexInfoV1() error: %v", err)
	}
	_ = resp
}

func TestListFBSRatingIndexPostingsV1(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ListFBSRatingIndexPostingsV1(ctx, &V1ListFBSRatingIndexPostingsV1Request{})
	if err != nil {
		t.Fatalf("ListFBSRatingIndexPostingsV1() error: %v", err)
	}
	_ = resp
}

package returns

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

func TestReturnsRfbsGetV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReturnsRfbsGetV2(ctx, &V2ReturnsRfbsGetRequest{})
	if err != nil {
		t.Fatalf("ReturnsRfbsGetV2() error: %v", err)
	}
	_ = resp
}

func TestReturnsRfbsListV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReturnsRfbsListV2(ctx, &V2ReturnsRfbsListRequest{})
	if err != nil {
		t.Fatalf("ReturnsRfbsListV2() error: %v", err)
	}
	_ = resp
}

func TestConditionalCancellationApproveV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ConditionalCancellationApproveV2(ctx)
	_ = err
}

func TestConditionalCancellationRejectV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ConditionalCancellationRejectV2(ctx)
	_ = err
}

func TestReturnsRfbsActionSet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ReturnsRfbsActionSet(ctx, &V1ReturnsRfbsActionSetRequest{})
	_ = err
}

func TestGetConditionalCancellationListV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetConditionalCancellationListV2(ctx, &V2GetConditionalCancellationListV2Request{})
	if err != nil {
		t.Fatalf("GetConditionalCancellationListV2() error: %v", err)
	}
	_ = resp
}

func TestReturnsList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReturnsList(ctx, &V1GetReturnsListRequest{})
	if err != nil {
		t.Fatalf("ReturnsList() error: %v", err)
	}
	_ = resp
}

package pass

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

func TestReturnPassCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReturnPassCreate(ctx, &ArrivalpassArrivalPassCreateRequest{})
	if err != nil {
		t.Fatalf("ReturnPassCreate() error: %v", err)
	}
	_ = resp
}

func TestReturnPassUpdate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ReturnPassUpdate(ctx, &ArrivalpassArrivalPassUpdateRequest{})
	_ = err
}

func TestCarriagePassUpdate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.CarriagePassUpdate(ctx, &SellerAPIArrivalPassUpdateRequest{})
	_ = err
}

func TestReturnPassDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ReturnPassDelete(ctx, &ArrivalpassArrivalPassDeleteRequest{})
	_ = err
}

func TestPassList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PassList(ctx, &ArrivalpassArrivalPassListRequest{})
	if err != nil {
		t.Fatalf("PassList() error: %v", err)
	}
	_ = resp
}

func TestReturnsCompanyFBSInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReturnsCompanyFBSInfo(ctx, &V1ReturnsCompanyFbsInfoRequest{})
	if err != nil {
		t.Fatalf("ReturnsCompanyFBSInfo() error: %v", err)
	}
	_ = resp
}

func TestCarriagePassCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CarriagePassCreate(ctx, &SellerAPIArrivalPassCreateRequest{})
	if err != nil {
		t.Fatalf("CarriagePassCreate() error: %v", err)
	}
	_ = resp
}

func TestCarriagePassDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.CarriagePassDelete(ctx, &SellerAPIArrivalPassDeleteRequest{})
	_ = err
}

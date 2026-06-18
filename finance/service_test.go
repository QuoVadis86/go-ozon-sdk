package finance

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

func TestFinanceTransactionTotalV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FinanceTransactionTotalV3(ctx, &V3FinanceTransactionTotalsV3Request{})
	if err != nil {
		t.Fatalf("FinanceTransactionTotalV3() error: %v", err)
	}
	_ = resp
}

func TestGetRealizationReportV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetRealizationReportV2(ctx, &V2GetRealizationReportRequestV2{})
	if err != nil {
		t.Fatalf("GetRealizationReportV2() error: %v", err)
	}
	_ = resp
}

func TestGetRealizationReportV1(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetRealizationReportV1(ctx, &V1GetRealizationReportPostingRequest{})
	if err != nil {
		t.Fatalf("GetRealizationReportV1() error: %v", err)
	}
	_ = resp
}

func TestGetCompensationReport(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetCompensationReport(ctx, &V1GetCompensationReportRequest{})
	if err != nil {
		t.Fatalf("GetCompensationReport() error: %v", err)
	}
	_ = resp
}

func TestFinanceTransactionListV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FinanceTransactionListV3(ctx, &V3FinanceTransactionListV3Request{})
	if err != nil {
		t.Fatalf("FinanceTransactionListV3() error: %v", err)
	}
	_ = resp
}

func TestGetDecompensationReport(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetDecompensationReport(ctx, &V1GetDecompensationReportRequest{})
	if err != nil {
		t.Fatalf("GetDecompensationReport() error: %v", err)
	}
	_ = resp
}

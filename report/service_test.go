package report

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

func TestCreateCompanyMarkedProductsSalesReport(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CreateCompanyMarkedProductsSalesReport(ctx, &V1ReportMarkedProductsSalesCreateRequest{})
	if err != nil {
		t.Fatalf("CreateCompanyMarkedProductsSalesReport() error: %v", err)
	}
	_ = resp
}

func TestCreateCompanyProductsReport(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CreateCompanyProductsReport(ctx, &CreateCompanyProductsReportRequest{})
	if err != nil {
		t.Fatalf("CreateCompanyProductsReport() error: %v", err)
	}
	_ = resp
}

func TestReportList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReportList(ctx, &ReportListRequest{})
	if err != nil {
		t.Fatalf("ReportList() error: %v", err)
	}
	_ = resp
}

func TestFinanceCashFlowStatementList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FinanceCashFlowStatementList(ctx, &V3FinanceCashFlowStatementListRequest{})
	if err != nil {
		t.Fatalf("FinanceCashFlowStatementList() error: %v", err)
	}
	_ = resp
}

func TestCreateCompanyPostingsReport(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CreateCompanyPostingsReport(ctx, &CreateCompanyPostingsReportRequest{})
	if err != nil {
		t.Fatalf("CreateCompanyPostingsReport() error: %v", err)
	}
	_ = resp
}

func TestCreateStockByWarehouseReport(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CreateStockByWarehouseReport(ctx, &V1CreateStockByWarehouseReportRequest{})
	if err != nil {
		t.Fatalf("CreateStockByWarehouseReport() error: %v", err)
	}
	_ = resp
}

func TestReportInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ReportInfo(ctx, &ReportInfoRequest{})
	if err != nil {
		t.Fatalf("ReportInfo() error: %v", err)
	}
	_ = resp
}

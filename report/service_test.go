package report

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestCreateCompanyMarkedProductsSalesReport(t *testing.T) {
	handler := transport.MockHandler(200, CommonCreateReportResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.CreateCompanyMarkedProductsSalesReport(ctx, &V1ReportMarkedProductsSalesCreateRequest{})
	if err != nil {
		t.Fatalf("CreateCompanyMarkedProductsSalesReport() error: %v", err)
	}
	if resp == nil {
		t.Fatal("CreateCompanyMarkedProductsSalesReport() returned nil")
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
	_, err := svc.CreateCompanyMarkedProductsSalesReport(ctx, &V1ReportMarkedProductsSalesCreateRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

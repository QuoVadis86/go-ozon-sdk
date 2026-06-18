package report

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestCreateCompanyProductsReport(t *testing.T) {
	handler := transport.MockHandler(200, CreateReportResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.CreateCompanyProductsReport(ctx, &CreateCompanyProductsReportRequest{})
	if err != nil {
		t.Fatalf("CreateCompanyProductsReport() error: %v", err)
	}
	if resp == nil {
		t.Fatal("CreateCompanyProductsReport() returned nil")
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
	_, err := svc.CreateCompanyProductsReport(ctx, &CreateCompanyProductsReportRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

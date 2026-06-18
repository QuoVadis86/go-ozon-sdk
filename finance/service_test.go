package finance

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestGetDecompensationReport(t *testing.T) {
	handler := transport.MockHandler(200, CreateReportResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.GetDecompensationReport(ctx, &V1GetDecompensationReportRequest{})
	if err != nil {
		t.Fatalf("GetDecompensationReport() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetDecompensationReport() returned nil")
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
	_, err := svc.GetDecompensationReport(ctx, &V1GetDecompensationReportRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

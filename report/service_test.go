package report

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"testing"
)

var ctx = context.Background()

func TestReportInfo(t *testing.T) {
	handler := transport.MockHandler(200, ReportInfoResponse{})
	cl, srv := transport.NewTestClient(handler)
	defer srv.Close()
	svc := &Service{Client: cl}
	resp, err := svc.ReportInfo(ctx, &ReportInfoRequest{})
	if err != nil {
		t.Fatalf("ReportInfo() error: %v", err)
	}
	if resp == nil {
		t.Fatal("ReportInfo() returned nil")
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
	_, err := svc.ReportInfo(ctx, &ReportInfoRequest{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

package premium

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

func TestGetRealizationByDayReportV1(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetRealizationByDayReportV1(ctx, &V1GetRealizationReportByDayRequest{})
	if err != nil {
		t.Fatalf("GetRealizationByDayReportV1() error: %v", err)
	}
	_ = resp
}

func TestAnalyticsProductQueriesDetails(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.AnalyticsProductQueriesDetails(ctx, &V1AnalyticsProductQueriesDetailsRequest{})
	if err != nil {
		t.Fatalf("AnalyticsProductQueriesDetails() error: %v", err)
	}
	_ = resp
}

func TestAnalyticsProductQueries(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.AnalyticsProductQueries(ctx, &V1AnalyticsProductQueriesRequest{})
	if err != nil {
		t.Fatalf("AnalyticsProductQueries() error: %v", err)
	}
	_ = resp
}

func TestAnalyticsGetData(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.AnalyticsGetData(ctx, &AnalyticsAnalyticsGetDataRequest{})
	if err != nil {
		t.Fatalf("AnalyticsGetData() error: %v", err)
	}
	_ = resp
}

func TestSearchQueriesTop(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SearchQueriesTop(ctx, &V1SearchQueriesTopRequest{})
	if err != nil {
		t.Fatalf("SearchQueriesTop() error: %v", err)
	}
	_ = resp
}

func TestProductPricesDetails(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductPricesDetails(ctx, &V1ProductPricesDetailsRequest{})
	if err != nil {
		t.Fatalf("ProductPricesDetails() error: %v", err)
	}
	_ = resp
}

func TestSearchQueriesText(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SearchQueriesText(ctx, &V1SearchQueriesTextRequest{})
	if err != nil {
		t.Fatalf("SearchQueriesText() error: %v", err)
	}
	_ = resp
}

package pricing

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

func TestItemsAdd(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ItemsAdd(ctx, &V1AddStrategyItemsRequest{})
	if err != nil {
		t.Fatalf("ItemsAdd() error: %v", err)
	}
	_ = resp
}

func TestItemsDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ItemsDelete(ctx, &V1ItemIDsRequest{})
	if err != nil {
		t.Fatalf("ItemsDelete() error: %v", err)
	}
	_ = resp
}

func TestList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.List(ctx, &V1GetStrategyListRequest{})
	if err != nil {
		t.Fatalf("List() error: %v", err)
	}
	_ = resp
}

func TestCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Create(ctx, &V1CreatePricingStrategyRequest{})
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}
	_ = resp
}

func TestItemsInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ItemsInfo(ctx, &V1GetStrategyItemInfoRequest{})
	if err != nil {
		t.Fatalf("ItemsInfo() error: %v", err)
	}
	_ = resp
}

func TestIds(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Ids(ctx, &V1ItemIDsRequest{})
	if err != nil {
		t.Fatalf("Ids() error: %v", err)
	}
	_ = resp
}

func TestCompetitors(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Competitors(ctx, &V1GetCompetitorsRequest{})
	if err != nil {
		t.Fatalf("Competitors() error: %v", err)
	}
	_ = resp
}

func TestItemsList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ItemsList(ctx, &V1StrategyRequest{})
	if err != nil {
		t.Fatalf("ItemsList() error: %v", err)
	}
	_ = resp
}

func TestInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Info(ctx, &V1StrategyRequest{})
	if err != nil {
		t.Fatalf("Info() error: %v", err)
	}
	_ = resp
}

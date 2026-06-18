package prices

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

func TestGetProductInfoStocksByWarehouseFbsV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoStocksByWarehouseFbsV2(ctx, &V2GetProductInfoStocksByWarehouseFbsRequestV2{})
	if err != nil {
		t.Fatalf("GetProductInfoStocksByWarehouseFbsV2() error: %v", err)
	}
	_ = resp
}

func TestProductUpdateDiscount(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductUpdateDiscount(ctx, &V1ProductUpdateDiscountRequest{})
	if err != nil {
		t.Fatalf("ProductUpdateDiscount() error: %v", err)
	}
	_ = resp
}

func TestGetProductInfoStocks(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoStocks(ctx, &V4GetProductInfoStocksRequest{})
	if err != nil {
		t.Fatalf("GetProductInfoStocks() error: %v", err)
	}
	_ = resp
}

func TestActionTimerUpdate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.ActionTimerUpdate(ctx, &V1ProductActionTimerUpdateRequest{})
	_ = err
}

func TestGetProductInfoPrices(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoPrices(ctx, &V5GetProductInfoPricesV5Request{})
	if err != nil {
		t.Fatalf("GetProductInfoPrices() error: %v", err)
	}
	_ = resp
}

func TestProductStocksByWarehouseFbs(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductStocksByWarehouseFbs(ctx, &Sv1GetProductInfoStocksByWarehouseFbsRequest{})
	if err != nil {
		t.Fatalf("ProductStocksByWarehouseFbs() error: %v", err)
	}
	_ = resp
}

func TestProductsStocksV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductsStocksV2(ctx, &V2ProductsStocksRequest{})
	if err != nil {
		t.Fatalf("ProductsStocksV2() error: %v", err)
	}
	_ = resp
}

func TestImportProductsPrices(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ImportProductsPrices(ctx, &ImportProductsPricesRequest{})
	if err != nil {
		t.Fatalf("ImportProductsPrices() error: %v", err)
	}
	_ = resp
}

func TestGetProductInfoDiscounted(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoDiscounted(ctx, &V1GetProductInfoDiscountedRequest{})
	if err != nil {
		t.Fatalf("GetProductInfoDiscounted() error: %v", err)
	}
	_ = resp
}

func TestActionTimerStatus(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionTimerStatus(ctx, &V1ProductActionTimerStatusRequest{})
	if err != nil {
		t.Fatalf("ActionTimerStatus() error: %v", err)
	}
	_ = resp
}

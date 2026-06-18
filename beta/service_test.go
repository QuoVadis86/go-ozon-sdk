package beta

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

func TestFbsPostingProductExemplarStatusV5(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingProductExemplarStatusV5(ctx, &V5FbsPostingProductExemplarStatusV5Request{})
	if err != nil {
		t.Fatalf("FbsPostingProductExemplarStatusV5() error: %v", err)
	}
	_ = resp
}

func TestGetFinanceAccrualTypes(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFinanceAccrualTypes(ctx)
	if err != nil {
		t.Fatalf("GetFinanceAccrualTypes() error: %v", err)
	}
	_ = resp
}

func TestProductVisibilityInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductVisibilityInfo(ctx, &V1ProductVisibilityInfoRequest{})
	if err != nil {
		t.Fatalf("ProductVisibilityInfo() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingProductExemplarValidateV5(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingProductExemplarValidateV5(ctx, &V5FbsPostingProductExemplarValidateV5Request{})
	if err != nil {
		t.Fatalf("FbsPostingProductExemplarValidateV5() error: %v", err)
	}
	_ = resp
}

func TestGetDiscountTaskListV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetDiscountTaskListV2(ctx, &V2GetDiscountTaskListV2Request{})
	if err != nil {
		t.Fatalf("GetDiscountTaskListV2() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingProductExemplarSetV6(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.FbsPostingProductExemplarSetV6(ctx, &V6FbsPostingProductExemplarSetV6Request{})
	_ = err
}

func TestProductInfoWarehouseStocks(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductInfoWarehouseStocks(ctx, &V1ProductInfoWarehouseStocksRequest{})
	if err != nil {
		t.Fatalf("ProductInfoWarehouseStocks() error: %v", err)
	}
	_ = resp
}

func TestGetFinanceAccrualByDay(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFinanceAccrualByDay(ctx, &V1GetFinanceAccrualByDayRequest{})
	if err != nil {
		t.Fatalf("GetFinanceAccrualByDay() error: %v", err)
	}
	_ = resp
}

func TestSetProductStairwayDiscountByQuantity(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SetProductStairwayDiscountByQuantity(ctx, &V1SetProductStairwayDiscountByQuantityRequest{})
	if err != nil {
		t.Fatalf("SetProductStairwayDiscountByQuantity() error: %v", err)
	}
	_ = resp
}

func TestGetFinanceBalanceV1(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFinanceBalanceV1(ctx, &V1GetFinanceBalanceV1Request{})
	if err != nil {
		t.Fatalf("GetFinanceBalanceV1() error: %v", err)
	}
	_ = resp
}

func TestDescriptionCategoryTips(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.DescriptionCategoryTips(ctx, &V1DescriptionCategoryTipsRequest{})
	if err != nil {
		t.Fatalf("DescriptionCategoryTips() error: %v", err)
	}
	_ = resp
}

func TestFbsSplit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsSplit(ctx)
	if err != nil {
		t.Fatalf("FbsSplit() error: %v", err)
	}
	_ = resp
}

func TestGetProductStairwayDiscountByQuantity(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductStairwayDiscountByQuantity(ctx, &V1GetProductStairwayDiscountByQuantityRequest{})
	if err != nil {
		t.Fatalf("GetProductStairwayDiscountByQuantity() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingProductExemplarCreateOrGetV6(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingProductExemplarCreateOrGetV6(ctx, &V6FbsPostingProductExemplarCreateOrGetV6Request{})
	if err != nil {
		t.Fatalf("FbsPostingProductExemplarCreateOrGetV6() error: %v", err)
	}
	_ = resp
}

func TestProductVisibilitySet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductVisibilitySet(ctx, &V1ProductVisibilitySetRequest{})
	if err != nil {
		t.Fatalf("ProductVisibilitySet() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingProductExemplarUpdate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.FbsPostingProductExemplarUpdate(ctx, &V1FbsPostingProductExemplarUpdateRequest{})
	_ = err
}

func TestGetFinanceAccrualPostings(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFinanceAccrualPostings(ctx, &V1GetFinanceAccrualPostingsRequest{})
	if err != nil {
		t.Fatalf("GetFinanceAccrualPostings() error: %v", err)
	}
	_ = resp
}

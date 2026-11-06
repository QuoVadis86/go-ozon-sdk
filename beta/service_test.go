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

func TestSetProductStairwayDiscountByQuantity(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SetProductStairwayDiscountByQuantity(ctx, &V1SetProductStairwayDiscountByQuantityRequest{})
	if err != nil {
		t.Fatalf("SetProductStairwayDiscountByQuantity() error: %v", err)
	}
	_ = resp
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

func TestGetProductStairwayDiscountByQuantity(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductStairwayDiscountByQuantity(ctx, &V1GetProductStairwayDiscountByQuantityRequest{})
	if err != nil {
		t.Fatalf("GetProductStairwayDiscountByQuantity() error: %v", err)
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

func TestProductVisibilityInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductVisibilityInfo(ctx, &V1ProductVisibilityInfoRequest{})
	if err != nil {
		t.Fatalf("ProductVisibilityInfo() error: %v", err)
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

func TestGetFinanceAccrualTypes(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFinanceAccrualTypes(ctx)
	if err != nil {
		t.Fatalf("GetFinanceAccrualTypes() error: %v", err)
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

func TestGetFinanceAccrualPostings(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFinanceAccrualPostings(ctx, &V1GetFinanceAccrualPostingsRequest{})
	if err != nil {
		t.Fatalf("GetFinanceAccrualPostings() error: %v", err)
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

func TestProductVisibilitySet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductVisibilitySet(ctx, &V1ProductVisibilitySetRequest{})
	if err != nil {
		t.Fatalf("ProductVisibilitySet() error: %v", err)
	}
	_ = resp
}

func TestGetFbpPosting(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFbpPosting(ctx, &PostingV1GetFbpPostingRequest{})
	if err != nil {
		t.Fatalf("GetFbpPosting() error: %v", err)
	}
	_ = resp
}

func TestCreateCompanyFinanceRealizationPostingReport(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CreateCompanyFinanceRealizationPostingReport(ctx, &V1CreateCompanyFinanceRealizationPostingReportRequest{})
	if err != nil {
		t.Fatalf("CreateCompanyFinanceRealizationPostingReport() error: %v", err)
	}
	_ = resp
}

func TestBoxGetByBarcode(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.BoxGetByBarcode(ctx, &ReturnV1BoxGetByBarcodeRequest{})
	if err != nil {
		t.Fatalf("BoxGetByBarcode() error: %v", err)
	}
	_ = resp
}

func TestDescriptionCategoryDependentAttributes(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.DescriptionCategoryDependentAttributes(ctx, &DescriptionV1DescriptionCategoryDependentAttributesRequest{})
	if err != nil {
		t.Fatalf("DescriptionCategoryDependentAttributes() error: %v", err)
	}
	_ = resp
}

func TestDescriptionCategoryDependentAttributesValues(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.DescriptionCategoryDependentAttributesValues(ctx, &DescriptionV1DescriptionCategoryDependentAttributesValuesRequest{})
	if err != nil {
		t.Fatalf("DescriptionCategoryDependentAttributesValues() error: %v", err)
	}
	_ = resp
}

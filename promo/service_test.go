package promo

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

func TestPromosProductsDeactivate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PromosProductsDeactivate(ctx, &SellerApiProductIDsV1Request{})
	if err != nil {
		t.Fatalf("PromosProductsDeactivate() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsCreateInstallment(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsCreateInstallment(ctx, &V1SellerActionsCreateInstallmentRequest{})
	if err != nil {
		t.Fatalf("SellerActionsCreateInstallment() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsUpdateDiscount(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsUpdateDiscount(ctx, &V1SellerActionsUpdateDiscountRequest{})
	_ = err
}

func TestSellerActionsUpdateMultiLevelDiscount(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsUpdateMultiLevelDiscount(ctx, &V1SellerActionsUpdateMultiLevelDiscountRequest{})
	_ = err
}

func TestActionsAutoAddProductsCandidates(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsCandidates(ctx, &ActionsV1ActionsAutoAddProductsCandidatesRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsCandidates() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsProductsDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsProductsDelete(ctx, &V1SellerActionsProductsDeleteRequest{})
	_ = err
}

func TestActionsAutoAddProductsList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsList(ctx, &ActionsV1ActionsAutoAddProductsListRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsList() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsUpdateInstallment(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsUpdateInstallment(ctx, &V1SellerActionsUpdateInstallmentRequest{})
	_ = err
}

func TestTaskDecline(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.TaskDecline(ctx, &V1DeclineDiscountTasksRequest{})
	if err != nil {
		t.Fatalf("TaskDecline() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsUpdateDiscountWithCondition(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsUpdateDiscountWithCondition(ctx, &V1SellerActionsUpdateDiscountWithConditionRequest{})
	_ = err
}

func TestPromosProductsActivate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PromosProductsActivate(ctx, &SellerApiActivateProductV1Request{})
	if err != nil {
		t.Fatalf("PromosProductsActivate() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsUpdateVoucher(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsUpdateVoucher(ctx, &V1SellerActionsUpdateVoucherRequest{})
	_ = err
}

func TestTaskList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.TaskList(ctx, &V1GetDiscountTaskListRequest{})
	if err != nil {
		t.Fatalf("TaskList() error: %v", err)
	}
	_ = resp
}

func TestPromosProducts(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PromosProducts(ctx, &SellerApiGetSellerProductV1Request{})
	if err != nil {
		t.Fatalf("PromosProducts() error: %v", err)
	}
	_ = resp
}

func TestActionsAutoAddProductsDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsDelete(ctx, &ActionsV1ActionsAutoAddProductsDeleteRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsDelete() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsChangeActivity(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsChangeActivity(ctx, &V1SellerActionsChangeActivityRequest{})
	_ = err
}

func TestPromosCandidates(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PromosCandidates(ctx, &SellerApiGetSellerProductV1Request{})
	if err != nil {
		t.Fatalf("PromosCandidates() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsProductsAdd(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsProductsAdd(ctx, &V1SellerActionsProductsAddRequest{})
	_ = err
}

func TestSellerActionsCreateMultiLevelDiscount(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsCreateMultiLevelDiscount(ctx, &V1SellerActionsCreateMultiLevelDiscountRequest{})
	if err != nil {
		t.Fatalf("SellerActionsCreateMultiLevelDiscount() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsProductsList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsProductsList(ctx, &V1SellerActionsProductsListRequest{})
	if err != nil {
		t.Fatalf("SellerActionsProductsList() error: %v", err)
	}
	_ = resp
}

func TestPromos(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.Promos(ctx)
	if err != nil {
		t.Fatalf("Promos() error: %v", err)
	}
	_ = resp
}

func TestTaskApprove(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.TaskApprove(ctx, &V1ApproveDiscountTasksRequest{})
	if err != nil {
		t.Fatalf("TaskApprove() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsVoucherGet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsVoucherGet(ctx, &V1SellerActionsVoucherGetRequest{})
	if err != nil {
		t.Fatalf("SellerActionsVoucherGet() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsArchive(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.SellerActionsArchive(ctx, &V1SellerActionsArchiveRequest{})
	_ = err
}

func TestSellerActionsList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsList(ctx, &V1SellerActionsListRequest{})
	if err != nil {
		t.Fatalf("SellerActionsList() error: %v", err)
	}
	_ = resp
}

func TestActionsAutoAddProductsUpdate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsUpdate(ctx, &ActionsV1ActionsAutoAddProductsUpdateRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsUpdate() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsCreateDiscountWithCondition(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsCreateDiscountWithCondition(ctx, &V1SellerActionsCreateDiscountWithConditionRequest{})
	if err != nil {
		t.Fatalf("SellerActionsCreateDiscountWithCondition() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsCreateDiscount(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsCreateDiscount(ctx, &V1SellerActionsCreateDiscountRequest{})
	if err != nil {
		t.Fatalf("SellerActionsCreateDiscount() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsProductsCandidates(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsProductsCandidates(ctx, &V1SellerActionsProductsCandidatesRequest{})
	if err != nil {
		t.Fatalf("SellerActionsProductsCandidates() error: %v", err)
	}
	_ = resp
}

func TestSellerActionsCreateVoucher(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerActionsCreateVoucher(ctx, &V1SellerActionsCreateVoucherRequest{})
	if err != nil {
		t.Fatalf("SellerActionsCreateVoucher() error: %v", err)
	}
	_ = resp
}

func TestActionsAutoAddProductsCandidatesV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsCandidatesV2(ctx, &ActionsV2ActionsAutoAddProductsCandidatesRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsCandidatesV2() error: %v", err)
	}
	_ = resp
}

func TestActionsAutoAddProductsDeleteV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsDeleteV2(ctx, &ActionsV2ActionsAutoAddProductsDeleteRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsDeleteV2() error: %v", err)
	}
	_ = resp
}

func TestActionsAutoAddProductsListV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsListV2(ctx, &ActionsV2ActionsAutoAddProductsListRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsListV2() error: %v", err)
	}
	_ = resp
}

func TestActionsAutoAddProductsUpdateV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsAutoAddProductsUpdateV2(ctx, &ActionsV2ActionsAutoAddProductsUpdateRequest{})
	if err != nil {
		t.Fatalf("ActionsAutoAddProductsUpdateV2() error: %v", err)
	}
	_ = resp
}

func TestActionsCandidates(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsCandidates(ctx, &ActionsV2ActionsCandidatesRequest{})
	if err != nil {
		t.Fatalf("ActionsCandidates() error: %v", err)
	}
	_ = resp
}

func TestActionsProducts(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsProducts(ctx, &ActionsV2ActionsProductsRequest{})
	if err != nil {
		t.Fatalf("ActionsProducts() error: %v", err)
	}
	_ = resp
}

func TestActionsProductsDeactivate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsProductsDeactivate(ctx, &ActionsV2ActionsProductsDeactivateRequest{})
	if err != nil {
		t.Fatalf("ActionsProductsDeactivate() error: %v", err)
	}
	_ = resp
}

func TestActionsProductsUpdate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActionsProductsUpdate(ctx, &ActionsV1ActionsProductsUpdateRequest{})
	if err != nil {
		t.Fatalf("ActionsProductsUpdate() error: %v", err)
	}
	_ = resp
}

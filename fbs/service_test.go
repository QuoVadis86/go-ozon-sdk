package fbs

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

func TestShipFbsPostingV4(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ShipFbsPostingV4(ctx, &V4FbsPostingShipV4Request{})
	if err != nil {
		t.Fatalf("ShipFbsPostingV4() error: %v", err)
	}
	_ = resp
}

func TestCarriageApprove(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	err := svc.CarriageApprove(ctx, &V1CarriageApproveRequest{})
	_ = err
}

func TestCarriageCancel(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CarriageCancel(ctx, &V1CarriageCancelRequest{})
	if err != nil {
		t.Fatalf("CarriageCancel() error: %v", err)
	}
	_ = resp
}

func TestMoveFbsPostingToAwaitingDelivery(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.MoveFbsPostingToAwaitingDelivery(ctx, &V2MovePostingToAwaitingDeliveryRequest{})
	if err != nil {
		t.Fatalf("MoveFbsPostingToAwaitingDelivery() error: %v", err)
	}
	_ = resp
}

func TestListCountryProductFbsPostingV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ListCountryProductFbsPostingV2(ctx, &V2FbsPostingProductCountryListRequest{})
	if err != nil {
		t.Fatalf("ListCountryProductFbsPostingV2() error: %v", err)
	}
	_ = resp
}

func TestShipFbsPostingPackage(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ShipFbsPostingPackage(ctx, &V4FbsPostingShipPackageV4Request{})
	if err != nil {
		t.Fatalf("ShipFbsPostingPackage() error: %v", err)
	}
	_ = resp
}

func TestPostingFbsUnfulfilledList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PostingFbsUnfulfilledList(ctx, &PostingV4PostingFbsUnfulfilledListRequest{})
	if err != nil {
		t.Fatalf("PostingFbsUnfulfilledList() error: %v", err)
	}
	_ = resp
}

func TestActPostingList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ActPostingList(ctx, &V2PostingFBSActGetPostingsRequest{})
	if err != nil {
		t.Fatalf("ActPostingList() error: %v", err)
	}
	_ = resp
}

func TestPostingFBSActGetContainerLabels(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PostingFBSActGetContainerLabels(ctx, &PostingPostingFBSActGetContainerLabelsRequest{})
	if err != nil {
		t.Fatalf("PostingFBSActGetContainerLabels() error: %v", err)
	}
	_ = resp
}

func TestSetPostings(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SetPostings(ctx, &V1SetPostingsRequest{})
	if err != nil {
		t.Fatalf("SetPostings() error: %v", err)
	}
	_ = resp
}

func TestGetFbsPostingUnfulfilledList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFbsPostingUnfulfilledList(ctx, &Postingv3GetFbsPostingUnfulfilledListRequest{})
	if err != nil {
		t.Fatalf("GetFbsPostingUnfulfilledList() error: %v", err)
	}
	_ = resp
}

func TestGetFbsPostingByBarcode(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFbsPostingByBarcode(ctx, &PostingGetFbsPostingByBarcodeRequest{})
	if err != nil {
		t.Fatalf("GetFbsPostingByBarcode() error: %v", err)
	}
	_ = resp
}

func TestCancelFbsPostingProduct(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CancelFbsPostingProduct(ctx, &PostingPostingProductCancelRequest{})
	if err != nil {
		t.Fatalf("CancelFbsPostingProduct() error: %v", err)
	}
	_ = resp
}

func TestSetCountryProductFbsPostingV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SetCountryProductFbsPostingV2(ctx, &V2FbsPostingProductCountrySetRequest{})
	if err != nil {
		t.Fatalf("SetCountryProductFbsPostingV2() error: %v", err)
	}
	_ = resp
}

func TestPostingFbsList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PostingFbsList(ctx, &PostingV4PostingFbsListRequest{})
	if err != nil {
		t.Fatalf("PostingFbsList() error: %v", err)
	}
	_ = resp
}

func TestGetPostingFbsCancelReasonList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetPostingFbsCancelReasonList(ctx)
	if err != nil {
		t.Fatalf("GetPostingFbsCancelReasonList() error: %v", err)
	}
	_ = resp
}

func TestAssemblyFbsPostingList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.AssemblyFbsPostingList(ctx, &V1AssemblyFbsPostingListRequest{})
	if err != nil {
		t.Fatalf("AssemblyFbsPostingList() error: %v", err)
	}
	_ = resp
}

func TestAssemblyFbsProductList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.AssemblyFbsProductList(ctx, &V1AssemblyFbsProductListRequest{})
	if err != nil {
		t.Fatalf("AssemblyFbsProductList() error: %v", err)
	}
	_ = resp
}

func TestPostingFBSPackageLabel(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PostingFBSPackageLabel(ctx, &PostingPostingFBSPackageLabelRequest{})
	if err != nil {
		t.Fatalf("PostingFBSPackageLabel() error: %v", err)
	}
	_ = resp
}

func TestGetFbsPostingListV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFbsPostingListV3(ctx, &Postingv3GetFbsPostingListRequest{})
	if err != nil {
		t.Fatalf("GetFbsPostingListV3() error: %v", err)
	}
	_ = resp
}

func TestCancelFbsPosting(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CancelFbsPosting(ctx, &PostingCancelFbsPostingRequest{})
	if err != nil {
		t.Fatalf("CancelFbsPosting() error: %v", err)
	}
	_ = resp
}

func TestGetCarriageAvailableList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetCarriageAvailableList(ctx, &Postingv1GetCarriageAvailableListRequest{})
	if err != nil {
		t.Fatalf("GetCarriageAvailableList() error: %v", err)
	}
	_ = resp
}

func TestGetPostingFbsCancelReasonV1(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetPostingFbsCancelReasonV1(ctx, &PostingCancelReasonRequest{})
	if err != nil {
		t.Fatalf("GetPostingFbsCancelReasonV1() error: %v", err)
	}
	_ = resp
}

func TestAssemblyCarriagePostingList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.AssemblyCarriagePostingList(ctx, &V1AssemblyCarriagePostingListRequest{})
	if err != nil {
		t.Fatalf("AssemblyCarriagePostingList() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingTrackingNumberSet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingTrackingNumberSet(ctx, &PostingFbsPostingTrackingNumberSetRequest{})
	if err != nil {
		t.Fatalf("FbsPostingTrackingNumberSet() error: %v", err)
	}
	_ = resp
}

func TestCarriageGet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CarriageGet(ctx, &CarriageCarriageGetRequest{})
	if err != nil {
		t.Fatalf("CarriageGet() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingLastMile(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingLastMile(ctx, &PostingFbsPostingLastMileRequest{})
	if err != nil {
		t.Fatalf("FbsPostingLastMile() error: %v", err)
	}
	_ = resp
}

func TestSetPostingCutoff(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SetPostingCutoff(ctx, &V1SetPostingCutoffRequest{})
	if err != nil {
		t.Fatalf("SetPostingCutoff() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingDelivering(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingDelivering(ctx, &PostingFbsPostingDeliveringRequest{})
	if err != nil {
		t.Fatalf("FbsPostingDelivering() error: %v", err)
	}
	_ = resp
}

func TestGetFbsPostingV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetFbsPostingV3(ctx, &Postingv3GetFbsPostingRequest{})
	if err != nil {
		t.Fatalf("GetFbsPostingV3() error: %v", err)
	}
	_ = resp
}

func TestFbsPostingDelivered(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbsPostingDelivered(ctx, &PostingFbsPostingDeliveredRequest{})
	if err != nil {
		t.Fatalf("FbsPostingDelivered() error: %v", err)
	}
	_ = resp
}

func TestCarriageCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CarriageCreate(ctx, &V1CarriageCreateRequest{})
	if err != nil {
		t.Fatalf("CarriageCreate() error: %v", err)
	}
	_ = resp
}

func TestAssemblyCarriageProductList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.AssemblyCarriageProductList(ctx, &V1AssemblyCarriageProductListRequest{})
	if err != nil {
		t.Fatalf("AssemblyCarriageProductList() error: %v", err)
	}
	_ = resp
}

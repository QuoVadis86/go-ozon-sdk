package fbo

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

func TestSupplyOrderBundle(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SupplyOrderBundle(ctx, &V1GetSupplyOrderBundleRequest{})
	if err != nil {
		t.Fatalf("SupplyOrderBundle() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectSellerDlvCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectSellerDlvCreate(ctx, &V1FbpDraftDirectSellerDlvCreateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectSellerDlvCreate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftGet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftGet(ctx, &V1FbpDraftGetRequest{})
	if err != nil {
		t.Fatalf("FbpDraftGet() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectGetTimeslot(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectGetTimeslot(ctx, &V1FbpDraftDirectGetTimeslotRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectGetTimeslot() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffRegistrate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffRegistrate(ctx, &V1FbpDraftDropOffRegistrateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffRegistrate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffProductValidate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffProductValidate(ctx, &V1FbpDraftDropOffProductValidateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffProductValidate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftPickupCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftPickupCreate(ctx, &V1FbpDraftPickupCreateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftPickupCreate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffPointList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffPointList(ctx, &V1FbpDraftDropOffPointListRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffPointList() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftPickUpDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftPickUpDelete(ctx, &V1FbpDraftPickUpDeleteRequest{})
	if err != nil {
		t.Fatalf("FbpDraftPickUpDelete() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffDlvEdit(ctx, &V1FbpDraftDropOffDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffDlvEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffPointTimetable(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffPointTimetable(ctx, &V1FbpDraftDropOffPointTimetableRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffPointTimetable() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderPickUpCancel(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderPickUpCancel(ctx, &V1FbpOrderPickUpCancelRequest{})
	if err != nil {
		t.Fatalf("FbpOrderPickUpCancel() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderPickUpDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderPickUpDlvEdit(ctx, &V1FbpOrderPickUpDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpOrderPickUpDlvEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectRegistrate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectRegistrate(ctx, &V1FbpDraftDirectRegistrateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectRegistrate() error: %v", err)
	}
	_ = resp
}

func TestFbpCreateConsignmentNote(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpCreateConsignmentNote(ctx, &V1FbpCreateConsignmentNoteRequest{})
	if err != nil {
		t.Fatalf("FbpCreateConsignmentNote() error: %v", err)
	}
	_ = resp
}

func TestFbpCheckConsignmentNoteState(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpCheckConsignmentNoteState(ctx, &V1FbpCheckConsignmentNoteStateRequest{})
	if err != nil {
		t.Fatalf("FbpCheckConsignmentNoteState() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffProvinceList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffProvinceList(ctx, &V1FbpDraftDropOffProvinceListRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffProvinceList() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectSellerDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectSellerDlvEdit(ctx, &V1FbpDraftDirectSellerDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectSellerDlvEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpCreateAct(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpCreateAct(ctx, &V1FbpCreateActRequest{})
	if err != nil {
		t.Fatalf("FbpCreateAct() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderDirectCancel(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderDirectCancel(ctx, &V1FbpOrderDirectCancelRequest{})
	if err != nil {
		t.Fatalf("FbpOrderDirectCancel() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftPickUpProductValidate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftPickUpProductValidate(ctx, &V1FbpDraftPickUpProductValidateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftPickUpProductValidate() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderList(ctx, &V1FbpOrderListRequest{})
	if err != nil {
		t.Fatalf("FbpOrderList() error: %v", err)
	}
	_ = resp
}

func TestFbpWarehouseList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpWarehouseList(ctx)
	if err != nil {
		t.Fatalf("FbpWarehouseList() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderGet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderGet(ctx, &V1FbpOrderGetRequest{})
	if err != nil {
		t.Fatalf("FbpOrderGet() error: %v", err)
	}
	_ = resp
}

func TestPostingFbpList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.PostingFbpList(ctx, &PostingV1PostingFbpListRequest{})
	if err != nil {
		t.Fatalf("PostingFbpList() error: %v", err)
	}
	_ = resp
}

func TestFbpArchiveList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpArchiveList(ctx, &V1FbpArchiveListRequest{})
	if err != nil {
		t.Fatalf("FbpArchiveList() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftList(ctx, &V1FbpDraftListRequest{})
	if err != nil {
		t.Fatalf("FbpDraftList() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderDropOffTimetable(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderDropOffTimetable(ctx, &V1FbpOrderDropOffTimetableRequest{})
	if err != nil {
		t.Fatalf("FbpOrderDropOffTimetable() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffCreate(ctx, &V1FbpDraftDropOffCreateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffCreate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectCreate(ctx, &V1FbpDraftDirectCreateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectCreate() error: %v", err)
	}
	_ = resp
}

func TestFbpCheckActState(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpCheckActState(ctx, &V1FbpCheckActStateRequest{})
	if err != nil {
		t.Fatalf("FbpCheckActState() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftPickUpRegistrate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftPickUpRegistrate(ctx, &V1FbpDraftPickUpRegistrateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftPickUpRegistrate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectDelete(ctx, &V1FbpDraftDirectDeleteRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectDelete() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderDropOffCancel(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderDropOffCancel(ctx, &V1FbpOrderDropOffCancelRequest{})
	if err != nil {
		t.Fatalf("FbpOrderDropOffCancel() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectTplDlvCreate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectTplDlvCreate(ctx, &V1FbpDraftDirectTplDlvCreateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectTplDlvCreate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectProductValidate(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectProductValidate(ctx, &V1FbpDraftDirectProductValidateRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectProductValidate() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectTimeslotEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectTimeslotEdit(ctx, &V1FbpDraftDirectTimeslotEditRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectTimeslotEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpArchiveGet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpArchiveGet(ctx, &V1FbpArchiveGetRequest{})
	if err != nil {
		t.Fatalf("FbpArchiveGet() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDirectTplDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDirectTplDlvEdit(ctx, &V1FbpDraftDirectTplDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDirectTplDlvEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftDropOffDelete(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftDropOffDelete(ctx, &V1FbpDraftDropOffDeleteRequest{})
	if err != nil {
		t.Fatalf("FbpDraftDropOffDelete() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderDirectSellerDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderDirectSellerDlvEdit(ctx, &V1FbpOrderDirectSellerDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpOrderDirectSellerDlvEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpCreateLabel(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpCreateLabel(ctx, &V1FbpCreateLabelRequest{})
	if err != nil {
		t.Fatalf("FbpCreateLabel() error: %v", err)
	}
	_ = resp
}

func TestFbpEditTimeslot(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpEditTimeslot(ctx, &V1FbpEditTimeslotRequest{})
	if err != nil {
		t.Fatalf("FbpEditTimeslot() error: %v", err)
	}
	_ = resp
}

func TestFbpAvailableTimeslotList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpAvailableTimeslotList(ctx, &V1FbpAvailableTimeslotListRequest{})
	if err != nil {
		t.Fatalf("FbpAvailableTimeslotList() error: %v", err)
	}
	_ = resp
}

func TestFbpDraftPickupDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpDraftPickupDlvEdit(ctx, &V1FbpDraftPickupDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpDraftPickupDlvEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpOrderDropOffDlvEdit(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpOrderDropOffDlvEdit(ctx, &V1FbpOrderDropOffDlvEditRequest{})
	if err != nil {
		t.Fatalf("FbpOrderDropOffDlvEdit() error: %v", err)
	}
	_ = resp
}

func TestFbpGetLabel(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.FbpGetLabel(ctx, &V1FbpGetLabelRequest{})
	if err != nil {
		t.Fatalf("FbpGetLabel() error: %v", err)
	}
	_ = resp
}

package warehouse

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

func TestWarehouseInvalidProductsGet(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseInvalidProductsGet(ctx, &V1WarehouseInvalidProductsGetRequest{})
	if err != nil {
		t.Fatalf("WarehouseInvalidProductsGet() error: %v", err)
	}
	_ = resp
}

func TestListDropOffPointsForCreateFBSWarehouse(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ListDropOffPointsForCreateFBSWarehouse(ctx, &V1ListDropOffPointsForCreateFBSWarehouseRequest{})
	if err != nil {
		t.Fatalf("ListDropOffPointsForCreateFBSWarehouse() error: %v", err)
	}
	_ = resp
}

func TestGetWarehouseFBSOperationStatus(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetWarehouseFBSOperationStatus(ctx, &V1GetWarehouseFBSOperationStatusRequest{})
	if err != nil {
		t.Fatalf("GetWarehouseFBSOperationStatus() error: %v", err)
	}
	_ = resp
}

func TestWarehouseFbsCreatePickUpTimeslotList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseFbsCreatePickUpTimeslotList(ctx, &V1WarehouseFbsCreatePickUpTimeslotListRequest{})
	if err != nil {
		t.Fatalf("WarehouseFbsCreatePickUpTimeslotList() error: %v", err)
	}
	_ = resp
}

func TestUnarchiveWarehouseFBS(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.UnarchiveWarehouseFBS(ctx, &V1UnarchiveWarehouseFBSRequest{})
	if err != nil {
		t.Fatalf("UnarchiveWarehouseFBS() error: %v", err)
	}
	_ = resp
}

func TestDeliveryMethodList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.DeliveryMethodList(ctx, &DeliveryMethodListRequest{})
	if err != nil {
		t.Fatalf("DeliveryMethodList() error: %v", err)
	}
	_ = resp
}

func TestUpdateWarehouseFBSFirstMile(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.UpdateWarehouseFBSFirstMile(ctx, &V1UpdateWarehouseFBSFirstMileRequest{})
	if err != nil {
		t.Fatalf("UpdateWarehouseFBSFirstMile() error: %v", err)
	}
	_ = resp
}

func TestCreateWarehouseFBS(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.CreateWarehouseFBS(ctx, &V1CreateWarehouseFBSRequest{})
	if err != nil {
		t.Fatalf("CreateWarehouseFBS() error: %v", err)
	}
	_ = resp
}

func TestListDropOffPointsForUpdateFBSWarehouse(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ListDropOffPointsForUpdateFBSWarehouse(ctx, &V1ListDropOffPointsForUpdateFBSWarehouseRequest{})
	if err != nil {
		t.Fatalf("ListDropOffPointsForUpdateFBSWarehouse() error: %v", err)
	}
	_ = resp
}

func TestWarehouseFbsUpdateDropOffTimeslotList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseFbsUpdateDropOffTimeslotList(ctx, &V1WarehouseFbsUpdateDropOffTimeslotListRequest{})
	if err != nil {
		t.Fatalf("WarehouseFbsUpdateDropOffTimeslotList() error: %v", err)
	}
	_ = resp
}

func TestDeliveryMethodListV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.DeliveryMethodListV2(ctx, &V2DeliveryMethodListV2Request{})
	if err != nil {
		t.Fatalf("DeliveryMethodListV2() error: %v", err)
	}
	_ = resp
}

func TestWarehouseListV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseListV2(ctx, &V2WarehouseListV2Request{})
	if err != nil {
		t.Fatalf("WarehouseListV2() error: %v", err)
	}
	_ = resp
}

func TestArchiveWarehouseFBS(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ArchiveWarehouseFBS(ctx, &V1ArchiveWarehouseFBSRequest{})
	if err != nil {
		t.Fatalf("ArchiveWarehouseFBS() error: %v", err)
	}
	_ = resp
}

func TestWarehouseFbsUpdatePickUpTimeslotList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseFbsUpdatePickUpTimeslotList(ctx, &V1WarehouseFbsUpdatePickUpTimeslotListRequest{})
	if err != nil {
		t.Fatalf("WarehouseFbsUpdatePickUpTimeslotList() error: %v", err)
	}
	_ = resp
}

func TestWarehouseList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseList(ctx, &V1WarehouseListRequest{})
	if err != nil {
		t.Fatalf("WarehouseList() error: %v", err)
	}
	_ = resp
}

func TestWarehouseWithInvalidProducts(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseWithInvalidProducts(ctx)
	if err != nil {
		t.Fatalf("WarehouseWithInvalidProducts() error: %v", err)
	}
	_ = resp
}

func TestWarehouseFbsCreateDropOffTimeslotList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.WarehouseFbsCreateDropOffTimeslotList(ctx, &V1WarehouseFbsCreateDropOffTimeslotListRequest{})
	if err != nil {
		t.Fatalf("WarehouseFbsCreateDropOffTimeslotList() error: %v", err)
	}
	_ = resp
}

func TestUpdateWarehouseFBS(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.UpdateWarehouseFBS(ctx, &V1UpdateWarehouseFBSRequest{})
	if err != nil {
		t.Fatalf("UpdateWarehouseFBS() error: %v", err)
	}
	_ = resp
}

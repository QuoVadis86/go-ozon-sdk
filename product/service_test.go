package product

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

func TestGetProductRatingBySku(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductRatingBySku(ctx, &V1GetProductRatingBySkuRequest{})
	if err != nil {
		t.Fatalf("GetProductRatingBySku() error: %v", err)
	}
	_ = resp
}

func TestGetProductInfoDescription(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoDescription(ctx, &GetProductInfoDescriptionRequest{})
	if err != nil {
		t.Fatalf("GetProductInfoDescription() error: %v", err)
	}
	_ = resp
}

func TestProductUpdateAttributes(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductUpdateAttributes(ctx, &V1ProductUpdateAttributesRequest{})
	if err != nil {
		t.Fatalf("ProductUpdateAttributes() error: %v", err)
	}
	_ = resp
}

func TestProductArchive(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductArchive(ctx, &ProductArchiveRequest{})
	if err != nil {
		t.Fatalf("ProductArchive() error: %v", err)
	}
	_ = resp
}

func TestGetProductInfoList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoList(ctx, &V3GetProductInfoListRequest{})
	if err != nil {
		t.Fatalf("GetProductInfoList() error: %v", err)
	}
	_ = resp
}

func TestGetProductInfoSubscription(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductInfoSubscription(ctx, &V1GetProductInfoSubscriptionRequest{})
	if err != nil {
		t.Fatalf("GetProductInfoSubscription() error: %v", err)
	}
	_ = resp
}

func TestGetUploadQuota(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetUploadQuota(ctx)
	if err != nil {
		t.Fatalf("GetUploadQuota() error: %v", err)
	}
	_ = resp
}

func TestGetProductList(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductList(ctx, &V3GetProductListRequest{})
	if err != nil {
		t.Fatalf("GetProductList() error: %v", err)
	}
	_ = resp
}

func TestProductGetRelatedSKU(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductGetRelatedSKU(ctx, &V1ProductGetRelatedSKURequest{})
	if err != nil {
		t.Fatalf("ProductGetRelatedSKU() error: %v", err)
	}
	_ = resp
}

func TestImportProductsV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ImportProductsV3(ctx, &V3ImportProductsRequest{})
	if err != nil {
		t.Fatalf("ImportProductsV3() error: %v", err)
	}
	_ = resp
}

func TestImportProductsBySKU(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ImportProductsBySKU(ctx, &ImportProductsBySKURequest{})
	if err != nil {
		t.Fatalf("ImportProductsBySKU() error: %v", err)
	}
	_ = resp
}

func TestDeleteProducts(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.DeleteProducts(ctx, &V2DeleteProductsRequest{})
	if err != nil {
		t.Fatalf("DeleteProducts() error: %v", err)
	}
	_ = resp
}

func TestProductUpdateOfferID(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductUpdateOfferID(ctx, &V1ProductUpdateOfferIdRequest{})
	if err != nil {
		t.Fatalf("ProductUpdateOfferID() error: %v", err)
	}
	_ = resp
}

func TestProductInfoPicturesV2(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductInfoPicturesV2(ctx, &V2ProductInfoPicturesRequest{})
	if err != nil {
		t.Fatalf("ProductInfoPicturesV2() error: %v", err)
	}
	_ = resp
}

func TestGetImportProductsInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetImportProductsInfo(ctx, &GetImportProductsInfoRequest{})
	if err != nil {
		t.Fatalf("GetImportProductsInfo() error: %v", err)
	}
	_ = resp
}

func TestProductUnarchive(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductUnarchive(ctx, &ProductUnarchiveRequest{})
	if err != nil {
		t.Fatalf("ProductUnarchive() error: %v", err)
	}
	_ = resp
}

func TestGetProductAttributesV4(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductAttributesV4(ctx, &V4GetProductAttributesV4Request{})
	if err != nil {
		t.Fatalf("GetProductAttributesV4() error: %v", err)
	}
	_ = resp
}

func TestProductImportPictures(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductImportPictures(ctx, &V1ProductImportPicturesRequest{})
	if err != nil {
		t.Fatalf("ProductImportPictures() error: %v", err)
	}
	_ = resp
}

func TestProductInfoWrongVolume(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.ProductInfoWrongVolume(ctx, &V1ProductInfoWrongVolumeRequest{})
	if err != nil {
		t.Fatalf("ProductInfoWrongVolume() error: %v", err)
	}
	_ = resp
}

func TestGetProductAttributesV3(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetProductAttributesV3(ctx, &V3GetProductAttributesV3Request{})
	if err != nil {
		t.Fatalf("GetProductAttributesV3() error: %v", err)
	}
	_ = resp
}

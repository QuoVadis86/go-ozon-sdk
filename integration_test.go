package ozon

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/QuoVadis86/go-ozon-sdk/product"
	"github.com/QuoVadis86/go-ozon-sdk/seller"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"github.com/QuoVadis86/go-ozon-sdk/warehouse"
)

func skipIfNoCreds(t *testing.T) {
	t.Helper()
	if os.Getenv("OZON_CLIENT_ID") == "" || os.Getenv("OZON_API_KEY") == "" {
		t.Skip("set OZON_CLIENT_ID and OZON_API_KEY to run integration tests")
	}
}

func testClient(t *testing.T) *transport.Client {
	t.Helper()
	return transport.New(os.Getenv("OZON_CLIENT_ID"), os.Getenv("OZON_API_KEY"), nil)
}

func TestIntegration_APIFlow(t *testing.T) {
	skipIfNoCreds(t)
	ctx := context.Background()
	cl := testClient(t)

	// 1. Seller info
	t.Run("SellerInfo", func(t *testing.T) {
		svc := &seller.Service{Client: cl}
		resp, err := svc.SellerInfo(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			t.Fatal("nil response")
		}
		t.Logf("Company: %v", resp.Company)
	})

	// 2. Roles
	t.Run("Roles", func(t *testing.T) {
		svc := &seller.Service{Client: cl}
		resp, err := svc.RolesByToken(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			t.Fatal("nil response")
		}
		t.Logf("Roles: %d", len(resp.Roles))
	})

	// 3. Product list -> gather product IDs
	var productIDs []int64
	t.Run("ProductList", func(t *testing.T) {
		svc := &product.Service{Client: cl}
		resp, err := svc.GetProductList(ctx, &product.V3GetProductListRequest{
			Filter: product.V3GetProductListRequestFilter{Visibility: "ALL"},
			Limit:  10,
		})
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			t.Fatal("nil response")
		}
		t.Logf("Products: %d", len(resp.Result.Items))
		for _, p := range resp.Result.Items {
			productIDs = append(productIDs, p.ProductID)
			t.Logf("  product_id=%d offer_id=%s", p.ProductID, p.OfferID)
		}
	})

	// 4. Product info using IDs from step 3
	if len(productIDs) > 0 {
		t.Run("ProductInfo", func(t *testing.T) {
			svc := &product.Service{Client: cl}
			var skus []string
			for _, id := range productIDs {
				skus = append(skus, strconv.FormatInt(id, 10))
			}
			resp, err := svc.GetProductInfoList(ctx, &product.V3GetProductInfoListRequest{
				SKU: skus,
			})
			if err != nil {
				t.Fatal(err)
			}
			if resp == nil {
				t.Fatal("nil response")
			}
			t.Logf("Product info: %d items", len(resp.Items))
		})
	}

	// 5. Warehouse list
	t.Run("WarehouseList", func(t *testing.T) {
		svc := &warehouse.Service{Client: cl}
		resp, err := svc.WarehouseListV2(ctx, &warehouse.V2WarehouseListV2Request{
			Limit: 10,
		})
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			t.Fatal("nil response")
		}
		t.Logf("Warehouses: %d", len(resp.Warehouses))
		for _, w := range resp.Warehouses {
			t.Logf("  warehouse_id=%d name=%s", w.WarehouseID, w.Name)
		}
	})
}

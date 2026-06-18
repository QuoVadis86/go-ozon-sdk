package main

import (
	"context"
	"fmt"
	"log"

	ozon "github.com/QuoVadis86/go-ozon-sdk"
	"github.com/QuoVadis86/go-ozon-sdk/types"
)

func main() {
	client := ozon.NewClient("your-client-id", "your-api-key", nil)
	ctx := context.Background()

	info, err := client.Seller.SellerInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Company: %v\n", info.Company)
	fmt.Printf("Premium: %v\n", info.Subscription)

	products, err := client.Product.ListV3(ctx, &types.V3ProductListRequest{Limit: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Products: %d\n", len(products.Items))
}

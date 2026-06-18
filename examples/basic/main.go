package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuoVadis86/go-ozon-sdk/ozon"
)

func main() {
	client := ozon.NewClient("your-client-id", "your-api-key", nil)
	ctx := context.Background()

	// Get seller info
	info, err := client.Seller.SellerInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Company: %s\n", info.Company.Name)
	fmt.Printf("Premium: %v\n", info.Subscription.IsPremium)

	// Get category tree
	tree, err := client.Category.GetTree(ctx, &ozon.V1GetTreeRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Categories: %d\n", len(tree.Result))

	// List products
	products, err := client.Product.ListV3(ctx, &ozon.V3ProductListRequest{
		Limit: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Products found: %d\n", len(products.Items))
}

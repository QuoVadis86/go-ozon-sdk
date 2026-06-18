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

	info, err := client.Seller.SellerInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Company: %v\n", info.Company)
	fmt.Printf("Premium: %v\n", info.Subscription)
}

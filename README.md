# Ozon Seller API Go SDK

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)](https://golang.org/dl/)
[![Go Reference](https://pkg.go.dev/badge/github.com/QuoVadis86/go-ozon-sdk.svg)](https://pkg.go.dev/github.com/QuoVadis86/go-ozon-sdk)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

[📖 中文文档](README_CN.md)

A Go client library for the [Ozon Seller API](https://docs.ozon.ru/api/seller/).  
Covers **263 API endpoints** across **22 service modules** with **1180 generated types**.

## Features

- Complete coverage of Ozon Seller API (v1, v2, v3, v4, v5, v6)
- Strongly-typed request/response structures
- Built-in error handling with detailed API error types
- Context support for cancellation and timeouts
- No external dependencies beyond the Go standard library
- Concurrent-safe transport with connection pooling

## Installation

```bash
go get github.com/QuoVadis86/go-ozon-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    ozon "github.com/QuoVadis86/go-ozon-sdk"
)

func main() {
    // Create client with your credentials
    client := ozon.NewClient("your-client-id", "your-api-key", nil)

    // Get seller information
    ctx := context.Background()
    info, err := client.Seller.SellerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Company: %v\n", info.Company)
    fmt.Printf("Premium: %v\n", info.Subscription)
}
```

## API Coverage

| Service | Endpoints | Description |
|---------|-----------|-------------|
| [Seller](./seller/) | 3 | Account info, roles, API key info |
| [Category](./category/) | 4 | Category tree, attributes, values |
| [Product](./product/) | 20 | Create, update, list, archive products |
| [Prices](./prices/) | 10 | Prices, stocks, discounts |
| [Barcode](./barcode/) | 2 | Generate and bind barcodes |
| [Warehouse](./warehouse/) | 18 | Warehouse management, delivery methods |
| [FBS](./fbs/) | 33 | FBS orders, shipments, marks, delivery |
| [FBO](./fbo/) | 47 | FBO orders, supply requests, FBP drafts |
| [Finance](./finance/) | 6 | Financial reports, transactions |
| [Promo](./promo/) | 30 | Promotions, seller actions, pricing strategies |
| [Pricing](./pricing/) | 12 | Pricing strategies, competitors |
| [Report](./report/) | 8 | Report generation and management |
| [Review](./review/) | 16 | Reviews, questions and answers |
| [Chat](./chat/) | 6 | Customer chats |
| [Notification](./notification/) | - | Push notification configuration |
| [Returns](./returns/) | 12 | Returns management (FBO, FBS, rFBS) |
| [Pass](./pass/) | 8 | Access passes |
| [Rating](./rating/) | 2 | Seller rating |
| [Delivery](./delivery/) | 2 | Delivery methods, polygons |
| [Beta](./beta/) | 17 | Beta methods, analytics, visibility |
| [Premium](./premium/) | 7 | Premium analytics, chat, finance |

## Examples

### Manage Products

```go
import "github.com/QuoVadis86/go-ozon-sdk/product"

// Create or update products
resp, err := client.Product.ImportProductsV3(ctx, &product.V3ImportProductsRequest{
    Items: []product.V3ImportProductsRequestItem{
        {
            OfferID: "my-sku-001",
            Name:    "Product Name",
            Price:   "999.99",
            CurrencyCode: "RUB",
            // ... other fields
        },
    },
})
```

### Manage FBS Orders

```go
import "github.com/QuoVadis86/go-ozon-sdk/fbs"

// List FBS postings
postings, err := client.FBS.ListFBSV4(ctx, &fbs.V4FbsPostingListRequest{
    Limit: 100,
    Filter: fbs.V4FbsPostingListRequestFilter{
        Status: "awaiting_deliver",
    },
})
```

### Update Prices and Stocks

```go
import "github.com/QuoVadis86/go-ozon-sdk/prices"

// Update stocks
stocksResp, err := client.Prices.ProductsStocksV2(ctx, &prices.V2ProductsStocksRequest{
    Stocks: []prices.V2ProductsStocksRequestStock{
        {
            Stock:   50,
            OfferID: "my-sku-001",
            WarehouseID: 12345,
        },
    },
})
```

## Error Handling

All API errors return `*ozon.APIError` with status code and error details:

```go
import ozon "github.com/QuoVadis86/go-ozon-sdk"

resp, err := client.Product.ImportProductsV3(ctx, req)
if err != nil {
    var apiErr *ozon.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("Error %d: %s\n", apiErr.Code, apiErr.Message)
    }
}
```

## License

MIT

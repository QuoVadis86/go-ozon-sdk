# Ozon Seller API Go SDK

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)](https://golang.org/dl/)
[![Go Reference](https://pkg.go.dev/badge/github.com/QuoVadis86/go-ozon-sdk.svg)](https://pkg.go.dev/github.com/QuoVadis86/go-ozon-sdk)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

[📖 中文文档](README_CN.md)

**go-ozon-sdk** is a Go client library for the [Ozon Seller API](https://docs.ozon.ru/api/seller/) — Ozon marketplace API for products, FBO/FBS orders, prices, stocks, warehouses, finance and reviews.  
Covers **277 API endpoints** across **20 service modules** with **1284 Swagger schemas**.

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

## Response Language

Use `WithLanguage` to send `x-o3-language` only for selected calls on the same client:

```go
client := ozon.NewClient("your-client-id", "your-api-key", nil)
zhCtx := ozon.WithLanguage(ctx, ozon.LanguageZHHans)
resp, err := client.Product.GetProductInfoList(zhCtx, req)
```

Calls using the original `ctx` remain unchanged. `ClientOptions.Language` can still set a client default; `WithLanguage` overrides it for calls using the derived context. `WithLanguage(ctx, ozon.LanguageDefault)` omits the header even when a client default is set.

The available constants are `LanguageDefault`, `LanguageRU`, `LanguageEN`, `LanguageTR`, and `LanguageZHHans`. The SDK maps `LanguageZHHans` to the exact header value `zh-Hans`; the category API's `LanguageLanguageZHHans` is a separate request-body enum and cannot be used here.

On `/v3/product/info/list`, this header can localize attribute names such as `items[].errors[].texts.attribute_name`. Error messages in `texts.message` may still be English or Russian. Handle errors using structured fields such as `errors[].code` and `attribute_id`.

## API Coverage

| Service | Endpoints | Description |
|---------|-----------|-------------|
| [Seller](./seller/service.go) | 3 | Account info, roles, API key info |
| [Category](./category/service.go) | 4 | Category tree, attributes, values |
| [Product](./product/service.go) | 20 | Create, update, list, archive products |
| [Prices](./prices/service.go) | 10 | Prices, stocks, discounts |
| [Barcode](./barcode/service.go) | 3 | Generate and bind barcodes |
| [Warehouse](./warehouse/service.go) | 18 | Warehouse management, delivery methods |
| [FBS](./fbs/service.go) | 42 | FBS orders, shipments, marks, delivery |
| [FBO](./fbo/service.go) | 48 | FBO orders, supply requests, FBP drafts |
| [Finance](./finance/service.go) | 6 | Financial reports, transactions |
| [Promo](./promo/service.go) | 38 | Promotions, seller actions, pricing strategies |
| [Pricing](./pricing/service.go) | 12 | Pricing strategies, competitors |
| [Report](./report/service.go) | 8 | Report generation and management |
| [Review](./review/service.go) | 16 | Reviews, questions and answers |
| [Chat](./chat/service.go) | 6 | Customer chats |
| [Returns](./returns/service.go) | 7 | Returns management (FBO, FBS, rFBS) |
| [Pass](./pass/service.go) | 8 | Access passes |
| [Rating](./rating/service.go) | 2 | Seller rating |
| [Delivery](./delivery/service.go) | 2 | Delivery methods, polygons |
| [Beta](./beta/service.go) | 17 | Beta methods, analytics, visibility |
| [Premium](./premium/service.go) | 7 | Premium analytics, chat, finance |

## Regenerating the SDK

Save the latest Ozon Swagger JSON to `origin/swagger.json`, then run `go run gen.go`. The generator keeps the existing order of methods, types, and fields; new entries are appended in name order. Running it again with the same Swagger should leave the generated files unchanged.

For endpoint diffs, compatibility rules, and release checks, follow the [SDK maintenance guide](docs/SDK_MAINTENANCE.md).

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

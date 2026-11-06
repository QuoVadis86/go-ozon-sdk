# Ozon Seller API Go SDK

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)](https://golang.org/dl/)
[![Go Reference](https://pkg.go.dev/badge/github.com/QuoVadis86/go-ozon-sdk.svg)](https://pkg.go.dev/github.com/QuoVadis86/go-ozon-sdk)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

[English Documentation](README.md)

**go-ozon-sdk** 是 [Ozon Seller API](https://docs.ozon.ru/api/seller/) 的 Go 语言客户端库——俄罗斯电商平台 Ozon 的卖家 API，覆盖商品、FBO/FBS 订单、价格、库存、仓库、财务与评价。  
覆盖 **277 个 API 端点**，**20 个服务模块**，**1284 个 Swagger 模型**。

## 特性

- 完整覆盖 Ozon Seller API（v1、v2、v3、v4、v5、v6）
- 强类型请求/响应结构
- 内置错误处理与 API 错误详情
- 支持 Context 取消与超时控制
- 零外部依赖（仅使用 Go 标准库）
- 并发安全的 HTTP 连接池

## 安装

```bash
go get github.com/QuoVadis86/go-ozon-sdk
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"
    "log"

    ozon "github.com/QuoVadis86/go-ozon-sdk"
)

func main() {
    // 使用凭证创建客户端
    client := ozon.NewClient("your-client-id", "your-api-key", nil)

    // 获取卖家信息
    ctx := context.Background()
    info, err := client.Seller.SellerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("公司: %s\n", info.Company.Name)
    fmt.Printf("高级订阅: %v\n", info.Subscription.IsPremium)
}
```

## 响应语言

通过 `WithLanguage` 只为同一个 client 的指定调用设置 `x-o3-language`：

```go
client := ozon.NewClient("your-client-id", "your-api-key", nil)
zhCtx := ozon.WithLanguage(ctx, ozon.LanguageZHHans)
resp, err := client.Product.GetProductInfoList(zhCtx, req)
```

继续使用原来的 `ctx` 调用不会受到影响。`ClientOptions.Language` 仍可设置客户端默认语言；`WithLanguage` 只覆盖派生 context 对应的调用。即使客户端设置了默认语言，`WithLanguage(ctx, ozon.LanguageDefault)` 也会在该调用中省略语言请求头。

可选常量为 `LanguageDefault`、`LanguageRU`、`LanguageEN`、`LanguageTR`、`LanguageZHHans`。SDK 会将 `LanguageZHHans` 映射为准确的请求头值 `zh-Hans`；类目接口的 `LanguageLanguageZHHans` 是独立的请求体枚举，不能用于此处。

在 `/v3/product/info/list` 的实测中，该请求头可将 `items[].errors[].texts.attribute_name` 等属性名称本地化，但 `texts.message` 中的错误正文仍可能是英文或俄文。处理错误时应依据 `errors[].code`、`attribute_id` 等结构化字段。

## API 覆盖

| 服务 | 端点数 | 说明 |
|---------|-----------|------|
| [Seller](./seller/service.go) | 3 | 账户信息、角色、API 密钥 |
| [Category](./category/service.go) | 4 | 类目树、属性、属性值 |
| [Product](./product/service.go) | 20 | 创建、更新、列出、归档商品 |
| [Prices](./prices/service.go) | 10 | 价格、库存、折扣 |
| [Barcode](./barcode/service.go) | 3 | 生成和绑定条形码 |
| [Warehouse](./warehouse/service.go) | 18 | 仓库管理、配送方式 |
| [FBS](./fbs/service.go) | 42 | FBS 订单、发货、标签、配送 |
| [FBO](./fbo/service.go) | 48 | FBO 订单、供应请求、FBP 草稿 |
| [Finance](./finance/service.go) | 6 | 财务报表、交易 |
| [Promo](./promo/service.go) | 38 | 促销活动、卖家活动、定价策略 |
| [Pricing](./pricing/service.go) | 12 | 定价策略、竞争对手 |
| [Report](./report/service.go) | 8 | 报告生成与管理 |
| [Review](./review/service.go) | 16 | 评价、问答 |
| [Chat](./chat/service.go) | 6 | 客户聊天 |
| [Returns](./returns/service.go) | 7 | 退货管理（FBO、FBS、rFBS） |
| [Pass](./pass/service.go) | 8 | 通行证 |
| [Rating](./rating/service.go) | 2 | 卖家评级 |
| [Delivery](./delivery/service.go) | 2 | 配送方式、配送区域 |
| [Beta](./beta/service.go) | 17 | 测试方法、分析、可见性 |
| [Premium](./premium/service.go) | 7 | 高级分析、聊天、财务 |

## 重新生成 SDK

将最新 Ozon Swagger JSON 保存为 `origin/swagger.json`，然后运行 `go run gen.go`。生成器保留已有方法、类型和字段的顺序，新内容按名称排序并追加。同一份 Swagger 再运行一次，生成文件应保持不变。

新版 Swagger 的接口对比、兼容处理和发布检查请按 [SDK 维护手册](docs/SDK_MAINTENANCE.md) 执行。

## 示例

### 管理商品

```go
// 创建或更新商品
resp, err := client.Product.ImportProductsV3(ctx, &ozon.V3ImportProductsRequest{
    Items: []ozon.V3ImportProductsRequestItem{
        {
            OfferID: "my-sku-001",
            Name:    "商品名称",
            Price:   "999.99",
            CurrencyCode: "RUB",
        },
    },
})
```

### 管理 FBS 订单

```go
// 获取 FBS 发货列表
postings, err := client.FBS.ListFBSV4(ctx, &ozon.V4FbsPostingListRequest{
    Limit: 100,
    Filter: ozon.V4FbsPostingListRequestFilter{
        Status: "awaiting_deliver",
    },
})
```

### 更新价格和库存

```go
// 更新库存
stocksResp, err := client.Prices.ProductsStocksV2(ctx, &ozon.V2ProductsStocksRequest{
    Stocks: []ozon.V2ProductsStocksRequestStock{
        {
            Stock:   50,
            OfferID: "my-sku-001",
            WarehouseID: 12345,
        },
    },
})
```

## 错误处理

所有 API 错误返回 `*ozon.APIError`，包含状态码和错误详情：

```go
resp, err := client.Product.ImportProductsV3(ctx, req)
if err != nil {
    var apiErr *ozon.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("错误 %d: %s\n", apiErr.Code, apiErr.Message)
    }
}
```

## 许可证

MIT

# Ozon Seller API Go SDK

[![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go)](https://golang.org/dl/)
[![Go Reference](https://pkg.go.dev/badge/github.com/QuoVadis86/go-ozon-sdk.svg)](https://pkg.go.dev/github.com/QuoVadis86/go-ozon-sdk)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

Ozon Seller API 的 Go 语言客户端库。  
覆盖 **263 个 API 端点**，**22 个服务模块**，**1180 个生成类型**。

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

    "github.com/QuoVadis86/go-ozon-sdk/ozon"
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

## API 覆盖

| 服务 | 端点数 | 说明 |
|---------|-----------|------|
| Seller | 3 | 账户信息、角色、API 密钥 |
| Category | 4 | 类目树、属性、属性值 |
| Product | 20 | 创建、更新、列出、归档商品 |
| Prices | 10 | 价格、库存、折扣 |
| Barcode | 2 | 生成和绑定条形码 |
| Warehouse | 18 | 仓库管理、配送方式 |
| FBS | 33 | FBS 订单、发货、标签、配送 |
| FBO | 47 | FBO 订单、供应请求、FBP 草稿 |
| Finance | 6 | 财务报表、交易 |
| Promo | 30 | 促销活动、卖家活动、定价策略 |
| Pricing | 12 | 定价策略、竞争对手 |
| Report | 8 | 报告生成与管理 |
| Review | 16 | 评价、问答 |
| Chat | 6 | 客户聊天 |
| Notification | - | 推送通知配置 |
| Returns | 12 | 退货管理（FBO、FBS、rFBS） |
| Pass | 8 | 通行证 |
| Rating | 2 | 卖家评级 |
| Delivery | 2 | 配送方式、配送区域 |
| Beta | 17 | 测试方法、分析、可见性 |
| Premium | 7 | 高级分析、聊天、财务 |

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

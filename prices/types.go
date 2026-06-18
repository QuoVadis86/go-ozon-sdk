package prices

type V1ProductActionTimerStatusResponse struct {
	Statuses interface{} `json:"statuses"`
}

type ProductImportProductsPricesRequest struct {
	Prices []interface{} `json:"prices"` // 商品价格信息。
}

type ProductImportProductsPricesResponse struct {
	Result []interface{} `json:"result"` // 搜索结果。
}

type V1GetProductInfoDiscountedResponse struct {
	Items interface{} `json:"items"` // 关于减价和主要商品的信息。
}

type V4GetProductInfoStocksResponse struct {
	Cursor string `json:"cursor"` // 后续数据的选择标志。
	Items []interface{} `json:"items"` // 商品信息。
	Total int32 `json:"total"` // 显示库存信息的独特商品数量。
}

type V4GetProductInfoStocksRequest struct {
	Cursor string `json:"cursor"` // 后续数据的选择标志。
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"` // 页面上的值数量。最低为1，最高为1000。
}

type V1ProductUpdateDiscountResponse struct {
	Result bool `json:"result"` // 方式工作结果 `true`, 如果正确完成请求。
}

type V1GetProductInfoDiscountedRequest struct {
	DiscountedSkus interface{} `json:"discounted_skus"` // 降价的商品SKU清单。
}

// 搜索结果。
type Productv2ProductsStocksResponse struct {
	Result []interface{} `json:"result"`
}

type Productv5GetProductInfoPricesV5Request struct {
	Cursor string `json:"cursor"` // 用于选择数据的指针。
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"` // 每页显示的数值数量。
}

type V2GetProductInfoStocksByWarehouseFbsRequestV2 struct {
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	OfferID []interface{} `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	SKU []interface{} `json:"sku"` // Ozon系统中的商品标识符——SKU。
	Cursor string `json:"cursor"` // 后续数据的选择标志。
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2 struct {
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	HasNext bool `json:"has_next"` // 如果响应中未返回所有商品，则为`true`。
	Products []interface{} `json:"products"` // 商品库存。
}

type Productv2ProductsStocksRequest struct {
	Stocks []interface{} `json:"stocks"` // 仓库中商品的信息。
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponse struct {
	Result interface{} `json:"result"` // 该处理方法的结果。
}

type V1ProductActionTimerUpdateRequest struct {
	ProductIds []interface{} `json:"product_ids"` // 卖家系统中的商品识别符列表——`product_id`。
}

type Productv5GetProductInfoPricesV5Response struct {
	Cursor string `json:"cursor"` // 用于选择数据的指针。
	Items interface{} `json:"items"` // 商品列表。
	Total int32 `json:"total"` // 商品列表中的商品数量。
}

type V1ProductUpdateDiscountRequest struct {
	Discount int32 `json:"discount"` // 折扣力度：从3%到99%。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

type V1ProductActionTimerStatusRequest struct {
	ProductIds interface{} `json:"product_ids"` // 卖家系统中的商品识别符列表——`product_id`。
}

type Productsv1GetProductInfoStocksByWarehouseFbsRequest struct {
	SKU interface{} `json:"sku"` // Ozon系统中的商品识别码是SKU。
	OfferID interface{} `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
}

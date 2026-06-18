package prices

type V1ProductActionTimerStatusRequest struct {
	ProductIds interface{} `json:"product_ids"`
}

type V1GetProductInfoDiscountedResponse struct {
	Items interface{} `json:"items"`
}

type Productv2ProductsStocksResponse struct {
	Result []interface{} `json:"result"`
}

type Productv5GetProductInfoPricesV5Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type V4GetProductInfoStocksResponse struct {
	Cursor string `json:"cursor"`
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type V4GetProductInfoStocksRequest struct {
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
	Cursor string `json:"cursor"`
}

type Productv2ProductsStocksRequest struct {
	Stocks []interface{} `json:"stocks"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2 struct {
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
	Cursor string `json:"cursor"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsRequest struct {
	SKU interface{} `json:"sku"`
	OfferID interface{} `json:"offer_id"`
}

type V1GetProductInfoDiscountedRequest struct {
	DiscountedSkus interface{} `json:"discounted_skus"`
}

type ProductImportProductsPricesResponse struct {
	Result []interface{} `json:"result"`
}

type Productv5GetProductInfoPricesV5Response struct {
	Cursor string `json:"cursor"`
	Items interface{} `json:"items"`
	Total int32 `json:"total"`
}

type ProductImportProductsPricesRequest struct {
	Prices []interface{} `json:"prices"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponse struct {
	Result interface{} `json:"result"`
}

type V1ProductUpdateDiscountResponse struct {
	Result bool `json:"result"`
}

type V1ProductActionTimerUpdateRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V2GetProductInfoStocksByWarehouseFbsRequestV2 struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	OfferID []interface{} `json:"offer_id"`
	SKU []interface{} `json:"sku"`
}

type V1ProductActionTimerStatusResponse struct {
	Statuses interface{} `json:"statuses"`
}

type V1ProductUpdateDiscountRequest struct {
	Discount int32 `json:"discount"`
	ProductID int64 `json:"product_id"`
}

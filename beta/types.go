package beta

type FinanceV1GetFinanceAccrualPostingsResponse struct {
	PostingAccruals []interface{} `json:"posting_accruals"` // 按货件统计的应计项目列表。
}

type V1GetProductStairwayDiscountByQuantityRequest struct {
	Skus []interface{} `json:"skus"` // 需要返回内容评级的商品SKU列表。
}

type V5FbsPostingProductExemplarStatusV5Response struct {
	PostingNumber string `json:"posting_number"` // 发货号。
	Products []interface{} `json:"products"` // 商品清单。
	Status string `json:"status"` // 所有样件和备货可用性的验证状态： - `ship_available`——可以备货， - `ship_not_available`——无法备货， - `validation_in_process`——样件正在验证中， - `update_a...
}

type V1FbsPostingProductExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"` // 发货号。
}

type FinanceV1GetFinanceAccrualByDayResponse struct {
	Accruals []interface{} `json:"accruals"` // 应计项目列表。
	LastID string `json:"last_id"` // 页面中最后一个值的标识符。
}

type V1GetFinanceBalanceV1Request struct {
	DateFrom string `json:"date_from"` // 报告期开始日期，格式为 `YYYY-MM-DD`。
	DateTo string `json:"date_to"` // 报告期结束日期，格式为 `YYYY-MM-DD`。`date_from` 与 `date_to` 之间的最⻓间隔为30 天。
}

type V1GetProductStairwayDiscountByQuantityResponse struct {
	Stairways []interface{} `json:"stairways"` // 单个商品的按数量折扣信息。
}

type V1ProductInfoWarehouseStocksResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。 如果该参数为空，则没有更多数据了。
	HasNext bool `json:"has_next"` // 标记是否返回了所有商品： - `true`——请使用不同的`cursor`值重新请求，以获取剩余的值； - `false`——响应中已包含所有值。
	Stocks []interface{} `json:"stocks"` // 商品库存信息。
}

type V5FbsPostingProductExemplarValidateV5Request struct {
	Products []interface{} `json:"products"` // 商品清单。
	PostingNumber string `json:"posting_number"` // 发货号。
}

type V6FbsPostingProductExemplarSetV6Request struct {
	MultiBoxQty int32 `json:"multi_box_qty"` // 商品包装的箱子数量。
	PostingNumber string `json:"posting_number"` // 发货号。
	Products []interface{} `json:"products"` // 商品清单。
}

type V2GetDiscountTaskListV2Request struct {
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。首次请求请留空。
	Limit int64 `json:"limit"` // 每页最大申请数量。
	Status interface{} `json:"status"`
}

type FinanceV1GetFinanceAccrualByDayRequest struct {
	Date string `json:"date"` // 应计日期。最早可查询日期为2022年1月1日。
	LastID string `json:"last_id"` // 页面上最后一个值的标识符。首次请求请留空。 要获取后续值，请指定上一次请求响应中的 `last_id`。
}

// 余额报告。
type V1GetFinanceBalanceV1Response struct {
	Cashflows interface{} `json:"cashflows"`
	Total interface{} `json:"total"`
}

type V1SetProductStairwayDiscountByQuantityResponse struct {
	Accepted bool `json:"accepted"` // `true`，表示请求已接收。请使用方法[/v1/product/stairway-discount/by-quantity/get](#operation/ProductAPI_GetProductStairwayDiscountByQu...
	Errors []interface{} `json:"errors"` // 错误描述。
	Warnings []interface{} `json:"warnings"` // 警告描述。
}

type V2GetDiscountTaskListV2Response struct {
	Tasks []interface{} `json:"tasks"` // 申请列表。
}

type V1DescriptionCategoryTipsRequest struct {
	TypeID []interface{} `json:"type_id"` // 商品类型标识符。可通过方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

type V6FbsPostingProductExemplarCreateOrGetV6Request struct {
	PostingNumber string `json:"posting_number"` // 发货号。
}

type ProductV1ProductVisibilitySetRequest struct {
	ItemPlacement []interface{} `json:"item_placement"` // 商品可见性信息。
}

type V1SetProductStairwayDiscountByQuantityRequest struct {
	Stairways []interface{} `json:"stairways"` // 多个商品的按数量折扣信息。
	SuppressWarnings bool `json:"suppress_warnings"` // 传递 `true` 可忽略警告并设置折扣。
}

type V1DescriptionCategoryTipsResponse struct {
	Result []interface{} `json:"result"` // 提示列表。
}

type ProductV1ProductVisibilityInfoResponse struct {
	Items []interface{} `json:"items"` // 商品列表。
}

type V1PostingFbsSplitResponse struct {
	ParentPosting interface{} `json:"parent_posting"`
	Postings []interface{} `json:"postings"` // 订单被拆分后的货件列表。
}

type V5FbsPostingProductExemplarValidateV5Response struct {
	Products []interface{} `json:"products"` // 商品清单。
}

type FinanceV1GetFinanceAccrualPostingsRequest struct {
	PostingNumbers []interface{} `json:"posting_numbers"` // 货件编号。
}

type ProductV1ProductVisibilitySetResponse struct {
	Items []interface{} `json:"items"` // 商品可见性信息。
	ItemsErrors []interface{} `json:"items_errors"` // 存在错误的商品。
}

type V1ProductInfoWarehouseStocksRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Limit int64 `json:"limit"` // 每页显示的数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V5FbsPostingProductExemplarStatusV5Request struct {
	PostingNumber string `json:"posting_number"` // 发货号。
}

type ProductV1ProductVisibilityInfoRequest struct {
	Skus []interface{} `json:"skus"` // Ozon系统中的商品标识符—— SKU。
}

type V6FbsPostingProductExemplarCreateOrGetV6Response struct {
	MultiBoxQty int32 `json:"multi_box_qty"` // 商品包装的箱子数量。
	PostingNumber string `json:"posting_number"` // 发货号。
	Products []interface{} `json:"products"` // 商品清单。
}

type FinanceV1GetFinanceAccrualTypesResponse struct {
	AccrualTypes []interface{} `json:"accrual_types"` // 应计项目相关信息。
}

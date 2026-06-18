package promo

type SellerApiProductV1Response struct {
	Result interface{} `json:"result"`
}

// 商品清单。
type SellerApiGetSellerProductV1Request struct {
	LastID float64 `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。
	ActionID float64 `json:"action_id"` // 活动识别号。可以使用方法 [/v1/actions](#operation/Promos)获取。
	Limit float64 `json:"limit"` // 每页的答复数量。在默认情况下 — 100。
}

type V1SellerActionsArchiveRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
}

type V1SellerActionsCreateMultiLevelDiscountResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type V1SellerActionsCreateDiscountWithConditionResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type V1SellerActionsUpdateDiscountRequest struct {
	ActionParameters interface{} `json:"action_parameters"`
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type V1SellerActionsUpdateDiscountWithConditionRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	ActionParameters interface{} `json:"action_parameters"`
}

type V1DeclineDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"` // 申请列表。
}

type SellerApiProductV1ResponseDeactivate struct {
	Result interface{} `json:"result"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponse struct {
	Products []interface{} `json:"products"` // 可用于自动添加到促销活动中的商品列表。
	Total int64 `json:"total"` // 商品总数。
}

type V1SellerActionsCreateInstallmentRequest struct {
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
	Title string `json:"title"` // 促销活动名称。
}

type V1ApproveDeclineDiscountTasksResponse struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsProductsListResponse struct {
	Cursor int64 `json:"cursor"` // 用于选择下一批数据的指针。
	HasNext bool `json:"has_next"` // 响应中仅返回了部分值的标志： - `true`——请使用新的`cursor`参数重复请求，以获取其余值； - `false`——响应中已包含所有值。
	Products []interface{} `json:"products"` // 商品信息。
}

type V1SellerActionsCreateVoucherRequest struct {
	DateEnd string `json:"date_end"` // 促销活动结束日期与时间。
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"` // 折扣幅度。
	Title string `json:"title"` // 促销活动名称。
	UserIds []interface{} `json:"user_ids"` // 可使用该促销码的用户标识符列表。
	VoucherParameters interface{} `json:"voucher_parameters"`
	Budget int64 `json:"budget"` // 促销活动预算。预算用尽后，促销活动将停止。
}

type V1SellerActionsCreateVoucherResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type V1SellerActionsVoucherGetResponse struct {
	File string `json:"file"` // 促销码CSV文件链接。
}

type V1SellerActionsListResponse struct {
	Actions []interface{} `json:"actions"` // 促销活动列表。
	Total int64 `json:"total"` // 促销活动总数。
}

type V1SellerActionsCreateMultiLevelDiscountRequest struct {
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
	DiscountLevels []interface{} `json:"discount_levels"` // 折扣等级。
	DiscountType interface{} `json:"discount_type"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"` // `true`，表示促销活动仅面向法人实体。
	Title string `json:"title"` // 促销活动名称。
	DateEnd string `json:"date_end"` // 促销活动结束日期与时间。
}

type SellerApiGetSellerProductV1Response struct {
	Result interface{} `json:"result"`
}

type SellerApiGetSellerActionsV1Response struct {
	Result []interface{} `json:"result"` // 请求结果。
}

type ActionsV1ActionsAutoAddProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	AutoAddDate string `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
	ProductIds []interface{} `json:"product_ids"` // Ozon系统中的商品标识符，即`product_id`。
}

type V1SellerActionsUpdateMultiLevelDiscountRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	ActionParameters interface{} `json:"action_parameters"`
}

type V1SellerActionsUpdateInstallmentRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	ActionParameters interface{} `json:"action_parameters"`
}

type ActionsV1ActionsAutoAddProductsListRequest struct {
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	Offset int64 `json:"offset"` // 在响应中将被跳过的项目数量。例如，如果`offset = 10`，响应将从第11个找到的元素开始。
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	AutoAddDate string `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
}

type ActionsV1ActionsAutoAddProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	AutoAddDate string `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	Offset int64 `json:"offset"` // 在响应中将被跳过的项目数量。例如，如果`offset = 10`，响应将从第11个找到的元素开始。
}

type ActionsV1ActionsAutoAddProductsDeleteResponse struct {
	ProductIds []interface{} `json:"product_ids"` // 已从自动添加中删除的商品ID。
}

type SellerApiActivateProductV1Request struct {
	Products []interface{} `json:"products"` // 商品清单。
	ActionID float64 `json:"action_id"` // 活动识别号。可以使用方法 [/v1/actions](#operation/Promos)获取。
}

type V1SellerActionsUpdateVoucherRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	ActionParameters interface{} `json:"action_parameters"`
}

type V1SellerActionsProductsAddRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Products []interface{} `json:"products"` // 商品信息。
}

type V1SellerActionsCreateInstallmentResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type V1ApproveDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"` // 申请列表。
}

type V1SellerActionsCreateDiscountWithConditionRequest struct {
	Title string `json:"title"` // 促销活动名称。
	DateEnd string `json:"date_end"` // 促销活动结束日期与时间。
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"` // 折扣幅度。
	MinOrderAmount float64 `json:"min_order_amount"` // 折扣生效的订单金额门槛。
}

type V1GetDiscountTaskListResponse struct {
	Result []interface{} `json:"result"` // 申请列表。
}

type ActionsV1ActionsAutoAddProductsUpdateResponse struct {
	ExtremelyLowPrice []interface{} `json:"extremely_low_price"` // 折扣幅度超过70%的商品列表。
	FailedPrice []interface{} `json:"failed_price"` // 未通过价格校验的商品列表。
	Rejected []interface{} `json:"rejected"` // 未能添加或更新的商品ID。
	UpdatedIds []interface{} `json:"updated_ids"` // 已成功添加或更新的商品ID。
	BelowMinPrice []interface{} `json:"below_min_price"` // 价格低于最低价格的商品列表。
}

type ActionsV1ActionsAutoAddProductsListResponse struct {
	Products []interface{} `json:"products"` // 启用自动添加的商品列表。
	Total int64 `json:"total"` // 商品总数。
}

type V1SellerActionsChangeActivityRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	IsTurnOn bool `json:"is_turn_on"` // `true`，用于启用促销活动。
}

type V1SellerActionsProductsCandidatesResponse struct {
	Products []interface{} `json:"products"` // 商品信息。
	Cursor int64 `json:"cursor"` // 用于选择下一批数据的指针。
	HasNext bool `json:"has_next"` // 响应中仅返回了部分值的标志： - `true`——请使用新的`cursor`参数重复请求，以获取其余值； - `false`——响应中已包含所有值。
}

type SellerApiProductIDsV1Request struct {
	ActionID float64 `json:"action_id"` // 活动识别号。可以使用方法 [/v1/actions](#operation/Promos)获取。
	ProductIds []interface{} `json:"product_ids"` // 活动识别号清单。
}

type V1GetDiscountTaskListRequest struct {
	Status interface{} `json:"status"`
	Page int64 `json:"page"` // 需要从中下载折扣申请列表的页面。
	Limit int64 `json:"limit"` // 页面上申请最大数量。
}

type V1SellerActionsListRequest struct {
	ActionIds []interface{} `json:"action_ids"` // 促销活动标识符列表。
	ActionType []interface{} `json:"action_type"` // 促销活动机制： - `DISCOUNT`——折扣； - `VOUCHER_DISCOUNT`——促销码折扣； - `DISCOUNT_WITH_CONDITION`——基于订单总额的折扣； - `INSTALLMENT`——免息分期付款； ...
	Limit int64 `json:"limit"` // 每页显示的数量。
	Offset int64 `json:"offset"` // 在响应中将被跳过的项目数量。例如，当`offset = 10`时，响应将从第11个找到的元素开始。
	Search string `json:"search"` // 按促销活动名称搜索。
	Status []interface{} `json:"status"` // 促销活动状态： - `ACTIVE`—— 活跃； - `ENDED`——已结束； - `PLANNED`——已计划； - `PAUSED`——已暂停。
}

type V1SellerActionsVoucherGetRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
}

type V1SellerActionsProductsListRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Cursor int64 `json:"cursor"` // 用于选择下一批数据的指针。
	Limit int64 `json:"limit"` // 响应中的最大元素数量。
}

type V1SellerActionsCreateDiscountResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type V1SellerActionsCreateDiscountRequest struct {
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
	MinActionPercent float64 `json:"min_action_percent"` // 最低折扣百分比。
	Title string `json:"title"` // 促销活动名称。
	DateEnd string `json:"date_end"` // 促销活动结束日期与时间。
}

type V1SellerActionsProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Skus []interface{} `json:"skus"` // Ozon系统中的商品标识符——SKU。
}

type V1SellerActionsProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Cursor int64 `json:"cursor"` // 用于选择下一批数据的指针。
	Limit int64 `json:"limit"` // 响应中的最大元素数量。
}

type ActionsV1ActionsAutoAddProductsUpdateRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
	AutoAddDate string `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
	ToUpdate []interface{} `json:"to_update"` // 需要添加到自动添加中或在自动添加中更新的商品列表。
}

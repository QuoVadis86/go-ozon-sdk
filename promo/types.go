package promo

type V1SellerActionsProductsListRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Cursor   int64 `json:"cursor"`    // 用于选择下一批数据的指针。
	Limit    int64 `json:"limit"`     // 响应中的最大元素数量。
}

type V1SellerActionsProductsDeleteRequest struct {
	ActionID int64    `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Skus     []string `json:"skus"`      // Ozon系统中的商品标识符——SKU。
}

type V1SellerActionsCreateVoucherResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type ParameterDiscountLevels struct {
	DiscountValue float64 `json:"discount_value"` // 折扣值。
	OrderAmount   float64 `json:"order_amount"`   // 订单金额。
}

// 促销码类型： - `UNSPECIFIED`——未定义； - `ONE`——面向所有买家的促销码，可用于1个订单； - `MULTIPLE`——面向所有买家的促销码，可用于任意数量的订单； - `UNIQUE`——面向1位买家的促销码，可用...
type ActionParameterVoucherParameterTypeEnum string

// 促销码参数。
type ActionParameterVoucherParameter struct {
	CountCodes int64                                   `json:"count_codes"` // 促销码代码。
	IsPrivate  bool                                    `json:"is_private"`  // `true`，表示促销码不公开。
	Type       ActionParameterVoucherParameterTypeEnum `json:"type_"`
}

// 折扣申请状态： - `NEW` — 新的， - `SEEN` — 已查看的， - `APPROVED` — 已批准的， - `PARTLY_APPROVED` — 部分批准的， - `DECLINED` — 取消的， - `AUTO_DEC...
type V1DiscountTaskStatus string

type V1GetDiscountTaskListRequest struct {
	Page   int64                `json:"page"`  // 需要从中下载折扣申请列表的页面。
	Limit  int64                `json:"limit"` // 页面上申请最大数量。
	Status V1DiscountTaskStatus `json:"status"`
}

// 折扣类型： - `PERCENT`——百分比折扣； - `CURRENCY`——按金额折扣。
type V1SellerActionsCreateVoucherRequestDiscountTypeEnum string

// 群体类型： - `UNSPECIFIED`——未定义； - `OZON`——Ozon; - `SELLER`——卖家。
type PickedSegmentSegmentTypeEnum string

type PickedSegmentSegment struct {
	Description string                       `json:"description"` // 群体描述。
	ID          int64                        `json:"id"`          // 群体标识符。
	Name        string                       `json:"name"`        // 群体名称。
	Type        PickedSegmentSegmentTypeEnum `json:"type_"`
}

type ParameterPickedSegment struct {
	Segments []PickedSegmentSegment `json:"segments"` // 受众群体参数。
}

type V1SellerActionsArchiveRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
}

// 定量包装类型： - `UNSPECIFIED`——未指定， - `BOX`——箱， - `PALLET`——托盘， - `GENERAL`——商品。
type V1SellerActionsProductsCandidatesResponseProductQuantTypeEnum string

type V1SellerActionsProductsCandidatesResponseProduct struct {
	QuantType       V1SellerActionsProductsCandidatesResponseProductQuantTypeEnum `json:"quant_type"`
	SKU             []string                                                      `json:"sku"`              // Ozon系统中的商品标识符——SKU。
	ActionPrice     float64                                                       `json:"action_price"`     // 计入促销活动后的商品价格。
	Currency        string                                                        `json:"currency"`         // 货币单位。
	DiscountPercent float64                                                       `json:"discount_percent"` // 折扣百分比。
	IsActive        bool                                                          `json:"is_active"`        // `true`，表示商品参与活动。
	MinSellerPrice  float64                                                       `json:"min_seller_price"` // 用于将商品自动添加到促销活动中的最低价格。
	Name            string                                                        `json:"name"`             // 商品名称。
	OfferID         string                                                        `json:"offer_id"`         // 卖家系统中的商品标识符——货号。
	ProductID       int64                                                         `json:"product_id"`       // Ozon系统中的商品标识符——`product_id`。
	BasePrice       float64                                                       `json:"base_price"`       // 商品未参与促销活动时在Ozon上的基本价格。
	Price           float64                                                       `json:"price"`            // 针对买家的商品价格。
	QuantSize       int64                                                         `json:"quant_size"`       // 定量包装大小。
}

type V1SellerActionsProductsCandidatesResponse struct {
	Cursor   int64                                              `json:"cursor"`   // 用于选择下一批数据的指针。
	HasNext  bool                                               `json:"has_next"` // 响应中仅返回了部分值的标志： - `true`——请使用新的`cursor`参数重复请求，以获取其余值； - `false`——响应中已包含所有值。
	Products []V1SellerActionsProductsCandidatesResponseProduct `json:"products"` // 商品信息。
}

// 促销码类型： - `ONE`——面向所有买家的促销码，可用于1个订单； - `MULTIPLE`——面向所有买家的促销码，可用于任意数量的订单； - `UNIQUE`——面向1位买家的促销码，可用于1个订单。
type V1SellerActionsCreateVoucherRequestVoucherParameterTypeEnum string

// 促销码参数。
type V1SellerActionsCreateVoucherRequestVoucherParameter struct {
	CountCodes int64                                                       `json:"count_codes"` // 促销码数量。
	IsPrivate  bool                                                        `json:"is_private"`  // `true`，表示促销码公开可用。
	Type       V1SellerActionsCreateVoucherRequestVoucherParameterTypeEnum `json:"type_"`
}

type V1SellerActionsCreateVoucherRequest struct {
	VoucherParameters V1SellerActionsCreateVoucherRequestVoucherParameter `json:"voucher_parameters"`
	Budget            int64                                               `json:"budget"`     // 促销活动预算。预算用尽后，促销活动将停止。
	DateEnd           string                                              `json:"date_end"`   // 促销活动结束日期与时间。
	DateStart         string                                              `json:"date_start"` // 促销活动开始日期与时间。
	DiscountType      V1SellerActionsCreateVoucherRequestDiscountTypeEnum `json:"discount_type"`
	DiscountValue     float64                                             `json:"discount_value"` // 折扣幅度。
	Title             string                                              `json:"title"`          // 促销活动名称。
	UserIds           []string                                            `json:"user_ids"`       // 可使用该促销码的用户标识符列表。
}

// 促销活动状态： - `ACTIVE`——活跃； - `ENDED`——已结束； - `PLANNED`——已计划； - `PAUSED`——已暂停。
type ActionsListRequestStatusEnum string

type V1SellerActionsCreateDiscountWithConditionResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

// 货币： - `RUB`——俄罗斯卢布； - `BYN`——白俄罗斯卢布； - `KZT`——坚戈； - `EUR`——欧元； - `USD`——美元； - `CNY`——人民币。
type CurrencyEnum string

type ApproveDeclineDiscountTasksResponseFailDetail struct {
	TaskID       int64  `json:"task_id"`        // 申请ID。
	ErrorForUser string `json:"error_for_user"` // 错误文本。
}

// 方法操作结果。
type V1ApproveDeclineDiscountTasksResponseResult struct {
	FailDetails  []ApproveDeclineDiscountTasksResponseFailDetail `json:"fail_details"`  // 创建申请时的错误。
	SuccessCount int32                                           `json:"success_count"` // 状态更改成功的申请数量。
	FailCount    int32                                           `json:"fail_count"`    // 未能更改状态的申请数量。
}

type ActionsV1ActionsAutoAddProductsUpdateResponseExtremelyLowPrice struct {
	Key   int64   `json:"key"`   // Ozon系统中的商品标识符——`product_id`。
	Value float64 `json:"value"` // 商品价格。
}

// 促销活动参数。
type V1SellerActionsUpdateDiscountWithConditionRequestActionParameters struct {
	DateEnd        string  `json:"date_end"`         // 促销活动结束日期与时间。
	DateStart      string  `json:"date_start"`       // 促销活动开始日期与时间。
	DiscountValue  float64 `json:"discount_value"`   // 折扣幅度。
	MinOrderAmount float64 `json:"min_order_amount"` // 获得折扣所需的最低订单金额。
	Title          string  `json:"title"`            // 促销活动名称。
}

type V1ApproveDiscountTasksRequestTask struct {
	SellerComment       string  `json:"seller_comment"`        // 卖家对申请的评论。
	ApprovedQuantityMin int64   `json:"approved_quantity_min"` // 批准的最小商品数量。
	ApprovedQuantityMax int64   `json:"approved_quantity_max"` // 批准的最大商品数量。
	ID                  int64   `json:"id"`                    // 申请ID。可以使用方法 [/v1/actions/discounts-task/list](#operation/promos_task_list)获取。
	ApprovedPrice       float64 `json:"approved_price"`        // 同意的价格。
}

type V1ApproveDiscountTasksRequest struct {
	Tasks []V1ApproveDiscountTasksRequestTask `json:"tasks"` // 申请列表。
}

type V1SellerActionsCreateInstallmentRequest struct {
	Title     string `json:"title"`      // 促销活动名称。
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
}

type SellerApiProductV1ResponseProductDeactivate struct {
	ProductID float64 `json:"product_id"` // 商品识别号。
	Reason    string  `json:"reason"`     // 该商品未从促销活动中删除的原因。
}

// 促销活动参数。
type V1SellerActionsUpdateVoucherRequestActionParameters struct {
	Budget        int64    `json:"budget"`         // 促销活动预算。在修改预算前，请先通过方法[/v1/seller-actions/change-activity](#operation/SellerActionsChangeActivity)关闭该促销活动。
	DateEnd       string   `json:"date_end"`       // 促销活动结束日期与时间。
	DateStart     string   `json:"date_start"`     // 促销活动开始日期与时间。
	DiscountValue float64  `json:"discount_value"` // 折扣幅度。
	Title         string   `json:"title"`          // 促销活动名称。
	UserIds       []string `json:"user_ids"`       // 可使用该促销码的用户标识符列表。
}

type V1SellerActionsUpdateVoucherRequest struct {
	ActionID         int64                                               `json:"action_id"` // 促销活动标识符。
	ActionParameters V1SellerActionsUpdateVoucherRequestActionParameters `json:"action_parameters"`
}

type GetSellerActionsV1ResponseAction struct {
	Description                string   `json:"description"`                  // 活动描述。
	DateStart                  string   `json:"date_start"`                   // 活动开始日期。
	FreezeDate                 string   `json:"freeze_date"`                  // 活动暂停的日期。 如果该空白被填写，卖家就不能提高价格，改变商品清单或减少促销活动的单位数量。 卖方可以降低价格，增加促销的单位数量。
	IsParticipating            bool     `json:"is_participating"`             // 无论你是否参加这项活动。
	WithTargeting              bool     `json:"with_targeting"`               // 此迹象表明该活动是与目标受众一起进行的。
	ID                         float64  `json:"id"`                           // 活动识别号。
	DateEnd                    string   `json:"date_end"`                     // 活动结束日期。
	IsVoucherAction            bool     `json:"is_voucher_action"`            // 此迹象表明买家需要促销代码才能参加。
	BannedProductsCount        float64  `json:"banned_products_count"`        // 被封商品数量。
	OrderAmount                float64  `json:"order_amount"`                 // 预定金额。
	Title                      string   `json:"title"`                        // 活动名称。
	ParticipatingProductsCount float64  `json:"participating_products_count"` // 参加促销的商品数量。
	DiscountType               string   `json:"discount_type"`                // 折扣类型。
	DiscountValue              float64  `json:"discount_value"`               // 折扣力度。
	ActionType                 string   `json:"action_type"`                  // 活动类型。
	AutoAddDates               []string `json:"auto_add_dates"`               // 商品自动添加到促销活动中的日期和时间。
	PotentialProductsCount     float64  `json:"potential_products_count"`     // 可供活动的商品数量。
}

type SellerApiGetSellerActionsV1Response struct {
	Result []GetSellerActionsV1ResponseAction `json:"result"` // 请求结果。
}

type SellerApiProductV1ResponseProduct struct {
	ProductID float64 `json:"product_id"` // 商品识别号。
	Reason    string  `json:"reason"`     // 该商品未被加入促销活动的原因。
}

// 请求结果。
type SellerApiProductV1ResponseResult struct {
	ProductIds []float64                           `json:"product_ids"` // 已添加到促销活动中的商品ID列表。
	Rejected   []SellerApiProductV1ResponseProduct `json:"rejected"`    // 无法添加到促销活动中的商品列表。
}

type SellerApiProductV1Response struct {
	Result SellerApiProductV1ResponseResult `json:"result"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseBelowMinPrice struct {
	Key   int64   `json:"key"`   // Ozon系统中的商品标识符——`product_id`。
	Value float64 `json:"value"` // 商品价格。
}

type ActionsV1ActionsAutoAddProductsUpdateResponseFailedPrice struct {
	Key   int64   `json:"key"`   // Ozon系统中的商品标识符——`product_id`。
	Value float64 `json:"value"` // 问题价格的值： - 如果商品价格高于`max_discount_price`，则为参数`max_discount_price`的值； - 如果商品折扣幅度超过95%，则为优惠95%后的价格。
}

// 错误代码： - `NOT_FOUND`——未找到商品； - `NO_CHANGES`——无变更； - `STOCK_REQUIRED`——商品库存不足，无法参与活动； - `INVALID_ACTION_PRICE`——促销活动价格不正确；...
type ActionsV1ActionsAutoAddProductsUpdateResponseRejectedCodeEnum string

type ActionsV1ActionsAutoAddProductsUpdateResponseRejected struct {
	Code      ActionsV1ActionsAutoAddProductsUpdateResponseRejectedCodeEnum `json:"code"`
	ProductID int64                                                         `json:"product_id"` // Ozon系统中的商品标识符——`product_id`。
	Reason    string                                                        `json:"reason"`     // 未能添加或更新商品的原因。
}

type ActionsV1ActionsAutoAddProductsUpdateResponse struct {
	FailedPrice       []ActionsV1ActionsAutoAddProductsUpdateResponseFailedPrice       `json:"failed_price"`        // 未通过价格校验的商品列表。
	Rejected          []ActionsV1ActionsAutoAddProductsUpdateResponseRejected          `json:"rejected"`            // 未能添加或更新的商品ID。
	UpdatedIds        []string                                                         `json:"updated_ids"`         // 已成功添加或更新的商品ID。
	BelowMinPrice     []ActionsV1ActionsAutoAddProductsUpdateResponseBelowMinPrice     `json:"below_min_price"`     // 价格低于最低价格的商品列表。
	ExtremelyLowPrice []ActionsV1ActionsAutoAddProductsUpdateResponseExtremelyLowPrice `json:"extremely_low_price"` // 折扣幅度超过70%的商品列表。
}

type V1SellerActionsCreateInstallmentResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

// 请求结果。
type SellerApiProductV1ResponseResultDeactivate struct {
	ProductIds []float64                                     `json:"product_ids"` // 已从促销活动中删除的商品ID列表。
	Rejected   []SellerApiProductV1ResponseProductDeactivate `json:"rejected"`    // 不能从促销活动中删除的商品清单。
}

type V1SellerActionsCreateMultiLevelDiscountResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type SellerApiProduct struct {
	ID                        float64 `json:"id"`                            // 商品识别号。
	Price                     float64 `json:"price"`                         // 不含折扣的商品的当前价格。
	CurrentBoost              float64 `json:"current_boost"`                 // 商品提升幅度。
	PriceMinElastic           float64 `json:"price_min_elastic"`             // 最小提升幅度对应的商品价格。
	PriceMaxElastic           float64 `json:"price_max_elastic"`             // 最大提升幅度对应的商品价格。
	MaxBoost                  float64 `json:"max_boost"`                     // 最大提升幅度（%）。
	ActionPrice               float64 `json:"action_price"`                  // 促销价格。
	AlertMaxActionPriceFailed bool    `json:"alert_max_action_price_failed"` // 如果商品价格高于建议价，则为`true`。商品被标记为红色，可能会被排除在促销活动之外。
	AlertMaxActionPrice       float64 `json:"alert_max_action_price"`        // 促销商品建议价格。
	MaxActionPrice            float64 `json:"max_action_price"`              // 可能的最高促销价格。
	AddMode                   string  `json:"add_mode"`                      // 添加到促销中的商品类型：自动或由卖家手动添加。
	MinStock                  float64 `json:"min_stock"`                     // 库存折扣促销中的最小商品数。
	Stock                     float64 `json:"stock"`                         // 《库存折扣》促销中的商品单位数量。
	MinBoost                  float64 `json:"min_boost"`                     // 最小提升幅度（%）。
}

type V1SellerActionsChangeActivityRequest struct {
	ActionID int64 `json:"action_id"`  // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	IsTurnOn bool  `json:"is_turn_on"` // `true`，用于启用促销活动。
}

// 折扣类型： - `PERCENT`——百分比折扣； - `CURRENCY`——按金额折扣。
type V1SellerActionsCreateMultiLevelDiscountRequestDiscountTypeEnum string

// 定量包装类型： - `UNSPECIFIED`——未指定， - `BOX`——箱， - `PALLET`——托盘， - `GENERAL`——商品。
type V1SellerActionsProductsListResponseProductQuantTypeEnum string

type V1SellerActionsProductsListResponseProduct struct {
	IsActive        bool                                                    `json:"is_active"`  // `true`，表示商品参与活动。
	QuantSize       int64                                                   `json:"quant_size"` // 定量包装大小。
	QuantType       V1SellerActionsProductsListResponseProductQuantTypeEnum `json:"quant_type"`
	SKU             []string                                                `json:"sku"`              // Ozon系统中的商品标识符——SKU。
	BasePrice       float64                                                 `json:"base_price"`       // 商品未参与促销活动时在Ozon上的基本价格。
	DiscountPercent float64                                                 `json:"discount_percent"` // 折扣百分比。
	MinSellerPrice  float64                                                 `json:"min_seller_price"` // 用于将商品自动添加到促销活动中的最低价格。
	Name            string                                                  `json:"name"`             // 商品名称。
	OfferID         string                                                  `json:"offer_id"`         // 卖家系统中的商品标识符——货号。
	Price           float64                                                 `json:"price"`            // 针对买家的商品价格。
	ProductID       int64                                                   `json:"product_id"`       // Ozon系统中的商品标识符——`product_id`。
	ActionPrice     float64                                                 `json:"action_price"`     // 计入促销活动后的商品价格。
	Currency        string                                                  `json:"currency"`         // 货币单位。
}

// 促销活动参数。
type V1SellerActionsUpdateInstallmentRequestActionParameters struct {
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
	Title     string `json:"title"`      // 促销活动名称。
}

// 折扣类型： - `PERCENT`——百分比折扣； - `CURRENCY`——按金额折扣。
type V1SellerActionsCreateDiscountWithConditionRequestDiscountTypeEnum string

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"` // 折扣幅度。
	OrderAmount   float64 `json:"order_amount"`   // 订单最低金额。
}

type SellerApiProductIDsV1Request struct {
	ActionID   float64   `json:"action_id"`   // 活动识别号。可以使用方法 [/v1/actions](#operation/Promos)获取。
	ProductIds []float64 `json:"product_ids"` // 活动识别号清单。
}

// AddMode values
type AddMode string

const (
	AddModeAuto   AddMode = "AUTO"   // 自动；
	AddModeManual AddMode = "MANUAL" // 卖家手动。
)

type ActionsV1ActionsAutoAddProductsListResponseProducts struct {
	AddMode                AddMode `json:"add_mode"`                 // 商品加入促销活动的方式： - `AUTO`——自动； - `MANUAL`——卖家手动。
	Currency               string  `json:"currency"`                 // 价格货币。
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"` // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	MinSellerPrice         float64 `json:"min_seller_price"`         // 应用促销活动后的商品最低价格。
	Name                   string  `json:"name"`                     // 商品名称。
	ProductID              int64   `json:"product_id"`               // Ozon系统中的商品标识符——`product_id`。
	QuantityToAutoAdd      int64   `json:"quantity_to_auto_add"`     // 促销活动中的商品数量。
	SKU                    int64   `json:"sku"`                      // Ozon系统中的商品标识符，即SKU。
	ActionPriceToAutoAdd   float64 `json:"action_price_to_auto_add"` // 商品的促销价格。
	MaxDiscountPrice       float64 `json:"max_discount_price"`       // 商品可自动添加到促销活动中的最高价格。
	MinActionQuantity      int64   `json:"min_action_quantity"`      // “库存折扣”促销活动类型中的最低商品数量。
	OfferID                string  `json:"offer_id"`                 // 卖家系统中的商品标识符，即货号。
	Price                  float64 `json:"price"`                    // 商品无折扣价格。
}

type V1SellerActionsProductsAddRequestProduct struct {
	DiscountPercent float64      `json:"discount_percent"` // 百分比折扣值。如果促销活动机制为"折扣"，请传入该参数。
	SKU             int64        `json:"sku"`              // Ozon系统中的商品标识符——SKU。
	Currency        CurrencyEnum `json:"currency"`
}

type V1SellerActionsCreateMultiLevelDiscountRequest struct {
	DateStart              string                                                         `json:"date_start"`      // 促销活动开始日期与时间。
	DiscountLevels         []V1SellerActionsCreateMultiLevelDiscountRequestDiscountLevel  `json:"discount_levels"` // 折扣等级。
	DiscountType           V1SellerActionsCreateMultiLevelDiscountRequestDiscountTypeEnum `json:"discount_type"`
	IsLegalEntitiesSegment bool                                                           `json:"is_legal_entities_segment"` // `true`，表示促销活动仅面向法人实体。
	Title                  string                                                         `json:"title"`                     // 促销活动名称。
	DateEnd                string                                                         `json:"date_end"`                  // 促销活动结束日期与时间。
}

type V1DeclineDiscountTasksRequestTask struct {
	ID            int64  `json:"id"`             // 申请ID。
	SellerComment string `json:"seller_comment"` // 卖家对申请的评价。
}

type ActionsV1ActionsAutoAddProductsListResponse struct {
	Products []ActionsV1ActionsAutoAddProductsListResponseProducts `json:"products"` // 启用自动添加的商品列表。
	Total    int64                                                 `json:"total"`    // 商品总数。
}

type SellerApiProductPrice struct {
	ProductID   float64 `json:"product_id"`   // 商品识别号。
	ActionPrice float64 `json:"action_price"` // 商品活动期间的价格。
	Stock       float64 `json:"stock"`        // 《库存折扣》促销中的商品单位数量。
}

type ActionsV1ActionsAutoAddProductsDeleteResponse struct {
	ProductIds []string `json:"product_ids"` // 已从自动添加中删除的商品ID。
}

type V1SellerActionsVoucherGetRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
}

// 促销活动停止原因： - `UNSPECIFIED`——未定义； - `BUDGET_EXCEEDED`——预算已用尽； - `FAST_BUDGET_SPENDING`——前几小时预算消耗过快； - `BUDGET_SPENDING_IS_...
type ParameterAutoStopActionReasonEnum string

// 折扣类型： - `UNSPECIFIED`——未定义； - `RUB`——以卢布计的折扣； - `PERCENT`——百分比折扣； - `FINAL_PRICE`——最终价格； - `INSTALLED_PERIOD`——分期付款周期； -...
type ActionParameterDiscountTypeEnum string

type ActionsV1ActionsAutoAddProductsUpdateRequestToUpdate struct {
	Quantity    int64   `json:"quantity"`     // 促销活动中的商品数量。
	ActionPrice float64 `json:"action_price"` // 商品的促销价格。
	Currency    string  `json:"currency"`     // 价格货币。
	ProductID   int64   `json:"product_id"`   // 需要添加或更新的商品ID。
}

type V1SellerActionsCreateDiscountRequest struct {
	DateStart        string  `json:"date_start"`         // 促销活动开始日期与时间。
	MinActionPercent float64 `json:"min_action_percent"` // 最低折扣百分比。
	Title            string  `json:"title"`              // 促销活动名称。
	DateEnd          string  `json:"date_end"`           // 促销活动结束日期与时间。
}

type V1SellerActionsUpdateDiscountWithConditionRequest struct {
	ActionID         int64                                                             `json:"action_id"` // 促销活动标识符。
	ActionParameters V1SellerActionsUpdateDiscountWithConditionRequestActionParameters `json:"action_parameters"`
}

// 促销活动参数。
type V1SellerActionsUpdateDiscountRequestActionParameters struct {
	Title     string `json:"title"`      // 促销活动名称。
	DateEnd   string `json:"date_end"`   // 促销活动结束日期与时间。
	DateStart string `json:"date_start"` // 促销活动开始日期与时间。
}

type V1SellerActionsUpdateDiscountRequest struct {
	ActionID         int64                                                `json:"action_id"` // 促销活动标识符。
	ActionParameters V1SellerActionsUpdateDiscountRequestActionParameters `json:"action_parameters"`
}

type V1GetDiscountTaskListResponseTask struct {
	UserComment             string  `json:"user_comment"`               // 买家对申请的评论。
	MinAutoPrice            float64 `json:"min_auto_price"`             // 自动应用折扣和促销后的最低价格值。
	ApprovedDiscount        float64 `json:"approved_discount"`          // 卖家同意的以卢布显示的折扣。 如果卖家不批准订单，则传递值“0”。
	IsPurchased             bool    `json:"is_purchased"`               // 用户是否购买了商品。 `true`，如果购买。
	IsAutoModerated         bool    `json:"is_auto_moderated"`          // 申请是否自动审核。 `true`，如果自动审核。
	ModeratedAt             string  `json:"moderated_at"`               // 审核日期：审核、批准或拒绝申请。
	LastName                string  `json:"last_name"`                  // 处理申请的卖家员工姓氏。
	Patronymic              string  `json:"patronymic"`                 // 处理请求的卖家员工父称。
	ApprovedQuantityMin     int64   `json:"approved_quantity_min"`      // 商品数量批准的最小值。
	RequestedQuantityMin    int64   `json:"requested_quantity_min"`     // 商品请求数量最小值。
	RequestedPriceWithFee   float64 `json:"requested_price_with_fee"`   // 带有区域加价的价格申请。
	EndAt                   string  `json:"end_at"`                     // 申请到期时间。
	RequestedPrice          float64 `json:"requested_price"`            // 申请价格。
	BasePrice               float64 `json:"base_price"`                 // 如果不参与促销活动，商品在Ozon上销售的基础价。
	FirstName               string  `json:"first_name"`                 // 处理请求的卖家员工姓名。
	ID                      int64   `json:"id"`                         // 申请ID。
	SellerComment           string  `json:"seller_comment"`             // 卖家对申请的评论。
	ApprovedPrice           float64 `json:"approved_price"`             // 批准的价格。
	PrevTaskID              int64   `json:"prev_task_id"`               // 该商品买家先前申请ID。
	ApprovedPriceWithFee    float64 `json:"approved_price_with_fee"`    // 批准的含区域加价的价格。
	ApprovedPriceFeePercent float64 `json:"approved_price_fee_percent"` // 按百分比显示的区域加价。
	OfferID                 string  `json:"offer_id"`                   // 卖家系统中的商品标识符是商品货号。
	EditedTill              string  `json:"edited_till"`                // 决定改变时间。
	OriginalPrice           float64 `json:"original_price"`             // 折扣前的商品价格。
	Discount                float64 `json:"discount"`                   // 卢布折扣。
	IsDamaged               bool    `json:"is_damaged"`                 // 商品是否打折。如果打折，`true`。
	ApprovedQuantityMax     int64   `json:"approved_quantity_max"`      // 商品数量批准的最大值。
	DiscountPercent         float64 `json:"discount_percent"`           // 折扣百分比。
	ApprovedDiscountPercent float64 `json:"approved_discount_percent"`  // 卖家批准的折扣百分比。请传递值 `0`，如果卖家不批准申请。
	RequestedQuantityMax    int64   `json:"requested_quantity_max"`     // 商品请求数量最大值。
	CustomerName            string  `json:"customer_name"`              // 买家姓名。
	SKU                     int64   `json:"sku"`                        // Ozon系统中的商品ID —— SKU。
	Email                   string  `json:"email"`                      // 处理请求的卖家员工电子邮件地址。
	CreatedAt               string  `json:"created_at"`                 // 申请创建日期。
	Status                  string  `json:"status"`                     // 申请状态。
}

type V1ApproveDeclineDiscountTasksResponse struct {
	Result V1ApproveDeclineDiscountTasksResponseResult `json:"result"`
}

type V1SellerActionsVoucherGetResponse struct {
	File string `json:"file"` // 促销码CSV文件链接。
}

type ActionsV1ActionsAutoAddProductsListRequest struct {
	ActionID    int64  `json:"action_id"`     // 促销活动标识符。
	AutoAddDate string `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
	Limit       int64  `json:"limit"`         // 响应中返回的值数量。
	Offset      int64  `json:"offset"`        // 在响应中将被跳过的项目数量。例如，如果`offset = 10`，响应将从第11个找到的元素开始。
}

type ActionsV1ActionsAutoAddProductsCandidatesRequest struct {
	Limit       int64  `json:"limit"`         // 响应中返回的值数量。
	Offset      int64  `json:"offset"`        // 在响应中将被跳过的项目数量。例如，如果`offset = 10`，响应将从第11个找到的元素开始。
	ActionID    int64  `json:"action_id"`     // 促销活动标识符。
	AutoAddDate string `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
}

// 促销活动机制： - `DISCOUNT`——折扣； - `VOUCHER_DISCOUNT`——促销码折扣； - `DISCOUNT_WITH_CONDITION`——基于订单总额的折扣； - `INSTALLMENT`——免息分期付款； ...
type ActionsListRequestActionTypeEnum string

// ActionType values
type ActionType string

const (
	ActionTypeDiscount                     ActionType = "DISCOUNT"                        // 折扣；
	ActionTypeVoucherDiscount              ActionType = "VOUCHER_DISCOUNT"                // 促销码折扣；
	ActionTypeDiscountWithCondition        ActionType = "DISCOUNT_WITH_CONDITION"         // 基于订单总额的折扣；
	ActionTypeInstallment                  ActionType = "INSTALLMENT"                     // 免息分期付款；
	ActionTypeIndividualDiscountBYProducts ActionType = "INDIVIDUAL_DISCOUNT_BY_PRODUCTS" // 卖家积分；
	ActionTypeOzonAccountDiscount          ActionType = "OZON_ACCOUNT_DISCOUNT"           // Ozon银行卡专享额外折扣；
	ActionTypeMultiLevelDiscountONAmount   ActionType = "MULTI_LEVEL_DISCOUNT_ON_AMOUNT"  // 多级满额折扣。
)

// Status values
type Status string

const (
	StatusActive  Status = "ACTIVE"  // 活跃；
	StatusEnded   Status = "ENDED"   // 已结束；
	StatusPlanned Status = "PLANNED" // 已计划；
	StatusPaused  Status = "PAUSED"  // 已暂停。
)

type V1SellerActionsListRequest struct {
	Status     []ActionsListRequestStatusEnum     `json:"status"`      // 促销活动状态： - `ACTIVE`—— 活跃； - `ENDED`——已结束； - `PLANNED`——已计划； - `PAUSED`——已暂停。
	ActionIds  []string                           `json:"action_ids"`  // 促销活动标识符列表。
	ActionType []ActionsListRequestActionTypeEnum `json:"action_type"` // 促销活动机制： - `DISCOUNT`——折扣； - `VOUCHER_DISCOUNT`——促销码折扣； - `DISCOUNT_WITH_CONDITION`——基于订单总额的折扣； - `INSTALLMENT`——免息分期付款； ...
	Limit      int64                              `json:"limit"`       // 每页显示的数量。
	Offset     int64                              `json:"offset"`      // 在响应中将被跳过的项目数量。例如，当`offset = 10`时，响应将从第11个找到的元素开始。
	Search     string                             `json:"search"`      // 按促销活动名称搜索。
}

type V1SellerActionsCreateDiscountResponse struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。
}

type ActionsV1ActionsAutoAddProductsUpdateRequest struct {
	ActionID    int64                                                  `json:"action_id"`     // 促销活动标识符。
	AutoAddDate string                                                 `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
	ToUpdate    []ActionsV1ActionsAutoAddProductsUpdateRequestToUpdate `json:"to_update"`     // 需要添加到自动添加中或在自动添加中更新的商品列表。
}

type V1SellerActionsCreateDiscountWithConditionRequest struct {
	DateEnd        string                                                            `json:"date_end"`   // 促销活动结束日期与时间。
	DateStart      string                                                            `json:"date_start"` // 促销活动开始日期与时间。
	DiscountType   V1SellerActionsCreateDiscountWithConditionRequestDiscountTypeEnum `json:"discount_type"`
	DiscountValue  float64                                                           `json:"discount_value"`   // 折扣幅度。
	MinOrderAmount float64                                                           `json:"min_order_amount"` // 折扣生效的订单金额门槛。
	Title          string                                                            `json:"title"`            // 促销活动名称。
}

type ActionsUpdateMultiLevelDiscountRequestActionParametersDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"` // 折扣幅度。
	OrderAmount   float64 `json:"order_amount"`   // 获得折扣所需的最低订单金额。
}

// 促销活动参数。
type V1SellerActionsUpdateMultiLevelDiscountRequestActionParameters struct {
	Title                  string                                                                `json:"title"`                     // 促销活动名称。
	DateEnd                string                                                                `json:"date_end"`                  // 促销活动结束日期与时间。
	DateStart              string                                                                `json:"date_start"`                // 促销活动开始日期与时间。
	DiscountLevels         []ActionsUpdateMultiLevelDiscountRequestActionParametersDiscountLevel `json:"discount_levels"`           // 折扣等级。
	IsLegalEntitiesSegment bool                                                                  `json:"is_legal_entities_segment"` // `true`，表示促销活动仅面向法人实体。
}

type V1SellerActionsUpdateMultiLevelDiscountRequest struct {
	ActionID         int64                                                          `json:"action_id"` // 促销活动标识符。
	ActionParameters V1SellerActionsUpdateMultiLevelDiscountRequestActionParameters `json:"action_parameters"`
}

type V1SellerActionsProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Cursor   int64 `json:"cursor"`    // 用于选择下一批数据的指针。
	Limit    int64 `json:"limit"`     // 响应中的最大元素数量。
}

type V1GetDiscountTaskListResponse struct {
	Result []V1GetDiscountTaskListResponseTask `json:"result"` // 申请列表。
}

// 促销活动参数。
type ActionParameter struct {
	Type                   ActionsListRequestActionTypeEnum  `json:"type_"`
	VoucherParameters      ActionParameterVoucherParameter   `json:"voucher_parameters"`
	Addresses              []string                          `json:"addresses"` // 地址列表。
	AutoStopActionReason   ParameterAutoStopActionReasonEnum `json:"auto_stop_action_reason"`
	DiscountValue          float64                           `json:"discount_value"`            // 折扣幅度。 如果活动机制为"免息分期付款"，将返回分期周期。
	PickedSegments         []ParameterPickedSegment          `json:"picked_segments"`           // 受众群体列表。
	Budget                 float64                           `json:"budget"`                    // 促销活动预算。预算用尽后，促销活动将停止。
	DateStart              string                            `json:"date_start"`                // 促销活动开始日期与时间。
	MinOrderAmount         float64                           `json:"min_order_amount"`          // 折扣生效的订单金额门槛。
	Warehouses             []string                          `json:"warehouses"`                // 仓库列表。
	BudgetSpent            float64                           `json:"budget_spent"`              // 促销活动预算支出。
	DateEnd                string                            `json:"date_end"`                  // 促销活动结束日期与时间。
	IsLegalEntitiesSegment bool                              `json:"is_legal_entities_segment"` // `true`，表示该促销活动面向企业客户。
	MinActionPercent       float64                           `json:"min_action_percent"`        // 最低折扣百分比。
	Title                  string                            `json:"title"`                     // 促销活动名称。
	DiscountLevels         []ParameterDiscountLevels         `json:"discount_levels"`           // 折扣等级。
	DiscountType           ActionParameterDiscountTypeEnum   `json:"discount_type"`
	Status                 ActionsListRequestStatusEnum      `json:"status"`
}

type ActionsListResponseAction struct {
	IsEditable       bool            `json:"is_editable"`     // `true`，表示该促销活动可编辑。
	IsParticipated   bool            `json:"is_participated"` // `true`，表示该促销活动中至少已添加1件商品。
	IsTurnOn         bool            `json:"is_turn_on"`      // `true`，表示该活动已启用。
	SKUCount         int64           `json:"sku_count"`       // 促销活动中的商品总数。
	ActionID         int64           `json:"action_id"`       // 促销活动标识符。
	ActionParameters ActionParameter `json:"action_parameters"`
	AllowDelete      bool            `json:"allow_delete"`  // `true`，表示该促销活动可删除。
	HighlightURL     string          `json:"highlight_url"` // 高亮链接。
}

type V1SellerActionsListResponse struct {
	Actions []ActionsListResponseAction `json:"actions"` // 促销活动列表。
	Total   int64                       `json:"total"`   // 促销活动总数。
}

type V1SellerActionsProductsListResponse struct {
	HasNext  bool                                         `json:"has_next"` // 响应中仅返回了部分值的标志： - `true`——请使用新的`cursor`参数重复请求，以获取其余值； - `false`——响应中已包含所有值。
	Products []V1SellerActionsProductsListResponseProduct `json:"products"` // 商品信息。
	Cursor   int64                                        `json:"cursor"`   // 用于选择下一批数据的指针。
}

type SellerApiActivateProductV1Request struct {
	ActionID float64                 `json:"action_id"` // 活动识别号。可以使用方法 [/v1/actions](#operation/Promos)获取。
	Products []SellerApiProductPrice `json:"products"`  // 商品清单。
}

type V1SellerActionsProductsAddRequest struct {
	ActionID int64                                      `json:"action_id"` // 促销活动标识符。请通过方法[/v1/seller-actions/list](#operation/SellerActionsList)获取该参数的值。
	Products []V1SellerActionsProductsAddRequestProduct `json:"products"`  // 商品信息。
}

type V1SellerActionsUpdateInstallmentRequest struct {
	ActionID         int64                                                   `json:"action_id"` // 促销活动标识符。
	ActionParameters V1SellerActionsUpdateInstallmentRequestActionParameters `json:"action_parameters"`
}

type SellerApiProductV1ResponseDeactivate struct {
	Result SellerApiProductV1ResponseResultDeactivate `json:"result"`
}

type V1DeclineDiscountTasksRequest struct {
	Tasks []V1DeclineDiscountTasksRequestTask `json:"tasks"` // 申请列表。
}

type ActionsV1ActionsAutoAddProductsCandidatesResponseProducts struct {
	OfferID                string  `json:"offer_id"`                 // 卖家系统中的商品标识符，即货号。
	Price                  float64 `json:"price"`                    // 商品无折扣价格。
	ActionPriceToAutoAdd   float64 `json:"action_price_to_auto_add"` // 商品的促销价格。
	BasePrice              float64 `json:"base_price"`               // 商品折扣前价格。
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"` // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	MinActionQuantity      int64   `json:"min_action_quantity"`      // “库存折扣”促销活动类型中的最低商品数量。
	ProductID              int64   `json:"product_id"`               // Ozon系统中的商品标识符——`product_id`。
	QuantityToAutoAdd      int64   `json:"quantity_to_auto_add"`     // 促销活动中的商品数量。
	SKU                    int64   `json:"sku"`                      // Ozon系统中的商品标识符，即SKU。
	Currency               string  `json:"currency"`                 // 价格货币。
	MaxDiscountPrice       float64 `json:"max_discount_price"`       // 商品可自动添加到促销活动中的最高价格。
	MinSellerPrice         float64 `json:"min_seller_price"`         // 应用促销活动后的商品最低价格。
	Name                   string  `json:"name"`                     // 商品名称。
}

type ActionsV1ActionsAutoAddProductsCandidatesResponse struct {
	Products []ActionsV1ActionsAutoAddProductsCandidatesResponseProducts `json:"products"` // 可用于自动添加到促销活动中的商品列表。
	Total    int64                                                       `json:"total"`    // 商品总数。
}

// 请求结果。
type SellerApiGetSellerProductV1ResponseResult struct {
	Products []SellerApiProduct `json:"products"` // 商品清单。
	Total    float64            `json:"total"`    // 可用于活动的商品总数。
	LastID   float64            `json:"last_id"`  // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
}

type SellerApiGetSellerProductV1Response struct {
	Result SellerApiGetSellerProductV1ResponseResult `json:"result"`
}

// 商品清单。
type SellerApiGetSellerProductV1Request struct {
	LastID   float64 `json:"last_id"`   // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。
	ActionID float64 `json:"action_id"` // 活动识别号。可以使用方法 [/v1/actions](#operation/Promos)获取。
	Limit    float64 `json:"limit"`     // 每页的答复数量。在默认情况下 — 100。
}

type ActionsV1ActionsAutoAddProductsDeleteRequest struct {
	ActionID    int64    `json:"action_id"`     // 促销活动标识符。
	AutoAddDate string   `json:"auto_add_date"` // 方法[/v1/actions](#operation/Promos)响应中`result.auto_add_dates`参数里的商品自动添加到促销活动中的日期和时间。
	ProductIds  []string `json:"product_ids"`   // Ozon系统中的商品标识符，即`product_id`。
}

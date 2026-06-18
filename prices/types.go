package prices

type V1ProductActionTimerStatusResponseStatuses struct {
	ExpiredAt                     string `json:"expired_at"`                         // 计时器结束时间。如果该参数为空，则没有有效的计时器。
	MinPriceForAutoActionsEnabled bool   `json:"min_price_for_auto_actions_enabled"` // 如果Ozon在添加商品至促销活动时会参考最低价格，则返回值为`true`。
	ProductID                     int64  `json:"product_id"`                         // 卖家系统中的商品识别符——`product_id`。
}

type V2ProductsStocksRequestStock struct {
	Stock       int64  `json:"stock"`        // 扣除预留库存后的可售商品数量。
	WarehouseID int64  `json:"warehouse_id"` // 得出的仓库编号的方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)。
	OfferID     string `json:"offer_id"`     // 卖家系统中的商品编号是 — 商品代码。
	ProductID   int64  `json:"product_id"`   // Ozon系统中商品的标识符 — `product_id`。
}

type V2ProductsStocksRequest struct {
	Stocks []V2ProductsStocksRequestStock `json:"stocks"` // 仓库中商品的信息。
}

type Sv1GetProductInfoStocksByWarehouseFbsResponseResult struct {
	WarehouseName string `json:"warehouse_name"` // 仓库名称。
	SKU           int64  `json:"sku"`            // Ozon系统中的商品识别码是SKU。
	OfferID       string `json:"offer_id"`       // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	Present       int64  `json:"present"`        // 库存商品总量。
	ProductID     int64  `json:"product_id"`     // 卖家系统中的Ozon系统中商品的标识符 — `product_id`。
	Reserved      int64  `json:"reserved"`       // 仓库中的保留商品的数量。
	WarehouseID   int64  `json:"warehouse_id"`   // 仓库编号。
}

type Sv1GetProductInfoStocksByWarehouseFbsResponse struct {
	Result []Sv1GetProductInfoStocksByWarehouseFbsResponseResult `json:"result"` // 该处理方法的结果。
}

// AutoActionEnabled values
type AutoActionEnabled string

const (
	AutoActionEnabledEnabled  AutoActionEnabled = "ENABLED"  // 启用
	AutoActionEnabledDisabled AutoActionEnabled = "DISABLED" // 关闭
	AutoActionEnabledUnknown  AutoActionEnabled = "UNKNOWN"  // 不做任何更改，默认赋值
)

// Vat values
type Vat string

const (
	VatV005 Vat = "0.05" // 5%
	VatV007 Vat = "0.07" // 7%
	VatV01  Vat = "0.1"  // 10%
	VatV02  Vat = "0.2"  // 20%
	VatV022 Vat = "0.22" // 22%
)

// CurrencyCode values
type CurrencyCode string

const (
	CurrencyCodeRUB CurrencyCode = "RUB" // 俄罗斯卢布
	CurrencyCodeBYN CurrencyCode = "BYN" // 白俄罗斯卢布
	CurrencyCodeKZT CurrencyCode = "KZT" // 坚戈
	CurrencyCodeUSD CurrencyCode = "USD" // 美元
	CurrencyCodeCNY CurrencyCode = "CNY" // 人民币
)

// PriceStrategyEnabled values
type PriceStrategyEnabled string

const (
	PriceStrategyEnabledEnabled  PriceStrategyEnabled = "ENABLED"  // 启用
	PriceStrategyEnabledDisabled PriceStrategyEnabled = "DISABLED" // 关闭
	PriceStrategyEnabledUnknown  PriceStrategyEnabled = "UNKNOWN"  // 不做任何更改，默认赋值
)

type ImportProductsPricesRequestPrice struct {
	MinPriceForAutoActionsEnabled     bool                 `json:"min_price_for_auto_actions_enabled"`    // `true`，如果为 Ozon 在添加到促销时会考虑最低价格。如果不传递任何值，价格计算的状态将保持不变。
	NetPrice                          string               `json:"net_price"`                             // 产品成本价。
	OfferID                           string               `json:"offer_id"`                              // 卖家系统中的商品标识符是商品货号。
	Price                             string               `json:"price"`                                 // 商品的价格，包括折扣，都显示在商品详情页上。 若参数`old_price`值大于0,则`price`与`old_price`之间应为具体差额。 差额取决于`price`值。 | `price`值 | 最小差额 | |---|---| | <...
	PriceStrategyEnabled              PriceStrategyEnabled `json:"price_strategy_enabled"`                // 用于自动应用价格策略的属性： - `ENABLED` — 启用; - `DISABLED` — 关闭; - `UNKNOWN` — 不做任何更改，默认赋值。 如果您之前已启用价格策略的自动应用，并且不希望关闭，请在后续请求中赋值 `UNKN...
	ProductID                         int64                `json:"product_id"`                            // Ozon系统中商品的标识符 — `product_id`。
	AutoActionEnabled                 AutoActionEnabled    `json:"auto_action_enabled"`                   // 启用和禁用自动应用活动的属性： - `ENABLED` — 启用; - `DISABLED` — 关闭; - `UNKNOWN` — 不做任何更改，默认赋值。 例如，如果你以前启用了活动的自动应用，并且不想关闭它，请提交`UNKNOWN`。...
	ManageElasticBoostingThroughPrice bool                 `json:"manage_elastic_boosting_through_price"` // 管理参与“弹性提升”促销活动： - `true`——如果参数 `price` 的数值符合促销活动条件，系统会自动将商品加入促销活动，或在促销活动中提高折扣与提升比例； - `false`——修改参数 `price` 的数值不会影响商品参与促...
	MinPrice                          string               `json:"min_price"`                             // 应用促销活动后的商品最低价格。<br> [在卖家知识库中了解更多关于最低价格的信息](https://global-help.ozon.com/zh/prices/price-control/)
	OldPrice                          string               `json:"old_price"`                             // 折扣前的价格（在商品卡上划掉）价格以卢布表。小数部分用一个点隔开，点后最多两个字符。 如果商品没有折扣，在这个字段中输入 "0"，并将当前价格提交到 "price"栏中。
	Vat                               Vat                  `json:"vat"`                                   // 商品增值税税率： - `0` — 免除增值税, - `0.05` — 5%, - `0.07` — 7%, - `0.1` — 10%, - `0.2` — 20%, - `0.22` — 22%。 传递当前有效的出价值。
	CurrencyCode                      CurrencyCode         `json:"currency_code"`                         // 您的价格的货币。提交的数值必须与你个人主页设置中的货币相匹配。默认情况下，会提交 `RUB` — 俄罗斯卢布。 例如，如果您设置的相互结算的货币是人民币，请提交数值 `CNY`，否则将返回错误。 可填的数值： - `RUB` — 俄罗斯卢布...
}

type V1GetProductInfoDiscountedResponseItem struct {
	Condition            string `json:"condition"`              // 商品的状态 — 新的或二手的。
	ConditionEstimation  string `json:"condition_estimation"`   // 商品的状况，以1至7分为标准。 - 1 — 令人满意。 - 2 — 良好。 - 3 — 非常好。 - 4 — 优秀。 - 5-7 — 像新的一样。
	DiscountedSKU        int64  `json:"discounted_sku"`         // 折扣商品的SKU。
	MechanicalDamage     string `json:"mechanical_damage"`      // 机械性损坏的说明。
	ReasonDamaged        string `json:"reason_damaged"`         // 损害原因。
	Repair               string `json:"repair"`                 // 商品已被修理的痕迹。
	SKU                  int64  `json:"sku"`                    // 主要商品的SKU。
	CommentReasonDamaged string `json:"comment_reason_damaged"` // 对损坏原因的评论。
	Defects              string `json:"defects"`                // 商品缺陷。
	PackageDamage        string `json:"package_damage"`         // 包装损坏的说明。
	PackagingViolation   string `json:"packaging_violation"`    // 篡改包装的痕迹。
	Shortage             string `json:"shortage"`               // 表示商品不完整。
	WarrantyType         string `json:"warranty_type"`          // 商品有有效保修的证明。
}

type V1GetProductInfoDiscountedResponse struct {
	Items []V1GetProductInfoDiscountedResponseItem `json:"items"` // 关于减价和主要商品的信息。
}

// 包装类型： - `SHIPMENT_TYPE_GENERAL` — 普通商品； - `SHIPMENT_TYPE_BOX` — 盒子； - `SHIPMENT_TYPE_PALLET` — 托盘。
type StockShipmentType string

const (
	StockShipmentTypeShipmentTypeGeneral StockShipmentType = "SHIPMENT_TYPE_GENERAL"
	StockShipmentTypeShipmentTypeBOX     StockShipmentType = "SHIPMENT_TYPE_BOX"
	StockShipmentTypeShipmentTypePallet  StockShipmentType = "SHIPMENT_TYPE_PALLET"
)

// Type values
type Type string

const (
	TypeFbs  Type = "fbs"  // 卖家仓库，由 Ozon 配送；
	TypeRfbs Type = "rfbs" // 卖家仓库，由卖家自行配送；
	TypeFbo  Type = "fbo"  // Ozon 仓库；
	TypeFbp  Type = "fbp"  // 合作伙伴仓库
)

type GetProductInfoStocksResponseStock struct {
	Reserved     int32             `json:"reserved"` // 已预定。
	ShipmentType StockShipmentType `json:"shipment_type"`
	SKU          int64             `json:"sku"`           // Ozon系统中的商品识别符——SKU。
	Type         Type              `json:"type_"`         // 仓库类型： - `fbs`——卖家仓库，由 Ozon 配送； - `rfbs`——卖家仓库，由卖家自行配送； - `fbo`——Ozon 仓库； - `fbp`——合作伙伴仓库。
	WarehouseIds []int64           `json:"warehouse_ids"` // 存放或曾经存放该商品的仓库标识符。
	Present      int32             `json:"present"`       // 现在在仓库中。
}

type V4GetProductInfoStocksResponseItem struct {
	OfferID   string                              `json:"offer_id"`   // 卖家系统中的商品识别符——货号。
	ProductID int64                               `json:"product_id"` // 商品识别符。
	Stocks    []GetProductInfoStocksResponseStock `json:"stocks"`     // 库存信息。
}

type V4GetProductInfoStocksResponse struct {
	Cursor string                               `json:"cursor"` // 后续数据的选择标志。
	Items  []V4GetProductInfoStocksResponseItem `json:"items"`  // 商品信息。
	Total  int32                                `json:"total"`  // 显示库存信息的独特商品数量。
}

// 其他平台上的竞争对手价格。
type PriceIndexesIndexExternalData struct {
	MinPrice         float64 `json:"min_price"`          // 竞争对手在其他平台上的最低商品价格。
	MinPriceCurrency string  `json:"min_price_currency"` // 价格的货币单位。
	PriceIndexValue  float64 `json:"price_index_value"`  // 价格指数值。
}

type V1ProductUpdateDiscountRequest struct {
	Discount  int32 `json:"discount"`   // 折扣力度：从3%到99%。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

// 商品可见性筛选: - `ALL` —— 所有的商品，除了已存档的。 - `VISIBLE` —— 买家可见的商品。 - `INVISIBLE` —— 买家不可见的商品。 - `EMPTY_STOCK` —— 库存未指定的商品。 - `NOT...
type V5GetProductListRequestFilterFilterVisibility string

const (
	V5GetProductListRequestFilterFilterVisibilityALL                    V5GetProductListRequestFilterFilterVisibility = "ALL"
	V5GetProductListRequestFilterFilterVisibilityVisible                V5GetProductListRequestFilterFilterVisibility = "VISIBLE"
	V5GetProductListRequestFilterFilterVisibilityInvisible              V5GetProductListRequestFilterFilterVisibility = "INVISIBLE"
	V5GetProductListRequestFilterFilterVisibilityEmptyStock             V5GetProductListRequestFilterFilterVisibility = "EMPTY_STOCK"
	V5GetProductListRequestFilterFilterVisibilityNOTModerated           V5GetProductListRequestFilterFilterVisibility = "NOT_MODERATED"
	V5GetProductListRequestFilterFilterVisibilityModerated              V5GetProductListRequestFilterFilterVisibility = "MODERATED"
	V5GetProductListRequestFilterFilterVisibilityDisabled               V5GetProductListRequestFilterFilterVisibility = "DISABLED"
	V5GetProductListRequestFilterFilterVisibilityStateFailed            V5GetProductListRequestFilterFilterVisibility = "STATE_FAILED"
	V5GetProductListRequestFilterFilterVisibilityReadyTOSupply          V5GetProductListRequestFilterFilterVisibility = "READY_TO_SUPPLY"
	V5GetProductListRequestFilterFilterVisibilityValidationStatePending V5GetProductListRequestFilterFilterVisibility = "VALIDATION_STATE_PENDING"
	V5GetProductListRequestFilterFilterVisibilityValidationStateFail    V5GetProductListRequestFilterFilterVisibility = "VALIDATION_STATE_FAIL"
	V5GetProductListRequestFilterFilterVisibilityValidationStateSuccess V5GetProductListRequestFilterFilterVisibility = "VALIDATION_STATE_SUCCESS"
	V5GetProductListRequestFilterFilterVisibilityTOSupply               V5GetProductListRequestFilterFilterVisibility = "TO_SUPPLY"
	V5GetProductListRequestFilterFilterVisibilityINSale                 V5GetProductListRequestFilterFilterVisibility = "IN_SALE"
	V5GetProductListRequestFilterFilterVisibilityRemovedFromSale        V5GetProductListRequestFilterFilterVisibility = "REMOVED_FROM_SALE"
	V5GetProductListRequestFilterFilterVisibilityOverpriced             V5GetProductListRequestFilterFilterVisibility = "OVERPRICED"
	V5GetProductListRequestFilterFilterVisibilityCriticallyOverpriced   V5GetProductListRequestFilterFilterVisibility = "CRITICALLY_OVERPRICED"
	V5GetProductListRequestFilterFilterVisibilityEmptyBarcode           V5GetProductListRequestFilterFilterVisibility = "EMPTY_BARCODE"
	V5GetProductListRequestFilterFilterVisibilityBarcodeExists          V5GetProductListRequestFilterFilterVisibility = "BARCODE_EXISTS"
	V5GetProductListRequestFilterFilterVisibilityQuarantine             V5GetProductListRequestFilterFilterVisibility = "QUARANTINE"
	V5GetProductListRequestFilterFilterVisibilityArchived               V5GetProductListRequestFilterFilterVisibility = "ARCHIVED"
	V5GetProductListRequestFilterFilterVisibilityOverpricedWithStock    V5GetProductListRequestFilterFilterVisibility = "OVERPRICED_WITH_STOCK"
	V5GetProductListRequestFilterFilterVisibilityPartialApproved        V5GetProductListRequestFilterFilterVisibility = "PARTIAL_APPROVED"
	V5GetProductListRequestFilterFilterVisibilityAutoArchived           V5GetProductListRequestFilterFilterVisibility = "AUTO_ARCHIVED"
	V5GetProductListRequestFilterFilterVisibilityManualArchived         V5GetProductListRequestFilterFilterVisibility = "MANUAL_ARCHIVED"
	V5GetProductListRequestFilterFilterVisibilitySeasonalAutoArchived   V5GetProductListRequestFilterFilterVisibility = "SEASONAL_AUTO_ARCHIVED"
	V5GetProductListRequestFilterFilterVisibilityVisibleWithFBOStock    V5GetProductListRequestFilterFilterVisibility = "VISIBLE_WITH_FBO_STOCK"
)

// 商品筛选器。
type V5Filter struct {
	OfferID    []string                                      `json:"offer_id"`   // 基于 `offer_id` 参数的筛选（最多可传递 1000 个值）。
	ProductID  []string                                      `json:"product_id"` // 基于 `product_id` 参数的筛选（最多可传递 1000 个值）。
	Visibility V5GetProductListRequestFilterFilterVisibility `json:"visibility"`
}

type V5GetProductInfoPricesV5Request struct {
	Cursor string   `json:"cursor"` // 用于选择数据的指针。
	Filter V5Filter `json:"filter"`
	Limit  int32    `json:"limit"` // 每页显示的数值数量。
}

type V2ProductsStocksResponseError struct {
	Code    string `json:"code"`    // 错误代码。
	Message string `json:"message"` // 错误原因。
}

type V2ProductsStocksResponseResult struct {
	ProductID   int64                           `json:"product_id"`   // Ozon系统中商品的标识符 — `product_id`。
	Updated     bool                            `json:"updated"`      // 如果商品信息已被成功更新 — `true`。
	WarehouseID int64                           `json:"warehouse_id"` // 仓库编号。
	Errors      []V2ProductsStocksResponseError `json:"errors"`       // 在搜索处理过程中发生的数组错误。
	OfferID     string                          `json:"offer_id"`     // 卖家系统中的商品编号是 — 商品代码。
}

// 搜索结果。
type V2ProductsStocksResponse struct {
	Result []V2ProductsStocksResponseResult `json:"result"`
}

type ImportProductsPricesResponseError struct {
	Code    string `json:"code"`    // 错误代码。
	Message string `json:"message"` // 错误原因。
}

type ImportProductsPricesResponseProcessResult struct {
	Errors    []ImportProductsPricesResponseError `json:"errors"`     // 在搜索处理过程中发生的数组错误。
	OfferID   string                              `json:"offer_id"`   // 卖家系统中的商品编号是 — 商品代码。
	ProductID int64                               `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	Updated   bool                                `json:"updated"`    // 如果商品信息已被成功更新 — `true`。
}

type V1GetProductInfoDiscountedRequest struct {
	DiscountedSkus []string `json:"discounted_skus"` // 降价的商品SKU清单。
}

// 佣金信息。
type ItemCommissionsv5 struct {
	FBSDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"` // 干线运输至仓 (FBS)。
	FBSFirstMileMaxAmount       float64 `json:"fbs_first_mile_max_amount"`        // 货件处理最高佣金 (FBS)。
	FBSReturnFlowAmount         float64 `json:"fbs_return_flow_amount"`           // 退货及取消货件佣金，货件处理费用 (FBS)。
	SalesPercentFBO             float64 `json:"sales_percent_fbo"`                // 销售佣金比例 (FBO)。
	SalesPercentFBS             float64 `json:"sales_percent_fbs"`                // 销售佣金比例 (FBS)。
	FBODirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"` // 干线运输至仓 (FBO)。
	FBOReturnFlowAmount         float64 `json:"fbo_return_flow_amount"`           // 退货及取消订单手续费 (FBO)。
	FBSDelivToCustomerAmount    float64 `json:"fbs_deliv_to_customer_amount"`     // 尾程物流 (FBS)。
	FBSDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"` // 干线运输出仓 (FBS)。
	FBSFirstMileMinAmount       float64 `json:"fbs_first_mile_min_amount"`        // 货件处理最低佣金 (FBS)。
	SalesPercentFBP             float64 `json:"sales_percent_fbp"`                // 销售佣金比例 (FBP)。
	SalesPercentRFBS            float64 `json:"sales_percent_rfbs"`               // 销售佣金比例 (rFBS)。
	FBODelivToCustomerAmount    float64 `json:"fbo_deliv_to_customer_amount"`     // 尾程物流 (FBO)。
	FBODirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"` // 干线运输出仓 (FBO)。
}

type MarketingAction struct {
	DateFrom string `json:"date_from"` // 促销活动开始日期与时间。
	DateTo   string `json:"date_to"`   // 促销活动结束日期与时间。
	Title    string `json:"title"`     // 促销活动名称。
	Value    int32  `json:"value"`     // 促销折扣。
}

// 营销活动。
type ItemMarketing struct {
	Actions []MarketingAction `json:"actions"` // 营销活动。 `date_from`、`date_to`、`title`和`value`参数会针对每个促销活动分别指定。
}

// 商品价格。
type ItemPricev5 struct {
	RetailPrice          float64      `json:"retail_price"`           // 供应商按合同约定的价格。如果没有供货合同，该字段将返回为空。
	MarketingSellerPrice float64      `json:"marketing_seller_price"` // 包含卖家促销的商品价格。
	Price                float64      `json:"price"`                  // 包含折扣的商品价格（在商品卡片展示的实际售价）。
	Vat                  float64      `json:"vat"`                    // 商品的增值税税率。
	AutoActionEnabled    bool         `json:"auto_action_enabled"`    // 如果商品启用了促销活动的自动应用，则为 `true`。
	CurrencyCode         CurrencyCode `json:"currency_code"`          // 您的定价货币（与个人中心设置中指定的货币相同）。 可能的值： - `RUB` — 俄罗斯卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 人民...
	MinPrice             float64      `json:"min_price"`              // 商品在应用所有折扣后的最低价格。
	NetPrice             float64      `json:"net_price"`              // 商品成本。
	OldPrice             float64      `json:"old_price"`              // 未应用折扣前的价格（在商品卡片显示为划线价）。
}

// Ozon 上的竞争对手价格。
type PriceIndexesIndexOzonData struct {
	MinPrice         float64 `json:"min_price"`          // Ozon 上竞争对手的最低商品价格。
	MinPriceCurrency string  `json:"min_price_currency"` // 价格的货币单位。
	PriceIndexValue  float64 `json:"price_index_value"`  // 价格指数值。
}

// 您的商品在其他平台上的价格。
type PriceIndexesIndexSelfData struct {
	MinPriceCurrency string  `json:"min_price_currency"` // 价格的货币单位。
	PriceIndexValue  float64 `json:"price_index_value"`  // 价格指数值。
	MinPrice         float64 `json:"min_price"`          // 您的商品在其他平台上的最低价格。
}

// 商品价格指数。
// ColorIndex values
type ColorIndex string

const (
	ColorIndexWithoutIndex ColorIndex = "WITHOUT_INDEX" // 无价格指数
	ColorIndexSuper        ColorIndex = "SUPER"         // 超级有利
	ColorIndexGreen        ColorIndex = "GREEN"         // 有利
	ColorIndexYellow       ColorIndex = "YELLOW"        // 中等
	ColorIndexRED          ColorIndex = "RED"           // 不利
)

type GetProductInfoPricesResponseItemPriceIndexes struct {
	ColorIndex                ColorIndex                    `json:"color_index"` // 最终价格指数： - `WITHOUT_INDEX` — 无价格指数， - `SUPER` — 超级有利, - `GREEN` — 有利， - `YELLOW` — 中等， - `RED` — 不利。
	ExternalIndexData         PriceIndexesIndexExternalData `json:"external_index_data"`
	OzonIndexData             PriceIndexesIndexOzonData     `json:"ozon_index_data"`
	SelfMarketplacesIndexData PriceIndexesIndexSelfData     `json:"self_marketplaces_index_data"`
}

type GetProductInfoPricesV5ResponseItem struct {
	VolumeWeight     float64                                      `json:"volume_weight"` // 商品体积重量。
	Acquiring        float64                                      `json:"acquiring"`     // 最高收单手续费。
	Commissions      ItemCommissionsv5                            `json:"commissions"`
	MarketingActions ItemMarketing                                `json:"marketing_actions"` // 卖家营销活动。
	OfferID          string                                       `json:"offer_id"`          // 商品在卖家系统中的标识符 — 货号。
	Price            ItemPricev5                                  `json:"price"`
	PriceIndexes     GetProductInfoPricesResponseItemPriceIndexes `json:"price_indexes"`
	ProductID        int64                                        `json:"product_id"` // 商品在卖家系统中的标识符 — `product_id`。
}

type ImportProductsPricesRequest struct {
	Prices []ImportProductsPricesRequestPrice `json:"prices"` // 商品价格信息。
}

type V5GetProductInfoPricesV5Response struct {
	Cursor string                               `json:"cursor"` // 用于选择数据的指针。
	Items  []GetProductInfoPricesV5ResponseItem `json:"items"`  // 商品列表。
	Total  int32                                `json:"total"`  // 商品列表中的商品数量。
}

// 按商品曝光度过滤: - `ALL` —— 所有的商品，除了已存档的。 - `VISIBLE` —— 买家可见的商品。 - `INVISIBLE` —— 买家看不到的商品。 - `EMPTY_STOCK`—— 没有现货的商品。 - `NOT_...
type V4Visibility string

const (
	V4VisibilityALL                    V4Visibility = "ALL"
	V4VisibilityVisible                V4Visibility = "VISIBLE"
	V4VisibilityInvisible              V4Visibility = "INVISIBLE"
	V4VisibilityEmptyStock             V4Visibility = "EMPTY_STOCK"
	V4VisibilityNOTModerated           V4Visibility = "NOT_MODERATED"
	V4VisibilityModerated              V4Visibility = "MODERATED"
	V4VisibilityDisabled               V4Visibility = "DISABLED"
	V4VisibilityStateFailed            V4Visibility = "STATE_FAILED"
	V4VisibilityReadyTOSupply          V4Visibility = "READY_TO_SUPPLY"
	V4VisibilityValidationStatePending V4Visibility = "VALIDATION_STATE_PENDING"
	V4VisibilityValidationStateFail    V4Visibility = "VALIDATION_STATE_FAIL"
	V4VisibilityValidationStateSuccess V4Visibility = "VALIDATION_STATE_SUCCESS"
	V4VisibilityTOSupply               V4Visibility = "TO_SUPPLY"
	V4VisibilityINSale                 V4Visibility = "IN_SALE"
	V4VisibilityRemovedFromSale        V4Visibility = "REMOVED_FROM_SALE"
	V4VisibilityBanned                 V4Visibility = "BANNED"
	V4VisibilityOverpriced             V4Visibility = "OVERPRICED"
	V4VisibilityCriticallyOverpriced   V4Visibility = "CRITICALLY_OVERPRICED"
	V4VisibilityEmptyBarcode           V4Visibility = "EMPTY_BARCODE"
	V4VisibilityBarcodeExists          V4Visibility = "BARCODE_EXISTS"
	V4VisibilityQuarantine             V4Visibility = "QUARANTINE"
	V4VisibilityArchived               V4Visibility = "ARCHIVED"
	V4VisibilityOverpricedWithStock    V4Visibility = "OVERPRICED_WITH_STOCK"
	V4VisibilityPartialApproved        V4Visibility = "PARTIAL_APPROVED"
	V4VisibilityAutoArchived           V4Visibility = "AUTO_ARCHIVED"
	V4VisibilityManualArchived         V4Visibility = "MANUAL_ARCHIVED"
	V4VisibilitySeasonalAutoArchived   V4Visibility = "SEASONAL_AUTO_ARCHIVED"
	V4VisibilityVisibleWithFBOStock    V4Visibility = "VISIBLE_WITH_FBO_STOCK"
)

// “经济”费率商品。
type FilterWithQuant struct {
	Created bool `json:"created"` // 处于活跃状态的经济商品。
	Exists  bool `json:"exists"`  // 所有状态下的经济商品。
}

// 商品筛选器。
type V4GetProductInfoStocksRequestFilter struct {
	OfferID    []string        `json:"offer_id"`   // 基于参数 `offer_id` 的过滤。 可以提交数值列表。
	ProductID  []string        `json:"product_id"` // 基于参数 `product_id` 的过滤。 可以提交数值列表。
	Visibility V4Visibility    `json:"visibility"`
	WithQuant  FilterWithQuant `json:"with_quant"`
}

type V4GetProductInfoStocksRequest struct {
	Cursor string                              `json:"cursor"` // 后续数据的选择标志。
	Filter V4GetProductInfoStocksRequestFilter `json:"filter"`
	Limit  int32                               `json:"limit"` // 页面上的值数量。最低为1，最高为1000。
}

type V1ProductActionTimerStatusResponse struct {
	Statuses []V1ProductActionTimerStatusResponseStatuses `json:"statuses"`
}

type V1ProductActionTimerStatusRequest struct {
	ProductIds []string `json:"product_ids"` // 卖家系统中的商品识别符列表——`product_id`。
}

type V2GetProductInfoStocksByWarehouseFbsRequestV2 struct {
	Cursor  string   `json:"cursor"`   // 后续数据的选择标志。
	Limit   int64    `json:"limit"`    // 响应中返回的值数量。
	OfferID []string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	SKU     []string `json:"sku"`      // Ozon系统中的商品标识符——SKU。
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2Product struct {
	OfferID       string `json:"offer_id"`       // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	Present       int64  `json:"present"`        // 仓库中的商品总数量。
	ProductID     int64  `json:"product_id"`     // 卖家系统中的Ozon系统中商品的标识符 — `product_id`。
	Reserved      int64  `json:"reserved"`       // 仓库中已预留商品的数量。
	SKU           int64  `json:"sku"`            // Ozon系统中的商品识别码是SKU。
	WarehouseID   int64  `json:"warehouse_id"`   // 仓库标识符。
	WarehouseName string `json:"warehouse_name"` // 仓库名称。
	FreeStock     int64  `json:"free_stock"`     // 可售商品数量。
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2 struct {
	Cursor   string                                                  `json:"cursor"`   // 用于获取下一批数据的指针。
	HasNext  bool                                                    `json:"has_next"` // 如果响应中未返回所有商品，则为`true`。
	Products []V2GetProductInfoStocksByWarehouseFbsResponseV2Product `json:"products"` // 商品库存。
}

type Sv1GetProductInfoStocksByWarehouseFbsRequest struct {
	SKU     []string `json:"sku"`      // Ozon系统中的商品识别码是SKU。
	OfferID []string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
}

type ImportProductsPricesResponse struct {
	Result []ImportProductsPricesResponseProcessResult `json:"result"` // 搜索结果。
}

type V1ProductUpdateDiscountResponse struct {
	Result bool `json:"result"` // 方式工作结果 `true`, 如果正确完成请求。
}

type V1ProductActionTimerUpdateRequest struct {
	ProductIds []string `json:"product_ids"` // 卖家系统中的商品识别符列表——`product_id`。
}

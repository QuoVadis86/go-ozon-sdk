package report

// 细节。
type FinanceCashFlowStatementListResponseDetailsOthers struct {
	Name string `json:"name"` // 操作名称。可能的值： - `MarketplaceRedistributionOfAcquiringOperation` — 收单业务支付， - `MarketplaceSellerCompensationLossOfGoodsOperat...
	Price float64 `json:"price"` // 操作金额。
}

// 补贴和其他津贴。
type DetailsOthers struct {
	Total float64 `json:"total"` // 总数。
	Items []FinanceCashFlowStatementListResponseDetailsOthers `json:"items"` // 细节。
}

// 关于报告的信息。
type ReportReport struct {
	Error string `json:"error"` // 生成报告时的错误代码。
	ExpiresAt string `json:"expires_at"` // 报告链接的有效日期和时间。 如果报告生成于 2025 年 10 月 14 日之前，该字段将为空。
	File string `json:"file"` // XLSX文件的链接。 `SELLER_RETURNS` 类型的报告，链接有效期为5分钟。
	Params map[string]any `json:"params"` // 一个数组，包含卖家创建报告时指定的过滤器。
	ReportType string `json:"report_type"` // 报告类型： - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTINGS` — 发货报告， - `S...
	Status string `json:"status"` // 报告生成状态： - `waiting`—在等待队列中待处理， - `processing`—正在处理， - `success`—报告成功生成， - `failed` — 报告生成错误。
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
	CreatedAt string `json:"created_at"` // 报告创建日期。
}

// 请求结果。
type ReportListResponseResult struct {
	Reports []ReportReport `json:"reports"` // 包含所有生成的报告的数组。
	Total int32 `json:"total"` // 累计报告数。
}

// 阶段。
type V3FinanceCashFlowStatementListResponsePeriod struct {
	ID int64 `json:"id"` // 时期识别码。
	Begin string `json:"begin"` // 开始阶段。
	End string `json:"end"` // 结束阶段。
}

type FinanceCashFlowStatementListResponseCashFlow struct {
	ReturnsAmount float64 `json:"returns_amount"` // 退货价格总和。
	CommissionAmount float64 `json:"commission_amount"` // 商品销售Ozon佣金。
	ServicesAmount float64 `json:"services_amount"` // 附加服务数额。
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"` // 物流服务数额。
	CurrencyCode string `json:"currency_code"` // 佣金计算的货币代码。
	Period V3FinanceCashFlowStatementListResponsePeriod `json:"period"`
	OrdersAmount float64 `json:"orders_amount"` // 已成交商品的价格总和。
}

// rFBS模式转账金额。
type DetailsRfbsDetails struct {
	Total float64 `json:"total"` // 总额。
	TransferDelivery float64 `json:"transfer_delivery"` // 来自买家的转账金额。
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"` // 退还给买家的转账金额。
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"` // 物流转账金额补贴。
	PartialCompensation float64 `json:"partial_compensation"` // 将部分退款转移给买家。
	PartialCompensationReturn float64 `json:"partial_compensation_return"` // 退换部分退款。
}

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

// 额外的字段，需要添加到响应中。
type CreateCompanyPostingsReportRequestWith struct {
	JewelryCodes bool `json:"jewelry_codes"` // `true`，用于在响应中添加珠宝信息。
	AdditionalData bool `json:"additional_data"` // `true`，用于在响应中添加附加信息。
	AnalyticsData bool `json:"analytics_data"` // `true`，用于在响应中添加分析数据。请传递值 `filter.delivery_schema = fbs`，否则 将返回 错误。
	CustomerData bool `json:"customer_data"` // `true`，用于在响应中添加买家信息。
}

// 报告类型： - `ALL`— 所有报告， - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTING...
type ReportListRequestReportType string

// 细节。
type FinanceCashFlowStatementListResponseDeliveryService struct {
	Name string `json:"name"` // 操作名称。可能的值： - `MarketplaceServiceItemDirectFlowLogisticSum` — 物流， - `MarketplaceServiceItemDropoff` — Drop-off发货处理， - `Ma...
	Price float64 `json:"price"` // 操作总和。
}

// 加工费和运费报酬。
type DetailsServices struct {
	Total float64 `json:"total"` // 总额。
	Items []FinanceCashFlowStatementListResponseDeliveryService `json:"items"` // 细节。
}

// 请求结果。
type CreateReportResponseCode struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

// 关于报告的信息。
type ReportReportinfo struct {
	Params map[string]any `json:"params"` // 一个数组，包含卖家创建报告时指定的过滤器。
	ReportType string `json:"report_type"` // 报告类型： - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTINGS` — 发货报告， - `S...
	Status string `json:"status"` // 报告生成状态： - `waiting`—在等待队列中待处理， - `processing`—正在处理， - `success`， - `failed`。
	Code string `json:"code"` // 报告的唯一识别码。
	CreatedAt string `json:"created_at"` // 报告创建日期。
	Error string `json:"error"` // 生成报告时的错误代码。
	ExpiresAt string `json:"expires_at"` // 报告链接的有效日期和时间。 如果报告生成于 2025 年 10 月 14 日之前，该字段将为空。
	File string `json:"file"` // XLSX文件的链接。 `SELLER_RETURNS` 类型的报告，链接有效期为5分钟。
}

// 回答所用语言： - `RU` — 俄语， - `EN` — 英语。
type ReportLanguage string

type V1CreateStockByWarehouseReportRequest struct {
	WarehouseId []string `json:"warehouseId"` // 仓库ID。 请求中参数值的限制。 最大值为 50。
	Language ReportLanguage `json:"language"`
}

// 已按时期支付。
type DetailsPayment struct {
	CurrencyCode string `json:"currency_code"` // 货币。
	Payment float64 `json:"payment"` // 支付金额。
}

// 细节。
type FinanceCashFlowStatementListResponseReturnService struct {
	Name string `json:"name"` // 操作名称。可能的值： - `MarketplaceServiceItemReturnAfterDelivToCustomer` — 退货处理， - `MarketplaceServiceItemReturnPartGoodsCustomer...
	Price float64 `json:"price"` // 操作数量。
}

// 退货和取消费用。
type DetailsReturns struct {
	Items []FinanceCashFlowStatementListResponseReturnService `json:"items"` // 细节。
	Total float64 `json:"total"` // 总量。
}

// 退货和取消订单。
type DetailsReturnDetails struct {
	Total float64 `json:"total"` // 总金额。
	Amount float64 `json:"amount"` // 佣金考虑在内的退款金额。
	ReturnServices DetailsReturns `json:"return_services"` // 退款和取消的费用。
}

// 细节。
type FinanceCashFlowStatementListResponseService struct {
	Name string `json:"name"` // 操作名称。可能的值： - `MarketplaceServiceItemElectronicServiceStencil` — “模板”服务， - `MarketplaceServiceItemElectronicServicesPromo...
	Price float64 `json:"price"` // 操作数量。
}

// 服务。
type DetailsService struct {
	Total float64 `json:"total"` // 总额。
	Items []FinanceCashFlowStatementListResponseService `json:"items"` // 细节。
}

// 订单。
type DetailsDeliveryDetails struct {
	DeliveryServices DetailsServices `json:"delivery_services"` // 加工费和运费。
	Total float64 `json:"total"` // 总额。
	Amount float64 `json:"amount"` // 考虑佣金商品购买的价格。
}

// 详细信息。
type FinanceCashFlowStatementListResponseDetails struct {
	BeginBalanceAmount float64 `json:"begin_balance_amount"` // 开始阶段的收支。
	InvoiceTransfer float64 `json:"invoice_transfer"` // 当期应付金额。
	Payments DetailsPayment `json:"payments"` // 期间已付清。
	Return DetailsReturnDetails `json:"return"`
	Services DetailsService `json:"services"` // 服务。
	Others DetailsOthers `json:"others"` // 补偿费和其他费用。
	EndBalanceAmount float64 `json:"end_balance_amount"` // 结束阶段的收支。
	Delivery DetailsDeliveryDetails `json:"delivery"` // 方法操作结果。
	Loan float64 `json:"loan"` // 根据贷款协议转账。
	Period V3FinanceCashFlowStatementListResponsePeriod `json:"period"` // 时期。
	RFBS DetailsRfbsDetails `json:"rfbs"` // rFBS框架下的转账金额。
}

// 方法操作结果。
type V3FinanceCashFlowStatementListResponseResult struct {
	PageCount int64 `json:"page_count"` // 含有报告的页数。
	CashFlows []FinanceCashFlowStatementListResponseCashFlow `json:"cash_flows"` // 报告清单。
	Details []FinanceCashFlowStatementListResponseDetails `json:"details"` // 细节信息。
}

type V3FinanceCashFlowStatementListResponse struct {
	Result V3FinanceCashFlowStatementListResponseResult `json:"result"`
}

// 按商品可见度过滤。 - `ALL`——除了档案中的所有商品； - `VALIDATION_STATE_FAIL`——预审时未被验证器检查的商品； - `TO_SUPPLY`——准备出售的货物； - `IN_SALE`——正在销售的商品； -...
type ReportCreateCompanyProductsReportRequestVisibility string

// 报告期限。
type Financev3Period struct {
	From string `json:"from"` // 计算报告的起始日期。
	To string `json:"to"` // 计算报告的停止日期。
}

type V3FinanceCashFlowStatementListRequest struct {
	WithDetails bool `json:"with_details"` // `true`，如果需要在响应中添加附加参数。
	PageSize int32 `json:"page_size"` // 页面上的元素数量。
	Date Financev3Period `json:"date"`
	Page int32 `json:"page"` // 请求返回中的页码。
}

type ReportReportInfoRequest struct {
	Code string `json:"code"` // 报告的唯一识别码。
}

// 报告生成周期。
type ReportMarkedProductsSalesCreateRequestDate struct {
	From string `json:"from"` // 报告期开始日期，格式为 `YYYY-MM-DD`。
	To string `json:"to"` // 报告期结束日期，格式为 `YYYY-MM-DD`。
}

type ReportCreateCompanyProductsReportRequest struct {
	Search string `json:"search"` // 在记录内容中搜索，检查现货。
	SKU []int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	Visibility ReportCreateCompanyProductsReportRequestVisibility `json:"visibility"`
	Language ReportLanguage `json:"language"`
	OfferID []string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
}

type ReportCreateDiscountedRequest any

type ReportReportListRequest struct {
	Page int32 `json:"page"` // 页数。
	PageSize int32 `json:"page_size"` // 每页的值的数量： - 默认值 — 100， - 最大值 — 1,000。
	ReportType ReportListRequestReportType `json:"report_type"`
}

type ReportCreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

// 过滤器。
type ReportCreateCompanyPostingsReportRequestFilter struct {
	ProcessedAtTo string `json:"processed_at_to"` // 订单出现在个人账户的时间。
	Title string `json:"title"` // 商品名称。
	DeliveryMethodID []int64 `json:"delivery_method_id"` // 配送方法标识符。 可通过方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
	SKU []int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	StatusAlias []string `json:"status_alias"` // 状态文本。
	Statuses []int64 `json:"statuses"` // 数值状况。
	WarehouseID []int64 `json:"warehouse_id"` // 仓库标识符。
	IsExpress any `json:"is_express"` // 快递配送： - `true`—仅包含使用 Ozon Express 快速配送的货件； - `false`—仅包含未使用 Ozon Express 快速配送的货件。 如果未传递任何值，将返回所有货件记录。
	CancelReasonID []int64 `json:"cancel_reason_id"` // 取消原因的识别码。
	DeliverySchema []string `json:"delivery_schema"` // 运作机制是FBO或FBS。 对于海外卖家来说，只有FBS方案可用，所以在参数中提交数值`fbs`。
	ProcessedAtFrom string `json:"processed_at_from"` // 订单进入处理程序的时间。
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date ReportMarkedProductsSalesCreateRequestDate `json:"date"`
}

type ReportReportInfoResponse struct {
	Result ReportReportinfo `json:"result"`
}

type ReportCreateCompanyPostingsReportRequest struct {
	Filter ReportCreateCompanyPostingsReportRequestFilter `json:"filter"`
	Language ReportLanguage `json:"language"`
	With CreateCompanyPostingsReportRequestWith `json:"with"`
}

type ReportReportListResponse struct {
	Result ReportListResponseResult `json:"result"`
}

type CommonCreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

package report

// 细节。
// 操作名称
type FinanceCashFlowStatementListResponseReturnServiceNameEnum string
const (
	FinanceCashFlowStatementListResponseReturnServiceNameEnum_MarketplaceServiceItemReturnAfterDelivToCustomer FinanceCashFlowStatementListResponseReturnServiceNameEnum = "MarketplaceServiceItemReturnAfterDelivToCustomer"
	FinanceCashFlowStatementListResponseReturnServiceNameEnum_MarketplaceServiceItemReturnPartGoodsCustomer FinanceCashFlowStatementListResponseReturnServiceNameEnum = "MarketplaceServiceItemReturnPartGoodsCustomer"
	FinanceCashFlowStatementListResponseReturnServiceNameEnum_MarketplaceServiceItemReturnNotDelivToCustomer FinanceCashFlowStatementListResponseReturnServiceNameEnum = "MarketplaceServiceItemReturnNotDelivToCustomer"
	FinanceCashFlowStatementListResponseReturnServiceNameEnum_MarketplaceServiceItemReturnFlowLogistic FinanceCashFlowStatementListResponseReturnServiceNameEnum = "MarketplaceServiceItemReturnFlowLogistic"
)

type FinanceCashFlowStatementListResponseReturnService struct {
	Name FinanceCashFlowStatementListResponseReturnServiceNameEnum `json:"name"` // 操作名称。可能的值： - `MarketplaceServiceItemReturnAfterDelivToCustomer` — 退货处理， - `MarketplaceServiceItemReturnPartGoodsCustomer...
	Price float64 `json:"price"` // 操作数量。
}

// 退货和取消费用。
type DetailsReturns struct {
	Total float64 `json:"total"` // 总量。
	Items []FinanceCashFlowStatementListResponseReturnService `json:"items"` // 细节。
}

// 退货和取消订单。
type DetailsReturnDetails struct {
	Total float64 `json:"total"` // 总金额。
	Amount float64 `json:"amount"` // 佣金考虑在内的退款金额。
	ReturnServices DetailsReturns `json:"return_services"` // 退款和取消的费用。
}

// 细节。
// 操作名称
type FinanceCashFlowStatementListResponseServiceNameEnum string
const (
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemElectronicServiceStencil FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemElectronicServiceStencil"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemElectronicServicesPromotionInSearch FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemElectronicServicesPromotionInSearch"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemElectronicServicesBrandShelf FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemElectronicServicesBrandShelf"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceBrandPromotion FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceBrandPromotion"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceBrandCommission FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceBrandCommission"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemMarketingServices FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemMarketingServices"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemOtherMarketAndTechService FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemOtherMarketAndTechService"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceReturnStorageServiceAtThePickupPointFbsItem FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceReturnStorageServiceAtThePickupPointFbsItem"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceSaleReviewsItem FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceSaleReviewsItem"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServicePremiumCashbackIndividualPoints FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServicePremiumCashbackIndividualPoints"
	FinanceCashFlowStatementListResponseServiceNameEnum_OperationMarketplaceServiceStorage FinanceCashFlowStatementListResponseServiceNameEnum = "OperationMarketplaceServiceStorage"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceStockDisposal FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceStockDisposal"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceReturnDisposalServiceFbsItem FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceReturnDisposalServiceFbsItem"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemFlexiblePaymentSchedule FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemFlexiblePaymentSchedule"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceProcessingSpoilage FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceProcessingSpoilage"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceProcessingIdentifiedSurplus FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceProcessingIdentifiedSurplus"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceProcessingIdentifiedDiscrepancies FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceProcessingIdentifiedDiscrepancies"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemInternetSiteAdvertising FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemInternetSiteAdvertising"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceServiceItemSubscribtionPremium FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceServiceItemSubscribtionPremium"
	FinanceCashFlowStatementListResponseServiceNameEnum_MarketplaceAgencyFeeAggregator3PLGlobalItem FinanceCashFlowStatementListResponseServiceNameEnum = "MarketplaceAgencyFeeAggregator3PLGlobalItem"
)

type FinanceCashFlowStatementListResponseService struct {
	Name FinanceCashFlowStatementListResponseServiceNameEnum `json:"name"` // 操作名称。可能的值： - `MarketplaceServiceItemElectronicServiceStencil` — “模板”服务， - `MarketplaceServiceItemElectronicServicesPromo...
	Price float64 `json:"price"` // 操作数量。
}

// 请求结果。
type CreateReportResponseCode struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

type ReportReportInfoRequest struct {
	Code string `json:"code"` // 报告的唯一识别码。
}

// rFBS模式转账金额。
type DetailsRfbsDetails struct {
	PartialCompensationReturn float64 `json:"partial_compensation_return"` // 退换部分退款。
	Total float64 `json:"total"` // 总额。
	TransferDelivery float64 `json:"transfer_delivery"` // 来自买家的转账金额。
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"` // 退还给买家的转账金额。
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"` // 物流转账金额补贴。
	PartialCompensation float64 `json:"partial_compensation"` // 将部分退款转移给买家。
}

// 额外的字段，需要添加到响应中。
type CreateCompanyPostingsReportRequestWith struct {
	CustomerData bool `json:"customer_data"` // `true`，用于在响应中添加买家信息。
	JewelryCodes bool `json:"jewelry_codes"` // `true`，用于在响应中添加珠宝信息。
	AdditionalData bool `json:"additional_data"` // `true`，用于在响应中添加附加信息。
	AnalyticsData bool `json:"analytics_data"` // `true`，用于在响应中添加分析数据。请传递值 `filter.delivery_schema = fbs`，否则 将返回 错误。
}

// 回答所用语言： - `RU` — 俄语， - `EN` — 英语。
type ReportLanguage string

// 按商品可见度过滤。 - `ALL`——除了档案中的所有商品； - `VALIDATION_STATE_FAIL`——预审时未被验证器检查的商品； - `TO_SUPPLY`——准备出售的货物； - `IN_SALE`——正在销售的商品； -...
type ReportCreateCompanyProductsReportRequestVisibility string

type ReportCreateCompanyProductsReportRequest struct {
	Search string `json:"search"` // 在记录内容中搜索，检查现货。
	SKU []int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	Visibility ReportCreateCompanyProductsReportRequestVisibility `json:"visibility"`
	Language ReportLanguage `json:"language"`
	OfferID []string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
}

// 关于报告的信息。
// 报告类型：
type ReportReportReportTypeEnum string
const (
	ReportReportReportTypeEnum_SELLERPRODUCTS ReportReportReportTypeEnum = "SELLER_PRODUCTS"
	ReportReportReportTypeEnum_SELLERSTOCK ReportReportReportTypeEnum = "SELLER_STOCK"
	ReportReportReportTypeEnum_SELLERRETURNS ReportReportReportTypeEnum = "SELLER_RETURNS"
	ReportReportReportTypeEnum_SELLERPOSTINGS ReportReportReportTypeEnum = "SELLER_POSTINGS"
	ReportReportReportTypeEnum_SELLERDISCOUNTED ReportReportReportTypeEnum = "SELLER_DISCOUNTED"
	ReportReportReportTypeEnum_MUTUALSETTLEMENT ReportReportReportTypeEnum = "MUTUAL_SETTLEMENT"
	ReportReportReportTypeEnum_DOCUMENTB2BSALES ReportReportReportTypeEnum = "DOCUMENT_B2B_SALES"
	ReportReportReportTypeEnum_COMPENSATIONREPORT ReportReportReportTypeEnum = "COMPENSATION_REPORT"
	ReportReportReportTypeEnum_DECOMPENSATIONREPORT ReportReportReportTypeEnum = "DECOMPENSATION_REPORT"
	ReportReportReportTypeEnum_MARKEDPRODUCTSSALES ReportReportReportTypeEnum = "MARKED_PRODUCTS_SALES"
	ReportReportReportTypeEnum_SELLERPLACEMENTBYPRODUCTS ReportReportReportTypeEnum = "SELLER_PLACEMENT_BY_PRODUCTS"
	ReportReportReportTypeEnum_SELLERPLACEMENTBYSUPPLIES ReportReportReportTypeEnum = "SELLER_PLACEMENT_BY_SUPPLIES"
)

// 报告生成状态：
type ReportReportStatusEnum string
const (
	ReportReportStatusEnum_Waiting ReportReportStatusEnum = "waiting"
	ReportReportStatusEnum_Processing ReportReportStatusEnum = "processing"
	ReportReportStatusEnum_Success ReportReportStatusEnum = "success"
	ReportReportStatusEnum_Failed ReportReportStatusEnum = "failed"
)

type ReportReport struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
	CreatedAt string `json:"created_at"` // 报告创建日期。
	Error string `json:"error"` // 生成报告时的错误代码。
	ExpiresAt string `json:"expires_at"` // 报告链接的有效日期和时间。 如果报告生成于 2025 年 10 月 14 日之前，该字段将为空。
	File string `json:"file"` // XLSX文件的链接。 `SELLER_RETURNS` 类型的报告，链接有效期为5分钟。
	Params map[string]any `json:"params"` // 一个数组，包含卖家创建报告时指定的过滤器。
	ReportType ReportReportReportTypeEnum `json:"report_type"` // 报告类型： - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTINGS` — 发货报告， - `S...
	Status ReportReportStatusEnum `json:"status"` // 报告生成状态： - `waiting`—在等待队列中待处理， - `processing`—正在处理， - `success`—报告成功生成， - `failed` — 报告生成错误。
}

// 请求结果。
type ReportListResponseResult struct {
	Reports []ReportReport `json:"reports"` // 包含所有生成的报告的数组。
	Total int32 `json:"total"` // 累计报告数。
}

type ReportReportListResponse struct {
	Result ReportListResponseResult `json:"result"`
}

// 服务。
type DetailsService struct {
	Items []FinanceCashFlowStatementListResponseService `json:"items"` // 细节。
	Total float64 `json:"total"` // 总额。
}

type ReportCreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

// 报告类型： - `ALL`— 所有报告， - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTING...
type ReportListRequestReportType string

// 阶段。
type V3FinanceCashFlowStatementListResponsePeriod struct {
	Begin string `json:"begin"` // 开始阶段。
	End string `json:"end"` // 结束阶段。
	ID int64 `json:"id"` // 时期识别码。
}

type FinanceCashFlowStatementListResponseCashFlow struct {
	CurrencyCode string `json:"currency_code"` // 佣金计算的货币代码。
	Period V3FinanceCashFlowStatementListResponsePeriod `json:"period"`
	OrdersAmount float64 `json:"orders_amount"` // 已成交商品的价格总和。
	ReturnsAmount float64 `json:"returns_amount"` // 退货价格总和。
	CommissionAmount float64 `json:"commission_amount"` // 商品销售Ozon佣金。
	ServicesAmount float64 `json:"services_amount"` // 附加服务数额。
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"` // 物流服务数额。
}

// 已按时期支付。
type DetailsPayment struct {
	CurrencyCode string `json:"currency_code"` // 货币。
	Payment float64 `json:"payment"` // 支付金额。
}

// 细节。
// 操作名称
type FinanceCashFlowStatementListResponseDetailsOthersNameEnum string
const (
	FinanceCashFlowStatementListResponseDetailsOthersNameEnum_MarketplaceRedistributionOfAcquiringOperation FinanceCashFlowStatementListResponseDetailsOthersNameEnum = "MarketplaceRedistributionOfAcquiringOperation"
	FinanceCashFlowStatementListResponseDetailsOthersNameEnum_MarketplaceSellerCompensationLossOfGoodsOperation FinanceCashFlowStatementListResponseDetailsOthersNameEnum = "MarketplaceSellerCompensationLossOfGoodsOperation"
	FinanceCashFlowStatementListResponseDetailsOthersNameEnum_MarketplaceSellerCorrectionOperation FinanceCashFlowStatementListResponseDetailsOthersNameEnum = "MarketplaceSellerCorrectionOperation"
	FinanceCashFlowStatementListResponseDetailsOthersNameEnum_OperationCorrectionSeller FinanceCashFlowStatementListResponseDetailsOthersNameEnum = "OperationCorrectionSeller"
	FinanceCashFlowStatementListResponseDetailsOthersNameEnum_OperationMarketplaceWithHoldingForUndeliverableGoods FinanceCashFlowStatementListResponseDetailsOthersNameEnum = "OperationMarketplaceWithHoldingForUndeliverableGoods"
	FinanceCashFlowStatementListResponseDetailsOthersNameEnum_OperationClaim FinanceCashFlowStatementListResponseDetailsOthersNameEnum = "OperationClaim"
)

type FinanceCashFlowStatementListResponseDetailsOthers struct {
	Name FinanceCashFlowStatementListResponseDetailsOthersNameEnum `json:"name"` // 操作名称。可能的值： - `MarketplaceRedistributionOfAcquiringOperation` — 收单业务支付， - `MarketplaceSellerCompensationLossOfGoodsOperat...
	Price float64 `json:"price"` // 操作金额。
}

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

// 过滤器。
type ReportCreateCompanyPostingsReportRequestFilter struct {
	WarehouseID []int64 `json:"warehouse_id"` // 仓库标识符。
	CancelReasonID []int64 `json:"cancel_reason_id"` // 取消原因的识别码。
	DeliverySchema []string `json:"delivery_schema"` // 运作机制是FBO或FBS。 对于海外卖家来说，只有FBS方案可用，所以在参数中提交数值`fbs`。
	ProcessedAtTo string `json:"processed_at_to"` // 订单出现在个人账户的时间。
	StatusAlias []string `json:"status_alias"` // 状态文本。
	DeliveryMethodID []int64 `json:"delivery_method_id"` // 配送方法标识符。 可通过方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	IsExpress any `json:"is_express"` // 快递配送： - `true`—仅包含使用 Ozon Express 快速配送的货件； - `false`—仅包含未使用 Ozon Express 快速配送的货件。 如果未传递任何值，将返回所有货件记录。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
	ProcessedAtFrom string `json:"processed_at_from"` // 订单进入处理程序的时间。
	SKU []int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	Statuses []int64 `json:"statuses"` // 数值状况。
	Title string `json:"title"` // 商品名称。
}

type ReportCreateCompanyPostingsReportRequest struct {
	Filter ReportCreateCompanyPostingsReportRequestFilter `json:"filter"`
	Language ReportLanguage `json:"language"`
	With CreateCompanyPostingsReportRequestWith `json:"with"`
}

type V1CreateStockByWarehouseReportRequest struct {
	Language ReportLanguage `json:"language"`
	WarehouseId []string `json:"warehouseId"` // 仓库ID。 请求中参数值的限制。 最大值为 50。
}

// 关于报告的信息。
// 报告类型：
type ReportReportinfoReportTypeEnum string
const (
	ReportReportinfoReportTypeEnum_SELLERPRODUCTS ReportReportinfoReportTypeEnum = "SELLER_PRODUCTS"
	ReportReportinfoReportTypeEnum_SELLERSTOCK ReportReportinfoReportTypeEnum = "SELLER_STOCK"
	ReportReportinfoReportTypeEnum_SELLERRETURNS ReportReportinfoReportTypeEnum = "SELLER_RETURNS"
	ReportReportinfoReportTypeEnum_SELLERPOSTINGS ReportReportinfoReportTypeEnum = "SELLER_POSTINGS"
	ReportReportinfoReportTypeEnum_SELLERDISCOUNTED ReportReportinfoReportTypeEnum = "SELLER_DISCOUNTED"
	ReportReportinfoReportTypeEnum_MUTUALSETTLEMENT ReportReportinfoReportTypeEnum = "MUTUAL_SETTLEMENT"
	ReportReportinfoReportTypeEnum_DOCUMENTB2BSALES ReportReportinfoReportTypeEnum = "DOCUMENT_B2B_SALES"
	ReportReportinfoReportTypeEnum_COMPENSATIONREPORT ReportReportinfoReportTypeEnum = "COMPENSATION_REPORT"
	ReportReportinfoReportTypeEnum_DECOMPENSATIONREPORT ReportReportinfoReportTypeEnum = "DECOMPENSATION_REPORT"
	ReportReportinfoReportTypeEnum_MARKEDPRODUCTSSALES ReportReportinfoReportTypeEnum = "MARKED_PRODUCTS_SALES"
	ReportReportinfoReportTypeEnum_SELLERPLACEMENTBYPRODUCTS ReportReportinfoReportTypeEnum = "SELLER_PLACEMENT_BY_PRODUCTS"
	ReportReportinfoReportTypeEnum_SELLERPLACEMENTBYSUPPLIES ReportReportinfoReportTypeEnum = "SELLER_PLACEMENT_BY_SUPPLIES"
)

// 报告生成状态：
type ReportReportinfoStatusEnum string
const (
	ReportReportinfoStatusEnum_Waiting ReportReportinfoStatusEnum = "waiting"
	ReportReportinfoStatusEnum_Processing ReportReportinfoStatusEnum = "processing"
	ReportReportinfoStatusEnum_Success ReportReportinfoStatusEnum = "success"
	ReportReportinfoStatusEnum_Failed ReportReportinfoStatusEnum = "failed"
)

type ReportReportinfo struct {
	Code string `json:"code"` // 报告的唯一识别码。
	CreatedAt string `json:"created_at"` // 报告创建日期。
	Error string `json:"error"` // 生成报告时的错误代码。
	ExpiresAt string `json:"expires_at"` // 报告链接的有效日期和时间。 如果报告生成于 2025 年 10 月 14 日之前，该字段将为空。
	File string `json:"file"` // XLSX文件的链接。 `SELLER_RETURNS` 类型的报告，链接有效期为5分钟。
	Params map[string]any `json:"params"` // 一个数组，包含卖家创建报告时指定的过滤器。
	ReportType ReportReportinfoReportTypeEnum `json:"report_type"` // 报告类型： - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTINGS` — 发货报告， - `S...
	Status ReportReportinfoStatusEnum `json:"status"` // 报告生成状态： - `waiting`—在等待队列中待处理， - `processing`—正在处理， - `success`， - `failed`。
}

// 报告生成周期。
type ReportMarkedProductsSalesCreateRequestDate struct {
	From string `json:"from"` // 报告期开始日期，格式为 `YYYY-MM-DD`。
	To string `json:"to"` // 报告期结束日期，格式为 `YYYY-MM-DD`。
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date ReportMarkedProductsSalesCreateRequestDate `json:"date"`
}

// 补贴和其他津贴。
type DetailsOthers struct {
	Total float64 `json:"total"` // 总数。
	Items []FinanceCashFlowStatementListResponseDetailsOthers `json:"items"` // 细节。
}

// 细节。
// 操作名称
type FinanceCashFlowStatementListResponseDeliveryServiceNameEnum string
const (
	FinanceCashFlowStatementListResponseDeliveryServiceNameEnum_MarketplaceServiceItemDirectFlowLogisticSum FinanceCashFlowStatementListResponseDeliveryServiceNameEnum = "MarketplaceServiceItemDirectFlowLogisticSum"
	FinanceCashFlowStatementListResponseDeliveryServiceNameEnum_MarketplaceServiceItemDropoff FinanceCashFlowStatementListResponseDeliveryServiceNameEnum = "MarketplaceServiceItemDropoff"
	FinanceCashFlowStatementListResponseDeliveryServiceNameEnum_MarketplaceServiceItemDelivToCustomer FinanceCashFlowStatementListResponseDeliveryServiceNameEnum = "MarketplaceServiceItemDelivToCustomer"
)

type FinanceCashFlowStatementListResponseDeliveryService struct {
	Name FinanceCashFlowStatementListResponseDeliveryServiceNameEnum `json:"name"` // 操作名称。可能的值： - `MarketplaceServiceItemDirectFlowLogisticSum` — 物流， - `MarketplaceServiceItemDropoff` — Drop-off发货处理， - `Ma...
	Price float64 `json:"price"` // 操作总和。
}

// 加工费和运费报酬。
type DetailsServices struct {
	Total float64 `json:"total"` // 总额。
	Items []FinanceCashFlowStatementListResponseDeliveryService `json:"items"` // 细节。
}

// 订单。
type DetailsDeliveryDetails struct {
	Total float64 `json:"total"` // 总额。
	Amount float64 `json:"amount"` // 考虑佣金商品购买的价格。
	DeliveryServices DetailsServices `json:"delivery_services"` // 加工费和运费。
}

// 详细信息。
type FinanceCashFlowStatementListResponseDetails struct {
	Delivery DetailsDeliveryDetails `json:"delivery"` // 方法操作结果。
	Payments DetailsPayment `json:"payments"` // 期间已付清。
	Return DetailsReturnDetails `json:"return"`
	RFBS DetailsRfbsDetails `json:"rfbs"` // rFBS框架下的转账金额。
	Services DetailsService `json:"services"` // 服务。
	Others DetailsOthers `json:"others"` // 补偿费和其他费用。
	EndBalanceAmount float64 `json:"end_balance_amount"` // 结束阶段的收支。
	BeginBalanceAmount float64 `json:"begin_balance_amount"` // 开始阶段的收支。
	InvoiceTransfer float64 `json:"invoice_transfer"` // 当期应付金额。
	Loan float64 `json:"loan"` // 根据贷款协议转账。
	Period V3FinanceCashFlowStatementListResponsePeriod `json:"period"` // 时期。
}

// 方法操作结果。
type V3FinanceCashFlowStatementListResponseResult struct {
	CashFlows []FinanceCashFlowStatementListResponseCashFlow `json:"cash_flows"` // 报告清单。
	Details []FinanceCashFlowStatementListResponseDetails `json:"details"` // 细节信息。
	PageCount int64 `json:"page_count"` // 含有报告的页数。
}

type ReportReportInfoResponse struct {
	Result ReportReportinfo `json:"result"`
}

type ReportReportListRequest struct {
	Page int32 `json:"page"` // 页数。
	PageSize int32 `json:"page_size"` // 每页的值的数量： - 默认值 — 100， - 最大值 — 1,000。
	ReportType ReportListRequestReportType `json:"report_type"`
}

// 报告期限。
type Financev3Period struct {
	From string `json:"from"` // 计算报告的起始日期。
	To string `json:"to"` // 计算报告的停止日期。
}

type V3FinanceCashFlowStatementListRequest struct {
	Date Financev3Period `json:"date"`
	Page int32 `json:"page"` // 请求返回中的页码。
	WithDetails bool `json:"with_details"` // `true`，如果需要在响应中添加附加参数。
	PageSize int32 `json:"page_size"` // 页面上的元素数量。
}

type ReportCreateDiscountedRequest any

type CommonCreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

type V3FinanceCashFlowStatementListResponse struct {
	Result V3FinanceCashFlowStatementListResponseResult `json:"result"`
}

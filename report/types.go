package report

// 细节。
// Name values
type Name string

const (
	NameMarketplaceServiceItemReturnAfterDelivToCustomer Name = "MarketplaceServiceItemReturnAfterDelivToCustomer" // 退货处理
	NameMarketplaceServiceItemReturnPartGoodsCustomer    Name = "MarketplaceServiceItemReturnPartGoodsCustomer"    // 部分非赎回处理
	NameMarketplaceServiceItemReturnNotDelivToCustomer   Name = "MarketplaceServiceItemReturnNotDelivToCustomer"   // 处理已取消和无人认领的商品
	NameMarketplaceServiceItemReturnFlowLogistic         Name = "MarketplaceServiceItemReturnFlowLogistic"         // 回程物流
)

type CashFlowStatementListResponseReturnService struct {
	Name  Name    `json:"name"`  // 操作名称。可能的值： - `MarketplaceServiceItemReturnAfterDelivToCustomer` — 退货处理， - `MarketplaceServiceItemReturnPartGoodsCustomer...
	Price float64 `json:"price"` // 操作数量。
}

// 退货和取消费用。
type DetailsReturns struct {
	Total float64                                      `json:"total"` // 总量。
	Items []CashFlowStatementListResponseReturnService `json:"items"` // 细节。
}

// 退货和取消订单。
type DetailsReturnDetails struct {
	Total          float64        `json:"total"`           // 总金额。
	Amount         float64        `json:"amount"`          // 佣金考虑在内的退款金额。
	ReturnServices DetailsReturns `json:"return_services"` // 退款和取消的费用。
}

// rFBS模式转账金额。
type DetailsRfbsDetails struct {
	Total                      float64 `json:"total"`                        // 总额。
	TransferDelivery           float64 `json:"transfer_delivery"`            // 来自买家的转账金额。
	TransferDeliveryReturn     float64 `json:"transfer_delivery_return"`     // 退还给买家的转账金额。
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"` // 物流转账金额补贴。
	PartialCompensation        float64 `json:"partial_compensation"`         // 将部分退款转移给买家。
	PartialCompensationReturn  float64 `json:"partial_compensation_return"`  // 退换部分退款。
}

// 细节。
type CashFlowStatementListResponseService struct {
	Name  Name    `json:"name"`  // 操作名称。可能的值： - `MarketplaceServiceItemElectronicServiceStencil` — “模板”服务， - `MarketplaceServiceItemElectronicServicesPromo...
	Price float64 `json:"price"` // 操作数量。
}

// 服务。
type DetailsService struct {
	Total float64                                `json:"total"` // 总额。
	Items []CashFlowStatementListResponseService `json:"items"` // 细节。
}

// 细节。
type CashFlowStatementListResponseDetailsOthers struct {
	Name  Name    `json:"name"`  // 操作名称。可能的值： - `MarketplaceRedistributionOfAcquiringOperation` — 收单业务支付， - `MarketplaceSellerCompensationLossOfGoodsOperat...
	Price float64 `json:"price"` // 操作金额。
}

// 补贴和其他津贴。
type DetailsOthers struct {
	Total float64                                      `json:"total"` // 总数。
	Items []CashFlowStatementListResponseDetailsOthers `json:"items"` // 细节。
}

// 细节。
type CashFlowStatementListResponseDeliveryService struct {
	Name  Name    `json:"name"`  // 操作名称。可能的值： - `MarketplaceServiceItemDirectFlowLogisticSum` — 物流， - `MarketplaceServiceItemDropoff` — Drop-off发货处理， - `Ma...
	Price float64 `json:"price"` // 操作总和。
}

// 加工费和运费报酬。
type DetailsServices struct {
	Total float64                                        `json:"total"` // 总额。
	Items []CashFlowStatementListResponseDeliveryService `json:"items"` // 细节。
}

// 订单。
type DetailsDeliveryDetails struct {
	DeliveryServices DetailsServices `json:"delivery_services"` // 加工费和运费。
	Total            float64         `json:"total"`             // 总额。
	Amount           float64         `json:"amount"`            // 考虑佣金商品购买的价格。
}

// 已按时期支付。
type DetailsPayment struct {
	CurrencyCode string  `json:"currency_code"` // 货币。
	Payment      float64 `json:"payment"`       // 支付金额。
}

// 阶段。
type V3FinanceCashFlowStatementListResponsePeriod struct {
	ID    int64  `json:"id"`    // 时期识别码。
	Begin string `json:"begin"` // 开始阶段。
	End   string `json:"end"`   // 结束阶段。
}

// 详细信息。
type CashFlowStatementListResponseDetails struct {
	BeginBalanceAmount float64                                      `json:"begin_balance_amount"` // 开始阶段的收支。
	InvoiceTransfer    float64                                      `json:"invoice_transfer"`     // 当期应付金额。
	Loan               float64                                      `json:"loan"`                 // 根据贷款协议转账。
	RFBS               DetailsRfbsDetails                           `json:"rfbs"`                 // rFBS框架下的转账金额。
	Services           DetailsService                               `json:"services"`             // 服务。
	Others             DetailsOthers                                `json:"others"`               // 补偿费和其他费用。
	Delivery           DetailsDeliveryDetails                       `json:"delivery"`             // 方法操作结果。
	Payments           DetailsPayment                               `json:"payments"`             // 期间已付清。
	Period             V3FinanceCashFlowStatementListResponsePeriod `json:"period"`               // 时期。
	Return             DetailsReturnDetails                         `json:"return"`
	EndBalanceAmount   float64                                      `json:"end_balance_amount"` // 结束阶段的收支。
}

type CashFlowStatementListResponseCashFlow struct {
	CurrencyCode                string                                       `json:"currency_code"` // 佣金计算的货币代码。
	Period                      V3FinanceCashFlowStatementListResponsePeriod `json:"period"`
	OrdersAmount                float64                                      `json:"orders_amount"`                   // 已成交商品的价格总和。
	ReturnsAmount               float64                                      `json:"returns_amount"`                  // 退货价格总和。
	CommissionAmount            float64                                      `json:"commission_amount"`               // 商品销售Ozon佣金。
	ServicesAmount              float64                                      `json:"services_amount"`                 // 附加服务数额。
	ItemDeliveryAndReturnAmount float64                                      `json:"item_delivery_and_return_amount"` // 物流服务数额。
}

// 请求结果。
type CreateReportResponseCode struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

type CommonCreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

// 报告生成周期。
type MarkedProductsSalesCreateRequestDate struct {
	From string `json:"from"` // 报告期开始日期，格式为 `YYYY-MM-DD`。
	To   string `json:"to"`   // 报告期结束日期，格式为 `YYYY-MM-DD`。
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date MarkedProductsSalesCreateRequestDate `json:"date"`
}

// 关于报告的信息。
// Status values
type Status string

const (
	StatusWaiting    Status = "waiting"    // 在等待队列中待处理
	StatusProcessing Status = "processing" // 正在处理
	StatusSuccess    Status = "success"    // 报告成功生成
	StatusFailed     Status = "failed"     // 报告生成错误
)

// ReportType values
type ReportType string

const (
	ReportTypeSellerProducts            ReportType = "SELLER_PRODUCTS"              // 商品报告
	ReportTypeSellerStock               ReportType = "SELLER_STOCK"                 // 商品库存报告
	ReportTypeSellerReturns             ReportType = "SELLER_RETURNS"               // 退货报告
	ReportTypeSellerPostings            ReportType = "SELLER_POSTINGS"              // 发货报告
	ReportTypeSellerDiscounted          ReportType = "SELLER_DISCOUNTED"            // 减价商品报告
	ReportTypeMutualSettlement          ReportType = "MUTUAL_SETTLEMENT"            // 结算报告
	ReportTypeDocumentB2BSales          ReportType = "DOCUMENT_B2B_SALES"           // 面向法人客户的销售报告
	ReportTypeCompensationReport        ReportType = "COMPENSATION_REPORT"          // 赔偿报告
	ReportTypeDecompensationReport      ReportType = "DECOMPENSATION_REPORT"        // 赔偿返还报告
	ReportTypeMarkedProductsSales       ReportType = "MARKED_PRODUCTS_SALES"        // 标签销售报告
	ReportTypeSellerPlacementBYProducts ReportType = "SELLER_PLACEMENT_BY_PRODUCTS" // 按商品维度的存储服务费用报告
	ReportTypeSellerPlacementBYSupplies ReportType = "SELLER_PLACEMENT_BY_SUPPLIES" // 按交货维度的存储服务费用报告
)

type Report struct {
	Code       string         `json:"code"`        // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
	CreatedAt  string         `json:"created_at"`  // 报告创建日期。
	Error      string         `json:"error"`       // 生成报告时的错误代码。
	ExpiresAt  string         `json:"expires_at"`  // 报告链接的有效日期和时间。 如果报告生成于 2025 年 10 月 14 日之前，该字段将为空。
	File       string         `json:"file"`        // XLSX文件的链接。 `SELLER_RETURNS` 类型的报告，链接有效期为5分钟。
	Params     map[string]any `json:"params"`      // 一个数组，包含卖家创建报告时指定的过滤器。
	ReportType ReportType     `json:"report_type"` // 报告类型： - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTINGS` — 发货报告， - `S...
	Status     Status         `json:"status"`      // 报告生成状态： - `waiting`—在等待队列中待处理， - `processing`—正在处理， - `success`—报告成功生成， - `failed` — 报告生成错误。
}

// 请求结果。
type ListResponseResult struct {
	Reports []Report `json:"reports"` // 包含所有生成的报告的数组。
	Total   int32    `json:"total"`   // 累计报告数。
}

type ReportListResponse struct {
	Result ListResponseResult `json:"result"`
}

type CreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

type ReportInfoRequest struct {
	Code string `json:"code"` // 报告的唯一识别码。
}

// 按商品可见度过滤。 - `ALL`——除了档案中的所有商品； - `VALIDATION_STATE_FAIL`——预审时未被验证器检查的商品； - `TO_SUPPLY`——准备出售的货物； - `IN_SALE`——正在销售的商品； -...
type CreateCompanyProductsReportRequestVisibility string

// 回答所用语言： - `RU` — 俄语， - `EN` — 英语。
type Language string

type CreateCompanyProductsReportRequest struct {
	Language   Language                                     `json:"language"`
	OfferID    []string                                     `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
	Search     string                                       `json:"search"`   // 在记录内容中搜索，检查现货。
	SKU        []int64                                      `json:"sku"`      // Ozon 系统中的商品标识符（SKU）。
	Visibility CreateCompanyProductsReportRequestVisibility `json:"visibility"`
}

// 额外的字段，需要添加到响应中。
type CreateCompanyPostingsReportRequestWith struct {
	AdditionalData bool `json:"additional_data"` // `true`，用于在响应中添加附加信息。
	AnalyticsData  bool `json:"analytics_data"`  // `true`，用于在响应中添加分析数据。请传递值 `filter.delivery_schema = fbs`，否则 将返回 错误。
	CustomerData   bool `json:"customer_data"`   // `true`，用于在响应中添加买家信息。
	JewelryCodes   bool `json:"jewelry_codes"`   // `true`，用于在响应中添加珠宝信息。
}

// 关于报告的信息。
type Reportinfo struct {
	Error      string         `json:"error"`       // 生成报告时的错误代码。
	ExpiresAt  string         `json:"expires_at"`  // 报告链接的有效日期和时间。 如果报告生成于 2025 年 10 月 14 日之前，该字段将为空。
	File       string         `json:"file"`        // XLSX文件的链接。 `SELLER_RETURNS` 类型的报告，链接有效期为5分钟。
	Params     map[string]any `json:"params"`      // 一个数组，包含卖家创建报告时指定的过滤器。
	ReportType ReportType     `json:"report_type"` // 报告类型： - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTINGS` — 发货报告， - `S...
	Status     Status         `json:"status"`      // 报告生成状态： - `waiting`—在等待队列中待处理， - `processing`—正在处理， - `success`， - `failed`。
	Code       string         `json:"code"`        // 报告的唯一识别码。
	CreatedAt  string         `json:"created_at"`  // 报告创建日期。
}

type CreateDiscountedResponse struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

// 报告类型： - `ALL`— 所有报告， - `SELLER_PRODUCTS` — 商品报告， - `SELLER_STOCK` — 商品库存报告， - `SELLER_RETURNS` — 退货报告， - `SELLER_POSTING...
type ListRequestReportType string

type ReportListRequest struct {
	Page       int32                 `json:"page"`      // 页数。
	PageSize   int32                 `json:"page_size"` // 每页的值的数量： - 默认值 — 100， - 最大值 — 1,000。
	ReportType ListRequestReportType `json:"report_type"`
}

// 方法操作结果。
type V3FinanceCashFlowStatementListResponseResult struct {
	Details   []CashFlowStatementListResponseDetails  `json:"details"`    // 细节信息。
	PageCount int64                                   `json:"page_count"` // 含有报告的页数。
	CashFlows []CashFlowStatementListResponseCashFlow `json:"cash_flows"` // 报告清单。
}

type V3FinanceCashFlowStatementListResponse struct {
	Result V3FinanceCashFlowStatementListResponseResult `json:"result"`
}

// 报告期限。
type V3Period struct {
	To   string `json:"to"`   // 计算报告的停止日期。
	From string `json:"from"` // 计算报告的起始日期。
}

type CreateDiscountedRequest any

// 过滤器。
type CreateCompanyPostingsReportRequestFilter struct {
	OfferID          string   `json:"offer_id"`           // 卖家系统中的商品标识符是商品货号。
	ProcessedAtFrom  string   `json:"processed_at_from"`  // 订单进入处理程序的时间。
	ProcessedAtTo    string   `json:"processed_at_to"`    // 订单出现在个人账户的时间。
	SKU              []int64  `json:"sku"`                // Ozon 系统中的商品标识符（SKU）。
	Statuses         []int64  `json:"statuses"`           // 数值状况。
	Title            string   `json:"title"`              // 商品名称。
	StatusAlias      []string `json:"status_alias"`       // 状态文本。
	WarehouseID      []int64  `json:"warehouse_id"`       // 仓库标识符。
	DeliveryMethodID []int64  `json:"delivery_method_id"` // 配送方法标识符。 可通过方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	IsExpress        any      `json:"is_express"`         // 快递配送： - `true`—仅包含使用 Ozon Express 快速配送的货件； - `false`—仅包含未使用 Ozon Express 快速配送的货件。 如果未传递任何值，将返回所有货件记录。
	CancelReasonID   []int64  `json:"cancel_reason_id"`   // 取消原因的识别码。
	DeliverySchema   []string `json:"delivery_schema"`    // 运作机制是FBO或FBS。 对于海外卖家来说，只有FBS方案可用，所以在参数中提交数值`fbs`。
}

type ReportInfoResponse struct {
	Result Reportinfo `json:"result"`
}

type CreateCompanyPostingsReportRequest struct {
	Language Language                                 `json:"language"`
	With     CreateCompanyPostingsReportRequestWith   `json:"with"`
	Filter   CreateCompanyPostingsReportRequestFilter `json:"filter"`
}

type V3FinanceCashFlowStatementListRequest struct {
	Date        V3Period `json:"date"`
	Page        int32    `json:"page"`         // 请求返回中的页码。
	WithDetails bool     `json:"with_details"` // `true`，如果需要在响应中添加附加参数。
	PageSize    int32    `json:"page_size"`    // 页面上的元素数量。
}

type V1CreateStockByWarehouseReportRequest struct {
	WarehouseId []string `json:"warehouseId"` // 仓库ID。 请求中参数值的限制。 最大值为 50。
	Language    Language `json:"language"`
}

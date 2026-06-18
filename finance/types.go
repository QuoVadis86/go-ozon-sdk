package finance

// 按日期过滤。
type FinanceTransactionTotalsV3RequestDate struct {
	From string `json:"from"` // 阶段开始。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
	To string `json:"to"` // 阶段结束。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
}

type Financev3FinanceTransactionTotalsV3Request struct {
	Date FinanceTransactionTotalsV3RequestDate `json:"date"`
	PostingNumber string `json:"posting_number"` // 发货号。
	TransactionType string `json:"transaction_type"` // 操作类型： - `all` — 所有, - `orders` — 订单, - `returns` — 退货和取消, - `services` — 服务费, - `compensation` — 补贴, - `transferDelivery...
}

// 关于向法人销售的信息。
type RowItemLegalEntityDocument struct {
	Number string `json:"number"` // 发票编号。
	SaleDate string `json:"sale_date"` // 日期格式为`YYYY-MM-DD`。
}

// 配送佣金。
type RowItemCommission struct {
	Compensation float64 `json:"compensation"` // Ozon 负责的补付额。 适用于 2024 年 4 月 30 日之前生成的报告。
	PricePerInstance float64 `json:"price_per_instance"` // 每件价格。
	Quantity int32 `json:"quantity"` // 商品数量。
	StandardFee float64 `json:"standard_fee"` // Ozon基础奖励。
	BankCoinvestment float64 `json:"bank_coinvestment"` // 合作伙伴忠诚机制付款：绿色价格。
	Total float64 `json:"total"` // 应计总额。
	Commission float64 `json:"commission"` // 将折扣和附加费考虑在内的总佣金。 适用于 2024 年 4 月 30 日之前生成的报告。
	Stars float64 `json:"stars"` // 合作伙伴忠诚度机制付款：星星。
	Amount float64 `json:"amount"` // 金额。
	Bonus float64 `json:"bonus"` // 折扣积分。
}

// 商品信息。
type RowItem struct {
	SKU int64 `json:"sku"` // Ozon系统中的商品识别码是SKU。
	Barcode string `json:"barcode"` // 商品条形码。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
}

// 商品退货佣金。
type RowItemCommissionReturn struct {
	PricePerInstance float64 `json:"price_per_instance"` // 每件价格。
	Quantity int32 `json:"quantity"` // 商品数量。
	BankCoinvestment float64 `json:"bank_coinvestment"` // 合作伙伴忠诚机制付款：绿色价格。
	Stars float64 `json:"stars"` // 合作伙伴忠诚度机制付款：星星。
	Commission float64 `json:"commission"` // 将折扣和附加费考虑在内的总佣金。 适用于 2024 年 4 月 30 日之前生成的报告。
	Compensation float64 `json:"compensation"` // Ozon 负责的补付额。 适用于 2024 年 4 月 30 日之前生成的报告。
	StandardFee float64 `json:"standard_fee"` // Ozon基础奖励。
	Total float64 `json:"total"` // 应计总额。
	Amount float64 `json:"amount"` // 金额。
	Bonus float64 `json:"bonus"` // 折扣积分。
}

// 订单信息。
type RowItemOrder struct {
	CreatedDate string `json:"created_date"` // 订单日期格式为`YYYY-MM-DD`。
	PostingNumber string `json:"posting_number"` // 货件编号。
}

type V1GetRealizationReportPostingResponseRow struct {
	LegalEntityDocument RowItemLegalEntityDocument `json:"legal_entity_document"`
	CommissionRatio float64 `json:"commission_ratio"` // 按类目划分的销售佣金比例。
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
	Item RowItem `json:"item"`
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"row_number"` // 报告中的行号。
	SellerPricePerInstance float64 `json:"seller_price_per_instance"` // 考虑折扣后的卖家价格。
	Order RowItemOrder `json:"order"`
}

type GetRealizationReportResponseV2Row struct {
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
	Item RowItem `json:"item"`
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"` // 报告中的行编号。
	SellerPricePerInstance float64 `json:"seller_price_per_instance"` // 将折扣考虑在内的卖家价格。
	CommissionRatio float64 `json:"commission_ratio"` // 按类目分类的销售佣金份额。
}

// 日期过滤。
type FilterPeriod struct {
	From string `json:"from"` // 阶段开始。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
	To string `json:"to"` // 阶段结束。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
}

// 过滤器。
type FinanceTransactionListV3RequestFilter struct {
	Date FilterPeriod `json:"date"`
	OperationType []string `json:"operation_type"` // 交易类型: - `ClientReturnAgentOperation` — 收到买家的退货、取消订单、非赎金； - `MarketplaceMarketingActionCostOperation` — 商品促销服务； - `Market...
	PostingNumber string `json:"posting_number"` // 发货号。
	TransactionType string `json:"transaction_type"` // 收费类型： - `all` — 所有, - `orders` — 订单, - `returns` — 退货和取消, - `services` — 服务费, - `compensation` — 补贴, - `transferDelivery...
}

type Financev3FinanceTransactionListV3Request struct {
	Filter FinanceTransactionListV3RequestFilter `json:"filter"`
	Page int64 `json:"page"` // 请求中返回的页码。
	PageSize int64 `json:"page_size"` // 每页的元素数。
}

// 报告语言： - `RU` — 俄语， - `EN` — 英语。
type CompensationReportLanguage string

type V1GetCompensationReportRequest struct {
	Date string `json:"date"` // 报告周期格式为 `YYYY-MM`。
	Language CompensationReportLanguage `json:"language"`
}

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"` // 月。
	Year int32 `json:"year"` // 年。
}

// 报告标题页。
type GetRealizationReportResponseV2Header struct {
	StartDate string `json:"start_date"` // 期间开始。
	StopDate string `json:"stop_date"` // 期间结束。
	ContractDate string `json:"contract_date"` // 合同签订日期。
	ContractNumber string `json:"contract_number"` // 合同编号。
	DocDate string `json:"doc_date"` // 文件生成日期。
	PayerInn string `json:"payer_inn"` // 付款人的纳税人识别号。
	PayerKpp string `json:"payer_kpp"` // 付款人的纳税人登记原因代码。
	ReceiverInn string `json:"receiver_inn"` // 收款人的纳税人识别号。
	ReceiverKpp string `json:"receiver_kpp"` // 收款人的纳税人登记原因代码。
	ReceiverName string `json:"receiver_name"` // 收款人名称。
	CurrencySysName string `json:"currency_sys_name"` // 货币。
	Number string `json:"number"` // 销售报告编号。
	PayerName string `json:"payer_name"` // 付款人名称。
}

// 运输信息。
type OperationPosting struct {
	DeliverySchema string `json:"delivery_schema"` // 发货方案： - `FBO` — 从Ozon仓库发货， - `FBS` — 从您的仓库发货， - `RFBS` — 按买方选择发货， - `CROSSBORDER` — 跨境配送， - `FBP` — 通过 Ozon 合作仓库配送， - `F...
	OrderDate string `json:"order_date"` // 接收处理货物日期。
	PostingNumber string `json:"posting_number"` // 发货号。
	WarehouseID int64 `json:"warehouse_id"` // 仓库ID。
}

type OperationService struct {
	Name string `json:"name"` // 服务名称： - `MarketplaceNotDeliveredCostItem` — 将买家处将无人认领的货物退回仓库。 - `MarketplaceReturnAfterDeliveryCostItem` — 快递签收后从买家处退回仓库...
	Price float64 `json:"price"` // 价格。
}

type OperationItem struct {
	Name string `json:"name"` // 商品名称。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
}

type FinanceTransactionListV3ResponseOperation struct {
	Amount float64 `json:"amount"` // 交易总额。
	DeliveryCharge float64 `json:"delivery_charge"` // 适用于2021年2月1日之前有效的关税以及大件商品的运费。
	OperationDate string `json:"operation_date"` // 操作日期。
	OperationID int64 `json:"operation_id"` // 操作ID。
	OperationType string `json:"operation_type"` // 操作类型。
	Type string `json:"type_"` // 收费类型： - `all` — 所有, - `orders` — 订单, - `returns` — 退货和取消订单, - `services` — 服务费, - `compensation` — 补贴, - `transferDelive...
	AccrualsForSale float64 `json:"accruals_for_sale"` // 考虑到卖家折扣的商品成本。
	Items []OperationItem `json:"items"` // 商品信息。
	OperationTypeName string `json:"operation_type_name"` // 操作类型名称。
	Posting OperationPosting `json:"posting"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"` // 退货和取消订单费用适用于2021年2月1日之前有效的费率，以及超大商品的费用。
	SaleCommission float64 `json:"sale_commission"` // 销售提成或销售提成返还。
	Services []OperationService `json:"services"` // 附加服务。
}

// 询问结果。
type Financev3FinanceTransactionListV3ResponseResult struct {
	Operations []FinanceTransactionListV3ResponseOperation `json:"operations"` // 操作信息。
	PageCount int64 `json:"page_count"` // 页数。如果为0，则说明已无页面。
	RowCount int64 `json:"row_count"` // 所有页面上的交易数量。如果为0，说明已无交易。
}

// 询问结果。
type Financev3FinanceTransactionTotalsV3ResponseResult struct {
	OthersAmount float64 `json:"others_amount"` // 其他应计费用。
	ProcessingAndDelivery float64 `json:"processing_and_delivery"` // 运输处理、订单装配、干线、最后一英里以及自2021年2月1日起引入新的佣金和费率前的快递服务费。 干线 —— 集群之间的货物交付。 最后一英里 —— 从订单交付点、自提点和快递员到买家处的快递。
	RefundsAndCancellations float64 `json:"refunds_and_cancellations"` // 干线返回、退货处理、取消和非赎回、2021年2月1日起引入新佣金和税率之前退货价格。 干线 —— 集群之间的货物交付。 最后一英里 —— 从订单交付点、自提点和快递员到买家处的快递。
	SaleCommission float64 `json:"sale_commission"` // 商品预售时预扣的佣金数额，退货时返还的佣金数。
	ServicesAmount float64 `json:"services_amount"` // 与商品交付和退货没有直接关系的附加服务成本。例如，促销或商品放置。
	AccrualsForSale float64 `json:"accruals_for_sale"` // 指定期间内商品的总成本和退货。
	CompensationAmount float64 `json:"compensation_amount"` // 补贴。
	MoneyTransfer float64 `json:"money_transfer"` // 根据“卖方选择交货”计划工作时的交货和退货费用。
}

// 请求结果。
type CreateReportResponseCodeNoDeadline struct {
	Code string `json:"code"` // 报告的唯一标识符。要获取报告，请将该值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

// 请求结果。
type GetRealizationReportResponseV2Result struct {
	Header GetRealizationReportResponseV2Header `json:"header"`
	Rows []GetRealizationReportResponseV2Row `json:"rows"` // 报告表格。
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result Financev3FinanceTransactionTotalsV3ResponseResult `json:"result"`
}

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"` // 报告周期格式为 `YYYY-MM`。
	Language CompensationReportLanguage `json:"language"`
}

type V1GetRealizationReportPostingRequest struct {
	Month int32 `json:"month"` // 月。
	Year int32 `json:"year"` // 年。
}

// 查询结果。
type V1GetRealizationReportPostingResponse struct {
	Header GetRealizationReportResponseV2Header `json:"header"`
	Rows []V1GetRealizationReportPostingResponseRow `json:"rows"` // 报告表格。
}

type CreateReportResponse struct {
	Result CreateReportResponseCodeNoDeadline `json:"result"`
}

type V2GetRealizationReportResponseV2 struct {
	Result GetRealizationReportResponseV2Result `json:"result"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result Financev3FinanceTransactionListV3ResponseResult `json:"result"`
}

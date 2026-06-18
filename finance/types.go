package finance

// 配送佣金。
type RowItemCommission struct {
	Amount float64 `json:"amount"` // 金额。
	Commission float64 `json:"commission"` // 将折扣和附加费考虑在内的总佣金。 适用于 2024 年 4 月 30 日之前生成的报告。
	Compensation float64 `json:"compensation"` // Ozon 负责的补付额。 适用于 2024 年 4 月 30 日之前生成的报告。
	Quantity int32 `json:"quantity"` // 商品数量。
	StandardFee float64 `json:"standard_fee"` // Ozon基础奖励。
	Total float64 `json:"total"` // 应计总额。
	Bonus float64 `json:"bonus"` // 折扣积分。
	PricePerInstance float64 `json:"price_per_instance"` // 每件价格。
	BankCoinvestment float64 `json:"bank_coinvestment"` // 合作伙伴忠诚机制付款：绿色价格。
	Stars float64 `json:"stars"` // 合作伙伴忠诚度机制付款：星星。
}

// 商品信息。
type RowItem struct {
	Barcode string `json:"barcode"` // 商品条形码。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
	SKU int64 `json:"sku"` // Ozon系统中的商品识别码是SKU。
}

// 商品退货佣金。
type RowItemCommissionReturn struct {
	StandardFee float64 `json:"standard_fee"` // Ozon基础奖励。
	Stars float64 `json:"stars"` // 合作伙伴忠诚度机制付款：星星。
	Total float64 `json:"total"` // 应计总额。
	Amount float64 `json:"amount"` // 金额。
	Bonus float64 `json:"bonus"` // 折扣积分。
	Commission float64 `json:"commission"` // 将折扣和附加费考虑在内的总佣金。 适用于 2024 年 4 月 30 日之前生成的报告。
	Compensation float64 `json:"compensation"` // Ozon 负责的补付额。 适用于 2024 年 4 月 30 日之前生成的报告。
	BankCoinvestment float64 `json:"bank_coinvestment"` // 合作伙伴忠诚机制付款：绿色价格。
	PricePerInstance float64 `json:"price_per_instance"` // 每件价格。
	Quantity int32 `json:"quantity"` // 商品数量。
}

type GetRealizationReportResponseV2Row struct {
	SellerPricePerInstance float64 `json:"seller_price_per_instance"` // 将折扣考虑在内的卖家价格。
	CommissionRatio float64 `json:"commission_ratio"` // 按类目分类的销售佣金份额。
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
	Item RowItem `json:"item"`
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"` // 报告中的行编号。
}

// 运输信息。
// 发货方案：
type OperationPostingDeliverySchemaEnum string
const (
	OperationPostingDeliverySchemaEnum_FBO OperationPostingDeliverySchemaEnum = "FBO"
	OperationPostingDeliverySchemaEnum_FBS OperationPostingDeliverySchemaEnum = "FBS"
	OperationPostingDeliverySchemaEnum_RFBS OperationPostingDeliverySchemaEnum = "RFBS"
	OperationPostingDeliverySchemaEnum_CROSSBORDER OperationPostingDeliverySchemaEnum = "CROSSBORDER"
	OperationPostingDeliverySchemaEnum_FBP OperationPostingDeliverySchemaEnum = "FBP"
	OperationPostingDeliverySchemaEnum_FBOECONOMY OperationPostingDeliverySchemaEnum = "FBOECONOMY"
	OperationPostingDeliverySchemaEnum_FBSECONOMY OperationPostingDeliverySchemaEnum = "FBSECONOMY"
)

type OperationPosting struct {
	PostingNumber string `json:"posting_number"` // 发货号。
	WarehouseID int64 `json:"warehouse_id"` // 仓库ID。
	DeliverySchema OperationPostingDeliverySchemaEnum `json:"delivery_schema"` // 发货方案： - `FBO` — 从Ozon仓库发货， - `FBS` — 从您的仓库发货， - `RFBS` — 按买方选择发货， - `CROSSBORDER` — 跨境配送， - `FBP` — 通过 Ozon 合作仓库配送， - `F...
	OrderDate string `json:"order_date"` // 接收处理货物日期。
}

// 服务名称：
type OperationServiceNameEnum string
const (
	OperationServiceNameEnum_MarketplaceNotDeliveredCostItem OperationServiceNameEnum = "MarketplaceNotDeliveredCostItem"
	OperationServiceNameEnum_MarketplaceReturnAfterDeliveryCostItem OperationServiceNameEnum = "MarketplaceReturnAfterDeliveryCostItem"
	OperationServiceNameEnum_MarketplaceDeliveryCostItem OperationServiceNameEnum = "MarketplaceDeliveryCostItem"
	OperationServiceNameEnum_MarketplaceSaleReviewsItem OperationServiceNameEnum = "MarketplaceSaleReviewsItem"
	OperationServiceNameEnum_ItemAdvertisementForSupplierLogistic OperationServiceNameEnum = "ItemAdvertisementForSupplierLogistic"
	OperationServiceNameEnum_OperationMarketplaceServiceStorage OperationServiceNameEnum = "OperationMarketplaceServiceStorage"
	OperationServiceNameEnum_MarketplaceMarketingActionCostItem OperationServiceNameEnum = "MarketplaceMarketingActionCostItem"
	OperationServiceNameEnum_MarketplaceServiceItemInstallment OperationServiceNameEnum = "MarketplaceServiceItemInstallment"
	OperationServiceNameEnum_MarketplaceServiceItemMarkingItems OperationServiceNameEnum = "MarketplaceServiceItemMarkingItems"
	OperationServiceNameEnum_MarketplaceServiceItemFlexiblePaymentSchedule OperationServiceNameEnum = "MarketplaceServiceItemFlexiblePaymentSchedule"
	OperationServiceNameEnum_MarketplaceServiceItemReturnFromStock OperationServiceNameEnum = "MarketplaceServiceItemReturnFromStock"
	OperationServiceNameEnum_ItemAdvertisementForSupplierLogisticSeller OperationServiceNameEnum = "ItemAdvertisementForSupplierLogisticSeller"
	OperationServiceNameEnum_ItemAgentServiceStarsMembership OperationServiceNameEnum = "ItemAgentServiceStarsMembership"
	OperationServiceNameEnum_MarketplaceServiceItemDelivToCustomer OperationServiceNameEnum = "MarketplaceServiceItemDelivToCustomer"
	OperationServiceNameEnum_MarketplaceServiceItemDirectFlowTrans OperationServiceNameEnum = "MarketplaceServiceItemDirectFlowTrans"
	OperationServiceNameEnum_MarketplaceServiceItemDropoffFF OperationServiceNameEnum = "MarketplaceServiceItemDropoffFF"
	OperationServiceNameEnum_MarketplaceServiceItemDropoffPVZ OperationServiceNameEnum = "MarketplaceServiceItemDropoffPVZ"
	OperationServiceNameEnum_MarketplaceServiceItemDropoffSC OperationServiceNameEnum = "MarketplaceServiceItemDropoffSC"
	OperationServiceNameEnum_MarketplaceServiceItemFulfillment OperationServiceNameEnum = "MarketplaceServiceItemFulfillment"
	OperationServiceNameEnum_MarketplaceServiceItemPickup OperationServiceNameEnum = "MarketplaceServiceItemPickup"
	OperationServiceNameEnum_MarketplaceServiceItemReturnAfterDelivToCustomer OperationServiceNameEnum = "MarketplaceServiceItemReturnAfterDelivToCustomer"
	OperationServiceNameEnum_MarketplaceServiceItemReturnFlowTrans OperationServiceNameEnum = "MarketplaceServiceItemReturnFlowTrans"
	OperationServiceNameEnum_MarketplaceServiceItemReturnNotDelivToCustomer OperationServiceNameEnum = "MarketplaceServiceItemReturnNotDelivToCustomer"
	OperationServiceNameEnum_MarketplaceServiceItemReturnPartGoodsCustomer OperationServiceNameEnum = "MarketplaceServiceItemReturnPartGoodsCustomer"
	OperationServiceNameEnum_MarketplaceServiceItemDropoffPPZ OperationServiceNameEnum = "MarketplaceServiceItemDropoffPPZ"
	OperationServiceNameEnum_MarketplaceServiceItemRedistributionReturnsPVZ OperationServiceNameEnum = "MarketplaceServiceItemRedistributionReturnsPVZ"
)

type OperationService struct {
	Name OperationServiceNameEnum `json:"name"` // 服务名称： - `MarketplaceNotDeliveredCostItem` — 将买家处将无人认领的货物退回仓库。 - `MarketplaceReturnAfterDeliveryCostItem` — 快递签收后从买家处退回仓库...
	Price float64 `json:"price"` // 价格。
}

type OperationItem struct {
	Name string `json:"name"` // 商品名称。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
}

// 收费类型：
type FinanceTransactionListV3ResponseOperationTypeEnum string
const (
	FinanceTransactionListV3ResponseOperationTypeEnum_All FinanceTransactionListV3ResponseOperationTypeEnum = "all"
	FinanceTransactionListV3ResponseOperationTypeEnum_Orders FinanceTransactionListV3ResponseOperationTypeEnum = "orders"
	FinanceTransactionListV3ResponseOperationTypeEnum_Returns FinanceTransactionListV3ResponseOperationTypeEnum = "returns"
	FinanceTransactionListV3ResponseOperationTypeEnum_Services FinanceTransactionListV3ResponseOperationTypeEnum = "services"
	FinanceTransactionListV3ResponseOperationTypeEnum_Compensation FinanceTransactionListV3ResponseOperationTypeEnum = "compensation"
	FinanceTransactionListV3ResponseOperationTypeEnum_TransferDelivery FinanceTransactionListV3ResponseOperationTypeEnum = "transferDelivery"
	FinanceTransactionListV3ResponseOperationTypeEnum_Other FinanceTransactionListV3ResponseOperationTypeEnum = "other"
)

type FinanceTransactionListV3ResponseOperation struct {
	AccrualsForSale float64 `json:"accruals_for_sale"` // 考虑到卖家折扣的商品成本。
	Items []OperationItem `json:"items"` // 商品信息。
	OperationID int64 `json:"operation_id"` // 操作ID。
	OperationTypeName string `json:"operation_type_name"` // 操作类型名称。
	Posting OperationPosting `json:"posting"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"` // 退货和取消订单费用适用于2021年2月1日之前有效的费率，以及超大商品的费用。
	SaleCommission float64 `json:"sale_commission"` // 销售提成或销售提成返还。
	Type FinanceTransactionListV3ResponseOperationTypeEnum `json:"type_"` // 收费类型： - `all` — 所有, - `orders` — 订单, - `returns` — 退货和取消订单, - `services` — 服务费, - `compensation` — 补贴, - `transferDelive...
	Amount float64 `json:"amount"` // 交易总额。
	DeliveryCharge float64 `json:"delivery_charge"` // 适用于2021年2月1日之前有效的关税以及大件商品的运费。
	OperationDate string `json:"operation_date"` // 操作日期。
	OperationType string `json:"operation_type"` // 操作类型。
	Services []OperationService `json:"services"` // 附加服务。
}

type V1GetRealizationReportPostingRequest struct {
	Month int32 `json:"month"` // 月。
	Year int32 `json:"year"` // 年。
}

// 报告标题页。
type GetRealizationReportResponseV2Header struct {
	ReceiverInn string `json:"receiver_inn"` // 收款人的纳税人识别号。
	ContractDate string `json:"contract_date"` // 合同签订日期。
	ContractNumber string `json:"contract_number"` // 合同编号。
	CurrencySysName string `json:"currency_sys_name"` // 货币。
	ReceiverKpp string `json:"receiver_kpp"` // 收款人的纳税人登记原因代码。
	ReceiverName string `json:"receiver_name"` // 收款人名称。
	StartDate string `json:"start_date"` // 期间开始。
	StopDate string `json:"stop_date"` // 期间结束。
	DocDate string `json:"doc_date"` // 文件生成日期。
	Number string `json:"number"` // 销售报告编号。
	PayerInn string `json:"payer_inn"` // 付款人的纳税人识别号。
	PayerKpp string `json:"payer_kpp"` // 付款人的纳税人登记原因代码。
	PayerName string `json:"payer_name"` // 付款人名称。
}

// 订单信息。
type RowItemOrder struct {
	PostingNumber string `json:"posting_number"` // 货件编号。
	CreatedDate string `json:"created_date"` // 订单日期格式为`YYYY-MM-DD`。
}

// 关于向法人销售的信息。
type RowItemLegalEntityDocument struct {
	Number string `json:"number"` // 发票编号。
	SaleDate string `json:"sale_date"` // 日期格式为`YYYY-MM-DD`。
}

type V1GetRealizationReportPostingResponseRow struct {
	Item RowItem `json:"item"`
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"row_number"` // 报告中的行号。
	SellerPricePerInstance float64 `json:"seller_price_per_instance"` // 考虑折扣后的卖家价格。
	Order RowItemOrder `json:"order"`
	LegalEntityDocument RowItemLegalEntityDocument `json:"legal_entity_document"`
	CommissionRatio float64 `json:"commission_ratio"` // 按类目划分的销售佣金比例。
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
}

// 查询结果。
type V1GetRealizationReportPostingResponse struct {
	Header GetRealizationReportResponseV2Header `json:"header"`
	Rows []V1GetRealizationReportPostingResponseRow `json:"rows"` // 报告表格。
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

type V2GetRealizationReportResponseV2 struct {
	Result GetRealizationReportResponseV2Result `json:"result"`
}

// 日期过滤。
// 阶段开始
type FilterPeriodFromEnum string
const (
	FilterPeriodFromEnum_YYYYMMDDTHHMmSsSssZ FilterPeriodFromEnum = "YYYY-MM-DDTHH:mm:ss.sssZ"
	FilterPeriodFromEnum_V20191125T10430651 FilterPeriodFromEnum = "2019-11-25T10:43:06.51"
)

// 阶段结束
type FilterPeriodToEnum string
const (
	FilterPeriodToEnum_YYYYMMDDTHHMmSsSssZ FilterPeriodToEnum = "YYYY-MM-DDTHH:mm:ss.sssZ"
	FilterPeriodToEnum_V20191125T10430651 FilterPeriodToEnum = "2019-11-25T10:43:06.51"
)

type FilterPeriod struct {
	From FilterPeriodFromEnum `json:"from"` // 阶段开始。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
	To FilterPeriodToEnum `json:"to"` // 阶段结束。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
}

// 过滤器。
// 交易类型:
type FinanceTransactionListV3RequestFilterOperationTypeEnum string
const (
	FinanceTransactionListV3RequestFilterOperationTypeEnum_ClientReturnAgentOperation FinanceTransactionListV3RequestFilterOperationTypeEnum = "ClientReturnAgentOperation"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceMarketingActionCostOperation FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceMarketingActionCostOperation"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceSaleReviewsOperation FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceSaleReviewsOperation"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceSellerCompensationOperation FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceSellerCompensationOperation"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationAgentDeliveredToCustomer FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationAgentDeliveredToCustomer"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationAgentDeliveredToCustomerCanceled FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationAgentDeliveredToCustomerCanceled"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationAgentStornoDeliveredToCustomer FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationAgentStornoDeliveredToCustomer"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationClaim FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationClaim"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationCorrectionSeller FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationCorrectionSeller"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationDefectiveWriteOff FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationDefectiveWriteOff"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationItemReturn FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationItemReturn"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationLackWriteOff FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationLackWriteOff"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationMarketplaceCrossDockServiceWriteOff FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationMarketplaceCrossDockServiceWriteOff"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationMarketplaceServiceStorage FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationMarketplaceServiceStorage"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationSetOff FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationSetOff"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceSellerReexposureDeliveryReturnOperation FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceSellerReexposureDeliveryReturnOperation"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationReturnGoodsFBSofRMS FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationReturnGoodsFBSofRMS"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_ReturnAgentOperationRFBS FinanceTransactionListV3RequestFilterOperationTypeEnum = "ReturnAgentOperationRFBS"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_ItemAgentServiceStarsMembership FinanceTransactionListV3RequestFilterOperationTypeEnum = "ItemAgentServiceStarsMembership"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceSellerShippingCompensationReturnOperation FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceSellerShippingCompensationReturnOperation"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationMarketplaceServicePremiumCashback FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationMarketplaceServicePremiumCashback"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceServicePremiumPromotion FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceServicePremiumPromotion"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceRedistributionOfAcquiringOperation FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceRedistributionOfAcquiringOperation"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceReturnStorageServiceAtThePickupPointFbsItem FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceReturnStorageServiceAtThePickupPointFbsItem"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceReturnStorageServiceInTheWarehouseFbsItem FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceReturnStorageServiceInTheWarehouseFbsItem"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceServiceItemDeliveryKGT FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceServiceItemDeliveryKGT"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceServiceItemDirectFlowLogistic FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceServiceItemDirectFlowLogistic"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceServiceItemReturnFlowLogistic FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceServiceItemReturnFlowLogistic"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceServicePremiumCashbackIndividualPoints FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceServicePremiumCashbackIndividualPoints"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationMarketplaceWithHoldingForUndeliverableGoods FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationMarketplaceWithHoldingForUndeliverableGoods"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_MarketplaceServiceItemRedistributionReturnsPVZ FinanceTransactionListV3RequestFilterOperationTypeEnum = "MarketplaceServiceItemRedistributionReturnsPVZ"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationElectronicServiceStencil FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationElectronicServiceStencil"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationElectronicServicesPromotionInSearch FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationElectronicServicesPromotionInSearch"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationMarketplaceServiceItemElectronicServicesBrandShelf FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationMarketplaceServiceItemElectronicServicesBrandShelf"
	FinanceTransactionListV3RequestFilterOperationTypeEnum_OperationSubscriptionPremium FinanceTransactionListV3RequestFilterOperationTypeEnum = "OperationSubscriptionPremium"
)

// 收费类型：
type FinanceTransactionListV3RequestFilterTransactionTypeEnum string
const (
	FinanceTransactionListV3RequestFilterTransactionTypeEnum_All FinanceTransactionListV3RequestFilterTransactionTypeEnum = "all"
	FinanceTransactionListV3RequestFilterTransactionTypeEnum_Orders FinanceTransactionListV3RequestFilterTransactionTypeEnum = "orders"
	FinanceTransactionListV3RequestFilterTransactionTypeEnum_Returns FinanceTransactionListV3RequestFilterTransactionTypeEnum = "returns"
	FinanceTransactionListV3RequestFilterTransactionTypeEnum_Services FinanceTransactionListV3RequestFilterTransactionTypeEnum = "services"
	FinanceTransactionListV3RequestFilterTransactionTypeEnum_Compensation FinanceTransactionListV3RequestFilterTransactionTypeEnum = "compensation"
	FinanceTransactionListV3RequestFilterTransactionTypeEnum_TransferDelivery FinanceTransactionListV3RequestFilterTransactionTypeEnum = "transferDelivery"
	FinanceTransactionListV3RequestFilterTransactionTypeEnum_Other FinanceTransactionListV3RequestFilterTransactionTypeEnum = "other"
)

type FinanceTransactionListV3RequestFilter struct {
	OperationType []string `json:"operation_type"` // 交易类型: - `ClientReturnAgentOperation` — 收到买家的退货、取消订单、非赎金； - `MarketplaceMarketingActionCostOperation` — 商品促销服务； - `Market...
	PostingNumber string `json:"posting_number"` // 发货号。
	TransactionType FinanceTransactionListV3RequestFilterTransactionTypeEnum `json:"transaction_type"` // 收费类型： - `all` — 所有, - `orders` — 订单, - `returns` — 退货和取消, - `services` — 服务费, - `compensation` — 补贴, - `transferDelivery...
	Date FilterPeriod `json:"date"`
}

// 报告语言： - `RU` — 俄语， - `EN` — 英语。
type CompensationReportLanguage string

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"` // 报告周期格式为 `YYYY-MM`。
	Language CompensationReportLanguage `json:"language"`
}

type Financev3FinanceTransactionListV3Request struct {
	Filter FinanceTransactionListV3RequestFilter `json:"filter"`
	Page int64 `json:"page"` // 请求中返回的页码。
	PageSize int64 `json:"page_size"` // 每页的元素数。
}

// 按日期过滤。
// 阶段开始
type FinanceTransactionTotalsV3RequestDateFromEnum string
const (
	FinanceTransactionTotalsV3RequestDateFromEnum_YYYYMMDDTHHMmSsSssZ FinanceTransactionTotalsV3RequestDateFromEnum = "YYYY-MM-DDTHH:mm:ss.sssZ"
	FinanceTransactionTotalsV3RequestDateFromEnum_V20191125T10430651 FinanceTransactionTotalsV3RequestDateFromEnum = "2019-11-25T10:43:06.51"
)

// 阶段结束
type FinanceTransactionTotalsV3RequestDateToEnum string
const (
	FinanceTransactionTotalsV3RequestDateToEnum_YYYYMMDDTHHMmSsSssZ FinanceTransactionTotalsV3RequestDateToEnum = "YYYY-MM-DDTHH:mm:ss.sssZ"
	FinanceTransactionTotalsV3RequestDateToEnum_V20191125T10430651 FinanceTransactionTotalsV3RequestDateToEnum = "2019-11-25T10:43:06.51"
)

type FinanceTransactionTotalsV3RequestDate struct {
	From FinanceTransactionTotalsV3RequestDateFromEnum `json:"from"` // 阶段开始。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
	To FinanceTransactionTotalsV3RequestDateToEnum `json:"to"` // 阶段结束。 格式: `YYYY-MM-DDTHH:mm:ss.sssZ`.<br> 例子: `2019-11-25T10:43:06.51`.
}

// 询问结果。
type Financev3FinanceTransactionTotalsV3ResponseResult struct {
	ServicesAmount float64 `json:"services_amount"` // 与商品交付和退货没有直接关系的附加服务成本。例如，促销或商品放置。
	AccrualsForSale float64 `json:"accruals_for_sale"` // 指定期间内商品的总成本和退货。
	CompensationAmount float64 `json:"compensation_amount"` // 补贴。
	MoneyTransfer float64 `json:"money_transfer"` // 根据“卖方选择交货”计划工作时的交货和退货费用。
	OthersAmount float64 `json:"others_amount"` // 其他应计费用。
	ProcessingAndDelivery float64 `json:"processing_and_delivery"` // 运输处理、订单装配、干线、最后一英里以及自2021年2月1日起引入新的佣金和费率前的快递服务费。 干线 —— 集群之间的货物交付。 最后一英里 —— 从订单交付点、自提点和快递员到买家处的快递。
	RefundsAndCancellations float64 `json:"refunds_and_cancellations"` // 干线返回、退货处理、取消和非赎回、2021年2月1日起引入新佣金和税率之前退货价格。 干线 —— 集群之间的货物交付。 最后一英里 —— 从订单交付点、自提点和快递员到买家处的快递。
	SaleCommission float64 `json:"sale_commission"` // 商品预售时预扣的佣金数额，退货时返还的佣金数。
}

// 询问结果。
type Financev3FinanceTransactionListV3ResponseResult struct {
	RowCount int64 `json:"row_count"` // 所有页面上的交易数量。如果为0，说明已无交易。
	Operations []FinanceTransactionListV3ResponseOperation `json:"operations"` // 操作信息。
	PageCount int64 `json:"page_count"` // 页数。如果为0，则说明已无页面。
}

type CreateReportResponse struct {
	Result CreateReportResponseCodeNoDeadline `json:"result"`
}

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"` // 月。
	Year int32 `json:"year"` // 年。
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result Financev3FinanceTransactionTotalsV3ResponseResult `json:"result"`
}

// 操作类型：
type Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum string
const (
	Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum_All Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum = "all"
	Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum_Orders Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum = "orders"
	Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum_Returns Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum = "returns"
	Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum_Services Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum = "services"
	Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum_Compensation Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum = "compensation"
	Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum_TransferDelivery Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum = "transferDelivery"
	Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum_Other Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum = "other"
)

type Financev3FinanceTransactionTotalsV3Request struct {
	Date FinanceTransactionTotalsV3RequestDate `json:"date"`
	PostingNumber string `json:"posting_number"` // 发货号。
	TransactionType Financev3FinanceTransactionTotalsV3RequestTransactionTypeEnum `json:"transaction_type"` // 操作类型： - `all` — 所有, - `orders` — 订单, - `returns` — 退货和取消, - `services` — 服务费, - `compensation` — 补贴, - `transferDelivery...
}

type V1GetCompensationReportRequest struct {
	Date string `json:"date"` // 报告周期格式为 `YYYY-MM`。
	Language CompensationReportLanguage `json:"language"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result Financev3FinanceTransactionListV3ResponseResult `json:"result"`
}

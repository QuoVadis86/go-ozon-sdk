package fbs

// 发运的计费信息。
// 新的费率开始生效的日期和时间
type V3FbsTarifficationNextTariffStartsAtEnum string
const (
	V3FbsTarifficationNextTariffStartsAtEnum_YYYYMMDDThhMmSsMcsZ V3FbsTarifficationNextTariffStartsAtEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V3FbsTarifficationNextTariffStartsAtEnum_V20231113T080557657Z V3FbsTarifficationNextTariffStartsAtEnum = "2023-11-13T08:05:57.657Z"
)

type V3FbsTariffication struct {
	CurrentTariffCharge string `json:"current_tariff_charge"` // 当前的折扣或附加费金额。
	CurrentTariffChargeCurrencyCode string `json:"current_tariff_charge_currency_code"` // 金额的货币单位。
	NextTariffRate float64 `json:"next_tariff_rate"` // 在参数 `next_tariff_starts_at` 指定的时间后，将按此百分比进行计费。
	NextTariffChargeCurrencyCode string `json:"next_tariff_charge_currency_code"` // 新费率的货币单位。
	CurrentTariffRate float64 `json:"current_tariff_rate"` // 当前运费的百分比。
	CurrentTariffType string `json:"current_tariff_type"` // 当前的计费类型 — 折扣或附加费。
	NextTariffType string `json:"next_tariff_type"` // 在参数 `next_tariff_starts_at` 指定的时间后，将按此类型计费 — 折扣或附加费。
	NextTariffCharge string `json:"next_tariff_charge"` // 下一步计费中的折扣或附加金额。
	NextTariffStartsAt V3FbsTarifficationNextTariffStartsAtEnum `json:"next_tariff_starts_at"` // 新的费率开始生效的日期和时间。 格式：`YYYY-MM-DDThh:mm:ss.mcsZ`. 示例：`2023-11-13T08:05:57.657Z`.
}

// 价格货币，其与个人中心中设置的币种相匹配
type PostingFinancialDataProductCurrencyCodeEnum string
const (
	PostingFinancialDataProductCurrencyCodeEnum_RUB PostingFinancialDataProductCurrencyCodeEnum = "RUB"
	PostingFinancialDataProductCurrencyCodeEnum_BYN PostingFinancialDataProductCurrencyCodeEnum = "BYN"
	PostingFinancialDataProductCurrencyCodeEnum_KZT PostingFinancialDataProductCurrencyCodeEnum = "KZT"
	PostingFinancialDataProductCurrencyCodeEnum_EUR PostingFinancialDataProductCurrencyCodeEnum = "EUR"
	PostingFinancialDataProductCurrencyCodeEnum_USD PostingFinancialDataProductCurrencyCodeEnum = "USD"
	PostingFinancialDataProductCurrencyCodeEnum_CNY PostingFinancialDataProductCurrencyCodeEnum = "CNY"
)

type PostingFinancialDataProduct struct {
	TotalDiscountPercent float64 `json:"total_discount_percent"` // 折扣百分比。
	TotalDiscountValue float64 `json:"total_discount_value"` // 折扣数量。
	Actions []string `json:"actions"` // 活动清单。
	CurrencyCode PostingFinancialDataProductCurrencyCodeEnum `json:"currency_code"` // 价格货币，其与个人中心中设置的币种相匹配。 可能的值： - `RUB` — 俄罗斯卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
	CustomerCurrencyCode string `json:"customer_currency_code"` // 买家货币代码。
	CommissionPercent int64 `json:"commission_percent"` // 佣金百分比。
	CommissionsCurrencyCode string `json:"commissions_currency_code"` // 计算佣金的币种代码。
	OldPrice float64 `json:"old_price"` // 打折前价格。在商品卡片上将被显示划掉。
	CustomerPrice float64 `json:"customer_price"` // 包含卖家与 Ozon 折扣的买家价格。
	ProductID int64 `json:"product_id"` // Ozon系统中的商品标识符——SKU。
	CommissionAmount float64 `json:"commission_amount"` // 商品佣金大小。
	Payout float64 `json:"payout"` // 支付给卖方。
	Price float64 `json:"price"` // 您的价格。包含卖家促销（如有），不含 Ozon 资助的促销。
	Quantity int64 `json:"quantity"` // 运输商品数量。
}

// 有关商品成本、折扣幅度、付款和佣金的信息。
type V3PostingFinancialData struct {
	ClusterFrom string `json:"cluster_from"` // 订单发送区域代码。
	ClusterTo string `json:"cluster_to"` // 订单接受区域代码。
	Products []PostingFinancialDataProduct `json:"products"` // 订单中的商品列表。
}

// 收件人联系方式。
type V3AddresseeFbsLists struct {
	Name string `json:"name"` // 收件人姓名。
	Phone string `json:"phone"` // 收件人联系电话。总是返回空字符串 `""`。要获取替代联系电话，请使用方法 [/v3/posting/fbs/get](#operation/PostingAPI_GetFbsPostingV3)。
}

// 取消原因。
// 取消货运的发起者：
type V3CancellationCancellationInitiatorEnum string
const (
	V3CancellationCancellationInitiatorEnum_V V3CancellationCancellationInitiatorEnum = "卖家"
	V3CancellationCancellationInitiatorEnum_V_1 V3CancellationCancellationInitiatorEnum = "客户"
	V3CancellationCancellationInitiatorEnum_V_2 V3CancellationCancellationInitiatorEnum = "买家"
	V3CancellationCancellationInitiatorEnum_Ozon V3CancellationCancellationInitiatorEnum = "Ozon"
	V3CancellationCancellationInitiatorEnum_V_3 V3CancellationCancellationInitiatorEnum = "系统"
	V3CancellationCancellationInitiatorEnum_V_4 V3CancellationCancellationInitiatorEnum = "配送服务"
)

// 货运取消类型：
type V3CancellationCancellationTypeEnum string
const (
	V3CancellationCancellationTypeEnum_Seller V3CancellationCancellationTypeEnum = "seller"
	V3CancellationCancellationTypeEnum_Client V3CancellationCancellationTypeEnum = "client"
	V3CancellationCancellationTypeEnum_Customer V3CancellationCancellationTypeEnum = "customer"
	V3CancellationCancellationTypeEnum_Ozon V3CancellationCancellationTypeEnum = "ozon"
	V3CancellationCancellationTypeEnum_System V3CancellationCancellationTypeEnum = "system"
	V3CancellationCancellationTypeEnum_Delivery V3CancellationCancellationTypeEnum = "delivery"
)

type V3Cancellation struct {
	CancellationInitiator V3CancellationCancellationInitiatorEnum `json:"cancellation_initiator"` // 取消货运的发起者： - `卖家`, - `客户` 或`买家`, - `Ozon`, - `系统`, - `配送服务`。
	CancellationType V3CancellationCancellationTypeEnum `json:"cancellation_type"` // 货运取消类型： - `seller` — 卖家取消； - `client` 或 `customer` — 买家取消； - `ozon` — Ozon取消； - `system`— 系统取消； - `delivery` — 配送服务取消。
	CancelledAfterShip bool `json:"cancelled_after_ship"` // 如果订单在装运后取消 — `true`。
	AffectCancellationRating bool `json:"affect_cancellation_rating"` // 如果取消影响买家排行 — `true`。
	CancelReason string `json:"cancel_reason"` // 取消原因。
	CancelReasonID int64 `json:"cancel_reason_id"` // 取消货运的原因ID。
}

// 货件条码。
type V3Barcodes struct {
	LowerBarcode string `json:"lower_barcode"` // 货件标签的下条码。
	UpperBarcode string `json:"upper_barcode"` // 货件标签的上条码。
}

// 分析数据。
// 付款方法：
type V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum string
const (
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_V V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "在线银行卡支付"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_Ozon V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "Ozon银行卡"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_Ozon_1 V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "取货时自动从Ozon银行卡收费"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_V_1 V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "收货时从已保存的银行卡收费"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_V_2 V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "快速支付系统"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_Ozon_2 V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "Ozon分期付款"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_V_3 V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "支付至结算账户"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_SberPay V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "SberPay"
	V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum_V_4 V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum = "外部卖家方的预付款"
)

type V3FbsPostingAnalyticsData struct {
	PaymentTypeGroupName V3FbsPostingAnalyticsDataPaymentTypeGroupNameEnum `json:"payment_type_group_name"` // 付款方法： - `在线银行卡支付`， - `Ozon银行卡`， - `取货时自动从Ozon银行卡收费`， - `收货时从已保存的银行卡收费`， - `快速支付系统`， - `Ozon分期付款`， - `支付至结算账户`， - `SberPa...
	Warehouse string `json:"warehouse"` // 订单发送仓库名称。
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"` // 向客户开始配送的日期和时间。仅适用于通过Ozon配送创建的货件。
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"` // 订单将送达客户的截止日期。仅适用于通过Ozon配送创建的货件。
	City string `json:"city"` // 快递城市。仅适用于rFBS货件和独联体卖家。
	DeliveryDateBegin string `json:"delivery_date_begin"` // 快递开始日期和时间。
	DeliveryType string `json:"delivery_type"` // 快递方式。
	IsPremium bool `json:"is_premium"` // 有Premium订阅。
	Region string `json:"region"` // 快递地区。
	TPLProvider string `json:"tpl_provider"` // 快递服务。
	TPLProviderID int64 `json:"tpl_provider_id"` // 快递服务ID。
	WarehouseID int64 `json:"warehouse_id"` // 仓库ID。
	DeliveryDateEnd string `json:"delivery_date_end"` // 快递结束日期和时间。
	IsLegal bool `json:"is_legal"` // 收件人是法人的标志： - `true` — 法人， - `false` — 自然人。
}

// 带有附加特征的商品列表。
type V3FbsPostingDetailOptional struct {
	ProductsWithPossibleMandatoryMark []any `json:"products_with_possible_mandatory_mark"` // 带有可能标志的商品列表。
}

// 快递地址信息。
type V3Address struct {
	City string `json:"city"` // 快递城市。
	Comment string `json:"comment"` // 订单评价。
	District string `json:"district"` // 快递地区。
	Latitude float64 `json:"latitude"` // 宽。
	ProviderPvzCode string `json:"provider_pvz_code"` // 3PL提供商的订单提货点的代码。
	PvzCode int64 `json:"pvz_code"` // 订单取货点代码。
	Region string `json:"region"` // 快递区域。
	ZipCode string `json:"zip_code"` // 收件人邮编。
	AddressTail string `json:"address_tail"` // 文本格式的地址。
	Country string `json:"country"` // 快递国家。
	Longitude float64 `json:"longitude"` // （时间的）长度。
}

// 买家信息。
type V3CustomerFbsLists struct {
	Name string `json:"name"` // 买家姓名。
	Phone string `json:"phone"` // 买家联系电话。始终返回空字符串 `""`。要获取替代联系电话，请使用方法 [/v3/posting/fbs/get](#operation/PostingAPI_GetFbsPostingV3)。
	Address V3Address `json:"address"`
	CustomerEmail string `json:"customer_email"` // 买家的电子邮箱地址。
	CustomerID int64 `json:"customer_id"` // 买家ID。
}

// 价格货币，其与个人中心中设置的币种相匹配
type V3FbsPostingProductCurrencyCodeEnum string
const (
	V3FbsPostingProductCurrencyCodeEnum_RUB V3FbsPostingProductCurrencyCodeEnum = "RUB"
	V3FbsPostingProductCurrencyCodeEnum_BYN V3FbsPostingProductCurrencyCodeEnum = "BYN"
	V3FbsPostingProductCurrencyCodeEnum_KZT V3FbsPostingProductCurrencyCodeEnum = "KZT"
	V3FbsPostingProductCurrencyCodeEnum_EUR V3FbsPostingProductCurrencyCodeEnum = "EUR"
	V3FbsPostingProductCurrencyCodeEnum_USD V3FbsPostingProductCurrencyCodeEnum = "USD"
	V3FbsPostingProductCurrencyCodeEnum_CNY V3FbsPostingProductCurrencyCodeEnum = "CNY"
)

type V3FbsPostingProduct struct {
	SKU int64 `json:"sku"` // 在Ozon系统中的商品ID — SKU。
	CurrencyCode V3FbsPostingProductCurrencyCodeEnum `json:"currency_code"` // 价格货币，其与个人中心中设置的币种相匹配。 可能的值： - `RUB` — 俄罗斯卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
	Imei []string `json:"imei"` // 移动设备的 IMEI 列表。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 在卖家系统中的商品ID — 货号。
	Price string `json:"price"` // 商品价格。
	Quantity int32 `json:"quantity"` // 运输中的商品数量。
}

// 买方的法律信息。
type V2FboSinglePostingLegalInfo struct {
	Kpp string `json:"kpp"` // 税务登记原因代码（KPP）。
	CompanyName string `json:"company_name"` // 公司名称。
	Inn string `json:"inn"` // 纳税人识别号（INN）。
}

// 最低折扣或附加费用。
type MoneyMoneyCurrentTariffMinCharge struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 折扣或附加费用。
type MoneyMoneyCurrentTariffCharge struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

type PostingV4PostingFbsListResponsePostingsTarifficationStep struct {
	MinCharge MoneyMoneyCurrentTariffMinCharge `json:"min_charge"`
	TariffCharge MoneyMoneyCurrentTariffCharge `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"` // 计费阶段结束的日期和时间。该日期后将自动进入下一计费阶段。
	TariffRate float64 `json:"tariff_rate"` // 折扣或附加费用百分比。
	TariffType string `json:"tariff_type"` // 计费类型。
}

// 快递方式。
type V3DeliveryMethod struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库ID。
	ID int64 `json:"id"` // 快递方式ID。
	Name string `json:"name"` // 快递方式名称。
	TPLProvider string `json:"tpl_provider"` // 快递服务。
	TPLProviderID int64 `json:"tpl_provider_id"` // 快递服务ID。
	Warehouse string `json:"warehouse"` // 仓库名称。
}

// 需提供制造国、货运报关单号、商品批次登记号、“诚实标志”、其他标识或重量的商品列表，以便将货件状态更新至下一阶段。
type V3FbsPostingRequirementsV3 struct {
	ProductsRequiringGTD []string `json:"products_requiring_gtd"` // 必须提供货运报关单号（CCD）的商品 ID（SKU）列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProduct...
	ProductsRequiringCountry []string `json:"products_requiring_country"` // 需要上传制造国信息的商品ID列表 (SKU)。 要配货，请上传上述商品的制造国信息，通过方法 [/v2/posting/fbs/product/country/set](#operation/PostingAPI_SetCountryPro...
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"` // 需要提供“诚实标志”标签的商品 ID（SKU）列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExe...
	ProductsRequiringJwUin []string `json:"products_requiring_jw_uin"` // 需要提供首饰唯一识别号（UIN）的商品列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExempla...
	ProductsRequiringRnpt []string `json:"products_requiring_rnpt"` // 需要提供商品批次注册号（RNPT）的商品 ID（SKU）列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProdu...
	ProductsRequiringImei []string `json:"products_requiring_imei"` // 需要提供 IMEI 的商品 ID 列表。
	ProductsRequiringChangeCountry []string `json:"products_requiring_change_country"` // 需要修改生产国家的商品（SKU）编号列表。要修改生产国家，请使用方法 [/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountryProductFbsPos...
}

// 可用的操作和货件信息包括：
type V3FbsPostingAvailableActionsEnum string
const (
	V3FbsPostingAvailableActionsEnum_Arbitration V3FbsPostingAvailableActionsEnum = "arbitration"
	V3FbsPostingAvailableActionsEnum_AwaitingDelivery V3FbsPostingAvailableActionsEnum = "awaiting_delivery"
	V3FbsPostingAvailableActionsEnum_CanCreateChat V3FbsPostingAvailableActionsEnum = "can_create_chat"
	V3FbsPostingAvailableActionsEnum_Cancel V3FbsPostingAvailableActionsEnum = "cancel"
	V3FbsPostingAvailableActionsEnum_ClickTrackNumber V3FbsPostingAvailableActionsEnum = "click_track_number"
	V3FbsPostingAvailableActionsEnum_CustomerPhoneAvailable V3FbsPostingAvailableActionsEnum = "customer_phone_available"
	V3FbsPostingAvailableActionsEnum_HasWeightProducts V3FbsPostingAvailableActionsEnum = "has_weight_products"
	V3FbsPostingAvailableActionsEnum_HideRegionAndCity V3FbsPostingAvailableActionsEnum = "hide_region_and_city"
	V3FbsPostingAvailableActionsEnum_InvoiceGet V3FbsPostingAvailableActionsEnum = "invoice_get"
	V3FbsPostingAvailableActionsEnum_InvoiceSend V3FbsPostingAvailableActionsEnum = "invoice_send"
	V3FbsPostingAvailableActionsEnum_InvoiceUpdate V3FbsPostingAvailableActionsEnum = "invoice_update"
	V3FbsPostingAvailableActionsEnum_LabelDownloadBig V3FbsPostingAvailableActionsEnum = "label_download_big"
	V3FbsPostingAvailableActionsEnum_LabelDownloadSmall V3FbsPostingAvailableActionsEnum = "label_download_small"
	V3FbsPostingAvailableActionsEnum_LabelDownload V3FbsPostingAvailableActionsEnum = "label_download"
	V3FbsPostingAvailableActionsEnum_NonIntDelivered V3FbsPostingAvailableActionsEnum = "non_int_delivered"
	V3FbsPostingAvailableActionsEnum_NonIntDelivering V3FbsPostingAvailableActionsEnum = "non_int_delivering"
	V3FbsPostingAvailableActionsEnum_NonIntLastMile V3FbsPostingAvailableActionsEnum = "non_int_last_mile"
	V3FbsPostingAvailableActionsEnum_ProductCancel V3FbsPostingAvailableActionsEnum = "product_cancel"
	V3FbsPostingAvailableActionsEnum_SetCutoff V3FbsPostingAvailableActionsEnum = "set_cutoff"
	V3FbsPostingAvailableActionsEnum_SetTimeslot V3FbsPostingAvailableActionsEnum = "set_timeslot"
	V3FbsPostingAvailableActionsEnum_SetTrackNumber V3FbsPostingAvailableActionsEnum = "set_track_number"
	V3FbsPostingAvailableActionsEnum_ShipAsyncInProcess V3FbsPostingAvailableActionsEnum = "ship_async_in_process"
	V3FbsPostingAvailableActionsEnum_ShipAsyncRetry V3FbsPostingAvailableActionsEnum = "ship_async_retry"
	V3FbsPostingAvailableActionsEnum_ShipAsync V3FbsPostingAvailableActionsEnum = "ship_async"
	V3FbsPostingAvailableActionsEnum_ShipWithAdditionalInfo V3FbsPostingAvailableActionsEnum = "ship_with_additional_info"
	V3FbsPostingAvailableActionsEnum_Ship V3FbsPostingAvailableActionsEnum = "ship"
	V3FbsPostingAvailableActionsEnum_UpdateCis V3FbsPostingAvailableActionsEnum = "update_cis"
)

// 快递服务集成类型：
type V3FbsPostingTPLIntegrationTypeEnum string
const (
	V3FbsPostingTPLIntegrationTypeEnum_Ozon V3FbsPostingTPLIntegrationTypeEnum = "ozon"
	V3FbsPostingTPLIntegrationTypeEnum_V3plTracking V3FbsPostingTPLIntegrationTypeEnum = "3pl_tracking"
	V3FbsPostingTPLIntegrationTypeEnum_NonIntegrated V3FbsPostingTPLIntegrationTypeEnum = "non_integrated"
	V3FbsPostingTPLIntegrationTypeEnum_Aggregator V3FbsPostingTPLIntegrationTypeEnum = "aggregator"
	V3FbsPostingTPLIntegrationTypeEnum_Hybryd V3FbsPostingTPLIntegrationTypeEnum = "hybryd"
)

// 货运状态:
type V3FbsPostingStatusEnum string
const (
	V3FbsPostingStatusEnum_AcceptanceInProgress V3FbsPostingStatusEnum = "acceptance_in_progress"
	V3FbsPostingStatusEnum_Arbitration V3FbsPostingStatusEnum = "arbitration"
	V3FbsPostingStatusEnum_AwaitingApprove V3FbsPostingStatusEnum = "awaiting_approve"
	V3FbsPostingStatusEnum_AwaitingDeliver V3FbsPostingStatusEnum = "awaiting_deliver"
	V3FbsPostingStatusEnum_AwaitingPackaging V3FbsPostingStatusEnum = "awaiting_packaging"
	V3FbsPostingStatusEnum_AwaitingRegistration V3FbsPostingStatusEnum = "awaiting_registration"
	V3FbsPostingStatusEnum_AwaitingVerification V3FbsPostingStatusEnum = "awaiting_verification"
	V3FbsPostingStatusEnum_Cancelled V3FbsPostingStatusEnum = "cancelled"
	V3FbsPostingStatusEnum_CancelledFromSplitPending V3FbsPostingStatusEnum = "cancelled_from_split_pending"
	V3FbsPostingStatusEnum_ClientArbitration V3FbsPostingStatusEnum = "client_arbitration"
	V3FbsPostingStatusEnum_Delivering V3FbsPostingStatusEnum = "delivering"
	V3FbsPostingStatusEnum_DriverPickup V3FbsPostingStatusEnum = "driver_pickup"
	V3FbsPostingStatusEnum_NotAccepted V3FbsPostingStatusEnum = "not_accepted"
)

// 发货子状态：
type V3FbsPostingSubstatusEnum string
const (
	V3FbsPostingSubstatusEnum_PostingAcceptanceInProgress V3FbsPostingSubstatusEnum = "posting_acceptance_in_progress"
	V3FbsPostingSubstatusEnum_PostingInArbitration V3FbsPostingSubstatusEnum = "posting_in_arbitration"
	V3FbsPostingSubstatusEnum_PostingCreated V3FbsPostingSubstatusEnum = "posting_created"
	V3FbsPostingSubstatusEnum_PostingInCarriage V3FbsPostingSubstatusEnum = "posting_in_carriage"
	V3FbsPostingSubstatusEnum_PostingNotInCarriage V3FbsPostingSubstatusEnum = "posting_not_in_carriage"
	V3FbsPostingSubstatusEnum_PostingRegistered V3FbsPostingSubstatusEnum = "posting_registered"
	V3FbsPostingSubstatusEnum_PostingTransferringToDelivery V3FbsPostingSubstatusEnum = "posting_transferring_to_delivery"
	V3FbsPostingSubstatusEnum_StatusAwaitingDeliver V3FbsPostingSubstatusEnum = "status=awaiting_deliver"
	V3FbsPostingSubstatusEnum_PostingAwaitingPassportData V3FbsPostingSubstatusEnum = "posting_awaiting_passport_data"
	V3FbsPostingSubstatusEnum_PostingAwaitingRegistration V3FbsPostingSubstatusEnum = "posting_awaiting_registration"
	V3FbsPostingSubstatusEnum_PostingRegistrationError V3FbsPostingSubstatusEnum = "posting_registration_error"
	V3FbsPostingSubstatusEnum_StatusAwaitingRegistration V3FbsPostingSubstatusEnum = "status=awaiting_registration"
	V3FbsPostingSubstatusEnum_PostingSplitPending V3FbsPostingSubstatusEnum = "posting_split_pending"
	V3FbsPostingSubstatusEnum_PostingCanceled V3FbsPostingSubstatusEnum = "posting_canceled"
	V3FbsPostingSubstatusEnum_PostingInClientArbitration V3FbsPostingSubstatusEnum = "posting_in_client_arbitration"
	V3FbsPostingSubstatusEnum_PostingDelivered V3FbsPostingSubstatusEnum = "posting_delivered"
	V3FbsPostingSubstatusEnum_PostingReceived V3FbsPostingSubstatusEnum = "posting_received"
	V3FbsPostingSubstatusEnum_PostingConditionallyDelivered V3FbsPostingSubstatusEnum = "posting_conditionally_delivered"
	V3FbsPostingSubstatusEnum_PostingInCourierService V3FbsPostingSubstatusEnum = "posting_in_courier_service"
	V3FbsPostingSubstatusEnum_PostingInPickupPoint V3FbsPostingSubstatusEnum = "posting_in_pickup_point"
	V3FbsPostingSubstatusEnum_PostingOnWayToCity V3FbsPostingSubstatusEnum = "posting_on_way_to_city"
	V3FbsPostingSubstatusEnum_PostingOnWayToPickupPoint V3FbsPostingSubstatusEnum = "posting_on_way_to_pickup_point"
	V3FbsPostingSubstatusEnum_PostingReturnedToWarehouse V3FbsPostingSubstatusEnum = "posting_returned_to_warehouse"
	V3FbsPostingSubstatusEnum_PostingTransferredToCourierService V3FbsPostingSubstatusEnum = "posting_transferred_to_courier_service"
	V3FbsPostingSubstatusEnum_PostingDriverPickUp V3FbsPostingSubstatusEnum = "posting_driver_pick_up"
	V3FbsPostingSubstatusEnum_PostingNotInSortCenter V3FbsPostingSubstatusEnum = "posting_not_in_sort_center"
	V3FbsPostingSubstatusEnum_ShipFailed V3FbsPostingSubstatusEnum = "ship_failed"
)

type V3FbsPosting struct {
	AvailableActions []string `json:"available_actions"` // 可用的操作和货件信息包括： - `arbitration` — 提出争议； - `awaiting_delivery` — 转为“等待发运”状态； - `can_create_chat` — 与买家开启聊天； - `cancel` — 取消...
	Barcodes V3Barcodes `json:"barcodes"`
	DeliveringDate string `json:"delivering_date"` // 货件交付物流的时间。
	Tariffication V3FbsTariffication `json:"tariffication"`
	AnalyticsData V3FbsPostingAnalyticsData `json:"analytics_data"`
	FinancialData V3PostingFinancialData `json:"financial_data"`
	Optional V3FbsPostingDetailOptional `json:"optional"`
	ParentPostingNumber string `json:"parent_posting_number"` // 快递母件编号，从该母件中拆分出了当前货件。
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"` // 不逾期发运日期和时间。
	TPLIntegrationType V3FbsPostingTPLIntegrationTypeEnum `json:"tpl_integration_type"` // 快递服务集成类型： - `ozon` —— Ozon 快递服务。 - `3pl_tracking` —— 集成服务快递。 - `non_integrated` —— 第三方物流服务。 - `aggregator` —— 通过Ozon合作物流...
	Customer V3CustomerFbsLists `json:"customer"`
	ShipmentDate string `json:"shipment_date"` // 必须收取货件的日期和时间。 超出该时间后将适用新费率，相关信息请查看字段 `tariffication`。
	IsExpress bool `json:"is_express"` // 如果使用快速物流 Ozon Express —— `true`。
	IsPresortable bool `json:"is_presortable"` // 如果发运是预先分拣的，则为`true`。
	OrderID int64 `json:"order_id"` // 货件所属订单的ID。
	Status V3FbsPostingStatusEnum `json:"status"` // 货运状态: - `acceptance_in_progress` —— 正在验收， - `arbitration` —— 仲裁， - `awaiting_approve` —— 等待确认， - `awaiting_deliver` —— 等...
	Addressee V3AddresseeFbsLists `json:"addressee"`
	DestinationPlaceName string `json:"destination_place_name"` // 目的仓库的名称。
	Products []V3FbsPostingProduct `json:"products"` // 货运商品列表。
	DestinationPlaceID int64 `json:"destination_place_id"` // 目的仓库的标识符。
	Substatus V3FbsPostingSubstatusEnum `json:"substatus"` // 发货子状态： - `posting_acceptance_in_progress` —— 正在验收， - `posting_in_arbitration` —— 仲裁， - `posting_created` —— 已创建， - `post...
	Cancellation V3Cancellation `json:"cancellation"`
	InProcessAt string `json:"in_process_at"` // 开始处理货件的日期和时间。
	LegalInfo V2FboSinglePostingLegalInfo `json:"legal_info"`
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号。
	TarifficationSteps []PostingV4PostingFbsListResponsePostingsTarifficationStep `json:"tariffication_steps"` // 计费阶段。
	DeliveryMethod V3DeliveryMethod `json:"delivery_method"`
	OrderNumber string `json:"order_number"` // 货件所属的订单号。
	PostingNumber string `json:"posting_number"` // 货件号。
	Requirements V3FbsPostingRequirementsV3 `json:"requirements"`
}

// 取消货件ID：
type PostingCancelReasonTypeIDEnum string
const (
	PostingCancelReasonTypeIDEnum_Buyer PostingCancelReasonTypeIDEnum = "buyer"
	PostingCancelReasonTypeIDEnum_Seller PostingCancelReasonTypeIDEnum = "seller"
)

type PostingCancelReason struct {
	Title string `json:"title"` // 类别名称。
	TypeID PostingCancelReasonTypeIDEnum `json:"type_id"` // 取消货件ID： - `buyer` — 买家， - `seller` — 卖家。
	ID int64 `json:"id"` // 取消原因ID。
	IsAvailableForCancellation bool `json:"is_available_for_cancellation"` // 取消装运结果。 `true`, 如果请求可以取消。
}

type V1AssemblyFbsPostingListResponsePostingProduct struct {
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符——SKU。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符——货号。
	PictureURL string `json:"picture_url"` // 商品图片链接。
	ProductName string `json:"product_name"` // 商品名称。
	Quantity int32 `json:"quantity"` // 商品数量。
}

type V1AssemblyFbsPostingListResponsePosting struct {
	Products []V1AssemblyFbsPostingListResponsePostingProduct `json:"products"` // 商品列表。
	AssemblyCode string `json:"assembly_code"` // 拣货单代码。
	PostingNumber string `json:"posting_number"` // 货件编号。
}

type PostingProductCancelRequestItem struct {
	Quantity int32 `json:"quantity"` // 货运商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品ID — SKU。
}

type PostingPostingProductCancelRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"` // 货物取消发货原因ID。
	CancelReasonMessage string `json:"cancel_reason_message"` // 必填字段。关于取消的其他信息。
	Items []PostingProductCancelRequestItem `json:"items"` // 商品信息。
	PostingNumber string `json:"posting_number"` // 货运ID。
}

type FbsPostingMoveStatusResponseMoveStatus struct {
	PostingNumber string `json:"posting_number"` // 发货号。
	Result bool `json:"result"` // 如果执行请求无误 — `true`。
	Error string `json:"error"` // 处理请求时出错。
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []FbsPostingMoveStatusResponseMoveStatus `json:"result"` // 方法操作结果。
}

// 排序方向： - `ASC`——升序， - `DESC`——降序。
type V1AssemblyFbsProductListRequestSortDirEnum string

type Postingv1GetCarriageAvailableListRequest struct {
	DeliveryMethodID int64 `json:"delivery_method_id"` // 按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	DepartureDate string `json:"departure_date"` // 装运日期。默认 —— 当前日期。
}

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"` // 货运号。
}

// 要添加到响应的附加字段。
type Postingv3FbsPostingWithParams struct {
	AnalyticsData bool `json:"analytics_data"` // 将分析数据添加到响应中。
	Barcodes bool `json:"barcodes"` // 将货件条形码添加到响应中。
	FinancialData bool `json:"financial_data"` // 将财务数据添加到响应中。
	LegalInfo bool `json:"legal_info"` // 将法律信息添加到响应中。
	Translit bool `json:"translit"` // 完成返回值的拼写。
}

type V3FbsPostingProductExemplarInfoV3 struct {
	IsRnptAbsent bool `json:"is_rnpt_absent"` // 未指出商品批次注册号（Product Batch Registration Number）的标志。
	Imei []string `json:"imei"` // 移动设备的 IMEI 列表。
	ExemplarID int64 `json:"exemplar_id"` // 样件识别码。
	MandatoryMark string `json:"mandatory_mark"` // 强制性标记“诚实标记”。
	GTD string `json:"gtd"` // 货运报关单号（Cargo Customs Declaration）。
	IsGTDAbsent bool `json:"is_gtd_absent"` // 未指出货运报关单号（Cargo Customs Declaration）的标志。
	Rnpt string `json:"rnpt"` // 商品批次注册号（Product Batch Registration Number）。
}

type V4FbsPostingShipPackageV4RequestProduct struct {
	ExemplarsIds []string `json:"exemplarsIds"` // 商品外部识别码。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	Quantity int32 `json:"quantity"` // 样件数量。
}

type V4FbsPostingShipPackageV4Request struct {
	Products []V4FbsPostingShipPackageV4RequestProduct `json:"products"` // 商品清单。
	PostingNumber string `json:"posting_number"` // 发货号。
}

// 商品价格。
type MoneyPostingMoney struct {
	Currency string `json:"currency"` // 货币单位。
	Amount string `json:"amount"` // 金额。
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsProducts struct {
	Weight float64 `json:"weight"` // 商品含包装重量。
	Imei []string `json:"imei"` // 移动设备的 IMEI 列表。
	IsBlrTraceable bool `json:"is_blr_traceable"` // `true`，表示商品可追溯。
	Name string `json:"name"` // 商品名称。
	ProductColor string `json:"product_color"` // 商品颜色。
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"` // `true`，表示Ozon已购买该商品，用于配送至欧亚经济联盟及其他国家。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符，即货号。
	Price MoneyPostingMoney `json:"price"`
	Quantity int32 `json:"quantity"` // 货件中的商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符，即SKU。
}

// 外部平台订单信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"` // `true`，表示订单来自外部平台。
	PlatformName string `json:"platform_name"` // 下单平台名称。
}

// 价格币种
type Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum string
const (
	Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum_RUB Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum = "RUB"
	Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum_BYN Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum = "BYN"
	Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum_KZT Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum = "KZT"
	Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum_EUR Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum = "EUR"
	Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum_USD Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum = "USD"
	Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum_CNY Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum = "CNY"
)

type Fbsv4PostingProductDetailWithoutDimensions struct {
	MandatoryMark []string `json:"mandatory_mark"` // 强制标记“诚信标志”。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品ID — 货号。
	Price string `json:"price"` // 价格。
	Quantity int32 `json:"quantity"` // 发货商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品ID — SKU.
	CurrencyCode Fbsv4PostingProductDetailWithoutDimensionsCurrencyCodeEnum `json:"currency_code"` // 价格币种。这与您在个人中心中设置的币种 一 致。 可能的值： - `RUB` — 卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
}

type FbsPostingShipV4ResponseShipAdditionalData struct {
	PostingNumber string `json:"posting_number"` // 发货号。
	Products []Fbsv4PostingProductDetailWithoutDimensions `json:"products"` // 正在运输途中的商品列表。
}

type V2FbsPostingProduct struct {
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品ID — 货号。
	Price string `json:"price"` // 商品价格。
	Quantity int64 `json:"quantity"` // 货运商品数量。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
}

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"` // 备货后生成的货件号码。
}

// 货件状态最后一次发生变更的时间段。
type V3TimeRange struct {
	From string `json:"from"` // 时间段的开始日期。
	To string `json:"to"` // 时间段的结束日期。
}

// 请求过滤。 请按装运时间使用过滤器 — `cutoff`, 或者按照货件交付给快递的时间 — `delivering_date`。 如果一起使用，则会在响应中返回错误。 要按装运时间使用过滤器，请填 `cutoff_from` 和 `cut...
// 货件运输状态：
type Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum string
const (
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_AcceptanceInProgress Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "acceptance_in_progress"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_AwaitingApprove Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "awaiting_approve"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_AwaitingPackaging Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "awaiting_packaging"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_AwaitingRegistration Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "awaiting_registration"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_AwaitingDeliver Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "awaiting_deliver"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_Arbitration Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "arbitration"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_ClientArbitration Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "client_arbitration"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_Delivering Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "delivering"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_DriverPickup Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "driver_pickup"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum_NotAccepted Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum = "not_accepted"
)

// 从合作伙伴仓库（FBP）发货时的货件筛选器：
type Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum string
const (
	Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum_ALL Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum = "ALL"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum_ONLY Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum = "ONLY"
	Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum_WITHOUT Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum = "WITHOUT"
)

type Postingv3GetFbsPostingUnfulfilledListRequestFilter struct {
	CutoffFrom string `json:"cutoff_from"` // 按卖家需要收订单的时间进行筛选。 时间段开始。 格式： YYYY-MM-DDThh:mm:ss.mcsZ. 例子： 2020-03-18T07:34:50.359Z.
	CutoffTo string `json:"cutoff_to"` // 按卖家需要收订单的时间进行筛选。 时间段结束。 格式： YYYY-MM-DDThh:mm:ss.mcsZ. 例子： 2020-03-18T07:34:50.359Z.
	DeliveringDateFrom string `json:"delivering_date_from"` // 将货件交给物流的最快日期。
	DeliveryMethodID []int64 `json:"delivery_method_id"` // 快递方式ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	LastChangedStatusDate V3TimeRange `json:"last_changed_status_date"`
	ProviderID []int64 `json:"provider_id"` // 快递服务ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	Status Postingv3GetFbsPostingUnfulfilledListRequestFilterStatusEnum `json:"status"` // 货件运输状态： - `acceptance_in_progress` — 正在验收， - `awaiting_approve` — 等待确认， - `awaiting_packaging` — 等待包装， - `awaiting_regis...
	WarehouseID []int64 `json:"warehouse_id"` // 仓库ID。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	DeliveringDateTo string `json:"delivering_date_to"` // 将货件交给物流的最迟日期。
	FbpFilter Postingv3GetFbsPostingUnfulfilledListRequestFilterFbpFilterEnum `json:"fbpFilter"` // 从合作伙伴仓库（FBP）发货时的货件筛选器： - `ALL` — 响应中将显示所有符合其他筛选器条件的货件； - `ONLY` — 仅显示FBP货件； - `WITHOUT` — 显示除FBP外的所有货件。 默认值为 `ALL`。
}

// 取消原因ID：
type RelatedPostingCancelReasonsIDEnum string
const (
	RelatedPostingCancelReasonsIDEnum_V352 RelatedPostingCancelReasonsIDEnum = "352"
	RelatedPostingCancelReasonsIDEnum_V400 RelatedPostingCancelReasonsIDEnum = "400"
	RelatedPostingCancelReasonsIDEnum_V401 RelatedPostingCancelReasonsIDEnum = "401"
	RelatedPostingCancelReasonsIDEnum_V402 RelatedPostingCancelReasonsIDEnum = "402"
	RelatedPostingCancelReasonsIDEnum_V665 RelatedPostingCancelReasonsIDEnum = "665"
	RelatedPostingCancelReasonsIDEnum_V666 RelatedPostingCancelReasonsIDEnum = "666"
	RelatedPostingCancelReasonsIDEnum_V667 RelatedPostingCancelReasonsIDEnum = "667"
)

// 取消货运提出方：
type RelatedPostingCancelReasonsTypeIDEnum string
const (
	RelatedPostingCancelReasonsTypeIDEnum_Buyer RelatedPostingCancelReasonsTypeIDEnum = "buyer"
	RelatedPostingCancelReasonsTypeIDEnum_Seller RelatedPostingCancelReasonsTypeIDEnum = "seller"
)

type RelatedPostingCancelReasons struct {
	ID int64 `json:"id"` // 取消原因ID： - `352` — 在卖家仓库已无商品。 - `400` — 只剩下残次品。 - `401` — 卖家拒绝了仲裁。 - `402` — 其他（卖家错误）。 - `665` — 买家没有收货。 - `666` — 快递服务退货...
	Title string `json:"title"` // 描述取消原因。
	TypeID RelatedPostingCancelReasonsTypeIDEnum `json:"type_id"` // 取消货运提出方： - `buyer` — 买家， - `seller` — 卖家。
}

type RelatedPostingCancelReason struct {
	PostingNumber string `json:"posting_number"` // 货运号。
	Reasons []RelatedPostingCancelReasons `json:"reasons"` // 取消订单原因。
}

type PostingCancelReasonResponse struct {
	Result []RelatedPostingCancelReason `json:"result"` // 请求结果。
}

// 收件人联系方式。
type V3Addressee struct {
	Name string `json:"name"` // 买家姓名。
	Phone string `json:"phone"` // 收件人的替代联系电话。
}

// 配送方法信息。
type PostingV4PostingFbsListResponsePostingsDeliveryMethod struct {
	Name string `json:"name"` // 配送方法名称。
	TPLProvider string `json:"tpl_provider"` // 配送服务。
	TPLProviderID int64 `json:"tpl_provider_id"` // 配送服务标识符。
	Warehouse string `json:"warehouse"` // 仓库名称。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	ID int64 `json:"id"` // 配送方式标识符。
}

// 外部平台订单信息。
type PostingV4PostingFbsListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"` // `true`，表示订单来自外部平台。
	PlatformName string `json:"platform_name"` // 下单平台名称。
}

// 在`next_tariff_starts_at`参数指定的时间后生效的最低折扣或加价。
type MoneyMoneyNextTariffMinCharge struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 在`next_tariff_starts_at`参数指定的时间后生效的折扣或加价。
type MoneyMoneyNextTariffCharge struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 发运的计费信息。
type PostingV4PostingFbsListResponsePostingsTariffication struct {
	CurrentTariffCharge MoneyMoneyCurrentTariffCharge `json:"current_tariff_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"` // 计费百分比。
	NextTariffMinCharge MoneyMoneyNextTariffMinCharge `json:"next_tariff_min_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"` // 在`next_tariff_starts_at`参数指定时间之后，该货件将按此百分比计费。
	NextTariffType string `json:"next_tariff_type"` // 在`next_tariff_starts_at`参数指定时间之后的计费类型——折扣或加价。
	CurrentTariffMinCharge MoneyMoneyCurrentTariffMinCharge `json:"current_tariff_min_charge"`
	CurrentTariffType string `json:"current_tariff_type"` // 计费类型——折扣或加价。
	NextTariffCharge MoneyMoneyNextTariffCharge `json:"next_tariff_charge"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"` // 新费率开始生效的日期和时间。
}

type PostingV4PostingFbsListResponsePostingsProducts struct {
	Imei []string `json:"imei"` // 移动设备的 IMEI 列表。
	Price MoneyPostingMoney `json:"price"`
	ProductColor string `json:"product_color"` // 商品颜色。
	Quantity int32 `json:"quantity"` // 货件中的商品数量。
	IsBlrTraceable bool `json:"is_blr_traceable"` // `true`，表示商品可追溯。
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"` // `true`，表示Ozon已购买该商品，用于配送至欧亚经济联盟及其他国家。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符，即货号。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符，即SKU。
	Weight float64 `json:"weight"` // 商品含包装重量。
}

// 商品佣金。
type PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission struct {
	Percent int64 `json:"percent"` // 佣金比例。
	Amount float64 `json:"amount"` // 商品佣金金额。
	Currency string `json:"currency"` // 计算佣金所使用的货币代码。
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProducts struct {
	Commission PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission `json:"commission"`
	Payout float64 `json:"payout"` // 卖家应收款。
	Price float64 `json:"price"` // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	ProductID int64 `json:"product_id"` // Ozon系统中的商品标识符，即SKU。
	Quantity int64 `json:"quantity"` // 货件中的商品数量。
	TotalDiscountPercent float64 `json:"total_discount_percent"` // 折扣百分比。
	Actions []string `json:"actions"` // 促销活动列表。
	CustomerPrice MoneyPostingMoney `json:"customer_price"`
	OldPrice float64 `json:"old_price"` // 折扣前的价格。商品卡上会以划线价显示。
	TotalDiscountValue float64 `json:"total_discount_value"` // 折扣金额。
}

// 商品价格、折扣金额、卖家应收款和佣金信息。
type PostingV4PostingFbsListResponsePostingsFinancialData struct {
	Products []PostingV4PostingFbsListResponsePostingsFinancialDataProducts `json:"products"` // 订单中的商品列表。
	ClusterFrom string `json:"cluster_from"` // 订单发出地区代码。
	ClusterTo string `json:"cluster_to"` // 订单配送地区代码。
}

// 带有附加特征的商品列表。
type PostingV4PostingFbsListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []string `json:"products_with_possible_mandatory_mark"` // 可能需要标志的商品列表。
}

// 收件人联系方式。
type PostingV4PostingFbsListResponsePostingsAddressee struct {
	Name string `json:"name"` // 收件人姓名。
}

// 货件条形码。
type PostingV4PostingFbsListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"` // 货件标签上的下方条形码。
	UpperBarcode string `json:"upper_barcode"` // 货件标签上的上方条形码。
}

// 分析数据。
// 支付方法：
type PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum string
const (
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "在线银行卡支付"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_OzonBank PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "Ozon Bank银行卡"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_Ozon PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "取货时从Ozon卡自动扣款"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_1 PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "收货时用已保存的银行卡支付"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_2 PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "快速支付系统"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_Ozon_1 PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "Ozon分期付款"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_3 PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "支付至结算账户"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_SberPay PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "SberPay"
	PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_4 PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "通过外部卖家预付款"
)

type PostingV4PostingFbsListResponsePostingsAnalyticsData struct {
	TPLProvider string `json:"tpl_provider"` // 配送服务。
	TPLProviderID int64 `json:"tpl_provider_id"` // 配送服务标识符。
	DeliveryDateEnd string `json:"delivery_date_end"` // 配送结束日期和时间。
	DeliveryType string `json:"delivery_type"` // 配送方法。
	IsPremium bool `json:"is_premium"` // `true`，表示收件人有Premium订阅。
	PaymentTypeGroupName PostingV4PostingFbsListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum `json:"payment_type_group_name"` // 支付方法： - `在线银行卡支付`； - `Ozon Bank银行卡`； - `取货时从Ozon卡自动扣款`； - `收货时用已保存的银行卡支付`； - `快速支付系统`； - `Ozon分期付款`； - `支付至结算账户`； - `Sbe...
	Warehouse string `json:"warehouse"` // 订单发货仓库名称。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	City string `json:"city"` // 配送城市。仅适用于rFBS货件和独联体卖家。
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"` // 配送开始的日期和时间。仅适用于通过Ozon配送创建的货件。
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"` // 订单预计最晚送达日期。仅适用于通过Ozon配送创建的货件。
	DeliveryDateBegin string `json:"delivery_date_begin"` // 配送开始的日期和时间。
	IsLegal bool `json:"is_legal"` // `true`，表示收件人为法人。
	Region string `json:"region"` // 配送地区。仅适用于rFBS货件。
}

// 配送地址。
type PostingV4PostingFbsListResponsePostingsCustomerAddress struct {
	ZipCode string `json:"zip_code"` // 收件人邮政编码。
	AddressTail string `json:"address_tail"` // 文本格式地址。
	Comment string `json:"comment"` // 订单备注。
	Country string `json:"country"` // 配送国家。
	District string `json:"district"` // 配送区域。
	Latitude float64 `json:"latitude"` // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
	City string `json:"city"` // 配送城市。
	ProviderPvzCode string `json:"provider_pvz_code"` // 3PL服务商取货点代码。
	PvzCode int64 `json:"pvz_code"` // 取货点代码。
	Region string `json:"region"` // 配送地区。
}

// 买家信息。
type PostingV4PostingFbsListResponsePostingsCustomer struct {
	CustomerID int64 `json:"customer_id"` // 买家标识符。
	Name string `json:"name"` // 买家姓名。
	Phone string `json:"phone"` // 买家的替换联系电话。
	Address PostingV4PostingFbsListResponsePostingsCustomerAddress `json:"address"`
	CustomerEmail string `json:"customer_email"` // 买家电子邮箱。
}

// 取消信息。
// 取消发起方：
type PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum string
const (
	PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum_V PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum = "卖家"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum_V_1 PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum = "客户"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum_V_2 PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum = "买家"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum_Ozon PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum = "Ozon"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum_V_3 PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum = "系统"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum_V_4 PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum = "配送服务商"
)

// 取消类型：
type PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum string
const (
	PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum_Seller PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum = "seller"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum_Client PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum = "client"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum_Customer PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum = "customer"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum_Ozon PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum = "ozon"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum_System PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum = "system"
	PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum_Delivery PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum = "delivery"
)

type PostingV4PostingFbsListResponsePostingsCancellation struct {
	CancelledAfterShip bool `json:"cancelled_after_ship"` // `true`，表示取消发生在货件完成备货之后。
	AffectCancellationRating bool `json:"affect_cancellation_rating"` // `true`，表示取消会影响卖家评级。
	CancelReason string `json:"cancel_reason"` // 取消原因。
	CancelReasonID int64 `json:"cancel_reason_id"` // 货件取消原因标识符。
	CancellationInitiator PostingV4PostingFbsListResponsePostingsCancellationCancellationInitiatorEnum `json:"cancellation_initiator"` // 取消发起方： - `卖家`， - `客户`， - `买家`， - `Ozon`， - `系统`， - `配送服务商`。
	CancellationType PostingV4PostingFbsListResponsePostingsCancellationCancellationTypeEnum `json:"cancellation_type"` // 取消类型： - `seller`——由卖家取消； - `client`或`customer`——由买家取消； - `ozon`——由Ozon取消； - `system`——由系统取消； - `delivery`——由配送服务取消。
}

// 需要补充附加信息的商品。 要将货件转入下一状态，请传递： - 制造国； - 货运报关单编号； - 商品批次登记号； - “诚实标志”标记； - 其他标记； - 重量。
type PostingV4PostingFbsListResponsePostingsRequirements struct {
	ProductsRequiringRnpt []string `json:"products_requiring_rnpt"` // 需要传递商品批次登记号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringWeight []string `json:"products_requiring_weight"` // 需要传递重量的商品列表。
	ProductsRequiringChangeCountry []string `json:"products_requiring_change_country"` // 需要更改制造国的商品标识符（SKU）列表。要更改制造国，请使用方法[/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountryProductFbsPosti...
	ProductsRequiringCountry []string `json:"products_requiring_country"` // 需要传递制造国信息的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v2/posting/fbs/product/country/set](#operation/PostingAPI_SetCountryProductFbsPost...
	ProductsRequiringGTD []string `json:"products_requiring_gtd"` // 需要传递货运报关单编号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringImei []string `json:"products_requiring_imei"` // 需要传递IMEI的商品标识符列表。
	ProductsRequiringJwUin []string `json:"products_requiring_jw_uin"` // 需要传递珠宝制品唯一标识码的商品列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExemplarSe...
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"` // С需要传递“诚实标志”标识的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductEx...
}

// 买家的法务信息。
type PostingV4PostingFbsListResponsePostingsLegalInfo struct {
	Kpp string `json:"kpp"` // 纳税人登记原因代码。
	CompanyName string `json:"company_name"` // 公司名称。
	Inn string `json:"inn"` // 税号。
}

// 配送模式：
type PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum string
const (
	PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum_SDS PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum = "SDS"
	PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum_FBO PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum = "FBO"
	PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum_FBS PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum = "FBS"
	PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum_Crossborder PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum = "Crossborder"
)

// 与配送服务的集成类型：
type PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum string
const (
	PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum_Ozon PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum = "ozon"
	PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum_V3plTracking PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum = "3pl_tracking"
	PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum_NonIntegrated PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum = "non_integrated"
	PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum_Aggregator PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum = "aggregator"
	PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum_Hybryd PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum = "hybryd"
)

// 货件子状态：
type PostingV4PostingFbsListResponsePostingsSubstatusEnum string
const (
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingAcceptanceInProgress PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_acceptance_in_progress"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingInArbitration PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_in_arbitration"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingCreated PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_created"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingInCarriage PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_in_carriage"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingNotInCarriage PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_not_in_carriage"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingRegistered PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_registered"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingTransferringToDelivery PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_transferring_to_delivery"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_StatusAwaitingDeliver PostingV4PostingFbsListResponsePostingsSubstatusEnum = "status=awaiting_deliver"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingAwaitingPassportData PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_awaiting_passport_data"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingAwaitingRegistration PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_awaiting_registration"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingRegistrationError PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_registration_error"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_StatusAwaitingRegistration PostingV4PostingFbsListResponsePostingsSubstatusEnum = "status=awaiting_registration"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingSplitPending PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_split_pending"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingCanceled PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_canceled"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingInClientArbitration PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_in_client_arbitration"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingDelivered PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_delivered"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingReceived PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_received"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingConditionallyDelivered PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_conditionally_delivered"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingInCourierService PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_in_courier_service"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingInPickupPoint PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_in_pickup_point"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingOnWayToCity PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_on_way_to_city"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingOnWayToPickupPoint PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_on_way_to_pickup_point"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingReturnedToWarehouse PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_returned_to_warehouse"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingTransferredToCourierService PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_transferred_to_courier_service"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingDriverPickUp PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_driver_pick_up"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_PostingNotInSortCenter PostingV4PostingFbsListResponsePostingsSubstatusEnum = "posting_not_in_sort_center"
	PostingV4PostingFbsListResponsePostingsSubstatusEnum_ShipFailed PostingV4PostingFbsListResponsePostingsSubstatusEnum = "ship_failed"
)

// 货件的可用操作和信息：
type PostingV4PostingFbsListResponsePostingsAvailableActionsEnum string
const (
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_Arbitration PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "arbitration"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_AwaitingDelivery PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "awaiting_delivery"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_CanCreateChat PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "can_create_chat"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_Cancel PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "cancel"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_ClickTrackNumber PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "click_track_number"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_CustomerPhoneAvailable PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "customer_phone_available"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_HasWeightProducts PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "has_weight_products"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_HideRegionAndCity PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "hide_region_and_city"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_InvoiceGet PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "invoice_get"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_InvoiceSend PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "invoice_send"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_InvoiceUpdate PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "invoice_update"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_LabelDownloadBig PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "label_download_big"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_LabelDownloadSmall PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "label_download_small"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_LabelDownload PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "label_download"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_NonIntDelivered PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "non_int_delivered"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_NonIntDelivering PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "non_int_delivering"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_NonIntLastMile PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "non_int_last_mile"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_ProductCancel PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "product_cancel"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_SetCutoff PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "set_cutoff"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_ShipmentDate PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "shipment_date"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_SetTimeslot PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "set_timeslot"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_SetTrackNumber PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "set_track_number"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_ShipAsyncInProcess PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "ship_async_in_process"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_ShipAsyncRetry PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "ship_async_retry"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_ShipAsync PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "ship_async"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_ShipWithAdditionalInfo PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "ship_with_additional_info"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_Ship PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "ship"
	PostingV4PostingFbsListResponsePostingsAvailableActionsEnum_UpdateCis PostingV4PostingFbsListResponsePostingsAvailableActionsEnum = "update_cis"
)

// 货件状态：
type PostingV4PostingFbsListResponsePostingsStatusEnum string
const (
	PostingV4PostingFbsListResponsePostingsStatusEnum_AcceptanceInProgress PostingV4PostingFbsListResponsePostingsStatusEnum = "acceptance_in_progress"
	PostingV4PostingFbsListResponsePostingsStatusEnum_Arbitration PostingV4PostingFbsListResponsePostingsStatusEnum = "arbitration"
	PostingV4PostingFbsListResponsePostingsStatusEnum_AwaitingApprove PostingV4PostingFbsListResponsePostingsStatusEnum = "awaiting_approve"
	PostingV4PostingFbsListResponsePostingsStatusEnum_AwaitingDeliver PostingV4PostingFbsListResponsePostingsStatusEnum = "awaiting_deliver"
	PostingV4PostingFbsListResponsePostingsStatusEnum_AwaitingPackaging PostingV4PostingFbsListResponsePostingsStatusEnum = "awaiting_packaging"
	PostingV4PostingFbsListResponsePostingsStatusEnum_AwaitingRegistration PostingV4PostingFbsListResponsePostingsStatusEnum = "awaiting_registration"
	PostingV4PostingFbsListResponsePostingsStatusEnum_AwaitingVerification PostingV4PostingFbsListResponsePostingsStatusEnum = "awaiting_verification"
	PostingV4PostingFbsListResponsePostingsStatusEnum_Cancelled PostingV4PostingFbsListResponsePostingsStatusEnum = "cancelled"
	PostingV4PostingFbsListResponsePostingsStatusEnum_CancelledFromSplitPending PostingV4PostingFbsListResponsePostingsStatusEnum = "cancelled_from_split_pending"
	PostingV4PostingFbsListResponsePostingsStatusEnum_ClientArbitration PostingV4PostingFbsListResponsePostingsStatusEnum = "client_arbitration"
	PostingV4PostingFbsListResponsePostingsStatusEnum_Delivering PostingV4PostingFbsListResponsePostingsStatusEnum = "delivering"
	PostingV4PostingFbsListResponsePostingsStatusEnum_DriverPickup PostingV4PostingFbsListResponsePostingsStatusEnum = "driver_pickup"
	PostingV4PostingFbsListResponsePostingsStatusEnum_NotAccepted PostingV4PostingFbsListResponsePostingsStatusEnum = "not_accepted"
)

// 装卸服务代码：
type PostingV4PostingFbsListResponsePostingsPrrOptionEnum string
const (
	PostingV4PostingFbsListResponsePostingsPrrOptionEnum_Lift PostingV4PostingFbsListResponsePostingsPrrOptionEnum = "lift"
	PostingV4PostingFbsListResponsePostingsPrrOptionEnum_Stairs PostingV4PostingFbsListResponsePostingsPrrOptionEnum = "stairs"
	PostingV4PostingFbsListResponsePostingsPrrOptionEnum_None PostingV4PostingFbsListResponsePostingsPrrOptionEnum = "none"
	PostingV4PostingFbsListResponsePostingsPrrOptionEnum_DeliveryDefault PostingV4PostingFbsListResponsePostingsPrrOptionEnum = "delivery_default"
)

type PostingV4PostingFbsListResponsePostings struct {
	DeliveryMethod PostingV4PostingFbsListResponsePostingsDeliveryMethod `json:"delivery_method"`
	DestinationPlaceName string `json:"destination_place_name"` // 目的地名称。
	ExternalOrder PostingV4PostingFbsListResponsePostingsExternalOrder `json:"external_order"`
	InProcessAt string `json:"in_process_at"` // 货件开始处理的日期和时间。
	Optional PostingV4PostingFbsListResponsePostingsOptional `json:"optional"`
	Substatus PostingV4PostingFbsListResponsePostingsSubstatusEnum `json:"substatus"` // 货件子状态： - `posting_acceptance_in_progress`——验收中； - `posting_in_arbitration`——仲裁； - `posting_created`——已创建； - `posting_in_...
	Addressee PostingV4PostingFbsListResponsePostingsAddressee `json:"addressee"`
	Barcodes PostingV4PostingFbsListResponsePostingsBarcodes `json:"barcodes"`
	MultiBoxQty int32 `json:"multi_box_qty"` // 商品包装所用箱数。
	Tariffication PostingV4PostingFbsListResponsePostingsTariffication `json:"tariffication"`
	AvailableActions []string `json:"available_actions"` // 货件的可用操作和信息： - `arbitration`——提出争议； - `awaiting_delivery`——变更为“等待发运”状态； - `can_create_chat`——与买家发起聊天； - `cancel`——取消货件； -...
	DeliveringDate string `json:"delivering_date"` // 货件转移配送的日期。
	IsClickAndCollect bool `json:"is_click_and_collect"` // `true`，表示货件通过“到店自提”方式配送。
	ParentPostingNumber string `json:"parent_posting_number"` // 当前货件所拆分自的父货件编号。
	Status PostingV4PostingFbsListResponsePostingsStatusEnum `json:"status"` // 货件状态： - `acceptance_in_progress`——接收中； - `arbitration`——仲裁； - `awaiting_approve`——等待确认； - `awaiting_deliver`——等待发运； - `a...
	AnalyticsData PostingV4PostingFbsListResponsePostingsAnalyticsData `json:"analytics_data"`
	Customer PostingV4PostingFbsListResponsePostingsCustomer `json:"customer"`
	DestinationPlaceID int64 `json:"destination_place_id"` // 目的地标识符。
	TarifficationSteps []PostingV4PostingFbsListResponsePostingsTarifficationStep `json:"tariffication_steps"` // 计费阶段。
	VolumeWeight float64 `json:"volume_weight"` // 商品体积重量。
	Products []PostingV4PostingFbsListResponsePostingsProducts `json:"products"` // 货件中商品列表。
	PrrOption PostingV4PostingFbsListResponsePostingsPrrOptionEnum `json:"prr_option"` // 装卸服务代码： - `lift`——通过电梯搬运上楼； - `stairs`——通过楼梯搬运上楼； - `none`——买家拒绝该服务，无需将商品搬上楼； - `delivery_default`——配送费用已包含在价格中，根据要约条款需将...
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"` // `true`，表示需要填写可追踪属性。
	ShipmentDate string `json:"shipment_date"` // 需要在此日期和时间之前完成货件备货。系统会显示推荐的发运时间。超过该时间后将开始适用新费率，相关信息可在`tariffication`字段中获取。
	PostingNumber string `json:"posting_number"` // 货件编号。
	Cancellation PostingV4PostingFbsListResponsePostingsCancellation `json:"cancellation"`
	DeliverySchema PostingV4PostingFbsListResponsePostingsDeliverySchemaEnum `json:"delivery_schema"` // 配送模式： - `SDS`——统一SKU标识符； - `FBO`——从Ozon仓库销售的商品标识符； - `FBS`——从FBS仓库销售的商品标识符； - `Crossborder`——从境外销售的商品标识符。
	IsExpress bool `json:"is_express"` // `true`，表示使用了Ozon Express快速配送。
	IsPresortable bool `json:"is_presortable"` // `true`，表示商品为错误分级的商品。
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"` // 快递员代码校验成功的日期和时间。请通过方法 [/v1/posting/fbs/pick-up-code/verify](#operation/PostingAPI_PostingFBSPickupCodeVerify)校验快递员代码。
	QuantumID int64 `json:"quantum_id"` // 经济商品标识符。
	Requirements PostingV4PostingFbsListResponsePostingsRequirements `json:"requirements"`
	FinancialData PostingV4PostingFbsListResponsePostingsFinancialData `json:"financial_data"`
	OrderID int64 `json:"order_id"` // 该货件所属订单的标识符。
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"` // 不逾期发运的日期和时间。
	TPLIntegrationType PostingV4PostingFbsListResponsePostingsTPLIntegrationTypeEnum `json:"tpl_integration_type"` // 与配送服务的集成类型： - `ozon`——由Ozon配送； - `3pl_tracking`——由已集成服务配送； - `non_integrated`——由第三方服务配送； - `aggregator`——通过Ozon合作伙伴配送； -...
	IsMultibox bool `json:"is_multibox"` // 标记货件中是否包含多箱商品，以及是否需要传递其箱数： - `true`——备货前，请通过方法[/v3/posting/multiboxqty/set](#operation/PostingAPI_PostingMultiBoxQtySetV...
	LegalInfo PostingV4PostingFbsListResponsePostingsLegalInfo `json:"legal_info"`
	OrderNumber string `json:"order_number"` // 该货件所属订单的编号。
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号码。
}

type V1SetPostingsRequest struct {
	CarriageID int64 `json:"carriage_id"` // 发运识别符。
	PostingNumbers []string `json:"posting_numbers"` // 最新货件列表。
}

// 筛选器。
// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyFbsProductListRequestFilterCutoffFromEnum string
const (
	V1AssemblyFbsProductListRequestFilterCutoffFromEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyFbsProductListRequestFilterCutoffFromEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyFbsProductListRequestFilterCutoffFromEnum_V20200318T073450359Z V1AssemblyFbsProductListRequestFilterCutoffFromEnum = "2020-03-18T07:34:50.359Z"
)

// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyFbsProductListRequestFilterCutoffToEnum string
const (
	V1AssemblyFbsProductListRequestFilterCutoffToEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyFbsProductListRequestFilterCutoffToEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyFbsProductListRequestFilterCutoffToEnum_V20200318T073450359Z V1AssemblyFbsProductListRequestFilterCutoffToEnum = "2020-03-18T07:34:50.359Z"
)

type V1AssemblyFbsProductListRequestFilter struct {
	CutoffFrom V1AssemblyFbsProductListRequestFilterCutoffFromEnum `json:"cutoff_from"` // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	CutoffTo V1AssemblyFbsProductListRequestFilterCutoffToEnum `json:"cutoff_to"` // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	DeliveryMethodID int64 `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
}

// 货件状态最后一次发生变更的时间段。
type PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"` // 周期开始日期。
	To string `json:"to"` // 周期结束日期。
}

// 筛选器。
// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyCarriagePostingListRequestFilterCutoffFromEnum string
const (
	V1AssemblyCarriagePostingListRequestFilterCutoffFromEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyCarriagePostingListRequestFilterCutoffFromEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyCarriagePostingListRequestFilterCutoffFromEnum_V20200318T073450359Z V1AssemblyCarriagePostingListRequestFilterCutoffFromEnum = "2020-03-18T07:34:50.359Z"
)

// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyCarriagePostingListRequestFilterCutoffToEnum string
const (
	V1AssemblyCarriagePostingListRequestFilterCutoffToEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyCarriagePostingListRequestFilterCutoffToEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyCarriagePostingListRequestFilterCutoffToEnum_V20200318T073450359Z V1AssemblyCarriagePostingListRequestFilterCutoffToEnum = "2020-03-18T07:34:50.359Z"
)

type V1AssemblyCarriagePostingListRequestFilter struct {
	DeliveryMethodID int64 `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
	CutoffFrom V1AssemblyCarriagePostingListRequestFilterCutoffFromEnum `json:"cutoff_from"` // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	CutoffTo V1AssemblyCarriagePostingListRequestFilterCutoffToEnum `json:"cutoff_to"` // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
}

type V1AssemblyCarriagePostingListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter V1AssemblyCarriagePostingListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 每页显示的数量。
}

type V1AssemblyCarriageProductListResponseProduct struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符——货号。
	PictureURL string `json:"picture_url"` // 商品图片链接。
	PostingNumbers []string `json:"posting_numbers"` // 货件编号。
	ProductName string `json:"product_name"` // 商品名称。
	Quantity int64 `json:"quantity"` // 商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符——SKU。
}

type V1AssemblyCarriageProductListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Products []V1AssemblyCarriageProductListResponseProduct `json:"products"` // 商品列表。
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []string `json:"posting_number"` // 货件ID。
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []string `json:"related_posting_numbers"` // 货件号。
}

// 买家信息。
type V3Customer struct {
	Name string `json:"name"` // 买家姓名。
	Phone string `json:"phone"` // 买家的替代联系电话。
	Address V3Address `json:"address"`
	CustomerID int64 `json:"customer_id"` // 买家ID。
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []string `json:"posting_number"` // 货件ID。
}

// 货件条形码。
type PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"` // 货件标签上的下方条形码。
	UpperBarcode string `json:"upper_barcode"` // 货件标签上的上方条形码。
}

// 错误类型：
type ResultErrorStatusEnum string
const (
	ResultErrorStatusEnum_Warning ResultErrorStatusEnum = "warning"
	ResultErrorStatusEnum_Critical ResultErrorStatusEnum = "critical"
)

type ResultError struct {
	Code string `json:"code"` // 错误代码。
	Status ResultErrorStatusEnum `json:"status"` // 错误类型： - `warning` — 提醒； - `critical` — 严重错误。
}

type V2PostingFBSActGetPostingsRequest struct {
	ID any `json:"id"` // 单据标识符。请通过方法[/v1/carriage/create](#operation/CarriageAPI_CarriageCreate)获取参数值。
}

type V2FbsPostingProductCountrySetResponse struct {
	ProductID int64 `json:"product_id"` // 产品ID。
	IsGTDNeeded bool `json:"is_gtd_needed"` // 有必要传送产品和货运的货物报关单（Cargo Customs Declaration）编号的标志。
}

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type PostingV4PostingFbsUnfulfilledListRequestSortDirEnum string

// 需要添加到响应中的附加字段。
type PostingV4PostingFbsUnfulfilledListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"` // 若为`true`，则在响应中添加分析数据。
	Barcodes bool `json:"barcodes"` // `true`，表示要在响应中添加货件条形码。
	FinancialData bool `json:"financial_data"` // 若为`true`，则在响应中添加财务数据。
	LegalInfo bool `json:"legal_info"` // 若为`true`，则在响应中添加法务信息。
}

// 请求筛选器。 使用按备货时间的筛选器——`cutoff`，或按货件转移配送的日期筛选——`delivering_date`。 如果同时使用这两个筛选器，响应会返回错误。 要使用按备货时间的筛选器，请填写`cutoff_from`和`cuto...
// 货件状态：
type PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum string
const (
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_AcceptanceInProgress PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "acceptance_in_progress"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_AwaitingApprove PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "awaiting_approve"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_AwaitingPackaging PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "awaiting_packaging"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_AwaitingRegistration PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "awaiting_registration"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_AwaitingDeliver PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "awaiting_deliver"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_Arbitration PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "arbitration"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_ClientArbitration PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "client_arbitration"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_Delivering PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "delivering"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_DriverPickup PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "driver_pickup"
	PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum_NotAccepted PostingV4PostingFbsUnfulfilledListRequestFilterStatusesEnum = "not_accepted"
)

type PostingV4PostingFbsUnfulfilledListRequestFilter struct {
	DeliveringDateTo string `json:"delivering_date_to"` // 货件转移配送的最晚日期。
	DeliveryMethodIds []string `json:"delivery_method_ids"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	ProviderIds []string `json:"provider_ids"` // 配送服务标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	WarehouseIds []string `json:"warehouse_ids"` // 仓库标识符。可通过方法[/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	DeliveringDateFrom string `json:"delivering_date_from"` // 货件转移配送的最早日期。
	LastChangedStatusDate PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate `json:"last_changed_status_date"`
	Statuses []string `json:"statuses"` // 货件状态： - `acceptance_in_progress`——接收中； - `awaiting_approve`——等待确认； - `awaiting_packaging`——等待包装； - `awaiting_registratio...
	CutoffFrom string `json:"cutoff_from"` // 卖家必须完成订单备货的截止时间。时间段开始。
	CutoffTo string `json:"cutoff_to"` // 卖家必须完成订单备货的截止时间。时间段结束。
}

type PostingV4PostingFbsUnfulfilledListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter PostingV4PostingFbsUnfulfilledListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortDir PostingV4PostingFbsUnfulfilledListRequestSortDirEnum `json:"sort_dir"`
	Translit bool `json:"translit"` // 则启用将地址从西里尔字母转写为拉丁字母。
	With PostingV4PostingFbsUnfulfilledListRequestWith `json:"with"`
}

// 筛选器。
// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyCarriageProductListRequestFilterCutoffFromEnum string
const (
	V1AssemblyCarriageProductListRequestFilterCutoffFromEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyCarriageProductListRequestFilterCutoffFromEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyCarriageProductListRequestFilterCutoffFromEnum_V20200318T073450359Z V1AssemblyCarriageProductListRequestFilterCutoffFromEnum = "2020-03-18T07:34:50.359Z"
)

// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyCarriageProductListRequestFilterCutoffToEnum string
const (
	V1AssemblyCarriageProductListRequestFilterCutoffToEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyCarriageProductListRequestFilterCutoffToEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyCarriageProductListRequestFilterCutoffToEnum_V20200318T073450359Z V1AssemblyCarriageProductListRequestFilterCutoffToEnum = "2020-03-18T07:34:50.359Z"
)

type V1AssemblyCarriageProductListRequestFilter struct {
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
	CutoffFrom V1AssemblyCarriageProductListRequestFilterCutoffFromEnum `json:"cutoff_from"` // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	CutoffTo V1AssemblyCarriageProductListRequestFilterCutoffToEnum `json:"cutoff_to"` // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	DeliveryMethodID int64 `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
}

// 分析数据。
// 支付方法：
type PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "在线银行卡支付"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_OzonBank PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "Ozon Bank银行卡"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_Ozon PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "取货时从Ozon卡自动扣款"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_1 PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "收货时用已保存的银行卡支付"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_2 PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "快速支付系统"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_Ozon_1 PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "Ozon分期付款"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_3 PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "支付至结算账户"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_SberPay PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "SberPay"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum_V_4 PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum = "通过外部卖家预付款"
)

type PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData struct {
	IsLegal bool `json:"is_legal"` // `true`，表示收件人为法人。
	Region string `json:"region"` // 配送地区。仅适用于rFBS货件。
	City string `json:"city"` // 配送城市。仅适用于rFBS货件和独联体卖家。
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"` // 配送开始的日期和时间。仅适用于通过Ozon配送创建的货件。
	IsPremium bool `json:"is_premium"` // `true`，表示收件人有Premium订阅。
	PaymentTypeGroupName PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsDataPaymentTypeGroupNameEnum `json:"payment_type_group_name"` // 支付方法： - `在线银行卡支付`； - `Ozon Bank银行卡`； - `取货时从Ozon卡自动扣款`； - `收货时用已保存的银行卡支付`； - `快速支付系统`； - `Ozon分期付款`； - `支付至结算账户`； - `Sbe...
	TPLProvider string `json:"tpl_provider"` // 配送服务。
	TPLProviderID int64 `json:"tpl_provider_id"` // 配送服务标识符。
	Warehouse string `json:"warehouse"` // 订单发货仓库名称。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"` // 订单预计最晚送达日期。仅适用于通过Ozon配送创建的货件。
	DeliveryDateBegin string `json:"delivery_date_begin"` // 配送开始的日期和时间。
	DeliveryDateEnd string `json:"delivery_date_end"` // 配送结束日期和时间。
	DeliveryType string `json:"delivery_type"` // 配送方法。
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
}

type GetCarriageAvailableListResponseResult struct {
	MandatoryPackagedCount int32 `json:"mandatory_packaged_count"` // 收取货件数量。
	TPLProviderIconURL string `json:"tpl_provider_icon_url"` // 快递服务图标的链接。
	WarehouseName string `json:"warehouse_name"` // 仓库名称。
	DeliveryMethodID int64 `json:"delivery_method_id"` // 快递方式ID。
	FirstMileType string `json:"first_mile_type"` // 第一英里类型。
	RecommendedTimeLocal string `json:"recommended_time_local"` // 推荐的本地发运时间（订单接收点）。
	RecommendedTimeUtcOffsetInMinutes float64 `json:"recommended_time_utc_offset_in_minutes"` // 推荐发运时间与UTC-0的时区偏移量（以分钟为单位）。
	WarehouseCity string `json:"warehouse_city"` // 仓库所在城市。
	WarehouseID int64 `json:"warehouse_id"` // 仓库ID。
	CarriagePostingsCount int32 `json:"carriage_postings_count"` // 运输中的货件数量。
	CutoffAt string `json:"cutoff_at"` // 需要收取货件的日期和时间。
	DeliveryMethodName string `json:"delivery_method_name"` // 快递方式名称。
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"` // 信任接收的标志。 如果在仓库中启用了信任接收，则为“true”。
	MandatoryPostingsCount int32 `json:"mandatory_postings_count"` // 需要收取的货件数量。
	WarehouseTimezone string `json:"warehouse_timezone"` // 仓库所在时区。
	CarriageID int64 `json:"carriage_id"` // 运输ID（也是文件形成的任务编号）。
	CarriageStatus string `json:"carriage_status"` // 所请求的交付方式和装运日期的运输状态。
	Errors []ResultError `json:"errors"` // 错误列表。
	TPLProviderName string `json:"tpl_provider_name"` // 快递服务名称。
}

type Postingv1GetCarriageAvailableListResponse struct {
	Result []GetCarriageAvailableListResponseResult `json:"result"` // 方法操作结果。
}

// 相关货件。
type V3FbsPostingDetailRelatedPostings struct {
	RelatedPostingNumbers []string `json:"related_posting_numbers"` // 相关货件号码列表。
}

// 快递员信息。
type FbsPostingDetailCourier struct {
	CarModel string `json:"car_model"` // 汽车型号。
	CarNumber string `json:"car_number"` // 车牌号。
	Name string `json:"name"` // 快递员全名。
	Phone string `json:"phone"` // 快递员电话。 过时的参数，不再使用。并总是返回到空字符串 `""`。
}

// 商品尺寸。
type V3Dimensions struct {
	Height string `json:"height"` // 包装高度。
	Length string `json:"length"` // 商品长度。
	Weight string `json:"weight"` // 商品包装重量。
	Width string `json:"width"` // 包装宽度。
}

// 商品尺寸。
// 价格显示的货币，其与个人中心中设置的币种相匹配
type V3PostingProductDetailCurrencyCodeEnum string
const (
	V3PostingProductDetailCurrencyCodeEnum_RUB V3PostingProductDetailCurrencyCodeEnum = "RUB"
	V3PostingProductDetailCurrencyCodeEnum_BYN V3PostingProductDetailCurrencyCodeEnum = "BYN"
	V3PostingProductDetailCurrencyCodeEnum_KZT V3PostingProductDetailCurrencyCodeEnum = "KZT"
	V3PostingProductDetailCurrencyCodeEnum_EUR V3PostingProductDetailCurrencyCodeEnum = "EUR"
	V3PostingProductDetailCurrencyCodeEnum_USD V3PostingProductDetailCurrencyCodeEnum = "USD"
	V3PostingProductDetailCurrencyCodeEnum_CNY V3PostingProductDetailCurrencyCodeEnum = "CNY"
)

type V3PostingProductDetail struct {
	Dimensions V3Dimensions `json:"dimensions"`
	Name string `json:"name"` // 名称。
	Price string `json:"price"` // 折扣后商品价格 — 该值在商品卡片上显示。
	CurrencyCode V3PostingProductDetailCurrencyCodeEnum `json:"currency_code"` // 价格显示的货币，其与个人中心中设置的币种相匹配。 -`RUB` — 俄罗斯卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
	HasImei bool `json:"has_imei"` // 存在 IMEI。 若存在 IMEI，则为`true`。
	MandatoryMark []string `json:"mandatory_mark"` // 商品强制性标签。
	OfferID string `json:"offer_id"` // 卖家系统中的商品ID — 货号。
	Quantity int32 `json:"quantity"` // 商品数量。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
}

type V3AdditionalDataItem struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

// 商品和副本列表。
type V3FbsPostingExemplarProductV3 struct {
	Exemplars []V3FbsPostingProductExemplarInfoV3 `json:"exemplars"` // 按副本的信息。
	SKU int64 `json:"sku"` // 在Ozon系统中的产品ID — SKU。
}

// 有关产品及其副本的信息。 响应包含 `product_exemplars`字段, 如果在请求中传递标志 `with.product_exemplars = true`。
type V3FbsPostingProductExemplarsV3 struct {
	Products []V3FbsPostingExemplarProductV3 `json:"products"`
}

// 货件信息。
// 快递服务集成类型：
type V3FbsPostingDetailTPLIntegrationTypeEnum string
const (
	V3FbsPostingDetailTPLIntegrationTypeEnum_Ozon V3FbsPostingDetailTPLIntegrationTypeEnum = "ozon"
	V3FbsPostingDetailTPLIntegrationTypeEnum_Aggregator V3FbsPostingDetailTPLIntegrationTypeEnum = "aggregator"
	V3FbsPostingDetailTPLIntegrationTypeEnum_V3plTracking V3FbsPostingDetailTPLIntegrationTypeEnum = "3pl_tracking"
	V3FbsPostingDetailTPLIntegrationTypeEnum_NonIntegrated V3FbsPostingDetailTPLIntegrationTypeEnum = "non_integrated"
)

// 货件的前一个子状态
type V3FbsPostingDetailPreviousSubstatusEnum string
const (
	V3FbsPostingDetailPreviousSubstatusEnum_PostingAcceptanceInProgress V3FbsPostingDetailPreviousSubstatusEnum = "posting_acceptance_in_progress"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingInArbitration V3FbsPostingDetailPreviousSubstatusEnum = "posting_in_arbitration"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingCreated V3FbsPostingDetailPreviousSubstatusEnum = "posting_created"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingInCarriage V3FbsPostingDetailPreviousSubstatusEnum = "posting_in_carriage"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingNotInCarriage V3FbsPostingDetailPreviousSubstatusEnum = "posting_not_in_carriage"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingRegistered V3FbsPostingDetailPreviousSubstatusEnum = "posting_registered"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingTransferringToDelivery V3FbsPostingDetailPreviousSubstatusEnum = "posting_transferring_to_delivery"
	V3FbsPostingDetailPreviousSubstatusEnum_StatusAwaitingDeliver V3FbsPostingDetailPreviousSubstatusEnum = "status=awaiting_deliver"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingAwaitingPassportData V3FbsPostingDetailPreviousSubstatusEnum = "posting_awaiting_passport_data"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingAwaitingRegistration V3FbsPostingDetailPreviousSubstatusEnum = "posting_awaiting_registration"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingRegistrationError V3FbsPostingDetailPreviousSubstatusEnum = "posting_registration_error"
	V3FbsPostingDetailPreviousSubstatusEnum_StatusAwaitingRegistration V3FbsPostingDetailPreviousSubstatusEnum = "status=awaiting_registration"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingSplitPending V3FbsPostingDetailPreviousSubstatusEnum = "posting_split_pending"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingCanceled V3FbsPostingDetailPreviousSubstatusEnum = "posting_canceled"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingInClientArbitration V3FbsPostingDetailPreviousSubstatusEnum = "posting_in_client_arbitration"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingDelivered V3FbsPostingDetailPreviousSubstatusEnum = "posting_delivered"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingReceived V3FbsPostingDetailPreviousSubstatusEnum = "posting_received"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingConditionallyDelivered V3FbsPostingDetailPreviousSubstatusEnum = "posting_conditionally_delivered"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingInCourierService V3FbsPostingDetailPreviousSubstatusEnum = "posting_in_courier_service"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingInPickupPoint V3FbsPostingDetailPreviousSubstatusEnum = "posting_in_pickup_point"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingOnWayToCity V3FbsPostingDetailPreviousSubstatusEnum = "posting_on_way_to_city"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingOnWayToPickupPoint V3FbsPostingDetailPreviousSubstatusEnum = "posting_on_way_to_pickup_point"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingReturnedToWarehouse V3FbsPostingDetailPreviousSubstatusEnum = "posting_returned_to_warehouse"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingTransferredToCourierService V3FbsPostingDetailPreviousSubstatusEnum = "posting_transferred_to_courier_service"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingDriverPickUp V3FbsPostingDetailPreviousSubstatusEnum = "posting_driver_pick_up"
	V3FbsPostingDetailPreviousSubstatusEnum_PostingNotInSortCenter V3FbsPostingDetailPreviousSubstatusEnum = "posting_not_in_sort_center"
)

// 发货子状态：
type V3FbsPostingDetailSubstatusEnum string
const (
	V3FbsPostingDetailSubstatusEnum_PostingAcceptanceInProgress V3FbsPostingDetailSubstatusEnum = "posting_acceptance_in_progress"
	V3FbsPostingDetailSubstatusEnum_PostingInArbitration V3FbsPostingDetailSubstatusEnum = "posting_in_arbitration"
	V3FbsPostingDetailSubstatusEnum_PostingCreated V3FbsPostingDetailSubstatusEnum = "posting_created"
	V3FbsPostingDetailSubstatusEnum_PostingInCarriage V3FbsPostingDetailSubstatusEnum = "posting_in_carriage"
	V3FbsPostingDetailSubstatusEnum_PostingNotInCarriage V3FbsPostingDetailSubstatusEnum = "posting_not_in_carriage"
	V3FbsPostingDetailSubstatusEnum_PostingRegistered V3FbsPostingDetailSubstatusEnum = "posting_registered"
	V3FbsPostingDetailSubstatusEnum_PostingTransferringToDelivery V3FbsPostingDetailSubstatusEnum = "posting_transferring_to_delivery"
	V3FbsPostingDetailSubstatusEnum_StatusAwaitingDeliver V3FbsPostingDetailSubstatusEnum = "status=awaiting_deliver"
	V3FbsPostingDetailSubstatusEnum_PostingAwaitingPassportData V3FbsPostingDetailSubstatusEnum = "posting_awaiting_passport_data"
	V3FbsPostingDetailSubstatusEnum_PostingAwaitingRegistration V3FbsPostingDetailSubstatusEnum = "posting_awaiting_registration"
	V3FbsPostingDetailSubstatusEnum_PostingRegistrationError V3FbsPostingDetailSubstatusEnum = "posting_registration_error"
	V3FbsPostingDetailSubstatusEnum_StatusAwaitingRegistration V3FbsPostingDetailSubstatusEnum = "status=awaiting_registration"
	V3FbsPostingDetailSubstatusEnum_PostingSplitPending V3FbsPostingDetailSubstatusEnum = "posting_split_pending"
	V3FbsPostingDetailSubstatusEnum_PostingCanceled V3FbsPostingDetailSubstatusEnum = "posting_canceled"
	V3FbsPostingDetailSubstatusEnum_PostingInClientArbitration V3FbsPostingDetailSubstatusEnum = "posting_in_client_arbitration"
	V3FbsPostingDetailSubstatusEnum_PostingDelivered V3FbsPostingDetailSubstatusEnum = "posting_delivered"
	V3FbsPostingDetailSubstatusEnum_PostingReceived V3FbsPostingDetailSubstatusEnum = "posting_received"
	V3FbsPostingDetailSubstatusEnum_PostingConditionallyDelivered V3FbsPostingDetailSubstatusEnum = "posting_conditionally_delivered"
	V3FbsPostingDetailSubstatusEnum_PostingInCourierService V3FbsPostingDetailSubstatusEnum = "posting_in_courier_service"
	V3FbsPostingDetailSubstatusEnum_PostingInPickupPoint V3FbsPostingDetailSubstatusEnum = "posting_in_pickup_point"
	V3FbsPostingDetailSubstatusEnum_PostingOnWayToCity V3FbsPostingDetailSubstatusEnum = "posting_on_way_to_city"
	V3FbsPostingDetailSubstatusEnum_PostingOnWayToPickupPoint V3FbsPostingDetailSubstatusEnum = "posting_on_way_to_pickup_point"
	V3FbsPostingDetailSubstatusEnum_PostingReturnedToWarehouse V3FbsPostingDetailSubstatusEnum = "posting_returned_to_warehouse"
	V3FbsPostingDetailSubstatusEnum_PostingTransferredToCourierService V3FbsPostingDetailSubstatusEnum = "posting_transferred_to_courier_service"
	V3FbsPostingDetailSubstatusEnum_PostingDriverPickUp V3FbsPostingDetailSubstatusEnum = "posting_driver_pick_up"
	V3FbsPostingDetailSubstatusEnum_PostingNotInSortCenter V3FbsPostingDetailSubstatusEnum = "posting_not_in_sort_center"
	V3FbsPostingDetailSubstatusEnum_ShipFailed V3FbsPostingDetailSubstatusEnum = "ship_failed"
)

// 可用的操作和货件信息包括：
type V3FbsPostingDetailAvailableActionsEnum string
const (
	V3FbsPostingDetailAvailableActionsEnum_Arbitration V3FbsPostingDetailAvailableActionsEnum = "arbitration"
	V3FbsPostingDetailAvailableActionsEnum_AwaitingDelivery V3FbsPostingDetailAvailableActionsEnum = "awaiting_delivery"
	V3FbsPostingDetailAvailableActionsEnum_CanCreateChat V3FbsPostingDetailAvailableActionsEnum = "can_create_chat"
	V3FbsPostingDetailAvailableActionsEnum_Cancel V3FbsPostingDetailAvailableActionsEnum = "cancel"
	V3FbsPostingDetailAvailableActionsEnum_ClickTrackNumber V3FbsPostingDetailAvailableActionsEnum = "click_track_number"
	V3FbsPostingDetailAvailableActionsEnum_CustomerPhoneAvailable V3FbsPostingDetailAvailableActionsEnum = "customer_phone_available"
	V3FbsPostingDetailAvailableActionsEnum_HasWeightProducts V3FbsPostingDetailAvailableActionsEnum = "has_weight_products"
	V3FbsPostingDetailAvailableActionsEnum_HideRegionAndCity V3FbsPostingDetailAvailableActionsEnum = "hide_region_and_city"
	V3FbsPostingDetailAvailableActionsEnum_InvoiceGet V3FbsPostingDetailAvailableActionsEnum = "invoice_get"
	V3FbsPostingDetailAvailableActionsEnum_InvoiceSend V3FbsPostingDetailAvailableActionsEnum = "invoice_send"
	V3FbsPostingDetailAvailableActionsEnum_InvoiceUpdate V3FbsPostingDetailAvailableActionsEnum = "invoice_update"
	V3FbsPostingDetailAvailableActionsEnum_LabelDownloadBig V3FbsPostingDetailAvailableActionsEnum = "label_download_big"
	V3FbsPostingDetailAvailableActionsEnum_LabelDownloadSmall V3FbsPostingDetailAvailableActionsEnum = "label_download_small"
	V3FbsPostingDetailAvailableActionsEnum_LabelDownload V3FbsPostingDetailAvailableActionsEnum = "label_download"
	V3FbsPostingDetailAvailableActionsEnum_NonIntDelivered V3FbsPostingDetailAvailableActionsEnum = "non_int_delivered"
	V3FbsPostingDetailAvailableActionsEnum_NonIntDelivering V3FbsPostingDetailAvailableActionsEnum = "non_int_delivering"
	V3FbsPostingDetailAvailableActionsEnum_NonIntLastMile V3FbsPostingDetailAvailableActionsEnum = "non_int_last_mile"
	V3FbsPostingDetailAvailableActionsEnum_ProductCancel V3FbsPostingDetailAvailableActionsEnum = "product_cancel"
	V3FbsPostingDetailAvailableActionsEnum_SetCutoff V3FbsPostingDetailAvailableActionsEnum = "set_cutoff"
	V3FbsPostingDetailAvailableActionsEnum_SetTimeslot V3FbsPostingDetailAvailableActionsEnum = "set_timeslot"
	V3FbsPostingDetailAvailableActionsEnum_SetTrackNumber V3FbsPostingDetailAvailableActionsEnum = "set_track_number"
	V3FbsPostingDetailAvailableActionsEnum_ShipAsyncInProcess V3FbsPostingDetailAvailableActionsEnum = "ship_async_in_process"
	V3FbsPostingDetailAvailableActionsEnum_ShipAsyncRetry V3FbsPostingDetailAvailableActionsEnum = "ship_async_retry"
	V3FbsPostingDetailAvailableActionsEnum_ShipAsync V3FbsPostingDetailAvailableActionsEnum = "ship_async"
	V3FbsPostingDetailAvailableActionsEnum_ShipWithAdditionalInfo V3FbsPostingDetailAvailableActionsEnum = "ship_with_additional_info"
	V3FbsPostingDetailAvailableActionsEnum_Ship V3FbsPostingDetailAvailableActionsEnum = "ship"
	V3FbsPostingDetailAvailableActionsEnum_UpdateCis V3FbsPostingDetailAvailableActionsEnum = "update_cis"
)

// 货运状态:
type V3FbsPostingDetailStatusEnum string
const (
	V3FbsPostingDetailStatusEnum_AcceptanceInProgress V3FbsPostingDetailStatusEnum = "acceptance_in_progress"
	V3FbsPostingDetailStatusEnum_Arbitration V3FbsPostingDetailStatusEnum = "arbitration"
	V3FbsPostingDetailStatusEnum_AwaitingApprove V3FbsPostingDetailStatusEnum = "awaiting_approve"
	V3FbsPostingDetailStatusEnum_AwaitingDeliver V3FbsPostingDetailStatusEnum = "awaiting_deliver"
	V3FbsPostingDetailStatusEnum_AwaitingPackaging V3FbsPostingDetailStatusEnum = "awaiting_packaging"
	V3FbsPostingDetailStatusEnum_AwaitingRegistration V3FbsPostingDetailStatusEnum = "awaiting_registration"
	V3FbsPostingDetailStatusEnum_AwaitingVerification V3FbsPostingDetailStatusEnum = "awaiting_verification"
	V3FbsPostingDetailStatusEnum_Cancelled V3FbsPostingDetailStatusEnum = "cancelled"
	V3FbsPostingDetailStatusEnum_CancelledFromSplitPending V3FbsPostingDetailStatusEnum = "cancelled_from_split_pending"
	V3FbsPostingDetailStatusEnum_ClientArbitration V3FbsPostingDetailStatusEnum = "client_arbitration"
	V3FbsPostingDetailStatusEnum_Delivered V3FbsPostingDetailStatusEnum = "delivered"
	V3FbsPostingDetailStatusEnum_Delivering V3FbsPostingDetailStatusEnum = "delivering"
	V3FbsPostingDetailStatusEnum_DriverPickup V3FbsPostingDetailStatusEnum = "driver_pickup"
	V3FbsPostingDetailStatusEnum_NotAccepted V3FbsPostingDetailStatusEnum = "not_accepted"
)

type V3FbsPostingDetail struct {
	DeliveringDate string `json:"delivering_date"` // 货件交付物流的时间。
	OrderID int64 `json:"order_id"` // 货件所属的订单ID。
	OrderNumber string `json:"order_number"` // 货件所属的订单号。
	PostingNumber string `json:"posting_number"` // 货件号。
	Products []V3PostingProductDetail `json:"products"` // 货物装运的数组。
	DeliveryMethod V3DeliveryMethod `json:"delivery_method"`
	IsExpress bool `json:"is_express"` // 如果使用了快速物流Ozon Express —— `true`。
	LegalInfo V2FboSinglePostingLegalInfo `json:"legal_info"`
	Requirements V3FbsPostingRequirementsV3 `json:"requirements"`
	DeliveryPrice string `json:"delivery_price"` // 物流价格。
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号。
	AnalyticsData V3FbsPostingAnalyticsData `json:"analytics_data"`
	RelatedPostings V3FbsPostingDetailRelatedPostings `json:"related_postings"`
	TPLIntegrationType V3FbsPostingDetailTPLIntegrationTypeEnum `json:"tpl_integration_type"` // 快递服务集成类型： - `ozon` —— 通过Ozon物流的快递。 - `aggregator` —— 外部服务快递，Ozon注册订单。 - `3pl_tracking` —— 外部服务快递，卖家注册订单。 - `non_integrat...
	Tariffication V3FbsTariffication `json:"tariffication"`
	TarifficationSteps []PostingV4PostingFbsListResponsePostingsTarifficationStep `json:"tariffication_steps"` // 计费阶段。
	Barcodes V3Barcodes `json:"barcodes"`
	Courier FbsPostingDetailCourier `json:"courier"`
	FactDeliveryDate string `json:"fact_delivery_date"` // 货件实际转移配送的日期。
	ParentPostingNumber string `json:"parent_posting_number"` // 母件编号，由该母件拆分出了该货件。
	PreviousSubstatus V3FbsPostingDetailPreviousSubstatusEnum `json:"previous_substatus"` // 货件的前一个子状态。可能的取值： - `posting_acceptance_in_progress` —— 正在验收， - `posting_in_arbitration` —— 仲裁， - `posting_created` —— 已创...
	AdditionalData []V3AdditionalDataItem `json:"additional_data"`
	FinancialData V3PostingFinancialData `json:"financial_data"`
	InProcessAt string `json:"in_process_at"` // 开始处理货件的日期和时间。
	ProductExemplars V3FbsPostingProductExemplarsV3 `json:"product_exemplars"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"` // 不逾期发运日期和时间。
	Substatus V3FbsPostingDetailSubstatusEnum `json:"substatus"` // 发货子状态： - `posting_acceptance_in_progress` —— 正在验收， - `posting_in_arbitration` —— 仲裁， - `posting_created` —— 已创建， - `post...
	AvailableActions []string `json:"available_actions"` // 可用的操作和货件信息包括： - `arbitration` — 提出争议； - `awaiting_delivery` — 转为“等待发运”状态； - `can_create_chat` — 与买家开启聊天； - `cancel` — 取消...
	Cancellation V3Cancellation `json:"cancellation"`
	Customer V3Customer `json:"customer"`
	Optional V3FbsPostingDetailOptional `json:"optional"`
	ProviderStatus string `json:"provider_status"` // 快递服务状态。
	ShipmentDate string `json:"shipment_date"` // 必须收取货件的日期和时间。 超出该时间后将适用新费率，相关信息请查看字段 `tariffication`。
	Status V3FbsPostingDetailStatusEnum `json:"status"` // 货运状态: - `acceptance_in_progress` —— 正在验收， - `arbitration` —— 仲裁， - `awaiting_approve` —— 等待确认， - `awaiting_deliver` —— 等...
	Addressee V3Addressee `json:"addressee"`
}

type PostingV4PostingFbsListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	HasNext bool `json:"has_next"` // 若响应中未返回全部货件，则为`true`。
	Postings []PostingV4PostingFbsListResponsePostings `json:"postings"` // 货件列表。
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"` // 发运识别符。
}

// 货件最后一次更改过状态的时期。
type PostinglistV3status struct {
	From string `json:"from"` // 时期开始日期。
	To string `json:"to"` // 时期结束日期。
}

// 过滤器。
// 从合作伙伴仓库（FBP）发货时的货件筛选器：
type Postingv3GetFbsPostingListRequestFilterFbpFilterEnum string
const (
	Postingv3GetFbsPostingListRequestFilterFbpFilterEnum_ALL Postingv3GetFbsPostingListRequestFilterFbpFilterEnum = "ALL"
	Postingv3GetFbsPostingListRequestFilterFbpFilterEnum_ONLY Postingv3GetFbsPostingListRequestFilterFbpFilterEnum = "ONLY"
	Postingv3GetFbsPostingListRequestFilterFbpFilterEnum_WITHOUT Postingv3GetFbsPostingListRequestFilterFbpFilterEnum = "WITHOUT"
)

// 货件运输状态：
type Postingv3GetFbsPostingListRequestFilterStatusEnum string
const (
	Postingv3GetFbsPostingListRequestFilterStatusEnum_AwaitingRegistration Postingv3GetFbsPostingListRequestFilterStatusEnum = "awaiting_registration"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_AcceptanceInProgress Postingv3GetFbsPostingListRequestFilterStatusEnum = "acceptance_in_progress"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_AwaitingApprove Postingv3GetFbsPostingListRequestFilterStatusEnum = "awaiting_approve"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_AwaitingPackaging Postingv3GetFbsPostingListRequestFilterStatusEnum = "awaiting_packaging"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_AwaitingDeliver Postingv3GetFbsPostingListRequestFilterStatusEnum = "awaiting_deliver"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_Arbitration Postingv3GetFbsPostingListRequestFilterStatusEnum = "arbitration"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_ClientArbitration Postingv3GetFbsPostingListRequestFilterStatusEnum = "client_arbitration"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_Delivering Postingv3GetFbsPostingListRequestFilterStatusEnum = "delivering"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_DriverPickup Postingv3GetFbsPostingListRequestFilterStatusEnum = "driver_pickup"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_Delivered Postingv3GetFbsPostingListRequestFilterStatusEnum = "delivered"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_Cancelled Postingv3GetFbsPostingListRequestFilterStatusEnum = "cancelled"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_NotAccepted Postingv3GetFbsPostingListRequestFilterStatusEnum = "not_accepted"
	Postingv3GetFbsPostingListRequestFilterStatusEnum_SentBySeller Postingv3GetFbsPostingListRequestFilterStatusEnum = "sent_by_seller"
)

type Postingv3GetFbsPostingListRequestFilter struct {
	DeliveryMethodID []int64 `json:"delivery_method_id"` // 快递方式ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	OrderID int64 `json:"order_id"` // 订单ID。
	ProviderID []int64 `json:"provider_id"` // 快递服务ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	WarehouseID []string `json:"warehouse_id"` // 仓库ID。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	LastChangedStatusDate PostinglistV3status `json:"last_changed_status_date"`
	FbpFilter Postingv3GetFbsPostingListRequestFilterFbpFilterEnum `json:"fbpFilter"` // 从合作伙伴仓库（FBP）发货时的货件筛选器： - `ALL` — 响应中将显示所有符合其他筛选器条件的货件； - `ONLY` — 仅显示FBP货件； - `WITHOUT` — 显示除FBP外的所有货件。 默认值为 `ALL`。
	Since string `json:"since"` // 应收到货件清单时间段的开始日期。 UTC模式: ГГГГ-ММ-ДДTЧЧ:ММ:ССZ. 例子: 2019-08-24T14:15:22Z.
	To string `json:"to"` // 应收到货件清单时间段的结束日期。 UTC模式： ГГГГ-ММ-ДДTЧЧ:ММ:ССZ. 例子： 2019-08-24T14:15:22Z.
	Status Postingv3GetFbsPostingListRequestFilterStatusEnum `json:"status"` // 货件运输状态： - `awaiting_registration` — 等待注册， - `acceptance_in_progress` — 正在验收， - `awaiting_approve` — 等待确认， - `awaiting_pa...
}

type V1CarriageCancelResponse struct {
	Error string `json:"error"` // 错误描述。
	CarriageStatus string `json:"carriage_status"` // 发运状态。
}

type V2PostingFBSActGetProducts struct {
	Quantity int32 `json:"quantity"` // 货件中的商品数量。
	SKU int64 `json:"sku"` // 商品在Ozon系统中的标识符——SKU。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 商品在卖家系统中的标识符——货号。
	Price string `json:"price"` // 商品价格。
}

type V2PostingFBSActGetPostingsResult struct {
	PostingNumber string `json:"posting_number"` // 货件编号。
	Status string `json:"status"` // 货件状态。
	SellerError string `json:"seller_error"` // 错误代码说明。
	UpdatedAt string `json:"updated_at"` // 货件记录的更新日期和时间。
	CreatedAt string `json:"created_at"` // 货件记录的创建日期和时间。
	Products []V2PostingFBSActGetProducts `json:"products"` // 货件中商品列表。
	ID int64 `json:"id"` // 单据标识符。
	MultiBoxQty int32 `json:"multi_box_qty"` // 商品包装所用箱数。
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []V2PostingFBSActGetPostingsResult `json:"result"` // 货件信息。
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber string `json:"posting_number"` // 货运号。
	ProductID int64 `json:"product_id"` // 产品ID。
	CountryIsoCode string `json:"country_iso_code"` // 根据标准ISO_3166-1添加的国家的两个字母代码 制造国家列表及其ISO代码可以使用该方法获得[/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountr...
}

// 商品佣金。
type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission struct {
	Amount float64 `json:"amount"` // 商品佣金金额。
	Currency string `json:"currency"` // 计算佣金所使用的货币代码。
	Percent int64 `json:"percent"` // 佣金比例。
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts struct {
	Payout float64 `json:"payout"` // 卖家应收款。
	ProductID int64 `json:"product_id"` // Ozon系统中的商品标识符，即SKU。
	Quantity int64 `json:"quantity"` // 货件中的商品数量。
	TotalDiscountPercent float64 `json:"total_discount_percent"` // 折扣百分比。
	TotalDiscountValue float64 `json:"total_discount_value"` // 折扣金额。
	Actions []string `json:"actions"` // 促销活动列表。
	CustomerPrice MoneyPostingMoney `json:"customer_price"`
	OldPrice float64 `json:"old_price"` // 折扣前的价格。商品卡上会以划线价显示。
	Price float64 `json:"price"` // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	Commission PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission `json:"commission"`
}

// 配送方法信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod struct {
	ID int64 `json:"id"` // 配送方式标识符。
	Name string `json:"name"` // 配送方法名称。
	TPLProvider string `json:"tpl_provider"` // 配送服务。
	TPLProviderID int64 `json:"tpl_provider_id"` // 配送服务标识符。
	Warehouse string `json:"warehouse"` // 仓库名称。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

// 带有附加特征的商品列表。
type PostingV4PostingFbsUnfulfilledListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []string `json:"products_with_possible_mandatory_mark"` // 可能需要标志的商品列表。
}

// 是否可以取消。
type CarriageCarriageGetResponseCancelAvailability struct {
	IsCancelAvailable bool `json:"is_cancel_available"` // `true`, 如果运输可以取消。
	Reason string `json:"reason"` // 运输无法取消的原因。
}

// 配送地址。
type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress struct {
	Comment string `json:"comment"` // 订单备注。
	Country string `json:"country"` // 配送国家。
	District string `json:"district"` // 配送区域。
	Latitude float64 `json:"latitude"` // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
	Region string `json:"region"` // 配送地区。
	AddressTail string `json:"address_tail"` // 文本格式地址。
	City string `json:"city"` // 配送城市。
	ProviderPvzCode string `json:"provider_pvz_code"` // 3PL服务商取货点代码。
	PvzCode int64 `json:"pvz_code"` // 取货点代码。
	ZipCode string `json:"zip_code"` // 收件人邮政编码。
}

// 买家信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer struct {
	Address PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress `json:"address"`
	CustomerEmail string `json:"customer_email"` // 买家电子邮箱。
	CustomerID int64 `json:"customer_id"` // 买家标识符。
	Name string `json:"name"` // 买家姓名。
	Phone string `json:"phone"` // 买家的替换联系电话。
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []string `json:"posting_number"` // 货运ID。
}

// 运输的可用操作：
type CarriageCarriageGetResponseAvailableActionsEnum string
const (
	CarriageCarriageGetResponseAvailableActionsEnum_GetShippingList CarriageCarriageGetResponseAvailableActionsEnum = "get_shipping_list"
	CarriageCarriageGetResponseAvailableActionsEnum_GetActOfAcceptance CarriageCarriageGetResponseAvailableActionsEnum = "get_act_of_acceptance"
	CarriageCarriageGetResponseAvailableActionsEnum_GetWaybill CarriageCarriageGetResponseAvailableActionsEnum = "get_waybill"
	CarriageCarriageGetResponseAvailableActionsEnum_SetArrivalPasses CarriageCarriageGetResponseAvailableActionsEnum = "set_arrival_passes"
)

type CarriageCarriageGetResponse struct {
	Status string `json:"status"` // 运输状态。
	TPLProviderID int64 `json:"tpl_provider_id"` // 配送服务商标识符。
	ActType string `json:"act_type"` // 交接单类型。针对FBS卖家。
	CreatedAt string `json:"created_at"` // 运输创建日期。
	IsContainerLabelPrinted bool `json:"is_container_label_printed"` // `true`, 如果您已经打印了货位标签。
	PartialNum int64 `json:"partial_num"` // 部分运输序列号。
	AvailableActions []string `json:"available_actions"` // 运输的可用操作： - `get_shipping_list`——获取发运清单； - `get_act_of_acceptance`——获取验收证明书； - `get_waybill`——获取 PDF 格式的货单； - `set_arriva...
	CompanyID int64 `json:"company_id"` // 卖家标识符。
	ContainersCount int32 `json:"containers_count"` // 货位数量。
	FirstMileType string `json:"first_mile_type"` // 头程物流类型。
	RetryCount int32 `json:"retry_count"` // 运输创建重复尝试数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	CancelAvailability CarriageCarriageGetResponseCancelAvailability `json:"cancel_availability"`
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
	IntegrationType string `json:"integration_type"` // 运输类型。
	IsPartial bool `json:"is_partial"` // `true`, 如果是部分运输。
	UpdatedAt string `json:"updated_at"` // 运输信息最后一次更新日期。
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 为运输生成的通行证标识符列表。
	DeliveryMethodID int64 `json:"delivery_method_id"` // 物流方式标识符。
	DepartureDate string `json:"departure_date"` // 运输完成日期。
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"` // `true`, 如果有未能进行运输，但需要发运的货件。
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"` // 货运条形码。可以使用以下方法获取： [/v3/posting/fbs/get](#operation/PostingAPI_GetFbsPostingV3)、[/v3/posting/fbs/list](#operation/Posting...
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"` // 按行过滤。
}

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []string `json:"posting_number"` // 货运ID。一次请求中的最大数量——100。
}

type V1AssemblyCarriagePostingListResponsePostingProduct struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符——货号。
	PictureURL string `json:"picture_url"` // 商品图片链接。
	ProductName string `json:"product_name"` // 商品名称。
	Quantity int64 `json:"quantity"` // 商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符——SKU。
}

type V1AssemblyCarriageProductListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter V1AssemblyCarriageProductListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 每页显示的数量。
}

type ProductsPostings struct {
	PostingNumber string `json:"posting_number"` // 货件编号。
	Quantity int32 `json:"quantity"` // 货件中的商品数量。
}

type AssemblyFbsProductListResponseProducts struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符——货号。
	PictureURL string `json:"picture_url"` // 商品图片链接。
	Postings []ProductsPostings `json:"postings"` // 货件列表。
	ProductName string `json:"product_name"` // 商品名称。
	Quantity int32 `json:"quantity"` // 商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符——SKU。
}

type V1AssemblyFbsProductListResponse struct {
	HasNext bool `json:"has_next"` // 响应中是否包含全部商品： - `true`——请使用新的 `offset`值重新请求以获取剩余数据； - `false`——响应中已包含所有值。
	Products []AssemblyFbsProductListResponseProducts `json:"products"` // 商品列表。
	ProductsCount int32 `json:"products_count"` // 商品数量。
}

// 买家的法务信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo struct {
	CompanyName string `json:"company_name"` // 公司名称。
	Inn string `json:"inn"` // 税号。
	Kpp string `json:"kpp"` // 纳税人登记原因代码。
}

// 货件条形码。
type FbsPostingBarcodes struct {
	UpperBarcode string `json:"upper_barcode"` // 货件标签的上条形码。
	LowerBarcode string `json:"lower_barcode"` // 货件标签的下条形码。
}

// 请求结果。
type V2FbsPosting struct {
	Barcodes FbsPostingBarcodes `json:"barcodes"`
	CancelReasonID int64 `json:"cancel_reason_id"` // 取消装运原因ID。
	CreatedAt string `json:"created_at"` // 创建装运日期和时间。
	InProcessAt string `json:"in_process_at"` // 开始处理货件的日期和时间。
	OrderID int64 `json:"order_id"` // 货运所属订单ID。
	OrderNumber string `json:"order_number"` // 货运所属的订单号。
	PostingNumber string `json:"posting_number"` // 货运号。
	Products []V2FbsPostingProduct `json:"products"` // 货运商品列表。
	ShipmentDate string `json:"shipment_date"` // 必须收取货件的日期和时间。 如果在此日期之前未完成配货，则货运自动取消。
	Status string `json:"status"` // 货运状态。
}

type V1CarriageApproveRequest struct {
	CarriageID int64 `json:"carriage_id"` // 发运标识符。
	ContainersCount int32 `json:"containers_count"` // 货位数量。 如果您已开通信任验收，并按货位发运订单，请使用该参数。如果您未开通信任验收，请跳过该参数。
}

type V1AssemblyFbsProductListRequest struct {
	SortDir V1AssemblyFbsProductListRequestSortDirEnum `json:"sort_dir"`
	Filter V1AssemblyFbsProductListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 每页显示的数量。
	Offset int64 `json:"offset"` // 在响应中将被跳过的项目数量。例如，如果 `offset = 10`，则响应将从第 11 个找到的项目开始。
}

// 收件人联系方式。
type PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee struct {
	Name string `json:"name"` // 收件人姓名。
}

type V1CarriageCreateRequest struct {
	AllBlrTraceable bool `json:"all_blr_traceable"` // `true`，表示需要创建包含可追溯商品的发运。
	DeliveryMethodID int64 `json:"delivery_method_id"` // 配送方式标识符。
	DepartureDate string `json:"departure_date"` // 发运日期。默认值为当前日期。
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
}

// 发运的计费信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication struct {
	CurrentTariffCharge MoneyMoneyCurrentTariffCharge `json:"current_tariff_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"` // 计费百分比。
	CurrentTariffType string `json:"current_tariff_type"` // 计费类型——折扣或加价。
	NextTariffStartsAt string `json:"next_tariff_starts_at"` // 新费率开始生效的日期和时间。
	CurrentTariffMinCharge MoneyMoneyCurrentTariffMinCharge `json:"current_tariff_min_charge"`
	NextTariffCharge MoneyMoneyNextTariffCharge `json:"next_tariff_charge"`
	NextTariffMinCharge MoneyMoneyNextTariffMinCharge `json:"next_tariff_min_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"` // 在`next_tariff_starts_at`参数指定时间之后，该货件将按此百分比计费。
	NextTariffType string `json:"next_tariff_type"` // 在`next_tariff_starts_at`参数指定时间之后的计费类型——折扣或加价。
}

type PostingCancelFbsPostingRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"` // 取消运输的原因ID。
	CancelReasonMessage string `json:"cancel_reason_message"` // 关于取消的附加信息。如果`cancel_reason_id = 402`，参数是必须的。
	PostingNumber string `json:"posting_number"` // 货件ID。
}

// 需要添加到响应中的附加字段。
type PostingV4PostingFbsListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"` // 若为`true`，则在响应中添加分析数据。
	Barcodes bool `json:"barcodes"` // `true`，表示要在响应中添加货件条形码。
	FinancialData bool `json:"financial_data"` // 若为`true`，则在响应中添加财务数据。
	LegalInfo bool `json:"legal_info"` // 若为`true`，则在响应中添加法务信息。
}

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"` // `true`表示已设置新日期。
}

type FbsPostingTrackingNumberSetRequestTrackingNumber struct {
	PostingNumber string `json:"posting_number"` // 货件ID。
	TrackingNumber string `json:"tracking_number"` // 货件追踪号。
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []FbsPostingTrackingNumberSetRequestTrackingNumber `json:"tracking_numbers"` // 具有成对货运ID的数据 - 追踪号。
}

// 需要添加到响应中的附加字段。
type Postingv3FbsPostingWithParamsExamplars struct {
	Translit bool `json:"translit"` // 完成返回值的拼写。
	AnalyticsData bool `json:"analytics_data"` // 将分析数据添加到响应中。
	Barcodes bool `json:"barcodes"` // 将货件条形码添加到响应中。
	FinancialData bool `json:"financial_data"` // 将财务数据添加到响应中。
	LegalInfo bool `json:"legal_info"` // 将法律信息添加到响应中。
	ProductExemplars bool `json:"product_exemplars"` // 将有关产品及其份数的数据添加到响应中。
	RelatedPostings bool `json:"related_postings"` // 将相关货件数量添加到响应中。 相关货件是在组装期间将母快递拆分的快递。
}

type Postingv3GetFbsPostingRequest struct {
	PostingNumber string `json:"posting_number"` // 货件ID。
	With Postingv3FbsPostingWithParamsExamplars `json:"with"`
}

// 请求结果。
type Postingv3GetFbsPostingUnfulfilledListResponseResult struct {
	Count int64 `json:"count"` // 在响应中的元素计数器。
	Postings []V3FbsPosting `json:"postings"` // 货件清单和每个货物的详细信息。
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result Postingv3GetFbsPostingUnfulfilledListResponseResult `json:"result"`
}

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type PostingV4PostingFbsListRequestSortDirEnum string

type V3GetFbsPostingResponseV3 struct {
	Result V3FbsPostingDetail `json:"result"`
}

type PostingBooleanResponse struct {
	Result bool `json:"result"` // 处理请求的结果。 如果请求执行时无误，则为“true”。
}

// 货件状态最后一次发生变更的时间段。
type PostingV4PostingFbsListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"` // 周期开始日期。
	To string `json:"to"` // 周期结束日期。
}

// 分类方向：
type Postingv3GetFbsPostingListRequestDirEnum string
const (
	Postingv3GetFbsPostingListRequestDirEnum_Asc Postingv3GetFbsPostingListRequestDirEnum = "asc"
	Postingv3GetFbsPostingListRequestDirEnum_Desc Postingv3GetFbsPostingListRequestDirEnum = "desc"
)

type Postingv3GetFbsPostingListRequest struct {
	Limit int64 `json:"limit"` // 响应中值的数量： - 最大值 — 1000, - 最小值 — 1。
	Offset int64 `json:"offset"` // 将在响应中跳过的元素数。 例如，如果“offset=10”，那么响应将从找到的第11个元素开始。
	With Postingv3FbsPostingWithParams `json:"with"`
	Dir Postingv3GetFbsPostingListRequestDirEnum `json:"dir"` // 分类方向： - `asc` — 从小到大， - `desc` — 从大到小。
	Filter Postingv3GetFbsPostingListRequestFilter `json:"filter"`
}

type V1AssemblyCarriagePostingListResponsePosting struct {
	AssemblyCode string `json:"assembly_code"` // 拣货单代码。
	CanPrintLabel bool `json:"can_print_label"` // `true`，前提是可以打印标签。
	PostingNumber string `json:"posting_number"` // 货件编号。
	Products []V1AssemblyCarriagePostingListResponsePostingProduct `json:"products"` // 商品列表。
}

type V1AssemblyCarriagePostingListResponse struct {
	CanPrintMassLabel bool `json:"can_print_mass_label"` // `true`，前提是可以批量打印标签。
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Postings []V1AssemblyCarriagePostingListResponsePosting `json:"postings"` // 货件列表。
}

// 筛选器。
// 货件状态：
type PostingV4PostingFbsListRequestFilterStatusesEnum string
const (
	PostingV4PostingFbsListRequestFilterStatusesEnum_AwaitingRegistration PostingV4PostingFbsListRequestFilterStatusesEnum = "awaiting_registration"
	PostingV4PostingFbsListRequestFilterStatusesEnum_AcceptanceInProgress PostingV4PostingFbsListRequestFilterStatusesEnum = "acceptance_in_progress"
	PostingV4PostingFbsListRequestFilterStatusesEnum_AwaitingApprove PostingV4PostingFbsListRequestFilterStatusesEnum = "awaiting_approve"
	PostingV4PostingFbsListRequestFilterStatusesEnum_AwaitingPackaging PostingV4PostingFbsListRequestFilterStatusesEnum = "awaiting_packaging"
	PostingV4PostingFbsListRequestFilterStatusesEnum_AwaitingDeliver PostingV4PostingFbsListRequestFilterStatusesEnum = "awaiting_deliver"
	PostingV4PostingFbsListRequestFilterStatusesEnum_Arbitration PostingV4PostingFbsListRequestFilterStatusesEnum = "arbitration"
	PostingV4PostingFbsListRequestFilterStatusesEnum_ClientArbitration PostingV4PostingFbsListRequestFilterStatusesEnum = "client_arbitration"
	PostingV4PostingFbsListRequestFilterStatusesEnum_Delivering PostingV4PostingFbsListRequestFilterStatusesEnum = "delivering"
	PostingV4PostingFbsListRequestFilterStatusesEnum_DriverPickup PostingV4PostingFbsListRequestFilterStatusesEnum = "driver_pickup"
	PostingV4PostingFbsListRequestFilterStatusesEnum_Delivered PostingV4PostingFbsListRequestFilterStatusesEnum = "delivered"
	PostingV4PostingFbsListRequestFilterStatusesEnum_Cancelled PostingV4PostingFbsListRequestFilterStatusesEnum = "cancelled"
	PostingV4PostingFbsListRequestFilterStatusesEnum_NotAccepted PostingV4PostingFbsListRequestFilterStatusesEnum = "not_accepted"
	PostingV4PostingFbsListRequestFilterStatusesEnum_SentBySeller PostingV4PostingFbsListRequestFilterStatusesEnum = "sent_by_seller"
)

type PostingV4PostingFbsListRequestFilter struct {
	DeliveryMethodIds []string `json:"delivery_method_ids"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	OrderID int64 `json:"order_id"` // 订单标识符。
	OrderNumbers []string `json:"order_numbers"` // 货件所属订单的订单号。
	Statuses []string `json:"statuses"` // 货件状态： - `awaiting_registration`——等待登记； - `acceptance_in_progress`——接收 中； - `awaiting_approve`——等待 确认； - `awaiting_packag...
	To string `json:"to"` // 需要获取货件列表的周期结束日期。
	IsBlrTraceable bool `json:"is_blr_traceable"` // `true`，表示商品可追溯。
	LastChangedStatusDate PostingV4PostingFbsListRequestFilterLastChangedStatusDate `json:"last_changed_status_date"`
	ProviderIds []string `json:"provider_ids"` // 配送服务标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	Since string `json:"since"` // 需要获取货件列表的周期开始日期。
	WarehouseIds []string `json:"warehouse_ids"` // 仓库标识符。可通过方法[/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
}

type PostingV4PostingFbsListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter PostingV4PostingFbsListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortDir PostingV4PostingFbsListRequestSortDirEnum `json:"sort_dir"`
	Translit bool `json:"translit"` // 若为`true`，则启用将地址从西里尔字母转写为拉丁字母。
	With PostingV4PostingFbsListRequestWith `json:"with"`
}

type Fbsv4FbsPostingShipV4Response struct {
	AdditionalData []FbsPostingShipV4ResponseShipAdditionalData `json:"additional_data"` // 与发货有关的附加信息。
	Result []string `json:"result"` // 货运装配结果。
}

// 货运信息。
type V2FbsPostingResponse struct {
	Result V2FbsPosting `json:"result"`
}

type V1AssemblyFbsPostingListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Cutoff string `json:"cutoff"` // 卖家需在此时间前完成订单备货。
	Postings []V1AssemblyFbsPostingListResponsePosting `json:"postings"` // 货件列表。
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"` // 新发运日期。
	PostingNumber string `json:"posting_number"` // 货件编号。
}

// 取消信息。
// 取消发起方：
type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum_V PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum = "卖家"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum_V_1 PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum = "客户"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum_V_2 PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum = "买家"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum_Ozon PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum = "Ozon"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum_V_3 PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum = "系统"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum_V_4 PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum = "配送服务商"
)

// 取消类型：
type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum_Seller PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum = "seller"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum_Client PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum = "client"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum_Customer PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum = "customer"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum_Ozon PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum = "ozon"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum_System PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum = "system"
	PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum_Delivery PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum = "delivery"
)

type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation struct {
	AffectCancellationRating bool `json:"affect_cancellation_rating"` // `true`，表示取消会影响卖家评级。
	CancelReason string `json:"cancel_reason"` // 取消原因。
	CancelReasonID int64 `json:"cancel_reason_id"` // 货件取消原因标识符。
	CancellationInitiator PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationInitiatorEnum `json:"cancellation_initiator"` // 取消发起方： - `卖家`， - `客户`， - `买家`， - `Ozon`， - `系统`， - `配送服务商`。
	CancellationType PostingV4PostingFbsUnfulfilledListResponsePostingsCancellationCancellationTypeEnum `json:"cancellation_type"` // 取消类型： - `seller`——由卖家取消； - `client`或`customer`——由买家取消； - `ozon`——由Ozon取消； - `system`——由系统取消； - `delivery`——由配送服务取消。
	CancelledAfterShip bool `json:"cancelled_after_ship"` // `true`，表示取消发生在货件完成备货之后。
}

// 商品价格、折扣金额、卖家应收款和佣金信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData struct {
	ClusterTo string `json:"cluster_to"` // 订单配送地区代码。
	Products []PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts `json:"products"` // 订单中的商品列表。
	ClusterFrom string `json:"cluster_from"` // 订单发出地区代码。
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep struct {
	TariffCharge MoneyMoneyCurrentTariffCharge `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"` // 计费阶段结束的日期和时间。该日期后将自动进入下一计费阶段。
	TariffRate float64 `json:"tariff_rate"` // 折扣或附加费用百分比。
	TariffType string `json:"tariff_type"` // 计费类型。
	MinCharge MoneyMoneyCurrentTariffMinCharge `json:"min_charge"`
}

// 需要补充附加信息的商品。 要将货件转入下一状态，请传递： - 制造国； - 货运报关单编号； - 商品批次登记号； - “诚实标志”标记； - 其他标记； - 重量。
type PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements struct {
	ProductsRequiringRnpt []string `json:"products_requiring_rnpt"` // 需要传递商品批次登记号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringWeight []string `json:"products_requiring_weight"` // 需要传递重量的商品列表。
	ProductsRequiringChangeCountry []string `json:"products_requiring_change_country"` // 需要更改制造国的商品标识符（SKU）列表。要更改制造国，请使用方法[/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountryProductFbsPosti...
	ProductsRequiringCountry []string `json:"products_requiring_country"` // 需要传递制造国信息的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v2/posting/fbs/product/country/set](#operation/PostingAPI_SetCountryProductFbsPost...
	ProductsRequiringGTD []string `json:"products_requiring_gtd"` // 需要传递货运报关单编号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringImei []string `json:"products_requiring_imei"` // 需要传递IMEI的商品标识符列表。
	ProductsRequiringJwUin []string `json:"products_requiring_jw_uin"` // 需要传递珠宝制品唯一标识码的商品列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExemplarSe...
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"` // С需要传递“诚实标志”标识的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductEx...
}

// 与配送服务的集成类型：
type PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum_Ozon PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum = "ozon"
	PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum_V3plTracking PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum = "3pl_tracking"
	PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum_NonIntegrated PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum = "non_integrated"
	PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum_Aggregator PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum = "aggregator"
	PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum_Hybryd PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum = "hybryd"
)

// 装卸服务代码：
type PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum_Lift PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum = "lift"
	PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum_Stairs PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum = "stairs"
	PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum_None PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum = "none"
	PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum_DeliveryDefault PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum = "delivery_default"
)

// 货件状态：
type PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_AcceptanceInProgress PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "acceptance_in_progress"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_Arbitration PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "arbitration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_AwaitingApprove PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "awaiting_approve"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_AwaitingDeliver PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "awaiting_deliver"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_AwaitingPackaging PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "awaiting_packaging"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_AwaitingRegistration PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "awaiting_registration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_AwaitingVerification PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "awaiting_verification"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_Cancelled PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "cancelled"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_CancelledFromSplitPending PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "cancelled_from_split_pending"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_ClientArbitration PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "client_arbitration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_Delivering PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "delivering"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_DriverPickup PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "driver_pickup"
	PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum_NotAccepted PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum = "not_accepted"
)

// 货件子状态：
type PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingAcceptanceInProgress PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_acceptance_in_progress"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingInArbitration PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_in_arbitration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingCreated PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_created"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingInCarriage PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_in_carriage"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingNotInCarriage PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_not_in_carriage"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingRegistered PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_registered"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingTransferringToDelivery PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_transferring_to_delivery"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_StatusAwaitingDeliver PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "status=awaiting_deliver"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingAwaitingPassportData PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_awaiting_passport_data"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingAwaitingRegistration PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_awaiting_registration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingRegistrationError PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_registration_error"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_StatusAwaitingRegistration PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "status=awaiting_registration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingSplitPending PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_split_pending"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingCanceled PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_canceled"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingInClientArbitration PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_in_client_arbitration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingDelivered PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_delivered"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingReceived PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_received"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingConditionallyDelivered PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_conditionally_delivered"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingInCourierService PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_in_courier_service"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingInPickupPoint PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_in_pickup_point"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingOnWayToCity PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_on_way_to_city"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingOnWayToPickupPoint PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_on_way_to_pickup_point"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingReturnedToWarehouse PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_returned_to_warehouse"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingTransferredToCourierService PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_transferred_to_courier_service"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingDriverPickUp PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_driver_pick_up"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_PostingNotInSortCenter PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "posting_not_in_sort_center"
	PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum_ShipFailed PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum = "ship_failed"
)

// 货件的可用操作和信息：
type PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_Arbitration PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "arbitration"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_AwaitingDelivery PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "awaiting_delivery"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_CanCreateChat PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "can_create_chat"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_Cancel PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "cancel"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_ClickTrackNumber PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "click_track_number"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_CustomerPhoneAvailable PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "customer_phone_available"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_HasWeightProducts PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "has_weight_products"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_HideRegionAndCity PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "hide_region_and_city"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_InvoiceGet PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "invoice_get"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_InvoiceSend PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "invoice_send"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_InvoiceUpdate PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "invoice_update"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_LabelDownloadBig PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "label_download_big"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_LabelDownloadSmall PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "label_download_small"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_LabelDownload PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "label_download"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_NonIntDelivered PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "non_int_delivered"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_NonIntDelivering PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "non_int_delivering"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_NonIntLastMile PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "non_int_last_mile"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_ProductCancel PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "product_cancel"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_SetCutoff PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "set_cutoff"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_ShipmentDate PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "shipment_date"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_SetTimeslot PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "set_timeslot"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_SetTrackNumber PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "set_track_number"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_ShipAsyncInProcess PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "ship_async_in_process"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_ShipAsyncRetry PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "ship_async_retry"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_ShipAsync PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "ship_async"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_ShipWithAdditionalInfo PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "ship_with_additional_info"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_Ship PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "ship"
	PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum_UpdateCis PostingV4PostingFbsUnfulfilledListResponsePostingsAvailableActionsEnum = "update_cis"
)

// 配送模式：
type PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum string
const (
	PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum_SDS PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum = "SDS"
	PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum_FBO PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum = "FBO"
	PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum_FBS PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum = "FBS"
	PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum_Crossborder PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum = "Crossborder"
)

type PostingV4PostingFbsUnfulfilledListResponsePostings struct {
	DestinationPlaceName string `json:"destination_place_name"` // 目的地名称。
	ExternalOrder PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder `json:"external_order"`
	IsPresortable bool `json:"is_presortable"` // `true`，表示商品为错误分级的商品。
	ShipmentDate string `json:"shipment_date"` // 需要在此日期和时间之前完成货件备货。系统会显示推荐的发运时间。超过该时间后将开始适用新费率，相关信息可在`tariffication`字段中获取。
	IsClickAndCollect bool `json:"is_click_and_collect"` // `true`，表示货件通过“到店自提”方式配送。
	IsMultibox bool `json:"is_multibox"` // 标记货件中是否包含多箱商品，以及是否需要传递其箱数： - `true`——备货前，请通过方法[/v3/posting/multiboxqty/set](#operation/PostingAPI_PostingMultiBoxQtySetV...
	Barcodes PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes `json:"barcodes"`
	DeliveringDate string `json:"delivering_date"` // 货件转移配送的日期。
	DeliveryMethod PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod `json:"delivery_method"`
	TPLIntegrationType PostingV4PostingFbsUnfulfilledListResponsePostingsTPLIntegrationTypeEnum `json:"tpl_integration_type"` // 与配送服务的集成类型： - `ozon`——由Ozon配送； - `3pl_tracking`——由已集成服务配送； - `non_integrated`——由第三方服务配送； - `aggregator`——通过Ozon合作伙伴配送； -...
	Cancellation PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation `json:"cancellation"`
	OrderID int64 `json:"order_id"` // 该货件所属订单的标识符。
	OrderNumber string `json:"order_number"` // 该货件所属订单的编号。
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"` // 快递员代码校验成功的日期和时间。请通过方法 [/v1/posting/fbs/pick-up-code/verify](#operation/PostingAPI_PostingFBSPickupCodeVerify)校验快递员代码。
	Tariffication PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication `json:"tariffication"`
	DestinationPlaceID int64 `json:"destination_place_id"` // 目的地标识符。
	IsExpress bool `json:"is_express"` // `true`，表示使用了Ozon Express快速配送。
	LegalInfo PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo `json:"legal_info"`
	Optional PostingV4PostingFbsUnfulfilledListResponsePostingsOptional `json:"optional"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"` // 不逾期发运的日期和时间。
	InProcessAt string `json:"in_process_at"` // 货件开始处理的日期和时间。
	MultiBoxQty int32 `json:"multi_box_qty"` // 商品包装所用箱数。
	Products []PostingV4PostingFbsUnfulfilledListResponsePostingsProducts `json:"products"` // 货件中商品列表。
	PrrOption PostingV4PostingFbsUnfulfilledListResponsePostingsPrrOptionEnum `json:"prr_option"` // 装卸服务代码： - `lift`——通过电梯搬运上楼； - `stairs`——通过楼梯搬运上楼； - `none`——买家拒绝该服务，无需将商品搬上楼； - `delivery_default`——配送费用已包含在价格中，根据要约条款需将...
	QuantumID int64 `json:"quantum_id"` // 经济商品标识符。
	Status PostingV4PostingFbsUnfulfilledListResponsePostingsStatusEnum `json:"status"` // 货件状态： - `acceptance_in_progress`——接收中； - `arbitration`——仲裁； - `awaiting_approve`——等待确认； - `awaiting_deliver`——等待发运； - `a...
	Substatus PostingV4PostingFbsUnfulfilledListResponsePostingsSubstatusEnum `json:"substatus"` // 货件子状态： - `posting_acceptance_in_progress`——验收中； - `posting_in_arbitration`——仲裁； - `posting_created`——已创建； - `posting_in_...
	VolumeWeight float64 `json:"volume_weight"` // 商品体积重量。
	PostingNumber string `json:"posting_number"` // 货件编号。
	FinancialData PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData `json:"financial_data"`
	Addressee PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee `json:"addressee"`
	AnalyticsData PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData `json:"analytics_data"`
	AvailableActions []string `json:"available_actions"` // 货件的可用操作和信息： - `arbitration`——提出争议； - `awaiting_delivery`——变更为“等待发运”状态； - `can_create_chat`——与买家发起聊天； - `cancel`——取消货件； -...
	Customer PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer `json:"customer"`
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号码。
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"` // `true`，表示需要填写可追踪属性。
	Requirements PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements `json:"requirements"`
	DeliverySchema PostingV4PostingFbsUnfulfilledListResponsePostingsDeliverySchemaEnum `json:"delivery_schema"` // 配送模式： - `SDS`——统一SKU标识符； - `FBO`——从Ozon仓库销售的商品标识符； - `FBS`——从FBS仓库销售的商品标识符； - `Crossborder`——从境外销售的商品标识符。
	ParentPostingNumber string `json:"parent_posting_number"` // 当前货件所拆分自的父货件编号。
	TarifficationSteps []PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep `json:"tariffication_steps"` // 计费阶段。
}

// 分类方向：
type Postingv3GetFbsPostingUnfulfilledListRequestDirEnum string
const (
	Postingv3GetFbsPostingUnfulfilledListRequestDirEnum_Asc Postingv3GetFbsPostingUnfulfilledListRequestDirEnum = "asc"
	Postingv3GetFbsPostingUnfulfilledListRequestDirEnum_Desc Postingv3GetFbsPostingUnfulfilledListRequestDirEnum = "desc"
)

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	Dir Postingv3GetFbsPostingUnfulfilledListRequestDirEnum `json:"dir"` // 分类方向： - `asc` — 从小到大， - `desc` — 从大到小。
	Filter Postingv3GetFbsPostingUnfulfilledListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 响应中值的数量： - 最大值 — 1000， - 最小值 — 1。
	Offset int64 `json:"offset"` // 将在响应中跳过的元素数。 例如，如果“offset=10”，那么响应将从找到的第11个元素开始。
	With Postingv3FbsPostingWithParams `json:"with"`
}

type FbsPostingShipV4RequestPackageProduct struct {
	ProductID int64 `json:"product_id"` // Ozon系统中的商品识别码是SKU。
	Quantity int32 `json:"quantity"` // 副本数量。
}

type FbsPostingShipV4RequestPackage struct {
	Products []FbsPostingShipV4RequestPackageProduct `json:"products"` // 运输途中的商品清单。
}

// 附加信息。
type FbsPostingShipV4RequestWith struct {
	AdditionalData bool `json:"additional_data"` // 为获取附加信息，请点击 `true`。
}

type Fbsv4FbsPostingShipV4Request struct {
	PostingNumber string `json:"posting_number"` // 发货号。
	With FbsPostingShipV4RequestWith `json:"with"`
	Packages []FbsPostingShipV4RequestPackage `json:"packages"` // 包装清单。 每个包装都包含订单分成的发货清单。
}

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []string `json:"posting_number"` // 货件ID。
}

// 筛选器。
// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyFbsPostingListRequestFilterCutoffFromEnum string
const (
	V1AssemblyFbsPostingListRequestFilterCutoffFromEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyFbsPostingListRequestFilterCutoffFromEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyFbsPostingListRequestFilterCutoffFromEnum_V20200318T073450359Z V1AssemblyFbsPostingListRequestFilterCutoffFromEnum = "2020-03-18T07:34:50.359Z"
)

// 按卖家需完成订单备货的时间进行筛选
type V1AssemblyFbsPostingListRequestFilterCutoffToEnum string
const (
	V1AssemblyFbsPostingListRequestFilterCutoffToEnum_YYYYMMDDThhMmSsMcsZ V1AssemblyFbsPostingListRequestFilterCutoffToEnum = "YYYY-MM-DDThh:mm:ss.mcsZ"
	V1AssemblyFbsPostingListRequestFilterCutoffToEnum_V20200318T073450359Z V1AssemblyFbsPostingListRequestFilterCutoffToEnum = "2020-03-18T07:34:50.359Z"
)

type V1AssemblyFbsPostingListRequestFilter struct {
	CutoffFrom V1AssemblyFbsPostingListRequestFilterCutoffFromEnum `json:"cutoff_from"` // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	CutoffTo V1AssemblyFbsPostingListRequestFilterCutoffToEnum `json:"cutoff_to"` // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	DeliveryMethodID int64 `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
}

type V2FbsPostingProductCountryListResponseResult struct {
	Name string `json:"name"` // 国家俄语名称
	CountryIsoCode string `json:"country_iso_code"` // ISO国家代码。
}

type PostingV4PostingFbsUnfulfilledListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	HasNext bool `json:"has_next"` // 若响应中未返回全部货件，则为`true`。
	Postings []PostingV4PostingFbsUnfulfilledListResponsePostings `json:"postings"` // 货件列表。
	Count int64 `json:"count"` // 在响应中的元素计数器。
}

type SetPostingsResponseResult struct {
	Error string `json:"error"` // 错误描述。
	PostingNumber string `json:"posting_number"` // 货件编号。
	Result bool `json:"result"` // 请求处理结果：若请求处理成功，返回值为`true`。
}

type V1SetPostingsResponse struct {
	Result []SetPostingsResponseResult `json:"result"`
}

type PostingCancelReasonListResponse struct {
	Result []PostingCancelReason `json:"result"` // 方法操作结果。
}

type V2FbsPostingProductCountryListResponse struct {
	Result []V2FbsPostingProductCountryListResponseResult `json:"result"` // 制造国和ISO代码列表。
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"` // 来自方法[/v1/carriage/create](#operation/CarriageAPI_CarriageCreate)的文件生成任务编号（也是运输标识符）。
}

// 货运数组。
type V3GetFbsPostingListResponseV3Result struct {
	HasNext bool `json:"has_next"` // 响应中未返回整个货运数组的标志: - `true` — 必须提出含其他值 `offset`的新请求，以获得其他货运信息； - `false` — 响应中返回了在请求中提出的整个用于过滤的货运数组。
	Postings []V3FbsPosting `json:"postings"` // 货运信息。
}

// 排序方向： - `ASC`——升序， - `DESC`——降序。
type V1AssemblyFbsPostingListRequestSortDirEnum string

type V3GetFbsPostingListResponseV3 struct {
	Result V3GetFbsPostingListResponseV3Result `json:"result"`
}

type V1AssemblyFbsPostingListRequest struct {
	Limit int64 `json:"limit"` // 每页显示的数量。
	SortDir V1AssemblyFbsPostingListRequestSortDirEnum `json:"sort_dir"`
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter V1AssemblyFbsPostingListRequestFilter `json:"filter"`
}

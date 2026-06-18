package fbo

// 用于计算产品标签的仓库列表。
type GetSupplyOrderBundleRequestItemTagsCalculation struct {
	StorageWarehouseIds []string `json:"storage_warehouse_ids"` // 发货仓库标识符列表，不超过 25 个值。
	DropoffWarehouseID string `json:"dropoff_warehouse_id"` // 用于发货的仓库标识符。
}

type V1OrderDraftValidationErrorErrorType string

// 草稿状态: - `DRAFT_STATUS_UNSPECIFIED` — 未定义; - `NEW` — 新的; - `SUPPLY_VARIANT_CONFIRMATION` — 等待确认; - `SUPPLY_NOT_CONFIRMED`...
type V1DraftStatusEnum string

// 错误状态： - `STATUS_UNSPECIFIED`——未指定； - `CONFIRMATION`——等待申请取消确认； - `CANCELED`——取消已确认； - `NOT_CANCELED`——未收到取消确认。
type V1CancellationStateStatus string

// 取消错误代码： - `CODE_UNSPECIFIED`——未指定； - `NO_RESPONSE_FROM_3PF`——取消申请未确认，未收到第三方响应； - `ACCEPTANCE_ALREADY_STARTED`——取消申请未确认，已...
type CancellationErrorCode string

// 取消错误。
type CancellationStateCancellationError struct {
	ErrorCode CancellationErrorCode `json:"error_code"`
	Message string `json:"message"` // 错误描述。
}

// 取消原因。
type V1CancellationState struct {
	CancellationError CancellationStateCancellationError `json:"cancellation_error"`
	CancellationStatus V1CancellationStateStatus `json:"cancellation_status"`
}

type V1FbpDraftPickUpDeleteResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	CancellationState V1CancellationState `json:"cancellation_state"`
}

type V1FbpOrderListRequest struct {
	Count int32 `json:"count"` // 响应中的交货数量。
	LastID int64 `json:"last_id"` // 页面上最后一次交货的标识符。首次请求时请将此字段留空。 如需获取后续数据，请填写上一次请求响应中最后一次交货的`id`。
}

type V1OrderValidationErrorErrorType string

// 错误信息。
// 错误类型：
type V1OrderValidationErrorOrderErrorsEnum string
const (
	V1OrderValidationErrorOrderErrorsEnum_ERRORTYPEUNSPECIFIED V1OrderValidationErrorOrderErrorsEnum = "ERROR_TYPE_UNSPECIFIED"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYDRIVERNAMELENGTHMAXIMUMREACHED V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_DRIVER_NAME_LENGTH_MAXIMUM_REACHED"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYVEHICLEGENRELENGTHMAXIMUMREACHED V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_VEHICLE_GENRE_LENGTH_MAXIMUM_REACHED"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYVEHICLEREGISTRATIONPLATELENGTHMAXIMUMREACHED V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_VEHICLE_REGISTRATION_PLATE_LENGTH_MAXIMUM_REACHED"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYTPLNAMELENGTHMAXIMUMREACHED V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_TPL_NAME_LENGTH_MAXIMUM_REACHED"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYTRACKINGNUMBERLENGTHMAXIMUMREACHED V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_TRACKING_NUMBER_LENGTH_MAXIMUM_REACHED"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYDRIVERNAMEEMPTY V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_DRIVER_NAME_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYVEHICLEGENREEMPTY V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_VEHICLE_GENRE_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYVEHICLEREGISTRATIONPLATEEMPTY V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_VEHICLE_REGISTRATION_PLATE_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYTPLNAMEEMPTY V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_TPL_NAME_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYTRACKINGNUMBEREMPTY V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_TRACKING_NUMBER_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYBYSELLEREMPTY V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_BY_SELLER_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_DELIVERYBYTPLEMPTY V1OrderValidationErrorOrderErrorsEnum = "DELIVERY_BY_TPL_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_RECEIVEDATENOTSET V1OrderValidationErrorOrderErrorsEnum = "RECEIVE_DATE_NOT_SET"
	V1OrderValidationErrorOrderErrorsEnum_SUPPLYTYPENOTSUPPORTED V1OrderValidationErrorOrderErrorsEnum = "SUPPLY_TYPE_NOT_SUPPORTED"
	V1OrderValidationErrorOrderErrorsEnum_INVALIDBUSINESSFLOW V1OrderValidationErrorOrderErrorsEnum = "INVALID_BUSINESS_FLOW"
	V1OrderValidationErrorOrderErrorsEnum_ORDERLOCKED V1OrderValidationErrorOrderErrorsEnum = "ORDER_LOCKED"
	V1OrderValidationErrorOrderErrorsEnum_INVALIDTIMESLOT V1OrderValidationErrorOrderErrorsEnum = "INVALID_TIMESLOT"
	V1OrderValidationErrorOrderErrorsEnum_DROPOFFDETAILSEMPTY V1OrderValidationErrorOrderErrorsEnum = "DROP_OFF_DETAILS_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPADDRESSISEMPTY V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_ADDRESS_IS_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPSENDERNAMEISEMPTY V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_SENDER_NAME_IS_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPSENDERPHONEISEMPTY V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_SENDER_PHONE_IS_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPADDRESSISTOOLARGE V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_ADDRESS_IS_TOO_LARGE"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPSENDERNAMEISTOOLARGE V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_SENDER_NAME_IS_TOO_LARGE"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPSENDERPHONEISTOOLARGE V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_SENDER_PHONE_IS_TOO_LARGE"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPCOMMENTISTOOLARGE V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_COMMENT_IS_TOO_LARGE"
	V1OrderValidationErrorOrderErrorsEnum_PICKUPDETAILSEMPTY V1OrderValidationErrorOrderErrorsEnum = "PICK_UP_DETAILS_EMPTY"
	V1OrderValidationErrorOrderErrorsEnum_DROPOFFADDRESSNOTSET V1OrderValidationErrorOrderErrorsEnum = "DROP_OFF_ADDRESS_NOT_SET"
	V1OrderValidationErrorOrderErrorsEnum_INVALIDSTATE V1OrderValidationErrorOrderErrorsEnum = "INVALID_STATE"
)

type V1OrderValidationError struct {
	OrderErrors []V1OrderValidationErrorErrorType `json:"order_errors"` // 错误类型： - `ERROR_TYPE_UNSPECIFIED`——未定义； - `DELIVERY_DRIVER_NAME_LENGTH_MAXIMUM_REACHED`——司机姓名长度超限； - `DELIVERY_VEHICLE_GE...
}

type V1FbpOrderDropOffCancelResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 配送详情。
type V1FbpDraftDirectSellerDlvCreateRequestDirectDetails struct {
	VehicleType string `json:"vehicle_type"` // 车辆类型。
	DriverName string `json:"driver_name"` // 司机姓名。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
}

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	BundleID string `json:"bundle_id"` // 已验证商品清单的标识符。
	DeliveryDetails V1FbpDraftDirectSellerDlvCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库标识符。
}

type V1FbpGetLabelRequest struct {
	Code string `json:"code"` // 标签生成任务标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 排序参数： - `SKU`——SKU； - `NAME`——按商品名称； - `QUANTITY`——按数量； - `TOTAL_VOLUME_IN_LITRES`——按体积（升）。
type V1ItemSortField string

type V1GetSupplyOrderBundleRequest struct {
	Query string `json:"query"` // 搜索查询，例如按商品名称、货号或 SKU 搜索。
	SortField V1ItemSortField `json:"sort_field"`
	BundleIds []string `json:"bundle_ids"` // 交货商品组成的标识符。可通过方法 [/v3/supply-order/get](#operation/SupplyOrderGet) 获取。
	IsAsc bool `json:"is_asc"` // 传入 `true` 表示按升序排序。
	ItemTagsCalculation GetSupplyOrderBundleRequestItemTagsCalculation `json:"item_tags_calculation"`
	LastID string `json:"last_id"` // 当前页面中最后一个 SKU 值的标识符。
	Limit int32 `json:"limit"` // 每页商品数量。
}

type V1FbpOrderPickUpCancelResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 休息时间。
type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"` // 时间段结束时间。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProductsActions struct {
	Description string `json:"description"` // 促销活动名称。
	ActionID string `json:"action_id"` // 促销活动标识符。
	DateFrom string `json:"date_from"` // 促销活动开始日期。
	DateTo string `json:"date_to"` // 促销活动结束日期。
	DiscountPercent float64 `json:"discount_percent"` // 折扣百分比。
	DiscountValue float64 `json:"discount_value"` // 折扣金额。
	IsFromSeller bool `json:"is_from_seller"` // `true`，表示促销活动由卖家创建。
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProducts struct {
	TotalDiscountValue float64 `json:"total_discount_value"` // 折扣金额。
	Actions []PostingV1PostingFbpListResponsePostingsFinancialDataProductsActions `json:"actions"` // 促销活动列表。
	CommissionsCurrencyCode string `json:"commissions_currency_code"` // 佣金货币代码。
	OldPrice float64 `json:"old_price"` // 折扣前的价格。商品卡上会以划线价显示。
	Price float64 `json:"price"` // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	ProductID int64 `json:"product_id"` // Ozon系统中的商品标识符，即SKU。
	Quantity int64 `json:"quantity"` // 商品数量。
	TotalDiscountPercent float64 `json:"total_discount_percent"` // 折扣百分比。
}

type V1FbpEditTimeslotResponseReserveFailureType string

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 生成状态： - `STATE_TYPE_UNSPECIFIED` ——未定义； - `IN_PROGRESS` ——进行中； - `FINISHED` ——成功完成； - `FAILED` ——错误。
type FbpCheckConsignmentNoteStateResponseStateType string

type V1BundleItemErrorEnum string

// 错误：
type V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum string
const (
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_BUNDLEITEMERRORUNSPECIFIED V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "BUNDLE_ITEM_ERROR_UNSPECIFIED"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_OUTOFASSORTMENT V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "OUT_OF_ASSORTMENT"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INVALID V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INVALID"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INCOMPATIBLEWAREHOUSE V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INCOMPATIBLE_WAREHOUSE"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INVALIDBARCODE V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INVALID_BARCODE"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_MULTIPLICITY V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "MULTIPLICITY"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_NOPRICE V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "NO_PRICE"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_BANNED V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "BANNED"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_ZEROQUANTITY V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "ZERO_QUANTITY"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_QUANTITYGREATERTHENMAX V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "QUANTITY_GREATER_THEN_MAX"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_NOSALES V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "NO_SALES"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_SURPLUS V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "SURPLUS"
	V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_AVAILABILITYISEMPTY V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "AVAILABILITY_IS_EMPTY"
)

type V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError struct {
	Errors []V1BundleItemErrorEnum `json:"errors"` // 错误： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——商品不在交货品类中； - `INVALID`——状态不正确； - `INCOMPATIBLE_WAREHO...
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符——SKU。
}

// 交货注册错误： - `ORDER_ERROR_TYPE_UNSPECIFIED` — 未知订单错误类型； - `INVALID_NUMBER_OF_PACKAGE_UNITS` — 申请中货位数量错误； - `MAXIMUM_NUMBER_...
type V1OrderErrorTypeEnum string

// 错误。
type V1FbpDraftPickUpRegistrateResponseRegistrationError struct {
	BundleErrors []V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError `json:"bundle_errors"` // 商品验证列表错误。
	OrderError V1OrderErrorTypeEnum `json:"order_error"`
}

type V1FbpDraftPickUpRegistrateResponse struct {
	Error V1FbpDraftPickUpRegistrateResponseRegistrationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 计入Ozon折扣后的卖家价格。
type MoneyMoneySellerPrice struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

type V1OrderAttentionTypeEnum string

type V1FbpDraftPickUpProductValidateResponseApprovedItem struct {
	IconName string `json:"icon_name"` // 商品图片链接。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品货号。
	Quantity int32 `json:"quantity"` // 商品数量。
	SKU int64 `json:"sku"` // 商品标识符（SKU）。
	Volume float64 `json:"volume"` // 商品体积。
	Barcode string `json:"barcode"` // 条形码。
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	DriverName string `json:"driver_name"` // 司机姓名。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
	VehicleType string `json:"vehicle_type"` // 车辆类型。
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"` // 验收证明书标识符。
}

// 配送详细信息。
type V1FbpDraftDirectTplDlvCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"` // 时间段开始本地时间。
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号码。
	TransportCompanyName string `json:"transport_company_name"` // 物流公司名称。
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	BundleID string `json:"bundle_id"` // 套装标识符。
	DeliveryDetails V1FbpDraftDirectTplDlvCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

// 标签生成任务状态： - `UNSPECIFIED`：未指定； - `IN_PROGRESS`：生成中； - `FINISHED`：生成成功； - `FAILED`：生成失败。
type FbpGetLabelResponseLabelCreationStateTypeEnum string

type V1FbpDraftDropOffProductValidateResponseApprovedItem struct {
	Barcode string `json:"barcode"` // 商品条形码。
	IconName string `json:"icon_name"` // 商品图片链接。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符——货号。
	Quantity int32 `json:"quantity"` // 商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符—— SKU。
	Volume float64 `json:"volume"` // 商品体积。
}

// 错误：
type V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum string
const (
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_BUNDLEITEMERRORUNSPECIFIED V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "BUNDLE_ITEM_ERROR_UNSPECIFIED"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_OUTOFASSORTMENT V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "OUT_OF_ASSORTMENT"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INVALID V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INVALID"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INCOMPATIBLEWAREHOUSE V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INCOMPATIBLE_WAREHOUSE"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INVALIDBARCODE V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INVALID_BARCODE"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_MULTIPLICITY V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "MULTIPLICITY"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_NOPRICE V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "NO_PRICE"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_BANNED V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "BANNED"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_ZEROQUANTITY V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "ZERO_QUANTITY"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_QUANTITYGREATERTHENMAX V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "QUANTITY_GREATER_THEN_MAX"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_NOSALES V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "NO_SALES"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_SURPLUS V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "SURPLUS"
	V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_AVAILABILITYISEMPTY V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "AVAILABILITY_IS_EMPTY"
)

type V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError struct {
	Errors []V1BundleItemErrorEnum `json:"errors"` // 错误： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——商品不在交货品类中； - `INVALID`——状态不正确； - `INCOMPATIBLE_WAREHO...
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
}

// 错误。
type V1FbpDraftDropOffRegistrateResponseRegistrationError struct {
	BundleErrors []V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError `json:"bundle_errors"` // 商品验证列表错误。
	OrderError V1OrderErrorTypeEnum `json:"order_error"`
}

type V1FbpDraftDropOffRegistrateResponse struct {
	Error V1FbpDraftDropOffRegistrateResponseRegistrationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"` // 交货到揽收点的到达日期。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 财务数据。
type PostingV1PostingFbpListResponsePostingsFinancialData struct {
	ClusterFrom string `json:"cluster_from"` // 订单发出地区代码。
	ClusterTo string `json:"cluster_to"` // 订单配送地区代码。
	DeliveryAmount float64 `json:"delivery_amount"` // 配送费用。
	Products []PostingV1PostingFbpListResponsePostingsFinancialDataProducts `json:"products"` // 订单中的商品列表。
}

// 网站上的商品价格。
type MoneyMoneyCustomerPrice struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 商品价格。
type MoneyPostingMoney struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

type PostingV1PostingFbpListResponsePostingsProducts struct {
	Price MoneyPostingMoney `json:"price"`
	Quantity int32 `json:"quantity"` // 货件中的商品数量。
	SellerPrice MoneyMoneySellerPrice `json:"seller_price"`
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符，即SKU。
	CustomerPrice MoneyMoneyCustomerPrice `json:"customer_price"`
	Name string `json:"name"` // 订单中的商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符，即货号。
}

type PostingV1PostingFbpListResponsePostings struct {
	OrderID int64 `json:"order_id"` // 该货件所属订单的标识符。
	Products []PostingV1PostingFbpListResponsePostingsProducts `json:"products"` // 货件中商品列表。
	ProviderID int64 `json:"provider_id"` // 配送服务标识符。
	InProcessAt string `json:"in_process_at"` // 货件开始处理的日期和时间。
	OrderNumber string `json:"order_number"` // 该货件所属订单的编号。
	PostingNumber string `json:"posting_number"` // 货件编号。
	Status string `json:"status"` // 货件状态。
	FinancialData PostingV1PostingFbpListResponsePostingsFinancialData `json:"financial_data"`
	OrderDate string `json:"order_date"` // 订单的创建日期。
}

// 超级产品标签： - `ITEM_SFBO_ATTRIBUTE_NONE`——无标签； - `ITEM_SFBO_ATTRIBUTE_SUPER_FBO`——超级产品； - `ITEM_SFBO_ATTRIBUTE_ANTI_FBO`——滞销...
type V1ItemSfboAttribute string

// 包装类型： - `BUNDLE_ITEM_SHIPMENT_TYPE_GENERAL`——常规产品； - `BUNDLE_ITEM_SHIPMENT_TYPE_BOX`——盒装； - `BUNDLE_ITEM_SHIPMENT_TYPE_P...
type V1BundleItemShipmentType string

// 商品存放区域：
type V1ItemResponsePlacementZoneEnum string
const (
	V1ItemResponsePlacementZoneEnum_UNSPECIFIED V1ItemResponsePlacementZoneEnum = "UNSPECIFIED"
	V1ItemResponsePlacementZoneEnum_CLOSEDZONE V1ItemResponsePlacementZoneEnum = "CLOSED_ZONE"
	V1ItemResponsePlacementZoneEnum_DANGEROUSGOODS V1ItemResponsePlacementZoneEnum = "DANGEROUS_GOODS"
	V1ItemResponsePlacementZoneEnum_PRODUCTS V1ItemResponsePlacementZoneEnum = "PRODUCTS"
	V1ItemResponsePlacementZoneEnum_SORT V1ItemResponsePlacementZoneEnum = "SORT"
	V1ItemResponsePlacementZoneEnum_NONSORT V1ItemResponsePlacementZoneEnum = "NON_SORT"
	V1ItemResponsePlacementZoneEnum_OVERSIZE V1ItemResponsePlacementZoneEnum = "OVERSIZE"
	V1ItemResponsePlacementZoneEnum_JEWELRY V1ItemResponsePlacementZoneEnum = "JEWELRY"
	V1ItemResponsePlacementZoneEnum_UNRESOLVED V1ItemResponsePlacementZoneEnum = "UNRESOLVED"
)

// 交货或交货申请中的商品标签
type V1ItemResponseTagsEnum string
const (
	V1ItemResponseTagsEnum_EVSDREQUIRED V1ItemResponseTagsEnum = "EVSD_REQUIRED"
	V1ItemResponseTagsEnum_MARKINGREQUIRED V1ItemResponseTagsEnum = "MARKING_REQUIRED"
	V1ItemResponseTagsEnum_MARKINGPOSSIBLE V1ItemResponseTagsEnum = "MARKING_POSSIBLE"
	V1ItemResponseTagsEnum_JEWELRY V1ItemResponseTagsEnum = "JEWELRY"
	V1ItemResponseTagsEnum_TRACEABLE V1ItemResponseTagsEnum = "TRACEABLE"
	V1ItemResponseTagsEnum_ETTNREQUIRED V1ItemResponseTagsEnum = "ETTN_REQUIRED"
	V1ItemResponseTagsEnum_UNDEFINED V1ItemResponseTagsEnum = "UNDEFINED"
)

type V1ItemResponse struct {
	Barcode string `json:"barcode"` // 商品条形码。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	VolumeInLitres float64 `json:"volume_in_litres"` // 商品体积（升）。
	ContractorItemCode string `json:"contractor_item_code"` // 卖家系统中的商品标识符——货号。
	Quantity int32 `json:"quantity"` // 商品数量。
	IconPath string `json:"icon_path"` // 商品图片链接。
	SKU int64 `json:"sku"` // 商品在Ozon系统中的ID（SKU）。
	ShipmentType V1BundleItemShipmentType `json:"shipment_type"`
	Tags []string `json:"tags"` // 交货或交货申请中的商品标签。 可能的取值： - `EVSD_REQUIRED`——需要 Mercury 认证的商品； - `MARKING_REQUIRED`——需要 “诚实标志” 强制标志的商品； - `MARKING_POSSIBLE`...
	Name string `json:"name"` // 商品名称。
	Quant int32 `json:"quant"` // 单个包装中的商品数量。
	IsQuantEditable bool `json:"is_quant_editable"` // `true`，表示单个包装中的商品数量可修改。
	TotalVolumeInLitres float64 `json:"total_volume_in_litres"` // 所有商品总体积（升）。
	SfboAttribute V1ItemSfboAttribute `json:"sfbo_attribute"`
	PlacementZone V1ItemResponsePlacementZoneEnum `json:"placement_zone"` // 商品存放区域： - `UNSPECIFIED`——未指定； - `CLOSED_ZONE`——封闭区域； - `DANGEROUS_GOODS`——2–4 类危险品； - `PRODUCTS`——食品； - `SORT`——可分拣商品； -...
	OfferID string `json:"offer_id"` // 商品在卖家系统中的标识符 — 货号。
}

type V1GetSupplyOrderBundleResponse struct {
	Items []V1ItemResponse `json:"items"` // 交货申请中的商品列表。
	TotalCount int32 `json:"total_count"` // 申请中的商品数量。
	HasNext bool `json:"has_next"` // 响应中是否未返回全部商品： - `true`——请使用不同的 `last_id` 值再次请求，以获取其余数据； - `false`——响应已包含全部商品数据。
	LastID string `json:"last_id"` // 当前页面最后一个值的标识符。
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpDraftDirectTimeslotEditResponseReserveFailureType string

type V1FbpDraftPickUpProductValidateRequestSkuItem struct {
	Count int32 `json:"count"` // 交货商品数量。
	SKU int64 `json:"sku"` // 商品标识符（SKU）。
}

// 错误信息。
// 错误类型：
type V1OrderDraftValidationErrorErrorsEnum string
const (
	V1OrderDraftValidationErrorErrorsEnum_ERRORTYPEUNSPECIFIED V1OrderDraftValidationErrorErrorsEnum = "ERROR_TYPE_UNSPECIFIED"
	V1OrderDraftValidationErrorErrorsEnum_ORDERDRAFTLOCKED V1OrderDraftValidationErrorErrorsEnum = "ORDER_DRAFT_LOCKED"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYDRIVERNAMELENGTHMAXIMUMREACHED V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_DRIVER_NAME_LENGTH_MAXIMUM_REACHED"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYVEHICLEGENRELENGTHMAXIMUMREACHED V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_VEHICLE_GENRE_LENGTH_MAXIMUM_REACHED"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYVEHICLEREGISTRATIONPLATELENGTHMAXIMUMREACHED V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_VEHICLE_REGISTRATION_PLATE_LENGTH_MAXIMUM_REACHED"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYTPLNAMELENGTHMAXIMUMREACHED V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_TPL_NAME_LENGTH_MAXIMUM_REACHED"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYTRACKINGNUMBERLENGTHMAXIMUMREACHED V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_TRACKING_NUMBER_LENGTH_MAXIMUM_REACHED"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYDRIVERNAMEEMPTY V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_DRIVER_NAME_EMPTY"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYVEHICLEGENREEMPTY V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_VEHICLE_GENRE_EMPTY"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYVEHICLEREGISTRATIONPLATEEMPTY V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_VEHICLE_REGISTRATION_PLATE_EMPTY"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYTPLNAMEEMPTY V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_TPL_NAME_EMPTY"
	V1OrderDraftValidationErrorErrorsEnum_DELIVERYTRACKINGNUMBEREMPTY V1OrderDraftValidationErrorErrorsEnum = "DELIVERY_TRACKING_NUMBER_EMPTY"
	V1OrderDraftValidationErrorErrorsEnum_INVALIDBUSINESSFLOW V1OrderDraftValidationErrorErrorsEnum = "INVALID_BUSINESS_FLOW"
	V1OrderDraftValidationErrorErrorsEnum_SUPPLYTYPENOTSUPPORTED V1OrderDraftValidationErrorErrorsEnum = "SUPPLY_TYPE_NOT_SUPPORTED"
	V1OrderDraftValidationErrorErrorsEnum_INVALIDSTATE V1OrderDraftValidationErrorErrorsEnum = "INVALID_STATE"
)

type V1OrderDraftValidationError struct {
	Errors []V1OrderDraftValidationErrorErrorType `json:"errors"` // 错误类型： - `ERROR_TYPE_UNSPECIFIED`——未定义； - `ORDER_DRAFT_LOCKED`——草稿被锁定； - `DELIVERY_DRIVER_NAME_LENGTH_MAXIMUM_REACHED`——司...
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	Error V1OrderDraftValidationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type FbpDraftDropOffPointListResponseDropOffPoint struct {
	City string `json:"city"` // 城市。
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	NearestDropOffDate string `json:"nearest_drop_off_date"` // 最近的发运日期。
	PointAddress string `json:"point_address"` // 接收点地址。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
}

type FbpDraftDropOffProvinceListResponseProvince struct {
	Name string `json:"name"` // 省份名称。
	PointsCount int32 `json:"points_count"` // 地图上接收点数量。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []FbpDraftDropOffProvinceListResponseProvince `json:"provinces"` // 省份列表。
}

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type PostingV1PostingFbpListRequestSortDirEnum string

// 地址详情。
type FbpWarehouseListResponseAddressDetailing struct {
	City string `json:"city"` // 城市。
	Country string `json:"country"` // 国家。
	House string `json:"house"` // 门牌号。
	Region string `json:"region"` // 地区。
	Street string `json:"street"` // 街道。
	Zipcode string `json:"zipcode"` // 邮政编码。
}

type FbpWarehouseListResponseWarehouse struct {
	ID int64 `json:"id"` // 仓库标识符。
	IsBonded bool `json:"is_bonded"` // `true`，表示该仓库为保税仓。
	Name string `json:"name"` // 仓库名称。
	PartnerName string `json:"partner_name"` // 合作伙伴名称。
	SupplyTypes []int32 `json:"supply_types"` // 交货类型。
	TimezoneName string `json:"timezone_name"` // 仓库所在时区。
	AddressDetailing FbpWarehouseListResponseAddressDetailing `json:"address_detailing"`
}

type V1FbpWarehouseListResponse struct {
	Warehouses []FbpWarehouseListResponseWarehouse `json:"warehouses"` // 仓库列表。
}

type V1FbpDraftDirectTplDlvEditResponse struct {
	Error V1OrderDraftValidationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 交货类型： - `SUPPLY_TYPE_UNSPECIFIED`：未指定； - `DIRECT_BY_SELLER`：卖家自行送货到仓库； - `DIRECT_BY_TPL`：第三方物流公司送货到仓库； - `DROP_OFF`：送货到揽...
type DeliveryDetailsSupplyType string

type V1FbpDraftListRequest struct {
	Count int32 `json:"count"` // 响应中的商品数量。
	LastID int64 `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
}

type V1FbpDraftDropOffDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDirectProductValidateResponseApprovedItem struct {
	IconName string `json:"icon_name"` // 商品图片链接。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品货号。
	Quantity int32 `json:"quantity"` // 商品数量。
	SKU int64 `json:"sku"` // 商品标识符（SKU）。
	Volume float64 `json:"volume"` // 商品体积。
	Barcode string `json:"barcode"` // 条形码。
}

// 卖家自配送详情。
type DirectDetailsBySellerDetails struct {
	DriverName string `json:"driver_name"` // 司机姓名。
	VehicleRegistrationNumber string `json:"vehicle_registration_number"` // 车牌号码。
	VehicleType string `json:"vehicle_type"` // 运输工具类型。
}

// 第三方物流公司配送详情。
type DirectDetailsByTplDetails struct {
	TransportCompanyName string `json:"transport_company_name"` // 物流公司名称。
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号码。
}

// 时间段。
type V1fbpTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"` // 时间段结束时间（UTC）。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间（UTC）。
}

// 交货时间段详情。
type DirectDetailsTimeslotDetails struct {
	Timeslot V1fbpTimeslot `json:"timeslot"`
	TimeslotReservationID string `json:"timeslot_reservation_id"` // 交货时间段预定标识符。
}

// 卖家配送详情。
type V1DeliveryDetailsDirectDetails struct {
	BySellerDetails DirectDetailsBySellerDetails `json:"by_seller_details"`
	ByTPLDetails DirectDetailsByTplDetails `json:"by_tpl_details"`
	TimeslotDetails DirectDetailsTimeslotDetails `json:"timeslot_details"`
}

// 揽收点详情。
type DeliveryDetailsDropOffPointDetails struct {
	ID int64 `json:"id"` // 揽收点标识符。
	ProvinceUuid string `json:"province_uuid"` // 区域唯一标识符。
	Timeslot V1fbpTimeslot `json:"timeslot"`
}

// 取货点详情。
type V1DeliveryDetailsPickUpDetails struct {
	Date string `json:"date"` // 送货日期。
	SenderName string `json:"sender_name"` // 发件人姓名。
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
	Address string `json:"address"` // 地址。
	Comment string `json:"comment"` // 备注。
}

// 配送详情。
type Fbpv1DeliveryDetails struct {
	DirectDetails V1DeliveryDetailsDirectDetails `json:"direct_details"`
	DropOffPoint DeliveryDetailsDropOffPointDetails `json:"drop_off_point"`
	PickupDetails V1DeliveryDetailsPickUpDetails `json:"pickup_details"`
	SupplyType DeliveryDetailsSupplyType `json:"supply_type"`
}

// 发件人详细信息。
type V1FbpOrderPickUpDlvEditRequestPickUpDetails struct {
	SenderName string `json:"sender_name"` // 发件人姓名。
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
}

// 生成错误： - `ERROR_REASON_UNSPECIFIED` ——未定义； - `INVALID_COMPANY` ——公司无效； - `FILE_NOT_FOUND` ——文件未找到； - `GENERATE_TIMEOUT_RE...
type FbpCheckActStateResponseErrorReason string

// 营业时间。
type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening struct {
	TimeslotEnd string `json:"timeslot_end"` // 时间段结束时间。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
}

// 营业时间表。
type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem struct {
	BreakHours V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"` // `true`，表示节假日。
	OpeningHours V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening `json:"opening_hours"`
}

// 星期： - `DAY_OF_WEEK_UNSPECIFIED`——未指定； - `MONDAY`——星期一； - `TUESDAY`——星期二； - `WEDNESDAY`——星期三； - `THURSDAY`——星期四； - `FRIDA...
type V1DayOfWeek string

type V1FbpDraftDropOffPointTimetableResponseCalendar struct {
	CalendarItem V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem `json:"calendar_item"`
	DayOfWeek V1DayOfWeek `json:"day_of_week"`
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []V1FbpDraftDropOffPointTimetableResponseCalendar `json:"calendar"` // 接收点的营业时间表。
}

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 配送详细信息。
type V1FbpDraftDirectCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"` // 配送时间段开始。
}

// 交货商品汇总信息。
type ItemBundleSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"` // 商品总体积（升）。
	TotalItemCount int64 `json:"total_item_count"` // 交货中的SKU数量。
	TotalQuantity int64 `json:"total_quantity"` // 交货中的商品数量。
}

// 拒绝原因：
type V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum string
const (
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_BUNDLEITEMERRORUNSPECIFIED V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "BUNDLE_ITEM_ERROR_UNSPECIFIED"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_OUTOFASSORTMENT V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "OUT_OF_ASSORTMENT"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_INVALID V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "INVALID"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_INCOMPATIBLEWAREHOUSE V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "INCOMPATIBLE_WAREHOUSE"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_INVALIDBARCODE V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "INVALID_BARCODE"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_MULTIPLICITY V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "MULTIPLICITY"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_NOPRICE V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "NO_PRICE"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_BANNED V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "BANNED"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_ZEROQUANTITY V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "ZERO_QUANTITY"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_QUANTITYGREATERTHENMAX V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "QUANTITY_GREATER_THEN_MAX"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_NOSALES V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "NO_SALES"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_SURPLUS V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "SURPLUS"
	V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum_AVAILABILITYISEMPTY V1FbpDraftPickUpProductValidateResponseRejectedItemRejectionReasonsEnum = "AVAILABILITY_IS_EMPTY"
)

type V1FbpDraftPickUpProductValidateResponseRejectedItem struct {
	RejectionReasons []V1BundleItemErrorEnum `json:"rejection_reasons"` // 拒绝原因： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——未找到商品； - `INVALID`——商品未创建； - `INCOMPATIBLE_WAREHOUS...
	SKU int64 `json:"sku"` // 商品标识符（SKU）。
	Volume float64 `json:"volume"` // 商品体积。
	Barcode string `json:"barcode"` // 条形码。
	IconName string `json:"icon_name"` // 商品图片链接。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品货号。
	Quantity int32 `json:"quantity"` // 商品数量。
}

type V1FbpDraftPickUpProductValidateResponse struct {
	ApprovedItems []V1FbpDraftPickUpProductValidateResponseApprovedItem `json:"approved_items"` // 已确认商品。
	BundleGenerated bool `json:"bundle_generated"` // `true`，前提是已创建校验商品列表。
	BundleID string `json:"bundle_id"` // 校验商品列表标识符。
	RejectedItems []V1FbpDraftPickUpProductValidateResponseRejectedItem `json:"rejected_items"` // 被拒绝的商品。
}

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 营业时间信息。
type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"` // 开始时间。
	TimeslotStart string `json:"timeslot_start"` // 结束时间。
}

// 错误原因：
type V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum string
const (
	V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum_RESERVEFAILURETYPEUNSPECIFIED V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum = "RESERVE_FAILURE_TYPE_UNSPECIFIED"
	V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum_REQUESTVALIDATION V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum = "REQUEST_VALIDATION"
	V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum_INVALIDRESERVE V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum = "INVALID_RESERVE"
	V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum_LOGISTICSREASON V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum = "LOGISTICS_REASON"
	V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum_SCHEDULEREASON V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum = "SCHEDULE_REASON"
	V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum_NOCAPACITY V1FbpDraftDirectTimeslotEditResponseErrorReasonsEnum = "NO_CAPACITY"
)

type V1FbpDraftDirectTimeslotEditResponse struct {
	ErrorReasons []V1FbpDraftDirectTimeslotEditResponseReserveFailureType `json:"error_reasons"` // 错误原因： - `RESERVE_FAILURE_TYPE_UNSPECIFIED`——未定义； - `REQUEST_VALIDATION`——请求中填写了过去的预定日期； - `INVALID_RESERVE`——原始预留未找到、已失效...
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// Детали доставки.
type V1FbpDraftPickupDlvEditRequestDeliveryDetails struct {
	SenderName string `json:"sender_name"` // 发件人姓名。
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
	Address string `json:"address"` // 地址。
	Comment string `json:"comment"` // 备注。
	Date string `json:"date"` // 送货日期。
}

// 拒收原因：
type V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum string
const (
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_BUNDLEITEMERRORUNSPECIFIED V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "BUNDLE_ITEM_ERROR_UNSPECIFIED"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_OUTOFASSORTMENT V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "OUT_OF_ASSORTMENT"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_INVALID V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "INVALID"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_INCOMPATIBLEWAREHOUSE V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "INCOMPATIBLE_WAREHOUSE"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_INVALIDBARCODE V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "INVALID_BARCODE"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_MULTIPLICITY V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "MULTIPLICITY"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_NOPRICE V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "NO_PRICE"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_BANNED V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "BANNED"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_ZEROQUANTITY V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "ZERO_QUANTITY"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_QUANTITYGREATERTHENMAX V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "QUANTITY_GREATER_THEN_MAX"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_NOSALES V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "NO_SALES"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_SURPLUS V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "SURPLUS"
	V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum_AVAILABILITYISEMPTY V1FbpDraftDropOffProductValidateResponseRejectedItemRejectionReasonsEnum = "AVAILABILITY_IS_EMPTY"
)

type V1FbpDraftDropOffProductValidateResponseRejectedItem struct {
	Barcode string `json:"barcode"` // 商品条形码。
	IconName string `json:"icon_name"` // 商品图片链接。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符——货号。
	Quantity int32 `json:"quantity"` // 商品数量。
	RejectionReasons []V1BundleItemErrorEnum `json:"rejection_reasons"` // 拒收原因： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——商品不在交货范围内； - `INVALID`——状态不正确； - `INCOMPATIBLE_WARE...
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符—— SKU。
	Volume float64 `json:"volume"` // 商品体积。
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号码。
	TransportCompanyName string `json:"transport_company_name"` // 物流公司名称。
}

type V1FbpDraftDirectProductValidateRequestSkuItem struct {
	Count int64 `json:"count"` // 交货商品数量。
	SKU int64 `json:"sku"` // 商品标识符（SKU）。
}

type V1FbpDraftDirectProductValidateRequest struct {
	Skus []V1FbpDraftDirectProductValidateRequestSkuItem `json:"skus"` // 商品标识符（SKU）列表。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"` // 标签生成任务标识符。
}

// 拒绝原因。
type FbpDraftGetResponseDeclineReason struct {
	FailedSKUIds []string `json:"failed_sku_ids"` // 不正确的SKU标识符。
	Message string `json:"message"` // 拒绝文本。
}

type V1FbpDraftDirectGetTimeslotResponseTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"` // 时间段结束日期。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始日期。
}

type V1FbpAvailableTimeslotListRequest struct {
	IntervalStart string `json:"interval_start"` // 可用时间段所需区间的开始日期。
	SupplyID string `json:"supply_id"` // 交货标识符。
	IntervalEnd string `json:"interval_end"` // 可用时间段所需区间的结束日期。
}

type V1FbpCheckConsignmentNoteStateResponse struct {
	ErrorMessage string `json:"error_message"` // 错误描述。
	LabelURL string `json:"label_url"` // 交货标签链接。
	State FbpCheckConsignmentNoteStateResponseStateType `json:"state"`
}

// 拒绝交货的原因。
// 拒绝交货原因代码：
type V1ArchiveDeclineReasonCodeEnum string
const (
	V1ArchiveDeclineReasonCodeEnum_DECLINEREASONCODEUNSPECIFIED V1ArchiveDeclineReasonCodeEnum = "DECLINE_REASON_CODE_UNSPECIFIED"
	V1ArchiveDeclineReasonCodeEnum_CANNOTCREATESUPPLYONTPF V1ArchiveDeclineReasonCodeEnum = "CANNOT_CREATE_SUPPLY_ON_TPF"
	V1ArchiveDeclineReasonCodeEnum_DROPOFFPOINTCLOSED V1ArchiveDeclineReasonCodeEnum = "DROP_OFF_POINT_CLOSED"
	V1ArchiveDeclineReasonCodeEnum_CODESUPPLYLOST V1ArchiveDeclineReasonCodeEnum = "CODE_SUPPLY_LOST"
	V1ArchiveDeclineReasonCodeEnum_COURIERPICKUPREJECTEDBYSELLER V1ArchiveDeclineReasonCodeEnum = "COURIER_PICK_UP_REJECTED_BY_SELLER"
	V1ArchiveDeclineReasonCodeEnum_BONDEDDOCUMENTSREJECTEDBYWAREHOUSE V1ArchiveDeclineReasonCodeEnum = "BONDED_DOCUMENTS_REJECTED_BY_WAREHOUSE"
)

type V1ArchiveDeclineReason struct {
	Code V1ArchiveDeclineReasonCodeEnum `json:"code"` // 拒绝交货原因代码： - `DECLINE_REASON_CODE_UNSPECIFIED`：未指定； - `CANNOT_CREATE_SUPPLY_ON_TPF`：无法在3PF创建交货； - `DROP_OFF_POINT_CLOSED`...
	Message string `json:"message"` // 拒绝原因说明。
}

// 已完成的交货状态： - `ARCHIVE_STATUS_UNSPECIFIED`：未指定； - `COMPLETED`：已完成； - `REJECTED_AT_SUPPLY_WAREHOUSE`：被仓库拒绝； - `CANCELLED_BY...
type V1ArchiveStatusEnum string

// 交货商品汇总信息。
type V1ArchiveSkuSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"` // 商品总体积（升）。
	TotalItemsCount int64 `json:"total_items_count"` // 交货中的SKU数量。
	TotalQuantity int64 `json:"total_quantity"` // 交货中的商品数量。
}

type V1FbpArchiveGetResponse struct {
	BundleID string `json:"bundle_id"` // 已验证商品清单的标识符。
	OrderDraftID int64 `json:"order_draft_id"` // 交货草稿标识符。
	OrderNumber string `json:"order_number"` // 已完成交货标识符。
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	BundleSKUSummary V1ArchiveSkuSummary `json:"bundle_sku_summary"`
	BusinessFlowTypeID int64 `json:"business_flow_type_id"` // 交货类型标识符。
	DeclineReason V1ArchiveDeclineReason `json:"decline_reason"`
	HasAct bool `json:"has_act"` // `true`，前提是已生成交接单。
	ID int64 `json:"id"` // 档案记录编号。
	ReceiveDate string `json:"receive_date"` // 交货接收日期和时间。
	SupplyID string `json:"supply_id"` // 交货标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	ActFileUuid string `json:"act_file_uuid"` // 验收证明书标识符。
	CreatedDate string `json:"created_date"` // 交货申请创建日期和时间。
	HasLabel bool `json:"has_label"` // `true`，前提是已生成标签。
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	Status V1ArchiveStatusEnum `json:"status"`
}

type V1FbpArchiveListResponseItem struct {
	CreatedDate string `json:"created_date"` // 交货申请创建日期。
	ExternalOrderID string `json:"external_order_id"` // 合作仓库自身系统已完成交货的标识符。
	HasAct bool `json:"has_act"` // `true`，前提是已生成交接单。
	HasLabel bool `json:"has_label"` // `true`，前提是已生成标签。
	OrderDraftID int64 `json:"order_draft_id"` // 交货草稿标识符。
	ReceiveDate string `json:"receive_date"` // 交货接收日期和时间。
	Status V1ArchiveStatusEnum `json:"status"`
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	ActFileUuid string `json:"act_file_uuid"` // 验收证明书标识符。
	BundleID string `json:"bundle_id"` // 已验证商品清单的标识符。
	BundleSKUSummary V1ArchiveSkuSummary `json:"bundle_sku_summary"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	DeclineReason V1ArchiveDeclineReason `json:"decline_reason"`
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	SupplyID string `json:"supply_id"` // 交货标识符。
	WhcOrderID int64 `json:"whc_order_id"` // 合作仓库已完成交货的标识符。
}

type V1FbpDraftPickupCreateResponse struct {
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectGetTimeslotRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库标识符。
	BundleID string `json:"bundle_id"` // 已验证商品清单的标识符。
	IntervalEnd string `json:"interval_end"` // 可用时间段所需区间的结束日期。
	IntervalStart string `json:"interval_start"` // 可用时间段所需区间的开始日期。
}

// 订单状态： - `ORDER_STATUS_UNSPECIFIED`——未指定； - `READY_TO_SUPPLY`——准备发运； - `FILLING_DELIVERY_DETAILS`——填写交货数据； - `COURIER_ASS...
type V1OrderStatusEnum string

// 警告原因：
type V1FbpOrderGetResponseAttentionReasonsEnum string
const (
	V1FbpOrderGetResponseAttentionReasonsEnum_ORDERATTENTIONTYPEUNSPECIFIED V1FbpOrderGetResponseAttentionReasonsEnum = "ORDER_ATTENTION_TYPE_UNSPECIFIED"
	V1FbpOrderGetResponseAttentionReasonsEnum_OLD V1FbpOrderGetResponseAttentionReasonsEnum = "OLD"
	V1FbpOrderGetResponseAttentionReasonsEnum_TIMESLOTEXPIRED V1FbpOrderGetResponseAttentionReasonsEnum = "TIME_SLOT_EXPIRED"
)

type V1FbpOrderGetResponse struct {
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	HasConsignmentNote bool `json:"has_consignment_note"` // `true`，如果有已签署的文件。
	ID int64 `json:"id"` // 交货申请标识符。
	Locked bool `json:"locked"` // `true`，如果无法编辑交货。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
	BundleUuid string `json:"bundle_uuid"` // 组成商品标识符。
	CancellationState V1CancellationState `json:"cancellation_state"`
	CreatedDate string `json:"created_date"` // 交货创建日期。
	HasLabel bool `json:"has_label"` // `true`，如果有标签。
	OrderNumber string `json:"order_number"` // 交货编号。
	ReceiveDate string `json:"receive_date"` // 交货接收日期和时间。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	Status V1OrderStatusEnum `json:"status"`
	CanBeCancelled bool `json:"can_be_cancelled"` // `true`，如果申请可以取消。
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	AttentionReasons []V1OrderAttentionTypeEnum `json:"attention_reasons"` // 警告原因： - `ORDER_ATTENTION_TYPE_UNSPECIFIED`——未指定； - `OLD`——过期申请； - `TIME_SLOT_EXPIRED`——时间段已过期。
}

type V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason string

// 缺少时间段的原因：
type V1FbpDraftDirectGetTimeslotResponseReasonsEnum string
const (
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_EMPTYTIMESLOTSREASONUNSPECIFIED V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "EMPTY_TIMESLOTS_REASON_UNSPECIFIED"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_LOGISTICSUNKNOWN V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "LOGISTICS_UNKNOWN"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_NOROUTE V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "NO_ROUTE"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_NOROUTESCHEDULES V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "NO_ROUTE_SCHEDULES"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_NOLOGISTICSCAPACITY V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "NO_LOGISTICS_CAPACITY"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_SCHEDULEUNKNOWN V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "SCHEDULE_UNKNOWN"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_NOTENOUGHCAPACITY V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "NOT_ENOUGH_CAPACITY"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_NOTENOUGHTRUCKS V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "NOT_ENOUGH_TRUCKS"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_LIMITSNOTAVAILABLE V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "LIMITS_NOT_AVAILABLE"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_CROSSDOCKRESERVEMISSING V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "CROSS_DOCK_RESERVE_MISSING"
	V1FbpDraftDirectGetTimeslotResponseReasonsEnum_SCHEDULERESERVEMISSING V1FbpDraftDirectGetTimeslotResponseReasonsEnum = "SCHEDULE_RESERVE_MISSING"
)

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons []V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason `json:"reasons"` // 缺少时间段的原因： - `EMPTY_TIMESLOTS_REASON_UNSPECIFIED`——未定义； - `LOGISTICS_UNKNOWN`——物流方未知错误； - `NO_ROUTE`——没有路线； - `NO_ROUTE_S...
	Timeslots []V1FbpDraftDirectGetTimeslotResponseTimeslot `json:"timeslots"` // 可用时间段列表。
	WarehouseTimezoneName string `json:"warehouse_timezone_name"` // 卖家仓库的时区。
}

type V1FbpDraftDirectCreateRequest struct {
	BundleID string `json:"bundle_id"` // 已校验商品列表的标识符。要获取，请使用方法[/v1/fbp/draft/direct/product/validate](#operation/FbpDraftDirectProductValidate)。
	DeliveryDetails V1FbpDraftDirectCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 包装单位数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpGetLabelResponse struct {
	State FbpGetLabelResponseLabelCreationStateTypeEnum `json:"state"`
	LabelURL string `json:"label_url"` // 交货标签链接。
}

// 错误原因：
type V1FbpEditTimeslotResponseErrorReasonsEnum string
const (
	V1FbpEditTimeslotResponseErrorReasonsEnum_RESERVEFAILURETYPEUNSPECIFIED V1FbpEditTimeslotResponseErrorReasonsEnum = "RESERVE_FAILURE_TYPE_UNSPECIFIED"
	V1FbpEditTimeslotResponseErrorReasonsEnum_REQUESTVALIDATION V1FbpEditTimeslotResponseErrorReasonsEnum = "REQUEST_VALIDATION"
	V1FbpEditTimeslotResponseErrorReasonsEnum_INVALIDRESERVE V1FbpEditTimeslotResponseErrorReasonsEnum = "INVALID_RESERVE"
	V1FbpEditTimeslotResponseErrorReasonsEnum_LOGISTICSREASON V1FbpEditTimeslotResponseErrorReasonsEnum = "LOGISTICS_REASON"
	V1FbpEditTimeslotResponseErrorReasonsEnum_SCHEDULEREASON V1FbpEditTimeslotResponseErrorReasonsEnum = "SCHEDULE_REASON"
	V1FbpEditTimeslotResponseErrorReasonsEnum_NOCAPACITY V1FbpEditTimeslotResponseErrorReasonsEnum = "NO_CAPACITY"
)

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []V1FbpEditTimeslotResponseReserveFailureType `json:"error_reasons"` // 错误原因： - `RESERVE_FAILURE_TYPE_UNSPECIFIED`——未定义； - `REQUEST_VALIDATION`——请求中填写了过去的预定日期； - `INVALID_RESERVE`——原始预留未找到、已失效...
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 配送详细信息。
type V1fbpDeliveryDetails struct {
	SupplyType DeliveryDetailsSupplyType `json:"supply_type"`
	DirectDetails V1DeliveryDetailsDirectDetails `json:"direct_details"`
	DropOffPoint DeliveryDetailsDropOffPointDetails `json:"drop_off_point"`
	PickupDetails V1DeliveryDetailsPickUpDetails `json:"pickup_details"`
}

type V1FbpDraftListResponseItem struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	CancellationState V1CancellationState `json:"cancellation_state"`
	CreatedAt string `json:"created_at"` // 草稿创建日期。
	DeliveryDetails V1fbpDeliveryDetails `json:"delivery_details"`
	ID int64 `json:"id"` // 草稿标识符。
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	Status V1DraftStatusEnum `json:"status"`
	BundleID string `json:"bundle_id"` // 验证后商品的列表标识符。
	DeletedAt string `json:"deleted_at"` // 草稿删除日期。
	Editable bool `json:"editable"` // `true`，如果草稿可以修改。
	IsCancelable bool `json:"is_cancelable"` // `true`，如果草稿可以取消。
	IsDeletable bool `json:"is_deletable"` // `true`，如果草稿可以删除。
	Locked bool `json:"locked"` // `true`，如果草稿被封锁。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

// 休息时间信息。
type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak struct {
	TimeslotEnd string `json:"timeslot_end"` // 开始时间。
	TimeslotStart string `json:"timeslot_start"` // 结束时间。
}

type V1FbpCheckConsignmentNoteStateRequest struct {
	Code string `json:"code"` // 货物运单标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails V1FbpOrderPickUpDlvEditRequestPickUpDetails `json:"pickup_details"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderDirectCancelResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type FbpCreateActResponseCreateActErrorReason string

type V1FbpDraftDropOffProductValidateRequestSkuItem struct {
	Count int32 `json:"count"` // 数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符—— SKU。
}

type V1FbpDraftDirectTplDlvCreateResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
	DraftID int64 `json:"draft_id"` // 草稿标识符。
}

// 订单详情。
// 警告原因：
type V1FbpOrderListResponseItemAttentionReasonsEnum string
const (
	V1FbpOrderListResponseItemAttentionReasonsEnum_ORDERATTENTIONTYPEUNSPECIFIED V1FbpOrderListResponseItemAttentionReasonsEnum = "ORDER_ATTENTION_TYPE_UNSPECIFIED"
	V1FbpOrderListResponseItemAttentionReasonsEnum_OLD V1FbpOrderListResponseItemAttentionReasonsEnum = "OLD"
	V1FbpOrderListResponseItemAttentionReasonsEnum_TIMESLOTEXPIRED V1FbpOrderListResponseItemAttentionReasonsEnum = "TIME_SLOT_EXPIRED"
)

type V1FbpOrderListResponseItem struct {
	OrderNumber string `json:"order_number"` // 交货编号。
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	HasLabel bool `json:"has_label"` // `true`，如果有标签。
	ID int64 `json:"id"` // 交货申请标识符。
	Locked bool `json:"locked"` // `true`，如果无法编辑交货。
	Status V1OrderStatusEnum `json:"status"`
	AttentionReasons []V1OrderAttentionTypeEnum `json:"attention_reasons"` // 警告原因： - `ORDER_ATTENTION_TYPE_UNSPECIFIED`——未指定； - `OLD`——过期申请； - `TIME_SLOT_EXPIRED`——时间段已过期。
	BundleSummary ItemBundleSummary `json:"bundle_summary"`
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	ReceiveDate string `json:"receive_date"` // 交货接收日期和时间。
	CreatedDate string `json:"created_date"` // 交货创建日期。
	CanBeCancelled bool `json:"can_be_cancelled"` // `true`，如果申请可以取消。
	CancellationState V1CancellationState `json:"cancellation_state"`
	HasConsignmentNote bool `json:"has_consignment_note"` // `true`，如果有已签署的文件。
}

type V1FbpOrderPickUpDlvEditResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

// 拒绝原因：
type V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum string
const (
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_BUNDLEITEMERRORUNSPECIFIED V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "BUNDLE_ITEM_ERROR_UNSPECIFIED"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_OUTOFASSORTMENT V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "OUT_OF_ASSORTMENT"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_INVALID V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "INVALID"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_INCOMPATIBLEWAREHOUSE V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "INCOMPATIBLE_WAREHOUSE"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_INVALIDBARCODE V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "INVALID_BARCODE"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_MULTIPLICITY V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "MULTIPLICITY"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_NOPRICE V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "NO_PRICE"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_BANNED V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "BANNED"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_ZEROQUANTITY V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "ZERO_QUANTITY"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_QUANTITYGREATERTHENMAX V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "QUANTITY_GREATER_THEN_MAX"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_NOSALES V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "NO_SALES"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_SURPLUS V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "SURPLUS"
	V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum_AVAILABILITYISEMPTY V1FbpDraftDirectProductValidateResponseRejectedItemRejectionReasonsEnum = "AVAILABILITY_IS_EMPTY"
)

type V1FbpDraftDirectProductValidateResponseRejectedItem struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品货号。
	Quantity int32 `json:"quantity"` // 商品数量。
	RejectionReasons []V1BundleItemErrorEnum `json:"rejection_reasons"` // 拒绝原因： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——未找到商品； - `INVALID`——商品未创建； - `INCOMPATIBLE_WAREHOUS...
	SKU int64 `json:"sku"` // 商品标识符（SKU）。
	Volume float64 `json:"volume"` // 商品体积。
	Barcode string `json:"barcode"` // 条形码。
	IconName string `json:"icon_name"` // 商品图片链接。
	Name string `json:"name"` // 商品名称。
}

// 配送详细信息。
type V1FbpDraftPickupCreateRequestDeliveryDetails struct {
	SenderName string `json:"sender_name"` // 发件人姓名。
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
	Address string `json:"address"` // 地址。
	Comment string `json:"comment"` // 备注。
	Date string `json:"date"` // 送货日期。
}

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"` // 送货日期。
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	DropOffProvinceUuid string `json:"drop_off_province_uuid"` // 省份唯一标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"` // 货物运单标识符。
}

type V1FbpDraftDropOffRegistrateRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 错误：
type V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum string
const (
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_BUNDLEITEMERRORUNSPECIFIED V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "BUNDLE_ITEM_ERROR_UNSPECIFIED"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_OUTOFASSORTMENT V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "OUT_OF_ASSORTMENT"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INVALID V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INVALID"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INCOMPATIBLEWAREHOUSE V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INCOMPATIBLE_WAREHOUSE"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_INVALIDBARCODE V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "INVALID_BARCODE"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_MULTIPLICITY V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "MULTIPLICITY"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_NOPRICE V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "NO_PRICE"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_BANNED V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "BANNED"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_ZEROQUANTITY V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "ZERO_QUANTITY"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_QUANTITYGREATERTHENMAX V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "QUANTITY_GREATER_THEN_MAX"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_NOSALES V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "NO_SALES"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_SURPLUS V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "SURPLUS"
	V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum_AVAILABILITYISEMPTY V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleErrorErrorsEnum = "AVAILABILITY_IS_EMPTY"
)

type V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError struct {
	Errors []V1BundleItemErrorEnum `json:"errors"` // 错误： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——未找到商品； - `INVALID`——商品未创建； - `INCOMPATIBLE_WAREHOUSE`...
	SKU int64 `json:"sku"` // 商品标识符（SKU）。
}

// 错误。
type V1FbpDraftDirectRegistrateResponseRegistrationError struct {
	BundleErrors []V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError `json:"bundle_errors"` // 校验商品列表的错误。
	OrderError V1OrderErrorTypeEnum `json:"order_error"`
}

type V1FbpDraftDirectRegistrateResponse struct {
	Error V1FbpDraftDirectRegistrateResponseRegistrationError `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason string

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 日期信息。
type V1FbpOrderDropOffTimetableResponseCalendarCalendarItem struct {
	BreakHours V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"` // `true`，表示休息日。
	OpeningHours V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime `json:"opening_hours"`
}

// 生成状态： - `STATUS_UNSPECIFIED` ——未定义； - `NOT_EXIST` ——不存在； - `PROCESSING` ——处理中； - `EXIST` ——已完成； - `ERROR` ——错误。
type V1FbpCheckActStateResponseStatus string

type V1FbpCheckActStateResponse struct {
	CdnURL string `json:"cdn_url"` // 验收证明书链接。
	Error FbpCheckActStateResponseErrorReason `json:"error"`
	Status V1FbpCheckActStateResponseStatus `json:"status"`
}

type V1FbpDraftDirectProductValidateResponse struct {
	ApprovedItems []V1FbpDraftDirectProductValidateResponseApprovedItem `json:"approved_items"` // 已确认商品。
	BundleGenerated bool `json:"bundle_generated"` // `true`，前提是已创建校验商品列表。
	BundleID string `json:"bundle_id"` // 校验商品列表标识符。
	RejectedItems []V1FbpDraftDirectProductValidateResponseRejectedItem `json:"rejected_items"` // 被拒绝的商品。
}

type V1FbpOrderDropOffTimetableResponseCalendar struct {
	CalendarItem V1FbpOrderDropOffTimetableResponseCalendarCalendarItem `json:"calendar_item"`
	DayOfWeek V1DayOfWeek `json:"day_of_week"`
}

type V1FbpDraftDropOffPointListRequest struct {
	NextPageNumber int32 `json:"next_page_number"` // 下一页页码。
	PageSize int32 `json:"page_size"` // 每页包含的商品数量。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus []V1FbpDraftPickUpProductValidateRequestSkuItem `json:"skus"` // 商品标识符（SKU）列表。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
}

type V1FbpDraftGetResponse struct {
	BundleID string `json:"bundle_id"` // 验证后商品的列表标识符。
	DeliveryDetails V1fbpDeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	CreatedAt string `json:"created_at"` // 草稿创建日期。
	DeletedAt string `json:"deleted_at"` // 草稿删除日期。
	ID int64 `json:"id"` // 草稿标识符。
	IsCancelable bool `json:"is_cancelable"` // `true`，如果草稿可以取消。
	IsDeletable bool `json:"is_deletable"` // `true`，如果草稿可以删除。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	Status V1DraftStatusEnum `json:"status"`
	SupplyID string `json:"supply_id"` // 交货标识符。
	DeclineReason FbpDraftGetResponseDeclineReason `json:"decline_reason"`
	Editable bool `json:"editable"` // `true`，如果草稿可以修改。
	CancellationState V1CancellationState `json:"cancellation_state"`
	IsRegistrationAvailable bool `json:"is_registration_available"` // `true`，如果可注册。
	Locked bool `json:"locked"` // `true`，如果草稿被封锁。
}

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
}

type V1FbpArchiveListResponse struct {
	HasNext bool `json:"has_next"` // `true`，前提是本次响应未返回所有数据。
	Items []V1FbpArchiveListResponseItem `json:"items"` // 已完成交货。
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。
}

type V1FbpDraftDropOffProductValidateResponse struct {
	BundleID string `json:"bundle_id"` // 验证后的商品列表标识符。
	RejectedItems []V1FbpDraftDropOffProductValidateResponseRejectedItem `json:"rejected_items"` // 被拒绝的商品。
	ApprovedItems []V1FbpDraftDropOffProductValidateResponseApprovedItem `json:"approved_items"` // 已接收的商品。
	BundleGenerated bool `json:"bundle_generated"` // `true`，前提是已创建商品成分信息。
}

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []FbpDraftDropOffPointListResponseDropOffPoint `json:"drop_off_points"` // 接收点列表。
}

// 用于搜索货件的筛选器。
type PostingV1PostingFbpListRequestFilter struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符，即货号。
	PostingNumbers []string `json:"posting_numbers"` // 货件编号。
	Since string `json:"since"` // 时间段开始。
	Statuses []string `json:"statuses"` // 货件状态。
	To string `json:"to"` // 时间段结束。
	Name string `json:"name"` // 商品名称。
}

// 货件排序参数：
type PostingV1PostingFbpListRequestSortByEnum string
const (
	PostingV1PostingFbpListRequestSortByEnum_LastChangeStatusDate PostingV1PostingFbpListRequestSortByEnum = "last_change_status_date"
	PostingV1PostingFbpListRequestSortByEnum_InProcessAt PostingV1PostingFbpListRequestSortByEnum = "in_process_at"
)

type PostingV1PostingFbpListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter PostingV1PostingFbpListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortBy PostingV1PostingFbpListRequestSortByEnum `json:"sort_by"` // 货件排序参数： - `last_change_status_date`——按最后一次状态变更日期排序； - `in_process_at`——按开始处理日期排序。
	SortDir PostingV1PostingFbpListRequestSortDirEnum `json:"sort_dir"`
}

type Fbpv1Timeslot struct {
	TimeslotEnd string `json:"timeslot_end"` // 时间段结束日期。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始日期。
}

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDropOffProductValidateRequest struct {
	Skus []V1FbpDraftDropOffProductValidateRequestSkuItem `json:"skus"` // Ozon系统中的商品标识符—— SKU。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDropOffCreateResponse struct {
	SupplyID string `json:"supply_id"` // 交货申请标识符。
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftPickupCreateRequest struct {
	BundleID string `json:"bundle_id"` // 已校验商品列表的标识符。
	DeliveryDetails V1FbpDraftPickupCreateRequestDeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 包装单位数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

// 错误原因：
type V1FbpCreateActResponseErrorsEnum string
const (
	V1FbpCreateActResponseErrorsEnum_CREATEACTERRORREASONUNSPECIFIED V1FbpCreateActResponseErrorsEnum = "CREATE_ACT_ERROR_REASON_UNSPECIFIED"
	V1FbpCreateActResponseErrorsEnum_INVALIDORDERTYPE V1FbpCreateActResponseErrorsEnum = "INVALID_ORDER_TYPE"
)

type V1FbpCreateActResponse struct {
	Errors []FbpCreateActResponseCreateActErrorReason `json:"errors"` // 错误原因： - `CREATE_ACT_ERROR_REASON_UNSPECIFIED` ——未定义； - `INVALID_ORDER_TYPE` ——无法为指定标识符创建验收证明书。
	FileUuid string `json:"file_uuid"` // 验收证明书标识符。
	IsSuccess bool `json:"is_success"` // `true`，前提是请求中没有错误。
}

type V1FbpOrderListResponse struct {
	HasNext bool `json:"has_next"` // `true`，如果响应中未返回所有交货。
	Items []V1FbpOrderListResponseItem `json:"items"` // 交货。
	LastID int64 `json:"last_id"` // 页面上最后一次交货的标识符。
}

type V1FbpArchiveListRequest struct {
	LastID string `json:"last_id"` // 页面上最后一个值的标识符。首次请求时请留空。 如需获取后续数据，请填写上次响应中的 `last_id`。
	Count string `json:"count"` // 响应中的元素数量。
}

// 配送详情。
type V1FbpDraftDropOffCreateRequestDeliveryDetails struct {
	DropOffDate string `json:"drop_off_date"` // 送货日期。
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	DropOffProvinceUuid string `json:"drop_off_province_uuid"` // 省份唯一标识符。
}

type V1FbpDraftListResponse struct {
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。
	HasNext bool `json:"has_next"` // `true`，如果响应中没有返回所有值。
	Items []V1FbpDraftListResponseItem `json:"items"` // 草稿。
}

// 缺少时间段的原因：
type V1FbpAvailableTimeslotListResponseReasonsEnum string
const (
	V1FbpAvailableTimeslotListResponseReasonsEnum_EMPTYTIMESLOTSREASONUNSPECIFIED V1FbpAvailableTimeslotListResponseReasonsEnum = "EMPTY_TIMESLOTS_REASON_UNSPECIFIED"
	V1FbpAvailableTimeslotListResponseReasonsEnum_LOGISTICSUNKNOWN V1FbpAvailableTimeslotListResponseReasonsEnum = "LOGISTICS_UNKNOWN"
	V1FbpAvailableTimeslotListResponseReasonsEnum_NOROUTE V1FbpAvailableTimeslotListResponseReasonsEnum = "NO_ROUTE"
	V1FbpAvailableTimeslotListResponseReasonsEnum_NOROUTESCHEDULES V1FbpAvailableTimeslotListResponseReasonsEnum = "NO_ROUTE_SCHEDULES"
	V1FbpAvailableTimeslotListResponseReasonsEnum_NOLOGISTICSCAPACITY V1FbpAvailableTimeslotListResponseReasonsEnum = "NO_LOGISTICS_CAPACITY"
	V1FbpAvailableTimeslotListResponseReasonsEnum_SCHEDULEUNKNOWN V1FbpAvailableTimeslotListResponseReasonsEnum = "SCHEDULE_UNKNOWN"
	V1FbpAvailableTimeslotListResponseReasonsEnum_NOTENOUGHCAPACITY V1FbpAvailableTimeslotListResponseReasonsEnum = "NOT_ENOUGH_CAPACITY"
	V1FbpAvailableTimeslotListResponseReasonsEnum_NOTENOUGHTRUCKS V1FbpAvailableTimeslotListResponseReasonsEnum = "NOT_ENOUGH_TRUCKS"
	V1FbpAvailableTimeslotListResponseReasonsEnum_LIMITSNOTAVAILABLE V1FbpAvailableTimeslotListResponseReasonsEnum = "LIMITS_NOT_AVAILABLE"
	V1FbpAvailableTimeslotListResponseReasonsEnum_CROSSDOCKRESERVEMISSING V1FbpAvailableTimeslotListResponseReasonsEnum = "CROSS_DOCK_RESERVE_MISSING"
	V1FbpAvailableTimeslotListResponseReasonsEnum_SCHEDULERESERVEMISSING V1FbpAvailableTimeslotListResponseReasonsEnum = "SCHEDULE_RESERVE_MISSING"
)

type V1FbpAvailableTimeslotListResponse struct {
	Reasons []V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason `json:"reasons"` // 缺少时间段的原因： - `EMPTY_TIMESLOTS_REASON_UNSPECIFIED`——未定义； - `LOGISTICS_UNKNOWN`——物流方未知错误； - `NO_ROUTE`——没有路线； - `NO_ROUTE_S...
	Timeslots []Fbpv1Timeslot `json:"timeslots"` // 可用时间段列表。
	WarehouseTimezoneName string `json:"warehouse_timezone_name"` // 卖家仓库的时区。
}

type V1FbpDraftDirectRegistrateRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type PostingV1PostingFbpListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Postings []PostingV1PostingFbpListResponsePostings `json:"postings"` // 货件列表。
}

type V1FbpDraftDropOffCreateRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库标识符。
	BundleID string `json:"bundle_id"` // 验证后的商品列表标识符。
	DeliveryDetails V1FbpDraftDropOffCreateRequestDeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"` // 供货申请标识符。
}

type V1FbpOrderDropOffTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDirectCreateResponse struct {
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftPickupDlvEditRequest struct {
	PickupDetails V1FbpDraftPickupDlvEditRequestDeliveryDetails `json:"pickup_details"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectSellerDlvEditRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
	VehicleType string `json:"vehicle_type"` // 车辆类型。
	DriverName string `json:"driver_name"` // 司机姓名。
}

type V1FbpEditTimeslotRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []V1FbpOrderDropOffTimetableResponseCalendar `json:"calendar"` // 接收点的营业时间信息。
}

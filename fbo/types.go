package fbo

type V1FbpDraftPickUpProductValidateRequestSkuItem struct {
	Count int32 `json:"count"` // 交货商品数量。
	SKU   int64 `json:"sku"`   // 商品标识符（SKU）。
}

type V1BundleItemErrorEnum string

// Errors values
type Errors string

const (
	ErrorsBundleItemErrorUnspecified Errors = "BUNDLE_ITEM_ERROR_UNSPECIFIED" // 未指定；
	ErrorsOUTOFAssortment            Errors = "OUT_OF_ASSORTMENT"             // 商品不在交货品类中；
	ErrorsInvalid                    Errors = "INVALID"                       // 状态不正确；
	ErrorsIncompatibleWarehouse      Errors = "INCOMPATIBLE_WAREHOUSE"        // 仓库标识符不正确；
	ErrorsInvalidBarcode             Errors = "INVALID_BARCODE"               // 未填写条形码；
	ErrorsMultiplicity               Errors = "MULTIPLICITY"                  // 商品数量与包装不成倍数关系；
	ErrorsNOPrice                    Errors = "NO_PRICE"                      // 未填写价格；
	ErrorsBanned                     Errors = "BANNED"                        // 商品被封禁；
	ErrorsZeroQuantity               Errors = "ZERO_QUANTITY"                 // 商品数量不能为 0；
	ErrorsQuantityGreaterThenMAX     Errors = "QUANTITY_GREATER_THEN_MAX"     // 商品数量超过单个SKU的最大值；
	ErrorsNOSales                    Errors = "NO_SALES"                      // 商品超过60天无销售；
	ErrorsSurplus                    Errors = "SURPLUS"                       // 仓库中的商品库存可供销售90天；
	ErrorsAvailabilityISEmpty        Errors = "AVAILABILITY_IS_EMPTY"         // 无商品可售性信息。
)

type V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError struct {
	Errors []V1BundleItemErrorEnum `json:"errors"` // 错误： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——商品不在交货品类中； - `INVALID`——状态不正确； - `INCOMPATIBLE_WAREHO...
	SKU    int64                   `json:"sku"`    // Ozon 系统中的商品标识符（SKU）。
}

// 营业时间信息。
type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd   string `json:"timeslot_end"`   // 开始时间。
	TimeslotStart string `json:"timeslot_start"` // 结束时间。
}

type V1FbpOrderListRequest struct {
	Count  int32 `json:"count"`   // 响应中的交货数量。
	LastID int64 `json:"last_id"` // 页面上最后一次交货的标识符。首次请求时请将此字段留空。 如需获取后续数据，请填写上一次请求响应中最后一次交货的`id`。
}

type FbpCreateActResponseCreateActErrorReason string

// 时间段。
type V1fbpTimeslot struct {
	TimeslotEnd   string `json:"timeslot_end"`   // 时间段结束时间（UTC）。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间（UTC）。
}

// 揽收点详情。
type DeliveryDetailsDropOffPointDetails struct {
	ProvinceUuid string        `json:"province_uuid"` // 区域唯一标识符。
	Timeslot     V1fbpTimeslot `json:"timeslot"`
	ID           int64         `json:"id"` // 揽收点标识符。
}

// 取货点详情。
type V1DeliveryDetailsPickUpDetails struct {
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
	Address     string `json:"address"`      // 地址。
	Comment     string `json:"comment"`      // 备注。
	Date        string `json:"date"`         // 送货日期。
	SenderName  string `json:"sender_name"`  // 发件人姓名。
}

// 交货类型： - `SUPPLY_TYPE_UNSPECIFIED`：未指定； - `DIRECT_BY_SELLER`：卖家自行送货到仓库； - `DIRECT_BY_TPL`：第三方物流公司送货到仓库； - `DROP_OFF`：送货到揽...
type DeliveryDetailsSupplyType string

// 卖家自配送详情。
type DirectDetailsBySellerDetails struct {
	DriverName                string `json:"driver_name"`                 // 司机姓名。
	VehicleRegistrationNumber string `json:"vehicle_registration_number"` // 车牌号码。
	VehicleType               string `json:"vehicle_type"`                // 运输工具类型。
}

// 第三方物流公司配送详情。
type DirectDetailsByTplDetails struct {
	TrackingNumber       string `json:"tracking_number"`        // 货件跟踪号码。
	TransportCompanyName string `json:"transport_company_name"` // 物流公司名称。
}

// 交货时间段详情。
type DirectDetailsTimeslotDetails struct {
	Timeslot              V1fbpTimeslot `json:"timeslot"`
	TimeslotReservationID string        `json:"timeslot_reservation_id"` // 交货时间段预定标识符。
}

// 卖家配送详情。
type V1DeliveryDetailsDirectDetails struct {
	BySellerDetails DirectDetailsBySellerDetails `json:"by_seller_details"`
	ByTPLDetails    DirectDetailsByTplDetails    `json:"by_tpl_details"`
	TimeslotDetails DirectDetailsTimeslotDetails `json:"timeslot_details"`
}

// 配送详细信息。
type V1fbpDeliveryDetails struct {
	DropOffPoint  DeliveryDetailsDropOffPointDetails `json:"drop_off_point"`
	PickupDetails V1DeliveryDetailsPickUpDetails     `json:"pickup_details"`
	SupplyType    DeliveryDetailsSupplyType          `json:"supply_type"`
	DirectDetails V1DeliveryDetailsDirectDetails     `json:"direct_details"`
}

type V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason string

type Fbpv1Timeslot struct {
	TimeslotStart string `json:"timeslot_start"` // 时间段开始日期。
	TimeslotEnd   string `json:"timeslot_end"`   // 时间段结束日期。
}

// Reasons values
type Reasons string

const (
	ReasonsEmptyTimeslotsReasonUnspecified Reasons = "EMPTY_TIMESLOTS_REASON_UNSPECIFIED" // 未定义；
	ReasonsLogisticsUnknown                Reasons = "LOGISTICS_UNKNOWN"                  // 物流方未知错误；
	ReasonsNORoute                         Reasons = "NO_ROUTE"                           // 没有路线；
	ReasonsNORouteSchedules                Reasons = "NO_ROUTE_SCHEDULES"                 // 路线上没有排期；
	ReasonsNOLogisticsCapacity             Reasons = "NO_LOGISTICS_CAPACITY"              // 路线上可用的时段不足；
	ReasonsScheduleUnknown                 Reasons = "SCHEDULE_UNKNOWN"                   // 排期方未知错误；
	ReasonsNOTEnoughCapacity               Reasons = "NOT_ENOUGH_CAPACITY"                // 仓库可用时段不足；
	ReasonsNOTEnoughTrucks                 Reasons = "NOT_ENOUGH_TRUCKS"                  // 车辆车位不足；
	ReasonsLimitsNOTAvailable              Reasons = "LIMITS_NOT_AVAILABLE"               // 仓库未设置限制；
	ReasonsCrossDockReserveMissing         Reasons = "CROSS_DOCK_RESERVE_MISSING"         // 仓库未预留越库配送容量；
	ReasonsScheduleReserveMissing          Reasons = "SCHEDULE_RESERVE_MISSING"           // 缺少必要的排期预留。
)

type V1FbpAvailableTimeslotListResponse struct {
	Reasons               []V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason `json:"reasons"`                 // 缺少时间段的原因： - `EMPTY_TIMESLOTS_REASON_UNSPECIFIED`——未定义； - `LOGISTICS_UNKNOWN`——物流方未知错误； - `NO_ROUTE`——没有路线； - `NO_ROUTE_S...
	Timeslots             []Fbpv1Timeslot                                          `json:"timeslots"`               // 可用时间段列表。
	WarehouseTimezoneName string                                                   `json:"warehouse_timezone_name"` // 卖家仓库的时区。
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	VehicleType   string `json:"vehicle_type"`   // 车辆类型。
	DriverName    string `json:"driver_name"`    // 司机姓名。
	RowVersion    int64  `json:"row_version"`    // 草稿的当前版本标识符。
	SupplyID      string `json:"supply_id"`      // 供货申请标识符。
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
}

// 配送详细信息。
type V1FbpDraftDirectTplDlvCreateRequestDirectDetails struct {
	TimeslotStart        string `json:"timeslot_start"`         // 时间段开始本地时间。
	TrackingNumber       string `json:"tracking_number"`        // 货件跟踪号码。
	TransportCompanyName string `json:"transport_company_name"` // 物流公司名称。
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	BundleID          string                                           `json:"bundle_id"` // 套装标识符。
	DeliveryDetails   V1FbpDraftDirectTplDlvCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32                                            `json:"package_units_count"` // 货位数量。
	WarehouseID       int64                                            `json:"warehouse_id"`        // 仓库标识符。
}

type V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason string

// 配送详细信息。
type V1FbpDraftDirectCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"` // 配送时间段开始。
}

// 取消错误代码： - `CODE_UNSPECIFIED`——未指定； - `NO_RESPONSE_FROM_3PF`——取消申请未确认，未收到第三方响应； - `ACCEPTANCE_ALREADY_STARTED`——取消申请未确认，已...
type CancellationErrorCode string

// 取消错误。
type CancellationStateCancellationError struct {
	ErrorCode CancellationErrorCode `json:"error_code"`
	Message   string                `json:"message"` // 错误描述。
}

// 草稿状态: - `DRAFT_STATUS_UNSPECIFIED` — 未定义; - `NEW` — 新的; - `SUPPLY_VARIANT_CONFIRMATION` — 等待确认; - `SUPPLY_NOT_CONFIRMED`...
type V1DraftStatusEnum string

// 错误状态： - `STATUS_UNSPECIFIED`——未指定； - `CONFIRMATION`——等待申请取消确认； - `CANCELED`——取消已确认； - `NOT_CANCELED`——未收到取消确认。
type V1CancellationStateStatus string

// 取消原因。
type V1CancellationState struct {
	CancellationError  CancellationStateCancellationError `json:"cancellation_error"`
	CancellationStatus V1CancellationStateStatus          `json:"cancellation_status"`
}

type V1FbpDraftListResponseItem struct {
	ID                int64                `json:"id"`            // 草稿标识符。
	IsCancelable      bool                 `json:"is_cancelable"` // `true`，如果草稿可以取消。
	Status            V1DraftStatusEnum    `json:"status"`
	SupplyID          string               `json:"supply_id"`    // 交货标识符。
	WarehouseID       int64                `json:"warehouse_id"` // 仓库标识符。
	BundleID          string               `json:"bundle_id"`    // 验证后商品的列表标识符。
	CreatedAt         string               `json:"created_at"`   // 草稿创建日期。
	DeletedAt         string               `json:"deleted_at"`   // 草稿删除日期。
	DeliveryDetails   V1fbpDeliveryDetails `json:"delivery_details"`
	Editable          bool                 `json:"editable"`            // `true`，如果草稿可以修改。
	IsDeletable       bool                 `json:"is_deletable"`        // `true`，如果草稿可以删除。
	Locked            bool                 `json:"locked"`              // `true`，如果草稿被封锁。
	PackageUnitsCount int32                `json:"package_units_count"` // 货位数量。
	CancellationState V1CancellationState  `json:"cancellation_state"`
}

type V1FbpDraftListResponse struct {
	HasNext bool                         `json:"has_next"` // `true`，如果响应中没有返回所有值。
	Items   []V1FbpDraftListResponseItem `json:"items"`    // 草稿。
	LastID  int64                        `json:"last_id"`  // 页面上最后一个值的标识符。
}

// 配送详情。
type Fbpv1DeliveryDetails struct {
	DirectDetails V1DeliveryDetailsDirectDetails     `json:"direct_details"`
	DropOffPoint  DeliveryDetailsDropOffPointDetails `json:"drop_off_point"`
	PickupDetails V1DeliveryDetailsPickUpDetails     `json:"pickup_details"`
	SupplyType    DeliveryDetailsSupplyType          `json:"supply_type"`
}

// 拒绝交货的原因。
// Code values
type Code string

const (
	CodeDeclineReasonCodeUnspecified       Code = "DECLINE_REASON_CODE_UNSPECIFIED"
	CodeCannotCreateSupplyONTPF            Code = "CANNOT_CREATE_SUPPLY_ON_TPF"
	CodeDropOFFPointClosed                 Code = "DROP_OFF_POINT_CLOSED"
	CodeCodeSupplyLost                     Code = "CODE_SUPPLY_LOST"
	CodeCourierPickUPRejectedBYSeller      Code = "COURIER_PICK_UP_REJECTED_BY_SELLER"
	CodeBondedDocumentsRejectedBYWarehouse Code = "BONDED_DOCUMENTS_REJECTED_BY_WAREHOUSE"
)

type V1ArchiveDeclineReason struct {
	Code    Code   `json:"code"`    // 拒绝交货原因代码： - `DECLINE_REASON_CODE_UNSPECIFIED`：未指定； - `CANNOT_CREATE_SUPPLY_ON_TPF`：无法在3PF创建交货； - `DROP_OFF_POINT_CLOSED`...
	Message string `json:"message"` // 拒绝原因说明。
}

// 已完成的交货状态： - `ARCHIVE_STATUS_UNSPECIFIED`：未指定； - `COMPLETED`：已完成； - `REJECTED_AT_SUPPLY_WAREHOUSE`：被仓库拒绝； - `CANCELLED_BY...
type V1ArchiveStatusEnum string

// 交货商品汇总信息。
type V1ArchiveSkuSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"` // 商品总体积（升）。
	TotalItemsCount            int64   `json:"total_items_count"`              // 交货中的SKU数量。
	TotalQuantity              int64   `json:"total_quantity"`                 // 交货中的商品数量。
}

type V1FbpArchiveListResponseItem struct {
	CreatedDate       string                 `json:"created_date"` // 交货申请创建日期。
	DeliveryDetails   Fbpv1DeliveryDetails   `json:"delivery_details"`
	HasAct            bool                   `json:"has_act"`             // `true`，前提是已生成交接单。
	PackageUnitsCount int32                  `json:"package_units_count"` // 货位数量。
	WhcOrderID        int64                  `json:"whc_order_id"`        // 合作仓库已完成交货的标识符。
	RowVersion        int64                  `json:"row_version"`         // 草稿的当前版本标识符。
	ActFileUuid       string                 `json:"act_file_uuid"`       // 验收证明书标识符。
	BundleID          string                 `json:"bundle_id"`           // 已验证商品清单的标识符。
	DeclineReason     V1ArchiveDeclineReason `json:"decline_reason"`
	ExternalOrderID   string                 `json:"external_order_id"` // 合作仓库自身系统已完成交货的标识符。
	Status            V1ArchiveStatusEnum    `json:"status"`
	BundleSKUSummary  V1ArchiveSkuSummary    `json:"bundle_sku_summary"`
	HasLabel          bool                   `json:"has_label"`      // `true`，前提是已生成标签。
	OrderDraftID      int64                  `json:"order_draft_id"` // 交货草稿标识符。
	ReceiveDate       string                 `json:"receive_date"`   // 交货接收日期和时间。
	SupplyID          string                 `json:"supply_id"`      // 交货标识符。
	WarehouseID       int64                  `json:"warehouse_id"`   // 仓库标识符。
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	DraftID    int64  `json:"draft_id"`    // 草稿标识符。
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID   string `json:"supply_id"`   // 供货申请标识符。
}

// 网站上的商品价格。
type MoneyMoneyCustomerPrice struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 商品价格。
type MoneyPostingMoney struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 计入Ozon折扣后的卖家价格。
type MoneyMoneySellerPrice struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

type PostingV1PostingFbpListResponsePostingsProducts struct {
	SellerPrice   MoneyMoneySellerPrice   `json:"seller_price"`
	SKU           int64                   `json:"sku"` // Ozon系统中的商品标识符，即SKU。
	CustomerPrice MoneyMoneyCustomerPrice `json:"customer_price"`
	Name          string                  `json:"name"`     // 订单中的商品名称。
	OfferID       string                  `json:"offer_id"` // 卖家系统中的商品标识符，即货号。
	Price         MoneyPostingMoney       `json:"price"`
	Quantity      int32                   `json:"quantity"` // 货件中的商品数量。
}

// 生成状态： - `STATUS_UNSPECIFIED` ——未定义； - `NOT_EXIST` ——不存在； - `PROCESSING` ——处理中； - `EXIST` ——已完成； - `ERROR` ——错误。
type V1FbpCheckActStateResponseStatus string

// 生成错误： - `ERROR_REASON_UNSPECIFIED` ——未定义； - `INVALID_COMPANY` ——公司无效； - `FILE_NOT_FOUND` ——文件未找到； - `GENERATE_TIMEOUT_RE...
type FbpCheckActStateResponseErrorReason string

type V1FbpCheckActStateResponse struct {
	Status V1FbpCheckActStateResponseStatus    `json:"status"`
	CdnURL string                              `json:"cdn_url"` // 验收证明书链接。
	Error  FbpCheckActStateResponseErrorReason `json:"error"`
}

type V1FbpDraftDirectProductValidateRequestSkuItem struct {
	Count int64 `json:"count"` // 交货商品数量。
	SKU   int64 `json:"sku"`   // 商品标识符（SKU）。
}

type V1FbpDraftDirectProductValidateRequest struct {
	Skus        []V1FbpDraftDirectProductValidateRequestSkuItem `json:"skus"`         // 商品标识符（SKU）列表。
	WarehouseID int64                                           `json:"warehouse_id"` // 仓库标识符。
}

// 休息时间。
type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
	TimeslotEnd   string `json:"timeslot_end"`   // 时间段结束时间。
}

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1OrderValidationErrorErrorType string

// 错误信息。
// OrderErrors values
type OrderErrors string

const (
	OrderErrorsErrorTypeUnspecified                                 OrderErrors = "ERROR_TYPE_UNSPECIFIED"                                     // 未定义；
	OrderErrorsDeliveryDriverNameLengthMaximumReached               OrderErrors = "DELIVERY_DRIVER_NAME_LENGTH_MAXIMUM_REACHED"                // 司机姓名长度超限；
	OrderErrorsDeliveryVehicleGenreLengthMaximumReached             OrderErrors = "DELIVERY_VEHICLE_GENRE_LENGTH_MAXIMUM_REACHED"              // 车辆类型长度超限；
	OrderErrorsDeliveryVehicleRegistrationPlateLengthMaximumReached OrderErrors = "DELIVERY_VEHICLE_REGISTRATION_PLATE_LENGTH_MAXIMUM_REACHED" // 车牌号长度超限；
	OrderErrorsDeliveryTPLNameLengthMaximumReached                  OrderErrors = "DELIVERY_TPL_NAME_LENGTH_MAXIMUM_REACHED"                   // 第三方承运人名称长度超限；
	OrderErrorsDeliveryTrackingNumberLengthMaximumReached           OrderErrors = "DELIVERY_TRACKING_NUMBER_LENGTH_MAXIMUM_REACHED"            // 运单号长度超限；
	OrderErrorsDeliveryDriverNameEmpty                              OrderErrors = "DELIVERY_DRIVER_NAME_EMPTY"                                 // 未填写司机姓名；
	OrderErrorsDeliveryVehicleGenreEmpty                            OrderErrors = "DELIVERY_VEHICLE_GENRE_EMPTY"                               // 未填写车辆类型；
	OrderErrorsDeliveryVehicleRegistrationPlateEmpty                OrderErrors = "DELIVERY_VEHICLE_REGISTRATION_PLATE_EMPTY"                  // 未填写车牌号；
	OrderErrorsDeliveryTPLNameEmpty                                 OrderErrors = "DELIVERY_TPL_NAME_EMPTY"                                    // 未填写第三方承运人名称；
	OrderErrorsDeliveryTrackingNumberEmpty                          OrderErrors = "DELIVERY_TRACKING_NUMBER_EMPTY"                             // 未填写运单号；
	OrderErrorsDeliveryBYSellerEmpty                                OrderErrors = "DELIVERY_BY_SELLER_EMPTY"                                   // 未填写卖家自配送信息；
	OrderErrorsDeliveryBYTPLEmpty                                   OrderErrors = "DELIVERY_BY_TPL_EMPTY"                                      // 未填写第三方承运商配送信息；
	OrderErrorsReceiveDateNOTSET                                    OrderErrors = "RECEIVE_DATE_NOT_SET"                                       // 未指定收货日期；
	OrderErrorsSupplyTypeNOTSupported                               OrderErrors = "SUPPLY_TYPE_NOT_SUPPORTED"                                  // 不支持该交货类型；
	OrderErrorsInvalidBusinessFlow                                  OrderErrors = "INVALID_BUSINESS_FLOW"                                      // 业务流程无效；
	OrderErrorsOrderLocked                                          OrderErrors = "ORDER_LOCKED"                                               // 交货已被锁定；
	OrderErrorsInvalidTimeslot                                      OrderErrors = "INVALID_TIMESLOT"                                           // 时间段无效；
	OrderErrorsDropOFFDetailsEmpty                                  OrderErrors = "DROP_OFF_DETAILS_EMPTY"                                     // 未填写送至接收点的发运信息；
	OrderErrorsPickUPAddressISEmpty                                 OrderErrors = "PICK_UP_ADDRESS_IS_EMPTY"                                   // 未 填写 快递员 发运 地址；
	OrderErrorsPickUPSenderNameISEmpty                              OrderErrors = "PICK_UP_SENDER_NAME_IS_EMPTY"                               // 未填写发件人姓名；
	OrderErrorsPickUPSenderPhoneISEmpty                             OrderErrors = "PICK_UP_SENDER_PHONE_IS_EMPTY"                              // 未填写发件人电话；
	OrderErrorsPickUPAddressISTOOLarge                              OrderErrors = "PICK_UP_ADDRESS_IS_TOO_LARGE"                               // 快递员发运地址超出长度限制；
	OrderErrorsPickUPSenderNameISTOOLarge                           OrderErrors = "PICK_UP_SENDER_NAME_IS_TOO_LARGE"                           // 发件人姓名超出长度限制；
	OrderErrorsPickUPSenderPhoneISTOOLarge                          OrderErrors = "PICK_UP_SENDER_PHONE_IS_TOO_LARGE"                          // 发件人电话超出长度限制；
	OrderErrorsPickUPCommentISTOOLarge                              OrderErrors = "PICK_UP_COMMENT_IS_TOO_LARGE"                               // 给快递员的备注超出长度限制；
	OrderErrorsPickUPDetailsEmpty                                   OrderErrors = "PICK_UP_DETAILS_EMPTY"                                      // 未填写快递员发运的信息；
	OrderErrorsDropOFFAddressNOTSET                                 OrderErrors = "DROP_OFF_ADDRESS_NOT_SET"                                   // 未指定接收点地址；
	OrderErrorsInvalidState                                         OrderErrors = "INVALID_STATE"                                              // 状态无效。
)

type V1OrderValidationError struct {
	OrderErrors []V1OrderValidationErrorErrorType `json:"order_errors"` // 错误类型： - `ERROR_TYPE_UNSPECIFIED`——未定义； - `DELIVERY_DRIVER_NAME_LENGTH_MAXIMUM_REACHED`——司机姓名长度超限； - `DELIVERY_VEHICLE_GE...
}

type V1FbpOrderDropOffCancelResponse struct {
	Error      V1OrderValidationError `json:"error"`
	IsError    bool                   `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                  `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError struct {
	SKU    int64                   `json:"sku"`    // Ozon系统中的商品标识符——SKU。
	Errors []V1BundleItemErrorEnum `json:"errors"` // 错误： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——商品不在交货品类中； - `INVALID`——状态不正确； - `INCOMPATIBLE_WAREHO...
}

// 交货注册错误： - `ORDER_ERROR_TYPE_UNSPECIFIED` — 未知订单错误类型； - `INVALID_NUMBER_OF_PACKAGE_UNITS` — 申请中货位数量错误； - `MAXIMUM_NUMBER_...
type V1OrderErrorTypeEnum string

// 错误。
type V1FbpDraftPickUpRegistrateResponseRegistrationError struct {
	BundleErrors []V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError `json:"bundle_errors"` // 商品验证列表错误。
	OrderError   V1OrderErrorTypeEnum                                             `json:"order_error"`
}

type V1OrderAttentionTypeEnum string

// 订单状态： - `ORDER_STATUS_UNSPECIFIED`——未指定； - `READY_TO_SUPPLY`——准备发运； - `FILLING_DELIVERY_DETAILS`——填写交货数据； - `COURIER_ASS...
type V1OrderStatusEnum string

// AttentionReasons values
type AttentionReasons string

const (
	AttentionReasonsOrderAttentionTypeUnspecified AttentionReasons = "ORDER_ATTENTION_TYPE_UNSPECIFIED" // 未指定；
	AttentionReasonsOLD                           AttentionReasons = "OLD"                              // 过期申请；
	AttentionReasonsTimeSlotExpired               AttentionReasons = "TIME_SLOT_EXPIRED"                // 时间段已过期。
)

type V1FbpOrderGetResponse struct {
	CreatedDate        string                     `json:"created_date"`        // 交货创建日期。
	DraftID            int64                      `json:"draft_id"`            // 草稿标识符。
	HasLabel           bool                       `json:"has_label"`           // `true`，如果有标签。
	PackageUnitsCount  int32                      `json:"package_units_count"` // 货位数量。
	BundleUuid         string                     `json:"bundle_uuid"`         // 组成商品标识符。
	CancellationState  V1CancellationState        `json:"cancellation_state"`
	DeliveryDetails    Fbpv1DeliveryDetails       `json:"delivery_details"`
	Locked             bool                       `json:"locked"`               // `true`，如果无法编辑交货。
	RowVersion         int64                      `json:"row_version"`          // 草稿的当前版本标识符。
	AttentionReasons   []V1OrderAttentionTypeEnum `json:"attention_reasons"`    // 警告原因： - `ORDER_ATTENTION_TYPE_UNSPECIFIED`——未指定； - `OLD`——过期申请； - `TIME_SLOT_EXPIRED`——时间段已过期。
	CanBeCancelled     bool                       `json:"can_be_cancelled"`     // `true`，如果申请可以取消。
	HasConsignmentNote bool                       `json:"has_consignment_note"` // `true`，如果有已签署的文件。
	ReceiveDate        string                     `json:"receive_date"`         // 交货接收日期和时间。
	SupplyID           string                     `json:"supply_id"`            // 交货申请标识符。
	ID                 int64                      `json:"id"`                   // 交货申请标识符。
	OrderNumber        string                     `json:"order_number"`         // 交货编号。
	Status             V1OrderStatusEnum          `json:"status"`
	WarehouseID        int64                      `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpDraftDropOffProductValidateRequestSkuItem struct {
	Count int32 `json:"count"` // 数量。
	SKU   int64 `json:"sku"`   // Ozon系统中的商品标识符—— SKU。
}

// 营业时间。
type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening struct {
	TimeslotEnd   string `json:"timeslot_end"`   // 时间段结束时间。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
}

// 营业时间表。
type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem struct {
	OpeningHours V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening `json:"opening_hours"`
	BreakHours   V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime        `json:"break_hours"`
	IsHoliday    bool                                                                               `json:"is_holiday"` // `true`，表示节假日。
}

// 星期： - `DAY_OF_WEEK_UNSPECIFIED`——未指定； - `MONDAY`——星期一； - `TUESDAY`——星期二； - `WEDNESDAY`——星期三； - `THURSDAY`——星期四； - `FRIDA...
type V1DayOfWeek string

type V1FbpDraftDropOffPointTimetableResponseCalendar struct {
	CalendarItem V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem `json:"calendar_item"`
	DayOfWeek    V1DayOfWeek                                                 `json:"day_of_week"`
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []V1FbpDraftDropOffPointTimetableResponseCalendar `json:"calendar"` // 接收点的营业时间表。
}

type V1FbpDraftPickupCreateResponse struct {
	SupplyID   string `json:"supply_id"`   // 交货标识符。
	DraftID    int64  `json:"draft_id"`    // 草稿标识符。
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDirectProductValidateResponseApprovedItem struct {
	SKU      int64   `json:"sku"`       // 商品标识符（SKU）。
	Volume   float64 `json:"volume"`    // 商品体积。
	Barcode  string  `json:"barcode"`   // 条形码。
	IconName string  `json:"icon_name"` // 商品图片链接。
	Name     string  `json:"name"`      // 商品名称。
	OfferID  string  `json:"offer_id"`  // 卖家系统中的商品货号。
	Quantity int32   `json:"quantity"`  // 商品数量。
}

// RejectionReasons values
type RejectionReasons string

const (
	RejectionReasonsBundleItemErrorUnspecified RejectionReasons = "BUNDLE_ITEM_ERROR_UNSPECIFIED" // 未指定；
	RejectionReasonsOUTOFAssortment            RejectionReasons = "OUT_OF_ASSORTMENT"             // 未找到商品；
	RejectionReasonsInvalid                    RejectionReasons = "INVALID"                       // 商品未创建；
	RejectionReasonsIncompatibleWarehouse      RejectionReasons = "INCOMPATIBLE_WAREHOUSE"        // 仓库标识符错误
	RejectionReasonsInvalidBarcode             RejectionReasons = "INVALID_BARCODE"               // 未指定条形码；
	RejectionReasonsMultiplicity               RejectionReasons = "MULTIPLICITY"                  // 商品数量不是所需批量的数倍；
	RejectionReasonsNOPrice                    RejectionReasons = "NO_PRICE"                      // 未指定价格；
	RejectionReasonsBanned                     RejectionReasons = "BANNED"                        // 商品不可在选定仓库销售或交货；
	RejectionReasonsZeroQuantity               RejectionReasons = "ZERO_QUANTITY"                 // 交货商品数量为0；
	RejectionReasonsQuantityGreaterThenMAX     RejectionReasons = "QUANTITY_GREATER_THEN_MAX"     // 单个SKU商品数量超过最大值；
	RejectionReasonsNOSales                    RejectionReasons = "NO_SALES"                      // 商品超过60天无销售；
	RejectionReasonsSurplus                    RejectionReasons = "SURPLUS"                       // 仓库中的商品库存可供销售90天；
	RejectionReasonsAvailabilityISEmpty        RejectionReasons = "AVAILABILITY_IS_EMPTY"         // 无商品可售性信息。
)

type V1FbpDraftDirectProductValidateResponseRejectedItem struct {
	Volume           float64                 `json:"volume"`            // 商品体积。
	Barcode          string                  `json:"barcode"`           // 条形码。
	IconName         string                  `json:"icon_name"`         // 商品图片链接。
	Name             string                  `json:"name"`              // 商品名称。
	OfferID          string                  `json:"offer_id"`          // 卖家系统中的商品货号。
	Quantity         int32                   `json:"quantity"`          // 商品数量。
	RejectionReasons []V1BundleItemErrorEnum `json:"rejection_reasons"` // 拒绝原因： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——未找到商品； - `INVALID`——商品未创建； - `INCOMPATIBLE_WAREHOUS...
	SKU              int64                   `json:"sku"`               // 商品标识符（SKU）。
}

// 地址详情。
type FbpWarehouseListResponseAddressDetailing struct {
	Region  string `json:"region"`  // 地区。
	Street  string `json:"street"`  // 街道。
	Zipcode string `json:"zipcode"` // 邮政编码。
	City    string `json:"city"`    // 城市。
	Country string `json:"country"` // 国家。
	House   string `json:"house"`   // 门牌号。
}

type FbpWarehouseListResponseWarehouse struct {
	AddressDetailing FbpWarehouseListResponseAddressDetailing `json:"address_detailing"`
	ID               int64                                    `json:"id"`            // 仓库标识符。
	IsBonded         bool                                     `json:"is_bonded"`     // `true`，表示该仓库为保税仓。
	Name             string                                   `json:"name"`          // 仓库名称。
	PartnerName      string                                   `json:"partner_name"`  // 合作伙伴名称。
	SupplyTypes      []int32                                  `json:"supply_types"`  // 交货类型。
	TimezoneName     string                                   `json:"timezone_name"` // 仓库所在时区。
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 错误。
type V1FbpDraftDropOffRegistrateResponseRegistrationError struct {
	OrderError   V1OrderErrorTypeEnum                                              `json:"order_error"`
	BundleErrors []V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError `json:"bundle_errors"` // 商品验证列表错误。
}

// 休息时间信息。
type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak struct {
	TimeslotEnd   string `json:"timeslot_end"`   // 开始时间。
	TimeslotStart string `json:"timeslot_start"` // 结束时间。
}

// 日期信息。
type V1FbpOrderDropOffTimetableResponseCalendarCalendarItem struct {
	BreakHours   V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak `json:"break_hours"`
	IsHoliday    bool                                                                        `json:"is_holiday"` // `true`，表示休息日。
	OpeningHours V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime      `json:"opening_hours"`
}

type FbpDraftDropOffProvinceListResponseProvince struct {
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
	Name         string `json:"name"`          // 省份名称。
	PointsCount  int32  `json:"points_count"`  // 地图上接收点数量。
}

type V1FbpOrderDropOffTimetableRequest struct {
	ProvinceUuid   string `json:"province_uuid"`     // 省份唯一标识符。
	WarehouseID    int64  `json:"warehouse_id"`      // 仓库标识符。
	DropOffPointID int64  `json:"drop_off_point_id"` // 揽收点标识符。
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 配送详细信息。
type V1FbpDraftPickupCreateRequestDeliveryDetails struct {
	Comment     string `json:"comment"`      // 备注。
	Date        string `json:"date"`         // 送货日期。
	SenderName  string `json:"sender_name"`  // 发件人姓名。
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
	Address     string `json:"address"`      // 地址。
}

type V1FbpDraftPickupCreateRequest struct {
	PackageUnitsCount int32                                        `json:"package_units_count"` // 包装单位数量。
	WarehouseID       int64                                        `json:"warehouse_id"`        // 仓库标识符。
	BundleID          string                                       `json:"bundle_id"`           // 已校验商品列表的标识符。
	DeliveryDetails   V1FbpDraftPickupCreateRequestDeliveryDetails `json:"delivery_details"`
}

// 发件人详细信息。
type V1FbpOrderPickUpDlvEditRequestPickUpDetails struct {
	SenderName  string `json:"sender_name"`  // 发件人姓名。
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
}

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails V1FbpOrderPickUpDlvEditRequestPickUpDetails `json:"pickup_details"`
	RowVersion    int64                                       `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID      string                                      `json:"supply_id"`   // 交货标识符。
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"` // 验收证明书标识符。
}

type V1OrderDraftValidationErrorErrorType string

// 错误信息。
type V1OrderDraftValidationError struct {
	Errors []V1OrderDraftValidationErrorErrorType `json:"errors"` // 错误类型： - `ERROR_TYPE_UNSPECIFIED`——未定义； - `ORDER_DRAFT_LOCKED`——草稿被锁定； - `DELIVERY_DRIVER_NAME_LENGTH_MAXIMUM_REACHED`——司...
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	Error      V1OrderDraftValidationError `json:"error"`
	IsError    bool                        `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                       `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpWarehouseListResponse struct {
	Warehouses []FbpWarehouseListResponseWarehouse `json:"warehouses"` // 仓库列表。
}

type V1FbpCreateActResponse struct {
	Errors    []FbpCreateActResponseCreateActErrorReason `json:"errors"`     // 错误原因： - `CREATE_ACT_ERROR_REASON_UNSPECIFIED` ——未定义； - `INVALID_ORDER_TYPE` ——无法为指定标识符创建验收证明书。
	FileUuid  string                                     `json:"file_uuid"`  // 验收证明书标识符。
	IsSuccess bool                                       `json:"is_success"` // `true`，前提是请求中没有错误。
}

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 交货商品汇总信息。
type ItemBundleSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"` // 商品总体积（升）。
	TotalItemCount             int64   `json:"total_item_count"`               // 交货中的SKU数量。
	TotalQuantity              int64   `json:"total_quantity"`                 // 交货中的商品数量。
}

// 拒绝原因。
type FbpDraftGetResponseDeclineReason struct {
	FailedSKUIds []string `json:"failed_sku_ids"` // 不正确的SKU标识符。
	Message      string   `json:"message"`        // 拒绝文本。
}

type V1FbpDraftGetResponse struct {
	IsCancelable            bool                             `json:"is_cancelable"` // `true`，如果草稿可以取消。
	IsDeletable             bool                             `json:"is_deletable"`  // `true`，如果草稿可以删除。
	WarehouseID             int64                            `json:"warehouse_id"`  // 仓库标识符。
	BundleID                string                           `json:"bundle_id"`     // 验证后商品的列表标识符。
	DeclineReason           FbpDraftGetResponseDeclineReason `json:"decline_reason"`
	ID                      int64                            `json:"id"`                        // 草稿标识符。
	PackageUnitsCount       int32                            `json:"package_units_count"`       // 货位数量。
	RowVersion              int64                            `json:"row_version"`               // 草稿的当前版本标识符。
	DeletedAt               string                           `json:"deleted_at"`                // 草稿删除日期。
	IsRegistrationAvailable bool                             `json:"is_registration_available"` // `true`，如果可注册。
	Locked                  bool                             `json:"locked"`                    // `true`，如果草稿被封锁。
	Editable                bool                             `json:"editable"`                  // `true`，如果草稿可以修改。
	Status                  V1DraftStatusEnum                `json:"status"`
	SupplyID                string                           `json:"supply_id"` // 交货标识符。
	CancellationState       V1CancellationState              `json:"cancellation_state"`
	CreatedAt               string                           `json:"created_at"` // 草稿创建日期。
	DeliveryDetails         V1fbpDeliveryDetails             `json:"delivery_details"`
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpCheckConsignmentNoteStateRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
	Code     string `json:"code"`      // 货物运单标识符。
}

type V1FbpDraftDirectTplDlvCreateResponse struct {
	DraftID    int64  `json:"draft_id"`    // 草稿标识符。
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID   string `json:"supply_id"`   // 交货标识符。
}

type V1FbpDraftDirectTplDlvEditResponse struct {
	IsError    bool                        `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                       `json:"row_version"` // 草稿的当前版本标识符。
	Error      V1OrderDraftValidationError `json:"error"`
}

// 订单详情。
type V1FbpOrderListResponseItem struct {
	HasLabel           bool                       `json:"has_label"`    // `true`，如果有标签。
	ID                 int64                      `json:"id"`           // 交货申请标识符。
	Locked             bool                       `json:"locked"`       // `true`，如果无法编辑交货。
	OrderNumber        string                     `json:"order_number"` // 交货编号。
	ReceiveDate        string                     `json:"receive_date"` // 交货接收日期和时间。
	WarehouseID        int64                      `json:"warehouse_id"` // 仓库标识符。
	CancellationState  V1CancellationState        `json:"cancellation_state"`
	Status             V1OrderStatusEnum          `json:"status"`
	AttentionReasons   []V1OrderAttentionTypeEnum `json:"attention_reasons"`   // 警告原因： - `ORDER_ATTENTION_TYPE_UNSPECIFIED`——未指定； - `OLD`——过期申请； - `TIME_SLOT_EXPIRED`——时间段已过期。
	PackageUnitsCount  int32                      `json:"package_units_count"` // 货位数量。
	SupplyID           string                     `json:"supply_id"`           // 交货申请标识符。
	BundleSummary      ItemBundleSummary          `json:"bundle_summary"`
	CanBeCancelled     bool                       `json:"can_be_cancelled"` // `true`，如果申请可以取消。
	CreatedDate        string                     `json:"created_date"`     // 交货创建日期。
	DeliveryDetails    Fbpv1DeliveryDetails       `json:"delivery_details"`
	HasConsignmentNote bool                       `json:"has_consignment_note"` // `true`，如果有已签署的文件。
}

type V1FbpOrderListResponse struct {
	HasNext bool                         `json:"has_next"` // `true`，如果响应中未返回所有交货。
	Items   []V1FbpOrderListResponseItem `json:"items"`    // 交货。
	LastID  int64                        `json:"last_id"`  // 页面上最后一次交货的标识符。
}

// 生成状态： - `STATE_TYPE_UNSPECIFIED` ——未定义； - `IN_PROGRESS` ——进行中； - `FINISHED` ——成功完成； - `FAILED` ——错误。
type FbpCheckConsignmentNoteStateResponseStateType string

type V1FbpCheckConsignmentNoteStateResponse struct {
	LabelURL     string                                        `json:"label_url"` // 交货标签链接。
	State        FbpCheckConsignmentNoteStateResponseStateType `json:"state"`
	ErrorMessage string                                        `json:"error_message"` // 错误描述。
}

type V1FbpEditTimeslotRequest struct {
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
	RowVersion    int64  `json:"row_version"`    // 草稿的当前版本标识符。
	SupplyID      string `json:"supply_id"`      // 供货申请标识符。
}

// 用于搜索货件的筛选器。
type PostingV1PostingFbpListRequestFilter struct {
	PostingNumbers []string `json:"posting_numbers"` // 货件编号。
	Since          string   `json:"since"`           // 时间段开始。
	Statuses       []string `json:"statuses"`        // 货件状态。
	To             string   `json:"to"`              // 时间段结束。
	Name           string   `json:"name"`            // 商品名称。
	OfferID        string   `json:"offer_id"`        // 卖家系统中的商品标识符，即货号。
}

// Детали доставки.
type V1FbpDraftPickupDlvEditRequestDeliveryDetails struct {
	Address     string `json:"address"`      // 地址。
	Comment     string `json:"comment"`      // 备注。
	Date        string `json:"date"`         // 送货日期。
	SenderName  string `json:"sender_name"`  // 发件人姓名。
	SenderPhone string `json:"sender_phone"` // 发件人电话号码。
}

type V1FbpDraftPickupDlvEditRequest struct {
	SupplyID      string                                        `json:"supply_id"` // 交货标识符。
	PickupDetails V1FbpDraftPickupDlvEditRequestDeliveryDetails `json:"pickup_details"`
	RowVersion    int64                                         `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDirectSellerDlvEditRequest struct {
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
	VehicleType   string `json:"vehicle_type"`   // 车辆类型。
	DriverName    string `json:"driver_name"`    // 司机姓名。
	RowVersion    int64  `json:"row_version"`    // 草稿的当前版本标识符。
	SupplyID      string `json:"supply_id"`      // 供货申请标识符。
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectGetTimeslotResponseTimeslot struct {
	TimeslotEnd   string `json:"timeslot_end"`   // 时间段结束日期。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始日期。
}

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons               []V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason `json:"reasons"`                 // 缺少时间段的原因： - `EMPTY_TIMESLOTS_REASON_UNSPECIFIED`——未定义； - `LOGISTICS_UNKNOWN`——物流方未知错误； - `NO_ROUTE`——没有路线； - `NO_ROUTE_S...
	Timeslots             []V1FbpDraftDirectGetTimeslotResponseTimeslot             `json:"timeslots"`               // 可用时间段列表。
	WarehouseTimezoneName string                                                    `json:"warehouse_timezone_name"` // 卖家仓库的时区。
}

type V1FbpDraftPickUpRegistrateResponse struct {
	Error      V1FbpDraftPickUpRegistrateResponseRegistrationError `json:"error"`
	IsError    bool                                                `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                                               `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	RowVersion           int64  `json:"row_version"`            // 草稿的当前版本标识符。
	SupplyID             string `json:"supply_id"`              // 交货标识符。
	TrackingNumber       string `json:"tracking_number"`        // 货件跟踪号码。
	TransportCompanyName string `json:"transport_company_name"` // 物流公司名称。
}

// 配送详情。
type V1FbpDraftDirectSellerDlvCreateRequestDirectDetails struct {
	VehicleType   string `json:"vehicle_type"`   // 车辆类型。
	DriverName    string `json:"driver_name"`    // 司机姓名。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
}

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	BundleID          string                                              `json:"bundle_id"` // 已验证商品清单的标识符。
	DeliveryDetails   V1FbpDraftDirectSellerDlvCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32                                               `json:"package_units_count"` // 货位数量。
	WarehouseID       int64                                               `json:"warehouse_id"`        // 卖家仓库标识符。
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProductsActions struct {
	DiscountPercent float64 `json:"discount_percent"` // 折扣百分比。
	DiscountValue   float64 `json:"discount_value"`   // 折扣金额。
	IsFromSeller    bool    `json:"is_from_seller"`   // `true`，表示促销活动由卖家创建。
	Description     string  `json:"description"`      // 促销活动名称。
	ActionID        string  `json:"action_id"`        // 促销活动标识符。
	DateFrom        string  `json:"date_from"`        // 促销活动开始日期。
	DateTo          string  `json:"date_to"`          // 促销活动结束日期。
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProducts struct {
	OldPrice                float64                                                               `json:"old_price"`                 // 折扣前的价格。商品卡上会以划线价显示。
	Price                   float64                                                               `json:"price"`                     // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	ProductID               int64                                                                 `json:"product_id"`                // Ozon系统中的商品标识符，即SKU。
	Quantity                int64                                                                 `json:"quantity"`                  // 商品数量。
	TotalDiscountPercent    float64                                                               `json:"total_discount_percent"`    // 折扣百分比。
	TotalDiscountValue      float64                                                               `json:"total_discount_value"`      // 折扣金额。
	Actions                 []PostingV1PostingFbpListResponsePostingsFinancialDataProductsActions `json:"actions"`                   // 促销活动列表。
	CommissionsCurrencyCode string                                                                `json:"commissions_currency_code"` // 佣金货币代码。
}

// 财务数据。
type PostingV1PostingFbpListResponsePostingsFinancialData struct {
	ClusterFrom    string                                                         `json:"cluster_from"`    // 订单发出地区代码。
	ClusterTo      string                                                         `json:"cluster_to"`      // 订单配送地区代码。
	DeliveryAmount float64                                                        `json:"delivery_amount"` // 配送费用。
	Products       []PostingV1PostingFbpListResponsePostingsFinancialDataProducts `json:"products"`        // 订单中的商品列表。
}

type PostingV1PostingFbpListResponsePostings struct {
	Status        string                                               `json:"status"`       // 货件状态。
	OrderDate     string                                               `json:"order_date"`   // 订单的创建日期。
	OrderID       int64                                                `json:"order_id"`     // 该货件所属订单的标识符。
	OrderNumber   string                                               `json:"order_number"` // 该货件所属订单的编号。
	Products      []PostingV1PostingFbpListResponsePostingsProducts    `json:"products"`     // 货件中商品列表。
	FinancialData PostingV1PostingFbpListResponsePostingsFinancialData `json:"financial_data"`
	InProcessAt   string                                               `json:"in_process_at"`  // 货件开始处理的日期和时间。
	PostingNumber string                                               `json:"posting_number"` // 货件编号。
	ProviderID    int64                                                `json:"provider_id"`    // 配送服务标识符。
}

type PostingV1PostingFbpListResponse struct {
	Cursor   string                                    `json:"cursor"`   // 用于选择下一批数据的指针。
	Postings []PostingV1PostingFbpListResponsePostings `json:"postings"` // 货件列表。
}

type V1FbpDraftDirectCreateResponse struct {
	SupplyID   string `json:"supply_id"`   // 交货标识符。
	DraftID    int64  `json:"draft_id"`    // 草稿标识符。
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
}

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type PostingV1PostingFbpListRequestSortDirEnum string

// SortBy values
type SortBy string

const (
	SortByLastChangeStatusDate SortBy = "last_change_status_date" // 按最后一次状态变更日期排序；
	SortByInProcessAt          SortBy = "in_process_at"           // 按开始处理日期排序。
)

type PostingV1PostingFbpListRequest struct {
	Cursor  string                                    `json:"cursor"` // 用于选择下一批数据的指针。
	Filter  PostingV1PostingFbpListRequestFilter      `json:"filter"`
	Limit   int64                                     `json:"limit"`   // 响应中返回的值数量。
	SortBy  SortBy                                    `json:"sort_by"` // 货件排序参数： - `last_change_status_date`——按最后一次状态变更日期排序； - `in_process_at`——按开始处理日期排序。
	SortDir PostingV1PostingFbpListRequestSortDirEnum `json:"sort_dir"`
}

type V1FbpDraftDropOffDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion        int64               `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []FbpDraftDropOffProvinceListResponseProvince `json:"provinces"` // 省份列表。
}

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

// 超级产品标签： - `ITEM_SFBO_ATTRIBUTE_NONE`——无标签； - `ITEM_SFBO_ATTRIBUTE_SUPER_FBO`——超级产品； - `ITEM_SFBO_ATTRIBUTE_ANTI_FBO`——滞销...
type V1ItemSfboAttribute string

// 包装类型： - `BUNDLE_ITEM_SHIPMENT_TYPE_GENERAL`——常规产品； - `BUNDLE_ITEM_SHIPMENT_TYPE_BOX`——盒装； - `BUNDLE_ITEM_SHIPMENT_TYPE_P...
type V1BundleItemShipmentType string

// Tags values
type Tags string

const (
	TagsEvsdRequired    Tags = "EVSD_REQUIRED"    // 需要 Mercury 认证的商品；
	TagsMarkingRequired Tags = "MARKING_REQUIRED" // 需要 “诚实标志” 强制标志的商品；
	TagsMarkingPossible Tags = "MARKING_POSSIBLE" // 可能需要 “诚实标志” 标记的商品；
	TagsJewelry         Tags = "JEWELRY"          // 含珠宝属性的商品；
	TagsTraceable       Tags = "TRACEABLE"        // 可追溯商品；
	TagsEttnRequired    Tags = "ETTN_REQUIRED"    // 可追溯商品，且需要电子运输货运单；
	TagsUndefined       Tags = "UNDEFINED"        // 未知标签。
)

// PlacementZone values
type PlacementZone string

const (
	PlacementZoneUnspecified    PlacementZone = "UNSPECIFIED"     // 未指定；
	PlacementZoneClosedZone     PlacementZone = "CLOSED_ZONE"     // 封闭区域；
	PlacementZoneDangerousGoods PlacementZone = "DANGEROUS_GOODS" // 2–4 类危险品；
	PlacementZoneProducts       PlacementZone = "PRODUCTS"        // 食品；
	PlacementZoneSort           PlacementZone = "SORT"            // 可分拣商品；
	PlacementZoneNONSort        PlacementZone = "NON_SORT"        // 不可分拣商品；
	PlacementZoneOversize       PlacementZone = "OVERSIZE"        // 超大货物；
	PlacementZoneJewelry        PlacementZone = "JEWELRY"         // 珠宝类商品；
	PlacementZoneUnresolved     PlacementZone = "UNRESOLVED"      // 未知区域。
)

type V1ItemResponse struct {
	IconPath            string                   `json:"icon_path"`              // 商品图片链接。
	Name                string                   `json:"name"`                   // 商品名称。
	OfferID             string                   `json:"offer_id"`               // 商品在卖家系统中的标识符 — 货号。
	Quant               int32                    `json:"quant"`                  // 单个包装中的商品数量。
	TotalVolumeInLitres float64                  `json:"total_volume_in_litres"` // 所有商品总体积（升）。
	ContractorItemCode  string                   `json:"contractor_item_code"`   // 卖家系统中的商品标识符——货号。
	Tags                []string                 `json:"tags"`                   // 交货或交货申请中的商品标签。 可能的取值： - `EVSD_REQUIRED`——需要 Mercury 认证的商品； - `MARKING_REQUIRED`——需要 “诚实标志” 强制标志的商品； - `MARKING_POSSIBLE`...
	PlacementZone       PlacementZone            `json:"placement_zone"`         // 商品存放区域： - `UNSPECIFIED`——未指定； - `CLOSED_ZONE`——封闭区域； - `DANGEROUS_GOODS`——2–4 类危险品； - `PRODUCTS`——食品； - `SORT`——可分拣商品； -...
	SKU                 int64                    `json:"sku"`                    // 商品在Ozon系统中的ID（SKU）。
	ProductID           int64                    `json:"product_id"`             // Ozon系统中商品的标识符 — `product_id`。
	SfboAttribute       V1ItemSfboAttribute      `json:"sfbo_attribute"`
	Quantity            int32                    `json:"quantity"`          // 商品数量。
	Barcode             string                   `json:"barcode"`           // 商品条形码。
	IsQuantEditable     bool                     `json:"is_quant_editable"` // `true`，表示单个包装中的商品数量可修改。
	VolumeInLitres      float64                  `json:"volume_in_litres"`  // 商品体积（升）。
	ShipmentType        V1BundleItemShipmentType `json:"shipment_type"`
}

type V1FbpDraftDropOffProductValidateResponseRejectedItem struct {
	Name             string                  `json:"name"`              // 商品名称。
	OfferID          string                  `json:"offer_id"`          // 卖家系统中的商品标识符——货号。
	Quantity         int32                   `json:"quantity"`          // 商品数量。
	RejectionReasons []V1BundleItemErrorEnum `json:"rejection_reasons"` // 拒收原因： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——商品不在交货范围内； - `INVALID`——状态不正确； - `INCOMPATIBLE_WARE...
	SKU              int64                   `json:"sku"`               // Ozon系统中的商品标识符—— SKU。
	Volume           float64                 `json:"volume"`            // 商品体积。
	Barcode          string                  `json:"barcode"`           // 商品条形码。
	IconName         string                  `json:"icon_name"`         // 商品图片链接。
}

type V1FbpOrderDropOffTimetableResponseCalendar struct {
	CalendarItem V1FbpOrderDropOffTimetableResponseCalendarCalendarItem `json:"calendar_item"`
	DayOfWeek    V1DayOfWeek                                            `json:"day_of_week"`
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []V1FbpOrderDropOffTimetableResponseCalendar `json:"calendar"` // 接收点的营业时间信息。
}

type V1FbpArchiveListRequest struct {
	LastID string `json:"last_id"` // 页面上最后一个值的标识符。首次请求时请留空。 如需获取后续数据，请填写上次响应中的 `last_id`。
	Count  string `json:"count"`   // 响应中的元素数量。
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"` // 货物运单标识符。
}

type V1FbpDraftDropOffCreateResponse struct {
	DraftID    int64  `json:"draft_id"`    // 草稿标识符。
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID   string `json:"supply_id"`   // 交货申请标识符。
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus        []V1FbpDraftPickUpProductValidateRequestSkuItem `json:"skus"`         // 商品标识符（SKU）列表。
	WarehouseID int64                                           `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftListRequest struct {
	Count  int32 `json:"count"`   // 响应中的商品数量。
	LastID int64 `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
}

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffProvinceUuid string `json:"drop_off_province_uuid"` // 省份唯一标识符。
	RowVersion          int64  `json:"row_version"`            // 草稿的当前版本标识符。
	SupplyID            string `json:"supply_id"`              // 交货申请标识符。
	DropOffDate         string `json:"drop_off_date"`          // 送货日期。
	DropOffPointID      int64  `json:"drop_off_point_id"`      // 揽收点标识符。
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	DropOffPointID int64  `json:"drop_off_point_id"` // 揽收点标识符。
	ProvinceUuid   string `json:"province_uuid"`     // 省份唯一标识符。
	WarehouseID    int64  `json:"warehouse_id"`      // 仓库标识符。
}

type V1FbpDraftDropOffProductValidateRequest struct {
	WarehouseID int64                                            `json:"warehouse_id"` // 仓库标识符。
	Skus        []V1FbpDraftDropOffProductValidateRequestSkuItem `json:"skus"`         // Ozon系统中的商品标识符—— SKU。
}

type V1FbpDraftDropOffRegistrateResponse struct {
	Error      V1FbpDraftDropOffRegistrateResponseRegistrationError `json:"error"`
	IsError    bool                                                 `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                                                `json:"row_version"` // 草稿的当前版本标识符。
}

// 标签生成任务状态： - `UNSPECIFIED`：未指定； - `IN_PROGRESS`：生成中； - `FINISHED`：生成成功； - `FAILED`：生成失败。
type FbpGetLabelResponseLabelCreationStateTypeEnum string

type V1FbpDraftDropOffProductValidateResponseApprovedItem struct {
	Barcode  string  `json:"barcode"`   // 商品条形码。
	IconName string  `json:"icon_name"` // 商品图片链接。
	Name     string  `json:"name"`      // 商品名称。
	OfferID  string  `json:"offer_id"`  // 卖家系统中的商品标识符——货号。
	Quantity int32   `json:"quantity"`  // 商品数量。
	SKU      int64   `json:"sku"`       // Ozon系统中的商品标识符—— SKU。
	Volume   float64 `json:"volume"`    // 商品体积。
}

type V1FbpDraftDropOffProductValidateResponse struct {
	BundleID        string                                                 `json:"bundle_id"`        // 验证后的商品列表标识符。
	RejectedItems   []V1FbpDraftDropOffProductValidateResponseRejectedItem `json:"rejected_items"`   // 被拒绝的商品。
	ApprovedItems   []V1FbpDraftDropOffProductValidateResponseApprovedItem `json:"approved_items"`   // 已接收的商品。
	BundleGenerated bool                                                   `json:"bundle_generated"` // `true`，前提是已创建商品成分信息。
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectRegistrateRequest struct {
	SupplyID   string `json:"supply_id"`   // 交货标识符。
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"` // 标签生成任务标识符。
}

type V1FbpDraftPickUpProductValidateResponseApprovedItem struct {
	SKU      int64   `json:"sku"`       // 商品标识符（SKU）。
	Volume   float64 `json:"volume"`    // 商品体积。
	Barcode  string  `json:"barcode"`   // 条形码。
	IconName string  `json:"icon_name"` // 商品图片链接。
	Name     string  `json:"name"`      // 商品名称。
	OfferID  string  `json:"offer_id"`  // 卖家系统中的商品货号。
	Quantity int32   `json:"quantity"`  // 商品数量。
}

type V1FbpDraftPickUpProductValidateResponseRejectedItem struct {
	OfferID          string                  `json:"offer_id"`          // 卖家系统中的商品货号。
	Quantity         int32                   `json:"quantity"`          // 商品数量。
	RejectionReasons []V1BundleItemErrorEnum `json:"rejection_reasons"` // 拒绝原因： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——未找到商品； - `INVALID`——商品未创建； - `INCOMPATIBLE_WAREHOUS...
	SKU              int64                   `json:"sku"`               // 商品标识符（SKU）。
	Volume           float64                 `json:"volume"`            // 商品体积。
	Barcode          string                  `json:"barcode"`           // 条形码。
	IconName         string                  `json:"icon_name"`         // 商品图片链接。
	Name             string                  `json:"name"`              // 商品名称。
}

type V1FbpDraftDirectTimeslotEditResponseReserveFailureType string

// ErrorReasons values
type ErrorReasons string

const (
	ErrorReasonsReserveFailureTypeUnspecified ErrorReasons = "RESERVE_FAILURE_TYPE_UNSPECIFIED" // 未定义；
	ErrorReasonsRequestValidation             ErrorReasons = "REQUEST_VALIDATION"               // 请求中填写了过去的预定日期；
	ErrorReasonsInvalidReserve                ErrorReasons = "INVALID_RESERVE"                  // 原始预留未找到、已失效或已包含申请，但尝试覆盖；
	ErrorReasonsLogisticsReason               ErrorReasons = "LOGISTICS_REASON"                 // 物流方错误；
	ErrorReasonsScheduleReason                ErrorReasons = "SCHEDULE_REASON"                  // 排期方错误；
	ErrorReasonsNOCapacity                    ErrorReasons = "NO_CAPACITY"                      // 无可用预定时段。
)

type V1FbpDraftDirectTimeslotEditResponse struct {
	ErrorReasons []V1FbpDraftDirectTimeslotEditResponseReserveFailureType `json:"error_reasons"` // 错误原因： - `RESERVE_FAILURE_TYPE_UNSPECIFIED`——未定义； - `REQUEST_VALIDATION`——请求中填写了过去的预定日期； - `INVALID_RESERVE`——原始预留未找到、已失效...
	RowVersion   int64                                                    `json:"row_version"`   // 草稿的当前版本标识符。
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error      V1OrderValidationError `json:"error"`
	IsError    bool                   `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                  `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpArchiveGetResponse struct {
	ID                 int64                  `json:"id"`                    // 档案记录编号。
	OrderDraftID       int64                  `json:"order_draft_id"`        // 交货草稿标识符。
	ReceiveDate        string                 `json:"receive_date"`          // 交货接收日期和时间。
	BundleID           string                 `json:"bundle_id"`             // 已验证商品清单的标识符。
	BusinessFlowTypeID int64                  `json:"business_flow_type_id"` // 交货类型标识符。
	CreatedDate        string                 `json:"created_date"`          // 交货申请创建日期和时间。
	HasLabel           bool                   `json:"has_label"`             // `true`，前提是已生成标签。
	DeliveryDetails    Fbpv1DeliveryDetails   `json:"delivery_details"`
	OrderNumber        string                 `json:"order_number"`        // 已完成交货标识符。
	PackageUnitsCount  int32                  `json:"package_units_count"` // 货位数量。
	RowVersion         int64                  `json:"row_version"`         // 草稿的当前版本标识符。
	ActFileUuid        string                 `json:"act_file_uuid"`       // 验收证明书标识符。
	BundleSKUSummary   V1ArchiveSkuSummary    `json:"bundle_sku_summary"`
	DeclineReason      V1ArchiveDeclineReason `json:"decline_reason"`
	HasAct             bool                   `json:"has_act"` // `true`，前提是已生成交接单。
	Status             V1ArchiveStatusEnum    `json:"status"`
	SupplyID           string                 `json:"supply_id"`    // 交货标识符。
	WarehouseID        int64                  `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type FbpDraftDropOffPointListResponseDropOffPoint struct {
	City               string `json:"city"`                  // 城市。
	DropOffPointID     int64  `json:"drop_off_point_id"`     // 揽收点标识符。
	NearestDropOffDate string `json:"nearest_drop_off_date"` // 最近的发运日期。
	PointAddress       string `json:"point_address"`         // 接收点地址。
	ProvinceUuid       string `json:"province_uuid"`         // 省份唯一标识符。
}

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []FbpDraftDropOffPointListResponseDropOffPoint `json:"drop_off_points"` // 接收点列表。
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpArchiveListResponse struct {
	HasNext bool                           `json:"has_next"` // `true`，前提是本次响应未返回所有数据。
	Items   []V1FbpArchiveListResponseItem `json:"items"`    // 已完成交货。
	LastID  int64                          `json:"last_id"`  // 页面上最后一个值的标识符。
}

type V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError struct {
	Errors []V1BundleItemErrorEnum `json:"errors"` // 错误： - `BUNDLE_ITEM_ERROR_UNSPECIFIED`——未指定； - `OUT_OF_ASSORTMENT`——未找到商品； - `INVALID`——商品未创建； - `INCOMPATIBLE_WAREHOUSE`...
	SKU    int64                   `json:"sku"`    // 商品标识符（SKU）。
}

// 错误。
type V1FbpDraftDirectRegistrateResponseRegistrationError struct {
	BundleErrors []V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError `json:"bundle_errors"` // 校验商品列表的错误。
	OrderError   V1OrderErrorTypeEnum                                             `json:"order_error"`
}

type V1FbpDraftDirectProductValidateResponse struct {
	ApprovedItems   []V1FbpDraftDirectProductValidateResponseApprovedItem `json:"approved_items"`   // 已确认商品。
	BundleGenerated bool                                                  `json:"bundle_generated"` // `true`，前提是已创建校验商品列表。
	BundleID        string                                                `json:"bundle_id"`        // 校验商品列表标识符。
	RejectedItems   []V1FbpDraftDirectProductValidateResponseRejectedItem `json:"rejected_items"`   // 被拒绝的商品。
}

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion    int64  `json:"row_version"`    // 草稿的当前版本标识符。
	SupplyID      string `json:"supply_id"`      // 供货申请标识符。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
}

type V1FbpDraftDirectCreateRequest struct {
	PackageUnitsCount int32                                      `json:"package_units_count"` // 包装单位数量。
	WarehouseID       int64                                      `json:"warehouse_id"`        // 仓库标识符。
	BundleID          string                                     `json:"bundle_id"`           // 已校验商品列表的标识符。要获取，请使用方法[/v1/fbp/draft/direct/product/validate](#operation/FbpDraftDirectProductValidate)。
	DeliveryDetails   V1FbpDraftDirectCreateRequestDirectDetails `json:"delivery_details"`
}

// 用于计算产品标签的仓库列表。
type GetSupplyOrderBundleRequestItemTagsCalculation struct {
	DropoffWarehouseID  string   `json:"dropoff_warehouse_id"`  // 用于发货的仓库标识符。
	StorageWarehouseIds []string `json:"storage_warehouse_ids"` // 发货仓库标识符列表，不超过 25 个值。
}

// 排序参数： - `SKU`——SKU； - `NAME`——按商品名称； - `QUANTITY`——按数量； - `TOTAL_VOLUME_IN_LITRES`——按体积（升）。
type V1ItemSortField string

type V1GetSupplyOrderBundleRequest struct {
	BundleIds           []string                                       `json:"bundle_ids"` // 交货商品组成的标识符。可通过方法 [/v3/supply-order/get](#operation/SupplyOrderGet) 获取。
	IsAsc               bool                                           `json:"is_asc"`     // 传入 `true` 表示按升序排序。
	ItemTagsCalculation GetSupplyOrderBundleRequestItemTagsCalculation `json:"item_tags_calculation"`
	LastID              string                                         `json:"last_id"` // 当前页面中最后一个 SKU 值的标识符。
	Limit               int32                                          `json:"limit"`   // 每页商品数量。
	Query               string                                         `json:"query"`   // 搜索查询，例如按商品名称、货号或 SKU 搜索。
	SortField           V1ItemSortField                                `json:"sort_field"`
}

type V1FbpDraftDropOffPointListRequest struct {
	PageSize       int32  `json:"page_size"`        // 每页包含的商品数量。
	ProvinceUuid   string `json:"province_uuid"`    // 省份唯一标识符。
	WarehouseID    int64  `json:"warehouse_id"`     // 仓库标识符。
	NextPageNumber int32  `json:"next_page_number"` // 下一页页码。
}

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpEditTimeslotResponseReserveFailureType string

type V1FbpOrderPickUpDlvEditResponse struct {
	RowVersion int64                  `json:"row_version"` // 草稿的当前版本标识符。
	Error      V1OrderValidationError `json:"error"`
	IsError    bool                   `json:"is_error"` // `true`，前提是有错误。
}

type V1FbpDraftDropOffRegistrateRequest struct {
	SupplyID   string `json:"supply_id"`   // 交货申请标识符。
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderDirectCancelResponse struct {
	RowVersion int64                  `json:"row_version"` // 草稿的当前版本标识符。
	Error      V1OrderValidationError `json:"error"`
	IsError    bool                   `json:"is_error"` // `true`，前提是有错误。
}

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion        int64               `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftPickUpProductValidateResponse struct {
	BundleGenerated bool                                                  `json:"bundle_generated"` // `true`，前提是已创建校验商品列表。
	BundleID        string                                                `json:"bundle_id"`        // 校验商品列表标识符。
	RejectedItems   []V1FbpDraftPickUpProductValidateResponseRejectedItem `json:"rejected_items"`   // 被拒绝的商品。
	ApprovedItems   []V1FbpDraftPickUpProductValidateResponseApprovedItem `json:"approved_items"`   // 已确认商品。
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"` // 交货到揽收点的到达日期。
	RowVersion  int64  `json:"row_version"`   // 草稿的当前版本标识符。
	SupplyID    string `json:"supply_id"`     // 交货标识符。
}

type V1FbpDraftDirectRegistrateResponse struct {
	Error      V1FbpDraftDirectRegistrateResponseRegistrationError `json:"error"`
	IsError    bool                                                `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                                               `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []V1FbpEditTimeslotResponseReserveFailureType `json:"error_reasons"` // 错误原因： - `RESERVE_FAILURE_TYPE_UNSPECIFIED`——未定义； - `REQUEST_VALIDATION`——请求中填写了过去的预定日期； - `INVALID_RESERVE`——原始预留未找到、已失效...
	RowVersion   int64                                         `json:"row_version"`   // 草稿的当前版本标识符。
}

type V1GetSupplyOrderBundleResponse struct {
	TotalCount int32            `json:"total_count"` // 申请中的商品数量。
	HasNext    bool             `json:"has_next"`    // 响应中是否未返回全部商品： - `true`——请使用不同的 `last_id` 值再次请求，以获取其余数据； - `false`——响应已包含全部商品数据。
	LastID     string           `json:"last_id"`     // 当前页面最后一个值的标识符。
	Items      []V1ItemResponse `json:"items"`       // 交货申请中的商品列表。
}

type V1FbpAvailableTimeslotListRequest struct {
	IntervalEnd   string `json:"interval_end"`   // 可用时间段所需区间的结束日期。
	IntervalStart string `json:"interval_start"` // 可用时间段所需区间的开始日期。
	SupplyID      string `json:"supply_id"`      // 交货标识符。
}

type V1FbpDraftPickUpDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion        int64               `json:"row_version"` // 草稿的当前版本标识符。
}

// 配送详情。
type V1FbpDraftDropOffCreateRequestDeliveryDetails struct {
	DropOffPointID      int64  `json:"drop_off_point_id"`      // 揽收点标识符。
	DropOffProvinceUuid string `json:"drop_off_province_uuid"` // 省份唯一标识符。
	DropOffDate         string `json:"drop_off_date"`          // 送货日期。
}

type V1FbpDraftDropOffCreateRequest struct {
	BundleID          string                                        `json:"bundle_id"` // 验证后的商品列表标识符。
	DeliveryDetails   V1FbpDraftDropOffCreateRequestDeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32                                         `json:"package_units_count"` // 货位数量。
	WarehouseID       int64                                         `json:"warehouse_id"`        // 卖家仓库标识符。
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64  `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID   string `json:"supply_id"`   // 交货申请标识符。
}

type V1FbpOrderPickUpCancelResponse struct {
	Error      V1OrderValidationError `json:"error"`
	IsError    bool                   `json:"is_error"`    // `true`，前提是有错误。
	RowVersion int64                  `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpGetLabelRequest struct {
	Code     string `json:"code"`      // 标签生成任务标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"` // 供货申请标识符。
}

type V1FbpDraftDirectGetTimeslotRequest struct {
	BundleID      string `json:"bundle_id"`      // 已验证商品清单的标识符。
	IntervalEnd   string `json:"interval_end"`   // 可用时间段所需区间的结束日期。
	IntervalStart string `json:"interval_start"` // 可用时间段所需区间的开始日期。
	WarehouseID   int64  `json:"warehouse_id"`   // 卖家仓库标识符。
}

type V1FbpGetLabelResponse struct {
	LabelURL string                                        `json:"label_url"` // 交货标签链接。
	State    FbpGetLabelResponseLabelCreationStateTypeEnum `json:"state"`
}

package fbo

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"` // 接收点列表。
}

type PostingV1PostingFbpListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Postings []interface{} `json:"postings"` // 货件列表。
}

type V1FbpOrderListRequest struct {
	Count int32 `json:"count"` // 响应中的交货数量。
	LastID int64 `json:"last_id"` // 页面上最后一次交货的标识符。首次请求时请将此字段留空。 如需获取后续数据，请填写上一次请求响应中最后一次交货的`id`。
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderGetResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	ID int64 `json:"id"` // 交货申请标识符。
	Locked bool `json:"locked"` // `true`，如果无法编辑交货。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
	CanBeCancelled bool `json:"can_be_cancelled"` // `true`，如果申请可以取消。
	CreatedDate string `json:"created_date"` // 交货创建日期。
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	BundleUuid string `json:"bundle_uuid"` // 组成商品标识符。
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	ReceiveDate string `json:"receive_date"` // 交货接收日期和时间。
	DeliveryDetails interface{} `json:"delivery_details"`
	HasConsignmentNote bool `json:"has_consignment_note"` // `true`，如果有已签署的文件。
	HasLabel bool `json:"has_label"` // `true`，如果有标签。
	OrderNumber string `json:"order_number"` // 交货编号。
	AttentionReasons []interface{} `json:"attention_reasons"` // 警告原因： - `ORDER_ATTENTION_TYPE_UNSPECIFIED`——未指定； - `OLD`——过期申请； - `TIME_SLOT_EXPIRED`——时间段已过期。
}

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1GetSupplyOrderBundleRequest struct {
	ItemTagsCalculation interface{} `json:"item_tags_calculation"`
	LastID string `json:"last_id"` // 当前页面中最后一个 SKU 值的标识符。
	Limit int32 `json:"limit"` // 每页商品数量。
	Query string `json:"query"` // 搜索查询，例如按商品名称、货号或 SKU 搜索。
	SortField interface{} `json:"sort_field"`
	BundleIds []interface{} `json:"bundle_ids"` // 交货商品组成的标识符。可通过方法 [/v3/supply-order/get](#operation/SupplyOrderGet) 获取。
	IsAsc bool `json:"is_asc"` // 传入 `true` 表示按升序排序。
}

type V1FbpDraftDropOffRegistrateRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDirectTimeslotEditResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"` // 错误原因： - `RESERVE_FAILURE_TYPE_UNSPECIFIED`——未定义； - `REQUEST_VALIDATION`——请求中填写了过去的预定日期； - `INVALID_RESERVE`——原始预留未找到、已失效...
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftPickUpRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpArchiveListResponse struct {
	HasNext bool `json:"has_next"` // `true`，前提是本次响应未返回所有数据。
	Items []interface{} `json:"items"` // 已完成交货。
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpAvailableTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"` // 可用时间段列表。
	WarehouseTimezoneName string `json:"warehouse_timezone_name"` // 卖家仓库的时区。
	Reasons []interface{} `json:"reasons"` // 缺少时间段的原因： - `EMPTY_TIMESLOTS_REASON_UNSPECIFIED`——未定义； - `LOGISTICS_UNKNOWN`——物流方未知错误； - `NO_ROUTE`——没有路线； - `NO_ROUTE_S...
}

type V1FbpOrderListResponse struct {
	LastID int64 `json:"last_id"` // 页面上最后一次交货的标识符。
	HasNext bool `json:"has_next"` // `true`，如果响应中未返回所有交货。
	Items []interface{} `json:"items"` // 交货。
}

type V1FbpCheckConsignmentNoteStateRequest struct {
	Code string `json:"code"` // 货物运单标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpCreateActResponse struct {
	Errors []interface{} `json:"errors"` // 错误原因： - `CREATE_ACT_ERROR_REASON_UNSPECIFIED` ——未定义； - `INVALID_ORDER_TYPE` ——无法为指定标识符创建验收证明书。
	FileUuid string `json:"file_uuid"` // 验收证明书标识符。
	IsSuccess bool `json:"is_success"` // `true`，前提是请求中没有错误。
}

type V1FbpDraftPickUpDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderPickUpDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
}

type V1FbpCheckConsignmentNoteStateResponse struct {
	ErrorMessage string `json:"error_message"` // 错误描述。
	LabelURL string `json:"label_url"` // 交货标签链接。
	State interface{} `json:"state"`
}

type V1FbpDraftDirectTplDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpEditTimeslotRequest struct {
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDropOffPointListRequest struct {
	PageSize int32 `json:"page_size"` // 每页包含的商品数量。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	NextPageNumber int32 `json:"next_page_number"` // 下一页页码。
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []interface{} `json:"provinces"` // 省份列表。
}

type V1FbpAvailableTimeslotListRequest struct {
	IntervalStart string `json:"interval_start"` // 可用时间段所需区间的开始日期。
	SupplyID string `json:"supply_id"` // 交货标识符。
	IntervalEnd string `json:"interval_end"` // 可用时间段所需区间的结束日期。
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	VehicleType string `json:"vehicle_type"` // 车辆类型。
	DriverName string `json:"driver_name"` // 司机姓名。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []interface{} `json:"calendar"` // 接收点的营业时间表。
}

type V1FbpDraftListRequest struct {
	Count int32 `json:"count"` // 响应中的商品数量。
	LastID int64 `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
}

type V1FbpOrderDropOffCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderPickUpCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftDirectCreateResponse struct {
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectProductValidateRequest struct {
	Skus []interface{} `json:"skus"` // 商品标识符（SKU）列表。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDropOffRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftGetResponse struct {
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	DeliveryDetails interface{} `json:"delivery_details"`
	Locked bool `json:"locked"` // `true`，如果草稿被封锁。
	SupplyID string `json:"supply_id"` // 交货标识符。
	CancellationState interface{} `json:"cancellation_state"`
	DeclineReason interface{} `json:"decline_reason"`
	IsCancelable bool `json:"is_cancelable"` // `true`，如果草稿可以取消。
	IsDeletable bool `json:"is_deletable"` // `true`，如果草稿可以删除。
	IsRegistrationAvailable bool `json:"is_registration_available"` // `true`，如果可注册。
	BundleID string `json:"bundle_id"` // 验证后商品的列表标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	CreatedAt string `json:"created_at"` // 草稿创建日期。
	DeletedAt string `json:"deleted_at"` // 草稿删除日期。
	Editable bool `json:"editable"` // `true`，如果草稿可以修改。
	ID int64 `json:"id"` // 草稿标识符。
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
}

type V1FbpDraftDropOffCreateResponse struct {
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
	TrackingNumber string `json:"tracking_number"` // 货件跟踪号码。
	TransportCompanyName string `json:"transport_company_name"` // 物流公司名称。
}

type V1FbpGetLabelResponse struct {
	LabelURL string `json:"label_url"` // 交货标签链接。
	State interface{} `json:"state"`
}

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons []interface{} `json:"reasons"` // 缺少时间段的原因： - `EMPTY_TIMESLOTS_REASON_UNSPECIFIED`——未定义； - `LOGISTICS_UNKNOWN`——物流方未知错误； - `NO_ROUTE`——没有路线； - `NO_ROUTE_S...
	Timeslots []interface{} `json:"timeslots"` // 可用时间段列表。
	WarehouseTimezoneName string `json:"warehouse_timezone_name"` // 卖家仓库的时区。
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"` // 交货到揽收点的到达日期。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1GetSupplyOrderBundleResponse struct {
	TotalCount int32 `json:"total_count"` // 申请中的商品数量。
	HasNext bool `json:"has_next"` // 响应中是否未返回全部商品： - `true`——请使用不同的 `last_id` 值再次请求，以获取其余数据； - `false`——响应已包含全部商品数据。
	LastID string `json:"last_id"` // 当前页面最后一个值的标识符。
	Items []interface{} `json:"items"` // 交货申请中的商品列表。
}

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpWarehouseListResponse struct {
	Warehouses []interface{} `json:"warehouses"` // 仓库列表。
}

type V1FbpDraftPickupDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	TimeslotStart string `json:"timeslot_start"` // 时间段开始时间。
}

type V1FbpDraftDropOffProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"` // 已接收的商品。
	BundleGenerated bool `json:"bundle_generated"` // `true`，前提是已创建商品成分信息。
	BundleID string `json:"bundle_id"` // 验证后的商品列表标识符。
	RejectedItems []interface{} `json:"rejected_items"` // 被拒绝的商品。
}

type V1FbpDraftPickUpProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"` // 已确认商品。
	BundleGenerated bool `json:"bundle_generated"` // `true`，前提是已创建校验商品列表。
	BundleID string `json:"bundle_id"` // 校验商品列表标识符。
	RejectedItems []interface{} `json:"rejected_items"` // 被拒绝的商品。
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"` // 货物运单标识符。
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftListResponse struct {
	HasNext bool `json:"has_next"` // `true`，如果响应中没有返回所有值。
	Items []interface{} `json:"items"` // 草稿。
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	BundleID string `json:"bundle_id"` // 套装标识符。
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectRegistrateRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
}

type V1FbpDraftDirectTplDlvEditResponse struct {
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	Error interface{} `json:"error"`
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"` // 标签生成任务标识符。
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDropOffDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus []interface{} `json:"skus"` // 商品标识符（SKU）列表。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDirectRegistrateResponse struct {
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
}

type PostingV1PostingFbpListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortBy string `json:"sort_by"` // 货件排序参数： - `last_change_status_date`——按最后一次状态变更日期排序； - `in_process_at`——按开始处理日期排序。
	SortDir interface{} `json:"sort_dir"`
}

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	BundleID string `json:"bundle_id"` // 已验证商品清单的标识符。
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库标识符。
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDropOffCreateRequest struct {
	BundleID string `json:"bundle_id"` // 验证后的商品列表标识符。
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库标识符。
}

type V1FbpOrderDirectCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"` // `true`，前提是有错误。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"` // 错误原因： - `RESERVE_FAILURE_TYPE_UNSPECIFIED`——未定义； - `REQUEST_VALIDATION`——请求中填写了过去的预定日期； - `INVALID_RESERVE`——原始预留未找到、已失效...
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpCheckActStateResponse struct {
	CdnURL string `json:"cdn_url"` // 验收证明书链接。
	Error interface{} `json:"error"`
	Status interface{} `json:"status"`
}

type V1FbpDraftDirectGetTimeslotRequest struct {
	BundleID string `json:"bundle_id"` // 已验证商品清单的标识符。
	IntervalEnd string `json:"interval_end"` // 可用时间段所需区间的结束日期。
	IntervalStart string `json:"interval_start"` // 可用时间段所需区间的开始日期。
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库标识符。
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"` // 交货申请标识符。
}

type V1FbpDraftPickupCreateResponse struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
	DraftID int64 `json:"draft_id"` // 草稿标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpDraftPickupCreateRequest struct {
	BundleID string `json:"bundle_id"` // 已校验商品列表的标识符。
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 包装单位数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"` // 供货申请标识符。
}

type V1FbpOrderDropOffTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	ProvinceUuid string `json:"province_uuid"` // 省份唯一标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDropOffProductValidateRequest struct {
	Skus []interface{} `json:"skus"` // Ozon系统中的商品标识符—— SKU。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpGetLabelRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
	Code string `json:"code"` // 标签生成任务标识符。
}

type V1FbpArchiveGetResponse struct {
	ActFileUuid string `json:"act_file_uuid"` // 验收证明书标识符。
	BundleID string `json:"bundle_id"` // 已验证商品清单的标识符。
	BusinessFlowTypeID int64 `json:"business_flow_type_id"` // 交货类型标识符。
	OrderDraftID int64 `json:"order_draft_id"` // 交货草稿标识符。
	OrderNumber string `json:"order_number"` // 已完成交货标识符。
	Status interface{} `json:"status"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
	CreatedDate string `json:"created_date"` // 交货申请创建日期和时间。
	DeliveryDetails interface{} `json:"delivery_details"`
	ID int64 `json:"id"` // 档案记录编号。
	PackageUnitsCount int32 `json:"package_units_count"` // 货位数量。
	ReceiveDate string `json:"receive_date"` // 交货接收日期和时间。
	SupplyID string `json:"supply_id"` // 交货标识符。
	DeclineReason interface{} `json:"decline_reason"`
	HasAct bool `json:"has_act"` // `true`，前提是已生成交接单。
	HasLabel bool `json:"has_label"` // `true`，前提是已生成标签。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"` // 验收证明书标识符。
}

type V1FbpDraftDirectProductValidateResponse struct {
	RejectedItems []interface{} `json:"rejected_items"` // 被拒绝的商品。
	ApprovedItems []interface{} `json:"approved_items"` // 已确认商品。
	BundleGenerated bool `json:"bundle_generated"` // `true`，前提是已创建校验商品列表。
	BundleID string `json:"bundle_id"` // 校验商品列表标识符。
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectCreateRequest struct {
	BundleID string `json:"bundle_id"` // 已校验商品列表的标识符。要获取，请使用方法[/v1/fbp/draft/direct/product/validate](#operation/FbpDraftDirectProductValidate)。
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"` // 包装单位数量。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffProvinceUuid string `json:"drop_off_province_uuid"` // 省份唯一标识符。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 交货申请标识符。
	DropOffDate string `json:"drop_off_date"` // 送货日期。
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"` // 交货标识符。
}

type V1FbpDraftDirectSellerDlvEditRequest struct {
	VehicleType string `json:"vehicle_type"` // 车辆类型。
	DriverName string `json:"driver_name"` // 司机姓名。
	RowVersion int64 `json:"row_version"` // 草稿的当前版本标识符。
	SupplyID string `json:"supply_id"` // 供货申请标识符。
	VehicleNumber string `json:"vehicle_number"` // 车牌号。
}

type V1FbpArchiveListRequest struct {
	Count string `json:"count"` // 响应中的元素数量。
	LastID string `json:"last_id"` // 页面上最后一个值的标识符。首次请求时请留空。 如需获取后续数据，请填写上次响应中的 `last_id`。
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []interface{} `json:"calendar"` // 接收点的营业时间信息。
}

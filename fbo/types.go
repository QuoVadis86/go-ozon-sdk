package fbo

type V1FbpDraftDirectTplDlvCreateResponse struct {
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpCheckConsignmentNoteStateResponse struct {
	LabelURL string `json:"label_url"`
	State interface{} `json:"state"`
	ErrorMessage string `json:"error_message"`
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type V1FbpDraftDirectTplDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDirectProductValidateRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	Skus []interface{} `json:"skus"`
}

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftGetResponse struct {
	SupplyID string `json:"supply_id"`
	WarehouseID int64 `json:"warehouse_id"`
	IsRegistrationAvailable bool `json:"is_registration_available"`
	RowVersion int64 `json:"row_version"`
	CreatedAt string `json:"created_at"`
	Editable bool `json:"editable"`
	ID int64 `json:"id"`
	IsDeletable bool `json:"is_deletable"`
	Status interface{} `json:"status"`
	DeclineReason interface{} `json:"decline_reason"`
	Locked bool `json:"locked"`
	PackageUnitsCount int32 `json:"package_units_count"`
	BundleID string `json:"bundle_id"`
	CancellationState interface{} `json:"cancellation_state"`
	DeletedAt string `json:"deleted_at"`
	IsCancelable bool `json:"is_cancelable"`
	DeliveryDetails interface{} `json:"delivery_details"`
}

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"`
}

type PostingV1PostingFbpListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
}

type V1FbpDraftPickUpDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDirectCreateRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
}

type V1FbpOrderListResponse struct {
	LastID int64 `json:"last_id"`
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
}

type V1FbpDraftDropOffCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpArchiveListResponse struct {
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
	HasNext bool `json:"has_next"`
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type V1GetSupplyOrderBundleRequest struct {
	BundleIds []interface{} `json:"bundle_ids"`
	IsAsc bool `json:"is_asc"`
	ItemTagsCalculation interface{} `json:"item_tags_calculation"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	Query string `json:"query"`
	SortField interface{} `json:"sort_field"`
}

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDirectRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftPickupCreateRequest struct {
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDirectProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1FbpArchiveListRequest struct {
	Count string `json:"count"`
	LastID string `json:"last_id"`
}

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1FbpOrderListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type V1FbpDraftPickupDlvEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	PickupDetails interface{} `json:"pickup_details"`
}

type V1FbpCheckConsignmentNoteStateRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDropOffCreateResponse struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type V1FbpOrderDropOffCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDirectSellerDlvEditRequest struct {
	SupplyID string `json:"supply_id"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"`
}

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type V1FbpDraftDropOffDeleteResponse struct {
	RowVersion int64 `json:"row_version"`
	CancellationState interface{} `json:"cancellation_state"`
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftPickUpRegistrateResponse struct {
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
	Error interface{} `json:"error"`
}

type V1FbpArchiveGetResponse struct {
	ReceiveDate string `json:"receive_date"`
	Status interface{} `json:"status"`
	BusinessFlowTypeID int64 `json:"business_flow_type_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	SupplyID string `json:"supply_id"`
	WarehouseID int64 `json:"warehouse_id"`
	ActFileUuid string `json:"act_file_uuid"`
	CreatedDate string `json:"created_date"`
	HasAct bool `json:"has_act"`
	PackageUnitsCount int32 `json:"package_units_count"`
	BundleID string `json:"bundle_id"`
	OrderDraftID int64 `json:"order_draft_id"`
	OrderNumber string `json:"order_number"`
	RowVersion int64 `json:"row_version"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
	DeclineReason interface{} `json:"decline_reason"`
	HasLabel bool `json:"has_label"`
	ID int64 `json:"id"`
}

type V1FbpDraftDirectGetTimeslotRequest struct {
	BundleID string `json:"bundle_id"`
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftPickupCreateResponse struct {
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDirectCreateResponse struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
}

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpEditTimeslotRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDirectTimeslotEditResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"`
}

type V1FbpOrderDirectCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpAvailableTimeslotListRequest struct {
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
	SupplyID string `json:"supply_id"`
}

type V1FbpOrderPickUpDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpGetLabelResponse struct {
	LabelURL string `json:"label_url"`
	State interface{} `json:"state"`
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpOrderGetResponse struct {
	ReceiveDate string `json:"receive_date"`
	BundleUuid string `json:"bundle_uuid"`
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	DeliveryDetails interface{} `json:"delivery_details"`
	DraftID int64 `json:"draft_id"`
	HasLabel bool `json:"has_label"`
	Locked bool `json:"locked"`
	SupplyID string `json:"supply_id"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	CreatedDate string `json:"created_date"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	ID int64 `json:"id"`
	OrderNumber string `json:"order_number"`
	PackageUnitsCount int32 `json:"package_units_count"`
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftDirectRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpWarehouseListResponse struct {
	Warehouses []interface{} `json:"warehouses"`
}

type V1FbpGetLabelRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type PostingV1PostingFbpListResponse struct {
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type V1FbpDraftListRequest struct {
	LastID int64 `json:"last_id"`
	Count int32 `json:"count"`
}

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	TransportCompanyName string `json:"transport_company_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TrackingNumber string `json:"tracking_number"`
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []interface{} `json:"provinces"`
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpAvailableTimeslotListResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type V1FbpDraftListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpOrderPickUpCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpCheckActStateResponse struct {
	Status interface{} `json:"status"`
	CdnURL string `json:"cdn_url"`
	Error interface{} `json:"error"`
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"`
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDropOffPointListRequest struct {
	PageSize int32 `json:"page_size"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
	NextPageNumber int32 `json:"next_page_number"`
}

type V1GetSupplyOrderBundleResponse struct {
	Items []interface{} `json:"items"`
	TotalCount int32 `json:"total_count"`
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
}

type V1FbpDraftPickUpProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1FbpOrderDropOffTimetableRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
}

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpCreateActResponse struct {
	Errors []interface{} `json:"errors"`
	FileUuid string `json:"file_uuid"`
	IsSuccess bool `json:"is_success"`
}

package fbs

type V2FbsPostingProductCountrySetResponse struct {
	ProductID int64 `json:"product_id"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
}

type PostingV4PostingFbsUnfulfilledListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Translit bool `json:"translit"`
	With interface{} `json:"with"`
}

type V2FbsPostingResponse struct {
	Result interface{} `json:"result"`
}

type CarriageCarriageGetResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	AvailableActions []interface{} `json:"available_actions"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	IsContainerLabelPrinted bool `json:"is_container_label_printed"`
	IsPartial bool `json:"is_partial"`
	PartialNum int64 `json:"partial_num"`
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	ActType string `json:"act_type"`
	CancelAvailability interface{} `json:"cancel_availability"`
	DepartureDate string `json:"departure_date"`
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`
	IntegrationType string `json:"integration_type"`
	RetryCount int32 `json:"retry_count"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	UpdatedAt string `json:"updated_at"`
	CompanyID int64 `json:"company_id"`
	CreatedAt string `json:"created_at"`
	FirstMileType string `json:"first_mile_type"`
	CarriageID int64 `json:"carriage_id"`
	ContainersCount int32 `json:"containers_count"`
}

type V1SetPostingsRequest struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
	CarriageID int64 `json:"carriage_id"`
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"`
}

type PostingCancelReasonResponse struct {
	Result []interface{} `json:"result"`
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"`
}

type V1AssemblyFbsPostingListResponse struct {
	Cursor string `json:"cursor"`
	Cutoff string `json:"cutoff"`
	Postings []interface{} `json:"postings"`
}

type V1CarriageCancelResponse struct {
	CarriageStatus string `json:"carriage_status"`
	Error string `json:"error"`
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1AssemblyCarriageProductListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type Postingv3GetFbsPostingListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
	Dir string `json:"dir"`
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1AssemblyCarriageProductListResponse struct {
	Cursor string `json:"cursor"`
	Products []interface{} `json:"products"`
}

type Fbsv4FbsPostingShipV4Request struct {
	Packages interface{} `json:"packages"`
	PostingNumber string `json:"posting_number"`
	With interface{} `json:"with"`
}

type PostingBooleanResponse struct {
	Result bool `json:"result"`
}

type V2PostingFBSActGetPostingsRequest struct {
	ID interface{} `json:"id"`
}

type PostingV4PostingFbsListResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type Fbsv4FbsPostingShipV4Response struct {
	AdditionalData interface{} `json:"additional_data"`
	Result interface{} `json:"result"`
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []interface{} `json:"result"`
}

type V1AssemblyCarriagePostingListResponse struct {
	CanPrintMassLabel bool `json:"can_print_mass_label"`
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type PostingCancelReasonListResponse struct {
	Result []interface{} `json:"result"`
}

type V4FbsPostingShipPackageV4Request struct {
	Products []interface{} `json:"products"`
	PostingNumber string `json:"posting_number"`
}

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"`
}

type PostingV4PostingFbsUnfulfilledListResponse struct {
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
	Count int64 `json:"count"`
	Cursor string `json:"cursor"`
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1SetPostingsResponse struct {
	Result interface{} `json:"result"`
}

type Postingv3GetFbsPostingRequest struct {
	PostingNumber string `json:"posting_number"`
	With interface{} `json:"with"`
}

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type Postingv1GetCarriageAvailableListRequest struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type V3GetFbsPostingResponseV3 struct {
	Result interface{} `json:"result"`
}

type V1CarriageCreateRequest struct {
	AllBlrTraceable bool `json:"all_blr_traceable"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type V1CarriageApproveRequest struct {
	ContainersCount int32 `json:"containers_count"`
	CarriageID int64 `json:"carriage_id"`
}

type V3GetFbsPostingListResponseV3 struct {
	Result interface{} `json:"result"`
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []interface{} `json:"tracking_numbers"`
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"`
	PostingNumber string `json:"posting_number"`
}

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"`
}

type V1AssemblyFbsProductListRequest struct {
	SortDir interface{} `json:"sort_dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type PostingPostingProductCancelRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	Items []interface{} `json:"items"`
	PostingNumber string `json:"posting_number"`
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"`
}

type V1AssemblyFbsProductListResponse struct {
	Products []interface{} `json:"products"`
	ProductsCount int32 `json:"products_count"`
	HasNext bool `json:"has_next"`
}

type PostingCancelFbsPostingRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	PostingNumber string `json:"posting_number"`
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"`
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber string `json:"posting_number"`
	ProductID int64 `json:"product_id"`
	CountryIsoCode string `json:"country_iso_code"`
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []interface{} `json:"result"`
}

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	Dir string `json:"dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
}

type Postingv1GetCarriageAvailableListResponse struct {
	Result interface{} `json:"result"`
}

type V1AssemblyFbsPostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type PostingV4PostingFbsListRequest struct {
	With interface{} `json:"with"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Translit bool `json:"translit"`
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result interface{} `json:"result"`
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"`
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []interface{} `json:"related_posting_numbers"`
}

type V1AssemblyCarriagePostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type V2FbsPostingProductCountryListResponse struct {
	Result []interface{} `json:"result"`
}

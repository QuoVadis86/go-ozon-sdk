package types

type PostingV4PostingFbsListResponse struct {
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
	Cursor string `json:"cursor"`
}

type PostingV1PostingFbpListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	Cursor string `json:"cursor"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFees struct {
	Fees []interface{} `json:"fees"`
}

type PostingV4PostingFbsListResponsePostingsCancellation struct {
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseExtremelyLowPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type V1FbpDraftDirectGetTimeslotRequest struct {
	IntervalStart string `json:"interval_start"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	IntervalEnd string `json:"interval_end"`
}

type V1DescriptionCategoryTipsResponse struct {
	Result []interface{} `json:"result"`
}

type V1ArchiveWarehouseFBSRequest struct {
	Reason string `json:"reason"`
	WarehouseID int64 `json:"warehouse_id"`
}

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPoint struct {
	AddressCoordinates interface{} `json:"address_coordinates"`
	Code string `json:"code"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type ProductGetProductInfoPricesV5ResponseItem struct {
	PriceIndexes interface{} `json:"price_indexes"`
	ProductID int64 `json:"product_id"`
	VolumeWeight float64 `json:"volume_weight"`
	Acquiring float64 `json:"acquiring"`
	Commissions interface{} `json:"commissions"`
	MarketingActions interface{} `json:"marketing_actions"`
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
}

type CreateWarehouseFBSRequestWorkingDaysEnum string

type V2GetDiscountTaskListV2Response struct {
	Tasks []interface{} `json:"tasks"`
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type StockShipmentType string

type V1FbpEditTimeslotResponseReserveFailureType string

type GetFBSRatingIndexInfoV1ResponseIndexDynamics struct {
	IndexByDate float64 `json:"index_by_date"`
	ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"`
	Date string `json:"date"`
}

type V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError struct {
	SKU int64 `json:"sku"`
	Errors []interface{} `json:"errors"`
}

type V3ChatList struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type RowItemCommissionReturn struct {
	Bonus float64 `json:"bonus"`
	Commission float64 `json:"commission"`
	Compensation float64 `json:"compensation"`
	PricePerInstance float64 `json:"price_per_instance"`
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	Quantity int32 `json:"quantity"`
	StandardFee float64 `json:"standard_fee"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Stars float64 `json:"stars"`
}

type GetRealizationReportResponseV2Header struct {
	StartDate string `json:"start_date"`
	StopDate string `json:"stop_date"`
	ContractDate string `json:"contract_date"`
	ContractNumber string `json:"contract_number"`
	DocDate string `json:"doc_date"`
	Number string `json:"number"`
	PayerInn string `json:"payer_inn"`
	PayerKpp string `json:"payer_kpp"`
	ReceiverInn string `json:"receiver_inn"`
	ReceiverKpp string `json:"receiver_kpp"`
	CurrencySysName string `json:"currency_sys_name"`
	PayerName string `json:"payer_name"`
	ReceiverName string `json:"receiver_name"`
}

type V1FbpDraftDirectSellerDlvEditRequest struct {
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type ValidationResultValidationError struct {
	TemplateID int32 `json:"template_id"`
	Type interface{} `json:"type_"`
	Characteristic interface{} `json:"characteristic"`
	RestrictionPrice interface{} `json:"restriction_price"`
	RestrictionVwc float64 `json:"restriction_vwc"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequest struct {
	AutoAddDate string `json:"auto_add_date"`
	ToUpdate []interface{} `json:"to_update"`
	ActionID int64 `json:"action_id"`
}

type Productv3GetProductAttributesV3Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"`
}

type Postingv3GetFbsPostingRequest struct {
	PostingNumber string `json:"posting_number"`
	With interface{} `json:"with"`
}

type ProductV1ProductVisibilityInfoResponseItemShowcasesVisibilityEnum string

type V1GetRealizationReportPostingRequest struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type SellerInfoResponseRating struct {
	CurrentValue interface{} `json:"current_value"`
	Name string `json:"name"`
	PastValue interface{} `json:"past_value"`
	Rating string `json:"rating"`
	Status interface{} `json:"status"`
	ValueType interface{} `json:"value_type"`
}

type GetConditionalCancellationListV2ResponseState struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	State interface{} `json:"state"`
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"`
}

type V2ReturnsRfbsListRequest struct {
	Filter interface{} `json:"filter"`
	LastID int32 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type V1SellerActionsProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type Productv5GetProductInfoPricesV5Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type PostingV4PostingFbsListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

type ReviewInfoResponsePhoto struct {
	Height int32 `json:"height"`
	URL string `json:"url"`
	Width int32 `json:"width"`
}

type ResultError struct {
	Code string `json:"code"`
	Status string `json:"status"`
}

type MoneyMoneySalePrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type Postingv3FbsPostingWithParams struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
	Translit bool `json:"translit"`
}

type V3FbsPostingProductExemplarsV3 struct {
	Products []interface{} `json:"products"`
}

type DeliveryDetailsDropOffPointDetails struct {
	ID int64 `json:"id"`
	ProvinceUuid string `json:"province_uuid"`
	Timeslot interface{} `json:"timeslot"`
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1AssemblyCarriageProductListResponse struct {
	Cursor string `json:"cursor"`
	Products []interface{} `json:"products"`
}

type ApproveDeclineDiscountTasksResponseFailDetail struct {
	ErrorForUser string `json:"error_for_user"`
	TaskID int64 `json:"task_id"`
}

type V3FbsPostingProductExemplarInfoV3 struct {
	IsGTDAbsent bool `json:"is_gtd_absent"`
	Rnpt string `json:"rnpt"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Imei []interface{} `json:"imei"`
	ExemplarID int64 `json:"exemplar_id"`
	MandatoryMark string `json:"mandatory_mark"`
	GTD string `json:"gtd"`
}

type V3FbsPostingDetailOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type Productv4GetProductAttributesV4ResponseResult struct {
	Barcode string `json:"barcode"`
	Depth int64 `json:"depth"`
	ModelInfo interface{} `json:"model_info"`
	Attributes []interface{} `json:"attributes"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	ColorImage string `json:"color_image"`
	ID int64 `json:"id"`
	SKU string `json:"sku"`
	WeightUnit string `json:"weight_unit"`
	AttributesWithDefaults []interface{} `json:"attributes_with_defaults"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	DimensionUnit string `json:"dimension_unit"`
	Images interface{} `json:"images"`
	OfferID string `json:"offer_id"`
	PDFList []interface{} `json:"pdf_list"`
	TypeID int64 `json:"type_id"`
	Weight int64 `json:"weight"`
	Barcodes interface{} `json:"barcodes"`
	Height int64 `json:"height"`
	Name string `json:"name"`
	PrimaryImage string `json:"primary_image"`
	Width int64 `json:"width"`
}

type V1SellerActionsUpdateInstallmentRequestActionParameters struct {
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type ProductImportProductsPricesResponseError struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

type PostingV4PostingFbsListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1SellerActionsProductsListRequest struct {
	Limit int64 `json:"limit"`
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
}

type V1SellerActionsUpdateDiscountWithConditionRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
}

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementListEnum string

type V2CancellationStateEnumFilters string

type SearchAttributeValuesResponseValue struct {
	ID int64 `json:"id"`
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
}

type CreateReportResponseCodeNoDeadline struct {
	Code string `json:"code"`
}

type MoneyMoneySellerPrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"`
}

type V1SellerActionsCreateVoucherRequestDiscountTypeEnum string

type V1ProductUpdateOfferIdResponseError struct {
	Message string `json:"message"`
	OfferID string `json:"offer_id"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplar struct {
	ExemplarID int64 `json:"exemplar_id"`
	GTD string `json:"gtd"`
	Marks []interface{} `json:"marks"`
	GTDCheckStatus string `json:"gtd_check_status"`
	GTDErrorCodes []interface{} `json:"gtd_error_codes"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Rnpt string `json:"rnpt"`
	RnptCheckStatus string `json:"rnpt_check_status"`
	RnptErrorCodes []interface{} `json:"rnpt_error_codes"`
}

type V1GetFinanceBalanceV1Request struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1AnalyticsProductQueriesDetailsResponse struct {
	AnalyticsPeriod interface{} `json:"analytics_period"`
	PageCount int64 `json:"page_count"`
	Queries []interface{} `json:"queries"`
	Total int64 `json:"total"`
}

type FilterPeriod struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1SellerActionsProductsCandidatesResponseProductQuantTypeEnum string

type Polygonv1Empty interface{}

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V2PostingFBSActGetProducts struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
}

type RelatedPostingCancelReasons struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	TypeID string `json:"type_id"`
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type SellerActionsListRequestStatusEnum string

type Productv3GetProductListRequestFilterFilterVisibility string

type V1QuestionTopSkuRequest struct {
	Limit int64 `json:"limit"`
}

type V3FbsPostingRequirementsV3 struct {
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
}

type FilterWithQuant struct {
	Created bool `json:"created"`
	Exists bool `json:"exists"`
}

type V1SellerActionsCreateDiscountWithConditionRequestDiscountTypeEnum string

type Fbsv4FbsPostingShipV4Response struct {
	AdditionalData interface{} `json:"additional_data"`
	Result interface{} `json:"result"`
}

type AnalyticsSorting struct {
	Key string `json:"key"`
	Order interface{} `json:"order"`
}

type DetailsReturns struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type V2ProductInfoPicturesRequest struct {
	ProductID interface{} `json:"product_id"`
}

type V4GetUploadQuotaResponse struct {
	DailyCreate interface{} `json:"daily_create"`
	DailyUpdate interface{} `json:"daily_update"`
	OperationLimits interface{} `json:"operation_limits"`
	Total interface{} `json:"total"`
}

type V6FbsPostingProductExemplarSetV6Request struct {
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"`
}

type SellerSellerAPIArrivalPassCreateRequestArrivalPass struct {
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
}

type ProductGetImportProductsInfoRequest struct {
	TaskID int64 `json:"task_id"`
}

type V1QuestionListResponseQuestions struct {
	QuestionLink string `json:"question_link"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
	AnswersCount int64 `json:"answers_count"`
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
	PublishedAt interface{} `json:"published_at"`
	Status interface{} `json:"status"`
	ProductURL string `json:"product_url"`
}

type SellerApiProductV1ResponseResult struct {
	Rejected []interface{} `json:"rejected"`
	ProductIds []interface{} `json:"product_ids"`
}

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type V2WarehouseListV2Response struct {
	HasNext bool `json:"has_next"`
	Cursor string `json:"cursor"`
	Warehouses []interface{} `json:"warehouses"`
}

type V2ReturnsRfbsGetRequest struct {
	ReturnID int64 `json:"return_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements struct {
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
	ProductsRequiringWeight []interface{} `json:"products_requiring_weight"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
}

type V1CommentSort string

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason string

type Productsv1GetProductInfoStocksByWarehouseFbsRequest struct {
	SKU interface{} `json:"sku"`
	OfferID interface{} `json:"offer_id"`
}

type V1GetStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type FbsPostingProductExemplarSetV6RequestExemplars struct {
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	ExemplarID int64 `json:"exemplar_id"`
}

type V1ApproveDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type ProductV1ProductVisibilitySetResponseItemsSelectPermissionEnum string

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"`
}

type DeliveryMethodListV2RequestFilter struct {
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	ProviderIds []interface{} `json:"provider_ids"`
	Status []interface{} `json:"status"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1FbpDraftPickUpProductValidateResponseRejectedItem struct {
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
}

type V3FbsPostingDetail struct {
	Substatus string `json:"substatus"`
	Customer interface{} `json:"customer"`
	Status string `json:"status"`
	Cancellation interface{} `json:"cancellation"`
	Courier interface{} `json:"courier"`
	Products []interface{} `json:"products"`
	ProviderStatus string `json:"provider_status"`
	OrderID int64 `json:"order_id"`
	PostingNumber string `json:"posting_number"`
	ShipmentDate string `json:"shipment_date"`
	DeliveringDate string `json:"delivering_date"`
	DeliveryMethod interface{} `json:"delivery_method"`
	DeliveryPrice string `json:"delivery_price"`
	FinancialData interface{} `json:"financial_data"`
	TrackingNumber string `json:"tracking_number"`
	Tariffication interface{} `json:"tariffication"`
	IsExpress bool `json:"is_express"`
	LegalInfo interface{} `json:"legal_info"`
	ParentPostingNumber string `json:"parent_posting_number"`
	PreviousSubstatus string `json:"previous_substatus"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	Addressee interface{} `json:"addressee"`
	Barcodes interface{} `json:"barcodes"`
	OrderNumber string `json:"order_number"`
	ProductExemplars interface{} `json:"product_exemplars"`
	AnalyticsData interface{} `json:"analytics_data"`
	AvailableActions interface{} `json:"available_actions"`
	FactDeliveryDate string `json:"fact_delivery_date"`
	InProcessAt string `json:"in_process_at"`
	Optional interface{} `json:"optional"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	AdditionalData []interface{} `json:"additional_data"`
	RelatedPostings interface{} `json:"related_postings"`
	Requirements interface{} `json:"requirements"`
}

type V1FbpDraftPickUpProductValidateResponseApprovedItem struct {
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
}

type ReportListRequestReportType string

type ProductProductInfoPicturesResponsePicture struct {
	URL string `json:"url"`
	Is360 bool `json:"is_360"`
	IsColor bool `json:"is_color"`
	IsPrimary bool `json:"is_primary"`
	ProductID int64 `json:"product_id"`
	State string `json:"state"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProducts struct {
	Commission interface{} `json:"commission"`
	CustomerPrice interface{} `json:"customer_price"`
	OldPrice float64 `json:"old_price"`
	Payout float64 `json:"payout"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	Actions []interface{} `json:"actions"`
	Price float64 `json:"price"`
	TotalDiscountValue float64 `json:"total_discount_value"`
}

type V1GetStrategyItemInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsVoucherGetResponse struct {
	File string `json:"file"`
}

type FbpCheckActStateResponseErrorReason string

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type V3ImportProductsRequestComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type V1QuestionChangeStatusRequest struct {
	QuestionIds interface{} `json:"question_ids"`
	Status string `json:"status"`
}

type ArrivalpassArrivalPassCreateRequestArrivalPass struct {
	VehicleModel string `json:"vehicle_model"`
	WarehouseID int64 `json:"warehouse_id"`
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	DropoffPointID int64 `json:"dropoff_point_id"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
}

type V1GetProductInfoSubscriptionResponse struct {
	Result []interface{} `json:"result"`
}

type AnalyticsFilter struct {
	Key string `json:"key"`
	Op interface{} `json:"op"`
	Value string `json:"value"`
}

type V1WarehouseWithInvalidProductsResponse struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1FbpDraftDropOffProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type Financev3Period struct {
	To string `json:"to"`
	From string `json:"from"`
}

type PickedSegmentSegment struct {
	Type interface{} `json:"type_"`
	Description string `json:"description"`
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type V1GetStrategyResponse struct {
	Result interface{} `json:"result"`
}

type CarriageCarriageGetResponse struct {
	TPLProviderID int64 `json:"tpl_provider_id"`
	WarehouseID int64 `json:"warehouse_id"`
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CancelAvailability interface{} `json:"cancel_availability"`
	CarriageID int64 `json:"carriage_id"`
	CompanyID int64 `json:"company_id"`
	IsContainerLabelPrinted bool `json:"is_container_label_printed"`
	PartialNum int64 `json:"partial_num"`
	RetryCount int32 `json:"retry_count"`
	ActType string `json:"act_type"`
	ContainersCount int32 `json:"containers_count"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	FirstMileType string `json:"first_mile_type"`
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`
	CreatedAt string `json:"created_at"`
	DepartureDate string `json:"departure_date"`
	IntegrationType string `json:"integration_type"`
	IsPartial bool `json:"is_partial"`
	UpdatedAt string `json:"updated_at"`
	AvailableActions []interface{} `json:"available_actions"`
	Status string `json:"status"`
}

type StatusEnum string

type V1QuestionTopSkuResponse struct {
	SKU interface{} `json:"sku"`
}

type V1FbpDraftDropOffProductValidateResponseApprovedItem struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
}

type PostingV1PostingFbpListResponsePostingsFinancialData struct {
	Products []interface{} `json:"products"`
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	DeliveryAmount float64 `json:"delivery_amount"`
}

type V1WarehouseFbsCreatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDeliveryService struct {
	Accrued interface{} `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type V1fbpDeliveryDetails struct {
	PickupDetails interface{} `json:"pickup_details"`
	SupplyType interface{} `json:"supply_type"`
	DirectDetails interface{} `json:"direct_details"`
	DropOffPoint interface{} `json:"drop_off_point"`
}

type PostingV1PostingFbpListResponsePostings struct {
	OrderDate string `json:"order_date"`
	OrderNumber string `json:"order_number"`
	Products []interface{} `json:"products"`
	FinancialData interface{} `json:"financial_data"`
	OrderID int64 `json:"order_id"`
	PostingNumber string `json:"posting_number"`
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
	InProcessAt string `json:"in_process_at"`
}

type V2ReturnsRfbsGetResponse struct {
	Returns interface{} `json:"returns"`
}

type V1GetAttributesResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type V1SearchQueriesTopRequest struct {
	Limit string `json:"limit"`
	Offset string `json:"offset"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProduct struct {
	Delivery interface{} `json:"delivery"`
	SKU int64 `json:"sku"`
	Commission interface{} `json:"commission"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1GetAttributeValuesResponseDictionaryValue struct {
	ID int64 `json:"id"`
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
}

type GetReturnsListResponsePlaceNow struct {
	Name string `json:"name"`
	Address string `json:"address"`
	ID int64 `json:"id"`
}

type ReturnsCompanyFbsInfoResponsePassInfo struct {
	Count int32 `json:"count"`
	IsRequired bool `json:"is_required"`
}

type V3FinanceCashFlowStatementListResponse struct {
	Result interface{} `json:"result"`
}

type ProductV4GetUploadQuotaResponseTotalQuotaByCategory struct {
	Usage int64 `json:"usage"`
	CategoryID int64 `json:"category_id"`
	Limit int64 `json:"limit"`
}

type Productv4GetProductAttributesV4Request struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type DetailsServices struct {
	Items []interface{} `json:"items"`
	Total float64 `json:"total"`
}

type Productv3GetProductListRequestFilter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type PolygonBindRequestwhLocation struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type V3AdditionalDataItem struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpOrderDropOffTimetableResponseCalendar struct {
	CalendarItem interface{} `json:"calendar_item"`
	DayOfWeek interface{} `json:"day_of_week"`
}

type V3FinanceCashFlowStatementListResponseResult struct {
	CashFlows interface{} `json:"cash_flows"`
	Details []interface{} `json:"details"`
	PageCount int64 `json:"page_count"`
}

type MoneyMoneyCommission struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponseCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DetailsRfbsDetails struct {
	PartialCompensationReturn float64 `json:"partial_compensation_return"`
	Total float64 `json:"total"`
	TransferDelivery float64 `json:"transfer_delivery"`
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"`
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`
	PartialCompensation float64 `json:"partial_compensation"`
}

type AvailabilityReason struct {
	ID int64 `json:"id"`
	HumanText interface{} `json:"human_text"`
}

type V1AssemblyFbsProductListRequestSortDirEnum string

type V2ServiceType string

type V3GetFbsPostingResponseV3 struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftDirectCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type PostingV4PostingFbsListRequestWith struct {
	LegalInfo bool `json:"legal_info"`
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrual struct {
	Date string `json:"date"`
	ItemFees interface{} `json:"item_fees"`
	NonItemFee interface{} `json:"non_item_fee"`
	Posting interface{} `json:"posting"`
	TotalAmount interface{} `json:"total_amount"`
	AccrualID int64 `json:"accrual_id"`
	UnitNumber string `json:"unit_number"`
	AccruedCategory interface{} `json:"accrued_category"`
}

type ProductV1ProductVisibilitySetRequestItemPlacement struct {
	Placement interface{} `json:"placement"`
	SKU int64 `json:"sku"`
}

type Productv4Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
	Visibility interface{} `json:"visibility"`
}

type V1AnalyticsProductQueriesResponse struct {
	Total int64 `json:"total"`
	AnalyticsPeriod interface{} `json:"analytics_period"`
	Items []interface{} `json:"items"`
	PageCount int64 `json:"page_count"`
}

type GetUploadQuotaResponseDailyUpdate struct {
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
	Usage int64 `json:"usage"`
}

type MoneyMoneyAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1DeleteStrategyItemsResponseResult struct {
	FailedProductCount int32 `json:"failed_product_count"`
}

type ChatChatSendMessageResponse struct {
	Result string `json:"result"`
}

type V1QuestionListRequestFilter struct {
	Status string `json:"status"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type RowItem struct {
	SKU int64 `json:"sku"`
	Barcode string `json:"barcode"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type WarehouseListResponseWarehouse struct {
	WarehouseID int64 `json:"warehouse_id"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	IsAbleToSetPrice bool `json:"is_able_to_set_price"`
	IsPresorted bool `json:"is_presorted"`
	PostingsLimit int32 `json:"postings_limit"`
	FirstMileType interface{} `json:"first_mile_type"`
	IsKgt bool `json:"is_kgt"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
	Status string `json:"status"`
	IsRFBS bool `json:"is_rfbs"`
	IsKarantin bool `json:"is_karantin"`
	MinWorkingDays int64 `json:"min_working_days"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	Name string `json:"name"`
	CanPrintActInAdvance bool `json:"can_print_act_in_advance"`
	IsTimetableEditable bool `json:"is_timetable_editable"`
	WorkingDays interface{} `json:"working_days"`
}

type V1ReviewListRequest struct {
	Status string `json:"status"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortDir string `json:"sort_dir"`
}

type ChatMessageContext struct {
	OrderNumber string `json:"order_number"`
	SKU string `json:"sku"`
}

type V1SellerActionsCreateInstallmentRequest struct {
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendar struct {
	CalendarItem interface{} `json:"calendar_item"`
	DayOfWeek interface{} `json:"day_of_week"`
}

type V1FbpOrderListResponseItem struct {
	SupplyID string `json:"supply_id"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	BundleSummary interface{} `json:"bundle_summary"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	CancellationState interface{} `json:"cancellation_state"`
	Status interface{} `json:"status"`
	CreatedDate string `json:"created_date"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	OrderNumber string `json:"order_number"`
	Locked bool `json:"locked"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	HasLabel bool `json:"has_label"`
	ID int64 `json:"id"`
	ReceiveDate string `json:"receive_date"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseFailedPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type V1DeliveryDetailsPickUpDetails struct {
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
	Address string `json:"address"`
}

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPoint struct {
	LastTransitTimeLocal interface{} `json:"last_transit_time_local"`
	Type interface{} `json:"type_"`
	Address string `json:"address"`
	Coordinates interface{} `json:"coordinates"`
	DiscountPercent float64 `json:"discount_percent"`
	ID string `json:"id"`
}

type Financev3FinanceTransactionListV3ResponseResult struct {
	Operations []interface{} `json:"operations"`
	PageCount int64 `json:"page_count"`
	RowCount int64 `json:"row_count"`
}

type FbsPostingDetailCourier struct {
	CarModel string `json:"car_model"`
	CarNumber string `json:"car_number"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type ReportCreateCompanyPostingsReportRequestFilter struct {
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	IsExpress interface{} `json:"is_express"`
	CancelReasonID []interface{} `json:"cancel_reason_id"`
	ProcessedAtFrom string `json:"processed_at_from"`
	ProcessedAtTo string `json:"processed_at_to"`
	SKU []interface{} `json:"sku"`
	Title string `json:"title"`
	DeliverySchema []interface{} `json:"delivery_schema"`
	OfferID string `json:"offer_id"`
	StatusAlias []interface{} `json:"status_alias"`
	Statuses []interface{} `json:"statuses"`
	WarehouseID []interface{} `json:"warehouse_id"`
}

type V1ArchiveDeclineReason struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V3ChatDetailsInfo struct {
	CreatedAt string `json:"created_at"`
	ChatID string `json:"chat_id"`
	ChatStatus interface{} `json:"chat_status"`
	ChatType interface{} `json:"chat_type"`
}

type AnalyticsProductQueriesResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type ChatInfoChatTypeEnum string

type Fbsv4FbsPostingShipV4Request struct {
	Packages interface{} `json:"packages"`
	PostingNumber string `json:"posting_number"`
	With interface{} `json:"with"`
}

type V1SearchAttributeValuesRequest struct {
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	Value string `json:"value"`
}

type ActionsV1ActionsAutoAddProductsListRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type PostingBooleanResponse struct {
	Result bool `json:"result"`
}

type V1PostingFbsSplitResponse struct {
	ParentPosting interface{} `json:"parent_posting"`
	Postings []interface{} `json:"postings"`
}

type V1Barcode struct {
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

type V2GetDiscountTaskListV2Request struct {
	LastID int64 `json:"last_id"`
	Limit int64 `json:"limit"`
	Status interface{} `json:"status"`
}

type V1GetProductInfoDiscountedResponse struct {
	Items interface{} `json:"items"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountTypeEnum string

type V1ReturnsRfbsActionSetRequest struct {
	Comment string `json:"comment"`
	CompensationAmount float64 `json:"compensation_amount"`
	ID int32 `json:"id"`
	RejectionReasonID int32 `json:"rejection_reason_id"`
	ReturnForBackWay float64 `json:"return_for_back_way"`
	ReturnID int64 `json:"return_id"`
}

type PostingPostingFBSActGetContainerLabelsResponse struct {
	FileContent string `json:"file_content"`
	FileName string `json:"file_name"`
	ContentType string `json:"content_type"`
}

type V1UnarchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostings struct {
	ParentPostingNumber string `json:"parent_posting_number"`
	Tariffication interface{} `json:"tariffication"`
	Barcodes interface{} `json:"barcodes"`
	IsExpress bool `json:"is_express"`
	IsPresortable bool `json:"is_presortable"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	VolumeWeight float64 `json:"volume_weight"`
	ExternalOrder interface{} `json:"external_order"`
	DeliveryMethod interface{} `json:"delivery_method"`
	Requirements interface{} `json:"requirements"`
	AvailableActions []interface{} `json:"available_actions"`
	OrderNumber string `json:"order_number"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	Substatus string `json:"substatus"`
	TrackingNumber string `json:"tracking_number"`
	InProcessAt string `json:"in_process_at"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	PrrOption string `json:"prr_option"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	Cancellation interface{} `json:"cancellation"`
	Customer interface{} `json:"customer"`
	DeliveringDate string `json:"delivering_date"`
	DeliverySchema string `json:"delivery_schema"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	LegalInfo interface{} `json:"legal_info"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	Optional interface{} `json:"optional"`
	AnalyticsData interface{} `json:"analytics_data"`
	OrderID int64 `json:"order_id"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	PostingNumber string `json:"posting_number"`
	QuantumID int64 `json:"quantum_id"`
	ShipmentDate string `json:"shipment_date"`
	FinancialData interface{} `json:"financial_data"`
	Addressee interface{} `json:"addressee"`
	DestinationPlaceName string `json:"destination_place_name"`
	IsMultibox bool `json:"is_multibox"`
}

type PostingV4PostingFbsListRequestSortDirEnum string

type V1GetRealizationReportByDayRequest struct {
	Day int32 `json:"day"`
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type ProductGetProductInfoDescriptionResponseProduct struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Description string `json:"description"`
	ID int64 `json:"id"`
}

type V1GetCompetitorsResponse struct {
	Competitor []interface{} `json:"competitor"`
	Total int32 `json:"total"`
}

type ProductV4GetUploadQuotaResponseOperationLimits struct {
	LimitType interface{} `json:"limit_type"`
	Limit int64 `json:"limit"`
}

type FbsPostingBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type Productv3GetProductAttributesV3Request struct {
	SortDir string `json:"sort_dir"`
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
}

type V4GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"`
	Name string `json:"name"`
}

type FbsPostingShipV4ResponseShipAdditionalData struct {
	PostingNumber string `json:"posting_number"`
	Products interface{} `json:"products"`
}

type V1AnalyticsProductQueriesDetailsRequestSortBy string

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem struct {
	BreakHours interface{} `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"`
	OpeningHours interface{} `json:"opening_hours"`
}

type SetPostingsResponseResult struct {
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
	Result bool `json:"result"`
}

type PostingV1PostingFbpListRequestFilter struct {
	Since string `json:"since"`
	Statuses []interface{} `json:"statuses"`
	To string `json:"to"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1BundleItemErrorEnum string

type V1AssemblyCarriageProductListRequestFilter struct {
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
}

type SellerApiActivateProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type DetailsOthers struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type SellerReturnsv1MoneyProduct struct {
	Price float64 `json:"price"`
	CurrencyCode string `json:"currency_code"`
}

type ProductUpdateOfferIdRequestUpdateOfferId struct {
	NewOfferID string `json:"new_offer_id"`
	OfferID string `json:"offer_id"`
}

type V2ReturnsRfbsRejectRequest struct {
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
	RejectionReasonID int64 `json:"rejection_reason_id"`
}

type V1FbpDraftListResponseItem struct {
	DeliveryDetails interface{} `json:"delivery_details"`
	Editable bool `json:"editable"`
	ID int64 `json:"id"`
	IsCancelable bool `json:"is_cancelable"`
	IsDeletable bool `json:"is_deletable"`
	Locked bool `json:"locked"`
	PackageUnitsCount int32 `json:"package_units_count"`
	Status interface{} `json:"status"`
	SupplyID string `json:"supply_id"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	CancellationState interface{} `json:"cancellation_state"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}

type WarehouseAddressInfo struct {
	Address string `json:"address"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Utc string `json:"utc"`
}

type V1ProductInfoWrongVolumeRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProductsActions struct {
	DiscountPercent float64 `json:"discount_percent"`
	DiscountValue float64 `json:"discount_value"`
	IsFromSeller bool `json:"is_from_seller"`
	Description string `json:"description"`
	ActionID string `json:"action_id"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1FbpDraftPickUpProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1GetProductStairwayDiscountByQuantityRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDelivery struct {
	Services []interface{} `json:"services"`
	TotalAccrued interface{} `json:"total_accrued"`
}

type V1OrderAttentionTypeEnum string

type V5FbsPostingProductExemplarStatusV5Response struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductImportProductsBySKUResponse struct {
	Result interface{} `json:"result"`
}

type RowItemCommission struct {
	Amount float64 `json:"amount"`
	Commission float64 `json:"commission"`
	Compensation float64 `json:"compensation"`
	PricePerInstance float64 `json:"price_per_instance"`
	Quantity int32 `json:"quantity"`
	StandardFee float64 `json:"standard_fee"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Total float64 `json:"total"`
	Bonus float64 `json:"bonus"`
	Stars float64 `json:"stars"`
}

type SellerReturnsv1MoneyWithoutCommission struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type WarehouseTimetable struct {
	TimetableFrom string `json:"timetable_from"`
	TimetableTo string `json:"timetable_to"`
	WorkingHours []interface{} `json:"working_hours"`
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V1QuestionListResponse struct {
	Questions interface{} `json:"questions"`
	LastID string `json:"last_id"`
	HasNext bool `json:"has_next"`
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1QuestionListRequest struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type V1CarriageCreateRequest struct {
	AllBlrTraceable bool `json:"all_blr_traceable"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type GetProductInfoStocksResponseStock struct {
	Reserved int32 `json:"reserved"`
	ShipmentType interface{} `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Type string `json:"type_"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	Present int32 `json:"present"`
}

type V1FbpDraftDropOffPointListRequest struct {
	NextPageNumber int32 `json:"next_page_number"`
	PageSize int32 `json:"page_size"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GetProductInfoListResponseError struct {
	Level interface{} `json:"level"`
	State string `json:"state"`
	Texts interface{} `json:"texts"`
	AttributeID int64 `json:"attribute_id"`
	Code string `json:"code"`
	Field string `json:"field"`
}

type Postingv3GetFbsPostingUnfulfilledListResponseResult struct {
	Count int64 `json:"count"`
	Postings []interface{} `json:"postings"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequest struct {
	Coordinates interface{} `json:"coordinates"`
	CountryCode string `json:"country_code"`
	IsKgt bool `json:"is_kgt"`
	Search interface{} `json:"search"`
}

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type GetConditionalCancellationListV2ResponseResult struct {
	OrderDate string `json:"order_date"`
	PostingNumber string `json:"posting_number"`
	SourceID int64 `json:"source_id"`
	State interface{} `json:"state"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	ApproveComment string `json:"approve_comment"`
	ApproveDate string `json:"approve_date"`
	AutoApproveDate string `json:"auto_approve_date"`
	CancellationID int64 `json:"cancellation_id"`
	CancellationInitiator interface{} `json:"cancellation_initiator"`
	CancellationReason interface{} `json:"cancellation_reason"`
	CancellationReasonMessage string `json:"cancellation_reason_message"`
	CancelledAt string `json:"cancelled_at"`
}

type PostingProductCancelRequestItem struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date interface{} `json:"date"`
}

type V2WarehouseListV2Request struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
}

type V1FbpDraftPickUpDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type FbpDraftGetResponseDeclineReason struct {
	FailedSKUIds []interface{} `json:"failed_sku_ids"`
	Message string `json:"message"`
}

type V2ChatReadResponse struct {
	UnreadCount int64 `json:"unread_count"`
}

type ProductGetProductInfoDescriptionResponse struct {
	Result interface{} `json:"result"`
}

type V1CreateWarehouseFBSRequest struct {
	TimeslotID int64 `json:"timeslot_id"`
	WorkingDays []interface{} `json:"working_days"`
	CutInTime int64 `json:"cut_in_time"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	IsKgt bool `json:"is_kgt"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	AddressCoordinates interface{} `json:"address_coordinates"`
	FirstMileType interface{} `json:"first_mile_type"`
	Options interface{} `json:"options"`
}

type ValidationResultItemStateEnum string

type V1GetAttributeValuesResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type ReviewV2ReviewListV2ResponseReviewOrderStatusEnum string

type V6FbsPostingProductExemplarCreateOrGetV6Request struct {
	PostingNumber string `json:"posting_number"`
}

type GetReturnsListResponseCompensation struct {
	Status interface{} `json:"status"`
	ChangeMoment string `json:"change_moment"`
}

type V1GetFinanceBalanceV1ResponseRevenueMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V3User struct {
	ID string `json:"id"`
	Type string `json:"type_"`
}

type V1GetFinanceBalanceV1ResponseServicesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequestCoordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type ChatChatSendMessageRequest struct {
	Text string `json:"text"`
	ChatID string `json:"chat_id"`
}

type V1DeliveryDetailsDirectDetails struct {
	ByTPLDetails interface{} `json:"by_tpl_details"`
	TimeslotDetails interface{} `json:"timeslot_details"`
	BySellerDetails interface{} `json:"by_seller_details"`
}

type WarehouseWorkingDaysEnum string

type WarehouseDeliveryMethodListResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V3AddresseeFbsLists struct {
	Phone string `json:"phone"`
	Name string `json:"name"`
}

type PostingV4PostingFbsListResponsePostings struct {
	Addressee interface{} `json:"addressee"`
	ExternalOrder interface{} `json:"external_order"`
	QuantumID int64 `json:"quantum_id"`
	Substatus string `json:"substatus"`
	Cancellation interface{} `json:"cancellation"`
	DeliverySchema string `json:"delivery_schema"`
	DestinationPlaceName string `json:"destination_place_name"`
	IsMultibox bool `json:"is_multibox"`
	Requirements interface{} `json:"requirements"`
	Barcodes interface{} `json:"barcodes"`
	IsExpress bool `json:"is_express"`
	VolumeWeight float64 `json:"volume_weight"`
	DeliveringDate string `json:"delivering_date"`
	DeliveryMethod interface{} `json:"delivery_method"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	InProcessAt string `json:"in_process_at"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	Optional interface{} `json:"optional"`
	OrderID int64 `json:"order_id"`
	OrderNumber string `json:"order_number"`
	FinancialData interface{} `json:"financial_data"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	PostingNumber string `json:"posting_number"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	ShipmentDate string `json:"shipment_date"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	Tariffication interface{} `json:"tariffication"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	Customer interface{} `json:"customer"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
	AnalyticsData interface{} `json:"analytics_data"`
	IsPresortable bool `json:"is_presortable"`
	LegalInfo interface{} `json:"legal_info"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	TrackingNumber string `json:"tracking_number"`
	AvailableActions []interface{} `json:"available_actions"`
	ParentPostingNumber string `json:"parent_posting_number"`
	PrrOption string `json:"prr_option"`
}

type ProductImportProductsBySKURequestItem struct {
	CurrencyCode string `json:"currency_code"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	OldPrice string `json:"old_price"`
	Price string `json:"price"`
	SKU int64 `json:"sku"`
	Vat string `json:"vat"`
}

type ProductImportProductsPricesRequestPrice struct {
	ProductID int64 `json:"product_id"`
	CurrencyCode string `json:"currency_code"`
	MinPrice string `json:"min_price"`
	OldPrice string `json:"old_price"`
	Price string `json:"price"`
	PriceStrategyEnabled string `json:"price_strategy_enabled"`
	Vat string `json:"vat"`
	AutoActionEnabled string `json:"auto_action_enabled"`
	ManageElasticBoostingThroughPrice bool `json:"manage_elastic_boosting_through_price"`
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	NetPrice string `json:"net_price"`
	OfferID string `json:"offer_id"`
}

type V2FbsPostingProduct struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type GetConditionalCancellationListV2RequestFilters struct {
	CancellationInitiator []interface{} `json:"cancellation_initiator"`
	PostingNumber []interface{} `json:"posting_number"`
	State interface{} `json:"state"`
}

type V1UnarchiveWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type PostingV4PostingFbsUnfulfilledListRequestSortDirEnum string

type V2ProductInfoPicturesResponseItem struct {
	ProductID int64 `json:"product_id"`
	PrimaryPhoto []interface{} `json:"primary_photo"`
	Photo []interface{} `json:"photo"`
	ColorPhoto []interface{} `json:"color_photo"`
	Photo360 []interface{} `json:"photo_360"`
	Errors []interface{} `json:"errors"`
}

type Productv5GetProductListRequestFilterFilterVisibility string

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponse struct {
	Products []interface{} `json:"products"`
	Total int64 `json:"total"`
}

type Postingv3GetFbsPostingUnfulfilledListRequestFilter struct {
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	WarehouseID []interface{} `json:"warehouse_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveringDateFrom string `json:"delivering_date_from"`
	DeliveringDateTo string `json:"delivering_date_to"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	FbpFilter string `json:"fbpFilter"`
	ProviderID []interface{} `json:"provider_id"`
	Status string `json:"status"`
}

type V1GetProductStairwayDiscountByQuantityResponseStairways struct {
	Enabled bool `json:"enabled"`
	SKU int64 `json:"sku"`
	Stairway interface{} `json:"stairway"`
	Status interface{} `json:"status"`
}

type V1GetProductInfoSubscriptionRequest struct {
	Skus []interface{} `json:"skus"`
}

type Productv3GetProductListResponseItem struct {
	Quants []interface{} `json:"quants"`
	Archived bool `json:"archived"`
	HasFBOStocks bool `json:"has_fbo_stocks"`
	HasFBSStocks bool `json:"has_fbs_stocks"`
	IsDiscounted bool `json:"is_discounted"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
}

type RpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type ItemBundleSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemCount int64 `json:"total_item_count"`
	TotalQuantity int64 `json:"total_quantity"`
}

type V4GetProductInfoStocksRequestFilter struct {
	OfferID []interface{} `json:"offer_id"`
	ProductID []interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
	WithQuant interface{} `json:"with_quant"`
}

type GetProductInfoPricesResponseItemPriceIndexes struct {
	ExternalIndexData interface{} `json:"external_index_data"`
	OzonIndexData interface{} `json:"ozon_index_data"`
	SelfMarketplacesIndexData interface{} `json:"self_marketplaces_index_data"`
	ColorIndex string `json:"color_index"`
}

type V1FbpGetLabelResponse struct {
	LabelURL string `json:"label_url"`
	State interface{} `json:"state"`
}

type V1QuestionAnswerListResponseAnswers struct {
	PublishedAt interface{} `json:"published_at"`
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
	StatusPublication interface{} `json:"status_publication"`
	Text string `json:"text"`
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
}

type CommentListResponseComment struct {
	ID string `json:"id"`
	IsOfficial bool `json:"is_official"`
	IsOwner bool `json:"is_owner"`
	ParentCommentID string `json:"parent_comment_id"`
	PublishedAt string `json:"published_at"`
	Text string `json:"text"`
}

type V1GetFinanceBalanceV1ResponseSalesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type FilterStatusEnum string

type V3Customer struct {
	Address interface{} `json:"address"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type ExemplarsMarks struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"`
}

type ReportCreateCompanyProductsReportRequest struct {
	Language interface{} `json:"language"`
	OfferID []interface{} `json:"offer_id"`
	Search string `json:"search"`
	SKU []interface{} `json:"sku"`
	Visibility interface{} `json:"visibility"`
}

type V1RolesByTokenResponse struct {
	ExpiresAt string `json:"expires_at"`
	Roles []interface{} `json:"roles"`
}

type DeliveryMethodListRequestFilter struct {
	WarehouseID int64 `json:"warehouse_id"`
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftGetResponse struct {
	ID int64 `json:"id"`
	SupplyID string `json:"supply_id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	RowVersion int64 `json:"row_version"`
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	DeclineReason interface{} `json:"decline_reason"`
	DeliveryDetails interface{} `json:"delivery_details"`
	Locked bool `json:"locked"`
	BundleID string `json:"bundle_id"`
	IsCancelable bool `json:"is_cancelable"`
	IsDeletable bool `json:"is_deletable"`
	IsRegistrationAvailable bool `json:"is_registration_available"`
	PackageUnitsCount int32 `json:"package_units_count"`
	CancellationState interface{} `json:"cancellation_state"`
	Editable bool `json:"editable"`
}

type V1SellerActionsProductsAddRequestProduct struct {
	Currency interface{} `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	SKU int64 `json:"sku"`
}

type GetDiscountTaskListV2ResponseTaskDiscountTaskStatusEnum string

type ReviewInfoResponseVideo struct {
	URL string `json:"url"`
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	PreviewURL string `json:"preview_url"`
	ShortVideoPreviewURL string `json:"short_video_preview_url"`
}

type MarketingAction struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Title string `json:"title"`
	Value int32 `json:"value"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1GetStrategyItemsResponseResult struct {
	ProductID []interface{} `json:"product_id"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccruals struct {
	Accruals []interface{} `json:"accruals"`
	PostingNumber string `json:"posting_number"`
}

type CreateWarehouseFBSRequestOptions struct {
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	Comment string `json:"comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
}

type Productv5Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type V1FbpOrderPickUpCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V3GetProductInfoListResponse struct {
	Items []interface{} `json:"items"`
}

type GetProductInfoListResponseVisibilityDetails struct {
	HasStock bool `json:"has_stock"`
	HasPrice bool `json:"has_price"`
}

type V4GetProductInfoStocksResponse struct {
	Cursor string `json:"cursor"`
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type FbpWarehouseListResponseAddressDetailing struct {
	Zipcode string `json:"zipcode"`
	City string `json:"city"`
	Country string `json:"country"`
	House string `json:"house"`
	Region string `json:"region"`
	Street string `json:"street"`
}

type ReportReportInfoRequest struct {
	Code string `json:"code"`
}

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"`
}

type V1SetPostingsResponse struct {
	Result interface{} `json:"result"`
}

type AnalyticsAnalyticsGetDataRequest struct {
	Dimension []interface{} `json:"dimension"`
	Filters []interface{} `json:"filters"`
	Limit int64 `json:"limit"`
	Metrics []interface{} `json:"metrics"`
	Offset int64 `json:"offset"`
	Sort []interface{} `json:"sort"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1QuestionAnswerListResponse struct {
	Answers interface{} `json:"answers"`
	LastID string `json:"last_id"`
}

type V5FbsPostingProductExemplarValidateV5RequestProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type ParameterPickedSegment struct {
	Segments []interface{} `json:"segments"`
}

type OperationService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2Product struct {
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
	SKU int64 `json:"sku"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	FreeStock int64 `json:"free_stock"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
}

type V1CommentListRequest struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	ReviewID string `json:"review_id"`
	SortDir interface{} `json:"sort_dir"`
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V3FbsPostingDetailRelatedPostings struct {
	RelatedPostingNumbers interface{} `json:"related_posting_numbers"`
}

type PostingFinancialDataProduct struct {
	CommissionAmount float64 `json:"commission_amount"`
	CommissionsCurrencyCode string `json:"commissions_currency_code"`
	OldPrice float64 `json:"old_price"`
	Payout float64 `json:"payout"`
	Price float64 `json:"price"`
	CommissionPercent int64 `json:"commission_percent"`
	CustomerPrice float64 `json:"customer_price"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
	CurrencyCode string `json:"currency_code"`
	CustomerCurrencyCode string `json:"customer_currency_code"`
}

type GetCarriageAvailableListResponseResult struct {
	CarriagePostingsCount int32 `json:"carriage_postings_count"`
	DeliveryMethodName string `json:"delivery_method_name"`
	FirstMileType string `json:"first_mile_type"`
	MandatoryPostingsCount int32 `json:"mandatory_postings_count"`
	RecommendedTimeLocal string `json:"recommended_time_local"`
	WarehouseName string `json:"warehouse_name"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	RecommendedTimeUtcOffsetInMinutes float64 `json:"recommended_time_utc_offset_in_minutes"`
	TPLProviderIconURL string `json:"tpl_provider_icon_url"`
	CarriageID int64 `json:"carriage_id"`
	CutoffAt string `json:"cutoff_at"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseTimezone string `json:"warehouse_timezone"`
	CarriageStatus string `json:"carriage_status"`
	Errors interface{} `json:"errors"`
	MandatoryPackagedCount int32 `json:"mandatory_packaged_count"`
	TPLProviderName string `json:"tpl_provider_name"`
	WarehouseCity string `json:"warehouse_city"`
}

type V1ProductPricesDetailsResponsePrice struct {
	CustomerPrice interface{} `json:"customer_price"`
	DiscountPercent float64 `json:"discount_percent"`
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
	PriceIndexes []interface{} `json:"price_indexes"`
	SKU int64 `json:"sku"`
}

type ChatInfoChatStatusEnum string

type V4FbsPostingShipPackageV4RequestProduct struct {
	ExemplarsIds []interface{} `json:"exemplarsIds"`
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

type V1FbpDraftPickupDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type SellerApiProductV1Response struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsCreateVoucherResponse struct {
	ActionID int64 `json:"action_id"`
}

type AnalyticsProductQueriesDetailsResponseQuery struct {
	UniqueViewUsers int64 `json:"unique_view_users"`
	ViewConversion float64 `json:"view_conversion"`
	Currency string `json:"currency"`
	Gmv float64 `json:"gmv"`
	OrderCount int64 `json:"order_count"`
	Position float64 `json:"position"`
	Query string `json:"query"`
	SKU int64 `json:"sku"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
	QueryIndex int64 `json:"query_index"`
}

type GetFinanceBalanceV1ResponseReturns struct {
	Amount interface{} `json:"amount"`
	AmountDetails interface{} `json:"amount_details"`
	Fee interface{} `json:"fee"`
}

type AddStrategyItemsResponseError struct {
	Code string `json:"code"`
	Error string `json:"error"`
	ProductID int64 `json:"product_id"`
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []interface{} `json:"related_posting_numbers"`
}

type V1FbpDraftDropOffRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type FbsPostingMoveStatusResponseMoveStatus struct {
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
	Result bool `json:"result"`
}

type V1GetAttributesRequest struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	Language interface{} `json:"language"`
	TypeID int64 `json:"type_id"`
}

type V1ProductUpdateOfferIdResponse struct {
	Errors interface{} `json:"errors"`
}

type GetWarehouseFBSOperationStatusResponseResult struct {
	EntityID int64 `json:"entity_id"`
}

type V1QuestionCountResponse struct {
	Viewed int64 `json:"viewed"`
	All int64 `json:"all"`
	New int64 `json:"new"`
	Processed int64 `json:"processed"`
	Unprocessed int64 `json:"unprocessed"`
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type V1GetProductInfoDiscountedRequest struct {
	DiscountedSkus interface{} `json:"discounted_skus"`
}

type V1QuestionAnswerCreateDefault struct {
	Code int32 `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

type ReviewV2ReviewListV2RequestFiltersOrderStatusEnum string

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"`
}

type V1SearchQueriesTextResponseSearchQuery struct {
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
	SellersCount float64 `json:"sellers_count"`
	AddToCart float64 `json:"add_to_cart"`
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
	ConversionToCart float64 `json:"conversion_to_cart"`
}

type V1FbpDraftDirectProductValidateResponseApprovedItem struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type GetWarehouseFBSOperationStatusResponseStatusEnum string

type V1FbpDraftDirectProductValidateResponseRejectedItem struct {
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
}

type V1SellerActionsCreateVoucherRequestVoucherParameter struct {
	CountCodes int64 `json:"count_codes"`
	IsPrivate bool `json:"is_private"`
	Type interface{} `json:"type_"`
}

type V1ProductUpdateOfferIdRequest struct {
	UpdateOfferID interface{} `json:"update_offer_id"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearch struct {
	Address string `json:"address"`
	Types []interface{} `json:"types"`
}

type V1AssemblyCarriagePostingListResponsePosting struct {
	AssemblyCode string `json:"assembly_code"`
	CanPrintLabel bool `json:"can_print_label"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequestToUpdate struct {
	Currency string `json:"currency"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	ActionPrice float64 `json:"action_price"`
}

type V1FbpDraftDirectSellerDlvCreateRequestDirectDetails struct {
	DriverName string `json:"driver_name"`
	TimeslotStart string `json:"timeslot_start"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []interface{} `json:"tracking_numbers"`
}

type V1GetDiscountTaskListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
	Status interface{} `json:"status"`
}

type V1SellerActionsArchiveRequest struct {
	ActionID int64 `json:"action_id"`
}

type V1ProductGetRelatedSKURequest struct {
	SKU interface{} `json:"sku"`
}

type Productv1ProductImportPicturesRequest struct {
	ColorImage string `json:"color_image"`
	Images interface{} `json:"images"`
	Images360 interface{} `json:"images360"`
	ProductID int64 `json:"product_id"`
}

type CommonCreateReportResponse struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsCreateDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type SellerApiGetSellerProductV1Response struct {
	Result interface{} `json:"result"`
}

type AnalyticsDataRow struct {
	Dimensions []interface{} `json:"dimensions"`
	Metrics []interface{} `json:"metrics"`
}

type FbpCreateActResponseCreateActErrorReason string

type ArrivalpassArrivalPassListResponseArrivalPass struct {
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DropoffPointID int64 `json:"dropoff_point_id"`
	IsActive bool `json:"is_active"`
	VehicleModel string `json:"vehicle_model"`
	ArrivalPassID int64 `json:"arrival_pass_id"`
	ArrivalReasons []interface{} `json:"arrival_reasons"`
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftPickupCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type Fbsv4PostingProductDetailWithoutDimensions struct {
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	CurrencyCode string `json:"currency_code"`
	MandatoryMark interface{} `json:"mandatory_mark"`
	Name string `json:"name"`
}

type V1OrderDraftValidationErrorErrorType string

type V1CommentCreateResponse struct {
	CommentID string `json:"comment_id"`
}

type ErrorErrorLevel string

type FbsPostingShipV4RequestPackageProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

type Fbpv1DeliveryDetails struct {
	DirectDetails interface{} `json:"direct_details"`
	DropOffPoint interface{} `json:"drop_off_point"`
	PickupDetails interface{} `json:"pickup_details"`
	SupplyType interface{} `json:"supply_type"`
}

type ProductGetProductAttributesV3ResponseComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type ReviewV2ReviewListV2RequestFiltersStatusEnum string

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexDataSelf struct {
	MinPrice interface{} `json:"min_price"`
	PriceIndex float64 `json:"price_index"`
	URL string `json:"url"`
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"`
}

type ProductProductUnarchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type ReportCreateCompanyPostingsReportRequest struct {
	Filter interface{} `json:"filter"`
	Language interface{} `json:"language"`
	With interface{} `json:"with"`
}

type WarehouseFirstMile struct {
	TimeslotFrom string `json:"timeslot_from"`
	TimeslotID int64 `json:"timeslot_id"`
	TimeslotTo string `json:"timeslot_to"`
	Type interface{} `json:"type_"`
	DropoffPointID string `json:"dropoff_point_id"`
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
}

type V1ApproveDiscountTasksRequestTask struct {
	ID int64 `json:"id"`
	ApprovedPrice float64 `json:"approved_price"`
	SellerComment string `json:"seller_comment"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
}

type V3PostingFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
}

type ReviewV2ReviewListV2Response struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
}

type V1ItemSfboAttribute string

type Productv3GetProductListResponseItemQuant struct {
	QuantCode string `json:"quant_code"`
	QuantSize int64 `json:"quant_size"`
}

type GetReturnsListResponseReturnsItem struct {
	Exemplars []interface{} `json:"exemplars"`
	ID int64 `json:"id"`
	ReturnReasonName string `json:"return_reason_name"`
	Type string `json:"type_"`
	OrderID int64 `json:"order_id"`
	Product interface{} `json:"product"`
	ClearingID int64 `json:"clearing_id"`
	ReturnClearingID int64 `json:"return_clearing_id"`
	Schema string `json:"schema"`
	TargetPlace interface{} `json:"target_place"`
	Visual interface{} `json:"visual"`
	PostingNumber string `json:"posting_number"`
	CompensationStatus interface{} `json:"compensation_status"`
	CompanyID int64 `json:"company_id"`
	OrderNumber string `json:"order_number"`
	AdditionalInfo interface{} `json:"additional_info"`
	SourceID int64 `json:"source_id"`
	Place interface{} `json:"place"`
	Storage interface{} `json:"storage"`
	Logistic interface{} `json:"logistic"`
}

type V1FbpDraftPickupDlvEditRequestDeliveryDetails struct {
	SenderPhone string `json:"sender_phone"`
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
}

type V1PostingFbsSplitRequestPosting struct {
	Products []interface{} `json:"products"`
}

type SellerApiProduct struct {
	CurrentBoost float64 `json:"current_boost"`
	PriceMinElastic float64 `json:"price_min_elastic"`
	AlertMaxActionPrice float64 `json:"alert_max_action_price"`
	AddMode string `json:"add_mode"`
	Stock float64 `json:"stock"`
	PriceMaxElastic float64 `json:"price_max_elastic"`
	MinBoost float64 `json:"min_boost"`
	MaxBoost float64 `json:"max_boost"`
	ID float64 `json:"id"`
	Price float64 `json:"price"`
	ActionPrice float64 `json:"action_price"`
	AlertMaxActionPriceFailed bool `json:"alert_max_action_price_failed"`
	MaxActionPrice float64 `json:"max_action_price"`
	MinStock float64 `json:"min_stock"`
}

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnForBackWay int64 `json:"return_for_back_way"`
	ReturnID int64 `json:"return_id"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilter struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
	CutoffFrom string `json:"cutoff_from"`
	DeliveringDateFrom string `json:"delivering_date_from"`
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	ProviderIds []interface{} `json:"provider_ids"`
	Statuses []interface{} `json:"statuses"`
	CutoffTo string `json:"cutoff_to"`
	DeliveringDateTo string `json:"delivering_date_to"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
}

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPointDropOffPointTypeEnum string

type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
	CancelReason string `json:"cancel_reason"`
}

type Polygonv1PolygonBindRequest struct {
	DeliveryMethodID int32 `json:"delivery_method_id"`
	Polygons []interface{} `json:"polygons"`
	WarehouseLocation interface{} `json:"warehouse_location"`
}

type ProductImportProductsPricesResponse struct {
	Result []interface{} `json:"result"`
}

type RpcStatusV1PolygonCreate struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1WarehouseFbsCreateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type PostingV4PostingFbsListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type LanguageLanguage string

type V1OrderErrorTypeEnum string

type QuestionV1GetQuestionListRequestSortDirEnum string

type V4Visibility string

type FinanceV1GetFinanceAccrualByDayRequest struct {
	Date string `json:"date"`
	LastID string `json:"last_id"`
}

type V1AssemblyFbsPostingListResponsePostingProduct struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
}

type Productv3GetProductListResponseResult struct {
	Items interface{} `json:"items"`
	LastID string `json:"last_id"`
	Total int32 `json:"total"`
}

type V1FbpDraftListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"`
}

type ProductV1ProductVisibilitySetRequest struct {
	ItemPlacement []interface{} `json:"item_placement"`
}

type RowItemLegalEntityDocument struct {
	Number string `json:"number"`
	SaleDate string `json:"sale_date"`
}

type SellerReturnsv1MoneyStorage struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1GetSupplyOrderBundleRequest struct {
	Limit int32 `json:"limit"`
	Query string `json:"query"`
	SortField interface{} `json:"sort_field"`
	BundleIds []interface{} `json:"bundle_ids"`
	IsAsc bool `json:"is_asc"`
	ItemTagsCalculation interface{} `json:"item_tags_calculation"`
	LastID string `json:"last_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualNonItemFee struct {
	TypeID int32 `json:"type_id"`
	Accrued interface{} `json:"accrued"`
}

type V1QuestionAnswerDeleteRequest struct {
	AnswerID string `json:"answer_id"`
	SKU int64 `json:"sku"`
}

type V1SetPostingsRequest struct {
	CarriageID int64 `json:"carriage_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1CarriageCancelResponse struct {
	Error string `json:"error"`
	CarriageStatus string `json:"carriage_status"`
}

type ProductInfoWrongVolumeResponseWrongVolumeProduct struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	Weight int64 `json:"weight"`
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	Length int64 `json:"length"`
	Name string `json:"name"`
}

type GetWarehouseFBSOperationStatusResponseTypeEnum string

type V1TimeRangeReturnDate struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type ProductGetImportProductsInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1DraftStatusEnum string

type Productv2GetProductListRequestFilterFilterVisibility string

type V1ProductFbsSplit struct {
	Quantity int64 `json:"quantity"`
	ProductID int64 `json:"product_id"`
}

type GetCompetitorsResponseCompetitorInfo struct {
	Name string `json:"name"`
	ID int64 `json:"id"`
}

type V4GetProductAttributesResponseModelInfo struct {
	ModelID int64 `json:"model_id"`
	Count int64 `json:"count"`
}

type V3Addressee struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V1ArchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type ProductV4GetUploadQuotaResponseOperationLimitsLimitTypeEnum string

type V1TimeRangeStorageTariffication struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type V1AssemblyFbsProductListRequest struct {
	SortDir interface{} `json:"sort_dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ChatStartResponseResult struct {
	ChatID string `json:"chat_id"`
}

type V1FbpDraftDropOffCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type ReportReportListRequest struct {
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	ReportType interface{} `json:"report_type"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejected struct {
	Code interface{} `json:"code"`
	ProductID int64 `json:"product_id"`
	Reason string `json:"reason"`
}

type ReturnsRfbsGetV2ResponseAvailableAction struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type V1FbpDraftDirectTplDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDirectTimeslotEditResponseReserveFailureType string

type SellerSellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type FbsPostingTrackingNumberSetRequestTrackingNumber struct {
	PostingNumber string `json:"posting_number"`
	TrackingNumber string `json:"tracking_number"`
}

type GetProductRatingBySkuResponseProductRating struct {
	SKU int64 `json:"sku"`
	Rating float64 `json:"rating"`
	Groups interface{} `json:"groups"`
}

type AnalyticsAnalyticsGetDataResponse struct {
	Result interface{} `json:"result"`
	Timestamp string `json:"timestamp"`
}

type PostingV4PostingFbsListResponsePostingsRequirements struct {
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
	ProductsRequiringWeight []interface{} `json:"products_requiring_weight"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
}

type V2FbsPostingProductCountryListResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftPickupCreateRequest struct {
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
}

type V1AddStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type GetRealizationReportByDayResponseRow struct {
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission interface{} `json:"delivery_commission"`
}

type V1UpdateWarehouseFBSRequestWorkingDaysEnum string

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
	Valid bool `json:"valid"`
	Errors []interface{} `json:"errors"`
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"`
	PostingNumber string `json:"posting_number"`
}

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1SellerActionsCreateMultiLevelDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPosting struct {
	DeliverySpeed int32 `json:"delivery_speed"`
	Products []interface{} `json:"products"`
	DeliverySchema string `json:"delivery_schema"`
}

type V1FbsPostingProductExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"`
}

type CreateReportResponseCode struct {
	Code string `json:"code"`
}

type V1PostingFbsSplitResponsePostingParent struct {
	Products []interface{} `json:"products"`
	PostingNumber string `json:"posting_number"`
}

type RatingValuePast struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Formatted string `json:"formatted"`
	Status interface{} `json:"status"`
	Value float64 `json:"value"`
}

type SellerApiProductV1ResponseProductDeactivate struct {
	ProductID float64 `json:"product_id"`
	Reason string `json:"reason"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairwayStep struct {
	Discount int64 `json:"discount"`
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
}

type PostingErrorTypeEnum string

type V1FbpCheckConsignmentNoteStateResponse struct {
	ErrorMessage string `json:"error_message"`
	LabelURL string `json:"label_url"`
	State interface{} `json:"state"`
}

type V1GetTreeResponse struct {
	Result []interface{} `json:"result"`
}

type DetailsReturnDetails struct {
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	ReturnServices interface{} `json:"return_services"`
}

type MoneyMoneyBonus struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type ProductCurrencyEnum string

type V1GetFinanceBalanceV1ResponseClosingBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearch struct {
	Address string `json:"address"`
	Types []interface{} `json:"types"`
}

type V3GetProductInfoListRequest struct {
	ProductID []interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
	OfferID []interface{} `json:"offer_id"`
}

type V1FbpAvailableTimeslotListResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type FinanceCashFlowStatementListResponseReturnService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V3ChatListResponse struct {
	Chats interface{} `json:"chats"`
	TotalUnreadCount int64 `json:"total_unread_count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
}

type SellerOzonLogisticsInfoResponseAvailableSchemasEnum string

type V2ReturnsRfbsListResponse struct {
	Returns interface{} `json:"returns"`
}

type PostingPostingProductCancelRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	Items []interface{} `json:"items"`
	PostingNumber string `json:"posting_number"`
}

type ReturnsRfbsGetV2ResponseClientReturnMethodType struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type SellerApiProductPrice struct {
	Stock float64 `json:"stock"`
	ProductID float64 `json:"product_id"`
	ActionPrice float64 `json:"action_price"`
}

type V1AssemblyCarriagePostingListRequestFilter struct {
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1SellerActionsProductsAddRequest struct {
	ActionID int64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplar struct {
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	Valid bool `json:"valid"`
	Errors []interface{} `json:"errors"`
	GTD string `json:"gtd"`
}

type Productv5GetProductInfoPricesV5Response struct {
	Cursor string `json:"cursor"`
	Items interface{} `json:"items"`
	Total int32 `json:"total"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type V2ReturnsRfbsListV2ResponseState struct {
	GroupState string `json:"group_state"`
	MoneyReturnStateName string `json:"money_return_state_name"`
	State string `json:"state"`
	StateName string `json:"state_name"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairways struct {
	SKU int64 `json:"sku"`
	Stairway interface{} `json:"stairway"`
	Enabled bool `json:"enabled"`
}

type DropOffPointCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1ProductUpdateAttributesRequestItem struct {
	Attributes interface{} `json:"attributes"`
	OfferID string `json:"offer_id"`
}

type ListFBSRatingIndexPostingsV1ResponseError struct {
	ChargePercent float64 `json:"charge_percent"`
	ChargePriceCurrencyCode string `json:"charge_price_currency_code"`
	DeliverySchema string `json:"delivery_schema"`
	ErrorAt string `json:"error_at"`
	Index float64 `json:"index"`
	ProductPriceCurrencyCode string `json:"product_price_currency_code"`
	ChargePrice float64 `json:"charge_price"`
	HasGraceStatus bool `json:"has_grace_status"`
	PostingErrorType interface{} `json:"posting_error_type"`
	PostingNumber string `json:"posting_number"`
	ProductPrice float64 `json:"product_price"`
}

type V1ItemError struct {
	Field string `json:"field"`
	AttributeID int64 `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
	Code string `json:"code"`
	Message string `json:"message"`
	State string `json:"state"`
	Level string `json:"level"`
	Description string `json:"description"`
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"`
}

type PostingV4PostingFbsListResponsePostingsDeliveryMethod struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1SellerActionsChangeActivityRequest struct {
	IsTurnOn bool `json:"is_turn_on"`
	ActionID int64 `json:"action_id"`
}

type V1FbpOrderPickUpDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type SellerInfoResponseCompany struct {
	Currency string `json:"currency"`
	Inn string `json:"inn"`
	LegalName string `json:"legal_name"`
	Name string `json:"name"`
	Ogrn string `json:"ogrn"`
	OwnershipForm string `json:"ownership_form"`
	TaxSystem interface{} `json:"tax_system"`
	Country string `json:"country"`
}

type GetReturnsListResponseLogistic struct {
	CancelledWithCompensationMoment string `json:"cancelled_with_compensation_moment"`
	ReturnDate string `json:"return_date"`
	Barcode string `json:"barcode"`
	TechnicalReturnMoment string `json:"technical_return_moment"`
	FinalMoment string `json:"final_moment"`
}

type V1ProductActionTimerUpdateRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type ReturnsRfbsGetV2ResponseRejectionReason struct {
	Hint string `json:"hint"`
	ID int32 `json:"id"`
	IsCommentRequired bool `json:"is_comment_required"`
	Name string `json:"name"`
}

type V1FbpDraftPickupCreateRequestDeliveryDetails struct {
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type V2GetDiscountTaskListV2RequestDiscountTaskStatusEnum string

type V2ConditionalCancellationMoveV2Request struct {
	CancellationID int64 `json:"cancellation_id"`
	Comment string `json:"comment"`
}

type V1GetReturnsListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
	LastID int64 `json:"last_id"`
}

type SellerApiGetSellerActionsV1Response struct {
	Result []interface{} `json:"result"`
}

type V3PostingProductDetail struct {
	Price string `json:"price"`
	CurrencyCode string `json:"currency_code"`
	SKU int64 `json:"sku"`
	Quantity int32 `json:"quantity"`
	HasImei bool `json:"has_imei"`
	Dimensions interface{} `json:"dimensions"`
	MandatoryMark []interface{} `json:"mandatory_mark"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type Productv3GetProductListResponse struct {
	Result interface{} `json:"result"`
}

type V2CancellationStateEnum string

type V1DeleteStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type FbsPostingProductExemplarSetV6RequestProducts struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualAccruedCategoryEnum string

type V2FbsPostingResponse struct {
	Result interface{} `json:"result"`
}

type V2GetConditionalCancellationListV2Request struct {
	Filters interface{} `json:"filters"`
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
	With interface{} `json:"with"`
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []interface{} `json:"result"`
}

type V1SellerActionsCreateVoucherRequest struct {
	VoucherParameters interface{} `json:"voucher_parameters"`
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
}

type ValidationResultItem struct {
	SKU int64 `json:"sku"`
	WeightG float64 `json:"weight_g"`
	Size interface{} `json:"size"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type PolygonBindRequestpolygon struct {
	PolygonID int64 `json:"polygon_id"`
	Time int64 `json:"time"`
}

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
	Dir string `json:"dir"`
}

type V1SellerActionsCreateDiscountWithConditionResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1ProductInfoWrongVolumeResponse struct {
	Cursor string `json:"cursor"`
	Products interface{} `json:"products"`
}

type CreateWarehouseFBSRequestFirstMileTypeEnum string

type V1FbpDraftDropOffDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type CompensationReportLanguage string

type ProductExemplar struct {
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	ExemplarID int64 `json:"exemplar_id"`
}

type V1FbpDraftDropOffProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ItemIDsRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type MoneyMoneyCurrentTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1FbpArchiveListRequest struct {
	Count string `json:"count"`
	LastID string `json:"last_id"`
}

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1CreatePricingStrategyResponseResult struct {
	StrategyID string `json:"strategy_id"`
}

type FinanceTransactionListV3ResponseOperation struct {
	Posting interface{} `json:"posting"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"`
	AccrualsForSale float64 `json:"accruals_for_sale"`
	OperationID int64 `json:"operation_id"`
	SaleCommission float64 `json:"sale_commission"`
	Services []interface{} `json:"services"`
	Type string `json:"type_"`
	Amount float64 `json:"amount"`
	DeliveryCharge float64 `json:"delivery_charge"`
	Items []interface{} `json:"items"`
	OperationDate string `json:"operation_date"`
	OperationType string `json:"operation_type"`
	OperationTypeName string `json:"operation_type_name"`
}

type ProductGetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type GetProductInfoListResponseModelInfo struct {
	Count int64 `json:"count"`
	ModelID int64 `json:"model_id"`
}

type SellerApiProductV1ResponseProduct struct {
	Reason string `json:"reason"`
	ProductID float64 `json:"product_id"`
}

type V2ReturnsRfbsFilter struct {
	PostingNumber string `json:"posting_number"`
	GroupState []interface{} `json:"group_state"`
	CreatedAt interface{} `json:"created_at"`
	OfferID string `json:"offer_id"`
}

type V1AssemblyCarriageProductListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type V3ChatInfo struct {
	UnreadCount int64 `json:"unread_count"`
	Chat interface{} `json:"chat"`
	FirstUnreadMessageID int64 `json:"first_unread_message_id"`
	LastMessageID int64 `json:"last_message_id"`
}

type V1FbpEditTimeslotRequest struct {
	TimeslotStart string `json:"timeslot_start"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type DirectDetailsByTplDetails struct {
	TransportCompanyName string `json:"transport_company_name"`
	TrackingNumber string `json:"tracking_number"`
}

type V1SellerActionsCreateInstallmentResponse struct {
	ActionID int64 `json:"action_id"`
}

type ReturnsRfbsListResponseReturns struct {
	CreatedAt string `json:"created_at"`
	OrderNumber string `json:"order_number"`
	PostingNumber string `json:"posting_number"`
	Product interface{} `json:"product"`
	ReturnID int64 `json:"return_id"`
	ReturnNumber string `json:"return_number"`
	State interface{} `json:"state"`
	ClientName string `json:"client_name"`
}

type ActionsV1ActionsAutoAddProductsListResponse struct {
	Products []interface{} `json:"products"`
	Total int64 `json:"total"`
}

type PriceIndexesIndexOzonData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1FbpOrderDropOffCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type ProductProductArchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type GetProductRatingBySkuResponseRatingImproveAttribute struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type DeliveryMethodListResponseDeliveryMethod struct {
	TemplateID int64 `json:"template_id"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
	Cutoff string `json:"cutoff"`
	ID int64 `json:"id"`
	Status string `json:"status"`
	CompanyID int64 `json:"company_id"`
	CreatedAt string `json:"created_at"`
	Name string `json:"name"`
	ProviderID int64 `json:"provider_id"`
	SlaCutIn int64 `json:"sla_cut_in"`
}

type FinanceTransactionListV3RequestFilter struct {
	Date interface{} `json:"date"`
	OperationType []interface{} `json:"operation_type"`
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
}

type GetReturnsListResponsePlaceTarget struct {
	Name string `json:"name"`
	Address string `json:"address"`
	ID int64 `json:"id"`
}

type V1StrategyRequest struct {
	StrategyID string `json:"strategy_id"`
}

type V1FbpDraftPickUpProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1ItemResponse struct {
	Barcode string `json:"barcode"`
	VolumeInLitres float64 `json:"volume_in_litres"`
	PlacementZone string `json:"placement_zone"`
	Quantity int32 `json:"quantity"`
	Quant int32 `json:"quant"`
	ShipmentType interface{} `json:"shipment_type"`
	IconPath string `json:"icon_path"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	TotalVolumeInLitres float64 `json:"total_volume_in_litres"`
	ContractorItemCode string `json:"contractor_item_code"`
	Tags []interface{} `json:"tags"`
	ProductID int64 `json:"product_id"`
	IsQuantEditable bool `json:"is_quant_editable"`
	SfboAttribute interface{} `json:"sfbo_attribute"`
}

type ImportProductRequestPromotion struct {
	Operation string `json:"operation"`
	Type string `json:"type_"`
}

type V1AnalyticsProductQueriesDetailsResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1FbpDraftDropOffRegistrateResponse struct {
	RowVersion int64 `json:"row_version"`
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
}

type GetProductInfoListResponsePriceIndexes struct {
	ExternalIndexData interface{} `json:"external_index_data"`
	OzonIndexData interface{} `json:"ozon_index_data"`
	SelfMarketplacesIndexData interface{} `json:"self_marketplaces_index_data"`
	ColorIndex interface{} `json:"color_index"`
}

type V1AddStrategyItemsRequest struct {
	StrategyID string `json:"strategy_id"`
	ProductID []interface{} `json:"product_id"`
}

type DirectDetailsTimeslotDetails struct {
	Timeslot interface{} `json:"timeslot"`
	TimeslotReservationID string `json:"timeslot_reservation_id"`
}

type V1FbpDraftDirectProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type V1SellerActionsListResponse struct {
	Total int64 `json:"total"`
	Actions []interface{} `json:"actions"`
}

type MoneyMoneySaleCommission struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type ArrivalpassArrivalPassListResponse struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	Cursor string `json:"cursor"`
}

type V1SellerActionsUpdateDiscountRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type SellerApiProductIDsV1Request struct {
	ActionID float64 `json:"action_id"`
	ProductIds []interface{} `json:"product_ids"`
}

type MoneyMoneyNextTariffCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1DeclineDiscountTasksRequestTask struct {
	ID int64 `json:"id"`
	SellerComment string `json:"seller_comment"`
}

type V1GetStrategyResponseResult struct {
	Competitors []interface{} `json:"competitors"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Type string `json:"type_"`
	UpdateType string `json:"update_type"`
}

type V1ProductInfoWarehouseStocksRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ProductUpdateDiscountRequest struct {
	Discount int32 `json:"discount"`
	ProductID int64 `json:"product_id"`
}

type ReportCreateCompanyProductsReportRequestVisibility string

type PostingV4PostingFbsListResponsePostingsAnalyticsData struct {
	DeliveryType string `json:"delivery_type"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region string `json:"region"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	City string `json:"city"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
}

type V1SellerActionsProductsCandidatesResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V5FbsPostingProductExemplarValidateV5Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type ProductInfoWarehouseStocksResponseStocks struct {
	SKU int64 `json:"sku"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
	FreeStock int64 `json:"free_stock"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponse struct {
	BelowMinPrice []interface{} `json:"below_min_price"`
	ExtremelyLowPrice []interface{} `json:"extremely_low_price"`
	FailedPrice []interface{} `json:"failed_price"`
	Rejected []interface{} `json:"rejected"`
	UpdatedIds []interface{} `json:"updated_ids"`
}

type ProductV1ProductVisibilitySetResponseItems struct {
	Warnings []interface{} `json:"warnings"`
	SelectPermission interface{} `json:"select_permission"`
	SellerItemPlacement interface{} `json:"seller_item_placement"`
	SellerItemPlacementList []interface{} `json:"seller_item_placement_list"`
	ShowcasesVisibility interface{} `json:"showcases_visibility"`
	ShowcasesVisibilityList []interface{} `json:"showcases_visibility_list"`
	SKU int64 `json:"sku"`
}

type V1AnalyticsProductQueriesDetailsRequest struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	LimitBySKU int32 `json:"limit_by_sku"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Skus []interface{} `json:"skus"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
}

type TimetableWorkingHours struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type PostingV4PostingFbsListResponsePostingsCustomerAddress struct {
	Region string `json:"region"`
	AddressTail string `json:"address_tail"`
	City string `json:"city"`
	Comment string `json:"comment"`
	District string `json:"district"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
}

type FbpCheckConsignmentNoteStateResponseStateType string

type PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type V3DeliveryMethod struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
}

type PostingV4PostingFbsListResponsePostingsTarifficationStep struct {
	MinCharge interface{} `json:"min_charge"`
	TariffCharge interface{} `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"`
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductCommission struct {
	CommissionRatio string `json:"commission_ratio"`
	SaleAmount interface{} `json:"sale_amount"`
	SaleCommission interface{} `json:"sale_commission"`
	SalePrice interface{} `json:"sale_price"`
	SellerPrice interface{} `json:"seller_price"`
	Bonus interface{} `json:"bonus"`
	Coinvestment interface{} `json:"coinvestment"`
	Commission interface{} `json:"commission"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type V1SearchAttributeValuesResponse struct {
	Result []interface{} `json:"result"`
}

type V1SearchQueriesTextResponse struct {
	Total string `json:"total"`
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
}

type ReportReportListResponse struct {
	Result interface{} `json:"result"`
}

type SellerApiGetSellerProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Limit float64 `json:"limit"`
	LastID float64 `json:"last_id"`
}

type V3FbsTariffication struct {
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	CurrentTariffType string `json:"current_tariff_type"`
	CurrentTariffChargeCurrencyCode string `json:"current_tariff_charge_currency_code"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffType string `json:"next_tariff_type"`
	NextTariffCharge string `json:"next_tariff_charge"`
	CurrentTariffCharge string `json:"current_tariff_charge"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	NextTariffChargeCurrencyCode string `json:"next_tariff_charge_currency_code"`
}

type WarehouseFirstMileType struct {
	DropoffPointID string `json:"dropoff_point_id"`
	DropoffTimeslotID int64 `json:"dropoff_timeslot_id"`
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
	FirstMileType string `json:"first_mile_type"`
}

type V1ApproveDeclineDiscountTasksResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1GetFinanceBalanceV1ResponseAccruedMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type Productv2ProductsStocksResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason string

type ActionParameterVoucherParameter struct {
	CountCodes int64 `json:"count_codes"`
	IsPrivate bool `json:"is_private"`
	Type interface{} `json:"type_"`
}

type V1SetProductStairwayDiscountByQuantityResponseError struct {
	Data []interface{} `json:"data"`
	SKU int64 `json:"sku"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityEnum string

type V1ProductGetRelatedSKUResponseError struct {
	Code string `json:"code"`
	SKU int64 `json:"sku"`
	Message string `json:"message"`
}

type MoneyMoneySelf struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V3ImportProductsRequestItem struct {
	DimensionUnit string `json:"dimension_unit"`
	Promotions []interface{} `json:"promotions"`
	OfferID string `json:"offer_id"`
	PDFList []interface{} `json:"pdf_list"`
	Weight int32 `json:"weight"`
	GeoNames interface{} `json:"geo_names"`
	Height int32 `json:"height"`
	PrimaryImage string `json:"primary_image"`
	ServiceType interface{} `json:"service_type"`
	OldPrice string `json:"old_price"`
	WeightUnit string `json:"weight_unit"`
	Width int32 `json:"width"`
	Attributes []interface{} `json:"attributes"`
	NewDescriptionCategoryID int64 `json:"new_description_category_id"`
	ColorImage string `json:"color_image"`
	Images360 []interface{} `json:"images360"`
	Name string `json:"name"`
	TypeID int64 `json:"type_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	CurrencyCode string `json:"currency_code"`
	Images []interface{} `json:"images"`
	Price string `json:"price"`
	Vat string `json:"vat"`
	Barcode string `json:"barcode"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Depth int32 `json:"depth"`
}

type DeleteProductsResponseDeleteStatus struct {
	Error string `json:"error"`
	IsDeleted bool `json:"is_deleted"`
	OfferID string `json:"offer_id"`
}

type V1UpdateWarehouseFBSFirstMileRequestFirstMileTypeEnum string

type FilterOp interface{}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress struct {
	District string `json:"district"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	City string `json:"city"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Region string `json:"region"`
	ZipCode string `json:"zip_code"`
	AddressTail string `json:"address_tail"`
	Comment string `json:"comment"`
	Country string `json:"country"`
}

type V1GetRealizationReportPostingResponse struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"`
}

type DetailsService struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type V3ImportProductsRequest struct {
	Items []interface{} `json:"items"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityListEnum string

type V1ListDropOffPointsForCreateFBSWarehouseResponseDropOffPointTypeEnum string

type V1AssemblyFbsProductListRequestFilter struct {
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1AddStrategyItemsResponseResult struct {
	Errors []interface{} `json:"errors"`
	FailedProductCount int32 `json:"failed_product_count"`
}

type V1SellerActionsUpdateDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type PickedSegmentSegmentTypeEnum string

type V1AnalyticsProductQueriesDetailsRequestSortDir string

type ProductV1ProductVisibilityInfoResponse struct {
	Items []interface{} `json:"items"`
}

type ProductImportProductsPricesRequest struct {
	Prices []interface{} `json:"prices"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type MoneyPostingMoney struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type ItemMarketing struct {
	Actions []interface{} `json:"actions"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts struct {
	Quantity int64 `json:"quantity"`
	Commission interface{} `json:"commission"`
	ProductID int64 `json:"product_id"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
	CustomerPrice interface{} `json:"customer_price"`
	OldPrice float64 `json:"old_price"`
	Payout float64 `json:"payout"`
	Price float64 `json:"price"`
}

type V2ProductInfoPicturesResponse struct {
	Items []interface{} `json:"items"`
}

type V3Dimensions struct {
	Length string `json:"length"`
	Weight string `json:"weight"`
	Width string `json:"width"`
	Height string `json:"height"`
}

type V1SellerActionsUpdateVoucherRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V2PostingFBSActGetPostingsResult struct {
	SellerError string `json:"seller_error"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	Products []interface{} `json:"products"`
	ID int64 `json:"id"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Status string `json:"status"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate struct {
	To string `json:"to"`
	From string `json:"from"`
}

type ProductBooleanResponse struct {
	Result bool `json:"result"`
}

type DeliveryMethodListV2ResponseDeliveryMethod struct {
	Status string `json:"status"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	ProviderID int64 `json:"provider_id"`
	SlaCutIn int64 `json:"sla_cut_in"`
	TemplateID int64 `json:"template_id"`
	TPLDropoffPoint interface{} `json:"tpl_dropoff_point"`
	WarehouseID int64 `json:"warehouse_id"`
	Cutoff string `json:"cutoff"`
	IsExpress bool `json:"is_express"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type V2ReturnsRfbsGetV2ResponseState struct {
	State string `json:"state"`
	StateName string `json:"state_name"`
}

type V1DayOfWeek string

type V1ReviewChangeStatusRequest struct {
	ReviewIds []interface{} `json:"review_ids"`
	Status string `json:"status"`
}

type V1AssemblyCarriagePostingListResponsePostingProduct struct {
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
}

type V1GetProductInfoSubscriptionResponseResult struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V1SellerOzonLogisticsInfoResponse struct {
	AvailableSchemas []interface{} `json:"available_schemas"`
	OzonLogisticsEnabled bool `json:"ozon_logistics_enabled"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponseResult struct {
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}

type PriceIndexesIndexSelfData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1SellerActionsCreateMultiLevelDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	DiscountType interface{} `json:"discount_type"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
}

type V1FbpDraftDirectGetTimeslotResponseTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1ApproveDeclineDiscountTasksResponseResult struct {
	FailDetails []interface{} `json:"fail_details"`
	SuccessCount int32 `json:"success_count"`
	FailCount int32 `json:"fail_count"`
}

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
}

type V1UpdateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V1OrderStatusEnum string

type V1AssemblyCarriageProductListResponseProduct struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	PostingNumbers []interface{} `json:"posting_numbers"`
	ProductName string `json:"product_name"`
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type HumanTextsParam struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type Productv2DeleteProductsRequest struct {
	Products []interface{} `json:"products"`
}

type V1ItemSortField string

type V1GetStrategyListResponse struct {
	Strategies []interface{} `json:"strategies"`
	Total int32 `json:"total"`
}

type ChatChatStartRequest struct {
	PostingNumber string `json:"posting_number"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type AnalyticsDataRowDimension struct {
	Name string `json:"name"`
	ID string `json:"id"`
}

type CompanyTaxSystemEnum string

type V1ProductInfoWarehouseStocksResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Stocks []interface{} `json:"stocks"`
}

type V1GetDiscountTaskListResponse struct {
	Result []interface{} `json:"result"`
}

type WarehouseInvalidProductsGetResponseValidationResult struct {
	ValidationErrors []interface{} `json:"validation_errors"`
	Item interface{} `json:"item"`
	State interface{} `json:"state"`
}

type Financev3FinanceTransactionListV3Request struct {
	Filter interface{} `json:"filter"`
	Page int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type V1QuestionAnswerCreateRequest struct {
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
}

type V1FbpDraftDirectProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type SellerActionsListResponseAction struct {
	HighlightURL string `json:"highlight_url"`
	IsEditable bool `json:"is_editable"`
	IsParticipated bool `json:"is_participated"`
	IsTurnOn bool `json:"is_turn_on"`
	SKUCount int64 `json:"sku_count"`
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
	AllowDelete bool `json:"allow_delete"`
}

type PostingCancelFbsPostingRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	PostingNumber string `json:"posting_number"`
}

type V1PostingFbsSplitResponsePosting struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1FbpGetLabelRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type V1CreatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyName string `json:"strategy_name"`
}

type V1SellerActionsUpdateInstallmentRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type GetSellerActionsV1ResponseAction struct {
	WithTargeting bool `json:"with_targeting"`
	ActionType string `json:"action_type"`
	DateStart string `json:"date_start"`
	DateEnd string `json:"date_end"`
	AutoAddDates []interface{} `json:"auto_add_dates"`
	ParticipatingProductsCount float64 `json:"participating_products_count"`
	IsParticipating bool `json:"is_participating"`
	DiscountValue float64 `json:"discount_value"`
	Description string `json:"description"`
	IsVoucherAction bool `json:"is_voucher_action"`
	BannedProductsCount float64 `json:"banned_products_count"`
	OrderAmount float64 `json:"order_amount"`
	ID float64 `json:"id"`
	Title string `json:"title"`
	FreezeDate string `json:"freeze_date"`
	PotentialProductsCount float64 `json:"potential_products_count"`
	DiscountType string `json:"discount_type"`
}

type V1GetAttributeValuesRequest struct {
	Language interface{} `json:"language"`
	LastValueID int64 `json:"last_value_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
}

type GetFinanceBalanceV1ResponseSales struct {
	Fee interface{} `json:"fee"`
	Amount interface{} `json:"amount"`
	AmountDetails interface{} `json:"amount_details"`
}

type V1UpdateWarehouseFBSFirstMileResponse struct {
	OperationID string `json:"operation_id"`
}

type RowItemOrder struct {
	PostingNumber string `json:"posting_number"`
	CreatedDate string `json:"created_date"`
}

type V1FbpCreateActResponse struct {
	Errors []interface{} `json:"errors"`
	FileUuid string `json:"file_uuid"`
	IsSuccess bool `json:"is_success"`
}

type SellerInfoResponseSubscription struct {
	IsPremium bool `json:"is_premium"`
	Type interface{} `json:"type_"`
}

type MoneyMoneyTotalAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type ActionsV1ActionsAutoAddProductsCandidatesRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer struct {
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V1UpdateWarehouseFBSRequestOptions struct {
	CourierPhones []interface{} `json:"courier_phones"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	Comment string `json:"comment"`
}

type PriceIndexesIndexDataOzon struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type GetProductInfoListResponseAvailability struct {
	Availability string `json:"availability"`
	Reasons []interface{} `json:"reasons"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type V1FbpWarehouseListResponse struct {
	Warehouses []interface{} `json:"warehouses"`
}

type FinanceV1GetFinanceAccrualByDayResponse struct {
	Accruals []interface{} `json:"accruals"`
	LastID string `json:"last_id"`
}

type V1GetFinanceBalanceV1ResponsePaymentsMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V2Product struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	CurrencyCode string `json:"currency_code"`
	Price int32 `json:"price"`
	SKU int64 `json:"sku"`
}

type V1ArchiveStatusEnum string

type V1CommentListResponse struct {
	Offset int32 `json:"offset"`
	Comments []interface{} `json:"comments"`
}

type RatingValueCurrent struct {
	DateTo string `json:"date_to"`
	Formatted string `json:"formatted"`
	Status interface{} `json:"status"`
	Value float64 `json:"value"`
	DateFrom string `json:"date_from"`
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
}

type V1CreateStockByWarehouseReportRequest struct {
	Language interface{} `json:"language"`
	WarehouseId []interface{} `json:"warehouseId"`
}

type V1FbpDraftDirectTimeslotEditResponse struct {
	RowVersion int64 `json:"row_version"`
	ErrorReasons []interface{} `json:"error_reasons"`
}

type SourceShipmentType string

type V1SellerInfoResponse struct {
	Company interface{} `json:"company"`
	Ratings []interface{} `json:"ratings"`
	Subscription interface{} `json:"subscription"`
}

type FirstMileTypeEnum string

type PostingV4PostingFbsListRequestFilter struct {
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	Since string `json:"since"`
	Statuses []interface{} `json:"statuses"`
	To string `json:"to"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	OrderID int64 `json:"order_id"`
	OrderNumbers []interface{} `json:"order_numbers"`
	ProviderIds []interface{} `json:"provider_ids"`
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
}

type V1FbpOrderDirectCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V3TimeRange struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1AssemblyFbsPostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type V1GetFinanceBalanceV1ResponseOpeningBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V3GetProductInfoListResponseItem struct {
	Barcodes []interface{} `json:"barcodes"`
	ColorImage []interface{} `json:"color_image"`
	Errors []interface{} `json:"errors"`
	Images360 []interface{} `json:"images360"`
	Name string `json:"name"`
	Price string `json:"price"`
	IsDiscounted bool `json:"is_discounted"`
	ModelInfo interface{} `json:"model_info"`
	OldPrice string `json:"old_price"`
	PrimaryImage []interface{} `json:"primary_image"`
	Statuses interface{} `json:"statuses"`
	CreatedAt string `json:"created_at"`
	IsKgt bool `json:"is_kgt"`
	IsPrepaymentAllowed bool `json:"is_prepayment_allowed"`
	IsSuper bool `json:"is_super"`
	Vat string `json:"vat"`
	IsAutoarchived bool `json:"is_autoarchived"`
	MinPrice string `json:"min_price"`
	PriceIndexes interface{} `json:"price_indexes"`
	Promotions []interface{} `json:"promotions"`
	IsArchived bool `json:"is_archived"`
	HasDiscountedFBOItem bool `json:"has_discounted_fbo_item"`
	Sources []interface{} `json:"sources"`
	UpdatedAt string `json:"updated_at"`
	Commissions []interface{} `json:"commissions"`
	CurrencyCode string `json:"currency_code"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Images []interface{} `json:"images"`
	OfferID string `json:"offer_id"`
	SKU int64 `json:"sku"`
	Stocks interface{} `json:"stocks"`
	VisibilityDetails interface{} `json:"visibility_details"`
	DiscountedFBOStocks int32 `json:"discounted_fbo_stocks"`
	ID int64 `json:"id"`
	VolumeWeight float64 `json:"volume_weight"`
	Availabilities []interface{} `json:"availabilities"`
	TypeID int64 `json:"type_id"`
}

type V4GetProductInfoStocksResponseItem struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Stocks []interface{} `json:"stocks"`
}

type V2FbsPostingProductCountryListResponseResult struct {
	Name string `json:"name"`
	CountryIsoCode string `json:"country_iso_code"`
}

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
}

type V1SellerActionsProductsListResponseProductQuantTypeEnum string

type DeliveryMethodListV2RequestSortDirEnum string

type GetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueId int64 `json:"dictionaryValueId"`
	Value string `json:"value"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type Productv1ProductInfoPicturesResponseResult struct {
	Pictures interface{} `json:"pictures"`
}

type Productv2DeleteProductsResponse struct {
	Status []interface{} `json:"status"`
}

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairwayStep struct {
	Discount int64 `json:"discount"`
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
}

type V1AddBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type V1AssemblyFbsPostingListResponse struct {
	Cursor string `json:"cursor"`
	Cutoff string `json:"cutoff"`
	Postings []interface{} `json:"postings"`
}

type RolesByTokenResponseRoles struct {
	Name string `json:"name"`
	Methods []interface{} `json:"methods"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData struct {
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	City string `json:"city"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	IsPremium bool `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region string `json:"region"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
}

type GooglerpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type AnalyticsProductQueriesRequestSortDir string

type Productv3GetProductListRequest struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData struct {
	Products []interface{} `json:"products"`
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
}

type V1ProductActionTimerStatusRequest struct {
	ProductIds interface{} `json:"product_ids"`
}

type Postingv1GetCarriageAvailableListResponse struct {
	Result interface{} `json:"result"`
}

type SortingOrder string

type V1SetProductStairwayDiscountByQuantityResponse struct {
	Errors []interface{} `json:"errors"`
	Warnings []interface{} `json:"warnings"`
	Accepted bool `json:"accepted"`
}

type V1AnalyticsProductQueriesRequest struct {
	Skus []interface{} `json:"skus"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}

type FbpDraftDropOffPointListResponseDropOffPoint struct {
	City string `json:"city"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	NearestDropOffDate string `json:"nearest_drop_off_date"`
	PointAddress string `json:"point_address"`
	ProvinceUuid string `json:"province_uuid"`
}

type V1AssemblyFbsProductListResponse struct {
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
	ProductsCount int32 `json:"products_count"`
}

type V1AssemblyFbsPostingListRequestFilter struct {
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1FbpDraftDirectCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
}

type V2DeliveryMethodListV2Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type SellerInfoResponseRatingTypeEnum string

type SellerReturnsv1MoneyUtilization struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1GetTreeRequest struct {
	Language interface{} `json:"language"`
}

type QuestionV1QuestionAnswerListResponseAnswerStatusPublicationEnum string

type V1DescriptionCategoryTipsRequest struct {
	TypeID []interface{} `json:"type_id"`
}

type V2DeliveryMethodListV2Response struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	DeliveryMethods []interface{} `json:"delivery_methods"`
}

type ReportCreateDiscountedRequest interface{}

type PostingV4PostingFbsListResponsePostingsCustomer struct {
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
}

type Financev3FinanceTransactionTotalsV3Request struct {
	Date interface{} `json:"date"`
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
}

type ChatChatSendFileRequest struct {
	Base64Content string `json:"base64_content"`
	ChatID string `json:"chat_id"`
	Name string `json:"name"`
}

type FinanceTransactionTotalsV3RequestDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type ProductGetProductAttributesV4ResponseAttribute struct {
	ID int64 `json:"id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type ArrivalpassArrivalPassUpdateRequestArrivalPass struct {
	ArrivalPassID int64 `json:"arrival_pass_id"`
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
}

type FbpGetLabelResponseLabelCreationStateTypeEnum string

type V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type FbsPostingShipV4RequestPackage struct {
	Products []interface{} `json:"products"`
}

type GetReturnsListResponseVisual struct {
	ChangeMoment string `json:"change_moment"`
	Status interface{} `json:"status"`
}

type V1FbpOrderGetResponse struct {
	CanBeCancelled bool `json:"can_be_cancelled"`
	CreatedDate string `json:"created_date"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	CancellationState interface{} `json:"cancellation_state"`
	DraftID int64 `json:"draft_id"`
	Locked bool `json:"locked"`
	BundleUuid string `json:"bundle_uuid"`
	HasLabel bool `json:"has_label"`
	ID int64 `json:"id"`
	ReceiveDate string `json:"receive_date"`
	Status interface{} `json:"status"`
	SupplyID string `json:"supply_id"`
	WarehouseID int64 `json:"warehouse_id"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	OrderNumber string `json:"order_number"`
	RowVersion int64 `json:"row_version"`
}

type SellerActionsUpdateMultiLevelDiscountRequestActionParametersDiscountLevel struct {
	OrderAmount float64 `json:"order_amount"`
	DiscountValue float64 `json:"discount_value"`
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	Cursor string `json:"cursor"`
	Errors []interface{} `json:"errors"`
	HasNext bool `json:"has_next"`
}

type ProductImportProductsBySKURequest struct {
	Items []interface{} `json:"items"`
}

type V1FbpOrderListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type V1PostingFbsSplitRequest struct {
	PostingNumber string `json:"posting_number"`
	Postings []interface{} `json:"postings"`
}

type GetStrategyIDsByItemIDsResponseProductInfo struct {
	ProductID int64 `json:"product_id"`
	StrategyID string `json:"strategy_id"`
}

type GetRealizationReportResponseV2Row struct {
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission interface{} `json:"delivery_commission"`
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
}

type V1FbpDraftDirectRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type CancellationErrorCode string

type FinanceCashFlowStatementListResponseDetails struct {
	BeginBalanceAmount float64 `json:"begin_balance_amount"`
	Delivery interface{} `json:"delivery"`
	InvoiceTransfer float64 `json:"invoice_transfer"`
	Payments []interface{} `json:"payments"`
	RFBS interface{} `json:"rfbs"`
	Others interface{} `json:"others"`
	EndBalanceAmount float64 `json:"end_balance_amount"`
	Loan float64 `json:"loan"`
	Period interface{} `json:"period"`
	Return interface{} `json:"return"`
	Services interface{} `json:"services"`
}

type V3GetFbsPostingListResponseV3Result struct {
	HasNext bool `json:"has_next"`
	Postings []interface{} `json:"postings"`
}

type ReviewV2ReviewListV2RequestSortDirEnum string

type ReportReportInfoResponse struct {
	Result interface{} `json:"result"`
}

type MoneyMoneyCustomerPrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type AnalyticsGetDataResponseResult struct {
	Data []interface{} `json:"data"`
	Totals []interface{} `json:"totals"`
}

type FbpDraftDropOffProvinceListResponseProvince struct {
	PointsCount int32 `json:"points_count"`
	ProvinceUuid string `json:"province_uuid"`
	Name string `json:"name"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseBelowMinPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type AnalyticsMetric string

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type V3ChatHistoryResponse struct {
	HasNext bool `json:"has_next"`
	Messages []interface{} `json:"messages"`
}

type V1FbpDraftDropOffCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type PriceIndexesIndexDataSelf struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1SellerActionsUpdateVoucherRequestActionParameters struct {
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountValue float64 `json:"discount_value"`
}

type SellerSellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CarriageID int64 `json:"carriage_id"`
}

type V3CustomerFbsLists struct {
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductV1ProductVisibilitySetResponseItemsErrors struct {
	SKU int64 `json:"sku"`
	Code string `json:"code"`
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type V3GetProductInfoListResponsePromotion struct {
	IsEnabled bool `json:"is_enabled"`
	Type interface{} `json:"type_"`
}

type PostingCancelReason struct {
	ID int64 `json:"id"`
	IsAvailableForCancellation bool `json:"is_available_for_cancellation"`
	Title string `json:"title"`
	TypeID string `json:"type_id"`
}

type FinanceV1GetFinanceAccrualTypesResponse struct {
	AccrualTypes []interface{} `json:"accrual_types"`
}

type V3GetFbsPostingListResponseV3 struct {
	Result interface{} `json:"result"`
}

type CreateWarehouseFBSRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1AssemblyFbsPostingListResponsePosting struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	AssemblyCode string `json:"assembly_code"`
}

type V1FbpCheckActStateResponseStatus string

type ValidationResultValidationErrorTypeEnum string

type GetWarehouseFBSOperationStatusResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1CancellationState struct {
	CancellationError interface{} `json:"cancellation_error"`
	CancellationStatus interface{} `json:"cancellation_status"`
}

type V1FbpArchiveGetResponse struct {
	HasAct bool `json:"has_act"`
	HasLabel bool `json:"has_label"`
	Status interface{} `json:"status"`
	SupplyID string `json:"supply_id"`
	ID int64 `json:"id"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
	DeclineReason interface{} `json:"decline_reason"`
	OrderNumber string `json:"order_number"`
	PackageUnitsCount int32 `json:"package_units_count"`
	ActFileUuid string `json:"act_file_uuid"`
	BundleID string `json:"bundle_id"`
	BusinessFlowTypeID int64 `json:"business_flow_type_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	OrderDraftID int64 `json:"order_draft_id"`
	ReceiveDate string `json:"receive_date"`
	RowVersion int64 `json:"row_version"`
	CreatedDate string `json:"created_date"`
}

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPointAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PostingV4PostingFbsUnfulfilledListRequestWith struct {
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
}

type ReturnsCompanyFbsInfoRequestPagination struct {
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndex struct {
	ExternalIndexData interface{} `json:"external_index_data"`
	SelfIndexData interface{} `json:"self_index_data"`
}

type V1FbpDraftDropOffProductValidateResponseRejectedItem struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
}

type MoneyMoney struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type RelatedPostingCancelReason struct {
	PostingNumber string `json:"posting_number"`
	Reasons []interface{} `json:"reasons"`
}

type V1ProductPricesDetailsRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1CreatePricingStrategyResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyItemInfoRequest struct {
	ProductID int64 `json:"product_id"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type V1DeclineDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplarMark struct {
	ErrorCodes []interface{} `json:"error_codes"`
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
	CheckStatus string `json:"check_status"`
}

type Productv2ProductsStocksRequest struct {
	Stocks []interface{} `json:"stocks"`
}

type V1QuestionAnswerListRequest struct {
	LastID interface{} `json:"last_id"`
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
}

type V3FbsPostingExemplarProductV3 struct {
	Exemplars []interface{} `json:"exemplars"`
	SKU int64 `json:"sku"`
}

type PostingV4PostingFbsListResponsePostingsFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
}

type PostingV1PostingFbpListResponse struct {
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type ReviewListResponseReview struct {
	VideosAmount int32 `json:"videos_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	Status string `json:"status"`
	Text string `json:"text"`
	CommentsAmount int32 `json:"comments_amount"`
	ID string `json:"id"`
	OrderStatus string `json:"order_status"`
	PhotosAmount int32 `json:"photos_amount"`
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	SKU int64 `json:"sku"`
}

type V1Money struct {
	Currency string `json:"currency"`
	Value float64 `json:"value"`
}

type ReasonHumanText struct {
	Text string `json:"text"`
}

type ReturnsCompanyFbsInfoResponseDropOffPoints struct {
	PassInfo interface{} `json:"pass_info"`
	ID int64 `json:"id"`
	PlaceID int64 `json:"place_id"`
	ReturnsCount int32 `json:"returns_count"`
	UtcOffset string `json:"utc_offset"`
	WarehousesIds []interface{} `json:"warehouses_ids"`
	Address string `json:"address"`
	BoxCount int32 `json:"box_count"`
	Name string `json:"name"`
}

type ParameterAutoStopActionReasonEnum string

type GetFinanceBalanceV1ResponseReturnsDetails struct {
	PartnerPrograms interface{} `json:"partner_programs"`
	PointsForDiscounts string `json:"points_for_discounts"`
	Revenue interface{} `json:"revenue"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type Postingv3FbsPostingWithParamsExamplars struct {
	ProductExemplars bool `json:"product_exemplars"`
	RelatedPostings bool `json:"related_postings"`
	Translit bool `json:"translit"`
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"`
}

type PriceIndexesIndexDataExternal struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1FbpArchiveListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type PostingV4PostingFbsListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Translit bool `json:"translit"`
	With interface{} `json:"with"`
	Cursor string `json:"cursor"`
}

type V1GetStrategyItemInfoResponseResult struct {
	StrategyID string `json:"strategy_id"`
	IsEnabled bool `json:"is_enabled"`
	StrategyProductPrice int32 `json:"strategy_product_price"`
	PriceDownloadedAt string `json:"price_downloaded_at"`
	StrategyCompetitorID int64 `json:"strategy_competitor_id"`
	StrategyCompetitorProductURL string `json:"strategy_competitor_product_url"`
}

type GetProductRatingBySkuResponseRatingGroup struct {
	Conditions interface{} `json:"conditions"`
	ImproveAtLeast int64 `json:"improve_at_least"`
	ImproveAttributes interface{} `json:"improve_attributes"`
	Key string `json:"key"`
	Name string `json:"name"`
	Rating float64 `json:"rating"`
	Weight float64 `json:"weight"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Percent int64 `json:"percent"`
}

type V1SellerActionsProductsCandidatesResponseProduct struct {
	Price float64 `json:"price"`
	QuantSize int64 `json:"quant_size"`
	QuantType interface{} `json:"quant_type"`
	ActionPrice float64 `json:"action_price"`
	Currency string `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	IsActive bool `json:"is_active"`
	MinSellerPrice float64 `json:"min_seller_price"`
	ProductID int64 `json:"product_id"`
	SKU []interface{} `json:"sku"`
	BasePrice float64 `json:"base_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type V1WarehouseListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
}

type V1GetProductRatingBySkuRequest struct {
	Skus interface{} `json:"skus"`
}

type V2ReturnsRfbsCompensateRequest struct {
	CompensationAmount string `json:"compensation_amount"`
	ReturnID int64 `json:"return_id"`
}

type ReportMarkedProductsSalesCreateRequestDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result interface{} `json:"result"`
}

type GetProductInfoListResponseSource struct {
	CreatedAt string `json:"created_at"`
	QuantCode string `json:"quant_code"`
	ShipmentType interface{} `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
	Name string `json:"name"`
}

type ReportListResponseResult struct {
	Total int32 `json:"total"`
	Reports []interface{} `json:"reports"`
}

type ProductGetImportProductsInfoResponseResult struct {
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type V1CarriageApproveRequest struct {
	CarriageID int64 `json:"carriage_id"`
	ContainersCount int32 `json:"containers_count"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFee struct {
	Fees []interface{} `json:"fees"`
	SKU int64 `json:"sku"`
}

type V3FbsPostingAnalyticsData struct {
	City string `json:"city"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region string `json:"region"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
}

type RatingStatusEnum string

type V2GetConditionalCancellationListV2Response struct {
	Counter int64 `json:"counter"`
	LastID int64 `json:"last_id"`
	Result []interface{} `json:"result"`
}

type V1AnalyticsProductQueriesResponseItem struct {
	ViewConversion float64 `json:"view_conversion"`
	Gmv float64 `json:"gmv"`
	UniqueViewUsers int64 `json:"unique_view_users"`
	Category string `json:"category"`
	Currency string `json:"currency"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Position float64 `json:"position"`
	SKU int64 `json:"sku"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProducts struct {
	Actions []interface{} `json:"actions"`
	CommissionsCurrencyCode string `json:"commissions_currency_code"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type GetFinanceBalanceV1ResponseSalesDetails struct {
	Revenue interface{} `json:"revenue"`
	PartnerPrograms interface{} `json:"partner_programs"`
	PointsForDiscounts string `json:"points_for_discounts"`
}

type V3ImportProductsResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpAvailableTimeslotListRequest struct {
	SupplyID string `json:"supply_id"`
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
}

type ChatHistoryRequestFilter struct {
	MessageIds []interface{} `json:"message_ids"`
}

type V1ProductGetRelatedSKUResponseItem struct {
	DeletedAt string `json:"deleted_at"`
	DeliverySchema string `json:"delivery_schema"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	Availability string `json:"availability"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod struct {
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
}

type V1DiscountTaskStatus string

type V1AddBarcodeResult struct {
	Code string `json:"code"`
	Error string `json:"error"`
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type V1FbpDraftDropOffProductValidateResponse struct {
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
	ApprovedItems []interface{} `json:"approved_items"`
}

type ReturnsRfbsGetResponseReturns struct {
	ClientPhoto []interface{} `json:"client_photo"`
	CreatedAt string `json:"created_at"`
	OrderNumber string `json:"order_number"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientName string `json:"client_name"`
	Comment string `json:"comment"`
	RejectionComment string `json:"rejection_comment"`
	ReturnMethodDescription string `json:"return_method_description"`
	ReturnNumber string `json:"return_number"`
	Product interface{} `json:"product"`
	State interface{} `json:"state"`
	ClientReturnMethodType interface{} `json:"client_return_method_type"`
	PostingNumber string `json:"posting_number"`
	RejectionReason []interface{} `json:"rejection_reason"`
	ReturnReason interface{} `json:"return_reason"`
	RuPostTrackingNumber string `json:"ru_post_tracking_number"`
	AvailableActions []interface{} `json:"available_actions"`
}

type ArrivalpassArrivalPassListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type MoneyMoneyCoinvestment struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V3Cancellation struct {
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
}

type V3FinanceCashFlowStatementListRequest struct {
	Date interface{} `json:"date"`
	Page int32 `json:"page"`
	WithDetails bool `json:"with_details"`
	PageSize int32 `json:"page_size"`
}

type GetConditionalCancellationListV2ResponseCancellationReason struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type V1ProductGetRelatedSKUResponse struct {
	Items interface{} `json:"items"`
	Errors interface{} `json:"errors"`
}

type CreateCompanyPostingsReportRequestWith struct {
	AdditionalData bool `json:"additional_data"`
	AnalyticsData bool `json:"analytics_data"`
	CustomerData bool `json:"customer_data"`
	JewelryCodes bool `json:"jewelry_codes"`
}

type ErrorData struct {
	Field string `json:"field"`
	Message string `json:"message"`
	Step int64 `json:"step"`
	Value string `json:"value"`
	Code string `json:"code"`
}

type DescriptionCategoryTipsResponseResult struct {
	ImagesURL []interface{} `json:"images_url"`
	InfoURL string `json:"info_url"`
	TypeID int64 `json:"type_id"`
}

type SellerSellerAPIArrivalPassUpdateRequestArrivalPass struct {
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	ID int64 `json:"id"`
}

type Productv3Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type FinanceV1GetFinanceAccrualTypesResponseAccrualType struct {
	Description string `json:"description"`
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type ActionsV1ActionsAutoAddProductsDeleteResponse struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1ReviewCountResponse struct {
	Processed int32 `json:"processed"`
	Total int32 `json:"total"`
	Unprocessed int32 `json:"unprocessed"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponseTimeslot struct {
	ID int64 `json:"id"`
	To string `json:"to"`
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
}

type V1AssemblyCarriagePostingListResponse struct {
	CanPrintMassLabel bool `json:"can_print_mass_label"`
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftDropOffCreateRequestDeliveryDetails struct {
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
}

type MoneyMoneyNextTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type GetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type V1FbpDraftDirectRegistrateRequest struct {
	SupplyID string `json:"supply_id"`
	RowVersion int64 `json:"row_version"`
}

type V1AssemblyCarriagePostingListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
}

type GetProductInfoListResponseCommission struct {
	DeliveryAmount float64 `json:"delivery_amount"`
	Percent float64 `json:"percent"`
	ReturnAmount float64 `json:"return_amount"`
	SaleSchema string `json:"sale_schema"`
	Value float64 `json:"value"`
}

type V1QuestionAnswerCreateResponse struct {
	AnswerID string `json:"answer_id"`
}

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementEnum string

type V3ChatHistoryRequest struct {
	ChatID string `json:"chat_id"`
	Direction string `json:"direction"`
	Filter interface{} `json:"filter"`
	FromMessageID int64 `json:"from_message_id"`
	Limit int64 `json:"limit"`
}

type V1AddBarcodeRequest struct {
	Barcodes []interface{} `json:"barcodes"`
}

type GetReturnsListResponseVisualStatus struct {
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
	SysName string `json:"sys_name"`
}

type V5FbsPostingProductExemplarValidateV5Response struct {
	Products []interface{} `json:"products"`
}

type PostingV4PostingFbsListResponsePostingsLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type V1FbpDraftPickUpRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type FbsPostingProductExemplarCreateOrGetV6ResponseProduct struct {
	IsMandatoryMarkNeeded bool `json:"is_mandatory_mark_needed"`
	Quantity int32 `json:"quantity"`
	Exemplars []interface{} `json:"exemplars"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
	IsMandatoryMarkPossible bool `json:"is_mandatory_mark_possible"`
	IsRnptNeeded bool `json:"is_rnpt_needed"`
	ProductID int64 `json:"product_id"`
	HasImei bool `json:"has_imei"`
	IsJwUinNeeded bool `json:"is_jw_uin_needed"`
}

type GetSupplyOrderBundleRequestItemTagsCalculation struct {
	DropoffWarehouseID string `json:"dropoff_warehouse_id"`
	StorageWarehouseIds []interface{} `json:"storage_warehouse_ids"`
}

type ReviewV2ReviewListV2Request struct {
	Filters interface{} `json:"filters"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type V1SellerActionsProductsListResponseProduct struct {
	OfferID string `json:"offer_id"`
	SKU []interface{} `json:"sku"`
	BasePrice float64 `json:"base_price"`
	Currency string `json:"currency"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantSize int64 `json:"quant_size"`
	QuantType interface{} `json:"quant_type"`
	ActionPrice float64 `json:"action_price"`
	DiscountPercent float64 `json:"discount_percent"`
	IsActive bool `json:"is_active"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Name string `json:"name"`
}

type V1BundleItemShipmentType string

type ActionsV1ActionsAutoAddProductsListResponseProducts struct {
	MinActionQuantity int64 `json:"min_action_quantity"`
	Name string `json:"name"`
	ProductID int64 `json:"product_id"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
	SKU int64 `json:"sku"`
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	AddMode string `json:"add_mode"`
	Currency string `json:"currency"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
	MinSellerPrice float64 `json:"min_seller_price"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
}

type PostingV4PostingFbsListResponsePostingsProducts struct {
	Imei []interface{} `json:"imei"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	Name string `json:"name"`
	ProductColor string `json:"product_color"`
	Quantity int32 `json:"quantity"`
	Weight float64 `json:"weight"`
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
	SKU int64 `json:"sku"`
}

type V1GetWarehouseFBSOperationStatusRequest struct {
	OperationID string `json:"operation_id"`
}

type V1SellerActionsCreateVoucherRequestVoucherParameterTypeEnum string

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequest struct {
	IsKgt bool `json:"is_kgt"`
	AddressCoordinates interface{} `json:"address_coordinates"`
}

type V3ImportProductsResponseResult struct {
	TaskID int64 `json:"task_id"`
}

type ProductsPostings struct {
	PostingNumber string `json:"posting_number"`
	Quantity int32 `json:"quantity"`
}

type Productv2ProductsStocksResponseResult struct {
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Updated bool `json:"updated"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1GetSupplyOrderBundleResponse struct {
	Items []interface{} `json:"items"`
	TotalCount int32 `json:"total_count"`
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
}

type V1GetFinanceBalanceV1ResponsePartnerMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V3GetProductInfoListResponsePromotionType string

type AssemblyFbsProductListResponseProducts struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	Postings []interface{} `json:"postings"`
	ProductName string `json:"product_name"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type GetFinanceBalanceV1ResponseCashflows struct {
	Services []interface{} `json:"services"`
	Returns interface{} `json:"returns"`
	Sales interface{} `json:"sales"`
}

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID int64 `json:"return_id"`
	ReturnMethodDescription string `json:"return_method_description"`
}

type Productv2ProductsStocksResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type DetailsDeliveryDetails struct {
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	DeliveryServices interface{} `json:"delivery_services"`
}

type V1WarehouseInvalidProductsGetResponse struct {
	HasNext bool `json:"has_next"`
	LastID int64 `json:"last_id"`
	ValidationResults []interface{} `json:"validation_results"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1AssemblyFbsPostingListRequestSortDirEnum string

type V1FbpDraftDirectTplDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication struct {
	CurrentTariffCharge interface{} `json:"current_tariff_charge"`
	CurrentTariffMinCharge interface{} `json:"current_tariff_min_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffMinCharge interface{} `json:"next_tariff_min_charge"`
	NextTariffType string `json:"next_tariff_type"`
	NextTariffCharge interface{} `json:"next_tariff_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
}

type PostingCancelReasonResponse struct {
	Result []interface{} `json:"result"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejectedCodeEnum string

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFeeFee struct {
	Accrued interface{} `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type V1GetRealizationReportPostingResponseRow struct {
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
	RowNumber int32 `json:"row_number"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	Order interface{} `json:"order"`
	LegalEntityDocument interface{} `json:"legal_entity_document"`
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission interface{} `json:"delivery_commission"`
}

type V1GenerateBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type DetailsPayment struct {
	CurrencyCode string `json:"currency_code"`
	Payment float64 `json:"payment"`
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type PriceIndexesIndexExternalData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V6FbsPostingProductExemplarCreateOrGetV6Response struct {
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type ReportLanguage string

type TypeTimeOfDay struct {
	Hours int32 `json:"hours"`
	Minutes int32 `json:"minutes"`
	Nanos int32 `json:"nanos"`
	Seconds int32 `json:"seconds"`
}

type V1FbpArchiveListResponseItem struct {
	OrderDraftID int64 `json:"order_draft_id"`
	SupplyID string `json:"supply_id"`
	BundleID string `json:"bundle_id"`
	HasAct bool `json:"has_act"`
	WhcOrderID int64 `json:"whc_order_id"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	DeclineReason interface{} `json:"decline_reason"`
	DeliveryDetails interface{} `json:"delivery_details"`
	HasLabel bool `json:"has_label"`
	ReceiveDate string `json:"receive_date"`
	RowVersion int64 `json:"row_version"`
	Status interface{} `json:"status"`
	ActFileUuid string `json:"act_file_uuid"`
	CreatedDate string `json:"created_date"`
	ExternalOrderID string `json:"external_order_id"`
}

type SellerApiProductV1ResponseResultDeactivate struct {
	ProductIds []interface{} `json:"product_ids"`
	Rejected []interface{} `json:"rejected"`
}

type GetProductRatingBySkuResponseRatingCondition struct {
	Cost float64 `json:"cost"`
	Description string `json:"description"`
	Fulfilled bool `json:"fulfilled"`
	Key string `json:"key"`
}

type GetReturnsListResponseStorage struct {
	TarifficationStartDate string `json:"tariffication_start_date"`
	ArrivedMoment string `json:"arrived_moment"`
	Days int64 `json:"days"`
	UtilizationSum interface{} `json:"utilization_sum"`
	UtilizationForecastDate string `json:"utilization_forecast_date"`
	Sum interface{} `json:"sum"`
	TarifficationFirstDate string `json:"tariffication_first_date"`
}

type OperationPosting struct {
	DeliverySchema string `json:"delivery_schema"`
	OrderDate string `json:"order_date"`
	PostingNumber string `json:"posting_number"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductGetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type ItemPricev5 struct {
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	RetailPrice float64 `json:"retail_price"`
	AutoActionEnabled bool `json:"auto_action_enabled"`
	MarketingSellerPrice float64 `json:"marketing_seller_price"`
	Vat float64 `json:"vat"`
	CurrencyCode string `json:"currency_code"`
	MinPrice float64 `json:"min_price"`
	NetPrice float64 `json:"net_price"`
}

type V1ReturnsCompanyFbsInfoRequestFilter struct {
	PlaceID int64 `json:"place_id"`
}

type GetStrategyListResponseStrategy struct {
	CompetitorsCount int64 `json:"competitors_count"`
	Enabled bool `json:"enabled"`
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type_"`
	UpdateType string `json:"update_type"`
	UpdatedAt string `json:"updated_at"`
	ProductsCount int64 `json:"products_count"`
}

type V1Competitor struct {
	Coefficient float64 `json:"coefficient"`
	CompetitorID int64 `json:"competitor_id"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type V1ReturnsCompanyFbsInfoRequest struct {
	Pagination interface{} `json:"pagination"`
	Filter interface{} `json:"filter"`
}

type MoneyMoneySaleAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V3Address struct {
	ZipCode string `json:"zip_code"`
	AddressTail string `json:"address_tail"`
	City string `json:"city"`
	Comment string `json:"comment"`
	Country string `json:"country"`
	Longitude float64 `json:"longitude"`
	District string `json:"district"`
	Latitude float64 `json:"latitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	Region string `json:"region"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission struct {
	Currency string `json:"currency"`
	Percent int64 `json:"percent"`
	Amount float64 `json:"amount"`
}

type V1SellerActionsCreateDiscountRequest struct {
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	MinActionPercent float64 `json:"min_action_percent"`
}

type GetFinanceBalanceV1ResponseService struct {
	Amount interface{} `json:"amount"`
	Name string `json:"name"`
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type ItemSize struct {
	LengthMm int32 `json:"length_mm"`
	WidthMm int32 `json:"width_mm"`
	HeightMm int32 `json:"height_mm"`
}

type PostingPostingFBSPackageLabelResponse struct {
	FileName string `json:"file_name"`
	ContentType string `json:"content_type"`
	FileContent string `json:"file_content"`
}

type V1FbpDraftPickUpRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProduct struct {
	Error string `json:"error"`
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
	Valid bool `json:"valid"`
}

type GetReturnsListResponseAdditionalInfo struct {
	IsOpened bool `json:"is_opened"`
	IsSuperEconom bool `json:"is_super_econom"`
}

type DeleteProductsRequestProduct struct {
	OfferID string `json:"offer_id"`
}

type V3ImportProductsRequestDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type V1GetWarehouseFBSOperationStatusResponse struct {
	Error interface{} `json:"error"`
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Type interface{} `json:"type_"`
}

type V1GetTreeResponseItem struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	CategoryName string `json:"category_name"`
	Children []interface{} `json:"children"`
	Disabled bool `json:"disabled"`
	TypeID int64 `json:"type_id"`
	TypeName string `json:"type_name"`
}

type ProductV1ProductVisibilityInfoResponseItem struct {
	ShowcasesVisibility interface{} `json:"showcases_visibility"`
	SKU int64 `json:"sku"`
}

type GetDiscountTaskListV2ResponseTask struct {
	Name string `json:"name"`
	OriginalPrice float64 `json:"original_price"`
	RequestedDiscount float64 `json:"requested_discount"`
	CreatedAt string `json:"created_at"`
	EditedTillDuration int64 `json:"edited_till_duration"`
	FirstName string `json:"first_name"`
	ModeratedAt string `json:"moderated_at"`
	Patronymic string `json:"patronymic"`
	RequestedPrice float64 `json:"requested_price"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	SKU int64 `json:"sku"`
	AutoModeratedInfo interface{} `json:"auto_moderated_info"`
	EditedTill string `json:"edited_till"`
	ID int64 `json:"id"`
	IsAutoModerated bool `json:"is_auto_moderated"`
	ReductionFactor float64 `json:"reduction_factor"`
	Status interface{} `json:"status"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	Email string `json:"email"`
	EndAt string `json:"end_at"`
	ApprovedDiscount float64 `json:"approved_discount"`
	ApprovedPrice float64 `json:"approved_price"`
	EndAtDuration int64 `json:"end_at_duration"`
	LastName string `json:"last_name"`
	MinAutoPrice float64 `json:"min_auto_price"`
}

type GetProductAttributesV4ResponseAttribute struct {
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
	ID int64 `json:"id"`
}

type V1FbpDraftDirectProductValidateRequestSkuItem struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1GetStrategyIDsByItemIDsResponseResult struct {
	ProductsInfo []interface{} `json:"products_info"`
}

type ExemplarMark struct {
	MarkType string `json:"mark_type"`
	Mark string `json:"mark"`
}

type SellerActionsListRequestActionTypeEnum string

type GetReturnsListRequestFilter struct {
	ProductName string `json:"product_name"`
	OfferID string `json:"offer_id"`
	VisualStatusName string `json:"visual_status_name"`
	ReturnSchema string `json:"return_schema"`
	CompensationStatusID int32 `json:"compensation_status_id"`
	LogisticReturnDate interface{} `json:"logistic_return_date"`
	StorageTarifficationStartDate interface{} `json:"storage_tariffication_start_date"`
	VisualStatusChangeMoment interface{} `json:"visual_status_change_moment"`
	OrderID int64 `json:"order_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
	WarehouseID int64 `json:"warehouse_id"`
	Barcode string `json:"barcode"`
}

type PostingV4PostingFbsListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1SellerActionsProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	Skus []interface{} `json:"skus"`
}

type V1GetFBSRatingIndexInfoV1Response struct {
	ProcessingCostsSum float64 `json:"processing_costs_sum"`
	CurrencyCode string `json:"currency_code"`
	Defects []interface{} `json:"defects"`
	Index float64 `json:"index"`
	PeriodFrom string `json:"period_from"`
	PeriodTo string `json:"period_to"`
}

type V3ChatListRequestFilter struct {
	ChatStatus string `json:"chat_status"`
	UnreadOnly bool `json:"unread_only"`
}

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"`
}

type V1RatingStatus struct {
	Danger bool `json:"danger"`
	Premium bool `json:"premium"`
	Warning bool `json:"warning"`
}

type ActionParameterVoucherParameterTypeEnum string

type V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type DeliveryDetailsSupplyType string

type CreatedAt struct {
	From string `json:"from"`
	To string `json:"to"`
}

type ProductV1ProductVisibilitySetRequestItemPlacementPlacementEnum string

type WarehouseWarehouseListResponse struct {
	Result []interface{} `json:"result"`
}

type V3FbsPosting struct {
	OrderNumber string `json:"order_number"`
	DeliveryMethod interface{} `json:"delivery_method"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	Status string `json:"status"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	Barcodes interface{} `json:"barcodes"`
	Requirements interface{} `json:"requirements"`
	Substatus string `json:"substatus"`
	AnalyticsData interface{} `json:"analytics_data"`
	IsPresortable bool `json:"is_presortable"`
	Optional interface{} `json:"optional"`
	Tariffication interface{} `json:"tariffication"`
	Customer interface{} `json:"customer"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	Addressee interface{} `json:"addressee"`
	InProcessAt string `json:"in_process_at"`
	LegalInfo interface{} `json:"legal_info"`
	Cancellation interface{} `json:"cancellation"`
	IsExpress bool `json:"is_express"`
	ParentPostingNumber string `json:"parent_posting_number"`
	PostingNumber string `json:"posting_number"`
	DeliveringDate string `json:"delivering_date"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	Products []interface{} `json:"products"`
	ShipmentDate string `json:"shipment_date"`
	TrackingNumber string `json:"tracking_number"`
	AvailableActions interface{} `json:"available_actions"`
	DestinationPlaceName string `json:"destination_place_name"`
	FinancialData interface{} `json:"financial_data"`
	OrderID int64 `json:"order_id"`
}

type TaskAutoModeratedInfo struct {
	MaxPercent float64 `json:"max_percent"`
	MaxPrice float64 `json:"max_price"`
	MinPercent float64 `json:"min_percent"`
	MinPrice float64 `json:"min_price"`
}

type V1Empty interface{}

type PostingV4PostingFbsListResponsePostingsTariffication struct {
	CurrentTariffCharge interface{} `json:"current_tariff_charge"`
	CurrentTariffMinCharge interface{} `json:"current_tariff_min_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffCharge interface{} `json:"next_tariff_charge"`
	NextTariffMinCharge interface{} `json:"next_tariff_min_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	NextTariffType string `json:"next_tariff_type"`
}

type V3ImportProductsRequestAttribute struct {
	ComplexID int64 `json:"complex_id"`
	ID int64 `json:"id"`
	Values []interface{} `json:"values"`
}

type ListFBSRatingIndexPostingsV1RequestFilter struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type ProductGetProductInfoDescriptionRequest struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
}

type FinanceV1GetFinanceAccrualPostingsRequest struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V2GetProductInfoStocksByWarehouseFbsRequestV2 struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	OfferID []interface{} `json:"offer_id"`
	SKU []interface{} `json:"sku"`
}

type V3Barcodes struct {
	UpperBarcode string `json:"upper_barcode"`
	LowerBarcode string `json:"lower_barcode"`
}

type V1ReviewListResponse struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
}

type V2CancellationInitiatorEnum string

type SellerInfoResponseSubscriptionTypeEnum string

type V1SearchQueriesTextRequest struct {
	Offset string `json:"offset"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	Text string `json:"text"`
	Limit string `json:"limit"`
}

type GetProductAttributesResponseImage struct {
	Default bool `json:"default"`
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
}

type ActionParameter struct {
	AutoStopActionReason interface{} `json:"auto_stop_action_reason"`
	Budget float64 `json:"budget"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	DiscountType interface{} `json:"discount_type"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
	Warehouses []interface{} `json:"warehouses"`
	Addresses []interface{} `json:"addresses"`
	DateEnd string `json:"date_end"`
	MinActionPercent float64 `json:"min_action_percent"`
	PickedSegments []interface{} `json:"picked_segments"`
	Status interface{} `json:"status"`
	BudgetSpent float64 `json:"budget_spent"`
	DiscountValue float64 `json:"discount_value"`
	VoucherParameters interface{} `json:"voucher_parameters"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Type interface{} `json:"type_"`
}

type V1GetFinanceBalanceV1ResponseReturnsMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"`
}

type ReviewV2ReviewListV2ResponseReviewStatusEnum string

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V2PostingFBSActGetPostingsRequest struct {
	ID interface{} `json:"id"`
}

type V1QuestionInfoRequest struct {
	QuestionID string `json:"question_id"`
}

type GetImportProductsInfoResponseResultItem struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Status string `json:"status"`
	Errors []interface{} `json:"errors"`
}

type ProductImportProductsBySKUResponseResult struct {
	TaskID int64 `json:"task_id"`
	UnmatchedSKUList []interface{} `json:"unmatched_sku_list"`
}

type PostingCancelReasonListResponse struct {
	Result []interface{} `json:"result"`
}

type ProtobufAny struct {
	TypeUrl string `json:"typeUrl"`
	Value string `json:"value"`
}

type V1ProductUpdateAttributesRequest struct {
	Items interface{} `json:"items"`
}

type WarehouseCarriageLabelTypeEnum string

type V1SearchQueriesTopResponseSearchQuery struct {
	AddToCart float64 `json:"add_to_cart"`
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
	ConversionToCart float64 `json:"conversion_to_cart"`
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
	SellersCount float64 `json:"sellers_count"`
}

type V1WarehouseListRequestWith struct {
	AbleToSetPrice bool `json:"able_to_set_price"`
}

type V1UpdateWarehouseFBSRequest struct {
	WorkingDays []interface{} `json:"working_days"`
	AddressCoordinates interface{} `json:"address_coordinates"`
	Name string `json:"name"`
	Options interface{} `json:"options"`
	Phone string `json:"phone"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []interface{} `json:"provinces"`
}

type V3FinanceCashFlowStatementListResponsePeriod struct {
	ID int64 `json:"id"`
	Begin string `json:"begin"`
	End string `json:"end"`
}

type GetReturnsListResponseProduct struct {
	OfferID string `json:"offer_id"`
	Name string `json:"name"`
	Price interface{} `json:"price"`
	PriceWithoutCommission interface{} `json:"price_without_commission"`
	CommissionPercent float64 `json:"commission_percent"`
	Commission interface{} `json:"commission"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1GetDiscountTaskListResponseTask struct {
	FirstName string `json:"first_name"`
	RequestedPrice float64 `json:"requested_price"`
	DiscountPercent float64 `json:"discount_percent"`
	BasePrice float64 `json:"base_price"`
	RequestedQuantityMin int64 `json:"requested_quantity_min"`
	Status string `json:"status"`
	IsDamaged bool `json:"is_damaged"`
	ApprovedDiscountPercent float64 `json:"approved_discount_percent"`
	IsPurchased bool `json:"is_purchased"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	RequestedPriceWithFee float64 `json:"requested_price_with_fee"`
	SellerComment string `json:"seller_comment"`
	ApprovedPrice float64 `json:"approved_price"`
	LastName string `json:"last_name"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	OfferID string `json:"offer_id"`
	Email string `json:"email"`
	Patronymic string `json:"patronymic"`
	ID int64 `json:"id"`
	CustomerName string `json:"customer_name"`
	SKU int64 `json:"sku"`
	UserComment string `json:"user_comment"`
	OriginalPrice float64 `json:"original_price"`
	CreatedAt string `json:"created_at"`
	EndAt string `json:"end_at"`
	Discount float64 `json:"discount"`
	PrevTaskID int64 `json:"prev_task_id"`
	ApprovedPriceWithFee float64 `json:"approved_price_with_fee"`
	ApprovedPriceFeePercent float64 `json:"approved_price_fee_percent"`
	EditedTill string `json:"edited_till"`
	MinAutoPrice float64 `json:"min_auto_price"`
	ModeratedAt string `json:"moderated_at"`
	ApprovedDiscount float64 `json:"approved_discount"`
	IsAutoModerated bool `json:"is_auto_moderated"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseRequest struct {
	Search interface{} `json:"search"`
	WarehouseID int64 `json:"warehouse_id"`
}

type FinanceCashFlowStatementListResponseDeliveryService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type Productv3GetProductAttributesV3ResponseResult struct {
	Name string `json:"name"`
	WeightUnit string `json:"weight_unit"`
	Width int32 `json:"width"`
	Attributes []interface{} `json:"attributes"`
	Barcode string `json:"barcode"`
	Height int32 `json:"height"`
	OfferID string `json:"offer_id"`
	PDFList []interface{} `json:"pdf_list"`
	CategoryID int64 `json:"category_id"`
	ColorImage string `json:"color_image"`
	DimensionUnit string `json:"dimension_unit"`
	Images []interface{} `json:"images"`
	Images360 []interface{} `json:"images360"`
	Depth int32 `json:"depth"`
	Weight int32 `json:"weight"`
	TypeID int64 `json:"type_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	ID int64 `json:"id"`
	ImageGroupID string `json:"image_group_id"`
}

type V1ArchiveSkuSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemsCount int64 `json:"total_items_count"`
	TotalQuantity int64 `json:"total_quantity"`
}

type V1FbpCheckActStateResponse struct {
	Error interface{} `json:"error"`
	Status interface{} `json:"status"`
	CdnURL string `json:"cdn_url"`
}

type V1ProductPricesDetailsResponse struct {
	Prices []interface{} `json:"prices"`
}

type V1GetProductStairwayDiscountByQuantityResponse struct {
	Stairways []interface{} `json:"stairways"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
}

type V1UpdateStatusStrategyRequest struct {
	Enabled bool `json:"enabled"`
	StrategyID string `json:"strategy_id"`
}

type V1UpdateWarehouseFBSRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1SellerActionsUpdateDiscountWithConditionRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V4GetProductInfoStocksRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type Productv4GetProductAttributesV4Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type GetRealizationReportResponseV2Result struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"`
}

type FinanceCashFlowStatementListResponseService struct {
	Price float64 `json:"price"`
	Name string `json:"name"`
}

type V1CancellationStateStatus string

type V1ReviewInfoResponse struct {
	Videos []interface{} `json:"videos"`
	CommentsAmount int32 `json:"comments_amount"`
	ID string `json:"id"`
	LikesAmount int32 `json:"likes_amount"`
	Rating int32 `json:"rating"`
	SKU int64 `json:"sku"`
	PublishedAt string `json:"published_at"`
	VideosAmount int32 `json:"videos_amount"`
	DislikesAmount int32 `json:"dislikes_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	OrderStatus string `json:"order_status"`
	Photos []interface{} `json:"photos"`
	PhotosAmount int32 `json:"photos_amount"`
	Status string `json:"status"`
	Text string `json:"text"`
}

type V1FbpCheckConsignmentNoteStateRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"`
}

type ProductV1ProductVisibilityInfoRequest struct {
	Skus []interface{} `json:"skus"`
}

type PostingV1PostingFbpListResponsePostingsProducts struct {
	CustomerPrice interface{} `json:"customer_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
	Quantity int32 `json:"quantity"`
	SellerPrice interface{} `json:"seller_price"`
	SKU int64 `json:"sku"`
}

type V1ProductUpdateAttributesRequestValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type GetRealizationReportByDayResponse struct {
	Rows []interface{} `json:"rows"`
}

type V1SellerActionsListRequest struct {
	Offset int64 `json:"offset"`
	Search string `json:"search"`
	Status []interface{} `json:"status"`
	ActionIds []interface{} `json:"action_ids"`
	ActionType []interface{} `json:"action_type"`
	Limit int64 `json:"limit"`
}

type V1WarehouseFbsCreatePickUpTimeslotListResponse struct {
	IsPickupSupported bool `json:"is_pickup_supported"`
	Timeslots []interface{} `json:"timeslots"`
}

type V1SellerActionsVoucherGetRequest struct {
	ActionID int64 `json:"action_id"`
}

type PostingV4PostingFbsListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V1OrderDraftValidationError struct {
	Errors []interface{} `json:"errors"`
}

type Productv2ProductsStocksRequestStock struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Stock int64 `json:"stock"`
	WarehouseID int64 `json:"warehouse_id"`
}

type SellerSellerAPIArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type V2FboSinglePostingLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type V1FbpDraftDirectCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V4FbsPostingShipPackageV4Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1FbpOrderDropOffTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ChatChatStartResponse struct {
	Result interface{} `json:"result"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponseProducts struct {
	BasePrice float64 `json:"base_price"`
	Currency string `json:"currency"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
	MinActionQuantity int64 `json:"min_action_quantity"`
	MinSellerPrice float64 `json:"min_seller_price"`
	OfferID string `json:"offer_id"`
	SKU int64 `json:"sku"`
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
}

type GetUploadQuotaResponseTotal struct {
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
	QuotaByCategory interface{} `json:"quota_by_category"`
}

type V3FbsPostingProduct struct {
	Imei []interface{} `json:"imei"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	CurrencyCode string `json:"currency_code"`
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	TransportCompanyName string `json:"transport_company_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TrackingNumber string `json:"tracking_number"`
}

type GetReturnsListResponseCompensationStatus struct {
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
	SysName string `json:"sys_name"`
}

type ChatChatSendFileResponse struct {
	Result string `json:"result"`
}

type ProductV1ProductVisibilitySetResponse struct {
	Items []interface{} `json:"items"`
	ItemsErrors []interface{} `json:"items_errors"`
}

type Productv1ProductInfoPicturesResponse struct {
	Result interface{} `json:"result"`
}

type ActionParameterDiscountTypeEnum string

type ErrorHumanTexts struct {
	Message string `json:"message"`
	Params []interface{} `json:"params"`
	ShortDescription string `json:"short_description"`
	AttributeName string `json:"attribute_name"`
	Description string `json:"description"`
	HintCode string `json:"hint_code"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexData struct {
	MinPrice interface{} `json:"min_price"`
	PriceIndex float64 `json:"price_index"`
	URL string `json:"url"`
}

type WarehouseDeliveryMethodListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Filter interface{} `json:"filter"`
}

type V1ProductUpdateAttributesRequestAttribute struct {
	ComplexID int64 `json:"complex_id"`
	ID int64 `json:"id"`
	Values interface{} `json:"values"`
}

type V1GenerateBarcodeResult struct {
	Code string `json:"code"`
	Error string `json:"error"`
	Barcode string `json:"barcode"`
	ProductID int64 `json:"product_id"`
}

type V1GenerateBarcodeRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type CreateReportResponse struct {
	Result interface{} `json:"result"`
}

type V5FbsPostingProductExemplarStatusV5Request struct {
	PostingNumber string `json:"posting_number"`
}

type V1SellerActionsProductsListResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V1ProductUpdateDiscountResponse struct {
	Result bool `json:"result"`
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpOrderListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type PostingV1PostingFbpListRequestSortDirEnum string

type ListDropOffPointsForCreateFBSWarehouseResponseDropOffPoint struct {
	Type interface{} `json:"type_"`
	Address string `json:"address"`
	Coordinates interface{} `json:"coordinates"`
	DiscountPercent float64 `json:"discount_percent"`
	ID string `json:"id"`
	LastTransitTimeLocal interface{} `json:"last_transit_time_local"`
}

type V1GetStrategyIDsByItemIDsResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1ReviewInfoRequest struct {
	ReviewID string `json:"review_id"`
}

type PostingV4PostingFbsUnfulfilledListResponse struct {
	Count int64 `json:"count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplar struct {
	GTD string `json:"gtd"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
}

type FinanceV1GetFinanceAccrualPostingsResponse struct {
	PostingAccruals []interface{} `json:"posting_accruals"`
}

type ReportReport struct {
	Error string `json:"error"`
	ExpiresAt string `json:"expires_at"`
	File string `json:"file"`
	Params map[string]interface{} `json:"params"`
	ReportType string `json:"report_type"`
	Status string `json:"status"`
	Code string `json:"code"`
	CreatedAt string `json:"created_at"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result interface{} `json:"result"`
}

type ProductImportProductsPricesResponseProcessResult struct {
	ProductID int64 `json:"product_id"`
	Updated bool `json:"updated"`
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1fbpTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type GetProductInfoListResponseStocks struct {
	HasStock bool `json:"has_stock"`
	Stocks []interface{} `json:"stocks"`
}

type Financev3FinanceTransactionTotalsV3ResponseResult struct {
	CompensationAmount float64 `json:"compensation_amount"`
	MoneyTransfer float64 `json:"money_transfer"`
	OthersAmount float64 `json:"others_amount"`
	ProcessingAndDelivery float64 `json:"processing_and_delivery"`
	RefundsAndCancellations float64 `json:"refunds_and_cancellations"`
	SaleCommission float64 `json:"sale_commission"`
	ServicesAmount float64 `json:"services_amount"`
	AccrualsForSale float64 `json:"accruals_for_sale"`
}

type V2FbsPosting struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	InProcessAt string `json:"in_process_at"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
	Barcodes interface{} `json:"barcodes"`
	CreatedAt string `json:"created_at"`
	OrderID int64 `json:"order_id"`
	OrderNumber string `json:"order_number"`
	ShipmentDate string `json:"shipment_date"`
}

type V1FbpDraftDirectTplDlvCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type V1ProductActionTimerStatusResponseStatuses struct {
	ExpiredAt string `json:"expired_at"`
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	ProductID int64 `json:"product_id"`
}

type V1CommentDeleteRequest struct {
	CommentID string `json:"comment_id"`
}

type Postingv3GetFbsPostingListRequest struct {
	Dir string `json:"dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
}

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
	HasNext bool `json:"has_next"`
}

type V1FbpDraftDirectRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type V1SellerActionsCreateDiscountWithConditionRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
}

type SearchQueriesTextRequestSortDir string

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type SellerApiProductV1ResponseDeactivate struct {
	Result interface{} `json:"result"`
}

type ReportReportinfo struct {
	Params map[string]interface{} `json:"params"`
	ReportType string `json:"report_type"`
	Status string `json:"status"`
	Code string `json:"code"`
	CreatedAt string `json:"created_at"`
	Error string `json:"error"`
	ExpiresAt string `json:"expires_at"`
	File string `json:"file"`
}

type V1GetProductRatingBySkuResponse struct {
	Products interface{} `json:"products"`
}

type ValidationErrorCharacteristicEnum string

type AnalyticsProductQueriesRequestSortBy string

type V1SearchQueriesTopResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type GetProductAttributesResponseImage360 struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
}

type V1WarehouseInvalidProductsGetRequest struct {
	LastID int64 `json:"last_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ReviewV2ReviewListV2RequestFilters struct {
	Skus []interface{} `json:"skus"`
	Status interface{} `json:"status"`
	OrderStatus interface{} `json:"order_status"`
	PublishedFrom string `json:"published_from"`
	PublishedTo string `json:"published_to"`
}

type V1OrderValidationError struct {
	OrderErrors []interface{} `json:"order_errors"`
}

type FbsPostingShipV4RequestWith struct {
	AdditionalData bool `json:"additional_data"`
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1GetCompetitorsRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type ItemCommissionsv5 struct {
	FBSDelivToCustomerAmount float64 `json:"fbs_deliv_to_customer_amount"`
	FBSDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
	FBSFirstMileMinAmount float64 `json:"fbs_first_mile_min_amount"`
	SalesPercentFBP float64 `json:"sales_percent_fbp"`
	SalesPercentFBS float64 `json:"sales_percent_fbs"`
	SalesPercentRFBS float64 `json:"sales_percent_rfbs"`
	FBODelivToCustomerAmount float64 `json:"fbo_deliv_to_customer_amount"`
	FBODirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
	FBOReturnFlowAmount float64 `json:"fbo_return_flow_amount"`
	FBSDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
	FBSFirstMileMaxAmount float64 `json:"fbs_first_mile_max_amount"`
	FBSReturnFlowAmount float64 `json:"fbs_return_flow_amount"`
	SalesPercentFBO float64 `json:"sales_percent_fbo"`
	FBODirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
}

type FbpWarehouseListResponseWarehouse struct {
	IsBonded bool `json:"is_bonded"`
	Name string `json:"name"`
	PartnerName string `json:"partner_name"`
	SupplyTypes []interface{} `json:"supply_types"`
	TimezoneName string `json:"timezone_name"`
	AddressDetailing interface{} `json:"address_detailing"`
	ID int64 `json:"id"`
}

type V1OrderValidationErrorErrorType string

type Fbpv1Timeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1GetReturnsListResponse struct {
	Returns []interface{} `json:"returns"`
	HasNext bool `json:"has_next"`
}

type V1GetAttributesResponseAttribute struct {
	GroupID int64 `json:"group_id"`
	GroupName string `json:"group_name"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type_"`
	AttributeComplexID int64 `json:"attribute_complex_id"`
	Description string `json:"description"`
	IsAspect bool `json:"is_aspect"`
	IsCollection bool `json:"is_collection"`
	IsRequired bool `json:"is_required"`
	MaxValueCount int64 `json:"max_value_count"`
	ComplexIsCollection bool `json:"complex_is_collection"`
	CategoryDependent bool `json:"category_dependent"`
	DictionaryID int64 `json:"dictionary_id"`
}

type V2FbsPostingProductCountrySetResponse struct {
	ProductID int64 `json:"product_id"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
}

type V1GetFinanceBalanceV1Response struct {
	Cashflows interface{} `json:"cashflows"`
	Total interface{} `json:"total"`
}

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type ParameterDiscountLevels struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type V3ChatMessage struct {
	CreatedAt string `json:"created_at"`
	Data []interface{} `json:"data"`
	IsImage bool `json:"is_image"`
	IsRead bool `json:"is_read"`
	MessageId int64 `json:"messageId"`
	ModerateImageStatus interface{} `json:"moderate_image_status"`
	User interface{} `json:"user"`
	Context interface{} `json:"context"`
}

type SellerServiceanalyticsDimension string

type V2GetRealizationReportResponseV2 struct {
	Result interface{} `json:"result"`
}

type GetFinanceBalanceV1ResponseTotal struct {
	Payments []interface{} `json:"payments"`
	Accrued interface{} `json:"accrued"`
	ClosingBalance interface{} `json:"closing_balance"`
	OpeningBalance interface{} `json:"opening_balance"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccrualsAccrual struct {
	AccrualDate string `json:"accrual_date"`
	Accrued interface{} `json:"accrued"`
	Quantity int32 `json:"quantity"`
	SellerPrice interface{} `json:"seller_price"`
	SKU int64 `json:"sku"`
	TypeID int32 `json:"type_id"`
}

type V1UpdateWarehouseFBSFirstMileRequest struct {
	TimeslotID int64 `json:"timeslot_id"`
	WarehouseID int64 `json:"warehouse_id"`
	CutInTime int64 `json:"cut_in_time"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	FirstMileType interface{} `json:"first_mile_type"`
}

type WarehouseListV2ResponseWarehouse struct {
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsComfort bool `json:"is_comfort"`
	IsKgt bool `json:"is_kgt"`
	Status string `json:"status"`
	WarehouseType string `json:"warehouse_type"`
	CarriageLabelType interface{} `json:"carriage_label_type"`
	CreatedAt string `json:"created_at"`
	CutInTime int64 `json:"cut_in_time"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	Name string `json:"name"`
	SlaCutIn int64 `json:"sla_cut_in"`
	WorkingDays []interface{} `json:"working_days"`
	IsRFBS bool `json:"is_rfbs"`
	PostingsLimit int32 `json:"postings_limit"`
	Timetable interface{} `json:"timetable"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
	WithItemList bool `json:"with_item_list"`
	AddressInfo interface{} `json:"address_info"`
	CourierPhones []interface{} `json:"courier_phones"`
	FirstMile interface{} `json:"first_mile"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	IsExpress bool `json:"is_express"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
	Phone string `json:"phone"`
	CourierComment string `json:"courier_comment"`
}

type ChatMessageModerateImageStatus string

type Postingv1GetCarriageAvailableListRequest struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type GetConditionalCancellationListV2RequestWith struct {
	Counter bool `json:"counter"`
}

type GetUploadQuotaResponseDailyCreate struct {
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
	Usage int64 `json:"usage"`
}

type ActionsV1ActionsAutoAddProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	ProductIds []interface{} `json:"product_ids"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep struct {
	MinCharge interface{} `json:"min_charge"`
	TariffCharge interface{} `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"`
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
}

type CarriageCarriageGetResponseCancelAvailability struct {
	IsCancelAvailable bool `json:"is_cancel_available"`
	Reason string `json:"reason"`
}

type V1ProductUpdateAttributesResponse struct {
	TaskID int64 `json:"task_id"`
}

type CancellationStateCancellationError struct {
	ErrorCode interface{} `json:"error_code"`
	Message string `json:"message"`
}

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type PostinglistV3status struct {
	From string `json:"from"`
	To string `json:"to"`
}

type FinanceCashFlowStatementListResponseDetailsOthers struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type SellerSellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V1SetProductStairwayDiscountByQuantityRequest struct {
	Stairways []interface{} `json:"stairways"`
	SuppressWarnings bool `json:"suppress_warnings"`
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber string `json:"posting_number"`
	ProductID int64 `json:"product_id"`
	CountryIsoCode string `json:"country_iso_code"`
}

type SearchQueriesTextRequestSortBy string

type ReportCreateReportResponse struct {
	Result interface{} `json:"result"`
}

type V1CreateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type GetProductInfoListResponseStocksStock struct {
	Present int32 `json:"present"`
	Reserved int32 `json:"reserved"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type V2ProductInfoPicturesResponseError struct {
	Message string `json:"message"`
	URL string `json:"url"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type ImportProductsRequestPdfList struct {
	Index int64 `json:"index"`
	Name string `json:"name"`
	SrcURL string `json:"src_url"`
}

type MoneyMoneyCurrentTariffCharge struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type PostingV4PostingFbsUnfulfilledListRequest struct {
	SortDir interface{} `json:"sort_dir"`
	Translit bool `json:"translit"`
	With interface{} `json:"with"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type SellerApiGetSellerProductV1ResponseResult struct {
	Products []interface{} `json:"products"`
	Total float64 `json:"total"`
	LastID float64 `json:"last_id"`
}

type GetProductInfoListResponseStatuses struct {
	IsCreated bool `json:"is_created"`
	ModerateStatus string `json:"moderate_status"`
	Status string `json:"status"`
	StatusName string `json:"status_name"`
	StatusUpdatedAt string `json:"status_updated_at"`
	ValidationStatus string `json:"validation_status"`
	StatusDescription string `json:"status_description"`
	StatusFailed string `json:"status_failed"`
	StatusTooltip string `json:"status_tooltip"`
}

type V1GetProductInfoDiscountedResponseItem struct {
	Condition string `json:"condition"`
	ConditionEstimation string `json:"condition_estimation"`
	Defects string `json:"defects"`
	DiscountedSKU int64 `json:"discounted_sku"`
	PackageDamage string `json:"package_damage"`
	Repair string `json:"repair"`
	WarrantyType string `json:"warranty_type"`
	CommentReasonDamaged string `json:"comment_reason_damaged"`
	MechanicalDamage string `json:"mechanical_damage"`
	PackagingViolation string `json:"packaging_violation"`
	ReasonDamaged string `json:"reason_damaged"`
	Shortage string `json:"shortage"`
	SKU int64 `json:"sku"`
}

type ReturnsRfbsGetV2ResponseReturnReason struct {
	Name string `json:"name"`
	ID int32 `json:"id"`
	IsDefect bool `json:"is_defect"`
}

type DirectDetailsBySellerDetails struct {
	DriverName string `json:"driver_name"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
	VehicleType string `json:"vehicle_type"`
}

type Postingv3GetFbsPostingListRequestFilter struct {
	FbpFilter string `json:"fbpFilter"`
	ProviderID []interface{} `json:"provider_id"`
	Since string `json:"since"`
	To string `json:"to"`
	WarehouseID []interface{} `json:"warehouse_id"`
	OrderID int64 `json:"order_id"`
	Status string `json:"status"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
}

type V1GetCompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type PriceIndexesColorIndex string

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItem struct {
	BreakHours interface{} `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"`
	OpeningHours interface{} `json:"opening_hours"`
}

type ArrivalPassListRequestFilter struct {
	DropoffPointIds []interface{} `json:"dropoff_point_ids"`
	OnlyActivePasses bool `json:"only_active_passes"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	ArrivalReason string `json:"arrival_reason"`
}

type ReviewV2ReviewListV2ResponseReview struct {
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
	ID string `json:"id"`
	OrderStatus interface{} `json:"order_status"`
	PhotosAmount int32 `json:"photos_amount"`
	SKU int64 `json:"sku"`
	VideosAmount int32 `json:"videos_amount"`
	CommentsAmount int32 `json:"comments_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
}

type V1GetFinanceBalanceV1ResponseFeeMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1FbpDraftDropOffRegistrateResponseRegistrationError struct {
	OrderError interface{} `json:"order_error"`
	BundleErrors []interface{} `json:"bundle_errors"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type RpcStatusV1PolygonBind struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1QuestionInfoResponse struct {
	AnswersCount int64 `json:"answers_count"`
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
	PublishedAt interface{} `json:"published_at"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
	ProductURL string `json:"product_url"`
	QuestionLink string `json:"question_link"`
	SKU int64 `json:"sku"`
}

type MoneyMoneyTotalAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1UpdatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyID string `json:"strategy_id"`
	StrategyName string `json:"strategy_name"`
}

type OperationItem struct {
	Name string `json:"name"`
	SKU int64 `json:"sku"`
}

type V1FbpOrderPickUpDlvEditRequestPickUpDetails struct {
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type V1ProductActionTimerStatusResponse struct {
	Statuses interface{} `json:"statuses"`
}

type FinanceCashFlowStatementListResponseCashFlow struct {
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"`
	CurrencyCode string `json:"currency_code"`
	Period interface{} `json:"period"`
	OrdersAmount float64 `json:"orders_amount"`
	ReturnsAmount float64 `json:"returns_amount"`
	CommissionAmount float64 `json:"commission_amount"`
	ServicesAmount float64 `json:"services_amount"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsProducts struct {
	Quantity int32 `json:"quantity"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	Name string `json:"name"`
	Price interface{} `json:"price"`
	ProductColor string `json:"product_color"`
	SKU int64 `json:"sku"`
	Weight float64 `json:"weight"`
	Imei []interface{} `json:"imei"`
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	OfferID string `json:"offer_id"`
}

type ChatRead struct {
	FromMessageID int64 `json:"from_message_id"`
	ChatID string `json:"chat_id"`
}

type GetReturnsListResponseExemplar struct {
	ID int64 `json:"id"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2 struct {
	Products []interface{} `json:"products"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
}

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type V1TimeRangeVisualStatus struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type V1CommentCreateRequest struct {
	MarkReviewAsProcessed bool `json:"mark_review_as_processed"`
	ParentCommentID string `json:"parent_comment_id"`
	ReviewID string `json:"review_id"`
	Text string `json:"text"`
}

type SellerReturnsv1MoneyCommission struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

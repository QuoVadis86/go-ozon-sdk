package types

type V1AssemblyFbsProductListRequestSortDirEnum string

type AnalyticsProductQueriesDetailsResponseQuery struct {
	QueryIndex int64 `json:"query_index"`
	UniqueViewUsers int64 `json:"unique_view_users"`
	Position float64 `json:"position"`
	Query string `json:"query"`
	SKU int64 `json:"sku"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
	ViewConversion float64 `json:"view_conversion"`
	Currency string `json:"currency"`
	Gmv float64 `json:"gmv"`
	OrderCount int64 `json:"order_count"`
}

type V1ApproveDiscountTasksRequestTask struct {
	ID int64 `json:"id"`
	ApprovedPrice float64 `json:"approved_price"`
	SellerComment string `json:"seller_comment"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
}

type V1BundleItemShipmentType string

type RelatedPostingCancelReasons struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	TypeID string `json:"type_id"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplarMark struct {
	Errors []interface{} `json:"errors"`
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
	Valid bool `json:"valid"`
}

type PostingPostingFBSActGetContainerLabelsResponse struct {
	FileContent string `json:"file_content"`
	FileName string `json:"file_name"`
	ContentType string `json:"content_type"`
}

type V4GetProductInfoStocksRequestFilter struct {
	OfferID []interface{} `json:"offer_id"`
	ProductID []interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
	WithQuant interface{} `json:"with_quant"`
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type CancellationErrorCode string

type MoneyMoneyNextTariffCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1ProductActionTimerUpdateRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type ErrorHumanTexts struct {
	ShortDescription string `json:"short_description"`
	AttributeName string `json:"attribute_name"`
	Description string `json:"description"`
	HintCode string `json:"hint_code"`
	Message string `json:"message"`
	Params []interface{} `json:"params"`
}

type V1OrderDraftValidationErrorErrorType string

type V3ImportProductsResponseResult struct {
	TaskID int64 `json:"task_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPosting struct {
	DeliverySchema string `json:"delivery_schema"`
	DeliverySpeed int32 `json:"delivery_speed"`
	Products []interface{} `json:"products"`
}

type V1GetStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplar struct {
	ExemplarID int64 `json:"exemplar_id"`
	GTDErrorCodes []interface{} `json:"gtd_error_codes"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	RnptCheckStatus string `json:"rnpt_check_status"`
	RnptErrorCodes []interface{} `json:"rnpt_error_codes"`
	GTD string `json:"gtd"`
	GTDCheckStatus string `json:"gtd_check_status"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
}

type V2ProductInfoPicturesRequest struct {
	ProductID interface{} `json:"product_id"`
}

type V2PostingFBSActGetProducts struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type V2CancellationStateEnum string

type V1RolesByTokenResponse struct {
	ExpiresAt string `json:"expires_at"`
	Roles []interface{} `json:"roles"`
}

type MoneyMoneyTotalAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type CancellationStateCancellationError struct {
	ErrorCode interface{} `json:"error_code"`
	Message string `json:"message"`
}

type SellerInfoResponseRating struct {
	ValueType interface{} `json:"value_type"`
	CurrentValue interface{} `json:"current_value"`
	Name string `json:"name"`
	PastValue interface{} `json:"past_value"`
	Rating string `json:"rating"`
	Status interface{} `json:"status"`
}

type ArrivalpassArrivalPassListResponseArrivalPass struct {
	DropoffPointID int64 `json:"dropoff_point_id"`
	IsActive bool `json:"is_active"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	ArrivalPassID int64 `json:"arrival_pass_id"`
	ArrivalReasons []interface{} `json:"arrival_reasons"`
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	VehicleModel string `json:"vehicle_model"`
	WarehouseID int64 `json:"warehouse_id"`
}

type WarehouseCarriageLabelTypeEnum string

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementEnum string

type V1GetSupplyOrderBundleRequest struct {
	Query string `json:"query"`
	SortField interface{} `json:"sort_field"`
	BundleIds []interface{} `json:"bundle_ids"`
	IsAsc bool `json:"is_asc"`
	ItemTagsCalculation interface{} `json:"item_tags_calculation"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
}

type FinanceCashFlowStatementListResponseCashFlow struct {
	Period interface{} `json:"period"`
	OrdersAmount float64 `json:"orders_amount"`
	ReturnsAmount float64 `json:"returns_amount"`
	CommissionAmount float64 `json:"commission_amount"`
	ServicesAmount float64 `json:"services_amount"`
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"`
	CurrencyCode string `json:"currency_code"`
}

type ReportReportinfo struct {
	Error string `json:"error"`
	ExpiresAt string `json:"expires_at"`
	File string `json:"file"`
	Params map[string]interface{} `json:"params"`
	ReportType string `json:"report_type"`
	Status string `json:"status"`
	Code string `json:"code"`
	CreatedAt string `json:"created_at"`
}

type V1ItemError struct {
	Message string `json:"message"`
	State string `json:"state"`
	Level string `json:"level"`
	Description string `json:"description"`
	Field string `json:"field"`
	AttributeID int64 `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
	Code string `json:"code"`
}

type CarriageCarriageGetResponse struct {
	PartialNum int64 `json:"partial_num"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	CreatedAt string `json:"created_at"`
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`
	IntegrationType string `json:"integration_type"`
	WarehouseID int64 `json:"warehouse_id"`
	ContainersCount int32 `json:"containers_count"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	Status string `json:"status"`
	UpdatedAt string `json:"updated_at"`
	ActType string `json:"act_type"`
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	AvailableActions []interface{} `json:"available_actions"`
	CancelAvailability interface{} `json:"cancel_availability"`
	FirstMileType string `json:"first_mile_type"`
	IsPartial bool `json:"is_partial"`
	RetryCount int32 `json:"retry_count"`
	CarriageID int64 `json:"carriage_id"`
	CompanyID int64 `json:"company_id"`
	DepartureDate string `json:"departure_date"`
	IsContainerLabelPrinted bool `json:"is_container_label_printed"`
}

type V1DiscountTaskStatus string

type DeliveryDetailsSupplyType string

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type CreateWarehouseFBSRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type StatusEnum string

type V1StrategyRequest struct {
	StrategyID string `json:"strategy_id"`
}

type FbsPostingBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements struct {
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
	ProductsRequiringWeight []interface{} `json:"products_requiring_weight"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseExtremelyLowPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type ItemSize struct {
	HeightMm int32 `json:"height_mm"`
	LengthMm int32 `json:"length_mm"`
	WidthMm int32 `json:"width_mm"`
}

type GetStrategyListResponseStrategy struct {
	UpdatedAt string `json:"updated_at"`
	ProductsCount int64 `json:"products_count"`
	CompetitorsCount int64 `json:"competitors_count"`
	Enabled bool `json:"enabled"`
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type_"`
	UpdateType string `json:"update_type"`
}

type AnalyticsProductQueriesRequestSortDir string

type V1GetFinanceBalanceV1ResponseRevenueMoney struct {
	Value float64 `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type V1FbpDraftDirectCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type GetProductAttributesV4ResponseAttribute struct {
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
	ID int64 `json:"id"`
}

type ErrorData struct {
	Step int64 `json:"step"`
	Value string `json:"value"`
	Code string `json:"code"`
	Field string `json:"field"`
	Message string `json:"message"`
}

type ListDropOffPointsForCreateFBSWarehouseResponseDropOffPoint struct {
	DiscountPercent float64 `json:"discount_percent"`
	ID string `json:"id"`
	LastTransitTimeLocal interface{} `json:"last_transit_time_local"`
	Type interface{} `json:"type_"`
	Address string `json:"address"`
	Coordinates interface{} `json:"coordinates"`
}

type ReviewInfoResponseVideo struct {
	ShortVideoPreviewURL string `json:"short_video_preview_url"`
	URL string `json:"url"`
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	PreviewURL string `json:"preview_url"`
}

type RatingValuePast struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Formatted string `json:"formatted"`
	Status interface{} `json:"status"`
	Value float64 `json:"value"`
}

type V1RatingStatus struct {
	Premium bool `json:"premium"`
	Warning bool `json:"warning"`
	Danger bool `json:"danger"`
}

type V1FbpOrderListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []interface{} `json:"result"`
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

type PostingV4PostingFbsUnfulfilledListRequestWith struct {
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
}

type PostingCancelReasonListResponse struct {
	Result []interface{} `json:"result"`
}

type ItemBundleSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemCount int64 `json:"total_item_count"`
	TotalQuantity int64 `json:"total_quantity"`
}

type GetProductInfoListResponseModelInfo struct {
	Count int64 `json:"count"`
	ModelID int64 `json:"model_id"`
}

type GetSupplyOrderBundleRequestItemTagsCalculation struct {
	StorageWarehouseIds []interface{} `json:"storage_warehouse_ids"`
	DropoffWarehouseID string `json:"dropoff_warehouse_id"`
}

type PickedSegmentSegmentTypeEnum string

type V1OrderAttentionTypeEnum string

type V1DeliveryDetailsPickUpDetails struct {
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type PostingV1PostingFbpListResponsePostings struct {
	FinancialData interface{} `json:"financial_data"`
	InProcessAt string `json:"in_process_at"`
	OrderNumber string `json:"order_number"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	OrderDate string `json:"order_date"`
	OrderID int64 `json:"order_id"`
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
}

type ReviewV2ReviewListV2RequestFilters struct {
	OrderStatus interface{} `json:"order_status"`
	PublishedFrom string `json:"published_from"`
	PublishedTo string `json:"published_to"`
	Skus []interface{} `json:"skus"`
	Status interface{} `json:"status"`
}

type FbsPostingShipV4RequestWith struct {
	AdditionalData bool `json:"additional_data"`
}

type GetReturnsListResponseCompensation struct {
	Status interface{} `json:"status"`
	ChangeMoment string `json:"change_moment"`
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
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

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem struct {
	BreakHours interface{} `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"`
	OpeningHours interface{} `json:"opening_hours"`
}

type QuestionV1GetQuestionListRequestSortDirEnum string

type ReportReportInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequest struct {
	AddressCoordinates interface{} `json:"address_coordinates"`
	IsKgt bool `json:"is_kgt"`
}

type V1ArchiveSkuSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemsCount int64 `json:"total_items_count"`
	TotalQuantity int64 `json:"total_quantity"`
}

type PostingV4PostingFbsListRequest struct {
	Translit bool `json:"translit"`
	With interface{} `json:"with"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type V1ReturnsRfbsActionSetRequest struct {
	RejectionReasonID int32 `json:"rejection_reason_id"`
	ReturnForBackWay float64 `json:"return_for_back_way"`
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
	CompensationAmount float64 `json:"compensation_amount"`
	ID int32 `json:"id"`
}

type WarehouseWarehouseListResponse struct {
	Result []interface{} `json:"result"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityListEnum string

type FinanceV1GetFinanceAccrualByDayResponse struct {
	Accruals []interface{} `json:"accruals"`
	LastID string `json:"last_id"`
}

type V1PostingFbsSplitRequest struct {
	PostingNumber string `json:"posting_number"`
	Postings []interface{} `json:"postings"`
}

type V1GetFinanceBalanceV1ResponseFeeMoney struct {
	Value float64 `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type GetFinanceBalanceV1ResponseService struct {
	Name string `json:"name"`
	Amount interface{} `json:"amount"`
}

type V2PostingFBSActGetPostingsRequest struct {
	ID interface{} `json:"id"`
}

type V1ProductInfoWrongVolumeRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
}

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexDataSelf struct {
	PriceIndex float64 `json:"price_index"`
	URL string `json:"url"`
	MinPrice interface{} `json:"min_price"`
}

type ProductImportProductsBySKURequest struct {
	Items []interface{} `json:"items"`
}

type V1ProductInfoWarehouseStocksRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	WarehouseID int64 `json:"warehouse_id"`
}

type DetailsService struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type V1FbpOrderDropOffTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1UnarchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type Productv3GetProductListResponseItemQuant struct {
	QuantCode string `json:"quant_code"`
	QuantSize int64 `json:"quant_size"`
}

type V1UpdateWarehouseFBSRequestOptions struct {
	CourierPhones []interface{} `json:"courier_phones"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	Comment string `json:"comment"`
}

type ActionParameterDiscountTypeEnum string

type V3FbsPostingProduct struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	CurrencyCode string `json:"currency_code"`
	Imei []interface{} `json:"imei"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
}

type V1Money struct {
	Currency string `json:"currency"`
	Value float64 `json:"value"`
}

type PostingBooleanResponse struct {
	Result bool `json:"result"`
}

type V1FbpDraftDirectTimeslotEditResponseReserveFailureType string

type V3ChatListRequestFilter struct {
	UnreadOnly bool `json:"unread_only"`
	ChatStatus string `json:"chat_status"`
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type AnalyticsProductQueriesRequestSortBy string

type V1FbpDraftDirectRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type PostingCancelReason struct {
	ID int64 `json:"id"`
	IsAvailableForCancellation bool `json:"is_available_for_cancellation"`
	Title string `json:"title"`
	TypeID string `json:"type_id"`
}

type ProductV1ProductVisibilitySetResponseItemsSelectPermissionEnum string

type V4Visibility string

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type GetFBSRatingIndexInfoV1ResponseIndexDynamics struct {
	IndexByDate float64 `json:"index_by_date"`
	ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"`
	Date string `json:"date"`
}

type FinanceV1GetFinanceAccrualByDayRequest struct {
	Date string `json:"date"`
	LastID string `json:"last_id"`
}

type ProductImportProductsPricesRequestPrice struct {
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	Price string `json:"price"`
	ProductID int64 `json:"product_id"`
	AutoActionEnabled string `json:"auto_action_enabled"`
	CurrencyCode string `json:"currency_code"`
	NetPrice string `json:"net_price"`
	OfferID string `json:"offer_id"`
	OldPrice string `json:"old_price"`
	PriceStrategyEnabled string `json:"price_strategy_enabled"`
	Vat string `json:"vat"`
	ManageElasticBoostingThroughPrice bool `json:"manage_elastic_boosting_through_price"`
	MinPrice string `json:"min_price"`
}

type GetConditionalCancellationListV2RequestWith struct {
	Counter bool `json:"counter"`
}

type ActionParameterVoucherParameterTypeEnum string

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPointAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseFailedPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPointDropOffPointTypeEnum string

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"`
}

type FinanceCashFlowStatementListResponseDetailsOthers struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type ReturnsRfbsGetV2ResponseRejectionReason struct {
	IsCommentRequired bool `json:"is_comment_required"`
	Name string `json:"name"`
	Hint string `json:"hint"`
	ID int32 `json:"id"`
}

type V1OrderStatusEnum string

type GetRealizationReportByDayResponseRow struct {
	DeliveryCommission interface{} `json:"delivery_commission"`
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	CommissionRatio float64 `json:"commission_ratio"`
}

type SellerSellerAPIArrivalPassUpdateRequestArrivalPass struct {
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	ID int64 `json:"id"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
}

type Fbpv1Timeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1ProductInfoWarehouseStocksResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Stocks []interface{} `json:"stocks"`
}

type V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason string

type Postingv3GetFbsPostingUnfulfilledListRequestFilter struct {
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	Status string `json:"status"`
	WarehouseID []interface{} `json:"warehouse_id"`
	CutoffFrom string `json:"cutoff_from"`
	DeliveringDateFrom string `json:"delivering_date_from"`
	DeliveringDateTo string `json:"delivering_date_to"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	FbpFilter string `json:"fbpFilter"`
	ProviderID []interface{} `json:"provider_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication struct {
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffMinCharge interface{} `json:"next_tariff_min_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffType string `json:"next_tariff_type"`
	CurrentTariffMinCharge interface{} `json:"current_tariff_min_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	NextTariffCharge interface{} `json:"next_tariff_charge"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	CurrentTariffCharge interface{} `json:"current_tariff_charge"`
}

type V1FbpDraftPickUpRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type DeliveryDetailsDropOffPointDetails struct {
	ProvinceUuid string `json:"province_uuid"`
	Timeslot interface{} `json:"timeslot"`
	ID int64 `json:"id"`
}

type V1FbpCheckConsignmentNoteStateResponse struct {
	ErrorMessage string `json:"error_message"`
	LabelURL string `json:"label_url"`
	State interface{} `json:"state"`
}

type GetRealizationReportByDayResponse struct {
	Rows []interface{} `json:"rows"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsProducts struct {
	Imei []interface{} `json:"imei"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
	ProductColor string `json:"product_color"`
	SKU int64 `json:"sku"`
	Weight float64 `json:"weight"`
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	Quantity int32 `json:"quantity"`
}

type Productv3GetProductAttributesV3Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type V1FbpOrderGetResponse struct {
	AttentionReasons []interface{} `json:"attention_reasons"`
	HasLabel bool `json:"has_label"`
	Status interface{} `json:"status"`
	CancellationState interface{} `json:"cancellation_state"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	OrderNumber string `json:"order_number"`
	RowVersion int64 `json:"row_version"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	CreatedDate string `json:"created_date"`
	DeliveryDetails interface{} `json:"delivery_details"`
	DraftID int64 `json:"draft_id"`
	ID int64 `json:"id"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleUuid string `json:"bundle_uuid"`
	Locked bool `json:"locked"`
	PackageUnitsCount int32 `json:"package_units_count"`
	ReceiveDate string `json:"receive_date"`
	SupplyID string `json:"supply_id"`
}

type V1CarriageCancelResponse struct {
	Error string `json:"error"`
	CarriageStatus string `json:"carriage_status"`
}

type FbpCreateActResponseCreateActErrorReason string

type V1ReturnsCompanyFbsInfoRequestFilter struct {
	PlaceID int64 `json:"place_id"`
}

type V1WarehouseFbsCreateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexData struct {
	MinPrice interface{} `json:"min_price"`
	PriceIndex float64 `json:"price_index"`
	URL string `json:"url"`
}

type ReportReportListResponse struct {
	Result interface{} `json:"result"`
}

type PostingCancelFbsPostingRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	PostingNumber string `json:"posting_number"`
}

type V1GetProductInfoDiscountedResponse struct {
	Items interface{} `json:"items"`
}

type ValidationErrorCharacteristicEnum string

type V2GetDiscountTaskListV2Response struct {
	Tasks []interface{} `json:"tasks"`
}

type V1SellerActionsUpdateInstallmentRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type MoneyMoneyCoinvestment struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type Productv4GetProductAttributesV4Request struct {
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
}

type Productv3GetProductListResponse struct {
	Result interface{} `json:"result"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplarMark struct {
	CheckStatus string `json:"check_status"`
	ErrorCodes []interface{} `json:"error_codes"`
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type ImportProductRequestPromotion struct {
	Operation string `json:"operation"`
	Type string `json:"type_"`
}

type ProductProductInfoPicturesResponsePicture struct {
	Is360 bool `json:"is_360"`
	IsColor bool `json:"is_color"`
	IsPrimary bool `json:"is_primary"`
	ProductID int64 `json:"product_id"`
	State string `json:"state"`
	URL string `json:"url"`
}

type V1QuestionListRequestFilter struct {
	Status string `json:"status"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type SellerOzonLogisticsInfoResponseAvailableSchemasEnum string

type Productv2ProductsStocksRequest struct {
	Stocks []interface{} `json:"stocks"`
}

type V1CreatePricingStrategyResponse struct {
	Result interface{} `json:"result"`
}

type V3FbsPostingExemplarProductV3 struct {
	Exemplars []interface{} `json:"exemplars"`
	SKU int64 `json:"sku"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejectedCodeEnum string

type V1FbpDraftPickUpProductValidateResponseRejectedItem struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
}

type ProductProductArchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	PvzCode int64 `json:"pvz_code"`
	AddressTail string `json:"address_tail"`
	City string `json:"city"`
	Comment string `json:"comment"`
	Country string `json:"country"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	Region string `json:"region"`
	ZipCode string `json:"zip_code"`
	District string `json:"district"`
}

type DeliveryMethodListV2RequestSortDirEnum string

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"`
}

type ParameterPickedSegment struct {
	Segments []interface{} `json:"segments"`
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

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1AssemblyFbsProductListResponse struct {
	ProductsCount int32 `json:"products_count"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V1FbsPostingProductExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"`
}

type GetDiscountTaskListV2ResponseTaskDiscountTaskStatusEnum string

type Fbsv4FbsPostingShipV4Request struct {
	PostingNumber string `json:"posting_number"`
	With interface{} `json:"with"`
	Packages interface{} `json:"packages"`
}

type ArrivalpassArrivalPassUpdateRequestArrivalPass struct {
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	ArrivalPassID int64 `json:"arrival_pass_id"`
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
}

type V1UpdateStatusStrategyRequest struct {
	Enabled bool `json:"enabled"`
	StrategyID string `json:"strategy_id"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndex struct {
	ExternalIndexData interface{} `json:"external_index_data"`
	SelfIndexData interface{} `json:"self_index_data"`
}

type V3ChatHistoryResponse struct {
	HasNext bool `json:"has_next"`
	Messages []interface{} `json:"messages"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type GetRealizationReportResponseV2Row struct {
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission interface{} `json:"delivery_commission"`
}

type V1AssemblyFbsProductListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	SortDir interface{} `json:"sort_dir"`
}

type GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
	Name string `json:"name"`
}

type ListFBSRatingIndexPostingsV1ResponseError struct {
	ChargePercent float64 `json:"charge_percent"`
	ChargePriceCurrencyCode string `json:"charge_price_currency_code"`
	DeliverySchema string `json:"delivery_schema"`
	ErrorAt string `json:"error_at"`
	HasGraceStatus bool `json:"has_grace_status"`
	PostingNumber string `json:"posting_number"`
	ProductPrice float64 `json:"product_price"`
	ProductPriceCurrencyCode string `json:"product_price_currency_code"`
	ChargePrice float64 `json:"charge_price"`
	Index float64 `json:"index"`
	PostingErrorType interface{} `json:"posting_error_type"`
}

type FinanceV1GetFinanceAccrualTypesResponseAccrualType struct {
	Description string `json:"description"`
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type V3ImportProductsResponse struct {
	Result interface{} `json:"result"`
}

type GetConditionalCancellationListV2RequestFilters struct {
	CancellationInitiator []interface{} `json:"cancellation_initiator"`
	PostingNumber []interface{} `json:"posting_number"`
	State interface{} `json:"state"`
}

type V1FbpDraftPickupDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type FinanceTransactionTotalsV3RequestDate struct {
	To string `json:"to"`
	From string `json:"from"`
}

type V3GetFbsPostingListResponseV3 struct {
	Result interface{} `json:"result"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearch struct {
	Address string `json:"address"`
	Types []interface{} `json:"types"`
}

type SetPostingsResponseResult struct {
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
	Result bool `json:"result"`
}

type V1QuestionAnswerListRequest struct {
	SKU int64 `json:"sku"`
	LastID interface{} `json:"last_id"`
	QuestionID string `json:"question_id"`
}

type V2FbsPostingProductCountryListResponseResult struct {
	Name string `json:"name"`
	CountryIsoCode string `json:"country_iso_code"`
}

type V1UpdateWarehouseFBSFirstMileRequestFirstMileTypeEnum string

type V1CreateStockByWarehouseReportRequest struct {
	Language interface{} `json:"language"`
	WarehouseId []interface{} `json:"warehouseId"`
}

type Postingv1GetCarriageAvailableListResponse struct {
	Result interface{} `json:"result"`
}

type CompensationReportLanguage string

type ReportCreateCompanyProductsReportRequestVisibility string

type V1ItemSortField string

type V3FbsPostingDetailOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V1FbpOrderPickUpCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type ReturnsRfbsListResponseReturns struct {
	PostingNumber string `json:"posting_number"`
	Product interface{} `json:"product"`
	ReturnID int64 `json:"return_id"`
	ReturnNumber string `json:"return_number"`
	State interface{} `json:"state"`
	ClientName string `json:"client_name"`
	CreatedAt string `json:"created_at"`
	OrderNumber string `json:"order_number"`
}

type Productv2ProductsStocksRequestStock struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Stock int64 `json:"stock"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductCurrencyEnum string

type V1ReviewCountResponse struct {
	Processed int32 `json:"processed"`
	Total int32 `json:"total"`
	Unprocessed int32 `json:"unprocessed"`
}

type V1AddStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1SellerActionsUpdateInstallmentRequestActionParameters struct {
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type V2FboSinglePostingLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type ProductImportProductsPricesResponseProcessResult struct {
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Updated bool `json:"updated"`
}

type V2GetDiscountTaskListV2RequestDiscountTaskStatusEnum string

type FbsPostingTrackingNumberSetRequestTrackingNumber struct {
	TrackingNumber string `json:"tracking_number"`
	PostingNumber string `json:"posting_number"`
}

type V2WarehouseListV2Response struct {
	Warehouses []interface{} `json:"warehouses"`
	HasNext bool `json:"has_next"`
	Cursor string `json:"cursor"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponseCoordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type Productv2ProductsStocksResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1UpdateWarehouseFBSFirstMileRequest struct {
	CutInTime int64 `json:"cut_in_time"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	FirstMileType interface{} `json:"first_mile_type"`
	TimeslotID int64 `json:"timeslot_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ListFBSRatingIndexPostingsV1RequestFilter struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V3ImportProductsRequestDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type V1SellerActionsProductsCandidatesResponseProduct struct {
	OfferID string `json:"offer_id"`
	SKU []interface{} `json:"sku"`
	DiscountPercent float64 `json:"discount_percent"`
	IsActive bool `json:"is_active"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantSize int64 `json:"quant_size"`
	QuantType interface{} `json:"quant_type"`
	ActionPrice float64 `json:"action_price"`
	BasePrice float64 `json:"base_price"`
	Currency string `json:"currency"`
	Name string `json:"name"`
}

type V2Product struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	CurrencyCode string `json:"currency_code"`
	Price int32 `json:"price"`
	SKU int64 `json:"sku"`
}

type ProductGetProductInfoDescriptionResponse struct {
	Result interface{} `json:"result"`
}

type V1UpdateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type ProductV1ProductVisibilityInfoRequest struct {
	Skus []interface{} `json:"skus"`
}

type V4GetUploadQuotaResponse struct {
	DailyUpdate interface{} `json:"daily_update"`
	OperationLimits interface{} `json:"operation_limits"`
	Total interface{} `json:"total"`
	DailyCreate interface{} `json:"daily_create"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type PriceIndexesIndexDataSelf struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1SellerActionsProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1FbpDraftDirectRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDirectTimeslotEditResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"`
}

type V1GetProductStairwayDiscountByQuantityResponse struct {
	Stairways []interface{} `json:"stairways"`
}

type V1GetAttributeValuesResponseDictionaryValue struct {
	ID int64 `json:"id"`
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
}

type V1SearchAttributeValuesResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftDirectCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
}

type V1DescriptionCategoryTipsRequest struct {
	TypeID []interface{} `json:"type_id"`
}

type ProductGetProductAttributesV3ResponseComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type MoneyMoneyCurrentTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type PostingV4PostingFbsListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
}

type V2FbsPostingProductCountrySetResponse struct {
	ProductID int64 `json:"product_id"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
}

type FbsPostingMoveStatusResponseMoveStatus struct {
	Result bool `json:"result"`
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
}

type PostingV4PostingFbsListResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItem struct {
	BreakHours interface{} `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"`
	OpeningHours interface{} `json:"opening_hours"`
}

type ActionsV1ActionsAutoAddProductsListResponseProducts struct {
	Name string `json:"name"`
	SKU int64 `json:"sku"`
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	AddMode string `json:"add_mode"`
	Currency string `json:"currency"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MinActionQuantity int64 `json:"min_action_quantity"`
	MinSellerPrice float64 `json:"min_seller_price"`
}

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type ChatHistoryRequestFilter struct {
	MessageIds []interface{} `json:"message_ids"`
}

type V3FbsPostingAnalyticsData struct {
	DeliveryDateEnd string `json:"delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
	City string `json:"city"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region string `json:"region"`
	TPLProvider string `json:"tpl_provider"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponseProducts struct {
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	Currency string `json:"currency"`
	MinActionQuantity int64 `json:"min_action_quantity"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	BasePrice float64 `json:"base_price"`
}

type ReturnsRfbsGetResponseReturns struct {
	ClientReturnMethodType interface{} `json:"client_return_method_type"`
	RejectionComment string `json:"rejection_comment"`
	ReturnMethodDescription string `json:"return_method_description"`
	ReturnNumber string `json:"return_number"`
	ReturnReason interface{} `json:"return_reason"`
	RuPostTrackingNumber string `json:"ru_post_tracking_number"`
	WarehouseID int64 `json:"warehouse_id"`
	AvailableActions []interface{} `json:"available_actions"`
	ClientName string `json:"client_name"`
	PostingNumber string `json:"posting_number"`
	ClientPhoto []interface{} `json:"client_photo"`
	Product interface{} `json:"product"`
	RejectionReason []interface{} `json:"rejection_reason"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
	OrderNumber string `json:"order_number"`
	State interface{} `json:"state"`
}

type V1SellerActionsCreateDiscountWithConditionRequestDiscountTypeEnum string

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFeeFee struct {
	Accrued interface{} `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type V1SellerActionsListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Search string `json:"search"`
	Status []interface{} `json:"status"`
	ActionIds []interface{} `json:"action_ids"`
	ActionType []interface{} `json:"action_type"`
}

type V1AssemblyCarriagePostingListResponse struct {
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
	CanPrintMassLabel bool `json:"can_print_mass_label"`
}

type PostingV4PostingFbsUnfulfilledListResponse struct {
	Count int64 `json:"count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type V1GetAttributesResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftPickupCreateRequest struct {
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFees struct {
	Fees []interface{} `json:"fees"`
}

type AnalyticsMetric string

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type MoneyMoneySaleAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type ReasonHumanText struct {
	Text string `json:"text"`
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber string `json:"posting_number"`
	ProductID int64 `json:"product_id"`
	CountryIsoCode string `json:"country_iso_code"`
}

type V1ProductUpdateAttributesResponse struct {
	TaskID int64 `json:"task_id"`
}

type V3FbsTariffication struct {
	CurrentTariffType string `json:"current_tariff_type"`
	CurrentTariffCharge string `json:"current_tariff_charge"`
	CurrentTariffChargeCurrencyCode string `json:"current_tariff_charge_currency_code"`
	NextTariffType string `json:"next_tariff_type"`
	NextTariffChargeCurrencyCode string `json:"next_tariff_charge_currency_code"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffCharge string `json:"next_tariff_charge"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
}

type FinanceCashFlowStatementListResponseService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V2FbsPostingProductCountryListResponse struct {
	Result []interface{} `json:"result"`
}

type V1SellerActionsCreateVoucherResponse struct {
	ActionID int64 `json:"action_id"`
}

type V4GetProductInfoStocksResponseItem struct {
	ProductID int64 `json:"product_id"`
	Stocks []interface{} `json:"stocks"`
	OfferID string `json:"offer_id"`
}

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type WarehouseDeliveryMethodListResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V1AssemblyCarriagePostingListResponsePosting struct {
	AssemblyCode string `json:"assembly_code"`
	CanPrintLabel bool `json:"can_print_label"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1SellerActionsProductsAddRequest struct {
	ActionID int64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	Errors []interface{} `json:"errors"`
	HasNext bool `json:"has_next"`
	Cursor string `json:"cursor"`
}

type ChatRead struct {
	ChatID string `json:"chat_id"`
	FromMessageID int64 `json:"from_message_id"`
}

type V3GetProductInfoListResponsePromotionType string

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type Productv3GetProductListResponseResult struct {
	Items interface{} `json:"items"`
	LastID string `json:"last_id"`
	Total int32 `json:"total"`
}

type V1GetFinanceBalanceV1ResponseAccruedMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1GetSupplyOrderBundleResponse struct {
	LastID string `json:"last_id"`
	Items []interface{} `json:"items"`
	TotalCount int32 `json:"total_count"`
	HasNext bool `json:"has_next"`
}

type V1FbpDraftDropOffProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftPickupCreateRequestDeliveryDetails struct {
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
}

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnForBackWay int64 `json:"return_for_back_way"`
	ReturnID int64 `json:"return_id"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type V1AssemblyCarriagePostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result interface{} `json:"result"`
}

type V1GetFinanceBalanceV1ResponsePaymentsMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualAccruedCategoryEnum string

type V1FbpOrderPickUpDlvEditRequestPickUpDetails struct {
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type ReturnsCompanyFbsInfoResponseDropOffPoints struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	PassInfo interface{} `json:"pass_info"`
	ReturnsCount int32 `json:"returns_count"`
	UtcOffset string `json:"utc_offset"`
	Address string `json:"address"`
	PlaceID int64 `json:"place_id"`
	WarehousesIds []interface{} `json:"warehouses_ids"`
	BoxCount int32 `json:"box_count"`
}

type V1DraftStatusEnum string

type V1SellerActionsCreateMultiLevelDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	DiscountType interface{} `json:"discount_type"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDeliveryService struct {
	TypeID int32 `json:"type_id"`
	Accrued interface{} `json:"accrued"`
}

type ApproveDeclineDiscountTasksResponseFailDetail struct {
	TaskID int64 `json:"task_id"`
	ErrorForUser string `json:"error_for_user"`
}

type V1QuestionInfoResponse struct {
	AuthorName string `json:"author_name"`
	ProductURL string `json:"product_url"`
	QuestionLink string `json:"question_link"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
	ID string `json:"id"`
	PublishedAt interface{} `json:"published_at"`
	Status interface{} `json:"status"`
	AnswersCount int64 `json:"answers_count"`
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

type V1FbpDraftDirectCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ReviewV2ReviewListV2Request struct {
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Filters interface{} `json:"filters"`
}

type Productv5GetProductListRequestFilterFilterVisibility string

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	With interface{} `json:"with"`
	Dir string `json:"dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type V1GenerateBarcodeRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1SellerActionsCreateVoucherRequestVoucherParameter struct {
	Type interface{} `json:"type_"`
	CountCodes int64 `json:"count_codes"`
	IsPrivate bool `json:"is_private"`
}

type V1ProductGetRelatedSKUResponseError struct {
	Code string `json:"code"`
	SKU int64 `json:"sku"`
	Message string `json:"message"`
}

type ProductsPostings struct {
	PostingNumber string `json:"posting_number"`
	Quantity int32 `json:"quantity"`
}

type V1ReviewInfoResponse struct {
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	Videos []interface{} `json:"videos"`
	CommentsAmount int32 `json:"comments_amount"`
	LikesAmount int32 `json:"likes_amount"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
	ID string `json:"id"`
	OrderStatus string `json:"order_status"`
	Status string `json:"status"`
	VideosAmount int32 `json:"videos_amount"`
	DislikesAmount int32 `json:"dislikes_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	Photos []interface{} `json:"photos"`
	PhotosAmount int32 `json:"photos_amount"`
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type MoneyMoneyNextTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type ReviewV2ReviewListV2Response struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
}

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type FinanceTransactionListV3ResponseOperation struct {
	OperationID int64 `json:"operation_id"`
	DeliveryCharge float64 `json:"delivery_charge"`
	OperationType string `json:"operation_type"`
	OperationTypeName string `json:"operation_type_name"`
	Posting interface{} `json:"posting"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"`
	SaleCommission float64 `json:"sale_commission"`
	Services []interface{} `json:"services"`
	Type string `json:"type_"`
	AccrualsForSale float64 `json:"accruals_for_sale"`
	Amount float64 `json:"amount"`
	Items []interface{} `json:"items"`
	OperationDate string `json:"operation_date"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type RowItemCommission struct {
	Total float64 `json:"total"`
	Bonus float64 `json:"bonus"`
	Compensation float64 `json:"compensation"`
	PricePerInstance float64 `json:"price_per_instance"`
	Quantity int32 `json:"quantity"`
	Amount float64 `json:"amount"`
	Commission float64 `json:"commission"`
	StandardFee float64 `json:"standard_fee"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Stars float64 `json:"stars"`
}

type CommonCreateReportResponse struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsProductsAddRequestProduct struct {
	Currency interface{} `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	SKU int64 `json:"sku"`
}

type AnalyticsDataRowDimension struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type V1FbpDraftDropOffRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1SetProductStairwayDiscountByQuantityRequest struct {
	Stairways []interface{} `json:"stairways"`
	SuppressWarnings bool `json:"suppress_warnings"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type GetUploadQuotaResponseDailyUpdate struct {
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
	Usage int64 `json:"usage"`
}

type V1ProductGetRelatedSKUResponse struct {
	Items interface{} `json:"items"`
	Errors interface{} `json:"errors"`
}

type ReportMarkedProductsSalesCreateRequestDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1ProductPricesDetailsResponsePrice struct {
	CustomerPrice interface{} `json:"customer_price"`
	DiscountPercent float64 `json:"discount_percent"`
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
	PriceIndexes []interface{} `json:"price_indexes"`
	SKU int64 `json:"sku"`
}

type V1AddStrategyItemsResponseResult struct {
	Errors []interface{} `json:"errors"`
	FailedProductCount int32 `json:"failed_product_count"`
}

type V1GetRealizationReportByDayRequest struct {
	Day int32 `json:"day"`
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type V1AssemblyFbsPostingListResponse struct {
	Cutoff string `json:"cutoff"`
	Postings []interface{} `json:"postings"`
	Cursor string `json:"cursor"`
}

type SellerApiProductV1Response struct {
	Result interface{} `json:"result"`
}

type V1GetRealizationReportPostingResponse struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"`
}

type DetailsServices struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type OperationService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type FbpDraftDropOffPointListResponseDropOffPoint struct {
	City string `json:"city"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	NearestDropOffDate string `json:"nearest_drop_off_date"`
	PointAddress string `json:"point_address"`
	ProvinceUuid string `json:"province_uuid"`
}

type ReportCreateReportResponse struct {
	Result interface{} `json:"result"`
}

type V2FbsPostingProduct struct {
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int64 `json:"quantity"`
}

type RpcStatusV1PolygonBind struct {
	Details []interface{} `json:"details"`
	Message string `json:"message"`
	Code int32 `json:"code"`
}

type ReviewV2ReviewListV2RequestFiltersStatusEnum string

type ProductGetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []interface{} `json:"tracking_numbers"`
}

type Productv2ProductsStocksResponseResult struct {
	ProductID int64 `json:"product_id"`
	Updated bool `json:"updated"`
	WarehouseID int64 `json:"warehouse_id"`
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Percent int64 `json:"percent"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequestToUpdate struct {
	Currency string `json:"currency"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	ActionPrice float64 `json:"action_price"`
}

type V2ReturnsRfbsGetV2ResponseState struct {
	State string `json:"state"`
	StateName string `json:"state_name"`
}

type Productv3GetProductListRequestFilterFilterVisibility string

type Financev3Period struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V3FinanceCashFlowStatementListResponseResult struct {
	CashFlows interface{} `json:"cash_flows"`
	Details []interface{} `json:"details"`
	PageCount int64 `json:"page_count"`
}

type ValidationResultValidationErrorTypeEnum string

type Productv1ProductInfoPicturesResponseResult struct {
	Pictures interface{} `json:"pictures"`
}

type ActionsV1ActionsAutoAddProductsListResponse struct {
	Products []interface{} `json:"products"`
	Total int64 `json:"total"`
}

type PostingProductCancelRequestItem struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type GetFinanceBalanceV1ResponseSalesDetails struct {
	PartnerPrograms interface{} `json:"partner_programs"`
	PointsForDiscounts string `json:"points_for_discounts"`
	Revenue interface{} `json:"revenue"`
}

type GetFinanceBalanceV1ResponseTotal struct {
	Accrued interface{} `json:"accrued"`
	ClosingBalance interface{} `json:"closing_balance"`
	OpeningBalance interface{} `json:"opening_balance"`
	Payments []interface{} `json:"payments"`
}

type V1FbpDraftPickUpRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type V2ReturnsRfbsListV2ResponseState struct {
	StateName string `json:"state_name"`
	GroupState string `json:"group_state"`
	MoneyReturnStateName string `json:"money_return_state_name"`
	State string `json:"state"`
}

type V5FbsPostingProductExemplarValidateV5Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1PostingFbsSplitRequestPosting struct {
	Products []interface{} `json:"products"`
}

type V1GetStrategyItemsResponseResult struct {
	ProductID []interface{} `json:"product_id"`
}

type V1FbpArchiveListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPoint struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Address string `json:"address"`
	AddressCoordinates interface{} `json:"address_coordinates"`
}

type ProductV1ProductVisibilitySetRequestItemPlacementPlacementEnum string

type V2ReturnsRfbsListRequest struct {
	Filter interface{} `json:"filter"`
	LastID int32 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type V3GetProductInfoListRequest struct {
	OfferID []interface{} `json:"offer_id"`
	ProductID []interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
}

type V1AnalyticsProductQueriesResponseItem struct {
	Currency string `json:"currency"`
	Name string `json:"name"`
	Position float64 `json:"position"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
	UniqueViewUsers int64 `json:"unique_view_users"`
	ViewConversion float64 `json:"view_conversion"`
	Gmv float64 `json:"gmv"`
	OfferID string `json:"offer_id"`
	SKU int64 `json:"sku"`
	Category string `json:"category"`
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type V1AddBarcodeResult struct {
	Code string `json:"code"`
	Error string `json:"error"`
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type V1ApproveDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V1SearchQueriesTopResponseSearchQuery struct {
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
	ConversionToCart float64 `json:"conversion_to_cart"`
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
	SellersCount float64 `json:"sellers_count"`
	AddToCart float64 `json:"add_to_cart"`
}

type V1FbpOrderDirectCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1UpdatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyID string `json:"strategy_id"`
	StrategyName string `json:"strategy_name"`
}

type V1FbpWarehouseListResponse struct {
	Warehouses []interface{} `json:"warehouses"`
}

type GetConditionalCancellationListV2ResponseResult struct {
	ApproveComment string `json:"approve_comment"`
	CancellationID int64 `json:"cancellation_id"`
	CancellationReasonMessage string `json:"cancellation_reason_message"`
	CancelledAt string `json:"cancelled_at"`
	PostingNumber string `json:"posting_number"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	ApproveDate string `json:"approve_date"`
	AutoApproveDate string `json:"auto_approve_date"`
	CancellationInitiator interface{} `json:"cancellation_initiator"`
	CancellationReason interface{} `json:"cancellation_reason"`
	OrderDate string `json:"order_date"`
	SourceID int64 `json:"source_id"`
	State interface{} `json:"state"`
}

type V2ProductInfoPicturesResponseError struct {
	Message string `json:"message"`
	URL string `json:"url"`
}

type V1GetStrategyIDsByItemIDsResponse struct {
	Result interface{} `json:"result"`
}

type ValidationResultItemStateEnum string

type V1SellerOzonLogisticsInfoResponse struct {
	AvailableSchemas []interface{} `json:"available_schemas"`
	OzonLogisticsEnabled bool `json:"ozon_logistics_enabled"`
}

type V1fbpDeliveryDetails struct {
	SupplyType interface{} `json:"supply_type"`
	DirectDetails interface{} `json:"direct_details"`
	DropOffPoint interface{} `json:"drop_off_point"`
	PickupDetails interface{} `json:"pickup_details"`
}

type AnalyticsSorting struct {
	Key string `json:"key"`
	Order interface{} `json:"order"`
}

type RowItemCommissionReturn struct {
	Quantity int32 `json:"quantity"`
	StandardFee float64 `json:"standard_fee"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Stars float64 `json:"stars"`
	Total float64 `json:"total"`
	Bonus float64 `json:"bonus"`
	Compensation float64 `json:"compensation"`
	PricePerInstance float64 `json:"price_per_instance"`
	Amount float64 `json:"amount"`
	Commission float64 `json:"commission"`
}

type V1SellerActionsCreateVoucherRequest struct {
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
	VoucherParameters interface{} `json:"voucher_parameters"`
}

type V1ApproveDeclineDiscountTasksResponse struct {
	Result interface{} `json:"result"`
}

type V1CreateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type ValidationResultItem struct {
	Size interface{} `json:"size"`
	SKU int64 `json:"sku"`
	WeightG float64 `json:"weight_g"`
}

type V2PostingFBSActGetPostingsResult struct {
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	Products []interface{} `json:"products"`
	ID int64 `json:"id"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Status string `json:"status"`
	SellerError string `json:"seller_error"`
}

type Productv5GetProductInfoPricesV5Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type V1FbpDraftListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type V1CommentCreateRequest struct {
	ReviewID string `json:"review_id"`
	Text string `json:"text"`
	MarkReviewAsProcessed bool `json:"mark_review_as_processed"`
	ParentCommentID string `json:"parent_comment_id"`
}

type PriceIndexesIndexSelfData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type PostingV1PostingFbpListResponsePostingsFinancialData struct {
	Products []interface{} `json:"products"`
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	DeliveryAmount float64 `json:"delivery_amount"`
}

type V1OrderValidationError struct {
	OrderErrors []interface{} `json:"order_errors"`
}

type GetUploadQuotaResponseDailyCreate struct {
	Usage int64 `json:"usage"`
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
}

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type ProductV1ProductVisibilityInfoResponseItem struct {
	ShowcasesVisibility interface{} `json:"showcases_visibility"`
	SKU int64 `json:"sku"`
}

type ReportLanguage string

type V1DayOfWeek string

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type MoneyMoneySellerPrice struct {
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

type SellerApiProductV1ResponseProduct struct {
	ProductID float64 `json:"product_id"`
	Reason string `json:"reason"`
}

type V1GetAttributeValuesRequest struct {
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Language interface{} `json:"language"`
	LastValueID int64 `json:"last_value_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
}

type V1SellerActionsUpdateVoucherRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1SetPostingsRequest struct {
	CarriageID int64 `json:"carriage_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1FbpGetLabelRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type FinanceCashFlowStatementListResponseDeliveryService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type PostingV1PostingFbpListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	Cursor string `json:"cursor"`
}

type V3PostingFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
}

type V1QuestionInfoRequest struct {
	QuestionID string `json:"question_id"`
}

type V1SearchQueriesTopRequest struct {
	Limit string `json:"limit"`
	Offset string `json:"offset"`
}

type V1FbpDraftPickupDlvEditRequestDeliveryDetails struct {
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type V1DeleteStrategyItemsResponseResult struct {
	FailedProductCount int32 `json:"failed_product_count"`
}

type V1FbpCreateActResponse struct {
	Errors []interface{} `json:"errors"`
	FileUuid string `json:"file_uuid"`
	IsSuccess bool `json:"is_success"`
}

type ChatInfoChatStatusEnum string

type SellerSellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CarriageID int64 `json:"carriage_id"`
}

type V1SearchQueriesTextRequest struct {
	Text string `json:"text"`
	Limit string `json:"limit"`
	Offset string `json:"offset"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
}

type V4FbsPostingShipPackageV4Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1TimeRangeVisualStatus struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type GetFinanceBalanceV1ResponseReturns struct {
	Amount interface{} `json:"amount"`
	AmountDetails interface{} `json:"amount_details"`
	Fee interface{} `json:"fee"`
}

type V1GetDiscountTaskListResponseTask struct {
	SellerComment string `json:"seller_comment"`
	RequestedPrice float64 `json:"requested_price"`
	DiscountPercent float64 `json:"discount_percent"`
	BasePrice float64 `json:"base_price"`
	IsAutoModerated bool `json:"is_auto_moderated"`
	Email string `json:"email"`
	Patronymic string `json:"patronymic"`
	UserComment string `json:"user_comment"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	ID int64 `json:"id"`
	MinAutoPrice float64 `json:"min_auto_price"`
	IsPurchased bool `json:"is_purchased"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	ApprovedPriceFeePercent float64 `json:"approved_price_fee_percent"`
	CreatedAt string `json:"created_at"`
	Status string `json:"status"`
	PrevTaskID int64 `json:"prev_task_id"`
	OfferID string `json:"offer_id"`
	CustomerName string `json:"customer_name"`
	SKU int64 `json:"sku"`
	Discount float64 `json:"discount"`
	FirstName string `json:"first_name"`
	EditedTill string `json:"edited_till"`
	IsDamaged bool `json:"is_damaged"`
	ApprovedDiscountPercent float64 `json:"approved_discount_percent"`
	ApprovedPrice float64 `json:"approved_price"`
	OriginalPrice float64 `json:"original_price"`
	ModeratedAt string `json:"moderated_at"`
	RequestedQuantityMin int64 `json:"requested_quantity_min"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	RequestedPriceWithFee float64 `json:"requested_price_with_fee"`
	EndAt string `json:"end_at"`
	ApprovedDiscount float64 `json:"approved_discount"`
	LastName string `json:"last_name"`
	ApprovedPriceWithFee float64 `json:"approved_price_with_fee"`
}

type RpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type PostingCancelReasonResponse struct {
	Result []interface{} `json:"result"`
}

type ReportListRequestReportType string

type Postingv3GetFbsPostingRequest struct {
	PostingNumber string `json:"posting_number"`
	With interface{} `json:"with"`
}

type FilterOp interface{}

type PostingPostingFBSPackageLabelResponse struct {
	ContentType string `json:"content_type"`
	FileContent string `json:"file_content"`
	FileName string `json:"file_name"`
}

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
	HasNext bool `json:"has_next"`
}

type V1ArchiveWarehouseFBSRequest struct {
	Reason string `json:"reason"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1CommentSort string

type V2GetProductInfoStocksByWarehouseFbsRequestV2 struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	OfferID []interface{} `json:"offer_id"`
	SKU []interface{} `json:"sku"`
}

type WarehouseInvalidProductsGetResponseValidationResult struct {
	ValidationErrors []interface{} `json:"validation_errors"`
	Item interface{} `json:"item"`
	State interface{} `json:"state"`
}

type V2GetConditionalCancellationListV2Response struct {
	Result []interface{} `json:"result"`
	Counter int64 `json:"counter"`
	LastID int64 `json:"last_id"`
}

type SellerInfoResponseSubscriptionTypeEnum string

type LanguageLanguage string

type ProtobufAny struct {
	Value string `json:"value"`
	TypeUrl string `json:"typeUrl"`
}

type WarehouseWorkingDaysEnum string

type V3FinanceCashFlowStatementListRequest struct {
	Date interface{} `json:"date"`
	Page int32 `json:"page"`
	WithDetails bool `json:"with_details"`
	PageSize int32 `json:"page_size"`
}

type DetailsOthers struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type SellerApiGetSellerProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Limit float64 `json:"limit"`
	LastID float64 `json:"last_id"`
}

type DeliveryMethodListV2ResponseDeliveryMethod struct {
	ProviderID int64 `json:"provider_id"`
	TPLDropoffPoint interface{} `json:"tpl_dropoff_point"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
	CreatedAt string `json:"created_at"`
	Cutoff string `json:"cutoff"`
	ID int64 `json:"id"`
	SlaCutIn int64 `json:"sla_cut_in"`
	Status string `json:"status"`
	TemplateID int64 `json:"template_id"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	IsExpress bool `json:"is_express"`
	Name string `json:"name"`
}

type V1UpdateWarehouseFBSRequestWorkingDaysEnum string

type V1FbpDraftDropOffRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrual struct {
	TotalAmount interface{} `json:"total_amount"`
	AccrualID int64 `json:"accrual_id"`
	UnitNumber string `json:"unit_number"`
	AccruedCategory interface{} `json:"accrued_category"`
	Date string `json:"date"`
	ItemFees interface{} `json:"item_fees"`
	NonItemFee interface{} `json:"non_item_fee"`
	Posting interface{} `json:"posting"`
}

type V1AssemblyCarriageProductListResponse struct {
	Cursor string `json:"cursor"`
	Products []interface{} `json:"products"`
}

type ReturnsRfbsGetV2ResponseReturnReason struct {
	ID int32 `json:"id"`
	IsDefect bool `json:"is_defect"`
	Name string `json:"name"`
}

type V1QuestionAnswerListResponseAnswers struct {
	StatusPublication interface{} `json:"status_publication"`
	Text string `json:"text"`
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
	PublishedAt interface{} `json:"published_at"`
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
}

type WarehouseFirstMile struct {
	DropoffPointID string `json:"dropoff_point_id"`
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
	TimeslotFrom string `json:"timeslot_from"`
	TimeslotID int64 `json:"timeslot_id"`
	TimeslotTo string `json:"timeslot_to"`
	Type interface{} `json:"type_"`
}

type V3ImportProductsRequestComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type FinanceV1GetFinanceAccrualTypesResponse struct {
	AccrualTypes []interface{} `json:"accrual_types"`
}

type MoneyMoney struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData struct {
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	IsPremium bool `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	City string `json:"city"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
	Region string `json:"region"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftPickUpProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1FbpDraftDirectProductValidateResponseRejectedItem struct {
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
}

type GetCarriageAvailableListResponseResult struct {
	WarehouseName string `json:"warehouse_name"`
	WarehouseTimezone string `json:"warehouse_timezone"`
	CarriageID int64 `json:"carriage_id"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	FirstMileType string `json:"first_mile_type"`
	RecommendedTimeLocal string `json:"recommended_time_local"`
	WarehouseID int64 `json:"warehouse_id"`
	CutoffAt string `json:"cutoff_at"`
	Errors interface{} `json:"errors"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	WarehouseCity string `json:"warehouse_city"`
	CarriageStatus string `json:"carriage_status"`
	TPLProviderIconURL string `json:"tpl_provider_icon_url"`
	TPLProviderName string `json:"tpl_provider_name"`
	CarriagePostingsCount int32 `json:"carriage_postings_count"`
	DeliveryMethodName string `json:"delivery_method_name"`
	MandatoryPostingsCount int32 `json:"mandatory_postings_count"`
	MandatoryPackagedCount int32 `json:"mandatory_packaged_count"`
	RecommendedTimeUtcOffsetInMinutes float64 `json:"recommended_time_utc_offset_in_minutes"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2 struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type Postingv3GetFbsPostingListRequestFilter struct {
	FbpFilter string `json:"fbpFilter"`
	ProviderID []interface{} `json:"provider_id"`
	Since string `json:"since"`
	Status string `json:"status"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	OrderID int64 `json:"order_id"`
	To string `json:"to"`
	WarehouseID []interface{} `json:"warehouse_id"`
}

type FinanceV1GetFinanceAccrualPostingsRequest struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1FbpGetLabelResponse struct {
	LabelURL string `json:"label_url"`
	State interface{} `json:"state"`
}

type GetReturnsListResponseVisualStatus struct {
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
	SysName string `json:"sys_name"`
}

type V3GetFbsPostingListResponseV3Result struct {
	HasNext bool `json:"has_next"`
	Postings []interface{} `json:"postings"`
}

type FirstMileTypeEnum string

type GetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueId int64 `json:"dictionaryValueId"`
	Value string `json:"value"`
}

type ReviewInfoResponsePhoto struct {
	Height int32 `json:"height"`
	URL string `json:"url"`
	Width int32 `json:"width"`
}

type V4GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"`
	Name string `json:"name"`
}

type MoneyMoneyCommission struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1UpdateWarehouseFBSFirstMileResponse struct {
	OperationID string `json:"operation_id"`
}

type Postingv3FbsPostingWithParamsExamplars struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
	ProductExemplars bool `json:"product_exemplars"`
	RelatedPostings bool `json:"related_postings"`
	Translit bool `json:"translit"`
}

type V1SellerActionsProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	Skus []interface{} `json:"skus"`
}

type V1CancellationStateStatus string

type V1GetTreeRequest struct {
	Language interface{} `json:"language"`
}

type V3ImportProductsRequestAttribute struct {
	ComplexID int64 `json:"complex_id"`
	ID int64 `json:"id"`
	Values []interface{} `json:"values"`
}

type V1GetRealizationReportPostingRequest struct {
	Year int32 `json:"year"`
	Month int32 `json:"month"`
}

type V1FbpDraftDirectProductValidateResponse struct {
	RejectedItems []interface{} `json:"rejected_items"`
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
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

type V3FbsPostingDetailRelatedPostings struct {
	RelatedPostingNumbers interface{} `json:"related_posting_numbers"`
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1QuestionAnswerCreateDefault struct {
	Code int32 `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

type PostingV4PostingFbsListResponsePostingsFinancialData struct {
	Products []interface{} `json:"products"`
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
}

type V1GetDiscountTaskListRequest struct {
	Status interface{} `json:"status"`
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type GetProductInfoListResponseCommission struct {
	DeliveryAmount float64 `json:"delivery_amount"`
	Percent float64 `json:"percent"`
	ReturnAmount float64 `json:"return_amount"`
	SaleSchema string `json:"sale_schema"`
	Value float64 `json:"value"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseBelowMinPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type V1ProductUpdateDiscountRequest struct {
	Discount int32 `json:"discount"`
	ProductID int64 `json:"product_id"`
}

type V1AssemblyCarriagePostingListResponsePostingProduct struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type GetProductInfoListResponseVisibilityDetails struct {
	HasPrice bool `json:"has_price"`
	HasStock bool `json:"has_stock"`
}

type SellerServiceanalyticsDimension string

type V3FbsPostingProductExemplarsV3 struct {
	Products []interface{} `json:"products"`
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1GetProductRatingBySkuResponse struct {
	Products interface{} `json:"products"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequestCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V6FbsPostingProductExemplarCreateOrGetV6Response struct {
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V4GetProductInfoStocksResponse struct {
	Cursor string `json:"cursor"`
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type MoneyPostingMoney struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type V1QuestionAnswerCreateRequest struct {
	Text string `json:"text"`
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
}

type V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason string

type Productv3GetProductAttributesV3ResponseResult struct {
	DimensionUnit string `json:"dimension_unit"`
	ImageGroupID string `json:"image_group_id"`
	OfferID string `json:"offer_id"`
	PDFList []interface{} `json:"pdf_list"`
	Width int32 `json:"width"`
	Attributes []interface{} `json:"attributes"`
	CategoryID int64 `json:"category_id"`
	Height int32 `json:"height"`
	ID int64 `json:"id"`
	Weight int32 `json:"weight"`
	TypeID int64 `json:"type_id"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Images []interface{} `json:"images"`
	Name string `json:"name"`
	WeightUnit string `json:"weight_unit"`
	ColorImage string `json:"color_image"`
	Images360 []interface{} `json:"images360"`
	Barcode string `json:"barcode"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Depth int32 `json:"depth"`
}

type Productv5GetProductInfoPricesV5Response struct {
	Total int32 `json:"total"`
	Cursor string `json:"cursor"`
	Items interface{} `json:"items"`
}

type WarehouseAddressInfo struct {
	Address string `json:"address"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Utc string `json:"utc"`
}

type RpcStatusV1PolygonCreate struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V3FbsPosting struct {
	AvailableActions interface{} `json:"available_actions"`
	IsPresortable bool `json:"is_presortable"`
	OrderID int64 `json:"order_id"`
	OrderNumber string `json:"order_number"`
	LegalInfo interface{} `json:"legal_info"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	Addressee interface{} `json:"addressee"`
	Substatus string `json:"substatus"`
	Cancellation interface{} `json:"cancellation"`
	DestinationPlaceName string `json:"destination_place_name"`
	ParentPostingNumber string `json:"parent_posting_number"`
	Requirements interface{} `json:"requirements"`
	TrackingNumber string `json:"tracking_number"`
	IsExpress bool `json:"is_express"`
	Optional interface{} `json:"optional"`
	ShipmentDate string `json:"shipment_date"`
	AnalyticsData interface{} `json:"analytics_data"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	InProcessAt string `json:"in_process_at"`
	Products []interface{} `json:"products"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	Customer interface{} `json:"customer"`
	DeliveringDate string `json:"delivering_date"`
	DeliveryMethod interface{} `json:"delivery_method"`
	FinancialData interface{} `json:"financial_data"`
	PostingNumber string `json:"posting_number"`
	Tariffication interface{} `json:"tariffication"`
	Barcodes interface{} `json:"barcodes"`
	Status string `json:"status"`
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairwayStep struct {
	Discount int64 `json:"discount"`
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
}

type V1GetProductInfoDiscountedRequest struct {
	DiscountedSkus interface{} `json:"discounted_skus"`
}

type V1SearchAttributeValuesRequest struct {
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	Value string `json:"value"`
}

type SellerReturnsv1MoneyProduct struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type CommentListResponseComment struct {
	ParentCommentID string `json:"parent_comment_id"`
	PublishedAt string `json:"published_at"`
	Text string `json:"text"`
	ID string `json:"id"`
	IsOfficial bool `json:"is_official"`
	IsOwner bool `json:"is_owner"`
}

type RowItemOrder struct {
	PostingNumber string `json:"posting_number"`
	CreatedDate string `json:"created_date"`
}

type ExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V1QuestionChangeStatusRequest struct {
	QuestionIds interface{} `json:"question_ids"`
	Status string `json:"status"`
}

type V1SellerActionsCreateDiscountWithConditionResponse struct {
	ActionID int64 `json:"action_id"`
}

type ProductGetImportProductsInfoRequest struct {
	TaskID int64 `json:"task_id"`
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []interface{} `json:"related_posting_numbers"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponseTimeslot struct {
	ID int64 `json:"id"`
	To string `json:"to"`
	From string `json:"from"`
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"`
}

type Fbsv4PostingProductDetailWithoutDimensions struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	CurrencyCode string `json:"currency_code"`
	MandatoryMark interface{} `json:"mandatory_mark"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
}

type V1GetProductInfoDiscountedResponseItem struct {
	PackagingViolation string `json:"packaging_violation"`
	ReasonDamaged string `json:"reason_damaged"`
	Repair string `json:"repair"`
	SKU int64 `json:"sku"`
	WarrantyType string `json:"warranty_type"`
	CommentReasonDamaged string `json:"comment_reason_damaged"`
	Condition string `json:"condition"`
	Defects string `json:"defects"`
	PackageDamage string `json:"package_damage"`
	Shortage string `json:"shortage"`
	ConditionEstimation string `json:"condition_estimation"`
	DiscountedSKU int64 `json:"discounted_sku"`
	MechanicalDamage string `json:"mechanical_damage"`
}

type V2FbsPostingResponse struct {
	Result interface{} `json:"result"`
}

type V1GetProductInfoSubscriptionResponseResult struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type DetailsRfbsDetails struct {
	Total float64 `json:"total"`
	TransferDelivery float64 `json:"transfer_delivery"`
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"`
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`
	PartialCompensation float64 `json:"partial_compensation"`
	PartialCompensationReturn float64 `json:"partial_compensation_return"`
}

type ParameterDiscountLevels struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type GetRealizationReportResponseV2Result struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
}

type DetailsReturns struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type V1FbpDraftDropOffCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V5FbsPostingProductExemplarValidateV5Response struct {
	Products []interface{} `json:"products"`
}

type V2ChatReadResponse struct {
	UnreadCount int64 `json:"unread_count"`
}

type ActionsV1ActionsAutoAddProductsListRequest struct {
	Offset int64 `json:"offset"`
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
}

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"`
}

type GetReturnsListResponseReturnsItem struct {
	ReturnClearingID int64 `json:"return_clearing_id"`
	OrderID int64 `json:"order_id"`
	OrderNumber string `json:"order_number"`
	Storage interface{} `json:"storage"`
	CompensationStatus interface{} `json:"compensation_status"`
	Exemplars []interface{} `json:"exemplars"`
	ID int64 `json:"id"`
	ReturnReasonName string `json:"return_reason_name"`
	Place interface{} `json:"place"`
	TargetPlace interface{} `json:"target_place"`
	Product interface{} `json:"product"`
	Logistic interface{} `json:"logistic"`
	Visual interface{} `json:"visual"`
	AdditionalInfo interface{} `json:"additional_info"`
	PostingNumber string `json:"posting_number"`
	CompanyID int64 `json:"company_id"`
	Type string `json:"type_"`
	Schema string `json:"schema"`
	SourceID int64 `json:"source_id"`
	ClearingID int64 `json:"clearing_id"`
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"`
}

type DescriptionCategoryTipsResponseResult struct {
	ImagesURL []interface{} `json:"images_url"`
	InfoURL string `json:"info_url"`
	TypeID int64 `json:"type_id"`
}

type Financev3FinanceTransactionListV3ResponseResult struct {
	PageCount int64 `json:"page_count"`
	RowCount int64 `json:"row_count"`
	Operations []interface{} `json:"operations"`
}

type ReturnsRfbsGetV2ResponseAvailableAction struct {
	Name string `json:"name"`
	ID int32 `json:"id"`
}

type V1GetFinanceBalanceV1ResponseServicesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V3PostingProductDetail struct {
	OfferID string `json:"offer_id"`
	CurrencyCode string `json:"currency_code"`
	Quantity int32 `json:"quantity"`
	HasImei bool `json:"has_imei"`
	Dimensions interface{} `json:"dimensions"`
	MandatoryMark []interface{} `json:"mandatory_mark"`
	Price string `json:"price"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
}

type V1FbpOrderPickUpDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type PostingV4PostingFbsListResponsePostingsAnalyticsData struct {
	DeliveryDateEnd string `json:"delivery_date_end"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	Region string `json:"region"`
	TPLProvider string `json:"tpl_provider"`
	Warehouse string `json:"warehouse"`
	City string `json:"city"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
}

type V1AddBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type ArrivalpassArrivalPassCreateRequestArrivalPass struct {
	DriverPhone string `json:"driver_phone"`
	DropoffPointID int64 `json:"dropoff_point_id"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WarehouseID int64 `json:"warehouse_id"`
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
}

type ValidationResultValidationError struct {
	Characteristic interface{} `json:"characteristic"`
	RestrictionPrice interface{} `json:"restriction_price"`
	RestrictionVwc float64 `json:"restriction_vwc"`
	TemplateID int32 `json:"template_id"`
	Type interface{} `json:"type_"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponseDropOffPointTypeEnum string

type RatingValueCurrent struct {
	Formatted string `json:"formatted"`
	Status interface{} `json:"status"`
	Value float64 `json:"value"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type ReviewListResponseReview struct {
	PhotosAmount int32 `json:"photos_amount"`
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
	VideosAmount int32 `json:"videos_amount"`
	CommentsAmount int32 `json:"comments_amount"`
	ID string `json:"id"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	OrderStatus string `json:"order_status"`
	Status string `json:"status"`
}

type MoneyMoneySelf struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type SearchAttributeValuesResponseValue struct {
	ID int64 `json:"id"`
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
}

type V2ReturnsRfbsGetRequest struct {
	ReturnID int64 `json:"return_id"`
}

type V1GetStrategyIDsByItemIDsResponseResult struct {
	ProductsInfo []interface{} `json:"products_info"`
}

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementListEnum string

type GetProductInfoPricesResponseItemPriceIndexes struct {
	ColorIndex string `json:"color_index"`
	ExternalIndexData interface{} `json:"external_index_data"`
	OzonIndexData interface{} `json:"ozon_index_data"`
	SelfMarketplacesIndexData interface{} `json:"self_marketplaces_index_data"`
}

type DirectDetailsTimeslotDetails struct {
	Timeslot interface{} `json:"timeslot"`
	TimeslotReservationID string `json:"timeslot_reservation_id"`
}

type V3ImportProductsRequestItem struct {
	Attributes []interface{} `json:"attributes"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	ColorImage string `json:"color_image"`
	Depth int32 `json:"depth"`
	PDFList []interface{} `json:"pdf_list"`
	Weight int32 `json:"weight"`
	WeightUnit string `json:"weight_unit"`
	CurrencyCode string `json:"currency_code"`
	DimensionUnit string `json:"dimension_unit"`
	GeoNames interface{} `json:"geo_names"`
	Height int32 `json:"height"`
	Images []interface{} `json:"images"`
	PrimaryImage string `json:"primary_image"`
	ServiceType interface{} `json:"service_type"`
	Name string `json:"name"`
	Barcode string `json:"barcode"`
	Promotions []interface{} `json:"promotions"`
	Images360 []interface{} `json:"images360"`
	OfferID string `json:"offer_id"`
	OldPrice string `json:"old_price"`
	Price string `json:"price"`
	TypeID int64 `json:"type_id"`
	Vat string `json:"vat"`
	NewDescriptionCategoryID int64 `json:"new_description_category_id"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Width int32 `json:"width"`
}

type V3User struct {
	ID string `json:"id"`
	Type string `json:"type_"`
}

type V1FbpDraftDirectTplDlvCreateRequestDirectDetails struct {
	TransportCompanyName string `json:"transport_company_name"`
	TimeslotStart string `json:"timeslot_start"`
	TrackingNumber string `json:"tracking_number"`
}

type ReportReport struct {
	File string `json:"file"`
	Params map[string]interface{} `json:"params"`
	ReportType string `json:"report_type"`
	Status string `json:"status"`
	Code string `json:"code"`
	CreatedAt string `json:"created_at"`
	Error string `json:"error"`
	ExpiresAt string `json:"expires_at"`
}

type V1FbpCheckActStateResponseStatus string

type ReportCreateCompanyPostingsReportRequest struct {
	Filter interface{} `json:"filter"`
	Language interface{} `json:"language"`
	With interface{} `json:"with"`
}

type V1FbpDraftDirectTplDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpArchiveGetResponse struct {
	DeclineReason interface{} `json:"decline_reason"`
	HasAct bool `json:"has_act"`
	ActFileUuid string `json:"act_file_uuid"`
	DeliveryDetails interface{} `json:"delivery_details"`
	ReceiveDate string `json:"receive_date"`
	BusinessFlowTypeID int64 `json:"business_flow_type_id"`
	CreatedDate string `json:"created_date"`
	HasLabel bool `json:"has_label"`
	ID int64 `json:"id"`
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	OrderDraftID int64 `json:"order_draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	OrderNumber string `json:"order_number"`
	PackageUnitsCount int32 `json:"package_units_count"`
	BundleID string `json:"bundle_id"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
}

type AnalyticsFilter struct {
	Key string `json:"key"`
	Op interface{} `json:"op"`
	Value string `json:"value"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo struct {
	Kpp string `json:"kpp"`
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
}

type Productv3GetProductListRequest struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
}

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequest struct {
	IsKgt bool `json:"is_kgt"`
	Search interface{} `json:"search"`
	Coordinates interface{} `json:"coordinates"`
	CountryCode string `json:"country_code"`
}

type V3GetFbsPostingResponseV3 struct {
	Result interface{} `json:"result"`
}

type SellerApiProductV1ResponseResult struct {
	ProductIds []interface{} `json:"product_ids"`
	Rejected []interface{} `json:"rejected"`
}

type CreateWarehouseFBSRequestOptions struct {
	Comment string `json:"comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
}

type SellerActionsListRequestActionTypeEnum string

type V2ReturnsRfbsRejectRequest struct {
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
	RejectionReasonID int64 `json:"rejection_reason_id"`
}

type V1GetProductInfoSubscriptionRequest struct {
	Skus []interface{} `json:"skus"`
}

type ReviewV2ReviewListV2ResponseReviewOrderStatusEnum string

type PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep struct {
	MinCharge interface{} `json:"min_charge"`
	TariffCharge interface{} `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"`
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
}

type V1FbpDraftGetResponse struct {
	ID int64 `json:"id"`
	IsCancelable bool `json:"is_cancelable"`
	Status interface{} `json:"status"`
	CreatedAt string `json:"created_at"`
	IsDeletable bool `json:"is_deletable"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	IsRegistrationAvailable bool `json:"is_registration_available"`
	PackageUnitsCount int32 `json:"package_units_count"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	CancellationState interface{} `json:"cancellation_state"`
	DeletedAt string `json:"deleted_at"`
	Editable bool `json:"editable"`
	Locked bool `json:"locked"`
	DeclineReason interface{} `json:"decline_reason"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponse struct {
	Result interface{} `json:"result"`
}

type ProductUpdateOfferIdRequestUpdateOfferId struct {
	NewOfferID string `json:"new_offer_id"`
	OfferID string `json:"offer_id"`
}

type V1WarehouseListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
}

type SellerReturnsv1MoneyWithoutCommission struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1FbpDraftDropOffProductValidateResponse struct {
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
}

type V1GetStrategyItemInfoResponseResult struct {
	StrategyID string `json:"strategy_id"`
	IsEnabled bool `json:"is_enabled"`
	StrategyProductPrice int32 `json:"strategy_product_price"`
	PriceDownloadedAt string `json:"price_downloaded_at"`
	StrategyCompetitorID int64 `json:"strategy_competitor_id"`
	StrategyCompetitorProductURL string `json:"strategy_competitor_product_url"`
}

type RowItemLegalEntityDocument struct {
	Number string `json:"number"`
	SaleDate string `json:"sale_date"`
}

type V1CarriageApproveRequest struct {
	CarriageID int64 `json:"carriage_id"`
	ContainersCount int32 `json:"containers_count"`
}

type V1SearchQueriesTopResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	ToUpdate []interface{} `json:"to_update"`
}

type V1UnarchiveWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError struct {
	SKU int64 `json:"sku"`
	Errors []interface{} `json:"errors"`
}

type V1fbpTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1SellerActionsUpdateVoucherRequestActionParameters struct {
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountValue float64 `json:"discount_value"`
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
}

type SortingOrder string

type CreateWarehouseFBSRequestFirstMileTypeEnum string

type ActionsV1ActionsAutoAddProductsUpdateResponse struct {
	BelowMinPrice []interface{} `json:"below_min_price"`
	ExtremelyLowPrice []interface{} `json:"extremely_low_price"`
	FailedPrice []interface{} `json:"failed_price"`
	Rejected []interface{} `json:"rejected"`
	UpdatedIds []interface{} `json:"updated_ids"`
}

type PickedSegmentSegment struct {
	Description string `json:"description"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	Type interface{} `json:"type_"`
}

type ProductImportProductsBySKUResponse struct {
	Result interface{} `json:"result"`
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProducts struct {
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
	CommissionsCurrencyCode string `json:"commissions_currency_code"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
}

type AnalyticsAnalyticsGetDataResponse struct {
	Result interface{} `json:"result"`
	Timestamp string `json:"timestamp"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type V1FbpDraftDropOffCreateResponse struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
}

type SearchQueriesTextRequestSortDir string

type FilterPeriod struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1CommentCreateResponse struct {
	CommentID string `json:"comment_id"`
}

type V1SellerActionsProductsCandidatesResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type PostingV4PostingFbsListResponsePostingsDeliveryMethod struct {
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
}

type V1GetFBSRatingIndexInfoV1Response struct {
	PeriodTo string `json:"period_to"`
	ProcessingCostsSum float64 `json:"processing_costs_sum"`
	CurrencyCode string `json:"currency_code"`
	Defects []interface{} `json:"defects"`
	Index float64 `json:"index"`
	PeriodFrom string `json:"period_from"`
}

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"`
}

type Productv1ProductInfoPicturesResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpAvailableTimeslotListResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type FinanceV1GetFinanceAccrualPostingsResponse struct {
	PostingAccruals []interface{} `json:"posting_accruals"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1QuestionTopSkuRequest struct {
	Limit int64 `json:"limit"`
}

type V1SellerActionsCreateInstallmentResponse struct {
	ActionID int64 `json:"action_id"`
}

type SearchQueriesTextRequestSortBy string

type V1FbpOrderDropOffTimetableResponseCalendar struct {
	DayOfWeek interface{} `json:"day_of_week"`
	CalendarItem interface{} `json:"calendar_item"`
}

type V1ProductActionTimerStatusResponse struct {
	Statuses interface{} `json:"statuses"`
}

type SellerApiProductV1ResponseResultDeactivate struct {
	ProductIds []interface{} `json:"product_ids"`
	Rejected []interface{} `json:"rejected"`
}

type MoneyMoneyCustomerPrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V3FbsPostingRequirementsV3 struct {
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GetFinanceBalanceV1ResponseSales struct {
	Amount interface{} `json:"amount"`
	AmountDetails interface{} `json:"amount_details"`
	Fee interface{} `json:"fee"`
}

type V1QuestionCountResponse struct {
	All int64 `json:"all"`
	New int64 `json:"new"`
	Processed int64 `json:"processed"`
	Unprocessed int64 `json:"unprocessed"`
	Viewed int64 `json:"viewed"`
}

type GetProductInfoStocksResponseStock struct {
	Present int32 `json:"present"`
	Reserved int32 `json:"reserved"`
	ShipmentType interface{} `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Type string `json:"type_"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type ChatChatSendFileResponse struct {
	Result string `json:"result"`
}

type PostingV1PostingFbpListResponse struct {
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type ProductV4GetUploadQuotaResponseTotalQuotaByCategory struct {
	CategoryID int64 `json:"category_id"`
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
}

type DetailsPayment struct {
	CurrencyCode string `json:"currency_code"`
	Payment float64 `json:"payment"`
}

type V1AnalyticsProductQueriesDetailsRequestSortDir string

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFee struct {
	Fees []interface{} `json:"fees"`
	SKU int64 `json:"sku"`
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type GetWarehouseFBSOperationStatusResponseStatusEnum string

type V1GetFinanceBalanceV1ResponsePartnerMoney struct {
	Value float64 `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type ReviewV2ReviewListV2ResponseReviewStatusEnum string

type V3ChatList struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
}

type V1TimeRangeStorageTariffication struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type V1FbpDraftPickUpDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type ProductV1ProductVisibilitySetRequestItemPlacement struct {
	Placement interface{} `json:"placement"`
	SKU int64 `json:"sku"`
}

type V2ProductInfoPicturesResponse struct {
	Items []interface{} `json:"items"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilter struct {
	ProviderIds []interface{} `json:"provider_ids"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	CutoffFrom string `json:"cutoff_from"`
	Statuses []interface{} `json:"statuses"`
	CutoffTo string `json:"cutoff_to"`
	DeliveringDateFrom string `json:"delivering_date_from"`
	DeliveringDateTo string `json:"delivering_date_to"`
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
}

type ActionParameter struct {
	PickedSegments []interface{} `json:"picked_segments"`
	Status interface{} `json:"status"`
	Title string `json:"title"`
	VoucherParameters interface{} `json:"voucher_parameters"`
	Warehouses []interface{} `json:"warehouses"`
	MinActionPercent float64 `json:"min_action_percent"`
	AutoStopActionReason interface{} `json:"auto_stop_action_reason"`
	BudgetSpent float64 `json:"budget_spent"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Type interface{} `json:"type_"`
	Addresses []interface{} `json:"addresses"`
	Budget float64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
}

type V1FbpCheckActStateResponse struct {
	CdnURL string `json:"cdn_url"`
	Error interface{} `json:"error"`
	Status interface{} `json:"status"`
}

type V1WarehouseInvalidProductsGetResponse struct {
	HasNext bool `json:"has_next"`
	LastID int64 `json:"last_id"`
	ValidationResults []interface{} `json:"validation_results"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ApproveDeclineDiscountTasksResponseResult struct {
	FailDetails []interface{} `json:"fail_details"`
	SuccessCount int32 `json:"success_count"`
	FailCount int32 `json:"fail_count"`
}

type ChatMessageModerateImageStatus string

type Productv5Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type WarehouseListV2ResponseWarehouse struct {
	CourierComment string `json:"courier_comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	FirstMile interface{} `json:"first_mile"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	PostingsLimit int32 `json:"postings_limit"`
	Status string `json:"status"`
	Timetable interface{} `json:"timetable"`
	WarehouseID int64 `json:"warehouse_id"`
	CreatedAt string `json:"created_at"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	IsComfort bool `json:"is_comfort"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
	WarehouseType string `json:"warehouse_type"`
	IsExpress bool `json:"is_express"`
	IsRFBS bool `json:"is_rfbs"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	Name string `json:"name"`
	SlaCutIn int64 `json:"sla_cut_in"`
	WorkingDays []interface{} `json:"working_days"`
	AddressInfo interface{} `json:"address_info"`
	CarriageLabelType interface{} `json:"carriage_label_type"`
	CutInTime int64 `json:"cut_in_time"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	IsKgt bool `json:"is_kgt"`
	Phone string `json:"phone"`
	UpdatedAt string `json:"updated_at"`
	WithItemList bool `json:"with_item_list"`
}

type Productv2DeleteProductsRequest struct {
	Products []interface{} `json:"products"`
}

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V1GetWarehouseFBSOperationStatusRequest struct {
	OperationID string `json:"operation_id"`
}

type V1ProductActionTimerStatusResponseStatuses struct {
	ExpiredAt string `json:"expired_at"`
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	ProductID int64 `json:"product_id"`
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type V1FbpDraftPickUpProductValidateResponseApprovedItem struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type Financev3FinanceTransactionTotalsV3Request struct {
	TransactionType string `json:"transaction_type"`
	Date interface{} `json:"date"`
	PostingNumber string `json:"posting_number"`
}

type V1GetFinanceBalanceV1ResponseClosingBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1ProductPricesDetailsResponse struct {
	Prices []interface{} `json:"prices"`
}

type V1FbpDraftDropOffRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type V1FbpDraftPickUpProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type FbsPostingProductExemplarCreateOrGetV6ResponseProduct struct {
	IsMandatoryMarkPossible bool `json:"is_mandatory_mark_possible"`
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
	HasImei bool `json:"has_imei"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
	IsJwUinNeeded bool `json:"is_jw_uin_needed"`
	IsMandatoryMarkNeeded bool `json:"is_mandatory_mark_needed"`
	IsRnptNeeded bool `json:"is_rnpt_needed"`
	Exemplars []interface{} `json:"exemplars"`
}

type ProductImportProductsBySKUResponseResult struct {
	UnmatchedSKUList []interface{} `json:"unmatched_sku_list"`
	TaskID int64 `json:"task_id"`
}

type FbsPostingDetailCourier struct {
	CarModel string `json:"car_model"`
	CarNumber string `json:"car_number"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type PostingV4PostingFbsListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type V1Empty interface{}

type SellerSellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseRequest struct {
	Search interface{} `json:"search"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1CancellationState struct {
	CancellationError interface{} `json:"cancellation_error"`
	CancellationStatus interface{} `json:"cancellation_status"`
}

type PostingV4PostingFbsListRequestSortDirEnum string

type V1PostingFbsSplitResponsePostingParent struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V2WarehouseListV2Request struct {
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type Fbpv1DeliveryDetails struct {
	DirectDetails interface{} `json:"direct_details"`
	DropOffPoint interface{} `json:"drop_off_point"`
	PickupDetails interface{} `json:"pickup_details"`
	SupplyType interface{} `json:"supply_type"`
}

type V1AddBarcodeRequest struct {
	Barcodes []interface{} `json:"barcodes"`
}

type PostinglistV3status struct {
	From string `json:"from"`
	To string `json:"to"`
}

type ProductV4GetUploadQuotaResponseOperationLimitsLimitTypeEnum string

type V1SellerActionsCreateVoucherRequestVoucherParameterTypeEnum string

type PostingV4PostingFbsListResponsePostingsCustomer struct {
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V1WarehouseFbsCreatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type Productv2GetProductListRequestFilterFilterVisibility string

type ActionsV1ActionsAutoAddProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	ProductIds []interface{} `json:"product_ids"`
}

type V2ConditionalCancellationMoveV2Request struct {
	CancellationID int64 `json:"cancellation_id"`
	Comment string `json:"comment"`
}

type V1QuestionListResponseQuestions struct {
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
	ProductURL string `json:"product_url"`
	AnswersCount int64 `json:"answers_count"`
	PublishedAt interface{} `json:"published_at"`
	QuestionLink string `json:"question_link"`
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
}

type V3ImportProductsRequest struct {
	Items []interface{} `json:"items"`
}

type GetProductInfoListResponseStocks struct {
	HasStock bool `json:"has_stock"`
	Stocks []interface{} `json:"stocks"`
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V1OrderValidationErrorErrorType string

type V1GetProductInfoSubscriptionResponse struct {
	Result []interface{} `json:"result"`
}

type V3Barcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type GetProductInfoListResponseStocksStock struct {
	Source string `json:"source"`
	Present int32 `json:"present"`
	Reserved int32 `json:"reserved"`
	SKU int64 `json:"sku"`
}

type V1SellerActionsCreateDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	MinActionPercent float64 `json:"min_action_percent"`
	Title string `json:"title"`
}

type V1AssemblyCarriageProductListRequestFilter struct {
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod struct {
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
}

type V1QuestionAnswerDeleteRequest struct {
	AnswerID string `json:"answer_id"`
	SKU int64 `json:"sku"`
}

type ReviewV2ReviewListV2RequestFiltersOrderStatusEnum string

type ItemPricev5 struct {
	MarketingSellerPrice float64 `json:"marketing_seller_price"`
	MinPrice float64 `json:"min_price"`
	NetPrice float64 `json:"net_price"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	Vat float64 `json:"vat"`
	RetailPrice float64 `json:"retail_price"`
	AutoActionEnabled bool `json:"auto_action_enabled"`
	CurrencyCode string `json:"currency_code"`
}

type PostingV1PostingFbpListResponsePostingsProducts struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
	Quantity int32 `json:"quantity"`
	SellerPrice interface{} `json:"seller_price"`
	SKU int64 `json:"sku"`
	CustomerPrice interface{} `json:"customer_price"`
}

type FbpCheckConsignmentNoteStateResponseStateType string

type V1FbpDraftListResponseItem struct {
	CancellationState interface{} `json:"cancellation_state"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	DeliveryDetails interface{} `json:"delivery_details"`
	IsDeletable bool `json:"is_deletable"`
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	Editable bool `json:"editable"`
	ID int64 `json:"id"`
	IsCancelable bool `json:"is_cancelable"`
	Locked bool `json:"locked"`
	PackageUnitsCount int32 `json:"package_units_count"`
	SupplyID string `json:"supply_id"`
	BundleID string `json:"bundle_id"`
}

type V3ChatHistoryRequest struct {
	ChatID string `json:"chat_id"`
	Direction string `json:"direction"`
	Filter interface{} `json:"filter"`
	FromMessageID int64 `json:"from_message_id"`
	Limit int64 `json:"limit"`
}

type V1GetReturnsListResponse struct {
	Returns []interface{} `json:"returns"`
	HasNext bool `json:"has_next"`
}

type ParameterAutoStopActionReasonEnum string

type RowItem struct {
	OfferID string `json:"offer_id"`
	SKU int64 `json:"sku"`
	Barcode string `json:"barcode"`
	Name string `json:"name"`
}

type ChatChatStartResponse struct {
	Result interface{} `json:"result"`
}

type Productv3Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type ProductGetImportProductsInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftDirectTplDlvEditResponse struct {
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
	Error interface{} `json:"error"`
}

type V1SellerActionsCreateDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type Financev3FinanceTransactionListV3Request struct {
	PageSize int64 `json:"page_size"`
	Filter interface{} `json:"filter"`
	Page int64 `json:"page"`
}

type V4FbsPostingShipPackageV4RequestProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
	ExemplarsIds []interface{} `json:"exemplarsIds"`
}

type V1SellerActionsCreateDiscountWithConditionRequest struct {
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
}

type SellerApiProductPrice struct {
	ProductID float64 `json:"product_id"`
	ActionPrice float64 `json:"action_price"`
	Stock float64 `json:"stock"`
}

type MoneyMoneyCurrentTariffCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type CreateCompanyPostingsReportRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	CustomerData bool `json:"customer_data"`
	JewelryCodes bool `json:"jewelry_codes"`
	AdditionalData bool `json:"additional_data"`
}

type GetRealizationReportResponseV2Header struct {
	Number string `json:"number"`
	PayerInn string `json:"payer_inn"`
	ReceiverInn string `json:"receiver_inn"`
	DocDate string `json:"doc_date"`
	PayerKpp string `json:"payer_kpp"`
	PayerName string `json:"payer_name"`
	ReceiverKpp string `json:"receiver_kpp"`
	ReceiverName string `json:"receiver_name"`
	StartDate string `json:"start_date"`
	StopDate string `json:"stop_date"`
	ContractDate string `json:"contract_date"`
	ContractNumber string `json:"contract_number"`
	CurrencySysName string `json:"currency_sys_name"`
}

type PriceIndexesIndexExternalData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type DeliveryMethodListResponseDeliveryMethod struct {
	CompanyID int64 `json:"company_id"`
	CreatedAt string `json:"created_at"`
	ID int64 `json:"id"`
	SlaCutIn int64 `json:"sla_cut_in"`
	TemplateID int64 `json:"template_id"`
	UpdatedAt string `json:"updated_at"`
	Cutoff string `json:"cutoff"`
	Name string `json:"name"`
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ReviewListRequest struct {
	Limit int32 `json:"limit"`
	SortDir string `json:"sort_dir"`
	Status string `json:"status"`
	LastID string `json:"last_id"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplar struct {
	GTD string `json:"gtd"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	Valid bool `json:"valid"`
	Errors []interface{} `json:"errors"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type MoneyMoneySaleCommission struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1FbpDraftDirectSellerDlvEditRequest struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDropOffProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type ReturnsRfbsGetV2ResponseClientReturnMethodType struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type V1SetProductStairwayDiscountByQuantityResponseError struct {
	Data []interface{} `json:"data"`
	SKU int64 `json:"sku"`
}

type SellerApiProductV1ResponseProductDeactivate struct {
	ProductID float64 `json:"product_id"`
	Reason string `json:"reason"`
}

type V1FbpDraftPickupCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1ArchiveDeclineReason struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1ProductUpdateOfferIdRequest struct {
	UpdateOfferID interface{} `json:"update_offer_id"`
}

type V1ProductUpdateAttributesRequestItem struct {
	Attributes interface{} `json:"attributes"`
	OfferID string `json:"offer_id"`
}

type SellerActionsListRequestStatusEnum string

type V3ChatDetailsInfo struct {
	CreatedAt string `json:"created_at"`
	ChatID string `json:"chat_id"`
	ChatStatus interface{} `json:"chat_status"`
	ChatType interface{} `json:"chat_type"`
}

type V1FbpDraftListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type V1SellerActionsChangeActivityRequest struct {
	IsTurnOn bool `json:"is_turn_on"`
	ActionID int64 `json:"action_id"`
}

type V1SellerActionsProductsCandidatesResponseProductQuantTypeEnum string

type V1GetFinanceBalanceV1ResponseReturnsMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V2CancellationInitiatorEnum string

type GetFinanceBalanceV1ResponseCashflows struct {
	Returns interface{} `json:"returns"`
	Sales interface{} `json:"sales"`
	Services []interface{} `json:"services"`
}

type ProductImportProductsPricesResponse struct {
	Result []interface{} `json:"result"`
}

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V3FinanceCashFlowStatementListResponse struct {
	Result interface{} `json:"result"`
}

type PriceIndexesIndexDataOzon struct {
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
	MinimalPrice string `json:"minimal_price"`
}

type V1SellerActionsVoucherGetRequest struct {
	ActionID int64 `json:"action_id"`
}

type FbsPostingShipV4RequestPackage struct {
	Products []interface{} `json:"products"`
}

type SellerSellerAPIArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type ProductBooleanResponse struct {
	Result bool `json:"result"`
}

type V2FbsPosting struct {
	OrderNumber string `json:"order_number"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
	Barcodes interface{} `json:"barcodes"`
	OrderID int64 `json:"order_id"`
	PostingNumber string `json:"posting_number"`
	ShipmentDate string `json:"shipment_date"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CreatedAt string `json:"created_at"`
	InProcessAt string `json:"in_process_at"`
}

type OperationItem struct {
	SKU int64 `json:"sku"`
	Name string `json:"name"`
}

type V1FbpCheckConsignmentNoteStateRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type V1SellerActionsCreateMultiLevelDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1SellerActionsListResponse struct {
	Actions []interface{} `json:"actions"`
	Total int64 `json:"total"`
}

type V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccruals struct {
	Accruals []interface{} `json:"accruals"`
	PostingNumber string `json:"posting_number"`
}

type GetWarehouseFBSOperationStatusResponseTypeEnum string

type ProductImportProductsPricesResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type ChatStartResponseResult struct {
	ChatID string `json:"chat_id"`
}

type V1ProductUpdateAttributesRequest struct {
	Items interface{} `json:"items"`
}

type ArrivalPassListRequestFilter struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	ArrivalReason string `json:"arrival_reason"`
	DropoffPointIds []interface{} `json:"dropoff_point_ids"`
	OnlyActivePasses bool `json:"only_active_passes"`
}

type Postingv3GetFbsPostingUnfulfilledListResponseResult struct {
	Count int64 `json:"count"`
	Postings []interface{} `json:"postings"`
}

type GetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type V1ProductGetRelatedSKUResponseItem struct {
	Availability string `json:"availability"`
	DeletedAt string `json:"deleted_at"`
	DeliverySchema string `json:"delivery_schema"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
}

type HumanTextsParam struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type ReviewV2ReviewListV2ResponseReview struct {
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
	CommentsAmount int32 `json:"comments_amount"`
	ID string `json:"id"`
	OrderStatus interface{} `json:"order_status"`
	PhotosAmount int32 `json:"photos_amount"`
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	Text string `json:"text"`
	VideosAmount int32 `json:"videos_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type RolesByTokenResponseRoles struct {
	Name string `json:"name"`
	Methods []interface{} `json:"methods"`
}

type TaskAutoModeratedInfo struct {
	MinPercent float64 `json:"min_percent"`
	MinPrice float64 `json:"min_price"`
	MaxPercent float64 `json:"max_percent"`
	MaxPrice float64 `json:"max_price"`
}

type FinanceCashFlowStatementListResponseReturnService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V1UpdateWarehouseFBSRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1CreatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyName string `json:"strategy_name"`
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []interface{} `json:"result"`
}

type ExemplarsMarks struct {
	MarkType string `json:"mark_type"`
	Mark string `json:"mark"`
}

type V1FbpDraftDirectProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ReportReportInfoRequest struct {
	Code string `json:"code"`
}

type WarehouseListResponseWarehouse struct {
	WarehouseID int64 `json:"warehouse_id"`
	IsTimetableEditable bool `json:"is_timetable_editable"`
	PostingsLimit int32 `json:"postings_limit"`
	Status string `json:"status"`
	IsRFBS bool `json:"is_rfbs"`
	CanPrintActInAdvance bool `json:"can_print_act_in_advance"`
	FirstMileType interface{} `json:"first_mile_type"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	Name string `json:"name"`
	IsKgt bool `json:"is_kgt"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
	MinWorkingDays int64 `json:"min_working_days"`
	WorkingDays interface{} `json:"working_days"`
	IsKarantin bool `json:"is_karantin"`
	IsAbleToSetPrice bool `json:"is_able_to_set_price"`
	IsPresorted bool `json:"is_presorted"`
}

type V1ItemIDsRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V1GetStrategyResponseResult struct {
	Competitors []interface{} `json:"competitors"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Type string `json:"type_"`
	UpdateType string `json:"update_type"`
}

type ProductGetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type V1CommentListResponse struct {
	Comments []interface{} `json:"comments"`
	Offset int32 `json:"offset"`
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"`
	PostingNumber string `json:"posting_number"`
}

type AddStrategyItemsResponseError struct {
	Code string `json:"code"`
	Error string `json:"error"`
	ProductID int64 `json:"product_id"`
}

type V3Address struct {
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	Region string `json:"region"`
	ZipCode string `json:"zip_code"`
	Comment string `json:"comment"`
	District string `json:"district"`
	Latitude float64 `json:"latitude"`
	PvzCode int64 `json:"pvz_code"`
	AddressTail string `json:"address_tail"`
	City string `json:"city"`
	Country string `json:"country"`
}

type DirectDetailsByTplDetails struct {
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"`
}

type V1SellerActionsCreateVoucherRequestDiscountTypeEnum string

type MarketingAction struct {
	Value int32 `json:"value"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Title string `json:"title"`
}

type PostingV4PostingFbsListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type V1GetFinanceBalanceV1ResponseOpeningBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type GetReturnsListResponseVisual struct {
	Status interface{} `json:"status"`
	ChangeMoment string `json:"change_moment"`
}

type Fbsv4FbsPostingShipV4Response struct {
	AdditionalData interface{} `json:"additional_data"`
	Result interface{} `json:"result"`
}

type PostingV4PostingFbsListResponsePostings struct {
	IsPresortable bool `json:"is_presortable"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	ExternalOrder interface{} `json:"external_order"`
	Substatus string `json:"substatus"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	VolumeWeight float64 `json:"volume_weight"`
	Cancellation interface{} `json:"cancellation"`
	InProcessAt string `json:"in_process_at"`
	OrderNumber string `json:"order_number"`
	ParentPostingNumber string `json:"parent_posting_number"`
	Products []interface{} `json:"products"`
	AvailableActions []interface{} `json:"available_actions"`
	Barcodes interface{} `json:"barcodes"`
	DeliveringDate string `json:"delivering_date"`
	DestinationPlaceName string `json:"destination_place_name"`
	LegalInfo interface{} `json:"legal_info"`
	PostingNumber string `json:"posting_number"`
	Tariffication interface{} `json:"tariffication"`
	Addressee interface{} `json:"addressee"`
	IsExpress bool `json:"is_express"`
	IsMultibox bool `json:"is_multibox"`
	Optional interface{} `json:"optional"`
	OrderID int64 `json:"order_id"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	ShipmentDate string `json:"shipment_date"`
	Status string `json:"status"`
	DeliveryMethod interface{} `json:"delivery_method"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	TrackingNumber string `json:"tracking_number"`
	Customer interface{} `json:"customer"`
	DeliverySchema string `json:"delivery_schema"`
	QuantumID int64 `json:"quantum_id"`
	PrrOption string `json:"prr_option"`
	AnalyticsData interface{} `json:"analytics_data"`
	FinancialData interface{} `json:"financial_data"`
	Requirements interface{} `json:"requirements"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	DestinationPlaceID int64 `json:"destination_place_id"`
}

type ResultError struct {
	Code string `json:"code"`
	Status string `json:"status"`
}

type V3ChatMessage struct {
	IsRead bool `json:"is_read"`
	MessageId int64 `json:"messageId"`
	ModerateImageStatus interface{} `json:"moderate_image_status"`
	User interface{} `json:"user"`
	Context interface{} `json:"context"`
	CreatedAt string `json:"created_at"`
	Data []interface{} `json:"data"`
	IsImage bool `json:"is_image"`
}

type FbpDraftDropOffProvinceListResponseProvince struct {
	Name string `json:"name"`
	PointsCount int32 `json:"points_count"`
	ProvinceUuid string `json:"province_uuid"`
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type SellerApiProductV1ResponseDeactivate struct {
	Result interface{} `json:"result"`
}

type V1FbpArchiveListRequest struct {
	Count string `json:"count"`
	LastID string `json:"last_id"`
}

type V1GenerateBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type Postingv1GetCarriageAvailableListRequest struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type V1AssemblyFbsProductListRequestFilter struct {
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1SellerActionsProductsListRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1AssemblyFbsPostingListRequestSortDirEnum string

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductCommission struct {
	Commission interface{} `json:"commission"`
	CommissionRatio string `json:"commission_ratio"`
	SaleAmount interface{} `json:"sale_amount"`
	SaleCommission interface{} `json:"sale_commission"`
	SalePrice interface{} `json:"sale_price"`
	SellerPrice interface{} `json:"seller_price"`
	Bonus interface{} `json:"bonus"`
	Coinvestment interface{} `json:"coinvestment"`
}

type AvailabilityReason struct {
	ID int64 `json:"id"`
	HumanText interface{} `json:"human_text"`
}

type GetProductRatingBySkuResponseRatingCondition struct {
	Cost float64 `json:"cost"`
	Description string `json:"description"`
	Fulfilled bool `json:"fulfilled"`
	Key string `json:"key"`
}

type Productv4GetProductAttributesV4Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type V1SearchQueriesTextResponseSearchQuery struct {
	ConversionToCart float64 `json:"conversion_to_cart"`
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
	SellersCount float64 `json:"sellers_count"`
	AddToCart float64 `json:"add_to_cart"`
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
}

type PostingV4PostingFbsListRequestFilter struct {
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	OrderID int64 `json:"order_id"`
	OrderNumbers []interface{} `json:"order_numbers"`
	ProviderIds []interface{} `json:"provider_ids"`
	Since string `json:"since"`
	Statuses []interface{} `json:"statuses"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	To string `json:"to"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type ChatInfoChatTypeEnum string

type CarriageCarriageGetResponseCancelAvailability struct {
	IsCancelAvailable bool `json:"is_cancel_available"`
	Reason string `json:"reason"`
}

type SellerApiActivateProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type V1OrderErrorTypeEnum string

type V1QuestionListResponse struct {
	HasNext bool `json:"has_next"`
	Questions interface{} `json:"questions"`
	LastID string `json:"last_id"`
}

type DeliveryMethodListV2RequestFilter struct {
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	ProviderIds []interface{} `json:"provider_ids"`
	Status []interface{} `json:"status"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1AnalyticsProductQueriesDetailsRequest struct {
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	LimitBySKU int32 `json:"limit_by_sku"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Skus []interface{} `json:"skus"`
}

type V1CreateWarehouseFBSRequest struct {
	AddressCoordinates interface{} `json:"address_coordinates"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	IsKgt bool `json:"is_kgt"`
	Name string `json:"name"`
	Options interface{} `json:"options"`
	Phone string `json:"phone"`
	TimeslotID int64 `json:"timeslot_id"`
	WorkingDays []interface{} `json:"working_days"`
	CutInTime int64 `json:"cut_in_time"`
	FirstMileType interface{} `json:"first_mile_type"`
}

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairwayStep struct {
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
	Discount int64 `json:"discount"`
}

type V2GetRealizationReportResponseV2 struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftDirectProductValidateRequestSkuItem struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1FbpEditTimeslotRequest struct {
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccrualsAccrual struct {
	TypeID int32 `json:"type_id"`
	AccrualDate string `json:"accrual_date"`
	Accrued interface{} `json:"accrued"`
	Quantity int32 `json:"quantity"`
	SellerPrice interface{} `json:"seller_price"`
	SKU int64 `json:"sku"`
}

type V1GetTreeResponseItem struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	CategoryName string `json:"category_name"`
	Children []interface{} `json:"children"`
	Disabled bool `json:"disabled"`
	TypeID int64 `json:"type_id"`
	TypeName string `json:"type_name"`
}

type OperationPosting struct {
	OrderDate string `json:"order_date"`
	PostingNumber string `json:"posting_number"`
	WarehouseID int64 `json:"warehouse_id"`
	DeliverySchema string `json:"delivery_schema"`
}

type V2CancellationStateEnumFilters string

type ProductV4GetUploadQuotaResponseOperationLimits struct {
	Limit int64 `json:"limit"`
	LimitType interface{} `json:"limit_type"`
}

type V1ArchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type CreateReportResponse struct {
	Result interface{} `json:"result"`
}

type Productv3GetProductAttributesV3Request struct {
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
	Filter interface{} `json:"filter"`
}

type GetProductInfoListResponsePriceIndexes struct {
	OzonIndexData interface{} `json:"ozon_index_data"`
	SelfMarketplacesIndexData interface{} `json:"self_marketplaces_index_data"`
	ColorIndex interface{} `json:"color_index"`
	ExternalIndexData interface{} `json:"external_index_data"`
}

type PostingV4PostingFbsListResponsePostingsRequirements struct {
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
	ProductsRequiringWeight []interface{} `json:"products_requiring_weight"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostings struct {
	Products []interface{} `json:"products"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	Requirements interface{} `json:"requirements"`
	Status string `json:"status"`
	IsExpress bool `json:"is_express"`
	LegalInfo interface{} `json:"legal_info"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	OrderID int64 `json:"order_id"`
	QuantumID int64 `json:"quantum_id"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	DestinationPlaceName string `json:"destination_place_name"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	Substatus string `json:"substatus"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	DeliverySchema string `json:"delivery_schema"`
	InProcessAt string `json:"in_process_at"`
	Customer interface{} `json:"customer"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	FinancialData interface{} `json:"financial_data"`
	IsPresortable bool `json:"is_presortable"`
	ParentPostingNumber string `json:"parent_posting_number"`
	AvailableActions []interface{} `json:"available_actions"`
	Barcodes interface{} `json:"barcodes"`
	ExternalOrder interface{} `json:"external_order"`
	Optional interface{} `json:"optional"`
	ShipmentDate string `json:"shipment_date"`
	TrackingNumber string `json:"tracking_number"`
	AnalyticsData interface{} `json:"analytics_data"`
	IsMultibox bool `json:"is_multibox"`
	OrderNumber string `json:"order_number"`
	PrrOption string `json:"prr_option"`
	Tariffication interface{} `json:"tariffication"`
	VolumeWeight float64 `json:"volume_weight"`
	Cancellation interface{} `json:"cancellation"`
	DeliveringDate string `json:"delivering_date"`
	DeliveryMethod interface{} `json:"delivery_method"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	Addressee interface{} `json:"addressee"`
	PostingNumber string `json:"posting_number"`
}

type V1FbpDraftDirectRegistrateRequest struct {
	SupplyID string `json:"supply_id"`
	RowVersion int64 `json:"row_version"`
}

type V3FbsPostingDetail struct {
	OrderID int64 `json:"order_id"`
	AvailableActions interface{} `json:"available_actions"`
	Barcodes interface{} `json:"barcodes"`
	IsExpress bool `json:"is_express"`
	AdditionalData []interface{} `json:"additional_data"`
	Requirements interface{} `json:"requirements"`
	PreviousSubstatus string `json:"previous_substatus"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	Cancellation interface{} `json:"cancellation"`
	Courier interface{} `json:"courier"`
	DeliveringDate string `json:"delivering_date"`
	RelatedPostings interface{} `json:"related_postings"`
	DeliveryMethod interface{} `json:"delivery_method"`
	Optional interface{} `json:"optional"`
	ParentPostingNumber string `json:"parent_posting_number"`
	ProductExemplars interface{} `json:"product_exemplars"`
	ProviderStatus string `json:"provider_status"`
	Status string `json:"status"`
	InProcessAt string `json:"in_process_at"`
	TrackingNumber string `json:"tracking_number"`
	Addressee interface{} `json:"addressee"`
	Customer interface{} `json:"customer"`
	OrderNumber string `json:"order_number"`
	Products []interface{} `json:"products"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	ShipmentDate string `json:"shipment_date"`
	FactDeliveryDate string `json:"fact_delivery_date"`
	FinancialData interface{} `json:"financial_data"`
	LegalInfo interface{} `json:"legal_info"`
	PostingNumber string `json:"posting_number"`
	Substatus string `json:"substatus"`
	Tariffication interface{} `json:"tariffication"`
	AnalyticsData interface{} `json:"analytics_data"`
	DeliveryPrice string `json:"delivery_price"`
}

type V1ProductUpdateDiscountResponse struct {
	Result bool `json:"result"`
}

type ArrivalpassArrivalPassListResponse struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	Cursor string `json:"cursor"`
}

type V1AddStrategyItemsRequest struct {
	StrategyID string `json:"strategy_id"`
	ProductID []interface{} `json:"product_id"`
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"`
}

type ChatMessageContext struct {
	OrderNumber string `json:"order_number"`
	SKU string `json:"sku"`
}

type PostingV4PostingFbsListResponsePostingsLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type Productv4GetProductAttributesV4ResponseResult struct {
	AttributesWithDefaults []interface{} `json:"attributes_with_defaults"`
	Barcodes interface{} `json:"barcodes"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Depth int64 `json:"depth"`
	Images interface{} `json:"images"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	DimensionUnit string `json:"dimension_unit"`
	ModelInfo interface{} `json:"model_info"`
	PrimaryImage string `json:"primary_image"`
	ColorImage string `json:"color_image"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	ID int64 `json:"id"`
	SKU string `json:"sku"`
	TypeID int64 `json:"type_id"`
	Barcode string `json:"barcode"`
	Height int64 `json:"height"`
	PDFList []interface{} `json:"pdf_list"`
	Weight int64 `json:"weight"`
	WeightUnit string `json:"weight_unit"`
	Width int64 `json:"width"`
	Attributes []interface{} `json:"attributes"`
}

type V1GetStrategyItemInfoRequest struct {
	ProductID int64 `json:"product_id"`
}

type WarehouseFirstMileType struct {
	DropoffPointID string `json:"dropoff_point_id"`
	DropoffTimeslotID int64 `json:"dropoff_timeslot_id"`
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
	FirstMileType string `json:"first_mile_type"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation struct {
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
}

type V1AnalyticsProductQueriesDetailsResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type V2ProductInfoPicturesResponseItem struct {
	ColorPhoto []interface{} `json:"color_photo"`
	Photo360 []interface{} `json:"photo_360"`
	Errors []interface{} `json:"errors"`
	ProductID int64 `json:"product_id"`
	PrimaryPhoto []interface{} `json:"primary_photo"`
	Photo []interface{} `json:"photo"`
}

type V1ProductUpdateAttributesRequestValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type V1GetStrategyListResponse struct {
	Strategies []interface{} `json:"strategies"`
	Total int32 `json:"total"`
}

type SellerApiProduct struct {
	PriceMaxElastic float64 `json:"price_max_elastic"`
	MinBoost float64 `json:"min_boost"`
	MaxBoost float64 `json:"max_boost"`
	ID float64 `json:"id"`
	Price float64 `json:"price"`
	AlertMaxActionPriceFailed bool `json:"alert_max_action_price_failed"`
	AlertMaxActionPrice float64 `json:"alert_max_action_price"`
	AddMode string `json:"add_mode"`
	MinStock float64 `json:"min_stock"`
	CurrentBoost float64 `json:"current_boost"`
	PriceMinElastic float64 `json:"price_min_elastic"`
	ActionPrice float64 `json:"action_price"`
	MaxActionPrice float64 `json:"max_action_price"`
	Stock float64 `json:"stock"`
}

type CreateReportResponseCodeNoDeadline struct {
	Code string `json:"code"`
}

type V3Addressee struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2Product struct {
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	FreeStock int64 `json:"free_stock"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
	SKU int64 `json:"sku"`
}

type V2ReturnsRfbsCompensateRequest struct {
	CompensationAmount string `json:"compensation_amount"`
	ReturnID int64 `json:"return_id"`
}

type V3GetProductInfoListResponse struct {
	Items []interface{} `json:"items"`
}

type PostingV4PostingFbsListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

type MoneyMoneyTotalAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
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

type RatingStatusEnum string

type Productv2DeleteProductsResponse struct {
	Status []interface{} `json:"status"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejected struct {
	Code interface{} `json:"code"`
	ProductID int64 `json:"product_id"`
	Reason string `json:"reason"`
}

type PostingPostingProductCancelRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	Items []interface{} `json:"items"`
	PostingNumber string `json:"posting_number"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDelivery struct {
	Services []interface{} `json:"services"`
	TotalAccrued interface{} `json:"total_accrued"`
}

type V1ProductGetRelatedSKURequest struct {
	SKU interface{} `json:"sku"`
}

type V3Cancellation struct {
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
}

type MoneyMoneySalePrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type ReviewV2ReviewListV2RequestSortDirEnum string

type FbpDraftGetResponseDeclineReason struct {
	FailedSKUIds []interface{} `json:"failed_sku_ids"`
	Message string `json:"message"`
}

type FbsPostingProductExemplarSetV6RequestProducts struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type DropOffPointCoordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type V3FinanceCashFlowStatementListResponsePeriod struct {
	ID int64 `json:"id"`
	Begin string `json:"begin"`
	End string `json:"end"`
}

type FinanceTransactionListV3RequestFilter struct {
	Date interface{} `json:"date"`
	OperationType []interface{} `json:"operation_type"`
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
}

type V1GetWarehouseFBSOperationStatusResponse struct {
	Error interface{} `json:"error"`
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Type interface{} `json:"type_"`
}

type V1GetProductStairwayDiscountByQuantityRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1CommentDeleteRequest struct {
	CommentID string `json:"comment_id"`
}

type V1GetStrategyItemInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1SellerActionsUpdateDiscountRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type GetSellerActionsV1ResponseAction struct {
	ID float64 `json:"id"`
	Title string `json:"title"`
	DateStart string `json:"date_start"`
	AutoAddDates []interface{} `json:"auto_add_dates"`
	ParticipatingProductsCount float64 `json:"participating_products_count"`
	IsParticipating bool `json:"is_participating"`
	OrderAmount float64 `json:"order_amount"`
	ActionType string `json:"action_type"`
	DateEnd string `json:"date_end"`
	FreezeDate string `json:"freeze_date"`
	BannedProductsCount float64 `json:"banned_products_count"`
	DiscountType string `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	Description string `json:"description"`
	PotentialProductsCount float64 `json:"potential_products_count"`
	IsVoucherAction bool `json:"is_voucher_action"`
	WithTargeting bool `json:"with_targeting"`
}

type V2ReturnsRfbsFilter struct {
	OfferID string `json:"offer_id"`
	PostingNumber string `json:"posting_number"`
	GroupState []interface{} `json:"group_state"`
	CreatedAt interface{} `json:"created_at"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProduct struct {
	Error string `json:"error"`
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
	Valid bool `json:"valid"`
}

type GetConditionalCancellationListV2ResponseState struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	State interface{} `json:"state"`
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProductsActions struct {
	ActionID string `json:"action_id"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	DiscountPercent float64 `json:"discount_percent"`
	DiscountValue float64 `json:"discount_value"`
	IsFromSeller bool `json:"is_from_seller"`
	Description string `json:"description"`
}

type GetProductAttributesResponseImage struct {
	Index int64 `json:"index"`
	Default bool `json:"default"`
	FileName string `json:"file_name"`
}

type V3CustomerFbsLists struct {
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
}

type V2ReturnsRfbsGetResponse struct {
	Returns interface{} `json:"returns"`
}

type CreateWarehouseFBSRequestWorkingDaysEnum string

type V1GetStrategyResponse struct {
	Result interface{} `json:"result"`
}

type ProductV1ProductVisibilitySetResponseItems struct {
	SelectPermission interface{} `json:"select_permission"`
	SellerItemPlacement interface{} `json:"seller_item_placement"`
	SellerItemPlacementList []interface{} `json:"seller_item_placement_list"`
	ShowcasesVisibility interface{} `json:"showcases_visibility"`
	ShowcasesVisibilityList []interface{} `json:"showcases_visibility_list"`
	SKU int64 `json:"sku"`
	Warnings []interface{} `json:"warnings"`
}

type ReportCreateCompanyProductsReportRequest struct {
	Visibility interface{} `json:"visibility"`
	Language interface{} `json:"language"`
	OfferID []interface{} `json:"offer_id"`
	Search string `json:"search"`
	SKU []interface{} `json:"sku"`
}

type Postingv3GetFbsPostingListRequest struct {
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
	Dir string `json:"dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type V1FbpAvailableTimeslotListRequest struct {
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
	SupplyID string `json:"supply_id"`
}

type V3DeliveryMethod struct {
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
}

type V1GetCompetitorsResponse struct {
	Competitor []interface{} `json:"competitor"`
	Total int32 `json:"total"`
}

type SellerInfoResponseCompany struct {
	Inn string `json:"inn"`
	LegalName string `json:"legal_name"`
	Name string `json:"name"`
	Ogrn string `json:"ogrn"`
	OwnershipForm string `json:"ownership_form"`
	TaxSystem interface{} `json:"tax_system"`
	Country string `json:"country"`
	Currency string `json:"currency"`
}

type V3FbsPostingProductExemplarInfoV3 struct {
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	Rnpt string `json:"rnpt"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Imei []interface{} `json:"imei"`
	ExemplarID int64 `json:"exemplar_id"`
	MandatoryMark string `json:"mandatory_mark"`
}

type WarehouseDeliveryMethodListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type ReportCreateCompanyPostingsReportRequestFilter struct {
	IsExpress interface{} `json:"is_express"`
	DeliverySchema []interface{} `json:"delivery_schema"`
	OfferID string `json:"offer_id"`
	ProcessedAtFrom string `json:"processed_at_from"`
	SKU []interface{} `json:"sku"`
	StatusAlias []interface{} `json:"status_alias"`
	Statuses []interface{} `json:"statuses"`
	Title string `json:"title"`
	WarehouseID []interface{} `json:"warehouse_id"`
	CancelReasonID []interface{} `json:"cancel_reason_id"`
	ProcessedAtTo string `json:"processed_at_to"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
}

type PostingV1PostingFbpListRequestSortDirEnum string

type SellerSellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type V1Competitor struct {
	Coefficient float64 `json:"coefficient"`
	CompetitorID int64 `json:"competitor_id"`
}

type PriceIndexesIndexOzonData struct {
	PriceIndexValue float64 `json:"price_index_value"`
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
}

type V1AnalyticsProductQueriesResponse struct {
	Total int64 `json:"total"`
	AnalyticsPeriod interface{} `json:"analytics_period"`
	Items []interface{} `json:"items"`
	PageCount int64 `json:"page_count"`
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []interface{} `json:"provinces"`
}

type V1FbpOrderListResponseItem struct {
	BundleSummary interface{} `json:"bundle_summary"`
	CancellationState interface{} `json:"cancellation_state"`
	DeliveryDetails interface{} `json:"delivery_details"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	CreatedDate string `json:"created_date"`
	HasLabel bool `json:"has_label"`
	OrderNumber string `json:"order_number"`
	PackageUnitsCount int32 `json:"package_units_count"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	ID int64 `json:"id"`
	Locked bool `json:"locked"`
	ReceiveDate string `json:"receive_date"`
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	SupplyID string `json:"supply_id"`
}

type V1BundleItemErrorEnum string

type V1FbpDraftDropOffPointListRequest struct {
	PageSize int32 `json:"page_size"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
	NextPageNumber int32 `json:"next_page_number"`
}

type GetReturnsListResponseCompensationStatus struct {
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
	SysName string `json:"sys_name"`
}

type SellerReturnsv1MoneyUtilization struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1CarriageCreateRequest struct {
	AllBlrTraceable bool `json:"all_blr_traceable"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result interface{} `json:"result"`
}

type V1GetCompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type GetFinanceBalanceV1ResponseReturnsDetails struct {
	PointsForDiscounts string `json:"points_for_discounts"`
	Revenue interface{} `json:"revenue"`
	PartnerPrograms interface{} `json:"partner_programs"`
}

type TimetableWorkingHours struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type Polygonv1Empty interface{}

type GetReturnsListResponseLogistic struct {
	TechnicalReturnMoment string `json:"technical_return_moment"`
	FinalMoment string `json:"final_moment"`
	CancelledWithCompensationMoment string `json:"cancelled_with_compensation_moment"`
	ReturnDate string `json:"return_date"`
	Barcode string `json:"barcode"`
}

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V3GetProductInfoListResponsePromotion struct {
	Type interface{} `json:"type_"`
	IsEnabled bool `json:"is_enabled"`
}

type V3Customer struct {
	Address interface{} `json:"address"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type ActionsV1ActionsAutoAddProductsDeleteResponse struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1AssemblyFbsPostingListResponsePostingProduct struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type ProductV1ProductVisibilityInfoResponse struct {
	Items []interface{} `json:"items"`
}

type V2ServiceType string

type V1AnalyticsProductQueriesRequest struct {
	SortDir interface{} `json:"sort_dir"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Skus []interface{} `json:"skus"`
	SortBy interface{} `json:"sort_by"`
}

type RelatedPostingCancelReason struct {
	PostingNumber string `json:"posting_number"`
	Reasons []interface{} `json:"reasons"`
}

type AnalyticsProductQueriesResponseAnalyticsPeriod struct {
	DateTo string `json:"date_to"`
	DateFrom string `json:"date_from"`
}

type V1GetDiscountTaskListResponse struct {
	Result []interface{} `json:"result"`
}

type V1AssemblyCarriageProductListRequest struct {
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
}

type V1SellerActionsProductsListResponseProductQuantTypeEnum string

type V1ArchiveStatusEnum string

type PostingV4PostingFbsListResponsePostingsCustomerAddress struct {
	AddressTail string `json:"address_tail"`
	City string `json:"city"`
	Comment string `json:"comment"`
	Country string `json:"country"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	District string `json:"district"`
	Region string `json:"region"`
	ZipCode string `json:"zip_code"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountTypeEnum string

type V1GetProductStairwayDiscountByQuantityResponseStairways struct {
	SKU int64 `json:"sku"`
	Stairway interface{} `json:"stairway"`
	Status interface{} `json:"status"`
	Enabled bool `json:"enabled"`
}

type ReportCreateDiscountedRequest interface{}

type PostingV1PostingFbpListRequestFilter struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
	Since string `json:"since"`
	Statuses []interface{} `json:"statuses"`
	To string `json:"to"`
}

type V1GetProductRatingBySkuRequest struct {
	Skus interface{} `json:"skus"`
}

type V1GetFinanceBalanceV1ResponseSalesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type ChatChatSendFileRequest struct {
	Name string `json:"name"`
	Base64Content string `json:"base64_content"`
	ChatID string `json:"chat_id"`
}

type V1FbpDraftDropOffCreateRequestDeliveryDetails struct {
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
}

type CreatedAt struct {
	From string `json:"from"`
	To string `json:"to"`
}

type GooglerpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1SellerActionsUpdateDiscountWithConditionRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1AnalyticsProductQueriesDetailsResponse struct {
	AnalyticsPeriod interface{} `json:"analytics_period"`
	PageCount int64 `json:"page_count"`
	Queries []interface{} `json:"queries"`
	Total int64 `json:"total"`
}

type V1DescriptionCategoryTipsResponse struct {
	Result []interface{} `json:"result"`
}

type FilterWithQuant struct {
	Created bool `json:"created"`
	Exists bool `json:"exists"`
}

type V1AssemblyCarriageProductListResponseProduct struct {
	ProductName string `json:"product_name"`
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type WarehouseTimetable struct {
	WorkingHours []interface{} `json:"working_hours"`
	TimetableFrom string `json:"timetable_from"`
	TimetableTo string `json:"timetable_to"`
}

type PolygonBindRequestwhLocation struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type DirectDetailsBySellerDetails struct {
	DriverName string `json:"driver_name"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
	VehicleType string `json:"vehicle_type"`
}

type SellerInfoResponseRatingTypeEnum string

type V2GetConditionalCancellationListV2Request struct {
	Filters interface{} `json:"filters"`
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
	With interface{} `json:"with"`
}

type SellerActionsUpdateMultiLevelDiscountRequestActionParametersDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type V1GetReturnsListRequest struct {
	Limit int32 `json:"limit"`
	LastID int64 `json:"last_id"`
	Filter interface{} `json:"filter"`
}

type V6FbsPostingProductExemplarSetV6Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	MultiBoxQty int32 `json:"multi_box_qty"`
}

type V1FbpEditTimeslotResponseReserveFailureType string

type V1GetCompetitorsRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1AnalyticsProductQueriesDetailsRequestSortBy string

type PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

type V3ChatListResponse struct {
	Chats interface{} `json:"chats"`
	TotalUnreadCount int64 `json:"total_unread_count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
}

type AssemblyFbsProductListResponseProducts struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	Postings []interface{} `json:"postings"`
	ProductName string `json:"product_name"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type GetProductInfoListResponseError struct {
	Code string `json:"code"`
	Field string `json:"field"`
	Level interface{} `json:"level"`
	State string `json:"state"`
	Texts interface{} `json:"texts"`
	AttributeID int64 `json:"attribute_id"`
}

type FbsPostingShipV4ResponseShipAdditionalData struct {
	PostingNumber string `json:"posting_number"`
	Products interface{} `json:"products"`
}

type SellerApiGetSellerProductV1Response struct {
	Result interface{} `json:"result"`
}

type V1OrderDraftValidationError struct {
	Errors []interface{} `json:"errors"`
}

type V1ReturnsCompanyFbsInfoRequest struct {
	Filter interface{} `json:"filter"`
	Pagination interface{} `json:"pagination"`
}

type V4GetProductInfoStocksRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type DetailsDeliveryDetails struct {
	Amount float64 `json:"amount"`
	DeliveryServices interface{} `json:"delivery_services"`
	Total float64 `json:"total"`
}

type SellerSellerAPIArrivalPassCreateRequestArrivalPass struct {
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
}

type GetProductAttributesResponseImage360 struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V2DeliveryMethodListV2Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type ProductGetProductAttributesV4ResponseAttribute struct {
	ID int64 `json:"id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type GetCompetitorsResponseCompetitorInfo struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type V1ProductUpdateOfferIdResponseError struct {
	Message string `json:"message"`
	OfferID string `json:"offer_id"`
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"`
}

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
}

type ActionsV1ActionsAutoAddProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type PostingFinancialDataProduct struct {
	Actions []interface{} `json:"actions"`
	CurrencyCode string `json:"currency_code"`
	CustomerCurrencyCode string `json:"customer_currency_code"`
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
}

type V1WarehouseFbsUpdatePickUpTimeslotListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type Productv3GetProductListRequestFilter struct {
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
	OfferID interface{} `json:"offer_id"`
}

type V1SellerActionsProductsListResponseProduct struct {
	BasePrice float64 `json:"base_price"`
	Currency string `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
	SKU []interface{} `json:"sku"`
	ActionPrice float64 `json:"action_price"`
	IsActive bool `json:"is_active"`
	ProductID int64 `json:"product_id"`
	QuantSize int64 `json:"quant_size"`
	QuantType interface{} `json:"quant_type"`
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpOrderDropOffCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1QuestionAnswerListResponse struct {
	Answers interface{} `json:"answers"`
	LastID string `json:"last_id"`
}

type V1FbpDraftDropOffDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type Productv2ProductsStocksResponse struct {
	Result []interface{} `json:"result"`
}

type PostingV4PostingFbsListResponsePostingsTariffication struct {
	NextTariffMinCharge interface{} `json:"next_tariff_min_charge"`
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffCharge interface{} `json:"next_tariff_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	NextTariffType string `json:"next_tariff_type"`
	CurrentTariffCharge interface{} `json:"current_tariff_charge"`
	CurrentTariffMinCharge interface{} `json:"current_tariff_min_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
}

type TypeTimeOfDay struct {
	Minutes int32 `json:"minutes"`
	Nanos int32 `json:"nanos"`
	Seconds int32 `json:"seconds"`
	Hours int32 `json:"hours"`
}

type ReturnsCompanyFbsInfoResponsePassInfo struct {
	Count int32 `json:"count"`
	IsRequired bool `json:"is_required"`
}

type ProductGetProductInfoDescriptionResponseProduct struct {
	Description string `json:"description"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDirectGetTimeslotRequest struct {
	BundleID string `json:"bundle_id"`
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GetReturnsListResponsePlaceNow struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type V1ProductUpdateAttributesRequestAttribute struct {
	ComplexID int64 `json:"complex_id"`
	ID int64 `json:"id"`
	Values interface{} `json:"values"`
}

type Polygonv1PolygonBindRequest struct {
	DeliveryMethodID int32 `json:"delivery_method_id"`
	Polygons []interface{} `json:"polygons"`
	WarehouseLocation interface{} `json:"warehouse_location"`
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"`
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type PostingV4PostingFbsListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendar struct {
	CalendarItem interface{} `json:"calendar_item"`
	DayOfWeek interface{} `json:"day_of_week"`
}

type GetReturnsListResponseExemplar struct {
	ID int64 `json:"id"`
}

type V1DeclineDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V3Dimensions struct {
	Length string `json:"length"`
	Weight string `json:"weight"`
	Width string `json:"width"`
	Height string `json:"height"`
}

type AnalyticsDataRow struct {
	Metrics []interface{} `json:"metrics"`
	Dimensions []interface{} `json:"dimensions"`
}

type GetProductInfoListResponseSource struct {
	CreatedAt string `json:"created_at"`
	QuantCode string `json:"quant_code"`
	ShipmentType interface{} `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type ProductExemplar struct {
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	ExemplarID int64 `json:"exemplar_id"`
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"`
}

type SellerReturnsv1MoneyStorage struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type Productv1ProductImportPicturesRequest struct {
	ColorImage string `json:"color_image"`
	Images interface{} `json:"images"`
	Images360 interface{} `json:"images360"`
	ProductID int64 `json:"product_id"`
}

type V1GetFinanceBalanceV1Request struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearch struct {
	Address string `json:"address"`
	Types []interface{} `json:"types"`
}

type MoneyMoneyBonus struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type GetProductRatingBySkuResponseProductRating struct {
	Rating float64 `json:"rating"`
	Groups interface{} `json:"groups"`
	SKU int64 `json:"sku"`
}

type PostingV4PostingFbsUnfulfilledListRequestSortDirEnum string

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPoint struct {
	ID string `json:"id"`
	LastTransitTimeLocal interface{} `json:"last_transit_time_local"`
	Type interface{} `json:"type_"`
	Address string `json:"address"`
	Coordinates interface{} `json:"coordinates"`
	DiscountPercent float64 `json:"discount_percent"`
}

type PostingV4PostingFbsUnfulfilledListRequest struct {
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Translit bool `json:"translit"`
	With interface{} `json:"with"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualNonItemFee struct {
	Accrued interface{} `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type GetReturnsListResponseAdditionalInfo struct {
	IsOpened bool `json:"is_opened"`
	IsSuperEconom bool `json:"is_super_econom"`
}

type SellerActionsListResponseAction struct {
	AllowDelete bool `json:"allow_delete"`
	HighlightURL string `json:"highlight_url"`
	IsEditable bool `json:"is_editable"`
	IsParticipated bool `json:"is_participated"`
	IsTurnOn bool `json:"is_turn_on"`
	SKUCount int64 `json:"sku_count"`
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V3ChatInfo struct {
	LastMessageID int64 `json:"last_message_id"`
	UnreadCount int64 `json:"unread_count"`
	Chat interface{} `json:"chat"`
	FirstUnreadMessageID int64 `json:"first_unread_message_id"`
}

type SourceShipmentType string

type ReportListResponseResult struct {
	Reports []interface{} `json:"reports"`
	Total int32 `json:"total"`
}

type V1GetAttributesResponseAttribute struct {
	CategoryDependent bool `json:"category_dependent"`
	Description string `json:"description"`
	DictionaryID int64 `json:"dictionary_id"`
	IsAspect bool `json:"is_aspect"`
	IsCollection bool `json:"is_collection"`
	IsRequired bool `json:"is_required"`
	Name string `json:"name"`
	Type string `json:"type_"`
	GroupID int64 `json:"group_id"`
	GroupName string `json:"group_name"`
	ID int64 `json:"id"`
	AttributeComplexID int64 `json:"attribute_complex_id"`
	MaxValueCount int64 `json:"max_value_count"`
	ComplexIsCollection bool `json:"complex_is_collection"`
}

type V1FbpDraftDirectProductValidateResponseApprovedItem struct {
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
}

type V1ItemSfboAttribute string

type PriceIndexesColorIndex string

type V1GetFinanceBalanceV1Response struct {
	Cashflows interface{} `json:"cashflows"`
	Total interface{} `json:"total"`
}

type FbpCheckActStateResponseErrorReason string

type ProductV1ProductVisibilitySetResponseItemsErrors struct {
	SKU int64 `json:"sku"`
	Code string `json:"code"`
}

type GetDiscountTaskListV2ResponseTask struct {
	ApprovedDiscount float64 `json:"approved_discount"`
	AutoModeratedInfo interface{} `json:"auto_moderated_info"`
	CreatedAt string `json:"created_at"`
	Email string `json:"email"`
	EndAt string `json:"end_at"`
	LastName string `json:"last_name"`
	RequestedPrice float64 `json:"requested_price"`
	ApprovedPrice float64 `json:"approved_price"`
	IsAutoModerated bool `json:"is_auto_moderated"`
	MinAutoPrice float64 `json:"min_auto_price"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	EditedTill string `json:"edited_till"`
	EditedTillDuration int64 `json:"edited_till_duration"`
	FirstName string `json:"first_name"`
	ModeratedAt string `json:"moderated_at"`
	Name string `json:"name"`
	OriginalPrice float64 `json:"original_price"`
	Patronymic string `json:"patronymic"`
	EndAtDuration int64 `json:"end_at_duration"`
	ID int64 `json:"id"`
	ReductionFactor float64 `json:"reduction_factor"`
	RequestedDiscount float64 `json:"requested_discount"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
}

type V5FbsPostingProductExemplarStatusV5Request struct {
	PostingNumber string `json:"posting_number"`
}

type V1ProductUpdateOfferIdResponse struct {
	Errors interface{} `json:"errors"`
}

type ItemCommissionsv5 struct {
	FBODelivToCustomerAmount float64 `json:"fbo_deliv_to_customer_amount"`
	FBODirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
	FBSDelivToCustomerAmount float64 `json:"fbs_deliv_to_customer_amount"`
	FBSDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
	FBSFirstMileMaxAmount float64 `json:"fbs_first_mile_max_amount"`
	FBSFirstMileMinAmount float64 `json:"fbs_first_mile_min_amount"`
	SalesPercentFBO float64 `json:"sales_percent_fbo"`
	SalesPercentFBP float64 `json:"sales_percent_fbp"`
	FBODirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
	FBOReturnFlowAmount float64 `json:"fbo_return_flow_amount"`
	FBSDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
	FBSReturnFlowAmount float64 `json:"fbs_return_flow_amount"`
	SalesPercentFBS float64 `json:"sales_percent_fbs"`
	SalesPercentRFBS float64 `json:"sales_percent_rfbs"`
}

type V3TimeRange struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type ProductV1ProductVisibilityInfoResponseItemShowcasesVisibilityEnum string

type SellerApiGetSellerActionsV1Response struct {
	Result []interface{} `json:"result"`
}

type V1FbpArchiveListResponseItem struct {
	PackageUnitsCount int32 `json:"package_units_count"`
	RowVersion int64 `json:"row_version"`
	Status interface{} `json:"status"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
	ExternalOrderID string `json:"external_order_id"`
	ActFileUuid string `json:"act_file_uuid"`
	SupplyID string `json:"supply_id"`
	BundleID string `json:"bundle_id"`
	CreatedDate string `json:"created_date"`
	DeclineReason interface{} `json:"decline_reason"`
	DeliveryDetails interface{} `json:"delivery_details"`
	HasAct bool `json:"has_act"`
	HasLabel bool `json:"has_label"`
	WarehouseID int64 `json:"warehouse_id"`
	OrderDraftID int64 `json:"order_draft_id"`
	ReceiveDate string `json:"receive_date"`
	WhcOrderID int64 `json:"whc_order_id"`
}

type V1GetTreeResponse struct {
	Result []interface{} `json:"result"`
}

type PostingV4PostingFbsListResponsePostingsCancellation struct {
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
}

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type ProductV1ProductVisibilitySetRequest struct {
	ItemPlacement []interface{} `json:"item_placement"`
}

type FbsPostingShipV4RequestPackageProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

type SellerReturnsv1MoneyCommission struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1FbpOrderListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type V1SellerActionsVoucherGetResponse struct {
	File string `json:"file"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityEnum string

type V6FbsPostingProductExemplarCreateOrGetV6Request struct {
	PostingNumber string `json:"posting_number"`
}

type GetProductInfoListResponseStatuses struct {
	IsCreated bool `json:"is_created"`
	Status string `json:"status"`
	StatusDescription string `json:"status_description"`
	StatusFailed string `json:"status_failed"`
	StatusUpdatedAt string `json:"status_updated_at"`
	ModerateStatus string `json:"moderate_status"`
	StatusName string `json:"status_name"`
	StatusTooltip string `json:"status_tooltip"`
	ValidationStatus string `json:"validation_status"`
}

type ActionParameterVoucherParameter struct {
	CountCodes int64 `json:"count_codes"`
	IsPrivate bool `json:"is_private"`
	Type interface{} `json:"type_"`
}

type V1SearchQueriesTextResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type V1GetRealizationReportPostingResponseRow struct {
	DeliveryCommission interface{} `json:"delivery_commission"`
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
	RowNumber int32 `json:"row_number"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	Order interface{} `json:"order"`
	LegalEntityDocument interface{} `json:"legal_entity_document"`
	CommissionRatio float64 `json:"commission_ratio"`
}

type SellerApiProductIDsV1Request struct {
	ActionID float64 `json:"action_id"`
	ProductIds []interface{} `json:"product_ids"`
}

type GetProductRatingBySkuResponseRatingImproveAttribute struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type V2GetDiscountTaskListV2Request struct {
	LastID int64 `json:"last_id"`
	Limit int64 `json:"limit"`
	Status interface{} `json:"status"`
}

type CreateReportResponseCode struct {
	Code string `json:"code"`
}

type V1ProductInfoWrongVolumeResponse struct {
	Cursor string `json:"cursor"`
	Products interface{} `json:"products"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairways struct {
	Stairway interface{} `json:"stairway"`
	Enabled bool `json:"enabled"`
	SKU int64 `json:"sku"`
}

type StockShipmentType string

type ProductGetImportProductsInfoResponseResult struct {
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type V1Barcode struct {
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type V1SetPostingsResponse struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsProductsListResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts struct {
	CustomerPrice interface{} `json:"customer_price"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	Commission interface{} `json:"commission"`
	Payout float64 `json:"payout"`
	ProductID int64 `json:"product_id"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
}

type V1SellerActionsArchiveRequest struct {
	ActionID int64 `json:"action_id"`
}

type QuestionV1QuestionAnswerListResponseAnswerStatusPublicationEnum string

type V1TimeRangeReturnDate struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type V1ProductPricesDetailsRequest struct {
	Skus []interface{} `json:"skus"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProduct struct {
	Delivery interface{} `json:"delivery"`
	SKU int64 `json:"sku"`
	Commission interface{} `json:"commission"`
}

type V1DeclineDiscountTasksRequestTask struct {
	ID int64 `json:"id"`
	SellerComment string `json:"seller_comment"`
}

type FilterStatusEnum string

type V1SellerActionsCreateInstallmentRequest struct {
	Title string `json:"title"`
	DateStart string `json:"date_start"`
}

type FbpWarehouseListResponseWarehouse struct {
	SupplyTypes []interface{} `json:"supply_types"`
	TimezoneName string `json:"timezone_name"`
	AddressDetailing interface{} `json:"address_detailing"`
	ID int64 `json:"id"`
	IsBonded bool `json:"is_bonded"`
	Name string `json:"name"`
	PartnerName string `json:"partner_name"`
}

type V1ReviewListResponse struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
}

type ProductProductUnarchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V5FbsPostingProductExemplarValidateV5RequestProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type AnalyticsAnalyticsGetDataRequest struct {
	Offset int64 `json:"offset"`
	Sort []interface{} `json:"sort"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Dimension []interface{} `json:"dimension"`
	Filters []interface{} `json:"filters"`
	Limit int64 `json:"limit"`
	Metrics []interface{} `json:"metrics"`
}

type V1SetProductStairwayDiscountByQuantityResponse struct {
	Accepted bool `json:"accepted"`
	Errors []interface{} `json:"errors"`
	Warnings []interface{} `json:"warnings"`
}

type ArrivalpassArrivalPassListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type ErrorErrorLevel string

type GetReturnsListRequestFilter struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
	ProductName string `json:"product_name"`
	Barcode string `json:"barcode"`
	ReturnSchema string `json:"return_schema"`
	CompensationStatusID int32 `json:"compensation_status_id"`
	StorageTarifficationStartDate interface{} `json:"storage_tariffication_start_date"`
	VisualStatusChangeMoment interface{} `json:"visual_status_change_moment"`
	OrderID int64 `json:"order_id"`
	OfferID string `json:"offer_id"`
	VisualStatusName string `json:"visual_status_name"`
	WarehouseID int64 `json:"warehouse_id"`
	LogisticReturnDate interface{} `json:"logistic_return_date"`
}

type ProductV1ProductVisibilitySetResponse struct {
	ItemsErrors []interface{} `json:"items_errors"`
	Items []interface{} `json:"items"`
}

type V1DeleteStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1PostingFbsSplitResponsePosting struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type CompanyTaxSystemEnum string

type V4GetProductAttributesResponseModelInfo struct {
	ModelID int64 `json:"model_id"`
	Count int64 `json:"count"`
}

type V3GetProductInfoListResponseItem struct {
	TypeID int64 `json:"type_id"`
	Images360 []interface{} `json:"images360"`
	IsPrepaymentAllowed bool `json:"is_prepayment_allowed"`
	MinPrice string `json:"min_price"`
	ModelInfo interface{} `json:"model_info"`
	PriceIndexes interface{} `json:"price_indexes"`
	CurrencyCode string `json:"currency_code"`
	ID int64 `json:"id"`
	Price string `json:"price"`
	CreatedAt string `json:"created_at"`
	HasDiscountedFBOItem bool `json:"has_discounted_fbo_item"`
	SKU int64 `json:"sku"`
	ColorImage []interface{} `json:"color_image"`
	DiscountedFBOStocks int32 `json:"discounted_fbo_stocks"`
	Errors []interface{} `json:"errors"`
	IsSuper bool `json:"is_super"`
	OfferID string `json:"offer_id"`
	UpdatedAt string `json:"updated_at"`
	Barcodes []interface{} `json:"barcodes"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	IsAutoarchived bool `json:"is_autoarchived"`
	OldPrice string `json:"old_price"`
	Promotions []interface{} `json:"promotions"`
	VisibilityDetails interface{} `json:"visibility_details"`
	Availabilities []interface{} `json:"availabilities"`
	IsArchived bool `json:"is_archived"`
	IsDiscounted bool `json:"is_discounted"`
	IsKgt bool `json:"is_kgt"`
	Name string `json:"name"`
	Statuses interface{} `json:"statuses"`
	VolumeWeight float64 `json:"volume_weight"`
	Commissions []interface{} `json:"commissions"`
	Images []interface{} `json:"images"`
	PrimaryImage []interface{} `json:"primary_image"`
	Sources []interface{} `json:"sources"`
	Stocks interface{} `json:"stocks"`
	Vat string `json:"vat"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Percent int64 `json:"percent"`
}

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"`
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplar struct {
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	GTD string `json:"gtd"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer struct {
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V3AdditionalDataItem struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type FinanceCashFlowStatementListResponseDetails struct {
	Loan float64 `json:"loan"`
	Return interface{} `json:"return"`
	RFBS interface{} `json:"rfbs"`
	Services interface{} `json:"services"`
	Others interface{} `json:"others"`
	EndBalanceAmount float64 `json:"end_balance_amount"`
	BeginBalanceAmount float64 `json:"begin_balance_amount"`
	Delivery interface{} `json:"delivery"`
	InvoiceTransfer float64 `json:"invoice_transfer"`
	Payments []interface{} `json:"payments"`
	Period interface{} `json:"period"`
}

type Productv4Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
	Visibility interface{} `json:"visibility"`
}

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type GetWarehouseFBSOperationStatusResponseResult struct {
	EntityID int64 `json:"entity_id"`
}

type DeleteProductsRequestProduct struct {
	OfferID string `json:"offer_id"`
}

type V2DeliveryMethodListV2Response struct {
	HasNext bool `json:"has_next"`
	DeliveryMethods []interface{} `json:"delivery_methods"`
	Cursor string `json:"cursor"`
}

type V1FbpDraftDirectSellerDlvCreateRequestDirectDetails struct {
	VehicleType string `json:"vehicle_type"`
	DriverName string `json:"driver_name"`
	TimeslotStart string `json:"timeslot_start"`
	VehicleNumber string `json:"vehicle_number"`
}

type PostingV4PostingFbsListResponsePostingsProducts struct {
	Imei []interface{} `json:"imei"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	Name string `json:"name"`
	Price interface{} `json:"price"`
	ProductColor string `json:"product_color"`
	SKU int64 `json:"sku"`
	Weight float64 `json:"weight"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
}

type ItemMarketing struct {
	Actions []interface{} `json:"actions"`
}

type V1QuestionListRequest struct {
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
}

type SellerApiGetSellerProductV1ResponseResult struct {
	Products []interface{} `json:"products"`
	Total float64 `json:"total"`
	LastID float64 `json:"last_id"`
}

type DeliveryMethodListRequestFilter struct {
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	ProviderID int64 `json:"provider_id"`
}

type V1QuestionTopSkuResponse struct {
	SKU interface{} `json:"sku"`
}

type FbpGetLabelResponseLabelCreationStateTypeEnum string

type V1ProductFbsSplit struct {
	Quantity int64 `json:"quantity"`
	ProductID int64 `json:"product_id"`
}

type SellerInfoResponseSubscription struct {
	IsPremium bool `json:"is_premium"`
	Type interface{} `json:"type_"`
}

type GetReturnsListResponseProduct struct {
	CommissionPercent float64 `json:"commission_percent"`
	Commission interface{} `json:"commission"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	Name string `json:"name"`
	Price interface{} `json:"price"`
	PriceWithoutCommission interface{} `json:"price_without_commission"`
}

type V1SellerActionsUpdateDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProducts struct {
	Commission interface{} `json:"commission"`
	CustomerPrice interface{} `json:"customer_price"`
	OldPrice float64 `json:"old_price"`
	Payout float64 `json:"payout"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	Actions []interface{} `json:"actions"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
}

type V1WarehouseFbsCreatePickUpTimeslotListResponse struct {
	IsPickupSupported bool `json:"is_pickup_supported"`
	Timeslots []interface{} `json:"timeslots"`
}

type ImportProductsRequestPdfList struct {
	Index int64 `json:"index"`
	Name string `json:"name"`
	SrcURL string `json:"src_url"`
}

type V1WarehouseWithInvalidProductsResponse struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1AssemblyCarriagePostingListRequestFilter struct {
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1WarehouseListRequestWith struct {
	AbleToSetPrice bool `json:"able_to_set_price"`
}

type V1FbpDraftDirectGetTimeslotResponseTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1UpdateWarehouseFBSRequest struct {
	Options interface{} `json:"options"`
	Phone string `json:"phone"`
	WarehouseID int64 `json:"warehouse_id"`
	WorkingDays []interface{} `json:"working_days"`
	AddressCoordinates interface{} `json:"address_coordinates"`
	Name string `json:"name"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
}

type V1FbpDraftDropOffProductValidateResponseApprovedItem struct {
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V2ReturnsRfbsListResponse struct {
	Returns interface{} `json:"returns"`
}

type V1AssemblyFbsPostingListRequestFilter struct {
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type GetConditionalCancellationListV2ResponseCancellationReason struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type V1CreatePricingStrategyResponseResult struct {
	StrategyID string `json:"strategy_id"`
}

type GetStrategyIDsByItemIDsResponseProductInfo struct {
	ProductID int64 `json:"product_id"`
	StrategyID string `json:"strategy_id"`
}

type ProductInfoWrongVolumeResponseWrongVolumeProduct struct {
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	Length int64 `json:"length"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	Weight int64 `json:"weight"`
}

type V1AssemblyFbsPostingListResponsePosting struct {
	AssemblyCode string `json:"assembly_code"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1PostingFbsSplitResponse struct {
	ParentPosting interface{} `json:"parent_posting"`
	Postings []interface{} `json:"postings"`
}

type ProductImportProductsPricesRequest struct {
	Prices []interface{} `json:"prices"`
}

type V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError struct {
	SKU int64 `json:"sku"`
	Errors []interface{} `json:"errors"`
}

type V1SellerActionsUpdateDiscountWithConditionRequestActionParameters struct {
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
}

type V1AssemblyFbsPostingListRequest struct {
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
}

type V1GetAttributeValuesResponse struct {
	Result []interface{} `json:"result"`
	HasNext bool `json:"has_next"`
}

type ChatChatSendMessageResponse struct {
	Result string `json:"result"`
}

type V1ReviewChangeStatusRequest struct {
	ReviewIds []interface{} `json:"review_ids"`
	Status string `json:"status"`
}

type V1QuestionAnswerCreateResponse struct {
	AnswerID string `json:"answer_id"`
}

type ProductGetProductInfoDescriptionRequest struct {
	ProductID int64 `json:"product_id"`
	OfferID string `json:"offer_id"`
}

type PostingErrorTypeEnum string

type DeleteProductsResponseDeleteStatus struct {
	Error string `json:"error"`
	IsDeleted bool `json:"is_deleted"`
	OfferID string `json:"offer_id"`
}

type V1SellerInfoResponse struct {
	Ratings []interface{} `json:"ratings"`
	Subscription interface{} `json:"subscription"`
	Company interface{} `json:"company"`
}

type V1GetAttributesRequest struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	Language interface{} `json:"language"`
	TypeID int64 `json:"type_id"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsRequest struct {
	SKU interface{} `json:"sku"`
	OfferID interface{} `json:"offer_id"`
}

type PriceIndexesIndexDataExternal struct {
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
	MinimalPrice string `json:"minimal_price"`
}

type ChatChatStartRequest struct {
	PostingNumber string `json:"posting_number"`
}

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID int64 `json:"return_id"`
	ReturnMethodDescription string `json:"return_method_description"`
}

type GetProductInfoListResponseAvailability struct {
	SKU int64 `json:"sku"`
	Source string `json:"source"`
	Availability string `json:"availability"`
	Reasons []interface{} `json:"reasons"`
}

type PostingV4PostingFbsListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1DeliveryDetailsDirectDetails struct {
	BySellerDetails interface{} `json:"by_seller_details"`
	ByTPLDetails interface{} `json:"by_tpl_details"`
	TimeslotDetails interface{} `json:"timeslot_details"`
}

type FbpWarehouseListResponseAddressDetailing struct {
	Country string `json:"country"`
	House string `json:"house"`
	Region string `json:"region"`
	Street string `json:"street"`
	Zipcode string `json:"zipcode"`
	City string `json:"city"`
}

type GetWarehouseFBSOperationStatusResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1GenerateBarcodeResult struct {
	Barcode string `json:"barcode"`
	ProductID int64 `json:"product_id"`
	Code string `json:"code"`
	Error string `json:"error"`
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date interface{} `json:"date"`
}

type V1ReviewInfoRequest struct {
	ReviewID string `json:"review_id"`
}

type V1ItemResponse struct {
	Quant int32 `json:"quant"`
	IsQuantEditable bool `json:"is_quant_editable"`
	Quantity int32 `json:"quantity"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	Barcode string `json:"barcode"`
	VolumeInLitres float64 `json:"volume_in_litres"`
	ShipmentType interface{} `json:"shipment_type"`
	Tags []interface{} `json:"tags"`
	PlacementZone string `json:"placement_zone"`
	IconPath string `json:"icon_path"`
	OfferID string `json:"offer_id"`
	TotalVolumeInLitres float64 `json:"total_volume_in_litres"`
	ContractorItemCode string `json:"contractor_item_code"`
	SfboAttribute interface{} `json:"sfbo_attribute"`
}

type PostingV4PostingFbsListResponsePostingsTarifficationStep struct {
	TariffDeadlineAt string `json:"tariff_deadline_at"`
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
	MinCharge interface{} `json:"min_charge"`
	TariffCharge interface{} `json:"tariff_charge"`
}

type GetProductRatingBySkuResponseRatingGroup struct {
	Rating float64 `json:"rating"`
	Weight float64 `json:"weight"`
	Conditions interface{} `json:"conditions"`
	ImproveAtLeast int64 `json:"improve_at_least"`
	ImproveAttributes interface{} `json:"improve_attributes"`
	Key string `json:"key"`
	Name string `json:"name"`
}

type V1WarehouseInvalidProductsGetRequest struct {
	LastID int64 `json:"last_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type PolygonBindRequestpolygon struct {
	PolygonID int64 `json:"polygon_id"`
	Time int64 `json:"time"`
}

type V1CommentListRequest struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	ReviewID string `json:"review_id"`
	SortDir interface{} `json:"sort_dir"`
}

type MoneyMoneyAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type GetUploadQuotaResponseTotal struct {
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
	QuotaByCategory interface{} `json:"quota_by_category"`
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type V3AddresseeFbsLists struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"`
}

type V1ProductActionTimerStatusRequest struct {
	ProductIds interface{} `json:"product_ids"`
}

type GetReturnsListResponsePlaceTarget struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponseResult struct {
	Reserved int64 `json:"reserved"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponse struct {
	Products []interface{} `json:"products"`
	Total int64 `json:"total"`
}

type ChatChatSendMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text string `json:"text"`
}

type AnalyticsGetDataResponseResult struct {
	Data []interface{} `json:"data"`
	Totals []interface{} `json:"totals"`
}

type GetImportProductsInfoResponseResultItem struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Status string `json:"status"`
	Errors []interface{} `json:"errors"`
}

type V5FbsPostingProductExemplarStatusV5Response struct {
	Status string `json:"status"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type FbsPostingProductExemplarSetV6RequestExemplars struct {
	ExemplarID int64 `json:"exemplar_id"`
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
}

type ReturnsCompanyFbsInfoRequestPagination struct {
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type ReportReportListRequest struct {
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	ReportType interface{} `json:"report_type"`
}

type V1FbpDraftDropOffProductValidateResponseRejectedItem struct {
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
}

type DetailsReturnDetails struct {
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	ReturnServices interface{} `json:"return_services"`
}

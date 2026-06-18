package internal

type ProductImportProductsPricesRequest struct {
	Prices []interface{} `json:"prices"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponse struct {
	Rejected []interface{} `json:"rejected"`
	UpdatedIds []interface{} `json:"updated_ids"`
	BelowMinPrice []interface{} `json:"below_min_price"`
	ExtremelyLowPrice []interface{} `json:"extremely_low_price"`
	FailedPrice []interface{} `json:"failed_price"`
}

type PostingV4PostingFbsListResponsePostingsAnalyticsData struct {
	Region string `json:"region"`
	TPLProvider string `json:"tpl_provider"`
	WarehouseID int64 `json:"warehouse_id"`
	City string `json:"city"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result interface{} `json:"result"`
}

type GetProductInfoListResponseStocks struct {
	HasStock bool `json:"has_stock"`
	Stocks []interface{} `json:"stocks"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplarMark struct {
	MarkType string `json:"mark_type"`
	Valid bool `json:"valid"`
	Errors []interface{} `json:"errors"`
	Mark string `json:"mark"`
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1WarehouseWithInvalidProductsResponse struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type FilterPeriod struct {
	From string `json:"from"`
	To string `json:"to"`
}

type GetSellerActionsV1ResponseAction struct {
	Title string `json:"title"`
	AutoAddDates []interface{} `json:"auto_add_dates"`
	IsVoucherAction bool `json:"is_voucher_action"`
	OrderAmount float64 `json:"order_amount"`
	ID float64 `json:"id"`
	ActionType string `json:"action_type"`
	DateStart string `json:"date_start"`
	DateEnd string `json:"date_end"`
	BannedProductsCount float64 `json:"banned_products_count"`
	DiscountType string `json:"discount_type"`
	Description string `json:"description"`
	FreezeDate string `json:"freeze_date"`
	PotentialProductsCount float64 `json:"potential_products_count"`
	WithTargeting bool `json:"with_targeting"`
	ParticipatingProductsCount float64 `json:"participating_products_count"`
	IsParticipating bool `json:"is_participating"`
	DiscountValue float64 `json:"discount_value"`
}

type V1GetTreeResponse struct {
	Result []interface{} `json:"result"`
}

type V1CommentListRequest struct {
	ReviewID string `json:"review_id"`
	SortDir interface{} `json:"sort_dir"`
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ValidationResultItem struct {
	Size interface{} `json:"size"`
	SKU int64 `json:"sku"`
	WeightG float64 `json:"weight_g"`
}

type CreateReportResponseCodeNoDeadline struct {
	Code string `json:"code"`
}

type V1QuestionListRequestFilter struct {
	DateTo string `json:"date_to"`
	Status string `json:"status"`
	DateFrom string `json:"date_from"`
}

type ReturnsCompanyFbsInfoResponseDropOffPoints struct {
	Address string `json:"address"`
	BoxCount int32 `json:"box_count"`
	Name string `json:"name"`
	PassInfo interface{} `json:"pass_info"`
	PlaceID int64 `json:"place_id"`
	ReturnsCount int32 `json:"returns_count"`
	UtcOffset string `json:"utc_offset"`
	ID int64 `json:"id"`
	WarehousesIds []interface{} `json:"warehouses_ids"`
}

type V1fbpDeliveryDetails struct {
	DropOffPoint interface{} `json:"drop_off_point"`
	PickupDetails interface{} `json:"pickup_details"`
	SupplyType interface{} `json:"supply_type"`
	DirectDetails interface{} `json:"direct_details"`
}

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type PostingV1PostingFbpListResponsePostingsFinancialDataProductsActions struct {
	Description string `json:"description"`
	ActionID string `json:"action_id"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	DiscountPercent float64 `json:"discount_percent"`
	DiscountValue float64 `json:"discount_value"`
	IsFromSeller bool `json:"is_from_seller"`
}

type PostingV1PostingFbpListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
}

type SellerApiProductIDsV1Request struct {
	ActionID float64 `json:"action_id"`
	ProductIds []interface{} `json:"product_ids"`
}

type ReportReportListResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpCheckConsignmentNoteStateResponse struct {
	ErrorMessage string `json:"error_message"`
	LabelURL string `json:"label_url"`
	State interface{} `json:"state"`
}

type V1OrderStatusEnum string

type ReturnsRfbsGetV2ResponseReturnReason struct {
	ID int32 `json:"id"`
	IsDefect bool `json:"is_defect"`
	Name string `json:"name"`
}

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type V1GetAttributeValuesResponseDictionaryValue struct {
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
	ID int64 `json:"id"`
}

type ProtobufAny struct {
	TypeUrl string `json:"typeUrl"`
	Value string `json:"value"`
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"`
	PostingNumber string `json:"posting_number"`
}

type ReportReportInfoRequest struct {
	Code string `json:"code"`
}

type V1FbpGetLabelResponse struct {
	State interface{} `json:"state"`
	LabelURL string `json:"label_url"`
}

type AnalyticsSorting struct {
	Key string `json:"key"`
	Order interface{} `json:"order"`
}

type PostingV4PostingFbsListRequestSortDirEnum string

type V1GetFinanceBalanceV1ResponseOpeningBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type V1SellerActionsProductsCandidatesResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type FilterStatusEnum string

type V3GetFbsPostingListResponseV3 struct {
	Result interface{} `json:"result"`
}

type V5FbsPostingProductExemplarValidateV5Response struct {
	Products []interface{} `json:"products"`
}

type V1ReturnsRfbsActionSetRequest struct {
	ReturnForBackWay float64 `json:"return_for_back_way"`
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
	CompensationAmount float64 `json:"compensation_amount"`
	ID int32 `json:"id"`
	RejectionReasonID int32 `json:"rejection_reason_id"`
}

type V1SellerActionsUpdateDiscountWithConditionRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
}

type V1FbpDraftDirectCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type AddStrategyItemsResponseError struct {
	Code string `json:"code"`
	Error string `json:"error"`
	ProductID int64 `json:"product_id"`
}

type V1ProductInfoWarehouseStocksRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	WarehouseID int64 `json:"warehouse_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication struct {
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffCharge interface{} `json:"next_tariff_charge"`
	NextTariffType string `json:"next_tariff_type"`
	CurrentTariffCharge interface{} `json:"current_tariff_charge"`
	CurrentTariffMinCharge interface{} `json:"current_tariff_min_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	NextTariffMinCharge interface{} `json:"next_tariff_min_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
}

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"`
}

type V1AnalyticsProductQueriesDetailsResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1FbpOrderDropOffTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendar struct {
	CalendarItem interface{} `json:"calendar_item"`
	DayOfWeek interface{} `json:"day_of_week"`
}

type V1QuestionAnswerListRequest struct {
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
	LastID interface{} `json:"last_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements struct {
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
	ProductsRequiringWeight []interface{} `json:"products_requiring_weight"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
}

type V1FbpDraftPickupDlvEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	PickupDetails interface{} `json:"pickup_details"`
}

type ProductV1ProductVisibilityInfoResponseItemShowcasesVisibilityEnum string

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPointAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type FinanceCashFlowStatementListResponseCashFlow struct {
	ServicesAmount float64 `json:"services_amount"`
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"`
	CurrencyCode string `json:"currency_code"`
	Period interface{} `json:"period"`
	OrdersAmount float64 `json:"orders_amount"`
	ReturnsAmount float64 `json:"returns_amount"`
	CommissionAmount float64 `json:"commission_amount"`
}

type ReviewInfoResponsePhoto struct {
	Height int32 `json:"height"`
	URL string `json:"url"`
	Width int32 `json:"width"`
}

type V1ProductActionTimerStatusRequest struct {
	ProductIds interface{} `json:"product_ids"`
}

type GetReturnsListResponseCompensation struct {
	ChangeMoment string `json:"change_moment"`
	Status interface{} `json:"status"`
}

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"`
}

type V1GetProductInfoSubscriptionRequest struct {
	Skus []interface{} `json:"skus"`
}

type V2WarehouseListV2Request struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
}

type V1FbpDraftDirectRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type SellerActionsListResponseAction struct {
	ActionParameters interface{} `json:"action_parameters"`
	AllowDelete bool `json:"allow_delete"`
	HighlightURL string `json:"highlight_url"`
	IsEditable bool `json:"is_editable"`
	IsParticipated bool `json:"is_participated"`
	IsTurnOn bool `json:"is_turn_on"`
	SKUCount int64 `json:"sku_count"`
	ActionID int64 `json:"action_id"`
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GetProductInfoListResponseVisibilityDetails struct {
	HasPrice bool `json:"has_price"`
	HasStock bool `json:"has_stock"`
}

type ReportReport struct {
	Params map[string]interface{} `json:"params"`
	ReportType string `json:"report_type"`
	Status string `json:"status"`
	Code string `json:"code"`
	CreatedAt string `json:"created_at"`
	Error string `json:"error"`
	ExpiresAt string `json:"expires_at"`
	File string `json:"file"`
}

type V1GetTreeResponseItem struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	CategoryName string `json:"category_name"`
	Children []interface{} `json:"children"`
	Disabled bool `json:"disabled"`
	TypeID int64 `json:"type_id"`
	TypeName string `json:"type_name"`
}

type FinanceV1GetFinanceAccrualByDayRequest struct {
	Date string `json:"date"`
	LastID string `json:"last_id"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponse struct {
	Result interface{} `json:"result"`
}

type PostingV4PostingFbsListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V1GetTreeRequest struct {
	Language interface{} `json:"language"`
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

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type ChatRead struct {
	ChatID string `json:"chat_id"`
	FromMessageID int64 `json:"from_message_id"`
}

type WarehouseDeliveryMethodListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WarehouseAddressInfo struct {
	Longitude float64 `json:"longitude"`
	Utc string `json:"utc"`
	Address string `json:"address"`
	Latitude float64 `json:"latitude"`
}

type GetProductRatingBySkuResponseProductRating struct {
	Groups interface{} `json:"groups"`
	SKU int64 `json:"sku"`
	Rating float64 `json:"rating"`
}

type Postingv3GetFbsPostingListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
	Dir string `json:"dir"`
	Filter interface{} `json:"filter"`
}

type V1AssemblyCarriageProductListResponseProduct struct {
	ProductName string `json:"product_name"`
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V3DeliveryMethod struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpOrderPickUpDlvEditRequestPickUpDetails struct {
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type V1UpdateWarehouseFBSFirstMileRequest struct {
	CutInTime int64 `json:"cut_in_time"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	FirstMileType interface{} `json:"first_mile_type"`
	TimeslotID int64 `json:"timeslot_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1GetProductStairwayDiscountByQuantityRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1GetStrategyItemInfoResponseResult struct {
	StrategyID string `json:"strategy_id"`
	IsEnabled bool `json:"is_enabled"`
	StrategyProductPrice int32 `json:"strategy_product_price"`
	PriceDownloadedAt string `json:"price_downloaded_at"`
	StrategyCompetitorID int64 `json:"strategy_competitor_id"`
	StrategyCompetitorProductURL string `json:"strategy_competitor_product_url"`
}

type GetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type ErrorHumanTexts struct {
	ShortDescription string `json:"short_description"`
	AttributeName string `json:"attribute_name"`
	Description string `json:"description"`
	HintCode string `json:"hint_code"`
	Message string `json:"message"`
	Params []interface{} `json:"params"`
}

type GetWarehouseFBSOperationStatusResponseTypeEnum string

type PostingV4PostingFbsUnfulfilledListResponsePostings struct {
	Barcodes interface{} `json:"barcodes"`
	LegalInfo interface{} `json:"legal_info"`
	AvailableActions []interface{} `json:"available_actions"`
	Customer interface{} `json:"customer"`
	FinancialData interface{} `json:"financial_data"`
	Requirements interface{} `json:"requirements"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	Status string `json:"status"`
	VolumeWeight float64 `json:"volume_weight"`
	Cancellation interface{} `json:"cancellation"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	TrackingNumber string `json:"tracking_number"`
	InProcessAt string `json:"in_process_at"`
	IsMultibox bool `json:"is_multibox"`
	OrderID int64 `json:"order_id"`
	Products []interface{} `json:"products"`
	ShipmentDate string `json:"shipment_date"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	AnalyticsData interface{} `json:"analytics_data"`
	DeliverySchema string `json:"delivery_schema"`
	ExternalOrder interface{} `json:"external_order"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	ParentPostingNumber string `json:"parent_posting_number"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	DestinationPlaceName string `json:"destination_place_name"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	Optional interface{} `json:"optional"`
	PrrOption string `json:"prr_option"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	Tariffication interface{} `json:"tariffication"`
	DeliveringDate string `json:"delivering_date"`
	DeliveryMethod interface{} `json:"delivery_method"`
	IsExpress bool `json:"is_express"`
	QuantumID int64 `json:"quantum_id"`
	Substatus string `json:"substatus"`
	IsPresortable bool `json:"is_presortable"`
	OrderNumber string `json:"order_number"`
	PostingNumber string `json:"posting_number"`
	Addressee interface{} `json:"addressee"`
}

type ReviewV2ReviewListV2RequestFilters struct {
	Status interface{} `json:"status"`
	OrderStatus interface{} `json:"order_status"`
	PublishedFrom string `json:"published_from"`
	PublishedTo string `json:"published_to"`
	Skus []interface{} `json:"skus"`
}

type GetProductInfoPricesResponseItemPriceIndexes struct {
	ExternalIndexData interface{} `json:"external_index_data"`
	OzonIndexData interface{} `json:"ozon_index_data"`
	SelfMarketplacesIndexData interface{} `json:"self_marketplaces_index_data"`
	ColorIndex string `json:"color_index"`
}

type V1OrderDraftValidationError struct {
	Errors []interface{} `json:"errors"`
}

type V1UpdateWarehouseFBSRequestAddressCoordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type Fbpv1DeliveryDetails struct {
	DirectDetails interface{} `json:"direct_details"`
	DropOffPoint interface{} `json:"drop_off_point"`
	PickupDetails interface{} `json:"pickup_details"`
	SupplyType interface{} `json:"supply_type"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress struct {
	ZipCode string `json:"zip_code"`
	AddressTail string `json:"address_tail"`
	Country string `json:"country"`
	District string `json:"district"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	Region string `json:"region"`
	City string `json:"city"`
	Comment string `json:"comment"`
}

type V1FbpOrderDirectCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V2FbsPostingProductCountryListResponseResult struct {
	Name string `json:"name"`
	CountryIsoCode string `json:"country_iso_code"`
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	Cursor string `json:"cursor"`
	Errors []interface{} `json:"errors"`
	HasNext bool `json:"has_next"`
}

type V2FbsPostingProductCountryListResponse struct {
	Result []interface{} `json:"result"`
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type Productv4GetProductAttributesV4ResponseResult struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	PDFList []interface{} `json:"pdf_list"`
	WeightUnit string `json:"weight_unit"`
	Attributes []interface{} `json:"attributes"`
	Barcode string `json:"barcode"`
	ColorImage string `json:"color_image"`
	ModelInfo interface{} `json:"model_info"`
	AttributesWithDefaults []interface{} `json:"attributes_with_defaults"`
	Barcodes interface{} `json:"barcodes"`
	Height int64 `json:"height"`
	Images interface{} `json:"images"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	SKU string `json:"sku"`
	Depth int64 `json:"depth"`
	DimensionUnit string `json:"dimension_unit"`
	ID int64 `json:"id"`
	PrimaryImage string `json:"primary_image"`
	TypeID int64 `json:"type_id"`
	Weight int64 `json:"weight"`
	Width int64 `json:"width"`
}

type AssemblyFbsProductListResponseProducts struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	Postings []interface{} `json:"postings"`
	ProductName string `json:"product_name"`
}

type GetReturnsListRequestFilter struct {
	LogisticReturnDate interface{} `json:"logistic_return_date"`
	StorageTarifficationStartDate interface{} `json:"storage_tariffication_start_date"`
	VisualStatusChangeMoment interface{} `json:"visual_status_change_moment"`
	OrderID int64 `json:"order_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
	VisualStatusName string `json:"visual_status_name"`
	WarehouseID int64 `json:"warehouse_id"`
	CompensationStatusID int32 `json:"compensation_status_id"`
	ProductName string `json:"product_name"`
	OfferID string `json:"offer_id"`
	Barcode string `json:"barcode"`
	ReturnSchema string `json:"return_schema"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProduct struct {
	ProductID int64 `json:"product_id"`
	Valid bool `json:"valid"`
	Error string `json:"error"`
	Exemplars []interface{} `json:"exemplars"`
}

type ChatMessageContext struct {
	OrderNumber string `json:"order_number"`
	SKU string `json:"sku"`
}

type V1AddStrategyItemsResponseResult struct {
	Errors []interface{} `json:"errors"`
	FailedProductCount int32 `json:"failed_product_count"`
}

type V1UpdateWarehouseFBSFirstMileResponse struct {
	OperationID string `json:"operation_id"`
}

type V2DeliveryMethodListV2Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type V1SellerActionsCreateMultiLevelDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type PostingV4PostingFbsListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1AssemblyFbsProductListRequestSortDirEnum string

type V1GetAttributeValuesResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V1GetProductInfoSubscriptionResponseResult struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1FbpOrderDropOffTimetableResponseCalendar struct {
	CalendarItem interface{} `json:"calendar_item"`
	DayOfWeek interface{} `json:"day_of_week"`
}

type V3ChatInfo struct {
	Chat interface{} `json:"chat"`
	FirstUnreadMessageID int64 `json:"first_unread_message_id"`
	LastMessageID int64 `json:"last_message_id"`
	UnreadCount int64 `json:"unread_count"`
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequestActionParameters struct {
	DiscountLevels []interface{} `json:"discount_levels"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
}

type ProductGetImportProductsInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1CreateStockByWarehouseReportRequest struct {
	Language interface{} `json:"language"`
	WarehouseId []interface{} `json:"warehouseId"`
}

type V1FbpDraftDirectTplDlvCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError struct {
	SKU int64 `json:"sku"`
	Errors []interface{} `json:"errors"`
}

type V2ProductInfoPicturesResponseError struct {
	Message string `json:"message"`
	URL string `json:"url"`
}

type V1TimeRangeVisualStatus struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V3ImportProductsRequestAttribute struct {
	ID int64 `json:"id"`
	Values []interface{} `json:"values"`
	ComplexID int64 `json:"complex_id"`
}

type RatingValueCurrent struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Formatted string `json:"formatted"`
	Status interface{} `json:"status"`
	Value float64 `json:"value"`
}

type V1ProductPricesDetailsResponse struct {
	Prices []interface{} `json:"prices"`
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"`
}

type CarriageCarriageGetResponseCancelAvailability struct {
	IsCancelAvailable bool `json:"is_cancel_available"`
	Reason string `json:"reason"`
}

type V1SellerActionsProductsListResponse struct {
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
	Cursor int64 `json:"cursor"`
}

type MarketingAction struct {
	Value int32 `json:"value"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Title string `json:"title"`
}

type ChatInfoChatStatusEnum string

type GetFinanceBalanceV1ResponseTotal struct {
	Accrued interface{} `json:"accrued"`
	ClosingBalance interface{} `json:"closing_balance"`
	OpeningBalance interface{} `json:"opening_balance"`
	Payments []interface{} `json:"payments"`
}

type DeliveryMethodListV2RequestSortDirEnum string

type Productv5GetProductListRequestFilterFilterVisibility string

type SellerApiActivateProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type PostingV4PostingFbsUnfulfilledListRequestSortDirEnum string

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type GetReturnsListResponseProduct struct {
	Commission interface{} `json:"commission"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	Name string `json:"name"`
	Price interface{} `json:"price"`
	PriceWithoutCommission interface{} `json:"price_without_commission"`
	CommissionPercent float64 `json:"commission_percent"`
}

type ReviewV2ReviewListV2ResponseReview struct {
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
	VideosAmount int32 `json:"videos_amount"`
	ID string `json:"id"`
	PhotosAmount int32 `json:"photos_amount"`
	PublishedAt string `json:"published_at"`
	CommentsAmount int32 `json:"comments_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	OrderStatus interface{} `json:"order_status"`
	Rating int32 `json:"rating"`
}

type PriceIndexesIndexOzonData struct {
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
	MinPrice float64 `json:"min_price"`
}

type ExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V2ReturnsRfbsListRequest struct {
	Filter interface{} `json:"filter"`
	LastID int32 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type V2FboSinglePostingLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type DetailsRfbsDetails struct {
	Total float64 `json:"total"`
	TransferDelivery float64 `json:"transfer_delivery"`
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"`
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`
	PartialCompensation float64 `json:"partial_compensation"`
	PartialCompensationReturn float64 `json:"partial_compensation_return"`
}

type V1WarehouseListRequest struct {
	With interface{} `json:"with"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type V2PostingFBSActGetPostingsRequest struct {
	ID interface{} `json:"id"`
}

type V1GetFinanceBalanceV1ResponseSalesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type MoneyMoneyCustomerPrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualAccruedCategoryEnum string

type PolygonBindRequestpolygon struct {
	PolygonID int64 `json:"polygon_id"`
	Time int64 `json:"time"`
}

type ProductExemplar struct {
	ExemplarID int64 `json:"exemplar_id"`
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
}

type V2FbsPosting struct {
	OrderNumber string `json:"order_number"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	ShipmentDate string `json:"shipment_date"`
	Status string `json:"status"`
	Barcodes interface{} `json:"barcodes"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CreatedAt string `json:"created_at"`
	InProcessAt string `json:"in_process_at"`
	OrderID int64 `json:"order_id"`
}

type GetUploadQuotaResponseDailyUpdate struct {
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
	Usage int64 `json:"usage"`
}

type V1UpdateWarehouseFBSRequest struct {
	AddressCoordinates interface{} `json:"address_coordinates"`
	Name string `json:"name"`
	Options interface{} `json:"options"`
	Phone string `json:"phone"`
	WarehouseID int64 `json:"warehouse_id"`
	WorkingDays []interface{} `json:"working_days"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDelivery struct {
	Services []interface{} `json:"services"`
	TotalAccrued interface{} `json:"total_accrued"`
}

type V1Barcode struct {
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type SellerApiProductV1ResponseProduct struct {
	ProductID float64 `json:"product_id"`
	Reason string `json:"reason"`
}

type V1Empty interface{}

type Postingv1GetCarriageAvailableListResponse struct {
	Result interface{} `json:"result"`
}

type V1GetFinanceBalanceV1ResponsePaymentsMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1FbpDraftListResponseItem struct {
	PackageUnitsCount int32 `json:"package_units_count"`
	Status interface{} `json:"status"`
	SupplyID string `json:"supply_id"`
	WarehouseID int64 `json:"warehouse_id"`
	CancellationState interface{} `json:"cancellation_state"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	DeliveryDetails interface{} `json:"delivery_details"`
	Editable bool `json:"editable"`
	IsCancelable bool `json:"is_cancelable"`
	IsDeletable bool `json:"is_deletable"`
	Locked bool `json:"locked"`
	BundleID string `json:"bundle_id"`
	ID int64 `json:"id"`
}

type ActionsV1ActionsAutoAddProductsListResponseProducts struct {
	MinActionQuantity int64 `json:"min_action_quantity"`
	Name string `json:"name"`
	AddMode string `json:"add_mode"`
	Currency string `json:"currency"`
	MinSellerPrice float64 `json:"min_seller_price"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
	SKU int64 `json:"sku"`
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
}

type V1SellerActionsProductsListResponseProductQuantTypeEnum string

type V1SellerActionsChangeActivityRequest struct {
	ActionID int64 `json:"action_id"`
	IsTurnOn bool `json:"is_turn_on"`
}

type FbpWarehouseListResponseAddressDetailing struct {
	City string `json:"city"`
	Country string `json:"country"`
	House string `json:"house"`
	Region string `json:"region"`
	Street string `json:"street"`
	Zipcode string `json:"zipcode"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProducts struct {
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
	Commission interface{} `json:"commission"`
	CustomerPrice interface{} `json:"customer_price"`
	OldPrice float64 `json:"old_price"`
	Payout float64 `json:"payout"`
}

type Postingv1GetCarriageAvailableListRequest struct {
	DepartureDate string `json:"departure_date"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type PostingV1PostingFbpListResponsePostings struct {
	Products []interface{} `json:"products"`
	Status string `json:"status"`
	FinancialData interface{} `json:"financial_data"`
	InProcessAt string `json:"in_process_at"`
	OrderDate string `json:"order_date"`
	OrderID int64 `json:"order_id"`
	OrderNumber string `json:"order_number"`
	PostingNumber string `json:"posting_number"`
	ProviderID int64 `json:"provider_id"`
}

type ErrorErrorLevel string

type V1CarriageApproveRequest struct {
	ContainersCount int32 `json:"containers_count"`
	CarriageID int64 `json:"carriage_id"`
}

type HumanTextsParam struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type V1Money struct {
	Currency string `json:"currency"`
	Value float64 `json:"value"`
}

type V1FbpDraftDirectProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1GetStrategyItemInfoRequest struct {
	ProductID int64 `json:"product_id"`
}

type WarehouseTimetable struct {
	TimetableTo string `json:"timetable_to"`
	WorkingHours []interface{} `json:"working_hours"`
	TimetableFrom string `json:"timetable_from"`
}

type V1AssemblyCarriagePostingListResponsePostingProduct struct {
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
}

type V1CreatePricingStrategyResponseResult struct {
	StrategyID string `json:"strategy_id"`
}

type PostingBooleanResponse struct {
	Result bool `json:"result"`
}

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"`
}

type V1SearchQueriesTextRequest struct {
	Offset string `json:"offset"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	Text string `json:"text"`
	Limit string `json:"limit"`
}

type V1SellerActionsUpdateVoucherRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1CreateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseExtremelyLowPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type SellerApiProduct struct {
	AlertMaxActionPrice float64 `json:"alert_max_action_price"`
	AddMode string `json:"add_mode"`
	MinStock float64 `json:"min_stock"`
	PriceMinElastic float64 `json:"price_min_elastic"`
	MinBoost float64 `json:"min_boost"`
	ActionPrice float64 `json:"action_price"`
	MaxActionPrice float64 `json:"max_action_price"`
	Stock float64 `json:"stock"`
	CurrentBoost float64 `json:"current_boost"`
	PriceMaxElastic float64 `json:"price_max_elastic"`
	MaxBoost float64 `json:"max_boost"`
	ID float64 `json:"id"`
	Price float64 `json:"price"`
	AlertMaxActionPriceFailed bool `json:"alert_max_action_price_failed"`
}

type WarehouseWorkingDaysEnum string

type V2ProductInfoPicturesResponseItem struct {
	Errors []interface{} `json:"errors"`
	ProductID int64 `json:"product_id"`
	PrimaryPhoto []interface{} `json:"primary_photo"`
	Photo []interface{} `json:"photo"`
	ColorPhoto []interface{} `json:"color_photo"`
	Photo360 []interface{} `json:"photo_360"`
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type SellerApiProductV1Response struct {
	Result interface{} `json:"result"`
}

type V1WarehouseInvalidProductsGetRequest struct {
	LastID int64 `json:"last_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"`
}

type Productv3GetProductAttributesV3Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type V1SetPostingsRequest struct {
	CarriageID int64 `json:"carriage_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1AddStrategyItemsRequest struct {
	ProductID []interface{} `json:"product_id"`
	StrategyID string `json:"strategy_id"`
}

type V4GetProductInfoStocksResponse struct {
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
	Cursor string `json:"cursor"`
}

type ReturnsCompanyFbsInfoRequestPagination struct {
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualNonItemFee struct {
	Accrued interface{} `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type V1OrderValidationErrorErrorType string

type OperationItem struct {
	Name string `json:"name"`
	SKU int64 `json:"sku"`
}

type ActionParameter struct {
	AutoStopActionReason interface{} `json:"auto_stop_action_reason"`
	BudgetSpent float64 `json:"budget_spent"`
	DiscountLevels []interface{} `json:"discount_levels"`
	PickedSegments []interface{} `json:"picked_segments"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	Status interface{} `json:"status"`
	Type interface{} `json:"type_"`
	Addresses []interface{} `json:"addresses"`
	Budget float64 `json:"budget"`
	DiscountType interface{} `json:"discount_type"`
	MinOrderAmount float64 `json:"min_order_amount"`
	VoucherParameters interface{} `json:"voucher_parameters"`
	Warehouses []interface{} `json:"warehouses"`
	DiscountValue float64 `json:"discount_value"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	MinActionPercent float64 `json:"min_action_percent"`
	Title string `json:"title"`
}

type V3ImportProductsResponse struct {
	Result interface{} `json:"result"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
}

type FinanceTransactionListV3RequestFilter struct {
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
	Date interface{} `json:"date"`
	OperationType []interface{} `json:"operation_type"`
}

type GetConditionalCancellationListV2RequestWith struct {
	Counter bool `json:"counter"`
}

type GetProductInfoListResponseAvailability struct {
	Availability string `json:"availability"`
	Reasons []interface{} `json:"reasons"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type V1GetStrategyIDsByItemIDsResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpDraftDropOffDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type V1Competitor struct {
	Coefficient float64 `json:"coefficient"`
	CompetitorID int64 `json:"competitor_id"`
}

type V1FbpDraftDropOffProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1CarriageCreateRequest struct {
	DepartureDate string `json:"departure_date"`
	AllBlrTraceable bool `json:"all_blr_traceable"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1GetProductInfoSubscriptionResponse struct {
	Result []interface{} `json:"result"`
}

type RolesByTokenResponseRoles struct {
	Name string `json:"name"`
	Methods []interface{} `json:"methods"`
}

type V1SellerActionsUpdateVoucherRequestActionParameters struct {
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountValue float64 `json:"discount_value"`
}

type V1ApproveDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type SellerApiGetSellerProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Limit float64 `json:"limit"`
	LastID float64 `json:"last_id"`
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []interface{} `json:"provinces"`
}

type FbsPostingTrackingNumberSetRequestTrackingNumber struct {
	PostingNumber string `json:"posting_number"`
	TrackingNumber string `json:"tracking_number"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type ParameterPickedSegment struct {
	Segments []interface{} `json:"segments"`
}

type V1GetProductInfoDiscountedRequest struct {
	DiscountedSkus interface{} `json:"discounted_skus"`
}

type AnalyticsProductQueriesDetailsResponseQuery struct {
	Currency string `json:"currency"`
	Position float64 `json:"position"`
	SKU int64 `json:"sku"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
	UniqueViewUsers int64 `json:"unique_view_users"`
	Gmv float64 `json:"gmv"`
	OrderCount int64 `json:"order_count"`
	Query string `json:"query"`
	QueryIndex int64 `json:"query_index"`
	ViewConversion float64 `json:"view_conversion"`
}

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementEnum string

type V1SellerActionsProductsListResponseProduct struct {
	BasePrice float64 `json:"base_price"`
	IsActive bool `json:"is_active"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantSize int64 `json:"quant_size"`
	SKU []interface{} `json:"sku"`
	ActionPrice float64 `json:"action_price"`
	Currency string `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	MinSellerPrice float64 `json:"min_seller_price"`
	QuantType interface{} `json:"quant_type"`
}

type V1FbpArchiveListRequest struct {
	LastID string `json:"last_id"`
	Count string `json:"count"`
}

type V1GetCompetitorsRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V3FbsPostingProduct struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	CurrencyCode string `json:"currency_code"`
	Imei []interface{} `json:"imei"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
}

type ProductCurrencyEnum string

type V2ProductInfoPicturesRequest struct {
	ProductID interface{} `json:"product_id"`
}

type V1DiscountTaskStatus string

type V3FinanceCashFlowStatementListResponseResult struct {
	CashFlows interface{} `json:"cash_flows"`
	Details []interface{} `json:"details"`
	PageCount int64 `json:"page_count"`
}

type V3GetProductInfoListResponseItem struct {
	IsSuper bool `json:"is_super"`
	Name string `json:"name"`
	ColorImage []interface{} `json:"color_image"`
	Commissions []interface{} `json:"commissions"`
	Images360 []interface{} `json:"images360"`
	IsAutoarchived bool `json:"is_autoarchived"`
	IsPrepaymentAllowed bool `json:"is_prepayment_allowed"`
	VolumeWeight float64 `json:"volume_weight"`
	DiscountedFBOStocks int32 `json:"discounted_fbo_stocks"`
	IsDiscounted bool `json:"is_discounted"`
	Price string `json:"price"`
	SKU int64 `json:"sku"`
	Sources []interface{} `json:"sources"`
	Stocks interface{} `json:"stocks"`
	VisibilityDetails interface{} `json:"visibility_details"`
	Availabilities []interface{} `json:"availabilities"`
	ID int64 `json:"id"`
	Images []interface{} `json:"images"`
	MinPrice string `json:"min_price"`
	HasDiscountedFBOItem bool `json:"has_discounted_fbo_item"`
	IsKgt bool `json:"is_kgt"`
	PriceIndexes interface{} `json:"price_indexes"`
	Vat string `json:"vat"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	IsArchived bool `json:"is_archived"`
	ModelInfo interface{} `json:"model_info"`
	OldPrice string `json:"old_price"`
	TypeID int64 `json:"type_id"`
	Barcodes []interface{} `json:"barcodes"`
	CurrencyCode string `json:"currency_code"`
	Errors []interface{} `json:"errors"`
	Promotions []interface{} `json:"promotions"`
	Statuses interface{} `json:"statuses"`
	CreatedAt string `json:"created_at"`
	OfferID string `json:"offer_id"`
	PrimaryImage []interface{} `json:"primary_image"`
	UpdatedAt string `json:"updated_at"`
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date interface{} `json:"date"`
}

type ProductBooleanResponse struct {
	Result bool `json:"result"`
}

type V1ProductUpdateAttributesRequest struct {
	Items interface{} `json:"items"`
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type ReviewV2ReviewListV2Response struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
}

type V1AssemblyFbsProductListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	SortDir interface{} `json:"sort_dir"`
}

type V1GetProductInfoDiscountedResponseItem struct {
	SKU int64 `json:"sku"`
	WarrantyType string `json:"warranty_type"`
	Condition string `json:"condition"`
	PackageDamage string `json:"package_damage"`
	PackagingViolation string `json:"packaging_violation"`
	Repair string `json:"repair"`
	Shortage string `json:"shortage"`
	CommentReasonDamaged string `json:"comment_reason_damaged"`
	ConditionEstimation string `json:"condition_estimation"`
	Defects string `json:"defects"`
	DiscountedSKU int64 `json:"discounted_sku"`
	MechanicalDamage string `json:"mechanical_damage"`
	ReasonDamaged string `json:"reason_damaged"`
}

type PriceIndexesIndexExternalData struct {
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
	MinPrice float64 `json:"min_price"`
}

type CreatedAt struct {
	From string `json:"from"`
	To string `json:"to"`
}

type GetFinanceBalanceV1ResponseReturns struct {
	AmountDetails interface{} `json:"amount_details"`
	Fee interface{} `json:"fee"`
	Amount interface{} `json:"amount"`
}

type PostingPostingFBSPackageLabelResponse struct {
	FileContent string `json:"file_content"`
	FileName string `json:"file_name"`
	ContentType string `json:"content_type"`
}

type V1SellerActionsListResponse struct {
	Actions []interface{} `json:"actions"`
	Total int64 `json:"total"`
}

type ValidationErrorCharacteristicEnum string

type ProductV1ProductVisibilityInfoResponseItem struct {
	ShowcasesVisibility interface{} `json:"showcases_visibility"`
	SKU int64 `json:"sku"`
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type SellerApiGetSellerProductV1Response struct {
	Result interface{} `json:"result"`
}

type V1ProductInfoWrongVolumeResponse struct {
	Cursor string `json:"cursor"`
	Products interface{} `json:"products"`
}

type V2GetDiscountTaskListV2Response struct {
	Tasks []interface{} `json:"tasks"`
}

type GetImportProductsInfoResponseResultItem struct {
	ProductID int64 `json:"product_id"`
	Status string `json:"status"`
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
}

type AnalyticsProductQueriesRequestSortDir string

type V3FbsPostingExemplarProductV3 struct {
	Exemplars []interface{} `json:"exemplars"`
	SKU int64 `json:"sku"`
}

type V1GetAttributesResponseAttribute struct {
	GroupName string `json:"group_name"`
	ID int64 `json:"id"`
	IsRequired bool `json:"is_required"`
	AttributeComplexID int64 `json:"attribute_complex_id"`
	IsAspect bool `json:"is_aspect"`
	IsCollection bool `json:"is_collection"`
	Name string `json:"name"`
	Type string `json:"type_"`
	MaxValueCount int64 `json:"max_value_count"`
	ComplexIsCollection bool `json:"complex_is_collection"`
	CategoryDependent bool `json:"category_dependent"`
	Description string `json:"description"`
	DictionaryID int64 `json:"dictionary_id"`
	GroupID int64 `json:"group_id"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission struct {
	Percent int64 `json:"percent"`
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
}

type V1TimeRangeReturnDate struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type AnalyticsAnalyticsGetDataRequest struct {
	Metrics []interface{} `json:"metrics"`
	Offset int64 `json:"offset"`
	Sort []interface{} `json:"sort"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Dimension []interface{} `json:"dimension"`
	Filters []interface{} `json:"filters"`
	Limit int64 `json:"limit"`
}

type GetWarehouseFBSOperationStatusResponseResult struct {
	EntityID int64 `json:"entity_id"`
}

type ArrivalpassArrivalPassListResponse struct {
	Cursor string `json:"cursor"`
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type V1SellerActionsProductsCandidatesResponseProduct struct {
	Currency string `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	ActionPrice float64 `json:"action_price"`
	IsActive bool `json:"is_active"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	QuantSize int64 `json:"quant_size"`
	QuantType interface{} `json:"quant_type"`
	SKU []interface{} `json:"sku"`
	BasePrice float64 `json:"base_price"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearch struct {
	Address string `json:"address"`
	Types []interface{} `json:"types"`
}

type V3ChatDetailsInfo struct {
	ChatID string `json:"chat_id"`
	ChatStatus interface{} `json:"chat_status"`
	ChatType interface{} `json:"chat_type"`
	CreatedAt string `json:"created_at"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep struct {
	MinCharge interface{} `json:"min_charge"`
	TariffCharge interface{} `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"`
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
}

type SellerInfoResponseSubscription struct {
	IsPremium bool `json:"is_premium"`
	Type interface{} `json:"type_"`
}

type V1GetFBSRatingIndexInfoV1Response struct {
	PeriodTo string `json:"period_to"`
	ProcessingCostsSum float64 `json:"processing_costs_sum"`
	CurrencyCode string `json:"currency_code"`
	Defects []interface{} `json:"defects"`
	Index float64 `json:"index"`
	PeriodFrom string `json:"period_from"`
}

type V1GetFinanceBalanceV1ResponseClosingBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type Productv1ProductImportPicturesRequest struct {
	ColorImage string `json:"color_image"`
	Images interface{} `json:"images"`
	Images360 interface{} `json:"images360"`
	ProductID int64 `json:"product_id"`
}

type V1UpdatePricingStrategyRequest struct {
	StrategyID string `json:"strategy_id"`
	StrategyName string `json:"strategy_name"`
	Competitors []interface{} `json:"competitors"`
}

type V1SellerOzonLogisticsInfoResponse struct {
	AvailableSchemas []interface{} `json:"available_schemas"`
	OzonLogisticsEnabled bool `json:"ozon_logistics_enabled"`
}

type PostingV4PostingFbsUnfulfilledListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
}

type V2WarehouseListV2Response struct {
	Cursor string `json:"cursor"`
	Warehouses []interface{} `json:"warehouses"`
	HasNext bool `json:"has_next"`
}

type FbpGetLabelResponseLabelCreationStateTypeEnum string

type GetFinanceBalanceV1ResponseReturnsDetails struct {
	Revenue interface{} `json:"revenue"`
	PartnerPrograms interface{} `json:"partner_programs"`
	PointsForDiscounts string `json:"points_for_discounts"`
}

type WarehouseCarriageLabelTypeEnum string

type ProductGetImportProductsInfoRequest struct {
	TaskID int64 `json:"task_id"`
}

type V3Address struct {
	Country string `json:"country"`
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	Region string `json:"region"`
	City string `json:"city"`
	Comment string `json:"comment"`
	District string `json:"district"`
	Latitude float64 `json:"latitude"`
	ZipCode string `json:"zip_code"`
	AddressTail string `json:"address_tail"`
}

type Financev3FinanceTransactionTotalsV3Request struct {
	Date interface{} `json:"date"`
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
}

type GetProductAttributesResponseImage360 struct {
	Index int64 `json:"index"`
	FileName string `json:"file_name"`
}

type SellerReturnsv1MoneyCommission struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type GetRealizationReportByDayResponse struct {
	Rows []interface{} `json:"rows"`
}

type V1GenerateBarcodeRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1ArchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V1WarehouseFbsCreateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type FinanceV1GetFinanceAccrualTypesResponse struct {
	AccrualTypes []interface{} `json:"accrual_types"`
}

type V1DeliveryDetailsPickUpDetails struct {
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type V1AnalyticsProductQueriesResponseItem struct {
	Category string `json:"category"`
	Currency string `json:"currency"`
	Name string `json:"name"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
	UniqueViewUsers int64 `json:"unique_view_users"`
	ViewConversion float64 `json:"view_conversion"`
	Gmv float64 `json:"gmv"`
	OfferID string `json:"offer_id"`
	Position float64 `json:"position"`
	SKU int64 `json:"sku"`
}

type V3ImportProductsRequest struct {
	Items []interface{} `json:"items"`
}

type GetCompetitorsResponseCompetitorInfo struct {
	Name string `json:"name"`
	ID int64 `json:"id"`
}

type V3FbsPostingDetail struct {
	AnalyticsData interface{} `json:"analytics_data"`
	DeliveryPrice string `json:"delivery_price"`
	Optional interface{} `json:"optional"`
	TrackingNumber string `json:"tracking_number"`
	ShipmentDate string `json:"shipment_date"`
	Addressee interface{} `json:"addressee"`
	Barcodes interface{} `json:"barcodes"`
	IsExpress bool `json:"is_express"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	PostingNumber string `json:"posting_number"`
	ProductExemplars interface{} `json:"product_exemplars"`
	ParentPostingNumber string `json:"parent_posting_number"`
	Substatus string `json:"substatus"`
	AvailableActions interface{} `json:"available_actions"`
	Cancellation interface{} `json:"cancellation"`
	InProcessAt string `json:"in_process_at"`
	OrderNumber string `json:"order_number"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	Courier interface{} `json:"courier"`
	Customer interface{} `json:"customer"`
	DeliveryMethod interface{} `json:"delivery_method"`
	FactDeliveryDate string `json:"fact_delivery_date"`
	FinancialData interface{} `json:"financial_data"`
	LegalInfo interface{} `json:"legal_info"`
	OrderID int64 `json:"order_id"`
	ProviderStatus string `json:"provider_status"`
	RelatedPostings interface{} `json:"related_postings"`
	AdditionalData []interface{} `json:"additional_data"`
	DeliveringDate string `json:"delivering_date"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	Requirements interface{} `json:"requirements"`
	PreviousSubstatus string `json:"previous_substatus"`
	Tariffication interface{} `json:"tariffication"`
}

type V5FbsPostingProductExemplarValidateV5Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type PostingProductCancelRequestItem struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V4FbsPostingShipPackageV4RequestProduct struct {
	ExemplarsIds []interface{} `json:"exemplarsIds"`
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

type V1AddBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type V1GetReturnsListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
	LastID int64 `json:"last_id"`
}

type ArrivalpassArrivalPassListResponseArrivalPass struct {
	DriverPhone string `json:"driver_phone"`
	DropoffPointID int64 `json:"dropoff_point_id"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WarehouseID int64 `json:"warehouse_id"`
	ArrivalPassID int64 `json:"arrival_pass_id"`
	ArrivalReasons []interface{} `json:"arrival_reasons"`
	DriverName string `json:"driver_name"`
	IsActive bool `json:"is_active"`
	ArrivalTime string `json:"arrival_time"`
}

type PickedSegmentSegmentTypeEnum string

type V2GetDiscountTaskListV2Request struct {
	LastID int64 `json:"last_id"`
	Limit int64 `json:"limit"`
	Status interface{} `json:"status"`
}

type V1AddBarcodeResult struct {
	Code string `json:"code"`
	Error string `json:"error"`
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type SellerActionsListRequestStatusEnum string

type ReportListRequestReportType string

type V1CommentDeleteRequest struct {
	CommentID string `json:"comment_id"`
}

type PostingCancelReason struct {
	IsAvailableForCancellation bool `json:"is_available_for_cancellation"`
	Title string `json:"title"`
	TypeID string `json:"type_id"`
	ID int64 `json:"id"`
}

type V1GetFinanceBalanceV1Response struct {
	Cashflows interface{} `json:"cashflows"`
	Total interface{} `json:"total"`
}

type Productv2DeleteProductsRequest struct {
	Products []interface{} `json:"products"`
}

type V1GetStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type ReportCreateCompanyPostingsReportRequestFilter struct {
	SKU []interface{} `json:"sku"`
	StatusAlias []interface{} `json:"status_alias"`
	WarehouseID []interface{} `json:"warehouse_id"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	ProcessedAtTo string `json:"processed_at_to"`
	Statuses []interface{} `json:"statuses"`
	Title string `json:"title"`
	IsExpress interface{} `json:"is_express"`
	CancelReasonID []interface{} `json:"cancel_reason_id"`
	DeliverySchema []interface{} `json:"delivery_schema"`
	OfferID string `json:"offer_id"`
	ProcessedAtFrom string `json:"processed_at_from"`
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"`
}

type V1FbpArchiveListResponseItem struct {
	WhcOrderID int64 `json:"whc_order_id"`
	ActFileUuid string `json:"act_file_uuid"`
	CreatedDate string `json:"created_date"`
	ExternalOrderID string `json:"external_order_id"`
	HasAct bool `json:"has_act"`
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	ReceiveDate string `json:"receive_date"`
	BundleID string `json:"bundle_id"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
	DeliveryDetails interface{} `json:"delivery_details"`
	HasLabel bool `json:"has_label"`
	RowVersion int64 `json:"row_version"`
	DeclineReason interface{} `json:"decline_reason"`
	OrderDraftID int64 `json:"order_draft_id"`
	PackageUnitsCount int32 `json:"package_units_count"`
	SupplyID string `json:"supply_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation struct {
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPosting struct {
	DeliverySchema string `json:"delivery_schema"`
	DeliverySpeed int32 `json:"delivery_speed"`
	Products []interface{} `json:"products"`
}

type ProductsPostings struct {
	PostingNumber string `json:"posting_number"`
	Quantity int32 `json:"quantity"`
}

type V1SellerActionsCreateDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason string

type PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo struct {
	Kpp string `json:"kpp"`
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
}

type MoneyMoneySellerPrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type SellerApiProductV1ResponseDeactivate struct {
	Result interface{} `json:"result"`
}

type FbpCheckActStateResponseErrorReason string

type V1FbpDraftDirectTimeslotEditResponseReserveFailureType string

type V1RatingStatus struct {
	Warning bool `json:"warning"`
	Danger bool `json:"danger"`
	Premium bool `json:"premium"`
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"`
}

type V3FbsPostingAnalyticsData struct {
	DeliveryType string `json:"delivery_type"`
	Region string `json:"region"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	City string `json:"city"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbsPostingProductExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairwayStep struct {
	Discount int64 `json:"discount"`
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
}

type V5FbsPostingProductExemplarValidateV5RequestProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type V1FbpDraftPickupCreateRequestDeliveryDetails struct {
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type V1GetAttributesResponse struct {
	Result []interface{} `json:"result"`
}

type ReviewV2ReviewListV2Request struct {
	SortDir interface{} `json:"sort_dir"`
	Filters interface{} `json:"filters"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
}

type GetUploadQuotaResponseTotal struct {
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
	QuotaByCategory interface{} `json:"quota_by_category"`
}

type V2GetProductInfoStocksByWarehouseFbsRequestV2 struct {
	SKU []interface{} `json:"sku"`
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	OfferID []interface{} `json:"offer_id"`
}

type V3AddresseeFbsLists struct {
	Phone string `json:"phone"`
	Name string `json:"name"`
}

type GetProductInfoStocksResponseStock struct {
	ShipmentType interface{} `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Type string `json:"type_"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	Present int32 `json:"present"`
	Reserved int32 `json:"reserved"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseFailedPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type V1FbpCheckActStateResponseStatus string

type ReviewV2ReviewListV2ResponseReviewOrderStatusEnum string

type V6FbsPostingProductExemplarCreateOrGetV6Request struct {
	PostingNumber string `json:"posting_number"`
}

type V2DeliveryMethodListV2Response struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	DeliveryMethods []interface{} `json:"delivery_methods"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilter struct {
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	Statuses []interface{} `json:"statuses"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	ProviderIds []interface{} `json:"provider_ids"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	CutoffFrom string `json:"cutoff_from"`
	DeliveringDateFrom string `json:"delivering_date_from"`
	DeliveringDateTo string `json:"delivering_date_to"`
}

type GetProductInfoListResponseCommission struct {
	SaleSchema string `json:"sale_schema"`
	Value float64 `json:"value"`
	DeliveryAmount float64 `json:"delivery_amount"`
	Percent float64 `json:"percent"`
	ReturnAmount float64 `json:"return_amount"`
}

type V1AssemblyFbsPostingListRequestSortDirEnum string

type V1FbpDraftDirectSellerDlvEditRequest struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type GetRealizationReportResponseV2Row struct {
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission interface{} `json:"delivery_commission"`
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
}

type V1UpdateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type PostingV1PostingFbpListRequestSortDirEnum string

type V1SearchAttributeValuesRequest struct {
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	Value string `json:"value"`
}

type RowItemLegalEntityDocument struct {
	SaleDate string `json:"sale_date"`
	Number string `json:"number"`
}

type SellerReturnsv1MoneyWithoutCommission struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type V1GetFinanceBalanceV1ResponseRevenueMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type SellerApiProductV1ResponseProductDeactivate struct {
	ProductID float64 `json:"product_id"`
	Reason string `json:"reason"`
}

type Financev3FinanceTransactionListV3Request struct {
	Filter interface{} `json:"filter"`
	Page int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type ReportReportinfo struct {
	Status string `json:"status"`
	Code string `json:"code"`
	CreatedAt string `json:"created_at"`
	Error string `json:"error"`
	ExpiresAt string `json:"expires_at"`
	File string `json:"file"`
	Params map[string]interface{} `json:"params"`
	ReportType string `json:"report_type"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsRequest struct {
	OfferID interface{} `json:"offer_id"`
	SKU interface{} `json:"sku"`
}

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V3FbsPostingProductExemplarsV3 struct {
	Products []interface{} `json:"products"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejectedCodeEnum string

type PriceIndexesIndexDataExternal struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type AnalyticsFilter struct {
	Value string `json:"value"`
	Key string `json:"key"`
	Op interface{} `json:"op"`
}

type V1ProductUpdateDiscountRequest struct {
	Discount int32 `json:"discount"`
	ProductID int64 `json:"product_id"`
}

type ProductGetImportProductsInfoResponseResult struct {
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type FinanceV1GetFinanceAccrualPostingsResponse struct {
	PostingAccruals []interface{} `json:"posting_accruals"`
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"`
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type DirectDetailsTimeslotDetails struct {
	Timeslot interface{} `json:"timeslot"`
	TimeslotReservationID string `json:"timeslot_reservation_id"`
}

type V1SellerActionsProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	Skus []interface{} `json:"skus"`
}

type V1ArchiveDeclineReason struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1ArchiveSkuSummary struct {
	TotalQuantity int64 `json:"total_quantity"`
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemsCount int64 `json:"total_items_count"`
}

type ProductImportProductsPricesResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1QuestionTopSkuRequest struct {
	Limit int64 `json:"limit"`
}

type V1FbpWarehouseListResponse struct {
	Warehouses []interface{} `json:"warehouses"`
}

type ProductV1ProductVisibilitySetRequestItemPlacement struct {
	Placement interface{} `json:"placement"`
	SKU int64 `json:"sku"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequest struct {
	Coordinates interface{} `json:"coordinates"`
	CountryCode string `json:"country_code"`
	IsKgt bool `json:"is_kgt"`
	Search interface{} `json:"search"`
}

type V1ArchiveStatusEnum string

type RelatedPostingCancelReason struct {
	Reasons []interface{} `json:"reasons"`
	PostingNumber string `json:"posting_number"`
}

type V1CommentCreateRequest struct {
	ReviewID string `json:"review_id"`
	Text string `json:"text"`
	MarkReviewAsProcessed bool `json:"mark_review_as_processed"`
	ParentCommentID string `json:"parent_comment_id"`
}

type GetConditionalCancellationListV2ResponseResult struct {
	State interface{} `json:"state"`
	CancellationID int64 `json:"cancellation_id"`
	CancellationReasonMessage string `json:"cancellation_reason_message"`
	PostingNumber string `json:"posting_number"`
	SourceID int64 `json:"source_id"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	ApproveComment string `json:"approve_comment"`
	ApproveDate string `json:"approve_date"`
	AutoApproveDate string `json:"auto_approve_date"`
	CancellationInitiator interface{} `json:"cancellation_initiator"`
	CancellationReason interface{} `json:"cancellation_reason"`
	CancelledAt string `json:"cancelled_at"`
	OrderDate string `json:"order_date"`
}

type PostingV4PostingFbsListRequest struct {
	Translit bool `json:"translit"`
	With interface{} `json:"with"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type ActionParameterVoucherParameterTypeEnum string

type V1FbpCheckConsignmentNoteStateRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFee struct {
	Fees []interface{} `json:"fees"`
	SKU int64 `json:"sku"`
}

type FbsPostingDetailCourier struct {
	CarModel string `json:"car_model"`
	CarNumber string `json:"car_number"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V1SellerActionsCreateVoucherRequestVoucherParameter struct {
	CountCodes int64 `json:"count_codes"`
	IsPrivate bool `json:"is_private"`
	Type interface{} `json:"type_"`
}

type V1GenerateBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type FbsPostingBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type RowItemOrder struct {
	PostingNumber string `json:"posting_number"`
	CreatedDate string `json:"created_date"`
}

type ReportListResponseResult struct {
	Reports []interface{} `json:"reports"`
	Total int32 `json:"total"`
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1FbpDraftDirectProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type Postingv3GetFbsPostingUnfulfilledListRequestFilter struct {
	DeliveringDateTo string `json:"delivering_date_to"`
	Status string `json:"status"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	FbpFilter string `json:"fbpFilter"`
	ProviderID []interface{} `json:"provider_id"`
	WarehouseID []interface{} `json:"warehouse_id"`
	CutoffFrom string `json:"cutoff_from"`
	DeliveringDateFrom string `json:"delivering_date_from"`
}

type V3FbsPostingProductExemplarInfoV3 struct {
	Imei []interface{} `json:"imei"`
	ExemplarID int64 `json:"exemplar_id"`
	MandatoryMark string `json:"mandatory_mark"`
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	Rnpt string `json:"rnpt"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
}

type RpcStatusV1PolygonBind struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1AddStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type GetWarehouseFBSOperationStatusResponseStatusEnum string

type FbsPostingShipV4RequestPackageProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

type FbsPostingMoveStatusResponseMoveStatus struct {
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
	Result bool `json:"result"`
}

type PostingPostingFBSActGetContainerLabelsResponse struct {
	FileContent string `json:"file_content"`
	FileName string `json:"file_name"`
	ContentType string `json:"content_type"`
}

type OperationService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type PostingV4PostingFbsListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1GetDiscountTaskListRequest struct {
	Status interface{} `json:"status"`
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1ProductUpdateOfferIdResponseError struct {
	Message string `json:"message"`
	OfferID string `json:"offer_id"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequestCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type DetailsReturns struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type GetConditionalCancellationListV2ResponseCancellationReason struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
}

type V2GetConditionalCancellationListV2Request struct {
	Filters interface{} `json:"filters"`
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
	With interface{} `json:"with"`
}

type PriceIndexesColorIndex string

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairwayStep struct {
	Discount int64 `json:"discount"`
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
}

type V1UnarchiveWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type GetConditionalCancellationListV2RequestFilters struct {
	CancellationInitiator []interface{} `json:"cancellation_initiator"`
	PostingNumber []interface{} `json:"posting_number"`
	State interface{} `json:"state"`
}

type SellerApiProductPrice struct {
	ProductID float64 `json:"product_id"`
	ActionPrice float64 `json:"action_price"`
	Stock float64 `json:"stock"`
}

type AnalyticsDataRowDimension struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type V2ReturnsRfbsListResponse struct {
	Returns interface{} `json:"returns"`
}

type V3TimeRange struct {
	To string `json:"to"`
	From string `json:"from"`
}

type V1FbpDraftDropOffProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1ReviewChangeStatusRequest struct {
	ReviewIds []interface{} `json:"review_ids"`
	Status string `json:"status"`
}

type V1ProductActionTimerStatusResponse struct {
	Statuses interface{} `json:"statuses"`
}

type PostingErrorTypeEnum string

type V3AdditionalDataItem struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type V3ChatHistoryResponse struct {
	HasNext bool `json:"has_next"`
	Messages []interface{} `json:"messages"`
}

type V1SellerActionsProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type FinanceV1GetFinanceAccrualTypesResponseAccrualType struct {
	Description string `json:"description"`
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type V1QuestionInfoRequest struct {
	QuestionID string `json:"question_id"`
}

type V1ProductGetRelatedSKUResponseItem struct {
	DeliverySchema string `json:"delivery_schema"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	Availability string `json:"availability"`
	DeletedAt string `json:"deleted_at"`
}

type V1FbpDraftDirectGetTimeslotResponseTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type QuestionV1GetQuestionListRequestSortDirEnum string

type FinanceCashFlowStatementListResponseDeliveryService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type GetFinanceBalanceV1ResponseService struct {
	Amount interface{} `json:"amount"`
	Name string `json:"name"`
}

type V3Barcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1SellerActionsCreateDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	MinActionPercent float64 `json:"min_action_percent"`
	Title string `json:"title"`
}

type V1ProductUpdateAttributesRequestValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type GetProductInfoListResponseStocksStock struct {
	Present int32 `json:"present"`
	Reserved int32 `json:"reserved"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type SearchAttributeValuesResponseValue struct {
	ID int64 `json:"id"`
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
}

type GetProductAttributesResponseImage struct {
	Default bool `json:"default"`
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
}

type PostingV4PostingFbsListResponsePostingsLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type FbsPostingShipV4RequestPackage struct {
	Products []interface{} `json:"products"`
}

type CompanyTaxSystemEnum string

type GetDiscountTaskListV2ResponseTask struct {
	EditedTill string `json:"edited_till"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	Email string `json:"email"`
	EndAt string `json:"end_at"`
	FirstName string `json:"first_name"`
	Name string `json:"name"`
	Patronymic string `json:"patronymic"`
	ReductionFactor float64 `json:"reduction_factor"`
	ApprovedPrice float64 `json:"approved_price"`
	CreatedAt string `json:"created_at"`
	EndAtDuration int64 `json:"end_at_duration"`
	ID int64 `json:"id"`
	LastName string `json:"last_name"`
	RequestedDiscount float64 `json:"requested_discount"`
	RequestedPrice float64 `json:"requested_price"`
	EditedTillDuration int64 `json:"edited_till_duration"`
	IsAutoModerated bool `json:"is_auto_moderated"`
	MinAutoPrice float64 `json:"min_auto_price"`
	ModeratedAt string `json:"moderated_at"`
	OriginalPrice float64 `json:"original_price"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
	ApprovedDiscount float64 `json:"approved_discount"`
	AutoModeratedInfo interface{} `json:"auto_moderated_info"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V2ReturnsRfbsGetV2ResponseState struct {
	State string `json:"state"`
	StateName string `json:"state_name"`
}

type MoneyMoneyNextTariffCharge struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type V3FbsPosting struct {
	Optional interface{} `json:"optional"`
	OrderID int64 `json:"order_id"`
	OrderNumber string `json:"order_number"`
	Products []interface{} `json:"products"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	TrackingNumber string `json:"tracking_number"`
	DeliveringDate string `json:"delivering_date"`
	DeliveryMethod interface{} `json:"delivery_method"`
	IsPresortable bool `json:"is_presortable"`
	ShipmentDate string `json:"shipment_date"`
	AvailableActions interface{} `json:"available_actions"`
	Cancellation interface{} `json:"cancellation"`
	FinancialData interface{} `json:"financial_data"`
	LegalInfo interface{} `json:"legal_info"`
	ParentPostingNumber string `json:"parent_posting_number"`
	Requirements interface{} `json:"requirements"`
	Substatus string `json:"substatus"`
	AnalyticsData interface{} `json:"analytics_data"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	DestinationPlaceName string `json:"destination_place_name"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	Addressee interface{} `json:"addressee"`
	Barcodes interface{} `json:"barcodes"`
	Customer interface{} `json:"customer"`
	InProcessAt string `json:"in_process_at"`
	IsExpress bool `json:"is_express"`
	PostingNumber string `json:"posting_number"`
	Tariffication interface{} `json:"tariffication"`
	Status string `json:"status"`
}

type TypeTimeOfDay struct {
	Seconds int32 `json:"seconds"`
	Hours int32 `json:"hours"`
	Minutes int32 `json:"minutes"`
	Nanos int32 `json:"nanos"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1ItemSortField string

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails interface{} `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1SellerActionsCreateInstallmentResponse struct {
	ActionID int64 `json:"action_id"`
}

type PolygonBindRequestwhLocation struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type V1PostingFbsSplitResponsePosting struct {
	Products []interface{} `json:"products"`
	PostingNumber string `json:"posting_number"`
}

type PostingFinancialDataProduct struct {
	CurrencyCode string `json:"currency_code"`
	CommissionPercent int64 `json:"commission_percent"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	CustomerPrice float64 `json:"customer_price"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	CustomerCurrencyCode string `json:"customer_currency_code"`
	CommissionAmount float64 `json:"commission_amount"`
	CommissionsCurrencyCode string `json:"commissions_currency_code"`
	Payout float64 `json:"payout"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
}

type Productv3GetProductAttributesV3Request struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

type V1GetWarehouseFBSOperationStatusRequest struct {
	OperationID string `json:"operation_id"`
}

type V1GetFinanceBalanceV1ResponseAccruedMoney struct {
	Value float64 `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponseProducts struct {
	Name string `json:"name"`
	ProductID int64 `json:"product_id"`
	BasePrice float64 `json:"base_price"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
	MinSellerPrice float64 `json:"min_seller_price"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
	SKU int64 `json:"sku"`
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	Currency string `json:"currency"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MinActionQuantity int64 `json:"min_action_quantity"`
}

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"`
}

type DeliveryDetailsSupplyType string

type V1ProductPricesDetailsRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1SellerActionsProductsListRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1SetProductStairwayDiscountByQuantityResponseError struct {
	Data []interface{} `json:"data"`
	SKU int64 `json:"sku"`
}

type V1GetSupplyOrderBundleResponse struct {
	TotalCount int32 `json:"total_count"`
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Items []interface{} `json:"items"`
}

type Financev3Period struct {
	From string `json:"from"`
	To string `json:"to"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearch struct {
	Types []interface{} `json:"types"`
	Address string `json:"address"`
}

type GetProductInfoListResponseModelInfo struct {
	Count int64 `json:"count"`
	ModelID int64 `json:"model_id"`
}

type GetStrategyListResponseStrategy struct {
	ProductsCount int64 `json:"products_count"`
	CompetitorsCount int64 `json:"competitors_count"`
	Enabled bool `json:"enabled"`
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type_"`
	UpdateType string `json:"update_type"`
	UpdatedAt string `json:"updated_at"`
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1AnalyticsProductQueriesResponse struct {
	AnalyticsPeriod interface{} `json:"analytics_period"`
	Items []interface{} `json:"items"`
	PageCount int64 `json:"page_count"`
	Total int64 `json:"total"`
}

type ReportCreateCompanyProductsReportRequest struct {
	Language interface{} `json:"language"`
	OfferID []interface{} `json:"offer_id"`
	Search string `json:"search"`
	SKU []interface{} `json:"sku"`
	Visibility interface{} `json:"visibility"`
}

type Productv3GetProductListRequestFilterFilterVisibility string

type Productv2GetProductListRequestFilterFilterVisibility string

type PostingV4PostingFbsUnfulfilledListResponse struct {
	Count int64 `json:"count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type Productv5Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type TimetableWorkingHours struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type V1WarehouseListRequestWith struct {
	AbleToSetPrice bool `json:"able_to_set_price"`
}

type ArrivalpassArrivalPassCreateRequestArrivalPass struct {
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WarehouseID int64 `json:"warehouse_id"`
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	DropoffPointID int64 `json:"dropoff_point_id"`
}

type SellerSellerAPIArrivalPassUpdateRequest struct {
	CarriageID int64 `json:"carriage_id"`
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type V1FbpCreateActResponse struct {
	Errors []interface{} `json:"errors"`
	FileUuid string `json:"file_uuid"`
	IsSuccess bool `json:"is_success"`
}

type V1GetRealizationReportPostingResponseRow struct {
	LegalEntityDocument interface{} `json:"legal_entity_document"`
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission interface{} `json:"delivery_commission"`
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
	RowNumber int32 `json:"row_number"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	Order interface{} `json:"order"`
}

type V1QuestionChangeStatusRequest struct {
	Status string `json:"status"`
	QuestionIds interface{} `json:"question_ids"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2Product struct {
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
	SKU int64 `json:"sku"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	FreeStock int64 `json:"free_stock"`
	OfferID string `json:"offer_id"`
}

type MoneyMoneyNextTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1CancellationState struct {
	CancellationError interface{} `json:"cancellation_error"`
	CancellationStatus interface{} `json:"cancellation_status"`
}

type ReportReportListRequest struct {
	PageSize int32 `json:"page_size"`
	ReportType interface{} `json:"report_type"`
	Page int32 `json:"page"`
}

type FinanceV1GetFinanceAccrualByDayResponse struct {
	LastID string `json:"last_id"`
	Accruals []interface{} `json:"accruals"`
}

type WarehouseListResponseWarehouse struct {
	IsKgt bool `json:"is_kgt"`
	IsPresorted bool `json:"is_presorted"`
	Status string `json:"status"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
	PostingsLimit int32 `json:"postings_limit"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	IsKarantin bool `json:"is_karantin"`
	IsTimetableEditable bool `json:"is_timetable_editable"`
	WorkingDays interface{} `json:"working_days"`
	Name string `json:"name"`
	WarehouseID int64 `json:"warehouse_id"`
	CanPrintActInAdvance bool `json:"can_print_act_in_advance"`
	FirstMileType interface{} `json:"first_mile_type"`
	IsAbleToSetPrice bool `json:"is_able_to_set_price"`
	MinWorkingDays int64 `json:"min_working_days"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	IsRFBS bool `json:"is_rfbs"`
}

type V1DraftStatusEnum string

type Polygonv1Empty interface{}

type GetFinanceBalanceV1ResponseSales struct {
	Amount interface{} `json:"amount"`
	AmountDetails interface{} `json:"amount_details"`
	Fee interface{} `json:"fee"`
}

type V1AssemblyFbsPostingListResponsePosting struct {
	AssemblyCode string `json:"assembly_code"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type FinanceCashFlowStatementListResponseDetails struct {
	Period interface{} `json:"period"`
	Return interface{} `json:"return"`
	Services interface{} `json:"services"`
	Others interface{} `json:"others"`
	Loan float64 `json:"loan"`
	Payments []interface{} `json:"payments"`
	RFBS interface{} `json:"rfbs"`
	EndBalanceAmount float64 `json:"end_balance_amount"`
	BeginBalanceAmount float64 `json:"begin_balance_amount"`
	Delivery interface{} `json:"delivery"`
	InvoiceTransfer float64 `json:"invoice_transfer"`
}

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID int64 `json:"return_id"`
	ReturnMethodDescription string `json:"return_method_description"`
}

type ErrorData struct {
	Step int64 `json:"step"`
	Value string `json:"value"`
	Code string `json:"code"`
	Field string `json:"field"`
	Message string `json:"message"`
}

type MoneyPostingMoney struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type SellerApiGetSellerActionsV1Response struct {
	Result []interface{} `json:"result"`
}

type V1GetFinanceBalanceV1ResponseFeeMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type CompensationReportLanguage string

type V1GetFinanceBalanceV1ResponsePartnerMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type MoneyMoneyCurrentTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1FbpDraftDirectCreateRequest struct {
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
}

type Productv3GetProductAttributesV3ResponseResult struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	DimensionUnit string `json:"dimension_unit"`
	ID int64 `json:"id"`
	PDFList []interface{} `json:"pdf_list"`
	Attributes []interface{} `json:"attributes"`
	Barcode string `json:"barcode"`
	Height int32 `json:"height"`
	ImageGroupID string `json:"image_group_id"`
	TypeID int64 `json:"type_id"`
	CategoryID int64 `json:"category_id"`
	ColorImage string `json:"color_image"`
	Depth int32 `json:"depth"`
	Images []interface{} `json:"images"`
	OfferID string `json:"offer_id"`
	WeightUnit string `json:"weight_unit"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Images360 []interface{} `json:"images360"`
	Name string `json:"name"`
	Weight int32 `json:"weight"`
	Width int32 `json:"width"`
}

type V4GetProductInfoStocksRequestFilter struct {
	WithQuant interface{} `json:"with_quant"`
	OfferID []interface{} `json:"offer_id"`
	ProductID []interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ItemIDsRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V3Customer struct {
	Phone string `json:"phone"`
	Address interface{} `json:"address"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
}

type V1SellerActionsUpdateInstallmentRequestActionParameters struct {
	Title string `json:"title"`
	DateStart string `json:"date_start"`
}

type PostingPostingProductCancelRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	Items []interface{} `json:"items"`
	PostingNumber string `json:"posting_number"`
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type V1ProductUpdateOfferIdRequest struct {
	UpdateOfferID interface{} `json:"update_offer_id"`
}

type MoneyMoneySelf struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type Productv2ProductsStocksRequestStock struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Stock int64 `json:"stock"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GetProductAttributesV3ResponseDictionaryValue struct {
	Value string `json:"value"`
	DictionaryValueId int64 `json:"dictionaryValueId"`
}

type SellerApiProductV1ResponseResultDeactivate struct {
	ProductIds []interface{} `json:"product_ids"`
	Rejected []interface{} `json:"rejected"`
}

type MoneyMoneyCommission struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
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

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItem struct {
	IsHoliday bool `json:"is_holiday"`
	OpeningHours interface{} `json:"opening_hours"`
	BreakHours interface{} `json:"break_hours"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Percent int64 `json:"percent"`
}

type V1GetStrategyItemInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1WarehouseInvalidProductsGetResponse struct {
	HasNext bool `json:"has_next"`
	LastID int64 `json:"last_id"`
	ValidationResults []interface{} `json:"validation_results"`
	WarehouseID int64 `json:"warehouse_id"`
}

type RatingValuePast struct {
	Formatted string `json:"formatted"`
	Status interface{} `json:"status"`
	Value float64 `json:"value"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1OrderDraftValidationErrorErrorType string

type ProductGetProductAttributesV3ResponseComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type LanguageLanguage string

type V2ServiceType string

type V1FbpDraftDropOffRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type V3User struct {
	ID string `json:"id"`
	Type string `json:"type_"`
}

type ProductUpdateOfferIdRequestUpdateOfferId struct {
	NewOfferID string `json:"new_offer_id"`
	OfferID string `json:"offer_id"`
}

type V1SearchQueriesTopResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type V5FbsPostingProductExemplarStatusV5Request struct {
	PostingNumber string `json:"posting_number"`
}

type V2ReturnsRfbsGetResponse struct {
	Returns interface{} `json:"returns"`
}

type V2ReturnsRfbsFilter struct {
	GroupState []interface{} `json:"group_state"`
	CreatedAt interface{} `json:"created_at"`
	OfferID string `json:"offer_id"`
	PostingNumber string `json:"posting_number"`
}

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V2ReturnsRfbsCompensateRequest struct {
	CompensationAmount string `json:"compensation_amount"`
	ReturnID int64 `json:"return_id"`
}

type CreateWarehouseFBSRequestWorkingDaysEnum string

type DeliveryMethodListResponseDeliveryMethod struct {
	Cutoff string `json:"cutoff"`
	Name string `json:"name"`
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	CompanyID int64 `json:"company_id"`
	ID int64 `json:"id"`
	SlaCutIn int64 `json:"sla_cut_in"`
	TemplateID int64 `json:"template_id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}

type DescriptionCategoryTipsResponseResult struct {
	ImagesURL []interface{} `json:"images_url"`
	InfoURL string `json:"info_url"`
	TypeID int64 `json:"type_id"`
}

type V1SellerActionsUpdateInstallmentRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1ArchiveWarehouseFBSRequest struct {
	Reason string `json:"reason"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductGetProductInfoDescriptionRequest struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
}

type PostingV4PostingFbsListResponsePostingsCustomerAddress struct {
	Latitude float64 `json:"latitude"`
	PvzCode int64 `json:"pvz_code"`
	AddressTail string `json:"address_tail"`
	Country string `json:"country"`
	District string `json:"district"`
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	Region string `json:"region"`
	ZipCode string `json:"zip_code"`
	City string `json:"city"`
	Comment string `json:"comment"`
}

type ReportMarkedProductsSalesCreateRequestDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccruals struct {
	Accruals []interface{} `json:"accruals"`
	PostingNumber string `json:"posting_number"`
}

type ProductV1ProductVisibilitySetResponseItemsErrors struct {
	Code string `json:"code"`
	SKU int64 `json:"sku"`
}

type V1DeleteStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftDropOffProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftPickUpDeleteResponse struct {
	CancellationState interface{} `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type DeliveryDetailsDropOffPointDetails struct {
	Timeslot interface{} `json:"timeslot"`
	ID int64 `json:"id"`
	ProvinceUuid string `json:"province_uuid"`
}

type V1BundleItemErrorEnum string

type SetPostingsResponseResult struct {
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
	Result bool `json:"result"`
}

type V1UpdateWarehouseFBSRequestOptions struct {
	Comment string `json:"comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
}

type ReportCreateReportResponse struct {
	Result interface{} `json:"result"`
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

type ReviewListResponseReview struct {
	CommentsAmount int32 `json:"comments_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	OrderStatus string `json:"order_status"`
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	VideosAmount int32 `json:"videos_amount"`
	ID string `json:"id"`
	PhotosAmount int32 `json:"photos_amount"`
	SKU int64 `json:"sku"`
	Status string `json:"status"`
	Text string `json:"text"`
}

type FbsPostingShipV4ResponseShipAdditionalData struct {
	PostingNumber string `json:"posting_number"`
	Products interface{} `json:"products"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type V1FbpAvailableTimeslotListRequest struct {
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
	SupplyID string `json:"supply_id"`
}

type ChatChatSendFileRequest struct {
	Base64Content string `json:"base64_content"`
	ChatID string `json:"chat_id"`
	Name string `json:"name"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type V3ImportProductsRequestDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type V1WarehouseFbsCreatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type ValidationResultValidationError struct {
	TemplateID int32 `json:"template_id"`
	Type interface{} `json:"type_"`
	Characteristic interface{} `json:"characteristic"`
	RestrictionPrice interface{} `json:"restriction_price"`
	RestrictionVwc float64 `json:"restriction_vwc"`
}

type V1FbpEditTimeslotRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type Productv2DeleteProductsResponse struct {
	Status []interface{} `json:"status"`
}

type ExemplarsMarks struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type ProductV1ProductVisibilityInfoResponse struct {
	Items []interface{} `json:"items"`
}

type V1GetRealizationReportPostingResponse struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"`
}

type V1OrderErrorTypeEnum string

type WarehouseDeliveryMethodListResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V1FbpCheckActStateResponse struct {
	CdnURL string `json:"cdn_url"`
	Error interface{} `json:"error"`
	Status interface{} `json:"status"`
}

type V3FbsPostingDetailOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V6FbsPostingProductExemplarSetV6Request struct {
	Products []interface{} `json:"products"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
}

type PostingV4PostingFbsListResponsePostingsCancellation struct {
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
}

type FbsPostingProductExemplarCreateOrGetV6ResponseProduct struct {
	IsMandatoryMarkPossible bool `json:"is_mandatory_mark_possible"`
	IsRnptNeeded bool `json:"is_rnpt_needed"`
	ProductID int64 `json:"product_id"`
	HasImei bool `json:"has_imei"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
	IsMandatoryMarkNeeded bool `json:"is_mandatory_mark_needed"`
	Quantity int32 `json:"quantity"`
	Exemplars []interface{} `json:"exemplars"`
	IsJwUinNeeded bool `json:"is_jw_uin_needed"`
}

type V1GetStrategyResponseResult struct {
	Type string `json:"type_"`
	UpdateType string `json:"update_type"`
	Competitors []interface{} `json:"competitors"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
}

type V2FbsPostingProductCountrySetResponse struct {
	ProductID int64 `json:"product_id"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
}

type Polygonv1PolygonBindRequest struct {
	DeliveryMethodID int32 `json:"delivery_method_id"`
	Polygons []interface{} `json:"polygons"`
	WarehouseLocation interface{} `json:"warehouse_location"`
}

type V1AssemblyFbsProductListRequestFilter struct {
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	CutoffFrom string `json:"cutoff_from"`
}

type ProductV4GetUploadQuotaResponseOperationLimitsLimitTypeEnum string

type V1FbpDraftDirectCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
}

type V1CreatePricingStrategyResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftDirectGetTimeslotRequest struct {
	BundleID string `json:"bundle_id"`
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
	WarehouseID int64 `json:"warehouse_id"`
}

type RowItemCommission struct {
	Total float64 `json:"total"`
	Bonus float64 `json:"bonus"`
	Commission float64 `json:"commission"`
	PricePerInstance float64 `json:"price_per_instance"`
	Quantity int32 `json:"quantity"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Amount float64 `json:"amount"`
	Compensation float64 `json:"compensation"`
	StandardFee float64 `json:"standard_fee"`
	Stars float64 `json:"stars"`
}

type CommentListResponseComment struct {
	ID string `json:"id"`
	IsOfficial bool `json:"is_official"`
	IsOwner bool `json:"is_owner"`
	ParentCommentID string `json:"parent_comment_id"`
	PublishedAt string `json:"published_at"`
	Text string `json:"text"`
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

type PostingCancelFbsPostingRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	PostingNumber string `json:"posting_number"`
}

type ActionsV1ActionsAutoAddProductsCandidatesRequest struct {
	Offset int64 `json:"offset"`
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
}

type Fbsv4FbsPostingShipV4Response struct {
	AdditionalData interface{} `json:"additional_data"`
	Result interface{} `json:"result"`
}

type V1SellerActionsVoucherGetRequest struct {
	ActionID int64 `json:"action_id"`
}

type GetProductInfoListResponsePriceIndexes struct {
	ExternalIndexData interface{} `json:"external_index_data"`
	OzonIndexData interface{} `json:"ozon_index_data"`
	SelfMarketplacesIndexData interface{} `json:"self_marketplaces_index_data"`
	ColorIndex interface{} `json:"color_index"`
}

type V1QuestionAnswerCreateResponse struct {
	AnswerID string `json:"answer_id"`
}

type V3GetProductInfoListResponse struct {
	Items []interface{} `json:"items"`
}

type FirstMileTypeEnum string

type V1CommentSort string

type PostingV4PostingFbsListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type ActionParameterDiscountTypeEnum string

type ReportReportInfoResponse struct {
	Result interface{} `json:"result"`
}

type Fbsv4PostingProductDetailWithoutDimensions struct {
	CurrencyCode string `json:"currency_code"`
	MandatoryMark interface{} `json:"mandatory_mark"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type ReturnsRfbsGetResponseReturns struct {
	AvailableActions []interface{} `json:"available_actions"`
	ReturnReason interface{} `json:"return_reason"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientPhoto []interface{} `json:"client_photo"`
	ClientReturnMethodType interface{} `json:"client_return_method_type"`
	PostingNumber string `json:"posting_number"`
	RejectionComment string `json:"rejection_comment"`
	ClientName string `json:"client_name"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
	OrderNumber string `json:"order_number"`
	RejectionReason []interface{} `json:"rejection_reason"`
	State interface{} `json:"state"`
	Product interface{} `json:"product"`
	ReturnMethodDescription string `json:"return_method_description"`
	ReturnNumber string `json:"return_number"`
	RuPostTrackingNumber string `json:"ru_post_tracking_number"`
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []interface{} `json:"result"`
}

type ProductGetProductAttributesV4ResponseAttribute struct {
	ID int64 `json:"id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type GetRealizationReportByDayResponseRow struct {
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission interface{} `json:"delivery_commission"`
	Item interface{} `json:"item"`
	ReturnCommission interface{} `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
}

type ReturnsRfbsGetV2ResponseAvailableAction struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type Postingv3GetFbsPostingListRequestFilter struct {
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	ProviderID []interface{} `json:"provider_id"`
	To string `json:"to"`
	WarehouseID []interface{} `json:"warehouse_id"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	FbpFilter string `json:"fbpFilter"`
	OrderID int64 `json:"order_id"`
	Since string `json:"since"`
	Status string `json:"status"`
}

type V3GetProductInfoListResponsePromotion struct {
	IsEnabled bool `json:"is_enabled"`
	Type interface{} `json:"type_"`
}

type PostingV4PostingFbsListRequestFilter struct {
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	LastChangedStatusDate interface{} `json:"last_changed_status_date"`
	OrderID int64 `json:"order_id"`
	OrderNumbers []interface{} `json:"order_numbers"`
	Since string `json:"since"`
	To string `json:"to"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	ProviderIds []interface{} `json:"provider_ids"`
	Statuses []interface{} `json:"statuses"`
}

type ProductImportProductsBySKURequestItem struct {
	Vat string `json:"vat"`
	CurrencyCode string `json:"currency_code"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	OldPrice string `json:"old_price"`
	Price string `json:"price"`
	SKU int64 `json:"sku"`
}

type SellerInfoResponseRatingTypeEnum string

type V1ProductFbsSplit struct {
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
}

type V2GetConditionalCancellationListV2Response struct {
	Counter int64 `json:"counter"`
	LastID int64 `json:"last_id"`
	Result []interface{} `json:"result"`
}

type SellerOzonLogisticsInfoResponseAvailableSchemasEnum string

type PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod struct {
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponseDropOffPointTypeEnum string

type V1GetProductInfoDiscountedResponse struct {
	Items interface{} `json:"items"`
}

type V1SellerInfoResponse struct {
	Company interface{} `json:"company"`
	Ratings []interface{} `json:"ratings"`
	Subscription interface{} `json:"subscription"`
}

type V1AnalyticsProductQueriesDetailsRequestSortBy string

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

type V1QuestionCountResponse struct {
	New int64 `json:"new"`
	Processed int64 `json:"processed"`
	Unprocessed int64 `json:"unprocessed"`
	Viewed int64 `json:"viewed"`
	All int64 `json:"all"`
}

type V1AssemblyCarriageProductListRequestFilter struct {
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1PostingFbsSplitRequestPosting struct {
	Products []interface{} `json:"products"`
}

type RatingStatusEnum string

type ProductImportProductsBySKURequest struct {
	Items []interface{} `json:"items"`
}

type PostingCancelReasonListResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftPickUpProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1GetProductRatingBySkuRequest struct {
	Skus interface{} `json:"skus"`
}

type PostingV4PostingFbsListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
}

type ReturnsCompanyFbsInfoResponsePassInfo struct {
	Count int32 `json:"count"`
	IsRequired bool `json:"is_required"`
}

type Fbpv1Timeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1FbpDraftPickUpProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1FbpDraftDirectTplDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type DetailsReturnDetails struct {
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	ReturnServices interface{} `json:"return_services"`
}

type WarehouseInvalidProductsGetResponseValidationResult struct {
	Item interface{} `json:"item"`
	State interface{} `json:"state"`
	ValidationErrors []interface{} `json:"validation_errors"`
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
}

type DeleteProductsResponseDeleteStatus struct {
	Error string `json:"error"`
	IsDeleted bool `json:"is_deleted"`
	OfferID string `json:"offer_id"`
}

type V2FbsPostingResponse struct {
	Result interface{} `json:"result"`
}

type V1ProductUpdateAttributesRequestItem struct {
	Attributes interface{} `json:"attributes"`
	OfferID string `json:"offer_id"`
}

type GetProductInfoListResponseError struct {
	Field string `json:"field"`
	Level interface{} `json:"level"`
	State string `json:"state"`
	Texts interface{} `json:"texts"`
	AttributeID int64 `json:"attribute_id"`
	Code string `json:"code"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1GetFinanceBalanceV1Request struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplar struct {
	Valid bool `json:"valid"`
	Errors []interface{} `json:"errors"`
	GTD string `json:"gtd"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
}

type PostingV4PostingFbsListResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type PostingV1PostingFbpListResponsePostingsProducts struct {
	OfferID string `json:"offer_id"`
	Price interface{} `json:"price"`
	Quantity int32 `json:"quantity"`
	SellerPrice interface{} `json:"seller_price"`
	SKU int64 `json:"sku"`
	CustomerPrice interface{} `json:"customer_price"`
	Name string `json:"name"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2 struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V1ApproveDeclineDiscountTasksResponseResult struct {
	FailDetails []interface{} `json:"fail_details"`
	SuccessCount int32 `json:"success_count"`
	FailCount int32 `json:"fail_count"`
}

type SortingOrder string

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type GetReturnsListResponsePlaceNow struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type PostingV1PostingFbpListRequestFilter struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
	Since string `json:"since"`
	Statuses []interface{} `json:"statuses"`
	To string `json:"to"`
}

type V2Product struct {
	CurrencyCode string `json:"currency_code"`
	Price int32 `json:"price"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type FbsPostingShipV4RequestWith struct {
	AdditionalData bool `json:"additional_data"`
}

type V1AnalyticsProductQueriesDetailsRequestSortDir string

type QuestionV1QuestionAnswerListResponseAnswerStatusPublicationEnum string

type Productsv1GetProductInfoStocksByWarehouseFbsResponseResult struct {
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
}

type FinanceCashFlowStatementListResponseReturnService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type GetStrategyIDsByItemIDsResponseProductInfo struct {
	ProductID int64 `json:"product_id"`
	StrategyID string `json:"strategy_id"`
}

type ProductGetProductInfoPricesV5ResponseItem struct {
	Price interface{} `json:"price"`
	PriceIndexes interface{} `json:"price_indexes"`
	ProductID int64 `json:"product_id"`
	VolumeWeight float64 `json:"volume_weight"`
	Acquiring float64 `json:"acquiring"`
	Commissions interface{} `json:"commissions"`
	MarketingActions interface{} `json:"marketing_actions"`
	OfferID string `json:"offer_id"`
}

type PostingV1PostingFbpListResponsePostingsFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	DeliveryAmount float64 `json:"delivery_amount"`
	Products []interface{} `json:"products"`
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

type MoneyMoney struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type DeliveryMethodListV2RequestFilter struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	ProviderIds []interface{} `json:"provider_ids"`
	Status []interface{} `json:"status"`
}

type V1SellerActionsListRequest struct {
	ActionIds []interface{} `json:"action_ids"`
	ActionType []interface{} `json:"action_type"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Search string `json:"search"`
	Status []interface{} `json:"status"`
}

type V1SearchQueriesTextResponseSearchQuery struct {
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
	ConversionToCart float64 `json:"conversion_to_cart"`
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
	SellersCount float64 `json:"sellers_count"`
	AddToCart float64 `json:"add_to_cart"`
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityEnum string

type V1ProductPricesDetailsResponsePrice struct {
	Price interface{} `json:"price"`
	PriceIndexes []interface{} `json:"price_indexes"`
	SKU int64 `json:"sku"`
	CustomerPrice interface{} `json:"customer_price"`
	DiscountPercent float64 `json:"discount_percent"`
	OfferID string `json:"offer_id"`
}

type V4GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"`
	Name string `json:"name"`
}

type V1PostingFbsSplitResponsePostingParent struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type MoneyMoneyBonus struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type DetailsDeliveryDetails struct {
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	DeliveryServices interface{} `json:"delivery_services"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequestToUpdate struct {
	ActionPrice float64 `json:"action_price"`
	Currency string `json:"currency"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
}

type V1AssemblyFbsProductListResponse struct {
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
	ProductsCount int32 `json:"products_count"`
}

type FbpCheckConsignmentNoteStateResponseStateType string

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"`
}

type ParameterDiscountLevels struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer struct {
	Phone string `json:"phone"`
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
}

type StatusEnum string

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData struct {
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
	ClusterFrom string `json:"cluster_from"`
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
}

type DetailsOthers struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type SourceShipmentType string

type GetRealizationReportResponseV2Result struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccrualsAccrual struct {
	SellerPrice interface{} `json:"seller_price"`
	SKU int64 `json:"sku"`
	TypeID int32 `json:"type_id"`
	AccrualDate string `json:"accrual_date"`
	Accrued interface{} `json:"accrued"`
	Quantity int32 `json:"quantity"`
}

type V3ImportProductsResponseResult struct {
	TaskID int64 `json:"task_id"`
}

type V1SellerActionsCreateVoucherRequestDiscountTypeEnum string

type ChatChatStartRequest struct {
	PostingNumber string `json:"posting_number"`
}

type ProductImportProductsPricesRequestPrice struct {
	OfferID string `json:"offer_id"`
	PriceStrategyEnabled string `json:"price_strategy_enabled"`
	Vat string `json:"vat"`
	AutoActionEnabled string `json:"auto_action_enabled"`
	CurrencyCode string `json:"currency_code"`
	ManageElasticBoostingThroughPrice bool `json:"manage_elastic_boosting_through_price"`
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	NetPrice string `json:"net_price"`
	OldPrice string `json:"old_price"`
	Price string `json:"price"`
	ProductID int64 `json:"product_id"`
	MinPrice string `json:"min_price"`
}

type SellerSellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsProducts struct {
	Imei []interface{} `json:"imei"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	Price interface{} `json:"price"`
	ProductColor string `json:"product_color"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Weight float64 `json:"weight"`
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type Postingv3FbsPostingWithParams struct {
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
	Translit bool `json:"translit"`
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
}

type V3ImportProductsRequestItem struct {
	TypeID int64 `json:"type_id"`
	Vat string `json:"vat"`
	Promotions []interface{} `json:"promotions"`
	Attributes []interface{} `json:"attributes"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	GeoNames interface{} `json:"geo_names"`
	Images360 []interface{} `json:"images360"`
	Name string `json:"name"`
	WeightUnit string `json:"weight_unit"`
	OldPrice string `json:"old_price"`
	Weight int32 `json:"weight"`
	Width int32 `json:"width"`
	Barcode string `json:"barcode"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Depth int32 `json:"depth"`
	DimensionUnit string `json:"dimension_unit"`
	PrimaryImage string `json:"primary_image"`
	OfferID string `json:"offer_id"`
	PDFList []interface{} `json:"pdf_list"`
	Price string `json:"price"`
	NewDescriptionCategoryID int64 `json:"new_description_category_id"`
	ColorImage string `json:"color_image"`
	CurrencyCode string `json:"currency_code"`
	Height int32 `json:"height"`
	Images []interface{} `json:"images"`
	ServiceType interface{} `json:"service_type"`
}

type V1UpdateWarehouseFBSRequestWorkingDaysEnum string

type PostingV4PostingFbsListResponsePostingsProducts struct {
	Price interface{} `json:"price"`
	ProductColor string `json:"product_color"`
	SKU int64 `json:"sku"`
	Weight float64 `json:"weight"`
	Name string `json:"name"`
	Quantity int32 `json:"quantity"`
	Imei []interface{} `json:"imei"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	OfferID string `json:"offer_id"`
}

type ChatChatStartResponse struct {
	Result interface{} `json:"result"`
}

type V1GetDiscountTaskListResponseTask struct {
	UserComment string `json:"user_comment"`
	OriginalPrice float64 `json:"original_price"`
	BasePrice float64 `json:"base_price"`
	IsAutoModerated bool `json:"is_auto_moderated"`
	CreatedAt string `json:"created_at"`
	DiscountPercent float64 `json:"discount_percent"`
	MinAutoPrice float64 `json:"min_auto_price"`
	IsDamaged bool `json:"is_damaged"`
	Email string `json:"email"`
	ApprovedPriceWithFee float64 `json:"approved_price_with_fee"`
	ApprovedPriceFeePercent float64 `json:"approved_price_fee_percent"`
	RequestedPrice float64 `json:"requested_price"`
	LastName string `json:"last_name"`
	ID int64 `json:"id"`
	EditedTill string `json:"edited_till"`
	Discount float64 `json:"discount"`
	ApprovedDiscount float64 `json:"approved_discount"`
	Patronymic string `json:"patronymic"`
	RequestedPriceWithFee float64 `json:"requested_price_with_fee"`
	Status string `json:"status"`
	CustomerName string `json:"customer_name"`
	SellerComment string `json:"seller_comment"`
	PrevTaskID int64 `json:"prev_task_id"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	EndAt string `json:"end_at"`
	SKU int64 `json:"sku"`
	ApprovedPrice float64 `json:"approved_price"`
	ModeratedAt string `json:"moderated_at"`
	ApprovedDiscountPercent float64 `json:"approved_discount_percent"`
	IsPurchased bool `json:"is_purchased"`
	FirstName string `json:"first_name"`
	OfferID string `json:"offer_id"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	RequestedQuantityMin int64 `json:"requested_quantity_min"`
}

type Productv3GetProductListResponseResult struct {
	Items interface{} `json:"items"`
	LastID string `json:"last_id"`
	Total int32 `json:"total"`
}

type PostingCancelReasonResponse struct {
	Result []interface{} `json:"result"`
}

type V2GetDiscountTaskListV2RequestDiscountTaskStatusEnum string

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPoint struct {
	Type interface{} `json:"type_"`
	Address string `json:"address"`
	Coordinates interface{} `json:"coordinates"`
	DiscountPercent float64 `json:"discount_percent"`
	ID string `json:"id"`
	LastTransitTimeLocal interface{} `json:"last_transit_time_local"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData struct {
	City string `json:"city"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region string `json:"region"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"`
}

type PostingV4PostingFbsListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

type V1ItemSfboAttribute string

type V1FbpDraftDirectTimeslotEditResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type GetReturnsListResponseExemplar struct {
	ID int64 `json:"id"`
}

type V3ChatListRequestFilter struct {
	ChatStatus string `json:"chat_status"`
	UnreadOnly bool `json:"unread_only"`
}

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	DropOffDate string `json:"drop_off_date"`
}

type V1GenerateBarcodeResult struct {
	Code string `json:"code"`
	Error string `json:"error"`
	Barcode string `json:"barcode"`
	ProductID int64 `json:"product_id"`
}

type V1RolesByTokenResponse struct {
	ExpiresAt string `json:"expires_at"`
	Roles []interface{} `json:"roles"`
}

type FinanceTransactionTotalsV3RequestDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type CreateReportResponseCode struct {
	Code string `json:"code"`
}

type MoneyMoneyCurrentTariffCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type PostingV4PostingFbsListResponsePostings struct {
	Addressee interface{} `json:"addressee"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	Optional interface{} `json:"optional"`
	OrderNumber string `json:"order_number"`
	ParentPostingNumber string `json:"parent_posting_number"`
	Requirements interface{} `json:"requirements"`
	TrackingNumber string `json:"tracking_number"`
	ExternalOrder interface{} `json:"external_order"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	Barcodes interface{} `json:"barcodes"`
	FinancialData interface{} `json:"financial_data"`
	IsMultibox bool `json:"is_multibox"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	DestinationPlaceName string `json:"destination_place_name"`
	Substatus string `json:"substatus"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	PrrOption string `json:"prr_option"`
	InProcessAt string `json:"in_process_at"`
	AnalyticsData interface{} `json:"analytics_data"`
	IsExpress bool `json:"is_express"`
	LegalInfo interface{} `json:"legal_info"`
	Products []interface{} `json:"products"`
	ShipmentDate string `json:"shipment_date"`
	DeliveryMethod interface{} `json:"delivery_method"`
	IsPresortable bool `json:"is_presortable"`
	OrderID int64 `json:"order_id"`
	PostingNumber string `json:"posting_number"`
	QuantumID int64 `json:"quantum_id"`
	VolumeWeight float64 `json:"volume_weight"`
	AvailableActions []interface{} `json:"available_actions"`
	Cancellation interface{} `json:"cancellation"`
	Customer interface{} `json:"customer"`
	DeliveringDate string `json:"delivering_date"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	DeliverySchema string `json:"delivery_schema"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	Status string `json:"status"`
	Tariffication interface{} `json:"tariffication"`
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

type ArrivalpassArrivalPassUpdateRequestArrivalPass struct {
	ArrivalPassID int64 `json:"arrival_pass_id"`
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
}

type V4FbsPostingShipPackageV4Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type ValidationResultItemStateEnum string

type V3ImportProductsRequestComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type GooglerpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
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

type ProductProductArchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V1CommentCreateResponse struct {
	CommentID string `json:"comment_id"`
}

type ListFBSRatingIndexPostingsV1ResponseError struct {
	ChargePrice float64 `json:"charge_price"`
	DeliverySchema string `json:"delivery_schema"`
	ErrorAt string `json:"error_at"`
	Index float64 `json:"index"`
	PostingErrorType interface{} `json:"posting_error_type"`
	PostingNumber string `json:"posting_number"`
	ProductPrice float64 `json:"product_price"`
	ChargePercent float64 `json:"charge_percent"`
	ChargePriceCurrencyCode string `json:"charge_price_currency_code"`
	HasGraceStatus bool `json:"has_grace_status"`
	ProductPriceCurrencyCode string `json:"product_price_currency_code"`
}

type Productv2ProductsStocksRequest struct {
	Stocks []interface{} `json:"stocks"`
}

type GetFBSRatingIndexInfoV1ResponseIndexDynamics struct {
	Date string `json:"date"`
	IndexByDate float64 `json:"index_by_date"`
	ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"`
}

type V3Addressee struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber string `json:"posting_number"`
	ProductID int64 `json:"product_id"`
	CountryIsoCode string `json:"country_iso_code"`
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

type FinanceV1GetFinanceAccrualPostingsRequest struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1QuestionListRequest struct {
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Filter interface{} `json:"filter"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDeliveryService struct {
	Accrued interface{} `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type ProductV1ProductVisibilitySetRequest struct {
	ItemPlacement []interface{} `json:"item_placement"`
}

type PriceIndexesIndexDataOzon struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1GetStrategyListResponse struct {
	Strategies []interface{} `json:"strategies"`
	Total int32 `json:"total"`
}

type V1SearchQueriesTopResponseSearchQuery struct {
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
	SellersCount float64 `json:"sellers_count"`
	AddToCart float64 `json:"add_to_cart"`
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
	ConversionToCart float64 `json:"conversion_to_cart"`
}

type V1GetWarehouseFBSOperationStatusResponse struct {
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Type interface{} `json:"type_"`
	Error interface{} `json:"error"`
}

type FbpDraftDropOffPointListResponseDropOffPoint struct {
	PointAddress string `json:"point_address"`
	ProvinceUuid string `json:"province_uuid"`
	City string `json:"city"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	NearestDropOffDate string `json:"nearest_drop_off_date"`
}

type V1SellerActionsCreateDiscountWithConditionResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1FbpDraftDropOffCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts struct {
	Commission interface{} `json:"commission"`
	CustomerPrice interface{} `json:"customer_price"`
	OldPrice float64 `json:"old_price"`
	Payout float64 `json:"payout"`
	Price float64 `json:"price"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
	ProductID int64 `json:"product_id"`
}

type V1FbpDraftDirectProductValidateResponseRejectedItem struct {
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
}

type V1DeleteStrategyItemsResponseResult struct {
	FailedProductCount int32 `json:"failed_product_count"`
}

type ProductGetProductInfoDescriptionResponse struct {
	Result interface{} `json:"result"`
}

type V1QuestionAnswerListResponse struct {
	Answers interface{} `json:"answers"`
	LastID string `json:"last_id"`
}

type FinanceCashFlowStatementListResponseService struct {
	Price float64 `json:"price"`
	Name string `json:"name"`
}

type MoneyMoneySaleCommission struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type GetReturnsListResponseVisual struct {
	Status interface{} `json:"status"`
	ChangeMoment string `json:"change_moment"`
}

type V1GetRealizationReportByDayRequest struct {
	Day int32 `json:"day"`
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type V1AnalyticsProductQueriesDetailsResponse struct {
	Queries []interface{} `json:"queries"`
	Total int64 `json:"total"`
	AnalyticsPeriod interface{} `json:"analytics_period"`
	PageCount int64 `json:"page_count"`
}

type V1BundleItemShipmentType string

type CarriageCarriageGetResponse struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
	IntegrationType string `json:"integration_type"`
	Status string `json:"status"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
	AvailableActions []interface{} `json:"available_actions"`
	CarriageID int64 `json:"carriage_id"`
	FirstMileType string `json:"first_mile_type"`
	IsPartial bool `json:"is_partial"`
	RetryCount int32 `json:"retry_count"`
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CompanyID int64 `json:"company_id"`
	CreatedAt string `json:"created_at"`
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	ActType string `json:"act_type"`
	CancelAvailability interface{} `json:"cancel_availability"`
	ContainersCount int32 `json:"containers_count"`
	IsContainerLabelPrinted bool `json:"is_container_label_printed"`
	PartialNum int64 `json:"partial_num"`
}

type ParameterAutoStopActionReasonEnum string

type V3CustomerFbsLists struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address interface{} `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
}

type V4Visibility string

type V1UpdateWarehouseFBSFirstMileRequestFirstMileTypeEnum string

type SellerInfoResponseRating struct {
	Name string `json:"name"`
	PastValue interface{} `json:"past_value"`
	Rating string `json:"rating"`
	Status interface{} `json:"status"`
	ValueType interface{} `json:"value_type"`
	CurrentValue interface{} `json:"current_value"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityListEnum string

type GetRealizationReportResponseV2Header struct {
	DocDate string `json:"doc_date"`
	Number string `json:"number"`
	PayerInn string `json:"payer_inn"`
	PayerName string `json:"payer_name"`
	ReceiverInn string `json:"receiver_inn"`
	StartDate string `json:"start_date"`
	StopDate string `json:"stop_date"`
	ContractDate string `json:"contract_date"`
	ContractNumber string `json:"contract_number"`
	PayerKpp string `json:"payer_kpp"`
	ReceiverKpp string `json:"receiver_kpp"`
	ReceiverName string `json:"receiver_name"`
	CurrencySysName string `json:"currency_sys_name"`
}

type ProductV1ProductVisibilityInfoRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1FbpDraftDropOffRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type OperationPosting struct {
	PostingNumber string `json:"posting_number"`
	WarehouseID int64 `json:"warehouse_id"`
	DeliverySchema string `json:"delivery_schema"`
	OrderDate string `json:"order_date"`
}

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
}

type V1FbpDraftListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type ReviewV2ReviewListV2ResponseReviewStatusEnum string

type V1ListDropOffPointsForCreateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type WarehouseFirstMile struct {
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
	TimeslotFrom string `json:"timeslot_from"`
	TimeslotID int64 `json:"timeslot_id"`
	TimeslotTo string `json:"timeslot_to"`
	Type interface{} `json:"type_"`
	DropoffPointID string `json:"dropoff_point_id"`
}

type V1FbpDraftDropOffPointListRequest struct {
	NextPageNumber int32 `json:"next_page_number"`
	PageSize int32 `json:"page_size"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type FbpCreateActResponseCreateActErrorReason string

type V1SellerActionsProductsAddRequestProduct struct {
	Currency interface{} `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	SKU int64 `json:"sku"`
}

type CancellationErrorCode string

type V2ConditionalCancellationMoveV2Request struct {
	CancellationID int64 `json:"cancellation_id"`
	Comment string `json:"comment"`
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V3GetFbsPostingResponseV3 struct {
	Result interface{} `json:"result"`
}

type SellerApiProductV1ResponseResult struct {
	ProductIds []interface{} `json:"product_ids"`
	Rejected []interface{} `json:"rejected"`
}

type V3ChatList struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
}

type V3PostingProductDetail struct {
	HasImei bool `json:"has_imei"`
	MandatoryMark []interface{} `json:"mandatory_mark"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	CurrencyCode string `json:"currency_code"`
	Dimensions interface{} `json:"dimensions"`
	Name string `json:"name"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type GetDiscountTaskListV2ResponseTaskDiscountTaskStatusEnum string

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
	HasNext bool `json:"has_next"`
}

type ActionParameterVoucherParameter struct {
	CountCodes int64 `json:"count_codes"`
	IsPrivate bool `json:"is_private"`
	Type interface{} `json:"type_"`
}

type ProductGetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type ReturnsRfbsGetV2ResponseClientReturnMethodType struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
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

type V1ApproveDiscountTasksRequestTask struct {
	SellerComment string `json:"seller_comment"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	ID int64 `json:"id"`
	ApprovedPrice float64 `json:"approved_price"`
}

type V1ReviewInfoResponse struct {
	Text string `json:"text"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	PhotosAmount int32 `json:"photos_amount"`
	CommentsAmount int32 `json:"comments_amount"`
	VideosAmount int32 `json:"videos_amount"`
	Photos []interface{} `json:"photos"`
	PublishedAt string `json:"published_at"`
	SKU int64 `json:"sku"`
	Videos []interface{} `json:"videos"`
	DislikesAmount int32 `json:"dislikes_amount"`
	ID string `json:"id"`
	LikesAmount int32 `json:"likes_amount"`
	OrderStatus string `json:"order_status"`
	Rating int32 `json:"rating"`
	Status string `json:"status"`
}

type V1AssemblyFbsPostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
}

type ImportProductRequestPromotion struct {
	Operation string `json:"operation"`
	Type string `json:"type_"`
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

type V1CreatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyName string `json:"strategy_name"`
}

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type ItemPricev5 struct {
	CurrencyCode string `json:"currency_code"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	Vat float64 `json:"vat"`
	AutoActionEnabled bool `json:"auto_action_enabled"`
	MarketingSellerPrice float64 `json:"marketing_seller_price"`
	MinPrice float64 `json:"min_price"`
	NetPrice float64 `json:"net_price"`
	RetailPrice float64 `json:"retail_price"`
}

type Productv1ProductInfoPicturesResponse struct {
	Result interface{} `json:"result"`
}

type Productv2ProductsStocksResponseResult struct {
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Updated bool `json:"updated"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductInfoWarehouseStocksResponseStocks struct {
	FreeStock int64 `json:"free_stock"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
	SKU int64 `json:"sku"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1SellerActionsVoucherGetResponse struct {
	File string `json:"file"`
}

type V3GetProductInfoListRequest struct {
	OfferID []interface{} `json:"offer_id"`
	ProductID []interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
}

type V1ItemError struct {
	Level string `json:"level"`
	Description string `json:"description"`
	Field string `json:"field"`
	AttributeID int64 `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
	Code string `json:"code"`
	Message string `json:"message"`
	State string `json:"state"`
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsCreateVoucherResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type GetWarehouseFBSOperationStatusResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type CommonCreateReportResponse struct {
	Result interface{} `json:"result"`
}

type MoneyMoneyAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1ProductUpdateOfferIdResponse struct {
	Errors interface{} `json:"errors"`
}

type V1ApproveDeclineDiscountTasksResponse struct {
	Result interface{} `json:"result"`
}

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnForBackWay int64 `json:"return_for_back_way"`
	ReturnID int64 `json:"return_id"`
}

type TaskAutoModeratedInfo struct {
	MaxPercent float64 `json:"max_percent"`
	MaxPrice float64 `json:"max_price"`
	MinPercent float64 `json:"min_percent"`
	MinPrice float64 `json:"min_price"`
}

type CreateWarehouseFBSRequestFirstMileTypeEnum string

type Postingv3GetFbsPostingRequest struct {
	With interface{} `json:"with"`
	PostingNumber string `json:"posting_number"`
}

type V1ProductUpdateAttributesRequestAttribute struct {
	ComplexID int64 `json:"complex_id"`
	ID int64 `json:"id"`
	Values interface{} `json:"values"`
}

type V1AnalyticsProductQueriesDetailsRequest struct {
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Skus []interface{} `json:"skus"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	LimitBySKU int32 `json:"limit_by_sku"`
}

type V1SellerActionsCreateDiscountWithConditionRequestDiscountTypeEnum string

type V2CancellationInitiatorEnum string

type CancellationStateCancellationError struct {
	ErrorCode interface{} `json:"error_code"`
	Message string `json:"message"`
}

type V3FinanceCashFlowStatementListRequest struct {
	Date interface{} `json:"date"`
	Page int32 `json:"page"`
	WithDetails bool `json:"with_details"`
	PageSize int32 `json:"page_size"`
}

type V1FbpDraftDropOffCreateRequestDeliveryDetails struct {
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
}

type ActionsV1ActionsAutoAddProductsListRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V2ChatReadResponse struct {
	UnreadCount int64 `json:"unread_count"`
}

type V1DescriptionCategoryTipsRequest struct {
	TypeID []interface{} `json:"type_id"`
}

type V1AssemblyFbsPostingListRequestFilter struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
}

type ReportLanguage string

type V1QuestionAnswerCreateDefault struct {
	Code int32 `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type Productv4GetProductAttributesV4Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type V1GetCompetitorsResponse struct {
	Competitor []interface{} `json:"competitor"`
	Total int32 `json:"total"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairways struct {
	Enabled bool `json:"enabled"`
	SKU int64 `json:"sku"`
	Stairway interface{} `json:"stairway"`
}

type V1GetStrategyResponse struct {
	Result interface{} `json:"result"`
}

type V3Cancellation struct {
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
}

type Productv4GetProductAttributesV4Request struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

type MoneyMoneySaleAmount struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type V1GetStrategyListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type FbpDraftDropOffProvinceListResponseProvince struct {
	Name string `json:"name"`
	PointsCount int32 `json:"points_count"`
	ProvinceUuid string `json:"province_uuid"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndex struct {
	ExternalIndexData interface{} `json:"external_index_data"`
	SelfIndexData interface{} `json:"self_index_data"`
}

type DeliveryMethodListRequestFilter struct {
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ReviewV2ReviewListV2RequestSortDirEnum string

type V5FbsPostingProductExemplarValidateV5RequestProductExemplar struct {
	Rnpt string `json:"rnpt"`
	GTD string `json:"gtd"`
	Marks []interface{} `json:"marks"`
}

type V1SellerActionsArchiveRequest struct {
	ActionID int64 `json:"action_id"`
}

type AvailabilityReason struct {
	HumanText interface{} `json:"human_text"`
	ID int64 `json:"id"`
}

type ActionsV1ActionsAutoAddProductsListResponse struct {
	Products []interface{} `json:"products"`
	Total int64 `json:"total"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFees struct {
	Fees []interface{} `json:"fees"`
}

type FilterOp interface{}

type V1FbpGetLabelRequest struct {
	SupplyID string `json:"supply_id"`
	Code string `json:"code"`
}

type V1FbpArchiveListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type FbpDraftGetResponseDeclineReason struct {
	FailedSKUIds []interface{} `json:"failed_sku_ids"`
	Message string `json:"message"`
}

type V2ProductInfoPicturesResponse struct {
	Items []interface{} `json:"items"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexData struct {
	MinPrice interface{} `json:"min_price"`
	PriceIndex float64 `json:"price_index"`
	URL string `json:"url"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFeeFee struct {
	Accrued interface{} `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type V1OrderValidationError struct {
	OrderErrors []interface{} `json:"order_errors"`
}

type DirectDetailsByTplDetails struct {
	TransportCompanyName string `json:"transport_company_name"`
	TrackingNumber string `json:"tracking_number"`
}

type FinanceTransactionListV3ResponseOperation struct {
	DeliveryCharge float64 `json:"delivery_charge"`
	Items []interface{} `json:"items"`
	OperationID int64 `json:"operation_id"`
	OperationType string `json:"operation_type"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"`
	SaleCommission float64 `json:"sale_commission"`
	Services []interface{} `json:"services"`
	AccrualsForSale float64 `json:"accruals_for_sale"`
	Amount float64 `json:"amount"`
	OperationDate string `json:"operation_date"`
	OperationTypeName string `json:"operation_type_name"`
	Posting interface{} `json:"posting"`
	Type string `json:"type_"`
}

type ReportCreateCompanyProductsReportRequestVisibility string

type V4GetProductAttributesResponseModelInfo struct {
	Count int64 `json:"count"`
	ModelID int64 `json:"model_id"`
}

type ProductGetProductInfoDescriptionResponseProduct struct {
	OfferID string `json:"offer_id"`
	Description string `json:"description"`
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseRequest struct {
	Search interface{} `json:"search"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ItemBundleSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemCount int64 `json:"total_item_count"`
	TotalQuantity int64 `json:"total_quantity"`
}

type V1SellerActionsCreateVoucherRequest struct {
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
	VoucherParameters interface{} `json:"voucher_parameters"`
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
}

type ProductImportProductsPricesResponse struct {
	Result []interface{} `json:"result"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplarMark struct {
	CheckStatus string `json:"check_status"`
	ErrorCodes []interface{} `json:"error_codes"`
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V1GetStrategyIDsByItemIDsResponseResult struct {
	ProductsInfo []interface{} `json:"products_info"`
}

type GetUploadQuotaResponseDailyCreate struct {
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
	Usage int64 `json:"usage"`
}

type V1QuestionListResponseQuestions struct {
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
	AnswersCount int64 `json:"answers_count"`
	PublishedAt interface{} `json:"published_at"`
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
	ProductURL string `json:"product_url"`
	QuestionLink string `json:"question_link"`
}

type V2FbsPostingProduct struct {
	Price string `json:"price"`
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type V3FinanceCashFlowStatementListResponsePeriod struct {
	ID int64 `json:"id"`
	Begin string `json:"begin"`
	End string `json:"end"`
}

type V1ProductUpdateDiscountResponse struct {
	Result bool `json:"result"`
}

type FilterWithQuant struct {
	Created bool `json:"created"`
	Exists bool `json:"exists"`
}

type MoneyMoneySalePrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1StrategyRequest struct {
	StrategyID string `json:"strategy_id"`
}

type ReportCreateDiscountedRequest interface{}

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPoint struct {
	Address string `json:"address"`
	AddressCoordinates interface{} `json:"address_coordinates"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type V4GetProductInfoStocksResponseItem struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Stocks []interface{} `json:"stocks"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejected struct {
	Code interface{} `json:"code"`
	ProductID int64 `json:"product_id"`
	Reason string `json:"reason"`
}

type Fbsv4FbsPostingShipV4Request struct {
	Packages interface{} `json:"packages"`
	PostingNumber string `json:"posting_number"`
	With interface{} `json:"with"`
}

type V1SellerActionsCreateVoucherRequestVoucherParameterTypeEnum string

type DirectDetailsBySellerDetails struct {
	DriverName string `json:"driver_name"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
	VehicleType string `json:"vehicle_type"`
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

type AnalyticsProductQueriesRequestSortBy string

type V1SellerActionsUpdateDiscountRequest struct {
	ActionParameters interface{} `json:"action_parameters"`
	ActionID int64 `json:"action_id"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem struct {
	BreakHours interface{} `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"`
	OpeningHours interface{} `json:"opening_hours"`
}

type V1ReviewListResponse struct {
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
	HasNext bool `json:"has_next"`
}

type SellerReturnsv1MoneyUtilization struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	Dir string `json:"dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
}

type V3Dimensions struct {
	Weight string `json:"weight"`
	Width string `json:"width"`
	Height string `json:"height"`
	Length string `json:"length"`
}

type DeleteProductsRequestProduct struct {
	OfferID string `json:"offer_id"`
}

type GetReturnsListResponseCompensationStatus struct {
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
	SysName string `json:"sys_name"`
}

type ProductV1ProductVisibilitySetResponseItemsSelectPermissionEnum string

type SellerReturnsv1MoneyProduct struct {
	Price float64 `json:"price"`
	CurrencyCode string `json:"currency_code"`
}

type Productv3GetProductListRequest struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V1SearchAttributeValuesResponse struct {
	Result []interface{} `json:"result"`
}

type RpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1PostingFbsSplitRequest struct {
	PostingNumber string `json:"posting_number"`
	Postings []interface{} `json:"postings"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponseCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1GetAttributeValuesRequest struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	Language interface{} `json:"language"`
	LastValueID int64 `json:"last_value_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	AttributeID int64 `json:"attribute_id"`
}

type V1DeclineDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V1FbpOrderListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type V1SellerActionsProductsCandidatesResponseProductQuantTypeEnum string

type RelatedPostingCancelReasons struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	TypeID string `json:"type_id"`
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []interface{} `json:"related_posting_numbers"`
}

type GetProductRatingBySkuResponseRatingGroup struct {
	ImproveAttributes interface{} `json:"improve_attributes"`
	Key string `json:"key"`
	Name string `json:"name"`
	Rating float64 `json:"rating"`
	Weight float64 `json:"weight"`
	Conditions interface{} `json:"conditions"`
	ImproveAtLeast int64 `json:"improve_at_least"`
}

type V1FbpDraftDropOffRegistrateResponse struct {
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
	Error interface{} `json:"error"`
}

type V1SetPostingsResponse struct {
	Result interface{} `json:"result"`
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"`
}

type V1AddBarcodeRequest struct {
	Barcodes []interface{} `json:"barcodes"`
}

type GetReturnsListResponseAdditionalInfo struct {
	IsOpened bool `json:"is_opened"`
	IsSuperEconom bool `json:"is_super_econom"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	ToUpdate []interface{} `json:"to_update"`
}

type V1AssemblyCarriagePostingListResponsePosting struct {
	CanPrintLabel bool `json:"can_print_label"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	AssemblyCode string `json:"assembly_code"`
}

type ChatChatSendMessageResponse struct {
	Result string `json:"result"`
}

type V4GetUploadQuotaResponse struct {
	DailyCreate interface{} `json:"daily_create"`
	DailyUpdate interface{} `json:"daily_update"`
	OperationLimits interface{} `json:"operation_limits"`
	Total interface{} `json:"total"`
}

type V2PostingFBSActGetProducts struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1AssemblyCarriagePostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type Productv3GetProductListResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftDirectProductValidateRequestSkuItem struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type Productv3GetProductListRequestFilter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type ArrivalpassArrivalPassListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type V1GetAttributesRequest struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	Language interface{} `json:"language"`
	TypeID int64 `json:"type_id"`
}

type V1FbpDraftDirectProductValidateResponseApprovedItem struct {
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
}

type V3PostingFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
}

type V1FbpDraftPickUpRegistrateResponse struct {
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
	Error interface{} `json:"error"`
}

type V3ChatHistoryRequest struct {
	Filter interface{} `json:"filter"`
	FromMessageID int64 `json:"from_message_id"`
	Limit int64 `json:"limit"`
	ChatID string `json:"chat_id"`
	Direction string `json:"direction"`
}

type V1AssemblyCarriagePostingListResponse struct {
	CanPrintMassLabel bool `json:"can_print_mass_label"`
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type V1AssemblyFbsPostingListResponsePostingProduct struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1GetFinanceBalanceV1ResponseServicesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountTypeEnum string

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1ProductInfoWrongVolumeRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1TimeRangeStorageTariffication struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type Postingv3GetFbsPostingUnfulfilledListResponseResult struct {
	Postings []interface{} `json:"postings"`
	Count int64 `json:"count"`
}

type GetCarriageAvailableListResponseResult struct {
	TPLProviderIconURL string `json:"tpl_provider_icon_url"`
	CarriageID int64 `json:"carriage_id"`
	FirstMileType string `json:"first_mile_type"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DeliveryMethodName string `json:"delivery_method_name"`
	MandatoryPackagedCount int32 `json:"mandatory_packaged_count"`
	TPLProviderName string `json:"tpl_provider_name"`
	WarehouseTimezone string `json:"warehouse_timezone"`
	RecommendedTimeUtcOffsetInMinutes float64 `json:"recommended_time_utc_offset_in_minutes"`
	CarriagePostingsCount int32 `json:"carriage_postings_count"`
	CarriageStatus string `json:"carriage_status"`
	CutoffAt string `json:"cutoff_at"`
	Errors interface{} `json:"errors"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	RecommendedTimeLocal string `json:"recommended_time_local"`
	WarehouseCity string `json:"warehouse_city"`
	MandatoryPostingsCount int32 `json:"mandatory_postings_count"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplar struct {
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	RnptCheckStatus string `json:"rnpt_check_status"`
	RnptErrorCodes []interface{} `json:"rnpt_error_codes"`
	ExemplarID int64 `json:"exemplar_id"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	GTD string `json:"gtd"`
	GTDCheckStatus string `json:"gtd_check_status"`
	GTDErrorCodes []interface{} `json:"gtd_error_codes"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
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

type V1CommentListResponse struct {
	Comments []interface{} `json:"comments"`
	Offset int32 `json:"offset"`
}

type ProductImportProductsBySKUResponse struct {
	Result interface{} `json:"result"`
}

type V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type V1SellerActionsProductsAddRequest struct {
	ActionID int64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type ChatMessageModerateImageStatus string

type V1SellerActionsUpdateMultiLevelDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type Productv3Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility interface{} `json:"visibility"`
}

type PriceIndexesIndexDataSelf struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type DetailsService struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type V1CreateWarehouseFBSRequest struct {
	CutInTime int64 `json:"cut_in_time"`
	FirstMileType interface{} `json:"first_mile_type"`
	IsKgt bool `json:"is_kgt"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	TimeslotID int64 `json:"timeslot_id"`
	AddressCoordinates interface{} `json:"address_coordinates"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	Options interface{} `json:"options"`
	WorkingDays []interface{} `json:"working_days"`
}

type DetailsServices struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type V1SearchQueriesTextResponse struct {
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
	Offset string `json:"offset"`
}

type V1ReturnsCompanyFbsInfoRequest struct {
	Filter interface{} `json:"filter"`
	Pagination interface{} `json:"pagination"`
}

type GetProductAttributesResponsePdf struct {
	Name string `json:"name"`
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
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

type FinanceV1GetFinanceAccrualByDayResponseAccrual struct {
	AccruedCategory interface{} `json:"accrued_category"`
	Date string `json:"date"`
	ItemFees interface{} `json:"item_fees"`
	NonItemFee interface{} `json:"non_item_fee"`
	Posting interface{} `json:"posting"`
	TotalAmount interface{} `json:"total_amount"`
	AccrualID int64 `json:"accrual_id"`
	UnitNumber string `json:"unit_number"`
}

type ProductV4GetUploadQuotaResponseTotalQuotaByCategory struct {
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
	CategoryID int64 `json:"category_id"`
}

type WarehouseFirstMileType struct {
	FirstMileType string `json:"first_mile_type"`
	DropoffPointID string `json:"dropoff_point_id"`
	DropoffTimeslotID int64 `json:"dropoff_timeslot_id"`
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
}

type SellerActionsUpdateMultiLevelDiscountRequestActionParametersDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type V1FbpOrderPickUpDlvEditResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"`
}

type V1ProductUpdateAttributesResponse struct {
	TaskID int64 `json:"task_id"`
}

type V1ProductGetRelatedSKUResponseError struct {
	Code string `json:"code"`
	SKU int64 `json:"sku"`
	Message string `json:"message"`
}

type V1ReviewInfoRequest struct {
	ReviewID string `json:"review_id"`
}

type V1SearchQueriesTopRequest struct {
	Limit string `json:"limit"`
	Offset string `json:"offset"`
}

type SellerInfoResponseCompany struct {
	TaxSystem interface{} `json:"tax_system"`
	Country string `json:"country"`
	Currency string `json:"currency"`
	Inn string `json:"inn"`
	LegalName string `json:"legal_name"`
	Name string `json:"name"`
	Ogrn string `json:"ogrn"`
	OwnershipForm string `json:"ownership_form"`
}

type V1GetStrategyItemsResponseResult struct {
	ProductID []interface{} `json:"product_id"`
}

type WarehouseListV2ResponseWarehouse struct {
	CarriageLabelType interface{} `json:"carriage_label_type"`
	FirstMile interface{} `json:"first_mile"`
	IsKgt bool `json:"is_kgt"`
	Name string `json:"name"`
	SlaCutIn int64 `json:"sla_cut_in"`
	UpdatedAt string `json:"updated_at"`
	WithItemList bool `json:"with_item_list"`
	CutInTime int64 `json:"cut_in_time"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsComfort bool `json:"is_comfort"`
	Timetable interface{} `json:"timetable"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseType string `json:"warehouse_type"`
	AddressInfo interface{} `json:"address_info"`
	CreatedAt string `json:"created_at"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	IsExpress bool `json:"is_express"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
	Phone string `json:"phone"`
	Status string `json:"status"`
	CourierComment string `json:"courier_comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	IsRFBS bool `json:"is_rfbs"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	PostingsLimit int32 `json:"postings_limit"`
	WorkingDays []interface{} `json:"working_days"`
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"`
}

type ProductGetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type V1FbpOrderPickUpCancelResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type Productv2ProductsStocksResponse struct {
	Result []interface{} `json:"result"`
}

type CreateWarehouseFBSRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ReturnsRfbsGetV2ResponseRejectionReason struct {
	Hint string `json:"hint"`
	ID int32 `json:"id"`
	IsCommentRequired bool `json:"is_comment_required"`
	Name string `json:"name"`
}

type V1QuestionAnswerDeleteRequest struct {
	AnswerID string `json:"answer_id"`
	SKU int64 `json:"sku"`
}

type Productv1ProductInfoPicturesResponseResult struct {
	Pictures interface{} `json:"pictures"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1FbpOrderDropOffCancelResponse struct {
	RowVersion int64 `json:"row_version"`
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
}

type SellerSellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CarriageID int64 `json:"carriage_id"`
}

type Productv3GetProductListResponseItemQuant struct {
	QuantCode string `json:"quant_code"`
	QuantSize int64 `json:"quant_size"`
}

type V3ChatListResponse struct {
	TotalUnreadCount int64 `json:"total_unread_count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Chats interface{} `json:"chats"`
}

type MoneyMoneyTotalAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type GetReturnsListResponsePlaceTarget struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type V1FbpDraftListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1AssemblyCarriagePostingListRequestFilter struct {
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1SellerActionsUpdateDiscountWithConditionRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1AssemblyFbsPostingListResponse struct {
	Cursor string `json:"cursor"`
	Cutoff string `json:"cutoff"`
	Postings []interface{} `json:"postings"`
}

type ActionsV1ActionsAutoAddProductsDeleteRequest struct {
	ProductIds []interface{} `json:"product_ids"`
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
}

type V1FbpDraftDirectTplDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type PriceIndexesIndexSelfData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type ChatChatSendMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text string `json:"text"`
}

type V1FbpDraftDropOffCreateRequest struct {
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
}

type Productv4Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
	Visibility interface{} `json:"visibility"`
}

type GetProductRatingBySkuResponseRatingImproveAttribute struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type SearchQueriesTextRequestSortDir string

type V1WarehouseFbsCreatePickUpTimeslotListRequest struct {
	AddressCoordinates interface{} `json:"address_coordinates"`
	IsKgt bool `json:"is_kgt"`
}

type FbsPostingProductExemplarSetV6RequestExemplars struct {
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	ExemplarID int64 `json:"exemplar_id"`
}

type PostingV4PostingFbsListResponsePostingsDeliveryMethod struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ItemResponse struct {
	TotalVolumeInLitres float64 `json:"total_volume_in_litres"`
	Tags []interface{} `json:"tags"`
	IconPath string `json:"icon_path"`
	Name string `json:"name"`
	Barcode string `json:"barcode"`
	ProductID int64 `json:"product_id"`
	Quant int32 `json:"quant"`
	VolumeInLitres float64 `json:"volume_in_litres"`
	ContractorItemCode string `json:"contractor_item_code"`
	PlacementZone string `json:"placement_zone"`
	SKU int64 `json:"sku"`
	SfboAttribute interface{} `json:"sfbo_attribute"`
	ShipmentType interface{} `json:"shipment_type"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	IsQuantEditable bool `json:"is_quant_editable"`
}

type V1FbpEditTimeslotResponseReserveFailureType string

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductCommission struct {
	SaleAmount interface{} `json:"sale_amount"`
	SaleCommission interface{} `json:"sale_commission"`
	SalePrice interface{} `json:"sale_price"`
	SellerPrice interface{} `json:"seller_price"`
	Bonus interface{} `json:"bonus"`
	Coinvestment interface{} `json:"coinvestment"`
	Commission interface{} `json:"commission"`
	CommissionRatio string `json:"commission_ratio"`
}

type V1GetCompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type ItemMarketing struct {
	Actions []interface{} `json:"actions"`
}

type StockShipmentType string

type ApproveDeclineDiscountTasksResponseFailDetail struct {
	TaskID int64 `json:"task_id"`
	ErrorForUser string `json:"error_for_user"`
}

type FbsPostingProductExemplarSetV6RequestProducts struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type SellerServiceanalyticsDimension string

type V1SellerActionsCreateMultiLevelDiscountRequest struct {
	DiscountType interface{} `json:"discount_type"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
}

type V1fbpTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
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

type V1WarehouseFbsUpdateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ValidationResultValidationErrorTypeEnum string

type V6FbsPostingProductExemplarCreateOrGetV6Response struct {
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type ListFBSRatingIndexPostingsV1RequestFilter struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1CancellationStateStatus string

type SellerSellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type V3FbsPostingDetailRelatedPostings struct {
	RelatedPostingNumbers interface{} `json:"related_posting_numbers"`
}

type ReviewV2ReviewListV2RequestFiltersStatusEnum string

type ChatInfoChatTypeEnum string

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1UnarchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V2ReturnsRfbsListV2ResponseState struct {
	GroupState string `json:"group_state"`
	MoneyReturnStateName string `json:"money_return_state_name"`
	State string `json:"state"`
	StateName string `json:"state_name"`
}

type V1GetReturnsListResponse struct {
	Returns []interface{} `json:"returns"`
	HasNext bool `json:"has_next"`
}

type PickedSegmentSegment struct {
	Description string `json:"description"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	Type interface{} `json:"type_"`
}

type V1ProductGetRelatedSKUResponse struct {
	Items interface{} `json:"items"`
	Errors interface{} `json:"errors"`
}

type V2ReturnsRfbsRejectRequest struct {
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
	RejectionReasonID int64 `json:"rejection_reason_id"`
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []interface{} `json:"tracking_numbers"`
}

type V1FbpOrderGetResponse struct {
	Locked bool `json:"locked"`
	Status interface{} `json:"status"`
	CancellationState interface{} `json:"cancellation_state"`
	CreatedDate string `json:"created_date"`
	OrderNumber string `json:"order_number"`
	ReceiveDate string `json:"receive_date"`
	RowVersion int64 `json:"row_version"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	DeliveryDetails interface{} `json:"delivery_details"`
	DraftID int64 `json:"draft_id"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	HasLabel bool `json:"has_label"`
	SupplyID string `json:"supply_id"`
	BundleUuid string `json:"bundle_uuid"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	ID int64 `json:"id"`
}

type V1WarehouseFbsCreatePickUpTimeslotListResponse struct {
	IsPickupSupported bool `json:"is_pickup_supported"`
	Timeslots []interface{} `json:"timeslots"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProduct struct {
	SKU int64 `json:"sku"`
	Commission interface{} `json:"commission"`
	Delivery interface{} `json:"delivery"`
}

type FinanceCashFlowStatementListResponseDetailsOthers struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V3GetProductInfoListResponsePromotionType string

type V1WarehouseFbsUpdateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type GetProductInfoListResponseSource struct {
	CreatedAt string `json:"created_at"`
	QuantCode string `json:"quant_code"`
	ShipmentType interface{} `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementListEnum string

type GetProductAttributesV4ResponseAttribute struct {
	ID int64 `json:"id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

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

type DropOffPointCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1FbpDraftPickupCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails interface{} `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ReviewV2ReviewListV2RequestFiltersOrderStatusEnum string

type V1GetProductStairwayDiscountByQuantityResponseStairways struct {
	Enabled bool `json:"enabled"`
	SKU int64 `json:"sku"`
	Stairway interface{} `json:"stairway"`
	Status interface{} `json:"status"`
}

type PostingV4PostingFbsListResponsePostingsTarifficationStep struct {
	TariffCharge interface{} `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"`
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
	MinCharge interface{} `json:"min_charge"`
}

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type V1FbpDraftPickUpRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type Productv5GetProductInfoPricesV5Request struct {
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
	Cursor string `json:"cursor"`
}

type V4GetProductInfoStocksRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
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

type V1AssemblyCarriageProductListResponse struct {
	Products []interface{} `json:"products"`
	Cursor string `json:"cursor"`
}

type GetReturnsListResponseVisualStatus struct {
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
	SysName string `json:"sys_name"`
}

type SellerApiGetSellerProductV1ResponseResult struct {
	LastID float64 `json:"last_id"`
	Products []interface{} `json:"products"`
	Total float64 `json:"total"`
}

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPointDropOffPointTypeEnum string

type AnalyticsGetDataResponseResult struct {
	Data []interface{} `json:"data"`
	Totals []interface{} `json:"totals"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexDataSelf struct {
	MinPrice interface{} `json:"min_price"`
	PriceIndex float64 `json:"price_index"`
	URL string `json:"url"`
}

type V5FbsPostingProductExemplarStatusV5Response struct {
	Products []interface{} `json:"products"`
	Status string `json:"status"`
	PostingNumber string `json:"posting_number"`
}

type ProductInfoWrongVolumeResponseWrongVolumeProduct struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	Weight int64 `json:"weight"`
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	Length int64 `json:"length"`
}

type V1SellerActionsCreateInstallmentRequest struct {
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type PostingV4PostingFbsListResponsePostingsCustomer struct {
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address interface{} `json:"address"`
}

type V1DeliveryDetailsDirectDetails struct {
	ByTPLDetails interface{} `json:"by_tpl_details"`
	TimeslotDetails interface{} `json:"timeslot_details"`
	BySellerDetails interface{} `json:"by_seller_details"`
}

type CreateWarehouseFBSRequestOptions struct {
	Comment string `json:"comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
}

type GetReturnsListResponseLogistic struct {
	TechnicalReturnMoment string `json:"technical_return_moment"`
	FinalMoment string `json:"final_moment"`
	CancelledWithCompensationMoment string `json:"cancelled_with_compensation_moment"`
	ReturnDate string `json:"return_date"`
	Barcode string `json:"barcode"`
}

type SellerSellerAPIArrivalPassUpdateRequestArrivalPass struct {
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	ID int64 `json:"id"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponse struct {
	Products []interface{} `json:"products"`
	Total int64 `json:"total"`
}

type AnalyticsAnalyticsGetDataResponse struct {
	Result interface{} `json:"result"`
	Timestamp string `json:"timestamp"`
}

type RowItem struct {
	Barcode string `json:"barcode"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	SKU int64 `json:"sku"`
}

type SearchQueriesTextRequestSortBy string

type V1AssemblyCarriageProductListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type ProductV1ProductVisibilitySetResponse struct {
	Items []interface{} `json:"items"`
	ItemsErrors []interface{} `json:"items_errors"`
}

type V3FbsTariffication struct {
	CurrentTariffChargeCurrencyCode string `json:"current_tariff_charge_currency_code"`
	NextTariffType string `json:"next_tariff_type"`
	NextTariffChargeCurrencyCode string `json:"next_tariff_charge_currency_code"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	CurrentTariffCharge string `json:"current_tariff_charge"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffCharge string `json:"next_tariff_charge"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	CurrentTariffType string `json:"current_tariff_type"`
}

type MoneyMoneyTotalAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1FbpOrderListResponseItem struct {
	OrderNumber string `json:"order_number"`
	ReceiveDate string `json:"receive_date"`
	Status interface{} `json:"status"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	BundleSummary interface{} `json:"bundle_summary"`
	DeliveryDetails interface{} `json:"delivery_details"`
	Locked bool `json:"locked"`
	CancellationState interface{} `json:"cancellation_state"`
	CreatedDate string `json:"created_date"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	HasLabel bool `json:"has_label"`
	PackageUnitsCount int32 `json:"package_units_count"`
	SupplyID string `json:"supply_id"`
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
}

type ItemSize struct {
	HeightMm int32 `json:"height_mm"`
	LengthMm int32 `json:"length_mm"`
	WidthMm int32 `json:"width_mm"`
}

type SellerInfoResponseSubscriptionTypeEnum string

type MoneyMoneyCoinvestment struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type V1FbpDraftDirectRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError interface{} `json:"order_error"`
}

type ChatHistoryRequestFilter struct {
	MessageIds []interface{} `json:"message_ids"`
}

type ActionsV1ActionsAutoAddProductsDeleteResponse struct {
	ProductIds []interface{} `json:"product_ids"`
}

type ProductImportProductsPricesResponseProcessResult struct {
	Updated bool `json:"updated"`
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
}

type V1ProductActionTimerUpdateRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1FbpArchiveGetResponse struct {
	SupplyID string `json:"supply_id"`
	BundleSKUSummary interface{} `json:"bundle_sku_summary"`
	HasAct bool `json:"has_act"`
	RowVersion int64 `json:"row_version"`
	Status interface{} `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	ActFileUuid string `json:"act_file_uuid"`
	DeclineReason interface{} `json:"decline_reason"`
	HasLabel bool `json:"has_label"`
	PackageUnitsCount int32 `json:"package_units_count"`
	ReceiveDate string `json:"receive_date"`
	BusinessFlowTypeID int64 `json:"business_flow_type_id"`
	BundleID string `json:"bundle_id"`
	CreatedDate string `json:"created_date"`
	DeliveryDetails interface{} `json:"delivery_details"`
	ID int64 `json:"id"`
	OrderDraftID int64 `json:"order_draft_id"`
	OrderNumber string `json:"order_number"`
}

type ReasonHumanText struct {
	Text string `json:"text"`
}

type Productv5GetProductInfoPricesV5Response struct {
	Cursor string `json:"cursor"`
	Items interface{} `json:"items"`
	Total int32 `json:"total"`
}

type V1FbpDraftDirectRegistrateResponse struct {
	Error interface{} `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1DeclineDiscountTasksRequestTask struct {
	SellerComment string `json:"seller_comment"`
	ID int64 `json:"id"`
}

type V1SetProductStairwayDiscountByQuantityRequest struct {
	Stairways []interface{} `json:"stairways"`
	SuppressWarnings bool `json:"suppress_warnings"`
}

type ProductProductInfoPicturesResponsePicture struct {
	Is360 bool `json:"is_360"`
	IsColor bool `json:"is_color"`
	IsPrimary bool `json:"is_primary"`
	ProductID int64 `json:"product_id"`
	State string `json:"state"`
	URL string `json:"url"`
}

type GetReturnsListResponseStorage struct {
	UtilizationSum interface{} `json:"utilization_sum"`
	UtilizationForecastDate string `json:"utilization_forecast_date"`
	Sum interface{} `json:"sum"`
	TarifficationFirstDate string `json:"tariffication_first_date"`
	TarifficationStartDate string `json:"tariffication_start_date"`
	ArrivedMoment string `json:"arrived_moment"`
	Days int64 `json:"days"`
}

type V1SetProductStairwayDiscountByQuantityResponse struct {
	Accepted bool `json:"accepted"`
	Errors []interface{} `json:"errors"`
	Warnings []interface{} `json:"warnings"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type GetProductRatingBySkuResponseRatingCondition struct {
	Fulfilled bool `json:"fulfilled"`
	Key string `json:"key"`
	Cost float64 `json:"cost"`
	Description string `json:"description"`
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
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

type DeliveryMethodListV2ResponseDeliveryMethod struct {
	CreatedAt string `json:"created_at"`
	IsExpress bool `json:"is_express"`
	Name string `json:"name"`
	ProviderID int64 `json:"provider_id"`
	SlaCutIn int64 `json:"sla_cut_in"`
	TemplateID int64 `json:"template_id"`
	TPLDropoffPoint interface{} `json:"tpl_dropoff_point"`
	UpdatedAt string `json:"updated_at"`
	Cutoff string `json:"cutoff"`
	ID int64 `json:"id"`
	Status string `json:"status"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftPickupDlvEditRequestDeliveryDetails struct {
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type ProductImportProductsBySKUResponseResult struct {
	TaskID int64 `json:"task_id"`
	UnmatchedSKUList []interface{} `json:"unmatched_sku_list"`
}

type V1ReturnsCompanyFbsInfoRequestFilter struct {
	PlaceID int64 `json:"place_id"`
}

type WarehouseWarehouseListResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type V1SellerActionsCreateDiscountWithConditionRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
}

type V1FbpOrderListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type V1QuestionAnswerCreateRequest struct {
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
}

type V3FinanceCashFlowStatementListResponse struct {
	Result interface{} `json:"result"`
}

type GetFinanceBalanceV1ResponseCashflows struct {
	Returns interface{} `json:"returns"`
	Sales interface{} `json:"sales"`
	Services []interface{} `json:"services"`
}

type V1UpdateStatusStrategyRequest struct {
	Enabled bool `json:"enabled"`
	StrategyID string `json:"strategy_id"`
}

type V1OrderAttentionTypeEnum string

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type V2CancellationStateEnum string

type GetSupplyOrderBundleRequestItemTagsCalculation struct {
	DropoffWarehouseID string `json:"dropoff_warehouse_id"`
	StorageWarehouseIds []interface{} `json:"storage_warehouse_ids"`
}

type Productv2ProductsStocksResponseError struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

type ResultError struct {
	Code string `json:"code"`
	Status string `json:"status"`
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V2ReturnsRfbsGetRequest struct {
	ReturnID int64 `json:"return_id"`
}

type V1FbpDraftPickupCreateResponse struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
}

type V2CancellationStateEnumFilters string

type PostingV4PostingFbsUnfulfilledListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir interface{} `json:"sort_dir"`
	Translit bool `json:"translit"`
	With interface{} `json:"with"`
}

type ReviewInfoResponseVideo struct {
	ShortVideoPreviewURL string `json:"short_video_preview_url"`
	URL string `json:"url"`
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	PreviewURL string `json:"preview_url"`
}

type V1GetRealizationReportPostingRequest struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type V1ProductActionTimerStatusResponseStatuses struct {
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	ProductID int64 `json:"product_id"`
	ExpiredAt string `json:"expired_at"`
}

type SellerReturnsv1MoneyStorage struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type ArrivalPassListRequestFilter struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	ArrivalReason string `json:"arrival_reason"`
	DropoffPointIds []interface{} `json:"dropoff_point_ids"`
	OnlyActivePasses bool `json:"only_active_passes"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1FbpDraftGetResponse struct {
	WarehouseID int64 `json:"warehouse_id"`
	RowVersion int64 `json:"row_version"`
	DeclineReason interface{} `json:"decline_reason"`
	DeliveryDetails interface{} `json:"delivery_details"`
	IsDeletable bool `json:"is_deletable"`
	CancellationState interface{} `json:"cancellation_state"`
	Editable bool `json:"editable"`
	PackageUnitsCount int32 `json:"package_units_count"`
	BundleID string `json:"bundle_id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	ID int64 `json:"id"`
	IsCancelable bool `json:"is_cancelable"`
	SupplyID string `json:"supply_id"`
	IsRegistrationAvailable bool `json:"is_registration_available"`
	Locked bool `json:"locked"`
	Status interface{} `json:"status"`
}

type AnalyticsMetric string

type ProductProductUnarchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type ChatStartResponseResult struct {
	ChatID string `json:"chat_id"`
}

type V1GetFinanceBalanceV1ResponseReturnsMoney struct {
	Value float64 `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type V1CarriageCancelResponse struct {
	Error string `json:"error"`
	CarriageStatus string `json:"carriage_status"`
}

type V1DayOfWeek string

type ProductV1ProductVisibilitySetRequestItemPlacementPlacementEnum string

type RpcStatusV1PolygonCreate struct {
	Details []interface{} `json:"details"`
	Message string `json:"message"`
	Code int32 `json:"code"`
}

type GetReturnsListResponseReturnsItem struct {
	Type string `json:"type_"`
	OrderID int64 `json:"order_id"`
	Schema string `json:"schema"`
	Place interface{} `json:"place"`
	Storage interface{} `json:"storage"`
	Logistic interface{} `json:"logistic"`
	AdditionalInfo interface{} `json:"additional_info"`
	ReturnClearingID int64 `json:"return_clearing_id"`
	Exemplars []interface{} `json:"exemplars"`
	ID int64 `json:"id"`
	CompanyID int64 `json:"company_id"`
	ReturnReasonName string `json:"return_reason_name"`
	OrderNumber string `json:"order_number"`
	TargetPlace interface{} `json:"target_place"`
	SourceID int64 `json:"source_id"`
	ClearingID int64 `json:"clearing_id"`
	Visual interface{} `json:"visual"`
	PostingNumber string `json:"posting_number"`
	CompensationStatus interface{} `json:"compensation_status"`
	Product interface{} `json:"product"`
}

type V1GetProductStairwayDiscountByQuantityResponse struct {
	Stairways []interface{} `json:"stairways"`
}

type PostingV4PostingFbsListResponsePostingsFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
}

type V1ProductInfoWarehouseStocksResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Stocks []interface{} `json:"stocks"`
}

type ReportCreateCompanyPostingsReportRequest struct {
	Filter interface{} `json:"filter"`
	Language interface{} `json:"language"`
	With interface{} `json:"with"`
}

type V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason string

type ListDropOffPointsForCreateFBSWarehouseResponseDropOffPoint struct {
	Address string `json:"address"`
	Coordinates interface{} `json:"coordinates"`
	DiscountPercent float64 `json:"discount_percent"`
	ID string `json:"id"`
	LastTransitTimeLocal interface{} `json:"last_transit_time_local"`
	Type interface{} `json:"type_"`
}

type V1ProductGetRelatedSKURequest struct {
	SKU interface{} `json:"sku"`
}

type V1FbpDraftDirectSellerDlvCreateRequestDirectDetails struct {
	DriverName string `json:"driver_name"`
	TimeslotStart string `json:"timeslot_start"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type PostinglistV3status struct {
	From string `json:"from"`
	To string `json:"to"`
}

type PostingV1PostingFbpListResponse struct {
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type V1QuestionListResponse struct {
	Questions interface{} `json:"questions"`
	LastID string `json:"last_id"`
	HasNext bool `json:"has_next"`
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []interface{} `json:"result"`
}

type GetConditionalCancellationListV2ResponseState struct {
	State interface{} `json:"state"`
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type SellerActionsListRequestActionTypeEnum string

type V1QuestionInfoResponse struct {
	AnswersCount int64 `json:"answers_count"`
	AuthorName string `json:"author_name"`
	ProductURL string `json:"product_url"`
	PublishedAt interface{} `json:"published_at"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
	ID string `json:"id"`
	QuestionLink string `json:"question_link"`
	SKU int64 `json:"sku"`
}

type V1FbpAvailableTimeslotListResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type AnalyticsDataRow struct {
	Metrics []interface{} `json:"metrics"`
	Dimensions []interface{} `json:"dimensions"`
}

type DetailsPayment struct {
	CurrencyCode string `json:"currency_code"`
	Payment float64 `json:"payment"`
}

type V1ReviewCountResponse struct {
	Total int32 `json:"total"`
	Unprocessed int32 `json:"unprocessed"`
	Processed int32 `json:"processed"`
}

type ItemCommissionsv5 struct {
	SalesPercentFBP float64 `json:"sales_percent_fbp"`
	SalesPercentFBS float64 `json:"sales_percent_fbs"`
	SalesPercentRFBS float64 `json:"sales_percent_rfbs"`
	FBODirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
	FBODirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
	FBSDelivToCustomerAmount float64 `json:"fbs_deliv_to_customer_amount"`
	FBSDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
	FBSDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
	FBSFirstMileMaxAmount float64 `json:"fbs_first_mile_max_amount"`
	FBSFirstMileMinAmount float64 `json:"fbs_first_mile_min_amount"`
	FBSReturnFlowAmount float64 `json:"fbs_return_flow_amount"`
	FBODelivToCustomerAmount float64 `json:"fbo_deliv_to_customer_amount"`
	FBOReturnFlowAmount float64 `json:"fbo_return_flow_amount"`
	SalesPercentFBO float64 `json:"sales_percent_fbo"`
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V1QuestionTopSkuResponse struct {
	SKU interface{} `json:"sku"`
}

type V1GetDiscountTaskListResponse struct {
	Result []interface{} `json:"result"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseBelowMinPrice struct {
	Value float64 `json:"value"`
	Key int64 `json:"key"`
}

type CreateReportResponse struct {
	Result interface{} `json:"result"`
}

type ProductV4GetUploadQuotaResponseOperationLimits struct {
	Limit int64 `json:"limit"`
	LimitType interface{} `json:"limit_type"`
}

type V1ReviewListRequest struct {
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortDir string `json:"sort_dir"`
	Status string `json:"status"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type Financev3FinanceTransactionListV3ResponseResult struct {
	Operations []interface{} `json:"operations"`
	PageCount int64 `json:"page_count"`
	RowCount int64 `json:"row_count"`
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V3GetFbsPostingListResponseV3Result struct {
	HasNext bool `json:"has_next"`
	Postings []interface{} `json:"postings"`
}

type V1DescriptionCategoryTipsResponse struct {
	Result []interface{} `json:"result"`
}

type AnalyticsProductQueriesResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type RowItemCommissionReturn struct {
	Amount float64 `json:"amount"`
	Compensation float64 `json:"compensation"`
	StandardFee float64 `json:"standard_fee"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Total float64 `json:"total"`
	Bonus float64 `json:"bonus"`
	Commission float64 `json:"commission"`
	PricePerInstance float64 `json:"price_per_instance"`
	Quantity int32 `json:"quantity"`
	Stars float64 `json:"stars"`
}

type V2GetRealizationReportResponseV2 struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsUpdateDiscountRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type GetFinanceBalanceV1ResponseSalesDetails struct {
	PartnerPrograms interface{} `json:"partner_programs"`
	PointsForDiscounts string `json:"points_for_discounts"`
	Revenue interface{} `json:"revenue"`
}

type GetProductInfoListResponseStatuses struct {
	ModerateStatus string `json:"moderate_status"`
	StatusFailed string `json:"status_failed"`
	StatusUpdatedAt string `json:"status_updated_at"`
	ValidationStatus string `json:"validation_status"`
	IsCreated bool `json:"is_created"`
	Status string `json:"status"`
	StatusDescription string `json:"status_description"`
	StatusName string `json:"status_name"`
	StatusTooltip string `json:"status_tooltip"`
}

type V1GetProductRatingBySkuResponse struct {
	Products interface{} `json:"products"`
}

type ImportProductsRequestPdfList struct {
	Index int64 `json:"index"`
	Name string `json:"name"`
	SrcURL string `json:"src_url"`
}

type ChatChatSendFileResponse struct {
	Result string `json:"result"`
}

type FbpWarehouseListResponseWarehouse struct {
	PartnerName string `json:"partner_name"`
	SupplyTypes []interface{} `json:"supply_types"`
	TimezoneName string `json:"timezone_name"`
	AddressDetailing interface{} `json:"address_detailing"`
	ID int64 `json:"id"`
	IsBonded bool `json:"is_bonded"`
	Name string `json:"name"`
}

type V1PostingFbsSplitResponse struct {
	ParentPosting interface{} `json:"parent_posting"`
	Postings []interface{} `json:"postings"`
}

type CreateCompanyPostingsReportRequestWith struct {
	AdditionalData bool `json:"additional_data"`
	AnalyticsData bool `json:"analytics_data"`
	CustomerData bool `json:"customer_data"`
	JewelryCodes bool `json:"jewelry_codes"`
}

type SellerSellerAPIArrivalPassCreateRequestArrivalPass struct {
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
	DriverName string `json:"driver_name"`
}

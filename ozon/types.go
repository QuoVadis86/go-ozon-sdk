package ozon

type PostingV4PostingFbsListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1GetFinanceBalanceV1ResponseRevenueMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1GetFinanceBalanceV1ResponsePartnerMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type GetFinanceBalanceV1ResponseSalesDetails struct {
	Revenue V1GetFinanceBalanceV1ResponseRevenueMoney `json:"revenue"`
	PartnerPrograms V1GetFinanceBalanceV1ResponsePartnerMoney `json:"partner_programs"`
	PointsForDiscounts string `json:"points_for_discounts"`
}

type CreatedAt struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V2ReturnsRfbsFilter struct {
	OfferID string `json:"offer_id"`
	PostingNumber string `json:"posting_number"`
	GroupState []interface{} `json:"group_state"`
	CreatedAt CreatedAt `json:"created_at"`
}

type V2ReturnsRfbsListRequest struct {
	Filter V2ReturnsRfbsFilter `json:"filter"`
	LastID int32 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type AnalyticsGetDataResponseResult struct {
	Data []interface{} `json:"data"`
	Totals []interface{} `json:"totals"`
}

type AnalyticsAnalyticsGetDataResponse struct {
	Result AnalyticsGetDataResponseResult `json:"result"`
	Timestamp string `json:"timestamp"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponseProducts struct {
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
	MinActionQuantity int64 `json:"min_action_quantity"`
	OfferID string `json:"offer_id"`
	Price float64 `json:"price"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
	SKU int64 `json:"sku"`
	BasePrice float64 `json:"base_price"`
	Currency string `json:"currency"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Name string `json:"name"`
	ProductID int64 `json:"product_id"`
}

type V2GetProductInfoStocksByWarehouseFbsResponseV2 struct {
	Products []interface{} `json:"products"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
}

type V1ReturnsCompanyFbsInfoRequestFilter struct {
	PlaceID int64 `json:"place_id"`
}

type V1GetProductRatingBySkuRequest struct {
	Skus interface{} `json:"skus"`
}

type ChatHistoryRequestFilter struct {
	MessageIds []interface{} `json:"message_ids"`
}

type V2ReturnsRfbsGetV2ResponseState struct {
	State string `json:"state"`
	StateName string `json:"state_name"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsRequest struct {
	SKU interface{} `json:"sku"`
	OfferID interface{} `json:"offer_id"`
}

type PostingV4PostingFbsUnfulfilledListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
}

type DirectDetailsByTplDetails struct {
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type GetReturnsListResponseLogistic struct {
	FinalMoment string `json:"final_moment"`
	CancelledWithCompensationMoment string `json:"cancelled_with_compensation_moment"`
	ReturnDate string `json:"return_date"`
	Barcode string `json:"barcode"`
	TechnicalReturnMoment string `json:"technical_return_moment"`
}

type MoneyMoneySaleAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type MoneyMoneySaleCommission struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type MoneyMoneySalePrice struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type MoneyMoneySellerPrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type MoneyMoneyBonus struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type MoneyMoneyCoinvestment struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type MoneyMoneyCommission struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductCommission struct {
	Bonus MoneyMoneyBonus `json:"bonus"`
	Coinvestment MoneyMoneyCoinvestment `json:"coinvestment"`
	Commission MoneyMoneyCommission `json:"commission"`
	CommissionRatio string `json:"commission_ratio"`
	SaleAmount MoneyMoneySaleAmount `json:"sale_amount"`
	SaleCommission MoneyMoneySaleCommission `json:"sale_commission"`
	SalePrice MoneyMoneySalePrice `json:"sale_price"`
	SellerPrice MoneyMoneySellerPrice `json:"seller_price"`
}

type CreateWarehouseFBSRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateWarehouseFBSRequestFirstMileTypeEnum string

type CreateWarehouseFBSRequestOptions struct {
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	Comment string `json:"comment"`
	CourierPhones []interface{} `json:"courier_phones"`
}

type V1CreateWarehouseFBSRequest struct {
	AddressCoordinates CreateWarehouseFBSRequestAddressCoordinates `json:"address_coordinates"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	FirstMileType CreateWarehouseFBSRequestFirstMileTypeEnum `json:"first_mile_type"`
	IsKgt bool `json:"is_kgt"`
	Name string `json:"name"`
	Options CreateWarehouseFBSRequestOptions `json:"options"`
	TimeslotID int64 `json:"timeslot_id"`
	CutInTime int64 `json:"cut_in_time"`
	Phone string `json:"phone"`
	WorkingDays []interface{} `json:"working_days"`
}

type V1SellerActionsUpdateDiscountRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type ChatChatSendFileResponse struct {
	Result string `json:"result"`
}

type V1FbpDraftGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type ActionParameterVoucherParameterTypeEnum string

type ActionParameterVoucherParameter struct {
	IsPrivate bool `json:"is_private"`
	Type ActionParameterVoucherParameterTypeEnum `json:"type_"`
	CountCodes int64 `json:"count_codes"`
}

type V1ApproveDeclineDiscountTasksResponseResult struct {
	FailCount int32 `json:"fail_count"`
	FailDetails []interface{} `json:"fail_details"`
	SuccessCount int32 `json:"success_count"`
}

type V1OrderValidationError struct {
	OrderErrors []interface{} `json:"order_errors"`
}

type V2CancellationStateEnum string

type GetConditionalCancellationListV2ResponseState struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	State V2CancellationStateEnum `json:"state"`
}

type V1GetProductStairwayDiscountByQuantityResponse struct {
	Stairways []interface{} `json:"stairways"`
}

type V1FbpCreateConsignmentNoteResponse struct {
	Code string `json:"code"`
}

type DetailsRfbsDetails struct {
	CompensationDeliveryReturn float64 `json:"compensation_delivery_return"`
	PartialCompensation float64 `json:"partial_compensation"`
	PartialCompensationReturn float64 `json:"partial_compensation_return"`
	Total float64 `json:"total"`
	TransferDelivery float64 `json:"transfer_delivery"`
	TransferDeliveryReturn float64 `json:"transfer_delivery_return"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProduct struct {
	Error string `json:"error"`
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
	Valid bool `json:"valid"`
}

type CreateReportResponseCodeNoDeadline struct {
	Code string `json:"code"`
}

type CreateReportResponse struct {
	Result CreateReportResponseCodeNoDeadline `json:"result"`
}

type V1FbpDraftDirectTplDlvEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type V3PostingFinancialData struct {
	Products []interface{} `json:"products"`
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
}

type V3GetProductInfoListRequest struct {
	ProductID []interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
	OfferID []interface{} `json:"offer_id"`
}

type RatingStatusEnum string

type SellerInfoResponseRatingTypeEnum string

type V1RatingStatus struct {
	Premium bool `json:"premium"`
	Warning bool `json:"warning"`
	Danger bool `json:"danger"`
}

type RatingValueCurrent struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Formatted string `json:"formatted"`
	Status V1RatingStatus `json:"status"`
	Value float64 `json:"value"`
}

type RatingValuePast struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Formatted string `json:"formatted"`
	Status V1RatingStatus `json:"status"`
	Value float64 `json:"value"`
}

type SellerInfoResponseRating struct {
	Name string `json:"name"`
	PastValue RatingValuePast `json:"past_value"`
	Rating string `json:"rating"`
	Status RatingStatusEnum `json:"status"`
	ValueType SellerInfoResponseRatingTypeEnum `json:"value_type"`
	CurrentValue RatingValueCurrent `json:"current_value"`
}

type Productv2ProductsStocksRequestStock struct {
	WarehouseID int64 `json:"warehouse_id"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Stock int64 `json:"stock"`
}

type V3FbsPostingDetailRelatedPostings struct {
	RelatedPostingNumbers interface{} `json:"related_posting_numbers"`
}

type GetProductInfoListResponseStocks struct {
	Stocks []interface{} `json:"stocks"`
	HasStock bool `json:"has_stock"`
}

type Productv2GetProductListRequestFilterFilterVisibility string

type Productv3Filter struct {
	ProductID interface{} `json:"product_id"`
	Visibility Productv2GetProductListRequestFilterFilterVisibility `json:"visibility"`
	OfferID interface{} `json:"offer_id"`
}

type Productv3GetProductAttributesV3Request struct {
	Filter Productv3Filter `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

type V1FbpCreateLabelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V5FbsPostingProductExemplarValidateV5Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type FilterPeriod struct {
	From string `json:"from"`
	To string `json:"to"`
}

type FinanceTransactionListV3RequestFilter struct {
	Date FilterPeriod `json:"date"`
	OperationType []interface{} `json:"operation_type"`
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
}

type ProductV1ProductVisibilitySetRequestItemPlacementPlacementEnum string

type V3FinanceCashFlowStatementListResponseResult struct {
	CashFlows interface{} `json:"cash_flows"`
	Details []interface{} `json:"details"`
	PageCount int64 `json:"page_count"`
}

type ArrivalpassArrivalPassCreateRequestArrivalPass struct {
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	DropoffPointID int64 `json:"dropoff_point_id"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ActionsV1ActionsAutoAddProductsListRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type V3FbsPostingDetailOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type FbsPostingShipV4RequestWith struct {
	AdditionalData bool `json:"additional_data"`
}

type Fbsv4FbsPostingShipV4Request struct {
	Packages interface{} `json:"packages"`
	PostingNumber string `json:"posting_number"`
	With FbsPostingShipV4RequestWith `json:"with"`
}

type ListFBSRatingIndexPostingsV1RequestFilter struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Cursor string `json:"cursor"`
	Filter ListFBSRatingIndexPostingsV1RequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
}

type V1FbpDraftDirectRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1GetStrategyIDsByItemIDsResponseResult struct {
	ProductsInfo []interface{} `json:"products_info"`
}

type V1GetStrategyIDsByItemIDsResponse struct {
	Result V1GetStrategyIDsByItemIDsResponseResult `json:"result"`
}

type ChatMessageContext struct {
	OrderNumber string `json:"order_number"`
	SKU string `json:"sku"`
}

type V2GetDiscountTaskListV2RequestDiscountTaskStatusEnum string

type V2GetDiscountTaskListV2Request struct {
	LastID int64 `json:"last_id"`
	Limit int64 `json:"limit"`
	Status V2GetDiscountTaskListV2RequestDiscountTaskStatusEnum `json:"status"`
}

type Financev3FinanceTransactionListV3ResponseResult struct {
	Operations []interface{} `json:"operations"`
	PageCount int64 `json:"page_count"`
	RowCount int64 `json:"row_count"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result Financev3FinanceTransactionListV3ResponseResult `json:"result"`
}

type ProductGetProductInfoDescriptionRequest struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
}

type V1StrategyRequest struct {
	StrategyID string `json:"strategy_id"`
}

type PolygonBindRequestpolygon struct {
	PolygonID int64 `json:"polygon_id"`
	Time int64 `json:"time"`
}

type ReviewV2ReviewListV2ResponseReviewStatusEnum string

type V1PostingFbsSplitResponsePostingParent struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1PostingFbsSplitResponse struct {
	ParentPosting V1PostingFbsSplitResponsePostingParent `json:"parent_posting"`
	Postings []interface{} `json:"postings"`
}

type AddStrategyItemsResponseError struct {
	Code string `json:"code"`
	Error string `json:"error"`
	ProductID int64 `json:"product_id"`
}

type V3FbsTariffication struct {
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	NextTariffType string `json:"next_tariff_type"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	CurrentTariffCharge string `json:"current_tariff_charge"`
	CurrentTariffChargeCurrencyCode string `json:"current_tariff_charge_currency_code"`
	NextTariffCharge string `json:"next_tariff_charge"`
	NextTariffChargeCurrencyCode string `json:"next_tariff_charge_currency_code"`
}

type AnalyticsProductQueriesDetailsResponseQuery struct {
	Position float64 `json:"position"`
	Query string `json:"query"`
	QueryIndex int64 `json:"query_index"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
	UniqueViewUsers int64 `json:"unique_view_users"`
	Currency string `json:"currency"`
	Gmv float64 `json:"gmv"`
	OrderCount int64 `json:"order_count"`
	SKU int64 `json:"sku"`
	ViewConversion float64 `json:"view_conversion"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFees struct {
	Fees []interface{} `json:"fees"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejectedCodeEnum string

type Productv3GetProductAttributesV3ResponseResult struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	Depth int32 `json:"depth"`
	Images360 []interface{} `json:"images360"`
	PDFList []interface{} `json:"pdf_list"`
	Weight int32 `json:"weight"`
	Attributes []interface{} `json:"attributes"`
	ImageGroupID string `json:"image_group_id"`
	Images []interface{} `json:"images"`
	WeightUnit string `json:"weight_unit"`
	TypeID int64 `json:"type_id"`
	Width int32 `json:"width"`
	Barcode string `json:"barcode"`
	ColorImage string `json:"color_image"`
	DimensionUnit string `json:"dimension_unit"`
	ID int64 `json:"id"`
	OfferID string `json:"offer_id"`
	CategoryID int64 `json:"category_id"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Height int32 `json:"height"`
	Name string `json:"name"`
}

type DirectDetailsBySellerDetails struct {
	DriverName string `json:"driver_name"`
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
	VehicleType string `json:"vehicle_type"`
}

type V1fbpTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type DirectDetailsTimeslotDetails struct {
	Timeslot V1fbpTimeslot `json:"timeslot"`
	TimeslotReservationID string `json:"timeslot_reservation_id"`
}

type V1DeliveryDetailsDirectDetails struct {
	BySellerDetails DirectDetailsBySellerDetails `json:"by_seller_details"`
	ByTPLDetails DirectDetailsByTplDetails `json:"by_tpl_details"`
	TimeslotDetails DirectDetailsTimeslotDetails `json:"timeslot_details"`
}

type DeliveryDetailsDropOffPointDetails struct {
	ID int64 `json:"id"`
	ProvinceUuid string `json:"province_uuid"`
	Timeslot V1fbpTimeslot `json:"timeslot"`
}

type V1DeliveryDetailsPickUpDetails struct {
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type DeliveryDetailsSupplyType string

type Fbpv1DeliveryDetails struct {
	SupplyType DeliveryDetailsSupplyType `json:"supply_type"`
	DirectDetails V1DeliveryDetailsDirectDetails `json:"direct_details"`
	DropOffPoint DeliveryDetailsDropOffPointDetails `json:"drop_off_point"`
	PickupDetails V1DeliveryDetailsPickUpDetails `json:"pickup_details"`
}

type V1OrderStatusEnum string

type CancellationErrorCode string

type CancellationStateCancellationError struct {
	ErrorCode CancellationErrorCode `json:"error_code"`
	Message string `json:"message"`
}

type V1CancellationStateStatus string

type V1CancellationState struct {
	CancellationError CancellationStateCancellationError `json:"cancellation_error"`
	CancellationStatus V1CancellationStateStatus `json:"cancellation_status"`
}

type V1FbpOrderGetResponse struct {
	HasLabel bool `json:"has_label"`
	ID int64 `json:"id"`
	Locked bool `json:"locked"`
	ReceiveDate string `json:"receive_date"`
	RowVersion int64 `json:"row_version"`
	WarehouseID int64 `json:"warehouse_id"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	BundleUuid string `json:"bundle_uuid"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	Status V1OrderStatusEnum `json:"status"`
	CreatedDate string `json:"created_date"`
	DraftID int64 `json:"draft_id"`
	SupplyID string `json:"supply_id"`
	CancellationState V1CancellationState `json:"cancellation_state"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	OrderNumber string `json:"order_number"`
}

type ProductGetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type SellerApiGetSellerProductV1ResponseResult struct {
	Products []interface{} `json:"products"`
	Total float64 `json:"total"`
	LastID float64 `json:"last_id"`
}

type SellerApiGetSellerProductV1Response struct {
	Result SellerApiGetSellerProductV1ResponseResult `json:"result"`
}

type V2PostingFBSActGetPostingsResult struct {
	CreatedAt string `json:"created_at"`
	Products []interface{} `json:"products"`
	ID int64 `json:"id"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Status string `json:"status"`
	SellerError string `json:"seller_error"`
	UpdatedAt string `json:"updated_at"`
}

type V1RolesByTokenResponse struct {
	ExpiresAt string `json:"expires_at"`
	Roles []interface{} `json:"roles"`
}

type ItemMarketing struct {
	Actions []interface{} `json:"actions"`
}

type V2WarehouseListV2Request struct {
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"`
}

type FinanceV1GetFinanceAccrualByDayResponse struct {
	Accruals []interface{} `json:"accruals"`
	LastID string `json:"last_id"`
}

type V1DeclineDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V1DiscountTaskStatus string

type MoneyMoneyAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualNonItemFee struct {
	Accrued MoneyMoneyAccrued `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type V1UnarchiveWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ItemSfboAttribute string

type V1FbpDraftDirectSellerDlvCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
	DriverName string `json:"driver_name"`
}

type V1AssemblyFbsPostingListResponsePostingProduct struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type FilterStatusEnum string

type FinanceV1GetFinanceAccrualByDayResponseAccrualAccruedCategoryEnum string

type V1GetFinanceBalanceV1ResponseReturnsMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type GetFinanceBalanceV1ResponseReturnsDetails struct {
	PartnerPrograms V1GetFinanceBalanceV1ResponsePartnerMoney `json:"partner_programs"`
	PointsForDiscounts string `json:"points_for_discounts"`
	Revenue V1GetFinanceBalanceV1ResponseRevenueMoney `json:"revenue"`
}

type V1GetFinanceBalanceV1ResponseFeeMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type GetFinanceBalanceV1ResponseReturns struct {
	AmountDetails GetFinanceBalanceV1ResponseReturnsDetails `json:"amount_details"`
	Fee V1GetFinanceBalanceV1ResponseFeeMoney `json:"fee"`
	Amount V1GetFinanceBalanceV1ResponseReturnsMoney `json:"amount"`
}

type Postingv3FbsPostingWithParamsExamplars struct {
	Translit bool `json:"translit"`
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
	ProductExemplars bool `json:"product_exemplars"`
	RelatedPostings bool `json:"related_postings"`
}

type Postingv3GetFbsPostingRequest struct {
	PostingNumber string `json:"posting_number"`
	With Postingv3FbsPostingWithParamsExamplars `json:"with"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequestAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1GetProductInfoDiscountedResponseItem struct {
	MechanicalDamage string `json:"mechanical_damage"`
	PackageDamage string `json:"package_damage"`
	ReasonDamaged string `json:"reason_damaged"`
	Repair string `json:"repair"`
	Shortage string `json:"shortage"`
	CommentReasonDamaged string `json:"comment_reason_damaged"`
	Condition string `json:"condition"`
	ConditionEstimation string `json:"condition_estimation"`
	PackagingViolation string `json:"packaging_violation"`
	SKU int64 `json:"sku"`
	WarrantyType string `json:"warranty_type"`
	Defects string `json:"defects"`
	DiscountedSKU int64 `json:"discounted_sku"`
}

type SellerActionsUpdateMultiLevelDiscountRequestActionParametersDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type FilterWithQuant struct {
	Created bool `json:"created"`
	Exists bool `json:"exists"`
}

type V6FbsPostingProductExemplarCreateOrGetV6Request struct {
	PostingNumber string `json:"posting_number"`
}

type GetProductAttributesResponseImage360 struct {
	Index int64 `json:"index"`
	FileName string `json:"file_name"`
}

type V1WarehouseListRequestWith struct {
	AbleToSetPrice bool `json:"able_to_set_price"`
}

type V1WarehouseListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With V1WarehouseListRequestWith `json:"with"`
}

type V1FbpDraftDropOffProvinceListResponse struct {
	Provinces []interface{} `json:"provinces"`
}

type GetReturnsListResponsePlaceNow struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type GetReturnsListResponsePlaceTarget struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
}

type SellerReturnsv1MoneyStorage struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type SellerReturnsv1MoneyUtilization struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type GetReturnsListResponseStorage struct {
	Sum SellerReturnsv1MoneyStorage `json:"sum"`
	TarifficationFirstDate string `json:"tariffication_first_date"`
	TarifficationStartDate string `json:"tariffication_start_date"`
	ArrivedMoment string `json:"arrived_moment"`
	Days int64 `json:"days"`
	UtilizationSum SellerReturnsv1MoneyUtilization `json:"utilization_sum"`
	UtilizationForecastDate string `json:"utilization_forecast_date"`
}

type SellerReturnsv1MoneyProduct struct {
	Price float64 `json:"price"`
	CurrencyCode string `json:"currency_code"`
}

type SellerReturnsv1MoneyWithoutCommission struct {
	Price float64 `json:"price"`
	CurrencyCode string `json:"currency_code"`
}

type SellerReturnsv1MoneyCommission struct {
	CurrencyCode string `json:"currency_code"`
	Price float64 `json:"price"`
}

type GetReturnsListResponseProduct struct {
	Commission SellerReturnsv1MoneyCommission `json:"commission"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	Name string `json:"name"`
	Price SellerReturnsv1MoneyProduct `json:"price"`
	PriceWithoutCommission SellerReturnsv1MoneyWithoutCommission `json:"price_without_commission"`
	CommissionPercent float64 `json:"commission_percent"`
}

type GetReturnsListResponseVisualStatus struct {
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
	SysName string `json:"sys_name"`
}

type GetReturnsListResponseVisual struct {
	Status GetReturnsListResponseVisualStatus `json:"status"`
	ChangeMoment string `json:"change_moment"`
}

type GetReturnsListResponseCompensationStatus struct {
	SysName string `json:"sys_name"`
	ID int32 `json:"id"`
	DisplayName string `json:"display_name"`
}

type GetReturnsListResponseCompensation struct {
	Status GetReturnsListResponseCompensationStatus `json:"status"`
	ChangeMoment string `json:"change_moment"`
}

type GetReturnsListResponseAdditionalInfo struct {
	IsOpened bool `json:"is_opened"`
	IsSuperEconom bool `json:"is_super_econom"`
}

type GetReturnsListResponseReturnsItem struct {
	Exemplars []interface{} `json:"exemplars"`
	Schema string `json:"schema"`
	TargetPlace GetReturnsListResponsePlaceTarget `json:"target_place"`
	Storage GetReturnsListResponseStorage `json:"storage"`
	Product GetReturnsListResponseProduct `json:"product"`
	Visual GetReturnsListResponseVisual `json:"visual"`
	SourceID int64 `json:"source_id"`
	CompensationStatus GetReturnsListResponseCompensation `json:"compensation_status"`
	CompanyID int64 `json:"company_id"`
	OrderID int64 `json:"order_id"`
	ReturnClearingID int64 `json:"return_clearing_id"`
	ID int64 `json:"id"`
	ReturnReasonName string `json:"return_reason_name"`
	AdditionalInfo GetReturnsListResponseAdditionalInfo `json:"additional_info"`
	PostingNumber string `json:"posting_number"`
	ClearingID int64 `json:"clearing_id"`
	Type string `json:"type_"`
	OrderNumber string `json:"order_number"`
	Place GetReturnsListResponsePlaceNow `json:"place"`
	Logistic GetReturnsListResponseLogistic `json:"logistic"`
}

type V1QuestionListRequestFilter struct {
	DateTo string `json:"date_to"`
	Status string `json:"status"`
	DateFrom string `json:"date_from"`
}

type QuestionV1GetQuestionListRequestSortDirEnum string

type V1QuestionListRequest struct {
	Filter V1QuestionListRequestFilter `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortDir QuestionV1GetQuestionListRequestSortDirEnum `json:"sort_dir"`
}

type ParameterPickedSegment struct {
	Segments []interface{} `json:"segments"`
}

type V1GetAttributeValuesResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V1DeclineDiscountTasksRequestTask struct {
	ID int64 `json:"id"`
	SellerComment string `json:"seller_comment"`
}

type V1FbpDraftDirectGetTimeslotResponseTimeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type FinanceV1GetFinanceAccrualPostingsResponse struct {
	PostingAccruals []interface{} `json:"posting_accruals"`
}

type V2Product struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	CurrencyCode string `json:"currency_code"`
	Price int32 `json:"price"`
	SKU int64 `json:"sku"`
}

type Productv3GetProductListRequestFilterFilterVisibility string

type V1ItemSortField string

type V1BundleItemShipmentType string

type V2FboSinglePostingLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type ArrivalpassArrivalPassListResponse struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	Cursor string `json:"cursor"`
}

type WarehouseAddressInfo struct {
	Address string `json:"address"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Utc string `json:"utc"`
}

type V1CommentSort string

type SellerApiGetSellerActionsV1Response struct {
	Result []interface{} `json:"result"`
}

type V1CarriageApproveRequest struct {
	CarriageID int64 `json:"carriage_id"`
	ContainersCount int32 `json:"containers_count"`
}

type SellerServiceanalyticsDimension string

type DropOffPointCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V2ReturnsRfbsListV2ResponseState struct {
	GroupState string `json:"group_state"`
	MoneyReturnStateName string `json:"money_return_state_name"`
	State string `json:"state"`
	StateName string `json:"state_name"`
}

type ReturnsRfbsListResponseReturns struct {
	State V2ReturnsRfbsListV2ResponseState `json:"state"`
	ClientName string `json:"client_name"`
	CreatedAt string `json:"created_at"`
	OrderNumber string `json:"order_number"`
	PostingNumber string `json:"posting_number"`
	Product V2Product `json:"product"`
	ReturnID int64 `json:"return_id"`
	ReturnNumber string `json:"return_number"`
}

type V1SellerActionsCreateVoucherResponse struct {
	ActionID int64 `json:"action_id"`
}

type GetSupplyOrderBundleRequestItemTagsCalculation struct {
	DropoffWarehouseID string `json:"dropoff_warehouse_id"`
	StorageWarehouseIds []interface{} `json:"storage_warehouse_ids"`
}

type V1GetSupplyOrderBundleRequest struct {
	BundleIds []interface{} `json:"bundle_ids"`
	IsAsc bool `json:"is_asc"`
	ItemTagsCalculation GetSupplyOrderBundleRequestItemTagsCalculation `json:"item_tags_calculation"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	Query string `json:"query"`
	SortField V1ItemSortField `json:"sort_field"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccruals struct {
	Accruals []interface{} `json:"accruals"`
	PostingNumber string `json:"posting_number"`
}

type ProductExemplar struct {
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	ExemplarID int64 `json:"exemplar_id"`
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearch struct {
	Address string `json:"address"`
	Types []interface{} `json:"types"`
}

type ReviewV2ReviewListV2RequestFiltersOrderStatusEnum string

type ReviewV2ReviewListV2RequestFiltersStatusEnum string

type ReviewV2ReviewListV2RequestFilters struct {
	OrderStatus ReviewV2ReviewListV2RequestFiltersOrderStatusEnum `json:"order_status"`
	PublishedFrom string `json:"published_from"`
	PublishedTo string `json:"published_to"`
	Skus []interface{} `json:"skus"`
	Status ReviewV2ReviewListV2RequestFiltersStatusEnum `json:"status"`
}

type ReviewV2ReviewListV2RequestSortDirEnum string

type ReviewV2ReviewListV2Request struct {
	Filters ReviewV2ReviewListV2RequestFilters `json:"filters"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortDir ReviewV2ReviewListV2RequestSortDirEnum `json:"sort_dir"`
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

type PostingV1PostingFbpListRequestSortDirEnum string

type V1FbpDraftPickupDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1FbpOrderDropOffTimetableResponseCalendarCalendarItem struct {
	BreakHours V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTimeBreak `json:"break_hours"`
	IsHoliday bool `json:"is_holiday"`
	OpeningHours V1FbpOrderDropOffTimetableResponseCalendarCalendarItemTimeslotWithTime `json:"opening_hours"`
}

type V3GetProductInfoListResponse struct {
	Items []interface{} `json:"items"`
}

type ProductV4GetUploadQuotaResponseOperationLimitsLimitTypeEnum string

type ProductV4GetUploadQuotaResponseOperationLimits struct {
	Limit int64 `json:"limit"`
	LimitType ProductV4GetUploadQuotaResponseOperationLimitsLimitTypeEnum `json:"limit_type"`
}

type MoneyPostingMoney struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type V2WarehouseListV2Response struct {
	Cursor string `json:"cursor"`
	Warehouses []interface{} `json:"warehouses"`
	HasNext bool `json:"has_next"`
}

type ReturnsRfbsGetV2ResponseReturnReason struct {
	Name string `json:"name"`
	ID int32 `json:"id"`
	IsDefect bool `json:"is_defect"`
}

type CarriageCarriageGetResponseCancelAvailability struct {
	IsCancelAvailable bool `json:"is_cancel_available"`
	Reason string `json:"reason"`
}

type CarriageCarriageGetResponse struct {
	RetryCount int32 `json:"retry_count"`
	ActType string `json:"act_type"`
	CancelAvailability CarriageCarriageGetResponseCancelAvailability `json:"cancel_availability"`
	CompanyID int64 `json:"company_id"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
	FirstMileType string `json:"first_mile_type"`
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"`
	IsContainerLabelPrinted bool `json:"is_container_label_printed"`
	AvailableActions []interface{} `json:"available_actions"`
	ContainersCount int32 `json:"containers_count"`
	Status string `json:"status"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CarriageID int64 `json:"carriage_id"`
	CreatedAt string `json:"created_at"`
	DepartureDate string `json:"departure_date"`
	IntegrationType string `json:"integration_type"`
	IsPartial bool `json:"is_partial"`
	PartialNum int64 `json:"partial_num"`
}

type ReportLanguage string

type V2ProductInfoPicturesResponseItem struct {
	Photo []interface{} `json:"photo"`
	ColorPhoto []interface{} `json:"color_photo"`
	Photo360 []interface{} `json:"photo_360"`
	Errors []interface{} `json:"errors"`
	ProductID int64 `json:"product_id"`
	PrimaryPhoto []interface{} `json:"primary_photo"`
}

type GetDiscountTaskListV2ResponseTaskDiscountTaskStatusEnum string

type V1GetFinanceBalanceV1Request struct {
	DateTo string `json:"date_to"`
	DateFrom string `json:"date_from"`
}

type V1GetProductInfoDiscountedResponse struct {
	Items interface{} `json:"items"`
}

type V1SetPostingsResponse struct {
	Result interface{} `json:"result"`
}

type V1GetWarehouseFBSOperationStatusRequest struct {
	OperationID string `json:"operation_id"`
}

type ChatInfoChatStatusEnum string

type V2GetDiscountTaskListV2Response struct {
	Tasks []interface{} `json:"tasks"`
}

type V1OrderErrorTypeEnum string

type V1FbpDraftPickUpRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError V1OrderErrorTypeEnum `json:"order_error"`
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

type AnalyticsDataRow struct {
	Dimensions []interface{} `json:"dimensions"`
	Metrics []interface{} `json:"metrics"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type MoneyMoneyCurrentTariffCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1FbpOrderPickUpDlvEditRequestPickUpDetails struct {
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

type V1FbpOrderPickUpDlvEditRequest struct {
	PickupDetails V1FbpOrderPickUpDlvEditRequestPickUpDetails `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1UpdateWarehouseFBSRequestAddressCoordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}

type V1UpdateWarehouseFBSRequestOptions struct {
	Comment string `json:"comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
}

type V1UpdateWarehouseFBSRequest struct {
	AddressCoordinates V1UpdateWarehouseFBSRequestAddressCoordinates `json:"address_coordinates"`
	Name string `json:"name"`
	Options V1UpdateWarehouseFBSRequestOptions `json:"options"`
	Phone string `json:"phone"`
	WarehouseID int64 `json:"warehouse_id"`
	WorkingDays []interface{} `json:"working_days"`
}

type V1AnalyticsProductQueriesDetailsResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1AnalyticsProductQueriesDetailsResponse struct {
	Queries []interface{} `json:"queries"`
	Total int64 `json:"total"`
	AnalyticsPeriod V1AnalyticsProductQueriesDetailsResponseAnalyticsPeriod `json:"analytics_period"`
	PageCount int64 `json:"page_count"`
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

type GetUploadQuotaResponseTotal struct {
	QuotaByCategory ProductV4GetUploadQuotaResponseTotalQuotaByCategory `json:"quota_by_category"`
	Limit int64 `json:"limit"`
	Usage int64 `json:"usage"`
}

type Financev3FinanceTransactionTotalsV3ResponseResult struct {
	ServicesAmount float64 `json:"services_amount"`
	AccrualsForSale float64 `json:"accruals_for_sale"`
	CompensationAmount float64 `json:"compensation_amount"`
	MoneyTransfer float64 `json:"money_transfer"`
	OthersAmount float64 `json:"others_amount"`
	ProcessingAndDelivery float64 `json:"processing_and_delivery"`
	RefundsAndCancellations float64 `json:"refunds_and_cancellations"`
	SaleCommission float64 `json:"sale_commission"`
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result Financev3FinanceTransactionTotalsV3ResponseResult `json:"result"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountLevel struct {
	DiscountValue float64 `json:"discount_value"`
	OrderAmount float64 `json:"order_amount"`
}

type V1GetProductInfoSubscriptionResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftDropOffDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type CreateReportResponseCode struct {
	Code string `json:"code"`
}

type ReasonHumanText struct {
	Text string `json:"text"`
}

type AvailabilityReason struct {
	HumanText ReasonHumanText `json:"human_text"`
	ID int64 `json:"id"`
}

type ChatMessageModerateImageStatus string

type FbpCheckConsignmentNoteStateResponseStateType string

type V1DeleteStrategyItemsResponseResult struct {
	FailedProductCount int32 `json:"failed_product_count"`
}

type ChatChatSendMessageRequest struct {
	Text string `json:"text"`
	ChatID string `json:"chat_id"`
}

type V3FbsPostingAnalyticsData struct {
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	City string `json:"city"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	IsPremium bool `json:"is_premium"`
	Region string `json:"region"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
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

type TypeTimeOfDay struct {
	Minutes int32 `json:"minutes"`
	Nanos int32 `json:"nanos"`
	Seconds int32 `json:"seconds"`
	Hours int32 `json:"hours"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponseDropOffPointTypeEnum string

type V1ListDropOffPointsForCreateFBSWarehouseResponseCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ListDropOffPointsForCreateFBSWarehouseResponseDropOffPoint struct {
	Address string `json:"address"`
	Coordinates V1ListDropOffPointsForCreateFBSWarehouseResponseCoordinates `json:"coordinates"`
	DiscountPercent float64 `json:"discount_percent"`
	ID string `json:"id"`
	LastTransitTimeLocal TypeTimeOfDay `json:"last_transit_time_local"`
	Type V1ListDropOffPointsForCreateFBSWarehouseResponseDropOffPointTypeEnum `json:"type_"`
}

type SourceShipmentType string

type GetProductInfoListResponseSource struct {
	CreatedAt string `json:"created_at"`
	QuantCode string `json:"quant_code"`
	ShipmentType SourceShipmentType `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type V5FbsPostingProductExemplarValidateV5RequestProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type V1SellerActionsUpdateDiscountWithConditionRequestActionParameters struct {
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
}

type FinanceCashFlowStatementListResponseReturnService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V1FbpDraftDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V1AssemblyCarriagePostingListResponsePosting struct {
	AssemblyCode string `json:"assembly_code"`
	CanPrintLabel bool `json:"can_print_label"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type RpcStatusV1PolygonCreate struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1UpdateStatusStrategyRequest struct {
	Enabled bool `json:"enabled"`
	StrategyID string `json:"strategy_id"`
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"`
}

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID int64 `json:"return_id"`
	ReturnMethodDescription string `json:"return_method_description"`
}

type RowItemOrder struct {
	PostingNumber string `json:"posting_number"`
	CreatedDate string `json:"created_date"`
}

type ChatStartResponseResult struct {
	ChatID string `json:"chat_id"`
}

type ChatChatStartResponse struct {
	Result ChatStartResponseResult `json:"result"`
}

type V1ArchiveWarehouseFBSRequest struct {
	Reason string `json:"reason"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1GetAttributeValuesResponseDictionaryValue struct {
	ID int64 `json:"id"`
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
}

type PostingV4PostingFbsListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type PostingV4PostingFbsListRequestFilter struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
	OrderID int64 `json:"order_id"`
	ProviderIds []interface{} `json:"provider_ids"`
	Since string `json:"since"`
	To string `json:"to"`
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	LastChangedStatusDate PostingV4PostingFbsListRequestFilterLastChangedStatusDate `json:"last_changed_status_date"`
	OrderNumbers []interface{} `json:"order_numbers"`
	Statuses []interface{} `json:"statuses"`
}

type PostingV4PostingFbsListRequestSortDirEnum string

type PostingV4PostingFbsListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	LegalInfo bool `json:"legal_info"`
}

type PostingV4PostingFbsListRequest struct {
	Filter PostingV4PostingFbsListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir PostingV4PostingFbsListRequestSortDirEnum `json:"sort_dir"`
	Translit bool `json:"translit"`
	With PostingV4PostingFbsListRequestWith `json:"with"`
	Cursor string `json:"cursor"`
}

type SellerApiProduct struct {
	MinBoost float64 `json:"min_boost"`
	MaxBoost float64 `json:"max_boost"`
	ID float64 `json:"id"`
	Price float64 `json:"price"`
	ActionPrice float64 `json:"action_price"`
	AlertMaxActionPrice float64 `json:"alert_max_action_price"`
	MaxActionPrice float64 `json:"max_action_price"`
	AddMode string `json:"add_mode"`
	Stock float64 `json:"stock"`
	CurrentBoost float64 `json:"current_boost"`
	AlertMaxActionPriceFailed bool `json:"alert_max_action_price_failed"`
	MinStock float64 `json:"min_stock"`
	PriceMinElastic float64 `json:"price_min_elastic"`
	PriceMaxElastic float64 `json:"price_max_elastic"`
}

type DeliveryMethodListV2RequestFilter struct {
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	ProviderIds []interface{} `json:"provider_ids"`
	Status []interface{} `json:"status"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequestCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type PostingV4PostingFbsUnfulfilledListRequestFilter struct {
	DeliveringDateTo string `json:"delivering_date_to"`
	DeliveryMethodIds []interface{} `json:"delivery_method_ids"`
	ProviderIds []interface{} `json:"provider_ids"`
	Statuses []interface{} `json:"statuses"`
	LastChangedStatusDate PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate `json:"last_changed_status_date"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveringDateFrom string `json:"delivering_date_from"`
}

type Productv3GetProductListRequestFilter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility Productv3GetProductListRequestFilterFilterVisibility `json:"visibility"`
}

type V1GetFinanceBalanceV1ResponsePaymentsMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1GetFinanceBalanceV1ResponseSalesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type GetFinanceBalanceV1ResponseSales struct {
	Amount V1GetFinanceBalanceV1ResponseSalesMoney `json:"amount"`
	AmountDetails GetFinanceBalanceV1ResponseSalesDetails `json:"amount_details"`
	Fee V1GetFinanceBalanceV1ResponseFeeMoney `json:"fee"`
}

type GetFinanceBalanceV1ResponseCashflows struct {
	Returns GetFinanceBalanceV1ResponseReturns `json:"returns"`
	Sales GetFinanceBalanceV1ResponseSales `json:"sales"`
	Services []interface{} `json:"services"`
}

type V1GetFinanceBalanceV1ResponseAccruedMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1GetFinanceBalanceV1ResponseClosingBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type V1GetFinanceBalanceV1ResponseOpeningBalanceMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type GetFinanceBalanceV1ResponseTotal struct {
	Accrued V1GetFinanceBalanceV1ResponseAccruedMoney `json:"accrued"`
	ClosingBalance V1GetFinanceBalanceV1ResponseClosingBalanceMoney `json:"closing_balance"`
	OpeningBalance V1GetFinanceBalanceV1ResponseOpeningBalanceMoney `json:"opening_balance"`
	Payments []interface{} `json:"payments"`
}

type V1GetFinanceBalanceV1Response struct {
	Cashflows GetFinanceBalanceV1ResponseCashflows `json:"cashflows"`
	Total GetFinanceBalanceV1ResponseTotal `json:"total"`
}

type FinanceV1GetFinanceAccrualPostingsResponsePostingAccrualsAccrual struct {
	AccrualDate string `json:"accrual_date"`
	Accrued MoneyMoneyAccrued `json:"accrued"`
	Quantity int32 `json:"quantity"`
	SellerPrice MoneyMoneySellerPrice `json:"seller_price"`
	SKU int64 `json:"sku"`
	TypeID int32 `json:"type_id"`
}

type V1SellerActionsCreateDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type SellerActionsListRequestStatusEnum string

type SellerActionsListRequestActionTypeEnum string

type ParameterAutoStopActionReasonEnum string

type ActionParameterDiscountTypeEnum string

type ActionParameter struct {
	Status SellerActionsListRequestStatusEnum `json:"status"`
	Budget float64 `json:"budget"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Type SellerActionsListRequestActionTypeEnum `json:"type_"`
	VoucherParameters ActionParameterVoucherParameter `json:"voucher_parameters"`
	BudgetSpent float64 `json:"budget_spent"`
	DiscountValue float64 `json:"discount_value"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	PickedSegments []interface{} `json:"picked_segments"`
	Addresses []interface{} `json:"addresses"`
	AutoStopActionReason ParameterAutoStopActionReasonEnum `json:"auto_stop_action_reason"`
	DiscountType ActionParameterDiscountTypeEnum `json:"discount_type"`
	MinActionPercent float64 `json:"min_action_percent"`
	DateEnd string `json:"date_end"`
	Title string `json:"title"`
	Warehouses []interface{} `json:"warehouses"`
}

type SellerActionsListResponseAction struct {
	AllowDelete bool `json:"allow_delete"`
	HighlightURL string `json:"highlight_url"`
	IsEditable bool `json:"is_editable"`
	IsParticipated bool `json:"is_participated"`
	IsTurnOn bool `json:"is_turn_on"`
	SKUCount int64 `json:"sku_count"`
	ActionID int64 `json:"action_id"`
	ActionParameters ActionParameter `json:"action_parameters"`
}

type GetProductRatingBySkuResponseRatingCondition struct {
	Description string `json:"description"`
	Fulfilled bool `json:"fulfilled"`
	Key string `json:"key"`
	Cost float64 `json:"cost"`
}

type ChatChatSendFileRequest struct {
	Base64Content string `json:"base64_content"`
	ChatID string `json:"chat_id"`
	Name string `json:"name"`
}

type PriceIndexesIndexDataSelf struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type AnalyticsProductQueriesResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

type V1AnalyticsProductQueriesResponse struct {
	Total int64 `json:"total"`
	AnalyticsPeriod AnalyticsProductQueriesResponseAnalyticsPeriod `json:"analytics_period"`
	Items []interface{} `json:"items"`
	PageCount int64 `json:"page_count"`
}

type V1AssemblyCarriageProductListRequestFilter struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
}

type V1AssemblyCarriageProductListRequest struct {
	Cursor string `json:"cursor"`
	Filter V1AssemblyCarriageProductListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
}

type GetRealizationReportResponseV2Header struct {
	ReceiverName string `json:"receiver_name"`
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
}

type GetRealizationReportResponseV2Result struct {
	Rows []interface{} `json:"rows"`
	Header GetRealizationReportResponseV2Header `json:"header"`
}

type V2GetRealizationReportResponseV2 struct {
	Result GetRealizationReportResponseV2Result `json:"result"`
}

type V1ProductActionTimerStatusResponse struct {
	Statuses interface{} `json:"statuses"`
}

type Productv5GetProductListRequestFilterFilterVisibility string

type Productv5Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	Visibility Productv5GetProductListRequestFilterFilterVisibility `json:"visibility"`
}

type DeleteProductsResponseDeleteStatus struct {
	Error string `json:"error"`
	IsDeleted bool `json:"is_deleted"`
	OfferID string `json:"offer_id"`
}

type V1SearchAttributeValuesRequest struct {
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Limit int64 `json:"limit"`
	TypeID int64 `json:"type_id"`
	Value string `json:"value"`
}

type V1FbpDraftDirectProductValidateResponseApprovedItem struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
}

type FirstMileTypeEnum string

type V1fbpDeliveryDetails struct {
	DirectDetails V1DeliveryDetailsDirectDetails `json:"direct_details"`
	DropOffPoint DeliveryDetailsDropOffPointDetails `json:"drop_off_point"`
	PickupDetails V1DeliveryDetailsPickUpDetails `json:"pickup_details"`
	SupplyType DeliveryDetailsSupplyType `json:"supply_type"`
}

type V1DraftStatusEnum string

type FbpDraftGetResponseDeclineReason struct {
	FailedSKUIds []interface{} `json:"failed_sku_ids"`
	Message string `json:"message"`
}

type V1FbpDraftGetResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	DeletedAt string `json:"deleted_at"`
	DeliveryDetails V1fbpDeliveryDetails `json:"delivery_details"`
	RowVersion int64 `json:"row_version"`
	Status V1DraftStatusEnum `json:"status"`
	SupplyID string `json:"supply_id"`
	Editable bool `json:"editable"`
	IsRegistrationAvailable bool `json:"is_registration_available"`
	Locked bool `json:"locked"`
	PackageUnitsCount int32 `json:"package_units_count"`
	CreatedAt string `json:"created_at"`
	DeclineReason FbpDraftGetResponseDeclineReason `json:"decline_reason"`
	ID int64 `json:"id"`
	BundleID string `json:"bundle_id"`
	IsCancelable bool `json:"is_cancelable"`
	IsDeletable bool `json:"is_deletable"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GooglerpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type DescriptionCategoryTipsResponseResult struct {
	TypeID int64 `json:"type_id"`
	ImagesURL []interface{} `json:"images_url"`
	InfoURL string `json:"info_url"`
}

type ArrivalpassArrivalPassListResponseArrivalPass struct {
	ArrivalReasons []interface{} `json:"arrival_reasons"`
	ArrivalTime string `json:"arrival_time"`
	DriverPhone string `json:"driver_phone"`
	DropoffPointID int64 `json:"dropoff_point_id"`
	IsActive bool `json:"is_active"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	ArrivalPassID int64 `json:"arrival_pass_id"`
	DriverName string `json:"driver_name"`
	VehicleModel string `json:"vehicle_model"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GetProductInfoListResponseCommission struct {
	SaleSchema string `json:"sale_schema"`
	Value float64 `json:"value"`
	DeliveryAmount float64 `json:"delivery_amount"`
	Percent float64 `json:"percent"`
	ReturnAmount float64 `json:"return_amount"`
}

type V2ProductInfoPicturesResponseError struct {
	Message string `json:"message"`
	URL string `json:"url"`
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"`
	PostingNumber string `json:"posting_number"`
}

type FbpGetLabelResponseLabelCreationStateTypeEnum string

type ReportMarkedProductsSalesCreateRequestDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date ReportMarkedProductsSalesCreateRequestDate `json:"date"`
}

type V1ReviewInfoResponse struct {
	DislikesAmount int32 `json:"dislikes_amount"`
	Status string `json:"status"`
	VideosAmount int32 `json:"videos_amount"`
	SKU int64 `json:"sku"`
	CommentsAmount int32 `json:"comments_amount"`
	ID string `json:"id"`
	LikesAmount int32 `json:"likes_amount"`
	PhotosAmount int32 `json:"photos_amount"`
	PublishedAt string `json:"published_at"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	Videos []interface{} `json:"videos"`
	Text string `json:"text"`
	OrderStatus string `json:"order_status"`
	Photos []interface{} `json:"photos"`
	Rating int32 `json:"rating"`
}

type V1AssemblyCarriagePostingListResponse struct {
	CanPrintMassLabel bool `json:"can_print_mass_label"`
	Cursor string `json:"cursor"`
	Postings []interface{} `json:"postings"`
}

type V1AddStrategyItemsRequest struct {
	ProductID []interface{} `json:"product_id"`
	StrategyID string `json:"strategy_id"`
}

type V3GetProductInfoListResponsePromotionType string

type V3GetProductInfoListResponsePromotion struct {
	IsEnabled bool `json:"is_enabled"`
	Type V3GetProductInfoListResponsePromotionType `json:"type_"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1DeleteStrategyItemsResponse struct {
	Result V1DeleteStrategyItemsResponseResult `json:"result"`
}

type V1FbpDraftDirectTimeslotEditResponse struct {
	RowVersion int64 `json:"row_version"`
	ErrorReasons []interface{} `json:"error_reasons"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1FbpDraftDirectGetTimeslotResponse struct {
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
}

type QuestionV1QuestionAnswerListResponseAnswerStatusPublicationEnum string

type V1QuestionAnswerListResponseAnswers struct {
	SKU int64 `json:"sku"`
	StatusPublication QuestionV1QuestionAnswerListResponseAnswerStatusPublicationEnum `json:"status_publication"`
	Text string `json:"text"`
	AuthorName string `json:"author_name"`
	ID string `json:"id"`
	PublishedAt interface{} `json:"published_at"`
	QuestionID string `json:"question_id"`
}

type GetUploadQuotaResponseDailyUpdate struct {
	Usage int64 `json:"usage"`
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
}

type V1FbpDraftPickUpRegistrateResponse struct {
	Error V1FbpDraftPickUpRegistrateResponseRegistrationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftDirectCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFeeFee struct {
	Accrued MoneyMoneyAccrued `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type SellerSellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type V4GetProductAttributesResponseModelInfo struct {
	ModelID int64 `json:"model_id"`
	Count int64 `json:"count"`
}

type ProductCurrencyEnum string

type AnalyticsProductQueriesRequestSortBy string

type V1FbpOrderPickUpDlvEditResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1SellerActionsUpdateInstallmentRequestActionParameters struct {
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type V1GetStrategyItemInfoRequest struct {
	ProductID int64 `json:"product_id"`
}

type V1AddStrategyItemsResponseResult struct {
	Errors []interface{} `json:"errors"`
	FailedProductCount int32 `json:"failed_product_count"`
}

type MoneyMoney struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexData struct {
	PriceIndex float64 `json:"price_index"`
	URL string `json:"url"`
	MinPrice MoneyMoney `json:"min_price"`
}

type MoneyMoneySelf struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexDataSelf struct {
	URL string `json:"url"`
	MinPrice MoneyMoneySelf `json:"min_price"`
	PriceIndex float64 `json:"price_index"`
}

type ProductV1ProductPricesDetailsResponsePricePriceIndex struct {
	ExternalIndexData ProductV1ProductPricesDetailsResponsePricePriceIndexIndexData `json:"external_index_data"`
	SelfIndexData ProductV1ProductPricesDetailsResponsePricePriceIndexIndexDataSelf `json:"self_index_data"`
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V2DeliveryMethodListV2Response struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	DeliveryMethods []interface{} `json:"delivery_methods"`
}

type MoneyMoneyCustomerPrice struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V2FbsPostingProductCountrySetResponse struct {
	ProductID int64 `json:"product_id"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
}

type V2ReturnsRfbsGetRequest struct {
	ReturnID int64 `json:"return_id"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplar struct {
	Errors []interface{} `json:"errors"`
	GTD string `json:"gtd"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	Valid bool `json:"valid"`
}

type FbsPostingProductExemplarCreateOrGetV6ResponseProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	IsJwUinNeeded bool `json:"is_jw_uin_needed"`
	IsMandatoryMarkNeeded bool `json:"is_mandatory_mark_needed"`
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
	HasImei bool `json:"has_imei"`
	IsGTDNeeded bool `json:"is_gtd_needed"`
	IsMandatoryMarkPossible bool `json:"is_mandatory_mark_possible"`
	IsRnptNeeded bool `json:"is_rnpt_needed"`
}

type PostingV4PostingFbsListResponsePostingsFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
}

type PostingPostingProductCancelRequest struct {
	CancelReasonMessage string `json:"cancel_reason_message"`
	Items []interface{} `json:"items"`
	PostingNumber string `json:"posting_number"`
	CancelReasonID int64 `json:"cancel_reason_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission struct {
	Amount float64 `json:"amount"`
	Currency string `json:"currency"`
	Percent int64 `json:"percent"`
}

type ReportCreateCompanyPostingsReportRequestFilter struct {
	CancelReasonID []interface{} `json:"cancel_reason_id"`
	DeliverySchema []interface{} `json:"delivery_schema"`
	ProcessedAtFrom string `json:"processed_at_from"`
	SKU []interface{} `json:"sku"`
	StatusAlias []interface{} `json:"status_alias"`
	Statuses []interface{} `json:"statuses"`
	Title string `json:"title"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	OfferID string `json:"offer_id"`
	ProcessedAtTo string `json:"processed_at_to"`
	WarehouseID []interface{} `json:"warehouse_id"`
	IsExpress interface{} `json:"is_express"`
}

type V3Address struct {
	AddressTail string `json:"address_tail"`
	Comment string `json:"comment"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	ZipCode string `json:"zip_code"`
	City string `json:"city"`
	Country string `json:"country"`
	District string `json:"district"`
	Region string `json:"region"`
}

type V3Customer struct {
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address V3Address `json:"address"`
}

type ProductV1ProductVisibilityInfoRequest struct {
	Skus []interface{} `json:"skus"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type SellerApiProductV1ResponseResult struct {
	ProductIds []interface{} `json:"product_ids"`
	Rejected []interface{} `json:"rejected"`
}

type V1SellerActionsUpdateDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters V1SellerActionsUpdateDiscountRequestActionParameters `json:"action_parameters"`
}

type ReturnsCompanyFbsInfoRequestPagination struct {
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type Fbsv4PostingProductDetailWithoutDimensions struct {
	MandatoryMark interface{} `json:"mandatory_mark"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	CurrencyCode string `json:"currency_code"`
}

type DetailsServices struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type ProductImportProductsPricesRequest struct {
	Prices []interface{} `json:"prices"`
}

type ProductV1ProductVisibilitySetRequest struct {
	ItemPlacement []interface{} `json:"item_placement"`
}

type PriceIndexesIndexDataOzon struct {
	MinimalPrice string `json:"minimal_price"`
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type PostingV4PostingFbsUnfulfilledListResponse struct {
	Count int64 `json:"count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type ReturnsRfbsGetV2ResponseClientReturnMethodType struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type ReturnsRfbsGetResponseReturns struct {
	AvailableActions []interface{} `json:"available_actions"`
	CreatedAt string `json:"created_at"`
	RejectionComment string `json:"rejection_comment"`
	ReturnReason ReturnsRfbsGetV2ResponseReturnReason `json:"return_reason"`
	Comment string `json:"comment"`
	Product V2Product `json:"product"`
	ReturnMethodDescription string `json:"return_method_description"`
	RuPostTrackingNumber string `json:"ru_post_tracking_number"`
	State V2ReturnsRfbsGetV2ResponseState `json:"state"`
	ClientName string `json:"client_name"`
	ClientPhoto []interface{} `json:"client_photo"`
	ClientReturnMethodType ReturnsRfbsGetV2ResponseClientReturnMethodType `json:"client_return_method_type"`
	PostingNumber string `json:"posting_number"`
	RejectionReason []interface{} `json:"rejection_reason"`
	WarehouseID int64 `json:"warehouse_id"`
	OrderNumber string `json:"order_number"`
	ReturnNumber string `json:"return_number"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequest struct {
	IsKgt bool `json:"is_kgt"`
	Search ListDropOffPointsForCreateFBSWarehouseRequestSearch `json:"search"`
	Coordinates V1ListDropOffPointsForCreateFBSWarehouseRequestCoordinates `json:"coordinates"`
	CountryCode string `json:"country_code"`
}

type V1FbpDraftDirectRegistrateResponseRegistrationError struct {
	OrderError V1OrderErrorTypeEnum `json:"order_error"`
	BundleErrors []interface{} `json:"bundle_errors"`
}

type V2PostingFBSActGetProducts struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
}

type PriceIndexesIndexDataExternal struct {
	MinimalPriceCurrency string `json:"minimal_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
	MinimalPrice string `json:"minimal_price"`
}

type PriceIndexesColorIndex string

type GetProductInfoListResponsePriceIndexes struct {
	ColorIndex PriceIndexesColorIndex `json:"color_index"`
	ExternalIndexData PriceIndexesIndexDataExternal `json:"external_index_data"`
	OzonIndexData PriceIndexesIndexDataOzon `json:"ozon_index_data"`
	SelfMarketplacesIndexData PriceIndexesIndexDataSelf `json:"self_marketplaces_index_data"`
}

type V1FbpEditTimeslotRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1ProductGetRelatedSKUResponse struct {
	Items interface{} `json:"items"`
	Errors interface{} `json:"errors"`
}

type LanguageLanguage string

type V1GetAttributesRequest struct {
	Language LanguageLanguage `json:"language"`
	TypeID int64 `json:"type_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
}

type V1FbpDraftPickUpRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPosting struct {
	Products []interface{} `json:"products"`
	DeliverySchema string `json:"delivery_schema"`
	DeliverySpeed int32 `json:"delivery_speed"`
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"`
}

type ItemBundleSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemCount int64 `json:"total_item_count"`
	TotalQuantity int64 `json:"total_quantity"`
}

type V2GetProductInfoStocksByWarehouseFbsRequestV2 struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	OfferID []interface{} `json:"offer_id"`
	SKU []interface{} `json:"sku"`
}

type CompanyTaxSystemEnum string

type SellerInfoResponseCompany struct {
	OwnershipForm string `json:"ownership_form"`
	TaxSystem CompanyTaxSystemEnum `json:"tax_system"`
	Country string `json:"country"`
	Currency string `json:"currency"`
	Inn string `json:"inn"`
	LegalName string `json:"legal_name"`
	Name string `json:"name"`
	Ogrn string `json:"ogrn"`
}

type GetRealizationReportByDayResponse struct {
	Rows []interface{} `json:"rows"`
}

type WarehouseFirstMileType struct {
	DropoffTimeslotID int64 `json:"dropoff_timeslot_id"`
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
	FirstMileType string `json:"first_mile_type"`
	DropoffPointID string `json:"dropoff_point_id"`
}

type Productv2ProductsStocksResponse struct {
	Result []interface{} `json:"result"`
}

type V1SellerActionsProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	Skus []interface{} `json:"skus"`
}

type V1FbpDraftDirectSellerDlvEditRequest struct {
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
}

type V1FbpDraftDropOffProductValidateResponseApprovedItem struct {
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
}

type MoneyMoneyTotalAccrued struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type Postingv3GetFbsPostingUnfulfilledListResponseResult struct {
	Count int64 `json:"count"`
	Postings []interface{} `json:"postings"`
}

type Productv3GetProductListRequest struct {
	Filter Productv3GetProductListRequestFilter `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
}

type CreateCompanyPostingsReportRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	CustomerData bool `json:"customer_data"`
	JewelryCodes bool `json:"jewelry_codes"`
	AdditionalData bool `json:"additional_data"`
}

type ReportCreateCompanyPostingsReportRequest struct {
	With CreateCompanyPostingsReportRequestWith `json:"with"`
	Filter ReportCreateCompanyPostingsReportRequestFilter `json:"filter"`
	Language ReportLanguage `json:"language"`
}

type V4Visibility string

type V3User struct {
	ID string `json:"id"`
	Type string `json:"type_"`
}

type V3ChatMessage struct {
	CreatedAt string `json:"created_at"`
	Data []interface{} `json:"data"`
	IsImage bool `json:"is_image"`
	IsRead bool `json:"is_read"`
	MessageId int64 `json:"messageId"`
	ModerateImageStatus ChatMessageModerateImageStatus `json:"moderate_image_status"`
	User V3User `json:"user"`
	Context ChatMessageContext `json:"context"`
}

type V3ImportProductsResponseResult struct {
	TaskID int64 `json:"task_id"`
}

type V1ProductUpdateOfferIdRequest struct {
	UpdateOfferID interface{} `json:"update_offer_id"`
}

type V1CommentListRequest struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
	ReviewID string `json:"review_id"`
	SortDir V1CommentSort `json:"sort_dir"`
}

type V1AddBarcodeResult struct {
	Code string `json:"code"`
	Error string `json:"error"`
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type StockShipmentType string

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1Barcode struct {
	Barcode string `json:"barcode"`
	SKU int64 `json:"sku"`
}

type FilterOp interface{}

type AnalyticsFilter struct {
	Key string `json:"key"`
	Op FilterOp `json:"op"`
	Value string `json:"value"`
}

type V1FbpEditTimeslotResponseReserveFailureType string

type V1ProductUpdateOfferIdResponseError struct {
	Message string `json:"message"`
	OfferID string `json:"offer_id"`
}

type V3TimeRange struct {
	From string `json:"from"`
	To string `json:"to"`
}

type Postingv3GetFbsPostingUnfulfilledListRequestFilter struct {
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveringDateTo string `json:"delivering_date_to"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	ProviderID []interface{} `json:"provider_id"`
	Status string `json:"status"`
	DeliveringDateFrom string `json:"delivering_date_from"`
	LastChangedStatusDate V3TimeRange `json:"last_changed_status_date"`
	FbpFilter string `json:"fbpFilter"`
	WarehouseID []interface{} `json:"warehouse_id"`
}

type Postingv3FbsPostingWithParams struct {
	LegalInfo bool `json:"legal_info"`
	Translit bool `json:"translit"`
	AnalyticsData bool `json:"analytics_data"`
	Barcodes bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
}

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	Filter Postingv3GetFbsPostingUnfulfilledListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With Postingv3FbsPostingWithParams `json:"with"`
	Dir string `json:"dir"`
}

type V1SellerActionsCreateDiscountWithConditionRequestDiscountTypeEnum string

type V1WarehouseFbsCreatePickUpTimeslotListResponse struct {
	IsPickupSupported bool `json:"is_pickup_supported"`
	Timeslots []interface{} `json:"timeslots"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequest struct {
	AutoAddDate string `json:"auto_add_date"`
	ToUpdate []interface{} `json:"to_update"`
	ActionID int64 `json:"action_id"`
}

type ReturnsCompanyFbsInfoResponsePassInfo struct {
	IsRequired bool `json:"is_required"`
	Count int32 `json:"count"`
}

type ReturnsCompanyFbsInfoResponseDropOffPoints struct {
	WarehousesIds []interface{} `json:"warehouses_ids"`
	Address string `json:"address"`
	BoxCount int32 `json:"box_count"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	PassInfo ReturnsCompanyFbsInfoResponsePassInfo `json:"pass_info"`
	ReturnsCount int32 `json:"returns_count"`
	PlaceID int64 `json:"place_id"`
	UtcOffset string `json:"utc_offset"`
}

type PostingV4PostingFbsUnfulfilledListRequestSortDirEnum string

type V1WarehouseFbsUpdateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ValidationResultValidationErrorTypeEnum string

type FinanceV1GetFinanceAccrualPostingsRequest struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type V1FbpDraftDirectCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails V1FbpDraftDirectCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1GenerateBarcodeRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1FbpOrderListResponseItem struct {
	ReceiveDate string `json:"receive_date"`
	BundleSummary ItemBundleSummary `json:"bundle_summary"`
	CreatedDate string `json:"created_date"`
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	WarehouseID int64 `json:"warehouse_id"`
	AttentionReasons []interface{} `json:"attention_reasons"`
	ID int64 `json:"id"`
	SupplyID string `json:"supply_id"`
	PackageUnitsCount int32 `json:"package_units_count"`
	CanBeCancelled bool `json:"can_be_cancelled"`
	CancellationState V1CancellationState `json:"cancellation_state"`
	HasConsignmentNote bool `json:"has_consignment_note"`
	Locked bool `json:"locked"`
	OrderNumber string `json:"order_number"`
	HasLabel bool `json:"has_label"`
	Status V1OrderStatusEnum `json:"status"`
}

type GetUploadQuotaResponseDailyCreate struct {
	Limit int64 `json:"limit"`
	ResetAt string `json:"reset_at"`
	Usage int64 `json:"usage"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDelivery struct {
	Services []interface{} `json:"services"`
	TotalAccrued MoneyMoneyTotalAccrued `json:"total_accrued"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProduct struct {
	Commission FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductCommission `json:"commission"`
	Delivery FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDelivery `json:"delivery"`
	SKU int64 `json:"sku"`
}

type FbsPostingDetailCourier struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	CarModel string `json:"car_model"`
	CarNumber string `json:"car_number"`
}

type ActionsV1ActionsAutoAddProductsListResponseProducts struct {
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantityToAutoAdd int64 `json:"quantity_to_auto_add"`
	SKU int64 `json:"sku"`
	ActionPriceToAutoAdd float64 `json:"action_price_to_auto_add"`
	MarketplaceSellerPrice float64 `json:"marketplace_seller_price"`
	MinActionQuantity int64 `json:"min_action_quantity"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	AddMode string `json:"add_mode"`
	Currency string `json:"currency"`
	MaxDiscountPrice float64 `json:"max_discount_price"`
}

type V2ServiceType string

type PriceIndexesIndexExternalData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1SetProductStairwayDiscountByQuantityResponseError struct {
	Data []interface{} `json:"data"`
	SKU int64 `json:"sku"`
}

type ActionsV1ActionsAutoAddProductsListResponse struct {
	Total int64 `json:"total"`
	Products []interface{} `json:"products"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData struct {
	ClusterTo string `json:"cluster_to"`
	Products []interface{} `json:"products"`
	ClusterFrom string `json:"cluster_from"`
}

type SearchQueriesTextRequestSortBy string

type V1FbpDraftDirectDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type SortingOrder string

type V1UpdateWarehouseFBSFirstMileRequestFirstMileTypeEnum string

type V1DayOfWeek string

type V1FbpOrderDropOffTimetableResponseCalendar struct {
	DayOfWeek V1DayOfWeek `json:"day_of_week"`
	CalendarItem V1FbpOrderDropOffTimetableResponseCalendarCalendarItem `json:"calendar_item"`
}

type V1WarehouseInvalidProductsGetRequest struct {
	LastID int64 `json:"last_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1SellerActionsProductsAddRequest struct {
	ActionID int64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type SetPostingsResponseResult struct {
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
	Result bool `json:"result"`
}

type V1FbpDraftPickUpProductValidateResponseApprovedItem struct {
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
}

type FbsPostingShipV4RequestPackage struct {
	Products []interface{} `json:"products"`
}

type V1OrderDraftValidationErrorErrorType string

type V2GetProductInfoStocksByWarehouseFbsResponseV2Product struct {
	Reserved int64 `json:"reserved"`
	SKU int64 `json:"sku"`
	WarehouseID int64 `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	FreeStock int64 `json:"free_stock"`
	OfferID string `json:"offer_id"`
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
}

type V2FbsPostingProductCountryListResponse struct {
	Result []interface{} `json:"result"`
}

type ItemCommissionsv5 struct {
	SalesPercentRFBS float64 `json:"sales_percent_rfbs"`
	FBODelivToCustomerAmount float64 `json:"fbo_deliv_to_customer_amount"`
	FBODirectFlowTransMaxAmount float64 `json:"fbo_direct_flow_trans_max_amount"`
	FBSDelivToCustomerAmount float64 `json:"fbs_deliv_to_customer_amount"`
	FBSDirectFlowTransMaxAmount float64 `json:"fbs_direct_flow_trans_max_amount"`
	FBSFirstMileMaxAmount float64 `json:"fbs_first_mile_max_amount"`
	SalesPercentFBO float64 `json:"sales_percent_fbo"`
	SalesPercentFBP float64 `json:"sales_percent_fbp"`
	FBODirectFlowTransMinAmount float64 `json:"fbo_direct_flow_trans_min_amount"`
	FBOReturnFlowAmount float64 `json:"fbo_return_flow_amount"`
	FBSDirectFlowTransMinAmount float64 `json:"fbs_direct_flow_trans_min_amount"`
	FBSFirstMileMinAmount float64 `json:"fbs_first_mile_min_amount"`
	FBSReturnFlowAmount float64 `json:"fbs_return_flow_amount"`
	SalesPercentFBS float64 `json:"sales_percent_fbs"`
}

type Productv4GetProductAttributesV4Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type ProductBooleanResponse struct {
	Result bool `json:"result"`
}

type FinanceTransactionTotalsV3RequestDate struct {
	From string `json:"from"`
	To string `json:"to"`
}

type Financev3FinanceTransactionTotalsV3Request struct {
	Date FinanceTransactionTotalsV3RequestDate `json:"date"`
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
}

type V1ArchiveStatusEnum string

type GetWarehouseFBSOperationStatusResponseTypeEnum string

type ActionsV1ActionsAutoAddProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	ProductIds []interface{} `json:"product_ids"`
}

type V1QuestionTopSkuResponse struct {
	SKU interface{} `json:"sku"`
}

type V1ApproveDeclineDiscountTasksResponse struct {
	Result V1ApproveDeclineDiscountTasksResponseResult `json:"result"`
}

type AnalyticsMetric string

type V1CommentCreateRequest struct {
	MarkReviewAsProcessed bool `json:"mark_review_as_processed"`
	ParentCommentID string `json:"parent_comment_id"`
	ReviewID string `json:"review_id"`
	Text string `json:"text"`
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V1Competitor struct {
	Coefficient float64 `json:"coefficient"`
	CompetitorID int64 `json:"competitor_id"`
}

type V1ArchiveDeclineReason struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type PickedSegmentSegmentTypeEnum string

type V3ChatListRequestFilter struct {
	ChatStatus string `json:"chat_status"`
	UnreadOnly bool `json:"unread_only"`
}

type HumanTextsParam struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type PostingV4PostingFbsListResponsePostingsCancellation struct {
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
}

type V1OrderDraftValidationError struct {
	Errors []interface{} `json:"errors"`
}

type V1GetProductInfoSubscriptionResponseResult struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type Productv4GetProductAttributesV4ResponseResult struct {
	DimensionUnit string `json:"dimension_unit"`
	ID int64 `json:"id"`
	Attributes []interface{} `json:"attributes"`
	Barcodes interface{} `json:"barcodes"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Depth int64 `json:"depth"`
	Name string `json:"name"`
	PrimaryImage string `json:"primary_image"`
	SKU string `json:"sku"`
	WeightUnit string `json:"weight_unit"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Images interface{} `json:"images"`
	ModelInfo V4GetProductAttributesResponseModelInfo `json:"model_info"`
	OfferID string `json:"offer_id"`
	Weight int64 `json:"weight"`
	Width int64 `json:"width"`
	AttributesWithDefaults []interface{} `json:"attributes_with_defaults"`
	Barcode string `json:"barcode"`
	ColorImage string `json:"color_image"`
	Height int64 `json:"height"`
	PDFList []interface{} `json:"pdf_list"`
	TypeID int64 `json:"type_id"`
}

type PostingV1PostingFbpListResponsePostingsProducts struct {
	Quantity int32 `json:"quantity"`
	SellerPrice MoneyMoneySellerPrice `json:"seller_price"`
	SKU int64 `json:"sku"`
	CustomerPrice MoneyMoneyCustomerPrice `json:"customer_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price MoneyPostingMoney `json:"price"`
}

type PostingPostingFBSPackageLabelResponse struct {
	ContentType string `json:"content_type"`
	FileContent string `json:"file_content"`
	FileName string `json:"file_name"`
}

type PolygonBindRequestwhLocation struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type V2FbsPostingProduct struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	Quantity int64 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1QuestionInfoResponse struct {
	Status interface{} `json:"status"`
	Text string `json:"text"`
	ID string `json:"id"`
	ProductURL string `json:"product_url"`
	PublishedAt interface{} `json:"published_at"`
	QuestionLink string `json:"question_link"`
	AnswersCount int64 `json:"answers_count"`
	AuthorName string `json:"author_name"`
	SKU int64 `json:"sku"`
}

type PostingFinancialDataProduct struct {
	CommissionPercent int64 `json:"commission_percent"`
	CommissionsCurrencyCode string `json:"commissions_currency_code"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	CustomerPrice float64 `json:"customer_price"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	CurrencyCode string `json:"currency_code"`
	CustomerCurrencyCode string `json:"customer_currency_code"`
	CommissionAmount float64 `json:"commission_amount"`
	Payout float64 `json:"payout"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	Actions []interface{} `json:"actions"`
}

type ImportProductsRequestPdfList struct {
	Name string `json:"name"`
	SrcURL string `json:"src_url"`
	Index int64 `json:"index"`
}

type FbsPostingProductExemplarSetV6RequestProducts struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type V1FbpDraftDirectSellerDlvEditResponse struct {
	Error V1OrderDraftValidationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type Polygonv1PolygonBindRequest struct {
	DeliveryMethodID int32 `json:"delivery_method_id"`
	Polygons []interface{} `json:"polygons"`
	WarehouseLocation PolygonBindRequestwhLocation `json:"warehouse_location"`
}

type Productv2ProductsStocksResponseResult struct {
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Updated bool `json:"updated"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductProductUnarchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V1ItemError struct {
	State string `json:"state"`
	Level string `json:"level"`
	Description string `json:"description"`
	Field string `json:"field"`
	AttributeID int64 `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
	Code string `json:"code"`
	Message string `json:"message"`
}

type Postingv1GetCarriageAvailableListResponse struct {
	Result interface{} `json:"result"`
}

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type ReportCreateDiscountedRequest interface{}

type V1WarehouseFbsCreatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type ProductImportProductsBySKUResponseResult struct {
	TaskID int64 `json:"task_id"`
	UnmatchedSKUList []interface{} `json:"unmatched_sku_list"`
}

type V1UpdateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
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

type ReportReportInfoResponse struct {
	Result ReportReportinfo `json:"result"`
}

type ArrivalpassArrivalPassUpdateRequestArrivalPass struct {
	ArrivalTime string `json:"arrival_time"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	ArrivalPassID int64 `json:"arrival_pass_id"`
}

type GetProductInfoListResponseVisibilityDetails struct {
	HasPrice bool `json:"has_price"`
	HasStock bool `json:"has_stock"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

type V2CancellationInitiatorEnum string

type GetConditionalCancellationListV2ResponseCancellationReason struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type GetConditionalCancellationListV2ResponseResult struct {
	OrderDate string `json:"order_date"`
	State GetConditionalCancellationListV2ResponseState `json:"state"`
	ApproveDate string `json:"approve_date"`
	CancelledAt string `json:"cancelled_at"`
	PostingNumber string `json:"posting_number"`
	SourceID int64 `json:"source_id"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	ApproveComment string `json:"approve_comment"`
	AutoApproveDate string `json:"auto_approve_date"`
	CancellationID int64 `json:"cancellation_id"`
	CancellationInitiator V2CancellationInitiatorEnum `json:"cancellation_initiator"`
	CancellationReason GetConditionalCancellationListV2ResponseCancellationReason `json:"cancellation_reason"`
	CancellationReasonMessage string `json:"cancellation_reason_message"`
}

type V3ImportProductsRequestDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type V1GetProductStairwayDiscountByQuantityRequest struct {
	Skus []interface{} `json:"skus"`
}

type V4FbsPostingShipPackageV4Request struct {
	Products []interface{} `json:"products"`
	PostingNumber string `json:"posting_number"`
}

type PostinglistV3status struct {
	From string `json:"from"`
	To string `json:"to"`
}

type ReturnsRfbsGetV2ResponseRejectionReason struct {
	Hint string `json:"hint"`
	ID int32 `json:"id"`
	IsCommentRequired bool `json:"is_comment_required"`
	Name string `json:"name"`
}

type V1FbpDraftDropOffProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"`
}

type V1FbpCheckConsignmentNoteStateRequest struct {
	Code string `json:"code"`
	SupplyID string `json:"supply_id"`
}

type V1FbpArchiveListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type MoneyMoneyTotalAmount struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrual struct {
	Date string `json:"date"`
	ItemFees FinanceV1GetFinanceAccrualByDayResponseAccrualItemFees `json:"item_fees"`
	NonItemFee FinanceV1GetFinanceAccrualByDayResponseAccrualNonItemFee `json:"non_item_fee"`
	Posting FinanceV1GetFinanceAccrualByDayResponseAccrualPosting `json:"posting"`
	TotalAmount MoneyMoneyTotalAmount `json:"total_amount"`
	AccrualID int64 `json:"accrual_id"`
	UnitNumber string `json:"unit_number"`
	AccruedCategory FinanceV1GetFinanceAccrualByDayResponseAccrualAccruedCategoryEnum `json:"accrued_category"`
}

type V3ImportProductsRequest struct {
	Items []interface{} `json:"items"`
}

type PostingErrorTypeEnum string

type SellerApiProductV1ResponseResultDeactivate struct {
	ProductIds []interface{} `json:"product_ids"`
	Rejected []interface{} `json:"rejected"`
}

type RowItem struct {
	Barcode string `json:"barcode"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	SKU int64 `json:"sku"`
}

type RowItemCommissionReturn struct {
	Bonus float64 `json:"bonus"`
	Compensation float64 `json:"compensation"`
	Quantity int32 `json:"quantity"`
	StandardFee float64 `json:"standard_fee"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Stars float64 `json:"stars"`
	Commission float64 `json:"commission"`
	PricePerInstance float64 `json:"price_per_instance"`
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
}

type RowItemLegalEntityDocument struct {
	Number string `json:"number"`
	SaleDate string `json:"sale_date"`
}

type RowItemCommission struct {
	Compensation float64 `json:"compensation"`
	Stars float64 `json:"stars"`
	Total float64 `json:"total"`
	Commission float64 `json:"commission"`
	PricePerInstance float64 `json:"price_per_instance"`
	Quantity int32 `json:"quantity"`
	StandardFee float64 `json:"standard_fee"`
	BankCoinvestment float64 `json:"bank_coinvestment"`
	Amount float64 `json:"amount"`
	Bonus float64 `json:"bonus"`
}

type V1GetRealizationReportPostingResponseRow struct {
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
	Item RowItem `json:"item"`
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"row_number"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	Order RowItemOrder `json:"order"`
	LegalEntityDocument RowItemLegalEntityDocument `json:"legal_entity_document"`
}

type FbsPostingProductExemplarSetV6RequestExemplars struct {
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
	ExemplarID int64 `json:"exemplar_id"`
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
}

type ProductGetProductInfoDescriptionResponseProduct struct {
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Description string `json:"description"`
	ID int64 `json:"id"`
}

type V1CreatePricingStrategyResponseResult struct {
	StrategyID string `json:"strategy_id"`
}

type V1CreatePricingStrategyResponse struct {
	Result V1CreatePricingStrategyResponseResult `json:"result"`
}

type V1ProductGetRelatedSKURequest struct {
	SKU interface{} `json:"sku"`
}

type V1GetFBSRatingIndexInfoV1Response struct {
	PeriodTo string `json:"period_to"`
	ProcessingCostsSum float64 `json:"processing_costs_sum"`
	CurrencyCode string `json:"currency_code"`
	Defects []interface{} `json:"defects"`
	Index float64 `json:"index"`
	PeriodFrom string `json:"period_from"`
}

type V1GetProductInfoDiscountedRequest struct {
	DiscountedSkus interface{} `json:"discounted_skus"`
}

type V1FbpDraftDropOffCreateRequestDeliveryDetails struct {
	DropOffDate string `json:"drop_off_date"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	DropOffProvinceUuid string `json:"drop_off_province_uuid"`
}

type GetProductInfoListResponseAvailability struct {
	Reasons []interface{} `json:"reasons"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
	Availability string `json:"availability"`
}

type ProductImportProductsPricesRequestPrice struct {
	AutoActionEnabled string `json:"auto_action_enabled"`
	CurrencyCode string `json:"currency_code"`
	MinPrice string `json:"min_price"`
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	NetPrice string `json:"net_price"`
	OfferID string `json:"offer_id"`
	OldPrice string `json:"old_price"`
	Price string `json:"price"`
	ManageElasticBoostingThroughPrice bool `json:"manage_elastic_boosting_through_price"`
	PriceStrategyEnabled string `json:"price_strategy_enabled"`
	ProductID int64 `json:"product_id"`
	Vat string `json:"vat"`
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []interface{} `json:"result"`
}

type V1SellerActionsCreateVoucherRequestDiscountTypeEnum string

type V1SellerActionsCreateVoucherRequestVoucherParameterTypeEnum string

type V1SellerActionsCreateVoucherRequestVoucherParameter struct {
	IsPrivate bool `json:"is_private"`
	Type V1SellerActionsCreateVoucherRequestVoucherParameterTypeEnum `json:"type_"`
	CountCodes int64 `json:"count_codes"`
}

type V1SellerActionsCreateVoucherRequest struct {
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
	VoucherParameters V1SellerActionsCreateVoucherRequestVoucherParameter `json:"voucher_parameters"`
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountType V1SellerActionsCreateVoucherRequestDiscountTypeEnum `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
}

type ParameterDiscountLevels struct {
	OrderAmount float64 `json:"order_amount"`
	DiscountValue float64 `json:"discount_value"`
}

type V1FbpOrderDirectSellerDlvEditResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1UpdateWarehouseFBSFirstMileRequest struct {
	FirstMileType V1UpdateWarehouseFBSFirstMileRequestFirstMileTypeEnum `json:"first_mile_type"`
	TimeslotID int64 `json:"timeslot_id"`
	WarehouseID int64 `json:"warehouse_id"`
	CutInTime int64 `json:"cut_in_time"`
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type V1FbpDraftDropOffPointTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ArchiveSkuSummary struct {
	RoundedTotalVolumeInLitres float64 `json:"rounded_total_volume_in_litres"`
	TotalItemsCount int64 `json:"total_items_count"`
	TotalQuantity int64 `json:"total_quantity"`
}

type RelatedPostingCancelReason struct {
	PostingNumber string `json:"posting_number"`
	Reasons []interface{} `json:"reasons"`
}

type SearchQueriesTextRequestSortDir string

type GetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type V1ProductPricesDetailsRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1FbpGetLabelResponse struct {
	LabelURL string `json:"label_url"`
	State FbpGetLabelResponseLabelCreationStateTypeEnum `json:"state"`
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

type FinanceV1GetFinanceAccrualByDayResponseAccrualPostingProductDeliveryService struct {
	Accrued MoneyMoneyAccrued `json:"accrued"`
	TypeID int32 `json:"type_id"`
}

type RelatedPostingCancelReasons struct {
	Title string `json:"title"`
	TypeID string `json:"type_id"`
	ID int64 `json:"id"`
}

type CompensationReportLanguage string

type SellerInfoResponseSubscriptionTypeEnum string

type SellerInfoResponseSubscription struct {
	IsPremium bool `json:"is_premium"`
	Type SellerInfoResponseSubscriptionTypeEnum `json:"type_"`
}

type V1SellerInfoResponse struct {
	Company SellerInfoResponseCompany `json:"company"`
	Ratings []interface{} `json:"ratings"`
	Subscription SellerInfoResponseSubscription `json:"subscription"`
}

type AnalyticsAnalyticsGetDataRequest struct {
	Sort []interface{} `json:"sort"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Dimension []interface{} `json:"dimension"`
	Filters []interface{} `json:"filters"`
	Limit int64 `json:"limit"`
	Metrics []interface{} `json:"metrics"`
	Offset int64 `json:"offset"`
}

type V1FbpOrderDropOffDlvEditRequest struct {
	DropOffDate string `json:"drop_off_date"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1SellerActionsCreateDiscountWithConditionResponse struct {
	ActionID int64 `json:"action_id"`
}

type Financev3FinanceTransactionListV3Request struct {
	Page int64 `json:"page"`
	PageSize int64 `json:"page_size"`
	Filter FinanceTransactionListV3RequestFilter `json:"filter"`
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpOrderDirectCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1GetStrategyItemsResponseResult struct {
	ProductID []interface{} `json:"product_id"`
}

type V1GetStrategyItemsResponse struct {
	Result V1GetStrategyItemsResponseResult `json:"result"`
}

type V1AnalyticsProductQueriesResponseItem struct {
	OfferID string `json:"offer_id"`
	SKU int64 `json:"sku"`
	ViewConversion float64 `json:"view_conversion"`
	Gmv float64 `json:"gmv"`
	Position float64 `json:"position"`
	UniqueSearchUsers int64 `json:"unique_search_users"`
	UniqueViewUsers int64 `json:"unique_view_users"`
	Category string `json:"category"`
	Currency string `json:"currency"`
	Name string `json:"name"`
}

type V1FbpDraftPickUpRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V2PostingFBSActGetPostingsRequest struct {
	ID interface{} `json:"id"`
}

type V1FbpDraftDropOffRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type ArrivalPassListRequestFilter struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	ArrivalReason string `json:"arrival_reason"`
	DropoffPointIds []interface{} `json:"dropoff_point_ids"`
	OnlyActivePasses bool `json:"only_active_passes"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1TimeRangeStorageTariffication struct {
	TimeTo string `json:"time_to"`
	TimeFrom string `json:"time_from"`
}

type V3Addressee struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V3FbsPostingRequirementsV3 struct {
	ProductsRequiringGTD []interface{} `json:"products_requiring_gtd"`
	ProductsRequiringCountry []interface{} `json:"products_requiring_country"`
	ProductsRequiringMandatoryMark []interface{} `json:"products_requiring_mandatory_mark"`
	ProductsRequiringJwUin []interface{} `json:"products_requiring_jw_uin"`
	ProductsRequiringRnpt []interface{} `json:"products_requiring_rnpt"`
	ProductsRequiringImei []interface{} `json:"products_requiring_imei"`
	ProductsRequiringChangeCountry []interface{} `json:"products_requiring_change_country"`
}

type V3Barcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V3FbsPostingProductExemplarsV3 struct {
	Products []interface{} `json:"products"`
}

type V3Cancellation struct {
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
}

type V3DeliveryMethod struct {
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
}

type V3FbsPostingDetail struct {
	Cancellation V3Cancellation `json:"cancellation"`
	FactDeliveryDate string `json:"fact_delivery_date"`
	OrderNumber string `json:"order_number"`
	ProviderStatus string `json:"provider_status"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	AdditionalData []interface{} `json:"additional_data"`
	DeliveringDate string `json:"delivering_date"`
	ShipmentDate string `json:"shipment_date"`
	PreviousSubstatus string `json:"previous_substatus"`
	FinancialData V3PostingFinancialData `json:"financial_data"`
	Courier FbsPostingDetailCourier `json:"courier"`
	Customer V3Customer `json:"customer"`
	DeliveryMethod V3DeliveryMethod `json:"delivery_method"`
	LegalInfo V2FboSinglePostingLegalInfo `json:"legal_info"`
	Optional V3FbsPostingDetailOptional `json:"optional"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	Addressee V3Addressee `json:"addressee"`
	InProcessAt string `json:"in_process_at"`
	IsExpress bool `json:"is_express"`
	OrderID int64 `json:"order_id"`
	ParentPostingNumber string `json:"parent_posting_number"`
	Requirements V3FbsPostingRequirementsV3 `json:"requirements"`
	Substatus string `json:"substatus"`
	TrackingNumber string `json:"tracking_number"`
	AnalyticsData V3FbsPostingAnalyticsData `json:"analytics_data"`
	AvailableActions interface{} `json:"available_actions"`
	Barcodes V3Barcodes `json:"barcodes"`
	ProductExemplars V3FbsPostingProductExemplarsV3 `json:"product_exemplars"`
	Tariffication V3FbsTariffication `json:"tariffication"`
	Status string `json:"status"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	DeliveryPrice string `json:"delivery_price"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	RelatedPostings V3FbsPostingDetailRelatedPostings `json:"related_postings"`
}

type V4FbsPostingShipPackageV4RequestProduct struct {
	ExemplarsIds []interface{} `json:"exemplarsIds"`
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

type SellerApiProductV1ResponseProduct struct {
	ProductID float64 `json:"product_id"`
	Reason string `json:"reason"`
}

type RpcStatus struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V3ChatList struct {
	Filter V3ChatListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
}

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairwayStep struct {
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
	Discount int64 `json:"discount"`
}

type V1ProductGetRelatedSKUResponseError struct {
	Code string `json:"code"`
	SKU int64 `json:"sku"`
	Message string `json:"message"`
}

type V1SellerActionsUpdateDiscountWithConditionRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters V1SellerActionsUpdateDiscountWithConditionRequestActionParameters `json:"action_parameters"`
}

type V1ProductActionTimerUpdateRequest struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1ReviewChangeStatusRequest struct {
	ReviewIds []interface{} `json:"review_ids"`
	Status string `json:"status"`
}

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementEnum string

type FinanceV1GetFinanceAccrualByDayRequest struct {
	Date string `json:"date"`
	LastID string `json:"last_id"`
}

type V1FbpOrderDropOffTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type WarehouseTimetable struct {
	TimetableTo string `json:"timetable_to"`
	WorkingHours []interface{} `json:"working_hours"`
	TimetableFrom string `json:"timetable_from"`
}

type DeliveryMethodListV2RequestSortDirEnum string

type ErrorErrorLevel string

type Fbpv1Timeslot struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type TaskAutoModeratedInfo struct {
	MaxPercent float64 `json:"max_percent"`
	MaxPrice float64 `json:"max_price"`
	MinPercent float64 `json:"min_percent"`
	MinPrice float64 `json:"min_price"`
}

type ExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type SellerSellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CarriageID int64 `json:"carriage_id"`
}

type V1CarriageCancelResponse struct {
	Error string `json:"error"`
	CarriageStatus string `json:"carriage_status"`
}

type V1Empty interface{}

type ErrorData struct {
	Code string `json:"code"`
	Field string `json:"field"`
	Message string `json:"message"`
	Step int64 `json:"step"`
	Value string `json:"value"`
}

type ProductUpdateOfferIdRequestUpdateOfferId struct {
	NewOfferID string `json:"new_offer_id"`
	OfferID string `json:"offer_id"`
}

type V1Money struct {
	Currency string `json:"currency"`
	Value float64 `json:"value"`
}

type ReportListResponseResult struct {
	Reports []interface{} `json:"reports"`
	Total int32 `json:"total"`
}

type ReportReportListResponse struct {
	Result ReportListResponseResult `json:"result"`
}

type ItemSize struct {
	HeightMm int32 `json:"height_mm"`
	LengthMm int32 `json:"length_mm"`
	WidthMm int32 `json:"width_mm"`
}

type ProductImportProductsPricesResponseProcessResult struct {
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Updated bool `json:"updated"`
}

type V1AnalyticsProductQueriesDetailsRequestSortBy string

type PostingCancelReasonListResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftDropOffCreateRequest struct {
	DeliveryDetails V1FbpDraftDropOffCreateRequestDeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
}

type V1FbpAvailableTimeslotListRequest struct {
	IntervalStart string `json:"interval_start"`
	SupplyID string `json:"supply_id"`
	IntervalEnd string `json:"interval_end"`
}

type V1FbpCreateActRequest struct {
	SupplyID string `json:"supply_id"`
}

type V2CancellationStateEnumFilters string

type V1OrderAttentionTypeEnum string

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts struct {
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
	OldPrice float64 `json:"old_price"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
	Commission PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission `json:"commission"`
	CustomerPrice MoneyPostingMoney `json:"customer_price"`
	Payout float64 `json:"payout"`
	Price float64 `json:"price"`
}

type V1BundleItemErrorEnum string

type V1UpdateWarehouseFBSRequestWorkingDaysEnum string

type V3GetFbsPostingListResponseV3Result struct {
	HasNext bool `json:"has_next"`
	Postings []interface{} `json:"postings"`
}

type SellerApiGetSellerProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Limit float64 `json:"limit"`
	LastID float64 `json:"last_id"`
}

type V1FbpDraftDropOffRegistrateResponseRegistrationError struct {
	BundleErrors []interface{} `json:"bundle_errors"`
	OrderError V1OrderErrorTypeEnum `json:"order_error"`
}

type V1FbpDraftDropOffRegistrateResponse struct {
	Error V1FbpDraftDropOffRegistrateResponseRegistrationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1AddBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type Productv3GetProductListResponseResult struct {
	Items interface{} `json:"items"`
	LastID string `json:"last_id"`
	Total int32 `json:"total"`
}

type Productv3GetProductListResponse struct {
	Result Productv3GetProductListResponseResult `json:"result"`
}

type V1ProductUpdateAttributesRequestItem struct {
	Attributes interface{} `json:"attributes"`
	OfferID string `json:"offer_id"`
}

type V1GetAttributeValuesRequest struct {
	TypeID int64 `json:"type_id"`
	AttributeID int64 `json:"attribute_id"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Language LanguageLanguage `json:"language"`
	LastValueID int64 `json:"last_value_id"`
	Limit int64 `json:"limit"`
}

type V1GetStrategyListResponse struct {
	Strategies []interface{} `json:"strategies"`
	Total int32 `json:"total"`
}

type V1ItemIDsRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V1ReviewListResponse struct {
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
	HasNext bool `json:"has_next"`
}

type ProductGetImportProductsInfoResponseResult struct {
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type ProductGetImportProductsInfoResponse struct {
	Result ProductGetImportProductsInfoResponseResult `json:"result"`
}

type V1ItemResponse struct {
	ProductID int64 `json:"product_id"`
	Quant int32 `json:"quant"`
	ShipmentType V1BundleItemShipmentType `json:"shipment_type"`
	Tags []interface{} `json:"tags"`
	PlacementZone string `json:"placement_zone"`
	IconPath string `json:"icon_path"`
	SKU int64 `json:"sku"`
	Name string `json:"name"`
	Barcode string `json:"barcode"`
	ContractorItemCode string `json:"contractor_item_code"`
	SfboAttribute V1ItemSfboAttribute `json:"sfbo_attribute"`
	OfferID string `json:"offer_id"`
	VolumeInLitres float64 `json:"volume_in_litres"`
	TotalVolumeInLitres float64 `json:"total_volume_in_litres"`
	Quantity int32 `json:"quantity"`
	IsQuantEditable bool `json:"is_quant_editable"`
}

type ProductGetProductInfoDescriptionResponse struct {
	Result ProductGetProductInfoDescriptionResponseProduct `json:"result"`
}

type FbpWarehouseListResponseAddressDetailing struct {
	City string `json:"city"`
	Country string `json:"country"`
	House string `json:"house"`
	Region string `json:"region"`
	Street string `json:"street"`
	Zipcode string `json:"zipcode"`
}

type FbpWarehouseListResponseWarehouse struct {
	AddressDetailing FbpWarehouseListResponseAddressDetailing `json:"address_detailing"`
	ID int64 `json:"id"`
	IsBonded bool `json:"is_bonded"`
	Name string `json:"name"`
	PartnerName string `json:"partner_name"`
	SupplyTypes []interface{} `json:"supply_types"`
	TimezoneName string `json:"timezone_name"`
}

type V1GetProductRatingBySkuResponse struct {
	Products interface{} `json:"products"`
}

type FinanceCashFlowStatementListResponseService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V1FbpCheckActStateRequest struct {
	FileUuid string `json:"file_uuid"`
}

type ItemPricev5 struct {
	NetPrice float64 `json:"net_price"`
	Price float64 `json:"price"`
	RetailPrice float64 `json:"retail_price"`
	AutoActionEnabled bool `json:"auto_action_enabled"`
	MarketingSellerPrice float64 `json:"marketing_seller_price"`
	MinPrice float64 `json:"min_price"`
	OldPrice float64 `json:"old_price"`
	Vat float64 `json:"vat"`
	CurrencyCode string `json:"currency_code"`
}

type V1FbpDraftDropOffCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type ActionsV1ActionsAutoAddProductsDeleteResponse struct {
	ProductIds []interface{} `json:"product_ids"`
}

type V1GetStrategyItemInfoResponseResult struct {
	PriceDownloadedAt string `json:"price_downloaded_at"`
	StrategyCompetitorID int64 `json:"strategy_competitor_id"`
	StrategyCompetitorProductURL string `json:"strategy_competitor_product_url"`
	StrategyID string `json:"strategy_id"`
	IsEnabled bool `json:"is_enabled"`
	StrategyProductPrice int32 `json:"strategy_product_price"`
}

type V1GetStrategyItemInfoResponse struct {
	Result V1GetStrategyItemInfoResponseResult `json:"result"`
}

type V1GetReturnsListResponse struct {
	Returns []interface{} `json:"returns"`
	HasNext bool `json:"has_next"`
}

type V1FbpDraftListResponseItem struct {
	SupplyID string `json:"supply_id"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
	CancellationState V1CancellationState `json:"cancellation_state"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	ID int64 `json:"id"`
	IsCancelable bool `json:"is_cancelable"`
	IsDeletable bool `json:"is_deletable"`
	Status V1DraftStatusEnum `json:"status"`
	DeliveryDetails V1fbpDeliveryDetails `json:"delivery_details"`
	Editable bool `json:"editable"`
	Locked bool `json:"locked"`
	PackageUnitsCount int32 `json:"package_units_count"`
}

type V1FbpOrderPickUpCancelResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type ValidationErrorCharacteristicEnum string

type ValidationResultValidationError struct {
	Type ValidationResultValidationErrorTypeEnum `json:"type_"`
	Characteristic ValidationErrorCharacteristicEnum `json:"characteristic"`
	RestrictionPrice V1Money `json:"restriction_price"`
	RestrictionVwc float64 `json:"restriction_vwc"`
	TemplateID int32 `json:"template_id"`
}

type V1FbpCheckActStateResponseStatus string

type ErrorHumanTexts struct {
	ShortDescription string `json:"short_description"`
	AttributeName string `json:"attribute_name"`
	Description string `json:"description"`
	HintCode string `json:"hint_code"`
	Message string `json:"message"`
	Params []interface{} `json:"params"`
}

type GetProductInfoListResponseError struct {
	Level ErrorErrorLevel `json:"level"`
	State string `json:"state"`
	Texts ErrorHumanTexts `json:"texts"`
	AttributeID int64 `json:"attribute_id"`
	Code string `json:"code"`
	Field string `json:"field"`
}

type V1SellerActionsCreateInstallmentResponse struct {
	ActionID int64 `json:"action_id"`
}

type GetConditionalCancellationListV2RequestFilters struct {
	CancellationInitiator []interface{} `json:"cancellation_initiator"`
	PostingNumber []interface{} `json:"posting_number"`
	State V2CancellationStateEnumFilters `json:"state"`
}

type GetConditionalCancellationListV2RequestWith struct {
	Counter bool `json:"counter"`
}

type V2GetConditionalCancellationListV2Request struct {
	Filters GetConditionalCancellationListV2RequestFilters `json:"filters"`
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
	With GetConditionalCancellationListV2RequestWith `json:"with"`
}

type GetProductAttributesV4ResponseAttribute struct {
	ID int64 `json:"id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type V1AddBarcodeRequest struct {
	Barcodes []interface{} `json:"barcodes"`
}

type V1FbpOrderPickUpCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1GetTreeResponse struct {
	Result []interface{} `json:"result"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime struct {
	TimeslotEnd string `json:"timeslot_end"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem struct {
	IsHoliday bool `json:"is_holiday"`
	OpeningHours V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTimeOpening `json:"opening_hours"`
	BreakHours V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItemTimeslotWithTime `json:"break_hours"`
}

type V1FbpDraftDropOffPointTimetableResponseCalendar struct {
	CalendarItem V1FbpDraftDropOffPointTimetableResponseCalendarCalendarItem `json:"calendar_item"`
	DayOfWeek V1DayOfWeek `json:"day_of_week"`
}

type V1UnarchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V3ImportProductsResponse struct {
	Result V3ImportProductsResponseResult `json:"result"`
}

type V1AssemblyFbsProductListRequestSortDirEnum string

type V1GetSupplyOrderBundleResponse struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Items []interface{} `json:"items"`
	TotalCount int32 `json:"total_count"`
}

type PostingV4PostingFbsListResponsePostingsCustomerAddress struct {
	City string `json:"city"`
	Comment string `json:"comment"`
	Country string `json:"country"`
	District string `json:"district"`
	Latitude float64 `json:"latitude"`
	AddressTail string `json:"address_tail"`
	Longitude float64 `json:"longitude"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	Region string `json:"region"`
	ZipCode string `json:"zip_code"`
}

type V1SellerActionsProductsCandidatesRequest struct {
	Limit int64 `json:"limit"`
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
}

type ApproveDeclineDiscountTasksResponseFailDetail struct {
	ErrorForUser string `json:"error_for_user"`
	TaskID int64 `json:"task_id"`
}

type SellerApiProductPrice struct {
	ProductID float64 `json:"product_id"`
	ActionPrice float64 `json:"action_price"`
	Stock float64 `json:"stock"`
}

type V3ChatHistoryResponse struct {
	HasNext bool `json:"has_next"`
	Messages []interface{} `json:"messages"`
}

type GetRealizationReportResponseV2Row struct {
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
	Item RowItem `json:"item"`
}

type V1GenerateBarcodeResponse struct {
	Errors []interface{} `json:"errors"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsProducts struct {
	Imei []interface{} `json:"imei"`
	OfferID string `json:"offer_id"`
	Price MoneyPostingMoney `json:"price"`
	ProductColor string `json:"product_color"`
	SKU int64 `json:"sku"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	Name string `json:"name"`
	Quantity int32 `json:"quantity"`
	Weight float64 `json:"weight"`
}

type GetFBSRatingIndexInfoV1ResponseIndexDynamics struct {
	Date string `json:"date"`
	IndexByDate float64 `json:"index_by_date"`
	ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"`
}

type V1ProductPricesDetailsResponse struct {
	Prices []interface{} `json:"prices"`
}

type V1FbpDraftDirectProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type MoneyMoneyCurrentTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequestActionParameters struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type V1FbpDraftDropOffProvinceListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type PostingBooleanResponse struct {
	Result bool `json:"result"`
}

type DetailsService struct {
	Items []interface{} `json:"items"`
	Total float64 `json:"total"`
}

type FinanceCashFlowStatementListResponseDetailsOthers struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type PickedSegmentSegment struct {
	Description string `json:"description"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	Type PickedSegmentSegmentTypeEnum `json:"type_"`
}

type FbpDraftDropOffPointListResponseDropOffPoint struct {
	City string `json:"city"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	NearestDropOffDate string `json:"nearest_drop_off_date"`
	PointAddress string `json:"point_address"`
	ProvinceUuid string `json:"province_uuid"`
}

type V1ReviewListRequest struct {
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortDir string `json:"sort_dir"`
	Status string `json:"status"`
}

type V1TimeRangeReturnDate struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type V1TimeRangeVisualStatus struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type GetReturnsListRequestFilter struct {
	VisualStatusChangeMoment V1TimeRangeVisualStatus `json:"visual_status_change_moment"`
	OrderID int64 `json:"order_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
	VisualStatusName string `json:"visual_status_name"`
	Barcode string `json:"barcode"`
	ProductName string `json:"product_name"`
	OfferID string `json:"offer_id"`
	WarehouseID int64 `json:"warehouse_id"`
	ReturnSchema string `json:"return_schema"`
	CompensationStatusID int32 `json:"compensation_status_id"`
	LogisticReturnDate V1TimeRangeReturnDate `json:"logistic_return_date"`
	StorageTarifficationStartDate V1TimeRangeStorageTariffication `json:"storage_tariffication_start_date"`
}

type V1GetReturnsListRequest struct {
	Limit int32 `json:"limit"`
	LastID int64 `json:"last_id"`
	Filter GetReturnsListRequestFilter `json:"filter"`
}

type ImportProductRequestPromotion struct {
	Type string `json:"type_"`
	Operation string `json:"operation"`
}

type V1FbpOrderListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type PostingV4PostingFbsListResponsePostingsProducts struct {
	IsMarketplaceBuyout bool `json:"is_marketplace_buyout"`
	Name string `json:"name"`
	ProductColor string `json:"product_color"`
	Imei []interface{} `json:"imei"`
	OfferID string `json:"offer_id"`
	Price MoneyPostingMoney `json:"price"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	Weight float64 `json:"weight"`
	IsBlrTraceable bool `json:"is_blr_traceable"`
}

type V1CarriageCreateRequest struct {
	AllBlrTraceable bool `json:"all_blr_traceable"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseExtremelyLowPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type ResultError struct {
	Code string `json:"code"`
	Status string `json:"status"`
}

type FbsPostingShipV4ResponseShipAdditionalData struct {
	Products interface{} `json:"products"`
	PostingNumber string `json:"posting_number"`
}

type WarehouseFirstMile struct {
	TimeslotTo string `json:"timeslot_to"`
	Type FirstMileTypeEnum `json:"type_"`
	DropoffPointID string `json:"dropoff_point_id"`
	FirstMileIsChanging bool `json:"first_mile_is_changing"`
	TimeslotFrom string `json:"timeslot_from"`
	TimeslotID int64 `json:"timeslot_id"`
}

type V1CreateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type ProductV1ProductVisibilitySetResponse struct {
	Items []interface{} `json:"items"`
	ItemsErrors []interface{} `json:"items_errors"`
}

type V1FbpDraftPickUpDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type ReviewV2ReviewListV2Response struct {
	HasNext bool `json:"has_next"`
	LastID string `json:"last_id"`
	Reviews []interface{} `json:"reviews"`
}

type ProductGetImportProductsInfoRequest struct {
	TaskID int64 `json:"task_id"`
}

type V1FbpOrderDirectSellerDlvEditRequest struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleType string `json:"vehicle_type"`
	DriverName string `json:"driver_name"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPointDropOffPointTypeEnum string

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPoint struct {
	LastTransitTimeLocal TypeTimeOfDay `json:"last_transit_time_local"`
	Type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPointDropOffPointTypeEnum `json:"type_"`
	Address string `json:"address"`
	Coordinates DropOffPointCoordinates `json:"coordinates"`
	DiscountPercent float64 `json:"discount_percent"`
	ID string `json:"id"`
}

type V1FbpDraftDirectSellerDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1FbpWarehouseListResponse struct {
	Warehouses []interface{} `json:"warehouses"`
}

type V1GetCompensationReportRequest struct {
	Date string `json:"date"`
	Language CompensationReportLanguage `json:"language"`
}

type DetailsDeliveryDetails struct {
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	DeliveryServices DetailsServices `json:"delivery_services"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequest struct {
	AddressCoordinates V1WarehouseFbsCreatePickUpTimeslotListRequestAddressCoordinates `json:"address_coordinates"`
	IsKgt bool `json:"is_kgt"`
}

type V1ReturnsCompanyFbsInfoRequest struct {
	Filter V1ReturnsCompanyFbsInfoRequestFilter `json:"filter"`
	Pagination ReturnsCompanyFbsInfoRequestPagination `json:"pagination"`
}

type GetProductInfoListResponseStatuses struct {
	ValidationStatus string `json:"validation_status"`
	ModerateStatus string `json:"moderate_status"`
	Status string `json:"status"`
	StatusDescription string `json:"status_description"`
	StatusName string `json:"status_name"`
	StatusTooltip string `json:"status_tooltip"`
	StatusUpdatedAt string `json:"status_updated_at"`
	IsCreated bool `json:"is_created"`
	StatusFailed string `json:"status_failed"`
}

type GetProductInfoListResponseModelInfo struct {
	Count int64 `json:"count"`
	ModelID int64 `json:"model_id"`
}

type V3GetProductInfoListResponseItem struct {
	ModelInfo GetProductInfoListResponseModelInfo `json:"model_info"`
	CurrencyCode string `json:"currency_code"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	Stocks GetProductInfoListResponseStocks `json:"stocks"`
	HasDiscountedFBOItem bool `json:"has_discounted_fbo_item"`
	Images360 []interface{} `json:"images360"`
	IsSuper bool `json:"is_super"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	Price string `json:"price"`
	PrimaryImage []interface{} `json:"primary_image"`
	TypeID int64 `json:"type_id"`
	UpdatedAt string `json:"updated_at"`
	VisibilityDetails GetProductInfoListResponseVisibilityDetails `json:"visibility_details"`
	DiscountedFBOStocks int32 `json:"discounted_fbo_stocks"`
	IsArchived bool `json:"is_archived"`
	IsAutoarchived bool `json:"is_autoarchived"`
	IsDiscounted bool `json:"is_discounted"`
	IsKgt bool `json:"is_kgt"`
	PriceIndexes GetProductInfoListResponsePriceIndexes `json:"price_indexes"`
	ColorImage []interface{} `json:"color_image"`
	Statuses GetProductInfoListResponseStatuses `json:"statuses"`
	Availabilities []interface{} `json:"availabilities"`
	Images []interface{} `json:"images"`
	MinPrice string `json:"min_price"`
	OfferID string `json:"offer_id"`
	Promotions []interface{} `json:"promotions"`
	SKU int64 `json:"sku"`
	Sources []interface{} `json:"sources"`
	VolumeWeight float64 `json:"volume_weight"`
	Barcodes []interface{} `json:"barcodes"`
	Commissions []interface{} `json:"commissions"`
	ID int64 `json:"id"`
	OldPrice string `json:"old_price"`
	Vat string `json:"vat"`
	Errors []interface{} `json:"errors"`
	IsPrepaymentAllowed bool `json:"is_prepayment_allowed"`
}

type MoneyMoneyNextTariffCharge struct {
	Currency string `json:"currency"`
	Amount string `json:"amount"`
}

type MoneyMoneyNextTariffMinCharge struct {
	Amount string `json:"amount"`
	Currency string `json:"currency"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication struct {
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	CurrentTariffCharge MoneyMoneyCurrentTariffCharge `json:"current_tariff_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	NextTariffCharge MoneyMoneyNextTariffCharge `json:"next_tariff_charge"`
	NextTariffMinCharge MoneyMoneyNextTariffMinCharge `json:"next_tariff_min_charge"`
	NextTariffType string `json:"next_tariff_type"`
	CurrentTariffMinCharge MoneyMoneyCurrentTariffMinCharge `json:"current_tariff_min_charge"`
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffRate float64 `json:"next_tariff_rate"`
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type V3ChatHistoryRequest struct {
	ChatID string `json:"chat_id"`
	Direction string `json:"direction"`
	Filter ChatHistoryRequestFilter `json:"filter"`
	FromMessageID int64 `json:"from_message_id"`
	Limit int64 `json:"limit"`
}

type V1ProductInfoWrongVolumeRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1GetCompetitorsResponse struct {
	Competitor []interface{} `json:"competitor"`
	Total int32 `json:"total"`
}

type V1SetProductStairwayDiscountByQuantityRequest struct {
	Stairways []interface{} `json:"stairways"`
	SuppressWarnings bool `json:"suppress_warnings"`
}

type V1FbpDraftPickupDlvEditRequestDeliveryDetails struct {
	Date string `json:"date"`
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
	Address string `json:"address"`
	Comment string `json:"comment"`
}

type RolesByTokenResponseRoles struct {
	Name string `json:"name"`
	Methods []interface{} `json:"methods"`
}

type PostingV4PostingFbsListResponsePostingsDeliveryMethod struct {
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
}

type SellerSellerAPIArrivalPassCreateRequestArrivalPass struct {
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
}

type Financev3Period struct {
	From string `json:"from"`
	To string `json:"to"`
}

type V1AssemblyFbsPostingListRequestFilter struct {
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
}

type V1AssemblyFbsPostingListRequestSortDirEnum string

type V1AssemblyFbsPostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter V1AssemblyFbsPostingListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir V1AssemblyFbsPostingListRequestSortDirEnum `json:"sort_dir"`
}

type V1FbpAvailableTimeslotListResponse struct {
	WarehouseTimezoneName string `json:"warehouse_timezone_name"`
	Reasons []interface{} `json:"reasons"`
	Timeslots []interface{} `json:"timeslots"`
}

type WarehouseWarehouseListResponse struct {
	Result []interface{} `json:"result"`
}

type V3AddresseeFbsLists struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type V1FbpDraftDropOffRegistrateRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type V1ProductUpdateDiscountResponse struct {
	Result bool `json:"result"`
}

type FbpCheckActStateResponseErrorReason string

type V1FbpCheckActStateResponse struct {
	CdnURL string `json:"cdn_url"`
	Error FbpCheckActStateResponseErrorReason `json:"error"`
	Status V1FbpCheckActStateResponseStatus `json:"status"`
}

type V1SearchAttributeValuesResponse struct {
	Result []interface{} `json:"result"`
}

type V1SellerActionsProductsListResponseProductQuantTypeEnum string

type V1SellerActionsProductsListResponseProduct struct {
	MinSellerPrice float64 `json:"min_seller_price"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	QuantSize int64 `json:"quant_size"`
	DiscountPercent float64 `json:"discount_percent"`
	IsActive bool `json:"is_active"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantType V1SellerActionsProductsListResponseProductQuantTypeEnum `json:"quant_type"`
	SKU []interface{} `json:"sku"`
	ActionPrice float64 `json:"action_price"`
	BasePrice float64 `json:"base_price"`
	Currency string `json:"currency"`
}

type PostingV4PostingFbsListResponsePostingsLegalInfo struct {
	CompanyName string `json:"company_name"`
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
}

type PostingV4PostingFbsListResponsePostingsExternalOrder struct {
	IsExternal bool `json:"is_external"`
	PlatformName string `json:"platform_name"`
}

type GetProductInfoStocksResponseStock struct {
	Reserved int32 `json:"reserved"`
	ShipmentType StockShipmentType `json:"shipment_type"`
	SKU int64 `json:"sku"`
	Type string `json:"type_"`
	WarehouseIds []interface{} `json:"warehouse_ids"`
	Present int32 `json:"present"`
}

type PostingV4PostingFbsListResponsePostingsCustomer struct {
	Address PostingV4PostingFbsListResponsePostingsCustomerAddress `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
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

type V2ProductInfoPicturesRequest struct {
	ProductID interface{} `json:"product_id"`
}

type V1FbsPostingProductExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"`
}

type V4GetUploadQuotaResponse struct {
	DailyCreate GetUploadQuotaResponseDailyCreate `json:"daily_create"`
	DailyUpdate GetUploadQuotaResponseDailyUpdate `json:"daily_update"`
	OperationLimits ProductV4GetUploadQuotaResponseOperationLimits `json:"operation_limits"`
	Total GetUploadQuotaResponseTotal `json:"total"`
}

type V1FbpGetLabelRequest struct {
	SupplyID string `json:"supply_id"`
	Code string `json:"code"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearch struct {
	Types []interface{} `json:"types"`
	Address string `json:"address"`
}

type PostingCancelReasonResponse struct {
	Result []interface{} `json:"result"`
}

type V1QuestionChangeStatusRequest struct {
	QuestionIds interface{} `json:"question_ids"`
	Status string `json:"status"`
}

type Productv3GetProductListResponseItemQuant struct {
	QuantCode string `json:"quant_code"`
	QuantSize int64 `json:"quant_size"`
}

type ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementListEnum string

type ChatChatSendMessageResponse struct {
	Result string `json:"result"`
}

type AnalyticsDataRowDimension struct {
	Name string `json:"name"`
	ID string `json:"id"`
}

type OperationPosting struct {
	DeliverySchema string `json:"delivery_schema"`
	OrderDate string `json:"order_date"`
	PostingNumber string `json:"posting_number"`
	WarehouseID int64 `json:"warehouse_id"`
}

type FbsPostingBarcodes struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type V2FbsPosting struct {
	PostingNumber string `json:"posting_number"`
	Barcodes FbsPostingBarcodes `json:"barcodes"`
	OrderNumber string `json:"order_number"`
	Products []interface{} `json:"products"`
	ShipmentDate string `json:"shipment_date"`
	Status string `json:"status"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CreatedAt string `json:"created_at"`
	InProcessAt string `json:"in_process_at"`
	OrderID int64 `json:"order_id"`
}

type V2FbsPostingResponse struct {
	Result V2FbsPosting `json:"result"`
}

type PostingProductCancelRequestItem struct {
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProduct struct {
	Exemplars []interface{} `json:"exemplars"`
	ProductID int64 `json:"product_id"`
}

type ReportCreateCompanyProductsReportRequestVisibility string

type ValidationResultItem struct {
	SKU int64 `json:"sku"`
	WeightG float64 `json:"weight_g"`
	Size ItemSize `json:"size"`
}

type ValidationResultItemStateEnum string

type WarehouseInvalidProductsGetResponseValidationResult struct {
	State ValidationResultItemStateEnum `json:"state"`
	ValidationErrors []interface{} `json:"validation_errors"`
	Item ValidationResultItem `json:"item"`
}

type GetRealizationReportByDayResponseRow struct {
	CommissionRatio float64 `json:"commission_ratio"`
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
	Item RowItem `json:"item"`
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"`
	SellerPricePerInstance float64 `json:"seller_price_per_instance"`
}

type ExemplarsMarks struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V1FbpOrderDropOffCancelRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1SellerActionsCreateDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	MinActionPercent float64 `json:"min_action_percent"`
	Title string `json:"title"`
}

type V1SearchQueriesTextResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type V1FbpArchiveListResponseItem struct {
	ActFileUuid string `json:"act_file_uuid"`
	DeclineReason V1ArchiveDeclineReason `json:"decline_reason"`
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	ExternalOrderID string `json:"external_order_id"`
	HasAct bool `json:"has_act"`
	ReceiveDate string `json:"receive_date"`
	RowVersion int64 `json:"row_version"`
	BundleID string `json:"bundle_id"`
	WarehouseID int64 `json:"warehouse_id"`
	CreatedDate string `json:"created_date"`
	OrderDraftID int64 `json:"order_draft_id"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WhcOrderID int64 `json:"whc_order_id"`
	BundleSKUSummary V1ArchiveSkuSummary `json:"bundle_sku_summary"`
	HasLabel bool `json:"has_label"`
	Status V1ArchiveStatusEnum `json:"status"`
	SupplyID string `json:"supply_id"`
}

type Polygonv1Empty interface{}

type PostingV4PostingFbsListResponsePostingsTariffication struct {
	CurrentTariffCharge MoneyMoneyCurrentTariffCharge `json:"current_tariff_charge"`
	CurrentTariffMinCharge MoneyMoneyCurrentTariffMinCharge `json:"current_tariff_min_charge"`
	CurrentTariffRate float64 `json:"current_tariff_rate"`
	NextTariffRate float64 `json:"next_tariff_rate"`
	CurrentTariffType string `json:"current_tariff_type"`
	NextTariffCharge MoneyMoneyNextTariffCharge `json:"next_tariff_charge"`
	NextTariffMinCharge MoneyMoneyNextTariffMinCharge `json:"next_tariff_min_charge"`
	NextTariffStartsAt string `json:"next_tariff_starts_at"`
	NextTariffType string `json:"next_tariff_type"`
}

type PostingV4PostingFbsListResponsePostingsAddressee struct {
	Name string `json:"name"`
}

type PostingV4PostingFbsListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []interface{} `json:"products_with_possible_mandatory_mark"`
}

type PostingV4PostingFbsListResponsePostingsAnalyticsData struct {
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryDateEnd string `json:"delivery_date_end"`
	DeliveryType string `json:"delivery_type"`
	IsLegal bool `json:"is_legal"`
	IsPremium bool `json:"is_premium"`
	Warehouse string `json:"warehouse"`
	City string `json:"city"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region string `json:"region"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	WarehouseID int64 `json:"warehouse_id"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
}

type PostingV4PostingFbsListResponsePostings struct {
	DeliveryMethod PostingV4PostingFbsListResponsePostingsDeliveryMethod `json:"delivery_method"`
	Status string `json:"status"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	Requirements PostingV4PostingFbsListResponsePostingsRequirements `json:"requirements"`
	DestinationPlaceName string `json:"destination_place_name"`
	IsMultibox bool `json:"is_multibox"`
	Products []interface{} `json:"products"`
	PrrOption string `json:"prr_option"`
	ShipmentDate string `json:"shipment_date"`
	Tariffication PostingV4PostingFbsListResponsePostingsTariffication `json:"tariffication"`
	TrackingNumber string `json:"tracking_number"`
	Addressee PostingV4PostingFbsListResponsePostingsAddressee `json:"addressee"`
	IsPresortable bool `json:"is_presortable"`
	AvailableActions []interface{} `json:"available_actions"`
	Barcodes PostingV4PostingFbsListResponsePostingsBarcodes `json:"barcodes"`
	ExternalOrder PostingV4PostingFbsListResponsePostingsExternalOrder `json:"external_order"`
	IsExpress bool `json:"is_express"`
	LegalInfo PostingV4PostingFbsListResponsePostingsLegalInfo `json:"legal_info"`
	OrderID int64 `json:"order_id"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	Cancellation PostingV4PostingFbsListResponsePostingsCancellation `json:"cancellation"`
	DeliverySchema string `json:"delivery_schema"`
	InProcessAt string `json:"in_process_at"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	OrderNumber string `json:"order_number"`
	ParentPostingNumber string `json:"parent_posting_number"`
	QuantumID int64 `json:"quantum_id"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	Customer PostingV4PostingFbsListResponsePostingsCustomer `json:"customer"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	VolumeWeight float64 `json:"volume_weight"`
	DeliveringDate string `json:"delivering_date"`
	FinancialData PostingV4PostingFbsListResponsePostingsFinancialData `json:"financial_data"`
	Optional PostingV4PostingFbsListResponsePostingsOptional `json:"optional"`
	PostingNumber string `json:"posting_number"`
	Substatus string `json:"substatus"`
	AnalyticsData PostingV4PostingFbsListResponsePostingsAnalyticsData `json:"analytics_data"`
}

type ChatRead struct {
	ChatID string `json:"chat_id"`
	FromMessageID int64 `json:"from_message_id"`
}

type PostingV4PostingFbsListResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Postings interface{} `json:"postings"`
}

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairway struct {
	Steps []interface{} `json:"steps"`
}

type StatusEnum string

type V1GetProductStairwayDiscountByQuantityResponseStairways struct {
	Stairway V1GetProductStairwayDiscountByQuantityResponseStairwaysStairway `json:"stairway"`
	Status StatusEnum `json:"status"`
	Enabled bool `json:"enabled"`
	SKU int64 `json:"sku"`
}

type SearchAttributeValuesResponseValue struct {
	Info string `json:"info"`
	Picture string `json:"picture"`
	Value string `json:"value"`
	ID int64 `json:"id"`
}

type GetSellerActionsV1ResponseAction struct {
	DateStart string `json:"date_start"`
	OrderAmount float64 `json:"order_amount"`
	DiscountValue float64 `json:"discount_value"`
	ID float64 `json:"id"`
	FreezeDate string `json:"freeze_date"`
	PotentialProductsCount float64 `json:"potential_products_count"`
	ParticipatingProductsCount float64 `json:"participating_products_count"`
	IsVoucherAction bool `json:"is_voucher_action"`
	Title string `json:"title"`
	Description string `json:"description"`
	WithTargeting bool `json:"with_targeting"`
	DiscountType string `json:"discount_type"`
	ActionType string `json:"action_type"`
	DateEnd string `json:"date_end"`
	AutoAddDates []interface{} `json:"auto_add_dates"`
	IsParticipating bool `json:"is_participating"`
	BannedProductsCount float64 `json:"banned_products_count"`
}

type ChatInfoChatTypeEnum string

type V1ApproveDiscountTasksRequestTask struct {
	ID int64 `json:"id"`
	ApprovedPrice float64 `json:"approved_price"`
	SellerComment string `json:"seller_comment"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
}

type PostingPostingFBSActGetContainerLabelsResponse struct {
	FileName string `json:"file_name"`
	ContentType string `json:"content_type"`
	FileContent string `json:"file_content"`
}

type ProductImportProductsPricesResponseError struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

type V1SellerActionsCreateDiscountWithConditionRequest struct {
	DateStart string `json:"date_start"`
	DiscountType V1SellerActionsCreateDiscountWithConditionRequestDiscountTypeEnum `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
}

type ReportListRequestReportType string

type V4GetProductInfoStocksRequestFilter struct {
	Visibility V4Visibility `json:"visibility"`
	WithQuant FilterWithQuant `json:"with_quant"`
	OfferID []interface{} `json:"offer_id"`
	ProductID []interface{} `json:"product_id"`
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation struct {
	CancelReason string `json:"cancel_reason"`
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancellationInitiator string `json:"cancellation_initiator"`
	CancellationType string `json:"cancellation_type"`
	CancelledAfterShip bool `json:"cancelled_after_ship"`
	AffectCancellationRating bool `json:"affect_cancellation_rating"`
}

type V4GetProductInfoStocksResponseItem struct {
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	Stocks []interface{} `json:"stocks"`
}

type FinanceV1GetFinanceAccrualTypesResponse struct {
	AccrualTypes []interface{} `json:"accrual_types"`
}

type ProductImportProductsBySKURequestItem struct {
	Price string `json:"price"`
	SKU int64 `json:"sku"`
	Vat string `json:"vat"`
	CurrencyCode string `json:"currency_code"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	OldPrice string `json:"old_price"`
}

type GetProductInfoListResponseStocksStock struct {
	Present int32 `json:"present"`
	Reserved int32 `json:"reserved"`
	SKU int64 `json:"sku"`
	Source string `json:"source"`
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"`
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"`
}

type PriceIndexesIndexSelfData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type V1GetRealizationReportPostingResponse struct {
	Header GetRealizationReportResponseV2Header `json:"header"`
	Rows []interface{} `json:"rows"`
}

type DeliveryMethodListRequestFilter struct {
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
}

type WarehouseDeliveryMethodListRequest struct {
	Offset int64 `json:"offset"`
	Filter DeliveryMethodListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
}

type V1FbpDraftDirectRegistrateResponseRegistrationErrorBundleError struct {
	Errors []interface{} `json:"errors"`
	SKU int64 `json:"sku"`
}

type V1FbpDraftPickUpProductValidateRequest struct {
	Skus []interface{} `json:"skus"`
	WarehouseID int64 `json:"warehouse_id"`
}

type GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
	Name string `json:"name"`
}

type CommentListResponseComment struct {
	IsOwner bool `json:"is_owner"`
	ParentCommentID string `json:"parent_comment_id"`
	PublishedAt string `json:"published_at"`
	Text string `json:"text"`
	ID string `json:"id"`
	IsOfficial bool `json:"is_official"`
}

type V1GetStrategyResponseResult struct {
	Competitors []interface{} `json:"competitors"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Type string `json:"type_"`
	UpdateType string `json:"update_type"`
}

type Postingv1GetCarriageAvailableListRequest struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	DepartureDate string `json:"departure_date"`
}

type GetStrategyIDsByItemIDsResponseProductInfo struct {
	StrategyID string `json:"strategy_id"`
	ProductID int64 `json:"product_id"`
}

type DetailsOthers struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type DetailsReturns struct {
	Total float64 `json:"total"`
	Items []interface{} `json:"items"`
}

type V1SellerActionsArchiveRequest struct {
	ActionID int64 `json:"action_id"`
}

type V1GetProductInfoSubscriptionRequest struct {
	Skus []interface{} `json:"skus"`
}

type Productv4Filter struct {
	OfferID interface{} `json:"offer_id"`
	ProductID interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
	Visibility Productv2GetProductListRequestFilterFilterVisibility `json:"visibility"`
}

type V1GetCompetitorsRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type V3ChatDetailsInfo struct {
	ChatStatus ChatInfoChatStatusEnum `json:"chat_status"`
	ChatType ChatInfoChatTypeEnum `json:"chat_type"`
	CreatedAt string `json:"created_at"`
	ChatID string `json:"chat_id"`
}

type V3ChatInfo struct {
	Chat V3ChatDetailsInfo `json:"chat"`
	FirstUnreadMessageID int64 `json:"first_unread_message_id"`
	LastMessageID int64 `json:"last_message_id"`
	UnreadCount int64 `json:"unread_count"`
}

type ReportReportInfoRequest struct {
	Code string `json:"code"`
}

type V1FbpDraftDirectProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type V1AddStrategyItemsResponse struct {
	Result V1AddStrategyItemsResponseResult `json:"result"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress struct {
	AddressTail string `json:"address_tail"`
	City string `json:"city"`
	Country string `json:"country"`
	District string `json:"district"`
	ProviderPvzCode string `json:"provider_pvz_code"`
	PvzCode int64 `json:"pvz_code"`
	Region string `json:"region"`
	ZipCode string `json:"zip_code"`
	Comment string `json:"comment"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type PostingV4PostingFbsListResponsePostingsTarifficationStep struct {
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
	MinCharge MoneyMoneyCurrentTariffMinCharge `json:"min_charge"`
	TariffCharge MoneyMoneyCurrentTariffCharge `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"`
}

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPointAddressCoordinates struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityEnum string

type V1SearchQueriesTextRequest struct {
	Limit string `json:"limit"`
	Offset string `json:"offset"`
	SortBy SearchQueriesTextRequestSortBy `json:"sort_by"`
	SortDir SearchQueriesTextRequestSortDir `json:"sort_dir"`
	Text string `json:"text"`
}

type V1SearchQueriesTopResponseSearchQuery struct {
	SellersCount float64 `json:"sellers_count"`
	AddToCart float64 `json:"add_to_cart"`
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
	ConversionToCart float64 `json:"conversion_to_cart"`
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
}

type ProductV1ProductVisibilitySetRequestItemPlacement struct {
	Placement ProductV1ProductVisibilitySetRequestItemPlacementPlacementEnum `json:"placement"`
	SKU int64 `json:"sku"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData struct {
	IsLegal bool `json:"is_legal"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
	City string `json:"city"`
	ClientDeliveryDateBegin string `json:"client_delivery_date_begin"`
	DeliveryDateBegin string `json:"delivery_date_begin"`
	DeliveryType string `json:"delivery_type"`
	IsPremium bool `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region string `json:"region"`
	TPLProvider string `json:"tpl_provider"`
	ClientDeliveryDateEnd string `json:"client_delivery_date_end"`
	DeliveryDateEnd string `json:"delivery_date_end"`
}

type V2ConditionalCancellationMoveV2Request struct {
	CancellationID int64 `json:"cancellation_id"`
	Comment string `json:"comment"`
}

type SellerSellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type V1SetProductStairwayDiscountByQuantityResponse struct {
	Accepted bool `json:"accepted"`
	Errors []interface{} `json:"errors"`
	Warnings []interface{} `json:"warnings"`
}

type V1CommentDeleteRequest struct {
	CommentID string `json:"comment_id"`
}

type V3ImportProductsRequestItem struct {
	DescriptionCategoryID int64 `json:"description_category_id"`
	GeoNames interface{} `json:"geo_names"`
	ServiceType V2ServiceType `json:"service_type"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	ColorImage string `json:"color_image"`
	PrimaryImage string `json:"primary_image"`
	Promotions []interface{} `json:"promotions"`
	Images360 []interface{} `json:"images360"`
	OldPrice string `json:"old_price"`
	NewDescriptionCategoryID int64 `json:"new_description_category_id"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	Height int32 `json:"height"`
	Images []interface{} `json:"images"`
	PDFList []interface{} `json:"pdf_list"`
	TypeID int64 `json:"type_id"`
	Vat string `json:"vat"`
	Weight int32 `json:"weight"`
	CurrencyCode string `json:"currency_code"`
	Depth int32 `json:"depth"`
	DimensionUnit string `json:"dimension_unit"`
	WeightUnit string `json:"weight_unit"`
	Width int32 `json:"width"`
	Attributes []interface{} `json:"attributes"`
	Barcode string `json:"barcode"`
}

type V3CustomerFbsLists struct {
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Address V3Address `json:"address"`
	CustomerEmail string `json:"customer_email"`
}

type V3FbsPosting struct {
	FinancialData V3PostingFinancialData `json:"financial_data"`
	Cancellation V3Cancellation `json:"cancellation"`
	OrderNumber string `json:"order_number"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	Tariffication V3FbsTariffication `json:"tariffication"`
	Customer V3CustomerFbsLists `json:"customer"`
	Optional V3FbsPostingDetailOptional `json:"optional"`
	OrderID int64 `json:"order_id"`
	Status string `json:"status"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	DeliveringDate string `json:"delivering_date"`
	IsPresortable bool `json:"is_presortable"`
	LegalInfo V2FboSinglePostingLegalInfo `json:"legal_info"`
	AvailableActions interface{} `json:"available_actions"`
	Barcodes V3Barcodes `json:"barcodes"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	IsExpress bool `json:"is_express"`
	Requirements V3FbsPostingRequirementsV3 `json:"requirements"`
	ShipmentDate string `json:"shipment_date"`
	Substatus string `json:"substatus"`
	AnalyticsData V3FbsPostingAnalyticsData `json:"analytics_data"`
	DestinationPlaceName string `json:"destination_place_name"`
	PostingNumber string `json:"posting_number"`
	TrackingNumber string `json:"tracking_number"`
	Addressee V3AddresseeFbsLists `json:"addressee"`
	DeliveryMethod V3DeliveryMethod `json:"delivery_method"`
	InProcessAt string `json:"in_process_at"`
	ParentPostingNumber string `json:"parent_posting_number"`
	Products []interface{} `json:"products"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission struct {
	Currency string `json:"currency"`
	Percent int64 `json:"percent"`
	Amount float64 `json:"amount"`
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProducts struct {
	Commission PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission `json:"commission"`
	OldPrice float64 `json:"old_price"`
	Price float64 `json:"price"`
	Quantity int64 `json:"quantity"`
	TotalDiscountValue float64 `json:"total_discount_value"`
	Actions []interface{} `json:"actions"`
	CustomerPrice MoneyPostingMoney `json:"customer_price"`
	Payout float64 `json:"payout"`
	ProductID int64 `json:"product_id"`
	TotalDiscountPercent float64 `json:"total_discount_percent"`
}

type V1AssemblyCarriageProductListResponseProduct struct {
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	PostingNumbers []interface{} `json:"posting_numbers"`
	ProductName string `json:"product_name"`
	Quantity int64 `json:"quantity"`
}

type AnalyticsProductQueriesRequestSortDir string

type V1AnalyticsProductQueriesRequest struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Skus []interface{} `json:"skus"`
	SortBy AnalyticsProductQueriesRequestSortBy `json:"sort_by"`
	SortDir AnalyticsProductQueriesRequestSortDir `json:"sort_dir"`
}

type V3Dimensions struct {
	Height string `json:"height"`
	Length string `json:"length"`
	Weight string `json:"weight"`
	Width string `json:"width"`
}

type V3PostingProductDetail struct {
	OfferID string `json:"offer_id"`
	Price string `json:"price"`
	CurrencyCode string `json:"currency_code"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
	MandatoryMark []interface{} `json:"mandatory_mark"`
	HasImei bool `json:"has_imei"`
	Dimensions V3Dimensions `json:"dimensions"`
	Name string `json:"name"`
}

type Productv1ProductInfoPicturesResponseResult struct {
	Pictures interface{} `json:"pictures"`
}

type Productv1ProductInfoPicturesResponse struct {
	Result Productv1ProductInfoPicturesResponseResult `json:"result"`
}

type V1FbpDraftDropOffPointListRequest struct {
	PageSize int32 `json:"page_size"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
	NextPageNumber int32 `json:"next_page_number"`
}

type PostingCancelFbsPostingRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"`
	CancelReasonMessage string `json:"cancel_reason_message"`
	PostingNumber string `json:"posting_number"`
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []interface{} `json:"related_posting_numbers"`
}

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type ProductInfoWarehouseStocksResponseStocks struct {
	Present int64 `json:"present"`
	ProductID int64 `json:"product_id"`
	Reserved int64 `json:"reserved"`
	SKU int64 `json:"sku"`
	UpdatedAt string `json:"updated_at"`
	WarehouseID int64 `json:"warehouse_id"`
	FreeStock int64 `json:"free_stock"`
	OfferID string `json:"offer_id"`
}

type V1CreatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyName string `json:"strategy_name"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairwayStep struct {
	Discount int64 `json:"discount"`
	Quantity int64 `json:"quantity"`
	Step int64 `json:"step"`
}

type V1SellerActionsProductsAddRequestProduct struct {
	Currency ProductCurrencyEnum `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	SKU int64 `json:"sku"`
}

type GetProductAttributesResponseImage struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
	Default bool `json:"default"`
}

type V5FbsPostingProductExemplarValidateV5Response struct {
	Products []interface{} `json:"products"`
}

type V1ProductGetRelatedSKUResponseItem struct {
	DeletedAt string `json:"deleted_at"`
	DeliverySchema string `json:"delivery_schema"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
	Availability string `json:"availability"`
}

type V1FbpDraftPickUpProductValidateResponseRejectedItem struct {
	Volume float64 `json:"volume"`
	Barcode string `json:"barcode"`
	IconName string `json:"icon_name"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	Quantity int32 `json:"quantity"`
	RejectionReasons []interface{} `json:"rejection_reasons"`
	SKU int64 `json:"sku"`
}

type V1ReviewInfoRequest struct {
	ReviewID string `json:"review_id"`
}

type FinanceV1GetFinanceAccrualByDayResponseAccrualItemFeesItemFee struct {
	Fees []interface{} `json:"fees"`
	SKU int64 `json:"sku"`
}

type ProtobufAny struct {
	TypeUrl string `json:"typeUrl"`
	Value string `json:"value"`
}

type V1AssemblyFbsPostingListResponsePosting struct {
	AssemblyCode string `json:"assembly_code"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1SellerActionsProductsCandidatesResponseProductQuantTypeEnum string

type V3GetFbsPostingResponseV3 struct {
	Result V3FbsPostingDetail `json:"result"`
}

type V1FbpDraftPickupCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type GetProductRatingBySkuResponseRatingImproveAttribute struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type V1QuestionListResponse struct {
	Questions interface{} `json:"questions"`
	LastID string `json:"last_id"`
	HasNext bool `json:"has_next"`
}

type V1ProductUpdateDiscountRequest struct {
	Discount int32 `json:"discount"`
	ProductID int64 `json:"product_id"`
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

type V1SellerActionsCreateInstallmentRequest struct {
	DateStart string `json:"date_start"`
	Title string `json:"title"`
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

type ArrivalpassArrivalPassListRequest struct {
	Cursor string `json:"cursor"`
	Filter ArrivalPassListRequestFilter `json:"filter"`
	Limit int32 `json:"limit"`
}

type V1SellerActionsCreateMultiLevelDiscountRequestDiscountTypeEnum string

type V1GetDiscountTaskListResponseTask struct {
	RequestedQuantityMin int64 `json:"requested_quantity_min"`
	Patronymic string `json:"patronymic"`
	CustomerName string `json:"customer_name"`
	SKU int64 `json:"sku"`
	SellerComment string `json:"seller_comment"`
	RequestedPriceWithFee float64 `json:"requested_price_with_fee"`
	CreatedAt string `json:"created_at"`
	Status string `json:"status"`
	Discount float64 `json:"discount"`
	PrevTaskID int64 `json:"prev_task_id"`
	IsAutoModerated bool `json:"is_auto_moderated"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	ApprovedPriceFeePercent float64 `json:"approved_price_fee_percent"`
	ModeratedAt string `json:"moderated_at"`
	OfferID string `json:"offer_id"`
	FirstName string `json:"first_name"`
	ApprovedPriceWithFee float64 `json:"approved_price_with_fee"`
	EditedTill string `json:"edited_till"`
	BasePrice float64 `json:"base_price"`
	ApprovedQuantityMin int64 `json:"approved_quantity_min"`
	UserComment string `json:"user_comment"`
	IsDamaged bool `json:"is_damaged"`
	IsPurchased bool `json:"is_purchased"`
	LastName string `json:"last_name"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	RequestedPrice float64 `json:"requested_price"`
	OriginalPrice float64 `json:"original_price"`
	MinAutoPrice float64 `json:"min_auto_price"`
	ApprovedDiscountPercent float64 `json:"approved_discount_percent"`
	ID int64 `json:"id"`
	EndAt string `json:"end_at"`
	ApprovedPrice float64 `json:"approved_price"`
	DiscountPercent float64 `json:"discount_percent"`
	ApprovedDiscount float64 `json:"approved_discount"`
	Email string `json:"email"`
}

type GetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueId int64 `json:"dictionaryValueId"`
	Value string `json:"value"`
}

type V1SellerOzonLogisticsInfoResponse struct {
	AvailableSchemas []interface{} `json:"available_schemas"`
	OzonLogisticsEnabled bool `json:"ozon_logistics_enabled"`
}

type V1ProductUpdateAttributesRequest struct {
	Items interface{} `json:"items"`
}

type V1FbpOrderDropOffCancelResponse struct {
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
	Error V1OrderValidationError `json:"error"`
}

type V1FbpArchiveGetResponse struct {
	BundleID string `json:"bundle_id"`
	BusinessFlowTypeID int64 `json:"business_flow_type_id"`
	DeliveryDetails Fbpv1DeliveryDetails `json:"delivery_details"`
	RowVersion int64 `json:"row_version"`
	ActFileUuid string `json:"act_file_uuid"`
	HasAct bool `json:"has_act"`
	HasLabel bool `json:"has_label"`
	ID int64 `json:"id"`
	OrderDraftID int64 `json:"order_draft_id"`
	PackageUnitsCount int32 `json:"package_units_count"`
	ReceiveDate string `json:"receive_date"`
	SupplyID string `json:"supply_id"`
	BundleSKUSummary V1ArchiveSkuSummary `json:"bundle_sku_summary"`
	CreatedDate string `json:"created_date"`
	DeclineReason V1ArchiveDeclineReason `json:"decline_reason"`
	OrderNumber string `json:"order_number"`
	Status V1ArchiveStatusEnum `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1SellerActionsProductsListResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V1DescriptionCategoryTipsRequest struct {
	TypeID []interface{} `json:"type_id"`
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
}

type V3GetFbsPostingListResponseV3 struct {
	Result V3GetFbsPostingListResponseV3Result `json:"result"`
}

type V1SellerActionsCreateMultiLevelDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type DeliveryMethodListResponseDeliveryMethod struct {
	CompanyID int64 `json:"company_id"`
	CreatedAt string `json:"created_at"`
	ID int64 `json:"id"`
	Name string `json:"name"`
	ProviderID int64 `json:"provider_id"`
	SlaCutIn int64 `json:"sla_cut_in"`
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	Cutoff string `json:"cutoff"`
	TemplateID int64 `json:"template_id"`
	UpdatedAt string `json:"updated_at"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod struct {
	Name string `json:"name"`
	TPLProvider string `json:"tpl_provider"`
	TPLProviderID int64 `json:"tpl_provider_id"`
	Warehouse string `json:"warehouse"`
	WarehouseID int64 `json:"warehouse_id"`
	ID int64 `json:"id"`
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

type V1QuestionTopSkuRequest struct {
	Limit int64 `json:"limit"`
}

type Productv2DeleteProductsResponse struct {
	Status []interface{} `json:"status"`
}

type V3ImportProductsRequestComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type V1CommentCreateResponse struct {
	CommentID string `json:"comment_id"`
}

type V1FbpDraftPickUpDeleteResponse struct {
	CancellationState V1CancellationState `json:"cancellation_state"`
	RowVersion int64 `json:"row_version"`
}

type SellerOzonLogisticsInfoResponseAvailableSchemasEnum string

type SellerApiProductV1Response struct {
	Result SellerApiProductV1ResponseResult `json:"result"`
}

type Polygonv1PolygonCreateResponse struct {
	PolygonID int64 `json:"polygon_id"`
}

type GetWarehouseFBSOperationStatusResponseStatusEnum string

type ReviewV2ReviewListV2ResponseReviewOrderStatusEnum string

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
	HasNext bool `json:"has_next"`
}

type V2FbsPostingProductCountryListResponseResult struct {
	Name string `json:"name"`
	CountryIsoCode string `json:"country_iso_code"`
}

type V1FbpDraftListResponse struct {
	HasNext bool `json:"has_next"`
	Items []interface{} `json:"items"`
	LastID int64 `json:"last_id"`
}

type V1AssemblyCarriageProductListResponse struct {
	Cursor string `json:"cursor"`
	Products []interface{} `json:"products"`
}

type V1AnalyticsProductQueriesDetailsRequestSortDir string

type WarehouseWorkingDaysEnum string

type V1GetRealizationReportByDayRequest struct {
	Day int32 `json:"day"`
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type V6FbsPostingProductExemplarCreateOrGetV6Response struct {
	Products []interface{} `json:"products"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
}

type ProductImportProductsBySKUResponse struct {
	Result ProductImportProductsBySKUResponseResult `json:"result"`
}

type V1GetTreeRequest struct {
	Language LanguageLanguage `json:"language"`
}

type FbsPostingTrackingNumberSetRequestTrackingNumber struct {
	PostingNumber string `json:"posting_number"`
	TrackingNumber string `json:"tracking_number"`
}

type V1AssemblyFbsProductListRequestFilter struct {
	CutoffTo string `json:"cutoff_to"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	CutoffFrom string `json:"cutoff_from"`
}

type WarehouseCarriageLabelTypeEnum string

type V1WarehouseFbsUpdateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceEndTimeLocal string `json:"acceptance_end_time_local"`
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"`
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type ProductProductInfoPicturesResponsePicture struct {
	IsColor bool `json:"is_color"`
	IsPrimary bool `json:"is_primary"`
	ProductID int64 `json:"product_id"`
	State string `json:"state"`
	URL string `json:"url"`
	Is360 bool `json:"is_360"`
}

type GetWarehouseFBSOperationStatusResponseResult struct {
	EntityID int64 `json:"entity_id"`
}

type DetailsPayment struct {
	CurrencyCode string `json:"currency_code"`
	Payment float64 `json:"payment"`
}

type V3FinanceCashFlowStatementListResponsePeriod struct {
	ID int64 `json:"id"`
	Begin string `json:"begin"`
	End string `json:"end"`
}

type DetailsReturnDetails struct {
	Total float64 `json:"total"`
	Amount float64 `json:"amount"`
	ReturnServices DetailsReturns `json:"return_services"`
}

type FinanceCashFlowStatementListResponseDetails struct {
	Others DetailsOthers `json:"others"`
	BeginBalanceAmount float64 `json:"begin_balance_amount"`
	Payments DetailsPayment `json:"payments"`
	Period V3FinanceCashFlowStatementListResponsePeriod `json:"period"`
	Services DetailsService `json:"services"`
	EndBalanceAmount float64 `json:"end_balance_amount"`
	Delivery DetailsDeliveryDetails `json:"delivery"`
	InvoiceTransfer float64 `json:"invoice_transfer"`
	Loan float64 `json:"loan"`
	Return DetailsReturnDetails `json:"return"`
	RFBS DetailsRfbsDetails `json:"rfbs"`
}

type ReviewV2ReviewListV2ResponseReview struct {
	CommentsAmount int32 `json:"comments_amount"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	OrderStatus ReviewV2ReviewListV2ResponseReviewOrderStatusEnum `json:"order_status"`
	Status ReviewV2ReviewListV2ResponseReviewStatusEnum `json:"status"`
	VideosAmount int32 `json:"videos_amount"`
	ID string `json:"id"`
	PhotosAmount int32 `json:"photos_amount"`
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	SKU int64 `json:"sku"`
	Text string `json:"text"`
}

type V1QuestionAnswerDeleteRequest struct {
	AnswerID string `json:"answer_id"`
	SKU int64 `json:"sku"`
}

type V1SellerActionsChangeActivityRequest struct {
	ActionID int64 `json:"action_id"`
	IsTurnOn bool `json:"is_turn_on"`
}

type V1FbpDraftDropOffDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1QuestionAnswerCreateDefault struct {
	Code int32 `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

type V1SearchQueriesTopRequest struct {
	Limit string `json:"limit"`
	Offset string `json:"offset"`
}

type V1SellerActionsVoucherGetRequest struct {
	ActionID int64 `json:"action_id"`
}

type ReviewListResponseReview struct {
	CommentsAmount int32 `json:"comments_amount"`
	ID string `json:"id"`
	PhotosAmount int32 `json:"photos_amount"`
	PublishedAt string `json:"published_at"`
	Rating int32 `json:"rating"`
	SKU int64 `json:"sku"`
	Status string `json:"status"`
	Text string `json:"text"`
	IsRatingParticipant bool `json:"is_rating_participant"`
	OrderStatus string `json:"order_status"`
	VideosAmount int32 `json:"videos_amount"`
}

type AssemblyFbsProductListResponseProducts struct {
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	Postings []interface{} `json:"postings"`
	ProductName string `json:"product_name"`
	Quantity int32 `json:"quantity"`
	SKU int64 `json:"sku"`
}

type V1AnalyticsProductQueriesDetailsRequest struct {
	SortBy V1AnalyticsProductQueriesDetailsRequestSortBy `json:"sort_by"`
	SortDir V1AnalyticsProductQueriesDetailsRequestSortDir `json:"sort_dir"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	LimitBySKU int32 `json:"limit_by_sku"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Skus []interface{} `json:"skus"`
}

type ReviewInfoResponsePhoto struct {
	Height int32 `json:"height"`
	URL string `json:"url"`
	Width int32 `json:"width"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplarMark struct {
	ErrorCodes []interface{} `json:"error_codes"`
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
	CheckStatus string `json:"check_status"`
}

type ProductGetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type V1ProductInfoWarehouseStocksRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ProductUpdateOfferIdResponse struct {
	Errors interface{} `json:"errors"`
}

type PriceIndexesIndexOzonData struct {
	MinPrice float64 `json:"min_price"`
	MinPriceCurrency string `json:"min_price_currency"`
	PriceIndexValue float64 `json:"price_index_value"`
}

type GetProductInfoPricesResponseItemPriceIndexes struct {
	ColorIndex string `json:"color_index"`
	ExternalIndexData PriceIndexesIndexExternalData `json:"external_index_data"`
	OzonIndexData PriceIndexesIndexOzonData `json:"ozon_index_data"`
	SelfMarketplacesIndexData PriceIndexesIndexSelfData `json:"self_marketplaces_index_data"`
}

type ProductGetProductInfoPricesV5ResponseItem struct {
	MarketingActions ItemMarketing `json:"marketing_actions"`
	OfferID string `json:"offer_id"`
	Price ItemPricev5 `json:"price"`
	PriceIndexes GetProductInfoPricesResponseItemPriceIndexes `json:"price_indexes"`
	ProductID int64 `json:"product_id"`
	VolumeWeight float64 `json:"volume_weight"`
	Acquiring float64 `json:"acquiring"`
	Commissions ItemCommissionsv5 `json:"commissions"`
}

type V1AssemblyCarriagePostingListRequestFilter struct {
	DeliveryMethodID int64 `json:"delivery_method_id"`
	CarriageID int64 `json:"carriage_id"`
	CutoffFrom string `json:"cutoff_from"`
	CutoffTo string `json:"cutoff_to"`
}

type V1AssemblyCarriagePostingListRequest struct {
	Cursor string `json:"cursor"`
	Filter V1AssemblyCarriagePostingListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
}

type MarketingAction struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Title string `json:"title"`
	Value int32 `json:"value"`
}

type CreateWarehouseFBSRequestWorkingDaysEnum string

type V1FbpDraftDirectGetTimeslotRequest struct {
	IntervalEnd string `json:"interval_end"`
	IntervalStart string `json:"interval_start"`
	WarehouseID int64 `json:"warehouse_id"`
	BundleID string `json:"bundle_id"`
}

type V1FbpDraftPickUpProductValidateResponse struct {
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
}

type ProductV1ProductVisibilitySetResponseItemsSelectPermissionEnum string

type ProductV1ProductVisibilitySetResponseItems struct {
	SelectPermission ProductV1ProductVisibilitySetResponseItemsSelectPermissionEnum `json:"select_permission"`
	SellerItemPlacement ProductV1ProductVisibilitySetResponseItemsSellerItemPlacementEnum `json:"seller_item_placement"`
	SellerItemPlacementList []interface{} `json:"seller_item_placement_list"`
	ShowcasesVisibility ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityEnum `json:"showcases_visibility"`
	ShowcasesVisibilityList []interface{} `json:"showcases_visibility_list"`
	SKU int64 `json:"sku"`
	Warnings []interface{} `json:"warnings"`
}

type V1GetDiscountTaskListResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetStrategyListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type ReturnsRfbsGetV2ResponseAvailableAction struct {
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type V1UpdatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyID string `json:"strategy_id"`
	StrategyName string `json:"strategy_name"`
}

type V1AssemblyFbsPostingListResponse struct {
	Cursor string `json:"cursor"`
	Cutoff string `json:"cutoff"`
	Postings []interface{} `json:"postings"`
}

type Productv5GetProductInfoPricesV5Request struct {
	Cursor string `json:"cursor"`
	Filter Productv5Filter `json:"filter"`
	Limit int32 `json:"limit"`
}

type V1QuestionInfoRequest struct {
	QuestionID string `json:"question_id"`
}

type PostingV1PostingFbpListRequestFilter struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
	Since string `json:"since"`
	Statuses []interface{} `json:"statuses"`
	To string `json:"to"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
}

type PostingV1PostingFbpListRequest struct {
	Cursor string `json:"cursor"`
	Filter PostingV1PostingFbpListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir PostingV1PostingFbpListRequestSortDirEnum `json:"sort_dir"`
}

type V1FbpEditTimeslotResponse struct {
	ErrorReasons []interface{} `json:"error_reasons"`
	RowVersion int64 `json:"row_version"`
}

type V4GetProductInfoStocksRequest struct {
	Cursor string `json:"cursor"`
	Filter V4GetProductInfoStocksRequestFilter `json:"filter"`
	Limit int32 `json:"limit"`
}

type V1QuestionAnswerListResponse struct {
	Answers interface{} `json:"answers"`
	LastID string `json:"last_id"`
}

type ReportReport struct {
	CreatedAt string `json:"created_at"`
	Error string `json:"error"`
	ExpiresAt string `json:"expires_at"`
	File string `json:"file"`
	Params map[string]interface{} `json:"params"`
	ReportType string `json:"report_type"`
	Status string `json:"status"`
	Code string `json:"code"`
}

type FbpCreateActResponseCreateActErrorReason string

type V1FbpDraftDirectGetTimeslotResponseEmptyTimeslotsReason string

type FinanceCashFlowStatementListResponseCashFlow struct {
	ReturnsAmount float64 `json:"returns_amount"`
	CommissionAmount float64 `json:"commission_amount"`
	ServicesAmount float64 `json:"services_amount"`
	ItemDeliveryAndReturnAmount float64 `json:"item_delivery_and_return_amount"`
	CurrencyCode string `json:"currency_code"`
	Period V3FinanceCashFlowStatementListResponsePeriod `json:"period"`
	OrdersAmount float64 `json:"orders_amount"`
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplar struct {
	GTD string `json:"gtd"`
	Marks []interface{} `json:"marks"`
	Rnpt string `json:"rnpt"`
}

type GetWarehouseFBSOperationStatusResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1GetWarehouseFBSOperationStatusResponse struct {
	Result GetWarehouseFBSOperationStatusResponseResult `json:"result"`
	Status GetWarehouseFBSOperationStatusResponseStatusEnum `json:"status"`
	Type GetWarehouseFBSOperationStatusResponseTypeEnum `json:"type_"`
	Error GetWarehouseFBSOperationStatusResponseError `json:"error"`
}

type V1GetFinanceBalanceV1ResponseServicesMoney struct {
	CurrencyCode string `json:"currency_code"`
	Value float64 `json:"value"`
}

type GetFinanceBalanceV1ResponseService struct {
	Amount V1GetFinanceBalanceV1ResponseServicesMoney `json:"amount"`
	Name string `json:"name"`
}

type ReportReportListRequest struct {
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	ReportType ReportListRequestReportType `json:"report_type"`
}

type V1ApproveDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V1FbpArchiveGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V2ReturnsRfbsGetResponse struct {
	Returns ReturnsRfbsGetResponseReturns `json:"returns"`
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

type V1DescriptionCategoryTipsResponse struct {
	Result []interface{} `json:"result"`
}

type ReportCreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

type V2ReturnsRfbsListResponse struct {
	Returns ReturnsRfbsListResponseReturns `json:"returns"`
}

type V1QuestionAnswerListRequest struct {
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
	LastID interface{} `json:"last_id"`
}

type V1FbpDraftDropOffProductValidateResponse struct {
	BundleID string `json:"bundle_id"`
	RejectedItems []interface{} `json:"rejected_items"`
	ApprovedItems []interface{} `json:"approved_items"`
	BundleGenerated bool `json:"bundle_generated"`
}

type V1FbpDraftPickupCreateRequestDeliveryDetails struct {
	SenderName string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
	Address string `json:"address"`
	Comment string `json:"comment"`
	Date string `json:"date"`
}

type V1FbpDraftPickupCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails V1FbpDraftPickupCreateRequestDeliveryDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductGetProductAttributesV4ResponseAttribute struct {
	ID int64 `json:"id"`
	ComplexID int64 `json:"complex_id"`
	Values []interface{} `json:"values"`
}

type Productv5GetProductInfoPricesV5Response struct {
	Cursor string `json:"cursor"`
	Items interface{} `json:"items"`
	Total int32 `json:"total"`
}

type Productsv1GetProductInfoStocksByWarehouseFbsResponse struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsCreateMultiLevelDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	DiscountType V1SellerActionsCreateMultiLevelDiscountRequestDiscountTypeEnum `json:"discount_type"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplarMark struct {
	Mark string `json:"mark"`
	MarkType string `json:"mark_type"`
	Valid bool `json:"valid"`
	Errors []interface{} `json:"errors"`
}

type V1SellerActionsProductsListRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1PostingFbsSplitRequest struct {
	PostingNumber string `json:"posting_number"`
	Postings []interface{} `json:"postings"`
}

type V1ReviewCountResponse struct {
	Processed int32 `json:"processed"`
	Total int32 `json:"total"`
	Unprocessed int32 `json:"unprocessed"`
}

type Polygonv1PolygonCreateRequest struct {
	Coordinates string `json:"coordinates"`
}

type V3FbsPostingProductExemplarInfoV3 struct {
	ExemplarID int64 `json:"exemplar_id"`
	MandatoryMark string `json:"mandatory_mark"`
	GTD string `json:"gtd"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	Rnpt string `json:"rnpt"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Imei []interface{} `json:"imei"`
}

type V1GetStrategyResponse struct {
	Result V1GetStrategyResponseResult `json:"result"`
}

type FbsPostingShipV4RequestPackageProduct struct {
	ProductID int64 `json:"product_id"`
	Quantity int32 `json:"quantity"`
}

type Productv3GetProductAttributesV3Response struct {
	Total string `json:"total"`
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
}

type V3AdditionalDataItem struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type V1SellerActionsListRequest struct {
	ActionIds []interface{} `json:"action_ids"`
	ActionType []interface{} `json:"action_type"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Search string `json:"search"`
	Status []interface{} `json:"status"`
}

type OperationItem struct {
	Name string `json:"name"`
	SKU int64 `json:"sku"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1FbpDraftListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type ReviewInfoResponseVideo struct {
	Height int64 `json:"height"`
	PreviewURL string `json:"preview_url"`
	ShortVideoPreviewURL string `json:"short_video_preview_url"`
	URL string `json:"url"`
	Width int64 `json:"width"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []interface{} `json:"posting_number"`
}

type Productv2DeleteProductsRequest struct {
	Products []interface{} `json:"products"`
}

type V1FbpCheckConsignmentNoteStateResponse struct {
	ErrorMessage string `json:"error_message"`
	LabelURL string `json:"label_url"`
	State FbpCheckConsignmentNoteStateResponseStateType `json:"state"`
}

type V1GetAttributesResponseAttribute struct {
	ComplexIsCollection bool `json:"complex_is_collection"`
	CategoryDependent bool `json:"category_dependent"`
	Description string `json:"description"`
	DictionaryID int64 `json:"dictionary_id"`
	GroupID int64 `json:"group_id"`
	IsAspect bool `json:"is_aspect"`
	IsCollection bool `json:"is_collection"`
	IsRequired bool `json:"is_required"`
	Name string `json:"name"`
	GroupName string `json:"group_name"`
	ID int64 `json:"id"`
	Type string `json:"type_"`
	AttributeComplexID int64 `json:"attribute_complex_id"`
	MaxValueCount int64 `json:"max_value_count"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters V1SellerActionsUpdateMultiLevelDiscountRequestActionParameters `json:"action_parameters"`
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	HasNext bool `json:"has_next"`
	Cursor string `json:"cursor"`
	Errors []interface{} `json:"errors"`
}

type V1FbpCreateActResponse struct {
	Errors []interface{} `json:"errors"`
	FileUuid string `json:"file_uuid"`
	IsSuccess bool `json:"is_success"`
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

type ProductGetProductAttributesV3ResponseComplexAttribute struct {
	Attributes []interface{} `json:"attributes"`
}

type Productv4GetProductAttributesV4Request struct {
	Filter Productv4Filter `json:"filter"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

type SellerSellerAPIArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type V3FbsPostingExemplarProductV3 struct {
	Exemplars []interface{} `json:"exemplars"`
	SKU int64 `json:"sku"`
}

type WarehouseListV2ResponseWarehouse struct {
	Phone string `json:"phone"`
	PostingsLimit int32 `json:"postings_limit"`
	SlaCutIn int64 `json:"sla_cut_in"`
	Status string `json:"status"`
	CourierComment string `json:"courier_comment"`
	CourierPhones []interface{} `json:"courier_phones"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	IsExpress bool `json:"is_express"`
	IsWaybillEnabled bool `json:"is_waybill_enabled"`
	WarehouseType string `json:"warehouse_type"`
	CutInTime int64 `json:"cut_in_time"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	IsAutoAssembly bool `json:"is_auto_assembly"`
	IsComfort bool `json:"is_comfort"`
	IsKgt bool `json:"is_kgt"`
	Timetable WarehouseTimetable `json:"timetable"`
	UpdatedAt string `json:"updated_at"`
	WorkingDays []interface{} `json:"working_days"`
	AddressInfo WarehouseAddressInfo `json:"address_info"`
	CarriageLabelType WarehouseCarriageLabelTypeEnum `json:"carriage_label_type"`
	CreatedAt string `json:"created_at"`
	FirstMile WarehouseFirstMile `json:"first_mile"`
	IsRFBS bool `json:"is_rfbs"`
	Name string `json:"name"`
	WarehouseID int64 `json:"warehouse_id"`
	WithItemList bool `json:"with_item_list"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
}

type ProductV1ProductVisibilityInfoResponseItemShowcasesVisibilityEnum string

type ProductV1ProductVisibilityInfoResponseItem struct {
	ShowcasesVisibility ProductV1ProductVisibilityInfoResponseItemShowcasesVisibilityEnum `json:"showcases_visibility"`
	SKU int64 `json:"sku"`
}

type V1FbpDraftPickUpProductValidateRequestSkuItem struct {
	Count int32 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1FbpDraftPickupDlvEditRequest struct {
	PickupDetails V1FbpDraftPickupDlvEditRequestDeliveryDetails `json:"pickup_details"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type SellerApiProductIDsV1Request struct {
	ProductIds []interface{} `json:"product_ids"`
	ActionID float64 `json:"action_id"`
}

type ProductProductArchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type ProductImportProductsBySKURequest struct {
	Items []interface{} `json:"items"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"`
	ID int64 `json:"id"`
	To string `json:"to"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponse struct {
	Products []interface{} `json:"products"`
	Total int64 `json:"total"`
}

type V3FinanceCashFlowStatementListRequest struct {
	Date Financev3Period `json:"date"`
	Page int32 `json:"page"`
	WithDetails bool `json:"with_details"`
	PageSize int32 `json:"page_size"`
}

type V1AssemblyFbsProductListResponse struct {
	ProductsCount int32 `json:"products_count"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V2GetConditionalCancellationListV2Response struct {
	Counter int64 `json:"counter"`
	LastID int64 `json:"last_id"`
	Result []interface{} `json:"result"`
}

type SellerSellerAPIArrivalPassUpdateRequestArrivalPass struct {
	DriverName string `json:"driver_name"`
	DriverPhone string `json:"driver_phone"`
	ID int64 `json:"id"`
	VehicleLicensePlate string `json:"vehicle_license_plate"`
	VehicleModel string `json:"vehicle_model"`
	WithReturns bool `json:"with_returns"`
}

type FinanceV1GetFinanceAccrualTypesResponseAccrualType struct {
	Description string `json:"description"`
	ID int32 `json:"id"`
	Name string `json:"name"`
}

type PostingV1PostingFbpListResponsePostingsFinancialData struct {
	ClusterFrom string `json:"cluster_from"`
	ClusterTo string `json:"cluster_to"`
	DeliveryAmount float64 `json:"delivery_amount"`
	Products []interface{} `json:"products"`
}

type ListFBSRatingIndexPostingsV1ResponseError struct {
	ChargePercent float64 `json:"charge_percent"`
	DeliverySchema string `json:"delivery_schema"`
	ErrorAt string `json:"error_at"`
	HasGraceStatus bool `json:"has_grace_status"`
	Index float64 `json:"index"`
	PostingErrorType PostingErrorTypeEnum `json:"posting_error_type"`
	PostingNumber string `json:"posting_number"`
	ProductPrice float64 `json:"product_price"`
	ChargePrice float64 `json:"charge_price"`
	ChargePriceCurrencyCode string `json:"charge_price_currency_code"`
	ProductPriceCurrencyCode string `json:"product_price_currency_code"`
}

type V1FbpDraftDirectTimeslotEditRequest struct {
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
	TimeslotStart string `json:"timeslot_start"`
}

type V1CommentListResponse struct {
	Comments []interface{} `json:"comments"`
	Offset int32 `json:"offset"`
}

type Postingv3GetFbsPostingListRequestFilter struct {
	FbpFilter string `json:"fbpFilter"`
	Since string `json:"since"`
	Status string `json:"status"`
	WarehouseID []interface{} `json:"warehouse_id"`
	DeliveryMethodID []interface{} `json:"delivery_method_id"`
	OrderID int64 `json:"order_id"`
	ProviderID []interface{} `json:"provider_id"`
	To string `json:"to"`
	LastChangedStatusDate PostinglistV3status `json:"last_changed_status_date"`
}

type V1SearchQueriesTextResponseSearchQuery struct {
	AddToCart float64 `json:"add_to_cart"`
	AvgPrice float64 `json:"avg_price"`
	ClientCount float64 `json:"client_count"`
	ConversionToCart float64 `json:"conversion_to_cart"`
	ItemsViews float64 `json:"items_views"`
	Query string `json:"query"`
	SellersCount float64 `json:"sellers_count"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequestToUpdate struct {
	ActionPrice float64 `json:"action_price"`
	Currency string `json:"currency"`
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
}

type Productv2ProductsStocksRequest struct {
	Stocks []interface{} `json:"stocks"`
}

type ProductInfoWrongVolumeResponseWrongVolumeProduct struct {
	Weight int64 `json:"weight"`
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	Length int64 `json:"length"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
	SKU int64 `json:"sku"`
}

type V1ProductUpdateAttributesResponse struct {
	TaskID int64 `json:"task_id"`
}

type V1SetProductStairwayDiscountByQuantityRequestStairways struct {
	Stairway V1SetProductStairwayDiscountByQuantityRequestStairwaysStairway `json:"stairway"`
	Enabled bool `json:"enabled"`
	SKU int64 `json:"sku"`
}

type V1WarehouseWithInvalidProductsResponse struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type PostingV1PostingFbpListResponsePostings struct {
	InProcessAt string `json:"in_process_at"`
	OrderDate string `json:"order_date"`
	OrderID int64 `json:"order_id"`
	PostingNumber string `json:"posting_number"`
	ProviderID int64 `json:"provider_id"`
	Status string `json:"status"`
	FinancialData PostingV1PostingFbpListResponsePostingsFinancialData `json:"financial_data"`
	OrderNumber string `json:"order_number"`
	Products []interface{} `json:"products"`
}

type PostingV4PostingFbsUnfulfilledListRequest struct {
	With PostingV4PostingFbsUnfulfilledListRequestWith `json:"with"`
	Cursor string `json:"cursor"`
	Filter PostingV4PostingFbsUnfulfilledListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir PostingV4PostingFbsUnfulfilledListRequestSortDirEnum `json:"sort_dir"`
	Translit bool `json:"translit"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo struct {
	Inn string `json:"inn"`
	Kpp string `json:"kpp"`
	CompanyName string `json:"company_name"`
}

type V1QuestionCountResponse struct {
	All int64 `json:"all"`
	New int64 `json:"new"`
	Processed int64 `json:"processed"`
	Unprocessed int64 `json:"unprocessed"`
	Viewed int64 `json:"viewed"`
}

type RpcStatusV1PolygonBind struct {
	Code int32 `json:"code"`
	Details []interface{} `json:"details"`
	Message string `json:"message"`
}

type V1QuestionAnswerCreateRequest struct {
	Text string `json:"text"`
	QuestionID string `json:"question_id"`
	SKU int64 `json:"sku"`
}

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnID int64 `json:"return_id"`
	ReturnForBackWay int64 `json:"return_for_back_way"`
}

type V1FbpDraftDirectTplDlvCreateRequestDirectDetails struct {
	TimeslotStart string `json:"timeslot_start"`
	TrackingNumber string `json:"tracking_number"`
	TransportCompanyName string `json:"transport_company_name"`
}

type V1FbpDraftDirectTplDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails V1FbpDraftDirectTplDlvCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1SellerActionsUpdateInstallmentRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters V1SellerActionsUpdateInstallmentRequestActionParameters `json:"action_parameters"`
}

type V1AssemblyCarriagePostingListResponsePostingProduct struct {
	SKU int64 `json:"sku"`
	OfferID string `json:"offer_id"`
	PictureURL string `json:"picture_url"`
	ProductName string `json:"product_name"`
	Quantity int64 `json:"quantity"`
}

type ProductV1ProductVisibilityInfoResponse struct {
	Items []interface{} `json:"items"`
}

type V1FbpDraftDirectTplDlvEditResponse struct {
	Error V1OrderDraftValidationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type V1PostingFbsSplitResponsePosting struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearchDropOffPointTypeEnum string

type TimetableWorkingHours struct {
	TimeFrom string `json:"time_from"`
	TimeTo string `json:"time_to"`
}

type GetCompetitorsResponseCompetitorInfo struct {
	Name string `json:"name"`
	ID int64 `json:"id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep struct {
	MinCharge MoneyMoneyCurrentTariffMinCharge `json:"min_charge"`
	TariffCharge MoneyMoneyCurrentTariffCharge `json:"tariff_charge"`
	TariffDeadlineAt string `json:"tariff_deadline_at"`
	TariffRate float64 `json:"tariff_rate"`
	TariffType string `json:"tariff_type"`
}

type V1FbpDraftDirectSellerDlvCreateRequest struct {
	BundleID string `json:"bundle_id"`
	DeliveryDetails V1FbpDraftDirectSellerDlvCreateRequestDirectDetails `json:"delivery_details"`
	PackageUnitsCount int32 `json:"package_units_count"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ProductUpdateAttributesRequestAttribute struct {
	ComplexID int64 `json:"complex_id"`
	ID int64 `json:"id"`
	Values interface{} `json:"values"`
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

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"`
	Language CompensationReportLanguage `json:"language"`
}

type V1FbpOrderListRequest struct {
	Count int32 `json:"count"`
	LastID int64 `json:"last_id"`
}

type ActionsV1ActionsAutoAddProductsCandidatesRequest struct {
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	ActionID int64 `json:"action_id"`
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result Postingv3GetFbsPostingUnfulfilledListResponseResult `json:"result"`
}

type V1FbpOrderDropOffTimetableRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	ProvinceUuid string `json:"province_uuid"`
	WarehouseID int64 `json:"warehouse_id"`
}

type FinanceTransactionListV3ResponseOperation struct {
	AccrualsForSale float64 `json:"accruals_for_sale"`
	Amount float64 `json:"amount"`
	OperationDate string `json:"operation_date"`
	OperationType string `json:"operation_type"`
	OperationTypeName string `json:"operation_type_name"`
	Posting OperationPosting `json:"posting"`
	SaleCommission float64 `json:"sale_commission"`
	Type string `json:"type_"`
	DeliveryCharge float64 `json:"delivery_charge"`
	Items []interface{} `json:"items"`
	OperationID int64 `json:"operation_id"`
	ReturnDeliveryCharge float64 `json:"return_delivery_charge"`
	Services []interface{} `json:"services"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponse struct {
	BelowMinPrice []interface{} `json:"below_min_price"`
	ExtremelyLowPrice []interface{} `json:"extremely_low_price"`
	FailedPrice []interface{} `json:"failed_price"`
	Rejected []interface{} `json:"rejected"`
	UpdatedIds []interface{} `json:"updated_ids"`
}

type V1WarehouseFbsCreateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type V1FbpCreateConsignmentNoteRequest struct {
	SupplyID string `json:"supply_id"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseRejected struct {
	Code ActionsV1ActionsAutoAddProductsUpdateResponseRejectedCodeEnum `json:"code"`
	ProductID int64 `json:"product_id"`
	Reason string `json:"reason"`
}

type V1GetDiscountTaskListRequest struct {
	Limit int64 `json:"limit"`
	Status V1DiscountTaskStatus `json:"status"`
	Page int64 `json:"page"`
}

type V4GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"`
	Name string `json:"name"`
}

type V1ProductActionTimerStatusRequest struct {
	ProductIds interface{} `json:"product_ids"`
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"`
}

type V1ArchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V2DeliveryMethodListV2Request struct {
	Cursor string `json:"cursor"`
	Filter DeliveryMethodListV2RequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	SortDir DeliveryMethodListV2RequestSortDirEnum `json:"sort_dir"`
}

type GetImportProductsInfoResponseResultItem struct {
	Status string `json:"status"`
	Errors []interface{} `json:"errors"`
	OfferID string `json:"offer_id"`
	ProductID int64 `json:"product_id"`
}

type V1FbpOrderGetRequest struct {
	SupplyID string `json:"supply_id"`
}

type V1QuestionAnswerCreateResponse struct {
	AnswerID string `json:"answer_id"`
}

type ProductImportProductsPricesResponse struct {
	Result []interface{} `json:"result"`
}

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"`
}

type CommonCreateReportResponse struct {
	Result CreateReportResponseCode `json:"result"`
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []interface{} `json:"tracking_numbers"`
}

type V1ProductFbsSplit struct {
	ProductID int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
}

type V2ReturnsRfbsRejectRequest struct {
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
	RejectionReasonID int64 `json:"rejection_reason_id"`
}

type V1FbpDraftDirectRegistrateResponse struct {
	Error V1FbpDraftDirectRegistrateResponseRegistrationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type WarehouseListResponseWarehouse struct {
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	IsRFBS bool `json:"is_rfbs"`
	IsKarantin bool `json:"is_karantin"`
	IsKgt bool `json:"is_kgt"`
	MinPostingsLimit int32 `json:"min_postings_limit"`
	MinWorkingDays int64 `json:"min_working_days"`
	Status string `json:"status"`
	WarehouseID int64 `json:"warehouse_id"`
	IsAbleToSetPrice bool `json:"is_able_to_set_price"`
	IsPresorted bool `json:"is_presorted"`
	IsTimetableEditable bool `json:"is_timetable_editable"`
	WorkingDays interface{} `json:"working_days"`
	Name string `json:"name"`
	CanPrintActInAdvance bool `json:"can_print_act_in_advance"`
	FirstMileType WarehouseFirstMileType `json:"first_mile_type"`
	HasPostingsLimit bool `json:"has_postings_limit"`
	PostingsLimit int32 `json:"postings_limit"`
}

type V1GetRealizationReportPostingRequest struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type FbsPostingMoveStatusResponseMoveStatus struct {
	Error string `json:"error"`
	PostingNumber string `json:"posting_number"`
	Result bool `json:"result"`
}

type V4GetProductInfoStocksResponse struct {
	Cursor string `json:"cursor"`
	Items []interface{} `json:"items"`
	Total int32 `json:"total"`
}

type Productv1ProductImportPicturesRequest struct {
	ProductID int64 `json:"product_id"`
	ColorImage string `json:"color_image"`
	Images interface{} `json:"images"`
	Images360 interface{} `json:"images360"`
}

type AnalyticsSorting struct {
	Key string `json:"key"`
	Order SortingOrder `json:"order"`
}

type SellerApiActivateProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type V1SellerActionsProductsCandidatesResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V1ProductUpdateAttributesRequestValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"`
	Value string `json:"value"`
}

type V1FbpDraftDropOffPointListResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
}

type V2ChatReadResponse struct {
	UnreadCount int64 `json:"unread_count"`
}

type V1PostingFbsSplitRequestPosting struct {
	Products []interface{} `json:"products"`
}

type V1ProductPricesDetailsResponsePrice struct {
	Price MoneyMoney `json:"price"`
	PriceIndexes []interface{} `json:"price_indexes"`
	SKU int64 `json:"sku"`
	CustomerPrice MoneyMoneyCustomerPrice `json:"customer_price"`
	DiscountPercent float64 `json:"discount_percent"`
	OfferID string `json:"offer_id"`
}

type V1OrderValidationErrorErrorType string

type V1SellerActionsProductsCandidatesResponseProduct struct {
	ActionPrice float64 `json:"action_price"`
	BasePrice float64 `json:"base_price"`
	Currency string `json:"currency"`
	DiscountPercent float64 `json:"discount_percent"`
	IsActive bool `json:"is_active"`
	Name string `json:"name"`
	OfferID string `json:"offer_id"`
	QuantSize int64 `json:"quant_size"`
	MinSellerPrice float64 `json:"min_seller_price"`
	Price float64 `json:"price"`
	ProductID int64 `json:"product_id"`
	QuantType V1SellerActionsProductsCandidatesResponseProductQuantTypeEnum `json:"quant_type"`
	SKU []interface{} `json:"sku"`
}

type ProductV1ProductVisibilitySetResponseItemsShowcasesVisibilityListEnum string

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type FinanceCashFlowStatementListResponseDeliveryService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V1FbpDraftDirectTimeslotEditResponseReserveFailureType string

type V1SetPostingsRequest struct {
	CarriageID int64 `json:"carriage_id"`
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type FbpDraftDropOffProvinceListResponseProvince struct {
	Name string `json:"name"`
	PointsCount int32 `json:"points_count"`
	ProvinceUuid string `json:"province_uuid"`
}

type WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPoint struct {
	Address string `json:"address"`
	AddressCoordinates WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPointAddressCoordinates `json:"address_coordinates"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type DeliveryMethodListV2ResponseDeliveryMethod struct {
	UpdatedAt string `json:"updated_at"`
	ID int64 `json:"id"`
	ProviderID int64 `json:"provider_id"`
	TemplateID int64 `json:"template_id"`
	TPLDropoffPoint WarehouseV2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPoint `json:"tpl_dropoff_point"`
	WarehouseID int64 `json:"warehouse_id"`
	CreatedAt string `json:"created_at"`
	Cutoff string `json:"cutoff"`
	IsExpress bool `json:"is_express"`
	Name string `json:"name"`
	SlaCutIn int64 `json:"sla_cut_in"`
	Status string `json:"status"`
	TPLIntegrationType string `json:"tpl_integration_type"`
}

type GetReturnsListResponseExemplar struct {
	ID int64 `json:"id"`
}

type V1ProductActionTimerStatusResponseStatuses struct {
	ExpiredAt string `json:"expired_at"`
	MinPriceForAutoActionsEnabled bool `json:"min_price_for_auto_actions_enabled"`
	ProductID int64 `json:"product_id"`
}

type V1FbpDraftDirectCreateResponse struct {
	SupplyID string `json:"supply_id"`
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
}

type V1FbpDraftDropOffPointTimetableResponse struct {
	Calendar []interface{} `json:"calendar"`
}

type DeleteProductsRequestProduct struct {
	OfferID string `json:"offer_id"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseBelowMinPrice struct {
	Key int64 `json:"key"`
	Value float64 `json:"value"`
}

type V1FbpDraftDirectTplDlvCreateResponse struct {
	DraftID int64 `json:"draft_id"`
	RowVersion int64 `json:"row_version"`
	SupplyID string `json:"supply_id"`
}

type OperationService struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type V5FbsPostingProductExemplarStatusV5Response struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
}

type WarehouseDeliveryMethodListResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V1ProductInfoWrongVolumeResponse struct {
	Cursor string `json:"cursor"`
	Products interface{} `json:"products"`
}

type V3FinanceCashFlowStatementListResponse struct {
	Result V3FinanceCashFlowStatementListResponseResult `json:"result"`
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplar struct {
	ExemplarID int64 `json:"exemplar_id"`
	GTD string `json:"gtd"`
	GTDCheckStatus string `json:"gtd_check_status"`
	IsGTDAbsent bool `json:"is_gtd_absent"`
	IsRnptAbsent bool `json:"is_rnpt_absent"`
	Rnpt string `json:"rnpt"`
	RnptCheckStatus string `json:"rnpt_check_status"`
	RnptErrorCodes []interface{} `json:"rnpt_error_codes"`
	GTDErrorCodes []interface{} `json:"gtd_error_codes"`
	Marks []interface{} `json:"marks"`
}

type V1FbpArchiveListRequest struct {
	Count string `json:"count"`
	LastID string `json:"last_id"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponseFailedPrice struct {
	Value float64 `json:"value"`
	Key int64 `json:"key"`
}

type V1WarehouseInvalidProductsGetResponse struct {
	HasNext bool `json:"has_next"`
	LastID int64 `json:"last_id"`
	ValidationResults []interface{} `json:"validation_results"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V2ReturnsRfbsCompensateRequest struct {
	CompensationAmount string `json:"compensation_amount"`
	ReturnID int64 `json:"return_id"`
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber string `json:"posting_number"`
	ProductID int64 `json:"product_id"`
	CountryIsoCode string `json:"country_iso_code"`
}

type V1QuestionListResponseQuestions struct {
	ID string `json:"id"`
	ProductURL string `json:"product_url"`
	QuestionLink string `json:"question_link"`
	SKU int64 `json:"sku"`
	Status interface{} `json:"status"`
	Text string `json:"text"`
	AnswersCount int64 `json:"answers_count"`
	AuthorName string `json:"author_name"`
	PublishedAt interface{} `json:"published_at"`
}

type SellerApiProductV1ResponseDeactivate struct {
	Result SellerApiProductV1ResponseResultDeactivate `json:"result"`
}

type V1FbpOrderDirectCancelResponse struct {
	Error V1OrderValidationError `json:"error"`
	IsError bool `json:"is_error"`
	RowVersion int64 `json:"row_version"`
}

type Fbsv4FbsPostingShipV4Response struct {
	AdditionalData interface{} `json:"additional_data"`
	Result interface{} `json:"result"`
}

type V1FbpDraftDirectDeleteRequest struct {
	SupplyID string `json:"supply_id"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer struct {
	Address PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress `json:"address"`
	CustomerEmail string `json:"customer_email"`
	CustomerID int64 `json:"customer_id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type PostingV4PostingFbsUnfulfilledListResponsePostings struct {
	InProcessAt string `json:"in_process_at"`
	PrrOption string `json:"prr_option"`
	Substatus string `json:"substatus"`
	Addressee PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee `json:"addressee"`
	AnalyticsData PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData `json:"analytics_data"`
	ExternalOrder PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder `json:"external_order"`
	IsMultibox bool `json:"is_multibox"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	QuantumID int64 `json:"quantum_id"`
	Customer PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer `json:"customer"`
	DeliveryMethod PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod `json:"delivery_method"`
	RequireBlrTraceableAttrs bool `json:"require_blr_traceable_attrs"`
	ShipmentDate string `json:"shipment_date"`
	Status string `json:"status"`
	AvailableActions []interface{} `json:"available_actions"`
	Barcodes PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes `json:"barcodes"`
	DeliverySchema string `json:"delivery_schema"`
	FinancialData PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData `json:"financial_data"`
	Optional PostingV4PostingFbsUnfulfilledListResponsePostingsOptional `json:"optional"`
	VolumeWeight float64 `json:"volume_weight"`
	PickupCodeVerifiedAt string `json:"pickup_code_verified_at"`
	ShipmentDateWithoutDelay string `json:"shipment_date_without_delay"`
	Tariffication PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication `json:"tariffication"`
	DestinationPlaceName string `json:"destination_place_name"`
	IsClickAndCollect bool `json:"is_click_and_collect"`
	IsPresortable bool `json:"is_presortable"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	TarifficationSteps []interface{} `json:"tariffication_steps"`
	TrackingNumber string `json:"tracking_number"`
	LegalInfo PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo `json:"legal_info"`
	ParentPostingNumber string `json:"parent_posting_number"`
	TPLIntegrationType string `json:"tpl_integration_type"`
	DestinationPlaceID int64 `json:"destination_place_id"`
	IsExpress bool `json:"is_express"`
	OrderID int64 `json:"order_id"`
	OrderNumber string `json:"order_number"`
	Requirements PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements `json:"requirements"`
	Cancellation PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation `json:"cancellation"`
	DeliveringDate string `json:"delivering_date"`
}

type V1CreateStockByWarehouseReportRequest struct {
	Language ReportLanguage `json:"language"`
	WarehouseId []interface{} `json:"warehouseId"`
}

type V1SellerActionsListResponse struct {
	Actions []interface{} `json:"actions"`
	Total int64 `json:"total"`
}

type Postingv3GetFbsPostingListRequest struct {
	Dir string `json:"dir"`
	Filter Postingv3GetFbsPostingListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With Postingv3FbsPostingWithParams `json:"with"`
}

type V1SellerActionsUpdateVoucherRequestActionParameters struct {
	Title string `json:"title"`
	UserIds []interface{} `json:"user_ids"`
	Budget int64 `json:"budget"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountValue float64 `json:"discount_value"`
}

type V1SellerActionsUpdateVoucherRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters V1SellerActionsUpdateVoucherRequestActionParameters `json:"action_parameters"`
}

type Productv2ProductsStocksResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type V1GetAttributesResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetTreeResponseItem struct {
	TypeID int64 `json:"type_id"`
	TypeName string `json:"type_name"`
	DescriptionCategoryID int64 `json:"description_category_id"`
	CategoryName string `json:"category_name"`
	Children []interface{} `json:"children"`
	Disabled bool `json:"disabled"`
}

type V1FbpAvailableTimeslotListResponseEmptyTimeslotsReason string

type V1FbpCreateLabelResponse struct {
	Code string `json:"code"`
}

type V1AssemblyFbsProductListRequest struct {
	SortDir V1AssemblyFbsProductListRequestSortDirEnum `json:"sort_dir"`
	Filter V1AssemblyFbsProductListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type SellerApiProductV1ResponseProductDeactivate struct {
	ProductID float64 `json:"product_id"`
	Reason string `json:"reason"`
}

type V2ProductInfoPicturesResponse struct {
	Items []interface{} `json:"items"`
}

type V5FbsPostingProductExemplarStatusV5Request struct {
	PostingNumber string `json:"posting_number"`
}

type PostingCancelReason struct {
	ID int64 `json:"id"`
	IsAvailableForCancellation bool `json:"is_available_for_cancellation"`
	Title string `json:"title"`
	TypeID string `json:"type_id"`
}

type V1ReturnsRfbsActionSetRequest struct {
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
	CompensationAmount float64 `json:"compensation_amount"`
	ID int32 `json:"id"`
	RejectionReasonID int32 `json:"rejection_reason_id"`
	ReturnForBackWay float64 `json:"return_for_back_way"`
}

type ProductV1ProductVisibilitySetResponseItemsErrors struct {
	Code string `json:"code"`
	SKU int64 `json:"sku"`
}

type V1UpdateWarehouseFBSFirstMileResponse struct {
	OperationID string `json:"operation_id"`
}

type V1GenerateBarcodeResult struct {
	ProductID int64 `json:"product_id"`
	Code string `json:"code"`
	Error string `json:"error"`
	Barcode string `json:"barcode"`
}

type V3ChatListResponse struct {
	TotalUnreadCount int64 `json:"total_unread_count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Chats interface{} `json:"chats"`
}

type V1SellerActionsVoucherGetResponse struct {
	File string `json:"file"`
}

type V1ProductInfoWarehouseStocksResponse struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	Stocks []interface{} `json:"stocks"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseRequest struct {
	Search ListDropOffPointsForUpdateFBSWarehouseRequestSearch `json:"search"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1FbpDraftDirectProductValidateRequestSkuItem struct {
	Count int64 `json:"count"`
	SKU int64 `json:"sku"`
}

type V1FbpOrderDropOffDlvEditResponse struct {
	RowVersion int64 `json:"row_version"`
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"`
}

type GetProductRatingBySkuResponseProductRating struct {
	Rating float64 `json:"rating"`
	Groups interface{} `json:"groups"`
	SKU int64 `json:"sku"`
}

type V1SearchQueriesTopResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type GetCarriageAvailableListResponseResult struct {
	CarriagePostingsCount int32 `json:"carriage_postings_count"`
	CutoffAt string `json:"cutoff_at"`
	Errors interface{} `json:"errors"`
	MandatoryPackagedCount int32 `json:"mandatory_packaged_count"`
	WarehouseID int64 `json:"warehouse_id"`
	TPLProviderIconURL string `json:"tpl_provider_icon_url"`
	CarriageID int64 `json:"carriage_id"`
	DeliveryMethodName string `json:"delivery_method_name"`
	RecommendedTimeUtcOffsetInMinutes float64 `json:"recommended_time_utc_offset_in_minutes"`
	TPLProviderName string `json:"tpl_provider_name"`
	WarehouseCity string `json:"warehouse_city"`
	FirstMileType string `json:"first_mile_type"`
	WarehouseTimezone string `json:"warehouse_timezone"`
	CarriageStatus string `json:"carriage_status"`
	DeliveryMethodID int64 `json:"delivery_method_id"`
	HasEntrustedAcceptance bool `json:"has_entrusted_acceptance"`
	MandatoryPostingsCount int32 `json:"mandatory_postings_count"`
	RecommendedTimeLocal string `json:"recommended_time_local"`
	WarehouseName string `json:"warehouse_name"`
}

type ProductsPostings struct {
	PostingNumber string `json:"posting_number"`
	Quantity int32 `json:"quantity"`
}

type GetDiscountTaskListV2ResponseTask struct {
	CreatedAt string `json:"created_at"`
	MinAutoPrice float64 `json:"min_auto_price"`
	OriginalPrice float64 `json:"original_price"`
	Patronymic string `json:"patronymic"`
	ReductionFactor float64 `json:"reduction_factor"`
	RequestedPrice float64 `json:"requested_price"`
	RequestedQuantityMax int64 `json:"requested_quantity_max"`
	Status GetDiscountTaskListV2ResponseTaskDiscountTaskStatusEnum `json:"status"`
	ApprovedDiscount float64 `json:"approved_discount"`
	AutoModeratedInfo TaskAutoModeratedInfo `json:"auto_moderated_info"`
	Email string `json:"email"`
	EndAt string `json:"end_at"`
	EndAtDuration int64 `json:"end_at_duration"`
	ID int64 `json:"id"`
	IsAutoModerated bool `json:"is_auto_moderated"`
	SKU int64 `json:"sku"`
	ApprovedPrice float64 `json:"approved_price"`
	ApprovedQuantityMax int64 `json:"approved_quantity_max"`
	EditedTill string `json:"edited_till"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ModeratedAt string `json:"moderated_at"`
	EditedTillDuration int64 `json:"edited_till_duration"`
	Name string `json:"name"`
	RequestedDiscount float64 `json:"requested_discount"`
}

type V6FbsPostingProductExemplarSetV6Request struct {
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type ChatChatStartRequest struct {
	PostingNumber string `json:"posting_number"`
}

type ReportCreateCompanyProductsReportRequest struct {
	Language ReportLanguage `json:"language"`
	OfferID []interface{} `json:"offer_id"`
	Search string `json:"search"`
	SKU []interface{} `json:"sku"`
	Visibility ReportCreateCompanyProductsReportRequestVisibility `json:"visibility"`
}

type V3ImportProductsRequestAttribute struct {
	ComplexID int64 `json:"complex_id"`
	ID int64 `json:"id"`
	Values []interface{} `json:"values"`
}

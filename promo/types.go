package promo

type V1SellerActionsProductsListRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1GetDiscountTaskListRequest struct {
	Status interface{} `json:"status"`
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1SellerActionsVoucherGetResponse struct {
	File string `json:"file"`
}

type V1SellerActionsUpdateMultiLevelDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1SellerActionsUpdateDiscountRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type ActionsV1ActionsAutoAddProductsListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
}

type ActionsV1ActionsAutoAddProductsListResponse struct {
	Total int64 `json:"total"`
	Products []interface{} `json:"products"`
}

type ActionsV1ActionsAutoAddProductsUpdateResponse struct {
	FailedPrice []interface{} `json:"failed_price"`
	Rejected []interface{} `json:"rejected"`
	UpdatedIds []interface{} `json:"updated_ids"`
	BelowMinPrice []interface{} `json:"below_min_price"`
	ExtremelyLowPrice []interface{} `json:"extremely_low_price"`
}

type V1SellerActionsCreateInstallmentResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1SellerActionsUpdateDiscountWithConditionRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type SellerApiProductV1ResponseDeactivate struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsCreateDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	MinActionPercent float64 `json:"min_action_percent"`
	Title string `json:"title"`
}

type V1SellerActionsCreateMultiLevelDiscountRequest struct {
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountLevels []interface{} `json:"discount_levels"`
	DiscountType interface{} `json:"discount_type"`
	IsLegalEntitiesSegment bool `json:"is_legal_entities_segment"`
	Title string `json:"title"`
}

type SellerApiGetSellerProductV1Response struct {
	Result interface{} `json:"result"`
}

type SellerApiActivateProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type V1SellerActionsProductsAddRequest struct {
	ActionID int64 `json:"action_id"`
	Products []interface{} `json:"products"`
}

type V1SellerActionsCreateVoucherResponse struct {
	ActionID int64 `json:"action_id"`
}

type ActionsV1ActionsAutoAddProductsUpdateRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	ToUpdate []interface{} `json:"to_update"`
}

type V1DeclineDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V1GetDiscountTaskListResponse struct {
	Result []interface{} `json:"result"`
}

type V1SellerActionsCreateDiscountWithConditionRequest struct {
	MinOrderAmount float64 `json:"min_order_amount"`
	Title string `json:"title"`
	DateEnd string `json:"date_end"`
	DateStart string `json:"date_start"`
	DiscountType interface{} `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
}

type V1SellerActionsVoucherGetRequest struct {
	ActionID int64 `json:"action_id"`
}

type ActionsV1ActionsAutoAddProductsCandidatesResponse struct {
	Total int64 `json:"total"`
	Products []interface{} `json:"products"`
}

type V1SellerActionsArchiveRequest struct {
	ActionID int64 `json:"action_id"`
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

type V1SellerActionsCreateMultiLevelDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1SellerActionsUpdateInstallmentRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type ActionsV1ActionsAutoAddProductsDeleteResponse struct {
	ProductIds []interface{} `json:"product_ids"`
}

type SellerApiGetSellerProductV1Request struct {
	ActionID float64 `json:"action_id"`
	Limit float64 `json:"limit"`
	LastID float64 `json:"last_id"`
}

type SellerApiGetSellerActionsV1Response struct {
	Result []interface{} `json:"result"`
}

type ActionsV1ActionsAutoAddProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type V1ApproveDeclineDiscountTasksResponse struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsProductsCandidatesResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type V1SellerActionsProductsListResponse struct {
	Cursor int64 `json:"cursor"`
	HasNext bool `json:"has_next"`
	Products []interface{} `json:"products"`
}

type ActionsV1ActionsAutoAddProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	AutoAddDate string `json:"auto_add_date"`
	ProductIds []interface{} `json:"product_ids"`
}

type V1SellerActionsChangeActivityRequest struct {
	ActionID int64 `json:"action_id"`
	IsTurnOn bool `json:"is_turn_on"`
}

type V1SellerActionsListRequest struct {
	ActionIds []interface{} `json:"action_ids"`
	ActionType []interface{} `json:"action_type"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Search string `json:"search"`
	Status []interface{} `json:"status"`
}

type V1SellerActionsCreateDiscountWithConditionResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1SellerActionsUpdateVoucherRequest struct {
	ActionID int64 `json:"action_id"`
	ActionParameters interface{} `json:"action_parameters"`
}

type V1ApproveDiscountTasksRequest struct {
	Tasks []interface{} `json:"tasks"`
}

type V1SellerActionsProductsCandidatesRequest struct {
	ActionID int64 `json:"action_id"`
	Cursor int64 `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1SellerActionsCreateInstallmentRequest struct {
	DateStart string `json:"date_start"`
	Title string `json:"title"`
}

type V1SellerActionsProductsDeleteRequest struct {
	ActionID int64 `json:"action_id"`
	Skus []interface{} `json:"skus"`
}

type SellerApiProductIDsV1Request struct {
	ActionID float64 `json:"action_id"`
	ProductIds []interface{} `json:"product_ids"`
}

type SellerApiProductV1Response struct {
	Result interface{} `json:"result"`
}

type V1SellerActionsCreateDiscountResponse struct {
	ActionID int64 `json:"action_id"`
}

type V1SellerActionsListResponse struct {
	Actions []interface{} `json:"actions"`
	Total int64 `json:"total"`
}

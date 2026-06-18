package beta

type V1PostingFbsSplitResponse struct {
	ParentPosting interface{} `json:"parent_posting"`
	Postings []interface{} `json:"postings"`
}

type V2GetDiscountTaskListV2Request struct {
	LastID int64 `json:"last_id"`
	Limit int64 `json:"limit"`
	Status interface{} `json:"status"`
}

type V1SetProductStairwayDiscountByQuantityResponse struct {
	Accepted bool `json:"accepted"`
	Errors []interface{} `json:"errors"`
	Warnings []interface{} `json:"warnings"`
}

type FinanceV1GetFinanceAccrualPostingsResponse struct {
	PostingAccruals []interface{} `json:"posting_accruals"`
}

type ProductV1ProductVisibilitySetRequest struct {
	ItemPlacement []interface{} `json:"item_placement"`
}

type V1GetProductStairwayDiscountByQuantityResponse struct {
	Stairways []interface{} `json:"stairways"`
}

type V1SetProductStairwayDiscountByQuantityRequest struct {
	SuppressWarnings bool `json:"suppress_warnings"`
	Stairways []interface{} `json:"stairways"`
}

type V1DescriptionCategoryTipsResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetFinanceBalanceV1Response struct {
	Cashflows interface{} `json:"cashflows"`
	Total interface{} `json:"total"`
}

type FinanceV1GetFinanceAccrualPostingsRequest struct {
	PostingNumbers []interface{} `json:"posting_numbers"`
}

type ProductV1ProductVisibilityInfoRequest struct {
	Skus []interface{} `json:"skus"`
}

type V6FbsPostingProductExemplarCreateOrGetV6Request struct {
	PostingNumber string `json:"posting_number"`
}

type V1FbsPostingProductExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"`
}

type V5FbsPostingProductExemplarValidateV5Request struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V5FbsPostingProductExemplarStatusV5Request struct {
	PostingNumber string `json:"posting_number"`
}

type V5FbsPostingProductExemplarStatusV5Response struct {
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
	Status string `json:"status"`
}

type V1ProductInfoWarehouseStocksRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
	WarehouseID int64 `json:"warehouse_id"`
}

type ProductV1ProductVisibilitySetResponse struct {
	Items []interface{} `json:"items"`
	ItemsErrors []interface{} `json:"items_errors"`
}

type V6FbsPostingProductExemplarCreateOrGetV6Response struct {
	Products []interface{} `json:"products"`
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
}

type FinanceV1GetFinanceAccrualTypesResponse struct {
	AccrualTypes []interface{} `json:"accrual_types"`
}

type V5FbsPostingProductExemplarValidateV5Response struct {
	Products []interface{} `json:"products"`
}

type V1ProductInfoWarehouseStocksResponse struct {
	Stocks []interface{} `json:"stocks"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
}

type V2GetDiscountTaskListV2Response struct {
	Tasks []interface{} `json:"tasks"`
}

type FinanceV1GetFinanceAccrualByDayResponse struct {
	Accruals []interface{} `json:"accruals"`
	LastID string `json:"last_id"`
}

type V1GetProductStairwayDiscountByQuantityRequest struct {
	Skus []interface{} `json:"skus"`
}

type V1DescriptionCategoryTipsRequest struct {
	TypeID []interface{} `json:"type_id"`
}

type ProductV1ProductVisibilityInfoResponse struct {
	Items []interface{} `json:"items"`
}

type FinanceV1GetFinanceAccrualByDayRequest struct {
	Date string `json:"date"`
	LastID string `json:"last_id"`
}

type V6FbsPostingProductExemplarSetV6Request struct {
	MultiBoxQty int32 `json:"multi_box_qty"`
	PostingNumber string `json:"posting_number"`
	Products []interface{} `json:"products"`
}

type V1GetFinanceBalanceV1Request struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}

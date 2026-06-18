package returns

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnID int64 `json:"return_id"`
	ReturnForBackWay int64 `json:"return_for_back_way"`
}

type V2GetConditionalCancellationListV2Response struct {
	Counter int64 `json:"counter"`
	LastID int64 `json:"last_id"`
	Result []interface{} `json:"result"`
}

type V2ReturnsRfbsRejectRequest struct {
	Comment string `json:"comment"`
	RejectionReasonID int64 `json:"rejection_reason_id"`
	ReturnID int64 `json:"return_id"`
}

type V2ReturnsRfbsGetRequest struct {
	ReturnID int64 `json:"return_id"`
}

type V2ReturnsRfbsListRequest struct {
	Filter interface{} `json:"filter"`
	LastID int32 `json:"last_id"`
	Limit int32 `json:"limit"`
}

type V2ReturnsRfbsListResponse struct {
	Returns interface{} `json:"returns"`
}

type V1GetReturnsListResponse struct {
	Returns []interface{} `json:"returns"`
	HasNext bool `json:"has_next"`
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"`
}

type V1ReturnsRfbsActionSetRequest struct {
	CompensationAmount float64 `json:"compensation_amount"`
	ID int32 `json:"id"`
	RejectionReasonID int32 `json:"rejection_reason_id"`
	ReturnForBackWay float64 `json:"return_for_back_way"`
	ReturnID int64 `json:"return_id"`
	Comment string `json:"comment"`
}

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID int64 `json:"return_id"`
	ReturnMethodDescription string `json:"return_method_description"`
}

type V2ReturnsRfbsGetResponse struct {
	Returns interface{} `json:"returns"`
}

type V2GetConditionalCancellationListV2Request struct {
	Filters interface{} `json:"filters"`
	LastID int64 `json:"last_id"`
	Limit int32 `json:"limit"`
	With interface{} `json:"with"`
}

type V2ReturnsRfbsCompensateRequest struct {
	ReturnID int64 `json:"return_id"`
	CompensationAmount string `json:"compensation_amount"`
}

type V1GetReturnsListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
	LastID int64 `json:"last_id"`
}

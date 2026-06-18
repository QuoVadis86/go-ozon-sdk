package pass

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"`
	HasNext bool `json:"has_next"`
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type SellerSellerAPIArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
}

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type ArrivalpassArrivalPassListResponse struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	Cursor string `json:"cursor"`
}

type SellerSellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"`
	CarriageID int64 `json:"carriage_id"`
}

type SellerSellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
}

type ArrivalpassArrivalPassListRequest struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"`
}

type V1ReturnsCompanyFbsInfoRequest struct {
	Filter interface{} `json:"filter"`
	Pagination interface{} `json:"pagination"`
}

type SellerSellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"`
	CarriageID int64 `json:"carriage_id"`
}

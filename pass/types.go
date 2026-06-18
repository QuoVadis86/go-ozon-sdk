package pass

type V1ReturnsCompanyFbsInfoRequest struct {
	Filter interface{} `json:"filter"`
	Pagination interface{} `json:"pagination"`
}

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []interface{} `json:"drop_off_points"` // 揽收点信息。
	HasNext bool `json:"has_next"` // 是否还有其他揽收点等待卖家退货的标志。
}

type SellerSellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"` // 通行证列表。
	CarriageID int64 `json:"carriage_id"` // 运输ID。
}

type SellerSellerAPIArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"` // 通行证列表。
	CarriageID int64 `json:"carriage_id"` // 运输ID。
}

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"` // 通行证列表。
}

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"` // 通行证ID。
}

// 请求结果。
type SellerSellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"` // 通行证ID。
}

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"` // 通行证列表。
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"` // 通行证ID。
}

type ArrivalpassArrivalPassListRequest struct {
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Filter interface{} `json:"filter"`
	Limit int32 `json:"limit"` // 响应中记录数量的限制。
}

type ArrivalpassArrivalPassListResponse struct {
	ArrivalPasses []interface{} `json:"arrival_passes"` // 运输通行证列表。
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。 如果此参数为空，则没有更多数据。
}

type SellerSellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []interface{} `json:"arrival_passes"` // 通行证列表。
	CarriageID int64 `json:"carriage_id"` // 运输ID。
}

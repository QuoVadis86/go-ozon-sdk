package rating

type V1GetFBSRatingIndexInfoV1Response struct {
	Index float64 `json:"index"`
	PeriodFrom string `json:"period_from"`
	PeriodTo string `json:"period_to"`
	ProcessingCostsSum float64 `json:"processing_costs_sum"`
	CurrencyCode string `json:"currency_code"`
	Defects []interface{} `json:"defects"`
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	Cursor string `json:"cursor"`
	Errors []interface{} `json:"errors"`
	HasNext bool `json:"has_next"`
}

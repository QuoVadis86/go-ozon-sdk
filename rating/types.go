package rating

type V1GetFBSRatingIndexInfoV1Response struct {
	PeriodFrom string `json:"period_from"` // 计算周期开始日期（格式`YYYY-MM-DD`）。
	PeriodTo string `json:"period_to"` // 计算周期结束日期（格式`YYYY-MM-DD`）。
	ProcessingCostsSum float64 `json:"processing_costs_sum"` // 周期内的错误处理费用。
	CurrencyCode string `json:"currency_code"` // 错误处理费用的币种代码。
	Defects []interface{} `json:"defects"` // 按天计算的错误指数。
	Index float64 `json:"index"` // 周期内的错误指数数值。
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Limit int64 `json:"limit"` // 返回结果中的数值数量。
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Filter interface{} `json:"filter"`
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	HasNext bool `json:"has_next"` // `true`，表示查询结果未包含所有货件。
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Errors []interface{} `json:"errors"` // 影响错误指数的货件。
}

package premium

type GetRealizationReportByDayResponse struct {
	Rows []interface{} `json:"rows"`
}

type V1SearchQueriesTopRequest struct {
	Limit string `json:"limit"`
	Offset string `json:"offset"`
}

type V1SearchQueriesTopResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type V1SearchQueriesTextRequest struct {
	Offset string `json:"offset"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	Text string `json:"text"`
	Limit string `json:"limit"`
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

type V1AnalyticsProductQueriesResponse struct {
	AnalyticsPeriod interface{} `json:"analytics_period"`
	Items []interface{} `json:"items"`
	PageCount int64 `json:"page_count"`
	Total int64 `json:"total"`
}

type V1ProductPricesDetailsRequest struct {
	Skus []interface{} `json:"skus"`
}

type AnalyticsAnalyticsGetDataRequest struct {
	Limit int64 `json:"limit"`
	Metrics []interface{} `json:"metrics"`
	Offset int64 `json:"offset"`
	Sort []interface{} `json:"sort"`
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Dimension []interface{} `json:"dimension"`
	Filters []interface{} `json:"filters"`
}

type V1GetRealizationReportByDayRequest struct {
	Day int32 `json:"day"`
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type V1ProductPricesDetailsResponse struct {
	Prices []interface{} `json:"prices"`
}

type V1AnalyticsProductQueriesDetailsResponse struct {
	Queries []interface{} `json:"queries"`
	Total int64 `json:"total"`
	AnalyticsPeriod interface{} `json:"analytics_period"`
	PageCount int64 `json:"page_count"`
}

type AnalyticsAnalyticsGetDataResponse struct {
	Result interface{} `json:"result"`
	Timestamp string `json:"timestamp"`
}

type V1SearchQueriesTextResponse struct {
	Offset string `json:"offset"`
	SearchQueries []interface{} `json:"search_queries"`
	Total string `json:"total"`
}

type V1AnalyticsProductQueriesRequest struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Skus []interface{} `json:"skus"`
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
}

package premium

type AnalyticsAnalyticsGetDataRequest struct {
	Dimension []interface{} `json:"dimension"` // 报告中的分组数据。 所有卖家可用的分组方法： - `unknownDimension` — 未知商品标识符； - `sku` — 商品标识符； - `spu` — 商品标识符 — 统一商品卡片； - `day` — 日； - `week` ...
	Filters []interface{} `json:"filters"` // 过滤器。
	Limit int64 `json:"limit"` // 响应的值个数： - 最大值 — 1000， - 最小值 — 1.
	Metrics []interface{} `json:"metrics"` // 最多指定14个指标。如有更多，您将收到 `InvalidArgument`的错误。 生成报告所依据的指标列表。 所有卖家可用的指标： - `revenue` — 订购的金额， - `ordered_units` — 订购的商品。 仅对Pre...
	Offset int64 `json:"offset"` // 响应中要跳过的元素数字。例如，如果 `offset = 10`, 那么答案将从找到的第11个元素开始。
	Sort []interface{} `json:"sort"` // 报告排列设置。
	DateFrom string `json:"date_from"` // 数据将出现在报告中的日期。 若您没有Premium订阅，请指定过去三个月内的日期。
	DateTo string `json:"date_to"` // 数据将出现在报告中的截止日期。
}

type V1GetRealizationReportByDayRequest struct {
	Year int32 `json:"year"` // 年。
	Day int32 `json:"day"` // 日。
	Month int32 `json:"month"` // 月。
}

type V1AnalyticsProductQueriesRequest struct {
	DateFrom string `json:"date_from"` // 分析数据的起始日期。
	DateTo string `json:"date_to"` // 分析数据的结束日期。
	Page int32 `json:"page"` // 请求返回的页码。
	PageSize int32 `json:"page_size"` // 每页包含的商品数量。
	Skus []interface{} `json:"skus"` // SKU 列表，即 Ozon 系统中的商品标识符。根据这些 SKU 返回搜索查询的分析数据。最多可查询 1000 个 SKU。
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
}

type V1AnalyticsProductQueriesResponse struct {
	AnalyticsPeriod interface{} `json:"analytics_period"`
	Items []interface{} `json:"items"` // 商品列表。
	PageCount int64 `json:"page_count"` // 总页数。
	Total int64 `json:"total"` // 搜索请求的总数。
}

type V1SearchQueriesTextRequest struct {
	Text string `json:"text"` // 按文本搜索。
	Limit string `json:"limit"` // 每页的结果数量。
	Offset string `json:"offset"` // 响应中将被跳过的项目数量。
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
}

type V1SearchQueriesTextResponse struct {
	Offset string `json:"offset"` // К每页显示的搜索查询数量。
	SearchQueries []interface{} `json:"search_queries"` // 搜索查询信息。
	Total string `json:"total"` // 搜索查询总数。
}

type V1ProductPricesDetailsResponse struct {
	Prices []interface{} `json:"prices"` // 商品价格。
}

type V1AnalyticsProductQueriesDetailsRequest struct {
	Page int32 `json:"page"` // 请求返回的页码。最小值为0。
	PageSize int32 `json:"page_size"` // 每页包含的商品数量。最大值为100。
	Skus []interface{} `json:"skus"` // SKU 列表，即 Ozon 系统中的商品标识符。根据这些 SKU 返回搜索查询的分析数据。最多可查询 1000 个 SKU。
	SortBy interface{} `json:"sort_by"`
	SortDir interface{} `json:"sort_dir"`
	DateFrom string `json:"date_from"` // 分析数据的起始日期。
	DateTo string `json:"date_to"` // 分析数据的结束日期。
	LimitBySKU int32 `json:"limit_by_sku"` // 单个SKU的查询数量限制。最大值为15次查询。
}

type V1AnalyticsProductQueriesDetailsResponse struct {
	PageCount int64 `json:"page_count"` // 总页数。
	Queries []interface{} `json:"queries"` // 查询列表。
	Total int64 `json:"total"` // 搜索请求的总数。
	AnalyticsPeriod interface{} `json:"analytics_period"`
}

// 查询结果。
type GetRealizationReportByDayResponse struct {
	Rows []interface{} `json:"rows"` // 报告表格。
}

type V1SearchQueriesTopRequest struct {
	Limit string `json:"limit"` // 每页的结果数量。
	Offset string `json:"offset"` // 响应中将被跳过的项目数量。
}

type AnalyticsAnalyticsGetDataResponse struct {
	Result interface{} `json:"result"`
	Timestamp string `json:"timestamp"` // 报告创建时间。
}

type V1ProductPricesDetailsRequest struct {
	Skus []interface{} `json:"skus"` // SKU列表。
}

type V1SearchQueriesTopResponse struct {
	Offset string `json:"offset"` // 每页显示的搜索查询数量。
	SearchQueries []interface{} `json:"search_queries"` // 搜索查询信息。
	Total string `json:"total"` // 搜索查询总数。
}

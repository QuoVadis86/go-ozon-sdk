package finance

type Financev3FinanceTransactionListV3Request struct {
	Filter interface{} `json:"filter"`
	Page int64 `json:"page"` // 请求中返回的页码。
	PageSize int64 `json:"page_size"` // 每页的元素数。
}

type CreateReportResponse struct {
	Result interface{} `json:"result"`
}

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"` // 月。
	Year int32 `json:"year"` // 年。
}

type V2GetRealizationReportResponseV2 struct {
	Result interface{} `json:"result"`
}

type Financev3FinanceTransactionTotalsV3Request struct {
	Date interface{} `json:"date"`
	PostingNumber string `json:"posting_number"` // 发货号。
	TransactionType string `json:"transaction_type"` // 操作类型： - `all` — 所有, - `orders` — 订单, - `returns` — 退货和取消, - `services` — 服务费, - `compensation` — 补贴, - `transferDelivery...
}

// 查询结果。
type V1GetRealizationReportPostingResponse struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"` // 报告表格。
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result interface{} `json:"result"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result interface{} `json:"result"`
}

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"` // 报告周期格式为 `YYYY-MM`。
	Language interface{} `json:"language"`
}

type V1GetCompensationReportRequest struct {
	Date string `json:"date"` // 报告周期格式为 `YYYY-MM`。
	Language interface{} `json:"language"`
}

type V1GetRealizationReportPostingRequest struct {
	Month int32 `json:"month"` // 月。
	Year int32 `json:"year"` // 年。
}

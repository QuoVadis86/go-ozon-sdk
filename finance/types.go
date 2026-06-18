package finance

type V2GetRealizationReportRequestV2 struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type V1GetDecompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type Financev3FinanceTransactionListV3Response struct {
	Result interface{} `json:"result"`
}

type V1GetRealizationReportPostingRequest struct {
	Month int32 `json:"month"`
	Year int32 `json:"year"`
}

type Financev3FinanceTransactionTotalsV3Response struct {
	Result interface{} `json:"result"`
}

type V1GetCompensationReportRequest struct {
	Date string `json:"date"`
	Language interface{} `json:"language"`
}

type V1GetRealizationReportPostingResponse struct {
	Header interface{} `json:"header"`
	Rows []interface{} `json:"rows"`
}

type V2GetRealizationReportResponseV2 struct {
	Result interface{} `json:"result"`
}

type Financev3FinanceTransactionTotalsV3Request struct {
	Date interface{} `json:"date"`
	PostingNumber string `json:"posting_number"`
	TransactionType string `json:"transaction_type"`
}

type CreateReportResponse struct {
	Result interface{} `json:"result"`
}

type Financev3FinanceTransactionListV3Request struct {
	Filter interface{} `json:"filter"`
	Page int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

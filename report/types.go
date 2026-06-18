package report

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"`
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date interface{} `json:"date"`
}

type ReportReportInfoRequest struct {
	Code string `json:"code"`
}

type ReportCreateCompanyProductsReportRequest struct {
	SKU []interface{} `json:"sku"`
	Visibility interface{} `json:"visibility"`
	Language interface{} `json:"language"`
	OfferID []interface{} `json:"offer_id"`
	Search string `json:"search"`
}

type ReportCreateDiscountedRequest interface{}

type V3FinanceCashFlowStatementListResponse struct {
	Result interface{} `json:"result"`
}

type ReportReportListRequest struct {
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	ReportType interface{} `json:"report_type"`
}

type ReportCreateCompanyPostingsReportRequest struct {
	Filter interface{} `json:"filter"`
	Language interface{} `json:"language"`
	With interface{} `json:"with"`
}

type V1CreateStockByWarehouseReportRequest struct {
	Language interface{} `json:"language"`
	WarehouseId []interface{} `json:"warehouseId"`
}

type ReportReportListResponse struct {
	Result interface{} `json:"result"`
}

type CommonCreateReportResponse struct {
	Result interface{} `json:"result"`
}

type V3FinanceCashFlowStatementListRequest struct {
	Date interface{} `json:"date"`
	Page int32 `json:"page"`
	WithDetails bool `json:"with_details"`
	PageSize int32 `json:"page_size"`
}

type ReportReportInfoResponse struct {
	Result interface{} `json:"result"`
}

type ReportCreateReportResponse struct {
	Result interface{} `json:"result"`
}

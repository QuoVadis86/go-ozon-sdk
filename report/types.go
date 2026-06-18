package report

type ReportReportInfoResponse struct {
	Result interface{} `json:"result"`
}

type ReportReportListResponse struct {
	Result interface{} `json:"result"`
}

type ReportCreateDiscountedRequest interface{}

type ReportReportInfoRequest struct {
	Code string `json:"code"` // 报告的唯一识别码。
}

type ReportCreateCompanyPostingsReportRequest struct {
	Filter interface{} `json:"filter"`
	Language interface{} `json:"language"`
	With interface{} `json:"with"`
}

type ReportCreateDiscountedResponse struct {
	Code string `json:"code"` // 报告的唯一识别码。要获取报告，请将此值传递到方法 [/v1/report/info](#operation/ReportAPI_ReportInfo)。
}

type ReportCreateCompanyProductsReportRequest struct {
	Search string `json:"search"` // 在记录内容中搜索，检查现货。
	SKU []interface{} `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	Visibility interface{} `json:"visibility"`
	Language interface{} `json:"language"`
	OfferID []interface{} `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
}

type ReportReportListRequest struct {
	Page int32 `json:"page"` // 页数。
	PageSize int32 `json:"page_size"` // 每页的值的数量： - 默认值 — 100， - 最大值 — 1,000。
	ReportType interface{} `json:"report_type"`
}

type V3FinanceCashFlowStatementListRequest struct {
	Date interface{} `json:"date"`
	Page int32 `json:"page"` // 请求返回中的页码。
	WithDetails bool `json:"with_details"` // `true`，如果需要在响应中添加附加参数。
	PageSize int32 `json:"page_size"` // 页面上的元素数量。
}

type V3FinanceCashFlowStatementListResponse struct {
	Result interface{} `json:"result"`
}

type V1ReportMarkedProductsSalesCreateRequest struct {
	Date interface{} `json:"date"`
}

type V1CreateStockByWarehouseReportRequest struct {
	Language interface{} `json:"language"`
	WarehouseId []interface{} `json:"warehouseId"` // 仓库ID。 请求中参数值的限制。 最大值为 50。
}

type ReportCreateReportResponse struct {
	Result interface{} `json:"result"`
}

type CommonCreateReportResponse struct {
	Result interface{} `json:"result"`
}

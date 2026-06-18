package rating

type GetFBSRatingIndexInfoV1ResponseIndexDynamics struct {
	Date string `json:"date"` // 日期格式为`YYYY-MM-DD`。
	IndexByDate float64 `json:"index_by_date"` // 错误指数的数值。
	ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"` // 错误处理费用。
}

type V1GetFBSRatingIndexInfoV1Response struct {
	PeriodTo string `json:"period_to"` // 计算周期结束日期（格式`YYYY-MM-DD`）。
	ProcessingCostsSum float64 `json:"processing_costs_sum"` // 周期内的错误处理费用。
	CurrencyCode string `json:"currency_code"` // 错误处理费用的币种代码。
	Defects []GetFBSRatingIndexInfoV1ResponseIndexDynamics `json:"defects"` // 按天计算的错误指数。
	Index float64 `json:"index"` // 周期内的错误指数数值。
	PeriodFrom string `json:"period_from"` // 计算周期开始日期（格式`YYYY-MM-DD`）。
}

// 筛选器。
type ListFBSRatingIndexPostingsV1RequestFilter struct {
	DateFrom string `json:"date_from"` // 周期开始日期。
	DateTo string `json:"date_to"` // 周期结束日期。
	PostingNumbers []string `json:"posting_numbers"` // 货件编号。
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Limit int64 `json:"limit"` // 返回结果中的数值数量。
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Filter ListFBSRatingIndexPostingsV1RequestFilter `json:"filter"`
}

// 错误类型： - `UNSPECIFIED`——未指定； - `SELLER_CANCELLATION`——卖家取消； - `SELLER_DELAY`——卖家逾期。
type PostingErrorTypeEnum string

// 商品价值的币种代码：
type ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum string
const (
	ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum_RUB ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum = "RUB"
	ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum_BYN ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum = "BYN"
	ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum_KZT ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum = "KZT"
	ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum_EUR ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum = "EUR"
	ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum_USD ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum = "USD"
	ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum_CNY ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum = "CNY"
)

// 错误处理费用的币种代码：
type ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum string
const (
	ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum_RUB ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum = "RUB"
	ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum_BYN ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum = "BYN"
	ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum_KZT ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum = "KZT"
	ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum_EUR ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum = "EUR"
	ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum_USD ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum = "USD"
	ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum_CNY ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum = "CNY"
)

// 配送方案：
type ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum string
const (
	ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum_FBS ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum = "FBS"
	ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum_RFBS ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum = "rFBS"
	ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum_ErFBS ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum = "erFBS"
)

type ListFBSRatingIndexPostingsV1ResponseError struct {
	ChargePercent float64 `json:"charge_percent"` // 处理费用占货件价值的百分比。
	ChargePrice float64 `json:"charge_price"` // 错误处理费用金额。
	HasGraceStatus bool `json:"has_grace_status"` // `true`，表示货件享有优惠状态。
	PostingNumber string `json:"posting_number"` // 货件编号。
	ProductPrice float64 `json:"product_price"` // 货件中的商品价值。
	ProductPriceCurrencyCode ListFBSRatingIndexPostingsV1ResponseErrorProductPriceCurrencyCodeEnum `json:"product_price_currency_code"` // 商品价值的币种代码： - `RUB`——俄罗斯卢布， - `BYN`——白俄罗斯卢布， - `KZT`——坚戈， - `EUR`——欧元， - `USD`——美元， - `CNY`——人民币。
	ChargePriceCurrencyCode ListFBSRatingIndexPostingsV1ResponseErrorChargePriceCurrencyCodeEnum `json:"charge_price_currency_code"` // 错误处理费用的币种代码： - `RUB`——俄罗斯卢布， - `BYN`——白俄罗斯卢布， - `KZT`——坚戈， - `EUR`——欧元， - `USD`——美元， - `CNY`——人民币。
	DeliverySchema ListFBSRatingIndexPostingsV1ResponseErrorDeliverySchemaEnum `json:"delivery_schema"` // 配送方案： - `FBS`, - `rFBS`, - `erFBS`.
	ErrorAt string `json:"error_at"` // 出现错误的日期。
	Index float64 `json:"index"` // 错误指数的数值。
	PostingErrorType PostingErrorTypeEnum `json:"posting_error_type"`
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Errors []ListFBSRatingIndexPostingsV1ResponseError `json:"errors"` // 影响错误指数的货件。
	HasNext bool `json:"has_next"` // `true`，表示查询结果未包含所有货件。
}

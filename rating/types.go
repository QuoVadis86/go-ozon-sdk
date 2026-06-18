package rating

// 筛选器。
type ListFBSRatingIndexPostingsV1RequestFilter struct {
	DateFrom       string   `json:"date_from"`       // 周期开始日期。
	DateTo         string   `json:"date_to"`         // 周期结束日期。
	PostingNumbers []string `json:"posting_numbers"` // 货件编号。
}

type V1ListFBSRatingIndexPostingsV1Request struct {
	Filter ListFBSRatingIndexPostingsV1RequestFilter `json:"filter"`
	Limit  int64                                     `json:"limit"`  // 返回结果中的数值数量。
	Cursor string                                    `json:"cursor"` // 用于获取下一批数据的指针。
}

// 错误类型： - `UNSPECIFIED`——未指定； - `SELLER_CANCELLATION`——卖家取消； - `SELLER_DELAY`——卖家逾期。
type PostingErrorTypeEnum string

// DeliverySchema values
type DeliverySchema string

const (
	DeliverySchemaFBS   DeliverySchema = "FBS"
	DeliverySchemaRFBS  DeliverySchema = "rFBS"
	DeliverySchemaErFBS DeliverySchema = "erFBS"
)

// ProductPriceCurrencyCode values
type ProductPriceCurrencyCode string

const (
	ProductPriceCurrencyCodeRUB ProductPriceCurrencyCode = "RUB" // 俄罗斯卢布
	ProductPriceCurrencyCodeBYN ProductPriceCurrencyCode = "BYN" // 白俄罗斯卢布
	ProductPriceCurrencyCodeKZT ProductPriceCurrencyCode = "KZT" // 坚戈
	ProductPriceCurrencyCodeEUR ProductPriceCurrencyCode = "EUR" // 欧元
	ProductPriceCurrencyCodeUSD ProductPriceCurrencyCode = "USD" // 美元
	ProductPriceCurrencyCodeCNY ProductPriceCurrencyCode = "CNY" // 人民币
)

// ChargePriceCurrencyCode values
type ChargePriceCurrencyCode string

const (
	ChargePriceCurrencyCodeRUB ChargePriceCurrencyCode = "RUB" // 俄罗斯卢布
	ChargePriceCurrencyCodeBYN ChargePriceCurrencyCode = "BYN" // 白俄罗斯卢布
	ChargePriceCurrencyCodeKZT ChargePriceCurrencyCode = "KZT" // 坚戈
	ChargePriceCurrencyCodeEUR ChargePriceCurrencyCode = "EUR" // 欧元
	ChargePriceCurrencyCodeUSD ChargePriceCurrencyCode = "USD" // 美元
	ChargePriceCurrencyCodeCNY ChargePriceCurrencyCode = "CNY" // 人民币
)

type ListFBSRatingIndexPostingsV1ResponseError struct {
	ChargePercent            float64                  `json:"charge_percent"`             // 处理费用占货件价值的百分比。
	ChargePriceCurrencyCode  ChargePriceCurrencyCode  `json:"charge_price_currency_code"` // 错误处理费用的币种代码： - `RUB`——俄罗斯卢布， - `BYN`——白俄罗斯卢布， - `KZT`——坚戈， - `EUR`——欧元， - `USD`——美元， - `CNY`——人民币。
	ErrorAt                  string                   `json:"error_at"`                   // 出现错误的日期。
	HasGraceStatus           bool                     `json:"has_grace_status"`           // `true`，表示货件享有优惠状态。
	Index                    float64                  `json:"index"`                      // 错误指数的数值。
	PostingNumber            string                   `json:"posting_number"`             // 货件编号。
	ProductPrice             float64                  `json:"product_price"`              // 货件中的商品价值。
	ChargePrice              float64                  `json:"charge_price"`               // 错误处理费用金额。
	DeliverySchema           DeliverySchema           `json:"delivery_schema"`            // 配送方案： - `FBS`, - `rFBS`, - `erFBS`.
	PostingErrorType         PostingErrorTypeEnum     `json:"posting_error_type"`
	ProductPriceCurrencyCode ProductPriceCurrencyCode `json:"product_price_currency_code"` // 商品价值的币种代码： - `RUB`——俄罗斯卢布， - `BYN`——白俄罗斯卢布， - `KZT`——坚戈， - `EUR`——欧元， - `USD`——美元， - `CNY`——人民币。
}

type V1ListFBSRatingIndexPostingsV1Response struct {
	Errors  []ListFBSRatingIndexPostingsV1ResponseError `json:"errors"`   // 影响错误指数的货件。
	HasNext bool                                        `json:"has_next"` // `true`，表示查询结果未包含所有货件。
	Cursor  string                                      `json:"cursor"`   // 用于获取下一批数据的指针。
}

type GetFBSRatingIndexInfoV1ResponseIndexDynamics struct {
	ProcessingCostsSumByDate float64 `json:"processing_costs_sum_by_date"` // 错误处理费用。
	Date                     string  `json:"date"`                         // 日期格式为`YYYY-MM-DD`。
	IndexByDate              float64 `json:"index_by_date"`                // 错误指数的数值。
}

type V1GetFBSRatingIndexInfoV1Response struct {
	ProcessingCostsSum float64                                        `json:"processing_costs_sum"` // 周期内的错误处理费用。
	CurrencyCode       string                                         `json:"currency_code"`        // 错误处理费用的币种代码。
	Defects            []GetFBSRatingIndexInfoV1ResponseIndexDynamics `json:"defects"`              // 按天计算的错误指数。
	Index              float64                                        `json:"index"`                // 周期内的错误指数数值。
	PeriodFrom         string                                         `json:"period_from"`          // 计算周期开始日期（格式`YYYY-MM-DD`）。
	PeriodTo           string                                         `json:"period_to"`            // 计算周期结束日期（格式`YYYY-MM-DD`）。
}

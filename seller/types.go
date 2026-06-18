package seller

// 评级状态。
type V1RatingStatus struct {
	Danger  bool `json:"danger"`  // 指示是否已超过封禁评级阈值。
	Premium bool `json:"premium"` // 指示是否已达到参与高级计划的阈值。
	Warning bool `json:"warning"` // 指示是否已发出关于可能超过封禁阈值的警告。
}

// 评级值。
type ValueCurrent struct {
	DateTo    string         `json:"date_to"`   // 评级统计结束日期。
	Formatted string         `json:"formatted"` // 格式化后的评级值。
	Status    V1RatingStatus `json:"status"`
	Value     float64        `json:"value"`     // 系统中的评级值。
	DateFrom  string         `json:"date_from"` // 评级统计开始日期。
}

// 上一次的评级值。
type ValuePast struct {
	DateFrom  string         `json:"date_from"` // 评级统计开始日期。
	DateTo    string         `json:"date_to"`   // 评级统计结束日期。
	Formatted string         `json:"formatted"` // 格式化后的评级值。
	Status    V1RatingStatus `json:"status"`
	Value     float64        `json:"value"` // 系统中的评级值。
}

// 评级状态： - `UNKNOWN`——未指定； - `OK`——良好； - `WARNING`——指标需关注； - `CRITICAL`——严重。
type StatusEnum string

// 值的类型： - `UNKNOWN`——未指定； - `INDEX`——指数； - `PERCENT`——百分比； - `TIME`——时间； - `RATIO`——系数； - `REVIEW_SCORE`——评分； - `COUNT`——数...
type InfoResponseRatingTypeEnum string

type InfoResponseRating struct {
	Status       StatusEnum                 `json:"status"`
	ValueType    InfoResponseRatingTypeEnum `json:"value_type"`
	CurrentValue ValueCurrent               `json:"current_value"`
	Name         string                     `json:"name"` // 评级名称。
	PastValue    ValuePast                  `json:"past_value"`
	Rating       string                     `json:"rating"` // 系统中的评级名称。
}

// 税收制度： - `UNKNOWN`——未知； - `UNSPECIFIED`——未指定； - `OSNO`——普通税制（OSNO）； - `USN`——简化税制（USN）； - `NPD`——职业收入税（NPD）； - `AUSN`——自动...
type CompanyTaxSystemEnum string

// 公司。
// Currency values
type Currency string

const (
	CurrencyRUB Currency = "RUB" // 俄罗斯卢布；
	CurrencyEUR Currency = "EUR" // 欧元；
	CurrencyUSD Currency = "USD" // 美元；
	CurrencyCNY Currency = "CNY" // 人民币；
	CurrencyBYN Currency = "BYN" // 白俄罗斯卢布；
	CurrencyKZT Currency = "KZT" // 坚戈；
	CurrencyKGS Currency = "KGS" // 吉尔吉斯斯坦索姆
)

type InfoResponseCompany struct {
	OwnershipForm string               `json:"ownership_form"` // 所有制形式。
	TaxSystem     CompanyTaxSystemEnum `json:"tax_system"`
	Country       string               `json:"country"`    // 国家。
	Currency      Currency             `json:"currency"`   // 货币： - `RUB`——俄罗斯卢布； - `EUR`——欧元； - `USD`——美元； - `CNY`——人民币； - `BYN`——白俄罗斯卢布； - `KZT`——坚戈； - `KGS`——吉尔吉斯斯坦索姆。
	Inn           string               `json:"inn"`        // 税号（INN）。
	LegalName     string               `json:"legal_name"` // 法人名称。
	Name          string               `json:"name"`       // Ozon上的公司名称。
	Ogrn          string               `json:"ogrn"`       // 国家基本登记号（OGRN）。
}

// 订阅类型： - `UNKNOWN`——未指定； - `UNSPECIFIED`——无需订阅； - `PREMIUM`——Premium； - `PREMIUM_LITE`——Premium Lite； - `PREMIUM_PLUS`——P...
type InfoResponseSubscriptionTypeEnum string

// 订阅。
type InfoResponseSubscription struct {
	IsPremium bool                             `json:"is_premium"` // `true`，前提是有订阅。
	Type      InfoResponseSubscriptionTypeEnum `json:"type_"`
}

type V1SellerInfoResponse struct {
	Company      InfoResponseCompany      `json:"company"`
	Ratings      []InfoResponseRating     `json:"ratings"` // 评级列表。
	Subscription InfoResponseSubscription `json:"subscription"`
}

type RolesByTokenResponseRoles struct {
	Name    string   `json:"name"`    // 角色名称。
	Methods []string `json:"methods"` // 角色可用的方式。
}

type V1RolesByTokenResponse struct {
	ExpiresAt string                      `json:"expires_at"` // 密钥到期日期。
	Roles     []RolesByTokenResponseRoles `json:"roles"`      // 可用角色和方式信息。
}

type OzonLogisticsInfoResponseAvailableSchemasEnum string

// AvailableSchemas values
type AvailableSchemas string

const (
	AvailableSchemasUnknown AvailableSchemas = "UNKNOWN" // 未指定；
	AvailableSchemasFBO     AvailableSchemas = "FBO"     // 从Ozon仓库配送；
	AvailableSchemasFBS     AvailableSchemas = "FBS"     // 从自有仓库配送
)

type V1SellerOzonLogisticsInfoResponse struct {
	AvailableSchemas     []OzonLogisticsInfoResponseAvailableSchemasEnum `json:"available_schemas"`      // 可用模式类型： - `UNKNOWN`——未指定； - `FBO`——从Ozon仓库配送； - `FBS`——从自有仓库配送。
	OzonLogisticsEnabled bool                                            `json:"ozon_logistics_enabled"` // `true`，表示Ozon配送已开通。
}

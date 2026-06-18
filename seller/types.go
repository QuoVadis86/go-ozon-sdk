package seller

import "github.com/QuoVadis86/go-ozon-sdk/enum"

// 评级状态： - `UNKNOWN`——未指定； - `OK`——良好； - `WARNING`——指标需关注； - `CRITICAL`——严重。
type StatusEnum string

const (
	StatusEnumUnknown  StatusEnum = "UNKNOWN"
	StatusEnumOK       StatusEnum = "OK"
	StatusEnumWarning  StatusEnum = "WARNING"
	StatusEnumCritical StatusEnum = "CRITICAL"
)

// 评级状态。
type V1RatingStatus struct {
	Danger  bool `json:"danger"`  // 指示是否已超过封禁评级阈值。
	Premium bool `json:"premium"` // 指示是否已达到参与高级计划的阈值。
	Warning bool `json:"warning"` // 指示是否已发出关于可能超过封禁阈值的警告。
}

// 评级值。
type ValueCurrent struct {
	DateFrom  string         `json:"date_from"` // 评级统计开始日期。
	DateTo    string         `json:"date_to"`   // 评级统计结束日期。
	Formatted string         `json:"formatted"` // 格式化后的评级值。
	Status    V1RatingStatus `json:"status"`
	Value     float64        `json:"value"` // 系统中的评级值。
}

type OzonLogisticsInfoResponseAvailableSchemas string

const (
	OzonLogisticsInfoResponseAvailableSchemasUnknown OzonLogisticsInfoResponseAvailableSchemas = "UNKNOWN"
	OzonLogisticsInfoResponseAvailableSchemasFBO     OzonLogisticsInfoResponseAvailableSchemas = "FBO"
	OzonLogisticsInfoResponseAvailableSchemasFBS     OzonLogisticsInfoResponseAvailableSchemas = "FBS"
)

// AvailableSchemas values
type AvailableSchemas string

const (
	AvailableSchemasUnknown AvailableSchemas = "UNKNOWN" // 未指定；
	AvailableSchemasFBO     AvailableSchemas = "FBO"     // 从Ozon仓库配送；
	AvailableSchemasFBS     AvailableSchemas = "FBS"     // 从自有仓库配送
)

type V1SellerOzonLogisticsInfoResponse struct {
	OzonLogisticsEnabled bool                                        `json:"ozon_logistics_enabled"` // `true`，表示Ozon配送已开通。
	AvailableSchemas     []OzonLogisticsInfoResponseAvailableSchemas `json:"available_schemas"`      // 可用模式类型： - `UNKNOWN`——未指定； - `FBO`——从Ozon仓库配送； - `FBS`——从自有仓库配送。
}

// 上一次的评级值。
type ValuePast struct {
	DateFrom  string         `json:"date_from"` // 评级统计开始日期。
	DateTo    string         `json:"date_to"`   // 评级统计结束日期。
	Formatted string         `json:"formatted"` // 格式化后的评级值。
	Status    V1RatingStatus `json:"status"`
	Value     float64        `json:"value"` // 系统中的评级值。
}

// 值的类型： - `UNKNOWN`——未指定； - `INDEX`——指数； - `PERCENT`——百分比； - `TIME`——时间； - `RATIO`——系数； - `REVIEW_SCORE`——评分； - `COUNT`——数...
type InfoResponseRatingType string

const (
	InfoResponseRatingTypeUnknown     InfoResponseRatingType = "UNKNOWN"
	InfoResponseRatingTypeIndex       InfoResponseRatingType = "INDEX"
	InfoResponseRatingTypePercent     InfoResponseRatingType = "PERCENT"
	InfoResponseRatingTypeTime        InfoResponseRatingType = "TIME"
	InfoResponseRatingTypeRatio       InfoResponseRatingType = "RATIO"
	InfoResponseRatingTypeReviewScore InfoResponseRatingType = "REVIEW_SCORE"
	InfoResponseRatingTypeCount       InfoResponseRatingType = "COUNT"
)

type InfoResponseRating struct {
	Rating       string                 `json:"rating"` // 系统中的评级名称。
	Status       StatusEnum             `json:"status"`
	ValueType    InfoResponseRatingType `json:"value_type"`
	CurrentValue ValueCurrent           `json:"current_value"`
	Name         string                 `json:"name"` // 评级名称。
	PastValue    ValuePast              `json:"past_value"`
}

type RolesByTokenResponseRoles struct {
	Name    string   `json:"name"`    // 角色名称。
	Methods []string `json:"methods"` // 角色可用的方式。
}

// 税收制度： - `UNKNOWN`——未知； - `UNSPECIFIED`——未指定； - `OSNO`——普通税制（OSNO）； - `USN`——简化税制（USN）； - `NPD`——职业收入税（NPD）； - `AUSN`——自动...
type CompanyTaxSystem string

const (
	CompanyTaxSystemUnknown     CompanyTaxSystem = "UNKNOWN"
	CompanyTaxSystemUnspecified CompanyTaxSystem = "UNSPECIFIED"
	CompanyTaxSystemOsno        CompanyTaxSystem = "OSNO"
	CompanyTaxSystemUSN         CompanyTaxSystem = "USN"
	CompanyTaxSystemNPD         CompanyTaxSystem = "NPD"
	CompanyTaxSystemAusn        CompanyTaxSystem = "AUSN"
	CompanyTaxSystemPSN         CompanyTaxSystem = "PSN"
)

// 公司。
type InfoResponseCompany struct {
	LegalName     string            `json:"legal_name"`     // 法人名称。
	Name          string            `json:"name"`           // Ozon上的公司名称。
	Ogrn          string            `json:"ogrn"`           // 国家基本登记号（OGRN）。
	OwnershipForm string            `json:"ownership_form"` // 所有制形式。
	TaxSystem     CompanyTaxSystem  `json:"tax_system"`
	Country       string            `json:"country"`  // 国家。
	Currency      enum.CurrencyCode `json:"currency"` // 货币： - `RUB`——俄罗斯卢布； - `EUR`——欧元； - `USD`——美元； - `CNY`——人民币； - `BYN`——白俄罗斯卢布； - `KZT`——坚戈； - `KGS`——吉尔吉斯斯坦索姆。
	Inn           string            `json:"inn"`      // 税号（INN）。
}

// 订阅类型： - `UNKNOWN`——未指定； - `UNSPECIFIED`——无需订阅； - `PREMIUM`——Premium； - `PREMIUM_LITE`——Premium Lite； - `PREMIUM_PLUS`——P...
type InfoResponseSubscriptionType string

const (
	InfoResponseSubscriptionTypeUnknown     InfoResponseSubscriptionType = "UNKNOWN"
	InfoResponseSubscriptionTypeUnspecified InfoResponseSubscriptionType = "UNSPECIFIED"
	InfoResponseSubscriptionTypePremium     InfoResponseSubscriptionType = "PREMIUM"
	InfoResponseSubscriptionTypePremiumLite InfoResponseSubscriptionType = "PREMIUM_LITE"
	InfoResponseSubscriptionTypePremiumPlus InfoResponseSubscriptionType = "PREMIUM_PLUS"
	InfoResponseSubscriptionTypePremiumPRO  InfoResponseSubscriptionType = "PREMIUM_PRO"
)

// 订阅。
type InfoResponseSubscription struct {
	IsPremium bool                         `json:"is_premium"` // `true`，前提是有订阅。
	Type      InfoResponseSubscriptionType `json:"type_"`
}

type V1SellerInfoResponse struct {
	Company      InfoResponseCompany      `json:"company"`
	Ratings      []InfoResponseRating     `json:"ratings"` // 评级列表。
	Subscription InfoResponseSubscription `json:"subscription"`
}

type V1RolesByTokenResponse struct {
	ExpiresAt string                      `json:"expires_at"` // 密钥到期日期。
	Roles     []RolesByTokenResponseRoles `json:"roles"`      // 可用角色和方式信息。
}

package seller

type V1SellerInfoResponse struct {
	Company interface{} `json:"company"`
	Ratings []interface{} `json:"ratings"` // 评级列表。
	Subscription interface{} `json:"subscription"`
}

type V1RolesByTokenResponse struct {
	ExpiresAt string `json:"expires_at"` // 密钥到期日期。
	Roles []interface{} `json:"roles"` // 可用角色和方式信息。
}

type V1SellerOzonLogisticsInfoResponse struct {
	OzonLogisticsEnabled bool `json:"ozon_logistics_enabled"` // `true`，表示Ozon配送已开通。
	AvailableSchemas []interface{} `json:"available_schemas"` // 可用模式类型： - `UNKNOWN`——未指定； - `FBO`——从Ozon仓库配送； - `FBS`——从自有仓库配送。
}

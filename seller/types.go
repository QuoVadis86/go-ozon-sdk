package seller

type V1SellerOzonLogisticsInfoResponse struct {
	OzonLogisticsEnabled bool `json:"ozon_logistics_enabled"`
	AvailableSchemas []interface{} `json:"available_schemas"`
}

type V1RolesByTokenResponse struct {
	Roles []interface{} `json:"roles"`
	ExpiresAt string `json:"expires_at"`
}

type V1SellerInfoResponse struct {
	Company interface{} `json:"company"`
	Ratings []interface{} `json:"ratings"`
	Subscription interface{} `json:"subscription"`
}

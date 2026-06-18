package pricing

type V1GetCompetitorsResponse struct {
	Competitor []interface{} `json:"competitor"`
	Total int32 `json:"total"`
}

type V1CreatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyName string `json:"strategy_name"`
}

type V1GetStrategyItemInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyListResponse struct {
	Strategies []interface{} `json:"strategies"`
	Total int32 `json:"total"`
}

type V1GetStrategyIDsByItemIDsResponse struct {
	Result interface{} `json:"result"`
}

type V1UpdatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"`
	StrategyID string `json:"strategy_id"`
	StrategyName string `json:"strategy_name"`
}

type V1GetStrategyResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyItemInfoRequest struct {
	ProductID int64 `json:"product_id"`
}

type V1ItemIDsRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type V1DeleteStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1CreatePricingStrategyResponse struct {
	Result interface{} `json:"result"`
}

type V1StrategyRequest struct {
	StrategyID string `json:"strategy_id"`
}

type V1GetCompetitorsRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1GetStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyListRequest struct {
	Page int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type V1AddStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1AddStrategyItemsRequest struct {
	ProductID []interface{} `json:"product_id"`
	StrategyID string `json:"strategy_id"`
}

type V1UpdateStatusStrategyRequest struct {
	Enabled bool `json:"enabled"`
	StrategyID string `json:"strategy_id"`
}

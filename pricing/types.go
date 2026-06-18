package pricing

type V1GetStrategyListResponse struct {
	Strategies []interface{} `json:"strategies"` // 策略列表。
	Total int32 `json:"total"` // 策略总数。
}

type V1GetCompetitorsResponse struct {
	Competitor []interface{} `json:"competitor"` // 竞争对手列表。
	Total int32 `json:"total"` // 竞争对手总数。
}

type V1GetStrategyItemInfoRequest struct {
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

type V1GetStrategyListRequest struct {
	Page int64 `json:"page"` // 卸载策略的列表页面。 最小值为`1`。
	Limit int64 `json:"limit"` // 每页的最大策略数。有效值是从`1`到`50`。
}

type V1CreatePricingStrategyResponse struct {
	Result interface{} `json:"result"`
}

type V1AddStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyItemInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1GetStrategyIDsByItemIDsResponse struct {
	Result interface{} `json:"result"`
}

type V1CreatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"` // 竞争对手名单。
	StrategyName string `json:"strategy_name"` // 策略名称。
}

type V1GetStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1AddStrategyItemsRequest struct {
	ProductID []interface{} `json:"product_id"` // 商品ID列表。 最大数量为 50。
	StrategyID string `json:"strategy_id"` // 策略ID。
}

type V1UpdatePricingStrategyRequest struct {
	Competitors []interface{} `json:"competitors"` // 竞争对手列表。
	StrategyID string `json:"strategy_id"` // 策略ID。
	StrategyName string `json:"strategy_name"` // 策略名称。
}

type V1StrategyRequest struct {
	StrategyID string `json:"strategy_id"` // 策略ID。
}

type V1ItemIDsRequest struct {
	ProductID []interface{} `json:"product_id"` // 商品ID列表。最大数量 —— 50。
}

type V1GetCompetitorsRequest struct {
	Page int64 `json:"page"` // 需要下载竞争对手的列表页面。 最小值为`1`。
	Limit int64 `json:"limit"` // 每页的最大竞争对手数。有效值是从`1`到`50`。
}

type V1DeleteStrategyItemsResponse struct {
	Result interface{} `json:"result"`
}

type V1UpdateStatusStrategyRequest struct {
	Enabled bool `json:"enabled"` // 策略状态： - `true` —— 打开， - `false` —— 关闭。
	StrategyID string `json:"strategy_id"` // 策略ID。
}

type V1GetStrategyResponse struct {
	Result interface{} `json:"result"`
}

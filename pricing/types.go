package pricing

// 商品列表。
type V1GetStrategyItemsResponseResult struct {
	ProductID []string `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

type V1GetStrategyItemsResponse struct {
	Result V1GetStrategyItemsResponseResult `json:"result"`
}

// 方法操作列表。
type V1DeleteStrategyItemsResponseResult struct {
	FailedProductCount int32 `json:"failed_product_count"` // 有错误的商品数量。
}

type AddStrategyItemsResponseError struct {
	Code      string `json:"code"`       // 错误代码。
	Error     string `json:"error"`      // 错误文本。
	ProductID int64  `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

// 方法操作结果。
type V1AddStrategyItemsResponseResult struct {
	Errors             []AddStrategyItemsResponseError `json:"errors"`               // 有错误的商品。
	FailedProductCount int32                           `json:"failed_product_count"` // 有错误的商品数量。
}

type V1AddStrategyItemsResponse struct {
	Result V1AddStrategyItemsResponseResult `json:"result"`
}

type GetCompetitorsResponseCompetitorInfo struct {
	Name string `json:"name"` // 竞争对手名称。
	ID   int64  `json:"id"`   // 竞争对手ID。
}

type V1GetCompetitorsResponse struct {
	Competitor []GetCompetitorsResponseCompetitorInfo `json:"competitor"` // 竞争对手列表。
	Total      int32                                  `json:"total"`      // 竞争对手总数。
}

type V1Competitor struct {
	CompetitorID int64   `json:"competitor_id"` // 竞争对手ID。
	Coefficient  float64 `json:"coefficient"`   // 竞争对手之间的最低价格将乘以的系数。有效范围是`0.5`到`1.2`。
}

type V1CreatePricingStrategyRequest struct {
	StrategyName string         `json:"strategy_name"` // 策略名称。
	Competitors  []V1Competitor `json:"competitors"`   // 竞争对手名单。
}

type V1UpdateStatusStrategyRequest struct {
	Enabled    bool   `json:"enabled"`     // 策略状态： - `true` —— 打开， - `false` —— 关闭。
	StrategyID string `json:"strategy_id"` // 策略ID。
}

type V1GetStrategyItemInfoRequest struct {
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

type V1GetStrategyListRequest struct {
	Page  int64 `json:"page"`  // 卸载策略的列表页面。 最小值为`1`。
	Limit int64 `json:"limit"` // 每页的最大策略数。有效值是从`1`到`50`。
}

type GetStrategyIDsByItemIDsResponseProductInfo struct {
	ProductID  int64  `json:"product_id"`  // Ozon系统中商品的标识符 — `product_id`。
	StrategyID string `json:"strategy_id"` // 添加商品的策略ID。
}

type V1StrategyRequest struct {
	StrategyID string `json:"strategy_id"` // 策略ID。
}

// 方法操作结果。
// Type values
type Type string

const (
	TypeMINEXTPrice Type = "MIN_EXT_PRICE" // 系统策略
	TypeCompPrice   Type = "COMP_PRICE"    // 用户策略
)

// UpdateType values
type UpdateType string

const (
	UpdateTypeStrategyEnabled          UpdateType = "strategyEnabled"          // 恢复
	UpdateTypeStrategyDisabled         UpdateType = "strategyDisabled"         // 停止
	UpdateTypeStrategyChanged          UpdateType = "strategyChanged"          // 更新
	UpdateTypeStrategyCreated          UpdateType = "strategyCreated"          // 创建
	UpdateTypeStrategyItemsListChanged UpdateType = "strategyItemsListChanged" // 策略中的商品集合已更改
)

type V1GetStrategyResponseResult struct {
	Competitors []V1Competitor `json:"competitors"` // 竞争对手列表。
	Enabled     bool           `json:"enabled"`     // 策略状态： - `true` —— 打开， - `false` —— 关闭。
	Name        string         `json:"name"`        // 策略名称。
	Type        Type           `json:"type_"`       // 策略类型： - `MIN_EXT_PRICE` —— 系统策略， - `COMP_PRICE` —— 用户策略。
	UpdateType  UpdateType     `json:"update_type"` // 上次策略更改的类型： - `strategyEnabled` —— 恢复， - `strategyDisabled` —— 停止， - `strategyChanged` —— 更新， - `strategyCreated` —— 创建， ...
}

type V1GetStrategyResponse struct {
	Result V1GetStrategyResponseResult `json:"result"`
}

type V1DeleteStrategyItemsResponse struct {
	Result V1DeleteStrategyItemsResponseResult `json:"result"`
}

// 方法操作结果。
type V1CreatePricingStrategyResponseResult struct {
	StrategyID string `json:"strategy_id"` // 策略ID。
}

type V1GetCompetitorsRequest struct {
	Page  int64 `json:"page"`  // 需要下载竞争对手的列表页面。 最小值为`1`。
	Limit int64 `json:"limit"` // 每页的最大竞争对手数。有效值是从`1`到`50`。
}

type V1ItemIDsRequest struct {
	ProductID []string `json:"product_id"` // 商品ID列表。最大数量 —— 50。
}

// 方法操作结果。
type V1GetStrategyItemInfoResponseResult struct {
	IsEnabled                    bool   `json:"is_enabled"`                      // `true`, 如果商品参与定价策略。
	StrategyProductPrice         int32  `json:"strategy_product_price"`          // 定价策略。
	PriceDownloadedAt            string `json:"price_downloaded_at"`             // 定价策略设定日期。
	StrategyCompetitorID         int64  `json:"strategy_competitor_id"`          // 竞争对手ID。
	StrategyCompetitorProductURL string `json:"strategy_competitor_product_url"` // 竞争对手商品链接。
	StrategyID                   string `json:"strategy_id"`                     // 策略ID。
}

// 方法操作结果。
type V1GetStrategyIDsByItemIDsResponseResult struct {
	ProductsInfo []GetStrategyIDsByItemIDsResponseProductInfo `json:"products_info"` // 商品信息。
}

type V1GetStrategyIDsByItemIDsResponse struct {
	Result V1GetStrategyIDsByItemIDsResponseResult `json:"result"`
}

type V1CreatePricingStrategyResponse struct {
	Result V1CreatePricingStrategyResponseResult `json:"result"`
}

type V1UpdatePricingStrategyRequest struct {
	Competitors  []V1Competitor `json:"competitors"`   // 竞争对手列表。
	StrategyID   string         `json:"strategy_id"`   // 策略ID。
	StrategyName string         `json:"strategy_name"` // 策略名称。
}

type V1Empty any

type GetStrategyListResponseStrategy struct {
	ProductsCount    int64      `json:"products_count"`    // 策略中的商品数量。
	CompetitorsCount int64      `json:"competitors_count"` // 选择的竞争对手数量。
	Enabled          bool       `json:"enabled"`           // 策略状态： - `true` —— 开启， - `false` —— 关闭。
	ID               string     `json:"id"`                // 策略ID。
	Name             string     `json:"name"`              // 策略名称。
	Type             Type       `json:"type_"`             // 策略类型： - `MIN_EXT_PRICE` —— 系统性， - `COMP_PRICE` —— 用户性。
	UpdateType       UpdateType `json:"update_type"`       // 策略最后变化的类型: - `strategyEnabled` — 恢复， - `strategyDisabled` — 停止， - `strategyChanged` — 更新， - `strategyCreated` — 创建， - `s...
	UpdatedAt        string     `json:"updated_at"`        // 最后一次修改的日期。
}

type V1GetStrategyListResponse struct {
	Strategies []GetStrategyListResponseStrategy `json:"strategies"` // 策略列表。
	Total      int32                             `json:"total"`      // 策略总数。
}

type V1AddStrategyItemsRequest struct {
	ProductID  []string `json:"product_id"`  // 商品ID列表。 最大数量为 50。
	StrategyID string   `json:"strategy_id"` // 策略ID。
}

type V1GetStrategyItemInfoResponse struct {
	Result V1GetStrategyItemInfoResponseResult `json:"result"`
}

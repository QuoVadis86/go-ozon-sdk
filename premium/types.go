package premium

type V1SearchQueriesTextResponseSearchQuery struct {
	AvgPrice float64 `json:"avg_price"` // 商品的平均价格（以卢布计）。
	ClientCount float64 `json:"client_count"` // 通过该,搜索查询 查找商品的买家数量。
	ConversionToCart float64 `json:"conversion_to_cart"` // 将至少 1 件商品添加到购物车的买家比例。
	ItemsViews float64 `json:"items_views"` // 商品的浏览次数。
	Query string `json:"query"` // 搜索查询。
	SellersCount float64 `json:"sellers_count"` // 买家根据该搜索查询查看其商品的平均卖家数量。
	AddToCart float64 `json:"add_to_cart"` // 将至少 1 件商品添加到购物车的买家数量。
}

type V1SearchQueriesTextResponse struct {
	Offset string `json:"offset"` // К每页显示的搜索查询数量。
	SearchQueries []V1SearchQueriesTextResponseSearchQuery `json:"search_queries"` // 搜索查询信息。
	Total string `json:"total"` // 搜索查询总数。
}

// 配送佣金。
type RowItemCommission struct {
	Quantity int32 `json:"quantity"` // 商品数量。
	StandardFee float64 `json:"standard_fee"` // Ozon基础奖励。
	BankCoinvestment float64 `json:"bank_coinvestment"` // 合作伙伴忠诚机制付款：绿色价格。
	Total float64 `json:"total"` // 应计总额。
	Commission float64 `json:"commission"` // 将折扣和附加费考虑在内的总佣金。 适用于 2024 年 4 月 30 日之前生成的报告。
	Stars float64 `json:"stars"` // 合作伙伴忠诚度机制付款：星星。
	Amount float64 `json:"amount"` // 金额。
	Bonus float64 `json:"bonus"` // 折扣积分。
	Compensation float64 `json:"compensation"` // Ozon 负责的补付额。 适用于 2024 年 4 月 30 日之前生成的报告。
	PricePerInstance float64 `json:"price_per_instance"` // 每件价格。
}

// 商品信息。
type RowItem struct {
	Barcode string `json:"barcode"` // 商品条形码。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
	SKU int64 `json:"sku"` // Ozon系统中的商品识别码是SKU。
}

// 商品退货佣金。
type RowItemCommissionReturn struct {
	Amount float64 `json:"amount"` // 金额。
	Bonus float64 `json:"bonus"` // 折扣积分。
	PricePerInstance float64 `json:"price_per_instance"` // 每件价格。
	Quantity int32 `json:"quantity"` // 商品数量。
	BankCoinvestment float64 `json:"bank_coinvestment"` // 合作伙伴忠诚机制付款：绿色价格。
	Stars float64 `json:"stars"` // 合作伙伴忠诚度机制付款：星星。
	Commission float64 `json:"commission"` // 将折扣和附加费考虑在内的总佣金。 适用于 2024 年 4 月 30 日之前生成的报告。
	Compensation float64 `json:"compensation"` // Ozon 负责的补付额。 适用于 2024 年 4 月 30 日之前生成的报告。
	StandardFee float64 `json:"standard_fee"` // Ozon基础奖励。
	Total float64 `json:"total"` // 应计总额。
}

type GetRealizationReportByDayResponseRow struct {
	CommissionRatio float64 `json:"commission_ratio"` // 按类目划分的销售佣金比例。
	DeliveryCommission RowItemCommission `json:"delivery_commission"`
	Item RowItem `json:"item"`
	ReturnCommission RowItemCommissionReturn `json:"return_commission"`
	RowNumber int32 `json:"rowNumber"` // 报告中的行号。
	SellerPricePerInstance float64 `json:"seller_price_per_instance"` // 考虑折扣后的卖家价格。
}

// 查询结果。
type GetRealizationReportByDayResponse struct {
	Rows []GetRealizationReportByDayResponseRow `json:"rows"` // 报告表格。
}

// 商品价格（已包含促销活动或推广优惠）。
type MoneyMoney struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 竞争对手商品价格。
type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexData struct {
	MinPrice MoneyMoney `json:"min_price"`
	PriceIndex float64 `json:"price_index"` // 价格指数。
	URL string `json:"url"` // 竞争对手商品链接。
}

// 您商品的最低价格。
type MoneyMoneySelf struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 您的商品价格。
type ProductV1ProductPricesDetailsResponsePricePriceIndexIndexDataSelf struct {
	MinPrice MoneyMoneySelf `json:"min_price"`
	PriceIndex float64 `json:"price_index"` // 价格指数。
	URL string `json:"url"` // 您的商品链接。
}

type ProductV1ProductPricesDetailsResponsePricePriceIndex struct {
	SelfIndexData ProductV1ProductPricesDetailsResponsePricePriceIndexIndexDataSelf `json:"self_index_data"`
	ExternalIndexData ProductV1ProductPricesDetailsResponsePricePriceIndexIndexData `json:"external_index_data"`
}

// 网站上的商品价格。
type MoneyMoneyCustomerPrice struct {
	Amount string `json:"amount"` // 金额。
	Currency string `json:"currency"` // 货币单位。
}

type V1ProductPricesDetailsResponsePrice struct {
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符——SKU。
	CustomerPrice MoneyMoneyCustomerPrice `json:"customer_price"`
	DiscountPercent float64 `json:"discount_percent"` // 由 Ozon 承担的折扣比例。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符（商品货号）。
	Price MoneyMoney `json:"price"`
	PriceIndexes []ProductV1ProductPricesDetailsResponsePricePriceIndex `json:"price_indexes"` // 价格指数。
}

type AnalyticsProductQueriesDetailsResponseQuery struct {
	Gmv float64 `json:"gmv"` // 搜索查询的销售额。
	Position float64 `json:"position"` // 商品的平均排名。仅适用于[Premium](https://seller-edu.ozon.ru/seller-rating/about-rating/premium-program) 或 [Premium Plus](https://se...
	Query string `json:"query"` // 请求文本。
	QueryIndex int64 `json:"query_index"` // 查询序号。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	UniqueSearchUsers int64 `json:"unique_search_users"` // 在 Ozon 平台上搜索该商品的买家数量。
	UniqueViewUsers int64 `json:"unique_view_users"` // 在 Ozon 平台上看到该商品的买家数量。仅适用于[Premium](https://seller-edu.ozon.ru/seller-rating/about-rating/premium-program) 或 [Premium Plu...
	OrderCount int64 `json:"order_count"` // 根据查询的订单数量。
	ViewConversion float64 `json:"view_conversion"` // 商品的转化率。 仅适用于[Premium](https://seller-edu.ozon.ru/seller-rating/about-rating/premium-program) 或 [Premium Plus](https://se...
	Currency string `json:"currency"` // 货币单位。
}

// 数据分析的时间范围。
type V1AnalyticsProductQueriesDetailsResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"` // 分析数据的起始日期。
	DateTo string `json:"date_to"` // 分析数据的结束日期。
}

type V1AnalyticsProductQueriesDetailsResponse struct {
	AnalyticsPeriod V1AnalyticsProductQueriesDetailsResponseAnalyticsPeriod `json:"analytics_period"`
	PageCount int64 `json:"page_count"` // 总页数。
	Queries []AnalyticsProductQueriesDetailsResponseQuery `json:"queries"` // 查询列表。
	Total int64 `json:"total"` // 搜索请求的总数。
}

// 数据分析的时间范围。
type AnalyticsProductQueriesResponseAnalyticsPeriod struct {
	DateFrom string `json:"date_from"` // 分析数据的起始日期。
	DateTo string `json:"date_to"` // 分析数据的结束日期。
}

type AnalyticsMetric string

type AnalyticsDataRowDimension struct {
	ID string `json:"id"` // Ozon系统中的商品识别码是SKU。
	Name string `json:"name"` // 命名。
}

type AnalyticsDataRow struct {
	Dimensions []AnalyticsDataRowDimension `json:"dimensions"` // 报告数据分组。
	Metrics []float64 `json:"metrics"` // 指标值列表。
}

// 排序搜索查询的参数： - `CLIENT_COUNT`——查询的受欢迎程度； - `ADD_TO_CART`——添加到购物车的次数； - `CONVERSION_TO_CART`——购物车转化率； - `AVG_PRICE`——平均价格。
type SearchQueriesTextRequestSortBy string

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type SearchQueriesTextRequestSortDir string

type V1SearchQueriesTextRequest struct {
	Text string `json:"text"` // 按文本搜索。
	Limit string `json:"limit"` // 每页的结果数量。
	Offset string `json:"offset"` // 响应中将被跳过的项目数量。
	SortBy SearchQueriesTextRequestSortBy `json:"sort_by"`
	SortDir SearchQueriesTextRequestSortDir `json:"sort_dir"`
}

// 按具体参数对商品进行排序。可能的取值： - `BY_SEARCHES`— 按搜索次数； - `BY_VIEWS`— 按浏览量； - `BY_POSITION`— 按商品的平均排名； - `BY_CONVERSION`— 按转化率； - `B...
type V1AnalyticsProductQueriesDetailsRequestSortBy string

// 排序方向： - `DESCENDING`— 降序； - `ASCENDING`— 升序。
type V1AnalyticsProductQueriesDetailsRequestSortDir string

type V1AnalyticsProductQueriesDetailsRequest struct {
	DateFrom string `json:"date_from"` // 分析数据的起始日期。
	DateTo string `json:"date_to"` // 分析数据的结束日期。
	LimitBySKU int32 `json:"limit_by_sku"` // 单个SKU的查询数量限制。最大值为15次查询。
	Page int32 `json:"page"` // 请求返回的页码。最小值为0。
	PageSize int32 `json:"page_size"` // 每页包含的商品数量。最大值为100。
	Skus []string `json:"skus"` // SKU 列表，即 Ozon 系统中的商品标识符。根据这些 SKU 返回搜索查询的分析数据。最多可查询 1000 个 SKU。
	SortBy V1AnalyticsProductQueriesDetailsRequestSortBy `json:"sort_by"`
	SortDir V1AnalyticsProductQueriesDetailsRequestSortDir `json:"sort_dir"`
}

// 分类类型: - `ASC` — 升序， - `DESC` — 降序。
type SortingOrder string

type AnalyticsSorting struct {
	Key string `json:"key"` // 查询排序结果所依据的指标。
	Order SortingOrder `json:"order"`
}

type SellerServiceanalyticsDimension string

type FilterOp any

type AnalyticsFilter struct {
	Key string `json:"key"` // 排序参数。 可以传递`dimension` 和 `metric`中的任何属性, `brand`除外。
	Op FilterOp `json:"op"`
	Value string `json:"value"` // 用于对比的值。
}

type AnalyticsAnalyticsGetDataRequest struct {
	Metrics []AnalyticsMetric `json:"metrics"` // 最多指定14个指标。如有更多，您将收到 `InvalidArgument`的错误。 生成报告所依据的指标列表。 所有卖家可用的指标： - `revenue` — 订购的金额， - `ordered_units` — 订购的商品。 仅对Pre...
	Offset int64 `json:"offset"` // 响应中要跳过的元素数字。例如，如果 `offset = 10`, 那么答案将从找到的第11个元素开始。
	Sort []AnalyticsSorting `json:"sort"` // 报告排列设置。
	DateFrom string `json:"date_from"` // 数据将出现在报告中的日期。 若您没有Premium订阅，请指定过去三个月内的日期。
	DateTo string `json:"date_to"` // 数据将出现在报告中的截止日期。
	Dimension []SellerServiceanalyticsDimension `json:"dimension"` // 报告中的分组数据。 所有卖家可用的分组方法： - `unknownDimension` — 未知商品标识符； - `sku` — 商品标识符； - `spu` — 商品标识符 — 统一商品卡片； - `day` — 日； - `week` ...
	Filters []AnalyticsFilter `json:"filters"` // 过滤器。
	Limit int64 `json:"limit"` // 响应的值个数： - 最大值 — 1000， - 最小值 — 1.
}

type V1ProductPricesDetailsResponse struct {
	Prices []V1ProductPricesDetailsResponsePrice `json:"prices"` // 商品价格。
}

type V1ProductPricesDetailsRequest struct {
	Skus []string `json:"skus"` // SKU列表。
}

// 按具体参数对商品进行排序。可能的取值： - `BY_SEARCHES`— 按搜索次数； - `BY_VIEWS`— 按浏览量； - `BY_POSITION`— 按商品的平均排名； - `BY_CONVERSION`— 按转化率； - `B...
type AnalyticsProductQueriesRequestSortBy string

type V1AnalyticsProductQueriesResponseItem struct {
	Category string `json:"category"` // 类目名称。
	Currency string `json:"currency"` // 货币单位。
	Position float64 `json:"position"` // 商品的平均排名。仅适用于[Premium](https://seller-edu.ozon.ru/seller-rating/about-rating/premium-program) 或 [Premium Plus](https://se...
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	ViewConversion float64 `json:"view_conversion"` // 商品的转化率。 仅适用于[Premium](https://seller-edu.ozon.ru/seller-rating/about-rating/premium-program) 或 [Premium Plus](https://se...
	Gmv float64 `json:"gmv"` // 搜索查询的销售额。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符（商品编号）。
	UniqueSearchUsers int64 `json:"unique_search_users"` // 在 Ozon 平台上搜索该商品的买家数量。
	UniqueViewUsers int64 `json:"unique_view_users"` // 在 Ozon 平台上看到该商品的买家数量。仅适用于[Premium](https://seller-edu.ozon.ru/seller-rating/about-rating/premium-program) 或 [Premium Plu...
}

type V1AnalyticsProductQueriesResponse struct {
	AnalyticsPeriod AnalyticsProductQueriesResponseAnalyticsPeriod `json:"analytics_period"`
	Items []V1AnalyticsProductQueriesResponseItem `json:"items"` // 商品列表。
	PageCount int64 `json:"page_count"` // 总页数。
	Total int64 `json:"total"` // 搜索请求的总数。
}

// 查询结果。
type AnalyticsGetDataResponseResult struct {
	Data []AnalyticsDataRow `json:"data"` // 数据组。
	Totals []float64 `json:"totals"` // 指标总计和平均值。
}

type AnalyticsAnalyticsGetDataResponse struct {
	Timestamp string `json:"timestamp"` // 报告创建时间。
	Result AnalyticsGetDataResponseResult `json:"result"`
}

// 排序方向： - `DESCENDING`— 降序； - `ASCENDING`— 升序。
type AnalyticsProductQueriesRequestSortDir string

type V1AnalyticsProductQueriesRequest struct {
	Skus []string `json:"skus"` // SKU 列表，即 Ozon 系统中的商品标识符。根据这些 SKU 返回搜索查询的分析数据。最多可查询 1000 个 SKU。
	SortBy AnalyticsProductQueriesRequestSortBy `json:"sort_by"`
	SortDir AnalyticsProductQueriesRequestSortDir `json:"sort_dir"`
	DateFrom string `json:"date_from"` // 分析数据的起始日期。
	DateTo string `json:"date_to"` // 分析数据的结束日期。
	Page int32 `json:"page"` // 请求返回的页码。
	PageSize int32 `json:"page_size"` // 每页包含的商品数量。
}

type V1GetRealizationReportByDayRequest struct {
	Day int32 `json:"day"` // 日。
	Month int32 `json:"month"` // 月。
	Year int32 `json:"year"` // 年。
}

type V1SearchQueriesTopRequest struct {
	Limit string `json:"limit"` // 每页的结果数量。
	Offset string `json:"offset"` // 响应中将被跳过的项目数量。
}

type V1SearchQueriesTopResponseSearchQuery struct {
	AvgPrice float64 `json:"avg_price"` // 商品的平均价格（以卢布计）。
	ClientCount float64 `json:"client_count"` // 通过该,搜索查询 查找商品的买家数量。
	ConversionToCart float64 `json:"conversion_to_cart"` // 将至少 1 件商品添加到购物车的买家比例。
	ItemsViews float64 `json:"items_views"` // 商品的浏览次数。
	Query string `json:"query"` // 搜索查询。
	SellersCount float64 `json:"sellers_count"` // 买家根据该搜索查询查看其商品的平均卖家数量。
	AddToCart float64 `json:"add_to_cart"` // 将至少 1 件商品添加到购物车的买家数量。
}

type V1SearchQueriesTopResponse struct {
	Total string `json:"total"` // 搜索查询总数。
	Offset string `json:"offset"` // 每页显示的搜索查询数量。
	SearchQueries []V1SearchQueriesTopResponseSearchQuery `json:"search_queries"` // 搜索查询信息。
}

package product

type V1ProductGetRelatedSKUResponseItem struct {
	SKU int64 `json:"sku"` // Ozon系统中的商品ID —— SKU。
	Availability string `json:"availability"` // SKU商品的可用性表示： - `HIDDEN` —— 隐藏； - `AVAILABLE` —— 可用； - `UNAVAILABLE` —— 不可用，SKU已删除。
	DeletedAt string `json:"deleted_at"` // 删除日期和时间。
	DeliverySchema string `json:"delivery_schema"` // 配送计划： - `SDS` — Ozon 通用 SKU 编号； - `FBO` — 来自 Ozon 仓库的商品编号； - `FBS` — 来自 FBS 仓库的商品编号； - `Crossborder` — 跨境销售商品编号。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

type V1ProductGetRelatedSKUResponseError struct {
	Code string `json:"code"` // 代码有误。
	SKU int64 `json:"sku"` // 有错误的SKU。
	Message string `json:"message"` // 文本错误。
}

type V1ProductGetRelatedSKUResponse struct {
	Items []V1ProductGetRelatedSKUResponseItem `json:"items"` // 相关SKU信息。
	Errors []V1ProductGetRelatedSKUResponseError `json:"errors"` // 错误。
}

type V1ItemError struct {
	Code string `json:"code"` // 错误代码。
	Message string `json:"message"` // 技术描述错误。
	State string `json:"state"` // 发生错误的商品状态。
	Level string `json:"level"` // 错误水平。
	Description string `json:"description"` // 错误描述。
	Field string `json:"field"` // 发生错误的字段。
	AttributeID int64 `json:"attribute_id"` // 发生错误的属性。
	AttributeName string `json:"attribute_name"` // 发生错误的属性名称。
}

type GetImportProductsInfoResponseResultItem struct {
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	Status string `json:"status"` // 商品创建或更新状态。 商品信息在队列中处理。 可能的参数值: - `pending` — 商品等待排队处理; - `imported` — 商品已成功加载； - `failed` — 商品加载错误； - `skipped` — 商品未更新，...
	Errors []V1ItemError `json:"errors"` // 错误数组。
	OfferID string `json:"offer_id"` // 在卖方系统中的商品ID是货号。字段数值中行的最大长度为50个字符。
}

type DeleteProductsRequestProduct struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
}

type Productv2DeleteProductsRequest struct {
	Products []DeleteProductsRequestProduct `json:"products"` // Ozon系统中商品的标识符 — `product_id`。
}

type ProductProductUnarchiveRequest struct {
	ProductID []int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。您一次最多可以传递100个标识符。 在一天内，您最多可以从档案中恢复100件自动归档的商品。 限额在莫斯科时间03：00更新。 手动归档的商品没有解除归档的限制。
}

type GetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueId int64 `json:"dictionaryValueId"` // 词典中特征的识别码。
	Value string `json:"value"` // 一个商品特性的价值。
}

type GetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"` // 特征的识别码。
	ComplexID int64 `json:"complex_id"` // 支持嵌入属性的特征的识别码。例如，"处理器" 有嵌入属性的特征 "制造商" 和 "二级缓存"。每个嵌入属性都可以有几个值选项。
	Values []GetProductAttributesV3ResponseDictionaryValue `json:"values"` // 特征值的数组。
}

type ProductGetProductAttributesV3ResponseComplexAttribute struct {
	Attributes []GetProductAttributesV3ResponseAttribute `json:"attributes"` // 商品特征的数组。
}

type GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"` // 到PDF文件的路径。
	Index int64 `json:"index"` // 存储库中的设定了顺序的文件指数。
	Name string `json:"name"` // 文件名称。
}

type GetProductAttributesResponseImage struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
	Default bool `json:"default"`
}

type ProductGetProductAttributesV3ResponseDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"` // 词典中该特征的识别码。
	Value string `json:"value"` // 商品特性的意义。
}

type ProductGetProductAttributesV3ResponseAttribute struct {
	AttributeID int64 `json:"attribute_id"` // 特征的识别码。
	ComplexID int64 `json:"complex_id"` // 支持嵌入属性的特征的识别码。例如，"处理器" 有嵌入属性的特征 "制造商" 和 "二级缓存"。每个嵌入属性都可以有几个值选项。
	Values []ProductGetProductAttributesV3ResponseDictionaryValue `json:"values"` // 特征值的数组。
}

type GetProductAttributesResponseImage360 struct {
	FileName string `json:"file_name"`
	Index int64 `json:"index"`
}

type Productv3GetProductAttributesV3ResponseResult struct {
	Barcode string `json:"barcode"` // 条形码。
	ColorImage string `json:"color_image"` // 市场营销色彩。
	ID int64 `json:"id"` // 商品特性的识别码。
	Images []GetProductAttributesResponseImage `json:"images"` // 产品图片链接的数组。
	Weight int32 `json:"weight"` // 商品在包装中的重量。
	WeightUnit string `json:"weight_unit"` // 重量测量单位。
	TypeID int64 `json:"type_id"` // 商品类型的标识符。
	Width int32 `json:"width"` // 包装宽度。
	Attributes []ProductGetProductAttributesV3ResponseAttribute `json:"attributes"` // 商品特性的数组。
	DescriptionCategoryID int64 `json:"description_category_id"` // 类别ID。 请将其与以下方法结合使用：[/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes) и [/v1/descript...
	DimensionUnit string `json:"dimension_unit"` // 尺寸的测量单位。 - `mm` —— 毫米， - `cm` —— 厘米， - `in` —— 英寸。
	Height int32 `json:"height"` // 包装高度。
	ImageGroupID string `json:"image_group_id"` // 用于后续批量下载图像的识别码。
	Images360 []GetProductAttributesResponseImage360 `json:"images360"` // 图像数组360。
	Name string `json:"name"` // 商品名称。最多500个字符。
	ComplexAttributes []ProductGetProductAttributesV3ResponseComplexAttribute `json:"complex_attributes"` // 已录入的特性的数组。
	PDFList []GetProductAttributesResponsePdf `json:"pdf_list"` // PDF文件的阵列。
	CategoryID int64 `json:"category_id"` // 类别ID。 请将其与以下方法结合使用：[/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree), [/v1/description-category...
	Depth int32 `json:"depth"` // 深度。
	OfferID string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
}

type V1ProductUpdateAttributesRequestValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"` // 字典中的特征ID。
	Value string `json:"value"` // 商品特征值。
}

type V1ProductUpdateAttributesRequestAttribute struct {
	ComplexID int64 `json:"complex_id"` // 支持嵌套属性的特征ID。 每个嵌套特征都可以有多个版本的值。
	ID int64 `json:"id"` // 特征ID。
	Values []V1ProductUpdateAttributesRequestValue `json:"values"` // 一组嵌套的特征值。
}

type Productv1ProductImportPicturesRequest struct {
	Images []string `json:"images"` // 数组图片链接。 最多30件。 数组中的图像是按照它们在网站上出现的顺序排列的。 数组中的第一个图像将是主图像。
	Images360 []string `json:"images360"` // 360图片的数组。多达70张图片。 格式：公共云存储中的图像链接地址。链接图片的格式是JPG。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	ColorImage string `json:"color_image"` // 市场营销色彩。
}

// 错误级别描述: - `ERROR_LEVEL_UNSPECIFIED` — 未定义; - `ERROR_LEVEL_ERROR` — 严重错误，无法销售该商品； - `ERROR_LEVEL_INTERNAL` — 严重错误，无法销售该商品...
type ErrorErrorLevel string

// 包装类型: - `SHIPMENT_TYPE_UNSPECIFIED` — 未指定; - `SHIPMENT_TYPE_GENERAL` — 普通商品; - `SHIPMENT_TYPE_BOX` — 箱子; - `SHIPMENT_TYP...
type SourceShipmentType string

// 按商品曝光度过滤: - `ALL` —— 所有的商品，除了已存档的。 - `VISIBLE` —— 买家可见的商品。 - `INVISIBLE` —— 买家看不到的商品。 - `EMPTY_STOCK`—— 没有现货的商品。 - `NOT_...
type Productv3GetProductListRequestFilterFilterVisibility string

// 商品过滤。
type Productv3GetProductListRequestFilter struct {
	ProductID []string `json:"product_id"` // 基于参数 `product_id` 的过滤。 可以提交数值列表。
	Visibility Productv3GetProductListRequestFilterFilterVisibility `json:"visibility"`
	OfferID []string `json:"offer_id"` // 基于参数 `offer_id` 的过滤。 可以提交数值列表。
}

type Productv3GetProductListRequest struct {
	Limit int64 `json:"limit"` // 答复的信息数量。默认设置为1。最大值是1000。
	Filter Productv3GetProductListRequestFilter `json:"filter"`
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定 `last_id`。
}

type V1GetProductInfoSubscriptionResponseResult struct {
	Count int64 `json:"count"` // 订阅用户的数量。
	SKU int64 `json:"sku"` // Ozon 系统中的商品ID、SKU。
}

type DeleteProductsResponseDeleteStatus struct {
	IsDeleted bool `json:"is_deleted"` // 如果查询的执行没有错误且商品被删除 —— `true`。
	OfferID string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	Error string `json:"error"` // 处理该请求时发生错误的原因。
}

type V4GetProductAttributesResponsePdf struct {
	FileName string `json:"file_name"` // 到PDF文件的路径。
	Name string `json:"name"` // 文件名称。
}

type ProductGetProductAttributesV4ResponseAttribute struct {
	Values []ProductGetProductAttributesV3ResponseDictionaryValue `json:"values"` // 特征值数组。
	ID int64 `json:"id"` // 特征标识符。
	ComplexID int64 `json:"complex_id"` // 支持嵌入属性的特征的识别码。例如，"处理器" 有嵌入属性的特征 "制造商" 和 "二级缓存"。每个嵌入属性都可以有几个值选项。
}

// 型号信息。
type V4GetProductAttributesResponseModelInfo struct {
	ModelID int64 `json:"model_id"` // 型号标识符。
	Count int64 `json:"count"` // 该型号下合并商品的数量。
}

type GetProductAttributesV4ResponseAttribute struct {
	Values []GetProductAttributesV3ResponseDictionaryValue `json:"values"` // 特征值数组。
	ID int64 `json:"id"` // 特征标识符。
	ComplexID int64 `json:"complex_id"` // 支持嵌入属性的特征的识别码。例如，"处理器" 有嵌入属性的特征 "制造商" 和 "二级缓存"。每个嵌入属性都可以有几个值选项。
}

type Productv4GetProductAttributesV4ResponseResult struct {
	AttributesWithDefaults []int64 `json:"attributes_with_defaults"` // 具有默认值的特征标识符列表。
	Barcodes any `json:"barcodes"` // 商品的所有条形码。
	Images []GetProductAttributesResponseImage `json:"images"` // 商品图片链接数组。图片顺序与商品卡片中的顺序一致。
	ModelInfo V4GetProductAttributesResponseModelInfo `json:"model_info"`
	PrimaryImage string `json:"primary_image"` // 商品主图链接。
	TypeID int64 `json:"type_id"` // 商品类型的标识符。
	Barcode string `json:"barcode"` // 条形码。
	DescriptionCategoryID int64 `json:"description_category_id"` // 类目标识符。 请将其与以下方法结合使用：[/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes) 和 [/v1/descrip...
	Height int64 `json:"height"` // 包装高度。
	SKU string `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	WeightUnit string `json:"weight_unit"` // 重量测量单位。
	ColorImage string `json:"color_image"` // 市场营销色彩。
	ComplexAttributes []GetProductAttributesV4ResponseAttribute `json:"complex_attributes"` // 嵌套特征列表。
	Depth int64 `json:"depth"` // 深度。
	DimensionUnit string `json:"dimension_unit"` // 尺寸的测量单位。 - `mm` —— 毫米， - `cm` —— 厘米， - `in` —— 英寸。
	ID int64 `json:"id"` // Ozon系统中商品的标识符 — `product_id`。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符 — 货号。
	PDFList []V4GetProductAttributesResponsePdf `json:"pdf_list"` // PDF文件列表。
	Weight int64 `json:"weight"` // 商品在包装中的重量。
	Width int64 `json:"width"` // 包装宽度。
	Attributes []ProductGetProductAttributesV4ResponseAttribute `json:"attributes"` // 商品特性的数组。
}

type Productv4GetProductAttributesV4Response struct {
	Total string `json:"total"` // 列表中的商品数量。
	Result []Productv4GetProductAttributesV4ResponseResult `json:"result"` // 查询结果。
	LastID string `json:"last_id"` // 页面上最后一个值的ID。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
}

// Ozon上竞争对手的商品价格。
type PriceIndexesIndexDataOzon struct {
	MinimalPrice string `json:"minimal_price"` // Ozon上竞争对手商品的最低价格。
	MinimalPriceCurrency string `json:"minimal_price_currency"` // 货币价值。
	PriceIndexValue float64 `json:"price_index_value"` // 价格指数的值。
}

type ProductGetProductInfoDescriptionRequest struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

// 商品更新的每日限制。
type GetUploadQuotaResponseDailyUpdate struct {
	Limit int64 `json:"limit"` // 每天可以更新多少商品。若取值为 `-1`，表示无限制。
	ResetAt string `json:"reset_at"` // 当前天数计数器值将被重置时，采用 UTC 格式的时间。
	Usage int64 `json:"usage"` // 当前几天更新了多少商品。
}

// 当前限额： - `UNSPECIFIED`: 未指定； - `RATE_LIMIT_PER_MINUTE`: 每分钟限额。
type ProductV4GetUploadQuotaResponseOperationLimitsLimitTypeEnum string

// 商品操作的可用限额。
type ProductV4GetUploadQuotaResponseOperationLimits struct {
	Limit int64 `json:"limit"` // 每分钟可以创建多少商品。
	LimitType ProductV4GetUploadQuotaResponseOperationLimitsLimitTypeEnum `json:"limit_type"`
}

// 特定类目商品的创建限额。
type ProductV4GetUploadQuotaResponseTotalQuotaByCategory struct {
	Usage int64 `json:"usage"` // 已创建的商品数量。
	CategoryID int64 `json:"category_id"` // 限额适用类目的标识符。
	Limit int64 `json:"limit"` // 可在个人中心中创建的商品数量。
}

// 品类限制。
type GetUploadQuotaResponseTotal struct {
	Limit int64 `json:"limit"` // 在个人中心一共可以创建多少商品。若取值为 `-1`，表示无限制。
	Usage int64 `json:"usage"` // 多少商品已被创建。
	QuotaByCategory ProductV4GetUploadQuotaResponseTotalQuotaByCategory `json:"quota_by_category"`
}

// 创建商品的每日限制。
type GetUploadQuotaResponseDailyCreate struct {
	Limit int64 `json:"limit"` // 一天可以创建多少商品。若取值为 `-1`，表示无限制。
	ResetAt string `json:"reset_at"` // 当前天数计数器值将被重置时，采用 UTC 格式的时间。
	Usage int64 `json:"usage"` // 近几天创建了多少商品。
}

type V4GetUploadQuotaResponse struct {
	DailyCreate GetUploadQuotaResponseDailyCreate `json:"daily_create"`
	DailyUpdate GetUploadQuotaResponseDailyUpdate `json:"daily_update"`
	OperationLimits ProductV4GetUploadQuotaResponseOperationLimits `json:"operation_limits"`
	Total GetUploadQuotaResponseTotal `json:"total"`
}

type ProductImportProductsBySKURequestItem struct {
	Vat string `json:"vat"` // 商品增值税税率： - `0` — 免除增值税, - `0.05` — 5%, - `0.07` — 7%, - `0.1` — 10%, - `0.2` — 20%, - `0.22` — 22%。 传递当前有效的出价值。
	CurrencyCode string `json:"currency_code"` // 价格显示的货币。 传输值必须与个人中心设置中所设置的货币相匹配。 默认情况下，显示`RUB` — 俄罗斯卢布。 例如，如果您设人民币为结算货币，请赋值`CNY`，否则将返回错误。 可能的值： - `RUB` — 俄罗斯卢布, - `BYN`...
	Name string `json:"name"` // 商品名称。 小于500字符。
	OfferID string `json:"offer_id"` // 在卖家系统中的商品ID — 货号。至50字符。
	OldPrice string `json:"old_price"` // 折扣前的价格（将在商品卡片上划掉）。 以卢布表示。 小数点分隔符是一个点，在点之后最多两个字符。
	Price string `json:"price"` // 含折扣的商品价格显示在商品卡片上。 如果商品没有折扣，请在此参数中指定值`old_price`。
	SKU int64 `json:"sku"` // 在Ozon系统中的商品ID — SKU。
}

type ProductImportProductsBySKURequest struct {
	Items []ProductImportProductsBySKURequestItem `json:"items"` // 商品信息。
}

type Productv3GetProductListResponseItemQuant struct {
	QuantCode string `json:"quant_code"` // 经济商品标识符。
	QuantSize int64 `json:"quant_size"` // 定量包装大小。
}

type Productv3GetProductListResponseItem struct {
	IsDiscounted bool `json:"is_discounted"` // 减价商品。
	OfferID string `json:"offer_id"` // 经济商品标识符。
	ProductID int64 `json:"product_id"` // 商品识别号.
	Quants []Productv3GetProductListResponseItemQuant `json:"quants"` // 量子清单.
	Archived bool `json:"archived"` // 商品已归档。
	HasFBOStocks bool `json:"has_fbo_stocks"` // FBO 仓库有库存。
	HasFBSStocks bool `json:"has_fbs_stocks"` // FBS 仓库有库存。
}

// 结果。
type Productv3GetProductListResponseResult struct {
	Items []Productv3GetProductListResponseItem `json:"items"` // 品列表的。
	LastID string `json:"last_id"` // 页面上最后一个值的ID。 要检索以下数值，请从上一个查询的响应中指定 `last_id`。
	Total int32 `json:"total"` // 品牌总数。
}

type Productv3GetProductListResponse struct {
	Result Productv3GetProductListResponseResult `json:"result"`
}

type ProductInfoWrongVolumeResponseWrongVolumeProduct struct {
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符 —— SKU。
	Weight int64 `json:"weight"` // 商品包装重量。
	Width int64 `json:"width"` // 商品宽度。
	Height int64 `json:"height"` // 商品高度。
	Length int64 `json:"length"` // 商品长度。
	Name string `json:"name"` // 商品名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符 —— 货号。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

// 价格指数类型: - `COLOR_INDEX_UNSPECIFIED` — 未指定; - `COLOR_INDEX_WITHOUT_INDEX` — 无价格指数; - `COLOR_INDEX_SUPER` — 超级有利, - `COLOR...
type PriceIndexesColorIndex string

// 商品的可见性设置。
type GetProductInfoListResponseVisibilityDetails struct {
	HasPrice bool `json:"has_price"` // 如果设置了价格，则为 `true`。
	HasStock bool `json:"has_stock"` // 如果仓库有库存，则为 `true`。
}

// 原因信息。
type ReasonHumanText struct {
	Text string `json:"text"` // 原因描述。
}

type AvailabilityReason struct {
	ID int64 `json:"id"` // 原因标识符。
	HumanText ReasonHumanText `json:"human_text"`
}

type GetProductInfoListResponseAvailability struct {
	Availability string `json:"availability"` // 商品可售数量。
	Reasons []AvailabilityReason `json:"reasons"` // 商品被隐藏的原因。
	SKU int64 `json:"sku"` // 商品标识符。
	Source string `json:"source"` // 来源链接。
}

// 商品型号信息。
type GetProductInfoListResponseModelInfo struct {
	Count int64 `json:"count"` // 响应中的商品数量。
	ModelID int64 `json:"model_id"` // 商品型号标识符。
}

// 其他平台上竞争对手的商品价格。
type PriceIndexesIndexDataExternal struct {
	PriceIndexValue float64 `json:"price_index_value"` // 价格指数的值。
	MinimalPrice string `json:"minimal_price"` // 其他平台上竞争对手的最低商品价格。
	MinimalPriceCurrency string `json:"minimal_price_currency"` // 货币价格。
}

// 您的商品在其他平台的价格。
type PriceIndexesIndexDataSelf struct {
	MinimalPrice string `json:"minimal_price"` // 您的商品在其他网站上的最低价格。
	MinimalPriceCurrency string `json:"minimal_price_currency"` // 货币价格。
	PriceIndexValue float64 `json:"price_index_value"` // 价格指数的值。
}

// 商品价格指数。
type GetProductInfoListResponsePriceIndexes struct {
	ColorIndex PriceIndexesColorIndex `json:"color_index"`
	ExternalIndexData PriceIndexesIndexDataExternal `json:"external_index_data"`
	OzonIndexData PriceIndexesIndexDataOzon `json:"ozon_index_data"`
	SelfMarketplacesIndexData PriceIndexesIndexDataSelf `json:"self_marketplaces_index_data"`
}

// 商品状态信息。
type GetProductInfoListResponseStatuses struct {
	StatusDescription string `json:"status_description"` // 商品状态描述。
	StatusName string `json:"status_name"` // 商品状态名称。
	StatusTooltip string `json:"status_tooltip"` // 状态描述。
	StatusUpdatedAt string `json:"status_updated_at"` // 状态最后变更时间。
	ValidationStatus string `json:"validation_status"` // 验证状态。
	IsCreated bool `json:"is_created"` // 如果商品创建正确，则为 `true`。
	ModerateStatus string `json:"moderate_status"` // 审核状态。
	Status string `json:"status"` // 商品状态。
	StatusFailed string `json:"status_failed"` // 发生错误的商品状态。
}

type GetProductInfoListResponseCommission struct {
	DeliveryAmount float64 `json:"delivery_amount"` // 配送费用。
	Percent float64 `json:"percent"` // 佣金比例。
	ReturnAmount float64 `json:"return_amount"` // 退货费用。
	SaleSchema string `json:"sale_schema"` // 销售模式。
	Value float64 `json:"value"` // 佣金总额。
}

// 促销活动类型： - `UNSPECIFIED` — 未知类型; - `REVIEWS_PROMO` — “快速收集评价”促销活动。
type V3GetProductInfoListResponsePromotionType string

type V3GetProductInfoListResponsePromotion struct {
	IsEnabled bool `json:"is_enabled"` // 如果促销活动已启用，则为 `true`。
	Type V3GetProductInfoListResponsePromotionType `json:"type_"`
}

type HumanTextsParam struct {
	Name string `json:"name"` // 参数名称。
	Value string `json:"value"` // 参数值。
}

// 错误描述。
type ErrorHumanTexts struct {
	ShortDescription string `json:"short_description"` // 错误简要描述。
	AttributeName string `json:"attribute_name"` // 发生错误的属性名称。
	Description string `json:"description"` // 错误详情。
	HintCode string `json:"hint_code"` // Ozon系统中的错误代码。
	Message string `json:"message"` // 错误文本。
	Params []HumanTextsParam `json:"params"` // 发生错误的参数。
}

type GetProductInfoListResponseError struct {
	AttributeID int64 `json:"attribute_id"` // 属性标识符。
	Code string `json:"code"` // 错误代码。
	Field string `json:"field"` // 出现错误的字段。
	Level ErrorErrorLevel `json:"level"`
	State string `json:"state"` // 发生错误的商品状态。
	Texts ErrorHumanTexts `json:"texts"`
}

type GetProductInfoListResponseStocksStock struct {
	Present int32 `json:"present"` // 当前库存数量。
	Reserved int32 `json:"reserved"` // 预定商品数量。
	SKU int64 `json:"sku"` // Ozon系统中的商品标识符 — SKU。
	Source string `json:"source"` // 销售模式。
}

// 商品库存信息。
type GetProductInfoListResponseStocks struct {
	HasStock bool `json:"has_stock"` // 如果库存中有剩余，则为 `true`。
	Stocks []GetProductInfoListResponseStocksStock `json:"stocks"` // 库存状态。
}

type GetProductInfoListResponseSource struct {
	CreatedAt string `json:"created_at"` // 商品创建日期。
	QuantCode string `json:"quant_code"` // 包含商品的定量包装列表。
	ShipmentType SourceShipmentType `json:"shipment_type"`
	SKU int64 `json:"sku"` // Ozon平台上的商品标识符 — SKU。
	Source string `json:"source"` // 销售模式： - `SDS` — 适用于 FBO 和 FBS，且 SKU 相同； - `FBO`; - `FBS`。
}

type V3GetProductInfoListResponseItem struct {
	SKU int64 `json:"sku"` // 商品在 Ozon 系统中的标识符 — SKU。
	CreatedAt string `json:"created_at"` // 商品的创建日期和时间。
	MinPrice string `json:"min_price"` // 应用促销后的最低价格。
	Availabilities []GetProductInfoListResponseAvailability `json:"availabilities"` // 商品可售数量信息。
	HasDiscountedFBOItem bool `json:"has_discounted_fbo_item"` // 商品在 Ozon 仓库中是否有减价版的同款商品。
	IsAutoarchived bool `json:"is_autoarchived"` // 如果商品是自动归档的，则为 `true`。
	ModelInfo GetProductInfoListResponseModelInfo `json:"model_info"`
	Price string `json:"price"` // 商品的含折扣价格。该值将在商品卡片上显示。
	PriceIndexes GetProductInfoListResponsePriceIndexes `json:"price_indexes"`
	UpdatedAt string `json:"updated_at"` // 商品的最后更新时间。
	ColorImage []string `json:"color_image"` // 商品颜色图片。
	Images []string `json:"images"` // 图片链接数组。 图片在数组中的顺序与网站上的展示顺序一致。 如果 `primary_image` 参数未指定，则数组中的第一张图片为商品主图。
	IsKgt bool `json:"is_kgt"` // `true`，表示商品为超大。仅适用于FBS模式。
	Statuses GetProductInfoListResponseStatuses `json:"statuses"`
	Stocks GetProductInfoListResponseStocks `json:"stocks"`
	Vat string `json:"vat"` // 商品的增值税税率。
	IsArchived bool `json:"is_archived"` // 如果商品是手动归档的，则为 `true`。
	IsDiscounted bool `json:"is_discounted"` // 商品是否为减价商品： - 如果商品是由卖家作为减价商品创建的，则为 `true`。 - 如果商品不是减价商品或是由Ozon减价的，则为 `false`。
	IsSuper bool `json:"is_super"` // 超级商品标识。 [有关超级商品的详细信息，请参考卖家知识库](https://seller-edu.ozon.ru/fbo/rabota-so-stokom/super-tovary)
	OfferID string `json:"offer_id"` // 商品在卖家系统中的标识符 — 货号。
	Sources []GetProductInfoListResponseSource `json:"sources"` // 商品创建来源信息。自2023年7月1日起，卖家按SDS模式创建商品。
	Commissions []GetProductInfoListResponseCommission `json:"commissions"` // 佣金信息。
	Images360 []string `json:"images360"` // 360° 商品图片数组。
	DiscountedFBOStocks int32 `json:"discounted_fbo_stocks"` // Ozon 仓库中减价商品的库存。
	IsPrepaymentAllowed bool `json:"is_prepayment_allowed"` // 如果支持预付款，则为 `true`。
	Promotions []V3GetProductInfoListResponsePromotion `json:"promotions"` // 促销活动。
	TypeID int64 `json:"type_id"` // 商品类型标识符。
	VolumeWeight float64 `json:"volume_weight"` // 商品的体积重量。
	DescriptionCategoryID int64 `json:"description_category_id"` // 类目标识符。 可与 [/v1/description-category/attribute](#operation/DescriptionCategoryAPI_GetAttributes) 和 [/v1/description-categ...
	Name string `json:"name"` // 名称。
	OldPrice string `json:"old_price"` // 不含折扣价格。在商品卡片上显示为划线价。
	PrimaryImage []string `json:"primary_image"` // 商品的主图。
	VisibilityDetails GetProductInfoListResponseVisibilityDetails `json:"visibility_details"`
	Barcodes []string `json:"barcodes"` // 商品的所有条形码。
	CurrencyCode string `json:"currency_code"` // 货币单位。
	Errors []GetProductInfoListResponseError `json:"errors"` // 创建或验证商品时的错误信息。
	ID int64 `json:"id"` // 商品标识符。
}

type V2ProductInfoPicturesRequest struct {
	ProductID []string `json:"product_id"` // 商品识别码的清单。
}

type V3ImportProductsRequestDictionaryValue struct {
	DictionaryValueID int64 `json:"dictionary_value_id"` // 目录ID。
	Value string `json:"value"` // 目录中的值。
}

type V3ImportProductsRequestAttribute struct {
	ComplexID int64 `json:"complex_id"` // 支持嵌套属性的特征ID。 例如，"处理器"特征具有嵌套特征"制造商"，"L2缓存"等。 每个嵌套特征都可以具有多值变体。
	ID int64 `json:"id"` // 特征ID。
	Values []V3ImportProductsRequestDictionaryValue `json:"values"` // 嵌套特征值的数组。
}

type V3ImportProductsRequestComplexAttribute struct {
	Attributes []V3ImportProductsRequestAttribute `json:"attributes"` // 商品特性组。不同类别的特征不同 — 可见 [卖家知识库](https://global-help.ozon.com/zh/) 或通过 [API](https://docs.ozon.ru/api/seller/zh/)。
}

type Productv2DeleteProductsResponse struct {
	Status []DeleteProductsResponseDeleteStatus `json:"status"` // 请求的处理情况。
}

type V3GetProductInfoListResponse struct {
	Items []V3GetProductInfoListResponseItem `json:"items"` // 数据数组。
}

type ProductBooleanResponse struct {
	Result bool `json:"result"` // 查询的处理结果。`true`，如果查询的执行无误。
}

type ProductGetImportProductsInfoRequest struct {
	TaskID int64 `json:"task_id"` // 进口商品的任务代码。按照运输方式筛选。可以使用方法 [/v3/product/import](#operation/ProductAPI_ImportProductsV3)获取。
}

type GetProductRatingBySkuResponseRatingCondition struct {
	Fulfilled bool `json:"fulfilled"` // 表示满足条件。
	Key string `json:"key"` // 条件识别码。
	Cost float64 `json:"cost"` // 完成条件给出的内容评级积分的数量。
	Description string `json:"description"` // 条件描述。
}

type GetProductRatingBySkuResponseRatingImproveAttribute struct {
	Name string `json:"name"` // 属性名称。
	ID int64 `json:"id"` // 属性标识码。
}

type GetProductRatingBySkuResponseRatingGroup struct {
	Weight float64 `json:"weight"` // 小组特征对内容排名的影响比例。
	Conditions []GetProductRatingBySkuResponseRatingCondition `json:"conditions"` // 增加商品内容评级的条件清单。
	ImproveAtLeast int64 `json:"improve_at_least"` // 为获得该组特征的最高分，需要完成填写的属性数。
	ImproveAttributes []GetProductRatingBySkuResponseRatingImproveAttribute `json:"improve_attributes"` // 属性列表，完成这些属性可以增加商品的内容等级。
	Key string `json:"key"` // 小组的识别码。
	Name string `json:"name"` // 小组名称。
	Rating float64 `json:"rating"` // 小组中的排名。
}

type V1GetProductInfoSubscriptionRequest struct {
	Skus []string `json:"skus"` // Ozon 系统中的 SKU、商品标识符列表。
}

// 按商品曝光度过滤: - `ALL` —— 所有的商品，除了已存档的。 - `VISIBLE` —— 买家可见的商品。 - `INVISIBLE` —— 买家看不到的商品。 - `EMPTY_STOCK`—— 没有现货的商品。 - `NOT_...
type Productv2GetProductListRequestFilterFilterVisibility string

// 商品过滤。
type Productv3Filter struct {
	OfferID []string `json:"offer_id"` // 基于参数 `offer_id`的过滤。 可以提交数值列表。
	ProductID []string `json:"product_id"` // 基于参数 `product_id`的过滤。 可以提交数值列表。
	Visibility Productv2GetProductListRequestFilterFilterVisibility `json:"visibility"`
}

type V1ProductUpdateAttributesRequestItem struct {
	Attributes []V1ProductUpdateAttributesRequestAttribute `json:"attributes"` // 商品特征。
	OfferID string `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
}

type V1ProductUpdateAttributesRequest struct {
	Items []V1ProductUpdateAttributesRequestItem `json:"items"` // 需要更新的商品和特征。
}

type ProductGetProductInfoDescriptionResponseProduct struct {
	Description string `json:"description"` // 描述。
	ID int64 `json:"id"` // 识别码。
	Name string `json:"name"` // 名称。
	OfferID string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
}

type ProductGetProductInfoDescriptionResponse struct {
	Result ProductGetProductInfoDescriptionResponseProduct `json:"result"`
}

// 服务类型。 以大写形式传递其中的一个值: - `IS_CODE_SERVICE`, - `IS_NO_CODE_SERVICE`。
type V2ServiceType string

type ImportProductRequestPromotion struct {
	Type string `json:"type_"` // 促销活动类型： - `REVIEWS_PROMO` — “快速收集评价”促销活动。
	Operation string `json:"operation"` // 促销活动操作属性： - `ENABLE` — 启用， - `DISABLE` — 禁用， - `UNKNOWN` — 保持默认，不做修改。
}

type ImportProductsRequestPdfList struct {
	Index int64 `json:"index"` // 存储库中的文档索引，用于设置顺序。
	Name string `json:"name"` // 文件名称。
	SrcURL string `json:"src_url"` // 文件地址。
}

type V3ImportProductsRequestItem struct {
	PDFList []ImportProductsRequestPdfList `json:"pdf_list"` // PDF-文件清单。
	WeightUnit string `json:"weight_unit"` // 测重单位： - `g` — 克, - `kg` — 公斤, - `lb` — 磅。
	Attributes []V3ImportProductsRequestAttribute `json:"attributes"` // 商品特性组。 不同类别的特征不同 — 其可见 [卖家知识库](https://global-help.ozon.com/zh/) 或通过 [API](https://docs.ozon.ru/api/seller/zh/)。
	NewDescriptionCategoryID int64 `json:"new_description_category_id"` // 新的类目标识符。如果需要更改当前商品类目，请填写该标识符。
	Depth int32 `json:"depth"` // 包装厚度。
	GeoNames []string `json:"geo_names"` // 地理限制 —— 如有必要，请在个人中心中创建或编辑商品时填写该参数。 该参数为可选项。
	OfferID string `json:"offer_id"` // 卖方系统中的商品ID — 货号。至50字符。
	Vat string `json:"vat"` // 商品增值税税率： - `0` — 免除增值税, - `0.05` — 5%, - `0.07` — 7%, - `0.1` — 10%, - `0.2` — 20%, - `0.22` — 22%。 传递当前有效的出价值。
	ComplexAttributes []V3ImportProductsRequestComplexAttribute `json:"complex_attributes"` // 具有嵌套属性的特征组。
	CurrencyCode string `json:"currency_code"` // 价格显示的货币。 传输值必须与个人中心设置中所设置的货币相匹配。 默认情况下，显示`RUB` — 俄罗斯卢布。 例如，如果您设人民币为结算货币，请赋值`CNY`，否则将返回错误。 可能的值： - `RUB` — 俄罗斯卢布, - `BYN`...
	ServiceType V2ServiceType `json:"service_type"`
	Name string `json:"name"` // 商品名称。 至500字符。 [商品名称规则](https://global-help.ozon.com/zh/products/requirements/product-info/naming-requirements/)
	OldPrice string `json:"old_price"` // 折扣前的价格（将在产品卡上划掉）。 以卢布表示。分数分隔 符号 — 点, 在点之后最多两个字符。 如果您之前传递了 `old_price`, 那么在更新 `price` 时也请更新 `old_price`。
	Price string `json:"price"` // 商品价格，包括折扣，显示在商品卡上。 如果商品没有折扣，请指定 此参数中 `old_price` 的值。
	Width int32 `json:"width"` // 包装宽度。
	Height int32 `json:"height"` // 包装高度。
	PrimaryImage string `json:"primary_image"` // 链接到商品主图。
	Promotions []ImportProductRequestPromotion `json:"promotions"` // 促销活动。
	Images360 []string `json:"images360"` // 图像组360。至70件。 格式：公共云存储中图像链接的URL。 链接的图像格式为JPG。
	TypeID int64 `json:"type_id"` // 商品类型的标识符。 可以通过方法[/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)的响应参数 `type_id` 获取对应的值。 填写此参数时...
	Weight int32 `json:"weight"` // 含包装的商品重量。 限值为1000公斤或其他换算值 计量单位。
	Barcode string `json:"barcode"` // 商品条码。
	DescriptionCategoryID int64 `json:"description_category_id"` // 类别ID。可以使用方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。要更改类目，请使用参数`new_description_cat...
	ColorImage string `json:"color_image"` // 营销色彩。 格式：公共云存储中图像链接的URL。 链接的图像格式为JPG。
	DimensionUnit string `json:"dimension_unit"` // 尺寸测量单位: - `mm` — 毫米, - `cm` — 厘米, - `in` — 英寸。
	Images []string `json:"images"` // 图像组。 最多30件。 图像以与组中相同的顺序显示在网站上。 如果不传递值 `primary_image`, 组中的第一个图像将是商品的主图。 如果您传递了值 `primary_image`，则最多传递29个图像。 如果 `primary_...
}

type V3ImportProductsRequest struct {
	Items []V3ImportProductsRequestItem `json:"items"` // 数据组。
}

type ProductUpdateOfferIdRequestUpdateOfferId struct {
	OfferID string `json:"offer_id"` // 旧货号。
	NewOfferID string `json:"new_offer_id"` // 新货号。
}

type V1ProductUpdateOfferIdRequest struct {
	UpdateOfferID []ProductUpdateOfferIdRequestUpdateOfferId `json:"update_offer_id"` // 具有新旧货号价值的配对列表。
}

type GetProductRatingBySkuResponseProductRating struct {
	Groups []GetProductRatingBySkuResponseRatingGroup `json:"groups"` // 构成内容评级的各组特征。
	SKU int64 `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	Rating float64 `json:"rating"` // 产品内容评级: 从0到100。
}

type V1GetProductRatingBySkuResponse struct {
	Products []GetProductRatingBySkuResponseProductRating `json:"products"` // 商品内容分级。
}

type ProductImportProductsBySKUResponseResult struct {
	TaskID int64 `json:"task_id"` // 进口货物任务代码。
	UnmatchedSKUList []int64 `json:"unmatched_sku_list"` // 商品Id列表。
}

// 查询结果。
type V3ImportProductsResponseResult struct {
	TaskID int64 `json:"task_id"` // 装卸任务的编号。
}

type V3ImportProductsResponse struct {
	Result V3ImportProductsResponseResult `json:"result"`
}

// 商品过滤。
type Productv4Filter struct {
	OfferID []string `json:"offer_id"` // 基于参数 `offer_id`的过滤。 可以提交数值列表。
	ProductID []string `json:"product_id"` // 基于参数 `product_id`的过滤。 可以提交最多1000个数值。
	SKU []string `json:"sku"` // Ozon 系统中的商品标识符（SKU）。
	Visibility Productv2GetProductListRequestFilterFilterVisibility `json:"visibility"`
}

type Productv4GetProductAttributesV4Request struct {
	Limit int32 `json:"limit"` // К每页的值的数量。
	SortBy string `json:"sort_by"` // 商品排序参数： - `sku` — 按Ozon系统中的商品标识符排序； - `offer_id` — 按商品货号排序； - `id` — 按商品标识符排序； - `title` — 按商品名称排序。
	SortDir string `json:"sort_dir"` // 排序方向： - `asc` — 升序， - `desc` — 降序。
	Filter Productv4Filter `json:"filter"`
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
}

type V2ProductInfoPicturesResponseError struct {
	Message string `json:"message"` // 错误描述。
	URL string `json:"url"` // 图片链接。
}

type V2ProductInfoPicturesResponseItem struct {
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	PrimaryPhoto []string `json:"primary_photo"` // 主图链接。
	Photo []string `json:"photo"` // 商品照片链接。
	ColorPhoto []string `json:"color_photo"` // 上传的颜色样本链接。
	Photo360 []string `json:"photo_360"` // 360度图片链接。
	Errors []V2ProductInfoPicturesResponseError `json:"errors"` // 商品图片相关错误列表。
}

type V1ProductUpdateOfferIdResponseError struct {
	Message string `json:"message"` // 错误信息。
	OfferID string `json:"offer_id"` // 无法更改的卖家系统中的商品标识符是商品货号。
}

type V1ProductInfoWrongVolumeRequest struct {
	Limit int64 `json:"limit"` // 响应中记录数量的限制。
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
}

type V1GetProductInfoSubscriptionResponse struct {
	Result []V1GetProductInfoSubscriptionResponseResult `json:"result"` // 操作方法结果。
}

type ProductProductArchiveRequest struct {
	ProductID []int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。您一次最多可以传递100个标识符。
}

type V1ProductInfoWrongVolumeResponse struct {
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Products []ProductInfoWrongVolumeResponseWrongVolumeProduct `json:"products"` // 商品列表。
}

type ProductImportProductsBySKUResponse struct {
	Result ProductImportProductsBySKUResponseResult `json:"result"`
}

type Productv3GetProductAttributesV3Request struct {
	SortDir string `json:"sort_dir"` // 分类方向。
	Filter Productv3Filter `json:"filter"`
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Limit int64 `json:"limit"` // 每页的值的数量。最小 —— 1，最大 —— 1000。
	SortBy string `json:"sort_by"` // 对商品进行分类的参数。
}

type V1ProductUpdateOfferIdResponse struct {
	Errors []V1ProductUpdateOfferIdResponseError `json:"errors"` // 错误清单。
}

type V2ProductInfoPicturesResponse struct {
	Items []V2ProductInfoPicturesResponseItem `json:"items"` // 商品图片。
}

type Productv3GetProductAttributesV3Response struct {
	Result []Productv3GetProductAttributesV3ResponseResult `json:"result"` // 查询结果。
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Total string `json:"total"` // 列表中的商品数量。
}

type V1ProductUpdateAttributesResponse struct {
	TaskID int64 `json:"task_id"` // 商品更新任务号码。 为检查更新状态，请将接收到的值传至方法 [/v1/product/import/info](#operation/ProductAPI_GetImportProductsInfo)。
}

type ProductGetImportProductsInfoResponseResult struct {
	Items []GetImportProductsInfoResponseResultItem `json:"items"` // 商品有关信息。
	Total int32 `json:"total"` // 在卖方系统中的卖家系统中的商品标识符 — `product_id`。
}

type ProductGetImportProductsInfoResponse struct {
	Result ProductGetImportProductsInfoResponseResult `json:"result"`
}

type ProductProductInfoPicturesResponsePicture struct {
	Is360 bool `json:"is_360"` // 表示该图片是一个360图像。
	IsColor bool `json:"is_color"` // 表示图片是样品颜色。
	IsPrimary bool `json:"is_primary"` // 表示该图片是主图像。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	State string `json:"state"` // 图片上传状态。 如果方法[/v1/product/pictures/import](#operation/ProductAPI_ProductImportPictures)被使用，方法的响应将总是`imported` —— 图像未被处理。 ...
	URL string `json:"url"` // 公共云存储中的图像的链接地址。链接图片的格式是JPG或PNG。
}

// 该方法的结果。
type Productv1ProductInfoPicturesResponseResult struct {
	Pictures []ProductProductInfoPicturesResponsePicture `json:"pictures"`
}

type Productv1ProductInfoPicturesResponse struct {
	Result Productv1ProductInfoPicturesResponseResult `json:"result"`
}

type V3GetProductInfoListRequest struct {
	SKU []string `json:"sku"` // 商品在 Ozon 系统中的标识符 — SKU。
	OfferID []string `json:"offer_id"` // 商品在卖家系统中的标识符 — 货号。
	ProductID []string `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

type V1GetProductRatingBySkuRequest struct {
	Skus []string `json:"skus"` // 需要返回内容评级的商品SKU列表。
}

type V1ProductGetRelatedSKURequest struct {
	SKU []string `json:"sku"` // SKU列表。
}

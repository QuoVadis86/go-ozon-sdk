package product

type V1GetProductInfoSubscriptionResponse struct {
	Result []interface{} `json:"result"` // 操作方法结果。
}

type V1ProductUpdateOfferIdResponse struct {
	Errors interface{} `json:"errors"` // 错误清单。
}

type Productv3GetProductListRequest struct {
	Limit int64 `json:"limit"` // 答复的信息数量。默认设置为1。最大值是1000。
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定 `last_id`。
}

type V1GetProductRatingBySkuRequest struct {
	Skus interface{} `json:"skus"` // 需要返回内容评级的商品SKU列表。
}

type V1ProductGetRelatedSKUResponse struct {
	Items interface{} `json:"items"` // 相关SKU信息。
	Errors interface{} `json:"errors"` // 错误。
}

type V2ProductInfoPicturesResponse struct {
	Items []interface{} `json:"items"` // 商品图片。
}

type Productv3GetProductAttributesV3Request struct {
	SortBy string `json:"sort_by"` // 对商品进行分类的参数。
	SortDir string `json:"sort_dir"` // 分类方向。
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Limit int64 `json:"limit"` // 每页的值的数量。最小 —— 1，最大 —— 1000。
}

type ProductImportProductsBySKURequest struct {
	Items []interface{} `json:"items"` // 商品信息。
}

type Productv2DeleteProductsRequest struct {
	Products []interface{} `json:"products"` // Ozon系统中商品的标识符 — `product_id`。
}

type ProductProductUnarchiveRequest struct {
	ProductID []interface{} `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。您一次最多可以传递100个标识符。 在一天内，您最多可以从档案中恢复100件自动归档的商品。 限额在莫斯科时间03：00更新。 手动归档的商品没有解除归档的限制。
}

type Productv3GetProductListResponse struct {
	Result interface{} `json:"result"`
}

type Productv2DeleteProductsResponse struct {
	Status []interface{} `json:"status"` // 请求的处理情况。
}

type Productv4GetProductAttributesV4Request struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Limit int32 `json:"limit"` // К每页的值的数量。
	SortBy string `json:"sort_by"` // 商品排序参数： - `sku` — 按Ozon系统中的商品标识符排序； - `offer_id` — 按商品货号排序； - `id` — 按商品标识符排序； - `title` — 按商品名称排序。
	SortDir string `json:"sort_dir"` // 排序方向： - `asc` — 升序， - `desc` — 降序。
}

type Productv1ProductInfoPicturesResponse struct {
	Result interface{} `json:"result"`
}

type V3ImportProductsRequest struct {
	Items []interface{} `json:"items"` // 数据组。
}

type ProductGetImportProductsInfoResponse struct {
	Result interface{} `json:"result"`
}

type V1ProductInfoWrongVolumeRequest struct {
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Limit int64 `json:"limit"` // 响应中记录数量的限制。
}

type V1ProductUpdateOfferIdRequest struct {
	UpdateOfferID interface{} `json:"update_offer_id"` // 具有新旧货号价值的配对列表。
}

type ProductImportProductsBySKUResponse struct {
	Result interface{} `json:"result"`
}

type V1ProductInfoWrongVolumeResponse struct {
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Products interface{} `json:"products"` // 商品列表。
}

type V1ProductGetRelatedSKURequest struct {
	SKU interface{} `json:"sku"` // SKU列表。
}

type ProductProductArchiveRequest struct {
	ProductID []interface{} `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。您一次最多可以传递100个标识符。
}

type V1ProductUpdateAttributesResponse struct {
	TaskID int64 `json:"task_id"` // 商品更新任务号码。 为检查更新状态，请将接收到的值传至方法 [/v1/product/import/info](#operation/ProductAPI_GetImportProductsInfo)。
}

type Productv3GetProductAttributesV3Response struct {
	Total string `json:"total"` // 列表中的商品数量。
	Result []interface{} `json:"result"` // 查询结果。
	LastID string `json:"last_id"` // 页面上最后一个值的ID。运行第一个查询时，将此字段留空。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
}

type V1ProductUpdateAttributesRequest struct {
	Items interface{} `json:"items"` // 需要更新的商品和特征。
}

type V1GetProductInfoSubscriptionRequest struct {
	Skus []interface{} `json:"skus"` // Ozon 系统中的 SKU、商品标识符列表。
}

type ProductGetProductInfoDescriptionResponse struct {
	Result interface{} `json:"result"`
}

type ProductGetImportProductsInfoRequest struct {
	TaskID int64 `json:"task_id"` // 进口商品的任务代码。按照运输方式筛选。可以使用方法 [/v3/product/import](#operation/ProductAPI_ImportProductsV3)获取。
}

type Productv4GetProductAttributesV4Response struct {
	Result []interface{} `json:"result"` // 查询结果。
	LastID string `json:"last_id"` // 页面上最后一个值的ID。 要检索以下数值，请从上一个查询的响应中指定`last_id`。
	Total string `json:"total"` // 列表中的商品数量。
}

type V2ProductInfoPicturesRequest struct {
	ProductID interface{} `json:"product_id"` // 商品识别码的清单。
}

type V3GetProductInfoListResponse struct {
	Items []interface{} `json:"items"` // 数据数组。
}

type Productv1ProductImportPicturesRequest struct {
	Images interface{} `json:"images"` // 数组图片链接。 最多30件。 数组中的图像是按照它们在网站上出现的顺序排列的。 数组中的第一个图像将是主图像。
	Images360 interface{} `json:"images360"` // 360图片的数组。多达70张图片。 格式：公共云存储中的图像链接地址。链接图片的格式是JPG。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	ColorImage string `json:"color_image"` // 市场营销色彩。
}

type ProductGetProductInfoDescriptionRequest struct {
	OfferID string `json:"offer_id"` // 卖家系统中的商品识别码是卖家系统中的商品标识符是商品货号。
	ProductID int64 `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
}

type V4GetUploadQuotaResponse struct {
	DailyCreate interface{} `json:"daily_create"`
	DailyUpdate interface{} `json:"daily_update"`
	OperationLimits interface{} `json:"operation_limits"`
	Total interface{} `json:"total"`
}

type V3GetProductInfoListRequest struct {
	ProductID []interface{} `json:"product_id"` // Ozon系统中商品的标识符 — `product_id`。
	SKU []interface{} `json:"sku"` // 商品在 Ozon 系统中的标识符 — SKU。
	OfferID []interface{} `json:"offer_id"` // 商品在卖家系统中的标识符 — 货号。
}

type V3ImportProductsResponse struct {
	Result interface{} `json:"result"`
}

type ProductBooleanResponse struct {
	Result bool `json:"result"` // 查询的处理结果。`true`，如果查询的执行无误。
}

type V1GetProductRatingBySkuResponse struct {
	Products interface{} `json:"products"` // 商品内容分级。
}

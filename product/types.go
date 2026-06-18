package product

type V1ProductInfoWrongVolumeResponse struct {
	Cursor string `json:"cursor"`
	Products interface{} `json:"products"`
}

type Productv3GetProductAttributesV3Request struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

type V3GetProductInfoListResponse struct {
	Items []interface{} `json:"items"`
}

type V1GetProductRatingBySkuResponse struct {
	Products interface{} `json:"products"`
}

type V3GetProductInfoListRequest struct {
	OfferID []interface{} `json:"offer_id"`
	ProductID []interface{} `json:"product_id"`
	SKU []interface{} `json:"sku"`
}

type Productv3GetProductListRequest struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int64 `json:"limit"`
}

type ProductGetProductInfoDescriptionRequest struct {
	ProductID int64 `json:"product_id"`
	OfferID string `json:"offer_id"`
}

type V1ProductGetRelatedSKURequest struct {
	SKU interface{} `json:"sku"`
}

type Productv4GetProductAttributesV4Request struct {
	Filter interface{} `json:"filter"`
	LastID string `json:"last_id"`
	Limit int32 `json:"limit"`
	SortBy string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

type ProductGetImportProductsInfoResponse struct {
	Result interface{} `json:"result"`
}

type ProductImportProductsBySKUResponse struct {
	Result interface{} `json:"result"`
}

type V3ImportProductsResponse struct {
	Result interface{} `json:"result"`
}

type V1GetProductInfoSubscriptionRequest struct {
	Skus []interface{} `json:"skus"`
}

type ProductGetImportProductsInfoRequest struct {
	TaskID int64 `json:"task_id"`
}

type V1ProductUpdateAttributesRequest struct {
	Items interface{} `json:"items"`
}

type ProductProductUnarchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type ProductBooleanResponse struct {
	Result bool `json:"result"`
}

type V1ProductUpdateOfferIdRequest struct {
	UpdateOfferID interface{} `json:"update_offer_id"`
}

type ProductProductArchiveRequest struct {
	ProductID []interface{} `json:"product_id"`
}

type ProductGetProductInfoDescriptionResponse struct {
	Result interface{} `json:"result"`
}

type V1ProductUpdateOfferIdResponse struct {
	Errors interface{} `json:"errors"`
}

type Productv2DeleteProductsResponse struct {
	Status []interface{} `json:"status"`
}

type Productv1ProductInfoPicturesResponse struct {
	Result interface{} `json:"result"`
}

type V2ProductInfoPicturesRequest struct {
	ProductID interface{} `json:"product_id"`
}

type V1GetProductInfoSubscriptionResponse struct {
	Result []interface{} `json:"result"`
}

type V1GetProductRatingBySkuRequest struct {
	Skus interface{} `json:"skus"`
}

type V4GetUploadQuotaResponse struct {
	OperationLimits interface{} `json:"operation_limits"`
	Total interface{} `json:"total"`
	DailyCreate interface{} `json:"daily_create"`
	DailyUpdate interface{} `json:"daily_update"`
}

type Productv3GetProductAttributesV3Response struct {
	LastID string `json:"last_id"`
	Total string `json:"total"`
	Result []interface{} `json:"result"`
}

type V2ProductInfoPicturesResponse struct {
	Items []interface{} `json:"items"`
}

type V1ProductUpdateAttributesResponse struct {
	TaskID int64 `json:"task_id"`
}

type ProductImportProductsBySKURequest struct {
	Items []interface{} `json:"items"`
}

type V3ImportProductsRequest struct {
	Items []interface{} `json:"items"`
}

type Productv4GetProductAttributesV4Response struct {
	Result []interface{} `json:"result"`
	LastID string `json:"last_id"`
	Total string `json:"total"`
}

type Productv3GetProductListResponse struct {
	Result interface{} `json:"result"`
}

type Productv2DeleteProductsRequest struct {
	Products []interface{} `json:"products"`
}

type Productv1ProductImportPicturesRequest struct {
	ColorImage string `json:"color_image"`
	Images interface{} `json:"images"`
	Images360 interface{} `json:"images360"`
	ProductID int64 `json:"product_id"`
}

type V1ProductInfoWrongVolumeRequest struct {
	Cursor string `json:"cursor"`
	Limit int64 `json:"limit"`
}

type V1ProductGetRelatedSKUResponse struct {
	Items interface{} `json:"items"`
	Errors interface{} `json:"errors"`
}

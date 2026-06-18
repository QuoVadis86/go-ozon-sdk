package product

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 将商品归档
// Note: 请使用方法[/v3/product/info/list](#operation/ProductAPI_GetProductInfoList)
// Note: 如需查看商品归档后的实际状态，请使用方法[/v3/product/info/list](#operation/ProductAPI_GetProductInfoList)
func (s *Service) ProductArchive(ctx context.Context, req *ProductArchiveRequest) (*BooleanResponse, error) {
	var resp BooleanResponse
	err := s.Client.Post(ctx, "/v1/product/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 品类限制、商品的创建和更新
// Note: 每分钟可创建的商品数量
// Note: 无法创建新商品
func (s *Service) GetUploadQuota(ctx context.Context) (*V4GetUploadQuotaResponse, error) {
	var resp V4GetUploadQuotaResponse
	err := s.Client.Post(ctx, "/v4/product/info/limit", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取商品详细信息
func (s *Service) GetProductInfoDescription(ctx context.Context, req *GetProductInfoDescriptionRequest) (*GetProductInfoDescriptionResponse, error) {
	var resp GetProductInfoDescriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/description", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 按SKU获得商品的内容排名
func (s *Service) GetProductRatingBySku(ctx context.Context, req *V1GetProductRatingBySkuRequest) (*V1GetProductRatingBySkuResponse, error) {
	var resp V1GetProductRatingBySkuResponse
	err := s.Client.Post(ctx, "/v1/product/rating-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取商品特征描述
func (s *Service) GetProductAttributesV4(ctx context.Context, req *V4GetProductAttributesV4Request) (*V4GetProductAttributesV4Response, error) {
	var resp V4GetProductAttributesV4Response
	err := s.Client.Post(ctx, "/v4/product/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 体积重量特征不正确的商品列表
func (s *Service) ProductInfoWrongVolume(ctx context.Context, req *V1ProductInfoWrongVolumeRequest) (*V1ProductInfoWrongVolumeResponse, error) {
	var resp V1ProductInfoWrongVolumeResponse
	err := s.Client.Post(ctx, "/v1/product/info/wrong-volume", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建或更新商品
// Note: 请使用 [/v4/product/info/limit](#operation/ProductAPI_GetUploadQuota)。 如果商品下载和更新次数 超过限制，则
// Note: 每件商品生成条形码
// Note: 无法接收没有条形码的商品
func (s *Service) ImportProductsV3(ctx context.Context, req *V3ImportProductsRequest) (*V3ImportProductsResponse, error) {
	var resp V3ImportProductsResponse
	err := s.Client.Post(ctx, "/v3/product/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取商品图片
func (s *Service) ProductInfoPicturesV2(ctx context.Context, req *V2ProductInfoPicturesRequest) (*V2ProductInfoPicturesResponse, error) {
	var resp V2ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v2/product/pictures/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 从存档删除没有SKU的商品
func (s *Service) DeleteProducts(ctx context.Context, req *V2DeleteProductsRequest) (*V2DeleteProductsResponse, error) {
	var resp V2DeleteProductsResponse
	err := s.Client.Post(ctx, "/v2/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 从卖家的系统中改变商品货号
// Note: 每分钟和每天的商品操作数量有限制
func (s *Service) ProductUpdateOfferID(ctx context.Context, req *V1ProductUpdateOfferIdRequest) (*V1ProductUpdateOfferIdResponse, error) {
	var resp V1ProductUpdateOfferIdResponse
	err := s.Client.Post(ctx, "/v1/product/update/offer-id", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 上传或更新商品图片
// Note: 请使用 [/v3/product/info/list](#operation/ProductAPI_GetProductInfoList) 方法获取信息 —— 它显示了当前图像的
// Note: 每次调用该方法时，要传递所有应该出现在商品详情页上的图片
func (s *Service) ProductImportPictures(ctx context.Context, req *V1ProductImportPicturesRequest) (*V1ProductInfoPicturesResponse, error) {
	var resp V1ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v1/product/pictures/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 从档案中还原商品
func (s *Service) ProductUnarchive(ctx context.Context, req *ProductUnarchiveRequest) (*BooleanResponse, error) {
	var resp BooleanResponse
	err := s.Client.Post(ctx, "/v1/product/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 订阅该商品的用户数
func (s *Service) GetProductInfoSubscription(ctx context.Context, req *V1GetProductInfoSubscriptionRequest) (*V1GetProductInfoSubscriptionResponse, error) {
	var resp V1GetProductInfoSubscriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/subscription", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 品列表的
// Note: 每次请求只能使用一组标识符，且商品数量不能超过 1000 个
// Note: 不能超过 1000 个
func (s *Service) GetProductList(ctx context.Context, req *V3GetProductListRequest) (*V3GetProductListResponse, error) {
	var resp V3GetProductListResponse
	err := s.Client.Post(ctx, "/v3/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取相关SKU
func (s *Service) ProductGetRelatedSKU(ctx context.Context, req *V1ProductGetRelatedSKURequest) (*V1ProductGetRelatedSKUResponse, error) {
	var resp V1ProductGetRelatedSKUResponse
	err := s.Client.Post(ctx, "/v1/product/related-sku/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取商品特征描述
func (s *Service) GetProductAttributesV3(ctx context.Context, req *V3GetProductAttributesV3Request) (*V3GetProductAttributesV3Response, error) {
	var resp V3GetProductAttributesV3Response
	err := s.Client.Post(ctx, "/v3/products/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 查询商品添加或更新状态
func (s *Service) GetImportProductsInfo(ctx context.Context, req *GetImportProductsInfoRequest) (*GetImportProductsInfoResponse, error) {
	var resp GetImportProductsInfoResponse
	err := s.Client.Post(ctx, "/v1/product/import/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新商品特征
// Note: 请使用/v3/product/import
// Note: 每分钟和每天的商品操作数量有限制
// Note: 无法删除。 如需完整更新全部特征，请使用/v3/product/import
func (s *Service) ProductUpdateAttributes(ctx context.Context, req *V1ProductUpdateAttributesRequest) (*V1ProductUpdateAttributesResponse, error) {
	var resp V1ProductUpdateAttributesResponse
	err := s.Client.Post(ctx, "/v1/product/attributes/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 通过SKU创建商品
// Note: 每分钟和每天的商品操作数量有限制
// Note: 无法复制您自己的商品
func (s *Service) ImportProductsBySKU(ctx context.Context, req *ImportProductsBySKURequest) (*ImportProductsBySKUResponse, error) {
	var resp ImportProductsBySKUResponse
	err := s.Client.Post(ctx, "/v1/product/import-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 根据标识符获取商品信息
func (s *Service) GetProductInfoList(ctx context.Context, req *V3GetProductInfoListRequest) (*V3GetProductInfoListResponse, error) {
	var resp V3GetProductInfoListResponse
	err := s.Client.Post(ctx, "/v3/product/info/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

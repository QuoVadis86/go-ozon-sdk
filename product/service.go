package product

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) GetUploadQuota(ctx context.Context) (*types.V4GetUploadQuotaResponse, error) {
	var resp types.V4GetUploadQuotaResponse
	err := s.Client.Post(ctx, "/v4/product/info/limit", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoList(ctx context.Context, req *types.V3GetProductInfoListRequest) (*types.V3GetProductInfoListResponse, error) {
	var resp types.V3GetProductInfoListResponse
	err := s.Client.Post(ctx, "/v3/product/info/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsV3(ctx context.Context, req *types.V3ImportProductsRequest) (*types.V3ImportProductsResponse, error) {
	var resp types.V3ImportProductsResponse
	err := s.Client.Post(ctx, "/v3/product/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductImportPictures(ctx context.Context, req *types.Productv1ProductImportPicturesRequest) (*types.Productv1ProductInfoPicturesResponse, error) {
	var resp types.Productv1ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v1/product/pictures/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUnarchive(ctx context.Context, req *types.ProductProductUnarchiveRequest) (*types.ProductBooleanResponse, error) {
	var resp types.ProductBooleanResponse
	err := s.Client.Post(ctx, "/v1/product/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductInfoPicturesV2(ctx context.Context, req *types.V2ProductInfoPicturesRequest) (*types.V2ProductInfoPicturesResponse, error) {
	var resp types.V2ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v2/product/pictures/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductGetRelatedSKU(ctx context.Context, req *types.V1ProductGetRelatedSKURequest) (*types.V1ProductGetRelatedSKUResponse, error) {
	var resp types.V1ProductGetRelatedSKUResponse
	err := s.Client.Post(ctx, "/v1/product/related-sku/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductAttributesV4(ctx context.Context, req *types.Productv4GetProductAttributesV4Request) (*types.Productv4GetProductAttributesV4Response, error) {
	var resp types.Productv4GetProductAttributesV4Response
	err := s.Client.Post(ctx, "/v4/product/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoDescription(ctx context.Context, req *types.ProductGetProductInfoDescriptionRequest) (*types.ProductGetProductInfoDescriptionResponse, error) {
	var resp types.ProductGetProductInfoDescriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/description", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductArchive(ctx context.Context, req *types.ProductProductArchiveRequest) (*types.ProductBooleanResponse, error) {
	var resp types.ProductBooleanResponse
	err := s.Client.Post(ctx, "/v1/product/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductList(ctx context.Context, req *types.Productv3GetProductListRequest) (*types.Productv3GetProductListResponse, error) {
	var resp types.Productv3GetProductListResponse
	err := s.Client.Post(ctx, "/v3/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductAttributesV3(ctx context.Context, req *types.Productv3GetProductAttributesV3Request) (*types.Productv3GetProductAttributesV3Response, error) {
	var resp types.Productv3GetProductAttributesV3Response
	err := s.Client.Post(ctx, "/v3/products/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductRatingBySku(ctx context.Context, req *types.V1GetProductRatingBySkuRequest) (*types.V1GetProductRatingBySkuResponse, error) {
	var resp types.V1GetProductRatingBySkuResponse
	err := s.Client.Post(ctx, "/v1/product/rating-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetImportProductsInfo(ctx context.Context, req *types.ProductGetImportProductsInfoRequest) (*types.ProductGetImportProductsInfoResponse, error) {
	var resp types.ProductGetImportProductsInfoResponse
	err := s.Client.Post(ctx, "/v1/product/import/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoSubscription(ctx context.Context, req *types.V1GetProductInfoSubscriptionRequest) (*types.V1GetProductInfoSubscriptionResponse, error) {
	var resp types.V1GetProductInfoSubscriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/subscription", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateAttributes(ctx context.Context, req *types.V1ProductUpdateAttributesRequest) (*types.V1ProductUpdateAttributesResponse, error) {
	var resp types.V1ProductUpdateAttributesResponse
	err := s.Client.Post(ctx, "/v1/product/attributes/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DeleteProducts(ctx context.Context, req *types.Productv2DeleteProductsRequest) (*types.Productv2DeleteProductsResponse, error) {
	var resp types.Productv2DeleteProductsResponse
	err := s.Client.Post(ctx, "/v2/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsBySKU(ctx context.Context, req *types.ProductImportProductsBySKURequest) (*types.ProductImportProductsBySKUResponse, error) {
	var resp types.ProductImportProductsBySKUResponse
	err := s.Client.Post(ctx, "/v1/product/import-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductInfoWrongVolume(ctx context.Context, req *types.V1ProductInfoWrongVolumeRequest) (*types.V1ProductInfoWrongVolumeResponse, error) {
	var resp types.V1ProductInfoWrongVolumeResponse
	err := s.Client.Post(ctx, "/v1/product/info/wrong-volume", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateOfferID(ctx context.Context, req *types.V1ProductUpdateOfferIdRequest) (*types.V1ProductUpdateOfferIdResponse, error) {
	var resp types.V1ProductUpdateOfferIdResponse
	err := s.Client.Post(ctx, "/v1/product/update/offer-id", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

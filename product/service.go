package product

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

func (s *Service) GetProductAttributesV4(ctx context.Context, req *Productv4GetProductAttributesV4Request) (*Productv4GetProductAttributesV4Response, error) {
	var resp Productv4GetProductAttributesV4Response
	err := s.Client.Post(ctx, "/v4/product/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductImportPictures(ctx context.Context, req *Productv1ProductImportPicturesRequest) (*Productv1ProductInfoPicturesResponse, error) {
	var resp Productv1ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v1/product/pictures/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductGetRelatedSKU(ctx context.Context, req *V1ProductGetRelatedSKURequest) (*V1ProductGetRelatedSKUResponse, error) {
	var resp V1ProductGetRelatedSKUResponse
	err := s.Client.Post(ctx, "/v1/product/related-sku/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductArchive(ctx context.Context, req *ProductProductArchiveRequest) (*ProductBooleanResponse, error) {
	var resp ProductBooleanResponse
	err := s.Client.Post(ctx, "/v1/product/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductList(ctx context.Context, req *Productv3GetProductListRequest) (*Productv3GetProductListResponse, error) {
	var resp Productv3GetProductListResponse
	err := s.Client.Post(ctx, "/v3/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetImportProductsInfo(ctx context.Context, req *ProductGetImportProductsInfoRequest) (*ProductGetImportProductsInfoResponse, error) {
	var resp ProductGetImportProductsInfoResponse
	err := s.Client.Post(ctx, "/v1/product/import/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsV3(ctx context.Context, req *V3ImportProductsRequest) (*V3ImportProductsResponse, error) {
	var resp V3ImportProductsResponse
	err := s.Client.Post(ctx, "/v3/product/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductAttributesV3(ctx context.Context, req *Productv3GetProductAttributesV3Request) (*Productv3GetProductAttributesV3Response, error) {
	var resp Productv3GetProductAttributesV3Response
	err := s.Client.Post(ctx, "/v3/products/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetUploadQuota(ctx context.Context) (*V4GetUploadQuotaResponse, error) {
	var resp V4GetUploadQuotaResponse
	err := s.Client.Post(ctx, "/v4/product/info/limit", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductInfoWrongVolume(ctx context.Context, req *V1ProductInfoWrongVolumeRequest) (*V1ProductInfoWrongVolumeResponse, error) {
	var resp V1ProductInfoWrongVolumeResponse
	err := s.Client.Post(ctx, "/v1/product/info/wrong-volume", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUnarchive(ctx context.Context, req *ProductProductUnarchiveRequest) (*ProductBooleanResponse, error) {
	var resp ProductBooleanResponse
	err := s.Client.Post(ctx, "/v1/product/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductInfoPicturesV2(ctx context.Context, req *V2ProductInfoPicturesRequest) (*V2ProductInfoPicturesResponse, error) {
	var resp V2ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v2/product/pictures/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoSubscription(ctx context.Context, req *V1GetProductInfoSubscriptionRequest) (*V1GetProductInfoSubscriptionResponse, error) {
	var resp V1GetProductInfoSubscriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/subscription", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateOfferID(ctx context.Context, req *V1ProductUpdateOfferIdRequest) (*V1ProductUpdateOfferIdResponse, error) {
	var resp V1ProductUpdateOfferIdResponse
	err := s.Client.Post(ctx, "/v1/product/update/offer-id", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateAttributes(ctx context.Context, req *V1ProductUpdateAttributesRequest) (*V1ProductUpdateAttributesResponse, error) {
	var resp V1ProductUpdateAttributesResponse
	err := s.Client.Post(ctx, "/v1/product/attributes/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoDescription(ctx context.Context, req *ProductGetProductInfoDescriptionRequest) (*ProductGetProductInfoDescriptionResponse, error) {
	var resp ProductGetProductInfoDescriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/description", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DeleteProducts(ctx context.Context, req *Productv2DeleteProductsRequest) (*Productv2DeleteProductsResponse, error) {
	var resp Productv2DeleteProductsResponse
	err := s.Client.Post(ctx, "/v2/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoList(ctx context.Context, req *V3GetProductInfoListRequest) (*V3GetProductInfoListResponse, error) {
	var resp V3GetProductInfoListResponse
	err := s.Client.Post(ctx, "/v3/product/info/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductRatingBySku(ctx context.Context, req *V1GetProductRatingBySkuRequest) (*V1GetProductRatingBySkuResponse, error) {
	var resp V1GetProductRatingBySkuResponse
	err := s.Client.Post(ctx, "/v1/product/rating-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsBySKU(ctx context.Context, req *ProductImportProductsBySKURequest) (*ProductImportProductsBySKUResponse, error) {
	var resp ProductImportProductsBySKUResponse
	err := s.Client.Post(ctx, "/v1/product/import-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

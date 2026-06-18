package product

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) ProductInfoWrongVolume(ctx context.Context, req *internal.V1ProductInfoWrongVolumeRequest) (*internal.V1ProductInfoWrongVolumeResponse, error) {
	var resp internal.V1ProductInfoWrongVolumeResponse
	err := s.Client.Post(ctx, "/v1/product/info/wrong-volume", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductArchive(ctx context.Context, req *internal.ProductProductArchiveRequest) (*internal.ProductBooleanResponse, error) {
	var resp internal.ProductBooleanResponse
	err := s.Client.Post(ctx, "/v1/product/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductInfoPicturesV2(ctx context.Context, req *internal.V2ProductInfoPicturesRequest) (*internal.V2ProductInfoPicturesResponse, error) {
	var resp internal.V2ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v2/product/pictures/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductRatingBySku(ctx context.Context, req *internal.V1GetProductRatingBySkuRequest) (*internal.V1GetProductRatingBySkuResponse, error) {
	var resp internal.V1GetProductRatingBySkuResponse
	err := s.Client.Post(ctx, "/v1/product/rating-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsBySKU(ctx context.Context, req *internal.ProductImportProductsBySKURequest) (*internal.ProductImportProductsBySKUResponse, error) {
	var resp internal.ProductImportProductsBySKUResponse
	err := s.Client.Post(ctx, "/v1/product/import-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetUploadQuota(ctx context.Context) (*internal.V4GetUploadQuotaResponse, error) {
	var resp internal.V4GetUploadQuotaResponse
	err := s.Client.Post(ctx, "/v4/product/info/limit", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductGetRelatedSKU(ctx context.Context, req *internal.V1ProductGetRelatedSKURequest) (*internal.V1ProductGetRelatedSKUResponse, error) {
	var resp internal.V1ProductGetRelatedSKUResponse
	err := s.Client.Post(ctx, "/v1/product/related-sku/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductList(ctx context.Context, req *internal.Productv3GetProductListRequest) (*internal.Productv3GetProductListResponse, error) {
	var resp internal.Productv3GetProductListResponse
	err := s.Client.Post(ctx, "/v3/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUnarchive(ctx context.Context, req *internal.ProductProductUnarchiveRequest) (*internal.ProductBooleanResponse, error) {
	var resp internal.ProductBooleanResponse
	err := s.Client.Post(ctx, "/v1/product/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateAttributes(ctx context.Context, req *internal.V1ProductUpdateAttributesRequest) (*internal.V1ProductUpdateAttributesResponse, error) {
	var resp internal.V1ProductUpdateAttributesResponse
	err := s.Client.Post(ctx, "/v1/product/attributes/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DeleteProducts(ctx context.Context, req *internal.Productv2DeleteProductsRequest) (*internal.Productv2DeleteProductsResponse, error) {
	var resp internal.Productv2DeleteProductsResponse
	err := s.Client.Post(ctx, "/v2/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsV3(ctx context.Context, req *internal.V3ImportProductsRequest) (*internal.V3ImportProductsResponse, error) {
	var resp internal.V3ImportProductsResponse
	err := s.Client.Post(ctx, "/v3/product/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductImportPictures(ctx context.Context, req *internal.Productv1ProductImportPicturesRequest) (*internal.Productv1ProductInfoPicturesResponse, error) {
	var resp internal.Productv1ProductInfoPicturesResponse
	err := s.Client.Post(ctx, "/v1/product/pictures/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoList(ctx context.Context, req *internal.V3GetProductInfoListRequest) (*internal.V3GetProductInfoListResponse, error) {
	var resp internal.V3GetProductInfoListResponse
	err := s.Client.Post(ctx, "/v3/product/info/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoDescription(ctx context.Context, req *internal.ProductGetProductInfoDescriptionRequest) (*internal.ProductGetProductInfoDescriptionResponse, error) {
	var resp internal.ProductGetProductInfoDescriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/description", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductAttributesV4(ctx context.Context, req *internal.Productv4GetProductAttributesV4Request) (*internal.Productv4GetProductAttributesV4Response, error) {
	var resp internal.Productv4GetProductAttributesV4Response
	err := s.Client.Post(ctx, "/v4/product/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoSubscription(ctx context.Context, req *internal.V1GetProductInfoSubscriptionRequest) (*internal.V1GetProductInfoSubscriptionResponse, error) {
	var resp internal.V1GetProductInfoSubscriptionResponse
	err := s.Client.Post(ctx, "/v1/product/info/subscription", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetImportProductsInfo(ctx context.Context, req *internal.ProductGetImportProductsInfoRequest) (*internal.ProductGetImportProductsInfoResponse, error) {
	var resp internal.ProductGetImportProductsInfoResponse
	err := s.Client.Post(ctx, "/v1/product/import/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateOfferID(ctx context.Context, req *internal.V1ProductUpdateOfferIdRequest) (*internal.V1ProductUpdateOfferIdResponse, error) {
	var resp internal.V1ProductUpdateOfferIdResponse
	err := s.Client.Post(ctx, "/v1/product/update/offer-id", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductAttributesV3(ctx context.Context, req *internal.Productv3GetProductAttributesV3Request) (*internal.Productv3GetProductAttributesV3Response, error) {
	var resp internal.Productv3GetProductAttributesV3Response
	err := s.Client.Post(ctx, "/v3/products/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

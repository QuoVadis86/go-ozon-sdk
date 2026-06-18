package ozon

import "context"

func (s *ProductService) GetProductAttributesV4(ctx context.Context, req *Productv4GetProductAttributesV4Request) (*Productv4GetProductAttributesV4Response, error) {
	var resp Productv4GetProductAttributesV4Response
	_, err := s.t.Post(ctx, "/v4/product/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetProductAttributesV3(ctx context.Context, req *Productv3GetProductAttributesV3Request) (*Productv3GetProductAttributesV3Response, error) {
	var resp Productv3GetProductAttributesV3Response
	_, err := s.t.Post(ctx, "/v3/products/info/attributes", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductUpdateAttributes(ctx context.Context, req *V1ProductUpdateAttributesRequest) (*V1ProductUpdateAttributesResponse, error) {
	var resp V1ProductUpdateAttributesResponse
	_, err := s.t.Post(ctx, "/v1/product/attributes/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductArchive(ctx context.Context, req *ProductProductArchiveRequest) (*ProductBooleanResponse, error) {
	var resp ProductBooleanResponse
	_, err := s.t.Post(ctx, "/v1/product/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductUpdateOfferID(ctx context.Context, req *V1ProductUpdateOfferIdRequest) (*V1ProductUpdateOfferIdResponse, error) {
	var resp V1ProductUpdateOfferIdResponse
	_, err := s.t.Post(ctx, "/v1/product/update/offer-id", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetImportProductsInfo(ctx context.Context, req *ProductGetImportProductsInfoRequest) (*ProductGetImportProductsInfoResponse, error) {
	var resp ProductGetImportProductsInfoResponse
	_, err := s.t.Post(ctx, "/v1/product/import/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetProductList(ctx context.Context, req *Productv3GetProductListRequest) (*Productv3GetProductListResponse, error) {
	var resp Productv3GetProductListResponse
	_, err := s.t.Post(ctx, "/v3/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ImportProductsV3(ctx context.Context, req *V3ImportProductsRequest) (*V3ImportProductsResponse, error) {
	var resp V3ImportProductsResponse
	_, err := s.t.Post(ctx, "/v3/product/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductInfoWrongVolume(ctx context.Context, req *V1ProductInfoWrongVolumeRequest) (*V1ProductInfoWrongVolumeResponse, error) {
	var resp V1ProductInfoWrongVolumeResponse
	_, err := s.t.Post(ctx, "/v1/product/info/wrong-volume", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductInfoPicturesV2(ctx context.Context, req *V2ProductInfoPicturesRequest) (*V2ProductInfoPicturesResponse, error) {
	var resp V2ProductInfoPicturesResponse
	_, err := s.t.Post(ctx, "/v2/product/pictures/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductImportPictures(ctx context.Context, req *Productv1ProductImportPicturesRequest) (*Productv1ProductInfoPicturesResponse, error) {
	var resp Productv1ProductInfoPicturesResponse
	_, err := s.t.Post(ctx, "/v1/product/pictures/import", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ImportProductsBySKU(ctx context.Context, req *ProductImportProductsBySKURequest) (*ProductImportProductsBySKUResponse, error) {
	var resp ProductImportProductsBySKUResponse
	_, err := s.t.Post(ctx, "/v1/product/import-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetProductInfoSubscription(ctx context.Context, req *V1GetProductInfoSubscriptionRequest) (*V1GetProductInfoSubscriptionResponse, error) {
	var resp V1GetProductInfoSubscriptionResponse
	_, err := s.t.Post(ctx, "/v1/product/info/subscription", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetProductInfoList(ctx context.Context, req *V3GetProductInfoListRequest) (*V3GetProductInfoListResponse, error) {
	var resp V3GetProductInfoListResponse
	_, err := s.t.Post(ctx, "/v3/product/info/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductUnarchive(ctx context.Context, req *ProductProductUnarchiveRequest) (*ProductBooleanResponse, error) {
	var resp ProductBooleanResponse
	_, err := s.t.Post(ctx, "/v1/product/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetUploadQuota(ctx context.Context) (*V4GetUploadQuotaResponse, error) {
	var resp V4GetUploadQuotaResponse
	_, err := s.t.Post(ctx, "/v4/product/info/limit", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetProductInfoDescription(ctx context.Context, req *ProductGetProductInfoDescriptionRequest) (*ProductGetProductInfoDescriptionResponse, error) {
	var resp ProductGetProductInfoDescriptionResponse
	_, err := s.t.Post(ctx, "/v1/product/info/description", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) ProductGetRelatedSKU(ctx context.Context, req *V1ProductGetRelatedSKURequest) (*V1ProductGetRelatedSKUResponse, error) {
	var resp V1ProductGetRelatedSKUResponse
	_, err := s.t.Post(ctx, "/v1/product/related-sku/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) GetProductRatingBySku(ctx context.Context, req *V1GetProductRatingBySkuRequest) (*V1GetProductRatingBySkuResponse, error) {
	var resp V1GetProductRatingBySkuResponse
	_, err := s.t.Post(ctx, "/v1/product/rating-by-sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ProductService) DeleteProducts(ctx context.Context, req *Productv2DeleteProductsRequest) (*Productv2DeleteProductsResponse, error) {
	var resp Productv2DeleteProductsResponse
	_, err := s.t.Post(ctx, "/v2/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

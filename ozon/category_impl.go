package ozon

import "context"

func (s *CategoryService) GetTree(ctx context.Context, req *V1GetTreeRequest) (*V1GetTreeResponse, error) {
	var resp V1GetTreeResponse
	_, err := s.t.Post(ctx, "/v1/description-category/tree", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *CategoryService) GetAttributeValues(ctx context.Context, req *V1GetAttributeValuesRequest) (*V1GetAttributeValuesResponse, error) {
	var resp V1GetAttributeValuesResponse
	_, err := s.t.Post(ctx, "/v1/description-category/attribute/values", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *CategoryService) GetAttributes(ctx context.Context, req *V1GetAttributesRequest) (*V1GetAttributesResponse, error) {
	var resp V1GetAttributesResponse
	_, err := s.t.Post(ctx, "/v1/description-category/attribute", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *CategoryService) SearchAttributeValues(ctx context.Context, req *V1SearchAttributeValuesRequest) (*V1SearchAttributeValuesResponse, error) {
	var resp V1SearchAttributeValuesResponse
	_, err := s.t.Post(ctx, "/v1/description-category/attribute/values/search", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

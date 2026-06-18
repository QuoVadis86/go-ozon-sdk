package category

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) GetAttributes(ctx context.Context, req *internal.V1GetAttributesRequest) (*internal.V1GetAttributesResponse, error) {
	var resp internal.V1GetAttributesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SearchAttributeValues(ctx context.Context, req *internal.V1SearchAttributeValuesRequest) (*internal.V1SearchAttributeValuesResponse, error) {
	var resp internal.V1SearchAttributeValuesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute/values/search", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetTree(ctx context.Context, req *internal.V1GetTreeRequest) (*internal.V1GetTreeResponse, error) {
	var resp internal.V1GetTreeResponse
	err := s.Client.Post(ctx, "/v1/description-category/tree", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetAttributeValues(ctx context.Context, req *internal.V1GetAttributeValuesRequest) (*internal.V1GetAttributeValuesResponse, error) {
	var resp internal.V1GetAttributeValuesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute/values", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

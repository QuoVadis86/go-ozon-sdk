package category

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) GetAttributes(ctx context.Context, req *types.V1GetAttributesRequest) (*types.V1GetAttributesResponse, error) {
	var resp types.V1GetAttributesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SearchAttributeValues(ctx context.Context, req *types.V1SearchAttributeValuesRequest) (*types.V1SearchAttributeValuesResponse, error) {
	var resp types.V1SearchAttributeValuesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute/values/search", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetTree(ctx context.Context, req *types.V1GetTreeRequest) (*types.V1GetTreeResponse, error) {
	var resp types.V1GetTreeResponse
	err := s.Client.Post(ctx, "/v1/description-category/tree", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetAttributeValues(ctx context.Context, req *types.V1GetAttributeValuesRequest) (*types.V1GetAttributeValuesResponse, error) {
	var resp types.V1GetAttributeValuesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute/values", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

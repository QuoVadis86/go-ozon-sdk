package pricing

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *internal.Client }

func (s *Service) ItemsDelete(ctx context.Context, req *types.V1ItemIDsRequest) (*types.V1DeleteStrategyItemsResponse, error) {
	var resp types.V1DeleteStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) List(ctx context.Context, req *types.V1GetStrategyListRequest) (*types.V1GetStrategyListResponse, error) {
	var resp types.V1GetStrategyListResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Competitors(ctx context.Context, req *types.V1GetCompetitorsRequest) (*types.V1GetCompetitorsResponse, error) {
	var resp types.V1GetCompetitorsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/competitors/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Update(ctx context.Context, req *types.V1UpdatePricingStrategyRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Info(ctx context.Context, req *types.V1StrategyRequest) (*types.V1GetStrategyResponse, error) {
	var resp types.V1GetStrategyResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ItemsAdd(ctx context.Context, req *types.V1AddStrategyItemsRequest) (*types.V1AddStrategyItemsResponse, error) {
	var resp types.V1AddStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ItemsList(ctx context.Context, req *types.V1StrategyRequest) (*types.V1GetStrategyItemsResponse, error) {
	var resp types.V1GetStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Create(ctx context.Context, req *types.V1CreatePricingStrategyRequest) (*types.V1CreatePricingStrategyResponse, error) {
	var resp types.V1CreatePricingStrategyResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ItemsInfo(ctx context.Context, req *types.V1GetStrategyItemInfoRequest) (*types.V1GetStrategyItemInfoResponse, error) {
	var resp types.V1GetStrategyItemInfoResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/product/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Delete(ctx context.Context, req *types.V1StrategyRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Ids(ctx context.Context, req *types.V1ItemIDsRequest) (*types.V1GetStrategyIDsByItemIDsResponse, error) {
	var resp types.V1GetStrategyIDsByItemIDsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/strategy-ids-by-product-ids", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Status(ctx context.Context, req *types.V1UpdateStatusStrategyRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

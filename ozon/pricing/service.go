package pricing

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) Create(ctx context.Context, req *internal.V1CreatePricingStrategyRequest) (*internal.V1CreatePricingStrategyResponse, error) {
	var resp internal.V1CreatePricingStrategyResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Status(ctx context.Context, req *internal.V1UpdateStatusStrategyRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ItemsDelete(ctx context.Context, req *internal.V1ItemIDsRequest) (*internal.V1DeleteStrategyItemsResponse, error) {
	var resp internal.V1DeleteStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Update(ctx context.Context, req *internal.V1UpdatePricingStrategyRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Delete(ctx context.Context, req *internal.V1StrategyRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ItemsList(ctx context.Context, req *internal.V1StrategyRequest) (*internal.V1GetStrategyItemsResponse, error) {
	var resp internal.V1GetStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) List(ctx context.Context, req *internal.V1GetStrategyListRequest) (*internal.V1GetStrategyListResponse, error) {
	var resp internal.V1GetStrategyListResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ItemsAdd(ctx context.Context, req *internal.V1AddStrategyItemsRequest) (*internal.V1AddStrategyItemsResponse, error) {
	var resp internal.V1AddStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Competitors(ctx context.Context, req *internal.V1GetCompetitorsRequest) (*internal.V1GetCompetitorsResponse, error) {
	var resp internal.V1GetCompetitorsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/competitors/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Info(ctx context.Context, req *internal.V1StrategyRequest) (*internal.V1GetStrategyResponse, error) {
	var resp internal.V1GetStrategyResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Ids(ctx context.Context, req *internal.V1ItemIDsRequest) (*internal.V1GetStrategyIDsByItemIDsResponse, error) {
	var resp internal.V1GetStrategyIDsByItemIDsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/strategy-ids-by-product-ids", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ItemsInfo(ctx context.Context, req *internal.V1GetStrategyItemInfoRequest) (*internal.V1GetStrategyItemInfoResponse, error) {
	var resp internal.V1GetStrategyItemInfoResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/product/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

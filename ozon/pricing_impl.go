package ozon

import "context"

func (s *PricingService) Ids(ctx context.Context, req *V1ItemIDsRequest) (*V1GetStrategyIDsByItemIDsResponse, error) {
	var resp V1GetStrategyIDsByItemIDsResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/strategy-ids-by-product-ids", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) ItemsAdd(ctx context.Context, req *V1AddStrategyItemsRequest) (*V1AddStrategyItemsResponse, error) {
	var resp V1AddStrategyItemsResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/products/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) ItemsList(ctx context.Context, req *V1StrategyRequest) (*V1GetStrategyItemsResponse, error) {
	var resp V1GetStrategyItemsResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) ItemsInfo(ctx context.Context, req *V1GetStrategyItemInfoRequest) (*V1GetStrategyItemInfoResponse, error) {
	var resp V1GetStrategyItemInfoResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/product/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) Delete(ctx context.Context, req *V1StrategyRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) Create(ctx context.Context, req *V1CreatePricingStrategyRequest) (*V1CreatePricingStrategyResponse, error) {
	var resp V1CreatePricingStrategyResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) Update(ctx context.Context, req *V1UpdatePricingStrategyRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) Info(ctx context.Context, req *V1StrategyRequest) (*V1GetStrategyResponse, error) {
	var resp V1GetStrategyResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) ItemsDelete(ctx context.Context, req *V1ItemIDsRequest) (*V1DeleteStrategyItemsResponse, error) {
	var resp V1DeleteStrategyItemsResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) List(ctx context.Context, req *V1GetStrategyListRequest) (*V1GetStrategyListResponse, error) {
	var resp V1GetStrategyListResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) Competitors(ctx context.Context, req *V1GetCompetitorsRequest) (*V1GetCompetitorsResponse, error) {
	var resp V1GetCompetitorsResponse
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/competitors/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricingService) Status(ctx context.Context, req *V1UpdateStatusStrategyRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v1/pricing-strategy/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

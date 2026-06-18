package pricing

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 策略ID列表
func (s *Service) Ids(ctx context.Context, req *V1ItemIDsRequest) (*V1GetStrategyIDsByItemIDsResponse, error) {
	var resp V1GetStrategyIDsByItemIDsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/strategy-ids-by-product-ids", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 竞争对手  的商品价格
func (s *Service) ItemsInfo(ctx context.Context, req *V1GetStrategyItemInfoRequest) (*V1GetStrategyItemInfoResponse, error) {
	var resp V1GetStrategyItemInfoResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/product/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建策略
func (s *Service) Create(ctx context.Context, req *V1CreatePricingStrategyRequest) (*V1CreatePricingStrategyResponse, error) {
	var resp V1CreatePricingStrategyResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新策略
func (s *Service) Update(ctx context.Context, req *V1UpdatePricingStrategyRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 策略中的商品列表
func (s *Service) ItemsList(ctx context.Context, req *V1StrategyRequest) (*V1GetStrategyItemsResponse, error) {
	var resp V1GetStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将商品添加到策略
func (s *Service) ItemsAdd(ctx context.Context, req *V1AddStrategyItemsRequest) (*V1AddStrategyItemsResponse, error) {
	var resp V1AddStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 策略列表
func (s *Service) List(ctx context.Context, req *V1GetStrategyListRequest) (*V1GetStrategyListResponse, error) {
	var resp V1GetStrategyListResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 从策略中删除商品
func (s *Service) ItemsDelete(ctx context.Context, req *V1ItemIDsRequest) (*V1DeleteStrategyItemsResponse, error) {
	var resp V1DeleteStrategyItemsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 策略信息
func (s *Service) Info(ctx context.Context, req *V1StrategyRequest) (*V1GetStrategyResponse, error) {
	var resp V1GetStrategyResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 竞争对手名单
func (s *Service) Competitors(ctx context.Context, req *V1GetCompetitorsRequest) (*V1GetCompetitorsResponse, error) {
	var resp V1GetCompetitorsResponse
	err := s.Client.Post(ctx, "/v1/pricing-strategy/competitors/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更改策略状态
func (s *Service) Status(ctx context.Context, req *V1UpdateStatusStrategyRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 删除策略
func (s *Service) Delete(ctx context.Context, req *V1StrategyRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v1/pricing-strategy/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

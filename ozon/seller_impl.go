package ozon

import "context"

func (s *SellerService) SellerOzonLogisticsInfo(ctx context.Context) (*V1SellerOzonLogisticsInfoResponse, error) {
	var resp V1SellerOzonLogisticsInfoResponse
	_, err := s.t.Post(ctx, "/v1/seller/ozon-logistics/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *SellerService) RolesByToken(ctx context.Context) (*V1RolesByTokenResponse, error) {
	var resp V1RolesByTokenResponse
	_, err := s.t.Post(ctx, "/v1/roles", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *SellerService) SellerInfo(ctx context.Context) (*V1SellerInfoResponse, error) {
	var resp V1SellerInfoResponse
	_, err := s.t.Post(ctx, "/v1/seller/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

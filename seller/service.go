package seller

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) RolesByToken(ctx context.Context) (*types.V1RolesByTokenResponse, error) {
	var resp types.V1RolesByTokenResponse
	err := s.Client.Post(ctx, "/v1/roles", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerInfo(ctx context.Context) (*types.V1SellerInfoResponse, error) {
	var resp types.V1SellerInfoResponse
	err := s.Client.Post(ctx, "/v1/seller/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerOzonLogisticsInfo(ctx context.Context) (*types.V1SellerOzonLogisticsInfoResponse, error) {
	var resp types.V1SellerOzonLogisticsInfoResponse
	err := s.Client.Post(ctx, "/v1/seller/ozon-logistics/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

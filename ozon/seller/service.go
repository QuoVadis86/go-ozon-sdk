package seller

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) RolesByToken(ctx context.Context) (*internal.V1RolesByTokenResponse, error) {
	var resp internal.V1RolesByTokenResponse
	err := s.Client.Post(ctx, "/v1/roles", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerInfo(ctx context.Context) (*internal.V1SellerInfoResponse, error) {
	var resp internal.V1SellerInfoResponse
	err := s.Client.Post(ctx, "/v1/seller/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerOzonLogisticsInfo(ctx context.Context) (*internal.V1SellerOzonLogisticsInfoResponse, error) {
	var resp internal.V1SellerOzonLogisticsInfoResponse
	err := s.Client.Post(ctx, "/v1/seller/ozon-logistics/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

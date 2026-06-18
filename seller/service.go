package seller

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

// 卖家个人中心信息
func (s *Service) SellerInfo(ctx context.Context) (*V1SellerInfoResponse, error) {
	var resp V1SellerInfoResponse
	err := s.Client.Post(ctx, "/v1/seller/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Ozon配送开通信息
func (s *Service) SellerOzonLogisticsInfo(ctx context.Context) (*V1SellerOzonLogisticsInfoResponse, error) {
	var resp V1SellerOzonLogisticsInfoResponse
	err := s.Client.Post(ctx, "/v1/seller/ozon-logistics/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 使用API密钥获取角色和方式列表
func (s *Service) RolesByToken(ctx context.Context) (*V1RolesByTokenResponse, error) {
	var resp V1RolesByTokenResponse
	err := s.Client.Post(ctx, "/v1/roles", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

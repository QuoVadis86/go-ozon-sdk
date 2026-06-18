package beta

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

// 获取按数量折扣信息
func (s *Service) GetProductStairwayDiscountByQuantity(ctx context.Context, req *V1GetProductStairwayDiscountByQuantityRequest) (*V1GetProductStairwayDiscountByQuantityResponse, error) {
	var resp V1GetProductStairwayDiscountByQuantityResponse
	err := s.Client.Post(ctx, "/v1/product/stairway-discount/by-quantity/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取样件添加状态
func (s *Service) FbsPostingProductExemplarStatusV5(ctx context.Context, req *V5FbsPostingProductExemplarStatusV5Request) (*V5FbsPostingProductExemplarStatusV5Response, error) {
	var resp V5FbsPostingProductExemplarStatusV5Response
	err := s.Client.Post(ctx, "/v5/fbs/posting/product/exemplar/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取按货件统计的应计项目
func (s *Service) GetFinanceAccrualPostings(ctx context.Context, req *FinanceV1GetFinanceAccrualPostingsRequest) (*FinanceV1GetFinanceAccrualPostingsResponse, error) {
	var resp FinanceV1GetFinanceAccrualPostingsResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 标志代码验证
func (s *Service) FbsPostingProductExemplarValidateV5(ctx context.Context, req *V5FbsPostingProductExemplarValidateV5Request) (*V5FbsPostingProductExemplarValidateV5Response, error) {
	var resp V5FbsPostingProductExemplarValidateV5Response
	err := s.Client.Post(ctx, "/v5/fbs/posting/product/exemplar/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取某日应计项目
func (s *Service) GetFinanceAccrualByDay(ctx context.Context, req *FinanceV1GetFinanceAccrualByDayRequest) (*FinanceV1GetFinanceAccrualByDayResponse, error) {
	var resp FinanceV1GetFinanceAccrualByDayResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取已创建样件数据
func (s *Service) FbsPostingProductExemplarCreateOrGetV6(ctx context.Context, req *V6FbsPostingProductExemplarCreateOrGetV6Request) (*V6FbsPostingProductExemplarCreateOrGetV6Response, error) {
	var resp V6FbsPostingProductExemplarCreateOrGetV6Response
	err := s.Client.Post(ctx, "/v6/fbs/posting/product/exemplar/create-or-get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取用于确定商品类目的提示
func (s *Service) DescriptionCategoryTips(ctx context.Context, req *V1DescriptionCategoryTipsRequest) (*V1DescriptionCategoryTipsResponse, error) {
	var resp V1DescriptionCategoryTipsResponse
	err := s.Client.Post(ctx, "/v1/description-category/tips", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取商品可见性信息
func (s *Service) ProductVisibilityInfo(ctx context.Context, req *ProductV1ProductVisibilityInfoRequest) (*ProductV1ProductVisibilityInfoResponse, error) {
	var resp ProductV1ProductVisibilityInfoResponse
	err := s.Client.Post(ctx, "/v1/product/visibility/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取应计项目参考信息
func (s *Service) GetFinanceAccrualTypes(ctx context.Context) (*FinanceV1GetFinanceAccrualTypesResponse, error) {
	var resp FinanceV1GetFinanceAccrualTypesResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/types", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取折扣申请列表
func (s *Service) GetDiscountTaskListV2(ctx context.Context, req *V2GetDiscountTaskListV2Request) (*V2GetDiscountTaskListV2Response, error) {
	var resp V2GetDiscountTaskListV2Response
	err := s.Client.Post(ctx, "/v2/actions/discounts-task/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 新增了用于设置商品在Ozon和Ozon Select橱窗可见性的Beta方法。
func (s *Service) ProductVisibilitySet(ctx context.Context, req *ProductV1ProductVisibilitySetRequest) (*ProductV1ProductVisibilitySetResponse, error) {
	var resp ProductV1ProductVisibilitySetResponse
	err := s.Client.Post(ctx, "/v1/product/visibility/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 管理按数量折扣
func (s *Service) SetProductStairwayDiscountByQuantity(ctx context.Context, req *V1SetProductStairwayDiscountByQuantityRequest) (*V1SetProductStairwayDiscountByQuantityResponse, error) {
	var resp V1SetProductStairwayDiscountByQuantityResponse
	err := s.Client.Post(ctx, "/v1/product/stairway-discount/by-quantity/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取余额报告
func (s *Service) GetFinanceBalanceV1(ctx context.Context, req *V1GetFinanceBalanceV1Request) (*V1GetFinanceBalanceV1Response, error) {
	var resp V1GetFinanceBalanceV1Response
	err := s.Client.Post(ctx, "/v1/finance/balance", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Обновить данные экземпляров
func (s *Service) FbsPostingProductExemplarUpdate(ctx context.Context, req *V1FbsPostingProductExemplarUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/fbs/posting/product/exemplar/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 检查并保存份数数据
func (s *Service) FbsPostingProductExemplarSetV6(ctx context.Context, req *V6FbsPostingProductExemplarSetV6Request) error {
	err := s.Client.Post(ctx, "/v6/fbs/posting/product/exemplar/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 将订单拆分为不带备货的货件
func (s *Service) FbsSplit(ctx context.Context) (*V1PostingFbsSplitResponse, error) {
	var resp V1PostingFbsSplitResponse
	err := s.Client.Post(ctx, "/v1/posting/fbs/split", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取FBS和rFBS仓库库存信息
func (s *Service) ProductInfoWarehouseStocks(ctx context.Context, req *V1ProductInfoWarehouseStocksRequest) (*V1ProductInfoWarehouseStocksResponse, error) {
	var resp V1ProductInfoWarehouseStocksResponse
	err := s.Client.Post(ctx, "/v1/product/info/warehouse/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

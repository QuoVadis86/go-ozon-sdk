package beta

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *internal.Client }

func (s *Service) FbsPostingProductExemplarSetV6(ctx context.Context, req *types.V6FbsPostingProductExemplarSetV6Request) error {
	err := s.Client.Post(ctx, "/v6/fbs/posting/product/exemplar/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ProductVisibilitySet(ctx context.Context, req *types.ProductV1ProductVisibilitySetRequest) (*types.ProductV1ProductVisibilitySetResponse, error) {
	var resp types.ProductV1ProductVisibilitySetResponse
	err := s.Client.Post(ctx, "/v1/product/visibility/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductVisibilityInfo(ctx context.Context, req *types.ProductV1ProductVisibilityInfoRequest) (*types.ProductV1ProductVisibilityInfoResponse, error) {
	var resp types.ProductV1ProductVisibilityInfoResponse
	err := s.Client.Post(ctx, "/v1/product/visibility/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarUpdate(ctx context.Context, req *types.V1FbsPostingProductExemplarUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/fbs/posting/product/exemplar/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FbsSplit(ctx context.Context) (*types.V1PostingFbsSplitResponse, error) {
	var resp types.V1PostingFbsSplitResponse
	err := s.Client.Post(ctx, "/v1/posting/fbs/split", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SetProductStairwayDiscountByQuantity(ctx context.Context, req *types.V1SetProductStairwayDiscountByQuantityRequest) (*types.V1SetProductStairwayDiscountByQuantityResponse, error) {
	var resp types.V1SetProductStairwayDiscountByQuantityResponse
	err := s.Client.Post(ctx, "/v1/product/stairway-discount/by-quantity/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFinanceBalanceV1(ctx context.Context, req *types.V1GetFinanceBalanceV1Request) (*types.V1GetFinanceBalanceV1Response, error) {
	var resp types.V1GetFinanceBalanceV1Response
	err := s.Client.Post(ctx, "/v1/finance/balance", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFinanceAccrualByDay(ctx context.Context, req *types.FinanceV1GetFinanceAccrualByDayRequest) (*types.FinanceV1GetFinanceAccrualByDayResponse, error) {
	var resp types.FinanceV1GetFinanceAccrualByDayResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DescriptionCategoryTips(ctx context.Context, req *types.V1DescriptionCategoryTipsRequest) (*types.V1DescriptionCategoryTipsResponse, error) {
	var resp types.V1DescriptionCategoryTipsResponse
	err := s.Client.Post(ctx, "/v1/description-category/tips", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductStairwayDiscountByQuantity(ctx context.Context, req *types.V1GetProductStairwayDiscountByQuantityRequest) (*types.V1GetProductStairwayDiscountByQuantityResponse, error) {
	var resp types.V1GetProductStairwayDiscountByQuantityResponse
	err := s.Client.Post(ctx, "/v1/product/stairway-discount/by-quantity/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarStatusV5(ctx context.Context, req *types.V5FbsPostingProductExemplarStatusV5Request) (*types.V5FbsPostingProductExemplarStatusV5Response, error) {
	var resp types.V5FbsPostingProductExemplarStatusV5Response
	err := s.Client.Post(ctx, "/v5/fbs/posting/product/exemplar/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarCreateOrGetV6(ctx context.Context, req *types.V6FbsPostingProductExemplarCreateOrGetV6Request) (*types.V6FbsPostingProductExemplarCreateOrGetV6Response, error) {
	var resp types.V6FbsPostingProductExemplarCreateOrGetV6Response
	err := s.Client.Post(ctx, "/v6/fbs/posting/product/exemplar/create-or-get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarValidateV5(ctx context.Context, req *types.V5FbsPostingProductExemplarValidateV5Request) (*types.V5FbsPostingProductExemplarValidateV5Response, error) {
	var resp types.V5FbsPostingProductExemplarValidateV5Response
	err := s.Client.Post(ctx, "/v5/fbs/posting/product/exemplar/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFinanceAccrualTypes(ctx context.Context) (*types.FinanceV1GetFinanceAccrualTypesResponse, error) {
	var resp types.FinanceV1GetFinanceAccrualTypesResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/types", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductInfoWarehouseStocks(ctx context.Context, req *types.V1ProductInfoWarehouseStocksRequest) (*types.V1ProductInfoWarehouseStocksResponse, error) {
	var resp types.V1ProductInfoWarehouseStocksResponse
	err := s.Client.Post(ctx, "/v1/product/info/warehouse/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetDiscountTaskListV2(ctx context.Context, req *types.V2GetDiscountTaskListV2Request) (*types.V2GetDiscountTaskListV2Response, error) {
	var resp types.V2GetDiscountTaskListV2Response
	err := s.Client.Post(ctx, "/v2/actions/discounts-task/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFinanceAccrualPostings(ctx context.Context, req *types.FinanceV1GetFinanceAccrualPostingsRequest) (*types.FinanceV1GetFinanceAccrualPostingsResponse, error) {
	var resp types.FinanceV1GetFinanceAccrualPostingsResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

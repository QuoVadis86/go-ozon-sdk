package ozon

import "context"

func (s *BetaService) FbsPostingProductExemplarStatusV5(ctx context.Context, req *V5FbsPostingProductExemplarStatusV5Request) (*V5FbsPostingProductExemplarStatusV5Response, error) {
	var resp V5FbsPostingProductExemplarStatusV5Response
	_, err := s.t.Post(ctx, "/v5/fbs/posting/product/exemplar/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) SetProductStairwayDiscountByQuantity(ctx context.Context, req *V1SetProductStairwayDiscountByQuantityRequest) (*V1SetProductStairwayDiscountByQuantityResponse, error) {
	var resp V1SetProductStairwayDiscountByQuantityResponse
	_, err := s.t.Post(ctx, "/v1/product/stairway-discount/by-quantity/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) DescriptionCategoryTips(ctx context.Context, req *V1DescriptionCategoryTipsRequest) (*V1DescriptionCategoryTipsResponse, error) {
	var resp V1DescriptionCategoryTipsResponse
	_, err := s.t.Post(ctx, "/v1/description-category/tips", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) GetFinanceBalanceV1(ctx context.Context, req *V1GetFinanceBalanceV1Request) (*V1GetFinanceBalanceV1Response, error) {
	var resp V1GetFinanceBalanceV1Response
	_, err := s.t.Post(ctx, "/v1/finance/balance", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) ProductVisibilityInfo(ctx context.Context, req *ProductV1ProductVisibilityInfoRequest) (*ProductV1ProductVisibilityInfoResponse, error) {
	var resp ProductV1ProductVisibilityInfoResponse
	_, err := s.t.Post(ctx, "/v1/product/visibility/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) FbsPostingProductExemplarSetV6(ctx context.Context, req *V6FbsPostingProductExemplarSetV6Request) error {
	_, err := s.t.Post(ctx, "/v6/fbs/posting/product/exemplar/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *BetaService) GetFinanceAccrualByDay(ctx context.Context, req *FinanceV1GetFinanceAccrualByDayRequest) (*FinanceV1GetFinanceAccrualByDayResponse, error) {
	var resp FinanceV1GetFinanceAccrualByDayResponse
	_, err := s.t.Post(ctx, "/v1/finance/accrual/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) FbsPostingProductExemplarCreateOrGetV6(ctx context.Context, req *V6FbsPostingProductExemplarCreateOrGetV6Request) (*V6FbsPostingProductExemplarCreateOrGetV6Response, error) {
	var resp V6FbsPostingProductExemplarCreateOrGetV6Response
	_, err := s.t.Post(ctx, "/v6/fbs/posting/product/exemplar/create-or-get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) FbsPostingProductExemplarValidateV5(ctx context.Context, req *V5FbsPostingProductExemplarValidateV5Request) (*V5FbsPostingProductExemplarValidateV5Response, error) {
	var resp V5FbsPostingProductExemplarValidateV5Response
	_, err := s.t.Post(ctx, "/v5/fbs/posting/product/exemplar/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) ProductVisibilitySet(ctx context.Context, req *ProductV1ProductVisibilitySetRequest) (*ProductV1ProductVisibilitySetResponse, error) {
	var resp ProductV1ProductVisibilitySetResponse
	_, err := s.t.Post(ctx, "/v1/product/visibility/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) GetDiscountTaskListV2(ctx context.Context, req *V2GetDiscountTaskListV2Request) (*V2GetDiscountTaskListV2Response, error) {
	var resp V2GetDiscountTaskListV2Response
	_, err := s.t.Post(ctx, "/v2/actions/discounts-task/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) GetFinanceAccrualTypes(ctx context.Context) (*FinanceV1GetFinanceAccrualTypesResponse, error) {
	var resp FinanceV1GetFinanceAccrualTypesResponse
	_, err := s.t.Post(ctx, "/v1/finance/accrual/types", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) ProductInfoWarehouseStocks(ctx context.Context, req *V1ProductInfoWarehouseStocksRequest) (*V1ProductInfoWarehouseStocksResponse, error) {
	var resp V1ProductInfoWarehouseStocksResponse
	_, err := s.t.Post(ctx, "/v1/product/info/warehouse/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) GetProductStairwayDiscountByQuantity(ctx context.Context, req *V1GetProductStairwayDiscountByQuantityRequest) (*V1GetProductStairwayDiscountByQuantityResponse, error) {
	var resp V1GetProductStairwayDiscountByQuantityResponse
	_, err := s.t.Post(ctx, "/v1/product/stairway-discount/by-quantity/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) FbsSplit(ctx context.Context) (*V1PostingFbsSplitResponse, error) {
	var resp V1PostingFbsSplitResponse
	_, err := s.t.Post(ctx, "/v1/posting/fbs/split", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BetaService) FbsPostingProductExemplarUpdate(ctx context.Context, req *V1FbsPostingProductExemplarUpdateRequest) error {
	_, err := s.t.Post(ctx, "/v1/fbs/posting/product/exemplar/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *BetaService) GetFinanceAccrualPostings(ctx context.Context, req *FinanceV1GetFinanceAccrualPostingsRequest) (*FinanceV1GetFinanceAccrualPostingsResponse, error) {
	var resp FinanceV1GetFinanceAccrualPostingsResponse
	_, err := s.t.Post(ctx, "/v1/finance/accrual/postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

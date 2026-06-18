package beta

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) FbsPostingProductExemplarValidateV5(ctx context.Context, req *internal.V5FbsPostingProductExemplarValidateV5Request) (*internal.V5FbsPostingProductExemplarValidateV5Response, error) {
	var resp internal.V5FbsPostingProductExemplarValidateV5Response
	err := s.Client.Post(ctx, "/v5/fbs/posting/product/exemplar/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SetProductStairwayDiscountByQuantity(ctx context.Context, req *internal.V1SetProductStairwayDiscountByQuantityRequest) (*internal.V1SetProductStairwayDiscountByQuantityResponse, error) {
	var resp internal.V1SetProductStairwayDiscountByQuantityResponse
	err := s.Client.Post(ctx, "/v1/product/stairway-discount/by-quantity/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFinanceBalanceV1(ctx context.Context, req *internal.V1GetFinanceBalanceV1Request) (*internal.V1GetFinanceBalanceV1Response, error) {
	var resp internal.V1GetFinanceBalanceV1Response
	err := s.Client.Post(ctx, "/v1/finance/balance", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFinanceAccrualPostings(ctx context.Context, req *internal.FinanceV1GetFinanceAccrualPostingsRequest) (*internal.FinanceV1GetFinanceAccrualPostingsResponse, error) {
	var resp internal.FinanceV1GetFinanceAccrualPostingsResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DescriptionCategoryTips(ctx context.Context, req *internal.V1DescriptionCategoryTipsRequest) (*internal.V1DescriptionCategoryTipsResponse, error) {
	var resp internal.V1DescriptionCategoryTipsResponse
	err := s.Client.Post(ctx, "/v1/description-category/tips", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetDiscountTaskListV2(ctx context.Context, req *internal.V2GetDiscountTaskListV2Request) (*internal.V2GetDiscountTaskListV2Response, error) {
	var resp internal.V2GetDiscountTaskListV2Response
	err := s.Client.Post(ctx, "/v2/actions/discounts-task/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductVisibilityInfo(ctx context.Context, req *internal.ProductV1ProductVisibilityInfoRequest) (*internal.ProductV1ProductVisibilityInfoResponse, error) {
	var resp internal.ProductV1ProductVisibilityInfoResponse
	err := s.Client.Post(ctx, "/v1/product/visibility/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductVisibilitySet(ctx context.Context, req *internal.ProductV1ProductVisibilitySetRequest) (*internal.ProductV1ProductVisibilitySetResponse, error) {
	var resp internal.ProductV1ProductVisibilitySetResponse
	err := s.Client.Post(ctx, "/v1/product/visibility/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductStairwayDiscountByQuantity(ctx context.Context, req *internal.V1GetProductStairwayDiscountByQuantityRequest) (*internal.V1GetProductStairwayDiscountByQuantityResponse, error) {
	var resp internal.V1GetProductStairwayDiscountByQuantityResponse
	err := s.Client.Post(ctx, "/v1/product/stairway-discount/by-quantity/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarSetV6(ctx context.Context, req *internal.V6FbsPostingProductExemplarSetV6Request) error {
	err := s.Client.Post(ctx, "/v6/fbs/posting/product/exemplar/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetFinanceAccrualByDay(ctx context.Context, req *internal.FinanceV1GetFinanceAccrualByDayRequest) (*internal.FinanceV1GetFinanceAccrualByDayResponse, error) {
	var resp internal.FinanceV1GetFinanceAccrualByDayResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarStatusV5(ctx context.Context, req *internal.V5FbsPostingProductExemplarStatusV5Request) (*internal.V5FbsPostingProductExemplarStatusV5Response, error) {
	var resp internal.V5FbsPostingProductExemplarStatusV5Response
	err := s.Client.Post(ctx, "/v5/fbs/posting/product/exemplar/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarCreateOrGetV6(ctx context.Context, req *internal.V6FbsPostingProductExemplarCreateOrGetV6Request) (*internal.V6FbsPostingProductExemplarCreateOrGetV6Response, error) {
	var resp internal.V6FbsPostingProductExemplarCreateOrGetV6Response
	err := s.Client.Post(ctx, "/v6/fbs/posting/product/exemplar/create-or-get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFinanceAccrualTypes(ctx context.Context) (*internal.FinanceV1GetFinanceAccrualTypesResponse, error) {
	var resp internal.FinanceV1GetFinanceAccrualTypesResponse
	err := s.Client.Post(ctx, "/v1/finance/accrual/types", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingProductExemplarUpdate(ctx context.Context, req *internal.V1FbsPostingProductExemplarUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/fbs/posting/product/exemplar/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FbsSplit(ctx context.Context) (*internal.V1PostingFbsSplitResponse, error) {
	var resp internal.V1PostingFbsSplitResponse
	err := s.Client.Post(ctx, "/v1/posting/fbs/split", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductInfoWarehouseStocks(ctx context.Context, req *internal.V1ProductInfoWarehouseStocksRequest) (*internal.V1ProductInfoWarehouseStocksResponse, error) {
	var resp internal.V1ProductInfoWarehouseStocksResponse
	err := s.Client.Post(ctx, "/v1/product/info/warehouse/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

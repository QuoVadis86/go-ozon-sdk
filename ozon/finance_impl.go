package ozon

import "context"

func (s *FinanceService) GetDecompensationReport(ctx context.Context, req *V1GetDecompensationReportRequest) (*CreateReportResponse, error) {
	var resp CreateReportResponse
	_, err := s.t.Post(ctx, "/v1/finance/decompensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FinanceService) FinanceTransactionListV3(ctx context.Context, req *Financev3FinanceTransactionListV3Request) (*Financev3FinanceTransactionListV3Response, error) {
	var resp Financev3FinanceTransactionListV3Response
	_, err := s.t.Post(ctx, "/v3/finance/transaction/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FinanceService) GetCompensationReport(ctx context.Context, req *V1GetCompensationReportRequest) (*CreateReportResponse, error) {
	var resp CreateReportResponse
	_, err := s.t.Post(ctx, "/v1/finance/compensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FinanceService) GetRealizationReportV2(ctx context.Context, req *V2GetRealizationReportRequestV2) (*V2GetRealizationReportResponseV2, error) {
	var resp V2GetRealizationReportResponseV2
	_, err := s.t.Post(ctx, "/v2/finance/realization", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FinanceService) GetRealizationReportV1(ctx context.Context, req *V1GetRealizationReportPostingRequest) (*V1GetRealizationReportPostingResponse, error) {
	var resp V1GetRealizationReportPostingResponse
	_, err := s.t.Post(ctx, "/v1/finance/realization/posting", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FinanceService) FinanceTransactionTotalV3(ctx context.Context, req *Financev3FinanceTransactionTotalsV3Request) (*Financev3FinanceTransactionTotalsV3Response, error) {
	var resp Financev3FinanceTransactionTotalsV3Response
	_, err := s.t.Post(ctx, "/v3/finance/transaction/totals", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

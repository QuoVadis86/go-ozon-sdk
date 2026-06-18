package finance

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *internal.Client }

func (s *Service) FinanceTransactionListV3(ctx context.Context, req *types.Financev3FinanceTransactionListV3Request) (*types.Financev3FinanceTransactionListV3Response, error) {
	var resp types.Financev3FinanceTransactionListV3Response
	err := s.Client.Post(ctx, "/v3/finance/transaction/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetDecompensationReport(ctx context.Context, req *types.V1GetDecompensationReportRequest) (*types.CreateReportResponse, error) {
	var resp types.CreateReportResponse
	err := s.Client.Post(ctx, "/v1/finance/decompensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetRealizationReportV2(ctx context.Context, req *types.V2GetRealizationReportRequestV2) (*types.V2GetRealizationReportResponseV2, error) {
	var resp types.V2GetRealizationReportResponseV2
	err := s.Client.Post(ctx, "/v2/finance/realization", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FinanceTransactionTotalV3(ctx context.Context, req *types.Financev3FinanceTransactionTotalsV3Request) (*types.Financev3FinanceTransactionTotalsV3Response, error) {
	var resp types.Financev3FinanceTransactionTotalsV3Response
	err := s.Client.Post(ctx, "/v3/finance/transaction/totals", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetCompensationReport(ctx context.Context, req *types.V1GetCompensationReportRequest) (*types.CreateReportResponse, error) {
	var resp types.CreateReportResponse
	err := s.Client.Post(ctx, "/v1/finance/compensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetRealizationReportV1(ctx context.Context, req *types.V1GetRealizationReportPostingRequest) (*types.V1GetRealizationReportPostingResponse, error) {
	var resp types.V1GetRealizationReportPostingResponse
	err := s.Client.Post(ctx, "/v1/finance/realization/posting", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

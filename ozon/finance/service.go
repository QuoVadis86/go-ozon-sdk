package finance

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) GetRealizationReportV1(ctx context.Context, req *internal.V1GetRealizationReportPostingRequest) (*internal.V1GetRealizationReportPostingResponse, error) {
	var resp internal.V1GetRealizationReportPostingResponse
	err := s.Client.Post(ctx, "/v1/finance/realization/posting", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FinanceTransactionTotalV3(ctx context.Context, req *internal.Financev3FinanceTransactionTotalsV3Request) (*internal.Financev3FinanceTransactionTotalsV3Response, error) {
	var resp internal.Financev3FinanceTransactionTotalsV3Response
	err := s.Client.Post(ctx, "/v3/finance/transaction/totals", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FinanceTransactionListV3(ctx context.Context, req *internal.Financev3FinanceTransactionListV3Request) (*internal.Financev3FinanceTransactionListV3Response, error) {
	var resp internal.Financev3FinanceTransactionListV3Response
	err := s.Client.Post(ctx, "/v3/finance/transaction/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetCompensationReport(ctx context.Context, req *internal.V1GetCompensationReportRequest) (*internal.CreateReportResponse, error) {
	var resp internal.CreateReportResponse
	err := s.Client.Post(ctx, "/v1/finance/compensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetRealizationReportV2(ctx context.Context, req *internal.V2GetRealizationReportRequestV2) (*internal.V2GetRealizationReportResponseV2, error) {
	var resp internal.V2GetRealizationReportResponseV2
	err := s.Client.Post(ctx, "/v2/finance/realization", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetDecompensationReport(ctx context.Context, req *internal.V1GetDecompensationReportRequest) (*internal.CreateReportResponse, error) {
	var resp internal.CreateReportResponse
	err := s.Client.Post(ctx, "/v1/finance/decompensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

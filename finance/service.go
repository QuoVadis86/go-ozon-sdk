package finance

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

// 按订单细分的商品销售报告
// Note: 每笔订单的详细信息
func (s *Service) GetRealizationReportV1(ctx context.Context, req *V1GetRealizationReportPostingRequest) (*V1GetRealizationReportPostingResponse, error) {
	var resp V1GetRealizationReportPostingResponse
	err := s.Client.Post(ctx, "/v1/finance/realization/posting", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: use /v1/finance/accrual/postings instead
// 清单数目
func (s *Service) FinanceTransactionTotalV3(ctx context.Context, req *Financev3FinanceTransactionTotalsV3Request) (*Financev3FinanceTransactionTotalsV3Response, error) {
	var resp Financev3FinanceTransactionTotalsV3Response
	err := s.Client.Post(ctx, "/v3/finance/transaction/totals", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 赔偿返还报告
func (s *Service) GetDecompensationReport(ctx context.Context, req *V1GetDecompensationReportRequest) (*CreateReportResponse, error) {
	var resp CreateReportResponse
	err := s.Client.Post(ctx, "/v1/finance/decompensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 商品销售报告 （第2版）
func (s *Service) GetRealizationReportV2(ctx context.Context, req *V2GetRealizationReportRequestV2) (*V2GetRealizationReportResponseV2, error) {
	var resp V2GetRealizationReportResponseV2
	err := s.Client.Post(ctx, "/v2/finance/realization", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: use /v1/finance/accrual/postings instead
// 交易清单
// Note: 请使用顺序发送请求的方式
func (s *Service) FinanceTransactionListV3(ctx context.Context, req *Financev3FinanceTransactionListV3Request) (*Financev3FinanceTransactionListV3Response, error) {
	var resp Financev3FinanceTransactionListV3Response
	err := s.Client.Post(ctx, "/v3/finance/transaction/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 赔偿报告
func (s *Service) GetCompensationReport(ctx context.Context, req *V1GetCompensationReportRequest) (*CreateReportResponse, error) {
	var resp CreateReportResponse
	err := s.Client.Post(ctx, "/v1/finance/compensation", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

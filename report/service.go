package report

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *internal.Client }

func (s *Service) CreateCompanyProductsReport(ctx context.Context, req *types.ReportCreateCompanyProductsReportRequest) (*types.ReportCreateReportResponse, error) {
	var resp types.ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/products/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FinanceCashFlowStatementList(ctx context.Context, req *types.V3FinanceCashFlowStatementListRequest) (*types.V3FinanceCashFlowStatementListResponse, error) {
	var resp types.V3FinanceCashFlowStatementListResponse
	err := s.Client.Post(ctx, "/v1/finance/cash-flow-statement/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateCompanyPostingsReport(ctx context.Context, req *types.ReportCreateCompanyPostingsReportRequest) (*types.ReportCreateReportResponse, error) {
	var resp types.ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/postings/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReportInfo(ctx context.Context, req *types.ReportReportInfoRequest) (*types.ReportReportInfoResponse, error) {
	var resp types.ReportReportInfoResponse
	err := s.Client.Post(ctx, "/v1/report/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateCompanyMarkedProductsSalesReport(ctx context.Context, req *types.V1ReportMarkedProductsSalesCreateRequest) (*types.CommonCreateReportResponse, error) {
	var resp types.CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/marked-products-sales/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateStockByWarehouseReport(ctx context.Context, req *types.V1CreateStockByWarehouseReportRequest) (*types.CommonCreateReportResponse, error) {
	var resp types.CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/warehouse/stock", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReportList(ctx context.Context, req *types.ReportReportListRequest) (*types.ReportReportListResponse, error) {
	var resp types.ReportReportListResponse
	err := s.Client.Post(ctx, "/v1/report/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateDiscountedReport(ctx context.Context, req *types.ReportCreateDiscountedRequest) (*types.ReportCreateDiscountedResponse, error) {
	var resp types.ReportCreateDiscountedResponse
	err := s.Client.Post(ctx, "/v1/report/discounted/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

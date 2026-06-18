package report

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) CreateCompanyProductsReport(ctx context.Context, req *internal.ReportCreateCompanyProductsReportRequest) (*internal.ReportCreateReportResponse, error) {
	var resp internal.ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/products/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateCompanyMarkedProductsSalesReport(ctx context.Context, req *internal.V1ReportMarkedProductsSalesCreateRequest) (*internal.CommonCreateReportResponse, error) {
	var resp internal.CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/marked-products-sales/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReportList(ctx context.Context, req *internal.ReportReportListRequest) (*internal.ReportReportListResponse, error) {
	var resp internal.ReportReportListResponse
	err := s.Client.Post(ctx, "/v1/report/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateCompanyPostingsReport(ctx context.Context, req *internal.ReportCreateCompanyPostingsReportRequest) (*internal.ReportCreateReportResponse, error) {
	var resp internal.ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/postings/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateDiscountedReport(ctx context.Context, req *internal.ReportCreateDiscountedRequest) (*internal.ReportCreateDiscountedResponse, error) {
	var resp internal.ReportCreateDiscountedResponse
	err := s.Client.Post(ctx, "/v1/report/discounted/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReportInfo(ctx context.Context, req *internal.ReportReportInfoRequest) (*internal.ReportReportInfoResponse, error) {
	var resp internal.ReportReportInfoResponse
	err := s.Client.Post(ctx, "/v1/report/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateStockByWarehouseReport(ctx context.Context, req *internal.V1CreateStockByWarehouseReportRequest) (*internal.CommonCreateReportResponse, error) {
	var resp internal.CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/warehouse/stock", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FinanceCashFlowStatementList(ctx context.Context, req *internal.V3FinanceCashFlowStatementListRequest) (*internal.V3FinanceCashFlowStatementListResponse, error) {
	var resp internal.V3FinanceCashFlowStatementListResponse
	err := s.Client.Post(ctx, "/v1/finance/cash-flow-statement/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package report

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

func (s *Service) CreateCompanyProductsReport(ctx context.Context, req *ReportCreateCompanyProductsReportRequest) (*ReportCreateReportResponse, error) {
	var resp ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/products/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateDiscountedReport(ctx context.Context, req *ReportCreateDiscountedRequest) (*ReportCreateDiscountedResponse, error) {
	var resp ReportCreateDiscountedResponse
	err := s.Client.Post(ctx, "/v1/report/discounted/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateCompanyMarkedProductsSalesReport(ctx context.Context, req *V1ReportMarkedProductsSalesCreateRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/marked-products-sales/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FinanceCashFlowStatementList(ctx context.Context, req *V3FinanceCashFlowStatementListRequest) (*V3FinanceCashFlowStatementListResponse, error) {
	var resp V3FinanceCashFlowStatementListResponse
	err := s.Client.Post(ctx, "/v1/finance/cash-flow-statement/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateStockByWarehouseReport(ctx context.Context, req *V1CreateStockByWarehouseReportRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/warehouse/stock", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReportList(ctx context.Context, req *ReportReportListRequest) (*ReportReportListResponse, error) {
	var resp ReportReportListResponse
	err := s.Client.Post(ctx, "/v1/report/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReportInfo(ctx context.Context, req *ReportReportInfoRequest) (*ReportReportInfoResponse, error) {
	var resp ReportReportInfoResponse
	err := s.Client.Post(ctx, "/v1/report/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateCompanyPostingsReport(ctx context.Context, req *ReportCreateCompanyPostingsReportRequest) (*ReportCreateReportResponse, error) {
	var resp ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/postings/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

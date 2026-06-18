package ozon

import "context"

func (s *ReportService) CreateCompanyPostingsReport(ctx context.Context, req *ReportCreateCompanyPostingsReportRequest) (*ReportCreateReportResponse, error) {
	var resp ReportCreateReportResponse
	_, err := s.t.Post(ctx, "/v1/report/postings/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReportService) CreateStockByWarehouseReport(ctx context.Context, req *V1CreateStockByWarehouseReportRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	_, err := s.t.Post(ctx, "/v1/report/warehouse/stock", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReportService) ReportList(ctx context.Context, req *ReportReportListRequest) (*ReportReportListResponse, error) {
	var resp ReportReportListResponse
	_, err := s.t.Post(ctx, "/v1/report/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReportService) CreateDiscountedReport(ctx context.Context, req *ReportCreateDiscountedRequest) (*ReportCreateDiscountedResponse, error) {
	var resp ReportCreateDiscountedResponse
	_, err := s.t.Post(ctx, "/v1/report/discounted/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReportService) FinanceCashFlowStatementList(ctx context.Context, req *V3FinanceCashFlowStatementListRequest) (*V3FinanceCashFlowStatementListResponse, error) {
	var resp V3FinanceCashFlowStatementListResponse
	_, err := s.t.Post(ctx, "/v1/finance/cash-flow-statement/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReportService) ReportInfo(ctx context.Context, req *ReportReportInfoRequest) (*ReportReportInfoResponse, error) {
	var resp ReportReportInfoResponse
	_, err := s.t.Post(ctx, "/v1/report/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReportService) CreateCompanyProductsReport(ctx context.Context, req *ReportCreateCompanyProductsReportRequest) (*ReportCreateReportResponse, error) {
	var resp ReportCreateReportResponse
	_, err := s.t.Post(ctx, "/v1/report/products/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReportService) CreateCompanyMarkedProductsSalesReport(ctx context.Context, req *V1ReportMarkedProductsSalesCreateRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	_, err := s.t.Post(ctx, "/v1/report/marked-products-sales/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

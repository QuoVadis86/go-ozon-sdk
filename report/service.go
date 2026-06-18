package report

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

// 商品报告
// Note: 不能购买该商品
func (s *Service) CreateCompanyProductsReport(ctx context.Context, req *ReportCreateCompanyProductsReportRequest) (*ReportCreateReportResponse, error) {
	var resp ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/products/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 报告信息
func (s *Service) ReportInfo(ctx context.Context, req *ReportReportInfoRequest) (*ReportReportInfoResponse, error) {
	var resp ReportReportInfoResponse
	err := s.Client.Post(ctx, "/v1/report/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 关于FBS仓库库存报告
// Note: 要获取报告，请在 [/v1/report/info](#operation/ReportAPI_ReportInfo) 方法的请求中发送ID
func (s *Service) CreateStockByWarehouseReport(ctx context.Context, req *V1CreateStockByWarehouseReportRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/warehouse/stock", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 生成带有标记商品的销售报告
// Note: 每个报告最多可包含 50,000 个商品标签代码
// Note: 如需获取其余数据，请缩短报告生成周期
func (s *Service) CreateCompanyMarkedProductsSalesReport(ctx context.Context, req *V1ReportMarkedProductsSalesCreateRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/marked-products-sales/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 财务报告
func (s *Service) FinanceCashFlowStatementList(ctx context.Context, req *V3FinanceCashFlowStatementListRequest) (*V3FinanceCashFlowStatementListResponse, error) {
	var resp V3FinanceCashFlowStatementListResponse
	err := s.Client.Post(ctx, "/v1/finance/cash-flow-statement/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 减价商品报告
// Note: 每分钟可以发送1次请求
// Note: 要获取报告，请在 [/v1/report/info](#operation/ReportAPI_ReportInfo) 方法请求中发送ID
func (s *Service) CreateDiscountedReport(ctx context.Context, req *ReportCreateDiscountedRequest) (*ReportCreateDiscountedResponse, error) {
	var resp ReportCreateDiscountedResponse
	err := s.Client.Post(ctx, "/v1/report/discounted/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 发货报告
func (s *Service) CreateCompanyPostingsReport(ctx context.Context, req *ReportCreateCompanyPostingsReportRequest) (*ReportCreateReportResponse, error) {
	var resp ReportCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/postings/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 报告清单
func (s *Service) ReportList(ctx context.Context, req *ReportReportListRequest) (*ReportReportListResponse, error) {
	var resp ReportReportListResponse
	err := s.Client.Post(ctx, "/v1/report/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

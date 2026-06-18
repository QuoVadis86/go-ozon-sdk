package report

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 报告清单
func (s *Service) ReportList(ctx context.Context, req *ReportListRequest) (*ReportListResponse, error) {
	var resp ReportListResponse
	err := s.Client.Post(ctx, "/v1/report/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 生成带有标记商品的销售报告
// Note: 每个报告最多可包含 50,000 个商品标签代码
func (s *Service) CreateCompanyMarkedProductsSalesReport(ctx context.Context, req *V1ReportMarkedProductsSalesCreateRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/marked-products-sales/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 发货报告
func (s *Service) CreateCompanyPostingsReport(ctx context.Context, req *CreateCompanyPostingsReportRequest) (*CreateReportResponse, error) {
	var resp CreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/postings/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 关于FBS仓库库存报告
func (s *Service) CreateStockByWarehouseReport(ctx context.Context, req *V1CreateStockByWarehouseReportRequest) (*CommonCreateReportResponse, error) {
	var resp CommonCreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/warehouse/stock", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 减价商品报告
// Note: 从一个卖家账号每分钟可以发送1次请求
func (s *Service) CreateDiscountedReport(ctx context.Context, req *CreateDiscountedRequest) (*CreateDiscountedResponse, error) {
	var resp CreateDiscountedResponse
	err := s.Client.Post(ctx, "/v1/report/discounted/create", req, &resp)
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

// 报告信息
func (s *Service) ReportInfo(ctx context.Context, req *ReportInfoRequest) (*ReportInfoResponse, error) {
	var resp ReportInfoResponse
	err := s.Client.Post(ctx, "/v1/report/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 商品报告
func (s *Service) CreateCompanyProductsReport(ctx context.Context, req *CreateCompanyProductsReportRequest) (*CreateReportResponse, error) {
	var resp CreateReportResponse
	err := s.Client.Post(ctx, "/v1/report/products/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package premium

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) AnalyticsGetData(ctx context.Context, req *types.AnalyticsAnalyticsGetDataRequest) (*types.AnalyticsAnalyticsGetDataResponse, error) {
	var resp types.AnalyticsAnalyticsGetDataResponse
	err := s.Client.Post(ctx, "/v1/analytics/data", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SearchQueriesText(ctx context.Context, req *types.V1SearchQueriesTextRequest) (*types.V1SearchQueriesTextResponse, error) {
	var resp types.V1SearchQueriesTextResponse
	err := s.Client.Post(ctx, "/v1/search-queries/text", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AnalyticsProductQueriesDetails(ctx context.Context, req *types.V1AnalyticsProductQueriesDetailsRequest) (*types.V1AnalyticsProductQueriesDetailsResponse, error) {
	var resp types.V1AnalyticsProductQueriesDetailsResponse
	err := s.Client.Post(ctx, "/v1/analytics/product-queries/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AnalyticsProductQueries(ctx context.Context, req *types.V1AnalyticsProductQueriesRequest) (*types.V1AnalyticsProductQueriesResponse, error) {
	var resp types.V1AnalyticsProductQueriesResponse
	err := s.Client.Post(ctx, "/v1/analytics/product-queries", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetRealizationByDayReportV1(ctx context.Context, req *types.V1GetRealizationReportByDayRequest) (*types.GetRealizationReportByDayResponse, error) {
	var resp types.GetRealizationReportByDayResponse
	err := s.Client.Post(ctx, "/v1/finance/realization/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductPricesDetails(ctx context.Context, req *types.V1ProductPricesDetailsRequest) (*types.V1ProductPricesDetailsResponse, error) {
	var resp types.V1ProductPricesDetailsResponse
	err := s.Client.Post(ctx, "/v1/product/prices/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SearchQueriesTop(ctx context.Context, req *types.V1SearchQueriesTopRequest) (*types.V1SearchQueriesTopResponse, error) {
	var resp types.V1SearchQueriesTopResponse
	err := s.Client.Post(ctx, "/v1/search-queries/top", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

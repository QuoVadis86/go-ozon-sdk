package premium

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

func (s *Service) SearchQueriesTop(ctx context.Context, req *V1SearchQueriesTopRequest) (*V1SearchQueriesTopResponse, error) {
	var resp V1SearchQueriesTopResponse
	err := s.Client.Post(ctx, "/v1/search-queries/top", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductPricesDetails(ctx context.Context, req *V1ProductPricesDetailsRequest) (*V1ProductPricesDetailsResponse, error) {
	var resp V1ProductPricesDetailsResponse
	err := s.Client.Post(ctx, "/v1/product/prices/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SearchQueriesText(ctx context.Context, req *V1SearchQueriesTextRequest) (*V1SearchQueriesTextResponse, error) {
	var resp V1SearchQueriesTextResponse
	err := s.Client.Post(ctx, "/v1/search-queries/text", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AnalyticsProductQueriesDetails(ctx context.Context, req *V1AnalyticsProductQueriesDetailsRequest) (*V1AnalyticsProductQueriesDetailsResponse, error) {
	var resp V1AnalyticsProductQueriesDetailsResponse
	err := s.Client.Post(ctx, "/v1/analytics/product-queries/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AnalyticsGetData(ctx context.Context, req *AnalyticsAnalyticsGetDataRequest) (*AnalyticsAnalyticsGetDataResponse, error) {
	var resp AnalyticsAnalyticsGetDataResponse
	err := s.Client.Post(ctx, "/v1/analytics/data", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AnalyticsProductQueries(ctx context.Context, req *V1AnalyticsProductQueriesRequest) (*V1AnalyticsProductQueriesResponse, error) {
	var resp V1AnalyticsProductQueriesResponse
	err := s.Client.Post(ctx, "/v1/analytics/product-queries", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetRealizationByDayReportV1(ctx context.Context, req *V1GetRealizationReportByDayRequest) (*GetRealizationReportByDayResponse, error) {
	var resp GetRealizationReportByDayResponse
	err := s.Client.Post(ctx, "/v1/finance/realization/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package premium

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) AnalyticsGetData(ctx context.Context, req *internal.AnalyticsAnalyticsGetDataRequest) (*internal.AnalyticsAnalyticsGetDataResponse, error) {
	var resp internal.AnalyticsAnalyticsGetDataResponse
	err := s.Client.Post(ctx, "/v1/analytics/data", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SearchQueriesTop(ctx context.Context, req *internal.V1SearchQueriesTopRequest) (*internal.V1SearchQueriesTopResponse, error) {
	var resp internal.V1SearchQueriesTopResponse
	err := s.Client.Post(ctx, "/v1/search-queries/top", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetRealizationByDayReportV1(ctx context.Context, req *internal.V1GetRealizationReportByDayRequest) (*internal.GetRealizationReportByDayResponse, error) {
	var resp internal.GetRealizationReportByDayResponse
	err := s.Client.Post(ctx, "/v1/finance/realization/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AnalyticsProductQueries(ctx context.Context, req *internal.V1AnalyticsProductQueriesRequest) (*internal.V1AnalyticsProductQueriesResponse, error) {
	var resp internal.V1AnalyticsProductQueriesResponse
	err := s.Client.Post(ctx, "/v1/analytics/product-queries", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AnalyticsProductQueriesDetails(ctx context.Context, req *internal.V1AnalyticsProductQueriesDetailsRequest) (*internal.V1AnalyticsProductQueriesDetailsResponse, error) {
	var resp internal.V1AnalyticsProductQueriesDetailsResponse
	err := s.Client.Post(ctx, "/v1/analytics/product-queries/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductPricesDetails(ctx context.Context, req *internal.V1ProductPricesDetailsRequest) (*internal.V1ProductPricesDetailsResponse, error) {
	var resp internal.V1ProductPricesDetailsResponse
	err := s.Client.Post(ctx, "/v1/product/prices/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SearchQueriesText(ctx context.Context, req *internal.V1SearchQueriesTextRequest) (*internal.V1SearchQueriesTextResponse, error) {
	var resp internal.V1SearchQueriesTextResponse
	err := s.Client.Post(ctx, "/v1/search-queries/text", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

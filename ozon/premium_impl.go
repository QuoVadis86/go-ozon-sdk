package ozon

import "context"

func (s *PremiumService) SearchQueriesText(ctx context.Context, req *V1SearchQueriesTextRequest) (*V1SearchQueriesTextResponse, error) {
	var resp V1SearchQueriesTextResponse
	_, err := s.t.Post(ctx, "/v1/search-queries/text", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PremiumService) SearchQueriesTop(ctx context.Context, req *V1SearchQueriesTopRequest) (*V1SearchQueriesTopResponse, error) {
	var resp V1SearchQueriesTopResponse
	_, err := s.t.Post(ctx, "/v1/search-queries/top", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PremiumService) AnalyticsProductQueriesDetails(ctx context.Context, req *V1AnalyticsProductQueriesDetailsRequest) (*V1AnalyticsProductQueriesDetailsResponse, error) {
	var resp V1AnalyticsProductQueriesDetailsResponse
	_, err := s.t.Post(ctx, "/v1/analytics/product-queries/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PremiumService) AnalyticsProductQueries(ctx context.Context, req *V1AnalyticsProductQueriesRequest) (*V1AnalyticsProductQueriesResponse, error) {
	var resp V1AnalyticsProductQueriesResponse
	_, err := s.t.Post(ctx, "/v1/analytics/product-queries", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PremiumService) ProductPricesDetails(ctx context.Context, req *V1ProductPricesDetailsRequest) (*V1ProductPricesDetailsResponse, error) {
	var resp V1ProductPricesDetailsResponse
	_, err := s.t.Post(ctx, "/v1/product/prices/details", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PremiumService) AnalyticsGetData(ctx context.Context, req *AnalyticsAnalyticsGetDataRequest) (*AnalyticsAnalyticsGetDataResponse, error) {
	var resp AnalyticsAnalyticsGetDataResponse
	_, err := s.t.Post(ctx, "/v1/analytics/data", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PremiumService) GetRealizationByDayReportV1(ctx context.Context, req *V1GetRealizationReportByDayRequest) (*GetRealizationReportByDayResponse, error) {
	var resp GetRealizationReportByDayResponse
	_, err := s.t.Post(ctx, "/v1/finance/realization/by-day", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

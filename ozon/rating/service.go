package rating

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) ListFBSRatingIndexPostingsV1(ctx context.Context, req *internal.V1ListFBSRatingIndexPostingsV1Request) (*internal.V1ListFBSRatingIndexPostingsV1Response, error) {
	var resp internal.V1ListFBSRatingIndexPostingsV1Response
	err := s.Client.Post(ctx, "/v1/rating/index/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFBSRatingIndexInfoV1(ctx context.Context) (*internal.V1GetFBSRatingIndexInfoV1Response, error) {
	var resp internal.V1GetFBSRatingIndexInfoV1Response
	err := s.Client.Post(ctx, "/v1/rating/index/fbs/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

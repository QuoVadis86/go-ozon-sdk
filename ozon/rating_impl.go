package ozon

import "context"

func (s *RatingService) ListFBSRatingIndexPostingsV1(ctx context.Context, req *V1ListFBSRatingIndexPostingsV1Request) (*V1ListFBSRatingIndexPostingsV1Response, error) {
	var resp V1ListFBSRatingIndexPostingsV1Response
	_, err := s.t.Post(ctx, "/v1/rating/index/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *RatingService) GetFBSRatingIndexInfoV1(ctx context.Context) (*V1GetFBSRatingIndexInfoV1Response, error) {
	var resp V1GetFBSRatingIndexInfoV1Response
	_, err := s.t.Post(ctx, "/v1/rating/index/fbs/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

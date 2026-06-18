package rating

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) GetFBSRatingIndexInfoV1(ctx context.Context) (*types.V1GetFBSRatingIndexInfoV1Response, error) {
	var resp types.V1GetFBSRatingIndexInfoV1Response
	err := s.Client.Post(ctx, "/v1/rating/index/fbs/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListFBSRatingIndexPostingsV1(ctx context.Context, req *types.V1ListFBSRatingIndexPostingsV1Request) (*types.V1ListFBSRatingIndexPostingsV1Response, error) {
	var resp types.V1ListFBSRatingIndexPostingsV1Response
	err := s.Client.Post(ctx, "/v1/rating/index/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

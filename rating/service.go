package rating

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 影响错误指数的货件列表：FBS 和 rFBS
func (s *Service) ListFBSRatingIndexPostingsV1(ctx context.Context, req *V1ListFBSRatingIndexPostingsV1Request) (*V1ListFBSRatingIndexPostingsV1Response, error) {
	var resp V1ListFBSRatingIndexPostingsV1Response
	err := s.Client.Post(ctx, "/v1/rating/index/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取错误指数：FBS 和 rFBS
func (s *Service) GetFBSRatingIndexInfoV1(ctx context.Context) (*V1GetFBSRatingIndexInfoV1Response, error) {
	var resp V1GetFBSRatingIndexInfoV1Response
	err := s.Client.Post(ctx, "/v1/rating/index/fbs/info", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

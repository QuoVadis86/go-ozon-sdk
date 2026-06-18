package review

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 按状态统计问题数量
func (s *Service) Count(ctx context.Context) (*V1QuestionCountResponse, error) {
	var resp V1QuestionCountResponse
	err := s.Client.Post(ctx, "/v1/question/count", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 评价的评论列表
func (s *Service) CommentList(ctx context.Context, req *V1CommentListRequest) (*V1CommentListResponse, error) {
	var resp V1CommentListResponse
	err := s.Client.Post(ctx, "/v1/review/comment/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 问题回答列表
func (s *Service) List(ctx context.Context, req *V1QuestionAnswerListRequest) (*V1QuestionAnswerListResponse, error) {
	var resp V1QuestionAnswerListResponse
	err := s.Client.Post(ctx, "/v1/question/answer/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 提问数量最多的商品
func (s *Service) TopSku(ctx context.Context, req *V1QuestionTopSkuRequest) (*V1QuestionTopSkuResponse, error) {
	var resp V1QuestionTopSkuResponse
	err := s.Client.Post(ctx, "/v1/question/top_sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更改问题状态
func (s *Service) ChangeStatus(ctx context.Context, req *V1QuestionChangeStatusRequest) error {
	err := s.Client.Post(ctx, "/v1/question/change_status", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 删除对评价的评论
func (s *Service) CommentDelete(ctx context.Context, req *V1CommentDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/review/comment/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 获取评价信息
func (s *Service) ReviewInfo(ctx context.Context, req *V1ReviewInfoRequest) (*V1ReviewInfoResponse, error) {
	var resp V1ReviewInfoResponse
	err := s.Client.Post(ctx, "/v1/review/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: use /v2/review/list instead
// 获取评价列表
func (s *Service) ReviewList(ctx context.Context, req *V1ReviewListRequest) (*V1ReviewListResponse, error) {
	var resp V1ReviewListResponse
	err := s.Client.Post(ctx, "/v1/review/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 问题列表
func (s *Service) ListV1(ctx context.Context, req *V1QuestionListRequest) (*V1QuestionListResponse, error) {
	var resp V1QuestionListResponse
	err := s.Client.Post(ctx, "/v1/question/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 问题详情
func (s *Service) Info(ctx context.Context, req *V1QuestionInfoRequest) (*V1QuestionInfoResponse, error) {
	var resp V1QuestionInfoResponse
	err := s.Client.Post(ctx, "/v1/question/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建对问题的回答
func (s *Service) Create(ctx context.Context, req *V1QuestionAnswerCreateRequest) (*V1QuestionAnswerCreateResponse, error) {
	var resp V1QuestionAnswerCreateResponse
	err := s.Client.Post(ctx, "/v1/question/answer/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更改评价状态
func (s *Service) ReviewChangeStatus(ctx context.Context, req *V1ReviewChangeStatusRequest) error {
	err := s.Client.Post(ctx, "/v1/review/change-status", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 获取评价列表
func (s *Service) ReviewListV2(ctx context.Context, req *V2ReviewListV2Request) (*V2ReviewListV2Response, error) {
	var resp V2ReviewListV2Response
	err := s.Client.Post(ctx, "/v2/review/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 对评价留下评论
func (s *Service) CommentCreate(ctx context.Context, req *V1CommentCreateRequest) (*V1CommentCreateResponse, error) {
	var resp V1CommentCreateResponse
	err := s.Client.Post(ctx, "/v1/review/comment/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 根据状态统计的评价数量
func (s *Service) ReviewCount(ctx context.Context) (*V1ReviewCountResponse, error) {
	var resp V1ReviewCountResponse
	err := s.Client.Post(ctx, "/v1/review/count", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 删除问题回答
func (s *Service) Delete(ctx context.Context, req *V1QuestionAnswerDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/question/answer/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

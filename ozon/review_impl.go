package ozon

import "context"

func (s *ReviewService) Create(ctx context.Context, req *V1QuestionAnswerCreateRequest) (*V1QuestionAnswerCreateResponse, error) {
	var resp V1QuestionAnswerCreateResponse
	_, err := s.t.Post(ctx, "/v1/question/answer/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) ReviewCount(ctx context.Context) (*V1ReviewCountResponse, error) {
	var resp V1ReviewCountResponse
	_, err := s.t.Post(ctx, "/v1/review/count", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) ReviewInfo(ctx context.Context, req *V1ReviewInfoRequest) (*V1ReviewInfoResponse, error) {
	var resp V1ReviewInfoResponse
	_, err := s.t.Post(ctx, "/v1/review/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) Delete(ctx context.Context, req *V1QuestionAnswerDeleteRequest) error {
	_, err := s.t.Post(ctx, "/v1/question/answer/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReviewService) CommentDelete(ctx context.Context, req *V1CommentDeleteRequest) error {
	_, err := s.t.Post(ctx, "/v1/review/comment/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReviewService) ReviewChangeStatus(ctx context.Context, req *V1ReviewChangeStatusRequest) error {
	_, err := s.t.Post(ctx, "/v1/review/change-status", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReviewService) List(ctx context.Context, req *V1QuestionListRequest) (*V1QuestionListResponse, error) {
	var resp V1QuestionListResponse
	_, err := s.t.Post(ctx, "/v1/question/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) TopSku(ctx context.Context, req *V1QuestionTopSkuRequest) (*V1QuestionTopSkuResponse, error) {
	var resp V1QuestionTopSkuResponse
	_, err := s.t.Post(ctx, "/v1/question/top_sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) ListList(ctx context.Context, req *V1QuestionAnswerListRequest) (*V1QuestionAnswerListResponse, error) {
	var resp V1QuestionAnswerListResponse
	_, err := s.t.Post(ctx, "/v1/question/answer/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) CommentCreate(ctx context.Context, req *V1CommentCreateRequest) (*V1CommentCreateResponse, error) {
	var resp V1CommentCreateResponse
	_, err := s.t.Post(ctx, "/v1/review/comment/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) Count(ctx context.Context) (*V1QuestionCountResponse, error) {
	var resp V1QuestionCountResponse
	_, err := s.t.Post(ctx, "/v1/question/count", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) ReviewList(ctx context.Context, req *V1ReviewListRequest) (*V1ReviewListResponse, error) {
	var resp V1ReviewListResponse
	_, err := s.t.Post(ctx, "/v1/review/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) ReviewListV2(ctx context.Context, req *ReviewV2ReviewListV2Request) (*ReviewV2ReviewListV2Response, error) {
	var resp ReviewV2ReviewListV2Response
	_, err := s.t.Post(ctx, "/v2/review/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) ChangeStatus(ctx context.Context, req *V1QuestionChangeStatusRequest) error {
	_, err := s.t.Post(ctx, "/v1/question/change_status", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReviewService) CommentList(ctx context.Context, req *V1CommentListRequest) (*V1CommentListResponse, error) {
	var resp V1CommentListResponse
	_, err := s.t.Post(ctx, "/v1/review/comment/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReviewService) Info(ctx context.Context, req *V1QuestionInfoRequest) (*V1QuestionInfoResponse, error) {
	var resp V1QuestionInfoResponse
	_, err := s.t.Post(ctx, "/v1/question/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

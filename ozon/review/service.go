package review

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) List(ctx context.Context, req *internal.V1QuestionListRequest) (*internal.V1QuestionListResponse, error) {
	var resp internal.V1QuestionListResponse
	err := s.Client.Post(ctx, "/v1/question/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReviewCount(ctx context.Context) (*internal.V1ReviewCountResponse, error) {
	var resp internal.V1ReviewCountResponse
	err := s.Client.Post(ctx, "/v1/review/count", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReviewInfo(ctx context.Context, req *internal.V1ReviewInfoRequest) (*internal.V1ReviewInfoResponse, error) {
	var resp internal.V1ReviewInfoResponse
	err := s.Client.Post(ctx, "/v1/review/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Create(ctx context.Context, req *internal.V1QuestionAnswerCreateRequest) (*internal.V1QuestionAnswerCreateResponse, error) {
	var resp internal.V1QuestionAnswerCreateResponse
	err := s.Client.Post(ctx, "/v1/question/answer/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Info(ctx context.Context, req *internal.V1QuestionInfoRequest) (*internal.V1QuestionInfoResponse, error) {
	var resp internal.V1QuestionInfoResponse
	err := s.Client.Post(ctx, "/v1/question/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CommentDelete(ctx context.Context, req *internal.V1CommentDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/review/comment/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReviewList(ctx context.Context, req *internal.V1ReviewListRequest) (*internal.V1ReviewListResponse, error) {
	var resp internal.V1ReviewListResponse
	err := s.Client.Post(ctx, "/v1/review/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChangeStatus(ctx context.Context, req *internal.V1QuestionChangeStatusRequest) error {
	err := s.Client.Post(ctx, "/v1/question/change_status", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) TopSku(ctx context.Context, req *internal.V1QuestionTopSkuRequest) (*internal.V1QuestionTopSkuResponse, error) {
	var resp internal.V1QuestionTopSkuResponse
	err := s.Client.Post(ctx, "/v1/question/top_sku", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Delete(ctx context.Context, req *internal.V1QuestionAnswerDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/question/answer/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CommentCreate(ctx context.Context, req *internal.V1CommentCreateRequest) (*internal.V1CommentCreateResponse, error) {
	var resp internal.V1CommentCreateResponse
	err := s.Client.Post(ctx, "/v1/review/comment/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListV1(ctx context.Context, req *internal.V1QuestionAnswerListRequest) (*internal.V1QuestionAnswerListResponse, error) {
	var resp internal.V1QuestionAnswerListResponse
	err := s.Client.Post(ctx, "/v1/question/answer/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReviewChangeStatus(ctx context.Context, req *internal.V1ReviewChangeStatusRequest) error {
	err := s.Client.Post(ctx, "/v1/review/change-status", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CommentList(ctx context.Context, req *internal.V1CommentListRequest) (*internal.V1CommentListResponse, error) {
	var resp internal.V1CommentListResponse
	err := s.Client.Post(ctx, "/v1/review/comment/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReviewListV2(ctx context.Context, req *internal.ReviewV2ReviewListV2Request) (*internal.ReviewV2ReviewListV2Response, error) {
	var resp internal.ReviewV2ReviewListV2Response
	err := s.Client.Post(ctx, "/v2/review/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Count(ctx context.Context) (*internal.V1QuestionCountResponse, error) {
	var resp internal.V1QuestionCountResponse
	err := s.Client.Post(ctx, "/v1/question/count", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

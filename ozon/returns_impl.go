package ozon

import "context"

func (s *ReturnsService) ReturnsRfbsListV2(ctx context.Context, req *V2ReturnsRfbsListRequest) (*V2ReturnsRfbsListResponse, error) {
	var resp V2ReturnsRfbsListResponse
	_, err := s.t.Post(ctx, "/v2/returns/rfbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) ReturnsRfbsCompensateV2(ctx context.Context, req *V2ReturnsRfbsCompensateRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v2/returns/rfbs/compensate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) ReturnsRfbsVerifyV2(ctx context.Context, req *V2ReturnsRfbsVerifyRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v2/returns/rfbs/verify", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) ReturnsRfbsActionSet(ctx context.Context, req *V1ReturnsRfbsActionSetRequest) error {
	_, err := s.t.Post(ctx, "/v1/returns/rfbs/action/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReturnsService) ReturnsRfbsGetV2(ctx context.Context, req *V2ReturnsRfbsGetRequest) (*V2ReturnsRfbsGetResponse, error) {
	var resp V2ReturnsRfbsGetResponse
	_, err := s.t.Post(ctx, "/v2/returns/rfbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) ConditionalCancellationApproveV2(ctx context.Context) error {
	_, err := s.t.Post(ctx, "/v2/conditional-cancellation/approve", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReturnsService) ReturnsRfbsRejectV2(ctx context.Context, req *V2ReturnsRfbsRejectRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v2/returns/rfbs/reject", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) ReturnsList(ctx context.Context, req *V1GetReturnsListRequest) (*V1GetReturnsListResponse, error) {
	var resp V1GetReturnsListResponse
	_, err := s.t.Post(ctx, "/v1/returns/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) GetConditionalCancellationListV2(ctx context.Context, req *V2GetConditionalCancellationListV2Request) (*V2GetConditionalCancellationListV2Response, error) {
	var resp V2GetConditionalCancellationListV2Response
	_, err := s.t.Post(ctx, "/v2/conditional-cancellation/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) ReturnsRfbsReceiveReturnV2(ctx context.Context, req *V2ReturnsRfbsReceiveReturnRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v2/returns/rfbs/receive-return", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ReturnsService) ConditionalCancellationRejectV2(ctx context.Context) error {
	_, err := s.t.Post(ctx, "/v2/conditional-cancellation/reject", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *ReturnsService) ReturnsRfbsReturnMoneyV2(ctx context.Context, req *V2ReturnsRfbsReturnMoneyRequest) (*V1Empty, error) {
	var resp V1Empty
	_, err := s.t.Post(ctx, "/v2/returns/rfbs/return-money", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

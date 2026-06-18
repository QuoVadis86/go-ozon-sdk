package returns

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) ConditionalCancellationApproveV2(ctx context.Context) error {
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/approve", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetConditionalCancellationListV2(ctx context.Context, req *internal.V2GetConditionalCancellationListV2Request) (*internal.V2GetConditionalCancellationListV2Response, error) {
	var resp internal.V2GetConditionalCancellationListV2Response
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsActionSet(ctx context.Context, req *internal.V1ReturnsRfbsActionSetRequest) error {
	err := s.Client.Post(ctx, "/v1/returns/rfbs/action/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnsRfbsGetV2(ctx context.Context, req *internal.V2ReturnsRfbsGetRequest) (*internal.V2ReturnsRfbsGetResponse, error) {
	var resp internal.V2ReturnsRfbsGetResponse
	err := s.Client.Post(ctx, "/v2/returns/rfbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsList(ctx context.Context, req *internal.V1GetReturnsListRequest) (*internal.V1GetReturnsListResponse, error) {
	var resp internal.V1GetReturnsListResponse
	err := s.Client.Post(ctx, "/v1/returns/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsReceiveReturnV2(ctx context.Context, req *internal.V2ReturnsRfbsReceiveReturnRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/receive-return", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsVerifyV2(ctx context.Context, req *internal.V2ReturnsRfbsVerifyRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/verify", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsRejectV2(ctx context.Context, req *internal.V2ReturnsRfbsRejectRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/reject", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsListV2(ctx context.Context, req *internal.V2ReturnsRfbsListRequest) (*internal.V2ReturnsRfbsListResponse, error) {
	var resp internal.V2ReturnsRfbsListResponse
	err := s.Client.Post(ctx, "/v2/returns/rfbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsCompensateV2(ctx context.Context, req *internal.V2ReturnsRfbsCompensateRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/compensate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsReturnMoneyV2(ctx context.Context, req *internal.V2ReturnsRfbsReturnMoneyRequest) (*internal.V1Empty, error) {
	var resp internal.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/return-money", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ConditionalCancellationRejectV2(ctx context.Context) error {
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/reject", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

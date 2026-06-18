package returns

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) ConditionalCancellationApproveV2(ctx context.Context) error {
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/approve", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnsRfbsActionSet(ctx context.Context, req *V1ReturnsRfbsActionSetRequest) error {
	err := s.Client.Post(ctx, "/v1/returns/rfbs/action/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnsRfbsReturnMoneyV2(ctx context.Context, req *V2ReturnsRfbsReturnMoneyRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/return-money", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsListV2(ctx context.Context, req *V2ReturnsRfbsListRequest) (*V2ReturnsRfbsListResponse, error) {
	var resp V2ReturnsRfbsListResponse
	err := s.Client.Post(ctx, "/v2/returns/rfbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsVerifyV2(ctx context.Context, req *V2ReturnsRfbsVerifyRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/verify", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetConditionalCancellationListV2(ctx context.Context, req *V2GetConditionalCancellationListV2Request) (*V2GetConditionalCancellationListV2Response, error) {
	var resp V2GetConditionalCancellationListV2Response
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsRejectV2(ctx context.Context, req *V2ReturnsRfbsRejectRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/reject", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsList(ctx context.Context, req *V1GetReturnsListRequest) (*V1GetReturnsListResponse, error) {
	var resp V1GetReturnsListResponse
	err := s.Client.Post(ctx, "/v1/returns/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsReceiveReturnV2(ctx context.Context, req *V2ReturnsRfbsReceiveReturnRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/receive-return", req, &resp)
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

func (s *Service) ReturnsRfbsCompensateV2(ctx context.Context, req *V2ReturnsRfbsCompensateRequest) (*types.V1Empty, error) {
	var resp types.V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/compensate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsRfbsGetV2(ctx context.Context, req *V2ReturnsRfbsGetRequest) (*V2ReturnsRfbsGetResponse, error) {
	var resp V2ReturnsRfbsGetResponse
	err := s.Client.Post(ctx, "/v2/returns/rfbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

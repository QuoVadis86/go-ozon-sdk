package pass

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) CarriagePassDelete(ctx context.Context, req *types.SellerSellerAPIArrivalPassDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnPassCreate(ctx context.Context, req *types.ArrivalpassArrivalPassCreateRequest) (*types.ArrivalpassArrivalPassCreateResponse, error) {
	var resp types.ArrivalpassArrivalPassCreateResponse
	err := s.Client.Post(ctx, "/v1/return/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriagePassCreate(ctx context.Context, req *types.SellerSellerAPIArrivalPassCreateRequest) (*types.SellerSellerAPIArrivalPassCreateResponse, error) {
	var resp types.SellerSellerAPIArrivalPassCreateResponse
	err := s.Client.Post(ctx, "/v1/carriage/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriagePassUpdate(ctx context.Context, req *types.SellerSellerAPIArrivalPassUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnPassDelete(ctx context.Context, req *types.ArrivalpassArrivalPassDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/return/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnPassUpdate(ctx context.Context, req *types.ArrivalpassArrivalPassUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/return/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnsCompanyFBSInfo(ctx context.Context, req *types.V1ReturnsCompanyFbsInfoRequest) (*types.V1ReturnsCompanyFbsInfoResponse, error) {
	var resp types.V1ReturnsCompanyFbsInfoResponse
	err := s.Client.Post(ctx, "/v1/returns/company/fbs/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PassList(ctx context.Context, req *types.ArrivalpassArrivalPassListRequest) (*types.ArrivalpassArrivalPassListResponse, error) {
	var resp types.ArrivalpassArrivalPassListResponse
	err := s.Client.Post(ctx, "/v1/pass/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

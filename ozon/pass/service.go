package pass

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) ReturnPassDelete(ctx context.Context, req *internal.ArrivalpassArrivalPassDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/return/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnPassUpdate(ctx context.Context, req *internal.ArrivalpassArrivalPassUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/return/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CarriagePassDelete(ctx context.Context, req *internal.SellerSellerAPIArrivalPassDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CarriagePassCreate(ctx context.Context, req *internal.SellerSellerAPIArrivalPassCreateRequest) (*internal.SellerSellerAPIArrivalPassCreateResponse, error) {
	var resp internal.SellerSellerAPIArrivalPassCreateResponse
	err := s.Client.Post(ctx, "/v1/carriage/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PassList(ctx context.Context, req *internal.ArrivalpassArrivalPassListRequest) (*internal.ArrivalpassArrivalPassListResponse, error) {
	var resp internal.ArrivalpassArrivalPassListResponse
	err := s.Client.Post(ctx, "/v1/pass/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnPassCreate(ctx context.Context, req *internal.ArrivalpassArrivalPassCreateRequest) (*internal.ArrivalpassArrivalPassCreateResponse, error) {
	var resp internal.ArrivalpassArrivalPassCreateResponse
	err := s.Client.Post(ctx, "/v1/return/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnsCompanyFBSInfo(ctx context.Context, req *internal.V1ReturnsCompanyFbsInfoRequest) (*internal.V1ReturnsCompanyFbsInfoResponse, error) {
	var resp internal.V1ReturnsCompanyFbsInfoResponse
	err := s.Client.Post(ctx, "/v1/returns/company/fbs/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriagePassUpdate(ctx context.Context, req *internal.SellerSellerAPIArrivalPassUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

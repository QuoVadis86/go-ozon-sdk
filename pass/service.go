package pass

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

func (s *Service) PassList(ctx context.Context, req *ArrivalpassArrivalPassListRequest) (*ArrivalpassArrivalPassListResponse, error) {
	var resp ArrivalpassArrivalPassListResponse
	err := s.Client.Post(ctx, "/v1/pass/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriagePassUpdate(ctx context.Context, req *SellerSellerAPIArrivalPassUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnPassUpdate(ctx context.Context, req *ArrivalpassArrivalPassUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/return/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnsCompanyFBSInfo(ctx context.Context, req *V1ReturnsCompanyFbsInfoRequest) (*V1ReturnsCompanyFbsInfoResponse, error) {
	var resp V1ReturnsCompanyFbsInfoResponse
	err := s.Client.Post(ctx, "/v1/returns/company/fbs/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriagePassCreate(ctx context.Context, req *SellerSellerAPIArrivalPassCreateRequest) (*SellerSellerAPIArrivalPassCreateResponse, error) {
	var resp SellerSellerAPIArrivalPassCreateResponse
	err := s.Client.Post(ctx, "/v1/carriage/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriagePassDelete(ctx context.Context, req *SellerSellerAPIArrivalPassDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ReturnPassCreate(ctx context.Context, req *ArrivalpassArrivalPassCreateRequest) (*ArrivalpassArrivalPassCreateResponse, error) {
	var resp ArrivalpassArrivalPassCreateResponse
	err := s.Client.Post(ctx, "/v1/return/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ReturnPassDelete(ctx context.Context, req *ArrivalpassArrivalPassDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/return/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

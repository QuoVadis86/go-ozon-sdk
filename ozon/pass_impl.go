package ozon

import "context"

func (s *PassService) ReturnPassCreate(ctx context.Context, req *ArrivalpassArrivalPassCreateRequest) (*ArrivalpassArrivalPassCreateResponse, error) {
	var resp ArrivalpassArrivalPassCreateResponse
	_, err := s.t.Post(ctx, "/v1/return/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PassService) PassList(ctx context.Context, req *ArrivalpassArrivalPassListRequest) (*ArrivalpassArrivalPassListResponse, error) {
	var resp ArrivalpassArrivalPassListResponse
	_, err := s.t.Post(ctx, "/v1/pass/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PassService) CarriagePassCreate(ctx context.Context, req *SellerSellerAPIArrivalPassCreateRequest) (*SellerSellerAPIArrivalPassCreateResponse, error) {
	var resp SellerSellerAPIArrivalPassCreateResponse
	_, err := s.t.Post(ctx, "/v1/carriage/pass/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PassService) CarriagePassUpdate(ctx context.Context, req *SellerSellerAPIArrivalPassUpdateRequest) error {
	_, err := s.t.Post(ctx, "/v1/carriage/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PassService) ReturnPassUpdate(ctx context.Context, req *ArrivalpassArrivalPassUpdateRequest) error {
	_, err := s.t.Post(ctx, "/v1/return/pass/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PassService) ReturnsCompanyFBSInfo(ctx context.Context, req *V1ReturnsCompanyFbsInfoRequest) (*V1ReturnsCompanyFbsInfoResponse, error) {
	var resp V1ReturnsCompanyFbsInfoResponse
	_, err := s.t.Post(ctx, "/v1/returns/company/fbs/info", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PassService) CarriagePassDelete(ctx context.Context, req *SellerSellerAPIArrivalPassDeleteRequest) error {
	_, err := s.t.Post(ctx, "/v1/carriage/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PassService) ReturnPassDelete(ctx context.Context, req *ArrivalpassArrivalPassDeleteRequest) error {
	_, err := s.t.Post(ctx, "/v1/return/pass/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

package fbo

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) FbpEditTimeslot(ctx context.Context, req *types.V1FbpEditTimeslotRequest) (*types.V1FbpEditTimeslotResponse, error) {
	var resp types.V1FbpEditTimeslotResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/timeslot/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectSellerDlvEdit(ctx context.Context, req *types.V1FbpDraftDirectSellerDlvEditRequest) (*types.V1FbpDraftDirectSellerDlvEditResponse, error) {
	var resp types.V1FbpDraftDirectSellerDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/seller-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderDirectCancel(ctx context.Context, req *types.V1FbpOrderDirectCancelRequest) (*types.V1FbpOrderDirectCancelResponse, error) {
	var resp types.V1FbpOrderDirectCancelResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpWarehouseList(ctx context.Context) (*types.V1FbpWarehouseListResponse, error) {
	var resp types.V1FbpWarehouseListResponse
	err := s.Client.Post(ctx, "/v1/fbp/warehouse/list", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderDropOffCancel(ctx context.Context, req *types.V1FbpOrderDropOffCancelRequest) (*types.V1FbpOrderDropOffCancelResponse, error) {
	var resp types.V1FbpOrderDropOffCancelResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/drop-off/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpGetLabel(ctx context.Context, req *types.V1FbpGetLabelRequest) (*types.V1FbpGetLabelResponse, error) {
	var resp types.V1FbpGetLabelResponse
	err := s.Client.Post(ctx, "/v1/fbp/label/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectSellerDlvCreate(ctx context.Context, req *types.V1FbpDraftDirectSellerDlvCreateRequest) (*types.V1FbpDraftDirectSellerDlvCreateResponse, error) {
	var resp types.V1FbpDraftDirectSellerDlvCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/seller-dlv/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpCreateAct(ctx context.Context, req *types.V1FbpCreateActRequest) (*types.V1FbpCreateActResponse, error) {
	var resp types.V1FbpCreateActResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-from/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectDelete(ctx context.Context, req *types.V1FbpDraftDirectDeleteRequest) (*types.V1FbpDraftDirectDeleteResponse, error) {
	var resp types.V1FbpDraftDirectDeleteResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftPickupCreate(ctx context.Context, req *types.V1FbpDraftPickupCreateRequest) (*types.V1FbpDraftPickupCreateResponse, error) {
	var resp types.V1FbpDraftPickupCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderGet(ctx context.Context, req *types.V1FbpOrderGetRequest) (*types.V1FbpOrderGetResponse, error) {
	var resp types.V1FbpOrderGetResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpCreateLabel(ctx context.Context, req *types.V1FbpCreateLabelRequest) (*types.V1FbpCreateLabelResponse, error) {
	var resp types.V1FbpCreateLabelResponse
	err := s.Client.Post(ctx, "/v1/fbp/label/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectCreate(ctx context.Context, req *types.V1FbpDraftDirectCreateRequest) (*types.V1FbpDraftDirectCreateResponse, error) {
	var resp types.V1FbpDraftDirectCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectTplDlvEdit(ctx context.Context, req *types.V1FbpDraftDirectTplDlvEditRequest) (*types.V1FbpDraftDirectTplDlvEditResponse, error) {
	var resp types.V1FbpDraftDirectTplDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/tpl-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFbpList(ctx context.Context, req *types.PostingV1PostingFbpListRequest) (*types.PostingV1PostingFbpListResponse, error) {
	var resp types.PostingV1PostingFbpListResponse
	err := s.Client.Post(ctx, "/v1/posting/fbp/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectGetTimeslot(ctx context.Context, req *types.V1FbpDraftDirectGetTimeslotRequest) (*types.V1FbpDraftDirectGetTimeslotResponse, error) {
	var resp types.V1FbpDraftDirectGetTimeslotResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/timeslot/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderPickUpCancel(ctx context.Context, req *types.V1FbpOrderPickUpCancelRequest) (*types.V1FbpOrderPickUpCancelResponse, error) {
	var resp types.V1FbpOrderPickUpCancelResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/pick-up/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffDelete(ctx context.Context, req *types.V1FbpDraftDropOffDeleteRequest) (*types.V1FbpDraftDropOffDeleteResponse, error) {
	var resp types.V1FbpDraftDropOffDeleteResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffProductValidate(ctx context.Context, req *types.V1FbpDraftDropOffProductValidateRequest) (*types.V1FbpDraftDropOffProductValidateResponse, error) {
	var resp types.V1FbpDraftDropOffProductValidateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftPickUpDelete(ctx context.Context, req *types.V1FbpDraftPickUpDeleteRequest) (*types.V1FbpDraftPickUpDeleteResponse, error) {
	var resp types.V1FbpDraftPickUpDeleteResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectProductValidate(ctx context.Context, req *types.V1FbpDraftDirectProductValidateRequest) (*types.V1FbpDraftDirectProductValidateResponse, error) {
	var resp types.V1FbpDraftDirectProductValidateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpCheckConsignmentNoteState(ctx context.Context, req *types.V1FbpCheckConsignmentNoteStateRequest) (*types.V1FbpCheckConsignmentNoteStateResponse, error) {
	var resp types.V1FbpCheckConsignmentNoteStateResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-to/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectTplDlvCreate(ctx context.Context, req *types.V1FbpDraftDirectTplDlvCreateRequest) (*types.V1FbpDraftDirectTplDlvCreateResponse, error) {
	var resp types.V1FbpDraftDirectTplDlvCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/tpl-dlv/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffRegistrate(ctx context.Context, req *types.V1FbpDraftDropOffRegistrateRequest) (*types.V1FbpDraftDropOffRegistrateResponse, error) {
	var resp types.V1FbpDraftDropOffRegistrateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpCreateConsignmentNote(ctx context.Context, req *types.V1FbpCreateConsignmentNoteRequest) (*types.V1FbpCreateConsignmentNoteResponse, error) {
	var resp types.V1FbpCreateConsignmentNoteResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-to/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftPickupDlvEdit(ctx context.Context, req *types.V1FbpDraftPickupDlvEditRequest) (*types.V1FbpDraftPickupDlvEditResponse, error) {
	var resp types.V1FbpDraftPickupDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffPointTimetable(ctx context.Context, req *types.V1FbpDraftDropOffPointTimetableRequest) (*types.V1FbpDraftDropOffPointTimetableResponse, error) {
	var resp types.V1FbpDraftDropOffPointTimetableResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/point/timetable", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectTimeslotEdit(ctx context.Context, req *types.V1FbpDraftDirectTimeslotEditRequest) (*types.V1FbpDraftDirectTimeslotEditResponse, error) {
	var resp types.V1FbpDraftDirectTimeslotEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/timeslot/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffProvinceList(ctx context.Context, req *types.V1FbpDraftDropOffProvinceListRequest) (*types.V1FbpDraftDropOffProvinceListResponse, error) {
	var resp types.V1FbpDraftDropOffProvinceListResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/province/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SupplyOrderBundle(ctx context.Context, req *types.V1GetSupplyOrderBundleRequest) (*types.V1GetSupplyOrderBundleResponse, error) {
	var resp types.V1GetSupplyOrderBundleResponse
	err := s.Client.Post(ctx, "/v1/supply-order/bundle", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderDropOffDlvEdit(ctx context.Context, req *types.V1FbpOrderDropOffDlvEditRequest) (*types.V1FbpOrderDropOffDlvEditResponse, error) {
	var resp types.V1FbpOrderDropOffDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/drop-off/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpAvailableTimeslotList(ctx context.Context, req *types.V1FbpAvailableTimeslotListRequest) (*types.V1FbpAvailableTimeslotListResponse, error) {
	var resp types.V1FbpAvailableTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffPointList(ctx context.Context, req *types.V1FbpDraftDropOffPointListRequest) (*types.V1FbpDraftDropOffPointListResponse, error) {
	var resp types.V1FbpDraftDropOffPointListResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/point/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffCreate(ctx context.Context, req *types.V1FbpDraftDropOffCreateRequest) (*types.V1FbpDraftDropOffCreateResponse, error) {
	var resp types.V1FbpDraftDropOffCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDirectRegistrate(ctx context.Context, req *types.V1FbpDraftDirectRegistrateRequest) (*types.V1FbpDraftDirectRegistrateResponse, error) {
	var resp types.V1FbpDraftDirectRegistrateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftDropOffDlvEdit(ctx context.Context, req *types.V1FbpDraftDropOffDlvEditRequest) (*types.V1FbpDraftDropOffDlvEditResponse, error) {
	var resp types.V1FbpDraftDropOffDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpArchiveList(ctx context.Context, req *types.V1FbpArchiveListRequest) (*types.V1FbpArchiveListResponse, error) {
	var resp types.V1FbpArchiveListResponse
	err := s.Client.Post(ctx, "/v1/fbp/archive/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftPickUpRegistrate(ctx context.Context, req *types.V1FbpDraftPickUpRegistrateRequest) (*types.V1FbpDraftPickUpRegistrateResponse, error) {
	var resp types.V1FbpDraftPickUpRegistrateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderList(ctx context.Context, req *types.V1FbpOrderListRequest) (*types.V1FbpOrderListResponse, error) {
	var resp types.V1FbpOrderListResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftPickUpProductValidate(ctx context.Context, req *types.V1FbpDraftPickUpProductValidateRequest) (*types.V1FbpDraftPickUpProductValidateResponse, error) {
	var resp types.V1FbpDraftPickUpProductValidateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftList(ctx context.Context, req *types.V1FbpDraftListRequest) (*types.V1FbpDraftListResponse, error) {
	var resp types.V1FbpDraftListResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderDropOffTimetable(ctx context.Context, req *types.V1FbpOrderDropOffTimetableRequest) (*types.V1FbpOrderDropOffTimetableResponse, error) {
	var resp types.V1FbpOrderDropOffTimetableResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/drop-off/timetable", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpCheckActState(ctx context.Context, req *types.V1FbpCheckActStateRequest) (*types.V1FbpCheckActStateResponse, error) {
	var resp types.V1FbpCheckActStateResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-from/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpArchiveGet(ctx context.Context, req *types.V1FbpArchiveGetRequest) (*types.V1FbpArchiveGetResponse, error) {
	var resp types.V1FbpArchiveGetResponse
	err := s.Client.Post(ctx, "/v1/fbp/archive/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderPickUpDlvEdit(ctx context.Context, req *types.V1FbpOrderPickUpDlvEditRequest) (*types.V1FbpOrderPickUpDlvEditResponse, error) {
	var resp types.V1FbpOrderPickUpDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/pick-up/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpOrderDirectSellerDlvEdit(ctx context.Context, req *types.V1FbpOrderDirectSellerDlvEditRequest) (*types.V1FbpOrderDirectSellerDlvEditResponse, error) {
	var resp types.V1FbpOrderDirectSellerDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/seller-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbpDraftGet(ctx context.Context, req *types.V1FbpDraftGetRequest) (*types.V1FbpDraftGetResponse, error) {
	var resp types.V1FbpDraftGetResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

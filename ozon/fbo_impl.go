package ozon

import "context"

func (s *FBOService) FbpDraftGet(ctx context.Context, req *V1FbpDraftGetRequest) (*V1FbpDraftGetResponse, error) {
	var resp V1FbpDraftGetResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffProvinceList(ctx context.Context, req *V1FbpDraftDropOffProvinceListRequest) (*V1FbpDraftDropOffProvinceListResponse, error) {
	var resp V1FbpDraftDropOffProvinceListResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/province/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffProductValidate(ctx context.Context, req *V1FbpDraftDropOffProductValidateRequest) (*V1FbpDraftDropOffProductValidateResponse, error) {
	var resp V1FbpDraftDropOffProductValidateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectRegistrate(ctx context.Context, req *V1FbpDraftDirectRegistrateRequest) (*V1FbpDraftDirectRegistrateResponse, error) {
	var resp V1FbpDraftDirectRegistrateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftPickUpDelete(ctx context.Context, req *V1FbpDraftPickUpDeleteRequest) (*V1FbpDraftPickUpDeleteResponse, error) {
	var resp V1FbpDraftPickUpDeleteResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/pick-up/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectTplDlvEdit(ctx context.Context, req *V1FbpDraftDirectTplDlvEditRequest) (*V1FbpDraftDirectTplDlvEditResponse, error) {
	var resp V1FbpDraftDirectTplDlvEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/tpl-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftPickUpProductValidate(ctx context.Context, req *V1FbpDraftPickUpProductValidateRequest) (*V1FbpDraftPickUpProductValidateResponse, error) {
	var resp V1FbpDraftPickUpProductValidateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/pick-up/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffDlvEdit(ctx context.Context, req *V1FbpDraftDropOffDlvEditRequest) (*V1FbpDraftDropOffDlvEditResponse, error) {
	var resp V1FbpDraftDropOffDlvEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectCreate(ctx context.Context, req *V1FbpDraftDirectCreateRequest) (*V1FbpDraftDirectCreateResponse, error) {
	var resp V1FbpDraftDirectCreateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpCheckConsignmentNoteState(ctx context.Context, req *V1FbpCheckConsignmentNoteStateRequest) (*V1FbpCheckConsignmentNoteStateResponse, error) {
	var resp V1FbpCheckConsignmentNoteStateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/act-to/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderDropOffTimetable(ctx context.Context, req *V1FbpOrderDropOffTimetableRequest) (*V1FbpOrderDropOffTimetableResponse, error) {
	var resp V1FbpOrderDropOffTimetableResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/drop-off/timetable", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftPickUpRegistrate(ctx context.Context, req *V1FbpDraftPickUpRegistrateRequest) (*V1FbpDraftPickUpRegistrateResponse, error) {
	var resp V1FbpDraftPickUpRegistrateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/pick-up/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpCheckActState(ctx context.Context, req *V1FbpCheckActStateRequest) (*V1FbpCheckActStateResponse, error) {
	var resp V1FbpCheckActStateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/act-from/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectTplDlvCreate(ctx context.Context, req *V1FbpDraftDirectTplDlvCreateRequest) (*V1FbpDraftDirectTplDlvCreateResponse, error) {
	var resp V1FbpDraftDirectTplDlvCreateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/tpl-dlv/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffPointList(ctx context.Context, req *V1FbpDraftDropOffPointListRequest) (*V1FbpDraftDropOffPointListResponse, error) {
	var resp V1FbpDraftDropOffPointListResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/point/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpArchiveGet(ctx context.Context, req *V1FbpArchiveGetRequest) (*V1FbpArchiveGetResponse, error) {
	var resp V1FbpArchiveGetResponse
	_, err := s.t.Post(ctx, "/v1/fbp/archive/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderPickUpCancel(ctx context.Context, req *V1FbpOrderPickUpCancelRequest) (*V1FbpOrderPickUpCancelResponse, error) {
	var resp V1FbpOrderPickUpCancelResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/pick-up/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderDropOffDlvEdit(ctx context.Context, req *V1FbpOrderDropOffDlvEditRequest) (*V1FbpOrderDropOffDlvEditResponse, error) {
	var resp V1FbpOrderDropOffDlvEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/drop-off/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectDelete(ctx context.Context, req *V1FbpDraftDirectDeleteRequest) (*V1FbpDraftDirectDeleteResponse, error) {
	var resp V1FbpDraftDirectDeleteResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftList(ctx context.Context, req *V1FbpDraftListRequest) (*V1FbpDraftListResponse, error) {
	var resp V1FbpDraftListResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpEditTimeslot(ctx context.Context, req *V1FbpEditTimeslotRequest) (*V1FbpEditTimeslotResponse, error) {
	var resp V1FbpEditTimeslotResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/direct/timeslot/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpWarehouseList(ctx context.Context) (*V1FbpWarehouseListResponse, error) {
	var resp V1FbpWarehouseListResponse
	_, err := s.t.Post(ctx, "/v1/fbp/warehouse/list", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectTimeslotEdit(ctx context.Context, req *V1FbpDraftDirectTimeslotEditRequest) (*V1FbpDraftDirectTimeslotEditResponse, error) {
	var resp V1FbpDraftDirectTimeslotEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/timeslot/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftPickupDlvEdit(ctx context.Context, req *V1FbpDraftPickupDlvEditRequest) (*V1FbpDraftPickupDlvEditResponse, error) {
	var resp V1FbpDraftPickupDlvEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/pick-up/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffRegistrate(ctx context.Context, req *V1FbpDraftDropOffRegistrateRequest) (*V1FbpDraftDropOffRegistrateResponse, error) {
	var resp V1FbpDraftDropOffRegistrateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftPickupCreate(ctx context.Context, req *V1FbpDraftPickupCreateRequest) (*V1FbpDraftPickupCreateResponse, error) {
	var resp V1FbpDraftPickupCreateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/pick-up/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffPointTimetable(ctx context.Context, req *V1FbpDraftDropOffPointTimetableRequest) (*V1FbpDraftDropOffPointTimetableResponse, error) {
	var resp V1FbpDraftDropOffPointTimetableResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/point/timetable", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpGetLabel(ctx context.Context, req *V1FbpGetLabelRequest) (*V1FbpGetLabelResponse, error) {
	var resp V1FbpGetLabelResponse
	_, err := s.t.Post(ctx, "/v1/fbp/label/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderList(ctx context.Context, req *V1FbpOrderListRequest) (*V1FbpOrderListResponse, error) {
	var resp V1FbpOrderListResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderPickUpDlvEdit(ctx context.Context, req *V1FbpOrderPickUpDlvEditRequest) (*V1FbpOrderPickUpDlvEditResponse, error) {
	var resp V1FbpOrderPickUpDlvEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/pick-up/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpAvailableTimeslotList(ctx context.Context, req *V1FbpAvailableTimeslotListRequest) (*V1FbpAvailableTimeslotListResponse, error) {
	var resp V1FbpAvailableTimeslotListResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/direct/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffCreate(ctx context.Context, req *V1FbpDraftDropOffCreateRequest) (*V1FbpDraftDropOffCreateResponse, error) {
	var resp V1FbpDraftDropOffCreateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderDropOffCancel(ctx context.Context, req *V1FbpOrderDropOffCancelRequest) (*V1FbpOrderDropOffCancelResponse, error) {
	var resp V1FbpOrderDropOffCancelResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/drop-off/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderGet(ctx context.Context, req *V1FbpOrderGetRequest) (*V1FbpOrderGetResponse, error) {
	var resp V1FbpOrderGetResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) PostingFbpList(ctx context.Context, req *PostingV1PostingFbpListRequest) (*PostingV1PostingFbpListResponse, error) {
	var resp PostingV1PostingFbpListResponse
	_, err := s.t.Post(ctx, "/v1/posting/fbp/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderDirectSellerDlvEdit(ctx context.Context, req *V1FbpOrderDirectSellerDlvEditRequest) (*V1FbpOrderDirectSellerDlvEditResponse, error) {
	var resp V1FbpOrderDirectSellerDlvEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/direct/seller-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDropOffDelete(ctx context.Context, req *V1FbpDraftDropOffDeleteRequest) (*V1FbpDraftDropOffDeleteResponse, error) {
	var resp V1FbpDraftDropOffDeleteResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/drop-off/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpOrderDirectCancel(ctx context.Context, req *V1FbpOrderDirectCancelRequest) (*V1FbpOrderDirectCancelResponse, error) {
	var resp V1FbpOrderDirectCancelResponse
	_, err := s.t.Post(ctx, "/v1/fbp/order/direct/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectProductValidate(ctx context.Context, req *V1FbpDraftDirectProductValidateRequest) (*V1FbpDraftDirectProductValidateResponse, error) {
	var resp V1FbpDraftDirectProductValidateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectGetTimeslot(ctx context.Context, req *V1FbpDraftDirectGetTimeslotRequest) (*V1FbpDraftDirectGetTimeslotResponse, error) {
	var resp V1FbpDraftDirectGetTimeslotResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/timeslot/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpCreateLabel(ctx context.Context, req *V1FbpCreateLabelRequest) (*V1FbpCreateLabelResponse, error) {
	var resp V1FbpCreateLabelResponse
	_, err := s.t.Post(ctx, "/v1/fbp/label/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpCreateAct(ctx context.Context, req *V1FbpCreateActRequest) (*V1FbpCreateActResponse, error) {
	var resp V1FbpCreateActResponse
	_, err := s.t.Post(ctx, "/v1/fbp/act-from/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpCreateConsignmentNote(ctx context.Context, req *V1FbpCreateConsignmentNoteRequest) (*V1FbpCreateConsignmentNoteResponse, error) {
	var resp V1FbpCreateConsignmentNoteResponse
	_, err := s.t.Post(ctx, "/v1/fbp/act-to/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectSellerDlvEdit(ctx context.Context, req *V1FbpDraftDirectSellerDlvEditRequest) (*V1FbpDraftDirectSellerDlvEditResponse, error) {
	var resp V1FbpDraftDirectSellerDlvEditResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/seller-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpArchiveList(ctx context.Context, req *V1FbpArchiveListRequest) (*V1FbpArchiveListResponse, error) {
	var resp V1FbpArchiveListResponse
	_, err := s.t.Post(ctx, "/v1/fbp/archive/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) FbpDraftDirectSellerDlvCreate(ctx context.Context, req *V1FbpDraftDirectSellerDlvCreateRequest) (*V1FbpDraftDirectSellerDlvCreateResponse, error) {
	var resp V1FbpDraftDirectSellerDlvCreateResponse
	_, err := s.t.Post(ctx, "/v1/fbp/draft/direct/seller-dlv/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBOService) SupplyOrderBundle(ctx context.Context, req *V1GetSupplyOrderBundleRequest) (*V1GetSupplyOrderBundleResponse, error) {
	var resp V1GetSupplyOrderBundleResponse
	_, err := s.t.Post(ctx, "/v1/supply-order/bundle", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

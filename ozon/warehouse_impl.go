package ozon

import "context"

func (s *WarehouseService) UnarchiveWarehouseFBS(ctx context.Context, req *V1UnarchiveWarehouseFBSRequest) (*V1UnarchiveWarehouseFBSResponse, error) {
	var resp V1UnarchiveWarehouseFBSResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) DeliveryMethodListV2(ctx context.Context, req *V2DeliveryMethodListV2Request) (*V2DeliveryMethodListV2Response, error) {
	var resp V2DeliveryMethodListV2Response
	_, err := s.t.Post(ctx, "/v2/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseList(ctx context.Context, req *V1WarehouseListRequest) (*WarehouseWarehouseListResponse, error) {
	var resp WarehouseWarehouseListResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) ArchiveWarehouseFBS(ctx context.Context, req *V1ArchiveWarehouseFBSRequest) (*V1ArchiveWarehouseFBSResponse, error) {
	var resp V1ArchiveWarehouseFBSResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) GetWarehouseFBSOperationStatus(ctx context.Context, req *V1GetWarehouseFBSOperationStatusRequest) (*V1GetWarehouseFBSOperationStatusResponse, error) {
	var resp V1GetWarehouseFBSOperationStatusResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/operation/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseFbsCreatePickUpTimeslotList(ctx context.Context, req *V1WarehouseFbsCreatePickUpTimeslotListRequest) (*V1WarehouseFbsCreatePickUpTimeslotListResponse, error) {
	var resp V1WarehouseFbsCreatePickUpTimeslotListResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/create/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseFbsCreateDropOffTimeslotList(ctx context.Context, req *V1WarehouseFbsCreateDropOffTimeslotListRequest) (*V1WarehouseFbsCreateDropOffTimeslotListResponse, error) {
	var resp V1WarehouseFbsCreateDropOffTimeslotListResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/create/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) CreateWarehouseFBS(ctx context.Context, req *V1CreateWarehouseFBSRequest) (*V1CreateWarehouseFBSResponse, error) {
	var resp V1CreateWarehouseFBSResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) DeliveryMethodList(ctx context.Context, req *WarehouseDeliveryMethodListRequest) (*WarehouseDeliveryMethodListResponse, error) {
	var resp WarehouseDeliveryMethodListResponse
	_, err := s.t.Post(ctx, "/v1/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) UpdateWarehouseFBS(ctx context.Context, req *V1UpdateWarehouseFBSRequest) (*V1UpdateWarehouseFBSResponse, error) {
	var resp V1UpdateWarehouseFBSResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseListV2(ctx context.Context, req *V2WarehouseListV2Request) (*V2WarehouseListV2Response, error) {
	var resp V2WarehouseListV2Response
	_, err := s.t.Post(ctx, "/v2/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) ListDropOffPointsForUpdateFBSWarehouse(ctx context.Context, req *V1ListDropOffPointsForUpdateFBSWarehouseRequest) (*V1ListDropOffPointsForUpdateFBSWarehouseResponse, error) {
	var resp V1ListDropOffPointsForUpdateFBSWarehouseResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/update/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseInvalidProductsGet(ctx context.Context, req *V1WarehouseInvalidProductsGetRequest) (*V1WarehouseInvalidProductsGetResponse, error) {
	var resp V1WarehouseInvalidProductsGetResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/invalid-products/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseWithInvalidProducts(ctx context.Context) (*V1WarehouseWithInvalidProductsResponse, error) {
	var resp V1WarehouseWithInvalidProductsResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/warehouses-with-invalid-products", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) ListDropOffPointsForCreateFBSWarehouse(ctx context.Context, req *V1ListDropOffPointsForCreateFBSWarehouseRequest) (*V1ListDropOffPointsForCreateFBSWarehouseResponse, error) {
	var resp V1ListDropOffPointsForCreateFBSWarehouseResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/create/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) UpdateWarehouseFBSFirstMile(ctx context.Context, req *V1UpdateWarehouseFBSFirstMileRequest) (*V1UpdateWarehouseFBSFirstMileResponse, error) {
	var resp V1UpdateWarehouseFBSFirstMileResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/first-mile/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseFbsUpdateDropOffTimeslotList(ctx context.Context, req *V1WarehouseFbsUpdateDropOffTimeslotListRequest) (*V1WarehouseFbsUpdateDropOffTimeslotListResponse, error) {
	var resp V1WarehouseFbsUpdateDropOffTimeslotListResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/update/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *WarehouseService) WarehouseFbsUpdatePickUpTimeslotList(ctx context.Context, req *V1WarehouseFbsUpdatePickUpTimeslotListRequest) (*V1WarehouseFbsUpdatePickUpTimeslotListResponse, error) {
	var resp V1WarehouseFbsUpdatePickUpTimeslotListResponse
	_, err := s.t.Post(ctx, "/v1/warehouse/fbs/update/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

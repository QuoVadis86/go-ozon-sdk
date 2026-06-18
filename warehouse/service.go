package warehouse

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) WarehouseInvalidProductsGet(ctx context.Context, req *types.V1WarehouseInvalidProductsGetRequest) (*types.V1WarehouseInvalidProductsGetResponse, error) {
	var resp types.V1WarehouseInvalidProductsGetResponse
	err := s.Client.Post(ctx, "/v1/warehouse/invalid-products/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseListV2(ctx context.Context, req *types.V2WarehouseListV2Request) (*types.V2WarehouseListV2Response, error) {
	var resp types.V2WarehouseListV2Response
	err := s.Client.Post(ctx, "/v2/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseList(ctx context.Context, req *types.V1WarehouseListRequest) (*types.WarehouseWarehouseListResponse, error) {
	var resp types.WarehouseWarehouseListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DeliveryMethodListV2(ctx context.Context, req *types.V2DeliveryMethodListV2Request) (*types.V2DeliveryMethodListV2Response, error) {
	var resp types.V2DeliveryMethodListV2Response
	err := s.Client.Post(ctx, "/v2/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetWarehouseFBSOperationStatus(ctx context.Context, req *types.V1GetWarehouseFBSOperationStatusRequest) (*types.V1GetWarehouseFBSOperationStatusResponse, error) {
	var resp types.V1GetWarehouseFBSOperationStatusResponse
	err := s.Client.Post(ctx, "/v1/warehouse/operation/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsCreateDropOffTimeslotList(ctx context.Context, req *types.V1WarehouseFbsCreateDropOffTimeslotListRequest) (*types.V1WarehouseFbsCreateDropOffTimeslotListResponse, error) {
	var resp types.V1WarehouseFbsCreateDropOffTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateWarehouseFBS(ctx context.Context, req *types.V1CreateWarehouseFBSRequest) (*types.V1CreateWarehouseFBSResponse, error) {
	var resp types.V1CreateWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseWithInvalidProducts(ctx context.Context) (*types.V1WarehouseWithInvalidProductsResponse, error) {
	var resp types.V1WarehouseWithInvalidProductsResponse
	err := s.Client.Post(ctx, "/v1/warehouse/warehouses-with-invalid-products", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) UpdateWarehouseFBSFirstMile(ctx context.Context, req *types.V1UpdateWarehouseFBSFirstMileRequest) (*types.V1UpdateWarehouseFBSFirstMileResponse, error) {
	var resp types.V1UpdateWarehouseFBSFirstMileResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/first-mile/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ArchiveWarehouseFBS(ctx context.Context, req *types.V1ArchiveWarehouseFBSRequest) (*types.V1ArchiveWarehouseFBSResponse, error) {
	var resp types.V1ArchiveWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) UpdateWarehouseFBS(ctx context.Context, req *types.V1UpdateWarehouseFBSRequest) (*types.V1UpdateWarehouseFBSResponse, error) {
	var resp types.V1UpdateWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsUpdatePickUpTimeslotList(ctx context.Context, req *types.V1WarehouseFbsUpdatePickUpTimeslotListRequest) (*types.V1WarehouseFbsUpdatePickUpTimeslotListResponse, error) {
	var resp types.V1WarehouseFbsUpdatePickUpTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsCreatePickUpTimeslotList(ctx context.Context, req *types.V1WarehouseFbsCreatePickUpTimeslotListRequest) (*types.V1WarehouseFbsCreatePickUpTimeslotListResponse, error) {
	var resp types.V1WarehouseFbsCreatePickUpTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListDropOffPointsForCreateFBSWarehouse(ctx context.Context, req *types.V1ListDropOffPointsForCreateFBSWarehouseRequest) (*types.V1ListDropOffPointsForCreateFBSWarehouseResponse, error) {
	var resp types.V1ListDropOffPointsForCreateFBSWarehouseResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DeliveryMethodList(ctx context.Context, req *types.WarehouseDeliveryMethodListRequest) (*types.WarehouseDeliveryMethodListResponse, error) {
	var resp types.WarehouseDeliveryMethodListResponse
	err := s.Client.Post(ctx, "/v1/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsUpdateDropOffTimeslotList(ctx context.Context, req *types.V1WarehouseFbsUpdateDropOffTimeslotListRequest) (*types.V1WarehouseFbsUpdateDropOffTimeslotListResponse, error) {
	var resp types.V1WarehouseFbsUpdateDropOffTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) UnarchiveWarehouseFBS(ctx context.Context, req *types.V1UnarchiveWarehouseFBSRequest) (*types.V1UnarchiveWarehouseFBSResponse, error) {
	var resp types.V1UnarchiveWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListDropOffPointsForUpdateFBSWarehouse(ctx context.Context, req *types.V1ListDropOffPointsForUpdateFBSWarehouseRequest) (*types.V1ListDropOffPointsForUpdateFBSWarehouseResponse, error) {
	var resp types.V1ListDropOffPointsForUpdateFBSWarehouseResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

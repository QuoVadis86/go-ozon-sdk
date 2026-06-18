package warehouse

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) ArchiveWarehouseFBS(ctx context.Context, req *internal.V1ArchiveWarehouseFBSRequest) (*internal.V1ArchiveWarehouseFBSResponse, error) {
	var resp internal.V1ArchiveWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsUpdatePickUpTimeslotList(ctx context.Context, req *internal.V1WarehouseFbsUpdatePickUpTimeslotListRequest) (*internal.V1WarehouseFbsUpdatePickUpTimeslotListResponse, error) {
	var resp internal.V1WarehouseFbsUpdatePickUpTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseWithInvalidProducts(ctx context.Context) (*internal.V1WarehouseWithInvalidProductsResponse, error) {
	var resp internal.V1WarehouseWithInvalidProductsResponse
	err := s.Client.Post(ctx, "/v1/warehouse/warehouses-with-invalid-products", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) UnarchiveWarehouseFBS(ctx context.Context, req *internal.V1UnarchiveWarehouseFBSRequest) (*internal.V1UnarchiveWarehouseFBSResponse, error) {
	var resp internal.V1UnarchiveWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListDropOffPointsForUpdateFBSWarehouse(ctx context.Context, req *internal.V1ListDropOffPointsForUpdateFBSWarehouseRequest) (*internal.V1ListDropOffPointsForUpdateFBSWarehouseResponse, error) {
	var resp internal.V1ListDropOffPointsForUpdateFBSWarehouseResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseListV2(ctx context.Context, req *internal.V2WarehouseListV2Request) (*internal.V2WarehouseListV2Response, error) {
	var resp internal.V2WarehouseListV2Response
	err := s.Client.Post(ctx, "/v2/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsUpdateDropOffTimeslotList(ctx context.Context, req *internal.V1WarehouseFbsUpdateDropOffTimeslotListRequest) (*internal.V1WarehouseFbsUpdateDropOffTimeslotListResponse, error) {
	var resp internal.V1WarehouseFbsUpdateDropOffTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsCreatePickUpTimeslotList(ctx context.Context, req *internal.V1WarehouseFbsCreatePickUpTimeslotListRequest) (*internal.V1WarehouseFbsCreatePickUpTimeslotListResponse, error) {
	var resp internal.V1WarehouseFbsCreatePickUpTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DeliveryMethodList(ctx context.Context, req *internal.WarehouseDeliveryMethodListRequest) (*internal.WarehouseDeliveryMethodListResponse, error) {
	var resp internal.WarehouseDeliveryMethodListResponse
	err := s.Client.Post(ctx, "/v1/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseInvalidProductsGet(ctx context.Context, req *internal.V1WarehouseInvalidProductsGetRequest) (*internal.V1WarehouseInvalidProductsGetResponse, error) {
	var resp internal.V1WarehouseInvalidProductsGetResponse
	err := s.Client.Post(ctx, "/v1/warehouse/invalid-products/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseFbsCreateDropOffTimeslotList(ctx context.Context, req *internal.V1WarehouseFbsCreateDropOffTimeslotListRequest) (*internal.V1WarehouseFbsCreateDropOffTimeslotListResponse, error) {
	var resp internal.V1WarehouseFbsCreateDropOffTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) DeliveryMethodListV2(ctx context.Context, req *internal.V2DeliveryMethodListV2Request) (*internal.V2DeliveryMethodListV2Response, error) {
	var resp internal.V2DeliveryMethodListV2Response
	err := s.Client.Post(ctx, "/v2/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreateWarehouseFBS(ctx context.Context, req *internal.V1CreateWarehouseFBSRequest) (*internal.V1CreateWarehouseFBSResponse, error) {
	var resp internal.V1CreateWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListDropOffPointsForCreateFBSWarehouse(ctx context.Context, req *internal.V1ListDropOffPointsForCreateFBSWarehouseRequest) (*internal.V1ListDropOffPointsForCreateFBSWarehouseResponse, error) {
	var resp internal.V1ListDropOffPointsForCreateFBSWarehouseResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) WarehouseList(ctx context.Context, req *internal.V1WarehouseListRequest) (*internal.WarehouseWarehouseListResponse, error) {
	var resp internal.WarehouseWarehouseListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) UpdateWarehouseFBS(ctx context.Context, req *internal.V1UpdateWarehouseFBSRequest) (*internal.V1UpdateWarehouseFBSResponse, error) {
	var resp internal.V1UpdateWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetWarehouseFBSOperationStatus(ctx context.Context, req *internal.V1GetWarehouseFBSOperationStatusRequest) (*internal.V1GetWarehouseFBSOperationStatusResponse, error) {
	var resp internal.V1GetWarehouseFBSOperationStatusResponse
	err := s.Client.Post(ctx, "/v1/warehouse/operation/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) UpdateWarehouseFBSFirstMile(ctx context.Context, req *internal.V1UpdateWarehouseFBSFirstMileRequest) (*internal.V1UpdateWarehouseFBSFirstMileResponse, error) {
	var resp internal.V1UpdateWarehouseFBSFirstMileResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/first-mile/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

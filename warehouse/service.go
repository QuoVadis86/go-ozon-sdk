package warehouse

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

// 获取用于创建仓库的揽收点列表
func (s *Service) ListDropOffPointsForCreateFBSWarehouse(ctx context.Context, req *V1ListDropOffPointsForCreateFBSWarehouseRequest) (*V1ListDropOffPointsForCreateFBSWarehouseResponse, error) {
	var resp V1ListDropOffPointsForCreateFBSWarehouseResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取用于创建drop-off发运仓库的时间段列表
func (s *Service) WarehouseFbsCreateDropOffTimeslotList(ctx context.Context, req *V1WarehouseFbsCreateDropOffTimeslotListRequest) (*V1WarehouseFbsCreateDropOffTimeslotListResponse, error) {
	var resp V1WarehouseFbsCreateDropOffTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 仓库清单
func (s *Service) WarehouseList(ctx context.Context, req *V1WarehouseListRequest) (*WarehouseWarehouseListResponse, error) {
	var resp WarehouseWarehouseListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 仓库列表
func (s *Service) WarehouseListV2(ctx context.Context, req *V2WarehouseListV2Request) (*V2WarehouseListV2Response, error) {
	var resp V2WarehouseListV2Response
	err := s.Client.Post(ctx, "/v2/warehouse/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取用于修改仓库信息的揽收点列表
func (s *Service) ListDropOffPointsForUpdateFBSWarehouse(ctx context.Context, req *V1ListDropOffPointsForUpdateFBSWarehouseRequest) (*V1ListDropOffPointsForUpdateFBSWarehouseResponse, error) {
	var resp V1ListDropOffPointsForUpdateFBSWarehouseResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/drop-off/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新头程物流
func (s *Service) UpdateWarehouseFBSFirstMile(ctx context.Context, req *V1UpdateWarehouseFBSFirstMileRequest) (*V1UpdateWarehouseFBSFirstMileResponse, error) {
	var resp V1UpdateWarehouseFBSFirstMileResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/first-mile/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// realFBS仓库的配送方式列表
func (s *Service) DeliveryMethodListV2(ctx context.Context, req *V2DeliveryMethodListV2Request) (*V2DeliveryMethodListV2Response, error) {
	var resp V2DeliveryMethodListV2Response
	err := s.Client.Post(ctx, "/v2/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 仓库物流方式清单
func (s *Service) DeliveryMethodList(ctx context.Context, req *WarehouseDeliveryMethodListRequest) (*WarehouseDeliveryMethodListResponse, error) {
	var resp WarehouseDeliveryMethodListResponse
	err := s.Client.Post(ctx, "/v1/delivery-method/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将仓库解除归档
func (s *Service) UnarchiveWarehouseFBS(ctx context.Context, req *V1UnarchiveWarehouseFBSRequest) (*V1UnarchiveWarehouseFBSResponse, error) {
	var resp V1UnarchiveWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/unarchive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新仓库
func (s *Service) UpdateWarehouseFBS(ctx context.Context, req *V1UpdateWarehouseFBSRequest) (*V1UpdateWarehouseFBSResponse, error) {
	var resp V1UpdateWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取用于更新pick-up发运仓库的时间段列表
func (s *Service) WarehouseFbsUpdatePickUpTimeslotList(ctx context.Context, req *V1WarehouseFbsUpdatePickUpTimeslotListRequest) (*V1WarehouseFbsUpdatePickUpTimeslotListResponse, error) {
	var resp V1WarehouseFbsUpdatePickUpTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取用于创建pick-up发运仓库的时间段列表
func (s *Service) WarehouseFbsCreatePickUpTimeslotList(ctx context.Context, req *V1WarehouseFbsCreatePickUpTimeslotListRequest) (*V1WarehouseFbsCreatePickUpTimeslotListResponse, error) {
	var resp V1WarehouseFbsCreatePickUpTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create/pick-up/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取操作状态
func (s *Service) GetWarehouseFBSOperationStatus(ctx context.Context, req *V1GetWarehouseFBSOperationStatusRequest) (*V1GetWarehouseFBSOperationStatusResponse, error) {
	var resp V1GetWarehouseFBSOperationStatusResponse
	err := s.Client.Post(ctx, "/v1/warehouse/operation/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建仓库
func (s *Service) CreateWarehouseFBS(ctx context.Context, req *V1CreateWarehouseFBSRequest) (*V1CreateWarehouseFBSResponse, error) {
	var resp V1CreateWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将仓库归档
func (s *Service) ArchiveWarehouseFBS(ctx context.Context, req *V1ArchiveWarehouseFBSRequest) (*V1ArchiveWarehouseFBSResponse, error) {
	var resp V1ArchiveWarehouseFBSResponse
	err := s.Client.Post(ctx, "/v1/warehouse/archive", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取含有配送受限商品的仓库列表
func (s *Service) WarehouseWithInvalidProducts(ctx context.Context) (*V1WarehouseWithInvalidProductsResponse, error) {
	var resp V1WarehouseWithInvalidProductsResponse
	err := s.Client.Post(ctx, "/v1/warehouse/warehouses-with-invalid-products", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取配送受限商品列表
func (s *Service) WarehouseInvalidProductsGet(ctx context.Context, req *V1WarehouseInvalidProductsGetRequest) (*V1WarehouseInvalidProductsGetResponse, error) {
	var resp V1WarehouseInvalidProductsGetResponse
	err := s.Client.Post(ctx, "/v1/warehouse/invalid-products/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取用于更新drop-off发运仓库的时间段列表
func (s *Service) WarehouseFbsUpdateDropOffTimeslotList(ctx context.Context, req *V1WarehouseFbsUpdateDropOffTimeslotListRequest) (*V1WarehouseFbsUpdateDropOffTimeslotListResponse, error) {
	var resp V1WarehouseFbsUpdateDropOffTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/warehouse/fbs/update/drop-off/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package warehouse

type WarehouseDeliveryMethodListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type V2DeliveryMethodListV2Request struct {
	SortDir interface{} `json:"sort_dir"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"`
}

type V1GetWarehouseFBSOperationStatusRequest struct {
	OperationID string `json:"operation_id"`
}

type V1ListDropOffPointsForCreateFBSWarehouseRequest struct {
	Coordinates interface{} `json:"coordinates"`
	CountryCode string `json:"country_code"`
	IsKgt bool `json:"is_kgt"`
	Search interface{} `json:"search"`
}

type V1CreateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V1WarehouseFbsCreatePickUpTimeslotListRequest struct {
	AddressCoordinates interface{} `json:"address_coordinates"`
	IsKgt bool `json:"is_kgt"`
}

type V1UnarchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V2DeliveryMethodListV2Response struct {
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
	DeliveryMethods []interface{} `json:"delivery_methods"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V1UpdateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V2WarehouseListV2Response struct {
	Cursor string `json:"cursor"`
	Warehouses []interface{} `json:"warehouses"`
	HasNext bool `json:"has_next"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type V2WarehouseListV2Request struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
}

type V1ArchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"`
}

type V1WarehouseWithInvalidProductsResponse struct {
	WarehouseIds []interface{} `json:"warehouse_ids"`
}

type V1WarehouseFbsCreateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseRequest struct {
	Search interface{} `json:"search"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1GetWarehouseFBSOperationStatusResponse struct {
	Error interface{} `json:"error"`
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Type interface{} `json:"type_"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1CreateWarehouseFBSRequest struct {
	WorkingDays []interface{} `json:"working_days"`
	AddressCoordinates interface{} `json:"address_coordinates"`
	CutInTime int64 `json:"cut_in_time"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	FirstMileType interface{} `json:"first_mile_type"`
	Name string `json:"name"`
	TimeslotID int64 `json:"timeslot_id"`
	IsKgt bool `json:"is_kgt"`
	Options interface{} `json:"options"`
	Phone string `json:"phone"`
}

type V1UpdateWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	WorkingDays []interface{} `json:"working_days"`
	AddressCoordinates interface{} `json:"address_coordinates"`
	Name string `json:"name"`
	Options interface{} `json:"options"`
	Phone string `json:"phone"`
}

type V1WarehouseInvalidProductsGetResponse struct {
	HasNext bool `json:"has_next"`
	LastID int64 `json:"last_id"`
	ValidationResults []interface{} `json:"validation_results"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1UpdateWarehouseFBSFirstMileRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	CutInTime int64 `json:"cut_in_time"`
	DropOffPointID int64 `json:"drop_off_point_id"`
	FirstMileType interface{} `json:"first_mile_type"`
	TimeslotID int64 `json:"timeslot_id"`
}

type V1WarehouseInvalidProductsGetRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
	LastID int64 `json:"last_id"`
}

type V1UnarchiveWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"`
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
}

type V1WarehouseFbsCreatePickUpTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"`
	IsPickupSupported bool `json:"is_pickup_supported"`
}

type V1UpdateWarehouseFBSFirstMileResponse struct {
	OperationID string `json:"operation_id"`
}

type WarehouseDeliveryMethodListResponse struct {
	HasNext bool `json:"has_next"`
	Result []interface{} `json:"result"`
}

type V1ArchiveWarehouseFBSRequest struct {
	Reason string `json:"reason"`
	WarehouseID int64 `json:"warehouse_id"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"`
}

type V1WarehouseListRequest struct {
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	With interface{} `json:"with"`
}

type WarehouseWarehouseListResponse struct {
	Result []interface{} `json:"result"`
}

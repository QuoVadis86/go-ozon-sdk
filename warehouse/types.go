package warehouse

type V2WarehouseListV2Request struct {
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	Cursor string `json:"cursor"` // 后续数据的选择标志。
	WarehouseIds []interface{} `json:"warehouse_ids"` // 仓库识别符。
}

type V1GetWarehouseFBSOperationStatusResponse struct {
	Error interface{} `json:"error"`
	Result interface{} `json:"result"`
	Status interface{} `json:"status"`
	Type interface{} `json:"type_"`
}

type V1WarehouseListRequest struct {
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	Offset int64 `json:"offset"` // 在响应中将被跳过的项目数量。例如，如果 `offset = 10`，则响应将从第 11 个找到的项目开始。
	With interface{} `json:"with"`
}

type V1ArchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 操作识别符。请使用 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 方式获取操作状态。
}

type V1CreateWarehouseFBSRequest struct {
	CutInTime int64 `json:"cut_in_time"` // 接单时间（分钟）。例如，如果传入 `3000`，则接单将在传入后50小时内结束。
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	Options interface{} `json:"options"`
	Phone string `json:"phone"` // 仓库联系电话。请按格式填写：+7 (XXX) XXX-XX-XX。
	TimeslotID int64 `json:"timeslot_id"` // 时间段标识符。
	WorkingDays []interface{} `json:"working_days"` // 仓库工作日： - `MONDAY` — 星期一， - `TUESDAY` — 星期二， - `WEDNESDAY` — 星期三， - `THURSDAY` — 星期四， - `FRIDAY` — 星期五， - `SATURDAY` — 星期...
	AddressCoordinates interface{} `json:"address_coordinates"`
	FirstMileType interface{} `json:"first_mile_type"`
	IsKgt bool `json:"is_kgt"` // `true`表示商品为超大。
	Name string `json:"name"` // 仓库名称。
}

type V1WarehouseFbsCreateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"` // 时间段列表。
}

type V1WarehouseFbsCreatePickUpTimeslotListResponse struct {
	IsPickupSupported bool `json:"is_pickup_supported"` // 支持pick-up发运标志。
	Timeslots []interface{} `json:"timeslots"` // 时间段列表。
}

type V1CreateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 创建FBS仓库的操作标识符。要获取操作状态，请使用方法 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus)。
}

type V2DeliveryMethodListV2Request struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortDir interface{} `json:"sort_dir"`
}

type V1ListDropOffPointsForUpdateFBSWarehouseRequest struct {
	Search interface{} `json:"search"`
	WarehouseID int64 `json:"warehouse_id"` // 现有FBS仓库的筛选条件。
}

type V1ListDropOffPointsForCreateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"` // 点位列表。
}

type V1WarehouseWithInvalidProductsResponse struct {
	WarehouseIds []interface{} `json:"warehouse_ids"` // 包含至少 1 件无法从该仓库发运的受限商品的仓库标识符列表。如需获取受限商品列表，请使用方法 [/v1/warehouse/invalid-products/get](#operation/WarehouseInvalidProductsG...
}

type V1UpdateWarehouseFBSFirstMileRequest struct {
	CutInTime int64 `json:"cut_in_time"` // 接单时间（分钟）。例如，如果传入 `3000`，则接单将在传入后50小时内结束。
	DropOffPointID int64 `json:"drop_off_point_id"` // 订单发运点ID。如果` first_mile_type = DROP_OFF`，该参数必填。
	FirstMileType interface{} `json:"first_mile_type"`
	TimeslotID int64 `json:"timeslot_id"` // 时间段标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库ID。
}

type V2WarehouseListV2Response struct {
	Cursor string `json:"cursor"` // 后续数据的选择标志。
	Warehouses []interface{} `json:"warehouses"` // 仓库列表。
	HasNext bool `json:"has_next"` // `true`，前提是本次响应未返回所有数据。
}

type V1UpdateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 操作ID。通过方法 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 获取操作状态。
}

type V1WarehouseFbsCreateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
}

type V1UnarchiveWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库识别符。
}

type V1WarehouseFbsCreatePickUpTimeslotListRequest struct {
	AddressCoordinates interface{} `json:"address_coordinates"`
	IsKgt bool `json:"is_kgt"` // 超大货物标志。
}

type V1WarehouseInvalidProductsGetResponse struct {
	HasNext bool `json:"has_next"` // 若响应中未包含全部商品，则为`true`。
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。要获取下一个批次的数据，请在下一个请求的 `last_id` 参数中传递上次获取的值。
	ValidationResults []interface{} `json:"validation_results"` // 检查结果。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"` // 时间段列表。
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponse struct {
	Timeslots []interface{} `json:"timeslots"` // 时间段列表。
}

type V1ListDropOffPointsForCreateFBSWarehouseRequest struct {
	Search interface{} `json:"search"`
	Coordinates interface{} `json:"coordinates"`
	CountryCode string `json:"country_code"` // ISO 2格式的国家代码。
	IsKgt bool `json:"is_kgt"` // `true`表示商品为超大。
}

type V2DeliveryMethodListV2Response struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	HasNext bool `json:"has_next"` // `true`，表示响应中未返回全部配送方式。
	DeliveryMethods []interface{} `json:"delivery_methods"` // 配送方式。
}

type V1ListDropOffPointsForUpdateFBSWarehouseResponse struct {
	Points []interface{} `json:"points"` // 点位列表。
}

type V1WarehouseFbsUpdatePickUpTimeslotListRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1ArchiveWarehouseFBSRequest struct {
	Reason string `json:"reason"` // 仓库归档原因。
	WarehouseID int64 `json:"warehouse_id"` // 仓库识别符。
}

type V1WarehouseFbsUpdateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1UpdateWarehouseFBSRequest struct {
	Options interface{} `json:"options"`
	Phone string `json:"phone"` // 仓库电话号码。
	WarehouseID int64 `json:"warehouse_id"` // 仓库ID。
	WorkingDays []interface{} `json:"working_days"` // 仓库工作日： - `MONDAY` — 周一； - `TUESDAY` — 周二； - `WEDNESDAY` — 周三； - `THURSDAY` — 周四； - `FRIDAY` — 周五； - `SATURDAY` — 周六； - `...
	AddressCoordinates interface{} `json:"address_coordinates"`
	Name string `json:"name"` // 仓库名称。
}

type V1UpdateWarehouseFBSFirstMileResponse struct {
	OperationID string `json:"operation_id"` // 操作ID。通过方法 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 获取操作状态。
}

type V1UnarchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 操作识别符。请使用 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 方式获取操作状态。
}

type WarehouseDeliveryMethodListRequest struct {
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 回答中的元素数量。最多50，最少1。
	Offset int64 `json:"offset"` // 回答中会被略过的元素数量。例如，如果`offset = 10`，回答将从发现的第11个元素开始。
}

type V1GetWarehouseFBSOperationStatusRequest struct {
	OperationID string `json:"operation_id"` // 操作ID。
}

type V1WarehouseInvalidProductsGetRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。请通过方法[/v1/warehouse/warehouses-with-invalid-products](#operation/WarehouseWithInvalidProducts)获取该参数值。
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。首次请求时请将此字段留空。 如需获取后续数据，请填写上次响应中的 `last_id`。
}

type WarehouseDeliveryMethodListResponse struct {
	Result []interface{} `json:"result"` // 查询结果。
	HasNext bool `json:"has_next"` // 以下该迹象会表明在查询中只送回了部分快递方式。 - `true` — 请用新的 `offset` 参数重新请求，以获得剩余的方式； - `false` — 回答中包含了所有应要求的快递方式。
}

type WarehouseWarehouseListResponse struct {
	Result []interface{} `json:"result"` // 仓库清单。
}

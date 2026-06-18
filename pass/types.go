package pass

// 筛选器。
// 按入场目的筛选：
type ArrivalPassListRequestFilterArrivalReasonEnum string
const (
	ArrivalPassListRequestFilterArrivalReasonEnum_FBSDELIVERY ArrivalPassListRequestFilterArrivalReasonEnum = "FBS_DELIVERY"
	ArrivalPassListRequestFilterArrivalReasonEnum_FBSRETURN ArrivalPassListRequestFilterArrivalReasonEnum = "FBS_RETURN"
)

type ArrivalPassListRequestFilter struct {
	OnlyActivePasses bool `json:"only_active_passes"` // `true`, 以获取仅活跃的通行证申请。
	WarehouseIds []string `json:"warehouse_ids"` // 按卖家仓库筛选。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 按通行证ID筛选。
	ArrivalReason ArrivalPassListRequestFilterArrivalReasonEnum `json:"arrival_reason"` // 按入场目的筛选： - `FBS_DELIVERY` — 发运。 - `FBS_RETURN` — 运出退货。 如果未指定此参数，则考虑两种目的。 指定的原因必须在通行证的原因列表中。
	DropoffPointIds []string `json:"dropoff_point_ids"` // 按发运点筛选。
}

// 通行证信息。
type ReturnsCompanyFbsInfoResponsePassInfo struct {
	Count int32 `json:"count"` // 每个揽收点的通行证数量。
	IsRequired bool `json:"is_required"` // 是否需要揽收点通行证的标志。
}

type ReturnsCompanyFbsInfoResponseDropOffPoints struct {
	PassInfo ReturnsCompanyFbsInfoResponsePassInfo `json:"pass_info"`
	PlaceID int64 `json:"place_id"` // 到货仓库的ID。
	UtcOffset string `json:"utc_offset"` // 发运时间与UTC-0的时区偏移量。
	Address string `json:"address"` // 揽收点地址。
	BoxCount int32 `json:"box_count"` // 在揽收点的箱数。
	ID int64 `json:"id"` // 揽收点ID。
	Name string `json:"name"` // 揽收点名称。
	ReturnsCount int32 `json:"returns_count"` // 揽收点的退货数量。
	WarehousesIds []string `json:"warehouses_ids"` // 卖家仓库ID。
}

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []ReturnsCompanyFbsInfoResponseDropOffPoints `json:"drop_off_points"` // 揽收点信息。
	HasNext bool `json:"has_next"` // 是否还有其他揽收点等待卖家退货的标志。
}

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证ID。
}

// 请求结果。
type SellerSellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证ID。
}

type ArrivalpassArrivalPassListRequest struct {
	Limit int32 `json:"limit"` // 响应中记录数量的限制。
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。
	Filter ArrivalPassListRequestFilter `json:"filter"`
}

// 筛选器。
type V1ReturnsCompanyFbsInfoRequestFilter struct {
	PlaceID int64 `json:"place_id"` // 按揽收点ID筛选。
}

// 方法响应的分割。
type ReturnsCompanyFbsInfoRequestPagination struct {
	Limit int32 `json:"limit"` // 页面上揽收点的数量。最大值为500。
	LastID int64 `json:"last_id"` // 页面上最后一个揽收点的ID。对于第一个请求，请将此字段留空。 要获取后续的值，请指定上一个请求响应中最后一个揽收点的`id`。
}

type V1ReturnsCompanyFbsInfoRequest struct {
	Pagination ReturnsCompanyFbsInfoRequestPagination `json:"pagination"`
	Filter V1ReturnsCompanyFbsInfoRequestFilter `json:"filter"`
}

type SellerSellerAPIArrivalPassCreateRequestArrivalPass struct {
	DriverPhone string `json:"driver_phone"` // 司机的电话号码。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
	VehicleModel string `json:"vehicle_model"` // 车辆型号。
	WithReturns bool `json:"with_returns"` // 如果要运出退货，则为`true`。 默认值为`false`。
	DriverName string `json:"driver_name"` // 司机的姓名。
}

type SellerSellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []SellerSellerAPIArrivalPassCreateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
	CarriageID int64 `json:"carriage_id"` // 运输ID。
}

type SellerSellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证列表。
	CarriageID int64 `json:"carriage_id"` // 运输ID。
}

type ArrivalpassArrivalPassUpdateRequestArrivalPass struct {
	DriverPhone string `json:"driver_phone"` // 司机的电话号码。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
	VehicleModel string `json:"vehicle_model"` // 车辆型号。
	ArrivalPassID int64 `json:"arrival_pass_id"` // 通行证ID。
	ArrivalTime string `json:"arrival_time"` // 入场时间，UTC格式。 此时通行证将开始生效。 要更改入场时间，请使用方法 [/v1/carriage/pass/update](#operation/carriagePassUpdate)。
	DriverName string `json:"driver_name"` // 司机的姓名。
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证ID。
}

type ArrivalpassArrivalPassCreateRequestArrivalPass struct {
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库ID。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	ArrivalTime string `json:"arrival_time"` // 入场时间，UTC格式。 此时通行证将开始生效。
	DriverName string `json:"driver_name"` // 司机姓名。
	DriverPhone string `json:"driver_phone"` // 司机电话号码。
	DropoffPointID int64 `json:"dropoff_point_id"` // 通行证适用的仓库ID。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
	VehicleModel string `json:"vehicle_model"` // 车辆型号。
}

type SellerSellerAPIArrivalPassUpdateRequestArrivalPass struct {
	VehicleModel string `json:"vehicle_model"` // 车辆型号。
	WithReturns bool `json:"with_returns"` // 如果要运出退货，则为`true`。 默认值为`false`。
	DriverName string `json:"driver_name"` // 司机的姓名。
	DriverPhone string `json:"driver_phone"` // 司机的电话号码。
	ID int64 `json:"id"` // 通行证ID。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
}

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []ArrivalpassArrivalPassUpdateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
}

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []ArrivalpassArrivalPassCreateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
}

type ArrivalpassArrivalPassListResponseArrivalPass struct {
	ArrivalReasons []string `json:"arrival_reasons"` // 入场目的。
	DriverName string `json:"driver_name"` // 司机的姓名。
	DropoffPointID int64 `json:"dropoff_point_id"` // 发运点ID。
	ArrivalPassID int64 `json:"arrival_pass_id"` // 通行证ID。
	ArrivalTime string `json:"arrival_time"` // 入场日期和时间，UTC格式。
	DriverPhone string `json:"driver_phone"` // 司机的电话号码。
	IsActive bool `json:"is_active"` // 如果申请是活跃的，则为`true`。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
	VehicleModel string `json:"vehicle_model"` // 车辆型号。
	WarehouseID int64 `json:"warehouse_id"` // 卖家仓库ID。
}

type ArrivalpassArrivalPassListResponse struct {
	ArrivalPasses []ArrivalpassArrivalPassListResponseArrivalPass `json:"arrival_passes"` // 运输通行证列表。
	Cursor string `json:"cursor"` // 用于获取下一批数据的指针。 如果此参数为空，则没有更多数据。
}

type SellerSellerAPIArrivalPassUpdateRequest struct {
	ArrivalPasses []SellerSellerAPIArrivalPassUpdateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
	CarriageID int64 `json:"carriage_id"` // 运输ID。
}

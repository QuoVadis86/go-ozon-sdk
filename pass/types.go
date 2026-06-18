package pass

// 通行证信息。
type CompanyFbsInfoResponsePassInfo struct {
	Count      int32 `json:"count"`       // 每个揽收点的通行证数量。
	IsRequired bool  `json:"is_required"` // 是否需要揽收点通行证的标志。
}

type CompanyFbsInfoResponseDropOffPoints struct {
	Address       string                         `json:"address"`    // 揽收点地址。
	ID            int64                          `json:"id"`         // 揽收点ID。
	Name          string                         `json:"name"`       // 揽收点名称。
	UtcOffset     string                         `json:"utc_offset"` // 发运时间与UTC-0的时区偏移量。
	BoxCount      int32                          `json:"box_count"`  // 在揽收点的箱数。
	PassInfo      CompanyFbsInfoResponsePassInfo `json:"pass_info"`
	PlaceID       int64                          `json:"place_id"`       // 到货仓库的ID。
	ReturnsCount  int32                          `json:"returns_count"`  // 揽收点的退货数量。
	WarehousesIds []string                       `json:"warehouses_ids"` // 卖家仓库ID。
}

type SellerAPIArrivalPassCreateRequestArrivalPass struct {
	VehicleModel        string `json:"vehicle_model"`         // 车辆型号。
	WithReturns         bool   `json:"with_returns"`          // 如果要运出退货，则为`true`。 默认值为`false`。
	DriverName          string `json:"driver_name"`           // 司机的姓名。
	DriverPhone         string `json:"driver_phone"`          // 司机的电话号码。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
}

type SellerAPIArrivalPassUpdateRequestArrivalPass struct {
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
	VehicleModel        string `json:"vehicle_model"`         // 车辆型号。
	WithReturns         bool   `json:"with_returns"`          // 如果要运出退货，则为`true`。 默认值为`false`。
	DriverName          string `json:"driver_name"`           // 司机的姓名。
	DriverPhone         string `json:"driver_phone"`          // 司机的电话号码。
	ID                  int64  `json:"id"`                    // 通行证ID。
}

type SellerAPIArrivalPassDeleteRequest struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证列表。
	CarriageID     int64    `json:"carriage_id"`      // 运输ID。
}

type ArrivalpassArrivalPassCreateRequestArrivalPass struct {
	VehicleModel        string `json:"vehicle_model"`         // 车辆型号。
	WarehouseID         int64  `json:"warehouse_id"`          // 卖家仓库ID。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	ArrivalTime         string `json:"arrival_time"`          // 入场时间，UTC格式。 此时通行证将开始生效。
	DriverName          string `json:"driver_name"`           // 司机姓名。
	DriverPhone         string `json:"driver_phone"`          // 司机电话号码。
	DropoffPointID      int64  `json:"dropoff_point_id"`      // 通行证适用的仓库ID。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
}

// 筛选器。
// ArrivalReason values
type ArrivalReason string

const (
	ArrivalReasonFBSDelivery ArrivalReason = "FBS_DELIVERY" // 发运
	ArrivalReasonFBSReturn   ArrivalReason = "FBS_RETURN"   // 运出退货
)

type ArrivalPassListRequestFilter struct {
	DropoffPointIds  []string      `json:"dropoff_point_ids"`  // 按发运点筛选。
	OnlyActivePasses bool          `json:"only_active_passes"` // `true`, 以获取仅活跃的通行证申请。
	WarehouseIds     []string      `json:"warehouse_ids"`      // 按卖家仓库筛选。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	ArrivalPassIds   []string      `json:"arrival_pass_ids"`   // 按通行证ID筛选。
	ArrivalReason    ArrivalReason `json:"arrival_reason"`     // 按入场目的筛选： - `FBS_DELIVERY` — 发运。 - `FBS_RETURN` — 运出退货。 如果未指定此参数，则考虑两种目的。 指定的原因必须在通行证的原因列表中。
}

type ArrivalpassArrivalPassListRequest struct {
	Cursor string                       `json:"cursor"` // 用于获取下一批数据的指针。
	Filter ArrivalPassListRequestFilter `json:"filter"`
	Limit  int32                        `json:"limit"` // 响应中记录数量的限制。
}

type ArrivalpassArrivalPassListResponseArrivalPass struct {
	VehicleModel        string   `json:"vehicle_model"`         // 车辆型号。
	DriverName          string   `json:"driver_name"`           // 司机的姓名。
	VehicleLicensePlate string   `json:"vehicle_license_plate"` // 车辆牌照号码。
	WarehouseID         int64    `json:"warehouse_id"`          // 卖家仓库ID。
	ArrivalPassID       int64    `json:"arrival_pass_id"`       // 通行证ID。
	ArrivalReasons      []string `json:"arrival_reasons"`       // 入场目的。
	ArrivalTime         string   `json:"arrival_time"`          // 入场日期和时间，UTC格式。
	DriverPhone         string   `json:"driver_phone"`          // 司机的电话号码。
	DropoffPointID      int64    `json:"dropoff_point_id"`      // 发运点ID。
	IsActive            bool     `json:"is_active"`             // 如果申请是活跃的，则为`true`。
}

type ArrivalpassArrivalPassListResponse struct {
	ArrivalPasses []ArrivalpassArrivalPassListResponseArrivalPass `json:"arrival_passes"` // 运输通行证列表。
	Cursor        string                                          `json:"cursor"`         // 用于获取下一批数据的指针。 如果此参数为空，则没有更多数据。
}

// 筛选器。
type V1ReturnsCompanyFbsInfoRequestFilter struct {
	PlaceID int64 `json:"place_id"` // 按揽收点ID筛选。
}

type ArrivalpassArrivalPassUpdateRequestArrivalPass struct {
	ArrivalPassID       int64  `json:"arrival_pass_id"`       // 通行证ID。
	ArrivalTime         string `json:"arrival_time"`          // 入场时间，UTC格式。 此时通行证将开始生效。 要更改入场时间，请使用方法 [/v1/carriage/pass/update](#operation/carriagePassUpdate)。
	DriverName          string `json:"driver_name"`           // 司机的姓名。
	DriverPhone         string `json:"driver_phone"`          // 司机的电话号码。
	VehicleLicensePlate string `json:"vehicle_license_plate"` // 车辆牌照号码。
	VehicleModel        string `json:"vehicle_model"`         // 车辆型号。
}

type SellerAPIArrivalPassUpdateRequest struct {
	ArrivalPasses []SellerAPIArrivalPassUpdateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
	CarriageID    int64                                          `json:"carriage_id"`    // 运输ID。
}

// 请求结果。
type SellerAPIArrivalPassCreateResponse struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证ID。
}

// 方法响应的分割。
type CompanyFbsInfoRequestPagination struct {
	LastID int64 `json:"last_id"` // 页面上最后一个揽收点的ID。对于第一个请求，请将此字段留空。 要获取后续的值，请指定上一个请求响应中最后一个揽收点的`id`。
	Limit  int32 `json:"limit"`   // 页面上揽收点的数量。最大值为500。
}

type ArrivalpassArrivalPassDeleteRequest struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证ID。
}

type V1ReturnsCompanyFbsInfoRequest struct {
	Filter     V1ReturnsCompanyFbsInfoRequestFilter `json:"filter"`
	Pagination CompanyFbsInfoRequestPagination      `json:"pagination"`
}

type V1ReturnsCompanyFbsInfoResponse struct {
	DropOffPoints []CompanyFbsInfoResponseDropOffPoints `json:"drop_off_points"` // 揽收点信息。
	HasNext       bool                                  `json:"has_next"`        // 是否还有其他揽收点等待卖家退货的标志。
}

type ArrivalpassArrivalPassUpdateRequest struct {
	ArrivalPasses []ArrivalpassArrivalPassUpdateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
}

type SellerAPIArrivalPassCreateRequest struct {
	ArrivalPasses []SellerAPIArrivalPassCreateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
	CarriageID    int64                                          `json:"carriage_id"`    // 运输ID。
}

type ArrivalpassArrivalPassCreateResponse struct {
	ArrivalPassIds []string `json:"arrival_pass_ids"` // 通行证ID。
}

type ArrivalpassArrivalPassCreateRequest struct {
	ArrivalPasses []ArrivalpassArrivalPassCreateRequestArrivalPass `json:"arrival_passes"` // 通行证列表。
}

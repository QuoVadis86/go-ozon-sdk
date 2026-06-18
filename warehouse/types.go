package warehouse

// 仓库工作时间。
type TimetableWorkingHours struct {
	TimeFrom string `json:"time_from"` // 工作开始时间。
	TimeTo   string `json:"time_to"`   // 工作结束时间。
}

// 仓库工作时间表。
type Timetable struct {
	TimetableFrom string                  `json:"timetable_from"` // 仓库工作开始日期。
	TimetableTo   string                  `json:"timetable_to"`   // 仓库工作结束日期。
	WorkingHours  []TimetableWorkingHours `json:"working_hours"`  // 仓库工作时间。
}

type WorkingDaysEnum string

// 仓库位置信息。
type AddressInfo struct {
	Address   string  `json:"address"`   // 仓库地址。
	Latitude  float64 `json:"latitude"`  // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
	Utc       string  `json:"utc"`       // 时区。
}

// 标签类型： - `UNSPECIFIED` — 未知类型； - `BIG` — 大标签； - `SMALL` — 小标签。
type CarriageLabelTypeEnum string

// 头程物流类型 — `PICK_UP` 或 `DROP_OFF`。
type FirstMileTypeEnum string

// 头程物流。
type FirstMile struct {
	TimeslotTo          string            `json:"timeslot_to"` // 时间段结束时间。
	Type                FirstMileTypeEnum `json:"type_"`
	DropoffPointID      string            `json:"dropoff_point_id"`       // 揽收点识别符。
	FirstMileIsChanging bool              `json:"first_mile_is_changing"` // 仓库设置更新的标识。
	TimeslotFrom        string            `json:"timeslot_from"`          // 时间段开始时间。
	TimeslotID          int64             `json:"timeslot_id"`            // 时间段识别符。
}

// WorkingDays values
type WorkingDays string

const (
	WorkingDaysUnspecified WorkingDays = "UNSPECIFIED" // 值未确定；
	WorkingDaysMonday      WorkingDays = "MONDAY"      // 星期一；
	WorkingDaysTuesday     WorkingDays = "TUESDAY"     // 星期二；
	WorkingDaysWednesday   WorkingDays = "WEDNESDAY"   // 星期三；
	WorkingDaysThursday    WorkingDays = "THURSDAY"    // 星期四；
	WorkingDaysFriday      WorkingDays = "FRIDAY"      // 星期五；
	WorkingDaysSaturday    WorkingDays = "SATURDAY"    // 星期六；
	WorkingDaysSunday      WorkingDays = "SUNDAY"      // 星期日
)

type ListV2ResponseWarehouse struct {
	HasPostingsLimit       bool                  `json:"has_postings_limit"` // 存在最低订单数量限额的标识：若有限额，则返回值为 `true`。
	IsComfort              bool                  `json:"is_comfort"`         // сomfort配送标志。送达买家时间大于等于60分钟。
	IsExpress              bool                  `json:"is_express"`         // express配送标志。送达买家时间不超过60分钟。
	PostingsLimit          int32                 `json:"postings_limit"`     // 订单限额。如果没有限额，则返回值为 `-1`。
	UpdatedAt              string                `json:"updated_at"`         // 仓库信息最后一次更新的日期和时间。
	WarehouseType          string                `json:"warehouse_type"`     // 仓库类型。
	WorkingDays            []WorkingDaysEnum     `json:"working_days"`       // 仓库工作日： - `UNSPECIFIED` — 值未确定； - `MONDAY` — 星期一； - `TUESDAY` — 星期二； - `WEDNESDAY` — 星期三； - `THURSDAY` — 星期四； - `FRIDAY` ...
	IsWaybillEnabled       bool                  `json:"is_waybill_enabled"` // 运单打印开启的标识。
	Phone                  string                `json:"phone"`              // 仓库电话号码。
	WithItemList           bool                  `json:"with_item_list"`     // 拣货单打印开启的标识。
	IsKgt                  bool                  `json:"is_kgt"`             // 仓库接收超大货物的标识。
	IsRFBS                 bool                  `json:"is_rfbs"`            // 仓库在 rFBS 工作模式下工作的标识。
	AddressInfo            AddressInfo           `json:"address_info"`
	CarriageLabelType      CarriageLabelTypeEnum `json:"carriage_label_type"`
	CutInTime              int64                 `json:"cut_in_time"`        // 发运所需时间（分钟）。
	IsAutoAssembly         bool                  `json:"is_auto_assembly"`   // 自动备货开启的标识。
	MinPostingsLimit       int32                 `json:"min_postings_limit"` // 一次交货中可以运来的最低订单数量。
	Name                   string                `json:"name"`               // 仓库名称。
	CourierPhones          []string              `json:"courier_phones"`     // 用于与快递员联系的电话号码。
	CreatedAt              string                `json:"created_at"`         // 仓库创建日期和时间。
	FirstMile              FirstMile             `json:"first_mile"`
	HasEntrustedAcceptance bool                  `json:"has_entrusted_acceptance"` // 信任验收开通标识。
	SlaCutIn               int64                 `json:"sla_cut_in"`               // 以分钟为单位的订单备货最低时间。
	Status                 string                `json:"status"`                   // 仓库状态。
	Timetable              Timetable             `json:"timetable"`
	WarehouseID            int64                 `json:"warehouse_id"`    // 仓库识别符。
	CourierComment         string                `json:"courier_comment"` // 给快递员的评论。
}

type V2WarehouseListV2Response struct {
	Warehouses []ListV2ResponseWarehouse `json:"warehouses"` // 仓库列表。
	HasNext    bool                      `json:"has_next"`   // `true`，前提是本次响应未返回所有数据。
	Cursor     string                    `json:"cursor"`     // 后续数据的选择标志。
}

// 需要添加到响应中的附加字段。
type V1WarehouseListRequestWith struct {
	AbleToSetPrice bool `json:"able_to_set_price"` // `true`，用于在响应中添加关于价格设置能力的信息。
}

type V1WarehouseListRequest struct {
	Limit  int64                      `json:"limit"`  // 响应中返回的值数量。
	Offset int64                      `json:"offset"` // 在响应中将被跳过的项目数量。例如，如果 `offset = 10`，则响应将从第 11 个找到的项目开始。
	With   V1WarehouseListRequestWith `json:"with"`
}

// 商品尺寸信息。
type ItemSize struct {
	HeightMm int32 `json:"height_mm"` // 商品高度（毫米）。
	LengthMm int32 `json:"length_mm"` // 商品长度（毫米）。
	WidthMm  int32 `json:"width_mm"`  // 商品宽度（毫米）。
}

// 坐标。
type V1ListDropOffPointsForCreateFBSWarehouseRequestCoordinates struct {
	Longitude float64 `json:"longitude"` // 经度。
	Latitude  float64 `json:"latitude"`  // 纬度。
}

type ListDropOffPointsForCreateFBSWarehouseRequestSearchDropOffPointTypeEnum string

// 搜索参数。
// Types values
type Types string

const (
	TypesPVZ Types = "PVZ" // 订单取货点
	TypesPPZ Types = "PPZ" // 订单接收点
	TypesSC  Types = "SC"  // 分拣中心
)

type ListDropOffPointsForCreateFBSWarehouseRequestSearch struct {
	Address string                                                                    `json:"address"` // 揽收点地址。
	Types   []ListDropOffPointsForCreateFBSWarehouseRequestSearchDropOffPointTypeEnum `json:"types"`   // 揽收点类型： - `PVZ` — 订单取货点， - `PPZ` — 订单接收点， - `SC` — 分拣中心。
}

type V1ListDropOffPointsForCreateFBSWarehouseRequest struct {
	Coordinates V1ListDropOffPointsForCreateFBSWarehouseRequestCoordinates `json:"coordinates"`
	CountryCode string                                                     `json:"country_code"` // ISO 2格式的国家代码。
	IsKgt       bool                                                       `json:"is_kgt"`       // `true`表示商品为超大。
	Search      ListDropOffPointsForCreateFBSWarehouseRequestSearch        `json:"search"`
}

type ListDropOffPointsForUpdateFBSWarehouseRequestSearchDropOffPointTypeEnum string

// 搜索参数。
type ListDropOffPointsForUpdateFBSWarehouseRequestSearch struct {
	Address string                                                                    `json:"address"` // 根据揽收点地址搜索。
	Types   []ListDropOffPointsForUpdateFBSWarehouseRequestSearchDropOffPointTypeEnum `json:"types"`   // 揽收点类型： - `PVZ` — 订单取货点， - `PPZ` — 订单接收点， - `SC` — 分拣中心。
}

type V1ListDropOffPointsForUpdateFBSWarehouseRequest struct {
	Search      ListDropOffPointsForUpdateFBSWarehouseRequestSearch `json:"search"`
	WarehouseID int64                                               `json:"warehouse_id"` // 现有FBS仓库的筛选条件。
}

// 操作状态： - `UNSPECIFIED` — 未定义； - `IN_PROGRESS` — 进行中； - `SUCCESS` — 已完成； - `ERROR` — 出错结束。
type GetWarehouseFBSOperationStatusResponseStatusEnum string

// 商品信息。
type ValidationResultItem struct {
	Size    ItemSize `json:"size"`
	SKU     int64    `json:"sku"`      // 商品在Ozon系统中的标识符——SKU。
	WeightG float64  `json:"weight_g"` // 商品重量（克）。
}

type CreateWarehouseFBSRequestWorkingDaysEnum string

// 第一英里 FBS。
type FirstMileType struct {
	DropoffPointID      string `json:"dropoff_point_id"`       // DropOff-点识别号。
	DropoffTimeslotID   int64  `json:"dropoff_timeslot_id"`    // DropOff临时存档识别号。
	FirstMileIsChanging bool   `json:"first_mile_is_changing"` // 正在更新仓库设置的迹象。
	FirstMileType       string `json:"first_mile_type"`        // 第一英里的类型 — `DropOff` 或 `Pickup`。
}

type ListResponseWarehouse struct {
	IsAbleToSetPrice       bool          `json:"is_able_to_set_price"`     // 如果可以设置价格，则为`true`。
	MinPostingsLimit       int32         `json:"min_postings_limit"`       // 限制的最小值是指在一次供货中可以带来的订单数量。
	PostingsLimit          int32         `json:"postings_limit"`           // 极限值。 `-1`, 如果没有限制
	HasEntrustedAcceptance bool          `json:"has_entrusted_acceptance"` // 会表明受信任的接受。`true`, 如果库存中启用了受信任的接受方式。
	FirstMileType          FirstMileType `json:"first_mile_type"`
	IsKgt                  bool          `json:"is_kgt"`                   // 该迹象表明仓库接受大宗商品。
	IsTimetableEditable    bool          `json:"is_timetable_editable"`    // 该迹象表明可以改变仓库运行时间表。
	Name                   string        `json:"name"`                     // 仓库名称。
	CanPrintActInAdvance   bool          `json:"can_print_act_in_advance"` // 有可能提前打印收发证书。 `true`, 如果可以提前打印的话。
	IsKarantin             bool          `json:"is_karantin"`              // 该迹象表明仓库因隔离而停止运作。
	MinWorkingDays         int64         `json:"min_working_days"`         // 仓库运行天数。
	WarehouseID            int64         `json:"warehouse_id"`             // 仓库识别号。
	HasPostingsLimit       bool          `json:"has_postings_limit"`       // 该迹象表明对最小订单数有限制。`true`, 如果有限制。
	IsPresorted            bool          `json:"is_presorted"`             // 如果发运是预先分拣的，则为`true`。
	Status                 string        `json:"status"`                   // 仓库状况。 仓库状态与个人账户中的状态的对应关系: | 状态 Seller&nbsp;API | 个人账户中的状态 | |---|---| | `new` | 正在激活 | | `created` | 已激活 | | `disabled` ...
	WorkingDays            []string      `json:"working_days"`             // 仓库运行天数。
	IsRFBS                 bool          `json:"is_rfbs"`                  // 在 rFBS 计划下运作的仓库的标志： - `true` — 仓库在 rFBS 计划下运行； - `false` — 仓库没有在 rFBS 计划下运行。
}

type WarehouseListResponse struct {
	Result []ListResponseWarehouse `json:"result"` // 仓库清单。
}

// 揽收点类型： - `PVZ` — 订单取货点， - `PPZ` — 订单接收点， - `SC` — 分拣中心。
type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPointDropOffPointTypeEnum string

type V1WarehouseFbsCreateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceEndTimeLocal   string `json:"acceptance_end_time_local"`   // 订单接收结束本地时间。
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"` // 订单接收开始本地时间。
	From                     string `json:"from"`                        // 时间段开始时间。
	ID                       int64  `json:"id"`                          // 时间段标识符。
	To                       string `json:"to"`                          // 时间段结束时间。
}

// 查找快递方式的过滤器。
// Status values
type Status string

const (
	StatusNEW      Status = "NEW"      // 已创建
	StatusEdited   Status = "EDITED"   // 正在编辑
	StatusActive   Status = "ACTIVE"   // 已激活
	StatusDisabled Status = "DISABLED" // 未激活
)

type MethodListRequestFilter struct {
	WarehouseID int64  `json:"warehouse_id"` // 仓库识别号。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	ProviderID  int64  `json:"provider_id"`  // 快递服务识别号。
	Status      Status `json:"status"`       // 快递方式状态: - `NEW` — 已创建, - `EDITED` — 正在编辑, - `ACTIVE` — 已激活, - `DISABLED` — 未激活。
}

type DeliveryMethodListRequest struct {
	Filter MethodListRequestFilter `json:"filter"`
	Limit  int64                   `json:"limit"`  // 回答中的元素数量。最多50，最少1。
	Offset int64                   `json:"offset"` // 回答中会被略过的元素数量。例如，如果`offset = 10`，回答将从发现的第11个元素开始。
}

// 揽收点坐标。
type V1ListDropOffPointsForCreateFBSWarehouseResponseCoordinates struct {
	Latitude  float64 `json:"latitude"`  // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
}

type V1ArchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 操作识别符。请使用 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 方式获取操作状态。
}

type V1UpdateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 操作ID。通过方法 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 获取操作状态。
}

// 处理操作时出错。
type GetWarehouseFBSOperationStatusResponseError struct {
	Code    string `json:"code"`    // 错误代码。
	Message string `json:"message"` // 错误描述。
}

// 操作结果。
type GetWarehouseFBSOperationStatusResponseResult struct {
	EntityID int64 `json:"entity_id"` // 正在处理的实体ID。如果操作为 `CREATE_FBS_WAREHOUSE`，则返回仓库ID。
}

// D揽收点类型： - `PVZ` — 订单取货点， - `PPZ` — 订单接收点， - `SC` — 分拣中心。
type V1ListDropOffPointsForCreateFBSWarehouseResponseDropOffPointTypeEnum string

type V1WarehouseFbsCreatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"` // 时间段开始时间。
	ID   int64  `json:"id"`   // 时间段标识符。
	To   string `json:"to"`   // 时间段结束时间。
}

// 错误类型： - `UNSPECIFIED`——未指定； - `LESS_THAN_MIN`——小于最小值； - `GREATER_THAN_MAX`——大于最大值。
type ValidationResultValidationErrorTypeEnum string

// 揽收点坐标。
type V2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPointAddressCoordinates struct {
	Latitude  float64 `json:"latitude"`  // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
}

type V1WarehouseFbsCreatePickUpTimeslotListResponse struct {
	IsPickupSupported bool                                                     `json:"is_pickup_supported"` // 支持pick-up发运标志。
	Timeslots         []V1WarehouseFbsCreatePickUpTimeslotListResponseTimeslot `json:"timeslots"`           // 时间段列表。
}

type V1UpdateWarehouseFBSFirstMileResponse struct {
	OperationID string `json:"operation_id"` // 操作ID。通过方法 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 获取操作状态。
}

// 未通过检查的商品特征。 可能的限制类型： - `UNSPECIFIED`——未指定； - `LENGTH`——长度； - `WIDTH`——宽度； - `HEIGHT`——高度； - `WEIGHT`——重量； - `SUM_OF_DIME...
type ValidationErrorCharacteristicEnum string

// 价格限制。
type V1Money struct {
	Currency string  `json:"currency"` // 货币单位。
	Value    float64 `json:"value"`    // 价格数值。
}

type ValidationResultValidationError struct {
	TemplateID       int32                                   `json:"template_id"` // 设置配送限制的订单配送服务 ID。
	Type             ValidationResultValidationErrorTypeEnum `json:"type_"`
	Characteristic   ValidationErrorCharacteristicEnum       `json:"characteristic"`
	RestrictionPrice V1Money                                 `json:"restriction_price"`
	RestrictionVwc   float64                                 `json:"restriction_vwc"` // 体积和重量特征限制值——体积和重量特征。
}

type FilterStatusEnum string

// 用于搜索配送方式的筛选条件。
type MethodListV2RequestFilter struct {
	DeliveryMethodIds []string           `json:"delivery_method_ids"` // 配送方式标识符。
	ProviderIds       []string           `json:"provider_ids"`        // 配送服务标识符。
	Status            []FilterStatusEnum `json:"status"`              // 配送方式状态： - `NEW`：已创建， - `EDITED`：编辑中， - `ACTIVE`：已启用， - `DISABLED`：未启用， - `WAITING`：审核中， - `BROKEN`：异常。
	WarehouseIds      []string           `json:"warehouse_ids"`       // 仓库标识符。可通过方法[/v2/warehouse/list](#operation/WarehouseListV2)获取。
}

// 排序方向： - `ASC`：升序， - `DESC`：降序。
type MethodListV2RequestSortDirEnum string

type V2DeliveryMethodListV2Request struct {
	Cursor  string                         `json:"cursor"` // 用于选择下一批数据的指针。
	Filter  MethodListV2RequestFilter      `json:"filter"`
	Limit   int64                          `json:"limit"` // 响应中返回的值数量。
	SortDir MethodListV2RequestSortDirEnum `json:"sort_dir"`
}

type V1WarehouseFbsCreateDropOffTimeslotListResponse struct {
	Timeslots []V1WarehouseFbsCreateDropOffTimeslotListResponseTimeslot `json:"timeslots"` // 时间段列表。
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponseTimeslot struct {
	From string `json:"from"` // 时间段开始时间。
	ID   int64  `json:"id"`   // 时间段标识符。
	To   string `json:"to"`   // 时间段结束时间。
}

type V1CreateWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 创建FBS仓库的操作标识符。要获取操作状态，请使用方法 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus)。
}

type V1UnarchiveWarehouseFBSResponse struct {
	OperationID string `json:"operation_id"` // 操作识别符。请使用 [/v1/warehouse/operation/status](#operation/GetWarehouseFBSOperationStatus) 方式获取操作状态。
}

// 揽收点信息。
type V2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPoint struct {
	Address            string                                                                        `json:"address"` // 揽收点地址。
	AddressCoordinates V2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPointAddressCoordinates `json:"address_coordinates"`
	Code               string                                                                        `json:"code"` // 运输公司系统中的揽收点代码。
	Name               string                                                                        `json:"name"` // 揽收点名称。
}

// TPLIntegrationType values
type TPLIntegrationType string

const (
	TPLIntegrationTypeAggregator    TPLIntegrationType = "aggregator"     // ：由第三方配送，Ozon负责登记订单；
	TPLIntegrationTypeV3plTracking  TPLIntegrationType = "3pl_tracking"   // ：由第三方配送，卖家负责登记订单；
	TPLIntegrationTypeNonIntegrated TPLIntegrationType = "non_integrated" // ：由卖家自行配送；
	TPLIntegrationTypeHybrid        TPLIntegrationType = "hybrid"         // ：混合集成
)

type MethodListV2ResponseDeliveryMethod struct {
	CreatedAt          string                                                      `json:"created_at"`  // 配送方式创建日期。
	Cutoff             string                                                      `json:"cutoff"`      // 卖家需在此时间前完成订单备货。
	ID                 int64                                                       `json:"id"`          // 配送方式标识符。
	IsExpress          bool                                                        `json:"is_express"`  // `true`，表示可使用Ozon Express快速配送。
	ProviderID         int64                                                       `json:"provider_id"` // 配送服务标识符。
	Status             Status                                                      `json:"status"`      // 配送方式状态： - `NEW`：已创建， - `EDITED`：编辑中， - `ACTIVE`：已启用， - `DISABLED`：未启用， - `WAITING`：审核中， - `BROKEN`：异常。
	TPLDropoffPoint    V2DeliveryMethodListV2ResponseDeliveryMethodTPLDropOffPoint `json:"tpl_dropoff_point"`
	WarehouseID        int64                                                       `json:"warehouse_id"`         // 仓库标识符。
	Name               string                                                      `json:"name"`                 // 配送方式名称。
	SlaCutIn           int64                                                       `json:"sla_cut_in"`           // 根据仓库设置，订单备货的最短时间（以分钟为单位）。
	TemplateID         int64                                                       `json:"template_id"`          // 订单配送服务标识符。
	TPLIntegrationType TPLIntegrationType                                          `json:"tpl_integration_type"` // 与配送服务的集成类型： - `aggregator`：由第三方配送，Ozon负责登记订单； - `3pl_tracking`：由第三方配送，卖家负责登记订单； - `non_integrated`：由卖家自行配送； - `hybrid`：混...
	UpdatedAt          string                                                      `json:"updated_at"`           // 配送方式最后更新日期和时间。
}

// 头程物流类型： - `PICK_UP` — 订单交给快递员； - `DROP_OFF` — 订单送到接收点。
type CreateWarehouseFBSRequestFirstMileTypeEnum string

// 仓库地址坐标。
type CreateWarehouseFBSRequestAddressCoordinates struct {
	Latitude  float64 `json:"latitude"`  // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
}

// 仓库参数。
type CreateWarehouseFBSRequestOptions struct {
	IsAutoAssembly   bool     `json:"is_auto_assembly"`   // `true`表示启用自动备货。
	IsWaybillEnabled bool     `json:"is_waybill_enabled"` // `true`表示启用运单打印。
	Comment          string   `json:"comment"`            // 交货类型为 `PICK_UP` 时给快递员的备注。
	CourierPhones    []string `json:"courier_phones"`     // 交货类型为 `PICK_UP` 时给快递员的电话号码。请按格式填写：+7 (XXX) XXX-XX-XX。
}

type V1CreateWarehouseFBSRequest struct {
	CutInTime          int64                                       `json:"cut_in_time"` // 接单时间（分钟）。例如，如果传入 `3000`，则接单将在传入后50小时内结束。
	FirstMileType      CreateWarehouseFBSRequestFirstMileTypeEnum  `json:"first_mile_type"`
	IsKgt              bool                                        `json:"is_kgt"`      // `true`表示商品为超大。
	Name               string                                      `json:"name"`        // 仓库名称。
	Phone              string                                      `json:"phone"`       // 仓库联系电话。请按格式填写：+7 (XXX) XXX-XX-XX。
	TimeslotID         int64                                       `json:"timeslot_id"` // 时间段标识符。
	AddressCoordinates CreateWarehouseFBSRequestAddressCoordinates `json:"address_coordinates"`
	DropOffPointID     int64                                       `json:"drop_off_point_id"` // 揽收点标识符。
	Options            CreateWarehouseFBSRequestOptions            `json:"options"`
	WorkingDays        []CreateWarehouseFBSRequestWorkingDaysEnum  `json:"working_days"` // 仓库工作日： - `MONDAY` — 星期一， - `TUESDAY` — 星期二， - `WEDNESDAY` — 星期三， - `THURSDAY` — 星期四， - `FRIDAY` — 星期五， - `SATURDAY` — 星期...
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponseTimeslot struct {
	AcceptanceStartTimeLocal string `json:"acceptance_start_time_local"` // Sipariş kabul işleminin başladığı yerel saat.
	From                     string `json:"from"`                        // Zaman aralığı başlangıç saati.
	ID                       int64  `json:"id"`                          // Zaman aralığı kimliği.
	To                       string `json:"to"`                          // Zaman aralığı bitiş saati.
	AcceptanceEndTimeLocal   string `json:"acceptance_end_time_local"`   // Sipariş kabul işleminin sona erdiği yerel saat.
}

// 为获得发运折扣，需要在此时间前交付货件。
type TypeTimeOfDay struct {
	Hours   int32 `json:"hours"`   // 小时。
	Minutes int32 `json:"minutes"` // 分钟。
	Nanos   int32 `json:"nanos"`   // 纳秒。
	Seconds int32 `json:"seconds"` // 秒。
}

type ListDropOffPointsForCreateFBSWarehouseResponseDropOffPoint struct {
	Address              string                                                               `json:"address"` // 揽收点地址。
	Coordinates          V1ListDropOffPointsForCreateFBSWarehouseResponseCoordinates          `json:"coordinates"`
	DiscountPercent      float64                                                              `json:"discount_percent"` // 交付货件的折扣百分比。
	ID                   string                                                               `json:"id"`               // 揽收点标识符。
	LastTransitTimeLocal TypeTimeOfDay                                                        `json:"last_transit_time_local"`
	Type                 V1ListDropOffPointsForCreateFBSWarehouseResponseDropOffPointTypeEnum `json:"type_"`
}

type V1ListDropOffPointsForCreateFBSWarehouseResponse struct {
	Points []ListDropOffPointsForCreateFBSWarehouseResponseDropOffPoint `json:"points"` // 点位列表。
}

// 揽收点坐标。
type DropOffPointCoordinates struct {
	Latitude  float64 `json:"latitude"`  // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
}

type ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPoint struct {
	LastTransitTimeLocal TypeTimeOfDay                                                                  `json:"last_transit_time_local"`
	Type                 ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPointDropOffPointTypeEnum `json:"type_"`
	Address              string                                                                         `json:"address"` // 揽收点地址。
	Coordinates          DropOffPointCoordinates                                                        `json:"coordinates"`
	DiscountPercent      float64                                                                        `json:"discount_percent"` // 交付货件的折扣百分比。
	ID                   string                                                                         `json:"id"`               // 揽收点标识符。
}

type V1ListDropOffPointsForUpdateFBSWarehouseResponse struct {
	Points []ListDropOffPointsForUpdateFBSWarehouseResponseDropOffPoint `json:"points"` // 点位列表。
}

type V2DeliveryMethodListV2Response struct {
	Cursor          string                               `json:"cursor"`           // 用于选择下一批数据的指针。
	HasNext         bool                                 `json:"has_next"`         // `true`，表示响应中未返回全部配送方式。
	DeliveryMethods []MethodListV2ResponseDeliveryMethod `json:"delivery_methods"` // 配送方式。
}

type V1WarehouseFbsUpdateDropOffTimeslotListRequest struct {
	WarehouseID    int64 `json:"warehouse_id"`      // 仓库标识符。
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
}

type V1WarehouseFbsUpdateDropOffTimeslotListResponse struct {
	Timeslots []V1WarehouseFbsUpdateDropOffTimeslotListResponseTimeslot `json:"timeslots"` // 时间段列表。
}

// 操作类型： - `UNSPECIFIED` — 未定义； - `CREATE_FBS_WAREHOUSE` — 创建FBS仓库； - `UPDATE_FBS_WAREHOUSE` — 更新FBS仓库； - `SET_FIRST_MILE` ...
type GetWarehouseFBSOperationStatusResponseTypeEnum string

type V1GetWarehouseFBSOperationStatusResponse struct {
	Error  GetWarehouseFBSOperationStatusResponseError      `json:"error"`
	Result GetWarehouseFBSOperationStatusResponseResult     `json:"result"`
	Status GetWarehouseFBSOperationStatusResponseStatusEnum `json:"status"`
	Type   GetWarehouseFBSOperationStatusResponseTypeEnum   `json:"type_"`
}

// 检查状态： - `UNSPECIFIED`——未指定； - `NOT_VALID`——商品未通过检查。
type ValidationResultItemStateEnum string

type InvalidProductsGetResponseValidationResult struct {
	Item             ValidationResultItem              `json:"item"`
	State            ValidationResultItemStateEnum     `json:"state"`
	ValidationErrors []ValidationResultValidationError `json:"validation_errors"` // 错误信息。
}

type V1WarehouseInvalidProductsGetResponse struct {
	HasNext           bool                                         `json:"has_next"`           // 若响应中未包含全部商品，则为`true`。
	LastID            int64                                        `json:"last_id"`            // 页面上最后一个值的标识符。要获取下一个批次的数据，请在下一个请求的 `last_id` 参数中传递上次获取的值。
	ValidationResults []InvalidProductsGetResponseValidationResult `json:"validation_results"` // 检查结果。
	WarehouseID       int64                                        `json:"warehouse_id"`       // 仓库标识符。
}

type V1ArchiveWarehouseFBSRequest struct {
	Reason      string `json:"reason"`       // 仓库归档原因。
	WarehouseID int64  `json:"warehouse_id"` // 仓库识别符。
}

type V1UpdateWarehouseFBSRequestWorkingDaysEnum string

type MethodListResponseDeliveryMethod struct {
	CompanyID   int64  `json:"company_id"`   // 卖家识别号。
	CreatedAt   string `json:"created_at"`   // 创建快递方式的日期和时间。
	Cutoff      string `json:"cutoff"`       // 卖方必须在此之前备货的时间。
	ID          int64  `json:"id"`           // 快递方式识别号。
	ProviderID  int64  `json:"provider_id"`  // 快递服务识别号。
	SlaCutIn    int64  `json:"sla_cut_in"`   // 根据仓库设置，订单备货的最短时间（以分钟为单位）。
	Status      Status `json:"status"`       // 快递方式状态: - `NEW` — 已创建, - `EDITED` — 正在编辑, - `ACTIVE` — 已激活, - `DISABLED` — 未激活。
	TemplateID  int64  `json:"template_id"`  // 订单快递服务识别号。
	Name        string `json:"name"`         // 快递方式名称。
	UpdatedAt   string `json:"updated_at"`   // 快递方式最后更新的日期和时间。
	WarehouseID int64  `json:"warehouse_id"` // 仓库识别号。
}

type DeliveryMethodListResponse struct {
	HasNext bool                               `json:"has_next"` // 以下该迹象会表明在查询中只送回了部分快递方式。 - `true` — 请用新的 `offset` 参数重新请求，以获得剩余的方式； - `false` — 回答中包含了所有应要求的快递方式。
	Result  []MethodListResponseDeliveryMethod `json:"result"`   // 查询结果。
}

type V1GetWarehouseFBSOperationStatusRequest struct {
	OperationID string `json:"operation_id"` // 操作ID。
}

type V1WarehouseFbsUpdatePickUpTimeslotListResponse struct {
	Timeslots []V1WarehouseFbsUpdatePickUpTimeslotListResponseTimeslot `json:"timeslots"` // 时间段列表。
}

// 第一英里类型： - `PICK_UP` —发运给快递员； - `DROP_OFF` —发运至接收点。
type V1UpdateWarehouseFBSFirstMileRequestFirstMileTypeEnum string

// 仓库坐标。
type V1WarehouseFbsCreatePickUpTimeslotListRequestAddressCoordinates struct {
	Latitude  float64 `json:"latitude"`  // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
}

type V1WarehouseFbsCreatePickUpTimeslotListRequest struct {
	AddressCoordinates V1WarehouseFbsCreatePickUpTimeslotListRequestAddressCoordinates `json:"address_coordinates"`
	IsKgt              bool                                                            `json:"is_kgt"` // 超大货物标志。
}

// 仓库坐标。
type V1UpdateWarehouseFBSRequestAddressCoordinates struct {
	Latitude  float64 `json:"latitude"`  // 纬度。
	Longitude float64 `json:"longitude"` // 经度。
}

type V1WarehouseFbsCreateDropOffTimeslotListRequest struct {
	DropOffPointID int64 `json:"drop_off_point_id"` // 揽收点标识符。
}

type V1WarehouseInvalidProductsGetRequest struct {
	LastID      int64 `json:"last_id"`      // 页面上最后一个值的标识符。首次请求时请将此字段留空。 如需获取后续数据，请填写上次响应中的 `last_id`。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。请通过方法[/v1/warehouse/warehouses-with-invalid-products](#operation/WarehouseWithInvalidProducts)获取该参数值。
}

type V1UpdateWarehouseFBSFirstMileRequest struct {
	CutInTime      int64                                                 `json:"cut_in_time"`       // 接单时间（分钟）。例如，如果传入 `3000`，则接单将在传入后50小时内结束。
	DropOffPointID int64                                                 `json:"drop_off_point_id"` // 订单发运点ID。如果` first_mile_type = DROP_OFF`，该参数必填。
	FirstMileType  V1UpdateWarehouseFBSFirstMileRequestFirstMileTypeEnum `json:"first_mile_type"`
	TimeslotID     int64                                                 `json:"timeslot_id"`  // 时间段标识符。
	WarehouseID    int64                                                 `json:"warehouse_id"` // 仓库ID。
}

// 仓库参数。
type V1UpdateWarehouseFBSRequestOptions struct {
	Comment          string   `json:"comment"`            // 发运类型为 `PICK_UP` 时，给快递员的备注。
	CourierPhones    []string `json:"courier_phones"`     // 发运类型为 `PICK_UP` 时，快递员联系电话。
	IsAutoAssembly   bool     `json:"is_auto_assembly"`   // 启用自动备货标识。
	IsWaybillEnabled bool     `json:"is_waybill_enabled"` // 启用运单打印标识。
}

type V1UpdateWarehouseFBSRequest struct {
	AddressCoordinates V1UpdateWarehouseFBSRequestAddressCoordinates `json:"address_coordinates"`
	Name               string                                        `json:"name"` // 仓库名称。
	Options            V1UpdateWarehouseFBSRequestOptions            `json:"options"`
	Phone              string                                        `json:"phone"`        // 仓库电话号码。
	WarehouseID        int64                                         `json:"warehouse_id"` // 仓库ID。
	WorkingDays        []V1UpdateWarehouseFBSRequestWorkingDaysEnum  `json:"working_days"` // 仓库工作日： - `MONDAY` — 周一； - `TUESDAY` — 周二； - `WEDNESDAY` — 周三； - `THURSDAY` — 周四； - `FRIDAY` — 周五； - `SATURDAY` — 周六； - `...
}

type V1WarehouseFbsUpdatePickUpTimeslotListRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
}

type V1UnarchiveWarehouseFBSRequest struct {
	WarehouseID int64 `json:"warehouse_id"` // 仓库识别符。
}

type V2WarehouseListV2Request struct {
	WarehouseIds []string `json:"warehouse_ids"` // 仓库识别符。
	Limit        int64    `json:"limit"`         // 响应中返回的值数量。
	Cursor       string   `json:"cursor"`        // 后续数据的选择标志。
}

type V1WarehouseWithInvalidProductsResponse struct {
	WarehouseIds []string `json:"warehouse_ids"` // 包含至少 1 件无法从该仓库发运的受限商品的仓库标识符列表。如需获取受限商品列表，请使用方法 [/v1/warehouse/invalid-products/get](#operation/WarehouseInvalidProductsG...
}

package fbs

type V1AssemblyCarriagePostingListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 每页显示的数量。
}

type V3GetFbsPostingResponseV3 struct {
	Result interface{} `json:"result"`
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []interface{} `json:"result"` // 货件信息。
}

type Postingv1GetCarriageAvailableListResponse struct {
	Result interface{} `json:"result"` // 方法操作结果。
}

type Postingv3GetFbsPostingRequest struct {
	PostingNumber string `json:"posting_number"` // 货件ID。
	With interface{} `json:"with"`
}

type PostingPostingProductCancelRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"` // 货物取消发货原因ID。
	CancelReasonMessage string `json:"cancel_reason_message"` // 必填字段。关于取消的其他信息。
	Items []interface{} `json:"items"` // 商品信息。
	PostingNumber string `json:"posting_number"` // 货运ID。
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"` // 按行过滤。
}

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"` // 备货后生成的货件号码。
}

type PostingV4PostingFbsUnfulfilledListResponse struct {
	Count int64 `json:"count"` // 在响应中的元素计数器。
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	HasNext bool `json:"has_next"` // 若响应中未返回全部货件，则为`true`。
	Postings interface{} `json:"postings"` // 货件列表。
}

type V1CarriageApproveRequest struct {
	CarriageID int64 `json:"carriage_id"` // 发运标识符。
	ContainersCount int32 `json:"containers_count"` // 货位数量。 如果您已开通信任验收，并按货位发运订单，请使用该参数。如果您未开通信任验收，请跳过该参数。
}

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"` // `true`表示已设置新日期。
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []interface{} `json:"result"` // 方法操作结果。
}

type V1AssemblyFbsProductListRequest struct {
	Offset int64 `json:"offset"` // 在响应中将被跳过的项目数量。例如，如果 `offset = 10`，则响应将从第 11 个找到的项目开始。
	SortDir interface{} `json:"sort_dir"`
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 每页显示的数量。
}

type V2FbsPostingProductCountryListResponse struct {
	Result []interface{} `json:"result"` // 制造国和ISO代码列表。
}

type PostingCancelReasonResponse struct {
	Result []interface{} `json:"result"` // 请求结果。
}

type V1AssemblyCarriagePostingListResponse struct {
	CanPrintMassLabel bool `json:"can_print_mass_label"` // `true`，前提是可以批量打印标签。
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Postings []interface{} `json:"postings"` // 货件列表。
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []interface{} `json:"tracking_numbers"` // 具有成对货运ID的数据 - 追踪号。
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []interface{} `json:"posting_number"` // 货件ID。
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"` // 新发运日期。
	PostingNumber string `json:"posting_number"` // 货件编号。
}

type V1CarriageCreateRequest struct {
	DeliveryMethodID int64 `json:"delivery_method_id"` // 配送方式标识符。
	DepartureDate string `json:"departure_date"` // 发运日期。默认值为当前日期。
	AllBlrTraceable bool `json:"all_blr_traceable"` // `true`，表示需要创建包含可追溯商品的发运。
}

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []interface{} `json:"posting_number"` // 货运ID。一次请求中的最大数量——100。
}

type PostingCancelReasonListResponse struct {
	Result []interface{} `json:"result"` // 方法操作结果。
}

type V1AssemblyFbsPostingListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 每页显示的数量。
	SortDir interface{} `json:"sort_dir"`
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []interface{} `json:"posting_number"` // 货运ID。
}

type Fbsv4FbsPostingShipV4Response struct {
	AdditionalData interface{} `json:"additional_data"` // 与发货有关的附加信息。
	Result interface{} `json:"result"` // 货运装配结果。
}

type Postingv1GetCarriageAvailableListRequest struct {
	DeliveryMethodID int64 `json:"delivery_method_id"` // 按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	DepartureDate string `json:"departure_date"` // 装运日期。默认 —— 当前日期。
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result interface{} `json:"result"`
}

type PostingCancelFbsPostingRequest struct {
	CancelReasonID int64 `json:"cancel_reason_id"` // 取消运输的原因ID。
	CancelReasonMessage string `json:"cancel_reason_message"` // 关于取消的附加信息。如果`cancel_reason_id = 402`，参数是必须的。
	PostingNumber string `json:"posting_number"` // 货件ID。
}

type V1AssemblyCarriageProductListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Products []interface{} `json:"products"` // 商品列表。
}

type Postingv3GetFbsPostingListRequest struct {
	Dir string `json:"dir"` // 分类方向： - `asc` — 从小到大， - `desc` — 从大到小。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 响应中值的数量： - 最大值 — 1000, - 最小值 — 1。
	Offset int64 `json:"offset"` // 将在响应中跳过的元素数。 例如，如果“offset=10”，那么响应将从找到的第11个元素开始。
	With interface{} `json:"with"`
}

type V2PostingFBSActGetPostingsRequest struct {
	ID interface{} `json:"id"` // 单据标识符。请通过方法[/v1/carriage/create](#operation/CarriageAPI_CarriageCreate)获取参数值。
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"` // 发运识别符。
}

type V1AssemblyFbsPostingListResponse struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Cutoff string `json:"cutoff"` // 卖家需在此时间前完成订单备货。
	Postings []interface{} `json:"postings"` // 货件列表。
}

type CarriageCarriageGetResponse struct {
	DeliveryMethodID int64 `json:"delivery_method_id"` // 物流方式标识符。
	FirstMileType string `json:"first_mile_type"` // 头程物流类型。
	Status string `json:"status"` // 运输状态。
	ArrivalPassIds []interface{} `json:"arrival_pass_ids"` // 为运输生成的通行证标识符列表。
	AvailableActions []interface{} `json:"available_actions"` // 运输的可用操作： - `get_shipping_list`——获取发运清单； - `get_act_of_acceptance`——获取验收证明书； - `get_waybill`——获取 PDF 格式的货单； - `set_arriva...
	DepartureDate string `json:"departure_date"` // 运输完成日期。
	RetryCount int32 `json:"retry_count"` // 运输创建重复尝试数量。
	CompanyID int64 `json:"company_id"` // 卖家标识符。
	PartialNum int64 `json:"partial_num"` // 部分运输序列号。
	UpdatedAt string `json:"updated_at"` // 运输信息最后一次更新日期。
	ActType string `json:"act_type"` // 交接单类型。针对FBS卖家。
	CreatedAt string `json:"created_at"` // 运输创建日期。
	HasPostingsForNextCarriage bool `json:"has_postings_for_next_carriage"` // `true`, 如果有未能进行运输，但需要发运的货件。
	IntegrationType string `json:"integration_type"` // 运输类型。
	IsContainerLabelPrinted bool `json:"is_container_label_printed"` // `true`, 如果您已经打印了货位标签。
	IsPartial bool `json:"is_partial"` // `true`, 如果是部分运输。
	TPLProviderID int64 `json:"tpl_provider_id"` // 配送服务商标识符。
	WarehouseID int64 `json:"warehouse_id"` // 仓库标识符。
	CancelAvailability interface{} `json:"cancel_availability"`
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
	ContainersCount int32 `json:"containers_count"` // 货位数量。
}

type V3GetFbsPostingListResponseV3 struct {
	Result interface{} `json:"result"`
}

type V1SetPostingsRequest struct {
	CarriageID int64 `json:"carriage_id"` // 发运识别符。
	PostingNumbers []interface{} `json:"posting_numbers"` // 最新货件列表。
}

type V2FbsPostingProductCountrySetResponse struct {
	IsGTDNeeded bool `json:"is_gtd_needed"` // 有必要传送产品和货运的货物报关单（Cargo Customs Declaration）编号的标志。
	ProductID int64 `json:"product_id"` // 产品ID。
}

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	Dir string `json:"dir"` // 分类方向： - `asc` — 从小到大， - `desc` — 从大到小。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 响应中值的数量： - 最大值 — 1000， - 最小值 — 1。
	Offset int64 `json:"offset"` // 将在响应中跳过的元素数。 例如，如果“offset=10”，那么响应将从找到的第11个元素开始。
	With interface{} `json:"with"`
}

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"` // 货运号。
}

type V1SetPostingsResponse struct {
	Result interface{} `json:"result"`
}

type V1CarriageCancelResponse struct {
	Error string `json:"error"` // 错误描述。
	CarriageStatus string `json:"carriage_status"` // 发运状态。
}

type PostingV4PostingFbsListResponse struct {
	HasNext bool `json:"has_next"` // 若响应中未返回全部货件，则为`true`。
	Postings interface{} `json:"postings"` // 货件列表。
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
}

type V1AssemblyFbsProductListResponse struct {
	HasNext bool `json:"has_next"` // 响应中是否包含全部商品： - `true`——请使用新的 `offset`值重新请求以获取剩余数据； - `false`——响应中已包含所有值。
	Products []interface{} `json:"products"` // 商品列表。
	ProductsCount int32 `json:"products_count"` // 商品数量。
}

type PostingBooleanResponse struct {
	Result bool `json:"result"` // 处理请求的结果。 如果请求执行时无误，则为“true”。
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"` // 来自方法[/v1/carriage/create](#operation/CarriageAPI_CarriageCreate)的文件生成任务编号（也是运输标识符）。
}

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []interface{} `json:"posting_number"` // 货件ID。
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"` // 货运条形码。可以使用以下方法获取： [/v3/posting/fbs/get](#operation/PostingAPI_GetFbsPostingV3)、[/v3/posting/fbs/list](#operation/Posting...
}

// 货运信息。
type V2FbsPostingResponse struct {
	Result interface{} `json:"result"`
}

type V4FbsPostingShipPackageV4Request struct {
	PostingNumber string `json:"posting_number"` // 发货号。
	Products []interface{} `json:"products"` // 商品清单。
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber string `json:"posting_number"` // 货运号。
	ProductID int64 `json:"product_id"` // 产品ID。
	CountryIsoCode string `json:"country_iso_code"` // 根据标准ISO_3166-1添加的国家的两个字母代码 制造国家列表及其ISO代码可以使用该方法获得[/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountr...
}

type Fbsv4FbsPostingShipV4Request struct {
	Packages interface{} `json:"packages"` // 包装清单。 每个包装都包含订单分成的发货清单。
	PostingNumber string `json:"posting_number"` // 发货号。
	With interface{} `json:"with"`
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []interface{} `json:"posting_number"` // 货件ID。
}

type PostingV4PostingFbsListRequest struct {
	Translit bool `json:"translit"` // 若为`true`，则启用将地址从西里尔字母转写为拉丁字母。
	With interface{} `json:"with"`
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortDir interface{} `json:"sort_dir"`
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []interface{} `json:"related_posting_numbers"` // 货件号。
}

type PostingV4PostingFbsUnfulfilledListRequest struct {
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 响应中返回的值数量。
	SortDir interface{} `json:"sort_dir"`
	Translit bool `json:"translit"` // 则启用将地址从西里尔字母转写为拉丁字母。
	With interface{} `json:"with"`
}

type V1AssemblyCarriageProductListRequest struct {
	Limit int64 `json:"limit"` // 每页显示的数量。
	Cursor string `json:"cursor"` // 用于选择下一批数据的指针。
	Filter interface{} `json:"filter"`
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
}

package returns

// 其他信息。
type GetReturnsListResponseAdditionalInfo struct {
	IsSuperEconom bool `json:"is_super_econom"` // 如果退货为"超级经济"商品，显示`true`。
	IsOpened      bool `json:"is_opened"`       // 如果退货已开封，显示`true`。
}

// 商品信息。
// CurrencyCode values
type CurrencyCode string

const (
	CurrencyCodeRUB CurrencyCode = "RUB" // 俄罗斯卢布
	CurrencyCodeBYN CurrencyCode = "BYN" // 白俄罗斯卢布
	CurrencyCodeKZT CurrencyCode = "KZT" // 坚戈
	CurrencyCodeEUR CurrencyCode = "EUR" // 欧元
	CurrencyCodeUSD CurrencyCode = "USD" // 美元
	CurrencyCodeCNY CurrencyCode = "CNY" // 人民币
)

type V2Product struct {
	Price        int32        `json:"price"`         // 商品价格。
	SKU          int64        `json:"sku"`           // Ozon系统中的商品标识符 —— SKU。
	Name         string       `json:"name"`          // 商品名称。
	OfferID      string       `json:"offer_id"`      // 卖家系统中的商品标识符 —— 货号。
	CurrencyCode CurrencyCode `json:"currency_code"` // 您的价格所使用的货币。与您在个人中心中设置的货币相匹配。 可能的值: - `RUB` — 俄罗斯卢布, - `BYN` — 白俄罗斯卢布, - `KZT` — 坚戈, - `EUR` — 欧元, - `USD` — 美元, - `CNY` ...
}

// 退货申请和退款状态。
type V2ReturnsRfbsListV2ResponseState struct {
	GroupState           string `json:"group_state"`             // 根据应用的筛选器的申请状态。
	MoneyReturnStateName string `json:"money_return_state_name"` // 退款状态。
	State                string `json:"state"`                   // 申请状态。
	StateName            string `json:"state_name"`              // 退货申请状态的俄语名称。
}

// 申请信息。
type RfbsListResponseReturns struct {
	State         V2ReturnsRfbsListV2ResponseState `json:"state"`
	ClientName    string                           `json:"client_name"`    // 买家姓名。
	CreatedAt     string                           `json:"created_at"`     // 创建日期。
	OrderNumber   string                           `json:"order_number"`   // 订单号。
	PostingNumber string                           `json:"posting_number"` // 货件编号。
	Product       V2Product                        `json:"product"`
	ReturnID      int64                            `json:"return_id"`     // 退货申请的标识符。
	ReturnNumber  string                           `json:"return_number"` // 退货申请编号。
}

type V2ReturnsRfbsListResponse struct {
	Returns RfbsListResponseReturns `json:"returns"`
}

// 销毁费用。
type SellerReturnsv1MoneyUtilization struct {
	CurrencyCode string  `json:"currency_code"` // 货币。
	Price        float64 `json:"price"`         // 销毁费用。
}

// 存储费用。
type SellerReturnsv1MoneyStorage struct {
	CurrencyCode string  `json:"currency_code"` // 货币。
	Price        float64 `json:"price"`         // 存储费用。
}

// 存储信息。
type GetReturnsListResponseStorage struct {
	ArrivedMoment           string                          `json:"arrived_moment"` // 退货准备交付给卖家的日期。
	Days                    int64                           `json:"days"`           // 退货等待交付给卖家的天数。
	UtilizationSum          SellerReturnsv1MoneyUtilization `json:"utilization_sum"`
	UtilizationForecastDate string                          `json:"utilization_forecast_date"` // 预计销毁日期。
	Sum                     SellerReturnsv1MoneyStorage     `json:"sum"`
	TarifficationFirstDate  string                          `json:"tariffication_first_date"` // 计算存储费用的第一天。
	TarifficationStartDate  string                          `json:"tariffication_start_date"` // 计算存储费用的开始日期。
}

// 退货信息。
type GetReturnsListResponseLogistic struct {
	Barcode                         string `json:"barcode"`                            // 退货标签的条形码。
	TechnicalReturnMoment           string `json:"technical_return_moment"`            // 商品被标记为技术退货的日期。
	FinalMoment                     string `json:"final_moment"`                       // 退货到达履约中心或卖家收到退货的日期。
	CancelledWithCompensationMoment string `json:"cancelled_with_compensation_moment"` // 向卖家补偿退货的日期。
	ReturnDate                      string `json:"return_date"`                        // 买家退货的日期。
}

type GetReturnsListResponseExemplar struct {
	ID int64 `json:"id"` // 实例ID。
}

// 退货状态。
type GetReturnsListResponseVisualStatus struct {
	SysName     string `json:"sys_name"`     // 退货状态的系统名称。
	ID          int32  `json:"id"`           // 退货状态ID。
	DisplayName string `json:"display_name"` // 退货状态名称。
}

// 退货状态信息。
type GetReturnsListResponseVisual struct {
	Status       GetReturnsListResponseVisualStatus `json:"status"`
	ChangeMoment string                             `json:"change_moment"` // 退货状态的变更日期。
}

// 退货所在仓库的信息。
type GetReturnsListResponsePlaceNow struct {
	ID      int64  `json:"id"`      // 仓库ID。
	Name    string `json:"name"`    // 名称。
	Address string `json:"address"` // 地址。
}

// 商品价格。
type SellerReturnsv1MoneyProduct struct {
	CurrencyCode string  `json:"currency_code"` // 货币。
	Price        float64 `json:"price"`         // 商品价格。
}

// 不含佣金的商品价格。
type SellerReturnsv1MoneyWithoutCommission struct {
	CurrencyCode string  `json:"currency_code"` // 货币。
	Price        float64 `json:"price"`         // 不含佣金的商品价格。
}

// 佣金费用。
type SellerReturnsv1MoneyCommission struct {
	Price        float64 `json:"price"`         // 佣金费用。
	CurrencyCode string  `json:"currency_code"` // 货币。
}

// 商品信息。
type GetReturnsListResponseProduct struct {
	OfferID                string                                `json:"offer_id"` // 卖家系统中的商品标识符是商品货号。
	Name                   string                                `json:"name"`     // 商品名称。
	Price                  SellerReturnsv1MoneyProduct           `json:"price"`
	PriceWithoutCommission SellerReturnsv1MoneyWithoutCommission `json:"price_without_commission"`
	CommissionPercent      float64                               `json:"commission_percent"` // 佣金比例。
	Commission             SellerReturnsv1MoneyCommission        `json:"commission"`
	Quantity               int32                                 `json:"quantity"` // 产品数量。
	SKU                    int64                                 `json:"sku"`      // 商品在Ozon系统中的ID（SKU）。
}

// 赔偿状态。
// SysName values
type SysName string

const (
	SysNameSent               SysName = "Sent"               // 已发送；
	SysNameReceived           SysName = "Received"           // 已收到；
	SysNameCanceled           SysName = "Canceled"           // 已取消；
	SysNameDecompensationSent SysName = "DecompensationSent" // 已进行赔偿返还
)

type GetReturnsListResponseCompensationStatus struct {
	SysName     SysName `json:"sys_name"`     // 状态的系统名称： - `Sent`——已发送； - `Received`——已收到； - `Canceled`——已取消； - `DecompensationSent`——已进行赔偿返还。
	ID          int32   `json:"id"`           // 状态标识符。
	DisplayName string  `json:"display_name"` // 状态名称： - “发送进行赔偿”， - “您已收到赔偿”， - "赔偿已取消"， - "已进行赔偿返还"。
}

// 赔偿状态信息。
type GetReturnsListResponseCompensation struct {
	Status       GetReturnsListResponseCompensationStatus `json:"status"`
	ChangeMoment string                                   `json:"change_moment"` // 赔偿状态的变更日期。
}

// 退货运往的仓库信息。
type GetReturnsListResponsePlaceTarget struct {
	ID      int64  `json:"id"`      // 仓库ID。
	Name    string `json:"name"`    // 名称。
	Address string `json:"address"` // 地址。
}

type GetReturnsListResponseReturnsItem struct {
	CompensationStatus GetReturnsListResponseCompensation   `json:"compensation_status"`
	ID                 int64                                `json:"id"`                 // 退货ID。
	ReturnReasonName   string                               `json:"return_reason_name"` // 退货或取消的原因。
	Schema             string                               `json:"schema"`             // 退货方案： `FBS`； `FBO`。
	OrderID            int64                                `json:"order_id"`           // 订单ID。
	Exemplars          []GetReturnsListResponseExemplar     `json:"exemplars"`          // 退货实例信息。
	OrderNumber        string                               `json:"order_number"`       // 订单编号。
	Visual             GetReturnsListResponseVisual         `json:"visual"`
	AdditionalInfo     GetReturnsListResponseAdditionalInfo `json:"additional_info"`
	ReturnClearingID   int64                                `json:"return_clearing_id"` // 初始货件的退货条形码。
	CompanyID          int64                                `json:"company_id"`         // 卖家ID。
	TargetPlace        GetReturnsListResponsePlaceTarget    `json:"target_place"`
	Storage            GetReturnsListResponseStorage        `json:"storage"`
	Logistic           GetReturnsListResponseLogistic       `json:"logistic"`
	ClearingID         int64                                `json:"clearing_id"` // 初始货件条形码。
	Type               string                               `json:"type_"`       // 退货类型： `Cancellation` - 取消订单（交货前）； `FullReturn` - 完全拒收（交货时）； `PartialReturn` - 部分拒收（交货时）； `ClientReturn` - 客户退货（交货后）； `Un...
	Place              GetReturnsListResponsePlaceNow       `json:"place"`
	Product            GetReturnsListResponseProduct        `json:"product"`
	SourceID           int64                                `json:"source_id"`      // 先前的退货ID。
	PostingNumber      string                               `json:"posting_number"` // 货件编号。
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"` // 退货申请的标识符。
}

// 根据退货状态变更日期筛选。
type V1TimeRangeVisualStatus struct {
	TimeFrom string `json:"time_from"` // 开始日期。
	TimeTo   string `json:"time_to"`   // 结束日期。
}

// 根据退货创建日期筛选。
type V1TimeRangeReturnDate struct {
	TimeFrom string `json:"time_from"` // 开始日期。
	TimeTo   string `json:"time_to"`   // 结束日期。
}

// 根据存储费用开始日期筛选。
type V1TimeRangeStorageTariffication struct {
	TimeFrom string `json:"time_from"` // 开始日期。
	TimeTo   string `json:"time_to"`   // 结束日期。
}

// 筛选条件。在请求中仅使用一个筛选器：`logistic_return_date`、`storage_tariffication_start_date` 或`visual_status_change_moment`，否则会返回错误。
// CompensationStatusID values
type CompensationStatusID string

const (
	CompensationStatusIDV2 CompensationStatusID = "2" // 已收到；
	CompensationStatusIDV3 CompensationStatusID = "3" // 已取消；
	CompensationStatusIDV4 CompensationStatusID = "4" // 已进行赔偿返还
)

// VisualStatusName values
type VisualStatusName string

const (
	VisualStatusNameDisputeOpened                                 VisualStatusName = "DisputeOpened"                                 // 与买家有争议；
	VisualStatusNameOnSellerApproval                              VisualStatusName = "OnSellerApproval"                              // 等待卖家批准；
	VisualStatusNameArrivedAtReturnPlace                          VisualStatusName = "ArrivedAtReturnPlace"                          // 到达取货点；
	VisualStatusNameOnSellerClarification                         VisualStatusName = "OnSellerClarification"                         // 等待卖家确认；
	VisualStatusNameOnSellerClarificationAfterPartialCompensation VisualStatusName = "OnSellerClarificationAfterPartialCompensation" // 部分补偿后等待卖家确认；
	VisualStatusNameOfferedPartialCompensation                    VisualStatusName = "OfferedPartialCompensation"                    // 提供部分补偿；
	VisualStatusNameReturnMoneyApproved                           VisualStatusName = "ReturnMoneyApproved"                           // 退货款项已批准；
	VisualStatusNamePartialCompensationReturned                   VisualStatusName = "PartialCompensationReturned"                   // 部分款项已退还；
	VisualStatusNameCancelledDisputeNotOpen                       VisualStatusName = "CancelledDisputeNotOpen"                       // 退货被拒绝，争议未开启；
	VisualStatusNameRejected                                      VisualStatusName = "Rejected"                                      // 申请被拒绝；
	VisualStatusNameCrmRejected                                   VisualStatusName = "CrmRejected"                                   // Ozon拒绝申请；
	VisualStatusNameCancelled                                     VisualStatusName = "Cancelled"                                     // 申请已取消；
	VisualStatusNameApproved                                      VisualStatusName = "Approved"                                      // 卖家批准了申请；
	VisualStatusNameApprovedByOzon                                VisualStatusName = "ApprovedByOzon"                                // Ozon批准了申请；
	VisualStatusNameReceivedBySeller                              VisualStatusName = "ReceivedBySeller"                              // 卖家已收到退货；
	VisualStatusNameMovingToSeller                                VisualStatusName = "MovingToSeller"                                // 退货正在运往卖家；
	VisualStatusNameReturningToSellerByCourier                    VisualStatusName = "ReturningToSellerByCourier"                    // 快递员正在将退货送回卖家；
	VisualStatusNameUtilizing                                     VisualStatusName = "Utilizing"                                     // 正在销毁中；
	VisualStatusNameUtilized                                      VisualStatusName = "Utilized"                                      // 已销毁；
	VisualStatusNameMoneyReturned                                 VisualStatusName = "MoneyReturned"                                 // 已退还全款给买家；
	VisualStatusNamePartialCompensationInProcess                  VisualStatusName = "PartialCompensationInProcess"                  // 部分退款已批准；
	VisualStatusNameDisputeYouOpened                              VisualStatusName = "DisputeYouOpened"                              // 卖家提出争议；
	VisualStatusNameCompensationRejected                          VisualStatusName = "CompensationRejected"                          // 补偿被拒绝；
	VisualStatusNameDisputeOpening                                VisualStatusName = "DisputeOpening"                                // 已向客服发送请求；
	VisualStatusNameCompensationOffered                           VisualStatusName = "CompensationOffered"                           // 等待您对补偿的决定；
	VisualStatusNameWaitingCompensation                           VisualStatusName = "WaitingCompensation"                           // 等待补偿；
	VisualStatusNameSendingError                                  VisualStatusName = "SendingError"                                  // 发送客服请求时发生错误；
	VisualStatusNameCompensationRejectedBySla                     VisualStatusName = "CompensationRejectedBySla"                     // 补偿请求过期；
	VisualStatusNameCompensationRejectedBySeller                  VisualStatusName = "CompensationRejectedBySeller"                  // 卖家拒绝补偿；
	VisualStatusNameMovingToOzon                                  VisualStatusName = "MovingToOzon"                                  // 正在运往Ozon仓库；
	VisualStatusNameReturnedToOzon                                VisualStatusName = "ReturnedToOzon"                                // 已返回Ozon仓库；
	VisualStatusNameMoneyReturnedBySystem                         VisualStatusName = "MoneyReturnedBySystem"                         // 快速退款；
	VisualStatusNameWaitingShipment                               VisualStatusName = "WaitingShipment"                               // 等待发货
)

type GetReturnsListRequestFilter struct {
	LogisticReturnDate            V1TimeRangeReturnDate           `json:"logistic_return_date"`
	PostingNumbers                []string                        `json:"posting_numbers"`    // 根据货件编号筛选。请勿传递超过 50 个货盒单号。
	ProductName                   string                          `json:"product_name"`       // 根据商品名称筛选。
	VisualStatusName              VisualStatusName                `json:"visual_status_name"` // 根据退货状态筛选： - `DisputeOpened` — 与买家有争议； - `OnSellerApproval` — 等待卖家批准； - `ArrivedAtReturnPlace` — 到达取货点； - `OnSellerClarif...
	Barcode                       string                          `json:"barcode"`            // 根据退货标签条形码筛选。
	StorageTarifficationStartDate V1TimeRangeStorageTariffication `json:"storage_tariffication_start_date"`
	VisualStatusChangeMoment      V1TimeRangeVisualStatus         `json:"visual_status_change_moment"`
	OrderID                       int64                           `json:"order_id"`               // 根据订单ID筛选。
	OfferID                       string                          `json:"offer_id"`               // 根据商品货号筛选。
	WarehouseID                   int64                           `json:"warehouse_id"`           // 根据仓库ID筛选。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	ReturnSchema                  string                          `json:"return_schema"`          // 根据配送方案筛选：`FBS` 或`FBO`。
	CompensationStatusID          int32                           `json:"compensation_status_id"` // 按赔偿状态筛选： - `1`——已发送； - `2`——已收到； - `3`——已取消； - `4`——已进行赔偿返还。
}

type V1GetReturnsListRequest struct {
	Filter GetReturnsListRequestFilter `json:"filter"`
	Limit  int32                       `json:"limit"`   // 加载的退货数量。
	LastID int64                       `json:"last_id"` // 最后加载的退货ID。
}

type V1Empty any

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnForBackWay int64 `json:"return_for_back_way"` // 退还给买家的商品运费金额。
	ReturnID         int64 `json:"return_id"`           // 退货申请的标识符。
}

// 按取消申请状态筛选： - `ALL` — 所有状态的申请， - `ON_APPROVAL` — 审核中申请， - `APPROVED` — 已确认申请， - `REJECTED` — 已拒绝申请。
type V2CancellationStateEnumFilters string

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID                int64  `json:"return_id"`                 // 退货申请的标识符。
	ReturnMethodDescription string `json:"return_method_description"` // 商品退货方式。
}

// 申请创建期间。
type CreatedAt struct {
	From string `json:"from"` // 开始日期。
	To   string `json:"to"`   // 结束日期。
}

// 筛选器。
// GroupState values
type GroupState string

const (
	GroupStateAll         GroupState = "All"         // 所有申请
	GroupStateNew         GroupState = "New"         // 新申请
	GroupStateDelivering  GroupState = "Delivering"  // 在途中
	GroupStateCheckout    GroupState = "Checkout"    // 审核中
	GroupStateArbitration GroupState = "Arbitration" // 具争议
	GroupStateApproved    GroupState = "Approved"    // 已批准
	GroupStateRejected    GroupState = "Rejected"    // 已拒绝
)

type V2ReturnsRfbsFilter struct {
	PostingNumber string    `json:"posting_number"` // 货件编号。
	GroupState    []string  `json:"group_state"`    // 根据申请状态筛选: - `All` — 所有申请。 - `New` — 新申请。 - `Delivering` — 在途中。 - `Checkout` — 审核中。 - `Arbitration` — 具争议。 - `Approved` —...
	CreatedAt     CreatedAt `json:"created_at"`
	OfferID       string    `json:"offer_id"` // 卖家系统中的商品标识符 —— 货号。
}

type V2ReturnsRfbsListRequest struct {
	Filter V2ReturnsRfbsFilter `json:"filter"`
	LastID int32               `json:"last_id"` // 页面上最后一个值的标识符——`return_id`。在第一次请求时，请将此字段留空。
	Limit  int32               `json:"limit"`   // 响应中的值数量。
}

type V1GetReturnsListResponse struct {
	Returns []GetReturnsListResponseReturnsItem `json:"returns"`  // 退货信息。
	HasNext bool                                `json:"has_next"` // 如果卖家有其他退货，显示`true`。
}

type V2ReturnsRfbsRejectRequest struct {
	ReturnID          int64  `json:"return_id"`           // 退货申请的标识符。
	Comment           string `json:"comment"`             // 备注。 如果 [/v2/returns/rfbs/get](#operation/RFBSReturnsAPI_ReturnsRfbsGetV2) 方法的响应中 `rejection_reason.is_comment_required` ...
	RejectionReasonID int64  `json:"rejection_reason_id"` // 取消原因的标识符。 从 [/v2/returns/rfbs/get](#operation/RFBSReturnsAPI_ReturnsRfbsGetV2) 响应中获取的原因列表中传递标识符，参数为 `rejection_reason`。
}

// 取消发起人： - `SELLER` — 卖家， - `CLIENT` — 买家， - `OZON` — Ozon, - `SYSTEM` — 系统， - `DELIVERY` — 配送服务。
type V2CancellationInitiator string

// 退货状态信息。
type V2ReturnsRfbsGetV2ResponseState struct {
	StateName string `json:"state_name"` // 状态的俄语名称。
	State     string `json:"state"`      // 状态。
}

// 退货方式信息。
type RfbsGetV2ResponseClientReturnMethodType struct {
	ID   int32  `json:"id"`   // 标识符。
	Name string `json:"name"` // 名称。
}

// 取消原因。
type GetConditionalCancellationListV2ResponseCancellationReason struct {
	Name string `json:"name"` // 取消原因名称。
	ID   int64  `json:"id"`   // 取消原因标识符。
}

// 申请状态： - `ON_APPROVAL` — 审核中， - `APPROVED` — 已确认， - `REJECTED` — 已拒绝。
type V2CancellationState string

// 取消申请的状态。
type GetConditionalCancellationListV2ResponseState struct {
	ID    int64               `json:"id"`   // 状态标识符。
	Name  string              `json:"name"` // 状态名称。
	State V2CancellationState `json:"state"`
}

type GetConditionalCancellationListV2ResponseResult struct {
	AutoApproveDate           string                                                     `json:"auto_approve_date"` // 申请将在此日期后自动确认。
	CancellationID            int64                                                      `json:"cancellation_id"`   // 取消申请标识符。
	CancellationInitiator     V2CancellationInitiator                                    `json:"cancellation_initiator"`
	CancellationReason        GetConditionalCancellationListV2ResponseCancellationReason `json:"cancellation_reason"`
	CancellationReasonMessage string                                                     `json:"cancellation_reason_message"` // 取消申请中由取消发起人手动填写的备注。
	CancelledAt               string                                                     `json:"cancelled_at"`                // 取消申请的创建日期。
	OrderDate                 string                                                     `json:"order_date"`                  // 订单的创建日期。
	PostingNumber             string                                                     `json:"posting_number"`              // 货件编号。
	ApproveComment            string                                                     `json:"approve_comment"`             // 在确认或拒绝取消申请时填写的备注。
	ApproveDate               string                                                     `json:"approve_date"`                // 取消申请确认或拒绝的日期。
	SourceID                  int64                                                      `json:"source_id"`                   // 上一次取消申请的标识符。 用于保持向后兼容性。
	State                     GetConditionalCancellationListV2ResponseState              `json:"state"`
	TPLIntegrationType        string                                                     `json:"tpl_integration_type"` // 与配送服务的集成类型。
}

type V2GetConditionalCancellationListV2Response struct {
	Counter int64                                            `json:"counter"` // `ON_APPROVAL` 状态申请的计数器。
	LastID  int64                                            `json:"last_id"` // 页面上最后一个值的标识符。 要获取后续值，请指定上一次请求响应中的 `last_id`。
	Result  []GetConditionalCancellationListV2ResponseResult `json:"result"`  // 取消申请的详细信息。
}

// 过滤器。
// CancellationInitiator values
type CancellationInitiator string

const (
	CancellationInitiatorSeller   CancellationInitiator = "SELLER"   // 卖家
	CancellationInitiatorClient   CancellationInitiator = "CLIENT"   // 买家
	CancellationInitiatorOzon     CancellationInitiator = "OZON"     // Ozon
	CancellationInitiatorSystem   CancellationInitiator = "SYSTEM"   // 系统
	CancellationInitiatorDelivery CancellationInitiator = "DELIVERY" // 配送服务
)

// State values
type State string

const (
	StateALL        State = "ALL"         // 所有状态的申请
	StateONApproval State = "ON_APPROVAL" // 审核中申请
	StateApproved   State = "APPROVED"    // 已确认申请
	StateRejected   State = "REJECTED"    // 已拒绝申请
)

type GetConditionalCancellationListV2RequestFilters struct {
	PostingNumber         []string                       `json:"posting_number"`         // 按货件编号筛选。
	State                 V2CancellationStateEnumFilters `json:"state"`                  // 按取消申请状态筛选： - `ALL` — 所有状态的申请， - `ON_APPROVAL` — 审核中申请， - `APPROVED` — 已确认申请， - `REJECTED` — 已拒绝申请。
	CancellationInitiator []V2CancellationInitiator      `json:"cancellation_initiator"` // 取消发起人： - `SELLER` — 卖家， - `CLIENT` — 买家， - `OZON` — Ozon， - `SYSTEM` — 系统， - `DELIVERY` — 配送服务。
}

// 附加信息。
type GetConditionalCancellationListV2RequestWith struct {
	Counter bool `json:"counter"` // 表示需要在响应中返回处于 `ON_APPROVAL` 状态的申请数量的标志。
}

type V2GetConditionalCancellationListV2Request struct {
	Filters GetConditionalCancellationListV2RequestFilters `json:"filters"`
	LastID  int64                                          `json:"last_id"` // 页面上最后一个值的标识符。在首次请求时此字段留空。 要获取后续值，请指定上一次请求响应中的 `last_id`。
	Limit   int32                                          `json:"limit"`   // 响应中包含的申请总数。
	With    GetConditionalCancellationListV2RequestWith    `json:"with"`
}

type RfbsGetV2ResponseAvailableAction struct {
	ID   int32  `json:"id"`   // 操作标识符。
	Name string `json:"name"` // 操作名称。
}

// 退货原因信息。
type RfbsGetV2ResponseReturnReason struct {
	ID       int32  `json:"id"`        // 原因的标识符。
	IsDefect bool   `json:"is_defect"` // 指示商品是否有瑕疵。
	Name     string `json:"name"`      // 原因的描述。
}

type RfbsGetV2ResponseRejectionReason struct {
	ID                int32  `json:"id"`                  // 原因的标识符。
	IsCommentRequired bool   `json:"is_comment_required"` // 指示是否需要备注。
	Name              string `json:"name"`                // 原因的描述。
	Hint              string `json:"hint"`                // 有关退货的进一步操作的提示。
}

// 申请信息。
type RfbsGetResponseReturns struct {
	RejectionComment        string                                  `json:"rejection_comment"` // 有关申请被拒绝的备注。
	ReturnNumber            string                                  `json:"return_number"`     // 退货申请编号。
	State                   V2ReturnsRfbsGetV2ResponseState         `json:"state"`
	ClientName              string                                  `json:"client_name"` // 买家姓名。
	ReturnReason            RfbsGetV2ResponseReturnReason           `json:"return_reason"`
	ClientReturnMethodType  RfbsGetV2ResponseClientReturnMethodType `json:"client_return_method_type"`
	Comment                 string                                  `json:"comment"` // 买家评论。
	Product                 V2Product                               `json:"product"`
	RejectionReason         []RfbsGetV2ResponseRejectionReason      `json:"rejection_reason"`          // 申请被拒绝的原因的信息。
	RuPostTrackingNumber    string                                  `json:"ru_post_tracking_number"`   // 跟踪号码。
	ClientPhoto             []string                                `json:"client_photo"`              // 商品照片链接。
	OrderNumber             string                                  `json:"order_number"`              // 订单号。
	PostingNumber           string                                  `json:"posting_number"`            // 货件编号。
	ReturnMethodDescription string                                  `json:"return_method_description"` // 商品退货方式。
	WarehouseID             int64                                   `json:"warehouse_id"`              // 仓库标识符。
	AvailableActions        []RfbsGetV2ResponseAvailableAction      `json:"available_actions"`         // 申请的可用操作的信息。
	CreatedAt               string                                  `json:"created_at"`                // 申请创建日期。
}

type V2ReturnsRfbsGetResponse struct {
	Returns RfbsGetResponseReturns `json:"returns"`
}

type V2ReturnsRfbsCompensateRequest struct {
	CompensationAmount string `json:"compensation_amount"` // 赔偿金额。
	ReturnID           int64  `json:"return_id"`           // 退货申请的标识符。
}

type V2ReturnsRfbsGetRequest struct {
	ReturnID int64 `json:"return_id"` // 退货申请标识符。通过方法 [/v2/returns/rfbs/list](#operation/RFBSReturnsAPI_ReturnsRfbsListV2) 获取。
}

type V1ReturnsRfbsActionSetRequest struct {
	Comment            string  `json:"comment"`             // 卖家评论。 对于 `id: -1` 和 `id: -10`，备注为必填项。
	CompensationAmount float64 `json:"compensation_amount"` // 赔偿金额。 对于 `id: 1020`，备注也为必填项。
	ID                 int32   `json:"id"`                  // 操作标识符。 获取可用操作 `returns.available_actions` ，请使用方法 [/v2/returns/rfbs/get](#operation/RFBSReturnsAPI_ReturnsRfbsGetV2)。
	RejectionReasonID  int32   `json:"rejection_reason_id"` // 取消原因的标识符。 对于 `id: -1` 和 `id: -10`，备注为必填项。 获取可用取消原因 `returns.rejection_reason`，请使用方法 [/v2/returns/rfbs/get](#operation/RF...
	ReturnForBackWay   float64 `json:"return_for_back_way"` // 退还给买家的商品运费金额。 负值将被视为 `0`。
	ReturnID           int64   `json:"return_id"`           // 退货申请的标识符。
}

package fbs

// ReasonsID values
type ReasonsID string

const (
	ReasonsIDV352 ReasonsID = "352" // 在卖家仓库已无商品
	ReasonsIDV400 ReasonsID = "400" // 只剩下残次品
	ReasonsIDV401 ReasonsID = "401" // 卖家拒绝了仲裁
	ReasonsIDV402 ReasonsID = "402" // 其他（卖家错误）
	ReasonsIDV665 ReasonsID = "665" // 买家没有收货
	ReasonsIDV666 ReasonsID = "666" // 快递服务退货：在该区域没有快递
	ReasonsIDV667 ReasonsID = "667" // 订单被快递弄丢
)

// TypeID values
type TypeID string

const (
	TypeIDBuyer  TypeID = "buyer"  // 买家
	TypeIDSeller TypeID = "seller" // 卖家
)

type RelatedPostingCancelReasons struct {
	ID     int64  `json:"id"`      // 取消原因ID： - `352` — 在卖家仓库已无商品。 - `400` — 只剩下残次品。 - `401` — 卖家拒绝了仲裁。 - `402` — 其他（卖家错误）。 - `665` — 买家没有收货。 - `666` — 快递服务退货...
	Title  string `json:"title"`   // 描述取消原因。
	TypeID TypeID `json:"type_id"` // 取消货运提出方： - `buyer` — 买家， - `seller` — 卖家。
}

type RelatedPostingCancelReason struct {
	PostingNumber string                        `json:"posting_number"` // 货运号。
	Reasons       []RelatedPostingCancelReasons `json:"reasons"`        // 取消订单原因。
}

type PostingCancelReasonResponse struct {
	Result []RelatedPostingCancelReason `json:"result"` // 请求结果。
}

// CurrencyCode values
type CurrencyCode string

const (
	CurrencyCodeRUB CurrencyCode = "RUB" // 俄罗斯卢布
	CurrencyCodeBYN CurrencyCode = "BYN" // 白俄罗斯卢布
	CurrencyCodeKZT CurrencyCode = "KZT" // 坚戈
	CurrencyCodeEUR CurrencyCode = "EUR" // 欧元
	CurrencyCodeUSD CurrencyCode = "USD" // 美元
	CurrencyCodeCNY CurrencyCode = "CNY" // 元
)

type V3FbsPostingProduct struct {
	Price        string       `json:"price"`         // 商品价格。
	Quantity     int32        `json:"quantity"`      // 运输中的商品数量。
	SKU          int64        `json:"sku"`           // 在Ozon系统中的商品ID — SKU。
	CurrencyCode CurrencyCode `json:"currency_code"` // 价格货币，其与个人中心中设置的币种相匹配。 可能的值： - `RUB` — 俄罗斯卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
	Imei         []string     `json:"imei"`          // 移动设备的 IMEI 列表。
	Name         string       `json:"name"`          // 商品名称。
	OfferID      string       `json:"offer_id"`      // 在卖家系统中的商品ID — 货号。
}

// 快递方式。
type V3DeliveryMethod struct {
	TPLProviderID int64  `json:"tpl_provider_id"` // 快递服务ID。
	Warehouse     string `json:"warehouse"`       // 仓库名称。
	WarehouseID   int64  `json:"warehouse_id"`    // 仓库ID。
	ID            int64  `json:"id"`              // 快递方式ID。
	Name          string `json:"name"`            // 快递方式名称。
	TPLProvider   string `json:"tpl_provider"`    // 快递服务。
}

type SPostings struct {
	PostingNumber string `json:"posting_number"` // 货件编号。
	Quantity      int32  `json:"quantity"`       // 货件中的商品数量。
}

type PostingFbsPostingDeliveringRequest struct {
	PostingNumber []string `json:"posting_number"` // 货件ID。
}

// 快递地址信息。
type V3Address struct {
	ProviderPvzCode string  `json:"provider_pvz_code"` // 3PL提供商的订单提货点的代码。
	Region          string  `json:"region"`            // 快递区域。
	City            string  `json:"city"`              // 快递城市。
	District        string  `json:"district"`          // 快递地区。
	PvzCode         int64   `json:"pvz_code"`          // 订单取货点代码。
	ZipCode         string  `json:"zip_code"`          // 收件人邮编。
	AddressTail     string  `json:"address_tail"`      // 文本格式的地址。
	Comment         string  `json:"comment"`           // 订单评价。
	Country         string  `json:"country"`           // 快递国家。
	Latitude        float64 `json:"latitude"`          // 宽。
	Longitude       float64 `json:"longitude"`         // （时间的）长度。
}

// 配送方法信息。
type PostingV4PostingFbsListResponsePostingsDeliveryMethod struct {
	ID            int64  `json:"id"`              // 配送方式标识符。
	Name          string `json:"name"`            // 配送方法名称。
	TPLProvider   string `json:"tpl_provider"`    // 配送服务。
	TPLProviderID int64  `json:"tpl_provider_id"` // 配送服务标识符。
	Warehouse     string `json:"warehouse"`       // 仓库名称。
	WarehouseID   int64  `json:"warehouse_id"`    // 仓库标识符。
}

type V1CarriageCancelResponse struct {
	Error          string `json:"error"`           // 错误描述。
	CarriageStatus string `json:"carriage_status"` // 发运状态。
}

// 商品价格。
type MoneyPostingMoney struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 商品佣金。
type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission struct {
	Amount   float64 `json:"amount"`   // 商品佣金金额。
	Currency string  `json:"currency"` // 计算佣金所使用的货币代码。
	Percent  int64   `json:"percent"`  // 佣金比例。
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts struct {
	ProductID            int64                                                                             `json:"product_id"`           // Ozon系统中的商品标识符，即SKU。
	TotalDiscountValue   float64                                                                           `json:"total_discount_value"` // 折扣金额。
	CustomerPrice        MoneyPostingMoney                                                                 `json:"customer_price"`
	Payout               float64                                                                           `json:"payout"`                 // 卖家应收款。
	Price                float64                                                                           `json:"price"`                  // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	Quantity             int64                                                                             `json:"quantity"`               // 货件中的商品数量。
	TotalDiscountPercent float64                                                                           `json:"total_discount_percent"` // 折扣百分比。
	Actions              []string                                                                          `json:"actions"`                // 促销活动列表。
	Commission           PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProductsCommission `json:"commission"`
	OldPrice             float64                                                                           `json:"old_price"` // 折扣前的价格。商品卡上会以划线价显示。
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsProducts struct {
	Price               MoneyPostingMoney `json:"price"`
	ProductColor        string            `json:"product_color"`         // 商品颜色。
	Quantity            int32             `json:"quantity"`              // 货件中的商品数量。
	IsMarketplaceBuyout bool              `json:"is_marketplace_buyout"` // `true`，表示Ozon已购买该商品，用于配送至欧亚经济联盟及其他国家。
	SKU                 int64             `json:"sku"`                   // Ozon系统中的商品标识符，即SKU。
	Weight              float64           `json:"weight"`                // 商品含包装重量。
	Imei                []string          `json:"imei"`                  // 移动设备的 IMEI 列表。
	IsBlrTraceable      bool              `json:"is_blr_traceable"`      // `true`，表示商品可追溯。
	Name                string            `json:"name"`                  // 商品名称。
	OfferID             string            `json:"offer_id"`              // 卖家系统中的商品标识符，即货号。
}

// 排序方向： - `ASC`——升序， - `DESC`——降序。
type V1AssemblyFbsPostingListRequestSortDirEnum string

// 筛选器。
type V1AssemblyFbsPostingListRequestFilter struct {
	CutoffTo         string `json:"cutoff_to"`          // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	DeliveryMethodID int64  `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	CutoffFrom       string `json:"cutoff_from"`        // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
}

type V1AssemblyFbsPostingListRequest struct {
	Cursor  string                                     `json:"cursor"` // 用于选择下一批数据的指针。
	Filter  V1AssemblyFbsPostingListRequestFilter      `json:"filter"`
	Limit   int64                                      `json:"limit"` // 每页显示的数量。
	SortDir V1AssemblyFbsPostingListRequestSortDirEnum `json:"sort_dir"`
}

type PostingCancelFbsPostingRequest struct {
	CancelReasonID      int64  `json:"cancel_reason_id"`      // 取消运输的原因ID。
	CancelReasonMessage string `json:"cancel_reason_message"` // 关于取消的附加信息。如果`cancel_reason_id = 402`，参数是必须的。
	PostingNumber       string `json:"posting_number"`        // 货件ID。
}

type V1AssemblyCarriageProductListResponseProduct struct {
	OfferID        string   `json:"offer_id"`        // 卖家系统中的商品标识符——货号。
	PictureURL     string   `json:"picture_url"`     // 商品图片链接。
	PostingNumbers []string `json:"posting_numbers"` // 货件编号。
	ProductName    string   `json:"product_name"`    // 商品名称。
	Quantity       int64    `json:"quantity"`        // 商品数量。
	SKU            int64    `json:"sku"`             // Ozon系统中的商品标识符——SKU。
}

type V1AssemblyCarriageProductListResponse struct {
	Cursor   string                                         `json:"cursor"`   // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Products []V1AssemblyCarriageProductListResponseProduct `json:"products"` // 商品列表。
}

// 需提供制造国、货运报关单号、商品批次登记号、“诚实标志”、其他标识或重量的商品列表，以便将货件状态更新至下一阶段。
type V3FbsPostingRequirementsV3 struct {
	ProductsRequiringImei          []string `json:"products_requiring_imei"`           // 需要提供 IMEI 的商品 ID 列表。
	ProductsRequiringChangeCountry []string `json:"products_requiring_change_country"` // 需要修改生产国家的商品（SKU）编号列表。要修改生产国家，请使用方法 [/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountryProductFbsPos...
	ProductsRequiringGTD           []string `json:"products_requiring_gtd"`            // 必须提供货运报关单号（CCD）的商品 ID（SKU）列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProduct...
	ProductsRequiringCountry       []string `json:"products_requiring_country"`        // 需要上传制造国信息的商品ID列表 (SKU)。 要配货，请上传上述商品的制造国信息，通过方法 [/v2/posting/fbs/product/country/set](#operation/PostingAPI_SetCountryPro...
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"` // 需要提供“诚实标志”标签的商品 ID（SKU）列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExe...
	ProductsRequiringJwUin         []string `json:"products_requiring_jw_uin"`         // 需要提供首饰唯一识别号（UIN）的商品列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExempla...
	ProductsRequiringRnpt          []string `json:"products_requiring_rnpt"`           // 需要提供商品批次注册号（RNPT）的商品 ID（SKU）列表。 在配货之前，请通过方法 [/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProdu...
}

type V1AssemblyCarriagePostingListResponsePostingProduct struct {
	SKU         int64  `json:"sku"`          // Ozon系统中的商品标识符——SKU。
	OfferID     string `json:"offer_id"`     // 卖家系统中的商品标识符——货号。
	PictureURL  string `json:"picture_url"`  // 商品图片链接。
	ProductName string `json:"product_name"` // 商品名称。
	Quantity    int64  `json:"quantity"`     // 商品数量。
}

type V1AssemblyCarriagePostingListResponsePosting struct {
	AssemblyCode  string                                                `json:"assembly_code"`   // 拣货单代码。
	CanPrintLabel bool                                                  `json:"can_print_label"` // `true`，前提是可以打印标签。
	PostingNumber string                                                `json:"posting_number"`  // 货件编号。
	Products      []V1AssemblyCarriagePostingListResponsePostingProduct `json:"products"`        // 商品列表。
}

// 筛选器。
type V1AssemblyFbsProductListRequestFilter struct {
	CutoffFrom       string `json:"cutoff_from"`        // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	CutoffTo         string `json:"cutoff_to"`          // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	DeliveryMethodID int64  `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
}

// 排序方向： - `ASC`——升序， - `DESC`——降序。
type V1AssemblyFbsProductListRequestSortDirEnum string

type V1AssemblyFbsProductListRequest struct {
	Filter  V1AssemblyFbsProductListRequestFilter      `json:"filter"`
	Limit   int64                                      `json:"limit"`  // 每页显示的数量。
	Offset  int64                                      `json:"offset"` // 在响应中将被跳过的项目数量。例如，如果 `offset = 10`，则响应将从第 11 个找到的项目开始。
	SortDir V1AssemblyFbsProductListRequestSortDirEnum `json:"sort_dir"`
}

// 折扣或附加费用。
type MoneyMoneyCurrentTariffCharge struct {
	Currency string `json:"currency"` // 货币单位。
	Amount   string `json:"amount"`   // 金额。
}

// 买家信息。
type V3CustomerFbsLists struct {
	CustomerID    int64     `json:"customer_id"` // 买家ID。
	Name          string    `json:"name"`        // 买家姓名。
	Phone         string    `json:"phone"`       // 买家联系电话。始终返回空字符串 `""`。要获取替代联系电话，请使用方法 [/v3/posting/fbs/get](#operation/PostingAPI_GetFbsPostingV3)。
	Address       V3Address `json:"address"`
	CustomerEmail string    `json:"customer_email"` // 买家的电子邮箱地址。
}

type PostingV4PostingFbsListResponsePostingsProducts struct {
	Imei                []string          `json:"imei"`     // 移动设备的 IMEI 列表。
	OfferID             string            `json:"offer_id"` // 卖家系统中的商品标识符，即货号。
	Price               MoneyPostingMoney `json:"price"`
	ProductColor        string            `json:"product_color"`         // 商品颜色。
	Quantity            int32             `json:"quantity"`              // 货件中的商品数量。
	SKU                 int64             `json:"sku"`                   // Ozon系统中的商品标识符，即SKU。
	IsBlrTraceable      bool              `json:"is_blr_traceable"`      // `true`，表示商品可追溯。
	IsMarketplaceBuyout bool              `json:"is_marketplace_buyout"` // `true`，表示Ozon已购买该商品，用于配送至欧亚经济联盟及其他国家。
	Name                string            `json:"name"`                  // 商品名称。
	Weight              float64           `json:"weight"`                // 商品含包装重量。
}

type PostingBooleanResponse struct {
	Result bool `json:"result"` // 处理请求的结果。 如果请求执行时无误，则为“true”。
}

type V3FbsPostingProductExemplarInfoV3 struct {
	MandatoryMark string   `json:"mandatory_mark"` // 强制性标记“诚实标记”。
	GTD           string   `json:"gtd"`            // 货运报关单号（Cargo Customs Declaration）。
	IsGTDAbsent   bool     `json:"is_gtd_absent"`  // 未指出货运报关单号（Cargo Customs Declaration）的标志。
	Rnpt          string   `json:"rnpt"`           // 商品批次注册号（Product Batch Registration Number）。
	IsRnptAbsent  bool     `json:"is_rnpt_absent"` // 未指出商品批次注册号（Product Batch Registration Number）的标志。
	Imei          []string `json:"imei"`           // 移动设备的 IMEI 列表。
	ExemplarID    int64    `json:"exemplar_id"`    // 样件识别码。
}

// 商品和副本列表。
type V3FbsPostingExemplarProductV3 struct {
	Exemplars []V3FbsPostingProductExemplarInfoV3 `json:"exemplars"` // 按副本的信息。
	SKU       int64                               `json:"sku"`       // 在Ozon系统中的产品ID — SKU。
}

// 收件人联系方式。
type V3AddresseeFbsLists struct {
	Name  string `json:"name"`  // 收件人姓名。
	Phone string `json:"phone"` // 收件人联系电话。总是返回空字符串 `""`。要获取替代联系电话，请使用方法 [/v3/posting/fbs/get](#operation/PostingAPI_GetFbsPostingV3)。
}

type PostingMoveStatusResponseMoveStatus struct {
	Error         string `json:"error"`          // 处理请求时出错。
	PostingNumber string `json:"posting_number"` // 发货号。
	Result        bool   `json:"result"`         // 如果执行请求无误 — `true`。
}

// 收件人联系方式。
type PostingV4PostingFbsListResponsePostingsAddressee struct {
	Name string `json:"name"` // 收件人姓名。
}

// 货件最后一次更改过状态的时期。
type PostinglistV3status struct {
	From string `json:"from"` // 时期开始日期。
	To   string `json:"to"`   // 时期结束日期。
}

type V1AssemblyFbsPostingListResponsePostingProduct struct {
	OfferID     string `json:"offer_id"`     // 卖家系统中的商品标识符——货号。
	PictureURL  string `json:"picture_url"`  // 商品图片链接。
	ProductName string `json:"product_name"` // 商品名称。
	Quantity    int32  `json:"quantity"`     // 商品数量。
	SKU         int64  `json:"sku"`          // Ozon系统中的商品标识符——SKU。
}

type V1AssemblyFbsPostingListResponsePosting struct {
	AssemblyCode  string                                           `json:"assembly_code"`  // 拣货单代码。
	PostingNumber string                                           `json:"posting_number"` // 货件编号。
	Products      []V1AssemblyFbsPostingListResponsePostingProduct `json:"products"`       // 商品列表。
}

// 需要添加到响应中的附加字段。
type PostingV4PostingFbsListRequestWith struct {
	Barcodes      bool `json:"barcodes"`       // `true`，表示要在响应中添加货件条形码。
	FinancialData bool `json:"financial_data"` // 若为`true`，则在响应中添加财务数据。
	LegalInfo     bool `json:"legal_info"`     // 若为`true`，则在响应中添加法务信息。
	AnalyticsData bool `json:"analytics_data"` // 若为`true`，则在响应中添加分析数据。
}

// 取消原因。
// CancellationInitiator values
type CancellationInitiator string

const (
	CancellationInitiatorV    CancellationInitiator = "卖家"
	CancellationInitiatorV_1  CancellationInitiator = "客户" // 或`买家`
	CancellationInitiatorOzon CancellationInitiator = "Ozon"
	CancellationInitiatorV_2  CancellationInitiator = "系统"
	CancellationInitiatorV_3  CancellationInitiator = "配送服务"
)

// CancellationType values
type CancellationType string

const (
	CancellationTypeSeller   CancellationType = "seller"   // 卖家取消；
	CancellationTypeClient   CancellationType = "client"   // 或 `customer` — 买家取消；
	CancellationTypeOzon     CancellationType = "ozon"     // Ozon取消；
	CancellationTypeSystem   CancellationType = "system"   // 系统取消；
	CancellationTypeDelivery CancellationType = "delivery" // 配送服务取消
)

type V3Cancellation struct {
	CancellationInitiator    CancellationInitiator `json:"cancellation_initiator"`     // 取消货运的发起者： - `卖家`, - `客户` 或`买家`, - `Ozon`, - `系统`, - `配送服务`。
	CancellationType         CancellationType      `json:"cancellation_type"`          // 货运取消类型： - `seller` — 卖家取消； - `client` 或 `customer` — 买家取消； - `ozon` — Ozon取消； - `system`— 系统取消； - `delivery` — 配送服务取消。
	CancelledAfterShip       bool                  `json:"cancelled_after_ship"`       // 如果订单在装运后取消 — `true`。
	AffectCancellationRating bool                  `json:"affect_cancellation_rating"` // 如果取消影响买家排行 — `true`。
	CancelReason             string                `json:"cancel_reason"`              // 取消原因。
	CancelReasonID           int64                 `json:"cancel_reason_id"`           // 取消货运的原因ID。
}

// 买家的法务信息。
type PostingV4PostingFbsListResponsePostingsLegalInfo struct {
	Inn         string `json:"inn"`          // 税号。
	Kpp         string `json:"kpp"`          // 纳税人登记原因代码。
	CompanyName string `json:"company_name"` // 公司名称。
}

// 需要补充附加信息的商品。 要将货件转入下一状态，请传递： - 制造国； - 货运报关单编号； - 商品批次登记号； - “诚实标志”标记； - 其他标记； - 重量。
type PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements struct {
	ProductsRequiringCountry       []string `json:"products_requiring_country"`        // 需要传递制造国信息的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v2/posting/fbs/product/country/set](#operation/PostingAPI_SetCountryProductFbsPost...
	ProductsRequiringGTD           []string `json:"products_requiring_gtd"`            // 需要传递货运报关单编号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringImei          []string `json:"products_requiring_imei"`           // 需要传递IMEI的商品标识符列表。
	ProductsRequiringJwUin         []string `json:"products_requiring_jw_uin"`         // 需要传递珠宝制品唯一标识码的商品列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExemplarSe...
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"` // С需要传递“诚实标志”标识的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductEx...
	ProductsRequiringRnpt          []string `json:"products_requiring_rnpt"`           // 需要传递商品批次登记号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringWeight        []string `json:"products_requiring_weight"`         // 需要传递重量的商品列表。
	ProductsRequiringChangeCountry []string `json:"products_requiring_change_country"` // 需要更改制造国的商品标识符（SKU）列表。要更改制造国，请使用方法[/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountryProductFbsPosti...
}

type PostingPostingFBSPackageLabelRequest struct {
	PostingNumber []string `json:"posting_number"` // 货运ID。
}

// 是否可以取消。
type CarriageCarriageGetResponseCancelAvailability struct {
	IsCancelAvailable bool   `json:"is_cancel_available"` // `true`, 如果运输可以取消。
	Reason            string `json:"reason"`              // 运输无法取消的原因。
}

type V2PostingFBSActGetProducts struct {
	SKU      int64  `json:"sku"`      // 商品在Ozon系统中的标识符——SKU。
	Name     string `json:"name"`     // 商品名称。
	OfferID  string `json:"offer_id"` // 商品在卖家系统中的标识符——货号。
	Price    string `json:"price"`    // 商品价格。
	Quantity int32  `json:"quantity"` // 货件中的商品数量。
}

type V2PostingFBSActGetPostingsResult struct {
	SellerError   string                       `json:"seller_error"`   // 错误代码说明。
	UpdatedAt     string                       `json:"updated_at"`     // 货件记录的更新日期和时间。
	CreatedAt     string                       `json:"created_at"`     // 货件记录的创建日期和时间。
	Products      []V2PostingFBSActGetProducts `json:"products"`       // 货件中商品列表。
	ID            int64                        `json:"id"`             // 单据标识符。
	MultiBoxQty   int32                        `json:"multi_box_qty"`  // 商品包装所用箱数。
	PostingNumber string                       `json:"posting_number"` // 货件编号。
	Status        string                       `json:"status"`         // 货件状态。
}

type V2PostingFBSActGetPostingsResponse struct {
	Result []V2PostingFBSActGetPostingsResult `json:"result"` // 货件信息。
}

// 附加信息。
type PostingShipV4RequestWith struct {
	AdditionalData bool `json:"additional_data"` // 为获取附加信息，请点击 `true`。
}

// 最低折扣或附加费用。
type MoneyMoneyCurrentTariffMinCharge struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 在`next_tariff_starts_at`参数指定的时间后生效的最低折扣或加价。
type MoneyMoneyNextTariffMinCharge struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 在`next_tariff_starts_at`参数指定的时间后生效的折扣或加价。
type MoneyMoneyNextTariffCharge struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 发运的计费信息。
type PostingV4PostingFbsListResponsePostingsTariffication struct {
	NextTariffRate         float64                          `json:"next_tariff_rate"`      // 在`next_tariff_starts_at`参数指定时间之后，该货件将按此百分比计费。
	NextTariffStartsAt     string                           `json:"next_tariff_starts_at"` // 新费率开始生效的日期和时间。
	NextTariffType         string                           `json:"next_tariff_type"`      // 在`next_tariff_starts_at`参数指定时间之后的计费类型——折扣或加价。
	CurrentTariffRate      float64                          `json:"current_tariff_rate"`   // 计费百分比。
	NextTariffCharge       MoneyMoneyNextTariffCharge       `json:"next_tariff_charge"`
	CurrentTariffCharge    MoneyMoneyCurrentTariffCharge    `json:"current_tariff_charge"`
	CurrentTariffMinCharge MoneyMoneyCurrentTariffMinCharge `json:"current_tariff_min_charge"`
	CurrentTariffType      string                           `json:"current_tariff_type"` // 计费类型——折扣或加价。
	NextTariffMinCharge    MoneyMoneyNextTariffMinCharge    `json:"next_tariff_min_charge"`
}

// 需要补充附加信息的商品。 要将货件转入下一状态，请传递： - 制造国； - 货运报关单编号； - 商品批次登记号； - “诚实标志”标记； - 其他标记； - 重量。
type PostingV4PostingFbsListResponsePostingsRequirements struct {
	ProductsRequiringJwUin         []string `json:"products_requiring_jw_uin"`         // 需要传递珠宝制品唯一标识码的商品列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExemplarSe...
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"` // С需要传递“诚实标志”标识的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductEx...
	ProductsRequiringRnpt          []string `json:"products_requiring_rnpt"`           // 需要传递商品批次登记号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringWeight        []string `json:"products_requiring_weight"`         // 需要传递重量的商品列表。
	ProductsRequiringChangeCountry []string `json:"products_requiring_change_country"` // 需要更改制造国的商品标识符（SKU）列表。要更改制造国，请使用方法[/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountryProductFbsPosti...
	ProductsRequiringCountry       []string `json:"products_requiring_country"`        // 需要传递制造国信息的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v2/posting/fbs/product/country/set](#operation/PostingAPI_SetCountryProductFbsPost...
	ProductsRequiringGTD           []string `json:"products_requiring_gtd"`            // 需要传递货运报关单编号的商品标识符（SKU）列表。 在备货货件前，请通过方法[/v6/fbs/posting/product/exemplar/set](#operation/PostingAPI_FbsPostingProductExem...
	ProductsRequiringImei          []string `json:"products_requiring_imei"`           // 需要传递IMEI的商品标识符列表。
}

type PostingV4PostingFbsListResponsePostingsTarifficationStep struct {
	MinCharge        MoneyMoneyCurrentTariffMinCharge `json:"min_charge"`
	TariffCharge     MoneyMoneyCurrentTariffCharge    `json:"tariff_charge"`
	TariffDeadlineAt string                           `json:"tariff_deadline_at"` // 计费阶段结束的日期和时间。该日期后将自动进入下一计费阶段。
	TariffRate       float64                          `json:"tariff_rate"`        // 折扣或附加费用百分比。
	TariffType       string                           `json:"tariff_type"`        // 计费类型。
}

// 配送地址。
type PostingV4PostingFbsListResponsePostingsCustomerAddress struct {
	AddressTail     string  `json:"address_tail"`      // 文本格式地址。
	District        string  `json:"district"`          // 配送区域。
	ProviderPvzCode string  `json:"provider_pvz_code"` // 3PL服务商取货点代码。
	PvzCode         int64   `json:"pvz_code"`          // 取货点代码。
	City            string  `json:"city"`              // 配送城市。
	Comment         string  `json:"comment"`           // 订单备注。
	Country         string  `json:"country"`           // 配送国家。
	Latitude        float64 `json:"latitude"`          // 纬度。
	Longitude       float64 `json:"longitude"`         // 经度。
	Region          string  `json:"region"`            // 配送地区。
	ZipCode         string  `json:"zip_code"`          // 收件人邮政编码。
}

// 买家信息。
type PostingV4PostingFbsListResponsePostingsCustomer struct {
	Phone         string                                                 `json:"phone"` // 买家的替换联系电话。
	Address       PostingV4PostingFbsListResponsePostingsCustomerAddress `json:"address"`
	CustomerEmail string                                                 `json:"customer_email"` // 买家电子邮箱。
	CustomerID    int64                                                  `json:"customer_id"`    // 买家标识符。
	Name          string                                                 `json:"name"`           // 买家姓名。
}

// 带有附加特征的商品列表。
type PostingV4PostingFbsListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []string `json:"products_with_possible_mandatory_mark"` // 可能需要标志的商品列表。
}

// 货件条形码。
type PostingV4PostingFbsListResponsePostingsBarcodes struct {
	LowerBarcode string `json:"lower_barcode"` // 货件标签上的下方条形码。
	UpperBarcode string `json:"upper_barcode"` // 货件标签上的上方条形码。
}

// 取消信息。
type PostingV4PostingFbsListResponsePostingsCancellation struct {
	CancelReasonID           int64                 `json:"cancel_reason_id"`           // 货件取消原因标识符。
	CancellationInitiator    CancellationInitiator `json:"cancellation_initiator"`     // 取消发起方： - `卖家`， - `客户`， - `买家`， - `Ozon`， - `系统`， - `配送服务商`。
	CancellationType         CancellationType      `json:"cancellation_type"`          // 取消类型： - `seller`——由卖家取消； - `client`或`customer`——由买家取消； - `ozon`——由Ozon取消； - `system`——由系统取消； - `delivery`——由配送服务取消。
	CancelledAfterShip       bool                  `json:"cancelled_after_ship"`       // `true`，表示取消发生在货件完成备货之后。
	AffectCancellationRating bool                  `json:"affect_cancellation_rating"` // `true`，表示取消会影响卖家评级。
	CancelReason             string                `json:"cancel_reason"`              // 取消原因。
}

// 商品佣金。
type PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission struct {
	Amount   float64 `json:"amount"`   // 商品佣金金额。
	Currency string  `json:"currency"` // 计算佣金所使用的货币代码。
	Percent  int64   `json:"percent"`  // 佣金比例。
}

type PostingV4PostingFbsListResponsePostingsFinancialDataProducts struct {
	OldPrice             float64                                                                `json:"old_price"`              // 折扣前的价格。商品卡上会以划线价显示。
	TotalDiscountValue   float64                                                                `json:"total_discount_value"`   // 折扣金额。
	Actions              []string                                                               `json:"actions"`                // 促销活动列表。
	Payout               float64                                                                `json:"payout"`                 // 卖家应收款。
	Price                float64                                                                `json:"price"`                  // 计入促销活动后的商品价格，不包括由Ozon承担费用的促销活动。
	ProductID            int64                                                                  `json:"product_id"`             // Ozon系统中的商品标识符，即SKU。
	Quantity             int64                                                                  `json:"quantity"`               // 货件中的商品数量。
	TotalDiscountPercent float64                                                                `json:"total_discount_percent"` // 折扣百分比。
	Commission           PostingV4PostingFbsListResponsePostingsFinancialDataProductsCommission `json:"commission"`
	CustomerPrice        MoneyPostingMoney                                                      `json:"customer_price"`
}

// 商品价格、折扣金额、卖家应收款和佣金信息。
type PostingV4PostingFbsListResponsePostingsFinancialData struct {
	ClusterTo   string                                                         `json:"cluster_to"`   // 订单配送地区代码。
	Products    []PostingV4PostingFbsListResponsePostingsFinancialDataProducts `json:"products"`     // 订单中的商品列表。
	ClusterFrom string                                                         `json:"cluster_from"` // 订单发出地区代码。
}

// 分析数据。
// PaymentTypeGroupName values
type PaymentTypeGroupName string

const (
	PaymentTypeGroupNameV        PaymentTypeGroupName = "在线银行卡支付"       // ；
	PaymentTypeGroupNameOzonBank PaymentTypeGroupName = "Ozon Bank银行卡"  // ；
	PaymentTypeGroupNameOzon     PaymentTypeGroupName = "取货时从Ozon卡自动扣款" // ；
	PaymentTypeGroupNameV_1      PaymentTypeGroupName = "收货时用已保存的银行卡支付" // ；
	PaymentTypeGroupNameV_2      PaymentTypeGroupName = "快速支付系统"        // ；
	PaymentTypeGroupNameOzon_1   PaymentTypeGroupName = "Ozon分期付款"      // ；
	PaymentTypeGroupNameV_3      PaymentTypeGroupName = "支付至结算账户"       // ；
	PaymentTypeGroupNameSberPay  PaymentTypeGroupName = "SberPay"       // ；
	PaymentTypeGroupNameV_4      PaymentTypeGroupName = "通过外部卖家预付款"
)

type PostingV4PostingFbsListResponsePostingsAnalyticsData struct {
	DeliveryDateEnd         string               `json:"delivery_date_end"`          // 配送结束日期和时间。
	DeliveryType            string               `json:"delivery_type"`              // 配送方法。
	TPLProvider             string               `json:"tpl_provider"`               // 配送服务。
	Warehouse               string               `json:"warehouse"`                  // 订单发货仓库名称。
	WarehouseID             int64                `json:"warehouse_id"`               // 仓库标识符。
	City                    string               `json:"city"`                       // 配送城市。仅适用于rFBS货件和独联体卖家。
	ClientDeliveryDateBegin string               `json:"client_delivery_date_begin"` // 配送开始的日期和时间。仅适用于通过Ozon配送创建的货件。
	IsLegal                 bool                 `json:"is_legal"`                   // `true`，表示收件人为法人。
	IsPremium               bool                 `json:"is_premium"`                 // `true`，表示收件人有Premium订阅。
	PaymentTypeGroupName    PaymentTypeGroupName `json:"payment_type_group_name"`    // 支付方法： - `在线银行卡支付`； - `Ozon Bank银行卡`； - `取货时从Ozon卡自动扣款`； - `收货时用已保存的银行卡支付`； - `快速支付系统`； - `Ozon分期付款`； - `支付至结算账户`； - `Sbe...
	Region                  string               `json:"region"`                     // 配送地区。仅适用于rFBS货件。
	TPLProviderID           int64                `json:"tpl_provider_id"`            // 配送服务标识符。
	ClientDeliveryDateEnd   string               `json:"client_delivery_date_end"`   // 订单预计最晚送达日期。仅适用于通过Ozon配送创建的货件。
	DeliveryDateBegin       string               `json:"delivery_date_begin"`        // 配送开始的日期和时间。
}

// 外部平台订单信息。
type PostingV4PostingFbsListResponsePostingsExternalOrder struct {
	IsExternal   bool   `json:"is_external"`   // `true`，表示订单来自外部平台。
	PlatformName string `json:"platform_name"` // 下单平台名称。
}

// Substatus values
type Substatus string

const (
	SubstatusPostingAcceptanceInProgress        Substatus = "posting_acceptance_in_progress"         // 验收中；
	SubstatusPostingInArbitration               Substatus = "posting_in_arbitration"                 // 仲裁；
	SubstatusPostingCreated                     Substatus = "posting_created"                        // 已创建；
	SubstatusPostingInCarriage                  Substatus = "posting_in_carriage"                    // 在运输中；
	SubstatusPostingNotInCarriage               Substatus = "posting_not_in_carriage"                // 未添加到运输；
	SubstatusPostingRegistered                  Substatus = "posting_registered"                     // 已登记；
	SubstatusPostingTransferringToDelivery      Substatus = "posting_transferring_to_delivery"       // ，当`status=awaiting_deliver`时——正在运送到物流公司；
	SubstatusPostingAwaitingPassportData        Substatus = "posting_awaiting_passport_data"         // 等待护照信息；
	SubstatusPostingAwaitingRegistration        Substatus = "posting_awaiting_registration"          // 等待登记；
	SubstatusPostingRegistrationError           Substatus = "posting_registration_error"             // 登记错误；
	SubstatusPostingSplitPending                Substatus = "posting_split_pending"                  // 已创建；
	SubstatusPostingCanceled                    Substatus = "posting_canceled"                       // 已取消；
	SubstatusPostingInClientArbitration         Substatus = "posting_in_client_arbitration"          // 客户配送仲裁；
	SubstatusPostingDelivered                   Substatus = "posting_delivered"                      // 已配送；
	SubstatusPostingReceived                    Substatus = "posting_received"                       // 已收到；
	SubstatusPostingConditionallyDelivered      Substatus = "posting_conditionally_delivered"        // 可能已收；
	SubstatusPostingInCourierService            Substatus = "posting_in_courier_service"             // 快递员派件中；
	SubstatusPostingInPickupPoint               Substatus = "posting_in_pickup_point"                // 在取货点；
	SubstatusPostingOnWayToCity                 Substatus = "posting_on_way_to_city"                 // 正在前往您的城市；
	SubstatusPostingOnWayToPickupPoint          Substatus = "posting_on_way_to_pickup_point"         // 正在前往取货点；
	SubstatusPostingReturnedToWarehouse         Substatus = "posting_returned_to_warehouse"          // 已退回仓库；
	SubstatusPostingTransferredToCourierService Substatus = "posting_transferred_to_courier_service" // 正在交给配送服务；
	SubstatusPostingDriverPickUp                Substatus = "posting_driver_pick_up"                 // 司机正在送货；
	SubstatusPostingNotInSortCenter             Substatus = "posting_not_in_sort_center"             // 未被分拣中心接收；
	SubstatusShipFailed                         Substatus = "ship_failed"                            // 货件备货失败
)

// Status values
type Status string

const (
	StatusAcceptanceInProgress      Status = "acceptance_in_progress"       // 接收中；
	StatusArbitration               Status = "arbitration"                  // 仲裁；
	StatusAwaitingApprove           Status = "awaiting_approve"             // 等待确认；
	StatusAwaitingDeliver           Status = "awaiting_deliver"             // 等待发运；
	StatusAwaitingPackaging         Status = "awaiting_packaging"           // 等待包装；
	StatusAwaitingRegistration      Status = "awaiting_registration"        // 等待登记；
	StatusAwaitingVerification      Status = "awaiting_verification"        // 已创建；
	StatusCancelled                 Status = "cancelled"                    // 已取消；
	StatusCancelledFromSplitPending Status = "cancelled_from_split_pending" // 因货件拆分而取消；
	StatusClientArbitration         Status = "client_arbitration"           // 客户配送仲裁；
	StatusDelivering                Status = "delivering"                   // 运输中；
	StatusDriverPickup              Status = "driver_pickup"                // 司机正在送货；
	StatusNotAccepted               Status = "not_accepted"                 // 在分拣中心尚未接收
)

// TPLIntegrationType values
type TPLIntegrationType string

const (
	TPLIntegrationTypeOzon          TPLIntegrationType = "ozon"           // 由Ozon配送；
	TPLIntegrationTypeV3plTracking  TPLIntegrationType = "3pl_tracking"   // 由已集成服务配送；
	TPLIntegrationTypeNonIntegrated TPLIntegrationType = "non_integrated" // 由第三方服务配送；
	TPLIntegrationTypeAggregator    TPLIntegrationType = "aggregator"     // 通过Ozon合作伙伴配送；
	TPLIntegrationTypeHybryd        TPLIntegrationType = "hybryd"         // 俄罗斯邮政配送方案
)

// AvailableActions values
type AvailableActions string

const (
	AvailableActionsArbitration            AvailableActions = "arbitration"               // 提出争议；
	AvailableActionsAwaitingDelivery       AvailableActions = "awaiting_delivery"         // 变更为“等待发运”状态；
	AvailableActionsCanCreateChat          AvailableActions = "can_create_chat"           // 与买家发起聊天；
	AvailableActionsCancel                 AvailableActions = "cancel"                    // 取消货件；
	AvailableActionsClickTrackNumber       AvailableActions = "click_track_number"        // 在个人中心通过追踪号码查看状态变更历史；
	AvailableActionsCustomerPhoneAvailable AvailableActions = "customer_phone_available"  // 买家 电话；
	AvailableActionsHasWeightProducts      AvailableActions = "has_weight_products"       // 货件中包含称重商品；
	AvailableActionsHideRegionAndCity      AvailableActions = "hide_region_and_city"      // 在报告中隐藏买家的地区和城市；
	AvailableActionsInvoiceGet             AvailableActions = "invoice_get"               // 获取发票信息；
	AvailableActionsInvoiceSend            AvailableActions = "invoice_send"              // 创建发票；
	AvailableActionsInvoiceUpdate          AvailableActions = "invoice_update"            // 编辑发票；
	AvailableActionsLabelDownloadBig       AvailableActions = "label_download_big"        // 下载大标签；
	AvailableActionsLabelDownloadSmall     AvailableActions = "label_download_small"      // 下载小标签；
	AvailableActionsLabelDownload          AvailableActions = "label_download"            // 下载标签；
	AvailableActionsNonIntDelivered        AvailableActions = "non_int_delivered"         // 变更为“可能已收”状态；
	AvailableActionsNonIntDelivering       AvailableActions = "non_int_delivering"        // 变更为“配送中”状态；
	AvailableActionsNonIntLastMile         AvailableActions = "non_int_last_mile"         // 变更为“快递员派件中”状态；
	AvailableActionsProductCancel          AvailableActions = "product_cancel"            // 取消货件中的部分商品；
	AvailableActionsSetCutoff              AvailableActions = "set_cutoff"                // 请通过方法[/v1/posting/cutoff/set](#operation/PostingAPI_SetPostingCutoff)设置发运日期，且不得晚于`shipment_date`参数中的日期；
	AvailableActionsSetTimeslot            AvailableActions = "set_timeslot"              // 修改给买家的配送时间；
	AvailableActionsSetTrackNumber         AvailableActions = "set_track_number"          // 填写或修改追踪号码；
	AvailableActionsShipAsyncInProcess     AvailableActions = "ship_async_in_process"     // 货件正在备货；
	AvailableActionsShipAsyncRetry         AvailableActions = "ship_async_retry"          // 在备货货件出错后重新备货货件；
	AvailableActionsShipAsync              AvailableActions = "ship_async"                // 备货货件；
	AvailableActionsShipWithAdditionalInfo AvailableActions = "ship_with_additional_info" // 请通过方法[/v6/fbs/posting/product/exemplar/set](https://docs.ozon.ru/api/seller/#operation/PostingAPI_FbsPostingProductExemp...
	AvailableActionsShip                   AvailableActions = "ship"                      // 备货货件；
	AvailableActionsUpdateCis              AvailableActions = "update_cis"                // 修改附加信息
)

// PrrOption values
type PrrOption string

const (
	PrrOptionLift            PrrOption = "lift"             // 通过电梯搬运上楼；
	PrrOptionStairs          PrrOption = "stairs"           // 通过楼梯搬运上楼；
	PrrOptionNone            PrrOption = "none"             // 买家拒绝该服务，无需将商品搬上楼；
	PrrOptionDeliveryDefault PrrOption = "delivery_default" // 配送费用已包含在价格中，根据要约条款需将商品送至楼层
)

// DeliverySchema values
type DeliverySchema string

const (
	DeliverySchemaSDS         DeliverySchema = "SDS"         // 统一SKU标识符；
	DeliverySchemaFBO         DeliverySchema = "FBO"         // 从Ozon仓库销售的商品标识符；
	DeliverySchemaFBS         DeliverySchema = "FBS"         // 从FBS仓库销售的商品标识符；
	DeliverySchemaCrossborder DeliverySchema = "Crossborder" // 从境外销售的商品标识符
)

type PostingV4PostingFbsListResponsePostings struct {
	DestinationPlaceName     string                                                     `json:"destination_place_name"` // 目的地名称。
	OrderNumber              string                                                     `json:"order_number"`           // 该货件所属订单的编号。
	Requirements             PostingV4PostingFbsListResponsePostingsRequirements        `json:"requirements"`
	Substatus                Substatus                                                  `json:"substatus"`           // 货件子状态： - `posting_acceptance_in_progress`——验收中； - `posting_in_arbitration`——仲裁； - `posting_created`——已创建； - `posting_in_...
	TarifficationSteps       []PostingV4PostingFbsListResponsePostingsTarifficationStep `json:"tariffication_steps"` // 计费阶段。
	InProcessAt              string                                                     `json:"in_process_at"`       // 货件开始处理的日期和时间。
	Barcodes                 PostingV4PostingFbsListResponsePostingsBarcodes            `json:"barcodes"`
	Cancellation             PostingV4PostingFbsListResponsePostingsCancellation        `json:"cancellation"`
	DeliveryMethod           PostingV4PostingFbsListResponsePostingsDeliveryMethod      `json:"delivery_method"`
	IsExpress                bool                                                       `json:"is_express"`           // `true`，表示使用了Ozon Express快速配送。
	OrderID                  int64                                                      `json:"order_id"`             // 该货件所属订单的标识符。
	Status                   Status                                                     `json:"status"`               // 货件状态： - `acceptance_in_progress`——接收中； - `arbitration`——仲裁； - `awaiting_approve`——等待确认； - `awaiting_deliver`——等待发运； - `a...
	TPLIntegrationType       TPLIntegrationType                                         `json:"tpl_integration_type"` // 与配送服务的集成类型： - `ozon`——由Ozon配送； - `3pl_tracking`——由已集成服务配送； - `non_integrated`——由第三方服务配送； - `aggregator`——通过Ozon合作伙伴配送； -...
	FinancialData            PostingV4PostingFbsListResponsePostingsFinancialData       `json:"financial_data"`
	Addressee                PostingV4PostingFbsListResponsePostingsAddressee           `json:"addressee"`
	AvailableActions         []string                                                   `json:"available_actions"` // 货件的可用操作和信息： - `arbitration`——提出争议； - `awaiting_delivery`——变更为“等待发运”状态； - `can_create_chat`——与买家发起聊天； - `cancel`——取消货件； -...
	Customer                 PostingV4PostingFbsListResponsePostingsCustomer            `json:"customer"`
	MultiBoxQty              int32                                                      `json:"multi_box_qty"`               // 商品包装所用箱数。
	PickupCodeVerifiedAt     string                                                     `json:"pickup_code_verified_at"`     // 快递员代码校验成功的日期和时间。请通过方法 [/v1/posting/fbs/pick-up-code/verify](#operation/PostingAPI_PostingFBSPickupCodeVerify)校验快递员代码。
	RequireBlrTraceableAttrs bool                                                       `json:"require_blr_traceable_attrs"` // `true`，表示需要填写可追踪属性。
	DeliveringDate           string                                                     `json:"delivering_date"`             // 货件转移配送的日期。
	IsMultibox               bool                                                       `json:"is_multibox"`                 // 标记货件中是否包含多箱商品，以及是否需要传递其箱数： - `true`——备货前，请通过方法[/v3/posting/multiboxqty/set](#operation/PostingAPI_PostingMultiBoxQtySetV...
	AnalyticsData            PostingV4PostingFbsListResponsePostingsAnalyticsData       `json:"analytics_data"`
	LegalInfo                PostingV4PostingFbsListResponsePostingsLegalInfo           `json:"legal_info"`
	ShipmentDateWithoutDelay string                                                     `json:"shipment_date_without_delay"` // 不逾期发运的日期和时间。
	DestinationPlaceID       int64                                                      `json:"destination_place_id"`        // 目的地标识符。
	ExternalOrder            PostingV4PostingFbsListResponsePostingsExternalOrder       `json:"external_order"`
	Optional                 PostingV4PostingFbsListResponsePostingsOptional            `json:"optional"`
	PostingNumber            string                                                     `json:"posting_number"`        // 货件编号。
	Products                 []PostingV4PostingFbsListResponsePostingsProducts          `json:"products"`              // 货件中商品列表。
	PrrOption                PrrOption                                                  `json:"prr_option"`            // 装卸服务代码： - `lift`——通过电梯搬运上楼； - `stairs`——通过楼梯搬运上楼； - `none`——买家拒绝该服务，无需将商品搬上楼； - `delivery_default`——配送费用已包含在价格中，根据要约条款需将...
	QuantumID                int64                                                      `json:"quantum_id"`            // 经济商品标识符。
	ShipmentDate             string                                                     `json:"shipment_date"`         // 需要在此日期和时间之前完成货件备货。系统会显示推荐的发运时间。超过该时间后将开始适用新费率，相关信息可在`tariffication`字段中获取。
	DeliverySchema           DeliverySchema                                             `json:"delivery_schema"`       // 配送模式： - `SDS`——统一SKU标识符； - `FBO`——从Ozon仓库销售的商品标识符； - `FBS`——从FBS仓库销售的商品标识符； - `Crossborder`——从境外销售的商品标识符。
	IsClickAndCollect        bool                                                       `json:"is_click_and_collect"`  // `true`，表示货件通过“到店自提”方式配送。
	IsPresortable            bool                                                       `json:"is_presortable"`        // `true`，表示商品为错误分级的商品。
	ParentPostingNumber      string                                                     `json:"parent_posting_number"` // 当前货件所拆分自的父货件编号。
	Tariffication            PostingV4PostingFbsListResponsePostingsTariffication       `json:"tariffication"`
	TrackingNumber           string                                                     `json:"tracking_number"` // 货件跟踪号码。
	VolumeWeight             float64                                                    `json:"volume_weight"`   // 商品体积重量。
}

type PostingV4PostingFbsListResponse struct {
	Cursor   string                                    `json:"cursor"`   // 用于选择下一批数据的指针。
	HasNext  bool                                      `json:"has_next"` // 若响应中未返回全部货件，则为`true`。
	Postings []PostingV4PostingFbsListResponsePostings `json:"postings"` // 货件列表。
}

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type PostingV4PostingFbsUnfulfilledListRequestSortDirEnum string

// 买家的法务信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo struct {
	CompanyName string `json:"company_name"` // 公司名称。
	Inn         string `json:"inn"`          // 税号。
	Kpp         string `json:"kpp"`          // 纳税人登记原因代码。
}

type V2FbsPostingProductCountryListResponseResult struct {
	CountryIsoCode string `json:"country_iso_code"` // ISO国家代码。
	Name           string `json:"name"`             // 国家俄语名称
}

type V2FbsPostingProductCountryListRequest struct {
	NameSearch string `json:"name_search"` // 按行过滤。
}

type PostingProductCancelRequestItem struct {
	Quantity int32 `json:"quantity"` // 货运商品数量。
	SKU      int64 `json:"sku"`      // Ozon系统中的商品ID — SKU。
}

// 发运的计费信息。
type V3FbsTariffication struct {
	NextTariffStartsAt              string  `json:"next_tariff_starts_at"`               // 新的费率开始生效的日期和时间。 格式：`YYYY-MM-DDThh:mm:ss.mcsZ`. 示例：`2023-11-13T08:05:57.657Z`.
	NextTariffRate                  float64 `json:"next_tariff_rate"`                    // 在参数 `next_tariff_starts_at` 指定的时间后，将按此百分比进行计费。
	NextTariffType                  string  `json:"next_tariff_type"`                    // 在参数 `next_tariff_starts_at` 指定的时间后，将按此类型计费 — 折扣或附加费。
	NextTariffCharge                string  `json:"next_tariff_charge"`                  // 下一步计费中的折扣或附加金额。
	NextTariffChargeCurrencyCode    string  `json:"next_tariff_charge_currency_code"`    // 新费率的货币单位。
	CurrentTariffRate               float64 `json:"current_tariff_rate"`                 // 当前运费的百分比。
	CurrentTariffType               string  `json:"current_tariff_type"`                 // 当前的计费类型 — 折扣或附加费。
	CurrentTariffCharge             string  `json:"current_tariff_charge"`               // 当前的折扣或附加费金额。
	CurrentTariffChargeCurrencyCode string  `json:"current_tariff_charge_currency_code"` // 金额的货币单位。
}

// 过滤器。
// FbpFilter values
type FbpFilter string

const (
	FbpFilterALL     FbpFilter = "ALL"     // 响应中将显示所有符合其他筛选器条件的货件；
	FbpFilterOnly    FbpFilter = "ONLY"    // 仅显示FBP货件；
	FbpFilterWithout FbpFilter = "WITHOUT" // 显示除FBP外的所有货件
)

type Postingv3GetFbsPostingListRequestFilter struct {
	ProviderID            []int64             `json:"provider_id"`        // 快递服务ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	Since                 string              `json:"since"`              // 应收到货件清单时间段的开始日期。 UTC模式: ГГГГ-ММ-ДДTЧЧ:ММ:ССZ. 例子: 2019-08-24T14:15:22Z.
	To                    string              `json:"to"`                 // 应收到货件清单时间段的结束日期。 UTC模式： ГГГГ-ММ-ДДTЧЧ:ММ:ССZ. 例子： 2019-08-24T14:15:22Z.
	Status                Status              `json:"status"`             // 货件运输状态： - `awaiting_registration` — 等待注册， - `acceptance_in_progress` — 正在验收， - `awaiting_approve` — 等待确认， - `awaiting_pa...
	WarehouseID           []string            `json:"warehouse_id"`       // 仓库ID。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	DeliveryMethodID      []int64             `json:"delivery_method_id"` // 快递方式ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	OrderID               int64               `json:"order_id"`           // 订单ID。
	LastChangedStatusDate PostinglistV3status `json:"last_changed_status_date"`
	FbpFilter             FbpFilter           `json:"fbpFilter"` // 从合作伙伴仓库（FBP）发货时的货件筛选器： - `ALL` — 响应中将显示所有符合其他筛选器条件的货件； - `ONLY` — 仅显示FBP货件； - `WITHOUT` — 显示除FBP外的所有货件。 默认值为 `ALL`。
}

// 要添加到响应的附加字段。
type Postingv3FbsPostingWithParams struct {
	FinancialData bool `json:"financial_data"` // 将财务数据添加到响应中。
	LegalInfo     bool `json:"legal_info"`     // 将法律信息添加到响应中。
	Translit      bool `json:"translit"`       // 完成返回值的拼写。
	AnalyticsData bool `json:"analytics_data"` // 将分析数据添加到响应中。
	Barcodes      bool `json:"barcodes"`       // 将货件条形码添加到响应中。
}

// Dir values
type Dir string

const (
	DirAsc  Dir = "asc"  // 从小到大
	DirDesc Dir = "desc" // 从大到小
)

type Postingv3GetFbsPostingListRequest struct {
	With   Postingv3FbsPostingWithParams           `json:"with"`
	Dir    Dir                                     `json:"dir"` // 分类方向： - `asc` — 从小到大， - `desc` — 从大到小。
	Filter Postingv3GetFbsPostingListRequestFilter `json:"filter"`
	Limit  int64                                   `json:"limit"`  // 响应中值的数量： - 最大值 — 1000, - 最小值 — 1。
	Offset int64                                   `json:"offset"` // 将在响应中跳过的元素数。 例如，如果“offset=10”，那么响应将从找到的第11个元素开始。
}

type V2FbsPostingProductCountrySetResponse struct {
	ProductID   int64 `json:"product_id"`    // 产品ID。
	IsGTDNeeded bool  `json:"is_gtd_needed"` // 有必要传送产品和货运的货物报关单（Cargo Customs Declaration）编号的标志。
}

// 买方的法律信息。
type V2FboSinglePostingLegalInfo struct {
	CompanyName string `json:"company_name"` // 公司名称。
	Inn         string `json:"inn"`          // 纳税人识别号（INN）。
	Kpp         string `json:"kpp"`          // 税务登记原因代码（KPP）。
}

// 货件条码。
type V3Barcodes struct {
	LowerBarcode string `json:"lower_barcode"` // 货件标签的下条码。
	UpperBarcode string `json:"upper_barcode"` // 货件标签的上条码。
}

// 分析数据。
type V3FbsPostingAnalyticsData struct {
	DeliveryDateBegin       string               `json:"delivery_date_begin"`        // 快递开始日期和时间。
	DeliveryType            string               `json:"delivery_type"`              // 快递方式。
	IsLegal                 bool                 `json:"is_legal"`                   // 收件人是法人的标志： - `true` — 法人， - `false` — 自然人。
	Warehouse               string               `json:"warehouse"`                  // 订单发送仓库名称。
	ClientDeliveryDateBegin string               `json:"client_delivery_date_begin"` // 向客户开始配送的日期和时间。仅适用于通过Ozon配送创建的货件。
	ClientDeliveryDateEnd   string               `json:"client_delivery_date_end"`   // 订单将送达客户的截止日期。仅适用于通过Ozon配送创建的货件。
	City                    string               `json:"city"`                       // 快递城市。仅适用于rFBS货件和独联体卖家。
	DeliveryDateEnd         string               `json:"delivery_date_end"`          // 快递结束日期和时间。
	IsPremium               bool                 `json:"is_premium"`                 // 有Premium订阅。
	PaymentTypeGroupName    PaymentTypeGroupName `json:"payment_type_group_name"`    // 付款方法： - `在线银行卡支付`， - `Ozon银行卡`， - `取货时自动从Ozon银行卡收费`， - `收货时从已保存的银行卡收费`， - `快速支付系统`， - `Ozon分期付款`， - `支付至结算账户`， - `SberPa...
	Region                  string               `json:"region"`                     // 快递地区。
	TPLProvider             string               `json:"tpl_provider"`               // 快递服务。
	TPLProviderID           int64                `json:"tpl_provider_id"`            // 快递服务ID。
	WarehouseID             int64                `json:"warehouse_id"`               // 仓库ID。
}

// 带有附加特征的商品列表。
type V3FbsPostingDetailOptional struct {
	ProductsWithPossibleMandatoryMark []any `json:"products_with_possible_mandatory_mark"` // 带有可能标志的商品列表。
}

type PostingFinancialDataProduct struct {
	Actions                 []string     `json:"actions"`                   // 活动清单。
	CurrencyCode            CurrencyCode `json:"currency_code"`             // 价格货币，其与个人中心中设置的币种相匹配。 可能的值： - `RUB` — 俄罗斯卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
	CustomerCurrencyCode    string       `json:"customer_currency_code"`    // 买家货币代码。
	CommissionAmount        float64      `json:"commission_amount"`         // 商品佣金大小。
	OldPrice                float64      `json:"old_price"`                 // 打折前价格。在商品卡片上将被显示划掉。
	Payout                  float64      `json:"payout"`                    // 支付给卖方。
	TotalDiscountPercent    float64      `json:"total_discount_percent"`    // 折扣百分比。
	CommissionPercent       int64        `json:"commission_percent"`        // 佣金百分比。
	CommissionsCurrencyCode string       `json:"commissions_currency_code"` // 计算佣金的币种代码。
	Price                   float64      `json:"price"`                     // 您的价格。包含卖家促销（如有），不含 Ozon 资助的促销。
	CustomerPrice           float64      `json:"customer_price"`            // 包含卖家与 Ozon 折扣的买家价格。
	ProductID               int64        `json:"product_id"`                // Ozon系统中的商品标识符——SKU。
	Quantity                int64        `json:"quantity"`                  // 运输商品数量。
	TotalDiscountValue      float64      `json:"total_discount_value"`      // 折扣数量。
}

// 有关商品成本、折扣幅度、付款和佣金的信息。
type V3PostingFinancialData struct {
	ClusterFrom string                        `json:"cluster_from"` // 订单发送区域代码。
	ClusterTo   string                        `json:"cluster_to"`   // 订单接受区域代码。
	Products    []PostingFinancialDataProduct `json:"products"`     // 订单中的商品列表。
}

type V3FbsPosting struct {
	AnalyticsData            V3FbsPostingAnalyticsData                                  `json:"analytics_data"`
	DeliveryMethod           V3DeliveryMethod                                           `json:"delivery_method"`
	IsPresortable            bool                                                       `json:"is_presortable"` // 如果发运是预先分拣的，则为`true`。
	Optional                 V3FbsPostingDetailOptional                                 `json:"optional"`
	Addressee                V3AddresseeFbsLists                                        `json:"addressee"`
	IsExpress                bool                                                       `json:"is_express"`    // 如果使用快速物流 Ozon Express —— `true`。
	OrderID                  int64                                                      `json:"order_id"`      // 货件所属订单的ID。
	Products                 []V3FbsPostingProduct                                      `json:"products"`      // 货运商品列表。
	Substatus                Substatus                                                  `json:"substatus"`     // 发货子状态： - `posting_acceptance_in_progress` —— 正在验收， - `posting_in_arbitration` —— 仲裁， - `posting_created` —— 已创建， - `post...
	InProcessAt              string                                                     `json:"in_process_at"` // 开始处理货件的日期和时间。
	ShipmentDate             string                                                     `json:"shipment_date"` // 必须收取货件的日期和时间。 超出该时间后将适用新费率，相关信息请查看字段 `tariffication`。
	Tariffication            V3FbsTariffication                                         `json:"tariffication"`
	Cancellation             V3Cancellation                                             `json:"cancellation"`
	LegalInfo                V2FboSinglePostingLegalInfo                                `json:"legal_info"`
	TrackingNumber           string                                                     `json:"tracking_number"`   // 货件跟踪号。
	AvailableActions         []string                                                   `json:"available_actions"` // 可用的操作和货件信息包括： - `arbitration` — 提出争议； - `awaiting_delivery` — 转为“等待发运”状态； - `can_create_chat` — 与买家开启聊天； - `cancel` — 取消...
	FinancialData            V3PostingFinancialData                                     `json:"financial_data"`
	OrderNumber              string                                                     `json:"order_number"` // 货件所属的订单号。
	Status                   Status                                                     `json:"status"`       // 货运状态: - `acceptance_in_progress` —— 正在验收， - `arbitration` —— 仲裁， - `awaiting_approve` —— 等待确认， - `awaiting_deliver` —— 等...
	Barcodes                 V3Barcodes                                                 `json:"barcodes"`
	DestinationPlaceID       int64                                                      `json:"destination_place_id"`        // 目的仓库的标识符。
	ParentPostingNumber      string                                                     `json:"parent_posting_number"`       // 快递母件编号，从该母件中拆分出了当前货件。
	PostingNumber            string                                                     `json:"posting_number"`              // 货件号。
	TarifficationSteps       []PostingV4PostingFbsListResponsePostingsTarifficationStep `json:"tariffication_steps"`         // 计费阶段。
	ShipmentDateWithoutDelay string                                                     `json:"shipment_date_without_delay"` // 不逾期发运日期和时间。
	TPLIntegrationType       TPLIntegrationType                                         `json:"tpl_integration_type"`        // 快递服务集成类型： - `ozon` —— Ozon 快递服务。 - `3pl_tracking` —— 集成服务快递。 - `non_integrated` —— 第三方物流服务。 - `aggregator` —— 通过Ozon合作物流...
	Customer                 V3CustomerFbsLists                                         `json:"customer"`
	DeliveringDate           string                                                     `json:"delivering_date"`        // 货件交付物流的时间。
	DestinationPlaceName     string                                                     `json:"destination_place_name"` // 目的仓库的名称。
	Requirements             V3FbsPostingRequirementsV3                                 `json:"requirements"`
}

// 货运数组。
type V3GetFbsPostingListResponseV3Result struct {
	HasNext  bool           `json:"has_next"` // 响应中未返回整个货运数组的标志: - `true` — 必须提出含其他值 `offset`的新请求，以获得其他货运信息； - `false` — 响应中返回了在请求中提出的整个用于过滤的货运数组。
	Postings []V3FbsPosting `json:"postings"` // 货运信息。
}

type V3GetFbsPostingListResponseV3 struct {
	Result V3GetFbsPostingListResponseV3Result `json:"result"`
}

type V4PostingProductDetailWithoutDimensions struct {
	CurrencyCode  CurrencyCode `json:"currency_code"`  // 价格币种。这与您在个人中心中设置的币种 一 致。 可能的值： - `RUB` — 卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
	MandatoryMark []string     `json:"mandatory_mark"` // 强制标记“诚信标志”。
	Name          string       `json:"name"`           // 商品名称。
	OfferID       string       `json:"offer_id"`       // 卖家系统中的商品ID — 货号。
	Price         string       `json:"price"`          // 价格。
	Quantity      int32        `json:"quantity"`       // 发货商品数量。
	SKU           int64        `json:"sku"`            // Ozon系统中的商品ID — SKU.
}

type PostingShipV4ResponseShipAdditionalData struct {
	Products      []V4PostingProductDetailWithoutDimensions `json:"products"`       // 正在运输途中的商品列表。
	PostingNumber string                                    `json:"posting_number"` // 发货号。
}

type V4FbsPostingShipV4Response struct {
	AdditionalData []PostingShipV4ResponseShipAdditionalData `json:"additional_data"` // 与发货有关的附加信息。
	Result         []string                                  `json:"result"`          // 货运装配结果。
}

type PostingShipV4RequestPackageProduct struct {
	ProductID int64 `json:"product_id"` // Ozon系统中的商品识别码是SKU。
	Quantity  int32 `json:"quantity"`   // 副本数量。
}

type PostingShipV4RequestPackage struct {
	Products []PostingShipV4RequestPackageProduct `json:"products"` // 运输途中的商品清单。
}

type V4FbsPostingShipV4Request struct {
	With          PostingShipV4RequestWith      `json:"with"`
	Packages      []PostingShipV4RequestPackage `json:"packages"`       // 包装清单。 每个包装都包含订单分成的发货清单。
	PostingNumber string                        `json:"posting_number"` // 发货号。
}

type V1AssemblyFbsPostingListResponse struct {
	Cursor   string                                    `json:"cursor"`   // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Cutoff   string                                    `json:"cutoff"`   // 卖家需在此时间前完成订单备货。
	Postings []V1AssemblyFbsPostingListResponsePosting `json:"postings"` // 货件列表。
}

// 货件状态最后一次发生变更的时间段。
type PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"` // 周期开始日期。
	To   string `json:"to"`   // 周期结束日期。
}

type V4FbsPostingShipPackageV4Response struct {
	Result string `json:"result"` // 备货后生成的货件号码。
}

type PostingPostingFBSPackageLabelResponse struct {
	FileContent string `json:"file_content"` // 二进制形式的文件内容。
	FileName    string `json:"file_name"`    // 文件名称。
	ContentType string `json:"content_type"` // 文件类型。
}

type PostingTrackingNumberSetRequestTrackingNumber struct {
	PostingNumber  string `json:"posting_number"`  // 货件ID。
	TrackingNumber string `json:"tracking_number"` // 货件追踪号。
}

type PostingFbsPostingTrackingNumberSetRequest struct {
	TrackingNumbers []PostingTrackingNumberSetRequestTrackingNumber `json:"tracking_numbers"` // 具有成对货运ID的数据 - 追踪号。
}

// 带有附加特征的商品列表。
type PostingV4PostingFbsUnfulfilledListResponsePostingsOptional struct {
	ProductsWithPossibleMandatoryMark []string `json:"products_with_possible_mandatory_mark"` // 可能需要标志的商品列表。
}

type V1SetPostingCutoffRequest struct {
	NewCutoffDate string `json:"new_cutoff_date"` // 新发运日期。
	PostingNumber string `json:"posting_number"`  // 货件编号。
}

// 商品尺寸。
type V3Dimensions struct {
	Width  string `json:"width"`  // 包装宽度。
	Height string `json:"height"` // 包装高度。
	Length string `json:"length"` // 商品长度。
	Weight string `json:"weight"` // 商品包装重量。
}

// 商品尺寸。
type V3PostingProductDetail struct {
	Dimensions    V3Dimensions `json:"dimensions"`
	CurrencyCode  CurrencyCode `json:"currency_code"`  // 价格显示的货币，其与个人中心中设置的币种相匹配。 -`RUB` — 俄罗斯卢布， - `BYN` — 白俄罗斯卢布， - `KZT` — 坚戈， - `EUR` — 欧元， - `USD` — 美元， - `CNY` — 元。
	Quantity      int32        `json:"quantity"`       // 商品数量。
	SKU           int64        `json:"sku"`            // Ozon 系统中的商品标识符（SKU）。
	HasImei       bool         `json:"has_imei"`       // 存在 IMEI。 若存在 IMEI，则为`true`。
	MandatoryMark []string     `json:"mandatory_mark"` // 商品强制性标签。
	Name          string       `json:"name"`           // 名称。
	OfferID       string       `json:"offer_id"`       // 卖家系统中的商品ID — 货号。
	Price         string       `json:"price"`          // 折扣后商品价格 — 该值在商品卡片上显示。
}

// 快递员信息。
type PostingDetailCourier struct {
	CarModel  string `json:"car_model"`  // 汽车型号。
	CarNumber string `json:"car_number"` // 车牌号。
	Name      string `json:"name"`       // 快递员全名。
	Phone     string `json:"phone"`      // 快递员电话。 过时的参数，不再使用。并总是返回到空字符串 `""`。
}

// 有关产品及其副本的信息。 响应包含 `product_exemplars`字段, 如果在请求中传递标志 `with.product_exemplars = true`。
type V3FbsPostingProductExemplarsV3 struct {
	Products []V3FbsPostingExemplarProductV3 `json:"products"`
}

// 收件人联系方式。
type V3Addressee struct {
	Name  string `json:"name"`  // 买家姓名。
	Phone string `json:"phone"` // 收件人的替代联系电话。
}

type V3AdditionalDataItem struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

// 买家信息。
type V3Customer struct {
	Address    V3Address `json:"address"`
	CustomerID int64     `json:"customer_id"` // 买家ID。
	Name       string    `json:"name"`        // 买家姓名。
	Phone      string    `json:"phone"`       // 买家的替代联系电话。
}

// 相关货件。
type V3FbsPostingDetailRelatedPostings struct {
	RelatedPostingNumbers []string `json:"related_posting_numbers"` // 相关货件号码列表。
}

// 货件信息。
// PreviousSubstatus values
type PreviousSubstatus string

const (
	PreviousSubstatusPostingAcceptanceInProgress        PreviousSubstatus = "posting_acceptance_in_progress"         // 正在验收
	PreviousSubstatusPostingInArbitration               PreviousSubstatus = "posting_in_arbitration"                 // 仲裁
	PreviousSubstatusPostingCreated                     PreviousSubstatus = "posting_created"                        // 已创建
	PreviousSubstatusPostingInCarriage                  PreviousSubstatus = "posting_in_carriage"                    // 在运输途中
	PreviousSubstatusPostingNotInCarriage               PreviousSubstatus = "posting_not_in_carriage"                // 未在运输中
	PreviousSubstatusPostingRegistered                  PreviousSubstatus = "posting_registered"                     // 已登记
	PreviousSubstatusPostingTransferringToDelivery      PreviousSubstatus = "posting_transferring_to_delivery"       // (`status=awaiting_deliver`) —— 移交给快递
	PreviousSubstatusPostingAwaitingPassportData        PreviousSubstatus = "posting_awaiting_passport_data"         // 等待护照资料
	PreviousSubstatusPostingAwaitingRegistration        PreviousSubstatus = "posting_awaiting_registration"          // 等待注册
	PreviousSubstatusPostingRegistrationError           PreviousSubstatus = "posting_registration_error"             // 注册错误
	PreviousSubstatusPostingSplitPending                PreviousSubstatus = "posting_split_pending"                  // 已创建
	PreviousSubstatusPostingCanceled                    PreviousSubstatus = "posting_canceled"                       // 已取消
	PreviousSubstatusPostingInClientArbitration         PreviousSubstatus = "posting_in_client_arbitration"          // 快递会员仲裁
	PreviousSubstatusPostingDelivered                   PreviousSubstatus = "posting_delivered"                      // 已送达
	PreviousSubstatusPostingReceived                    PreviousSubstatus = "posting_received"                       // 已收到
	PreviousSubstatusPostingConditionallyDelivered      PreviousSubstatus = "posting_conditionally_delivered"        // 暂时送到
	PreviousSubstatusPostingInCourierService            PreviousSubstatus = "posting_in_courier_service"             // 快递员正在路上
	PreviousSubstatusPostingInPickupPoint               PreviousSubstatus = "posting_in_pickup_point"                // 在取货点
	PreviousSubstatusPostingOnWayToCity                 PreviousSubstatus = "posting_on_way_to_city"                 // 发往城市途中
	PreviousSubstatusPostingOnWayToPickupPoint          PreviousSubstatus = "posting_on_way_to_pickup_point"         // 正发往取货点
	PreviousSubstatusPostingReturnedToWarehouse         PreviousSubstatus = "posting_returned_to_warehouse"          // 返回仓库
	PreviousSubstatusPostingTransferredToCourierService PreviousSubstatus = "posting_transferred_to_courier_service" // 转交给快递员
	PreviousSubstatusPostingDriverPickUp                PreviousSubstatus = "posting_driver_pick_up"                 // 在司机那儿
	PreviousSubstatusPostingNotInSortCenter             PreviousSubstatus = "posting_not_in_sort_center"             // 集散中心未收到
)

type V3FbsPostingDetail struct {
	ProductExemplars         V3FbsPostingProductExemplarsV3                             `json:"product_exemplars"`
	TarifficationSteps       []PostingV4PostingFbsListResponsePostingsTarifficationStep `json:"tariffication_steps"` // 计费阶段。
	FinancialData            V3PostingFinancialData                                     `json:"financial_data"`
	IsExpress                bool                                                       `json:"is_express"` // 如果使用了快速物流Ozon Express —— `true`。
	Optional                 V3FbsPostingDetailOptional                                 `json:"optional"`
	AnalyticsData            V3FbsPostingAnalyticsData                                  `json:"analytics_data"`
	Cancellation             V3Cancellation                                             `json:"cancellation"`
	Customer                 V3Customer                                                 `json:"customer"`
	OrderID                  int64                                                      `json:"order_id"`        // 货件所属的订单ID。
	Status                   Status                                                     `json:"status"`          // 货运状态: - `acceptance_in_progress` —— 正在验收， - `arbitration` —— 仲裁， - `awaiting_approve` —— 等待确认， - `awaiting_deliver` —— 等...
	TrackingNumber           string                                                     `json:"tracking_number"` // 货件跟踪号。
	Tariffication            V3FbsTariffication                                         `json:"tariffication"`
	Addressee                V3Addressee                                                `json:"addressee"`
	ProviderStatus           string                                                     `json:"provider_status"` // 快递服务状态。
	Requirements             V3FbsPostingRequirementsV3                                 `json:"requirements"`
	DeliveryPrice            string                                                     `json:"delivery_price"` // 物流价格。
	LegalInfo                V2FboSinglePostingLegalInfo                                `json:"legal_info"`
	RelatedPostings          V3FbsPostingDetailRelatedPostings                          `json:"related_postings"`
	TPLIntegrationType       TPLIntegrationType                                         `json:"tpl_integration_type"`        // 快递服务集成类型： - `ozon` —— 通过Ozon物流的快递。 - `aggregator` —— 外部服务快递，Ozon注册订单。 - `3pl_tracking` —— 外部服务快递，卖家注册订单。 - `non_integrat...
	FactDeliveryDate         string                                                     `json:"fact_delivery_date"`          // 货件实际转移配送的日期。
	Products                 []V3PostingProductDetail                                   `json:"products"`                    // 货物装运的数组。
	Substatus                Substatus                                                  `json:"substatus"`                   // 发货子状态： - `posting_acceptance_in_progress` —— 正在验收， - `posting_in_arbitration` —— 仲裁， - `posting_created` —— 已创建， - `post...
	PreviousSubstatus        PreviousSubstatus                                          `json:"previous_substatus"`          // 货件的前一个子状态。可能的取值： - `posting_acceptance_in_progress` —— 正在验收， - `posting_in_arbitration` —— 仲裁， - `posting_created` —— 已创...
	ParentPostingNumber      string                                                     `json:"parent_posting_number"`       // 母件编号，由该母件拆分出了该货件。
	ShipmentDate             string                                                     `json:"shipment_date"`               // 必须收取货件的日期和时间。 超出该时间后将适用新费率，相关信息请查看字段 `tariffication`。
	ShipmentDateWithoutDelay string                                                     `json:"shipment_date_without_delay"` // 不逾期发运日期和时间。
	AdditionalData           []V3AdditionalDataItem                                     `json:"additional_data"`
	AvailableActions         []string                                                   `json:"available_actions"` // 可用的操作和货件信息包括： - `arbitration` — 提出争议； - `awaiting_delivery` — 转为“等待发运”状态； - `can_create_chat` — 与买家开启聊天； - `cancel` — 取消...
	Barcodes                 V3Barcodes                                                 `json:"barcodes"`
	DeliveryMethod           V3DeliveryMethod                                           `json:"delivery_method"`
	Courier                  PostingDetailCourier                                       `json:"courier"`
	DeliveringDate           string                                                     `json:"delivering_date"` // 货件交付物流的时间。
	OrderNumber              string                                                     `json:"order_number"`    // 货件所属的订单号。
	InProcessAt              string                                                     `json:"in_process_at"`   // 开始处理货件的日期和时间。
	PostingNumber            string                                                     `json:"posting_number"`  // 货件号。
}

type V3GetFbsPostingResponseV3 struct {
	Result V3FbsPostingDetail `json:"result"`
}

type V1CarriageApproveRequest struct {
	ContainersCount int32 `json:"containers_count"` // 货位数量。 如果您已开通信任验收，并按货位发运订单，请使用该参数。如果您未开通信任验收，请跳过该参数。
	CarriageID      int64 `json:"carriage_id"`      // 发运标识符。
}

// 筛选器。
type V1AssemblyCarriageProductListRequestFilter struct {
	CarriageID       int64  `json:"carriage_id"`        // 运输标识符。
	CutoffFrom       string `json:"cutoff_from"`        // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	CutoffTo         string `json:"cutoff_to"`          // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	DeliveryMethodID int64  `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
}

// 货件状态最后一次发生变更的时间段。
type PostingV4PostingFbsListRequestFilterLastChangedStatusDate struct {
	From string `json:"from"` // 周期开始日期。
	To   string `json:"to"`   // 周期结束日期。
}

// 筛选器。
type V1AssemblyCarriagePostingListRequestFilter struct {
	DeliveryMethodID int64  `json:"delivery_method_id"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	CarriageID       int64  `json:"carriage_id"`        // 运输标识符。
	CutoffFrom       string `json:"cutoff_from"`        // 按卖家需完成订单备货的时间进行筛选。时间段开始。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
	CutoffTo         string `json:"cutoff_to"`          // 按卖家需完成订单备货的时间进行筛选。时间段结束。 格式： `YYYY-MM-DDThh:mm:ss.mcsZ`。示例：`2020-03-18T07:34:50.359Z`。
}

type V2PostingFBSActGetPostingsRequest struct {
	ID any `json:"id"` // 单据标识符。请通过方法[/v1/carriage/create](#operation/CarriageAPI_CarriageCreate)获取参数值。
}

type CarriageCarriageGetRequest struct {
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
}

type CarriageCarriageGetResponse struct {
	UpdatedAt                  string                                        `json:"updated_at"`                 // 运输信息最后一次更新日期。
	WarehouseID                int64                                         `json:"warehouse_id"`               // 仓库标识符。
	ActType                    string                                        `json:"act_type"`                   // 交接单类型。针对FBS卖家。
	FirstMileType              string                                        `json:"first_mile_type"`            // 头程物流类型。
	IsContainerLabelPrinted    bool                                          `json:"is_container_label_printed"` // `true`, 如果您已经打印了货位标签。
	TPLProviderID              int64                                         `json:"tpl_provider_id"`            // 配送服务商标识符。
	ArrivalPassIds             []string                                      `json:"arrival_pass_ids"`           // 为运输生成的通行证标识符列表。
	AvailableActions           []string                                      `json:"available_actions"`          // 运输的可用操作： - `get_shipping_list`——获取发运清单； - `get_act_of_acceptance`——获取验收证明书； - `get_waybill`——获取 PDF 格式的货单； - `set_arriva...
	DepartureDate              string                                        `json:"departure_date"`             // 运输完成日期。
	IntegrationType            string                                        `json:"integration_type"`           // 运输类型。
	IsPartial                  bool                                          `json:"is_partial"`                 // `true`, 如果是部分运输。
	RetryCount                 int32                                         `json:"retry_count"`                // 运输创建重复尝试数量。
	CancelAvailability         CarriageCarriageGetResponseCancelAvailability `json:"cancel_availability"`
	CompanyID                  int64                                         `json:"company_id"`                     // 卖家标识符。
	CreatedAt                  string                                        `json:"created_at"`                     // 运输创建日期。
	Status                     string                                        `json:"status"`                         // 运输状态。
	CarriageID                 int64                                         `json:"carriage_id"`                    // 运输标识符。
	ContainersCount            int32                                         `json:"containers_count"`               // 货位数量。
	DeliveryMethodID           int64                                         `json:"delivery_method_id"`             // 物流方式标识符。
	HasPostingsForNextCarriage bool                                          `json:"has_postings_for_next_carriage"` // `true`, 如果有未能进行运输，但需要发运的货件。
	PartialNum                 int64                                         `json:"partial_num"`                    // 部分运输序列号。
}

type PostingPostingFBSActGetContainerLabelsRequest struct {
	ID int64 `json:"id"` // 来自方法[/v1/carriage/create](#operation/CarriageAPI_CarriageCreate)的文件生成任务编号（也是运输标识符）。
}

// 请求结果。
type Postingv3GetFbsPostingUnfulfilledListResponseResult struct {
	Count    int64          `json:"count"`    // 在响应中的元素计数器。
	Postings []V3FbsPosting `json:"postings"` // 货件清单和每个货物的详细信息。
}

type ResultError struct {
	Code   string `json:"code"`   // 错误代码。
	Status Status `json:"status"` // 错误类型： - `warning` — 提醒； - `critical` — 严重错误。
}

// 配送方法信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod struct {
	Warehouse     string `json:"warehouse"`       // 仓库名称。
	WarehouseID   int64  `json:"warehouse_id"`    // 仓库标识符。
	ID            int64  `json:"id"`              // 配送方式标识符。
	Name          string `json:"name"`            // 配送方法名称。
	TPLProvider   string `json:"tpl_provider"`    // 配送服务。
	TPLProviderID int64  `json:"tpl_provider_id"` // 配送服务标识符。
}

// 货件状态最后一次发生变更的时间段。
type V3TimeRange struct {
	From string `json:"from"` // 时间段的开始日期。
	To   string `json:"to"`   // 时间段的结束日期。
}

// 请求过滤。 请按装运时间使用过滤器 — `cutoff`, 或者按照货件交付给快递的时间 — `delivering_date`。 如果一起使用，则会在响应中返回错误。 要按装运时间使用过滤器，请填 `cutoff_from` 和 `cut...
type Postingv3GetFbsPostingUnfulfilledListRequestFilter struct {
	CutoffFrom            string      `json:"cutoff_from"`          // 按卖家需要收订单的时间进行筛选。 时间段开始。 格式： YYYY-MM-DDThh:mm:ss.mcsZ. 例子： 2020-03-18T07:34:50.359Z.
	DeliveringDateFrom    string      `json:"delivering_date_from"` // 将货件交给物流的最快日期。
	DeliveringDateTo      string      `json:"delivering_date_to"`   // 将货件交给物流的最迟日期。
	DeliveryMethodID      []int64     `json:"delivery_method_id"`   // 快递方式ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	LastChangedStatusDate V3TimeRange `json:"last_changed_status_date"`
	FbpFilter             FbpFilter   `json:"fbpFilter"`    // 从合作伙伴仓库（FBP）发货时的货件筛选器： - `ALL` — 响应中将显示所有符合其他筛选器条件的货件； - `ONLY` — 仅显示FBP货件； - `WITHOUT` — 显示除FBP外的所有货件。 默认值为 `ALL`。
	ProviderID            []int64     `json:"provider_id"`  // 快递服务ID。按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	CutoffTo              string      `json:"cutoff_to"`    // 按卖家需要收订单的时间进行筛选。 时间段结束。 格式： YYYY-MM-DDThh:mm:ss.mcsZ. 例子： 2020-03-18T07:34:50.359Z.
	Status                Status      `json:"status"`       // 货件运输状态： - `acceptance_in_progress` — 正在验收， - `awaiting_approve` — 等待确认， - `awaiting_packaging` — 等待包装， - `awaiting_regis...
	WarehouseID           []int64     `json:"warehouse_id"` // 仓库ID。可以使用方法 [/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
}

type Postingv3GetFbsPostingUnfulfilledListRequest struct {
	Dir    Dir                                                `json:"dir"` // 分类方向： - `asc` — 从小到大， - `desc` — 从大到小。
	Filter Postingv3GetFbsPostingUnfulfilledListRequestFilter `json:"filter"`
	Limit  int64                                              `json:"limit"`  // 响应中值的数量： - 最大值 — 1000， - 最小值 — 1。
	Offset int64                                              `json:"offset"` // 将在响应中跳过的元素数。 例如，如果“offset=10”，那么响应将从找到的第11个元素开始。
	With   Postingv3FbsPostingWithParams                      `json:"with"`
}

type V2FbsPostingProduct struct {
	Name     string `json:"name"`     // 商品名称。
	OfferID  string `json:"offer_id"` // 卖家系统中的商品ID — 货号。
	Price    string `json:"price"`    // 商品价格。
	Quantity int64  `json:"quantity"` // 货运商品数量。
	SKU      int64  `json:"sku"`      // Ozon 系统中的商品标识符（SKU）。
}

// 货件条形码。
type PostingBarcodes struct {
	UpperBarcode string `json:"upper_barcode"` // 货件标签的上条形码。
	LowerBarcode string `json:"lower_barcode"` // 货件标签的下条形码。
}

// 请求结果。
type V2FbsPosting struct {
	Barcodes       PostingBarcodes       `json:"barcodes"`
	CancelReasonID int64                 `json:"cancel_reason_id"` // 取消装运原因ID。
	InProcessAt    string                `json:"in_process_at"`    // 开始处理货件的日期和时间。
	OrderID        int64                 `json:"order_id"`         // 货运所属订单ID。
	CreatedAt      string                `json:"created_at"`       // 创建装运日期和时间。
	OrderNumber    string                `json:"order_number"`     // 货运所属的订单号。
	PostingNumber  string                `json:"posting_number"`   // 货运号。
	Products       []V2FbsPostingProduct `json:"products"`         // 货运商品列表。
	ShipmentDate   string                `json:"shipment_date"`    // 必须收取货件的日期和时间。 如果在此日期之前未完成配货，则货运自动取消。
	Status         string                `json:"status"`           // 货运状态。
}

// 货运信息。
type V2FbsPostingResponse struct {
	Result V2FbsPosting `json:"result"`
}

type V2FbsPostingProductCountryListResponse struct {
	Result []V2FbsPostingProductCountryListResponseResult `json:"result"` // 制造国和ISO代码列表。
}

// 筛选器。
// Statuses values
type Statuses string

const (
	StatusesAwaitingRegistration Statuses = "awaiting_registration"  // 等待登记；
	StatusesAcceptanceInProgress Statuses = "acceptance_in_progress" // 接收 中；
	StatusesAwaitingApprove      Statuses = "awaiting_approve"       // 等待 确认；
	StatusesAwaitingPackaging    Statuses = "awaiting_packaging"     // 等待 包装；
	StatusesAwaitingDeliver      Statuses = "awaiting_deliver"       // 等待 发运；
	StatusesArbitration          Statuses = "arbitration"            // 仲裁；
	StatusesClientArbitration    Statuses = "client_arbitration"     // 客户配送仲裁；
	StatusesDelivering           Statuses = "delivering"             // 运输中；
	StatusesDriverPickup         Statuses = "driver_pickup"          // 司机正在送货；
	StatusesDelivered            Statuses = "delivered"              // 已配送；
	StatusesCancelled            Statuses = "cancelled"              // 已取消；
	StatusesNotAccepted          Statuses = "not_accepted"           // 在分拣中心尚未接收；
	StatusesSentBySeller         Statuses = "sent_by_seller"         // 已被卖家发送
)

type PostingV4PostingFbsListRequestFilter struct {
	IsBlrTraceable        bool                                                      `json:"is_blr_traceable"` // `true`，表示商品可追溯。
	LastChangedStatusDate PostingV4PostingFbsListRequestFilterLastChangedStatusDate `json:"last_changed_status_date"`
	OrderID               int64                                                     `json:"order_id"`            // 订单标识符。
	ProviderIds           []string                                                  `json:"provider_ids"`        // 配送服务标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	Since                 string                                                    `json:"since"`               // 需要获取货件列表的周期开始日期。
	To                    string                                                    `json:"to"`                  // 需要获取货件列表的周期结束日期。
	OrderNumbers          []string                                                  `json:"order_numbers"`       // 货件所属订单的订单号。
	Statuses              []string                                                  `json:"statuses"`            // 货件状态： - `awaiting_registration`——等待登记； - `acceptance_in_progress`——接收 中； - `awaiting_approve`——等待 确认； - `awaiting_packag...
	WarehouseIds          []string                                                  `json:"warehouse_ids"`       // 仓库标识符。可通过方法[/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	DeliveryMethodIds     []string                                                  `json:"delivery_method_ids"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
}

// 排序方向： - `ASC`——升序； - `DESC`——降序。
type PostingV4PostingFbsListRequestSortDirEnum string

type PostingV4PostingFbsListRequest struct {
	Limit    int64                                     `json:"limit"` // 响应中返回的值数量。
	SortDir  PostingV4PostingFbsListRequestSortDirEnum `json:"sort_dir"`
	Translit bool                                      `json:"translit"` // 若为`true`，则启用将地址从西里尔字母转写为拉丁字母。
	With     PostingV4PostingFbsListRequestWith        `json:"with"`
	Cursor   string                                    `json:"cursor"` // 用于选择下一批数据的指针。
	Filter   PostingV4PostingFbsListRequestFilter      `json:"filter"`
}

// 需要添加到响应中的附加字段。
type Postingv3FbsPostingWithParamsExamplars struct {
	Translit         bool `json:"translit"`          // 完成返回值的拼写。
	AnalyticsData    bool `json:"analytics_data"`    // 将分析数据添加到响应中。
	Barcodes         bool `json:"barcodes"`          // 将货件条形码添加到响应中。
	FinancialData    bool `json:"financial_data"`    // 将财务数据添加到响应中。
	LegalInfo        bool `json:"legal_info"`        // 将法律信息添加到响应中。
	ProductExemplars bool `json:"product_exemplars"` // 将有关产品及其份数的数据添加到响应中。
	RelatedPostings  bool `json:"related_postings"`  // 将相关货件数量添加到响应中。 相关货件是在组装期间将母快递拆分的快递。
}

type V2FbsPostingProductCountrySetRequest struct {
	PostingNumber  string `json:"posting_number"`   // 货运号。
	ProductID      int64  `json:"product_id"`       // 产品ID。
	CountryIsoCode string `json:"country_iso_code"` // 根据标准ISO_3166-1添加的国家的两个字母代码 制造国家列表及其ISO代码可以使用该方法获得[/v2/posting/fbs/product/country/list](#operation/PostingAPI_ListCountr...
}

type GetCarriageAvailableListResponseResult struct {
	HasEntrustedAcceptance            bool          `json:"has_entrusted_acceptance"`               // 信任接收的标志。 如果在仓库中启用了信任接收，则为“true”。
	MandatoryPackagedCount            int32         `json:"mandatory_packaged_count"`               // 收取货件数量。
	CarriageID                        int64         `json:"carriage_id"`                            // 运输ID（也是文件形成的任务编号）。
	DeliveryMethodName                string        `json:"delivery_method_name"`                   // 快递方式名称。
	WarehouseName                     string        `json:"warehouse_name"`                         // 仓库名称。
	RecommendedTimeUtcOffsetInMinutes float64       `json:"recommended_time_utc_offset_in_minutes"` // 推荐发运时间与UTC-0的时区偏移量（以分钟为单位）。
	CarriagePostingsCount             int32         `json:"carriage_postings_count"`                // 运输中的货件数量。
	CarriageStatus                    string        `json:"carriage_status"`                        // 所请求的交付方式和装运日期的运输状态。
	CutoffAt                          string        `json:"cutoff_at"`                              // 需要收取货件的日期和时间。
	MandatoryPostingsCount            int32         `json:"mandatory_postings_count"`               // 需要收取的货件数量。
	WarehouseCity                     string        `json:"warehouse_city"`                         // 仓库所在城市。
	DeliveryMethodID                  int64         `json:"delivery_method_id"`                     // 快递方式ID。
	TPLProviderName                   string        `json:"tpl_provider_name"`                      // 快递服务名称。
	WarehouseID                       int64         `json:"warehouse_id"`                           // 仓库ID。
	WarehouseTimezone                 string        `json:"warehouse_timezone"`                     // 仓库所在时区。
	RecommendedTimeLocal              string        `json:"recommended_time_local"`                 // 推荐的本地发运时间（订单接收点）。
	TPLProviderIconURL                string        `json:"tpl_provider_icon_url"`                  // 快递服务图标的链接。
	Errors                            []ResultError `json:"errors"`                                 // 错误列表。
	FirstMileType                     string        `json:"first_mile_type"`                        // 第一英里类型。
}

// 货件条形码。
type PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes struct {
	UpperBarcode string `json:"upper_barcode"` // 货件标签上的上方条形码。
	LowerBarcode string `json:"lower_barcode"` // 货件标签上的下方条形码。
}

// 收件人联系方式。
type PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee struct {
	Name string `json:"name"` // 收件人姓名。
}

type V1SetPostingCutoffResponse struct {
	Result bool `json:"result"` // `true`表示已设置新日期。
}

type SetPostingsResponseResult struct {
	Error         string `json:"error"`          // 错误描述。
	PostingNumber string `json:"posting_number"` // 货件编号。
	Result        bool   `json:"result"`         // 请求处理结果：若请求处理成功，返回值为`true`。
}

type V1SetPostingsResponse struct {
	Result []SetPostingsResponseResult `json:"result"`
}

type V1CarriageCreateResponse struct {
	CarriageID int64 `json:"carriage_id"` // 运输标识符。
}

type V1CarriageCancelRequest struct {
	CarriageID int64 `json:"carriage_id"` // 发运识别符。
}

type PostingCancelReason struct {
	Title                      string `json:"title"`                         // 类别名称。
	TypeID                     TypeID `json:"type_id"`                       // 取消货件ID： - `buyer` — 买家， - `seller` — 卖家。
	ID                         int64  `json:"id"`                            // 取消原因ID。
	IsAvailableForCancellation bool   `json:"is_available_for_cancellation"` // 取消装运结果。 `true`, 如果请求可以取消。
}

type PostingCancelReasonListResponse struct {
	Result []PostingCancelReason `json:"result"` // 方法操作结果。
}

type Postingv3GetFbsPostingRequest struct {
	PostingNumber string                                 `json:"posting_number"` // 货件ID。
	With          Postingv3FbsPostingWithParamsExamplars `json:"with"`
}

type PostingCancelReasonRequest struct {
	RelatedPostingNumbers []string `json:"related_posting_numbers"` // 货件号。
}

// 配送地址。
type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress struct {
	Region          string  `json:"region"`            // 配送地区。
	Comment         string  `json:"comment"`           // 订单备注。
	Country         string  `json:"country"`           // 配送国家。
	ZipCode         string  `json:"zip_code"`          // 收件人邮政编码。
	AddressTail     string  `json:"address_tail"`      // 文本格式地址。
	City            string  `json:"city"`              // 配送城市。
	District        string  `json:"district"`          // 配送区域。
	Latitude        float64 `json:"latitude"`          // 纬度。
	Longitude       float64 `json:"longitude"`         // 经度。
	ProviderPvzCode string  `json:"provider_pvz_code"` // 3PL服务商取货点代码。
	PvzCode         int64   `json:"pvz_code"`          // 取货点代码。
}

// 买家信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer struct {
	Phone         string                                                            `json:"phone"` // 买家的替换联系电话。
	Address       PostingV4PostingFbsUnfulfilledListResponsePostingsCustomerAddress `json:"address"`
	CustomerEmail string                                                            `json:"customer_email"` // 买家电子邮箱。
	CustomerID    int64                                                             `json:"customer_id"`    // 买家标识符。
	Name          string                                                            `json:"name"`           // 买家姓名。
}

type V1AssemblyCarriageProductListRequest struct {
	Cursor string                                     `json:"cursor"` // 用于选择下一批数据的指针。
	Filter V1AssemblyCarriageProductListRequestFilter `json:"filter"`
	Limit  int64                                      `json:"limit"` // 每页显示的数量。
}

type AssemblyFbsProductListResponseProducts struct {
	Postings    []SPostings `json:"postings"`     // 货件列表。
	ProductName string      `json:"product_name"` // 商品名称。
	Quantity    int32       `json:"quantity"`     // 商品数量。
	SKU         int64       `json:"sku"`          // Ozon系统中的商品标识符——SKU。
	OfferID     string      `json:"offer_id"`     // 卖家系统中的商品标识符——货号。
	PictureURL  string      `json:"picture_url"`  // 商品图片链接。
}

// 发运的计费信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication struct {
	CurrentTariffType      string                           `json:"current_tariff_type"` // 计费类型——折扣或加价。
	NextTariffCharge       MoneyMoneyNextTariffCharge       `json:"next_tariff_charge"`
	NextTariffStartsAt     string                           `json:"next_tariff_starts_at"` // 新费率开始生效的日期和时间。
	NextTariffType         string                           `json:"next_tariff_type"`      // 在`next_tariff_starts_at`参数指定时间之后的计费类型——折扣或加价。
	CurrentTariffCharge    MoneyMoneyCurrentTariffCharge    `json:"current_tariff_charge"`
	NextTariffMinCharge    MoneyMoneyNextTariffMinCharge    `json:"next_tariff_min_charge"`
	NextTariffRate         float64                          `json:"next_tariff_rate"` // 在`next_tariff_starts_at`参数指定时间之后，该货件将按此百分比计费。
	CurrentTariffMinCharge MoneyMoneyCurrentTariffMinCharge `json:"current_tariff_min_charge"`
	CurrentTariffRate      float64                          `json:"current_tariff_rate"` // 计费百分比。
}

type PostingFbsPostingLastMileRequest struct {
	PostingNumber []string `json:"posting_number"` // 货件ID。
}

type PostingPostingFBSActGetContainerLabelsResponse struct {
	FileContent string `json:"file_content"` // 文件内容，以二进制形式。
	FileName    string `json:"file_name"`    // 文件名称。
	ContentType string `json:"content_type"` // 文件类型。
}

type V4FbsPostingShipPackageV4RequestProduct struct {
	ExemplarsIds []string `json:"exemplarsIds"` // 商品外部识别码。
	ProductID    int64    `json:"product_id"`   // Ozon系统中商品的标识符 — `product_id`。
	Quantity     int32    `json:"quantity"`     // 样件数量。
}

// 取消信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation struct {
	AffectCancellationRating bool                  `json:"affect_cancellation_rating"` // `true`，表示取消会影响卖家评级。
	CancelReason             string                `json:"cancel_reason"`              // 取消原因。
	CancelReasonID           int64                 `json:"cancel_reason_id"`           // 货件取消原因标识符。
	CancellationInitiator    CancellationInitiator `json:"cancellation_initiator"`     // 取消发起方： - `卖家`， - `客户`， - `买家`， - `Ozon`， - `系统`， - `配送服务商`。
	CancellationType         CancellationType      `json:"cancellation_type"`          // 取消类型： - `seller`——由卖家取消； - `client`或`customer`——由买家取消； - `ozon`——由Ozon取消； - `system`——由系统取消； - `delivery`——由配送服务取消。
	CancelledAfterShip       bool                  `json:"cancelled_after_ship"`       // `true`，表示取消发生在货件完成备货之后。
}

type V1AssemblyCarriagePostingListResponse struct {
	CanPrintMassLabel bool                                           `json:"can_print_mass_label"` // `true`，前提是可以批量打印标签。
	Cursor            string                                         `json:"cursor"`               // 用于选择下一批数据的指针。如果该参数为空，则没有更多数据了。
	Postings          []V1AssemblyCarriagePostingListResponsePosting `json:"postings"`             // 货件列表。
}

type PostingGetFbsPostingByBarcodeRequest struct {
	Barcode string `json:"barcode"` // 货运条形码。可以使用以下方法获取： [/v3/posting/fbs/get](#operation/PostingAPI_GetFbsPostingV3)、[/v3/posting/fbs/list](#operation/Posting...
}

// 请求筛选器。 使用按备货时间的筛选器——`cutoff`，或按货件转移配送的日期筛选——`delivering_date`。 如果同时使用这两个筛选器，响应会返回错误。 要使用按备货时间的筛选器，请填写`cutoff_from`和`cuto...
type PostingV4PostingFbsUnfulfilledListRequestFilter struct {
	DeliveryMethodIds     []string                                                             `json:"delivery_method_ids"` // 配送方式标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	LastChangedStatusDate PostingV4PostingFbsUnfulfilledListRequestFilterLastChangedStatusDate `json:"last_changed_status_date"`
	CutoffFrom            string                                                               `json:"cutoff_from"`          // 卖家必须完成订单备货的截止时间。时间段开始。
	CutoffTo              string                                                               `json:"cutoff_to"`            // 卖家必须完成订单备货的截止时间。时间段结束。
	ProviderIds           []string                                                             `json:"provider_ids"`         // 配送服务标识符。可通过方法[/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	Statuses              []string                                                             `json:"statuses"`             // 货件状态： - `acceptance_in_progress`——接收中； - `awaiting_approve`——等待确认； - `awaiting_packaging`——等待包装； - `awaiting_registratio...
	WarehouseIds          []string                                                             `json:"warehouse_ids"`        // 仓库标识符。可通过方法[/v1/warehouse/list](#operation/WarehouseAPI_WarehouseList)获取。
	DeliveringDateFrom    string                                                               `json:"delivering_date_from"` // 货件转移配送的最早日期。
	DeliveringDateTo      string                                                               `json:"delivering_date_to"`   // 货件转移配送的最晚日期。
}

type V1AssemblyCarriagePostingListRequest struct {
	Cursor string                                     `json:"cursor"` // 用于选择下一批数据的指针。
	Filter V1AssemblyCarriagePostingListRequestFilter `json:"filter"`
	Limit  int64                                      `json:"limit"` // 每页显示的数量。
}

type PostingPostingProductCancelResponse struct {
	Result string `json:"result"` // 货运号。
}

type Postingv3GetFbsPostingUnfulfilledListResponse struct {
	Result Postingv3GetFbsPostingUnfulfilledListResponseResult `json:"result"`
}

type PostingFbsPostingMoveStatusResponse struct {
	Result []PostingMoveStatusResponseMoveStatus `json:"result"` // 方法操作结果。
}

type V2MovePostingToAwaitingDeliveryRequest struct {
	PostingNumber []string `json:"posting_number"` // 货运ID。一次请求中的最大数量——100。
}

// 分析数据。
type PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData struct {
	TPLProvider             string               `json:"tpl_provider"`               // 配送服务。
	TPLProviderID           int64                `json:"tpl_provider_id"`            // 配送服务标识符。
	Warehouse               string               `json:"warehouse"`                  // 订单发货仓库名称。
	WarehouseID             int64                `json:"warehouse_id"`               // 仓库标识符。
	ClientDeliveryDateEnd   string               `json:"client_delivery_date_end"`   // 订单预计最晚送达日期。仅适用于通过Ozon配送创建的货件。
	DeliveryDateBegin       string               `json:"delivery_date_begin"`        // 配送开始的日期和时间。
	DeliveryType            string               `json:"delivery_type"`              // 配送方法。
	IsPremium               bool                 `json:"is_premium"`                 // `true`，表示收件人有Premium订阅。
	PaymentTypeGroupName    PaymentTypeGroupName `json:"payment_type_group_name"`    // 支付方法： - `在线银行卡支付`； - `Ozon Bank银行卡`； - `取货时从Ozon卡自动扣款`； - `收货时用已保存的银行卡支付`； - `快速支付系统`； - `Ozon分期付款`； - `支付至结算账户`； - `Sbe...
	Region                  string               `json:"region"`                     // 配送地区。仅适用于rFBS货件。
	City                    string               `json:"city"`                       // 配送城市。仅适用于rFBS货件和独联体卖家。
	ClientDeliveryDateBegin string               `json:"client_delivery_date_begin"` // 配送开始的日期和时间。仅适用于通过Ozon配送创建的货件。
	DeliveryDateEnd         string               `json:"delivery_date_end"`          // 配送结束日期和时间。
	IsLegal                 bool                 `json:"is_legal"`                   // `true`，表示收件人为法人。
}

type PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep struct {
	MinCharge        MoneyMoneyCurrentTariffMinCharge `json:"min_charge"`
	TariffCharge     MoneyMoneyCurrentTariffCharge    `json:"tariff_charge"`
	TariffDeadlineAt string                           `json:"tariff_deadline_at"` // 计费阶段结束的日期和时间。该日期后将自动进入下一计费阶段。
	TariffRate       float64                          `json:"tariff_rate"`        // 折扣或附加费用百分比。
	TariffType       string                           `json:"tariff_type"`        // 计费类型。
}

// 外部平台订单信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder struct {
	IsExternal   bool   `json:"is_external"`   // `true`，表示订单来自外部平台。
	PlatformName string `json:"platform_name"` // 下单平台名称。
}

// 商品价格、折扣金额、卖家应收款和佣金信息。
type PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData struct {
	ClusterFrom string                                                                    `json:"cluster_from"` // 订单发出地区代码。
	ClusterTo   string                                                                    `json:"cluster_to"`   // 订单配送地区代码。
	Products    []PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialDataProducts `json:"products"`     // 订单中的商品列表。
}

type PostingV4PostingFbsUnfulfilledListResponsePostings struct {
	PrrOption                PrrOption                                                             `json:"prr_option"` // 装卸服务代码： - `lift`——通过电梯搬运上楼； - `stairs`——通过楼梯搬运上楼； - `none`——买家拒绝该服务，无需将商品搬上楼； - `delivery_default`——配送费用已包含在价格中，根据要约条款需将...
	Barcodes                 PostingV4PostingFbsUnfulfilledListResponsePostingsBarcodes            `json:"barcodes"`
	DeliverySchema           DeliverySchema                                                        `json:"delivery_schema"`             // 配送模式： - `SDS`——统一SKU标识符； - `FBO`——从Ozon仓库销售的商品标识符； - `FBS`——从FBS仓库销售的商品标识符； - `Crossborder`——从境外销售的商品标识符。
	DestinationPlaceID       int64                                                                 `json:"destination_place_id"`        // 目的地标识符。
	Products                 []PostingV4PostingFbsUnfulfilledListResponsePostingsProducts          `json:"products"`                    // 货件中商品列表。
	Substatus                Substatus                                                             `json:"substatus"`                   // 货件子状态： - `posting_acceptance_in_progress`——验收中； - `posting_in_arbitration`——仲裁； - `posting_created`——已创建； - `posting_in_...
	TarifficationSteps       []PostingV4PostingFbsUnfulfilledListResponsePostingsTarifficationStep `json:"tariffication_steps"`         // 计费阶段。
	TPLIntegrationType       TPLIntegrationType                                                    `json:"tpl_integration_type"`        // 与配送服务的集成类型： - `ozon`——由Ozon配送； - `3pl_tracking`——由已集成服务配送； - `non_integrated`——由第三方服务配送； - `aggregator`——通过Ozon合作伙伴配送； -...
	RequireBlrTraceableAttrs bool                                                                  `json:"require_blr_traceable_attrs"` // `true`，表示需要填写可追踪属性。
	ShipmentDateWithoutDelay string                                                                `json:"shipment_date_without_delay"` // 不逾期发运的日期和时间。
	TrackingNumber           string                                                                `json:"tracking_number"`             // 货件跟踪号码。
	VolumeWeight             float64                                                               `json:"volume_weight"`               // 商品体积重量。
	AvailableActions         []string                                                              `json:"available_actions"`           // 货件的可用操作和信息： - `arbitration`——提出争议； - `awaiting_delivery`——变更为“等待发运”状态； - `can_create_chat`——与买家发起聊天； - `cancel`——取消货件； -...
	LegalInfo                PostingV4PostingFbsUnfulfilledListResponsePostingsLegalInfo           `json:"legal_info"`
	OrderID                  int64                                                                 `json:"order_id"`     // 该货件所属订单的标识符。
	OrderNumber              string                                                                `json:"order_number"` // 该货件所属订单的编号。
	Tariffication            PostingV4PostingFbsUnfulfilledListResponsePostingsTariffication       `json:"tariffication"`
	Customer                 PostingV4PostingFbsUnfulfilledListResponsePostingsCustomer            `json:"customer"`
	IsPresortable            bool                                                                  `json:"is_presortable"` // `true`，表示商品为错误分级的商品。
	Requirements             PostingV4PostingFbsUnfulfilledListResponsePostingsRequirements        `json:"requirements"`
	DeliveringDate           string                                                                `json:"delivering_date"`        // 货件转移配送的日期。
	DestinationPlaceName     string                                                                `json:"destination_place_name"` // 目的地名称。
	Optional                 PostingV4PostingFbsUnfulfilledListResponsePostingsOptional            `json:"optional"`
	QuantumID                int64                                                                 `json:"quantum_id"` // 经济商品标识符。
	Cancellation             PostingV4PostingFbsUnfulfilledListResponsePostingsCancellation        `json:"cancellation"`
	IsExpress                bool                                                                  `json:"is_express"`              // `true`，表示使用了Ozon Express快速配送。
	PickupCodeVerifiedAt     string                                                                `json:"pickup_code_verified_at"` // 快递员代码校验成功的日期和时间。请通过方法 [/v1/posting/fbs/pick-up-code/verify](#operation/PostingAPI_PostingFBSPickupCodeVerify)校验快递员代码。
	ExternalOrder            PostingV4PostingFbsUnfulfilledListResponsePostingsExternalOrder       `json:"external_order"`
	AnalyticsData            PostingV4PostingFbsUnfulfilledListResponsePostingsAnalyticsData       `json:"analytics_data"`
	DeliveryMethod           PostingV4PostingFbsUnfulfilledListResponsePostingsDeliveryMethod      `json:"delivery_method"`
	IsMultibox               bool                                                                  `json:"is_multibox"`           // 标记货件中是否包含多箱商品，以及是否需要传递其箱数： - `true`——备货前，请通过方法[/v3/posting/multiboxqty/set](#operation/PostingAPI_PostingMultiBoxQtySetV...
	ParentPostingNumber      string                                                                `json:"parent_posting_number"` // 当前货件所拆分自的父货件编号。
	ShipmentDate             string                                                                `json:"shipment_date"`         // 需要在此日期和时间之前完成货件备货。系统会显示推荐的发运时间。超过该时间后将开始适用新费率，相关信息可在`tariffication`字段中获取。
	Addressee                PostingV4PostingFbsUnfulfilledListResponsePostingsAddressee           `json:"addressee"`
	FinancialData            PostingV4PostingFbsUnfulfilledListResponsePostingsFinancialData       `json:"financial_data"`
	InProcessAt              string                                                                `json:"in_process_at"`        // 货件开始处理的日期和时间。
	IsClickAndCollect        bool                                                                  `json:"is_click_and_collect"` // `true`，表示货件通过“到店自提”方式配送。
	MultiBoxQty              int32                                                                 `json:"multi_box_qty"`        // 商品包装所用箱数。
	PostingNumber            string                                                                `json:"posting_number"`       // 货件编号。
	Status                   Status                                                                `json:"status"`               // 货件状态： - `acceptance_in_progress`——接收中； - `arbitration`——仲裁； - `awaiting_approve`——等待确认； - `awaiting_deliver`——等待发运； - `a...
}

type PostingV4PostingFbsUnfulfilledListResponse struct {
	Count    int64                                                `json:"count"`    // 在响应中的元素计数器。
	Cursor   string                                               `json:"cursor"`   // 用于选择下一批数据的指针。
	HasNext  bool                                                 `json:"has_next"` // 若响应中未返回全部货件，则为`true`。
	Postings []PostingV4PostingFbsUnfulfilledListResponsePostings `json:"postings"` // 货件列表。
}

type V1CarriageCreateRequest struct {
	DeliveryMethodID int64  `json:"delivery_method_id"` // 配送方式标识符。
	DepartureDate    string `json:"departure_date"`     // 发运日期。默认值为当前日期。
	AllBlrTraceable  bool   `json:"all_blr_traceable"`  // `true`，表示需要创建包含可追溯商品的发运。
}

type V1SetPostingsRequest struct {
	PostingNumbers []string `json:"posting_numbers"` // 最新货件列表。
	CarriageID     int64    `json:"carriage_id"`     // 发运识别符。
}

// 需要添加到响应中的附加字段。
type PostingV4PostingFbsUnfulfilledListRequestWith struct {
	Barcodes      bool `json:"barcodes"`       // `true`，表示要在响应中添加货件条形码。
	FinancialData bool `json:"financial_data"` // 若为`true`，则在响应中添加财务数据。
	LegalInfo     bool `json:"legal_info"`     // 若为`true`，则在响应中添加法务信息。
	AnalyticsData bool `json:"analytics_data"` // 若为`true`，则在响应中添加分析数据。
}

type PostingV4PostingFbsUnfulfilledListRequest struct {
	Filter   PostingV4PostingFbsUnfulfilledListRequestFilter      `json:"filter"`
	Limit    int64                                                `json:"limit"` // 响应中返回的值数量。
	SortDir  PostingV4PostingFbsUnfulfilledListRequestSortDirEnum `json:"sort_dir"`
	Translit bool                                                 `json:"translit"` // 则启用将地址从西里尔字母转写为拉丁字母。
	With     PostingV4PostingFbsUnfulfilledListRequestWith        `json:"with"`
	Cursor   string                                               `json:"cursor"` // 用于选择下一批数据的指针。
}

type V1AssemblyFbsProductListResponse struct {
	HasNext       bool                                     `json:"has_next"`       // 响应中是否包含全部商品： - `true`——请使用新的 `offset`值重新请求以获取剩余数据； - `false`——响应中已包含所有值。
	Products      []AssemblyFbsProductListResponseProducts `json:"products"`       // 商品列表。
	ProductsCount int32                                    `json:"products_count"` // 商品数量。
}

type PostingPostingProductCancelRequest struct {
	CancelReasonID      int64                             `json:"cancel_reason_id"`      // 货物取消发货原因ID。
	CancelReasonMessage string                            `json:"cancel_reason_message"` // 必填字段。关于取消的其他信息。
	Items               []PostingProductCancelRequestItem `json:"items"`                 // 商品信息。
	PostingNumber       string                            `json:"posting_number"`        // 货运ID。
}

type PostingFbsPostingDeliveredRequest struct {
	PostingNumber []string `json:"posting_number"` // 货件ID。
}

type V4FbsPostingShipPackageV4Request struct {
	PostingNumber string                                    `json:"posting_number"` // 发货号。
	Products      []V4FbsPostingShipPackageV4RequestProduct `json:"products"`       // 商品清单。
}

type Postingv1GetCarriageAvailableListRequest struct {
	DeliveryMethodID int64  `json:"delivery_method_id"` // 按照运输方式筛选。可以使用方法 [/v1/delivery-method/list](#operation/WarehouseAPI_DeliveryMethodList)获取。
	DepartureDate    string `json:"departure_date"`     // 装运日期。默认 —— 当前日期。
}

type Postingv1GetCarriageAvailableListResponse struct {
	Result []GetCarriageAvailableListResponseResult `json:"result"` // 方法操作结果。
}

package beta

type V1ProductFbsSplit struct {
	ProductID int64 `json:"product_id"` // Ozon系统中的商品标识符 — SKU。
	Quantity  int64 `json:"quantity"`   // 数量。
}

// MarkType values
type MarkType string

const (
	MarkTypeMandatoryMark MarkType = "mandatory_mark" // “诚实标志”（Chestny ZNAK）强制性标志；
	MarkTypeJwUin         MarkType = "jw_uin"         // 珠宝制品的唯 一 识别编号；
	MarkTypeImei          MarkType = "imei"           // 移动设备 IMEI。
)

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplarMark struct {
	Errors   []string `json:"errors"`    // 检查控制识别码及其他标识时出现的错误。
	Mark     string   `json:"mark"`      // 标志代码的值。
	MarkType MarkType `json:"mark_type"` // 标志代码类型： - `mandatory_mark` — “诚实标志”（Chestny ZNAK）强制性标志； - `jw_uin` — 珠宝制品的唯 一 识别编号； - `imei` — 移动设备 IMEI。
	Valid    bool     `json:"valid"`     // 检查结果： 若控制识别码及其他标识符合要求，则为`true`。
}

type V5FbsPostingProductExemplarValidateV5ResponseProductExemplar struct {
	Rnpt   string                                                             `json:"rnpt"`   // 商品批次注册号 (Product Batch Registration Number)。
	Valid  bool                                                               `json:"valid"`  // 验证结果。如果样件代码都符合要求，那么结果将为 `true`。
	Errors []string                                                           `json:"errors"` // 样件验证错误。
	GTD    string                                                             `json:"gtd"`    // 货运报关单号码（Cargo Customs Declaration)。
	Marks  []V5FbsPostingProductExemplarValidateV5ResponseProductExemplarMark `json:"marks"`  // 检查控制识别码及其他标识时出现的错误。
}

// 服务应计金额。
type MoneyMoneyAccrued struct {
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
	Currency string `json:"currency"` // 货币单位。
}

// 不关联商品的卖家层面应计项目。
type V1GetFinanceAccrualByDayResponseAccrualNonItemFee struct {
	TypeID  int32             `json:"type_id"` // 应计类型标识符。可通过方法[/v1/finance/accrual/types](#operation/GetFinanceAccrualTypes)获取。
	Accrued MoneyMoneyAccrued `json:"accrued"`
}

// 因折扣积累的积分数。
type MoneyMoneyBonus struct {
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
	Currency string `json:"currency"` // 货币单位。
}

// 合作伙伴计划应计金额。
type MoneyMoneyCoinvestment struct {
	Currency string `json:"currency"` // 货币单位。
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
}

// 计入折扣和加价后的最终佣金。
type MoneyMoneyCommission struct {
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
	Currency string `json:"currency"` // 货币单位。
}

// 已销售金额。
type MoneyMoneySaleAmount struct {
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
	Currency string `json:"currency"` // 货币单位。
}

// 按价目表计算的佣金。
type MoneyMoneySaleCommission struct {
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
	Currency string `json:"currency"` // 货币单位。
}

// 买家价格。
type MoneyMoneySalePrice struct {
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
	Currency string `json:"currency"` // 货币单位。
}

// 单价。
type MoneyMoneySellerPrice struct {
	Currency string `json:"currency"` // 货币单位。
	Amount   string `json:"amount"`   // 金额。如果计入的是销售佣金，数值可以为负数。
}

// 计入折扣和加价后的最终佣金。
type V1GetFinanceAccrualByDayResponseAccrualPostingProductCommission struct {
	SellerPrice     MoneyMoneySellerPrice    `json:"seller_price"`
	Bonus           MoneyMoneyBonus          `json:"bonus"`
	Coinvestment    MoneyMoneyCoinvestment   `json:"coinvestment"`
	Commission      MoneyMoneyCommission     `json:"commission"`
	CommissionRatio string                   `json:"commission_ratio"` // 按类目划分的销售佣金比例。
	SaleAmount      MoneyMoneySaleAmount     `json:"sale_amount"`
	SaleCommission  MoneyMoneySaleCommission `json:"sale_commission"`
	SalePrice       MoneyMoneySalePrice      `json:"sale_price"`
}

// 服务应计总金额。
type MoneyMoneyTotalAccrued struct {
	Amount   string `json:"amount"`   // 金额。数值可以为负数。
	Currency string `json:"currency"` // 货币单位。
}

type V1GetFinanceAccrualByDayResponseAccrualPostingProductDeliveryService struct {
	Accrued MoneyMoneyAccrued `json:"accrued"`
	TypeID  int32             `json:"type_id"`
}

// 配送相关应计项目。
type V1GetFinanceAccrualByDayResponseAccrualPostingProductDelivery struct {
	Services     []V1GetFinanceAccrualByDayResponseAccrualPostingProductDeliveryService `json:"services"` // 其他服务相关应计项目。
	TotalAccrued MoneyMoneyTotalAccrued                                                 `json:"total_accrued"`
}

type V1GetFinanceAccrualByDayResponseAccrualPostingProduct struct {
	Commission V1GetFinanceAccrualByDayResponseAccrualPostingProductCommission `json:"commission"`
	Delivery   V1GetFinanceAccrualByDayResponseAccrualPostingProductDelivery   `json:"delivery"`
	SKU        int64                                                           `json:"sku"` // Ozon系统中的商品标识符——SKU。
}

// 按货件计算的应计项目。
type V1GetFinanceAccrualByDayResponseAccrualPosting struct {
	DeliverySchema string                                                  `json:"delivery_schema"` // 销售模式。
	DeliverySpeed  int32                                                   `json:"delivery_speed"`  // 配送速度。
	Products       []V1GetFinanceAccrualByDayResponseAccrualPostingProduct `json:"products"`        // 货件中的商品数据。
}

// 应计总金额。
type MoneyMoneyTotalAmount struct {
	Amount   string `json:"amount"`   // 金额。
	Currency string `json:"currency"` // 货币单位。
}

// 应计类型： - `UNSPECIFIED`——未指定； - `POSTING`——按货件计算的应计项目； - `ITEM`——按商品计算的应计项目； - `NON_ITEM`——不关联商品的卖家层面应计项目。
type V1GetFinanceAccrualByDayResponseAccrualAccruedCategoryEnum string

type V1GetFinanceAccrualByDayResponseAccrualItemFeesItemFeeFee struct {
	Accrued MoneyMoneyAccrued `json:"accrued"`
	TypeID  int32             `json:"type_id"` // 应计类型标识符。可通过方法[/v1/finance/accrual/types](#operation/GetFinanceAccrualTypes)获取。
}

type V1GetFinanceAccrualByDayResponseAccrualItemFeesItemFee struct {
	Fees []V1GetFinanceAccrualByDayResponseAccrualItemFeesItemFeeFee `json:"fees"` // 应计项目。
	SKU  int64                                                       `json:"sku"`  // Ozon系统中的商品标识符——SKU。
}

// 按商品计算的应计项目。
type V1GetFinanceAccrualByDayResponseAccrualItemFees struct {
	Fees []V1GetFinanceAccrualByDayResponseAccrualItemFeesItemFee `json:"fees"` // 某一商品的应计项目。
}

type V1GetFinanceAccrualByDayResponseAccrual struct {
	ItemFees        V1GetFinanceAccrualByDayResponseAccrualItemFees            `json:"item_fees"`
	NonItemFee      V1GetFinanceAccrualByDayResponseAccrualNonItemFee          `json:"non_item_fee"`
	Posting         V1GetFinanceAccrualByDayResponseAccrualPosting             `json:"posting"`
	TotalAmount     MoneyMoneyTotalAmount                                      `json:"total_amount"`
	AccrualID       int64                                                      `json:"accrual_id"`  // 应计项目标识符。
	UnitNumber      string                                                     `json:"unit_number"` // 例如：货件编号或广告合同编号。
	AccruedCategory V1GetFinanceAccrualByDayResponseAccrualAccruedCategoryEnum `json:"accrued_category"`
	Date            string                                                     `json:"date"` // 应计日期。
}

type V1GetFinanceAccrualByDayResponse struct {
	Accruals []V1GetFinanceAccrualByDayResponseAccrual `json:"accruals"` // 应计项目列表。
	LastID   string                                    `json:"last_id"`  // 页面中最后一个值的标识符。
}

type ErrorData struct {
	Code    string `json:"code"`    // 代码。
	Field   string `json:"field"`   // 原因。
	Message string `json:"message"` // 文字描述。
	Step    int64  `json:"step"`    // 折扣等级。
	Value   string `json:"value"`   // 出现错误字段的值。
}

type V1SetProductStairwayDiscountByQuantityResponseError struct {
	Data []ErrorData `json:"data"` // 错误或警告描述。
	SKU  int64       `json:"sku"`  // Ozon系统中的商品识别符——SKU。
}

type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairwayStep struct {
	Discount int64 `json:"discount"` // 折扣百分比。
	Quantity int64 `json:"quantity"` // 订单中用于应用折扣的商品数量。
	Step     int64 `json:"step"`     // 折扣等级。
}

// 按数量折扣等级信息。
type V1GetProductStairwayDiscountByQuantityResponseStairwaysStairway struct {
	Steps []V1GetProductStairwayDiscountByQuantityResponseStairwaysStairwayStep `json:"steps"` // 折扣等级信息。
}

// 周期内已支付金额。
type V1GetFinanceBalanceV1ResponsePaymentsMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplarMark struct {
	Mark     string   `json:"mark"`      // 标志代码的值。
	MarkType MarkType `json:"mark_type"` // 标志代码类型： - `mandatory_mark` — “诚实标志”（Chestny ZNAK）强制性标志； - `jw_uin` — 珠宝制品的唯 一 识别编号； - `imei` — 移动设备 IMEI。
}

type V5FbsPostingProductExemplarValidateV5RequestProductExemplar struct {
	Marks []V5FbsPostingProductExemplarValidateV5RequestProductExemplarMark `json:"marks"` // 单个商品实例中的控制识别码及其他标识列表。
	Rnpt  string                                                            `json:"rnpt"`  // 商品批次注册号 (Product Batch Registration Number)。
	GTD   string                                                            `json:"gtd"`   // 货运报关单号码（Cargo Customs Declaration)。
}

type V5FbsPostingProductExemplarValidateV5RequestProduct struct {
	Exemplars []V5FbsPostingProductExemplarValidateV5RequestProductExemplar `json:"exemplars"`  // 副本信息。
	ProductID int64                                                         `json:"product_id"` // Ozon系统中的商品ID — SKU。
}

// CheckStatus values
type CheckStatus string

const (
	CheckStatusProcessing CheckStatus = "processing" // 正在检查中；
	CheckStatusFailed     CheckStatus = "failed"     // 系统未能成功处理请求；
	CheckStatusPassed     CheckStatus = "passed"     // 订单已准备好进行备货。
)

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplarMark struct {
	CheckStatus CheckStatus `json:"check_status"` // 检查状态： - `processing`——正在检查中； - `failed`——系统未能成功处理请求； - `passed`——订单已准备好进行备货。
	ErrorCodes  []string    `json:"error_codes"`  // 检查控制识别码及其他标识时出现的错误。
	Mark        string      `json:"mark"`         // 标志代码的值。
	MarkType    MarkType    `json:"mark_type"`    // 标志代码类型： - `mandatory_mark` — “诚实标志”（Chestny ZNAK）强制性标志； - `jw_uin` — 珠宝制品的唯 一 识别编号； - `imei` — 移动设备 IMEI。
}

// 按数量折扣变更状态。可能的取值： - `ERROR`——修改折扣时出错。请再次调用方法 [/v1/product/stairway-discount/by-quantity/set](#operation/ProductAPI_SetPro...
type StatusEnum string

type V1GetProductStairwayDiscountByQuantityResponseStairways struct {
	Stairway V1GetProductStairwayDiscountByQuantityResponseStairwaysStairway `json:"stairway"`
	Status   StatusEnum                                                      `json:"status"`
	Enabled  bool                                                            `json:"enabled"` // `true`，表示数量折扣已启用。
	SKU      int64                                                           `json:"sku"`     // Ozon系统中的商品标识符——SKU。
}

type V1GetProductStairwayDiscountByQuantityResponse struct {
	Stairways []V1GetProductStairwayDiscountByQuantityResponseStairways `json:"stairways"` // 单个商品的按数量折扣信息。
}

// 销售金额。
type V1GetFinanceBalanceV1ResponseSalesMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

// 买家支付的金额。
type V1GetFinanceBalanceV1ResponseRevenueMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

// 合作伙伴忠诚度机制的付款。
type V1GetFinanceBalanceV1ResponsePartnerMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

// 销售金额明细。
type GetFinanceBalanceV1ResponseSalesDetails struct {
	PartnerPrograms    V1GetFinanceBalanceV1ResponsePartnerMoney `json:"partner_programs"`
	PointsForDiscounts string                                    `json:"points_for_discounts"` // 折扣积分。
	Revenue            V1GetFinanceBalanceV1ResponseRevenueMoney `json:"revenue"`
}

// Ozon 代理费金额。
type V1GetFinanceBalanceV1ResponseFeeMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

// 销售应计金额。
type GetFinanceBalanceV1ResponseSales struct {
	Amount        V1GetFinanceBalanceV1ResponseSalesMoney `json:"amount"`
	AmountDetails GetFinanceBalanceV1ResponseSalesDetails `json:"amount_details"`
	Fee           V1GetFinanceBalanceV1ResponseFeeMoney   `json:"fee"`
}

type V1ProductInfoWarehouseStocksRequest struct {
	Cursor      string `json:"cursor"`       // 用于选择下一批数据的指针。
	Limit       int64  `json:"limit"`        // 每页显示的数量。
	WarehouseID int64  `json:"warehouse_id"` // 仓库标识符。
}

type V5FbsPostingProductExemplarValidateV5ResponseProduct struct {
	Error     string                                                         `json:"error"`      // 错误代码。
	Exemplars []V5FbsPostingProductExemplarValidateV5ResponseProductExemplar `json:"exemplars"`  // 副本信息。
	ProductID int64                                                          `json:"product_id"` // Ozon系统中的商品ID — SKU。
	Valid     bool                                                           `json:"valid"`      // 验证结果。如果所有样件的代码都符合要求，那么结果将为 `true`。
}

type V5FbsPostingProductExemplarValidateV5Response struct {
	Products []V5FbsPostingProductExemplarValidateV5ResponseProduct `json:"products"` // 商品清单。
}

// 折扣申请状态： - `ALL`——全部状态， - `NEW`——新建， - `APPROVED`——已批准， - `DECLINED`——已拒绝。
type GetDiscountTaskListV2ResponseTaskDiscountTaskStatusEnum string

type V1GetFinanceAccrualPostingsResponsePostingAccrualsAccrual struct {
	TypeID      int32                 `json:"type_id"`      // 应计类型标识符。可通过方法[/v1/finance/accrual/types](#operation/GetFinanceAccrualTypes)获取。
	AccrualDate string                `json:"accrual_date"` // 应计日期。
	Accrued     MoneyMoneyAccrued     `json:"accrued"`
	Quantity    int32                 `json:"quantity"` // 商品数量。
	SellerPrice MoneyMoneySellerPrice `json:"seller_price"`
	SKU         int64                 `json:"sku"` // Ozon系统中的商品标识符——SKU。
}

// 周期内已应计金额。
type V1GetFinanceBalanceV1ResponseAccruedMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

type V1ProductVisibilitySetResponseItemsErrors struct {
	Code string `json:"code"` // 错误代码。
	SKU  int64  `json:"sku"`  // Ozon系统中的商品标识符——SKU。
}

// 期末余额。
type V1GetFinanceBalanceV1ResponseClosingBalanceMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairwayStep struct {
	Discount int64 `json:"discount"` // 折扣百分比。
	Quantity int64 `json:"quantity"` // 订单中用于应用折扣的商品数量。
	Step     int64 `json:"step"`     // 折扣等级。
}

// 按数量折扣等级信息。
type V1SetProductStairwayDiscountByQuantityRequestStairwaysStairway struct {
	Steps []V1SetProductStairwayDiscountByQuantityRequestStairwaysStairwayStep `json:"steps"` // 折扣等级信息。等级数量可为1到4。
}

type V1SetProductStairwayDiscountByQuantityRequestStairways struct {
	SKU      int64                                                          `json:"sku"` // Ozon系统中的商品标识符——SKU。
	Stairway V1SetProductStairwayDiscountByQuantityRequestStairwaysStairway `json:"stairway"`
	Enabled  bool                                                           `json:"enabled"` // `true`，表示启用折扣。
}

type V1SetProductStairwayDiscountByQuantityRequest struct {
	Stairways        []V1SetProductStairwayDiscountByQuantityRequestStairways `json:"stairways"`         // 多个商品的按数量折扣信息。
	SuppressWarnings bool                                                     `json:"suppress_warnings"` // 传递 `true` 可忽略警告并设置折扣。
}

// 退货金额。
type V1GetFinanceBalanceV1ResponseReturnsMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

// 退货金额明细。
type GetFinanceBalanceV1ResponseReturnsDetails struct {
	PartnerPrograms    V1GetFinanceBalanceV1ResponsePartnerMoney `json:"partner_programs"`
	PointsForDiscounts string                                    `json:"points_for_discounts"` // 折扣积分。
	Revenue            V1GetFinanceBalanceV1ResponseRevenueMoney `json:"revenue"`
}

// 退货应计金额。
type GetFinanceBalanceV1ResponseReturns struct {
	Amount        V1GetFinanceBalanceV1ResponseReturnsMoney `json:"amount"`
	AmountDetails GetFinanceBalanceV1ResponseReturnsDetails `json:"amount_details"`
	Fee           V1GetFinanceBalanceV1ResponseFeeMoney     `json:"fee"`
}

// 其他服务的应计金额。
type V1GetFinanceBalanceV1ResponseServicesMoney struct {
	CurrencyCode string  `json:"currency_code"` // 货币单位。
	Value        float64 `json:"value"`         // 金额。
}

type GetFinanceBalanceV1ResponseService struct {
	Amount V1GetFinanceBalanceV1ResponseServicesMoney `json:"amount"`
	Name   string                                     `json:"name"` // 服务的系统名称。
}

// 收入和支出信息。
type GetFinanceBalanceV1ResponseCashflows struct {
	Services []GetFinanceBalanceV1ResponseService `json:"services"` // 其他服务的应计金额。
	Returns  GetFinanceBalanceV1ResponseReturns   `json:"returns"`
	Sales    GetFinanceBalanceV1ResponseSales     `json:"sales"`
}

type ExemplarMark struct {
	Mark     string   `json:"mark"`      // 标志代码的值。
	MarkType MarkType `json:"mark_type"` // 标志代码类型： - `mandatory_mark` — “诚实标志”（Chestny ZNAK）强制性标志； - `jw_uin` — 珠宝制品的唯 一 识别编号； - `imei` — 移动设备 IMEI。
}

type Exemplar struct {
	Marks        []ExemplarMark `json:"marks"`          // 单个商品实例中的控制识别码及其他标识列表。
	Rnpt         string         `json:"rnpt"`           // 商品批次注册号 (Product Batch Registration Number)。
	ExemplarID   int64          `json:"exemplar_id"`    // 样件识别码。
	GTD          string         `json:"gtd"`            // 货运报关单号码（Cargo Customs Declaration)。
	IsGTDAbsent  bool           `json:"is_gtd_absent"`  // 不需要指出货运报关单（Cargo Customs Declaration）号码的标志。
	IsRnptAbsent bool           `json:"is_rnpt_absent"` // 不需要指出商品批次注册号(Product Batch Registration Number)的标志。
}

type PostingProductExemplarCreateOrGetV6ResponseProduct struct {
	Exemplars               []Exemplar `json:"exemplars"`                  // 副本信息。
	HasImei                 bool       `json:"has_imei"`                   // 存在 IMEI。 若存在 IMEI，则为`true`。
	IsGTDNeeded             bool       `json:"is_gtd_needed"`              // 说明需要递交产品和货件的货物报关单号码。
	IsMandatoryMarkPossible bool       `json:"is_mandatory_mark_possible"` // 是否可以填写“诚实标志”（Chestny ZNAK）信息。
	IsRnptNeeded            bool       `json:"is_rnpt_needed"`             // 说明需要递交商品批次号码。
	ProductID               int64      `json:"product_id"`                 // Ozon系统中的商品ID — SKU。
	Quantity                int32      `json:"quantity"`                   // 样件数量。
	IsJwUinNeeded           bool       `json:"is_jw_uin_needed"`           // 是否需要提供珠宝制品的唯一识别编号。
	IsMandatoryMarkNeeded   bool       `json:"is_mandatory_mark_needed"`   // 说明需要将标志递交给“诚实标志”。
}

type V6FbsPostingProductExemplarCreateOrGetV6Response struct {
	MultiBoxQty   int32                                                `json:"multi_box_qty"`  // 商品包装的箱子数量。
	PostingNumber string                                               `json:"posting_number"` // 发货号。
	Products      []PostingProductExemplarCreateOrGetV6ResponseProduct `json:"products"`       // 商品清单。
}

type ExemplarsMarks struct {
	Mark     string   `json:"mark"`      // 标志代码的值。
	MarkType MarkType `json:"mark_type"` // 标志代码类型： - `mandatory_mark` — “诚实标志”（Chestny ZNAK）强制性标志； - `jw_uin` — 珠宝制品的唯 一 识别编号； - `imei` — 移动设备 IMEI。
}

type PostingProductExemplarSetV6RequestExemplars struct {
	Rnpt         string           `json:"rnpt"`           // 商品批次注册号 (Product Batch Registration Number)。
	ExemplarID   int64            `json:"exemplar_id"`    // 样件识别码。
	GTD          string           `json:"gtd"`            // 货运报关单号码（Cargo Customs Declaration)。
	IsGTDAbsent  bool             `json:"is_gtd_absent"`  // 不需要指出货运报关单（Cargo Customs Declaration）号码的标志。
	IsRnptAbsent bool             `json:"is_rnpt_absent"` // 不需要指出商品批次注册号(Product Batch Registration Number)的标志。
	Marks        []ExemplarsMarks `json:"marks"`          // 单个商品实例中的控制识别码及其他标识列表。
}

type PostingProductExemplarSetV6RequestProducts struct {
	Exemplars []PostingProductExemplarSetV6RequestExemplars `json:"exemplars"`  // 副本信息。
	ProductID int64                                         `json:"product_id"` // Ozon系统中的商品ID — SKU。
}

type V1GetFinanceAccrualPostingsResponsePostingAccruals struct {
	Accruals      []V1GetFinanceAccrualPostingsResponsePostingAccrualsAccrual `json:"accruals"`       // 应计项目列表。
	PostingNumber string                                                      `json:"posting_number"` // 货件编号。
}

type V1GetProductStairwayDiscountByQuantityRequest struct {
	Skus []string `json:"skus"` // 需要返回内容评级的商品SKU列表。
}

// 应计项目信息。
type V1GetFinanceAccrualTypesResponseAccrualType struct {
	Description string `json:"description"` // 应计项目说明。
	ID          int32  `json:"id"`          // 应计项目标识符。
	Name        string `json:"name"`        // 应计项目名称。
}

type V1GetFinanceAccrualTypesResponse struct {
	AccrualTypes []V1GetFinanceAccrualTypesResponseAccrualType `json:"accrual_types"` // 应计项目相关信息。
}

// 期初余额。
type V1GetFinanceBalanceV1ResponseOpeningBalanceMoney struct {
	Value        float64 `json:"value"`         // 金额。
	CurrencyCode string  `json:"currency_code"` // 货币单位。
}

// 周期内的余额总体数据。
type GetFinanceBalanceV1ResponseTotal struct {
	Accrued        V1GetFinanceBalanceV1ResponseAccruedMoney        `json:"accrued"`
	ClosingBalance V1GetFinanceBalanceV1ResponseClosingBalanceMoney `json:"closing_balance"`
	OpeningBalance V1GetFinanceBalanceV1ResponseOpeningBalanceMoney `json:"opening_balance"`
	Payments       []V1GetFinanceBalanceV1ResponsePaymentsMoney     `json:"payments"` // 周期内的付款。
}

type V1FbsPostingProductExemplarUpdateRequest struct {
	PostingNumber string `json:"posting_number"` // 发货号。
}

type V1SetProductStairwayDiscountByQuantityResponse struct {
	Accepted bool                                                  `json:"accepted"` // `true`，表示请求已接收。请使用方法[/v1/product/stairway-discount/by-quantity/get](#operation/ProductAPI_GetProductStairwayDiscountByQu...
	Errors   []V1SetProductStairwayDiscountByQuantityResponseError `json:"errors"`   // 错误描述。
	Warnings []V1SetProductStairwayDiscountByQuantityResponseError `json:"warnings"` // 警告描述。
}

// 商品展示在哪些橱窗中： - `UNSPECIFIED`——未指定； - `OZON`——仅在Ozon展示； - `SELECT`——仅在Select展示； - `OZON_SELECT`——在Select和Ozon展示； - `NONE`—...
type V1ProductVisibilityInfoResponseItemShowcasesVisibilityEnum string

type V1ProductVisibilityInfoResponseItem struct {
	ShowcasesVisibility V1ProductVisibilityInfoResponseItemShowcasesVisibilityEnum `json:"showcases_visibility"`
	SKU                 int64                                                      `json:"sku"` // 商品在Ozon系统中的标识符——SKU。
}

type V5FbsPostingProductExemplarValidateV5Request struct {
	PostingNumber string                                                `json:"posting_number"` // 发货号。
	Products      []V5FbsPostingProductExemplarValidateV5RequestProduct `json:"products"`       // 商品清单。
}

type V5FbsPostingProductExemplarStatusV5ResponseProductExemplar struct {
	RnptErrorCodes  []string                                                         `json:"rnpt_error_codes"`  // 商品批次注册编号验证错误代码。
	ExemplarID      int64                                                            `json:"exemplar_id"`       // 样件识别码。
	GTD             string                                                           `json:"gtd"`               // 货运报关单号码（Cargo Customs Declaration)。
	GTDCheckStatus  string                                                           `json:"gtd_check_status"`  // 货物报关验证状态。
	IsRnptAbsent    bool                                                             `json:"is_rnpt_absent"`    // 不需要指出商品批次注册号(Product Batch Registration Number)的标志。
	Rnpt            string                                                           `json:"rnpt"`              // 商品批次注册号 (Product Batch Registration Number)。
	RnptCheckStatus string                                                           `json:"rnpt_check_status"` // 商品批次注册编号验证状态。
	GTDErrorCodes   []string                                                         `json:"gtd_error_codes"`   // 货物报关验证错误代码。
	IsGTDAbsent     bool                                                             `json:"is_gtd_absent"`     // 这说明没有输入货物报关单号。
	Marks           []V5FbsPostingProductExemplarStatusV5ResponseProductExemplarMark `json:"marks"`             // 单个商品实例中的控制识别码及其他标识列表。
}

type V5FbsPostingProductExemplarStatusV5ResponseProduct struct {
	Exemplars []V5FbsPostingProductExemplarStatusV5ResponseProductExemplar `json:"exemplars"`  // 副本信息。
	ProductID int64                                                        `json:"product_id"` // Ozon系统中的商品ID — SKU。
}

// Status values
type Status string

const (
	StatusShipAvailable       Status = "ship_available"        // 可以备货，
	StatusShipNotAvailable    Status = "ship_not_available"    // 无法备货，
	StatusValidationInProcess Status = "validation_in_process" // 样件正在验证中，
	StatusUpdateAvailable     Status = "update_available"      // 可以编辑商品实例信息，
	StatusUpdateNotAvailable  Status = "update_not_available"  // 无法编辑商品实例信息。
)

type V5FbsPostingProductExemplarStatusV5Response struct {
	PostingNumber string                                               `json:"posting_number"` // 发货号。
	Products      []V5FbsPostingProductExemplarStatusV5ResponseProduct `json:"products"`       // 商品清单。
	Status        Status                                               `json:"status"`         // 所有样件和备货可用性的验证状态： - `ship_available`——可以备货， - `ship_not_available`——无法备货， - `validation_in_process`——样件正在验证中， - `update_a...
}

// 商品展示在哪些橱窗中： - `UNSPECIFIED`——未指定； - `OZON`——仅在Ozon展示； - `SELECT`——仅在Select展示； - `OZON_SELECT`——在Select和Ozon展示； - `NONE`—...
type V1ProductVisibilitySetResponseItemsShowcasesVisibilityEnum string

// 余额报告。
type V1GetFinanceBalanceV1Response struct {
	Cashflows GetFinanceBalanceV1ResponseCashflows `json:"cashflows"`
	Total     GetFinanceBalanceV1ResponseTotal     `json:"total"`
}

// 商品在Ozon Select上的销售权限： - `UNSPECIFIED`——未指定； - `RESTRICTED`——商品不可销售； - `ALLOWED`——商品可以销售。
type V1ProductVisibilitySetResponseItemsSelectPermissionEnum string

// 卖家设置的可见性取值： - `UNSPECIFIED`——未指定； - `OZON`——仅在Ozon展示； - `SELECT`——仅在Select展示； - `OZON_SELECT`——在Select和Ozon展示。
type V1ProductVisibilitySetResponseItemsSellerItemPlacementEnum string

type V1ProductVisibilitySetResponseItemsSellerItemPlacementListEnum string

type V1ProductVisibilitySetResponseItemsShowcasesVisibilityListEnum string

// SellerItemPlacementList values
type SellerItemPlacementList string

const (
	SellerItemPlacementListUnspecified SellerItemPlacementList = "UNSPECIFIED" // 未指定；
	SellerItemPlacementListOzon        SellerItemPlacementList = "OZON"        // 仅在Ozon展示；
	SellerItemPlacementListSelect      SellerItemPlacementList = "SELECT"      // 仅在Select展示。
)

// ShowcasesVisibilityList values
type ShowcasesVisibilityList string

const (
	ShowcasesVisibilityListUnspecified ShowcasesVisibilityList = "UNSPECIFIED" // 未指定；
	ShowcasesVisibilityListOzon        ShowcasesVisibilityList = "OZON"        // 仅在Ozon展示；
	ShowcasesVisibilityListSelect      ShowcasesVisibilityList = "SELECT"      // 仅在Select展示。
)

type V1ProductVisibilitySetResponseItems struct {
	Warnings                []string                                                         `json:"warnings"` // 警告。
	SelectPermission        V1ProductVisibilitySetResponseItemsSelectPermissionEnum          `json:"select_permission"`
	SellerItemPlacement     V1ProductVisibilitySetResponseItemsSellerItemPlacementEnum       `json:"seller_item_placement"`
	SellerItemPlacementList []V1ProductVisibilitySetResponseItemsSellerItemPlacementListEnum `json:"seller_item_placement_list"` // 卖家设置的可见性取值列表： - `UNSPECIFIED`——未指定； - `OZON`——仅在Ozon展示； - `SELECT`——仅在Select展示。
	ShowcasesVisibility     V1ProductVisibilitySetResponseItemsShowcasesVisibilityEnum       `json:"showcases_visibility"`
	ShowcasesVisibilityList []V1ProductVisibilitySetResponseItemsShowcasesVisibilityListEnum `json:"showcases_visibility_list"` // 商品展示所在的橱窗列表： - `UNSPECIFIED`——未指定； - `OZON`——仅在Ozon展示； - `SELECT`——仅在Select展示。
	SKU                     int64                                                            `json:"sku"`                       // Ozon系统中的商品标识符——SKU。
}

type V1ProductVisibilitySetResponse struct {
	Items       []V1ProductVisibilitySetResponseItems       `json:"items"`        // 商品可见性信息。
	ItemsErrors []V1ProductVisibilitySetResponseItemsErrors `json:"items_errors"` // 存在错误的商品。
}

type V6FbsPostingProductExemplarCreateOrGetV6Request struct {
	PostingNumber string `json:"posting_number"` // 发货号。
}

// 商品投放平台： - `OZON`——仅在Ozon； - `SELECT`——仅在Select； - `OZON_SELECT`——在Select和Ozon。
type V1ProductVisibilitySetRequestItemPlacementPlacementEnum string

type V1ProductVisibilitySetRequestItemPlacement struct {
	Placement V1ProductVisibilitySetRequestItemPlacementPlacementEnum `json:"placement"`
	SKU       int64                                                   `json:"sku"` // Ozon系统中的商品标识符——SKU。
}

type V1ProductVisibilitySetRequest struct {
	ItemPlacement []V1ProductVisibilitySetRequestItemPlacement `json:"item_placement"` // 商品可见性信息。
}

type V1ProductVisibilityInfoResponse struct {
	Items []V1ProductVisibilityInfoResponseItem `json:"items"` // 商品列表。
}

type V1DescriptionCategoryTipsRequest struct {
	TypeID []string `json:"type_id"` // 商品类型标识符。可通过方法 [/v1/description-category/tree](#operation/DescriptionCategoryAPI_GetTree)获取。
}

// DateTo values
type DateTo string

const (
	DateToYyyyMMDD DateTo = "YYYY-MM-DD"
	DateToDateFrom DateTo = "date_from"
	DateToDateTo   DateTo = "date_to"
)

type V1GetFinanceBalanceV1Request struct {
	DateFrom string `json:"date_from"` // 报告期开始日期，格式为 `YYYY-MM-DD`。
	DateTo   DateTo `json:"date_to"`   // 报告期结束日期，格式为 `YYYY-MM-DD`。`date_from` 与 `date_to` 之间的最⻓间隔为30 天。
}

type V5FbsPostingProductExemplarStatusV5Request struct {
	PostingNumber string `json:"posting_number"` // 发货号。
}

type DescriptionCategoryTipsResponseResult struct {
	TypeID    int64    `json:"type_id"`    // 商品类型标识符。
	ImagesURL []string `json:"images_url"` // 相似商品图片链接。
	InfoURL   string   `json:"info_url"`   // 指向Ozon商品橱窗的链接，其中包含相似商品及其信息。
}

type InfoWarehouseStocksResponseStocks struct {
	OfferID     string `json:"offer_id"`     // 卖家系统中的商品标识符——`offer_id`。
	Present     int64  `json:"present"`      // 仓库中的商品总数量。
	ProductID   int64  `json:"product_id"`   // Ozon系统中商品的标识符 — `product_id`。
	Reserved    int64  `json:"reserved"`     // 仓库中已预留商品的数量。
	SKU         int64  `json:"sku"`          // Ozon系统中的商品标识符——SKU。
	UpdatedAt   string `json:"updated_at"`   // 商品的最后更新时间。
	WarehouseID int64  `json:"warehouse_id"` // 仓库标识符。
	FreeStock   int64  `json:"free_stock"`   // 仓库中可供下单的商品数量。
}

// 申请自动审核信息。
type TaskAutoModeratedInfo struct {
	MaxPercent float64 `json:"max_percent"` // 可批准的最大折扣。
	MaxPrice   float64 `json:"max_price"`   // 申请中的价格。
	MinPercent float64 `json:"min_percent"` // 可批准的最小折扣。
	MinPrice   float64 `json:"min_price"`   // 可批准的最低价格。
}

type GetDiscountTaskListV2ResponseTask struct {
	ApprovedQuantityMax  int64                                                   `json:"approved_quantity_max"`  // 批准的最大商品数量。
	CreatedAt            string                                                  `json:"created_at"`             // 申请创建日期。
	EditedTillDuration   int64                                                   `json:"edited_till_duration"`   // 可修改决定的时间（秒）。
	FirstName            string                                                  `json:"first_name"`             // 处理申请的卖家员工名字。
	IsAutoModerated      bool                                                    `json:"is_auto_moderated"`      // `true`，表示审核为自动审核。
	MinAutoPrice         float64                                                 `json:"min_auto_price"`         // 自动应用折扣与促销后的最低价格值。
	RequestedDiscount    float64                                                 `json:"requested_discount"`     // 折扣百分比。
	RequestedQuantityMax int64                                                   `json:"requested_quantity_max"` // 请求的最大商品数量。
	AutoModeratedInfo    TaskAutoModeratedInfo                                   `json:"auto_moderated_info"`
	EndAtDuration        int64                                                   `json:"end_at_duration"` // 申请有效期结束时间（秒）。
	ID                   int64                                                   `json:"id"`              // 申请标识符。
	LastName             string                                                  `json:"last_name"`       // 处理申请的卖家员工姓氏。
	Patronymic           string                                                  `json:"patronymic"`      // 处理申请的卖家员工父名（中间名）。
	SKU                  int64                                                   `json:"sku"`             // Ozon 系统中的商品标识符——SKU。
	Status               GetDiscountTaskListV2ResponseTaskDiscountTaskStatusEnum `json:"status"`
	ApprovedPrice        float64                                                 `json:"approved_price"`    // 批准价格。
	EditedTill           string                                                  `json:"edited_till"`       // 可修改决定的时间。
	Email                string                                                  `json:"email"`             // 处理申请的卖家员工邮箱地址。
	ModeratedAt          string                                                  `json:"moderated_at"`      // 审核日期：查看、批准或拒绝申请的日期。
	Name                 string                                                  `json:"name"`              // 商品名称。
	ApprovedDiscount     float64                                                 `json:"approved_discount"` // 卖家批准的折扣金额（卢布）。如果卖家未批准申请，请传入 `0`。
	EndAt                string                                                  `json:"end_at"`            // 申请有效期结束时间。
	OriginalPrice        float64                                                 `json:"original_price"`    // 商品在所有折扣前的价格。
	ReductionFactor      float64                                                 `json:"reduction_factor"`  // 创建申请时买家价格与卖家价格之间的差值。
	RequestedPrice       float64                                                 `json:"requested_price"`   // 申请价格。
}

type V1GetFinanceAccrualPostingsResponse struct {
	PostingAccruals []V1GetFinanceAccrualPostingsResponsePostingAccruals `json:"posting_accruals"` // 按货件统计的应计项目列表。
}

type V1ProductInfoWarehouseStocksResponse struct {
	Cursor  string                              `json:"cursor"`   // 用于选择下一批数据的指针。 如果该参数为空，则没有更多数据了。
	HasNext bool                                `json:"has_next"` // 标记是否返回了所有商品： - `true`——请使用不同的`cursor`值重新请求，以获取剩余的值； - `false`——响应中已包含所有值。
	Stocks  []InfoWarehouseStocksResponseStocks `json:"stocks"`   // 商品库存信息。
}

type V1DescriptionCategoryTipsResponse struct {
	Result []DescriptionCategoryTipsResponseResult `json:"result"` // 提示列表。
}

type V1PostingFbsSplitResponsePosting struct {
	PostingNumber string              `json:"posting_number"` // 货件编号。
	Products      []V1ProductFbsSplit `json:"products"`       // 货件中的商品列表。
}

// 原始货件的信息。
type V1PostingFbsSplitResponsePostingParent struct {
	PostingNumber string              `json:"posting_number"` // 原始货件编号。
	Products      []V1ProductFbsSplit `json:"products"`       // 货件中的商品列表。
}

type V1PostingFbsSplitResponse struct {
	ParentPosting V1PostingFbsSplitResponsePostingParent `json:"parent_posting"`
	Postings      []V1PostingFbsSplitResponsePosting     `json:"postings"` // 订单被拆分后的货件列表。
}

// 折扣申请状态： - `ALL`——全部状态， - `NEW`——新建， - `APPROVED`——已批准， - `DECLINED`——已拒绝。
type V2GetDiscountTaskListV2RequestDiscountTaskStatusEnum string

type V2GetDiscountTaskListV2Request struct {
	Status V2GetDiscountTaskListV2RequestDiscountTaskStatusEnum `json:"status"`
	LastID int64                                                `json:"last_id"` // 页面上最后一个值的标识符。首次请求请留空。
	Limit  int64                                                `json:"limit"`   // 每页最大申请数量。
}

type V1ProductVisibilityInfoRequest struct {
	Skus []string `json:"skus"` // Ozon系统中的商品标识符—— SKU。
}

type V2GetDiscountTaskListV2Response struct {
	Tasks []GetDiscountTaskListV2ResponseTask `json:"tasks"` // 申请列表。
}

type V6FbsPostingProductExemplarSetV6Request struct {
	Products      []PostingProductExemplarSetV6RequestProducts `json:"products"`       // 商品清单。
	MultiBoxQty   int32                                        `json:"multi_box_qty"`  // 商品包装的箱子数量。
	PostingNumber string                                       `json:"posting_number"` // 发货号。
}

type V1GetFinanceAccrualPostingsRequest struct {
	PostingNumbers []string `json:"posting_numbers"` // 货件编号。
}

type V1GetFinanceAccrualByDayRequest struct {
	Date   string `json:"date"`    // 应计日期。最早可查询日期为2022年1月1日。
	LastID string `json:"last_id"` // 页面上最后一个值的标识符。首次请求请留空。 要获取后续值，请指定上一次请求响应中的 `last_id`。
}

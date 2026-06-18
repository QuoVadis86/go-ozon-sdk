package returns

type V1ReturnsRfbsActionSetRequest struct {
	RejectionReasonID int32 `json:"rejection_reason_id"` // 取消原因的标识符。 对于 `id: -1` 和 `id: -10`，备注为必填项。 获取可用取消原因 `returns.rejection_reason`，请使用方法 [/v2/returns/rfbs/get](#operation/RF...
	ReturnForBackWay float64 `json:"return_for_back_way"` // 退还给买家的商品运费金额。 负值将被视为 `0`。
	ReturnID int64 `json:"return_id"` // 退货申请的标识符。
	Comment string `json:"comment"` // 卖家评论。 对于 `id: -1` 和 `id: -10`，备注为必填项。
	CompensationAmount float64 `json:"compensation_amount"` // 赔偿金额。 对于 `id: 1020`，备注也为必填项。
	ID int32 `json:"id"` // 操作标识符。 获取可用操作 `returns.available_actions` ，请使用方法 [/v2/returns/rfbs/get](#operation/RFBSReturnsAPI_ReturnsRfbsGetV2)。
}

type V2ReturnsRfbsCompensateRequest struct {
	ReturnID int64 `json:"return_id"` // 退货申请的标识符。
	CompensationAmount string `json:"compensation_amount"` // 赔偿金额。
}

type V2ReturnsRfbsGetResponse struct {
	Returns interface{} `json:"returns"`
}

type V1GetReturnsListResponse struct {
	Returns []interface{} `json:"returns"` // 退货信息。
	HasNext bool `json:"has_next"` // 如果卖家有其他退货，显示`true`。
}

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnID int64 `json:"return_id"` // 退货申请的标识符。
	ReturnForBackWay int64 `json:"return_for_back_way"` // 退还给买家的商品运费金额。
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"` // 退货申请的标识符。
}

type V2ReturnsRfbsListRequest struct {
	LastID int32 `json:"last_id"` // 页面上最后一个值的标识符——`return_id`。在第一次请求时，请将此字段留空。
	Limit int32 `json:"limit"` // 响应中的值数量。
	Filter interface{} `json:"filter"`
}

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID int64 `json:"return_id"` // 退货申请的标识符。
	ReturnMethodDescription string `json:"return_method_description"` // 商品退货方式。
}

type V1GetReturnsListRequest struct {
	Limit int32 `json:"limit"` // 加载的退货数量。
	LastID int64 `json:"last_id"` // 最后加载的退货ID。
	Filter interface{} `json:"filter"`
}

type V2GetConditionalCancellationListV2Response struct {
	Counter int64 `json:"counter"` // `ON_APPROVAL` 状态申请的计数器。
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。 要获取后续值，请指定上一次请求响应中的 `last_id`。
	Result []interface{} `json:"result"` // 取消申请的详细信息。
}

type V2ReturnsRfbsRejectRequest struct {
	Comment string `json:"comment"` // 备注。 如果 [/v2/returns/rfbs/get](#operation/RFBSReturnsAPI_ReturnsRfbsGetV2) 方法的响应中 `rejection_reason.is_comment_required` ...
	RejectionReasonID int64 `json:"rejection_reason_id"` // 取消原因的标识符。 从 [/v2/returns/rfbs/get](#operation/RFBSReturnsAPI_ReturnsRfbsGetV2) 响应中获取的原因列表中传递标识符，参数为 `rejection_reason`。
	ReturnID int64 `json:"return_id"` // 退货申请的标识符。
}

type V2GetConditionalCancellationListV2Request struct {
	Filters interface{} `json:"filters"`
	LastID int64 `json:"last_id"` // 页面上最后一个值的标识符。在首次请求时此字段留空。 要获取后续值，请指定上一次请求响应中的 `last_id`。
	Limit int32 `json:"limit"` // 响应中包含的申请总数。
	With interface{} `json:"with"`
}

type V2ReturnsRfbsListResponse struct {
	Returns interface{} `json:"returns"`
}

type V2ReturnsRfbsGetRequest struct {
	ReturnID int64 `json:"return_id"` // 退货申请标识符。通过方法 [/v2/returns/rfbs/list](#operation/RFBSReturnsAPI_ReturnsRfbsListV2) 获取。
}

package chat

// 聊天状态： - `UNSPECIFIED` — 未指定， - `All` — 所有聊天， - `OPENED` — 开放的聊天， - `CLOSED` — 不开放的聊天。
type ChatInfoChatStatusEnum string

// 聊天类型：' - `UNSPECIFIED` — 未指定； - `SELLER_SUPPORT` — 与帮助中心聊天； - `BUYER_SELLER` — 与买家聊天； - `BUYER_SELLER_SELECT` — 与买家的聊天：关...
type ChatInfoChatTypeEnum string

// 聊天信息。
type V3ChatDetailsInfo struct {
	CreatedAt string `json:"created_at"` // 聊天的创建日期。
	ChatID string `json:"chat_id"` // 聊天识别码。
	ChatStatus ChatInfoChatStatusEnum `json:"chat_status"`
	ChatType ChatInfoChatTypeEnum `json:"chat_type"`
}

type V3ChatInfo struct {
	LastMessageID int64 `json:"last_message_id"` // 最后一条聊天信息的识别码。
	UnreadCount int64 `json:"unread_count"` // 聊天中未读消息的数量。
	Chat V3ChatDetailsInfo `json:"chat"`
	FirstUnreadMessageID int64 `json:"first_unread_message_id"` // 第一条未读聊天信息的识别码。
}

type V3ChatListResponse struct {
	Chats []V3ChatInfo `json:"chats"` // 聊天数据。
	TotalUnreadCount int64 `json:"total_unread_count"` // 未读信息总数。
	Cursor string `json:"cursor"` // 后续数据的选择标志。
	HasNext bool `json:"has_next"` // 表示响应中未包含全部聊天: - `true` — 使用新的 `cursor` 参数重新发送请求以获取剩余聊天; - `false` — 响应中已包含匹配请求筛选器的所有聊天。
}

// 图片审核状态： - `SUCCESS` — 审核成功了; - `MODERATION` — 审核中; - `FAILED` — 审核失败了.
type ChatMessageModerateImageStatus string

// 按聊天过滤。
// ChatStatus values
type ChatStatus string
const (
	ChatStatusAll ChatStatus = "All" // 所有聊天。
	ChatStatusOpened ChatStatus = "Opened" // 开放的聊天。
	ChatStatusClosed ChatStatus = "Closed" // 不开放的聊天。
)

type V3ChatListRequestFilter struct {
	ChatStatus ChatStatus `json:"chat_status"` // 按聊天状态过滤： - `All` — 所有聊天。 - `Opened` — 开放的聊天。 - `Closed` — 不开放的聊天。 默认值：`All`。
	UnreadOnly bool `json:"unread_only"` // 按有未读信息的聊天过滤。
}

type ChatChatSendMessageResponse struct {
	Result string `json:"result"` // 请求的处理结果。
}

// 消息过滤器。
type ChatHistoryRequestFilter struct {
	MessageIds []string `json:"message_ids"` // 消息标识符。
}

type ChatChatStartRequest struct {
	PostingNumber string `json:"posting_number"` // 发货识别码。
}

type ChatChatSendFileResponse struct {
	Result string `json:"result"` // 请求的处理结果。
}

// 关于聊天参与者的信息。
// Type values
type Type string
const (
	TypeCustomer Type = "Customer" // 买家，
	TypeSeller Type = "Seller" // 卖家，
	TypeCrm Type = "Crm" // 系统信息，
	TypeCourier Type = "Courier" // 快递员，
	TypeSupport Type = "Support" // 客服，
	TypeNotificationUser Type = "NotificationUser" // 通知。
)

type V3User struct {
	ID string `json:"id"` // 聊天参与者的身份。
	Type Type `json:"type_"` // 聊天参与者类型： - `Customer` — 买家， - `Seller` — 卖家， - `Crm` — 系统信息， - `Courier` — 快递员， - `Support` — 客服， - `NotificationUser` —...
}

// 聊天室信息。
type ChatMessageContext struct {
	OrderNumber string `json:"order_number"` // 订单编号。
	SKU string `json:"sku"` // Ozon系统中的商品识别码是SKU。
}

type V3ChatMessage struct {
	User V3User `json:"user"`
	Context ChatMessageContext `json:"context"`
	CreatedAt string `json:"created_at"` // 信息创建日期。
	Data []string `json:"data"` // Markdown格式的带有信息内容的数组。
	IsImage bool `json:"is_image"` // 消息包含图片的标志。
	IsRead bool `json:"is_read"` // 表示信息已读。
	MessageId int64 `json:"messageId"` // 信息识别码。
	ModerateImageStatus ChatMessageModerateImageStatus `json:"moderate_image_status"`
}

type V3ChatHistoryResponse struct {
	HasNext bool `json:"has_next"` // 表示不是所有信息都在答复中返回。
	Messages []V3ChatMessage `json:"messages"` // 根据请求正文中的`direction`参数排序的信息数组。
}

// 方法运行结果。
type ChatStartResponseResult struct {
	ChatID string `json:"chat_id"` // 聊天识别码。
}

type ChatChatStartResponse struct {
	Result ChatStartResponseResult `json:"result"`
}

type V3ChatList struct {
	Filter V3ChatListRequestFilter `json:"filter"`
	Limit int64 `json:"limit"` // 回答中值的数量。默认值为30。最大值是100。
	Cursor string `json:"cursor"` // 后续数据的选择标志。
}

type ChatChatSendMessageRequest struct {
	ChatID string `json:"chat_id"` // 聊天识别码。
	Text string `json:"text"` // plain文本格式的信息文本1到1000个字符。
}

type ChatRead struct {
	ChatID string `json:"chat_id"` // 聊天识别码。
	FromMessageID int64 `json:"from_message_id"` // 信息识别码。
}

type V2ChatReadResponse struct {
	UnreadCount int64 `json:"unread_count"` // 聊天中未读消息的数量。
}

type ChatChatSendFileRequest struct {
	Base64Content string `json:"base64_content"` // 文件为 base64 行形式。
	ChatID string `json:"chat_id"` // 聊天识别码。
	Name string `json:"name"` // 带有扩展名的文件名。
}

// FromMessageID values
type FromMessageID string
const (
	FromMessageIDDirectionForward FromMessageID = "direction = Forward"
	FromMessageIDFromMessageId FromMessageID = "from_message_id"
)

// Direction values
type Direction string
const (
	DirectionForward Direction = "Forward" // 从旧到新。
	DirectionBackward Direction = "Backward" // 从新到旧。
	DirectionLimit Direction = "limit"
)

type V3ChatHistoryRequest struct {
	ChatID string `json:"chat_id"` // 聊天识别码。
	Direction Direction `json:"direction"` // 信息排序方向： - `Forward` — 从旧到新。 - `Backward` — 从新到旧。 默认值是 — `Backward`。消息的数量可以在 `limit`参数中设置。
	Filter ChatHistoryRequestFilter `json:"filter"`
	FromMessageID int64 `json:"from_message_id"` // 从该信息开始整理聊天记录的消息识别码。默认为从最后一条可见信息。 当 `direction = Forward` 时，`from_message_id` 参数为必填。
	Limit int64 `json:"limit"` // 答复的信息数量。默认设置为50。最大值是1000。
}

package chat

type ChatChatStartResponse struct {
	Result interface{} `json:"result"`
}

type V3ChatHistoryRequest struct {
	ChatID string `json:"chat_id"` // 聊天识别码。
	Direction string `json:"direction"` // 信息排序方向： - `Forward` — 从旧到新。 - `Backward` — 从新到旧。 默认值是 — `Backward`。消息的数量可以在 `limit`参数中设置。
	Filter interface{} `json:"filter"`
	FromMessageID int64 `json:"from_message_id"` // 从该信息开始整理聊天记录的消息识别码。默认为从最后一条可见信息。 当 `direction = Forward` 时，`from_message_id` 参数为必填。
	Limit int64 `json:"limit"` // 答复的信息数量。默认设置为50。最大值是1000。
}

type ChatChatSendMessageRequest struct {
	ChatID string `json:"chat_id"` // 聊天识别码。
	Text string `json:"text"` // plain文本格式的信息文本1到1000个字符。
}

type ChatRead struct {
	ChatID string `json:"chat_id"` // 聊天识别码。
	FromMessageID int64 `json:"from_message_id"` // 信息识别码。
}

type ChatChatSendMessageResponse struct {
	Result string `json:"result"` // 请求的处理结果。
}

type ChatChatSendFileRequest struct {
	Base64Content string `json:"base64_content"` // 文件为 base64 行形式。
	ChatID string `json:"chat_id"` // 聊天识别码。
	Name string `json:"name"` // 带有扩展名的文件名。
}

type ChatChatSendFileResponse struct {
	Result string `json:"result"` // 请求的处理结果。
}

type V3ChatList struct {
	Cursor string `json:"cursor"` // 后续数据的选择标志。
	Filter interface{} `json:"filter"`
	Limit int64 `json:"limit"` // 回答中值的数量。默认值为30。最大值是100。
}

type ChatChatStartRequest struct {
	PostingNumber string `json:"posting_number"` // 发货识别码。
}

type V3ChatHistoryResponse struct {
	Messages []interface{} `json:"messages"` // 根据请求正文中的`direction`参数排序的信息数组。
	HasNext bool `json:"has_next"` // 表示不是所有信息都在答复中返回。
}

type V2ChatReadResponse struct {
	UnreadCount int64 `json:"unread_count"` // 聊天中未读消息的数量。
}

type V3ChatListResponse struct {
	Cursor string `json:"cursor"` // 后续数据的选择标志。
	HasNext bool `json:"has_next"` // 表示响应中未包含全部聊天: - `true` — 使用新的 `cursor` 参数重新发送请求以获取剩余聊天; - `false` — 响应中已包含匹配请求筛选器的所有聊天。
	Chats interface{} `json:"chats"` // 聊天数据。
	TotalUnreadCount int64 `json:"total_unread_count"` // 未读信息总数。
}

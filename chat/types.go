package chat

type ChatChatSendMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text string `json:"text"`
}

type V3ChatHistoryResponse struct {
	HasNext bool `json:"has_next"`
	Messages []interface{} `json:"messages"`
}

type ChatChatStartResponse struct {
	Result interface{} `json:"result"`
}

type V3ChatHistoryRequest struct {
	ChatID string `json:"chat_id"`
	Direction string `json:"direction"`
	Filter interface{} `json:"filter"`
	FromMessageID int64 `json:"from_message_id"`
	Limit int64 `json:"limit"`
}

type ChatChatStartRequest struct {
	PostingNumber string `json:"posting_number"`
}

type V3ChatList struct {
	Limit int64 `json:"limit"`
	Cursor string `json:"cursor"`
	Filter interface{} `json:"filter"`
}

type V3ChatListResponse struct {
	Chats interface{} `json:"chats"`
	TotalUnreadCount int64 `json:"total_unread_count"`
	Cursor string `json:"cursor"`
	HasNext bool `json:"has_next"`
}

type ChatRead struct {
	ChatID string `json:"chat_id"`
	FromMessageID int64 `json:"from_message_id"`
}

type ChatChatSendMessageResponse struct {
	Result string `json:"result"`
}

type V2ChatReadResponse struct {
	UnreadCount int64 `json:"unread_count"`
}

type ChatChatSendFileRequest struct {
	Base64Content string `json:"base64_content"`
	ChatID string `json:"chat_id"`
	Name string `json:"name"`
}

type ChatChatSendFileResponse struct {
	Result string `json:"result"`
}

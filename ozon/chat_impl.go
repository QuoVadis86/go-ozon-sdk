package ozon

import "context"

func (s *ChatService) ChatStart(ctx context.Context, req *ChatChatStartRequest) (*ChatChatStartResponse, error) {
	var resp ChatChatStartResponse
	_, err := s.t.Post(ctx, "/v1/chat/start", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ChatService) ChatSendFile(ctx context.Context, req *ChatChatSendFileRequest) (*ChatChatSendFileResponse, error) {
	var resp ChatChatSendFileResponse
	_, err := s.t.Post(ctx, "/v1/chat/send/file", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ChatService) ChatSendMessage(ctx context.Context, req *ChatChatSendMessageRequest) (*ChatChatSendMessageResponse, error) {
	var resp ChatChatSendMessageResponse
	_, err := s.t.Post(ctx, "/v1/chat/send/message", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ChatService) ChatReadV2(ctx context.Context, req *ChatRead) (*V2ChatReadResponse, error) {
	var resp V2ChatReadResponse
	_, err := s.t.Post(ctx, "/v2/chat/read", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ChatService) ChatHistoryV3(ctx context.Context, req *V3ChatHistoryRequest) (*V3ChatHistoryResponse, error) {
	var resp V3ChatHistoryResponse
	_, err := s.t.Post(ctx, "/v3/chat/history", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *ChatService) ChatListV3(ctx context.Context, req *V3ChatList) (*V3ChatListResponse, error) {
	var resp V3ChatListResponse
	_, err := s.t.Post(ctx, "/v3/chat/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package chat

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) ChatSendMessage(ctx context.Context, req *internal.ChatChatSendMessageRequest) (*internal.ChatChatSendMessageResponse, error) {
	var resp internal.ChatChatSendMessageResponse
	err := s.Client.Post(ctx, "/v1/chat/send/message", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatListV3(ctx context.Context, req *internal.V3ChatList) (*internal.V3ChatListResponse, error) {
	var resp internal.V3ChatListResponse
	err := s.Client.Post(ctx, "/v3/chat/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatReadV2(ctx context.Context, req *internal.ChatRead) (*internal.V2ChatReadResponse, error) {
	var resp internal.V2ChatReadResponse
	err := s.Client.Post(ctx, "/v2/chat/read", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatSendFile(ctx context.Context, req *internal.ChatChatSendFileRequest) (*internal.ChatChatSendFileResponse, error) {
	var resp internal.ChatChatSendFileResponse
	err := s.Client.Post(ctx, "/v1/chat/send/file", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatStart(ctx context.Context, req *internal.ChatChatStartRequest) (*internal.ChatChatStartResponse, error) {
	var resp internal.ChatChatStartResponse
	err := s.Client.Post(ctx, "/v1/chat/start", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatHistoryV3(ctx context.Context, req *internal.V3ChatHistoryRequest) (*internal.V3ChatHistoryResponse, error) {
	var resp internal.V3ChatHistoryResponse
	err := s.Client.Post(ctx, "/v3/chat/history", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package chat

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

func (s *Service) ChatStart(ctx context.Context, req *ChatChatStartRequest) (*ChatChatStartResponse, error) {
	var resp ChatChatStartResponse
	err := s.Client.Post(ctx, "/v1/chat/start", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatListV3(ctx context.Context, req *V3ChatList) (*V3ChatListResponse, error) {
	var resp V3ChatListResponse
	err := s.Client.Post(ctx, "/v3/chat/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatSendFile(ctx context.Context, req *ChatChatSendFileRequest) (*ChatChatSendFileResponse, error) {
	var resp ChatChatSendFileResponse
	err := s.Client.Post(ctx, "/v1/chat/send/file", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatSendMessage(ctx context.Context, req *ChatChatSendMessageRequest) (*ChatChatSendMessageResponse, error) {
	var resp ChatChatSendMessageResponse
	err := s.Client.Post(ctx, "/v1/chat/send/message", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatHistoryV3(ctx context.Context, req *V3ChatHistoryRequest) (*V3ChatHistoryResponse, error) {
	var resp V3ChatHistoryResponse
	err := s.Client.Post(ctx, "/v3/chat/history", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatReadV2(ctx context.Context, req *ChatRead) (*V2ChatReadResponse, error) {
	var resp V2ChatReadResponse
	err := s.Client.Post(ctx, "/v2/chat/read", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

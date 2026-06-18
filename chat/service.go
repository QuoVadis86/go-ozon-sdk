package chat

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) ChatSendFile(ctx context.Context, req *types.ChatChatSendFileRequest) (*types.ChatChatSendFileResponse, error) {
	var resp types.ChatChatSendFileResponse
	err := s.Client.Post(ctx, "/v1/chat/send/file", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatSendMessage(ctx context.Context, req *types.ChatChatSendMessageRequest) (*types.ChatChatSendMessageResponse, error) {
	var resp types.ChatChatSendMessageResponse
	err := s.Client.Post(ctx, "/v1/chat/send/message", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatStart(ctx context.Context, req *types.ChatChatStartRequest) (*types.ChatChatStartResponse, error) {
	var resp types.ChatChatStartResponse
	err := s.Client.Post(ctx, "/v1/chat/start", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatListV3(ctx context.Context, req *types.V3ChatList) (*types.V3ChatListResponse, error) {
	var resp types.V3ChatListResponse
	err := s.Client.Post(ctx, "/v3/chat/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatHistoryV3(ctx context.Context, req *types.V3ChatHistoryRequest) (*types.V3ChatHistoryResponse, error) {
	var resp types.V3ChatHistoryResponse
	err := s.Client.Post(ctx, "/v3/chat/history", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ChatReadV2(ctx context.Context, req *types.ChatRead) (*types.V2ChatReadResponse, error) {
	var resp types.V2ChatReadResponse
	err := s.Client.Post(ctx, "/v2/chat/read", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

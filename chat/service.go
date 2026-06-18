package chat

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 将信息标记为已读
func (s *Service) ChatReadV2(ctx context.Context, req *Read) (*V2ChatReadResponse, error) {
	var resp V2ChatReadResponse
	err := s.Client.Post(ctx, "/v2/chat/read", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 聊天历史记录
func (s *Service) ChatHistoryV3(ctx context.Context, req *V3ChatHistoryRequest) (*V3ChatHistoryResponse, error) {
	var resp V3ChatHistoryResponse
	err := s.Client.Post(ctx, "/v3/chat/history", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建新聊天
// Note: 为了确定地址或商品型号
func (s *Service) ChatStart(ctx context.Context, req *ChatStartRequest) (*ChatStartResponse, error) {
	var resp ChatStartResponse
	err := s.Client.Post(ctx, "/v1/chat/start", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 发送信息
func (s *Service) ChatSendMessage(ctx context.Context, req *ChatSendMessageRequest) (*ChatSendMessageResponse, error) {
	var resp ChatSendMessageResponse
	err := s.Client.Post(ctx, "/v1/chat/send/message", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 聊天清单
func (s *Service) ChatListV3(ctx context.Context, req *V3ChatList) (*V3ChatListResponse, error) {
	var resp V3ChatListResponse
	err := s.Client.Post(ctx, "/v3/chat/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 发送文件
func (s *Service) ChatSendFile(ctx context.Context, req *ChatSendFileRequest) (*ChatSendFileResponse, error) {
	var resp ChatSendFileResponse
	err := s.Client.Post(ctx, "/v1/chat/send/file", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

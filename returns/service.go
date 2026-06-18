package returns

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 拒绝退货申请
func (s *Service) ReturnsRfbsRejectV2(ctx context.Context, req *V2ReturnsRfbsRejectRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/reject", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 确认收到待检查商品
func (s *Service) ReturnsRfbsReceiveReturnV2(ctx context.Context, req *V2ReturnsRfbsReceiveReturnRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/receive-return", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 确认 rFBS 取消申请
func (s *Service) ConditionalCancellationApproveV2(ctx context.Context) error {
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/approve", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// 退货申请列表
func (s *Service) ReturnsRfbsListV2(ctx context.Context, req *V2ReturnsRfbsListRequest) (*V2ReturnsRfbsListResponse, error) {
	var resp V2ReturnsRfbsListResponse
	err := s.Client.Post(ctx, "/v2/returns/rfbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取 rFBS 取消申请列表
func (s *Service) GetConditionalCancellationListV2(ctx context.Context, req *V2GetConditionalCancellationListV2Request) (*V2GetConditionalCancellationListV2Response, error) {
	var resp V2GetConditionalCancellationListV2Response
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 向买家退款
// Note: 请使用此方法：
func (s *Service) ReturnsRfbsReturnMoneyV2(ctx context.Context, req *V2ReturnsRfbsReturnMoneyRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/return-money", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FBO和FBS退货信息
func (s *Service) ReturnsList(ctx context.Context, req *V1GetReturnsListRequest) (*V1GetReturnsListResponse, error) {
	var resp V1GetReturnsListResponse
	err := s.Client.Post(ctx, "/v1/returns/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 拒绝 rFBS 取消申请
func (s *Service) ConditionalCancellationRejectV2(ctx context.Context) error {
	err := s.Client.Post(ctx, "/v2/conditional-cancellation/reject", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// 传递 rFBS 退货的可用操作
func (s *Service) ReturnsRfbsActionSet(ctx context.Context, req *V1ReturnsRfbsActionSetRequest) error {
	err := s.Client.Post(ctx, "/v1/returns/rfbs/action/set", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 退货申请信息
func (s *Service) ReturnsRfbsGetV2(ctx context.Context, req *V2ReturnsRfbsGetRequest) (*V2ReturnsRfbsGetResponse, error) {
	var resp V2ReturnsRfbsGetResponse
	err := s.Client.Post(ctx, "/v2/returns/rfbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 批准退货申请
// Note: 请使用[/v2/returns/rfbs/receive-return](#operation/RFBSReturnsAPI_ReturnsRfbsReceiveReturnV2)方法确认收到商品
func (s *Service) ReturnsRfbsVerifyV2(ctx context.Context, req *V2ReturnsRfbsVerifyRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/verify", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 退还部分商品金额
func (s *Service) ReturnsRfbsCompensateV2(ctx context.Context, req *V2ReturnsRfbsCompensateRequest) (*V1Empty, error) {
	var resp V1Empty
	err := s.Client.Post(ctx, "/v2/returns/rfbs/compensate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

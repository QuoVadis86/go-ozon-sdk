package fbs

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// Deprecated: use /v4/posting/fbs/unfulfilled/list instead
// 未处理货件列表
func (s *Service) GetFbsPostingUnfulfilledList(ctx context.Context, req *Postingv3GetFbsPostingUnfulfilledListRequest) (*Postingv3GetFbsPostingUnfulfilledListResponse, error) {
	var resp Postingv3GetFbsPostingUnfulfilledListResponse
	err := s.Client.Post(ctx, "/v3/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 打印标签
func (s *Service) PostingFBSPackageLabel(ctx context.Context, req *PostingPostingFBSPackageLabelRequest) (*PostingPostingFBSPackageLabelResponse, error) {
	var resp PostingPostingFBSPackageLabelResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/package-label", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 发运组成商品更改
func (s *Service) SetPostings(ctx context.Context, req *V1SetPostingsRequest) (*V1SetPostingsResponse, error) {
	var resp V1SetPostingsResponse
	err := s.Client.Post(ctx, "/v1/carriage/set-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 货位标签
func (s *Service) PostingFBSActGetContainerLabels(ctx context.Context, req *PostingPostingFBSActGetContainerLabelsRequest) (*PostingPostingFBSActGetContainerLabelsResponse, error) {
	var resp PostingPostingFBSActGetContainerLabelsResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/act/get-container-labels", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建发运
// Note: 请使用方法[/v2/posting/fbs/act/get-postings](#operation/PostingAPI_ActPostingList)。
// Note: 如需获取发运中的货件列表，请使用方法[/v2/posting/fbs/act/get-postings](#operation/PostingAPI_ActPostingList)。
func (s *Service) CarriageCreate(ctx context.Context, req *V1CarriageCreateRequest) (*V1CarriageCreateResponse, error) {
	var resp V1CarriageCreateResponse
	err := s.Client.Post(ctx, "/v1/carriage/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 按条形码获取有关货件的信息
func (s *Service) GetFbsPostingByBarcode(ctx context.Context, req *PostingGetFbsPostingByBarcodeRequest) (*V2FbsPostingResponse, error) {
	var resp V2FbsPostingResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/get-by-barcode", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 发运确认
func (s *Service) CarriageApprove(ctx context.Context, req *V1CarriageApproveRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/approve", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 获取发运中的商品列表
func (s *Service) AssemblyCarriageProductList(ctx context.Context, req *V1AssemblyCarriageProductListRequest) (*V1AssemblyCarriageProductListResponse, error) {
	var resp V1AssemblyCarriageProductListResponse
	err := s.Client.Post(ctx, "/v1/assembly/carriage/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 取消货运
// Note: 无法取消可能送达的包裹。
func (s *Service) CancelFbsPosting(ctx context.Context, req *PostingCancelFbsPostingRequest) (*PostingBooleanResponse, error) {
	var resp PostingBooleanResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取货件中的商品列表
func (s *Service) AssemblyFbsProductList(ctx context.Context, req *V1AssemblyFbsProductListRequest) (*V1AssemblyFbsProductListResponse, error) {
	var resp V1AssemblyFbsProductListResponse
	err := s.Client.Post(ctx, "/v1/assembly/fbs/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 状态改为“最后一英里”
// Note: 请使用 /v3/posting/fbs/get 方法检查当前货件状态。
// Note: 在更改状态前，请使用 /v3/posting/fbs/get 方法检查当前货件状态。
func (s *Service) FbsPostingLastMile(ctx context.Context, req *PostingFbsPostingLastMileRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/last-mile", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取货件列表
func (s *Service) AssemblyFbsPostingList(ctx context.Context, req *V1AssemblyFbsPostingListRequest) (*V1AssemblyFbsPostingListResponse, error) {
	var resp V1AssemblyFbsPostingListResponse
	err := s.Client.Post(ctx, "/v1/assembly/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 可供运输的列表
func (s *Service) GetCarriageAvailableList(ctx context.Context, req *Postingv1GetCarriageAvailableListRequest) (*Postingv1GetCarriageAvailableListResponse, error) {
	var resp Postingv1GetCarriageAvailableListResponse
	err := s.Client.Post(ctx, "/v1/posting/carriage-available/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将状态改成“运输中”
// Note: 请使用 /v3/posting/fbs/get 方法检查当前货件状态。
// Note: 在更改状态前，请使用 /v3/posting/fbs/get 方法检查当前货件状态。
func (s *Service) FbsPostingDelivering(ctx context.Context, req *PostingFbsPostingDeliveringRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/delivering", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取发运中的货件列表
func (s *Service) AssemblyCarriagePostingList(ctx context.Context, req *V1AssemblyCarriagePostingListRequest) (*V1AssemblyCarriagePostingListResponse, error) {
	var resp V1AssemblyCarriagePostingListResponse
	err := s.Client.Post(ctx, "/v1/assembly/carriage/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将状态改成“已送达”
// Note: 请使用 /v3/posting/fbs/get 方法检查当前货件状态。
// Note: 在更改状态前，请使用 /v3/posting/fbs/get 方法检查当前货件状态。
func (s *Service) FbsPostingDelivered(ctx context.Context, req *PostingFbsPostingDeliveredRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/delivered", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 按照ID获取货件信息
func (s *Service) GetFbsPostingV3(ctx context.Context, req *Postingv3GetFbsPostingRequest) (*V3GetFbsPostingResponseV3, error) {
	var resp V3GetFbsPostingResponseV3
	err := s.Client.Post(ctx, "/v3/posting/fbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 取消某些商品发货
// Note: 请使用该方法。
// Note: 无法从货件中发送部分产品，请使用该方法。
// Note: 为了在使用FBS或rFBS模式时获取取消原因的标识符`cancel_reason_id`，请使用方法[/v2/posting/fbs/cancel-reason/list](#operation/PostingAPI_GetPostingF...
func (s *Service) CancelFbsPostingProduct(ctx context.Context, req *PostingPostingProductCancelRequest) (*PostingPostingProductCancelResponse, error) {
	var resp PostingPostingProductCancelResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 货运取消原因
func (s *Service) GetPostingFbsCancelReasonV1(ctx context.Context, req *PostingCancelReasonRequest) (*PostingCancelReasonResponse, error) {
	var resp PostingCancelReasonResponse
	err := s.Client.Post(ctx, "/v1/posting/fbs/cancel-reason", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: use /v4/posting/fbs/list instead
// 货件列表
// Note: 要获取有关剩余货件的信息，请提出新的含 `offset`值的请求。
func (s *Service) GetFbsPostingListV3(ctx context.Context, req *Postingv3GetFbsPostingListRequest) (*V3GetFbsPostingListResponseV3, error) {
	var resp V3GetFbsPostingListResponseV3
	err := s.Client.Post(ctx, "/v3/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 货件的部分装配 (第4方案)
// Note: 请使用/v3/posting/fbs/get方式来检查，订单是否已完成备货。
func (s *Service) ShipFbsPostingPackage(ctx context.Context, req *V4FbsPostingShipPackageV4Request) (*V4FbsPostingShipPackageV4Response, error) {
	var resp V4FbsPostingShipPackageV4Response
	err := s.Client.Post(ctx, "/v4/posting/fbs/ship/package", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 发运删除
func (s *Service) CarriageCancel(ctx context.Context, req *V1CarriageCancelRequest) (*V1CarriageCancelResponse, error) {
	var resp V1CarriageCancelResponse
	err := s.Client.Post(ctx, "/v1/carriage/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 添加商品产地信息
func (s *Service) SetCountryProductFbsPostingV2(ctx context.Context, req *V2FbsPostingProductCountrySetRequest) (*V2FbsPostingProductCountrySetResponse, error) {
	var resp V2FbsPostingProductCountrySetResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/country/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 货件取消原因
func (s *Service) GetPostingFbsCancelReasonList(ctx context.Context) (*PostingCancelReasonListResponse, error) {
	var resp PostingCancelReasonListResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/cancel-reason/list", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取未处理货件列表
// Note: 如需获取最新发运日期，请定期更新货件信息，或接入[推送通知](#tag/push_start)。
func (s *Service) PostingFbsUnfulfilledList(ctx context.Context, req *PostingV4PostingFbsUnfulfilledListRequest) (*PostingV4PostingFbsUnfulfilledListResponse, error) {
	var resp PostingV4PostingFbsUnfulfilledListResponse
	err := s.Client.Post(ctx, "/v4/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 确认货件发运日期
func (s *Service) SetPostingCutoff(ctx context.Context, req *V1SetPostingCutoffRequest) (*V1SetPostingCutoffResponse, error) {
	var resp V1SetPostingCutoffResponse
	err := s.Client.Post(ctx, "/v1/posting/cutoff/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 运输信息
func (s *Service) CarriageGet(ctx context.Context, req *CarriageCarriageGetRequest) (*CarriageCarriageGetResponse, error) {
	var resp CarriageCarriageGetResponse
	err := s.Client.Post(ctx, "/v1/carriage/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取货件列表
// Note: 如需获取最新发运日期，请定期更新货件信息，或接入[推送通知](#tag/push_start)。
func (s *Service) PostingFbsList(ctx context.Context, req *PostingV4PostingFbsListRequest) (*PostingV4PostingFbsListResponse, error) {
	var resp PostingV4PostingFbsListResponse
	err := s.Client.Post(ctx, "/v4/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 可用产地名单
func (s *Service) ListCountryProductFbsPostingV2(ctx context.Context, req *V2FbsPostingProductCountryListRequest) (*V2FbsPostingProductCountryListResponse, error) {
	var resp V2FbsPostingProductCountryListResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/country/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 添加跟踪号
func (s *Service) FbsPostingTrackingNumberSet(ctx context.Context, req *PostingFbsPostingTrackingNumberSetRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/tracking-number/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 单据中的货件列表
func (s *Service) ActPostingList(ctx context.Context, req *V2PostingFBSActGetPostingsRequest) (*V2PostingFBSActGetPostingsResponse, error) {
	var resp V2PostingFBSActGetPostingsResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/act/get-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 货件装运
func (s *Service) MoveFbsPostingToAwaitingDelivery(ctx context.Context, req *V2MovePostingToAwaitingDeliveryRequest) (*PostingBooleanResponse, error) {
	var resp PostingBooleanResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/awaiting-delivery", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 搜集订单 (第4方案)
// Note: 请使用/v3/posting/fbs/get方式来检查，订单是否已完成备货。
// Note: 不可以放在一个包装里。
// Note: 如需拆分订单，请在`packages`数组中传递多个对象。
func (s *Service) ShipFbsPostingV4(ctx context.Context, req *V4FbsPostingShipV4Request) (*V4FbsPostingShipV4Response, error) {
	var resp V4FbsPostingShipV4Response
	err := s.Client.Post(ctx, "/v4/posting/fbs/ship", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

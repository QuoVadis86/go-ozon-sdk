package fbo

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 获取已完成交货信息
func (s *Service) FbpArchiveGet(ctx context.Context, req *V1FbpArchiveGetRequest) (*V1FbpArchiveGetResponse, error) {
	var resp V1FbpArchiveGetResponse
	err := s.Client.Post(ctx, "/v1/fbp/archive/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取标签生成任务状态
func (s *Service) FbpGetLabel(ctx context.Context, req *V1FbpGetLabelRequest) (*V1FbpGetLabelResponse, error) {
	var resp V1FbpGetLabelResponse
	err := s.Client.Post(ctx, "/v1/fbp/label/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 取消上门揽收交货
func (s *Service) FbpOrderPickUpCancel(ctx context.Context, req *V1FbpOrderPickUpCancelRequest) (*V1FbpOrderPickUpCancelResponse, error) {
	var resp V1FbpOrderPickUpCancelResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/pick-up/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 取消交货
func (s *Service) FbpOrderDirectCancel(ctx context.Context, req *V1FbpOrderDirectCancelRequest) (*V1FbpOrderDirectCancelResponse, error) {
	var resp V1FbpOrderDirectCancelResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 编辑采用第三方承运商配送方法的交货草稿
func (s *Service) FbpDraftDirectTplDlvEdit(ctx context.Context, req *V1FbpDraftDirectTplDlvEditRequest) (*V1FbpDraftDirectTplDlvEditResponse, error) {
	var resp V1FbpDraftDirectTplDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/tpl-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 编辑接收点配送草稿的配送详情
func (s *Service) FbpDraftDropOffDlvEdit(ctx context.Context, req *V1FbpDraftDropOffDlvEditRequest) (*V1FbpDraftDropOffDlvEditResponse, error) {
	var resp V1FbpDraftDropOffDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新草稿中由卖家配送的信息
func (s *Service) FbpDraftDirectSellerDlvEdit(ctx context.Context, req *V1FbpDraftDirectSellerDlvEditRequest) (*V1FbpDraftDirectSellerDlvEditResponse, error) {
	var resp V1FbpDraftDirectSellerDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/seller-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 取消 pick-up 交货申请草稿
func (s *Service) FbpDraftPickUpDelete(ctx context.Context, req *V1FbpDraftPickUpDeleteRequest) (*V1FbpDraftPickUpDeleteResponse, error) {
	var resp V1FbpDraftPickUpDeleteResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取关于特定交货的信息
func (s *Service) FbpOrderGet(ctx context.Context, req *V1FbpOrderGetRequest) (*V1FbpOrderGetResponse, error) {
	var resp V1FbpOrderGetResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取省份内接收点列表
func (s *Service) FbpDraftDropOffPointList(ctx context.Context, req *V1FbpDraftDropOffPointListRequest) (*V1FbpDraftDropOffPointListResponse, error) {
	var resp V1FbpDraftDropOffPointListResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/point/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取验收证明书生成状态
func (s *Service) FbpCheckActState(ctx context.Context, req *V1FbpCheckActStateRequest) (*V1FbpCheckActStateResponse, error) {
	var resp V1FbpCheckActStateResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-from/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建接收点配送草稿
func (s *Service) FbpDraftDropOffCreate(ctx context.Context, req *V1FbpDraftDropOffCreateRequest) (*V1FbpDraftDropOffCreateResponse, error) {
	var resp V1FbpDraftDropOffCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建由卖家配送的草稿
func (s *Service) FbpDraftDirectSellerDlvCreate(ctx context.Context, req *V1FbpDraftDirectSellerDlvCreateRequest) (*V1FbpDraftDirectSellerDlvCreateResponse, error) {
	var resp V1FbpDraftDirectSellerDlvCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/seller-dlv/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 检查合作伙伴仓库可接收的商品列表
func (s *Service) FbpDraftDropOffProductValidate(ctx context.Context, req *V1FbpDraftDropOffProductValidateRequest) (*V1FbpDraftDropOffProductValidateResponse, error) {
	var resp V1FbpDraftDropOffProductValidateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将草稿单转为正式交货
func (s *Service) FbpDraftDirectRegistrate(ctx context.Context, req *V1FbpDraftDirectRegistrateRequest) (*V1FbpDraftDirectRegistrateResponse, error) {
	var resp V1FbpDraftDirectRegistrateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建标签生成任务
func (s *Service) FbpCreateLabel(ctx context.Context, req *V1FbpCreateLabelRequest) (*V1FbpCreateLabelResponse, error) {
	var resp V1FbpCreateLabelResponse
	err := s.Client.Post(ctx, "/v1/fbp/label/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取货件列表
func (s *Service) PostingFbpList(ctx context.Context, req *PostingV1PostingFbpListRequest) (*PostingV1PostingFbpListResponse, error) {
	var resp PostingV1PostingFbpListResponse
	err := s.Client.Post(ctx, "/v1/posting/fbp/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取接收点的营业时间表
func (s *Service) FbpOrderDropOffTimetable(ctx context.Context, req *V1FbpOrderDropOffTimetableRequest) (*V1FbpOrderDropOffTimetableResponse, error) {
	var resp V1FbpOrderDropOffTimetableResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/drop-off/timetable", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 验证用于 pick-up 交货的商品列表
func (s *Service) FbpDraftPickUpProductValidate(ctx context.Context, req *V1FbpDraftPickUpProductValidateRequest) (*V1FbpDraftPickUpProductValidateResponse, error) {
	var resp V1FbpDraftPickUpProductValidateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建第三方物流公司配送的申请草稿
func (s *Service) FbpDraftDirectTplDlvCreate(ctx context.Context, req *V1FbpDraftDirectTplDlvCreateRequest) (*V1FbpDraftDirectTplDlvCreateResponse, error) {
	var resp V1FbpDraftDirectTplDlvCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/tpl-dlv/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 编辑交货申请中的时间段
func (s *Service) FbpEditTimeslot(ctx context.Context, req *V1FbpEditTimeslotRequest) (*V1FbpEditTimeslotResponse, error) {
	var resp V1FbpEditTimeslotResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/timeslot/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取交货时间段列表
func (s *Service) FbpAvailableTimeslotList(ctx context.Context, req *V1FbpAvailableTimeslotListRequest) (*V1FbpAvailableTimeslotListResponse, error) {
	var resp V1FbpAvailableTimeslotListResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/timeslot/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建 pick-up 交货申请草稿
func (s *Service) FbpDraftPickupCreate(ctx context.Context, req *V1FbpDraftPickupCreateRequest) (*V1FbpDraftPickupCreateResponse, error) {
	var resp V1FbpDraftPickupCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取合作伙伴仓库列表
func (s *Service) FbpWarehouseList(ctx context.Context) (*V1FbpWarehouseListResponse, error) {
	var resp V1FbpWarehouseListResponse
	err := s.Client.Post(ctx, "/v1/fbp/warehouse/list", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取货物运单生成状态
func (s *Service) FbpCheckConsignmentNoteState(ctx context.Context, req *V1FbpCheckConsignmentNoteStateRequest) (*V1FbpCheckConsignmentNoteStateResponse, error) {
	var resp V1FbpCheckConsignmentNoteStateResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-to/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取交货草稿信息
func (s *Service) FbpDraftGet(ctx context.Context, req *V1FbpDraftGetRequest) (*V1FbpDraftGetResponse, error) {
	var resp V1FbpDraftGetResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 取消 drop-off 交货
func (s *Service) FbpOrderDropOffCancel(ctx context.Context, req *V1FbpOrderDropOffCancelRequest) (*V1FbpOrderDropOffCancelResponse, error) {
	var resp V1FbpOrderDropOffCancelResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/drop-off/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建不指定配送方法的交货申请草稿
func (s *Service) FbpDraftDirectCreate(ctx context.Context, req *V1FbpDraftDirectCreateRequest) (*V1FbpDraftDirectCreateResponse, error) {
	var resp V1FbpDraftDirectCreateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取直供的时间段列表
func (s *Service) FbpDraftDirectGetTimeslot(ctx context.Context, req *V1FbpDraftDirectGetTimeslotRequest) (*V1FbpDraftDirectGetTimeslotResponse, error) {
	var resp V1FbpDraftDirectGetTimeslotResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/timeslot/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 编辑草稿中的时间段
func (s *Service) FbpDraftDirectTimeslotEdit(ctx context.Context, req *V1FbpDraftDirectTimeslotEditRequest) (*V1FbpDraftDirectTimeslotEditResponse, error) {
	var resp V1FbpDraftDirectTimeslotEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/timeslot/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取接收点的营业时间表
func (s *Service) FbpDraftDropOffPointTimetable(ctx context.Context, req *V1FbpDraftDropOffPointTimetableRequest) (*V1FbpDraftDropOffPointTimetableResponse, error) {
	var resp V1FbpDraftDropOffPointTimetableResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/point/timetable", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 删除接收点配送草稿
func (s *Service) FbpDraftDropOffDelete(ctx context.Context, req *V1FbpDraftDropOffDeleteRequest) (*V1FbpDraftDropOffDeleteResponse, error) {
	var resp V1FbpDraftDropOffDeleteResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 生成验收证明书
func (s *Service) FbpCreateAct(ctx context.Context, req *V1FbpCreateActRequest) (*V1FbpCreateActResponse, error) {
	var resp V1FbpCreateActResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-from/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更改取货地点信息
func (s *Service) FbpOrderPickUpDlvEdit(ctx context.Context, req *V1FbpOrderPickUpDlvEditRequest) (*V1FbpOrderPickUpDlvEditResponse, error) {
	var resp V1FbpOrderPickUpDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/pick-up/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取省份列表
func (s *Service) FbpDraftDropOffProvinceList(ctx context.Context, req *V1FbpDraftDropOffProvinceListRequest) (*V1FbpDraftDropOffProvinceListResponse, error) {
	var resp V1FbpDraftDropOffProvinceListResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/province/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 编辑收货点的送货信息
func (s *Service) FbpOrderDropOffDlvEdit(ctx context.Context, req *V1FbpOrderDropOffDlvEditRequest) (*V1FbpOrderDropOffDlvEditResponse, error) {
	var resp V1FbpOrderDropOffDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/drop-off/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新卖家自配送信息
func (s *Service) FbpOrderDirectSellerDlvEdit(ctx context.Context, req *V1FbpOrderDirectSellerDlvEditRequest) (*V1FbpOrderDirectSellerDlvEditResponse, error) {
	var resp V1FbpOrderDirectSellerDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/direct/seller-dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 交货草稿列表
func (s *Service) FbpDraftList(ctx context.Context, req *V1FbpDraftListRequest) (*V1FbpDraftListResponse, error) {
	var resp V1FbpDraftListResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 检查合作伙伴仓库商品列表
func (s *Service) FbpDraftDirectProductValidate(ctx context.Context, req *V1FbpDraftDirectProductValidateRequest) (*V1FbpDraftDirectProductValidateResponse, error) {
	var resp V1FbpDraftDirectProductValidateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/product/validate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取已完成交货列表
func (s *Service) FbpArchiveList(ctx context.Context, req *V1FbpArchiveListRequest) (*V1FbpArchiveListResponse, error) {
	var resp V1FbpArchiveListResponse
	err := s.Client.Post(ctx, "/v1/fbp/archive/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取交货列表
func (s *Service) FbpOrderList(ctx context.Context, req *V1FbpOrderListRequest) (*V1FbpOrderListResponse, error) {
	var resp V1FbpOrderListResponse
	err := s.Client.Post(ctx, "/v1/fbp/order/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 修改 pick-up 交货申请
func (s *Service) FbpDraftPickupDlvEdit(ctx context.Context, req *V1FbpDraftPickupDlvEditRequest) (*V1FbpDraftPickupDlvEditResponse, error) {
	var resp V1FbpDraftPickupDlvEditResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/dlv/edit", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将草稿单转为正式交货
func (s *Service) FbpDraftPickUpRegistrate(ctx context.Context, req *V1FbpDraftPickUpRegistrateRequest) (*V1FbpDraftPickUpRegistrateResponse, error) {
	var resp V1FbpDraftPickUpRegistrateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/pick-up/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 生成货物运单
func (s *Service) FbpCreateConsignmentNote(ctx context.Context, req *V1FbpCreateConsignmentNoteRequest) (*V1FbpCreateConsignmentNoteResponse, error) {
	var resp V1FbpCreateConsignmentNoteResponse
	err := s.Client.Post(ctx, "/v1/fbp/act-to/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 交货或交货申请的商品组成
func (s *Service) SupplyOrderBundle(ctx context.Context, req *V1GetSupplyOrderBundleRequest) (*V1GetSupplyOrderBundleResponse, error) {
	var resp V1GetSupplyOrderBundleResponse
	err := s.Client.Post(ctx, "/v1/supply-order/bundle", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 删除交货申请草稿
func (s *Service) FbpDraftDirectDelete(ctx context.Context, req *V1FbpDraftDirectDeleteRequest) (*V1FbpDraftDirectDeleteResponse, error) {
	var resp V1FbpDraftDirectDeleteResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/direct/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将草稿转为正式交货
func (s *Service) FbpDraftDropOffRegistrate(ctx context.Context, req *V1FbpDraftDropOffRegistrateRequest) (*V1FbpDraftDropOffRegistrateResponse, error) {
	var resp V1FbpDraftDropOffRegistrateResponse
	err := s.Client.Post(ctx, "/v1/fbp/draft/drop-off/registrate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

package promo

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

// 获取促销活动可用商品列表
func (s *Service) SellerActionsProductsCandidates(ctx context.Context, req *V1SellerActionsProductsCandidatesRequest) (*V1SellerActionsProductsCandidatesResponse, error) {
	var resp V1SellerActionsProductsCandidatesResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/products/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 在促销活动自动添加列表中添加或更新商品
func (s *Service) ActionsAutoAddProductsUpdate(ctx context.Context, req *ActionsV1ActionsAutoAddProductsUpdateRequest) (*ActionsV1ActionsAutoAddProductsUpdateResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsUpdateResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 从促销活动自动添加列表中删除商品
func (s *Service) ActionsAutoAddProductsDelete(ctx context.Context, req *ActionsV1ActionsAutoAddProductsDeleteRequest) (*ActionsV1ActionsAutoAddProductsDeleteResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsDeleteResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取促销活动自动添加列表中的商品列表
func (s *Service) ActionsAutoAddProductsList(ctx context.Context, req *ActionsV1ActionsAutoAddProductsListRequest) (*ActionsV1ActionsAutoAddProductsListResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsListResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取CSV格式的促销码文件
func (s *Service) SellerActionsVoucherGet(ctx context.Context, req *V1SellerActionsVoucherGetRequest) (*V1SellerActionsVoucherGetResponse, error) {
	var resp V1SellerActionsVoucherGetResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/voucher/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 取消折扣申请
func (s *Service) TaskDecline(ctx context.Context, req *V1DeclineDiscountTasksRequest) (*V1ApproveDeclineDiscountTasksResponse, error) {
	var resp V1ApproveDeclineDiscountTasksResponse
	err := s.Client.Post(ctx, "/v1/actions/discounts-task/decline", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新“免息分期付款”机制的促销活动
func (s *Service) SellerActionsUpdateInstallment(ctx context.Context, req *V1SellerActionsUpdateInstallmentRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/installment", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 更新“基于订单总额的折扣”机制的促销活动
func (s *Service) SellerActionsUpdateDiscountWithCondition(ctx context.Context, req *V1SellerActionsUpdateDiscountWithConditionRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/discount-with-condition", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 更新“多级满额折扣”机制的促销活动
func (s *Service) SellerActionsUpdateMultiLevelDiscount(ctx context.Context, req *V1SellerActionsUpdateMultiLevelDiscountRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/multi-level-discount", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 创建采用"促销码折扣"机制的促销活动
func (s *Service) SellerActionsCreateVoucher(ctx context.Context, req *V1SellerActionsCreateVoucherRequest) (*V1SellerActionsCreateVoucherResponse, error) {
	var resp V1SellerActionsCreateVoucherResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/voucher", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 从促销活动中移除商品
func (s *Service) SellerActionsProductsDelete(ctx context.Context, req *V1SellerActionsProductsDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/products/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 将商品添加到促销活动中
func (s *Service) SellerActionsProductsAdd(ctx context.Context, req *V1SellerActionsProductsAddRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/products/add", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 启用或关闭活动
func (s *Service) SellerActionsChangeActivity(ctx context.Context, req *V1SellerActionsChangeActivityRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/change-activity", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 更新“促销码折扣”机制的促销活动
func (s *Service) SellerActionsUpdateVoucher(ctx context.Context, req *V1SellerActionsUpdateVoucherRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/voucher", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 申请折扣列表
func (s *Service) TaskList(ctx context.Context, req *V1GetDiscountTaskListRequest) (*V1GetDiscountTaskListResponse, error) {
	var resp V1GetDiscountTaskListResponse
	err := s.Client.Post(ctx, "/v1/actions/discounts-task/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新“折扣”机制的促销活动
func (s *Service) SellerActionsUpdateDiscount(ctx context.Context, req *V1SellerActionsUpdateDiscountRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/discount", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 创建采用"免息分期付款"机制的促销活动
func (s *Service) SellerActionsCreateInstallment(ctx context.Context, req *V1SellerActionsCreateInstallmentRequest) (*V1SellerActionsCreateInstallmentResponse, error) {
	var resp V1SellerActionsCreateInstallmentResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/installment", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取参与活动的商品列表
func (s *Service) SellerActionsProductsList(ctx context.Context, req *V1SellerActionsProductsListRequest) (*V1SellerActionsProductsListResponse, error) {
	var resp V1SellerActionsProductsListResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 同意折扣申请
func (s *Service) TaskApprove(ctx context.Context, req *V1ApproveDiscountTasksRequest) (*V1ApproveDeclineDiscountTasksResponse, error) {
	var resp V1ApproveDeclineDiscountTasksResponse
	err := s.Client.Post(ctx, "/v1/actions/discounts-task/approve", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建采用"折扣"机制的促销活动
func (s *Service) SellerActionsCreateDiscount(ctx context.Context, req *V1SellerActionsCreateDiscountRequest) (*V1SellerActionsCreateDiscountResponse, error) {
	var resp V1SellerActionsCreateDiscountResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 可用的促销商品清单
func (s *Service) PromosCandidates(ctx context.Context, req *SellerApiGetSellerProductV1Request) (*SellerApiGetSellerProductV1Response, error) {
	var resp SellerApiGetSellerProductV1Response
	err := s.Client.Post(ctx, "/v1/actions/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建采用"多级满额折扣"机制的促销活动
func (s *Service) SellerActionsCreateMultiLevelDiscount(ctx context.Context, req *V1SellerActionsCreateMultiLevelDiscountRequest) (*V1SellerActionsCreateMultiLevelDiscountResponse, error) {
	var resp V1SellerActionsCreateMultiLevelDiscountResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/multi-level-discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 从活动中删除商品
func (s *Service) PromosProductsDeactivate(ctx context.Context, req *SellerApiProductIDsV1Request) (*SellerApiProductV1ResponseDeactivate, error) {
	var resp SellerApiProductV1ResponseDeactivate
	err := s.Client.Post(ctx, "/v1/actions/products/deactivate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 在促销活动中增加一个商品
func (s *Service) PromosProductsActivate(ctx context.Context, req *SellerApiActivateProductV1Request) (*SellerApiProductV1Response, error) {
	var resp SellerApiProductV1Response
	err := s.Client.Post(ctx, "/v1/actions/products/activate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取促销活动列表
func (s *Service) SellerActionsList(ctx context.Context, req *V1SellerActionsListRequest) (*V1SellerActionsListResponse, error) {
	var resp V1SellerActionsListResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 创建采用"基于订单总额的折扣"机制的促销活动
func (s *Service) SellerActionsCreateDiscountWithCondition(ctx context.Context, req *V1SellerActionsCreateDiscountWithConditionRequest) (*V1SellerActionsCreateDiscountWithConditionResponse, error) {
	var resp V1SellerActionsCreateDiscountWithConditionResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/discount-with-condition", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将促销活动归档
func (s *Service) SellerActionsArchive(ctx context.Context, req *V1SellerActionsArchiveRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/archive", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 参与 活动的商品列表
func (s *Service) PromosProducts(ctx context.Context, req *SellerApiGetSellerProductV1Request) (*SellerApiGetSellerProductV1Response, error) {
	var resp SellerApiGetSellerProductV1Response
	err := s.Client.Post(ctx, "/v1/actions/products", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取可自动添加到促销活动中的商品列表
func (s *Service) ActionsAutoAddProductsCandidates(ctx context.Context, req *ActionsV1ActionsAutoAddProductsCandidatesRequest) (*ActionsV1ActionsAutoAddProductsCandidatesResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsCandidatesResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 活动清单
func (s *Service) Promos(ctx context.Context) (*SellerApiGetSellerActionsV1Response, error) {
	var resp SellerApiGetSellerActionsV1Response
	err := s.Client.Get(ctx, "/v1/actions", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

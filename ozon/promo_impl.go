package ozon

import "context"

func (s *PromoService) SellerActionsProductsAdd(ctx context.Context, req *V1SellerActionsProductsAddRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/products/add", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) SellerActionsCreateDiscountWithCondition(ctx context.Context, req *V1SellerActionsCreateDiscountWithConditionRequest) (*V1SellerActionsCreateDiscountWithConditionResponse, error) {
	var resp V1SellerActionsCreateDiscountWithConditionResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/create/discount-with-condition", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsCreateVoucher(ctx context.Context, req *V1SellerActionsCreateVoucherRequest) (*V1SellerActionsCreateVoucherResponse, error) {
	var resp V1SellerActionsCreateVoucherResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/create/voucher", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsProductsCandidates(ctx context.Context, req *V1SellerActionsProductsCandidatesRequest) (*V1SellerActionsProductsCandidatesResponse, error) {
	var resp V1SellerActionsProductsCandidatesResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/products/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsUpdateInstallment(ctx context.Context, req *V1SellerActionsUpdateInstallmentRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/update/installment", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) SellerActionsVoucherGet(ctx context.Context, req *V1SellerActionsVoucherGetRequest) (*V1SellerActionsVoucherGetResponse, error) {
	var resp V1SellerActionsVoucherGetResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/voucher/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) PromosProducts(ctx context.Context, req *SellerApiGetSellerProductV1Request) (*SellerApiGetSellerProductV1Response, error) {
	var resp SellerApiGetSellerProductV1Response
	_, err := s.t.Post(ctx, "/v1/actions/products", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsUpdateDiscount(ctx context.Context, req *V1SellerActionsUpdateDiscountRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/update/discount", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) TaskList(ctx context.Context, req *V1GetDiscountTaskListRequest) (*V1GetDiscountTaskListResponse, error) {
	var resp V1GetDiscountTaskListResponse
	_, err := s.t.Post(ctx, "/v1/actions/discounts-task/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) TaskDecline(ctx context.Context, req *V1DeclineDiscountTasksRequest) (*V1ApproveDeclineDiscountTasksResponse, error) {
	var resp V1ApproveDeclineDiscountTasksResponse
	_, err := s.t.Post(ctx, "/v1/actions/discounts-task/decline", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsUpdateVoucher(ctx context.Context, req *V1SellerActionsUpdateVoucherRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/update/voucher", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) ActionsAutoAddProductsDelete(ctx context.Context, req *ActionsV1ActionsAutoAddProductsDeleteRequest) (*ActionsV1ActionsAutoAddProductsDeleteResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsDeleteResponse
	_, err := s.t.Post(ctx, "/v1/actions/auto-add/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsChangeActivity(ctx context.Context, req *V1SellerActionsChangeActivityRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/change-activity", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) Promos(ctx context.Context) (*SellerApiGetSellerActionsV1Response, error) {
	var resp SellerApiGetSellerActionsV1Response
	_, err := s.t.Get(ctx, "/v1/actions", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsCreateInstallment(ctx context.Context, req *V1SellerActionsCreateInstallmentRequest) (*V1SellerActionsCreateInstallmentResponse, error) {
	var resp V1SellerActionsCreateInstallmentResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/create/installment", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) ActionsAutoAddProductsCandidates(ctx context.Context, req *ActionsV1ActionsAutoAddProductsCandidatesRequest) (*ActionsV1ActionsAutoAddProductsCandidatesResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsCandidatesResponse
	_, err := s.t.Post(ctx, "/v1/actions/auto-add/products/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsUpdateMultiLevelDiscount(ctx context.Context, req *V1SellerActionsUpdateMultiLevelDiscountRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/update/multi-level-discount", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) SellerActionsProductsList(ctx context.Context, req *V1SellerActionsProductsListRequest) (*V1SellerActionsProductsListResponse, error) {
	var resp V1SellerActionsProductsListResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) TaskApprove(ctx context.Context, req *V1ApproveDiscountTasksRequest) (*V1ApproveDeclineDiscountTasksResponse, error) {
	var resp V1ApproveDeclineDiscountTasksResponse
	_, err := s.t.Post(ctx, "/v1/actions/discounts-task/approve", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsList(ctx context.Context, req *V1SellerActionsListRequest) (*V1SellerActionsListResponse, error) {
	var resp V1SellerActionsListResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsProductsDelete(ctx context.Context, req *V1SellerActionsProductsDeleteRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/products/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) ActionsAutoAddProductsUpdate(ctx context.Context, req *ActionsV1ActionsAutoAddProductsUpdateRequest) (*ActionsV1ActionsAutoAddProductsUpdateResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsUpdateResponse
	_, err := s.t.Post(ctx, "/v1/actions/auto-add/products/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsUpdateDiscountWithCondition(ctx context.Context, req *V1SellerActionsUpdateDiscountWithConditionRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/update/discount-with-condition", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) SellerActionsCreateDiscount(ctx context.Context, req *V1SellerActionsCreateDiscountRequest) (*V1SellerActionsCreateDiscountResponse, error) {
	var resp V1SellerActionsCreateDiscountResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/create/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsArchive(ctx context.Context, req *V1SellerActionsArchiveRequest) error {
	_, err := s.t.Post(ctx, "/v1/seller-actions/archive", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PromoService) PromosCandidates(ctx context.Context, req *SellerApiGetSellerProductV1Request) (*SellerApiGetSellerProductV1Response, error) {
	var resp SellerApiGetSellerProductV1Response
	_, err := s.t.Post(ctx, "/v1/actions/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) PromosProductsActivate(ctx context.Context, req *SellerApiActivateProductV1Request) (*SellerApiProductV1Response, error) {
	var resp SellerApiProductV1Response
	_, err := s.t.Post(ctx, "/v1/actions/products/activate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) PromosProductsDeactivate(ctx context.Context, req *SellerApiProductIDsV1Request) (*SellerApiProductV1ResponseDeactivate, error) {
	var resp SellerApiProductV1ResponseDeactivate
	_, err := s.t.Post(ctx, "/v1/actions/products/deactivate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) SellerActionsCreateMultiLevelDiscount(ctx context.Context, req *V1SellerActionsCreateMultiLevelDiscountRequest) (*V1SellerActionsCreateMultiLevelDiscountResponse, error) {
	var resp V1SellerActionsCreateMultiLevelDiscountResponse
	_, err := s.t.Post(ctx, "/v1/seller-actions/create/multi-level-discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PromoService) ActionsAutoAddProductsList(ctx context.Context, req *ActionsV1ActionsAutoAddProductsListRequest) (*ActionsV1ActionsAutoAddProductsListResponse, error) {
	var resp ActionsV1ActionsAutoAddProductsListResponse
	_, err := s.t.Post(ctx, "/v1/actions/auto-add/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

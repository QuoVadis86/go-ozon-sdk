package promo

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) TaskApprove(ctx context.Context, req *internal.V1ApproveDiscountTasksRequest) (*internal.V1ApproveDeclineDiscountTasksResponse, error) {
	var resp internal.V1ApproveDeclineDiscountTasksResponse
	err := s.Client.Post(ctx, "/v1/actions/discounts-task/approve", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsCreateInstallment(ctx context.Context, req *internal.V1SellerActionsCreateInstallmentRequest) (*internal.V1SellerActionsCreateInstallmentResponse, error) {
	var resp internal.V1SellerActionsCreateInstallmentResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/installment", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsArchive(ctx context.Context, req *internal.V1SellerActionsArchiveRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/archive", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) PromosProducts(ctx context.Context, req *internal.SellerApiGetSellerProductV1Request) (*internal.SellerApiGetSellerProductV1Response, error) {
	var resp internal.SellerApiGetSellerProductV1Response
	err := s.Client.Post(ctx, "/v1/actions/products", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) TaskList(ctx context.Context, req *internal.V1GetDiscountTaskListRequest) (*internal.V1GetDiscountTaskListResponse, error) {
	var resp internal.V1GetDiscountTaskListResponse
	err := s.Client.Post(ctx, "/v1/actions/discounts-task/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsUpdateDiscountWithCondition(ctx context.Context, req *internal.V1SellerActionsUpdateDiscountWithConditionRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/discount-with-condition", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SellerActionsProductsDelete(ctx context.Context, req *internal.V1SellerActionsProductsDeleteRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/products/delete", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SellerActionsUpdateDiscount(ctx context.Context, req *internal.V1SellerActionsUpdateDiscountRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/discount", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ActionsAutoAddProductsDelete(ctx context.Context, req *internal.ActionsV1ActionsAutoAddProductsDeleteRequest) (*internal.ActionsV1ActionsAutoAddProductsDeleteResponse, error) {
	var resp internal.ActionsV1ActionsAutoAddProductsDeleteResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/delete", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsUpdateMultiLevelDiscount(ctx context.Context, req *internal.V1SellerActionsUpdateMultiLevelDiscountRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/multi-level-discount", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ActionsAutoAddProductsList(ctx context.Context, req *internal.ActionsV1ActionsAutoAddProductsListRequest) (*internal.ActionsV1ActionsAutoAddProductsListResponse, error) {
	var resp internal.ActionsV1ActionsAutoAddProductsListResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionsAutoAddProductsUpdate(ctx context.Context, req *internal.ActionsV1ActionsAutoAddProductsUpdateRequest) (*internal.ActionsV1ActionsAutoAddProductsUpdateResponse, error) {
	var resp internal.ActionsV1ActionsAutoAddProductsUpdateResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/update", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) Promos(ctx context.Context) (*internal.SellerApiGetSellerActionsV1Response, error) {
	var resp internal.SellerApiGetSellerActionsV1Response
	err := s.Client.Get(ctx, "/v1/actions", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PromosProductsDeactivate(ctx context.Context, req *internal.SellerApiProductIDsV1Request) (*internal.SellerApiProductV1ResponseDeactivate, error) {
	var resp internal.SellerApiProductV1ResponseDeactivate
	err := s.Client.Post(ctx, "/v1/actions/products/deactivate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsChangeActivity(ctx context.Context, req *internal.V1SellerActionsChangeActivityRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/change-activity", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SellerActionsUpdateVoucher(ctx context.Context, req *internal.V1SellerActionsUpdateVoucherRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/voucher", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SellerActionsCreateVoucher(ctx context.Context, req *internal.V1SellerActionsCreateVoucherRequest) (*internal.V1SellerActionsCreateVoucherResponse, error) {
	var resp internal.V1SellerActionsCreateVoucherResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/voucher", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PromosCandidates(ctx context.Context, req *internal.SellerApiGetSellerProductV1Request) (*internal.SellerApiGetSellerProductV1Response, error) {
	var resp internal.SellerApiGetSellerProductV1Response
	err := s.Client.Post(ctx, "/v1/actions/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsProductsCandidates(ctx context.Context, req *internal.V1SellerActionsProductsCandidatesRequest) (*internal.V1SellerActionsProductsCandidatesResponse, error) {
	var resp internal.V1SellerActionsProductsCandidatesResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/products/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsProductsAdd(ctx context.Context, req *internal.V1SellerActionsProductsAddRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/products/add", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SellerActionsCreateMultiLevelDiscount(ctx context.Context, req *internal.V1SellerActionsCreateMultiLevelDiscountRequest) (*internal.V1SellerActionsCreateMultiLevelDiscountResponse, error) {
	var resp internal.V1SellerActionsCreateMultiLevelDiscountResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/multi-level-discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PromosProductsActivate(ctx context.Context, req *internal.SellerApiActivateProductV1Request) (*internal.SellerApiProductV1Response, error) {
	var resp internal.SellerApiProductV1Response
	err := s.Client.Post(ctx, "/v1/actions/products/activate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsUpdateInstallment(ctx context.Context, req *internal.V1SellerActionsUpdateInstallmentRequest) error {
	err := s.Client.Post(ctx, "/v1/seller-actions/update/installment", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SellerActionsList(ctx context.Context, req *internal.V1SellerActionsListRequest) (*internal.V1SellerActionsListResponse, error) {
	var resp internal.V1SellerActionsListResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) TaskDecline(ctx context.Context, req *internal.V1DeclineDiscountTasksRequest) (*internal.V1ApproveDeclineDiscountTasksResponse, error) {
	var resp internal.V1ApproveDeclineDiscountTasksResponse
	err := s.Client.Post(ctx, "/v1/actions/discounts-task/decline", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsVoucherGet(ctx context.Context, req *internal.V1SellerActionsVoucherGetRequest) (*internal.V1SellerActionsVoucherGetResponse, error) {
	var resp internal.V1SellerActionsVoucherGetResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/voucher/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsCreateDiscountWithCondition(ctx context.Context, req *internal.V1SellerActionsCreateDiscountWithConditionRequest) (*internal.V1SellerActionsCreateDiscountWithConditionResponse, error) {
	var resp internal.V1SellerActionsCreateDiscountWithConditionResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/discount-with-condition", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsProductsList(ctx context.Context, req *internal.V1SellerActionsProductsListRequest) (*internal.V1SellerActionsProductsListResponse, error) {
	var resp internal.V1SellerActionsProductsListResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/products/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SellerActionsCreateDiscount(ctx context.Context, req *internal.V1SellerActionsCreateDiscountRequest) (*internal.V1SellerActionsCreateDiscountResponse, error) {
	var resp internal.V1SellerActionsCreateDiscountResponse
	err := s.Client.Post(ctx, "/v1/seller-actions/create/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionsAutoAddProductsCandidates(ctx context.Context, req *internal.ActionsV1ActionsAutoAddProductsCandidatesRequest) (*internal.ActionsV1ActionsAutoAddProductsCandidatesResponse, error) {
	var resp internal.ActionsV1ActionsAutoAddProductsCandidatesResponse
	err := s.Client.Post(ctx, "/v1/actions/auto-add/products/candidates", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

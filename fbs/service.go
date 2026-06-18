package fbs

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) AssemblyCarriagePostingList(ctx context.Context, req *types.V1AssemblyCarriagePostingListRequest) (*types.V1AssemblyCarriagePostingListResponse, error) {
	var resp types.V1AssemblyCarriagePostingListResponse
	err := s.Client.Post(ctx, "/v1/assembly/carriage/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFbsPostingUnfulfilledList(ctx context.Context, req *types.Postingv3GetFbsPostingUnfulfilledListRequest) (*types.Postingv3GetFbsPostingUnfulfilledListResponse, error) {
	var resp types.Postingv3GetFbsPostingUnfulfilledListResponse
	err := s.Client.Post(ctx, "/v3/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AssemblyFbsPostingList(ctx context.Context, req *types.V1AssemblyFbsPostingListRequest) (*types.V1AssemblyFbsPostingListResponse, error) {
	var resp types.V1AssemblyFbsPostingListResponse
	err := s.Client.Post(ctx, "/v1/assembly/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriageCreate(ctx context.Context, req *types.V1CarriageCreateRequest) (*types.V1CarriageCreateResponse, error) {
	var resp types.V1CarriageCreateResponse
	err := s.Client.Post(ctx, "/v1/carriage/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SetPostingCutoff(ctx context.Context, req *types.V1SetPostingCutoffRequest) (*types.V1SetPostingCutoffResponse, error) {
	var resp types.V1SetPostingCutoffResponse
	err := s.Client.Post(ctx, "/v1/posting/cutoff/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) MoveFbsPostingToAwaitingDelivery(ctx context.Context, req *types.V2MovePostingToAwaitingDeliveryRequest) (*types.PostingBooleanResponse, error) {
	var resp types.PostingBooleanResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/awaiting-delivery", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingLastMile(ctx context.Context, req *types.PostingFbsPostingLastMileRequest) (*types.PostingFbsPostingMoveStatusResponse, error) {
	var resp types.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/last-mile", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFbsUnfulfilledList(ctx context.Context, req *types.PostingV4PostingFbsUnfulfilledListRequest) (*types.PostingV4PostingFbsUnfulfilledListResponse, error) {
	var resp types.PostingV4PostingFbsUnfulfilledListResponse
	err := s.Client.Post(ctx, "/v4/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetCarriageAvailableList(ctx context.Context, req *types.Postingv1GetCarriageAvailableListRequest) (*types.Postingv1GetCarriageAvailableListResponse, error) {
	var resp types.Postingv1GetCarriageAvailableListResponse
	err := s.Client.Post(ctx, "/v1/posting/carriage-available/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ShipFbsPostingV4(ctx context.Context, req *types.Fbsv4FbsPostingShipV4Request) (*types.Fbsv4FbsPostingShipV4Response, error) {
	var resp types.Fbsv4FbsPostingShipV4Response
	err := s.Client.Post(ctx, "/v4/posting/fbs/ship", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingDelivered(ctx context.Context, req *types.PostingFbsPostingDeliveredRequest) (*types.PostingFbsPostingMoveStatusResponse, error) {
	var resp types.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/delivered", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFbsList(ctx context.Context, req *types.PostingV4PostingFbsListRequest) (*types.PostingV4PostingFbsListResponse, error) {
	var resp types.PostingV4PostingFbsListResponse
	err := s.Client.Post(ctx, "/v4/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFBSActGetContainerLabels(ctx context.Context, req *types.PostingPostingFBSActGetContainerLabelsRequest) error {
	err := s.Client.Post(ctx, "/v2/posting/fbs/act/get-container-labels", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CarriageCancel(ctx context.Context, req *types.V1CarriageCancelRequest) (*types.V1CarriageCancelResponse, error) {
	var resp types.V1CarriageCancelResponse
	err := s.Client.Post(ctx, "/v1/carriage/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriageGet(ctx context.Context, req *types.CarriageCarriageGetRequest) (*types.CarriageCarriageGetResponse, error) {
	var resp types.CarriageCarriageGetResponse
	err := s.Client.Post(ctx, "/v1/carriage/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetPostingFbsCancelReasonList(ctx context.Context) (*types.PostingCancelReasonListResponse, error) {
	var resp types.PostingCancelReasonListResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/cancel-reason/list", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriageApprove(ctx context.Context, req *types.V1CarriageApproveRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/approve", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AssemblyCarriageProductList(ctx context.Context, req *types.V1AssemblyCarriageProductListRequest) (*types.V1AssemblyCarriageProductListResponse, error) {
	var resp types.V1AssemblyCarriageProductListResponse
	err := s.Client.Post(ctx, "/v1/assembly/carriage/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CancelFbsPostingProduct(ctx context.Context, req *types.PostingPostingProductCancelRequest) (*types.PostingPostingProductCancelResponse, error) {
	var resp types.PostingPostingProductCancelResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingTrackingNumberSet(ctx context.Context, req *types.PostingFbsPostingTrackingNumberSetRequest) (*types.PostingFbsPostingMoveStatusResponse, error) {
	var resp types.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/tracking-number/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFbsPostingV3(ctx context.Context, req *types.Postingv3GetFbsPostingRequest) (*types.V3GetFbsPostingResponseV3, error) {
	var resp types.V3GetFbsPostingResponseV3
	err := s.Client.Post(ctx, "/v3/posting/fbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ShipFbsPostingPackage(ctx context.Context, req *types.V4FbsPostingShipPackageV4Request) (*types.V4FbsPostingShipPackageV4Response, error) {
	var resp types.V4FbsPostingShipPackageV4Response
	err := s.Client.Post(ctx, "/v4/posting/fbs/ship/package", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SetCountryProductFbsPostingV2(ctx context.Context, req *types.V2FbsPostingProductCountrySetRequest) (*types.V2FbsPostingProductCountrySetResponse, error) {
	var resp types.V2FbsPostingProductCountrySetResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/country/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActPostingList(ctx context.Context, req *types.V2PostingFBSActGetPostingsRequest) (*types.V2PostingFBSActGetPostingsResponse, error) {
	var resp types.V2PostingFBSActGetPostingsResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/act/get-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFbsPostingByBarcode(ctx context.Context, req *types.PostingGetFbsPostingByBarcodeRequest) (*types.V2FbsPostingResponse, error) {
	var resp types.V2FbsPostingResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/get-by-barcode", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetPostingFbsCancelReasonV1(ctx context.Context, req *types.PostingCancelReasonRequest) (*types.PostingCancelReasonResponse, error) {
	var resp types.PostingCancelReasonResponse
	err := s.Client.Post(ctx, "/v1/posting/fbs/cancel-reason", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListCountryProductFbsPostingV2(ctx context.Context, req *types.V2FbsPostingProductCountryListRequest) (*types.V2FbsPostingProductCountryListResponse, error) {
	var resp types.V2FbsPostingProductCountryListResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/country/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AssemblyFbsProductList(ctx context.Context, req *types.V1AssemblyFbsProductListRequest) (*types.V1AssemblyFbsProductListResponse, error) {
	var resp types.V1AssemblyFbsProductListResponse
	err := s.Client.Post(ctx, "/v1/assembly/fbs/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFBSPackageLabel(ctx context.Context, req *types.PostingPostingFBSPackageLabelRequest) error {
	err := s.Client.Post(ctx, "/v2/posting/fbs/package-label", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetPostings(ctx context.Context, req *types.V1SetPostingsRequest) (*types.V1SetPostingsResponse, error) {
	var resp types.V1SetPostingsResponse
	err := s.Client.Post(ctx, "/v1/carriage/set-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFbsPostingListV3(ctx context.Context, req *types.Postingv3GetFbsPostingListRequest) (*types.V3GetFbsPostingListResponseV3, error) {
	var resp types.V3GetFbsPostingListResponseV3
	err := s.Client.Post(ctx, "/v3/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CancelFbsPosting(ctx context.Context, req *types.PostingCancelFbsPostingRequest) (*types.PostingBooleanResponse, error) {
	var resp types.PostingBooleanResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingDelivering(ctx context.Context, req *types.PostingFbsPostingDeliveringRequest) (*types.PostingFbsPostingMoveStatusResponse, error) {
	var resp types.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/delivering", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

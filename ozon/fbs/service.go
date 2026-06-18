package fbs

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) AssemblyCarriagePostingList(ctx context.Context, req *internal.V1AssemblyCarriagePostingListRequest) (*internal.V1AssemblyCarriagePostingListResponse, error) {
	var resp internal.V1AssemblyCarriagePostingListResponse
	err := s.Client.Post(ctx, "/v1/assembly/carriage/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFbsPostingByBarcode(ctx context.Context, req *internal.PostingGetFbsPostingByBarcodeRequest) (*internal.V2FbsPostingResponse, error) {
	var resp internal.V2FbsPostingResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/get-by-barcode", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingTrackingNumberSet(ctx context.Context, req *internal.PostingFbsPostingTrackingNumberSetRequest) (*internal.PostingFbsPostingMoveStatusResponse, error) {
	var resp internal.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/tracking-number/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFbsList(ctx context.Context, req *internal.PostingV4PostingFbsListRequest) (*internal.PostingV4PostingFbsListResponse, error) {
	var resp internal.PostingV4PostingFbsListResponse
	err := s.Client.Post(ctx, "/v4/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFBSActGetContainerLabels(ctx context.Context, req *internal.PostingPostingFBSActGetContainerLabelsRequest) error {
	err := s.Client.Post(ctx, "/v2/posting/fbs/act/get-container-labels", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AssemblyFbsProductList(ctx context.Context, req *internal.V1AssemblyFbsProductListRequest) (*internal.V1AssemblyFbsProductListResponse, error) {
	var resp internal.V1AssemblyFbsProductListResponse
	err := s.Client.Post(ctx, "/v1/assembly/fbs/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CancelFbsPostingProduct(ctx context.Context, req *internal.PostingPostingProductCancelRequest) (*internal.PostingPostingProductCancelResponse, error) {
	var resp internal.PostingPostingProductCancelResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActPostingList(ctx context.Context, req *internal.V2PostingFBSActGetPostingsRequest) (*internal.V2PostingFBSActGetPostingsResponse, error) {
	var resp internal.V2PostingFBSActGetPostingsResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/act/get-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFbsPostingUnfulfilledList(ctx context.Context, req *internal.Postingv3GetFbsPostingUnfulfilledListRequest) (*internal.Postingv3GetFbsPostingUnfulfilledListResponse, error) {
	var resp internal.Postingv3GetFbsPostingUnfulfilledListResponse
	err := s.Client.Post(ctx, "/v3/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFbsUnfulfilledList(ctx context.Context, req *internal.PostingV4PostingFbsUnfulfilledListRequest) (*internal.PostingV4PostingFbsUnfulfilledListResponse, error) {
	var resp internal.PostingV4PostingFbsUnfulfilledListResponse
	err := s.Client.Post(ctx, "/v4/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AssemblyCarriageProductList(ctx context.Context, req *internal.V1AssemblyCarriageProductListRequest) (*internal.V1AssemblyCarriageProductListResponse, error) {
	var resp internal.V1AssemblyCarriageProductListResponse
	err := s.Client.Post(ctx, "/v1/assembly/carriage/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetCarriageAvailableList(ctx context.Context, req *internal.Postingv1GetCarriageAvailableListRequest) (*internal.Postingv1GetCarriageAvailableListResponse, error) {
	var resp internal.Postingv1GetCarriageAvailableListResponse
	err := s.Client.Post(ctx, "/v1/posting/carriage-available/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriageCreate(ctx context.Context, req *internal.V1CarriageCreateRequest) (*internal.V1CarriageCreateResponse, error) {
	var resp internal.V1CarriageCreateResponse
	err := s.Client.Post(ctx, "/v1/carriage/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ShipFbsPostingV4(ctx context.Context, req *internal.Fbsv4FbsPostingShipV4Request) (*internal.Fbsv4FbsPostingShipV4Response, error) {
	var resp internal.Fbsv4FbsPostingShipV4Response
	err := s.Client.Post(ctx, "/v4/posting/fbs/ship", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetPostingFbsCancelReasonList(ctx context.Context) (*internal.PostingCancelReasonListResponse, error) {
	var resp internal.PostingCancelReasonListResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/cancel-reason/list", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriageGet(ctx context.Context, req *internal.CarriageCarriageGetRequest) (*internal.CarriageCarriageGetResponse, error) {
	var resp internal.CarriageCarriageGetResponse
	err := s.Client.Post(ctx, "/v1/carriage/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingLastMile(ctx context.Context, req *internal.PostingFbsPostingLastMileRequest) (*internal.PostingFbsPostingMoveStatusResponse, error) {
	var resp internal.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/last-mile", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) MoveFbsPostingToAwaitingDelivery(ctx context.Context, req *internal.V2MovePostingToAwaitingDeliveryRequest) (*internal.PostingBooleanResponse, error) {
	var resp internal.PostingBooleanResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/awaiting-delivery", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetPostingFbsCancelReasonV1(ctx context.Context, req *internal.PostingCancelReasonRequest) (*internal.PostingCancelReasonResponse, error) {
	var resp internal.PostingCancelReasonResponse
	err := s.Client.Post(ctx, "/v1/posting/fbs/cancel-reason", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ShipFbsPostingPackage(ctx context.Context, req *internal.V4FbsPostingShipPackageV4Request) (*internal.V4FbsPostingShipPackageV4Response, error) {
	var resp internal.V4FbsPostingShipPackageV4Response
	err := s.Client.Post(ctx, "/v4/posting/fbs/ship/package", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriageApprove(ctx context.Context, req *internal.V1CarriageApproveRequest) error {
	err := s.Client.Post(ctx, "/v1/carriage/approve", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetFbsPostingV3(ctx context.Context, req *internal.Postingv3GetFbsPostingRequest) (*internal.V3GetFbsPostingResponseV3, error) {
	var resp internal.V3GetFbsPostingResponseV3
	err := s.Client.Post(ctx, "/v3/posting/fbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetFbsPostingListV3(ctx context.Context, req *internal.Postingv3GetFbsPostingListRequest) (*internal.V3GetFbsPostingListResponseV3, error) {
	var resp internal.V3GetFbsPostingListResponseV3
	err := s.Client.Post(ctx, "/v3/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ListCountryProductFbsPostingV2(ctx context.Context, req *internal.V2FbsPostingProductCountryListRequest) (*internal.V2FbsPostingProductCountryListResponse, error) {
	var resp internal.V2FbsPostingProductCountryListResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/country/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingDelivering(ctx context.Context, req *internal.PostingFbsPostingDeliveringRequest) (*internal.PostingFbsPostingMoveStatusResponse, error) {
	var resp internal.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/delivering", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CancelFbsPosting(ctx context.Context, req *internal.PostingCancelFbsPostingRequest) (*internal.PostingBooleanResponse, error) {
	var resp internal.PostingBooleanResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) AssemblyFbsPostingList(ctx context.Context, req *internal.V1AssemblyFbsPostingListRequest) (*internal.V1AssemblyFbsPostingListResponse, error) {
	var resp internal.V1AssemblyFbsPostingListResponse
	err := s.Client.Post(ctx, "/v1/assembly/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SetPostings(ctx context.Context, req *internal.V1SetPostingsRequest) (*internal.V1SetPostingsResponse, error) {
	var resp internal.V1SetPostingsResponse
	err := s.Client.Post(ctx, "/v1/carriage/set-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CarriageCancel(ctx context.Context, req *internal.V1CarriageCancelRequest) (*internal.V1CarriageCancelResponse, error) {
	var resp internal.V1CarriageCancelResponse
	err := s.Client.Post(ctx, "/v1/carriage/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) PostingFBSPackageLabel(ctx context.Context, req *internal.PostingPostingFBSPackageLabelRequest) error {
	err := s.Client.Post(ctx, "/v2/posting/fbs/package-label", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetPostingCutoff(ctx context.Context, req *internal.V1SetPostingCutoffRequest) (*internal.V1SetPostingCutoffResponse, error) {
	var resp internal.V1SetPostingCutoffResponse
	err := s.Client.Post(ctx, "/v1/posting/cutoff/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) SetCountryProductFbsPostingV2(ctx context.Context, req *internal.V2FbsPostingProductCountrySetRequest) (*internal.V2FbsPostingProductCountrySetResponse, error) {
	var resp internal.V2FbsPostingProductCountrySetResponse
	err := s.Client.Post(ctx, "/v2/posting/fbs/product/country/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) FbsPostingDelivered(ctx context.Context, req *internal.PostingFbsPostingDeliveredRequest) (*internal.PostingFbsPostingMoveStatusResponse, error) {
	var resp internal.PostingFbsPostingMoveStatusResponse
	err := s.Client.Post(ctx, "/v2/fbs/posting/delivered", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

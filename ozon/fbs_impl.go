package ozon

import "context"

func (s *FBSService) GetFbsPostingListV3(ctx context.Context, req *Postingv3GetFbsPostingListRequest) (*V3GetFbsPostingListResponseV3, error) {
	var resp V3GetFbsPostingListResponseV3
	_, err := s.t.Post(ctx, "/v3/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) GetCarriageAvailableList(ctx context.Context, req *Postingv1GetCarriageAvailableListRequest) (*Postingv1GetCarriageAvailableListResponse, error) {
	var resp Postingv1GetCarriageAvailableListResponse
	_, err := s.t.Post(ctx, "/v1/posting/carriage-available/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) CarriageCreate(ctx context.Context, req *V1CarriageCreateRequest) (*V1CarriageCreateResponse, error) {
	var resp V1CarriageCreateResponse
	_, err := s.t.Post(ctx, "/v1/carriage/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) GetFbsPostingUnfulfilledList(ctx context.Context, req *Postingv3GetFbsPostingUnfulfilledListRequest) (*Postingv3GetFbsPostingUnfulfilledListResponse, error) {
	var resp Postingv3GetFbsPostingUnfulfilledListResponse
	_, err := s.t.Post(ctx, "/v3/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) CancelFbsPostingProduct(ctx context.Context, req *PostingPostingProductCancelRequest) (*PostingPostingProductCancelResponse, error) {
	var resp PostingPostingProductCancelResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/product/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) FbsPostingDelivering(ctx context.Context, req *PostingFbsPostingDeliveringRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	_, err := s.t.Post(ctx, "/v2/fbs/posting/delivering", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) GetPostingFbsCancelReasonList(ctx context.Context) (*PostingCancelReasonListResponse, error) {
	var resp PostingCancelReasonListResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/cancel-reason/list", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) ShipFbsPostingV4(ctx context.Context, req *Fbsv4FbsPostingShipV4Request) (*Fbsv4FbsPostingShipV4Response, error) {
	var resp Fbsv4FbsPostingShipV4Response
	_, err := s.t.Post(ctx, "/v4/posting/fbs/ship", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) SetPostings(ctx context.Context, req *V1SetPostingsRequest) (*V1SetPostingsResponse, error) {
	var resp V1SetPostingsResponse
	_, err := s.t.Post(ctx, "/v1/carriage/set-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) FbsPostingTrackingNumberSet(ctx context.Context, req *PostingFbsPostingTrackingNumberSetRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	_, err := s.t.Post(ctx, "/v2/fbs/posting/tracking-number/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) CarriageGet(ctx context.Context, req *CarriageCarriageGetRequest) (*CarriageCarriageGetResponse, error) {
	var resp CarriageCarriageGetResponse
	_, err := s.t.Post(ctx, "/v1/carriage/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) PostingFbsList(ctx context.Context, req *PostingV4PostingFbsListRequest) (*PostingV4PostingFbsListResponse, error) {
	var resp PostingV4PostingFbsListResponse
	_, err := s.t.Post(ctx, "/v4/posting/fbs/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) ActPostingList(ctx context.Context, req *V2PostingFBSActGetPostingsRequest) (*V2PostingFBSActGetPostingsResponse, error) {
	var resp V2PostingFBSActGetPostingsResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/act/get-postings", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) AssemblyFbsPostingList(ctx context.Context, req *V1AssemblyFbsPostingListRequest) (*V1AssemblyFbsPostingListResponse, error) {
	var resp V1AssemblyFbsPostingListResponse
	_, err := s.t.Post(ctx, "/v1/assembly/fbs/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) SetPostingCutoff(ctx context.Context, req *V1SetPostingCutoffRequest) (*V1SetPostingCutoffResponse, error) {
	var resp V1SetPostingCutoffResponse
	_, err := s.t.Post(ctx, "/v1/posting/cutoff/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) GetFbsPostingV3(ctx context.Context, req *Postingv3GetFbsPostingRequest) (*V3GetFbsPostingResponseV3, error) {
	var resp V3GetFbsPostingResponseV3
	_, err := s.t.Post(ctx, "/v3/posting/fbs/get", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) SetCountryProductFbsPostingV2(ctx context.Context, req *V2FbsPostingProductCountrySetRequest) (*V2FbsPostingProductCountrySetResponse, error) {
	var resp V2FbsPostingProductCountrySetResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/product/country/set", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) GetFbsPostingByBarcode(ctx context.Context, req *PostingGetFbsPostingByBarcodeRequest) (*V2FbsPostingResponse, error) {
	var resp V2FbsPostingResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/get-by-barcode", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) ListCountryProductFbsPostingV2(ctx context.Context, req *V2FbsPostingProductCountryListRequest) (*V2FbsPostingProductCountryListResponse, error) {
	var resp V2FbsPostingProductCountryListResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/product/country/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) MoveFbsPostingToAwaitingDelivery(ctx context.Context, req *V2MovePostingToAwaitingDeliveryRequest) (*PostingBooleanResponse, error) {
	var resp PostingBooleanResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/awaiting-delivery", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) GetPostingFbsCancelReasonV1(ctx context.Context, req *PostingCancelReasonRequest) (*PostingCancelReasonResponse, error) {
	var resp PostingCancelReasonResponse
	_, err := s.t.Post(ctx, "/v1/posting/fbs/cancel-reason", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) PostingFBSPackageLabel(ctx context.Context, req *PostingPostingFBSPackageLabelRequest) error {
	_, err := s.t.Post(ctx, "/v2/posting/fbs/package-label", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *FBSService) AssemblyCarriageProductList(ctx context.Context, req *V1AssemblyCarriageProductListRequest) (*V1AssemblyCarriageProductListResponse, error) {
	var resp V1AssemblyCarriageProductListResponse
	_, err := s.t.Post(ctx, "/v1/assembly/carriage/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) PostingFbsUnfulfilledList(ctx context.Context, req *PostingV4PostingFbsUnfulfilledListRequest) (*PostingV4PostingFbsUnfulfilledListResponse, error) {
	var resp PostingV4PostingFbsUnfulfilledListResponse
	_, err := s.t.Post(ctx, "/v4/posting/fbs/unfulfilled/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) AssemblyCarriagePostingList(ctx context.Context, req *V1AssemblyCarriagePostingListRequest) (*V1AssemblyCarriagePostingListResponse, error) {
	var resp V1AssemblyCarriagePostingListResponse
	_, err := s.t.Post(ctx, "/v1/assembly/carriage/posting/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) CancelFbsPosting(ctx context.Context, req *PostingCancelFbsPostingRequest) (*PostingBooleanResponse, error) {
	var resp PostingBooleanResponse
	_, err := s.t.Post(ctx, "/v2/posting/fbs/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) ShipFbsPostingPackage(ctx context.Context, req *V4FbsPostingShipPackageV4Request) (*V4FbsPostingShipPackageV4Response, error) {
	var resp V4FbsPostingShipPackageV4Response
	_, err := s.t.Post(ctx, "/v4/posting/fbs/ship/package", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) CarriageCancel(ctx context.Context, req *V1CarriageCancelRequest) (*V1CarriageCancelResponse, error) {
	var resp V1CarriageCancelResponse
	_, err := s.t.Post(ctx, "/v1/carriage/cancel", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) CarriageApprove(ctx context.Context, req *V1CarriageApproveRequest) error {
	_, err := s.t.Post(ctx, "/v1/carriage/approve", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *FBSService) AssemblyFbsProductList(ctx context.Context, req *V1AssemblyFbsProductListRequest) (*V1AssemblyFbsProductListResponse, error) {
	var resp V1AssemblyFbsProductListResponse
	_, err := s.t.Post(ctx, "/v1/assembly/fbs/product/list", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) FbsPostingDelivered(ctx context.Context, req *PostingFbsPostingDeliveredRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	_, err := s.t.Post(ctx, "/v2/fbs/posting/delivered", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *FBSService) PostingFBSActGetContainerLabels(ctx context.Context, req *PostingPostingFBSActGetContainerLabelsRequest) error {
	_, err := s.t.Post(ctx, "/v2/posting/fbs/act/get-container-labels", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *FBSService) FbsPostingLastMile(ctx context.Context, req *PostingFbsPostingLastMileRequest) (*PostingFbsPostingMoveStatusResponse, error) {
	var resp PostingFbsPostingMoveStatusResponse
	_, err := s.t.Post(ctx, "/v2/fbs/posting/last-mile", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

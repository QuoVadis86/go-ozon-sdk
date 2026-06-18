package ozon

import "context"

func (s *BarcodeService) GenerateBarcode(ctx context.Context, req *V1GenerateBarcodeRequest) (*V1GenerateBarcodeResponse, error) {
	var resp V1GenerateBarcodeResponse
	_, err := s.t.Post(ctx, "/v1/barcode/generate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BarcodeService) AddBarcode(ctx context.Context, req *V1AddBarcodeRequest) (*V1AddBarcodeResponse, error) {
	var resp V1AddBarcodeResponse
	_, err := s.t.Post(ctx, "/v1/barcode/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

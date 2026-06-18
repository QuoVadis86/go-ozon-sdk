package barcode

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) AddBarcode(ctx context.Context, req *internal.V1AddBarcodeRequest) (*internal.V1AddBarcodeResponse, error) {
	var resp internal.V1AddBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GenerateBarcode(ctx context.Context, req *internal.V1GenerateBarcodeRequest) (*internal.V1GenerateBarcodeResponse, error) {
	var resp internal.V1GenerateBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/generate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

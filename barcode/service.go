package barcode

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

func (s *Service) AddBarcode(ctx context.Context, req *V1AddBarcodeRequest) (*V1AddBarcodeResponse, error) {
	var resp V1AddBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GenerateBarcode(ctx context.Context, req *V1GenerateBarcodeRequest) (*V1GenerateBarcodeResponse, error) {
	var resp V1GenerateBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/generate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

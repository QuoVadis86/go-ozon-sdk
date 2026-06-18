package barcode

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) AddBarcode(ctx context.Context, req *types.V1AddBarcodeRequest) (*types.V1AddBarcodeResponse, error) {
	var resp types.V1AddBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GenerateBarcode(ctx context.Context, req *types.V1GenerateBarcodeRequest) (*types.V1GenerateBarcodeResponse, error) {
	var resp types.V1GenerateBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/generate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

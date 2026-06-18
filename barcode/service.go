package barcode

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 创建商品条形码
// Note: 每次请求最多可为 100 个商品生成条形码
// Note: 每个卖家账号每分钟最多可使用该方法 20 次
func (s *Service) GenerateBarcode(ctx context.Context, req *V1GenerateBarcodeRequest) (*V1GenerateBarcodeResponse, error) {
	var resp V1GenerateBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/generate", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 为商品绑定条形码
// Note: 每个商品最多可绑定 100 个条形码
// Note: 每个卖家账号每分钟最多可使用该方法 20 次
func (s *Service) AddBarcode(ctx context.Context, req *V1AddBarcodeRequest) (*V1AddBarcodeResponse, error) {
	var resp V1AddBarcodeResponse
	err := s.Client.Post(ctx, "/v1/barcode/add", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

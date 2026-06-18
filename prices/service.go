package prices

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

// 最低价格时效性计时器更新
func (s *Service) ActionTimerUpdate(ctx context.Context, req *V1ProductActionTimerUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/product/action/timer/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

// 关于卖家库存余额的信息
func (s *Service) ProductStocksByWarehouseFbs(ctx context.Context, req *Productsv1GetProductInfoStocksByWarehouseFbsRequest) (*Productsv1GetProductInfoStocksByWarehouseFbsResponse, error) {
	var resp Productsv1GetProductInfoStocksByWarehouseFbsResponse
	err := s.Client.Post(ctx, "/v1/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取已设置计时器状态
func (s *Service) ActionTimerStatus(ctx context.Context, req *V1ProductActionTimerStatusRequest) (*V1ProductActionTimerStatusResponse, error) {
	var resp V1ProductActionTimerStatusResponse
	err := s.Client.Post(ctx, "/v1/product/action/timer/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新库存商品的数量
// Note: 请使用以下方法检查已预留库存数量：/v1/product/info/st
// Note: 每30秒内只能为一组商品-仓库更新一次库存，否则在回复中的result.
// Note: 每分钟可以发送80 个请求
func (s *Service) ProductsStocksV2(ctx context.Context, req *Productv2ProductsStocksRequest) (*Productv2ProductsStocksResponse, error) {
	var resp Productv2ProductsStocksResponse
	err := s.Client.Post(ctx, "/v2/products/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 更新价格
// Note: 每个商品的价格每小时不能更新超过10次
// Note: 不能更新超过10次
func (s *Service) ImportProductsPrices(ctx context.Context, req *ProductImportProductsPricesRequest) (*ProductImportProductsPricesResponse, error) {
	var resp ProductImportProductsPricesResponse
	err := s.Client.Post(ctx, "/v1/product/import/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取商品价格信息
func (s *Service) GetProductInfoPrices(ctx context.Context, req *Productv5GetProductInfoPricesV5Request) (*Productv5GetProductInfoPricesV5Response, error) {
	var resp Productv5GetProductInfoPricesV5Response
	err := s.Client.Post(ctx, "/v5/product/info/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 关于商品数量的信息
func (s *Service) GetProductInfoStocks(ctx context.Context, req *V4GetProductInfoStocksRequest) (*V4GetProductInfoStocksResponse, error) {
	var resp V4GetProductInfoStocksResponse
	err := s.Client.Post(ctx, "/v4/product/info/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 通过减价商品的SKU查找减价商品和主商品的信息
func (s *Service) GetProductInfoDiscounted(ctx context.Context, req *V1GetProductInfoDiscountedRequest) (*V1GetProductInfoDiscountedResponse, error) {
	var resp V1GetProductInfoDiscountedResponse
	err := s.Client.Post(ctx, "/v1/product/info/discounted", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 为打折商品设置折扣
func (s *Service) ProductUpdateDiscount(ctx context.Context, req *V1ProductUpdateDiscountRequest) (*V1ProductUpdateDiscountResponse, error) {
	var resp V1ProductUpdateDiscountResponse
	err := s.Client.Post(ctx, "/v1/product/update/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 获取卖家仓库库存信息
func (s *Service) GetProductInfoStocksByWarehouseFbsV2(ctx context.Context, req *V2GetProductInfoStocksByWarehouseFbsRequestV2) (*V2GetProductInfoStocksByWarehouseFbsResponseV2, error) {
	var resp V2GetProductInfoStocksByWarehouseFbsResponseV2
	err := s.Client.Post(ctx, "/v2/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

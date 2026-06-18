package prices

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport")

type Service struct { Client *transport.Client }

func (s *Service) GetProductInfoDiscounted(ctx context.Context, req *V1GetProductInfoDiscountedRequest) (*V1GetProductInfoDiscountedResponse, error) {
	var resp V1GetProductInfoDiscountedResponse
	err := s.Client.Post(ctx, "/v1/product/info/discounted", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsPrices(ctx context.Context, req *ProductImportProductsPricesRequest) (*ProductImportProductsPricesResponse, error) {
	var resp ProductImportProductsPricesResponse
	err := s.Client.Post(ctx, "/v1/product/import/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductStocksByWarehouseFbs(ctx context.Context, req *Productsv1GetProductInfoStocksByWarehouseFbsRequest) (*Productsv1GetProductInfoStocksByWarehouseFbsResponse, error) {
	var resp Productsv1GetProductInfoStocksByWarehouseFbsResponse
	err := s.Client.Post(ctx, "/v1/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoStocksByWarehouseFbsV2(ctx context.Context, req *V2GetProductInfoStocksByWarehouseFbsRequestV2) (*V2GetProductInfoStocksByWarehouseFbsResponseV2, error) {
	var resp V2GetProductInfoStocksByWarehouseFbsResponseV2
	err := s.Client.Post(ctx, "/v2/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionTimerStatus(ctx context.Context, req *V1ProductActionTimerStatusRequest) (*V1ProductActionTimerStatusResponse, error) {
	var resp V1ProductActionTimerStatusResponse
	err := s.Client.Post(ctx, "/v1/product/action/timer/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoPrices(ctx context.Context, req *Productv5GetProductInfoPricesV5Request) (*Productv5GetProductInfoPricesV5Response, error) {
	var resp Productv5GetProductInfoPricesV5Response
	err := s.Client.Post(ctx, "/v5/product/info/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductsStocksV2(ctx context.Context, req *Productv2ProductsStocksRequest) (*Productv2ProductsStocksResponse, error) {
	var resp Productv2ProductsStocksResponse
	err := s.Client.Post(ctx, "/v2/products/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoStocks(ctx context.Context, req *V4GetProductInfoStocksRequest) (*V4GetProductInfoStocksResponse, error) {
	var resp V4GetProductInfoStocksResponse
	err := s.Client.Post(ctx, "/v4/product/info/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateDiscount(ctx context.Context, req *V1ProductUpdateDiscountRequest) (*V1ProductUpdateDiscountResponse, error) {
	var resp V1ProductUpdateDiscountResponse
	err := s.Client.Post(ctx, "/v1/product/update/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionTimerUpdate(ctx context.Context, req *V1ProductActionTimerUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/product/action/timer/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

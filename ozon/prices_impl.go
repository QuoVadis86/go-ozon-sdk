package ozon

import "context"

func (s *PricesService) ProductUpdateDiscount(ctx context.Context, req *V1ProductUpdateDiscountRequest) (*V1ProductUpdateDiscountResponse, error) {
	var resp V1ProductUpdateDiscountResponse
	_, err := s.t.Post(ctx, "/v1/product/update/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) ActionTimerStatus(ctx context.Context, req *V1ProductActionTimerStatusRequest) (*V1ProductActionTimerStatusResponse, error) {
	var resp V1ProductActionTimerStatusResponse
	_, err := s.t.Post(ctx, "/v1/product/action/timer/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) GetProductInfoStocks(ctx context.Context, req *V4GetProductInfoStocksRequest) (*V4GetProductInfoStocksResponse, error) {
	var resp V4GetProductInfoStocksResponse
	_, err := s.t.Post(ctx, "/v4/product/info/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) ImportProductsPrices(ctx context.Context, req *ProductImportProductsPricesRequest) (*ProductImportProductsPricesResponse, error) {
	var resp ProductImportProductsPricesResponse
	_, err := s.t.Post(ctx, "/v1/product/import/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) ProductsStocksV2(ctx context.Context, req *Productv2ProductsStocksRequest) (*Productv2ProductsStocksResponse, error) {
	var resp Productv2ProductsStocksResponse
	_, err := s.t.Post(ctx, "/v2/products/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) ActionTimerUpdate(ctx context.Context, req *V1ProductActionTimerUpdateRequest) error {
	_, err := s.t.Post(ctx, "/v1/product/action/timer/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *PricesService) ProductStocksByWarehouseFbs(ctx context.Context, req *Productsv1GetProductInfoStocksByWarehouseFbsRequest) (*Productsv1GetProductInfoStocksByWarehouseFbsResponse, error) {
	var resp Productsv1GetProductInfoStocksByWarehouseFbsResponse
	_, err := s.t.Post(ctx, "/v1/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) GetProductInfoStocksByWarehouseFbsV2(ctx context.Context, req *V2GetProductInfoStocksByWarehouseFbsRequestV2) (*V2GetProductInfoStocksByWarehouseFbsResponseV2, error) {
	var resp V2GetProductInfoStocksByWarehouseFbsResponseV2
	_, err := s.t.Post(ctx, "/v2/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) GetProductInfoDiscounted(ctx context.Context, req *V1GetProductInfoDiscountedRequest) (*V1GetProductInfoDiscountedResponse, error) {
	var resp V1GetProductInfoDiscountedResponse
	_, err := s.t.Post(ctx, "/v1/product/info/discounted", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *PricesService) GetProductInfoPrices(ctx context.Context, req *Productv5GetProductInfoPricesV5Request) (*Productv5GetProductInfoPricesV5Response, error) {
	var resp Productv5GetProductInfoPricesV5Response
	_, err := s.t.Post(ctx, "/v5/product/info/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

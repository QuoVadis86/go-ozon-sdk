package prices

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) GetProductInfoStocksByWarehouseFbsV2(ctx context.Context, req *internal.V2GetProductInfoStocksByWarehouseFbsRequestV2) (*internal.V2GetProductInfoStocksByWarehouseFbsResponseV2, error) {
	var resp internal.V2GetProductInfoStocksByWarehouseFbsResponseV2
	err := s.Client.Post(ctx, "/v2/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateDiscount(ctx context.Context, req *internal.V1ProductUpdateDiscountRequest) (*internal.V1ProductUpdateDiscountResponse, error) {
	var resp internal.V1ProductUpdateDiscountResponse
	err := s.Client.Post(ctx, "/v1/product/update/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductStocksByWarehouseFbs(ctx context.Context, req *internal.Productsv1GetProductInfoStocksByWarehouseFbsRequest) (*internal.Productsv1GetProductInfoStocksByWarehouseFbsResponse, error) {
	var resp internal.Productsv1GetProductInfoStocksByWarehouseFbsResponse
	err := s.Client.Post(ctx, "/v1/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionTimerUpdate(ctx context.Context, req *internal.V1ProductActionTimerUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/product/action/timer/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetProductInfoDiscounted(ctx context.Context, req *internal.V1GetProductInfoDiscountedRequest) (*internal.V1GetProductInfoDiscountedResponse, error) {
	var resp internal.V1GetProductInfoDiscountedResponse
	err := s.Client.Post(ctx, "/v1/product/info/discounted", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductsStocksV2(ctx context.Context, req *internal.Productv2ProductsStocksRequest) (*internal.Productv2ProductsStocksResponse, error) {
	var resp internal.Productv2ProductsStocksResponse
	err := s.Client.Post(ctx, "/v2/products/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionTimerStatus(ctx context.Context, req *internal.V1ProductActionTimerStatusRequest) (*internal.V1ProductActionTimerStatusResponse, error) {
	var resp internal.V1ProductActionTimerStatusResponse
	err := s.Client.Post(ctx, "/v1/product/action/timer/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsPrices(ctx context.Context, req *internal.ProductImportProductsPricesRequest) (*internal.ProductImportProductsPricesResponse, error) {
	var resp internal.ProductImportProductsPricesResponse
	err := s.Client.Post(ctx, "/v1/product/import/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoPrices(ctx context.Context, req *internal.Productv5GetProductInfoPricesV5Request) (*internal.Productv5GetProductInfoPricesV5Response, error) {
	var resp internal.Productv5GetProductInfoPricesV5Response
	err := s.Client.Post(ctx, "/v5/product/info/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoStocks(ctx context.Context, req *internal.V4GetProductInfoStocksRequest) (*internal.V4GetProductInfoStocksResponse, error) {
	var resp internal.V4GetProductInfoStocksResponse
	err := s.Client.Post(ctx, "/v4/product/info/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

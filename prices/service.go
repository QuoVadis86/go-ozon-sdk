package prices

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *internal.Client }

func (s *Service) ProductsStocksV2(ctx context.Context, req *types.Productv2ProductsStocksRequest) (*types.Productv2ProductsStocksResponse, error) {
	var resp types.Productv2ProductsStocksResponse
	err := s.Client.Post(ctx, "/v2/products/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductUpdateDiscount(ctx context.Context, req *types.V1ProductUpdateDiscountRequest) (*types.V1ProductUpdateDiscountResponse, error) {
	var resp types.V1ProductUpdateDiscountResponse
	err := s.Client.Post(ctx, "/v1/product/update/discount", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ImportProductsPrices(ctx context.Context, req *types.ProductImportProductsPricesRequest) (*types.ProductImportProductsPricesResponse, error) {
	var resp types.ProductImportProductsPricesResponse
	err := s.Client.Post(ctx, "/v1/product/import/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ProductStocksByWarehouseFbs(ctx context.Context, req *types.Productsv1GetProductInfoStocksByWarehouseFbsRequest) (*types.Productsv1GetProductInfoStocksByWarehouseFbsResponse, error) {
	var resp types.Productsv1GetProductInfoStocksByWarehouseFbsResponse
	err := s.Client.Post(ctx, "/v1/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoStocksByWarehouseFbsV2(ctx context.Context, req *types.V2GetProductInfoStocksByWarehouseFbsRequestV2) (*types.V2GetProductInfoStocksByWarehouseFbsResponseV2, error) {
	var resp types.V2GetProductInfoStocksByWarehouseFbsResponseV2
	err := s.Client.Post(ctx, "/v2/product/info/stocks-by-warehouse/fbs", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoDiscounted(ctx context.Context, req *types.V1GetProductInfoDiscountedRequest) (*types.V1GetProductInfoDiscountedResponse, error) {
	var resp types.V1GetProductInfoDiscountedResponse
	err := s.Client.Post(ctx, "/v1/product/info/discounted", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) GetProductInfoStocks(ctx context.Context, req *types.V4GetProductInfoStocksRequest) (*types.V4GetProductInfoStocksResponse, error) {
	var resp types.V4GetProductInfoStocksResponse
	err := s.Client.Post(ctx, "/v4/product/info/stocks", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionTimerUpdate(ctx context.Context, req *types.V1ProductActionTimerUpdateRequest) error {
	err := s.Client.Post(ctx, "/v1/product/action/timer/update", req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetProductInfoPrices(ctx context.Context, req *types.Productv5GetProductInfoPricesV5Request) (*types.Productv5GetProductInfoPricesV5Response, error) {
	var resp types.Productv5GetProductInfoPricesV5Response
	err := s.Client.Post(ctx, "/v5/product/info/prices", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) ActionTimerStatus(ctx context.Context, req *types.V1ProductActionTimerStatusRequest) (*types.V1ProductActionTimerStatusResponse, error) {
	var resp types.V1ProductActionTimerStatusResponse
	err := s.Client.Post(ctx, "/v1/product/action/timer/status", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

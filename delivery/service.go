package delivery

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/transport"; "github.com/QuoVadis86/go-ozon-sdk/types")

type Service struct { Client *transport.Client }

func (s *Service) BindPolygon(ctx context.Context, req *types.Polygonv1PolygonBindRequest) (*types.Polygonv1Empty, error) {
	var resp types.Polygonv1Empty
	err := s.Client.Post(ctx, "/v1/polygon/bind", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) CreatePolygon(ctx context.Context, req *types.Polygonv1PolygonCreateRequest) (*types.Polygonv1PolygonCreateResponse, error) {
	var resp types.Polygonv1PolygonCreateResponse
	err := s.Client.Post(ctx, "/v1/polygon/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

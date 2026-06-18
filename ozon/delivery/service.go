package delivery

import ("context"; "github.com/QuoVadis86/go-ozon-sdk/internal")

type Service struct { Client *internal.Client }

func (s *Service) CreatePolygon(ctx context.Context, req *internal.Polygonv1PolygonCreateRequest) (*internal.Polygonv1PolygonCreateResponse, error) {
	var resp internal.Polygonv1PolygonCreateResponse
	err := s.Client.Post(ctx, "/v1/polygon/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) BindPolygon(ctx context.Context, req *internal.Polygonv1PolygonBindRequest) (*internal.Polygonv1Empty, error) {
	var resp internal.Polygonv1Empty
	err := s.Client.Post(ctx, "/v1/polygon/bind", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

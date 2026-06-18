package ozon

import "context"

func (s *DeliveryService) CreatePolygon(ctx context.Context, req *Polygonv1PolygonCreateRequest) (*Polygonv1PolygonCreateResponse, error) {
	var resp Polygonv1PolygonCreateResponse
	_, err := s.t.Post(ctx, "/v1/polygon/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *DeliveryService) BindPolygon(ctx context.Context, req *Polygonv1PolygonBindRequest) (*Polygonv1Empty, error) {
	var resp Polygonv1Empty
	_, err := s.t.Post(ctx, "/v1/polygon/bind", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

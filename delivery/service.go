package delivery

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 创建一个快递的设施
func (s *Service) CreatePolygon(ctx context.Context, req *Polygonv1PolygonCreateRequest) (*Polygonv1PolygonCreateResponse, error) {
	var resp Polygonv1PolygonCreateResponse
	err := s.Client.Post(ctx, "/v1/polygon/create", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 将快递方式与快递设施联系起来
func (s *Service) BindPolygon(ctx context.Context, req *Polygonv1PolygonBindRequest) (*Polygonv1Empty, error) {
	var resp Polygonv1Empty
	err := s.Client.Post(ctx, "/v1/polygon/bind", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

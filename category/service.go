package category

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

type Service struct{ Client *transport.Client }

// 类别特征列表
func (s *Service) GetAttributes(ctx context.Context, req *V1GetAttributesRequest) (*V1GetAttributesResponse, error) {
	var resp V1GetAttributesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 特征值指南
func (s *Service) GetAttributeValues(ctx context.Context, req *V1GetAttributeValuesRequest) (*V1GetAttributeValuesResponse, error) {
	var resp V1GetAttributeValuesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute/values", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 商品类别和类型的树形图
func (s *Service) GetTree(ctx context.Context, req *V1GetTreeRequest) (*V1GetTreeResponse, error) {
	var resp V1GetTreeResponse
	err := s.Client.Post(ctx, "/v1/description-category/tree", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// 根据属性的参考值进行搜索
func (s *Service) SearchAttributeValues(ctx context.Context, req *V1SearchAttributeValuesRequest) (*V1SearchAttributeValuesResponse, error) {
	var resp V1SearchAttributeValuesResponse
	err := s.Client.Post(ctx, "/v1/description-category/attribute/values/search", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

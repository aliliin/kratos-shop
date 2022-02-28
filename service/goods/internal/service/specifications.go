package service

import (
	"golang.org/x/net/context"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

// CreateGoodsSpecification 创建商品规格版本
func (g *GoodsService) CreateGoodsSpecification(ctx context.Context, r *v1.SpecificationRequest) (*v1.SpecificationResponse, error) {
	id, err := g.s.CreateSpecification(ctx, &biz.Specification{
		TypeID:   r.TypeId,
		Name:     r.Name,
		Sort:     r.Sort,
		Status:   r.Status,
		IsSKU:    r.IsSku,
		IsSelect: r.IsSelect,
	})

	if err != nil {
		return nil, err
	}
	return &v1.SpecificationResponse{
		Id: id,
	}, nil
}

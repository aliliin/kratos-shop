package service

import (
	"context"

	v1 "goods/api/goods/v1"
	"goods/internal/domain"
)

// CreateGoodsSpecification 创建商品规格版本
func (g *GoodsService) CreateGoodsSpecification(ctx context.Context, r *v1.SpecificationRequest) (*v1.SpecificationResponse, error) {
	var value []*domain.SpecificationValue
	// 组织规格参数值
	if r.SpecificationValue != nil {
		for _, v := range r.SpecificationValue {
			res := &domain.SpecificationValue{
				Value: v.Value,
				Sort:  v.Sort,
			}
			value = append(value, res)
		}
	}

	id, err := g.s.CreateSpecification(ctx, &domain.Specification{
		TypeID:             r.TypeId,
		Name:               r.Name,
		Sort:               r.Sort,
		Status:             r.Status,
		IsSKU:              r.IsSku,
		IsSelect:           r.IsSelect,
		SpecificationValue: value,
	})

	if err != nil {
		return nil, err
	}
	return &v1.SpecificationResponse{
		Id: id,
	}, nil
}

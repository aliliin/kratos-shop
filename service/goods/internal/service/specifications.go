package service

import (
	"fmt"
	"golang.org/x/net/context"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

// CreateGoodsSpecification 创建商品规格版本
func (g *GoodsService) CreateGoodsSpecification(ctx context.Context, r *v1.SpecificationRequest) (*v1.SpecificationResponse, error) {
	fmt.Println(r.SpecificationValue)
	var value []*biz.SpecificationValue
	if r.SpecificationValue != nil {
		for _, v := range r.SpecificationValue {
			res := &biz.SpecificationValue{
				ID:     int64(v.Id),
				AttrId: int64(v.AttrId),
				Value:  v.Value,
				Sort:   v.Sort,
			}
			value = append(value, res)
		}
	}
	id, err := g.s.CreateSpecification(ctx, &biz.Specification{
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
		Id: int32(id),
	}, nil
}

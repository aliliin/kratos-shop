package service

import (
	"context"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateAttrGroup 创建属性组
func (g *GoodsService) CreateAttrGroup(ctx context.Context, r *v1.AttrGroupRequest) (*v1.AttrGroupResponse, error) {
	result, err := g.ga.CreateAttrGroup(ctx, &biz.AttrGroup{
		TypeID: r.TypeId,
		Title:  r.Title,
		Desc:   r.Desc,
		Status: r.Status,
		Sort:   r.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &v1.AttrGroupResponse{
		Id:     result.ID,
		TypeId: result.TypeID,
		Title:  result.Title,
		Desc:   result.Desc,
		Status: result.Status,
		Sort:   result.Sort,
	}, nil
}

// CreateAttrValue 创建属性名称和值
func (g *GoodsService) CreateAttrValue(ctx context.Context, r *v1.AttrValueRequest) (*emptypb.Empty, error) {

	var value []*biz.GoodsAttrValue

	for _, v := range r.AttrValue {
		res := &biz.GoodsAttrValue{
			GroupID: v.GroupId,
			Value:   v.Value,
		}
		value = append(value, res)
	}

	_, err := g.ga.CreateAttrValue(ctx, &biz.GoodsAttr{
		TypeID:         r.TypeId,
		GroupID:        r.GroupId,
		Title:          r.Title,
		Sort:           r.Sort,
		Status:         r.Status,
		Desc:           r.Desc,
		GoodsAttrValue: value,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

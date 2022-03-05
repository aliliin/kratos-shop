package service

import (
	"context"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
	"goods/internal/domain"
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
func (g *GoodsService) CreateAttrValue(ctx context.Context, r *v1.AttrValueRequest) (*v1.AttrResponse, error) {
	var value []*biz.GoodsAttrValue
	for _, v := range r.AttrValue {
		res := &biz.GoodsAttrValue{
			GroupID: v.GroupId,
			Value:   v.Value,
		}
		value = append(value, res)
	}

	info, err := g.ga.CreateAttrValue(ctx, &domain.GoodsAttr{
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
	var AttrValue []*v1.AttrValueResponse
	for _, v := range info.GoodsAttrValue {
		result := &v1.AttrValueResponse{
			Id:      v.ID,
			AttrId:  v.AttrId,
			GroupId: v.GroupID,
			Value:   v.Value,
		}

		AttrValue = append(AttrValue, result)
	}

	return &v1.AttrResponse{
		Id:        info.ID,
		TypeId:    info.TypeID,
		GroupId:   info.GroupID,
		Title:     info.Title,
		Desc:      info.Desc,
		Status:    info.Status,
		Sort:      info.Sort,
		AttrValue: AttrValue,
	}, nil
}

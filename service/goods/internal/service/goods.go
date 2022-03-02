package service

import (
	"golang.org/x/net/context"
	v1 "goods/api/goods/v1"
)

// CreateGoods 创建商品
func (g *GoodsService) CreateGoods(ctx context.Context, r *v1.CreateGoodsRequest) (*v1.CreateGoodsRequest, error) {
	//result, err := g.g.CreateAttrGroup(ctx, &biz.AttrGroup{
	//    TypeID: r.TypeId,
	//    Title:  r.Title,
	//    Desc:   r.Desc,
	//    Status: r.Status,
	//    Sort:   r.Sort,
	//})
	//if err != nil {
	//    return nil, err
	//}
	//
	//return &v1.AttrGroupResponse{
	//    Id:     result.ID,
	//    TypeId: result.TypeID,
	//    Title:  result.Title,
	//    Desc:   result.Desc,
	//    Status: result.Status,
	//    Sort:   result.Sort,
	//}, nil
	return nil, nil
}

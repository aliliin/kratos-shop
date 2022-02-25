package service

import (
	"context"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateBrand 创建品牌
func (g *GoodsService) CreateBrand(ctx context.Context, r *v1.BrandRequest) (*v1.BrandInfoResponse, error) {
	brand, err := g.bc.CreateBrand(ctx, toBiz(r))
	if err != nil {
		return nil, err
	}

	return &v1.BrandInfoResponse{
		Id:    brand.ID,
		Name:  brand.Name,
		Logo:  brand.Logo,
		Desc:  brand.Desc,
		IsTab: brand.IsTab,
		Sort:  brand.Sort,
	}, nil
}

func (g *GoodsService) UpdateBrand(ctx context.Context, r *v1.BrandRequest) (*emptypb.Empty, error) {
	err := g.bc.UpdateBrand(ctx, toBiz(r))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func toBiz(r *v1.BrandRequest) *biz.Brand {
	return &biz.Brand{
		ID:    r.Id,
		Name:  r.Name,
		Logo:  r.Logo,
		Desc:  r.Desc,
		IsTab: r.IsTab,
		Sort:  r.Sort,
	}
}

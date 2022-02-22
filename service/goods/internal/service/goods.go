package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-kratos/kratos/v2/log"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

// GoodsService is a goods service.
type GoodsService struct {
	v1.UnimplementedGoodsServer

	gc  *biz.GoodsUsecase
	cac *biz.CategoryUsecase
	bc  *biz.BrandUsecase
	gic *biz.GoodsImageUsecase
	log *log.Helper
}

// NewGoodsService new a goods service.
func NewGoodsService(gc *biz.GoodsUsecase, cac *biz.CategoryUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{
		gc:  gc,
		cac: cac,
		log: log.NewHelper(logger),
	}
}

func (g *GoodsService) GetAllCategoryList(ctx context.Context, r *emptypb.Empty) (*v1.CategoryListResponse, error) {
	_, err := g.cac.CategoryList(ctx)
	if err != nil {
		return nil, err
	}

	res := &v1.CategoryListResponse{
		Total:    0,
		Data:     nil,
		JsonData: "",
	}
	return res, nil
}

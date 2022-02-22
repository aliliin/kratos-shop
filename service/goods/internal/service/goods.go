package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
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
	cate, err := g.cac.CategoryList(ctx)
	if err != nil {
		return nil, err
	}
	jsonData, _ := json.Marshal(cate)
	res := &v1.CategoryListResponse{
		Total:    0,
		Data:     nil,
		JsonData: string(jsonData),
	}
	return res, nil
}

package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGoodsService)

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
func NewGoodsService(gc *biz.GoodsUsecase, cac *biz.CategoryUsecase, bc *biz.BrandUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{
		gc:  gc,
		cac: cac,
		bc:  bc,
		log: log.NewHelper(logger),
	}
}

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
	cac *biz.CategoryUsecase
	bc  *biz.BrandUsecase
	gt  *biz.GoodsTypeUsecase
	s   *biz.SpecificationUsecase
	ga  *biz.GoodsAttrUsecase
	g   *biz.GoodsUsecase
	log *log.Helper
}

// NewGoodsService new a goods service.
func NewGoodsService(bc *biz.BrandUsecase, cac *biz.CategoryUsecase, gt *biz.GoodsTypeUsecase, s *biz.SpecificationUsecase,
	ga *biz.GoodsAttrUsecase, gc *biz.GoodsUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{
		bc:  bc,
		cac: cac,
		gt:  gt,
		s:   s,
		ga:  ga,
		g:   gc,
		log: log.NewHelper(logger),
	}
}

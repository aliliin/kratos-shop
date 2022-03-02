package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Goods struct {
	Hello string
}

type GoodsRepo interface {
}

type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(logger)}
}

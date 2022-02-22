package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Goods struct {
	Hello string
}

type GoodsRepo interface {
	CreateGreeter(context.Context, *Goods) error
	UpdateGreeter(context.Context, *Goods) error
}

type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GoodsUsecase) Create(ctx context.Context, g *Goods) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GoodsUsecase) Update(ctx context.Context, g *Goods) error {
	return uc.repo.UpdateGreeter(ctx, g)
}

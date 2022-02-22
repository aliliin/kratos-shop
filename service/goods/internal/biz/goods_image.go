package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type GoodsImages struct {
	Hello string
}

type GoodsImagesRepo interface {
	CreateGreeter(context.Context, *Goods) error
	UpdateGreeter(context.Context, *Goods) error
}

type GoodsImageUsecase struct {
	repo GoodsImagesRepo
	log  *log.Helper
}

func NewGoodsImagesUsecase(repo GoodsRepo, logger log.Logger) *GoodsImageUsecase {
	return &GoodsImageUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GoodsImageUsecase) Create(ctx context.Context, g *Goods) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GoodsImageUsecase) Update(ctx context.Context, g *Goods) error {
	return uc.repo.UpdateGreeter(ctx, g)
}

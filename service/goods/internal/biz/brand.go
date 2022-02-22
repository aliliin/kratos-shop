package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Brand struct {
	Hello string
}

type BrandRepo interface {
	CreateGreeter(context.Context, *Goods) error
	UpdateGreeter(context.Context, *Goods) error
}

type BrandUsecase struct {
	repo BrandRepo
	log  *log.Helper
}

func NewBrandUsecase(repo BrandRepo, logger log.Logger) *BrandUsecase {
	return &BrandUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BrandUsecase) Create(ctx context.Context, g *Goods) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *BrandUsecase) Update(ctx context.Context, g *Goods) error {
	return uc.repo.UpdateGreeter(ctx, g)
}

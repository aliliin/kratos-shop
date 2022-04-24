package biz

import (
	"cart/internal/domain"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type CartRepo interface {
	Create(ctx context.Context, c *domain.ShopCart) (*domain.ShopCart, error)
}

type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) CreateCart(ctx context.Context, c *domain.ShopCart) (*domain.ShopCart, error) {
	return uc.repo.Create(ctx, c)
}

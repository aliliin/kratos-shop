package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type InventoryRepo interface {
	Create(context.Context, *domain.Inventory) (*domain.Inventory, error)
}

type InventoryUsecase struct {
	repo InventoryRepo
	log  *log.Helper
}

func NewInventoryUsecase(repo InventoryRepo, logger log.Logger) *InventoryUsecase {
	return &InventoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

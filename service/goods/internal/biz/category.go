package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Categories struct {
	ID               int32
	Name             string
	ParentCategoryID int32
	SubCategory      []*Categories
	Level            int32
	IsTab            bool
	Sort             int32
}

type CategoryRepo interface {
	Category(context.Context) ([]*Categories, error)
	//CreateGreeter(context.Context, *Goods) error
}

type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (c *CategoryUsecase) CategoryList(ctx context.Context) ([]*Categories, error) {
	return c.repo.Category(ctx)
}

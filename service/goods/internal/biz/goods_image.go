package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type GoodsImages struct {
	Hello string
}

type GoodsImagesRepo interface {
	CreateGreeter(context.Context, *GoodsInfo) error
	UpdateGreeter(context.Context, *GoodsInfo) error
}

type GoodsImageUsecase struct {
	repo GoodsImagesRepo
	log  *log.Helper
}

func NewGoodsImagesUsecase(repo GoodsImagesRepo, logger log.Logger) *GoodsImageUsecase {
	return &GoodsImageUsecase{repo: repo, log: log.NewHelper(logger)}
}

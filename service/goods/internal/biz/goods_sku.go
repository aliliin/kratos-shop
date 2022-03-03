package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Sku struct {
	GoodsID        int64
	GoodsSn        string
	GoodsName      string
	SkuName        string
	SkuCode        string
	BarCode        string
	Price          int64
	PromotionPrice int64
	Points         int64
	RemarksInfo    string
	Pic            string
	Num            int64
	OnSale         bool
}

type GoodsSkuRepo interface {
	Create(context.Context, *Sku) (*Sku, error)
}

type GoodsSkuUsecase struct {
	repo GoodsSkuRepo
	log  *log.Helper
}

func NewGoodsSkuUsecase(repo GoodsSkuRepo, logger log.Logger) *GoodsSkuUsecase {
	return &GoodsSkuUsecase{repo: repo, log: log.NewHelper(logger)}
}

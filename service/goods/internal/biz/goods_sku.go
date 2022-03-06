package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Sku struct {
	ID             int64
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
	Inventory      int64
	OnSale         bool
	AttrInfo       string
}

type GoodsSpecificationSku struct {
	ID              int64
	SkuID           int64
	SkuCode         string
	Sort            int32
	SpecificationId int64
	ValueId         int64
}

type GoodsSkuRepo interface {
	Create(context.Context, *Sku) (*Sku, error)
	CreateSkuRelation(context.Context, []*GoodsSpecificationSku) error
}

type GoodsSkuUsecase struct {
	repo GoodsSkuRepo
	log  *log.Helper
}

func NewGoodsSkuUsecase(repo GoodsSkuRepo, logger log.Logger) *GoodsSkuUsecase {
	return &GoodsSkuUsecase{repo: repo, log: log.NewHelper(logger)}
}

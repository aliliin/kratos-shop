package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
)

type GoodsAttrValue struct {
	ID     int32
	AttrId int32
	Value  string
	Sort   int32
}

type GoodsAttr struct {
	ID                 int32
	TypeID             int32
	Name               string
	Sort               int32
	Status             bool
	IsSKU              bool
	IsSelect           bool
	SpecificationValue []*SpecificationValue
}

type GoodsAttrRepo interface {
	CreateSpecification(context.Context, *Specification) (int32, error)
	CreateSpecificationValue(context.Context, int32, []*SpecificationValue) error
}

type GoodsAttrUsecase struct {
	repo    GoodsAttrRepo
	tx      Transaction
	BrandUc *BrandUsecase
	log     *log.Helper
}

func NewGoodsAttrUsecase(repo GoodsAttrRepo, tx Transaction, logger log.Logger) *GoodsAttrUsecase {
	return &GoodsAttrUsecase{
		repo: repo,
		tx:   tx,
		log:  log.NewHelper(logger),
	}
}

func (ga *GoodsAttrUsecase) Create(ctx context.Context, g *GoodsAttr) error {
	return nil
}

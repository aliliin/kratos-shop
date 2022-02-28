package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Specification struct {
	ID       int32
	TypeID   int32
	Name     string
	Sort     int32
	Status   bool
	IsSKU    bool
	IsSelect bool
}

type SpecificationRepo interface {
	CreateSpecification(context.Context, *Specification) (int32, error)
}

type SpecificationUsecase struct {
	repo    SpecificationRepo
	TypeUc  *GoodsTypeUsecase
	tx      Transaction
	BrandUc *BrandUsecase
	log     *log.Helper
}

func NewSpecificationUsecase(repo SpecificationRepo, TypeUc *GoodsTypeUsecase, tx Transaction, BrandUc *BrandUsecase,
	logger log.Logger) *SpecificationUsecase {

	return &SpecificationUsecase{
		repo:    repo,
		TypeUc:  TypeUc,
		tx:      tx,
		BrandUc: BrandUc,
		log:     log.NewHelper(logger),
	}
}

// CreateSpecification 创建商品规格
func (s *SpecificationUsecase) CreateSpecification(ctx context.Context, r *Specification) (int32, error) {
	var (
		id  int32
		err error
	)
	if r.TypeID == 0 {
		return id, errors.New("请选择商品类型进行绑定")
	}
	// 去查询有没有这个类型
	err = s.TypeUc.IsTypeByID(ctx, r.TypeID)
	if err != nil {
		return id, err
	}
	id, err = s.repo.CreateSpecification(ctx, r)
	return id, err
}

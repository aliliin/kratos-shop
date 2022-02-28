package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type GoodsType struct {
	ID        int32
	Name      string
	TypeCode  string
	NameAlias string
	IsVirtual bool
	Desc      string
	Sort      int32
	BrandIds  string
}

type GoodsTypeRepo interface {
	CreateGoodsType(context.Context, *GoodsType) (int32, error)
	CreateGoodsBrandType(context.Context, int32, string) error
	GetGoodsTypeByID(context.Context, int32) (*GoodsType, error)
}

type GoodsTypeUsecase struct {
	repo    GoodsTypeRepo
	tx      Transaction
	BrandUc *BrandUsecase
	log     *log.Helper
}

func NewGoodsTypeUsecase(repo GoodsTypeRepo, tx Transaction, BrandUc *BrandUsecase, logger log.Logger) *GoodsTypeUsecase {
	return &GoodsTypeUsecase{
		repo:    repo,
		tx:      tx,
		BrandUc: BrandUc,
		log:     log.NewHelper(logger),
	}
}

// GoosTypeCreate 创建商品类型
func (gt *GoodsTypeUsecase) GoosTypeCreate(ctx context.Context, r *GoodsType) (int32, error) {
	var (
		id  int32
		err error
	)
	if r.BrandIds == "" {
		return id, errors.New("请选择品牌进行绑定")
	}
	err = gt.BrandUc.IsBrand(ctx, r.BrandIds)
	if err != nil {
		return id, err
	}
	// 使用事务
	err = gt.tx.ExecTx(ctx, func(ctx context.Context) error {
		id, err := gt.repo.CreateGoodsType(ctx, r)
		if err != nil {
			return err
		}
		// 创建商品类型关联关系表
		err = gt.repo.CreateGoodsBrandType(ctx, id, r.BrandIds)
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (gt *GoodsTypeUsecase) IsTypeByID(ctx context.Context, id int32) error {
	_, err := gt.repo.GetGoodsTypeByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

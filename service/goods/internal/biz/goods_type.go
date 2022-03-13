package biz

import (
	"context"
	"goods/internal/domain"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type GoodsTypeRepo interface {
	CreateGoodsType(context.Context, *domain.GoodsType) (int64, error)
	CreateGoodsBrandType(context.Context, int64, string) error
	GetGoodsTypeByID(context.Context, int64) (*domain.GoodsType, error)
	IsExistsByID(context.Context, int64) (*domain.GoodsType, error)
}

type GoodsTypeUsecase struct {
	repo  GoodsTypeRepo
	bRepo BrandRepo
	tx    Transaction
	log   *log.Helper
}

func NewGoodsTypeUsecase(repo GoodsTypeRepo, tx Transaction, BrandUc BrandRepo, logger log.Logger) *GoodsTypeUsecase {
	return &GoodsTypeUsecase{
		repo:  repo,
		tx:    tx,
		bRepo: BrandUc,
		log:   log.NewHelper(logger),
	}
}

// GoosTypeCreate 创建商品类型
func (gt *GoodsTypeUsecase) GoosTypeCreate(ctx context.Context, r *domain.GoodsType) (int64, error) {
	var (
		id  int64
		err error
	)

	if r.IsEmpty() {
		return id, errors.InternalServer("TYPE_IS_EMPTY", "请选择品牌进行绑定")
	}

	i, err := r.FormatBrandIds()
	if err != nil {
		return 0, err
	}

	brand, err := gt.bRepo.ListByIds(ctx, i...)
	if err != nil {
		return 0, err
	}

	if !brand.CheckLength(len(i)) {
		return 0, errors.InternalServer("BRAND_IS_EMPTY", "品牌不存在")
	}

	// 使用事务
	err = gt.tx.ExecTx(ctx, func(ctx context.Context) error {
		id, err = gt.repo.CreateGoodsType(ctx, r)
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

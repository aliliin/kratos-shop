package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
	"strconv"
	"strings"
)

type GoodsTypeRepo interface {
	CreateGoodsType(context.Context, *domain.GoodsType) (int64, error)
	CreateGoodsBrandType(context.Context, int64, string) error
	GetGoodsTypeByID(context.Context, int64) (*domain.GoodsType, error)
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
	if r.BrandIds == "" {
		return id, errors.New("请选择品牌进行绑定")
	}
	ids := strings.Replace(r.BrandIds, "，", ",", -1)
	Ids := strings.Split(ids, ",")

	var i []int32
	for _, bid := range Ids {
		j, _ := strconv.ParseInt(bid, 10, 32)
		i = append(i, int32(j))
	}

	err = gt.bRepo.IsBrand(ctx, i)
	if err != nil {
		return id, err
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

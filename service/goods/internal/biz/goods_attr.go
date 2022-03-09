package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type GoodsAttrRepo interface {
	CreateGoodsGroupAttr(context.Context, *domain.AttrGroup) (*domain.AttrGroup, error)
	CreateGoodsAttr(context.Context, *domain.GoodsAttr) (*domain.GoodsAttr, error)
	CreateGoodsAttrValue(context.Context, []*domain.GoodsAttrValue) ([]*domain.GoodsAttrValue, error)
	GetAttrByIDs(ctx context.Context, id []*int64) error
	ListByIds(ctx context.Context, id ...*int64) (domain.GoodsAttrList, error)
}

type GoodsAttrUsecase struct {
	repo  GoodsAttrRepo
	gRepo GoodsTypeRepo
	tx    Transaction
	log   *log.Helper
}

func NewGoodsAttrUsecase(repo GoodsAttrRepo, tx Transaction, gRepo GoodsTypeRepo, logger log.Logger) *GoodsAttrUsecase {
	return &GoodsAttrUsecase{
		repo:  repo,
		tx:    tx,
		gRepo: gRepo,
		log:   log.NewHelper(logger),
	}
}

func (ga *GoodsAttrUsecase) CreateAttrGroup(ctx context.Context, r *domain.AttrGroup) (*domain.AttrGroup, error) {
	if r.IsTypeIDEmpty() {
		return nil, errors.New("请选择商品类型进行绑定")
	}
	// 去查询有没有这个类型
	_, err := ga.gRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, err
	}

	attr, err := ga.repo.CreateGoodsGroupAttr(ctx, r)
	if err != nil {
		return nil, err
	}
	return attr, nil
}

func (ga *GoodsAttrUsecase) CreateAttrValue(ctx context.Context, r *domain.GoodsAttr) (*domain.GoodsAttr, error) {
	var (
		attrInfo *domain.GoodsAttr
		err      error
	)
	if r.IsTypeIDEmpty() {
		return attrInfo, errors.New("请选择商品类型进行绑定")
	}
	// 去查询有没有这个类型
	_, err = ga.gRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, errors.New("请选择商品类型进行绑定")
	}

	err = ga.tx.ExecTx(ctx, func(ctx context.Context) error {
		attr, err := ga.repo.CreateGoodsAttr(ctx, r)
		if err != nil {
			return err
		}
		var value []*domain.GoodsAttrValue
		for _, attrValue := range r.GoodsAttrValue {
			res := &domain.GoodsAttrValue{
				AttrId:  attr.ID,
				GroupID: attrValue.GroupID,
				Value:   attrValue.Value,
			}
			value = append(value, res)
		}
		attrValue, err := ga.repo.CreateGoodsAttrValue(ctx, value)
		if err != nil {
			return err
		}
		attrInfo = &domain.GoodsAttr{
			ID:             attr.ID,
			TypeID:         attr.TypeID,
			GroupID:        attr.GroupID,
			Title:          attr.Title,
			Sort:           attr.Sort,
			Status:         attr.Status,
			Desc:           attr.Desc,
			GoodsAttrValue: attrValue,
		}
		return nil
	})
	return attrInfo, nil
}

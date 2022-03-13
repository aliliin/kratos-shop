package biz

import (
	"context"
	"goods/internal/domain"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type GoodsAttrRepo interface {
	CreateGoodsGroupAttr(context.Context, *domain.AttrGroup) (*domain.AttrGroup, error)
	IsExistsGroupByID(ctx context.Context, id int64) (*domain.AttrGroup, error)
	CreateGoodsAttr(context.Context, *domain.GoodsAttr) (*domain.GoodsAttr, error)
	CreateGoodsAttrValue(context.Context, []*domain.GoodsAttrValue) ([]*domain.GoodsAttrValue, error)
	GetAttrByIDs(ctx context.Context, id []*int64) error
	ListByIds(ctx context.Context, id ...int64) (domain.GoodsAttrList, error)
}

type GoodsAttrUsecase struct {
	repo     GoodsAttrRepo
	typeRepo GoodsTypeRepo // 引入goods type 的 repo
	tx       Transaction   // 引入事务
	log      *log.Helper
}

func NewGoodsAttrUsecase(repo GoodsAttrRepo, tx Transaction, gRepo GoodsTypeRepo, logger log.Logger) *GoodsAttrUsecase {
	return &GoodsAttrUsecase{
		repo:     repo,
		tx:       tx,
		typeRepo: gRepo,
		log:      log.NewHelper(logger),
	}
}

func (ga *GoodsAttrUsecase) CreateAttrGroup(ctx context.Context, r *domain.AttrGroup) (*domain.AttrGroup, error) {
	if r.IsTypeIDEmpty() {
		return nil, errors.InternalServer("TYPE_IS_EMPTY", "请选择商品类型进行绑定")
	}

	_, err := ga.typeRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, err
	}

	attr, err := ga.repo.CreateGoodsGroupAttr(ctx, r)
	if err != nil {
		return nil, err
	}
	return attr, nil
}

// CreateAttrValue 创建商品属性和属性信息
func (ga *GoodsAttrUsecase) CreateAttrValue(ctx context.Context, r *domain.GoodsAttr) (*domain.GoodsAttr, error) {
	var (
		attrInfo  *domain.GoodsAttr
		attrValue []*domain.GoodsAttrValue
		err       error
	)
	if r.IsTypeIDEmpty() {
		return nil, errors.InternalServer("TYPE_IS_EMPTY", "请选择商品类型进行绑定")
	}

	_, err = ga.typeRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, err
	}

	_, err = ga.repo.IsExistsGroupByID(ctx, r.GroupID)
	if err != nil {
		return nil, err
	}

	err = ga.tx.ExecTx(ctx, func(ctx context.Context) error {
		attrInfo, err = ga.repo.CreateGoodsAttr(ctx, r)
		if err != nil {
			return err
		}
		var value []*domain.GoodsAttrValue
		for _, v := range r.GoodsAttrValue {
			if v.IsValueEmpty() {
				return errors.InternalServer("ATTR_IS_EMPTY", "商品属性不能为空")
			}
			res := &domain.GoodsAttrValue{
				AttrId:  attrInfo.ID,
				GroupID: v.GroupID,
				Value:   v.Value,
			}
			value = append(value, res)
		}
		attrValue, err = ga.repo.CreateGoodsAttrValue(ctx, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &domain.GoodsAttr{
		ID:             attrInfo.ID,
		TypeID:         attrInfo.TypeID,
		GroupID:        attrInfo.GroupID,
		Title:          attrInfo.Title,
		Sort:           attrInfo.Sort,
		Status:         attrInfo.Status,
		Desc:           attrInfo.Desc,
		GoodsAttrValue: attrValue,
	}, nil
}

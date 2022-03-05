package biz

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
)

type AttrGroup struct {
	ID     int64
	TypeID int32
	Title  string
	Desc   string
	Status bool
	Sort   int32
}

type GoodsAttrValue struct {
	ID      int64
	AttrId  int64
	GroupID int64
	Value   string
}

type GoodsAttr struct {
	ID             int64
	TypeID         int32
	GroupID        int64
	Title          string
	Sort           int32
	Status         bool
	Desc           string
	GoodsAttrValue []*GoodsAttrValue
}

type GoodsAttrRepo interface {
	CreateGoodsGroupAttr(context.Context, *AttrGroup) (*AttrGroup, error)
	CreateGoodsAttr(context.Context, *GoodsAttr) (*GoodsAttr, error)
	CreateGoodsAttrValue(context.Context, []*GoodsAttrValue) ([]*GoodsAttrValue, error)
	GetAttrByIDs(ctx context.Context, id []*int64) error
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

func (ga *GoodsAttrUsecase) CreateAttrGroup(ctx context.Context, r *AttrGroup) (*AttrGroup, error) {
	if r.TypeID == 0 {
		return nil, errors.New("请选择商品类型进行绑定")
	}
	// 去查询有没有这个类型
	_, err := ga.gRepo.GetGoodsTypeByID(ctx, r.TypeID)
	if err != nil {
		return nil, errors.New("请选择商品类型进行绑定")
	}

	attr, err := ga.repo.CreateGoodsGroupAttr(ctx, r)
	if err != nil {
		return nil, err
	}
	return attr, nil
}

func (ga *GoodsAttrUsecase) CreateAttrValue(ctx context.Context, r *GoodsAttr) (*GoodsAttr, error) {
	var (
		attrInfo *GoodsAttr
		err      error
	)
	if r.TypeID == 0 {
		return attrInfo, errors.New("请选择商品类型进行绑定")
	}
	// 去查询有没有这个类型
	_, err = ga.gRepo.GetGoodsTypeByID(ctx, r.TypeID)
	if err != nil {
		return attrInfo, errors.New("请选择商品类型进行绑定")
	}

	err = ga.tx.ExecTx(ctx, func(ctx context.Context) error {
		attr, err := ga.repo.CreateGoodsAttr(ctx, r)
		if err != nil {
			return err
		}
		var value []*GoodsAttrValue
		for _, attrValue := range r.GoodsAttrValue {
			res := &GoodsAttrValue{
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
		attrInfo = &GoodsAttr{
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

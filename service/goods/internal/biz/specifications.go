package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type SpecificationValue struct {
	ID     int64
	AttrId int64
	Value  string
	Sort   int32
}

type Specification struct {
	ID                 int64
	TypeID             int32
	Name               string
	Sort               int32
	Status             bool
	IsSKU              bool
	IsSelect           bool
	SpecificationValue []*SpecificationValue
}

type SpecificationRepo interface {
	CreateSpecification(context.Context, *Specification) (int64, error)
	CreateSpecificationValue(context.Context, int64, []*SpecificationValue) error
	GetSpecificationByIDs(context.Context, []*Specification) error
}

type SpecificationUsecase struct {
	repo  SpecificationRepo
	gRepo GoodsTypeRepo
	tx    Transaction
	log   *log.Helper
}

func NewSpecificationUsecase(repo SpecificationRepo, TypeUc GoodsTypeRepo, tx Transaction,
	logger log.Logger) *SpecificationUsecase {

	return &SpecificationUsecase{
		repo:  repo,
		gRepo: TypeUc,
		tx:    tx,
		log:   log.NewHelper(logger),
	}
}

// CreateSpecification 创建商品规格
func (s *SpecificationUsecase) CreateSpecification(ctx context.Context, r *Specification) (int64, error) {
	var (
		id  int64
		err error
	)
	if r.TypeID == 0 {
		return id, errors.New("请选择商品类型进行绑定")
	}

	if r.SpecificationValue == nil {
		return id, errors.New("请填写商品规格下的参数")
	}
	// 去查询有没有这个类型
	typeInfo, err := s.gRepo.GetGoodsTypeByID(ctx, r.TypeID)
	if err != nil {
		return int64(typeInfo.ID), err
	}

	// 使用事务
	err = s.tx.ExecTx(ctx, func(ctx context.Context) error {
		id, err = s.repo.CreateSpecification(ctx, r) // 插入选项
		if err != nil {
			return err
		}

		err = s.repo.CreateSpecificationValue(ctx, id, r.SpecificationValue) // 插入选项对应的值
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

type Brand struct {
	ID    int32
	Name  string
	Logo  string
	Desc  string
	IsTab bool
	Sort  int32
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type BrandRepo interface {
	Create(context.Context, *Brand) (*Brand, error)
	GetBradByName(context.Context, string) (*Brand, error)
	Update(context.Context, *Brand) error
	List(context.Context, *Pagination) ([]*Brand, int64, error)
	IsBrand(context.Context, []int32) error
}

type BrandUsecase struct {
	repo BrandRepo
	log  *log.Helper
}

func NewBrandUsecase(repo BrandRepo, logger log.Logger) *BrandUsecase {
	return &BrandUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BrandUsecase) CreateBrand(ctx context.Context, b *Brand) (*Brand, error) {
	_, err := uc.repo.GetBradByName(ctx, b.Name)
	if err != nil {
		return uc.repo.Create(ctx, b)
	} else {
		return nil, errors.New("当前品牌已经存在")
	}
}

func (uc *BrandUsecase) UpdateBrand(ctx context.Context, b *Brand) error {
	err := uc.repo.Update(ctx, b)
	if err != nil {
		return err
	}
	return nil
}

func (uc *BrandUsecase) BrandList(ctx context.Context, b *Pagination) ([]*Brand, int64, error) {
	list, total, err := uc.repo.List(ctx, b)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil

}

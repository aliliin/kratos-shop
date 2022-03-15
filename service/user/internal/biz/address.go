package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/domain"
)

//go:generate mockgen -destination=../mocks/mrepo/address.go -package=mrepo . AddressRepo
type AddressRepo interface {
	CreateAddress(ctx context.Context, a *domain.Address) (*domain.Address, error)
	AddressListByUid(ctx context.Context, uid int64) ([]*domain.Address, error)
	UpdateAddress(ctx context.Context, a *domain.Address) error
	DefaultAddress(ctx context.Context, a *domain.Address) error
	DeleteAddress(ctx context.Context, a *domain.Address) error
}

type AddressUsecase struct {
	repo AddressRepo
	log  *log.Helper
}

func NewAddressUsecase(repo AddressRepo, logger log.Logger) *AddressUsecase {
	return &AddressUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AddressUsecase) AddAddress(ctx context.Context, a *domain.Address) (*domain.Address, error) {
	return uc.repo.CreateAddress(ctx, a)
}

func (uc *AddressUsecase) AddressListByUid(ctx context.Context, uid int64) ([]*domain.Address, error) {
	return uc.repo.AddressListByUid(ctx, uid)
}

func (uc *AddressUsecase) UpdateAddress(ctx context.Context, a *domain.Address) error {
	return uc.repo.UpdateAddress(ctx, a)
}

func (uc *AddressUsecase) DefaultAddress(ctx context.Context, a *domain.Address) error {
	return uc.repo.DefaultAddress(ctx, a)
}

func (uc *AddressUsecase) DeleteAddress(ctx context.Context, a *domain.Address) error {
	return uc.repo.DeleteAddress(ctx, a)
}

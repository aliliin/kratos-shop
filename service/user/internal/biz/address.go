package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Address struct {
	ID        int64
	UserID    int64
	IsDefault int
	Mobile    string
	Name      string
	Province  string
	City      string
	Districts string
	Address   string
	PostCode  string
}

//go:generate mockgen -destination=../mocks/mrepo/address.go -package=mrepo . AddressRepo
type AddressRepo interface {
	CreateAddress(ctx context.Context, a *Address) (*Address, error)
}

type AddressUsecase struct {
	repo AddressRepo
	log  *log.Helper
}

func NewAddressUsecase(repo AddressRepo, logger log.Logger) *AddressUsecase {
	return &AddressUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AddressUsecase) AddAddress(ctx context.Context, a *Address) (*Address, error) {
	return uc.repo.CreateAddress(ctx, a)
}

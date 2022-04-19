package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

//var (
//	// ErrUserNotFound is user not found.
//	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
//)
type Greeter struct {
}

type CartRepo interface {
	Create(context.Context, *Greeter) (*Greeter, error)
}

type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) CreateCart(ctx context.Context, g *Greeter) (*Greeter, error) {
	return uc.repo.Create(ctx, g)
}

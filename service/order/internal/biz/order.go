package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"order/internal/domain"
)

//go:generate mockgen -destination=../mocks/mrepo/order.go -package=mrepo . OrderRepo
type OrderRepo interface {
	//CreateOrder(context.Context, *s) (*s, error)
}

type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (oc *OrderUsecase) CreateOrder(ctx context.Context, o *domain.CreateOrder) {

}

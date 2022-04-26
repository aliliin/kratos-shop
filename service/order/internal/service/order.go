package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "order/api/order/v1"
	"order/internal/biz"
	"order/internal/domain"
)

type OrderService struct {
	v1.UnimplementedOrderServer

	oc  *biz.OrderUsecase
	log *log.Helper
}

func NewOrderService(o *biz.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{oc: o, log: log.NewHelper(logger)}
}

func (o *OrderService) CreateOrder(ctx context.Context, r *v1.OrderRequest) (*v1.OrderInfoResponse, error) {
	o.oc.CreateOrder(ctx, &domain.CreateOrder{
		UserId:    r.UserId,
		AddressId: r.AddressId,
	})
	return &v1.OrderInfoResponse{}, nil
}

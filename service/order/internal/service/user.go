package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "order/api/order/v1"
	"order/internal/biz"
)

type OrderService struct {
	v1.UnimplementedOrderServer

	oc  *biz.OrderUsecase
	log *log.Helper
}

// NewOrderService new a order service.
func NewOrderService(o *biz.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{oc: o, log: log.NewHelper(logger)}
}

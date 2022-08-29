package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "order/api/order/v1"
	"order/internal/domain"
	"order/internal/usecase"
)

type OrderService struct {
	v1.UnimplementedOrderServer

	oc  *usecase.OrderUsecase
	log *log.Helper
}

func NewOrderService(o *usecase.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{oc: o, log: log.NewHelper(logger)}
}

func (o *OrderService) CreateOrder(ctx context.Context, r *v1.OrderRequest) (*v1.OrderInfoResponse, error) {
	var cartItem []*domain.CartItem
	for _, cart := range r.CartItem {
		res := &domain.CartItem{
			CartId:   cart.CartId,
			SkuId:    cart.SkuId,
			SkuPrice: cart.SkuPrice,
			SkuNum:   cart.SkuNum,
		}
		cartItem = append(cartItem, res)
	}

	o.oc.CreateOrder(ctx, &domain.CreateOrder{
		UserId:    r.UserId,
		AddressId: r.Address,
		CartItem:  cartItem,
	})
	return &v1.OrderInfoResponse{}, nil
}

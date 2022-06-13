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
	// 跨服务 查询用户收货地址
	// 跨服务（立即购买）查询商品信息、（购物车ID null）查询购物车商品信息
	// 跨服务 查询库存信息

	// 删除购物车的数据 （分布式 rocketmq）

	// 支付抠库存（分布式 rocketmq）
}

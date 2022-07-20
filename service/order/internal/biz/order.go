package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	cartV1 "order/api/cart/v1"
	goodsV1 "order/api/goods/v1"
	userV1 "order/api/user/v1"
	"order/internal/domain"
)

//go:generate mockgen -destination=../mocks/mrepo/order.go -package=mrepo . OrderRepo
type OrderRepo interface {
	GetAddressByID(ctx context.Context, aid int64, userId int64) (*domain.OrderAddress, error)
}

type OrderUsecase struct {
	repo     OrderRepo
	userRPC  userV1.UserClient
	cartRPC  cartV1.CartClient
	goodsRPC goodsV1.GoodsClient
	log      *log.Helper
}

func NewOrderUsecase(repo OrderRepo, userRPC userV1.UserClient, cartRPC cartV1.CartClient, goodsRPC goodsV1.GoodsClient,
	logger log.Logger) *OrderUsecase {

	return &OrderUsecase{
		repo:     repo,
		userRPC:  userRPC,
		cartRPC:  cartRPC,
		goodsRPC: goodsRPC,
		log:      log.NewHelper(logger),
	}
}

func (oc *OrderUsecase) CreateOrder(ctx context.Context, order *domain.CreateOrder) {
	// 跨服务（购物车购买）查询商品信息
	{
		// 已选中，根据用户ID，查询这个用户的所有已经选中的购物车商品
		//cartList, err := oc.cartRPC.ListCart(ctx, &cartV1.ListCartRequest{UserId: order.UserId})
		//bytes.Compare()

		for _, item := range order.CartItem {
			fmt.Println(item.CartId)
			fmt.Println()
			fmt.Println(item.SkuId)
		}

		// 商品ID，去查询商品对比价格
		// 商品里面价格对不上，就返回价格已经更改
	}
	// 跨服务 查询用户收货地址
	{
		address, err := oc.userRPC.GetAddress(ctx, &userV1.AddressReq{
			Id:  order.AddressId,
			Uid: order.UserId,
		})
		if err != nil {
			return
		}
		fmt.Println(address)
	}

	// 跨服务 查询库存信息
	// 删除购物车的数据 （分布式 rocketmq）
	// 支付抠库存（分布式 rocketmq）
}

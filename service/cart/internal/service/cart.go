package service

import (
	"cart/internal/biz"
	"cart/internal/domain"
	"context"

	v1 "cart/api/cart/v1"
)

type CartService struct {
	v1.UnimplementedCartServer
	cart *biz.CartUsecase
}

func NewCartService(cart *biz.CartUsecase) *CartService {
	return &CartService{cart:cart}
}

func (s *CartService) CreateCart(ctx context.Context, req *v1.CreateCartRequest) (*v1.CartInfo, error) {

	rv, err := s.cart.CreateCart(ctx, &domain.ShopCart{
		UserId:     req.UserId,
		GoodsId:    req.GoodsId,
		SkuId:      req.SkuId,
		GoodsPrice: req.GoodsPrice,
		GoodsNum:   req.GoodsNum,
		GoodsSn:    req.GoodsSn,
		GoodsName:  req.GoodsName,
		IsSelect:   req.IsSelect,
	})

	if err != nil {
		return nil, err
	}

	return &v1.CartInfo{
		Id:         rv.ID,
		UserId:     rv.UserId,
		GoodsId:    rv.GoodsId,
		GoodsSn:    rv.GoodsSn,
		GoodsName:  rv.GoodsName,
		SkuId:      rv.SkuId,
		GoodsPrice: rv.GoodsPrice,
		GoodsNum:   rv.GoodsNum,
		IsSelect:   rv.IsSelect,
	}, nil
}
//func (s *CartService) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) (*pb.UpdateCartReply, error) {
//	return &pb.UpdateCartReply{}, nil
//}
//func (s *CartService) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartReply, error) {
//	return &pb.DeleteCartReply{}, nil
//}
//func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartReply, error) {
//	return &pb.GetCartReply{}, nil
//}
//func (s *CartService) ListCart(ctx context.Context, req *pb.ListCartRequest) (*pb.ListCartReply, error) {
//	return &pb.ListCartReply{}, nil
//}

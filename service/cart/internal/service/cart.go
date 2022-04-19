package service

import (
	"cart/internal/biz"
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
	return &v1.CartInfo{}, nil
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

package service

import (
	"context"

	pb "cart/api/cart/v1"
)

type CartService struct {
	pb.UnimplementedCartServer
}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) CreateCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CreateCartReply, error) {
	return &pb.CreateCartReply{}, nil
}
func (s *CartService) UpdateCart(ctx context.Context, req *pb.UpdateCartRequest) (*pb.UpdateCartReply, error) {
	return &pb.UpdateCartReply{}, nil
}
func (s *CartService) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartReply, error) {
	return &pb.DeleteCartReply{}, nil
}
func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartReply, error) {
	return &pb.GetCartReply{}, nil
}
func (s *CartService) ListCart(ctx context.Context, req *pb.ListCartRequest) (*pb.ListCartReply, error) {
	return &pb.ListCartReply{}, nil
}

package service

import (
	"context"
	v1 "shop/api/shop/v1"
)

func (s *ShopService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return s.uc.CreateUser(ctx, req)
}

func (s *ShopService) Login(ctx context.Context, req *v1.LoginReq) (*v1.RegisterReply, error) {
	return s.uc.PassWordLogin(ctx, req)
}

func (s *ShopService) Detail(ctx context.Context, req *v1.DetailReq) (*v1.UserDetailResponse, error) {
	return s.uc.UserDetailByID(ctx, req)
}

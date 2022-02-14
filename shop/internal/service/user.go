package service

import (
	"context"
	v1 "shop/api/shop/v1"
)

func (s *ShopService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return s.uc.CreateUser(ctx, req)
}

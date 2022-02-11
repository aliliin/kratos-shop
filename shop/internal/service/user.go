package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "shop/api/shop/v1"
	"shop/internal/biz"
)

// ShopService is a greeter service.
type ShopService struct {
	v1.UnimplementedShopServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewShopService(uc *biz.UserUsecase, logger log.Logger) *ShopService {
	return &ShopService{uc: uc, log: log.NewHelper(log.With(logger, "module", "service/interface"))}
}

func (s *ShopService) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return s.uc.CreateUser(ctx, req)
}

package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "shop/api/shop/v1"
	"shop/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopService)

// ShopService is a greeter service.
type ShopService struct {
	v1.UnimplementedShopServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewShopService new a greeter service.
func NewShopService(uc *biz.UserUsecase, logger log.Logger) *ShopService {
	return &ShopService{uc: uc, log: log.NewHelper(log.With(logger, "module", "service/shop"))}
}

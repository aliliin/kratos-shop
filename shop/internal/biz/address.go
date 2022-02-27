package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	v1 "shop/api/shop/v1"
	"shop/internal/conf"
)

type Address struct {
	ID        int64
	UserID    int64
	IsDefault int
	Mobile    string
	Name      string
	Province  string
	City      string
	Districts string
	Address   string
	PostCode  string
}

type AddressRepo interface {
	CreateAddress(ctx context.Context, a *Address) (*Address, error)
}

type AddressUsecase struct {
	uRepo      UserRepo
	aRepo      AddressRepo
	log        *log.Helper
	signingKey string
}

func NewAddressUsecase(repo UserRepo, arepo AddressRepo, logger log.Logger, conf *conf.Auth) *AddressUsecase {
	helper := log.NewHelper(log.With(logger, "module", "usecase/shop"))
	return &AddressUsecase{
		uRepo:      repo,
		aRepo:      arepo,
		log:        helper,
		signingKey: conf.JwtKey}
}

func (ua *AddressUsecase) CreateAddress(ctx context.Context, r *v1.CreateAddressReq) (*v1.AddressInfo, error) {
	// 在上下文 context 中取出 claims 对象
	var uId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		i, ok := c["ID"].(float64)
		if !ok {
			return nil, ErrAuthFailed
		}
		uId = int64(i)
	}

	req := Address{
		UserID:    uId,
		IsDefault: 0,
		Mobile:    r.Mobile,
		Name:      r.Name,
		Province:  r.Province,
		City:      r.City,
		Districts: r.Districts,
		Address:   r.Address,
		PostCode:  r.PostCode,
	}
	res, err := ua.aRepo.CreateAddress(ctx, &req)
	if err != nil {
		return nil, err
	}

	result := &v1.AddressInfo{
		Id:        res.ID,
		Name:      res.Name,
		Mobile:    res.Mobile,
		Province:  res.Province,
		City:      res.City,
		Districts: res.Districts,
		Address:   res.Address,
		PostCode:  res.PostCode,
		IsDefault: int32(res.IsDefault),
	}
	return result, nil
}

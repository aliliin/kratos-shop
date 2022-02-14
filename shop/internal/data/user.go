package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userService "shop/api/service/user/v1"
	"shop/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (u *userRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	return nil, nil
}

func (u *userRepo) UserByMobile(c context.Context, mobile string) (*biz.User, error) {

	byMobile, err := u.data.uc.GetUserByMobile(c, &userService.MobileRequest{Mobile: mobile})
	if err != nil {
		return nil, err
	}
	return &biz.User{Mobile: byMobile.Mobile}, nil
}

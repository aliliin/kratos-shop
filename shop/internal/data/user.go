package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"shop/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) *biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (u *userRepo) UserByMobile(c context.Context, user *biz.User) (*biz.User, error) {
	return nil, nil
}

func (u *userRepo) CreateUser(c context.Context, mobile string) (*biz.User, error) {
	return nil, nil
}

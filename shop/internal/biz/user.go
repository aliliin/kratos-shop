package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "shop/api/shop/v1"
	"time"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrMobileInvalid   = errors.New("mobile invalid")
	ErrUserNotFound    = errors.New("user not found")
	ErrLoginFailed     = errors.New("login failed")
)

type User struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday *time.Time
	Gender   string
	Role     int
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
}

type UserUsecase struct {
	uRepo UserRepo
	log   *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	// check mobile
	_, err := uc.repo.UserByMobile(ctx, req.Mobile)
	if !errors.Is(err, ErrUserNotFound) {
		return nil, status.Errorf(codes.AlreadyExists, "The phone number has been registered.")
	}

	// create user
	user, err := NewUser(req.Mobile, req.Username, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "create user failed: %s", err.Error())
	}

	userId, err := uc.uRepo.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "save user failed: %s", err.Error())
	}
	return &v1.RegisterReply{
		Id: userId.ID,
	}, nil
}

func NewUser(mobile, username, password string) (User, error) {
	// check mobile
	if len(mobile) <= 0 {
		return User{}, ErrMobileInvalid
	}
	// check username
	if len(username) <= 0 {
		return User{}, ErrUsernameInvalid
	}
	// check password
	if len(password) <= 0 {
		return User{}, ErrPasswordInvalid
	}
	return User{
		NickName: username,
		Password: password,
	}, nil
}

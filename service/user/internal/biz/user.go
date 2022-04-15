package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          int64
	Mobile      string
	Password    string
	NickName    string
	Birthday    *time.Time
	Gender      string
	Role        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeletedAt bool
}

//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error)
	UserByMobile(ctx context.Context, mobile string) (*User, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	tr   Transaction
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, tr Transaction, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, tr: tr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	var (
		err  error
		user *User
	)
	err = uc.tr.ExecTx(ctx, func(ctx context.Context) error {
		user, err = uc.repo.CreateUser(ctx, u)
		if err != nil {
			return err
		}
		fmt.Println("ddd1111", user)
		user.NickName = "2222222"
		updateUser, err := uc.repo.UpdateUser(ctx, user)
		if err != nil {
			return err
		}
		fmt.Println("ddd", updateUser)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
	//return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUsecase) List(ctx context.Context, pageNum, pageSize int) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUsecase) UserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUsecase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}

func (uc *UserUsecase) UserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}

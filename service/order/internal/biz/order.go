package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type order struct {
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

//go:generate mockgen -destination=../mocks/mrepo/order.go -package=mrepo . orderRepo
type OrderRepo interface {
	//CreateOrder(context.Context, *order) (*order, error)
}

type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

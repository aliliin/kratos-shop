package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"order/internal/biz"
	"order/internal/domain"
	"time"
)

type OrderAddress struct {
	ID      int64  `gorm:"primarykey"`
	User    int64  `gorm:"type:int;index"`
	OrderSn string `gorm:"type:varchar(30);index"` // 订单号，我们平台自己生成的订单号

	// 用户收货信息
	Address         string `gorm:"type:varchar(100)"`
	RecipientName   string `gorm:"type:varchar(20)"`
	RecipientMobile string `gorm:"type:varchar(11)"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (OrderAddress) TableName() string {
	return "order_address"
}

type orderAddressRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderAddressRepo .
func NewOrderAddressRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p *OrderAddress) ToDomain() *domain.Order {
	return &domain.Order{
		ID:           0,
		User:         0,
		OrderSn:      "",
		PayType:      "",
		Status:       "",
		TradeNo:      "",
		OrderMount:   0,
		PayTime:      time.Time{},
		Address:      "",
		SignerName:   "",
		SingerMobile: "",
		Post:         "",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    time.Time{},
	}
}

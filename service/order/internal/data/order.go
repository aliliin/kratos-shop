package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"order/internal/biz"
	"order/internal/domain"
	"time"
)

type Order struct {
	ID            int64     `gorm:"primarykey"`
	User          int64     `gorm:"type:int not null;default:0;index"`
	OrderSn       string    `gorm:"type:varchar(30) not null;default:'';index"` // 订单号，我们平台自己生成的订单号
	OrderAmount   int64     `gorm:"type:int not null;default:0;comment '订单金额'"`
	GoodsAmount   int64     `gorm:"type:int not null;default:0;comment '商品总金额'"`
	OrderType     int64     `gorm:"type:tinyint(1) unsigned not null;default:0;comment '订单状态:1待支付2已支付3已发货4已签收5已取消6交易完成'"`
	ExpressAmount int64     `gorm:"type:int comment '运费'"`
	DeliveryAt    time.Time `gorm:"column:delivery_at comment '发货时间'"`
	RefundTime    time.Time `gorm:"type:datetime comment '退款时间'"`
	Post          string    `gorm:"type:varchar(200)  comment '订单备注信息'"`

	// 优惠信息、赠品、买反、优惠卷
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt
}

func (Order) TableName() string {
	return "orders"
}

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// NewOrderRepo .
func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p *Order) ToDomain() *domain.Order {
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

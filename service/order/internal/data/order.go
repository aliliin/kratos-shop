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
	User          int64     `gorm:"type:int;index"`
	OrderSn       string    `gorm:"type:varchar(30);index"` // 订单号，我们平台自己生成的订单号
	OrderAmount   int64     `gorm:"type:int;index"`
	GoodsAmount   int64     `gorm:"type:int;index"`
	OrderType     int64     `gorm:"type:int comment '1待支付2已支付3已发货4已签收5已取消6已退款'"`
	RefundTime    time.Time `gorm:"type:datetime comment '退款时间'"`
	ExpressAmount int64     `gorm:"type:int comment '运费"`
	DeliveryAt    time.Time `gorm:"column:delivery_at comment '发货时间'"`
	Post          string    `gorm:"type:varchar(20)  comment '订单备注信息'"`

	// 支付信息
	PayType   string    `gorm:"type:int comment '1alipay(支付宝)， 2wechat(微信)'"`
	PayStatus string    `gorm:"type:int  comment '1PAYING(待支付), 2TRADE_SUCCESS(成功)， 3TRADE_CLOSED(超时关闭), 4WAIT_BUYER_PAY(交易创建), 5TRADE_FINISHED(交易结束)'"`
	TradeNo   string    `gorm:"type:varchar(100) comment '交易号'"` // 交易号就是支付宝的订单号 查账
	PayTime   time.Time `gorm:"type:datetime"`

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

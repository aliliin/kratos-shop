package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"order/internal/biz"
	"time"
)

type Order struct {
	ID      int64  `gorm:"primarykey"`
	User    int64  `gorm:"type:int;index"`
	OrderSn string `gorm:"type:varchar(30);index"` // 订单号，我们平台自己生成的订单号
	PayType string `gorm:"type:varchar(20) comment 'alipay(支付宝)， wechat(微信)'"`
	// status 大家可以考虑使用 iota 来做
	Status       string `gorm:"type:varchar(20)  comment 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)'"`
	TradeNo      string `gorm:"type:varchar(100) comment '交易号'"` // 交易号就是支付宝的订单号 查账
	OrderMount   int64
	PayTime      *time.Time `gorm:"type:datetime"`
	Address      string     `gorm:"type:varchar(100)"`
	SignerName   string     `gorm:"type:varchar(20)"`
	SingerMobile string     `gorm:"type:varchar(11)"`
	Post         string     `gorm:"type:varchar(20)  comment '订单备注信息'"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt
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

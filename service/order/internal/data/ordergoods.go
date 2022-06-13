package data

import (
	"gorm.io/gorm"
	"time"
)

type OrderGoods struct {
	ID       int64  `gorm:"primarykey;type:int" json:"id"`
	OrderId  int64  `gorm:"type:int;index"`
	SkuId    int64  `gorm:"type:int;index"`
	SkuName  string `gorm:"type:varchar(100);index"`
	SkuPrice int64
	Nums     int32 `gorm:"type:int"`
	// 商品快照信息

	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (OrderGoods) TableName() string {
	return "order_goods"
}

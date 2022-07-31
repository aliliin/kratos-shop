package domain

import (
	"time"
)

type Order struct {
	ID           int64
	User         int64
	OrderSn      string
	PayType      string
	Status       string
	TradeNo      string
	OrderMount   int64
	PayTime      time.Time
	Address      string
	SignerName   string
	SingerMobile string
	Post         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type CreateOrder struct {
	UserId    int64
	AddressId int64
	CartItem  CartItemList
}

type CartItem struct {
	CartId   int64
	SkuId    int64
	SkuPrice int64
	SkuNum   int32
}

type CartItemList []*CartItem

func (p CartItemList) FindById(id int64) *CartItem {
	for _, item := range p {
		if item.CartId == id {
			return item
		}
	}
	return nil
}

func (p CartItemList) SkuId() []int64 {
	var l []int64
	for _, item := range p {
		l = append(l, item.SkuId)
	}
	return l
}

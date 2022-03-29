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

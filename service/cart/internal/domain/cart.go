package domain

type ShopCart struct {
	ID         int64
	UserId     int64
	GoodsId    int64
	SkuId      int64
	GoodsPrice int64
	GoodsNum   int32
	GoodsSn    string
	GoodsName  string
	IsSelect   bool
}

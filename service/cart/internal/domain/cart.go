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

type ShopCarts []*ShopCart

type ShopCartList struct {
	Total int32
	List  []*ShopCart
}

func (p ShopCarts) SelectedLists() []*ShopCart {
	var list []*ShopCart
	for _, cart := range p {
		if cart.IsSelect == true {
			list = append(list, cart)
		}
	}
	return list
}

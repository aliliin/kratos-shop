package domain

import "github.com/olivere/elastic/v7"

type ESGoodsFilter struct {
	ID          int64
	CategoryID  int32
	BrandsID    int32
	Keywords    string
	OnSale      bool
	ShipFree    bool
	IsNew       bool
	IsHot       bool
	ClickNum    int64
	SoldNum     int64
	FavNum      int64
	MaxPrice    int64
	MinPrice    int64
	Pages       int64
	PagePerNums int64
}

type ESGoods struct {
	ID           int64   `json:"id"`
	CategoryID   int32   `json:"category_id"`
	CategoryName string  `json:"category_name"`
	BrandsID     int32   `json:"brands_id"`
	BrandName    string  `json:"brand_name"`
	TypeID       int64   `json:"type_id"`
	TypeName     string  `json:"type_name"`
	OnSale       bool    `json:"on_sale"`
	ShipFree     bool    `json:"ship_free"`
	IsNew        bool    `json:"is_new"`
	IsHot        bool    `json:"is_hot"`
	Name         string  `json:"name"`
	GoodsTags    string  `json:"goods_tags"`
	ClickNum     int64   `json:"click_num"`
	SoldNum      int64   `json:"sold_num"`
	FavNum       int64   `json:"fav_num"`
	MarketPrice  int64   `json:"market_price"`
	GoodsBrief   string  `json:"goods_brief"`
	Pages        int64   `json:"pages"`
	PagePerNums  int64   `json:"page_pre_num"`
	Sku          []EsSku `json:"sku"`
}
type EsSku struct {
	SkuID    int64  `json:"sku_id"`
	SkuName  string `json:"sku_name"`
	SkuPrice int64  `json:"sku_price"`
}

type EsSearch struct {
	MustQuery    []elastic.Query
	MustNotQuery []elastic.Query
	ShouldQuery  []elastic.Query
	Filters      []elastic.Query
	Sorters      []elastic.Sorter
	Form         int64 // 分页
	Size         int64
}

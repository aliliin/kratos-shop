package data

type ESGoods struct {
	ID          int32  `json:"id"`
	CategoryID  int32  `json:"category_id"`
	BrandsID    int32  `json:"brands_id"`
	TypeID      int64  `json:"type_id"`
	OnSale      bool   `json:"on_sale"`
	ShipFree    bool   `json:"ship_free"`
	IsNew       bool   `json:"is_new"`
	IsHot       bool   `json:"is_hot"`
	Name        string `json:"name"`
	ClickNum    int64  `json:"click_num"`
	SoldNum     int64  `json:"sold_num"`
	FavNum      int64  `json:"fav_num"`
	MarketPrice int64  `json:"market_price"`
	GoodsBrief  string `json:"goods_brief"`
}

func (ESGoods) GetIndexName() string {
	return "goods"
}

func (ESGoods) GetMapping() string {
	goodsMapping := `
	{
		"mappings" : {
			"properties" : {
				"id" : {
					"type" : "integer"
				},
				"brands_id" : {
					"type" : "integer"
				},
				"category_id" : {
					"type" : "integer"
				},
				"type_id" : {
					"type" : "integer"
				},
				"click_num" : {
					"type" : "integer"
				},
				"fav_num" : {
					"type" : "integer"
				},
				"is_hot" : {
					"type" : "boolean"
				},
				"is_new" : {
					"type" : "boolean"
				},
				"market_price" : {
					"type" : "integer"
				},
				"name" : {
					"type" : "text",
					"analyzer":"ik_max_word"
				},
				"goods_brief" : {
					"type" : "text",
					"analyzer":"ik_max_word"
				},
				"on_sale" : {
					"type" : "boolean"
				},
				"ship_free" : {
					"type" : "boolean"
				},
				"shop_price" : {
					"type" : "integer"
				},
				"sold_num" : {
					"type" : "integer"
				}
			}
		}
	}`
	return goodsMapping
}

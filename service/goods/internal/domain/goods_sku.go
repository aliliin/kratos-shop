package domain

type GoodsSku struct {
	ID             int64
	GoodsID        int64
	GoodsSn        string
	GoodsName      string
	SkuName        string
	SkuCode        string
	BarCode        string
	Price          int64
	PromotionPrice int64
	Points         int64
	RemarksInfo    string
	Pic            string
	Inventory      int64
	OnSale         bool
	AttrInfo       string
	Specification  []*SpecificationInfo
	GroupAttr      []*GroupAttr
}

type SpecificationInfo struct {
	SpecificationID      int64
	SpecificationValueID int64
}

type GroupAttr struct {
	GroupId   int64   `json:"group_id"`
	GroupName string  `json:"group_name"`
	Attr      []*Attr `json:"attr"`
}

type Attr struct {
	AttrID        int64  `json:"attr_id"`
	AttrName      string `json:"attr_name"`
	AttrValueID   int64  `json:"attr_value_id"`
	AttrValueName string `json:"attr_value_name"`
}

type GoodsSpecificationSku struct {
	ID              int64
	SkuID           int64
	SkuCode         string
	SpecificationId int64
	ValueId         int64
}

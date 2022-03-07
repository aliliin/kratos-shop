package domain

type GoodsSku struct {
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
	Num            int64
	OnSale         bool

	Specification []*SpecificationInfo
	GroupAttr     []*GroupAttr
}

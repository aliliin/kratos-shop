package service

import (
	"golang.org/x/net/context"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
	"goods/internal/domain"
)

// CreateGoods 创建商品
func (g *GoodsService) CreateGoods(ctx context.Context, r *v1.CreateGoodsRequest) (*v1.CreateGoodsRequest, error) {
	var goodsSku []*domain.GoodsSku
	for _, sku := range r.Sku {
		res := &domain.GoodsSku{
			GoodsName:      r.Name,
			GoodsSn:        r.GoodsSn,
			SkuName:        sku.SkuName,
			SkuCode:        sku.Code,
			BarCode:        sku.BarCode,
			Price:          sku.Price,
			PromotionPrice: sku.PromotionPrice,
			Points:         sku.Points,
			//Pic:            sku.Pic,
			Num:    sku.Inventory,
			OnSale: r.OnSale,
		}

		for _, specification := range sku.SpecificationInfo {
			s := &biz.SpecificationInfo{
				SpecificationID:      specification.SId,
				SpecificationValueID: specification.VId,
			}
			res.Specification = append(res.Specification, s)
		}
		for _, attrGroup := range sku.GroupAttrInfo {
			group := &biz.GroupAttr{
				GroupId: attrGroup.GroupId,
			}
			for _, attr := range attrGroup.AttrInfo {
				s := &biz.Attr{
					AttrID:      attr.AttrId,
					AttrValueID: attr.AttrValueId,
				}
				group.Attr = append(group.Attr, s)
			}
			res.GroupAttr = append(res.GroupAttr, group)
		}
		goodsSku = append(goodsSku, res)
	}

	goodsInfo := &biz.GoodsInfo{
		ID:              r.Id,
		CategoryID:      int32(r.CategoryId),
		BrandsID:        int32(r.BrandId),
		TypeID:          int32(r.TypeId),
		Name:            r.Name,
		NameAlias:       r.NameAlias,
		GoodsSn:         r.GoodsSn,
		GoodsTags:       r.GoodsTags,
		MarketPrice:     r.MarketPrice,
		GoodsBrief:      r.GoodsBrief,
		GoodsFrontImage: r.GoodsFrontImage,
		GoodsImages:     r.GoodsImages,
		OnSale:          r.OnSale,
		ShipFree:        r.ShipFree,
		ShipID:          int32(r.ShipId),
		IsNew:           r.IsNew,
		IsHot:           r.IsHot,
		Sku:             goodsSku,
	}

	_, err := g.g.CreateGoods(ctx, goodsInfo)

	if err != nil {
		return nil, err
	}

	return r, nil
}

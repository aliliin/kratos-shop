package service

import (
	"context"
	v1 "goods/api/goods/v1"
	"goods/internal/domain"
)

// CreateGoods 创建商品
func (g *GoodsService) CreateGoods(ctx context.Context, r *v1.CreateGoodsRequest) (*v1.CreateGoodsResponse, error) {
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
			Pic:            sku.Image,
			Inventory:      sku.Inventory,
			OnSale:         r.OnSale,
		}

		for _, specification := range sku.SpecificationInfo {
			s := &domain.SpecificationInfo{
				SpecificationID:      specification.SId,
				SpecificationValueID: specification.VId,
			}
			res.Specification = append(res.Specification, s)
		}
		for _, attrGroup := range sku.GroupAttrInfo {
			group := &domain.GroupAttr{
				GroupId:   attrGroup.GroupId,
				GroupName: attrGroup.GroupName,
			}
			for _, attr := range attrGroup.AttrInfo {
				s := &domain.Attr{
					AttrID:        attr.AttrId,
					AttrName:      attr.AttrName,
					AttrValueID:   attr.AttrValueId,
					AttrValueName: attr.AttrValueName,
				}
				group.Attr = append(group.Attr, s)
			}
			res.GroupAttr = append(res.GroupAttr, group)
		}
		goodsSku = append(goodsSku, res)
	}

	goodsInfo := &domain.Goods{
		ID:              r.Id,
		CategoryID:      r.CategoryId,
		BrandsID:        r.BrandId,
		TypeID:          r.TypeId,
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
		ShipID:          r.ShipId,
		IsNew:           r.IsNew,
		IsHot:           r.IsHot,
		Sku:             goodsSku,
	}

	result, err := g.g.CreateGoods(ctx, goodsInfo)
	if err != nil {
		return nil, err
	}
	return &v1.CreateGoodsResponse{ID: result.GoodsID}, nil

}

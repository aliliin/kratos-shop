package data_test

import (
	"cart/internal/biz"
	"cart/internal/data"
	"cart/internal/domain"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cart", func() {
	var ro biz.CartRepo
	BeforeEach(func() {
		ro = data.NewCartRepo(Db, nil)
	})
	// 设置 It 块来添加单个规格
	It("CreateCart", func() {
		cartData := domain.ShopCart{
			UserId:     1,
			GoodsId:    1,
			SkuId:      1,
			GoodsPrice: 1000,
			GoodsNum:   10,
			GoodsSn:    "20232232231",
			GoodsName:  "Mate 40 Pro",
			IsSelect:   true,
		}
		c, err := ro.Create(ctx, &cartData)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(c.UserId).Should(Equal(int64(1)))
		Ω(c.GoodsNum).Should(Equal(int32(10)))

		// 二次验证创建相同商品的数据，只增加商品数量
		cartData2 := domain.ShopCart{
			UserId:     1,
			GoodsId:    1,
			SkuId:      1,
			GoodsPrice: 1000,
			GoodsNum:   10,
			GoodsSn:    "20232232231",
			GoodsName:  "Mate 40 Pro",
			IsSelect:   true,
		}
		c2, err := ro.Create(ctx, &cartData2)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(c2.UserId).Should(Equal(int64(1)))
		Ω(c2.GoodsNum).Should(Equal(int32(20)))
	})

})

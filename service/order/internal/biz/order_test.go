package biz_test

import (
	. "github.com/onsi/ginkgo"
	"order/internal/biz"
	"order/internal/mocks/mrepo"
)

var _ = Describe("orderUsecase", func() {
	var orderCase *biz.OrderUsecase
	var mOrderRepo *mrepo.MockorderRepo

	BeforeEach(func() {
		mOrderRepo = mrepo.NewMockorderRepo(ctl)
		orderCase = biz.NewOrderUsecase(mOrderRepo, nil)
	})
})

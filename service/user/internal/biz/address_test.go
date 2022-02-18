package biz_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/mocks/mrepo"
	"user/internal/testdata"
)

var _ = Describe("AddressUsecase", func() {
	var addressCase *biz.AddressUsecase
	var mAddressRepo *mrepo.MockAddressRepo

	BeforeEach(func() {
		mAddressRepo = mrepo.NewMockAddressRepo(ctl)
		addressCase = biz.NewAddressUsecase(mAddressRepo, nil)
	})

	It("Create", func() {
		info := testdata.Address()
		mAddressRepo.EXPECT().CreateAddress(ctx, gomock.Any()).Return(info, nil)
		l, err := addressCase.AddAddress(ctx, info)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(err).ToNot(HaveOccurred())
		Ω(l.Mobile).To(Equal("13509876789"))
	})

})

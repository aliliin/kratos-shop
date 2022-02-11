package biz_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/mocks/mrepo"
	"user/internal/testdata"
)

var _ = Describe("UserUsecase", func() {

	var sr *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)

		sr = biz.NewUserUsecase(mUserRepo, nil)
	})

	It("Create", func() {
		info := testdata.User(1)
		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)

		l, err := sr.Create(ctx, info)
		Ω(err).ToNot(HaveOccurred())
		Ω(l.ID).To(Equal(int64(1)))
	})

})

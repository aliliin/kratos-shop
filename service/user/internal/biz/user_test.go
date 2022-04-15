package biz_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/mocks/mrepo"
	"user/internal/mocks/usecase"
	"user/internal/testdata"
)

var _ = Describe("UserUsecase", func() {
	var userCase *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo
	var transaction *usecase.MockTransaction

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)
		transaction = usecase.NewMockTransaction(ctl)
		userCase = biz.NewUserUsecase(mUserRepo, transaction, nil)
	})

	FIt("Create", func() {
		//info := testdata.User()
		transaction.EXPECT().ExecTx(ctx, gomock.Any()).Return(nil)
		//mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)
		//mUserRepo.EXPECT().UpdateUser(ctx, gomock.Any()).Return(true, nil)
		//l, err := userCase.Create(ctx, info)
		//Ω(err).ShouldNot(HaveOccurred())
		//Ω(err).ToNot(HaveOccurred())
		//fmt.Println(l)
		//Ω(l.ID).To(Equal(int64(1)))
		//Ω(l.Mobile).To(Equal("13509876789"))
	})

	It("List", func() {
		info := testdata.User()
		info1 := testdata.User()
		info1.ID = 2
		info1.Mobile = "2323232323"
		u := []*biz.User{
			info,
			info1,
		}
		mUserRepo.EXPECT().ListUser(ctx, 1, 1).Return(u, 2, nil)
		list, total, err := userCase.List(ctx, 1, 1)
		if err != nil {
			return
		}
		Ω(err).ToNot(HaveOccurred())
		Ω(total).Should(Equal(2))
		Ω(list).ShouldNot(BeEmpty())
		Ω(list[0].ID).To(Equal(int64(1)))
		Ω(list[1].ID).To(Equal(int64(2)))
		Ω(list[0].Mobile).To(Equal("13509876789"))
		Ω(list[1].Mobile).To(Equal("2323232323"))
	})
})

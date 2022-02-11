package service_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "user/api/user/v1"
	"user/internal/biz"
	"user/internal/mocks/mrepo"
	"user/internal/service"
	"user/internal/testdata"
)

var _ = Describe("UserService", func() {

	var sr *service.UserService
	var userUsecase *biz.UserUsecase
	var mUserRepo *mrepo.MockUserRepo

	BeforeEach(func() {
		mUserRepo = mrepo.NewMockUserRepo(ctl)

		userUsecase = biz.NewUserUsecase(mUserRepo, nil)
		sr = service.NewUserService(userUsecase, nil)
	})

	It("CreateUser", func() {
		info := testdata.User(1)
		mUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(info, nil)

		uinfo := &v1.CreateUserInfo{
			Mobile:   "11",
			NickName: "",
		}

		l, err := sr.CreateUser(ctx, uinfo)
		Ω(err).ToNot(HaveOccurred())
		Ω(l.NickName).To(Equal("aaa"))
	})

})

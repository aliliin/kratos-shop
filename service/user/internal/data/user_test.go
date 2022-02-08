package data_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/data"
	"user/internal/testdata"
)

var _ = Describe("User", func() {

	var ro biz.UserRepo

	BeforeEach(func() {
		ro = data.NewUserRepo(Db, nil)

		f := func() {
			testData := testdata.User()
			_, err := ro.CreateUser(ctx, testData)
			Ω(err).ShouldNot(HaveOccurred())
		}

		f()
	})

	It("CreateUser", func() {
		_, err := ro.CreateUser(ctx, testdata.User())
		Ω(err).ShouldNot(HaveOccurred())
	})

})

package biz_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBiz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "biz user test")
}

var ctl *gomock.Controller
var cleaner func()
var ctx context.Context

var _ = BeforeEach(func() {
	ctl = gomock.NewController(GinkgoT())
	cleaner = ctl.Finish
	ctx = context.Background()
})
var _ = AfterEach(func() {
	// remove any mocks
	cleaner()
})

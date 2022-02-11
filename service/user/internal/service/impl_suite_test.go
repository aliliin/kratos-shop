package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestImpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "ServiceImpl Suite", []Reporter{reporters.NewJUnitReporter("junit.xml")})
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

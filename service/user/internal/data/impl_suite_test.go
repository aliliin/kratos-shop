package data_test

import (
	"context"
	"fmt"
	"testing"
	"user/internal/conf"
	"user/internal/data"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestImpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "RepoImpl Suite", []Reporter{reporters.NewJUnitReporter("junit.xml")})
}

var cleaner func()
var Db *data.Data
var ctx context.Context

var dbi data.DbInitializer
var _ = BeforeSuite(func() {
	con, f := data.DockerMysql("mysql/mysql-server", "5.7")
	fmt.Println(con)
	cleaner = f
	config := &conf.Data{Database: &conf.Data_Database{
		Source: con,
	},
	}

	db := data.NewDB(config)
	dataData, _, err := data.NewData(config, nil, db)
	Db = dataData
	dbi = data.DbInitializer{Ds: db}
	err = dbi.Initialize(ctx)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	//cleaner()
})

var _ = BeforeEach(func() {
})
var _ = AfterEach(func() {
})

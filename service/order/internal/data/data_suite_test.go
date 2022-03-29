package data_test

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"order/internal/conf"
	"order/internal/data"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// 测试 data 方法
func TestData(t *testing.T) {
	//  Ginkgo 测试通过调用 Fail(description string) 功能来表示失败
	// 使用 RegisterFailHandler 将此函数传递给 Gomega 。这是 Ginkgo 和 Gomega 之间的唯一连接点
	RegisterFailHandler(Fail)
	// 通知 Ginkgo 启动测试套件。如果您的任何 specs 失败，Ginkgo 将自动使 testing.T 失败。
	RunSpecs(t, "biz data test order")
}

var cleaner func()      // 定义删除 mysql 容器的回调函数
var Db *data.Data       // 用于测试的 data
var ctx context.Context // 上下文

// initialize  AutoMigrate gorm自动建表
func initialize(db *gorm.DB) error {
	err := db.AutoMigrate(
		&data.Order{},
	)
	return errors.WithStack(err)
}

// ginkgo 使用 BeforeEach 为您的 Specs 设置状态
var _ = BeforeSuite(func() {
	// 执行测试数据库操作之前，链接之前 docker 容器创建的 mysql
	//con, f := data.DockerMysql("mysql", "latest")
	con, f := data.DockerMysql("mariadb", "latest")
	cleaner = f // 测试完成，关闭容器的回调方法
	config := &conf.Data{Database: &conf.Data_Database{Driver: "mysql", Source: con}}
	db := data.NewDB(config)
	mySQLDb, _, err := data.NewData(config, nil, db, nil)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	Db = mySQLDb
	err = initialize(db)
	if err != nil {
		return
	}
	Expect(err).NotTo(HaveOccurred())
})

// 测试结束后 通过回调函数，关闭并删除 docker 创建的容器
var _ = AfterSuite(func() {
	cleaner()
})

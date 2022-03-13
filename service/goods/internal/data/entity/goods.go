package main

import (
	"goods/internal/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 链接数据库
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/shop_goods?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(
		//&data.Brand{},
		//&data.Category{},
		//&data.GoodsCategoryBrand{},
		//&data.GoodsType{},
		//&data.GoodsTypeBrand{},
		//&data.SpecificationsAttr{},
		//&data.SpecificationsAttrValue{},
		//&data.GoodsAttrGroup{},
		//&data.GoodsAttr{},
		//&data.GoodsAttrValue{},
		//&data.Goods{},
		//&data.GoodsSku{},
		//&data.GoodsImages{},
		&data.GoodsSpecificationSku{},
		&data.GoodsInventory{},
	)
}

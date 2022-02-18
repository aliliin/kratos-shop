package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"user/internal/data"
)

// 链接数据库
func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/shop_user?charset=utf8mb4&parseTime=True&loc=Local"
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
	//_ = db.AutoMigrate(&data.User{})
	_ = db.AutoMigrate(&data.Address{})
}

package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	slog "log"
	"os"
	"time"
	"user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewUserRepo)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db, rdb: rdb}, cleanup, nil
}

// NewDB .
//func NewDB(c *conf.Data) *gorm.DB {
//	// 终端打印输入 sql 执行记录
//	newLogger := logger.New(
//		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold: time.Second, // 慢查询 SQL 阈值
//			Colorful:      true,        // 禁用彩色打印
//			//IgnoreRecordNotFoundError: false,
//			LogLevel: logger.Info, // Log lever
//		},
//	)
//
//	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
//		Logger: newLogger,
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true, // 表名是否加 s
//		},
//	})
//
//	if err != nil {
//		log.Errorf("failed opening connection to sqlite: %v", err)
//		panic("failed to connect database")
//	}
//	return db
//}

//func NewRedis(c *conf.Data) *redis.Client {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:         c.Redis.Addr,
//		Password:     c.Redis.Password,
//		DB:           int(c.Redis.Db),
//		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
//		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
//		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
//	})
//	rdb.AddHook(redisotel.TracingHook{})
//	if err := rdb.Close(); err != nil {
//		log.Error(err)
//	}
//	return rdb
//}

// NewDB .
func NewDB(conf *conf.Data, l log.Logger) (*Data, func(), error) {
	log.NewHelper(l).Info("closing the data resources")
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			//IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	client, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名是否加 s
		},
	})
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Error; err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"layout/pkg/log"
	sLog "log"
	"os"
	"time"
)

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

func NewDB(conf *viper.Viper) *gorm.DB {
	var loggerAdapter logger.Interface
	if global.Config.Debug {
		loggerAdapter = logger.New(
			//将标准输出作为Writer
			sLog.New(os.Stdout, "\r\n", sLog.LstdFlags),
			logger.Config{
				//设定慢查询时间阈值为1ms
				SlowThreshold: 1 * time.Microsecond,
				//设置日志级别，只有Warn和Info级别会输出慢查询日志
				LogLevel: logger.Info,
			},
		)
	} else {
		loggerAdapter = newGormLogger()
	}
	gormConf := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: loggerAdapter,
	}

	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), gormConf)
	if err != nil {
		panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		logx.Channel(logx.Default).Error("获取MySQL连接错误", err)
		os.Exit(1)
	}
	conn.SetMaxIdleConns(conf.MaxIdleConns)
	conn.SetMaxOpenConns(conf.MaxOpenConns)
	_ = db.Use(&TracePlugin{})
	return db
}

func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"),
		DB:       conf.GetInt("data.redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}

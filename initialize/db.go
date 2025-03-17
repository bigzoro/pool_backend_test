package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"pool/global"
	"pool/internal/database/redisx"
	"pool/models"
	"time"
)

func InitDB() {
	mysqlConf := global.Config.MySQLConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Name)

	log.Println("dsn: ", dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	var err error
	global.GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = global.GormDB.AutoMigrate(
		&models.Pool{},
		&models.User{},
		&models.InviteCode{},
		&models.Purchase{},
		&models.Block{},
		&models.GoogleAuth{},
		&models.Plan{},
		&models.PlanDetail{},
		&models.Notice{},
		&models.Addresses{},
		&models.RechargeRecord{},
	)
	if err != nil {
		panic(err)
	}
}

func InitRedisDB(ctx context.Context) {
	option := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Config.RedisConfig.Host, global.Config.RedisConfig.Port),
		Password: global.Config.RedisConfig.Password,
		DB:       global.Config.RedisConfig.DBName,
		PoolSize: global.Config.RedisConfig.PoolSize,
	}
	global.RedisClient = redisx.InitRedisConn(ctx, option)
}

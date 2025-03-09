package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"pool/config"
)

var (
	Config      config.ServerConfig
	GormDB      *gorm.DB
	RedisClient *redis.Client
	//Trans       ut.Translator
)

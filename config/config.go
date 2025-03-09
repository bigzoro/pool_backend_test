package config

import "pool/internal/logx"

type ServerConfig struct {
	Port int
	//path                 string
	//Service              Service              `yaml:"service"`              // 基础服务
	Logger         *logx.LogConfig `mapstructure:"log" yaml:"log"`     // 日志
	MySQLConfig    MysqlConfig     `mapstructure:"mysql" yaml:"mysql"` // mysql 数据库
	RedisConfig    RedisConfig     `mapstructure:"redis" yaml:"redis"` // redis 数据库
	AwsEmailConfig AwsEmailConfig  `mapstructure:"aws_email" yaml:"aws_email"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
	DBName   int    `mapstructure:"dbName" json:"dbName"`
	PoolSize int    `mapstructure:"poolSize" json:"poolSize"`
	Expire   int    `mapstructure:"expire" json:"expire"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"dbName" json:"dbName"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type AwsEmailConfig struct {
	AwsAccessKeyID     string `mapstructure:"aws_access_key_id" json:"aws_access_key_id"`
	AwsSecretAccessKey string `mapstructure:"aws_secret_access_key" json:"aws_secret_access_key"`
	Region             string `mapstructure:"region" json:"region"`
	Sender             string `mapstructure:"sender" json:"sender"`
}

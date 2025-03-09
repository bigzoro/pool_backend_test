package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"pool/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	currDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	configDir := filepath.Join(currDir, "config")
	configFileName := filepath.Join(configDir, "config.yaml")
	v := viper.New()

	v.SetConfigFile(configFileName)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}

	//zap.S().Infof("配置信息: &v", global.NacosConfig)
	fmt.Println(global.Config)
}

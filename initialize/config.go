package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"pool/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//confFilePrefix := "config"

	//debug := GetEnvInfo("MXSHOP_DEBUG")
	//configFileName := fmt.Sprintf("meihua/%s-pro.yaml", confFilePrefix)
	//if debug {
	//	configFileName = fmt.Sprintf("user-web/%s-debug.yaml", confFilePrefix)
	//}

	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFileName := currDir + "\\config\\config.yaml"
	//configFileName := fmt.Sprintf("\\%s\\%s.yaml", currDir, confFilePrefix)
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

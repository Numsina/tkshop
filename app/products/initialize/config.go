package initialize

import (
	"log"
	"os"
	"path"

	"github.com/spf13/viper"

	"github.com/Numsina/tkshop/app/user/config"
)

var Conf config.Config

func InitConfig() {
	v := viper.New()

	// 自动加载环境变量
	v.AutomaticEnv()

	pwd, _ := os.Getwd()

	configName := "app/user/config/config.yaml"
	current := path.Join(pwd, configName)
	log.Println(current)
	// TODO
	// 获取当前的环境是否为线上还是线下(之后使用配置中心来配置)
	// ok := v.GetBool("dev")
	// if !ok {
	// 	// 线上环境
	// }

	v.SetConfigType("yaml")
	v.SetConfigFile(current)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic(err)
	}

}

package core

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/core/internal"
	"github.com/kesilent/react-light-blog/global"
	"github.com/spf13/viper"
)

// 读取配置文件
func Viper() *viper.Viper {
	var config string

	if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
		switch gin.Mode() {
		case gin.DebugMode:
			config = internal.ConfigDefaultFile
		case gin.ReleaseMode:
			config = internal.ConfigReleaseFile
		case gin.TestMode:
			config = internal.ConfigTestFile
		}
		fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), config)
	} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
		config = configEnv
		fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//监控配置文件的修改
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
	if err = v.Unmarshal(&global.RLB_CONFIG); err != nil {
		panic(err)
	}

	return v
}

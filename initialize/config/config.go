package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"mailgo/global"
	"mailgo/utils"
	"os"
)

// 加载配置文件
func init() {

	var config string

	if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
		// 默认配置
		config = utils.ConfigFile
		fmt.Println("loading default config file.")
	} else {
		// 环境变量中设置了新的配置， 通过 BA_CONFIG 配置
		config = configEnv
		fmt.Println("loading %v", config)
	}

	// 使用 viper 解析配置
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		fmt.Errorf("fatal error config file: %s ", err)
	}

	// 监听配置
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.BA_CONFIG); err != nil {
			fmt.Errorf("fatal error config file: %s ", err)
		}
	})

	if err := v.Unmarshal(&global.BA_CONFIG); err != nil {
		fmt.Errorf("fatal error config file: %s ", err)
	}
}

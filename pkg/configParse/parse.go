package configParse

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"layout/config"
	"layout/global"
	"os"
)

func Reload(path string) {
	_ = os.Setenv("APP_CONF", path)
	InitConfig()
}

func InitConfig() *config.Config {
	path := os.Getenv("APP_CONF")
	if path == "" {
		flag.StringVar(&path, "conf", "config/local.yml", "config path, eg: -conf config/local.yml")
		flag.Parse()
	}
	fmt.Println("load conf file:", path)
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conf.WatchConfig()
	conf.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := conf.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := conf.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
	return global.Config
}

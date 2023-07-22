package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"layout/global"
	"os"
)

func init() {
	path := os.Getenv("APP_CONF")
	if path == "" {
		flag.StringVar(&path, "conf", "config/local.yml", "config path, eg: -conf config/local.yml")
		flag.Parse()
	}
	if path == "" {
		path = "local"
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
}

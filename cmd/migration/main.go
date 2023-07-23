package main

import (
	"layout/cmd/migration/wireinject"
	"layout/global"
	"layout/pkg/configParse"
	"layout/pkg/redis"
)

func main() {
	configParse.InitConfig()
	global.Redis = redis.InitRedis()
	app, cleanup, err := wireinject.NewApp()
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()
}

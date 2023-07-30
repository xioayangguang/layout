package main

import (
	"fmt"
	"layout/cmd/server/wireinject"
	"layout/global"
	"layout/pkg/configParse"
	"layout/pkg/http"
	"layout/pkg/logx"
	"layout/pkg/redis"
)

// go build -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"
var (
	gitHash   string
	buildTime string
	goVersion string
)

// @title YoYo API
// @version 0.0.1
// @description This is a YoYo Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	configParse.InitConfig()
	global.GitHash = gitHash
	global.BuildTime = buildTime
	global.GoVersion = goVersion
	global.Redis = redis.InitRedis()
	engine, cleanup, err := wireinject.NewApp()
	if err != nil {
		panic(err)
	}
	logx.Channel(logx.Default).Info("server start http://127.0.0.1:", global.Config.Http.Port)
	http.Run(engine, fmt.Sprintf(":%d", global.Config.Http.Port))
	defer cleanup()
}

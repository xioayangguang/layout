package main

import (
	"layout/cmd/job/wireinject"
	"layout/global"
	"layout/pkg/configParse"
	"layout/pkg/redis"
)

// go build -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"
var (
	gitHash   string
	buildTime string
	goVersion string
)

func main() {
	configParse.InitConfig()
	global.GitHash = gitHash
	global.BuildTime = buildTime
	global.GoVersion = goVersion
	global.Redis = redis.InitRedis()
	app, cleanup, err := wireinject.NewApp()
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()
}

package main

import (
	"fmt"
	"layout/cmd/server/wireinject"
	"layout/global"
	"layout/pkg/config"
	"layout/pkg/http"
	"layout/pkg/logx"
)

// go build -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"
var (
	gitHash   string
	buildTime string
	goVersion string
)

func main() {
	global.GitHash = gitHash
	global.BuildTime = buildTime
	global.GoVersion = goVersion
	app, cleanup, err := wireinject.NewApp(config.InitConfig())
	if err != nil {
		panic(err)
	}
	logx.Channel(logx.Default).Info("server start http://127.0.0.1:", global.Config.Http.Port)
	http.Run(app, fmt.Sprintf(":%d", global.Config.Http.Port))
	defer cleanup()
}

package main

import (
	"fmt"
	"go.uber.org/zap"
	"layout/global"
	"layout/pkg/config"
	"layout/pkg/http"
	"layout/pkg/log"
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
	conf := config.NewConfig()
	logger := log.NewLog(conf)
	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))
	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()
}

#!/bin/bash

cd ..
git pull
name="everyday-rpc"
go build -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'" ./cmd/server/main.go -o $name
kill -1 $(cat pid.log)

#kill -2 $(cat pid.log)  优雅关机
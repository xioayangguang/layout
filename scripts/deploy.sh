#!/bin/bash

cd ..
git pull
name="everyday-rpc"
go build cmd/server/main.go -o $name -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"

kill -1 $(cat pid.log)

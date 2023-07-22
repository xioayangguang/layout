git pull
name="everyday-rpc"
port=9005

#go build -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'"
go build ../cmd/server/main.go -o $name

kill -1 $pid
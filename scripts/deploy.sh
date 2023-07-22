git pull
name="everyday-rpc"
port=9005

go build ../cmd/server/main.go -o $name

kill -1 $pid
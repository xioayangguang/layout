//go:build !windows
// +build !windows

package http

import (
	"github.com/fvbock/endless"
	"log"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"time"
)

func initServer(address string, handler http.Handler) server {
	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	//执行kill -1 pid命令发送syscall.SIGINT来通知程序优雅重启，具体做法如下：
	s := endless.NewServer(address, handler)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	s.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
		os.WriteFile("./scripts/pid.log", []byte(strconv.Itoa(syscall.Getpid())), os.ModePerm)
	}
	return s
}

//go:build windows
// +build windows

package http

import (
	"net/http"
	"time"
)

func initServer(address string, handler http.Handler) server {
	return &http.Server{
		Addr:           address,
		Handler:        handler,
		ReadTimeout:    35 * time.Second,
		WriteTimeout:   35 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

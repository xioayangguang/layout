package http

import (
	"log"
	"net/http"
)

type server interface {
	ListenAndServe() error
}

func Run(handler http.Handler, addr string) {
	s := initServer(addr, handler)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}

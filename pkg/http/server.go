package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

type server interface {
	ListenAndServe() error
}

func Run(r *gin.Engine, addr string) {
	s := initServer(addr, r)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}

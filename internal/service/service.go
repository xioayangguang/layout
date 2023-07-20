package service

import (
	"layout/internal/middleware"
	"layout/pkg/helper/sid"
	"layout/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *middleware.JWT
}

func NewService(logger *log.Logger, sid *sid.Sid, jwt *middleware.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}

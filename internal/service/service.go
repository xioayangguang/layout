package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewService,
	NewUserService,
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

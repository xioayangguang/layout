package service

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewService,
	NewUserService,
)

type Service struct {
	db *gorm.DB
}

func (s *Service) transaction(ctx context.Context, callBack func(ctx context.Context, tx *gorm.DB) error) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, "tx", tx)
		return callBack(ctx, tx)
	})
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

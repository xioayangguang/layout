package repository

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"layout/pkg/db"
)

var ProviderSet = wire.NewSet(
	db.NewDB,
	NewRepository,
	NewUserRepository,
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

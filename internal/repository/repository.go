package repository

import (
	"context"
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

func (r *Repository) getDb(ctx context.Context) (dbHandler *gorm.DB) {
	dbHandler, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		dbHandler = r.db
	}
	dbHandler.WithContext(ctx)
	return dbHandler
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

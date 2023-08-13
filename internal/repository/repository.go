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

	//context.valueCtx{}
	//ctx.Value()

	//context.WithValue()
	//context.WithCancel()
	//context.WithDeadline()
	//ctx.

	if !ok {
		dbHandler = r.db
	}

	return dbHandler.WithContext(ctx)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

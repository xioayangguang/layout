package migration

import (
	"gorm.io/gorm"
	"layout/internal/model"
)

type Migrate struct {
	db *gorm.DB
}

func NewMigrate(db *gorm.DB) *Migrate {
	return &Migrate{
		db: db,
	}
}
func (m *Migrate) Run() {
	m.db.AutoMigrate(&model.User{})
}

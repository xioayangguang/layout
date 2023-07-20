package migration

import (
	"gorm.io/gorm"
	"layout/internal/model"
	"layout/pkg/log"
)

type Migrate struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMigrate(db *gorm.DB, log *log.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}
func (m *Migrate) Run() {
	m.db.AutoMigrate(&model.User{})
	m.log.Info("AutoMigrate end")
}

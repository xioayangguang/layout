// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"layout/internal/migration"
	"layout/internal/repository"
	"layout/pkg/log"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*migration.Migrate, func(), error) {
	db := repository.NewDB(viperViper)
	migrate := migration.NewMigrate(db, logger)
	return migrate, func() {
	}, nil
}

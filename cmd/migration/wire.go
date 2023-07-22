//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"layout/internal/migration"
	"layout/internal/repository"
	_ "layout/pkg/config"
	_ "layout/pkg/redis"
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRepository,
	repository.NewUserRepository,
)
var MigrateSet = wire.NewSet(migration.NewMigrate)

func newApp() (*migration.Migrate, func(), error) {
	panic(wire.Build(
		RepositorySet,
		MigrateSet,
	))
}

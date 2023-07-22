//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"layout/internal/job"
	_ "layout/pkg/config"
	_ "layout/pkg/redis"
)

var JobSet = wire.NewSet(job.NewJob)

func newApp() (*job.Job, func(), error) {
	panic(wire.Build(
		JobSet,
	))
}

//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/google/wire"
	"layout/internal/job"
	_ "layout/pkg/config"
	_ "layout/pkg/redis"
)

var JobSet = wire.NewSet(job.NewJob)

func NewApp() (*job.Job, func(), error) {
	panic(wire.Build(
		JobSet,
	))
}

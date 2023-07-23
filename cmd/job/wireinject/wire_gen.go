// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireinject

import (
	"github.com/google/wire"
	"layout/internal/job"
)

// Injectors from wire.go:

func NewApp() (*job.Job, func(), error) {
	jobJob := job.NewJob()
	return jobJob, func() {
	}, nil
}

// wire.go:

var JobSet = wire.NewSet(job.NewJob)

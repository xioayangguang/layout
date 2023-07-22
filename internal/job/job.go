package job

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"layout/pkg/logx"
	"time"
)

type Job struct {
}

func NewJob() *Job {
	return &Job{}
}
func (j *Job) Run() {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.CronWithSeconds("0/3 * * * * *").Do(func() {
		logx.Channel(logx.Job).Info("I'm a Task1.")
	})
	if err != nil {
		fmt.Println(err)
	}
	_, err = s.Every("3s").Do(func() {
		logx.Channel(logx.Job).Info("I'm a Task2.")
	})
	if err != nil {
		fmt.Println(err)
	}
	s.StartBlocking()
}

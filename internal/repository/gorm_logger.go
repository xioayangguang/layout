package repository

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"layout/pkg/logx"
	"time"
)

type gormLogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

func newGormLogger() *gormLogger {
	return &gormLogger{
		SkipErrRecordNotFound: true,
		Debug:                 true,
	}
}

func (l *gormLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	logx.Channel(logx.Database).WithContext(ctx).Warnf(s, args)
}

func (l *gormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	logx.Channel(logx.Database).WithContext(ctx).Warnf(s, args)
}

func (l *gormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	logx.Channel(logx.Database).WithContext(ctx).Warnf(s, args)
}

// Trace type Fields map[string]interface{}
func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		logx.Channel(logx.Database).WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		//utils2.Log.WithChannel("Database").WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	if l.Debug {
		logx.Channel(logx.Database).WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
	}
}

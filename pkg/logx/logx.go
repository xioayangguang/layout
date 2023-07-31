package logx

import (
	"github.com/sirupsen/logrus"
	"layout/global"
	"layout/pkg/helper/rotatelogs"
	"layout/pkg/logx/formatter"
	"layout/pkg/logx/hook"
	"sync"
	"time"
)

const (
	Default  = "default"
	Database = "database"
	Request  = "request"
	Job      = "job"
	Panic    = "panic"
)

var loggerFormatter = map[string]logrus.Formatter{
	Request: &formatter.OnlyMsgFormatter{},
	Panic:   &formatter.OnlyMsgFormatter{},
}

var m sync.Map

func Channel(channel string) *logrus.Logger {
	c, ok := m.Load(channel)
	if ok {
		return c.(*logrus.Logger)
	} else {
		var channelLogger = logrus.New()
		channelLogger.SetOutput(rotatelogs.GetRotateLogs(channel))
		channelLogger.SetLevel(getLevel(global.Config.LogLevel))
		if f, ok := loggerFormatter[channel]; ok {
			channelLogger.SetFormatter(f)
		} else {
			channelLogger.SetFormatter(&logrus.JSONFormatter{
				TimestampFormat:   time.RFC3339,
				PrettyPrint:       false,
				DisableHTMLEscape: true,
			})
		}
		channelLogger.Hooks.Add(&hook.ErrorHook{})
		m.Store(channel, channelLogger)
		return channelLogger
	}
}

func getLevel(level string) logrus.Level {
	switch level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

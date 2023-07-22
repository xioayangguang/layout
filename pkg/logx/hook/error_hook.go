package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"layout/pkg/helper/rotatelogs"
)

type ErrorHook struct {
}

var logf = rotatelogs.GetRotateLogs("error")

func (h *ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.PanicLevel,
		logrus.FatalLevel,
	}
}

// Fire 将异常日志写入到指定日志文件中
func (h *ErrorHook) Fire(entry *logrus.Entry) error {
	errorMsg := fmt.Sprintf(
		"{\"time\":\"%s\",\"msg\":\"%s\",\"level\":\"%s\"}\r\n",
		entry.Time.String(),
		entry.Message,
		entry.Level.String(),
	)
	if _, err := logf.Write([]byte(errorMsg)); err != nil {
		return err
	}
	return nil
}

package formatter

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

type OnlyMsgFormatter struct {
}

func (m *OnlyMsgFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	newLog = fmt.Sprintf("[%s]\r\n %s\n", timestamp, entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}

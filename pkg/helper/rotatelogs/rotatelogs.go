package rotatelogs

import "time"
import rotatelogs "github.com/lestrrat-go/file-rotatelogs"

func GetRotateLogs(dir string) *rotatelogs.RotateLogs {
	logf, _ := rotatelogs.New(
		"./log/"+dir+"/%Y-%m-%d.log",
		//rotatelogs.WithLinkName("./log/"+dir+"/request.log"),
		rotatelogs.WithMaxAge(720*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	return logf
}

package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"layout/internal/response"
	"layout/internal/validate"
	"layout/pkg/logx"
	"net/http/httputil"
	"runtime"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
	//logger    *log.Logger = nil
)

// Recover 提前拦截前端业务恐慌
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpRequest, _ := httputil.DumpRequest(c.Request, true)
		defer func() {
			if err := recover(); err != nil {
				if err, ok := err.(*validate.ValidateError); ok {
					//response.FailWithCode(c, response.ParameterError)
					response.ValidationErrors(c, err.Error())
				} else {
					stack := stack(3)
					logx.Channel(logx.Panic).Printf("[Recovery] %s\r\n\r\n panic recovered:\n%s\n%s%s\r\n\r\n", httpRequest, err, stack)
					response.FailWithCode(c, response.Error)
				}
				//c.Abort()
			}
		}()
		c.Next()
	}
}

func stack(skip int) []byte {
	buf := new(bytes.Buffer)
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		_, _ = fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

func source(lines [][]byte, n int) []byte {
	n--
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

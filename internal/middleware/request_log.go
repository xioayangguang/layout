package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"layout/pkg/db"
	"layout/pkg/helper/snowflake"
	"layout/pkg/logx"
	"net/http/httputil"
	"strings"
	"time"
)

var Status = map[int]string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-status",
	208: "Already Reported",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	306: "Switch Proxy",
	307: "Temporary Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Time-out",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request-URI Too Large",
	415: "Unsupported Media Type",
	416: "Requested range not satisfiable",
	417: "Expectation Failed",
	418: "I\"m a teapot",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Unordered Collection",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Time-out",
	505: "HTTP Version not supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	511: "Network Authentication Required",
}

type bodyWriter struct {
	gin.ResponseWriter
	bodyCache *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyCache.Write(b)
	return w.ResponseWriter.Write(b)
}

var logStr = ""

// RequestLog 请求响应日志处理
func RequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBefore(ctx)
		ctx.Next()
		requestAfter(ctx)
	}
}

func requestBefore(ctx *gin.Context) {
	requestID := snowflake.GlobalSnowflake.Generate().String()
	c := context.WithValue(ctx.Request.Context(), "Request-Id", requestID)
	ctx.Request = ctx.Request.WithContext(c)
	httpRequest, _ := httputil.DumpRequest(ctx.Request, true)
	logStr = fmt.Sprintf("requestID:%v\r\n", requestID)
	logStr = fmt.Sprintf("%v%v\r\n", logStr, string(httpRequest))
	ctx.Writer = &bodyWriter{bodyCache: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Header("Request-Id", requestID)
	ctx.Header("Date", time.Now().Format(time.RFC1123))
}

func requestAfter(ctx *gin.Context) {
	logStr = fmt.Sprintf("%vHTTP/1.1 %v %v\r\n", logStr, ctx.Writer.Status(), Status[ctx.Writer.Status()])
	for k, v := range ctx.Writer.Header() {
		logStr = fmt.Sprintf("%v%v: %v\r\n", logStr, k, strings.Join(v, ","))
	}
	logStr = fmt.Sprintf("%vContent-Length: %v\r\n\r\n", logStr, ctx.Writer.Size())
	bw, ok := ctx.Writer.(*bodyWriter)
	if ok {
		var strB strings.Builder
		strB.WriteString(string(bw.bodyCache.Bytes()))
		logStr = fmt.Sprintf("%v%v\r\n", logStr, strB.String())
		logStr = fmt.Sprintf("%v%v\r\n", logStr, strings.Repeat("-", 100))
	}
	c := ctx.Request.Context()
	if c != nil {
		if requestId, ok := c.Value("Request-Id").(string); ok {
			sqls := db.GetAllSql(requestId)
			if sqls != nil {
				sqlLog, _ := json.Marshal(sqls)
				logStr = fmt.Sprintf("%v%v\r\n\r\n", logStr, string(sqlLog))
				db.ClearSql(requestId)
			}
		}
	}
	logx.Channel(logx.Request).Warning(logStr)
}

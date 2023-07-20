package response

import (
	"horse/pkg/berror"
	"horse/pkg/logx"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, c *gin.Context) {
	language := strings.ToLower(c.Request.Header.Get("Language"))
	msg := ""
	if language == "zh-cn" {
		msg = ZhCnTranslate[code]
	} else if language == "zh-tw" {
		msg = ZhTwTranslate[code]
	} else {
		msg = EnTranslate[code]
	}
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(Success, map[string]interface{}{}, c)
}
func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, data, c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]interface{}{}, c)
}

func FailWithData(data interface{}, c *gin.Context) {
	Result(Error, data, c)
}

func FailWithCode(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, c)
}

func FailWithError(err error, c *gin.Context) {
	code := berror.GetCode(err)
	if code == -1 {
		logx.Channel(logx.Default).Error(err.Error())
		FailWithCode(Error, c)
	} else {
		FailWithCode(code, c)
	}
}

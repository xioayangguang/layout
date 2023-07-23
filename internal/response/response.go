package response

import (
	"layout/pkg/berror"
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
func OkWithData(c *gin.Context, data interface{}) {
	Result(Success, data, c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]interface{}{}, c)
}

func FailWithData(c *gin.Context, data interface{}) {
	Result(Error, data, c)
}

func FailWithCode(c *gin.Context, code int) {
	Result(code, map[string]interface{}{}, c)
}

func ValidationErrors(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		ParameterError,
		map[string]interface{}{},
		msg,
	})
}

func FailWithError(c *gin.Context, err error) {
	code := berror.GetCode(err)
	if code == -1 {
		FailWithCode(c, Error)
	} else {
		FailWithCode(c, code)
	}
}

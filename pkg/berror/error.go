package berror

import "fmt"

type CodeErr struct {
	code int
	msg  string
}

func (e CodeErr) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", e.code, e.msg)
}

func New(code int) error {
	return CodeErr{
		code: code,
		msg:  "",
	}
}
func NewWithMsg(code int, msg string) error {
	return CodeErr{
		code: code,
		msg:  msg,
	}
}

func GetCode(err error) int {
	if e, ok := err.(CodeErr); ok {
		return e.code
	}
	return -1
}

func GetMsg(err error) string {
	if e, ok := err.(CodeErr); ok {
		return e.msg
	}
	return ""
}

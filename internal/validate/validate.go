package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var Validator *validator.Validate
var Trans ut.Translator

// New 初始化验证器
func init() {
	Validator, _ = binding.Validator.Engine().(*validator.Validate)
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	//注册翻译器
	Trans, _ = ut.New(zh.New()).GetTranslator("zh")
	_ = zhTranslations.RegisterDefaultTranslations(Validator, Trans)
	addressValid()
	bankCard()
	idCard()
	mobile()
	ruleRealName()
}

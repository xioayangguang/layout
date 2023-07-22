package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func mobile() {
	_ = Validator.RegisterValidation("mobile", func(f validator.FieldLevel) bool {
		ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, f.Field().String())
		return ok
	})
	_ = Validator.RegisterTranslation("mobile", Trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "手机号格式不正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})
}

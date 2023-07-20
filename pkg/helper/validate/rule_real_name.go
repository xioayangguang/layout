package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ruleRealName() {
	_ = Validator.RegisterValidation("real_name", func(f validator.FieldLevel) bool {
		regRuler := "^[\\x{4e00}-\\x{9fa5}]{2,6}$"
		reg := regexp.MustCompile(regRuler)
		return reg.MatchString(f.Field().String())
	})
	_ = Validator.RegisterTranslation("real_name", Trans, func(ut ut.Translator) error {
		return ut.Add("real_name", "姓名格式不正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("real_name", fe.Field())
		return t
	})
}

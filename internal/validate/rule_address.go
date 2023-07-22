package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func addressValid() {
	_ = Validator.RegisterValidation("address", func(f validator.FieldLevel) bool {
		address := f.Field().String()
		re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
		return re.MatchString(address)
	})
	_ = Validator.RegisterTranslation("address", Trans, func(ut ut.Translator) error {
		return ut.Add("address_error", "钱包地址格式不正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("address_error", fe.Field())
		return t
	})
}

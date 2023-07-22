package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

func bankCard() {
	_ = Validator.RegisterValidation("bank_card", func(f validator.FieldLevel) bool {
		patternText := f.Field().String()
		sanitizedValue := strings.Replace(patternText, "-", "", -1)
		numberLen := len(sanitizedValue)
		sum := 0
		alternate := false
		if numberLen < 13 || numberLen > 19 {
			return false
		}
		for i := numberLen - 1; i > -1; i-- {
			mod := int(byte(sanitizedValue[i]) - '0')
			if alternate {
				mod *= 2
				if mod > 9 {
					mod = (mod % 10) + 1
				}
			}
			alternate = !alternate
			sum += mod
		}
		return sum%10 == 0
	})
	_ = Validator.RegisterTranslation("bank_card", Trans, func(ut ut.Translator) error {
		return ut.Add("bank_card_error", "银行卡格式不正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("bank_card_error", fe.Field())
		return t
	})
}

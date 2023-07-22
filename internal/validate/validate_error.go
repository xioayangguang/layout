package validate

func NewValidateError(text string) error {
	return &ValidateError{text}
}

type ValidateError struct {
	s string
}

func (e *ValidateError) Error() string {
	return e.s
}

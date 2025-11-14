package errorpkg

type ErrorStack struct {
	Errors []error
}

func (e *ErrorStack) Add(err error) {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}
}

func (e *ErrorStack) Error() string {
	combined := Combine(e.Errors)
	if combined != nil {
		return combined.Error()
	}
	return ""
}
func CombineErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return ValidationError{All: errs}
}

package rules

import (
	"strings"
)

// CombineErrors combines multiple errors into one
type ValidationError struct {
	All []error
}

func (v ValidationError) Error() string {
	var sb strings.Builder
	for i, err := range v.All {
		sb.WriteString(err.Error())
		if i != len(v.All)-1 {
			sb.WriteString("; ")
		}
	}
	return sb.String()
}

func CombineErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return ValidationError{All: errs}
}

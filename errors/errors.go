package errorpkg

import (
	"errors"
	"strings"
)

func Combine(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}

type ValidationError struct {
	All []error
}

func (ve ValidationError) Error() string {
	var out []string
	for _, e := range ve.All {
		out = append(out, e.Error())
	}
	return strings.Join(out, "; ")
}

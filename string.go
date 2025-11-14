package govalid

import (
	errorpkg "govalid/errors"
	"govalid/rules"
)

type StringValidator struct {
	value string
	errs  []error
}

func String(v string) *StringValidator {
	return &StringValidator{value: v}
}

//// BASIC RULES ////////////////////////////

func (s *StringValidator) Required() *StringValidator {
	if err := rules.StringRequired(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) Empty() *StringValidator {
	if err := rules.StringEmpty(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

//// LENGTH RULES ///////////////////////////

func (s *StringValidator) Min(n int) *StringValidator {
	if err := rules.StringMin(s.value, n); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) Max(n int) *StringValidator {
	if err := rules.StringMax(s.value, n); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) MinLength(n int) *StringValidator {
	if err := rules.MinLength(s.value, n); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) MaxLength(n int) *StringValidator {
	if err := rules.MaxLength(s.value, n); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) LengthRange(min, max int) *StringValidator {
	if err := rules.LengthRange(s.value, min, max); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

//// CONTENT RULES //////////////////////////

func (s *StringValidator) Alphanumeric() *StringValidator {
	if err := rules.StringAlphanumeric(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) Alphabetic() *StringValidator {
	if err := rules.StringAlphabetic(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) Numeric() *StringValidator {
	if err := rules.StringNumeric(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) Contains(substr string) *StringValidator {
	if err := rules.StringContains(s.value, substr); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) StartsWith(prefix string) *StringValidator {
	if err := rules.StringStartsWith(s.value, prefix); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) EndsWith(suffix string) *StringValidator {
	if err := rules.StringEndsWith(s.value, suffix); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) OneOf(list ...string) *StringValidator {
	if err := rules.StringOneOf(s.value, list); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

//// REGEX RULES ////////////////////////////

func (s *StringValidator) Match(regex string) *StringValidator {
	if err := rules.StringMatch(s.value, regex); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

//// FORMAT VALIDATIONS /////////////////////

func (s *StringValidator) Email() *StringValidator {
	if err := rules.StringEmail(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) URL() *StringValidator {
	if err := rules.StringURL(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) UUID() *StringValidator {
	if err := rules.StringUUID(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) IP() *StringValidator {
	if err := rules.StringIP(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) IPv4() *StringValidator {
	if err := rules.StringIPv4(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) IPv6() *StringValidator {
	if err := rules.StringIPv6(s.value); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

func (s *StringValidator) Date(format string) *StringValidator {
	if err := rules.StringDate(s.value, format); err != nil {
		s.errs = append(s.errs, err)
	}
	return s
}

//// FINAL ERROR RETURN /////////////////////

func (s *StringValidator) Error() error {
	return errorpkg.Combine(s.errs)
}

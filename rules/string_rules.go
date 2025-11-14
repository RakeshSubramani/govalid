package rules

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

// ----------------- BASIC -----------------

func StringRequired(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return errors.New("value is required")
	}
	return nil
}

func StringEmpty(s string) error {
	if len(s) != 0 {
		return errors.New("value must be empty")
	}
	return nil
}

// ----------------- LENGTH -----------------

func MinLength(s string, n int) error {
	if len(s) < n {
		return fmt.Errorf("minimum length is %d", n)
	}
	return nil
}

func MaxLength(s string, n int) error {
	if len(s) > n {
		return fmt.Errorf("maximum length is %d", n)
	}
	return nil
}

func LengthRange(s string, min, max int) error {
	l := len(s)
	if l < min || l > max {
		return fmt.Errorf("length must be between %d and %d", min, max)
	}
	return nil
}

func StringMin(s string, n int) error {
	if len(s) < n {
		return errors.New("value length is too short")
	}
	return nil
}

func StringMax(s string, n int) error {
	if len(s) > n {
		return errors.New("value length is too long")
	}
	return nil
}

// ----------------- CONTENT -----------------

func StringAlphanumeric(s string) error {
	for _, ch := range s {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return errors.New("value must be alphanumeric")
		}
	}
	return nil
}

func StringAlphabetic(s string) error {
	for _, ch := range s {
		if !unicode.IsLetter(ch) {
			return errors.New("value must contain only letters")
		}
	}
	return nil
}

func StringNumeric(s string) error {
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return errors.New("value must contain only digits")
		}
	}
	return nil
}

func StringContains(s, substr string) error {
	if !strings.Contains(s, substr) {
		return fmt.Errorf("value must contain '%s'", substr)
	}
	return nil
}

func StringStartsWith(s, prefix string) error {
	if !strings.HasPrefix(s, prefix) {
		return fmt.Errorf("value must start with '%s'", prefix)
	}
	return nil
}

func StringEndsWith(s, suffix string) error {
	if !strings.HasSuffix(s, suffix) {
		return fmt.Errorf("value must end with '%s'", suffix)
	}
	return nil
}

func StringOneOf(s string, list []string) error {
	for _, v := range list {
		if s == v {
			return nil
		}
	}
	return fmt.Errorf("value must be one of %v", list)
}

// ----------------- REGEX -----------------

func StringMatch(s string, pattern string) error {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return fmt.Errorf("invalid regex pattern: %v", err)
	}
	if !matched {
		return fmt.Errorf("value does not match pattern '%s'", pattern)
	}
	return nil
}

// ----------------- FORMAT -----------------

func StringEmail(s string) error {
	_, err := mail.ParseAddress(s)
	if err != nil {
		return errors.New("invalid email address")
	}
	return nil
}

func StringURL(s string) error {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return errors.New("invalid URL")
	}
	return nil
}

func StringUUID(s string) error {
	uuidRegex := `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`
	return StringMatch(s, uuidRegex)
}

func StringIP(s string) error {
	ipv4Regex := `^(\d{1,3}\.){3}\d{1,3}$`
	ipv6Regex := `^([0-9a-fA-F]{0,4}:){1,7}[0-9a-fA-F]{0,4}$`
	if err := StringMatch(s, ipv4Regex); err != nil {
		if err := StringMatch(s, ipv6Regex); err != nil {
			return errors.New("invalid IP address")
		}
	}
	return nil
}

func StringIPv4(s string) error {
	ipv4Regex := `^(\d{1,3}\.){3}\d{1,3}$`
	return StringMatch(s, ipv4Regex)
}

func StringIPv6(s string) error {
	ipv6Regex := `^([0-9a-fA-F]{0,4}:){1,7}[0-9a-fA-F]{0,4}$`
	return StringMatch(s, ipv6Regex)
}

func StringDate(s, format string) error {
	// optional: you can implement date parsing using time.Parse(format, s)
	// for simplicity, here just check non-empty
	if s == "" {
		return errors.New("date cannot be empty")
	}
	return nil
}

package govalid

import (
	"fmt"
	errorpkg "govalid/errors"
	"govalid/rules"
	"reflect"
	"strconv"
	"strings"
)

// ValidateStruct validates any struct using `valid` tags
func ValidateStruct(v interface{}) error {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// If pointer, get the element
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	var errs []error

	// Iterate over struct fields
	for i := 0; i < rv.NumField(); i++ {
		fieldVal := rv.Field(i)
		fieldType := rt.Field(i)

		tag := fieldType.Tag.Get("valid")
		if tag == "" {
			continue
		}

		rulesList := strings.Split(tag, ",")

		for _, rawRule := range rulesList {
			rule, param := parseRule(rawRule)

			switch rule {

			// ----- REQUIRED -----
			case "required":
				if err := rules.Required(fieldVal); err != nil {
					errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
				}

			// ----- STRING RULES -----
			case "min":
				if fieldVal.Kind() == reflect.String {
					n := atoi(param)
					if err := rules.StringMin(fieldVal.String(), n); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			case "max":
				if fieldVal.Kind() == reflect.String {
					n := atoi(param)
					if err := rules.StringMax(fieldVal.String(), n); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			case "email":
				if fieldVal.Kind() == reflect.String {
					if err := rules.StringEmail(fieldVal.String()); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			case "uuid":
				if fieldVal.Kind() == reflect.String {
					if err := rules.StringUUID(fieldVal.String()); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			case "url":
				if fieldVal.Kind() == reflect.String {
					if err := rules.StringURL(fieldVal.String()); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			// ----- NUMERIC RULES -----
			case "min_num":
				if isNumber(fieldVal) {
					n := atoi(param)
					if err := rules.MinNumber(fieldVal.Float(), float64(n)); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			case "max_num":
				if isNumber(fieldVal) {
					n := atoi(param)
					if err := rules.MaxNumber(fieldVal.Float(), float64(n)); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			// ----- SLICE VALIDATION -----
			case "each_email":
				if fieldVal.Kind() == reflect.Slice {
					for j := 0; j < fieldVal.Len(); j++ {
						elem := fieldVal.Index(j)
						if err := rules.StringEmail(elem.String()); err != nil {
							errs = append(errs, fmt.Errorf("%s[%d]: %v", fieldType.Name, j, err))
						}
					}
				}

			// ----- NESTED STRUCT VALIDATION -----
			case "nested":
				if fieldVal.Kind() == reflect.Struct {
					if err := ValidateStruct(fieldVal.Interface()); err != nil {
						errs = append(errs, fmt.Errorf("%s: %v", fieldType.Name, err))
					}
				}

			// ----- UNKNOWN RULES -----
			default:
				errs = append(errs, fmt.Errorf("unknown rule '%s' for field '%s'", rule, fieldType.Name))
			}
		}
	}

	return errorpkg.CombineErrors(errs)
}

////////////////////////////////////////////
// HELPERS
////////////////////////////////////////////

func parseRule(rule string) (string, string) {
	if strings.Contains(rule, ":") {
		parts := strings.Split(rule, ":")
		return parts[0], parts[1]
	}
	return rule, ""
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func isNumber(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int64, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

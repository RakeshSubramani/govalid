package rules

import (
	"errors"
	"fmt"
	errorpkg "govalid/errors"
	"reflect"
)

// Required validates that a field is not zero-valued
func Required(v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		if v.Len() == 0 {
			return errors.New("value is required")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() == 0 {
			return errors.New("value is required")
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.Uint() == 0 {
			return errors.New("value is required")
		}
	case reflect.Float32, reflect.Float64:
		if v.Float() == 0 {
			return errors.New("value is required")
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if v.Len() == 0 {
			return errors.New("value is required")
		}
	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			return errors.New("value is required")
		}
	case reflect.Struct:
		// check zero value
		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
			return errors.New("value is required")
		}
	default:
		return errors.New("unsupported type for required validation")
	}
	return nil
}

// ValidateSlice validates each element of a slice with a provided function
func ValidateSlice(slice reflect.Value, validateFunc func(reflect.Value) error) error {
	if slice.Kind() != reflect.Slice && slice.Kind() != reflect.Array {
		return fmt.Errorf("ValidateSlice requires slice/array type, got %s", slice.Kind())
	}

	var errs []error
	for i := 0; i < slice.Len(); i++ {
		elem := slice.Index(i)
		if err := validateFunc(elem); err != nil {
			errs = append(errs, fmt.Errorf("[%d]: %v", i, err))
		}
	}

	if len(errs) > 0 {
		return errorpkg.CombineErrors(errs)
	}
	return nil
}

// NestedStruct validates a nested struct using a validator function
func NestedStruct(v reflect.Value, validateFunc func(interface{}) error) error {
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("NestedStruct requires a struct type, got %s", v.Kind())
	}
	return validateFunc(v.Interface())
}

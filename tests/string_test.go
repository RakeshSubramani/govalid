package tests

import (
	"govalid/rules"
	"reflect"
	"testing"

	validator "govalid"
)

func TestStringValidator(t *testing.T) {

	t.Run("Required - fail", func(t *testing.T) {
		err := validator.String("").Required().Error()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Required - pass", func(t *testing.T) {
		err := validator.String("hello").Required().Error()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("MinLength - fail", func(t *testing.T) {
		err := validator.String("go").MinLength(3).Error()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("MinLength - pass", func(t *testing.T) {
		err := validator.String("hello").MinLength(3).Error()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("MaxLength - fail", func(t *testing.T) {
		err := validator.String("helloworld!").MaxLength(5).Error()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("MaxLength - pass", func(t *testing.T) {
		err := validator.String("hello").MaxLength(10).Error()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}

func TestStringRequired(t *testing.T) {
	val := reflect.ValueOf("")
	err := rules.Required(val)
	if err == nil {
		t.Error("expected required error, got nil")
	}
}


package tests

import (
	"testing"

	validator "govalid"
)

func TestNumberValidator(t *testing.T) {

	t.Run("Min - fail", func(t *testing.T) {
		err := validator.Number(3).Min(5).Error()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Min - pass", func(t *testing.T) {
		err := validator.Number(10).Min(5).Error()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("Max - fail", func(t *testing.T) {
		err := validator.Number(10).Max(5).Error()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Max - pass", func(t *testing.T) {
		err := validator.Number(3).Max(5).Error()
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}

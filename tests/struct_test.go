package tests

import (
	gv "govalid"
	"testing"
)

type User struct {
	Name  string `valid:"required,min:3"`
	Age   int    `valid:"min:18"`
	Email string `valid:"email"`
}

func TestStructValidation(t *testing.T) {
	u := User{
		Name:  "Ra",
		Age:   12,
		Email: "invalid-email",
	}

	err := gv.ValidateStruct(u)
	if err == nil {
		t.Fatal("expected validation errors, got nil")
	}
}

func TestStructValidationSuccess(t *testing.T) {
	u := User{
		Name:  "sample",
		Age:   22,
		Email: "sample@example.com",
	}

	err := gv.ValidateStruct(u)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

type Nested struct {
	City string `valid:"required"`
}

type Person struct {
	Name   string `valid:"required"`
	Detail Nested `valid:"required"`
}

func TestNestedStruct(t *testing.T) {
	p := Person{
		Name:   "",
		Detail: Nested{City: ""},
	}

	err := gv.ValidateStruct(p)
	if err == nil {
		t.Error("expected nested struct errors, got nil")
	}
}

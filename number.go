package govalid

import (
	"govalid/errors"
	"govalid/rules"
)

type NumberValidator struct {
	value float64
	errs  []error
}

func Number(value float64) *NumberValidator {
	return &NumberValidator{value: value}
}

///////////////////////////////////////////
// BASIC RANGE
///////////////////////////////////////////

func (n *NumberValidator) Min(v float64) *NumberValidator {
	if err := rules.MinNumber(n.value, v); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

func (n *NumberValidator) Max(v float64) *NumberValidator {
	if err := rules.MaxNumber(n.value, v); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

func (n *NumberValidator) Between(min, max float64) *NumberValidator {
	if err := rules.NumberBetween(n.value, min, max); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

///////////////////////////////////////////
// GREATER / LESS THAN
///////////////////////////////////////////

func (n *NumberValidator) GreaterThan(v float64) *NumberValidator {
	if err := rules.NumberGreaterThan(n.value, v); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

func (n *NumberValidator) LessThan(v float64) *NumberValidator {
	if err := rules.NumberLessThan(n.value, v); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

///////////////////////////////////////////
// POSITIVE / NEGATIVE
///////////////////////////////////////////

func (n *NumberValidator) Positive() *NumberValidator {
	if err := rules.NumberPositive(n.value); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

func (n *NumberValidator) Negative() *NumberValidator {
	if err := rules.NumberNegative(n.value); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

///////////////////////////////////////////
// NON-ZERO
///////////////////////////////////////////

func (n *NumberValidator) NonZero() *NumberValidator {
	if err := rules.NumberNonZero(n.value); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

///////////////////////////////////////////
// INTEGER / FLOAT CHECKS
///////////////////////////////////////////

func (n *NumberValidator) Integer() *NumberValidator {
	if err := rules.NumberInteger(n.value); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

func (n *NumberValidator) Float() *NumberValidator {
	if err := rules.NumberFloat(n.value); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

///////////////////////////////////////////
// MULTIPLE OF
///////////////////////////////////////////

func (n *NumberValidator) MultipleOf(v float64) *NumberValidator {
	if err := rules.NumberMultipleOf(n.value, v); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

///////////////////////////////////////////
// EQUAL / NOT EQUAL
///////////////////////////////////////////

func (n *NumberValidator) Equals(v float64) *NumberValidator {
	if err := rules.NumberEquals(n.value, v); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

func (n *NumberValidator) NotEquals(v float64) *NumberValidator {
	if err := rules.NumberNotEquals(n.value, v); err != nil {
		n.errs = append(n.errs, err)
	}
	return n
}

///////////////////////////////////////////
// FINAL ERROR
///////////////////////////////////////////

func (n *NumberValidator) Error() error {
	if len(n.errs) == 0 {
		return nil
	}
	return errorpkg.CombineErrors(n.errs)
}

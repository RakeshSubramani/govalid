package rules

import (
	"errors"
	"fmt"
	"math"
)

// ----------------- MIN / MAX -----------------

func MinNumber(v, min float64) error {
	if v < min {
		return fmt.Errorf("must be >= %v", min)
	}
	return nil
}

func MaxNumber(v, max float64) error {
	if v > max {
		return fmt.Errorf("must be <= %v", max)
	}
	return nil
}

func NumberBetween(v, min, max float64) error {
	if v < min || v > max {
		return fmt.Errorf("must be between %v and %v", min, max)
	}
	return nil
}

// ----------------- GREATER / LESS THAN -----------------

func NumberGreaterThan(v, min float64) error {
	if v <= min {
		return fmt.Errorf("must be > %v", min)
	}
	return nil
}

func NumberLessThan(v, max float64) error {
	if v >= max {
		return fmt.Errorf("must be < %v", max)
	}
	return nil
}

// ----------------- POSITIVE / NEGATIVE -----------------

func NumberPositive(v float64) error {
	if v <= 0 {
		return errors.New("must be positive")
	}
	return nil
}

func NumberNegative(v float64) error {
	if v >= 0 {
		return errors.New("must be negative")
	}
	return nil
}

func NumberNonZero(v float64) error {
	if v == 0 {
		return errors.New("must be non-zero")
	}
	return nil
}

// ----------------- INTEGER / FLOAT CHECK -----------------

func NumberInteger(v float64) error {
	if v != math.Trunc(v) {
		return errors.New("must be an integer")
	}
	return nil
}

func NumberFloat(v float64) error {
	if v == math.Trunc(v) {
		return errors.New("must be a float")
	}
	return nil
}

// ----------------- MULTIPLE OF -----------------

func NumberMultipleOf(v, n float64) error {
	if n == 0 {
		return errors.New("divisor cannot be zero")
	}
	if math.Mod(v, n) != 0 {
		return fmt.Errorf("must be a multiple of %v", n)
	}
	return nil
}

// ----------------- EQUAL / NOT EQUAL -----------------

func NumberEquals(v, n float64) error {
	if v != n {
		return fmt.Errorf("must be equal to %v", n)
	}
	return nil
}

func NumberNotEquals(v, n float64) error {
	if v == n {
		return fmt.Errorf("must not be equal to %v", n)
	}
	return nil
}

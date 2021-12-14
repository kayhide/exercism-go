package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	count int
}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.count)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	x, e := weightFodder.FodderAmount()
	switch {
	case x < 0:
		return 0.0, errors.New("Negative fodder")
	case cows == 0:
		return 0.0, errors.New("Division by zero")
	case cows < 0:
		return 0.0, &SillyNephewError{cows}
	case e == ErrScaleMalfunction:
		return 2.0 * x / float64(cows), nil
	case e == nil:
		return x / float64(cows), nil
	}
	return 0.0, e
}

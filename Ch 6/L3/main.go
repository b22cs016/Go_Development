package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return "can not divide " + fmt.Sprintf("%v", de.dividend) + " by zero"
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

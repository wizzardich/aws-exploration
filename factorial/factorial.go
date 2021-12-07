package main

import (
	"fmt"
)

type NegativeFactorialError struct {
	value int
}

type IntegerOverflowError struct{}

func (e NegativeFactorialError) Error() string {
	return fmt.Sprintf("cannot accept negative factorial value: %d", e.value)
}

func (e IntegerOverflowError) Error() string {
	return "runtime integer overflow occured"
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, NegativeFactorialError{n}
	}

	result := 1

	for i := 2; i <= n; i++ {
		result *= i
		if result <= 0 {
			return 0, IntegerOverflowError{}
		}
	}

	return result, nil
}

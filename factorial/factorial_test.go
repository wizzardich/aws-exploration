package main

import "testing"

func TestFactorial(t *testing.T) {
	result, _ := factorial(0)

	if result != 1 {
		t.Fatalf("factorial(0) == %d != 1", result)
	}

	result, _ = factorial(1)

	if result != 1 {
		t.Fatalf("factorial(1) == %d != 1", result)
	}

	result, _ = factorial(2)

	if result != 2 {
		t.Fatalf("factorial(2) == %d != 2", result)
	}

	result, _ = factorial(5)

	if result != 120 {
		t.Fatalf("factorial(5) == %d != 120", result)
	}

	result, _ = factorial(20)

	if result != 2432902008176640000 {
		t.Fatalf("factorial(20) == %d != 2432902008176640000", result)
	}

	result, err := factorial(21)

	switch typerr := err.(type) {
	case IntegerOverflowError:
		t.Log("received expected error")
	default:
		t.Fatalf("unexpected error type; expected: IntegerOverflowError, got %T", typerr)
	}

	result, err = factorial(130)

	switch typerr := err.(type) {
	case IntegerOverflowError:
		t.Log("received expected error")
	default:
		t.Fatalf("unexpected error type; expected: IntegerOverflowError, got %T", typerr)
	}

	result, err = factorial(-1)

	if result != 0 {
		t.Fatal("got non-zero factorial response along with negative factorial error")
	}

	switch typerr := err.(type) {
	case NegativeFactorialError:
		t.Log("received expected error")
	default:
		t.Fatalf("unexpected error type; expected: NegativeFactorialError, got %T", typerr)
	}
}
